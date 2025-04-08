package orchestrator

import (
	"context"
	"time"

	"cosmossdk.io/errors"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/avast/retry-go"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/metrics"
)

const (
	defaultLoopDur = 60 * time.Second
)

// PriceFeed provides token price for a given contract address
type PriceFeed interface {
	QueryUSDPrice(address gethcommon.Address) (float64, error)
}

type Config struct {
	CosmosAddr           cosmostypes.AccAddress
	HyperionId           uint64
	EthereumAddr         gethcommon.Address
	MinBatchFeeUSD       float64
	RelayValsetOffsetDur time.Duration
	RelayBatchOffsetDur  time.Duration
	RelayValsets         bool
	RelayBatches         bool
	RelayExternalDatas   bool
	RelayerMode          bool
	ChainParams          *hyperiontypes.CounterpartyChainParams
}

type Orchestrator struct {
	logger      log.Logger
	svcTags     metrics.Tags
	cfg         Config
	maxAttempts uint

	helios    helios.Network
	ethereum  ethereum.Network
	priceFeed PriceFeed
}

func NewOrchestrator(
	helios helios.Network,
	eth ethereum.Network,
	priceFeed PriceFeed,
	cfg Config,
) (*Orchestrator, error) {
	o := &Orchestrator{
		logger:      log.DefaultLogger,
		svcTags:     metrics.Tags{"svc": "hyperion_orchestrator"},
		helios:      helios,
		ethereum:    eth,
		priceFeed:   priceFeed,
		cfg:         cfg,
		maxAttempts: 10,
	}

	return o, nil
}

// Run starts all major loops required to make
// up the Orchestrator, all of these are async loops.
func (s *Orchestrator) Run(ctx context.Context, helios helios.Network, eth ethereum.Network) error {
	if s.cfg.RelayerMode {
		return s.startRelayerMode(ctx, helios, eth)
	}

	return s.startValidatorMode(ctx, helios, eth)
}

// startValidatorMode runs all orchestrator processes. This is called
// when hyperion is run alongside a validator helios node.
func (s *Orchestrator) startValidatorMode(ctx context.Context, helios helios.Network, eth ethereum.Network) error {
	log.Infoln("running orchestrator in validator mode")

	// get hyperion ID from contract
	delay, err := time.ParseDuration("1s")
	hyperionIDHash, err := loops.RetryFunction(ctx, func() (gethcommon.Hash, error) {
		return eth.GetHyperionID(ctx)
	}, delay)
	if err != nil {
		s.logger.WithError(err).Fatalln("unable to query hyperion ID from contract")
	}
	hyperionID := hyperionIDHash.Big().Uint64()

	log.Infoln("Our HyperionID", "is", hyperionID, "hash", hyperionIDHash.Hex())

	latestObservedHeight, err := s.helios.QueryGetLastObservedEthereumBlockHeight(ctx, s.cfg.HyperionId)
	if err != nil {
		s.logger.WithError(err).Fatalln("unable to query hyperion module params, is heliades running?")
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
		s.logger.WithError(err).Fatalln("unable to query hyperion module params, is heliades running?")
	}

	lastEventNonce, err := s.ethereum.GetLastEventNonce(ctx)
	if err != nil {
		s.logger.WithError(err).Fatalln("unable to query blockchain state_lastEventNonce, is rpc running?")
	}

	if lastEventNonce.Uint64() > lastObservedEventNonce {
		log.Infoln("lastEventNonce is greater than lastObservedEventNonce, starting from lastObservedEventNonce")
		// start from the next block after the last observed height
		ethereumBlockHeightWhereStart = latestObservedHeight.EthereumBlockHeight + 1
	}

	var pg loops.ParanoidGroup

	pg.Go(func() error { return s.runOracle(ctx, ethereumBlockHeightWhereStart) })
	pg.Go(func() error { return s.runSigner(ctx, hyperionIDHash) })
	pg.Go(func() error { return s.runBatchCreator(ctx) })
	pg.Go(func() error { return s.runRelayer(ctx) })

	return pg.Wait()
}

// startRelayerMode runs orchestrator processes that only relay specific
// messages that do not require a validator's signature. This mode is run
// alongside a non-validator helios node
func (s *Orchestrator) startRelayerMode(ctx context.Context, helios helios.Network, eth ethereum.Network) error {
	log.Infoln("running orchestrator in relayer mode")

	var pg loops.ParanoidGroup

	pg.Go(func() error { return s.runBatchCreator(ctx) })
	pg.Go(func() error { return s.runRelayer(ctx) })

	return pg.Wait()
}

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
			s.logger.WithError(err).Warningf("loop error, retrying... (#%d)", n+1)
		}))
}
