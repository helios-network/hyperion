package queries

import (
	"context"
	"fmt"
	"slices"
	"strings"
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
		MinBatchFeeHLS:       global.GetMinBatchFeeHLS(counterpartyChainParams.BridgeChainId),
		MinTxFeeHLS:          global.GetMinTxFeeHLS(counterpartyChainParams.BridgeChainId),
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
	runnerSet := make(chan struct{}, 1)
	startErr := make(chan error, 1)

	go func() {
		fmt.Println("run orchestrator")
		baseDelay, _ := time.ParseDuration("10s")
		maxDelay, _ := time.ParseDuration("5m")
		currentDelay := baseDelay
		consecutiveErrors := 0

		for {
			// Create new cancellable context for each iteration
			ctxCancellable, cancel := context.WithCancelCause(ctx)

			fmt.Println("Starting orchestrator")

			// Initialize new target network
			targetNetworks, err := global.InitTargetNetworks(counterpartyChainParams)
			if err != nil {
				fmt.Println("Error initializing target network:", err)
				cancel(errors.New("runner stopped", 1, "error initializing target network"))
				if !sleepCtx(ctx, currentDelay) {
					fmt.Println("Cancel during backoff")
					cancel(errors.New("runner stopped", 1, "context cancelled during backoff"))
					return
				}
				currentDelay = min(currentDelay*2, maxDelay)
				consecutiveErrors++
				if consecutiveErrors > 5 {
					fmt.Println("Too many consecutive errors, waiting for max delay")
					if !sleepCtx(ctx, maxDelay) {
						fmt.Println("Cancel during max delay")
						cancel(errors.New("runner stopped", 1, "context cancelled during max delay"))
						return
					}
				}
				if strings.Contains(err.Error(), "no rpcs found") {
					select {
					case startErr <- err:
					default:
					}
					return
				}
				continue
			}

			// Create new hyperion instance
			hyperion, err := orchestrator.NewOrchestrator(
				targetNetworks,
				pricefeed.NewCoingeckoPriceFeed(100, &pricefeed.Config{BaseURL: "https://api.coingecko.com/api/v3"}),
				orchestratorCfg,
				global,
			)
			if err != nil {
				fmt.Println("Error creating new orchestrator:", err)
				cancel(errors.New("runner stopped", 1, "error creating new orchestrator"))
				if !sleepCtx(ctx, currentDelay) {
					fmt.Println("Cancel during backoff")
					cancel(errors.New("runner stopped", 1, "context cancelled during backoff"))
					return
				}
				currentDelay = min(currentDelay*2, maxDelay)
				consecutiveErrors++
				if consecutiveErrors > 5 {
					fmt.Println("Too many consecutive errors, waiting for max delay")
					if !sleepCtx(ctx, maxDelay) {
						fmt.Println("Cancel during max delay")
						cancel(errors.New("runner stopped", 1, "context cancelled during max delay"))
						return
					}
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

				err := hyperion.Run(ctxCancellable)
				errChan <- err
			}()

			// Wait for either error or context cancellation
			select {
			case err := <-errChan:
				if err != nil {
					fmt.Printf("Orchestrator error: %v, restarting after delay...\n", err)
					cancel(errors.New("runner stopped", 1, "error running orchestrator"))
					if !sleepCtx(ctx, currentDelay) {
						fmt.Println("Cancel during backoff")
						cancel(errors.New("runner stopped", 1, "context cancelled during backoff"))
						return
					}
					currentDelay = min(currentDelay*2, maxDelay)
					consecutiveErrors++
					if consecutiveErrors > 5 {
						fmt.Println("Too many consecutive errors, waiting for max delay")
						if !sleepCtx(ctx, maxDelay) {
							fmt.Println("Cancel during max delay")
							cancel(errors.New("runner stopped", 1, "context cancelled during max delay"))
							return
						}
					}
				} else {
					fmt.Println("Orchestrator stopped normally, exiting...")
					cancel(errors.New("runner stopped", 1, "orchestrator stopped normally"))
					return
				}
			case <-ctx.Done():
				fmt.Println("Context cancelled, stopping orchestrator, cause:", context.Cause(ctx))
				cancel(errors.New("runner stopped", 1, "context cancelled"))
				return
			}
		}
	}()

	// Wait for the runner to be set or context cancellation
	select {
	case <-runnerSet:
		return nil
	case err := <-startErr:
		fmt.Println("Error starting orchestrator, exiting...", err)
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func sleepCtx(ctx context.Context, d time.Duration) bool {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return true // a dormi
	case <-ctx.Done():
		return false // annulé
	}
}
