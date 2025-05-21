package main

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	"github.com/xlab/closer"
	log "github.com/xlab/suplog"

	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/keystore"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/pricefeed"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcchainlist"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/version"
)

func runOrchestrator(
	ctx *context.Context,
	cfg *Config,
	heliosKeyring *helios.Keyring,
	heliosNetworkCfg *helios.NetworkConfig,
	counterpartyChainParams *hyperiontypes.CounterpartyChainParams,
	ethKeyFromAddress gethcommon.Address,
	signerFn bind.SignerFn,
	personalSignFn keystore.PersonalSignFn,
) error {

	log.WithFields(log.Fields{"addr": heliosKeyring.Addr.String(), "hex": heliosKeyring.HexAddr.String()}).Infoln("Initialized Helios keyring")

	heliosNetwork, err := helios.NewNetworkWithBroadcast(heliosKeyring, personalSignFn, *heliosNetworkCfg)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to initialize helios network for chain %d", counterpartyChainParams.BridgeChainId))
	}

	var (
		hyperionContractAddr = gethcommon.HexToAddress(counterpartyChainParams.BridgeCounterpartyAddress)
	)

	log.WithFields(log.Fields{"hyperion_contract": hyperionContractAddr.String()}).Infoln("loaded Hyperion module params")

	rpcs := counterpartyChainParams.Rpcs

	log.Println("cfg.evmRPCs", cfg.evmRPCs)

	formattedRPCs := formatRPCs(*cfg.evmRPCs)

	log.Println("formattedRPCs", formattedRPCs)

	if ok := formattedRPCs[fmt.Sprintf("%d", counterpartyChainParams.BridgeChainId)]; ok != nil {
		for _, rpc := range formattedRPCs[fmt.Sprintf("%d", counterpartyChainParams.BridgeChainId)] {
			log.Println("rpc", rpc)
			rpcs = append(rpcs, &hyperiontypes.Rpc{
				Url:            rpc,
				Reputation:     1,
				LastHeightUsed: 1,
			})
		}
	}

	rpcChainListFeed := rpcchainlist.NewRpcChainListFeed()

	rpcsFromChainList, err := rpcChainListFeed.QueryRpcs(counterpartyChainParams.BridgeChainId)
	if err == nil {
		for _, rpc := range rpcsFromChainList {
			log.Println("rpc", rpc)
			rpcs = append(rpcs, &hyperiontypes.Rpc{
				Url:            rpc,
				Reputation:     1,
				LastHeightUsed: 1,
			})
		}
	} else {
		log.Println("Error fetching rpcs from chain list", err)
	}

	ethNetwork, err := ethereum.NewNetwork(hyperionContractAddr, ethKeyFromAddress, signerFn, ethereum.NetworkConfig{
		EthNodeRPCs:           rpcs,
		GasPriceAdjustment:    *cfg.ethGasPriceAdjustment,
		MaxGasPrice:           *cfg.ethMaxGasPrice,
		PendingTxWaitDuration: *cfg.pendingTxWaitDuration,
	})
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to initialize ethereum network for chain %d", counterpartyChainParams.BridgeChainId))
	}

	if !ethNetwork.TestRpcs(*ctx) {
		return errors.New("failed to test rpc")
	}

	balance, err := ethNetwork.GetNativeBalance(*ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get native balance")
	}
	if balance.Cmp(big.NewInt(0)) == 0 {
		return errors.New("Please fund your account with native tokens")
	}

	log.WithFields(log.Fields{
		"chain_id":             counterpartyChainParams.BridgeChainId,
		"rpcs":                 counterpartyChainParams.Rpcs,
		"max_gas_price":        *cfg.ethMaxGasPrice,
		"gas_price_adjustment": *cfg.ethGasPriceAdjustment,
	}).Infoln("connected to Ethereum network")

	addr, isValidator := helios.HasRegisteredOrchestrator(heliosNetwork, uint64(counterpartyChainParams.HyperionId), ethKeyFromAddress)

	if !isValidator {
		return nil
	}

	// bech32Str, err := sdk.Bech32ifyAddressBytes("helios", ethKeyFromAddress.Bytes())
	// orShutdown(err)
	// log.Infoln("bech32Str", bech32Str)

	if *cfg.testnetAutoRegister {
		log.Printf("auto-registering validator %s with orchestrator %s\n", ethKeyFromAddress.String(), heliosKeyring.Addr.String())
		isValidator, err = helios.TestnetAutoRegisterValidator(*ctx, int(counterpartyChainParams.HyperionId), heliosNetwork, isValidator, addr, ethKeyFromAddress)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to auto-register validator %s with orchestrator %s", ethKeyFromAddress.String(), heliosKeyring.Addr.String()))
		}
	}

	if *cfg.testnetForceValset {
		log.Printf("force-updating valset for validator %s with orchestrator %s\n", ethKeyFromAddress.String(), heliosKeyring.Addr.String())
		err = helios.TestnetForceUpdateValset(*ctx, int(counterpartyChainParams.HyperionId), heliosNetwork, ethNetwork)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to force-update valset for validator %s with orchestrator %s", ethKeyFromAddress.String(), heliosKeyring.Addr.String()))
		}
	}

	if !isValidator {
		return errors.New(fmt.Sprintf("Currently Hyperion is only worked on valiator mode for chain %d", counterpartyChainParams.BridgeChainId))
	}

	var (
		valsetDur time.Duration
		batchDur  time.Duration
	)

	if *cfg.relayValsets {
		valsetDur, err = time.ParseDuration(*cfg.relayValsetOffsetDur)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to parse relay valset offset duration for chain %d", counterpartyChainParams.BridgeChainId))
		}
	}

	if *cfg.relayBatches {
		batchDur, err = time.ParseDuration(*cfg.relayBatchOffsetDur)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to parse relay batch offset duration for chain %d", counterpartyChainParams.BridgeChainId))
		}
	}

	orchestratorCfg := orchestrator.Config{
		EnabledLogs:          *cfg.enabledLogs,
		ChainId:              counterpartyChainParams.BridgeChainId,
		ChainName:            counterpartyChainParams.BridgeChainName,
		HyperionId:           uint64(counterpartyChainParams.HyperionId),
		CosmosAddr:           heliosKeyring.Addr,
		EthereumAddr:         ethKeyFromAddress,
		MinBatchFeeUSD:       *cfg.minBatchFeeUSD,
		RelayValsetOffsetDur: valsetDur,
		RelayBatchOffsetDur:  batchDur,
		RelayValsets:         *cfg.relayValsets,
		RelayBatches:         *cfg.relayBatches,
		RelayExternalDatas:   *cfg.relayExternalDatas,
		RelayerMode:          !isValidator,
		ChainParams:          counterpartyChainParams,
	}

	// Create hyperion and run it
	hyperion, err := orchestrator.NewOrchestrator(
		heliosNetwork,
		ethNetwork,
		pricefeed.NewCoingeckoPriceFeed(100, &pricefeed.Config{BaseURL: *cfg.coingeckoApi}),
		orchestratorCfg,
	)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to initialize orchestrator for chain %d", counterpartyChainParams.BridgeChainId))
	}

	// run orchestrator and retry if it fails
	go func() {
		delay, _ := time.ParseDuration("10s")
		loops.RetryFunction(*ctx, func() (error, error) {

			err = hyperion.Run(*ctx, heliosNetwork, ethNetwork)
			if err != nil {

				if strings.Contains(err.Error(), "connection refused") {
					return nil, err
				}
				log.Infoln("Error Removing RPC", ethNetwork.GetLastUsedRpc())
				ethNetwork.RemoveLastUsedRpc() // remove the last used rpc who is in cause of the error then retry
				return nil, err
			}
			return nil, nil
		}, delay)
	}()

	return nil
}

// startOrchestrator action runs an infinite loop,
// listening for events and performing hooks.
//
// $ hyperion orchestrator
func orchestratorCmd(cmd *cli.Cmd) {
	cmd.Before = func() {
		initMetrics(cmd)
	}

	cfg := initConfig(cmd)
	cmd.Action = func() {
		// ensure a clean exit
		defer closer.Close()

		var (
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
		)

		log.WithFields(log.Fields{
			"version":    version.AppVersion,
			"git":        version.GitCommit,
			"build_date": version.BuildDate,
			"go_version": version.GoVersion,
			"go_arch":    version.GoArch,
		}).Infoln("Hyperion - Hyperion module companion binary for bridging assets between Helios and Ethereum")

		ctx, cancelFn := context.WithCancel(context.Background())
		closer.Bind(cancelFn)

		heliosKeyring, err := helios.NewKeyring(heliosKeyringCfg)
		orShutdown(errors.Wrap(err, "failed to initialize Helios keyring"))

		heliosNetworkCfg.ValidatorAddress = heliosKeyring.Addr.String()

		heliosNetwork, err := helios.NewNetworkWithoutBroadcast(heliosKeyring, heliosNetworkCfg)
		orShutdown(err)
		log.WithFields(log.Fields{"chain_id": *cfg.heliosChainID, "gas_price": *cfg.heliosGasPrices}).Infoln("connected to Helios network")

		hyperionParams, err := heliosNetwork.HyperionParams(ctx)
		orShutdown(errors.Wrap(err, "failed to query hyperion params, is heliades running?"))

		networks, err := heliosNetwork.GetListOfNetworksWhereRegistered(ctx, gethcommon.HexToAddress(*cfg.ethKeyFrom))
		orShutdown(errors.Wrap(err, "failed to get list of networks where registered"))

		for _, counterpartyChainParams := range hyperionParams.CounterpartyChainParams {

			if !slices.Contains(networks, uint64(counterpartyChainParams.BridgeChainId)) {
				continue
			}

			ethKeyFromAddress, signerFn, personalSignFn, err := initEthereumAccountsManager(
				uint64(counterpartyChainParams.BridgeChainId),
				cfg.ethKeystoreDir,
				cfg.ethKeyFrom,
				cfg.ethPassphrase,
				cfg.ethPrivKey,
			)

			if err != nil {
				orShutdown(errors.Wrap(err, fmt.Sprintf("failed to initialize ethereum accounts manager for chain %d", counterpartyChainParams.BridgeChainId)))
			}

			go func() {
				delay, _ := time.ParseDuration("5s")
				loops.RetryFunction(ctx, func() (error, error) {
					return nil, runOrchestrator(&ctx, &cfg, &heliosKeyring, &heliosNetworkCfg, counterpartyChainParams, ethKeyFromAddress, signerFn, personalSignFn)
				}, delay)
			}()
		}

		closer.Hold()
	}
}
