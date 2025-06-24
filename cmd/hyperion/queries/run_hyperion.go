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
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}

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

	defaultBatchOffsetDur := "2m"
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
		ValidatorAddress:     cosmostypes.ValAddress(global.GetCosmosAddress().Bytes()),
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

	fmt.Println("run Hyperion FUNC")

	// Create a channel to signal when the runner is set
	runnerSet := make(chan struct{})

	go func() {
		fmt.Println("run orchestrator")
		baseDelay, _ := time.ParseDuration("10s")
		maxDelay, _ := time.ParseDuration("5m")
		currentDelay := baseDelay
		consecutiveErrors := 0

		for {
			// Create new cancellable context for each iteration
			ctxCancellable, cancel := context.WithCancel(ctx)

			fmt.Println("Starting orchestrator")

			heliosNetwork := global.GetHeliosNetwork()
			if heliosNetwork == nil { // should not happen
				fmt.Println("Error initializing helios network:", err)
				cancel()
				time.Sleep(currentDelay)
				currentDelay = min(currentDelay*2, maxDelay)
				consecutiveErrors++
				if consecutiveErrors > 5 {
					fmt.Println("Too many consecutive errors, waiting for max delay")
					time.Sleep(maxDelay)
				}
				continue
			}

			// Initialize new target network
			targetNetwork, err := global.InitTargetNetwork(counterpartyChainParams)
			if err != nil {
				fmt.Println("Error initializing target network:", err)
				cancel()
				time.Sleep(currentDelay)
				currentDelay = min(currentDelay*2, maxDelay)
				consecutiveErrors++
				if consecutiveErrors > 5 {
					fmt.Println("Too many consecutive errors, waiting for max delay")
					time.Sleep(maxDelay)
				}
				continue
			}

			// Create new hyperion instance
			hyperion, err := orchestrator.NewOrchestrator(
				*heliosNetwork,
				*targetNetwork,
				pricefeed.NewCoingeckoPriceFeed(100, &pricefeed.Config{BaseURL: "https://api.coingecko.com/api/v3"}),
				orchestratorCfg,
				global,
			)
			if err != nil {
				fmt.Println("Error creating new orchestrator:", err)
				cancel()
				time.Sleep(currentDelay)
				currentDelay = min(currentDelay*2, maxDelay)
				consecutiveErrors++
				if consecutiveErrors > 5 {
					fmt.Println("Too many consecutive errors, waiting for max delay")
					time.Sleep(maxDelay)
				}
				continue
			}

			// Reset delay and error count on successful initialization
			currentDelay = baseDelay
			consecutiveErrors = 0

			// Update runner in global state
			global.SetRunner(chainId, cancel, hyperion)

			// Signal that the runner is set
			select {
			case runnerSet <- struct{}{}:
			default:
			}

			// Run the orchestrator with panic recovery
			errChan := make(chan error, 1)
			go func() {
				defer func() {
					if r := recover(); r != nil {
						fmt.Printf("Recovered from panic in hyperion.Run: %v\n", r)
						errChan <- fmt.Errorf("panic: %v", r)
					}
				}()

				err := hyperion.Run(ctxCancellable, network, *targetNetwork)
				errChan <- err
			}()

			// Wait for either error or context cancellation
			select {
			case err := <-errChan:
				if err != nil {
					fmt.Printf("Orchestrator error: %v, restarting after delay...\n", err)
					cancel()
					time.Sleep(currentDelay)
					currentDelay = min(currentDelay*2, maxDelay)
					consecutiveErrors++
					if consecutiveErrors > 5 {
						fmt.Println("Too many consecutive errors, waiting for max delay")
						time.Sleep(maxDelay)
					}
				} else {
					fmt.Println("Orchestrator stopped normally, exiting...")
					cancel()
					return
				}
			case <-ctx.Done():
				fmt.Println("Context cancelled, stopping orchestrator")
				cancel()
				return
			}
		}
	}()

	// Wait for the runner to be set or context cancellation
	select {
	case <-runnerSet:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
