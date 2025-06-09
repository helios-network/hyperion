package queries

import (
	"context"
	"fmt"
	"slices"
	"time"

	"cosmossdk.io/errors"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/pricefeed"
)

func RunHyperion(ctx context.Context, global *global.Global, chainId uint64) error {
	registeredNetworks, _ := helios.GetListOfNetworksWhereRegistered(*global.GetHeliosNetwork(), global.GetAddress())

	if !slices.Contains(registeredNetworks, chainId) {
		return fmt.Errorf("chainId %d is not registered", chainId)
	}

	network := *global.GetHeliosNetwork()
	counterpartyChainParams, err := network.GetCounterpartyChainParamsByChainId(ctx, chainId)
	if err != nil {
		return err
	}

	// todo: get from config

	defaultValsetOffsetDur := "5m"
	valsetDur, err := time.ParseDuration(defaultValsetOffsetDur)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to parse relay valset offset duration for chain %d", counterpartyChainParams.BridgeChainId))
	}

	defaultBatchOffsetDur := "5m"
	batchDur, err := time.ParseDuration(defaultBatchOffsetDur)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to parse relay batch offset duration for chain %d", counterpartyChainParams.BridgeChainId))
	}

	orchestratorCfg := orchestrator.Config{
		EnabledLogs:          "signer,relayer,oracle,batch-creator",
		ChainId:              counterpartyChainParams.BridgeChainId,
		ChainName:            counterpartyChainParams.BridgeChainName,
		HyperionId:           uint64(counterpartyChainParams.HyperionId),
		CosmosAddr:           global.GetCosmosAddress(),
		EthereumAddr:         global.GetAddress(),
		MinBatchFeeUSD:       global.GetMinBatchFeeUSD(),
		RelayValsetOffsetDur: valsetDur,
		RelayBatchOffsetDur:  batchDur,
		RelayValsets:         true,
		RelayBatches:         true,
		RelayExternalDatas:   true,
		RelayerMode:          false,
		ChainParams:          counterpartyChainParams,
	}

	heliosNetwork, err := global.InitHeliosNetwork(counterpartyChainParams.BridgeChainId)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to initialize helios network for chain %d", counterpartyChainParams.BridgeChainId))
	}

	targetNetwork, err := global.InitTargetNetwork(counterpartyChainParams)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to initialize target network for chain %d", counterpartyChainParams.BridgeChainId))
	}

	// Create hyperion and run it
	hyperion, err := orchestrator.NewOrchestrator(
		*heliosNetwork,
		*targetNetwork,
		pricefeed.NewCoingeckoPriceFeed(100, &pricefeed.Config{BaseURL: "https://api.coingecko.com/api/v3"}),
		orchestratorCfg,
	)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to initialize orchestrator for chain %d", counterpartyChainParams.BridgeChainId))
	}

	ctxCancellable, cancel := context.WithCancel(ctx)
	// run orchestrator and retry if it fails
	go func() {
		fmt.Println("run orchestrator")
		err = hyperion.Run(ctxCancellable, network, *targetNetwork)
		if err != nil {
			fmt.Println("Error running orchestrator", err)
			time.Sleep(10 * time.Second)
			RunHyperion(ctx, global, counterpartyChainParams.BridgeChainId)
		}
	}()

	global.SetRunner(chainId, cancel)

	return nil
}
