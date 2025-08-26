package orchestrator

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cosmossdk.io/errors"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/avast/retry-go"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/metrics"
)

const (
	defaultLoopDur = 30 * time.Second
)

// PriceFeed provides token price for a given contract address
type PriceFeed interface {
	QueryUSDPrice(address gethcommon.Address) (float64, error)
}

type Global interface {
	GetRpcs(chainId uint64) ([]*hyperiontypes.Rpc, error)
	InitTargetNetwork(counterpartyChainParams *hyperiontypes.CounterpartyChainParams) (*ethereum.Network, error)
	GetMinBatchFeeHLS(chainId uint64) float64
	GetMinTxFeeHLS(chainId uint64) float64
}

type Config struct {
	EnabledLogs          string
	CosmosAddr           cosmostypes.AccAddress
	ValidatorAddress     cosmostypes.ValAddress
	HyperionId           uint64
	ChainName            string
	ChainId              uint64
	EthereumAddr         gethcommon.Address
	MinBatchFeeHLS       float64
	MinTxFeeHLS          float64
	RelayValsetOffsetDur time.Duration
	RelayBatchOffsetDur  time.Duration
	RelayValsets         bool
	RelayBatches         bool
	RelayExternalDatas   bool
	RelayerMode          bool
	ChainParams          *hyperiontypes.CounterpartyChainParams
}

type HyperionState struct {
	HyperionID             uint64
	Height                 uint64
	TargetHeight           uint64
	LastObservedHeight     uint64
	LastObservedEventNonce uint64
	LastEventNonce         uint64
	LastClaimBlockHeight   uint64
	LastClaimEventNonce    uint64

	BatchCount        int
	TxCount           int
	OutBridgedTxCount int
	InBridgedTxCount  int

	BatchCreatorStatus  string
	ExternalDataStatus  string
	OracleStatus        string
	RelayerStatus       string
	SignerStatus        string
	UpdaterStatus       string
	ValsetManagerStatus string

	BatchCreatorNextExecutionTimestamp  uint64
	ExternalDataNextExecutionTimestamp  uint64
	OracleNextExecutionTimestamp        uint64
	RelayerNextExecutionTimestamp       uint64
	SignerNextExecutionTimestamp        uint64
	UpdaterNextExecutionTimestamp       uint64
	ValsetManagerNextExecutionTimestamp uint64

	BatchCreatorLastExecutionFinishedTimestamp  uint64
	ExternalDataLastExecutionFinishedTimestamp  uint64
	OracleLastExecutionFinishedTimestamp        uint64
	RelayerLastExecutionFinishedTimestamp       uint64
	SignerLastExecutionFinishedTimestamp        uint64
	UpdaterLastExecutionFinishedTimestamp       uint64
	ValsetManagerLastExecutionFinishedTimestamp uint64
}

type Orchestrator struct {
	logger      log.Logger
	svcTags     metrics.Tags
	cfg         Config
	maxAttempts uint

	helios        helios.Network
	ethereum      ethereum.Network
	priceFeed     PriceFeed
	firstTimeSync bool
	global        Global

	height        uint64
	targetHeight  uint64
	valsetManager valsetManager

	HyperionState HyperionState

	CacheSymbol map[gethcommon.Address]string
}

func NewOrchestrator(
	helios helios.Network,
	eth ethereum.Network,
	priceFeed PriceFeed,
	cfg Config,
	global Global,
) (*Orchestrator, error) {
	o := &Orchestrator{
		logger: log.WithFields(log.Fields{
			"chain": cfg.ChainName,
		}),
		svcTags:       metrics.Tags{"svc": "hyperion_orchestrator"},
		helios:        helios,
		ethereum:      eth,
		priceFeed:     priceFeed,
		cfg:           cfg,
		maxAttempts:   100,
		firstTimeSync: false,
		global:        global,

		height:       0,
		targetHeight: 0,

		HyperionState: HyperionState{
			HyperionID:             cfg.HyperionId,
			Height:                 0,
			TargetHeight:           0,
			LastObservedHeight:     0,
			LastObservedEventNonce: 0,
			LastEventNonce:         0,
			LastClaimBlockHeight:   0,
			LastClaimEventNonce:    0,

			BatchCount:        0,
			OutBridgedTxCount: 0,
			InBridgedTxCount:  0,

			BatchCreatorStatus:  "idle",
			ExternalDataStatus:  "idle",
			OracleStatus:        "idle",
			RelayerStatus:       "idle",
			SignerStatus:        "idle",
			UpdaterStatus:       "idle",
			ValsetManagerStatus: "idle",

			BatchCreatorNextExecutionTimestamp:  0,
			ExternalDataNextExecutionTimestamp:  0,
			OracleNextExecutionTimestamp:        0,
			RelayerNextExecutionTimestamp:       0,
			SignerNextExecutionTimestamp:        0,
			UpdaterNextExecutionTimestamp:       0,
			ValsetManagerNextExecutionTimestamp: 0,

			BatchCreatorLastExecutionFinishedTimestamp:  0,
			ExternalDataLastExecutionFinishedTimestamp:  0,
			OracleLastExecutionFinishedTimestamp:        0,
			RelayerLastExecutionFinishedTimestamp:       0,
			SignerLastExecutionFinishedTimestamp:        0,
			UpdaterLastExecutionFinishedTimestamp:       0,
			ValsetManagerLastExecutionFinishedTimestamp: 0,
		},

		CacheSymbol: make(map[gethcommon.Address]string),
	}

	return o, nil
}

// Run starts all major loops required to make
// up the Orchestrator, all of these are async loops.
func (s *Orchestrator) Run(ctx context.Context, helios helios.Network, eth ethereum.Network) error {
	return s.startValidatorMode(ctx, eth)
}

func (s *Orchestrator) SetEthereum(eth ethereum.Network) {
	s.ethereum = eth
}

func (s *Orchestrator) SetValsetManager(valsetManager *valsetManager) {
	s.valsetManager = *valsetManager
}

func (s *Orchestrator) GetValsetManager() *valsetManager {
	return &s.valsetManager
}

func (s *Orchestrator) GetEthereum() ethereum.Network {
	return s.ethereum
}

func (s *Orchestrator) GetHelios() helios.Network {
	return s.helios
}

func (s *Orchestrator) GetConfig() Config {
	return s.cfg
}

func (s *Orchestrator) SetHeight(height uint64) {
	s.height = height
}

func (s *Orchestrator) SetTargetHeight(height uint64) {
	s.targetHeight = height
}

func (s *Orchestrator) GetHeight() uint64 {
	return s.height
}

func (s *Orchestrator) GetTargetHeight() uint64 {
	return s.targetHeight
}

func (s *Orchestrator) GetLogger() log.Logger {
	return s.logger
}

func (s *Orchestrator) GetStats() map[string]interface{} {
	return map[string]interface{}{
		"height":       s.height,
		"targetHeight": s.targetHeight,
	}
}

// startValidatorMode runs all orchestrator processes. This is called
// when hyperion is run alongside a validator helios node.
func (s *Orchestrator) startValidatorMode(ctx context.Context, eth ethereum.Network) error {
	s.logger.Infoln("running orchestrator in validator mode")

	fmt.Println("startValidatorMode")

	// get hyperion ID from contract
	hyperionIDHash, err := eth.GetHyperionID(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to query hyperion ID from contract")
	}
	hyperionID := hyperionIDHash.Big().Uint64()

	s.logger.Infoln("Our HyperionID", "is", hyperionID, "hash", hyperionIDHash.Hex())

	latestObservedHeight, err := s.helios.QueryGetLastObservedEthereumBlockHeight(ctx, s.cfg.HyperionId)
	if err != nil {
		return errors.Wrap(err, "unable to query hyperion module params, is heliades running?")
	}

	h, err := s.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "failed to get latest ethereum header")
	}
	latestHeight := h.Number.Uint64()

	// start from top of the blockchain
	ethereumBlockHeightWhereStart := latestHeight

	lastObservedEventNonce, err := s.helios.QueryGetLastObservedEventNonce(ctx, s.cfg.HyperionId)
	if err != nil {
		return errors.Wrap(err, "unable to query hyperion module params, is heliades running?")
	}

	lastEventNonce, err := s.ethereum.GetLastEventNonce(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to query blockchain state_lastEventNonce, is rpc running?")
	}

	if lastEventNonce.Uint64() > lastObservedEventNonce {
		s.logger.Infoln("lastEventNonce is greater than lastObservedEventNonce, starting from lastObservedEventNonce")
		// start from the next block after the last observed height
		ethereumBlockHeightWhereStart = latestObservedHeight.EthereumBlockHeight + 1
	}

	var pg loops.ParanoidGroup

	pg.Go(func() error { return s.runOracle(ctx, ethereumBlockHeightWhereStart) })
	pg.Go(func() error { return s.runSigner(ctx, hyperionIDHash) })
	pg.Go(func() error { return s.runBatchCreator(ctx) })
	pg.Go(func() error { return s.runRelayer(ctx) })
	pg.Go(func() error { return s.runUpdater(ctx) })
	pg.Go(func() error { return s.runExternalData(ctx) })
	if s.cfg.RelayValsets {
		pg.Go(func() error { return s.runValsetManager(ctx) })
	}

	return pg.Wait()
}

// startRelayerMode runs orchestrator processes that only relay specific
// messages that do not require a validator's signature. This mode is run
// alongside a non-validator helios node
// func (s *Orchestrator) startRelayerMode(ctx context.Context) error {
// 	log.Infoln("running orchestrator in relayer mode")

// 	var pg loops.ParanoidGroup

// 	pg.Go(func() error { return s.runBatchCreator(ctx) })
// 	pg.Go(func() error { return s.runRelayer(ctx) })
// 	pg.Go(func() error { return s.runUpdater(ctx) })

// 	return pg.Wait()
// }

func (s *Orchestrator) getLastClaimBlockHeight(ctx context.Context, helios helios.Network) (uint64, error) {
	claim, err := helios.LastClaimEventByAddr(ctx, s.cfg.HyperionId, s.cfg.CosmosAddr)
	if err != nil {
		s.logger.Info("SSSSS", "err", err)
		return 0, err
	}
	s.logger.Info("SSSSS2", "claim", claim.EthereumEventHeight)

	return claim.EthereumEventHeight, nil
}

func (s *Orchestrator) retry(ctx context.Context, fn func() error) error {
	return retry.Do(fn,
		retry.Context(ctx),
		retry.Delay(200*time.Millisecond),
		retry.Attempts(s.maxAttempts),
		retry.OnRetry(func(n uint, err error) {
			if strings.Contains(err.Error(), "unavailable on our public API") || strings.Contains(err.Error(), "no contract code at given address") || strings.Contains(err.Error(), "History has been pruned for this block") || strings.Contains(err.Error(), "public API") {
				usedRpc := provider.GetCurrentRPCURL(ctx)
				if usedRpc != "" {
					s.ethereum.PenalizeRpc(usedRpc, 1)
					s.logger.WithField("rpc", usedRpc).Debug("Penalized RPC for unavailable on our public API")
				}
				return
			}
			if strings.Contains(err.Error(), "no RPC clients available") {
				s.logger.Warningf("no RPC clients available, refreshing rpcs... (#%d)", n+1)
				rpcs, err := s.global.GetRpcs(s.cfg.ChainId)
				if err != nil {
					s.logger.WithError(err).Warningf("failed to get rpcs")
					return
				}
				s.ethereum.SetRpcs(rpcs)
				return
			}
			s.logger.WithError(err).Warningf("loop error, retrying... (#%d)", n+1)
		}))
}

func (s *Orchestrator) IsStaticRpcAnonymous() bool {
	settings, err := storage.GetChainSettings(s.cfg.ChainId, map[string]interface{}{})
	if err != nil {
		return false
	}
	if _, ok := settings["static_rpc_anonymous"]; !ok {
		return false
	}
	return settings["static_rpc_anonymous"].(bool)
}

func (s *Orchestrator) UpdateRpcs() {

	targetNetwork, err := s.global.InitTargetNetwork(s.cfg.ChainParams)
	if err != nil {
		s.logger.WithError(err).Warningf("failed to update rpcs")
		return
	}
	s.ethereum = *targetNetwork
}
