package orchestrator

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"cosmossdk.io/errors"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/avast/retry-go"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/utils"
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
	// GetRpcs(chainId uint64) ([]*hyperiontypes.Rpc, error)
	InitTargetNetworks(counterpartyChainParams *hyperiontypes.CounterpartyChainParams) ([]*ethereum.Network, error)
	GetMinBatchFeeHLS(chainId uint64) float64
	GetMinTxFeeHLS(chainId uint64) float64
	ResetHeliosClient()
	GetHeliosNetwork() *helios.Network
	SyncBroadcastMsgs(ctx context.Context, msgs []sdk.Msg) (*sdk.TxResponse, error)
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
	NativeBalance          string
	HyperionID             uint64
	Height                 uint64
	TargetHeight           uint64
	LastObservedHeight     uint64
	LastObservedEventNonce uint64
	LastEventNonce         uint64
	LastClaimBlockHeight   uint64
	LastClaimEventNonce    uint64
	GasPrice               string

	BatchCount           int
	TxCount              int
	OutBridgedTxCount    int
	InBridgedTxCount     int
	ValsetUpdateCount    int
	ERC20DeploymentCount int
	SkippedRetriedCount  int
	ExternalDataCount    int

	BatchCreatorStatus  string
	ExternalDataStatus  string
	OracleStatus        string
	RelayerStatus       string
	SignerStatus        string
	UpdaterStatus       string
	SkippedStatus       string
	ValsetManagerStatus string
	ErrorStatus         string

	BatchCreatorNextExecutionTimestamp  uint64
	ExternalDataNextExecutionTimestamp  uint64
	OracleNextExecutionTimestamp        uint64
	RelayerNextExecutionTimestamp       uint64
	SignerNextExecutionTimestamp        uint64
	UpdaterNextExecutionTimestamp       uint64
	SkippedNextExecutionTimestamp       uint64
	ValsetManagerNextExecutionTimestamp uint64

	BatchCreatorLastExecutionFinishedTimestamp  uint64
	ExternalDataLastExecutionFinishedTimestamp  uint64
	OracleLastExecutionFinishedTimestamp        uint64
	RelayerLastExecutionFinishedTimestamp       uint64
	SignerLastExecutionFinishedTimestamp        uint64
	UpdaterLastExecutionFinishedTimestamp       uint64
	SkippedLastExecutionFinishedTimestamp       uint64
	ValsetManagerLastExecutionFinishedTimestamp uint64

	IsDepositPaused    bool
	IsWithdrawalPaused bool
}

type Orchestrator struct {
	logger      log.Logger
	svcTags     metrics.Tags
	cfg         Config
	maxAttempts uint

	ethereum      ethereum.Network
	ethereums     []*ethereum.Network
	priceFeed     PriceFeed
	firstTimeSync bool
	global        Global

	height        uint64
	targetHeight  uint64
	valsetManager valsetManager

	HyperionState HyperionState

	CacheSymbol map[gethcommon.Address]string

	Oracle *oracle
}

func NewOrchestrator(
	eths []*ethereum.Network,
	priceFeed PriceFeed,
	cfg Config,
	global Global,
) (*Orchestrator, error) {
	o := &Orchestrator{
		logger: log.WithFields(log.Fields{
			"chain": cfg.ChainName,
		}),
		svcTags:       metrics.Tags{"svc": "hyperion_orchestrator"},
		ethereum:      *eths[0],
		ethereums:     eths,
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
			GasPrice:               "0.0",

			BatchCount:           0,
			OutBridgedTxCount:    0,
			InBridgedTxCount:     0,
			ValsetUpdateCount:    0,
			ERC20DeploymentCount: 0,
			SkippedRetriedCount:  0,
			ExternalDataCount:    0,

			BatchCreatorStatus:  "idle",
			ExternalDataStatus:  "idle",
			OracleStatus:        "idle",
			RelayerStatus:       "idle",
			SignerStatus:        "idle",
			UpdaterStatus:       "idle",
			SkippedStatus:       "idle",
			ValsetManagerStatus: "idle",
			ErrorStatus:         "okay",

			BatchCreatorNextExecutionTimestamp:  0,
			ExternalDataNextExecutionTimestamp:  0,
			OracleNextExecutionTimestamp:        0,
			RelayerNextExecutionTimestamp:       0,
			SignerNextExecutionTimestamp:        0,
			UpdaterNextExecutionTimestamp:       0,
			SkippedNextExecutionTimestamp:       0,
			ValsetManagerNextExecutionTimestamp: 0,

			BatchCreatorLastExecutionFinishedTimestamp:  0,
			ExternalDataLastExecutionFinishedTimestamp:  0,
			OracleLastExecutionFinishedTimestamp:        0,
			RelayerLastExecutionFinishedTimestamp:       0,
			SignerLastExecutionFinishedTimestamp:        0,
			UpdaterLastExecutionFinishedTimestamp:       0,
			SkippedLastExecutionFinishedTimestamp:       0,
			ValsetManagerLastExecutionFinishedTimestamp: 0,

			IsDepositPaused:    false,
			IsWithdrawalPaused: false,
		},

		CacheSymbol: make(map[gethcommon.Address]string),

		Oracle: nil,
	}

	return o, nil
}

// Run starts all major loops required to make
// up the Orchestrator, all of these are async loops.
func (s *Orchestrator) Run(ctx context.Context) error {
	return s.startValidatorMode(ctx)
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
	return *s.global.GetHeliosNetwork()
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
func (s *Orchestrator) startValidatorMode(ctx context.Context) error {
	s.logger.Infoln("running orchestrator in validator mode")

	fmt.Println("startValidatorMode")

	// update native balance
	err := s.UpdateNativeBalance(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to update native balance")
	}

	// get hyperion ID from contract
	hyperionIDHash, err := s.ethereum.GetHyperionID(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to query hyperion ID from contract")
	}
	hyperionID := hyperionIDHash.Big().Uint64()

	s.logger.Infoln("Our HyperionID", "is", hyperionID, "hash", hyperionIDHash.Hex(), "hyperionAddress", s.ethereum.GetHyperionContractAddress().Hex())

	latestObservedHeight, err := s.GetHelios().QueryGetLastObservedEthereumBlockHeight(ctx, s.cfg.HyperionId)
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

	lastObservedEventNonce, err := s.GetHelios().QueryGetLastObservedEventNonce(ctx, s.cfg.HyperionId)
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

	// check if deposit is paused
	isDepositPaused, err := s.ethereum.IsDepositPaused(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to query deposit pause status")
	}
	s.HyperionState.IsDepositPaused = isDepositPaused
	s.HyperionState.IsWithdrawalPaused = s.cfg.ChainParams.Paused

	var pg loops.ParanoidGroup

	pg.Go(func() error { return s.runOracle(ctx, ethereumBlockHeightWhereStart) })
	// pg.Go(func() error { return s.runSkipped(ctx) })
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
			if strings.Contains(err.Error(), "no RPC clients available") {
				s.logger.Warningf("no RPC clients available, refreshing rpcs... (#%d)", n+1)
				return
			}
			s.logger.WithError(err).Warningf("loop error, retrying... (#%d)", n+1)
		}))
}

func (s *Orchestrator) IsStaticRpcAnonymous() bool {
	settings, err := storage.GetChainSettings(s.cfg.ChainId)
	if err != nil {
		return false
	}
	if _, ok := settings["static_rpc_anonymous"]; !ok {
		return false
	}
	return settings["static_rpc_anonymous"].(bool)
}

func (s *Orchestrator) RotateRpc() {
	usedRpc := s.ethereum.GetRpc().Url

	lstOfclients := make([]*ethereum.Network, 0)

	for _, eth := range s.ethereums {
		if (*eth).GetRpc().Url != usedRpc {
			lstOfclients = append(lstOfclients, eth)
		}
	}
	if len(lstOfclients) == 0 {
		s.logger.Warning("No other rpcs available, using the same rpc")
		return
	}
	s.ethereum = *lstOfclients[rand.Intn(len(lstOfclients))]
	s.logger.Info("Rotated rpc to rpc ", s.ethereum.GetRpc().Url, " / total rpcs ", len(lstOfclients))
}

func (s *Orchestrator) UpdateNativeBalance(ctx context.Context) error {
	nativeBalance, err := s.ethereum.GetNativeBalance(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to get native balance")
	}
	s.HyperionState.NativeBalance = utils.FormatBigStringToFloat64(nativeBalance.String(), 18)
	return nil
}

func (s *Orchestrator) ResetHeliosClient() {
	fmt.Println("ResetHeliosClient called")
	s.global.ResetHeliosClient()
}

func (s *Orchestrator) ResetEthereum() {
	targetNetworks, err := s.global.InitTargetNetworks(s.cfg.ChainParams)
	if err != nil {
		s.logger.Error("Error resetting ethereum", "error", err)
		return
	}
	if len(targetNetworks) == 0 {
		s.logger.Error("No target networks found for chain", "chain", s.cfg.ChainName)
		return
	}
	s.logger.Info("Reset ethereum", "chain", s.cfg.ChainName, "new ethereum", (*targetNetworks[0]).GetRpc().Url)
	s.ethereums = targetNetworks
	s.ethereum = *targetNetworks[0]
}
