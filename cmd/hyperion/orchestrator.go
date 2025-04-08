package main

import (
	"context"
	"os"
	"time"

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
				EthNodeRPCs:           *cfg.ethNodeRPCs,
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

		// delay, err := time.ParseDuration("1s")
		// orShutdown(errors.Wrap(err, "failed to parse duration"))

		// loops.RetryFunction(ctx, func() (string, error) {
		// 	heliosNetwork.HyperionParams(ctx)
		// 	return "hello", nil
		// }, delay)
		// hyperionParams, err := loops.RetryFunction(ctx, func() (*hyperiontypes.Params, error) {
		// 	log.Println(heliosNetwork)
		// 	return heliosNetwork.HyperionParams(ctx)
		// }, delay)
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
		)

		log.WithFields(log.Fields{"hyperion_contract": hyperionContractAddr.String()}).Debugln("loaded Hyperion module params")

		// 2. Connect to ethereum network

		ethNetwork, err := ethereum.NewNetwork(hyperionContractAddr, ethKeyFromAddress, signerFn, ethNetworkCfg)
		orShutdown(err)

		log.WithFields(log.Fields{
			"chain_id":             *cfg.ethChainID,
			"rpcs":                 *cfg.ethNodeRPCs,
			"max_gas_price":        *cfg.ethMaxGasPrice,
			"gas_price_adjustment": *cfg.ethGasPriceAdjustment,
		}).Infoln("connected to Ethereum network")

		addr, isValidator := helios.HasRegisteredOrchestrator(heliosNetwork, uint64(*cfg.hyperionID), ethKeyFromAddress)
		
		if *cfg.testnetAutoRegister {
			log.Printf("auto-registering validator %s with orchestrator %s\n", ethKeyFromAddress.String(), heliosKeyring.Addr.String())
			isValidator, err = helios.TestnetAutoRegisterValidator(ctx, *cfg.hyperionID, heliosNetwork, isValidator, addr, ethKeyFromAddress)
			orShutdown(err)
		}

		if *cfg.testnetForceValset {
			log.Printf("force-updating valset for validator %s with orchestrator %s\n", ethKeyFromAddress.String(), heliosKeyring.Addr.String())
			err = helios.TestnetForceUpdateValset(ctx, *cfg.hyperionID, heliosNetwork, ethNetwork)
			orShutdown(err)
		}

		if !isValidator {
			orShutdown(errors.Wrap(err, "Currently Hyperion is only worked on valiator mode"))
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
