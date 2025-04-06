package main

import (
	"context"
	"os"
	"time"

	sdkmath "cosmossdk.io/math"
	gethcommon "github.com/ethereum/go-ethereum/common"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/pricefeed"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/version"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	chaintypes "github.com/Helios-Chain-Labs/sdk-go/chain/types"
)

// startOrchestrator action runs an infinite loop,
// listening for events and performing hooks.
//
// $ hyperion orchestrator
func orchestratorCmd(cmd *cli.Cmd) {
	cmd.Before = func() {
		initMetrics(cmd)
	}

	// config := sdk.GetConfig()
	// config.SetBech32PrefixForAccount("helios", "helios")
	// config.Seal()

	cmd.Action = func() {
		// ensure a clean exit
		defer closer.Close()

		var (
			cfg              = initConfig(cmd)
			heliosKeyringCfg = helios.KeyringConfig{
				KeyringDir:     *cfg.heliosKeyringDir,
				KeyringAppName: *cfg.heliosKeyringAppName,
				KeyringBackend: *cfg.heliosKeyringBackend,
				KeyFrom:        *cfg.heliosKeyFrom,
				KeyPassphrase:  *cfg.heliosKeyPassphrase,
				PrivateKey:     *cfg.heliosPrivKey,
			}
			heliosNetworkCfg = helios.NetworkConfig{
				ChainID:       *cfg.heliosChainID,
				HeliosGRPC:    *cfg.heliosGRPC,
				TendermintRPC: *cfg.tendermintRPC,
				GasPrice:      *cfg.heliosGasPrices,
				Gas:           *cfg.heliosGas,
			}
			ethNetworkCfg = ethereum.NetworkConfig{
				EthNodeRPC:            *cfg.ethNodeRPC,
				GasPriceAdjustment:    *cfg.ethGasPriceAdjustment,
				MaxGasPrice:           *cfg.ethMaxGasPrice,
				PendingTxWaitDuration: *cfg.pendingTxWaitDuration,
				EthNodeAlchemyWS:      *cfg.ethNodeAlchemyWS,
			}
		)

		log.WithFields(log.Fields{
			"version":    version.AppVersion,
			"git":        version.GitCommit,
			"build_date": version.BuildDate,
			"go_version": version.GoVersion,
			"go_arch":    version.GoArch,
		}).Infoln("Hyperion - Hyperion module companion binary for bridging assets between Helios and Ethereum")

		// 1. Connect to Helios network

		heliosKeyring, err := helios.NewKeyring(heliosKeyringCfg)
		orShutdown(errors.Wrap(err, "failed to initialize Helios keyring"))

		log.WithFields(log.Fields{"addr": heliosKeyring.Addr.String(), "hex": heliosKeyring.HexAddr.String()}).Infoln("Initialized Helios keyring")
		ethKeyFromAddress, signerFn, personalSignFn, err := initEthereumAccountsManager(
			uint64(*cfg.ethChainID),
			cfg.ethKeystoreDir,
			cfg.ethKeyFrom,
			cfg.ethPassphrase,
			cfg.ethPrivKey,
		)
		orShutdown(errors.Wrap(err, "failed to initialize Ethereum keyring"))
		log.WithFields(log.Fields{"address": ethKeyFromAddress.String()}).Infoln("Initialized Ethereum keyring")

		heliosNetworkCfg.ValidatorAddress = heliosKeyring.Addr.String()
		heliosNetwork, err := helios.NewNetwork(heliosKeyring, personalSignFn, heliosNetworkCfg)
		orShutdown(err)
		log.WithFields(log.Fields{"chain_id": *cfg.heliosChainID, "gas_price": *cfg.heliosGasPrices}).Infoln("connected to Helios network")

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		hyperionParams, err := heliosNetwork.HyperionParams(ctx)
		orShutdown(errors.Wrap(err, "failed to query hyperion params, is heliades running?"))

		// 1.1 Search HyperionId into CounterpartyChainParams

		counterpartyChainParamsCfg := &hyperiontypes.CounterpartyChainParams{
			HyperionId:                0,
			BridgeCounterpartyAddress: "",
		}

		for _, counterpartyChainParams := range hyperionParams.CounterpartyChainParams {
			if counterpartyChainParams.HyperionId == uint64(*cfg.hyperionID) {
				counterpartyChainParamsCfg = counterpartyChainParams
			}
		}

		if counterpartyChainParamsCfg.BridgeCounterpartyAddress == "" {
			log.Fatalln("Counterparty Chain Params not found for hyperionId =", *cfg.hyperionID)
		}

		cfg.chainParams = counterpartyChainParamsCfg

		//------------

		var (
			hyperionContractAddr = gethcommon.HexToAddress(cfg.chainParams.BridgeCounterpartyAddress)
			heliosTokenAddr      = gethcommon.HexToAddress(cfg.chainParams.CosmosCoinErc20Contract)
			erc20ContractMapping = map[gethcommon.Address]string{heliosTokenAddr: chaintypes.HeliosCoin}
		)

		log.WithFields(log.Fields{"hyperion_contract": hyperionContractAddr.String(), "helios_token_contract": heliosTokenAddr.String()}).Debugln("loaded Hyperion module params")

		// 2. Connect to ethereum network

		log.Info("hyperionContractAddr", hyperionContractAddr)
		ethNetwork, err := ethereum.NewNetwork(hyperionContractAddr, ethKeyFromAddress, signerFn, ethNetworkCfg)
		orShutdown(err)

		log.WithFields(log.Fields{
			"chain_id":             *cfg.ethChainID,
			"rpc":                  *cfg.ethNodeRPC,
			"max_gas_price":        *cfg.ethMaxGasPrice,
			"gas_price_adjustment": *cfg.ethGasPriceAdjustment,
		}).Infoln("connected to Ethereum network")

		addr, isValidator := helios.HasRegisteredOrchestrator(heliosNetwork, uint64(*cfg.hyperionID), ethKeyFromAddress)
		if isValidator {
			log.Debugln("provided ETH address is registered with an orchestrator", addr.String())
		} else {
			err := heliosNetwork.SendSetOrchestratorAddresses(ctx, uint64(*cfg.hyperionID), ethKeyFromAddress.String())
			orShutdown(err)
			addr, isValidator = helios.HasRegisteredOrchestrator(heliosNetwork, uint64(*cfg.hyperionID), ethKeyFromAddress)
			if isValidator {
				log.Debugln("provided ETH address is registered with an orchestrator", addr.String())
			}
		}

		// check if the helios hyperion is not synchronized
		nonce, err := heliosNetwork.QueryGetLastObservedEventNonce(ctx, uint64(*cfg.hyperionID))
		orShutdown(err)
		lastEventNonce, err := ethNetwork.GetLastEventNonce(ctx)
		orShutdown(err)
		if nonce == 0 && lastEventNonce.Uint64() > 1 { // not firstime

			height, err := ethNetwork.GetLastValsetUpdatedEventHeight(ctx)
			orShutdown(err)

			lastEventBlockHeight, err := ethNetwork.GetLastEventHeight(ctx)
			orShutdown(err)

			events, err := ethNetwork.GetValsetUpdatedEventsAtSpecificBlock(height.Uint64())
			orShutdown(err)

			if len(events) == 0 {
				log.Fatalln("helios hyperion is not synchronized, please wait for it to be synchronized")
			}

			event := events[0]

			valset := &hyperiontypes.Valset{
				Nonce:        event.NewValsetNonce.Uint64(),
				Members:      make([]*hyperiontypes.BridgeValidator, 0, len(event.Powers)),
				RewardAmount: sdkmath.NewIntFromBigInt(event.RewardAmount),
				RewardToken:  event.RewardToken.Hex(),
			}

			for idx, p := range event.Powers {
				valset.Members = append(valset.Members, &hyperiontypes.BridgeValidator{
					Power:           p.Uint64(),
					EthereumAddress: event.Validators[idx].Hex(),
				})
			}

			err = heliosNetwork.SendForceSetValsetAndLastObservedEventNonce(ctx, uint64(*cfg.hyperionID), lastEventNonce.Uint64(), lastEventBlockHeight.Uint64(), valset)
			orShutdown(err)

			log.Infoln("helios hyperion is now forcefully synchronized with ethereum hyperion")
		}

		var (
			valsetDur time.Duration
			batchDur  time.Duration
		)

		if *cfg.relayValsets {
			valsetDur, err = time.ParseDuration(*cfg.relayValsetOffsetDur)
			orShutdown(err)
		}

		if *cfg.relayBatches {
			batchDur, err = time.ParseDuration(*cfg.relayBatchOffsetDur)
			orShutdown(err)
		}

		orchestratorCfg := orchestrator.Config{
			HyperionId:           uint64(*cfg.hyperionID),
			CosmosAddr:           heliosKeyring.Addr,
			EthereumAddr:         ethKeyFromAddress,
			MinBatchFeeUSD:       *cfg.minBatchFeeUSD,
			ERC20ContractMapping: erc20ContractMapping,
			RelayValsetOffsetDur: valsetDur,
			RelayBatchOffsetDur:  batchDur,
			RelayValsets:         *cfg.relayValsets,
			RelayBatches:         *cfg.relayBatches,
			RelayExternalDatas:   *cfg.relayExternalDatas,
			RelayerMode:          !isValidator,
			ChainParams:          cfg.chainParams,
		}

		// Create hyperion and run it
		hyperion, err := orchestrator.NewOrchestrator(
			heliosNetwork,
			ethNetwork,
			pricefeed.NewCoingeckoPriceFeed(100, &pricefeed.Config{BaseURL: *cfg.coingeckoApi}),
			orchestratorCfg,
		)
		orShutdown(err)

		go func() {
			if err := hyperion.Run(ctx, heliosNetwork, ethNetwork); err != nil {
				log.Errorln(err)
				os.Exit(1)
			}
		}()

		closer.Hold()
	}
}
