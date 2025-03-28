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
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/cosmos"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
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
			cosmosKeyringCfg = cosmos.KeyringConfig{
				KeyringDir:     *cfg.cosmosKeyringDir,
				KeyringAppName: *cfg.cosmosKeyringAppName,
				KeyringBackend: *cfg.cosmosKeyringBackend,
				KeyFrom:        *cfg.cosmosKeyFrom,
				KeyPassphrase:  *cfg.cosmosKeyPassphrase,
				PrivateKey:     *cfg.cosmosPrivKey,
				UseLedger:      *cfg.cosmosUseLedger,
			}
			cosmosNetworkCfg = cosmos.NetworkConfig{
				ChainID:       *cfg.cosmosChainID,
				CosmosGRPC:    *cfg.cosmosGRPC,
				TendermintRPC: *cfg.tendermintRPC,
				GasPrice:      *cfg.cosmosGasPrices,
			}
			ethNetworkCfg = ethereum.NetworkConfig{
				EthNodeRPC:            *cfg.ethNodeRPC,
				GasPriceAdjustment:    *cfg.ethGasPriceAdjustment,
				MaxGasPrice:           *cfg.ethMaxGasPrice,
				PendingTxWaitDuration: *cfg.pendingTxWaitDuration,
				EthNodeAlchemyWS:      *cfg.ethNodeAlchemyWS,
			}
		)
		log.Println("cosmosKeyringCfg", cosmosKeyringCfg)
		if *cfg.cosmosUseLedger || *cfg.ethUseLedger {
			log.Fatalln("cannot use Ledger for orchestrator, since signatures must be realtime")
		}

		log.WithFields(log.Fields{
			"version":    version.AppVersion,
			"git":        version.GitCommit,
			"build_date": version.BuildDate,
			"go_version": version.GoVersion,
			"go_arch":    version.GoArch,
		}).Infoln("Hyperion - Hyperion module companion binary for bridging assets between Helios and Ethereum")

		// 1. Connect to Helios network

		cosmosKeyring, err := cosmos.NewKeyring(cosmosKeyringCfg)
		orShutdown(errors.Wrap(err, "failed to initialize Helios keyring"))
		log.Infoln("initialized Helios keyring", cosmosKeyring.Addr.String())

		log.WithFields(log.Fields{"ethChainID": *cfg.ethChainID, "ethKeystoreDir": *cfg.ethKeystoreDir, "ethKeyFrom": *cfg.ethKeyFrom, "ethPassphrase": *cfg.ethPassphrase, "ethPrivKey": *cfg.ethPrivKey, "ethUseLedger": *cfg.ethUseLedger}).Infoln("initializing Ethereum keyring")
		ethKeyFromAddress, signerFn, personalSignFn, err := initEthereumAccountsManager(
			uint64(*cfg.ethChainID),
			cfg.ethKeystoreDir,
			cfg.ethKeyFrom,
			cfg.ethPassphrase,
			cfg.ethPrivKey,
			cfg.ethUseLedger,
		)
		orShutdown(errors.Wrap(err, "failed to initialize Ethereum keyring"))
		log.Infoln("initialized Ethereum keyring", ethKeyFromAddress.String())

		cosmosNetworkCfg.ValidatorAddress = cosmosKeyring.Addr.String()
		cosmosNetwork, err := cosmos.NewNetwork(cosmosKeyring, personalSignFn, cosmosNetworkCfg)
		orShutdown(err)
		log.WithFields(log.Fields{"chain_id": *cfg.cosmosChainID, "gas_price": *cfg.cosmosGasPrices}).Infoln("connected to Helios network")

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		hyperionParams, err := cosmosNetwork.HyperionParams(ctx)
		orShutdown(errors.Wrap(err, "failed to query hyperion params, is heliades running?"))

		// for _, hyperionCounterpartyChainParams := range hyperionParams.CounterpartyChainParams {
		// 	var (
		// 		hyperionContractAddr = gethcommon.HexToAddress(hyperionCounterpartyChainParams.BridgeCounterpartyAddress)
		// 		heliosTokenAddr      = gethcommon.HexToAddress(hyperionCounterpartyChainParams.CosmosCoinErc20Contract)
		// 		// erc20ContractMapping = map[gethcommon.Address]string{heliosTokenAddr: chaintypes.HeliosCoin}
		// 	)

		// 	log.WithFields(log.Fields{"hyperion_id": hyperionCounterpartyChainParams.HyperionId, "hyperion_contract": hyperionContractAddr.String(), "helios_token_contract": heliosTokenAddr.String()}).Debugln("loaded Hyperion module params")

		// 	_, err := ethereum.NewNetwork(hyperionContractAddr, ethKeyFromAddress, signerFn, ethNetworkCfg)
		// 	orShutdown(err)
		// }
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

		var (
			hyperionContractAddr = gethcommon.HexToAddress(counterpartyChainParamsCfg.BridgeCounterpartyAddress)
			heliosTokenAddr      = gethcommon.HexToAddress(counterpartyChainParamsCfg.CosmosCoinErc20Contract)
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

		addr, isValidator := cosmos.HasRegisteredOrchestrator(cosmosNetwork, uint64(*cfg.hyperionID), ethKeyFromAddress)
		if isValidator {
			log.Debugln("provided ETH address is registered with an orchestrator", addr.String())
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
			CosmosAddr:           cosmosKeyring.Addr,
			EthereumAddr:         ethKeyFromAddress,
			MinBatchFeeUSD:       *cfg.minBatchFeeUSD,
			ERC20ContractMapping: erc20ContractMapping,
			RelayValsetOffsetDur: valsetDur,
			RelayBatchOffsetDur:  batchDur,
			RelayValsets:         *cfg.relayValsets,
			RelayBatches:         *cfg.relayBatches,
			RelayerMode:          !isValidator,
		}

		// Create hyperion and run it
		hyperion, err := orchestrator.NewOrchestrator(
			cosmosNetwork,
			ethNetwork,
			pricefeed.NewCoingeckoPriceFeed(100, &pricefeed.Config{BaseURL: *cfg.coingeckoApi}),
			orchestratorCfg,
		)
		orShutdown(err)

		go func() {
			if err := hyperion.Run(ctx, cosmosNetwork, ethNetwork); err != nil {
				log.Errorln(err)
				os.Exit(1)
			}
		}()

		closer.Hold()
	}
}
