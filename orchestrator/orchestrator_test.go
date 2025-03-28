package orchestrator

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	"cosmossdk.io/math"
	cometrpc "github.com/cometbft/cometbft/rpc/core/types"
	comettypes "github.com/cometbft/cometbft/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"

	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	hyperionsubgraphevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/HyperionSubgraph.sol"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

const maxLoopRetries = 1

func Test_BatchCreator(t *testing.T) {
	t.Parallel()

	heliosTokenAddress := gethcommon.HexToAddress("0xe28b3B32B6c345A34Ff64674606124Dd5Aceca30")

	testTable := []struct {
		name     string
		expected error
		orch     *Orchestrator
	}{
		{
			name:     "failed to get token fees",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					UnbatchedTokensWithFeesFn: func(_ context.Context) ([]*hyperiontypes.BatchFees, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:     "no unbatched token fees",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					UnbatchedTokensWithFeesFn: func(context.Context) ([]*hyperiontypes.BatchFees, error) {
						return nil, nil
					},
				},
			},
		},

		{
			name:     "token fee less than threshold",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				cfg: Config{
					MinBatchFeeUSD:       51.0,
					ERC20ContractMapping: map[gethcommon.Address]string{heliosTokenAddress: "helios"},
				},
				priceFeed: MockPriceFeed{QueryUSDPriceFn: func(_ gethcommon.Address) (float64, error) { return 1, nil }},
				helios: MockCosmosNetwork{
					SendRequestBatchFn: func(context.Context, uint64, string) error { return nil },
					UnbatchedTokensWithFeesFn: func(context.Context) ([]*hyperiontypes.BatchFees, error) {
						fees, _ := math.NewIntFromString("50000000000000000000")
						return []*hyperiontypes.BatchFees{
							{
								Token:     heliosTokenAddress.String(),
								TotalFees: fees,
							},
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					TokenDecimalsFn: func(_ context.Context, _ gethcommon.Address) (uint8, error) {
						return 18, nil
					},
				},
			},
		},

		{
			name:     "token fees exceed threshold",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				priceFeed:   MockPriceFeed{QueryUSDPriceFn: func(_ gethcommon.Address) (float64, error) { return 1, nil }},
				cfg: Config{
					MinBatchFeeUSD:       49.0,
					ERC20ContractMapping: map[gethcommon.Address]string{heliosTokenAddress: "helios"},
				},
				helios: MockCosmosNetwork{
					SendRequestBatchFn: func(context.Context, uint64, string) error { return nil },
					UnbatchedTokensWithFeesFn: func(_ context.Context) ([]*hyperiontypes.BatchFees, error) {
						fees, _ := math.NewIntFromString("50000000000000000000")
						return []*hyperiontypes.BatchFees{{
							Token:     heliosTokenAddress.String(),
							TotalFees: fees,
						}}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					TokenDecimalsFn: func(_ context.Context, _ gethcommon.Address) (uint8, error) {
						return 18, nil
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			bc := batchCreator{Orchestrator: tt.orch}
			assert.ErrorIs(t, bc.requestTokenBatches(context.Background()), tt.expected)
		})
	}
}

func Test_Oracle(t *testing.T) {
	t.Parallel()

	ethAddr1 := gethcommon.HexToAddress("0x76D2dDbb89C36FA39FAa5c5e7C61ee95AC4D76C4")
	ethAddr2 := gethcommon.HexToAddress("0x3959f5246c452463279F690301D923D5a75bbD88")

	testTable := []struct {
		name                  string
		expected              error
		orch                  *Orchestrator
		lastResyncWithHelios  time.Time
		lastObservedEthHeight uint64
	}{
		{
			name:     "failed to get current validator set",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:     "orchestrator not bonded",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr1},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},
				},
			},
			lastResyncWithHelios:  time.Time{},
			lastObservedEthHeight: 0,
		},

		{
			name:     "failed to get latest ethereum height",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return nil, errors.New("fail")
					},
				},
			},
		},

		{
			name:     "not enough block on ethereum",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil // minimum is 12
					},
				},
			},
		},

		{
			name:                  "failed to get ethereum events",
			expected:              errors.New("oops"),
			lastObservedEthHeight: 100,
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(2100)}, nil
					},
					GetSendToCosmosEventsFn: func(_, _ uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:                  "failed to get last claim event",
			expected:              errors.New("oops"),
			lastObservedEthHeight: 100,
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},

					LastClaimEventByAddrFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
						return nil, errors.New("oops")
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(2100)}, nil
					},
					GetSendToCosmosEventsFn: func(_, _ uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error) {
						return []*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent{
							{
								EventNonce: big.NewInt(100),
							},
						}, nil
					},

					GetValsetUpdatedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
						return nil, nil
					},
					GetSendToHeliosEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
						return nil, nil
					},
					GetTransactionBatchExecutedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error) {
						return nil, nil
					},
					GetHyperionERC20DeployedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error) {
						return nil, nil
					},
				},
			},
		},

		{
			name:                  "no new events",
			expected:              nil,
			lastObservedEthHeight: 100,
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},

					LastClaimEventByAddrFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
						return &hyperiontypes.LastClaimEvent{
							EthereumEventNonce: 101,
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(2100)}, nil
					},
					GetSendToCosmosEventsFn: func(_, _ uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error) {
						return []*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent{
							{
								EventNonce: big.NewInt(100),
							},
						}, nil
					},

					GetValsetUpdatedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
						return nil, nil
					},
					GetSendToHeliosEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
						return nil, nil
					},
					GetTransactionBatchExecutedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error) {
						return nil, nil
					},
					GetHyperionERC20DeployedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error) {
						return nil, nil
					},
				},
			},
		},

		{
			name:                  "missed events triggers resync",
			expected:              nil,
			lastObservedEthHeight: 100,
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},

					LastClaimEventByAddrFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
						return &hyperiontypes.LastClaimEvent{
							EthereumEventNonce:  102,
							EthereumEventHeight: 1000,
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(2100)}, nil
					},
					GetSendToCosmosEventsFn: func(_, _ uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error) {
						return []*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent{
							{
								EventNonce: big.NewInt(104),
							},
						}, nil
					},

					GetValsetUpdatedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
						return nil, nil
					},
					GetSendToHeliosEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
						return nil, nil
					},
					GetTransactionBatchExecutedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error) {
						return nil, nil
					},
					GetHyperionERC20DeployedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error) {
						return nil, nil
					},
				},
			},
		},

		{
			name:                  "sent new event claim",
			expected:              nil,
			lastObservedEthHeight: 100,
			lastResyncWithHelios:  time.Now(), // skip auto resync
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},

					LastClaimEventByAddrFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
						return &hyperiontypes.LastClaimEvent{
							EthereumEventNonce:  102,
							EthereumEventHeight: 1000,
						}, nil
					},

					SendOldDepositClaimFn: func(_ context.Context, _ uint64, _ *hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent) error {
						return nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(2100)}, nil
					},
					GetSendToCosmosEventsFn: func(_, _ uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error) {
						return []*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent{
							{
								EventNonce: big.NewInt(103),
							},
						}, nil
					},

					GetValsetUpdatedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
						return nil, nil
					},
					GetSendToHeliosEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
						return nil, nil
					},
					GetTransactionBatchExecutedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error) {
						return nil, nil
					},
					GetHyperionERC20DeployedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error) {
						return nil, nil
					},
				},
			},
		},

		{
			name:                  "auto resync",
			expected:              nil,
			lastObservedEthHeight: 100,
			orch: &Orchestrator{
				logger:      DummyLog,
				cfg:         Config{EthereumAddr: ethAddr2},
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					CurrentValsetFn: func(_ context.Context, _ uint64) (*hyperiontypes.Valset, error) {
						return &hyperiontypes.Valset{
							Members: []*hyperiontypes.BridgeValidator{
								{
									EthereumAddress: ethAddr2.String(),
								},
							},
						}, nil
					},

					LastClaimEventByAddrFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
						return &hyperiontypes.LastClaimEvent{
							EthereumEventNonce:  102,
							EthereumEventHeight: 1000,
						}, nil
					},

					SendOldDepositClaimFn: func(_ context.Context, _ uint64, _ *hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent) error {
						return nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(context.Context, *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(2100)}, nil
					},
					GetSendToCosmosEventsFn: func(_, _ uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error) {
						return []*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent{
							{
								EventNonce: big.NewInt(103),
							},
						}, nil
					},

					GetValsetUpdatedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
						return nil, nil
					},
					GetSendToHeliosEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
						return nil, nil
					},
					GetTransactionBatchExecutedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error) {
						return nil, nil
					},
					GetHyperionERC20DeployedEventsFn: func(_, _ uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error) {
						return nil, nil
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			o := oracle{
				Orchestrator:          tt.orch,
				lastResyncWithHelios:  tt.lastResyncWithHelios,
				lastObservedEthHeight: tt.lastObservedEthHeight,
			}

			err := o.observeEthEvents(context.Background())
			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_Relayer_Valsets(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		expected error
		orch     *Orchestrator
	}{
		{
			name:     "failed to get latest valset updates",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:     "failed to get valset confirmations",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil // non-empty will do
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:     "no new valset to relay",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return nil, nil
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return nil, nil
					},
				},
			},
		},

		{
			name:     "no new valset to relay",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				ethereum: MockEthereumNetwork{
					GetValsetNonceFn: func(_ context.Context) (*big.Int, error) {
						return nil, errors.New("oops")
					},
				},
				helios: MockCosmosNetwork{
					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil // non-empty will do
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return []*hyperiontypes.MsgValsetConfirm{{}}, nil // non-empty will do
					},
				},
			},
		},

		{
			name:     "valset already updated",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				ethereum: MockEthereumNetwork{
					GetValsetNonceFn: func(_ context.Context) (*big.Int, error) {
						return big.NewInt(101), nil
					},
				},
				helios: MockCosmosNetwork{
					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil // non-empty will do
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return []*hyperiontypes.MsgValsetConfirm{{}}, nil // non-empty will do
					},
				},
			},
		},

		{
			name:     "failed to get helios block",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				ethereum: MockEthereumNetwork{
					GetValsetNonceFn: func(_ context.Context) (*big.Int, error) {
						return big.NewInt(99), nil
					},
				},
				helios: MockCosmosNetwork{
					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return nil, errors.New("oops")
					},

					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil // non-empty will do
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return []*hyperiontypes.MsgValsetConfirm{{}}, nil // non-empty will do
					},
				},
			},
		},

		{
			name:     "relay valset offser duration not expired",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				cfg:         Config{RelayValsetOffsetDur: 10 * time.Second},
				ethereum: MockEthereumNetwork{
					GetValsetNonceFn: func(_ context.Context) (*big.Int, error) {
						return big.NewInt(99), nil
					},
				},
				helios: MockCosmosNetwork{
					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return &cometrpc.ResultBlock{
							Block: &comettypes.Block{
								Header: comettypes.Header{Time: time.Now()},
							},
						}, nil
					},

					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil // non-empty will do
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return []*hyperiontypes.MsgValsetConfirm{{}}, nil // non-empty will do
					},
				},
			},
		},

		{
			name:     "failed to send valset update",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				cfg:         Config{RelayValsetOffsetDur: 0},
				ethereum: MockEthereumNetwork{
					GetValsetNonceFn: func(_ context.Context) (*big.Int, error) {
						return big.NewInt(99), nil
					},

					SendEthValsetUpdateFn: func(_ context.Context, _ *hyperiontypes.Valset, _ *hyperiontypes.Valset, _ []*hyperiontypes.MsgValsetConfirm) (*gethcommon.Hash, error) {
						return nil, errors.New("oops")
					},
				},
				helios: MockCosmosNetwork{
					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return &cometrpc.ResultBlock{
							Block: &comettypes.Block{
								Header: comettypes.Header{Time: time.Now()},
							},
						}, nil
					},

					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil // non-empty will do
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return []*hyperiontypes.MsgValsetConfirm{{}}, nil // non-empty will do
					},
				},
			},
		},

		{
			name:     "sent valset update",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				svcTags:     metrics.Tags{"svc": "relayer"},
				cfg:         Config{RelayValsetOffsetDur: 0},
				ethereum: MockEthereumNetwork{
					GetValsetNonceFn: func(_ context.Context) (*big.Int, error) {
						return big.NewInt(99), nil
					},

					SendEthValsetUpdateFn: func(_ context.Context, _ *hyperiontypes.Valset, _ *hyperiontypes.Valset, _ []*hyperiontypes.MsgValsetConfirm) (*gethcommon.Hash, error) {
						return &gethcommon.Hash{}, nil
					},
				},
				helios: MockCosmosNetwork{
					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return &cometrpc.ResultBlock{
							Block: &comettypes.Block{
								Header: comettypes.Header{Time: time.Now()},
							},
						}, nil
					},

					LatestValsetsFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil // non-empty will do
					},

					AllValsetConfirmsFn: func(_ context.Context, _ uint64, _ uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
						return []*hyperiontypes.MsgValsetConfirm{{}}, nil // non-empty will do
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := relayer{tt.orch}

			err := r.relayValset(context.Background(), &hyperiontypes.Valset{Nonce: 101})
			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}

}

func Test_Relayer_Batches(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		expected error
		orch     *Orchestrator
	}{
		{
			name:     "failed to get token batches",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:     "failed to get token batch confirmations",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return []*hyperiontypes.OutgoingTxBatch{{
							BatchTimeout: 100,
						}}, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return nil, errors.New("oops")
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},
				},
			},
		},

		{
			name:     "no batch to relay",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return nil, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return []*hyperiontypes.MsgConfirmBatch{{}}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},
				},
			},
		},

		{
			name:     "failed to get latest batch nonce",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return []*hyperiontypes.OutgoingTxBatch{{}}, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return []*hyperiontypes.MsgConfirmBatch{{}}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetTxBatchNonceFn: func(_ context.Context, _ gethcommon.Address) (*big.Int, error) {
						return nil, errors.New("oops")
					},
					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},
				},
			},
		},

		{
			name:     "batch already updated",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return []*hyperiontypes.OutgoingTxBatch{{
							BatchNonce: 100,
						}}, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return []*hyperiontypes.MsgConfirmBatch{{}}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetTxBatchNonceFn: func(_ context.Context, _ gethcommon.Address) (*big.Int, error) {
						return big.NewInt(100), nil
					},
					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},
				},
			},
		},

		{
			name:     "failed to get helios block",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return []*hyperiontypes.OutgoingTxBatch{{
							BatchNonce: 101,
						}}, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return []*hyperiontypes.MsgConfirmBatch{{}}, nil
					},

					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return nil, errors.New("oops")
					},
				},
				ethereum: MockEthereumNetwork{
					GetTxBatchNonceFn: func(_ context.Context, _ gethcommon.Address) (*big.Int, error) {
						return big.NewInt(100), nil
					},
					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},
				},
			},
		},

		{
			name:     "batch relay offset not expired",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				cfg:         Config{RelayBatchOffsetDur: 10 * time.Second},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return []*hyperiontypes.OutgoingTxBatch{{
							BatchNonce: 101,
						}}, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return []*hyperiontypes.MsgConfirmBatch{{}}, nil
					},

					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return &cometrpc.ResultBlock{
							Block: &comettypes.Block{
								Header: comettypes.Header{Time: time.Now()},
							},
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetTxBatchNonceFn: func(_ context.Context, _ gethcommon.Address) (*big.Int, error) {
						return big.NewInt(100), nil
					},
					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},
				},
			},
		},

		{
			name:     "failed to send batch update",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				cfg:         Config{RelayBatchOffsetDur: 0},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return []*hyperiontypes.OutgoingTxBatch{{
							BatchNonce: 101,
						}}, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return []*hyperiontypes.MsgConfirmBatch{{}}, nil
					},

					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return &cometrpc.ResultBlock{
							Block: &comettypes.Block{
								Header: comettypes.Header{Time: time.Now()},
							},
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetTxBatchNonceFn: func(_ context.Context, _ gethcommon.Address) (*big.Int, error) {
						return big.NewInt(100), nil
					},

					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},

					SendTransactionBatchFn: func(_ context.Context, _ *hyperiontypes.Valset, _ *hyperiontypes.OutgoingTxBatch, _ []*hyperiontypes.MsgConfirmBatch) (*gethcommon.Hash, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:     "sent batch update",
			expected: nil,
			orch: &Orchestrator{
				maxAttempts: maxLoopRetries,
				logger:      DummyLog,
				svcTags:     metrics.Tags{"svc": "relayer"},
				cfg:         Config{RelayBatchOffsetDur: 0},
				helios: MockCosmosNetwork{
					LatestTransactionBatchesFn: func(_ context.Context, _ uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
						return []*hyperiontypes.OutgoingTxBatch{{
							BatchNonce: 101,
						}}, nil
					},

					TransactionBatchSignaturesFn: func(_ context.Context, _ uint64, _ uint64, _ gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
						return []*hyperiontypes.MsgConfirmBatch{{}}, nil
					},

					GetBlockFn: func(_ context.Context, _ int64) (*cometrpc.ResultBlock, error) {
						return &cometrpc.ResultBlock{
							Block: &comettypes.Block{
								Header: comettypes.Header{Time: time.Now()},
							},
						}, nil
					},
				},
				ethereum: MockEthereumNetwork{
					GetTxBatchNonceFn: func(_ context.Context, _ gethcommon.Address) (*big.Int, error) {
						return big.NewInt(100), nil
					},

					GetHeaderByNumberFn: func(_ context.Context, _ *big.Int) (*gethtypes.Header, error) {
						return &gethtypes.Header{Number: big.NewInt(10)}, nil
					},

					SendTransactionBatchFn: func(_ context.Context, _ *hyperiontypes.Valset, _ *hyperiontypes.OutgoingTxBatch, _ []*hyperiontypes.MsgConfirmBatch) (*gethcommon.Hash, error) {
						return &gethcommon.Hash{}, nil
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := relayer{Orchestrator: tt.orch}
			err := r.relayTokenBatch(context.Background(), &hyperiontypes.Valset{Nonce: 101})

			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_Signer_Valsets(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		expected error
		orch     *Orchestrator
	}{
		{
			name:     "failed to get unsigned valsets",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					OldestUnsignedValsetsFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error) {
						return nil, errors.New("oops")
					},
				},
			},
		},

		{
			name:     "no valset updates to sign",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					OldestUnsignedValsetsFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error) {
						return nil, nil
					},
				},
			},
		},

		{
			name:     "failed to send valset confirm",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					OldestUnsignedValsetsFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil
					},

					SendValsetConfirmFn: func(_ context.Context, _ uint64, _ gethcommon.Address, _ gethcommon.Hash, _ *hyperiontypes.Valset) error {
						return errors.New("oops")
					},
				},
			},
		},

		{
			name:     "sent valset confirm",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					OldestUnsignedValsetsFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error) {
						return []*hyperiontypes.Valset{{}}, nil
					},

					SendValsetConfirmFn: func(_ context.Context, _ uint64, _ gethcommon.Address, _ gethcommon.Hash, _ *hyperiontypes.Valset) error {
						return nil
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := signer{Orchestrator: tt.orch}
			err := s.signValidatorSets(context.Background())

			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func Test_Signer_Batches(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name     string
		expected error
		orch     *Orchestrator
	}{
		{
			name:     "failed to get unsigned batches/no batch to confirm",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					OldestUnsignedTransactionBatchFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error) {
						return nil, errors.New("ooops")
					},
				},
			},
		},

		{
			name:     "failed to send batch confirm",
			expected: errors.New("oops"),
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					OldestUnsignedTransactionBatchFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error) {
						return &hyperiontypes.OutgoingTxBatch{}, nil
					},

					SendBatchConfirmFn: func(_ context.Context, _ uint64, _ gethcommon.Address, _ gethcommon.Hash, _ *hyperiontypes.OutgoingTxBatch) error {
						return errors.New("oops")
					},
				},
			},
		},

		{
			name:     "sent batch confirm",
			expected: nil,
			orch: &Orchestrator{
				logger:      DummyLog,
				maxAttempts: maxLoopRetries,
				helios: MockCosmosNetwork{
					OldestUnsignedTransactionBatchFn: func(_ context.Context, _ uint64, _ cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error) {
						return &hyperiontypes.OutgoingTxBatch{}, nil
					},

					SendBatchConfirmFn: func(_ context.Context, _ uint64, _ gethcommon.Address, _ gethcommon.Hash, _ *hyperiontypes.OutgoingTxBatch) error {
						return nil
					},
				},
			},
		},
	}

	for _, tt := range testTable {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := signer{Orchestrator: tt.orch}
			err := s.signNewBatch(context.Background())

			if tt.expected == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
