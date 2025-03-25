package main

import (
	"context"
	// "math/big"
	"time"

	// sdkmath "cosmossdk.io/math"
	gethcommon "github.com/ethereum/go-ethereum/common"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/peggo/orchestrator"
	"github.com/Helios-Chain-Labs/peggo/orchestrator/cosmos"
	"github.com/Helios-Chain-Labs/peggo/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/peggo/orchestrator/pricefeed"
	"github.com/Helios-Chain-Labs/peggo/orchestrator/version"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	chaintypes "github.com/Helios-Chain-Labs/sdk-go/chain/types"
)

func testCmd(cmd *cli.Cmd) {
	cmd.Before = func() {
		initMetrics(cmd)
	}

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
		cosmosNetwork, hyperionQueryClient, chainClient, err := cosmos.NewNetworkV2(cosmosKeyring, personalSignFn, cosmosNetworkCfg)
		log.Debugln("chainClient", chainClient)
		log.Debugln("hyperionQueryClient", hyperionQueryClient)
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

		var (
			hyperionContractAddr = gethcommon.HexToAddress(hyperionParams.CounterpartyChainParams[*cfg.hyperionID].BridgeCounterpartyAddress)
			heliosTokenAddr      = gethcommon.HexToAddress(hyperionParams.CounterpartyChainParams[*cfg.hyperionID].CosmosCoinErc20Contract)
			erc20ContractMapping = map[gethcommon.Address]string{heliosTokenAddr: chaintypes.HeliosCoin}
		)
		log.Infoln("erc20ContractMapping", erc20ContractMapping)

		log.WithFields(log.Fields{"hyperion_contract": hyperionContractAddr.String(), "helios_token_contract": heliosTokenAddr.String()}).Debugln("loaded Hyperion module params")

		// 2. Connect to ethereum network

		log.Info("hyperionContractAddr", hyperionContractAddr)
		ethNetwork, err := ethereum.NewNetwork(hyperionContractAddr, ethKeyFromAddress, signerFn, ethNetworkCfg)
		log.Infoln("ethNetwork", ethNetwork)
		orShutdown(err)

		log.WithFields(log.Fields{
			"chain_id":             *cfg.ethChainID,
			"rpc":                  *cfg.ethNodeRPC,
			"max_gas_price":        *cfg.ethMaxGasPrice,
			"gas_price_adjustment": *cfg.ethGasPriceAdjustment,
		}).Infoln("connected to Ethereum network")

		addr, isValidator := cosmos.HasRegisteredOrchestrator(cosmosNetwork, ethKeyFromAddress, uint64(*cfg.hyperionID))
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

		log.Infoln("hyperion", hyperion)
		res, err := hyperionQueryClient.BatchFees(ctx, &hyperiontypes.QueryBatchFeeRequest{
			HyperionId: uint64(*cfg.hyperionID),
		})
		orShutdown(err)
		log.Infoln("res", res)
		// err = cosmosNetwork.SendRequestBatch(ctx, 2, "hyperion0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8")
		// orShutdown(err)

		oldestTxBatch, err := cosmosNetwork.OldestUnsignedTransactionBatch(ctx, cosmosKeyring.Addr, uint64(*cfg.hyperionID))
		orShutdown(err)
		log.Infoln("oldestTxBatch", oldestTxBatch)
		// err = cosmosNetwork.SendBatchConfirm(ctx, uint64(*cfg.hyperionID), ethKeyFromAddress, gethcommon.HexToHash("0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8"), oldestTxBatch)
		// orShutdown(err)

		// query batch confirm request
		// req := hyperiontypes.QueryBatchConfirmsRequest{
		// 	HyperionId: uint64(*cfg.hyperionID),
		// 	Nonce:      oldestTxBatch.BatchNonce,
		// 	ContractAddress: oldestTxBatch.TokenContract,
		// }
		// log.Infoln("req", req)
		// batchConfirmReq, err := hyperionQueryClient.BatchConfirms(ctx, &req)
		// orShutdown(err)
		// log.Infoln("batchConfirmReq", batchConfirmReq)

		latestTxBatches, err := cosmosNetwork.LatestTransactionBatches(ctx, uint64(*cfg.hyperionID))
		orShutdown(err)
		log.Infoln("latestTxBatches", latestTxBatches)
		// latestEthValset, err := hyperion.GetLatestEthValset(ctx)
		// orShutdown(err)
		// log.Infoln("latestEthValset", latestEthValset)

		for _, batch := range latestTxBatches {
			log.Infoln("batch", batch)
			sigs, err := cosmosNetwork.TransactionBatchSignatures(ctx, batch.BatchNonce, gethcommon.HexToAddress(batch.TokenContract), uint64(*cfg.hyperionID))
			orShutdown(err)
			log.Infoln("sigs", sigs)
			// hash, err := ethNetwork.SendTransactionBatch(ctx, latestEthValset, batch, sigs)
			// orShutdown(err)
			// log.Infoln("hash", hash)
			// chainClient.SyncBroadcastMsg(&hyperiontypes.MsgWithdrawClaim{
			// 	HyperionId:    uint64(*cfg.hyperionID),
			// 	EventNonce:    batch.EventNonce.Uint64(),
			// 	BatchNonce:    batch.BatchNonce.Uint64(),
			// 	BlockHeight:   batch.Raw.BlockNumber,
			// 	TokenContract: batch.Token.Hex(),
			// 	Orchestrator:  cosmosKeyring.Addr.String(),
			// })
		}

		chainClient.SyncBroadcastMsg(&hyperiontypes.MsgWithdrawClaim{
			HyperionId:    uint64(*cfg.hyperionID),
			EventNonce:    3,
			BatchNonce:    1,
			BlockHeight:   19588125,
			TokenContract: "0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8",
			Orchestrator:  cosmosKeyring.Addr.String(),
		})

		// log.Infoln("ethKeyFromAddress", ethKeyFromAddress.String())
		// log.Infoln("cosmosKeyring.Addr", cosmosKeyring.Addr.String())
		// chainClient.SyncBroadcastMsg(&hyperiontypes.MsgDepositClaim{
		// 	HyperionId:    uint64(*cfg.hyperionID),
		// 	EventNonce:    9,
		// 	BlockHeight:   19384013,
		// 	TokenContract: "0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8",
		// 	Orchestrator:  cosmosKeyring.Addr.String(),
		// 	Amount:         sdkmath.NewIntFromBigInt(big.NewInt(100000)),
		// 	EthereumSender: ethKeyFromAddress.String(),
		// 	CosmosReceiver: cosmosKeyring.Addr.String(),
		// 	Data:           "",
		// })

		// hash, err := ethNetwork.SendTransactionBatch(ctx, nil, latestTxBatches[0], batchConfirmReq.Confirms)
		// orShutdown(err)
		// log.Infoln("hash", hash)

		closer.Hold()
	}
}
