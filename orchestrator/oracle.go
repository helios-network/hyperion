package orchestrator

import (
	"context"
	"sort"
	"time"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

const (
	// Minimum number of confirmations for an Ethereum block to be considered valid
	ethBlockConfirmationDelay uint64 = 12

	// Maximum block range for Ethereum event query. If the orchestrator has been offline for a long time,
	// the oracle loop can potentially run longer than defaultLoopDur due to a surge of events. This usually happens
	// when there are more than ~50 events to claim in a single run.
	defaultBlocksToSearch uint64 = 2000

	// Auto re-sync to catch up the validator's last observed event nonce. Reasons why event nonce fall behind:
	// 1. It takes some time for events to be indexed on Ethereum. So if hyperion queried events immediately as block produced, there is a chance the event is missed.
	//  We need to re-scan this block to ensure events are not missed due to indexing delay.
	// 2. if validator was in UnBonding state, the claims broadcasted in last iteration are failed.
	// 3. if infura call failed while filtering events, the hyperion missed to broadcast claim events occured in last iteration.
	resyncInterval = 24 * time.Hour
)

// runOracle is responsible for making sure that Ethereum events are retrieved from the Ethereum blockchain
// and ferried over to Cosmos where they will be used to issue tokens or process batches.
func (s *Orchestrator) runOracle(ctx context.Context, lastObservedBlock uint64) error {
	oracle := oracle{
		Orchestrator:          s,
		lastObservedEthHeight: lastObservedBlock,
		lastResyncWithHelios:  time.Now(),
	}

	s.logger.WithField("loop_duration", defaultLoopDur.String()).Debugln("starting Oracle...")

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		log.Info("observing")
		err := oracle.observeEthEvents(ctx)
		log.Info("observing done")
		return err
	})
}

type oracle struct {
	*Orchestrator
	lastResyncWithHelios  time.Time
	lastObservedEthHeight uint64
}

func (l *oracle) Log() log.Logger {
	return l.logger.WithField("loop", "Oracle")
}

func (l *oracle) observeEthEvents(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	// check if validator is in the active set since claims will fail otherwise
	vs, err := l.helios.CurrentValset(ctx, l.cfg.HyperionId)
	if err != nil {
		l.logger.WithError(err).Warningln("failed to get active validator set on Helios")
		return err
	}

	latestObservedHeight, err := l.helios.QueryGetLastObservedEthereumBlockHeight(ctx, l.cfg.HyperionId)

	if err != nil {
		return errors.Wrap(err, "failed to get latest valsets on Helios")
	}

	if latestObservedHeight.EthereumBlockHeight <= l.cfg.ChainParams.BridgeContractStartHeight { // first time total sync needed
		l.lastObservedEthHeight = l.cfg.ChainParams.BridgeContractStartHeight
		l.logger.Info("First Time Hyperion total sync needed BridgeContractStartHeight: ", l.lastObservedEthHeight)
	}

	// ajouter la possibilite de forcer un restart depuis un certain blockHeight en config

	bonded := false
	for _, v := range vs.Members {
		if l.cfg.EthereumAddr.Hex() == v.EthereumAddress {
			bonded = true
		}
	}

	if !bonded {
		l.Log().WithFields(log.Fields{"latest_helios_block": vs.Height}).Warningln("validator not in active set, cannot make claims...")
		err := l.helios.SendSetOrchestratorAddresses(ctx, uint64(l.cfg.HyperionId), l.cfg.EthereumAddr.String())
		if err != nil {
			return err
		}
		return nil
	}

	latestHeight, err := l.getLatestEthHeight(ctx)
	if err != nil {
		return err
	}

	// not enough blocks on ethereum yet
	if latestHeight <= ethBlockConfirmationDelay {
		l.Log().Debugln("not enough blocks on Ethereum")
		return nil
	}

	// ensure that latest block has minimum confirmations
	latestHeight = latestHeight - ethBlockConfirmationDelay
	if latestHeight <= l.lastObservedEthHeight {
		l.Log().WithFields(log.Fields{"latest": latestHeight, "observed": l.lastObservedEthHeight}).Debugln("latest Ethereum height already observed")
		return nil
	}

	// ensure the block range is within defaultBlocksToSearch
	if latestHeight > l.lastObservedEthHeight+defaultBlocksToSearch {
		latestHeight = l.lastObservedEthHeight + defaultBlocksToSearch
	}

	l.Log().Infoln("GET ETHEREUM EVENTS FOR HEIGHT", latestHeight, latestHeight+defaultBlocksToSearch)

	events, err := l.getEthEvents(ctx, l.lastObservedEthHeight, latestHeight)
	if err != nil {
		return err
	}
	log.Info("events: ", events)

	lastObservedEventNonce, err := l.helios.QueryGetLastObservedEventNonce(ctx, l.cfg.HyperionId)
	if err != nil {
		return err
	}

	lastEventNonce, err := l.ethereum.GetLastEventNonce(ctx)
	if err != nil {
		return err
	}

	log.Info("lastObservedEventNonce: ", lastObservedEventNonce, " lastEventNonce: ", lastEventNonce)
	newEvents := filterEvents(events, lastObservedEventNonce)
	log.Info("newEvents: ", newEvents)

	sort.Slice(newEvents, func(i, j int) bool {
		return newEvents[i].Nonce() < newEvents[j].Nonce()
	})

	if len(newEvents) == 0 {
		l.Log().Infoln("NO EVENTS DETECTED 0")
		l.Log().WithFields(log.Fields{"last_observed_event_nonce": lastObservedEventNonce, "eth_block_start": l.lastObservedEthHeight, "eth_block_end": latestHeight}).Infoln("oracle no new events on Ethereum")
		l.lastObservedEthHeight = latestHeight
		return nil
	}

	if len(newEvents) > 0 {
		l.Log().Infoln("SOME EVENTS DETECTED %d", len(newEvents))
	}

	l.Log().WithFields(log.Fields{"event_helios_nonce": lastObservedEventNonce, "event_ethereum_nonce": newEvents[0].Nonce()}).Infoln("try oracle relay to helios")

	if newEvents[0].Nonce() > lastObservedEventNonce+1 {
		// we missed an event
		lastObservedHeight, err := l.helios.QueryGetLastObservedEthereumBlockHeight(ctx, l.cfg.HyperionId)
		if err != nil {
			return err
		}
		l.lastObservedEthHeight = lastObservedHeight.EthereumBlockHeight
		// move back to the last observed event height
		l.Log().WithFields(log.Fields{"current_helios_nonce": lastObservedEventNonce, "wanted_nonce": lastObservedEventNonce + 1, "actual_ethereum_nonce": newEvents[0].Nonce()}).Infoln("orchestrator missed an Ethereum event. Restarting block search from last observed claim...")
		return nil
	}

	if err := l.sendNewEventClaims(ctx, newEvents); err != nil {
		log.Info("err: ", err)
		return err
	}

	l.Log().WithFields(log.Fields{"claims": len(newEvents), "eth_block_start": l.lastObservedEthHeight, "eth_block_end": latestHeight}).Infoln("sent new event claims to Helios")
	l.lastObservedEthHeight = latestHeight

	if time.Since(l.lastResyncWithHelios) >= resyncInterval {
		if err := l.autoResync(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *oracle) getEthEvents(ctx context.Context, startBlock, endBlock uint64) ([]event, error) {
	var events []event
	scanEthEventsFn := func() error {
		events = nil // clear previous result in case a retry occurred

		depositEvents, err := l.ethereum.GetSendToHeliosEvents(startBlock, endBlock)
		if err != nil {
			return errors.Wrap(err, "failed to get SendToHelios events")
		}

		withdrawalEvents, err := l.ethereum.GetTransactionBatchExecutedEvents(startBlock, endBlock)
		if err != nil {
			return errors.Wrap(err, "failed to get TransactionBatchExecuted events")
		}

		erc20DeploymentEvents, err := l.ethereum.GetHyperionERC20DeployedEvents(startBlock, endBlock)
		if err != nil {
			return errors.Wrap(err, "failed to get ERC20Deployed events")
		}

		valsetUpdateEvents, err := l.ethereum.GetValsetUpdatedEvents(startBlock, endBlock)
		if err != nil {
			return errors.Wrap(err, "failed to get ValsetUpdated events")
		}

		for _, e := range depositEvents {
			ev := deposit(*e)
			events = append(events, &ev)
		}

		for _, e := range withdrawalEvents {
			ev := withdrawal(*e)
			events = append(events, &ev)
		}

		for _, e := range valsetUpdateEvents {
			ev := valsetUpdate(*e)
			events = append(events, &ev)
		}

		for _, e := range erc20DeploymentEvents {
			ev := erc20Deployment(*e)
			events = append(events, &ev)
		}

		return nil
	}

	if err := l.retry(ctx, scanEthEventsFn); err != nil {
		return nil, err
	}

	return events, nil
}

func (l *oracle) getLatestEthHeight(ctx context.Context) (uint64, error) {
	latestHeight := uint64(0)
	fn := func() error {
		h, err := l.ethereum.GetHeaderByNumber(ctx, nil)
		if err != nil {
			return errors.Wrap(err, "failed to get latest ethereum header")
		}

		latestHeight = h.Number.Uint64()
		return nil
	}

	if err := l.retry(ctx, fn); err != nil {
		return 0, err
	}

	return latestHeight, nil
}

func (l *oracle) getLastClaimEvent(ctx context.Context) (*hyperiontypes.LastClaimEvent, error) {
	var claim *hyperiontypes.LastClaimEvent
	fn := func() (err error) {
		claim, err = l.helios.LastClaimEventByAddr(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr)
		return
	}

	if err := l.retry(ctx, fn); err != nil {
		return nil, err
	}

	return claim, nil
}

func (l *oracle) sendNewEventClaims(ctx context.Context, events []event) error {
	sendEventsFn := func() error {
		// in case sending one of more claims fails, we reload the latest claimed nonce to filter processed events
		lastClaim, err := l.helios.LastClaimEventByAddr(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr)
		if err != nil {
			return err
		}

		newEvents := filterEvents(events, lastClaim.EthereumEventNonce)
		if len(newEvents) == 0 {
			return nil
		}

		for _, event := range newEvents {
			if err := l.sendEthEventClaim(ctx, event); err != nil {
				return err
			}

			// Considering block time ~1s on Helios chain, adding Sleep to make sure new event is sent
			// only after previous event is executed successfully. Otherwise it will through `non contiguous event nonce` failing CheckTx.
			time.Sleep(1100 * time.Millisecond)
		}

		return nil
	}

	if err := l.retry(ctx, sendEventsFn); err != nil {
		return err
	}

	return nil
}

func (l *oracle) autoResync(ctx context.Context) error {
	var height uint64
	fn := func() (err error) {
		height, err = l.getLastClaimBlockHeight(ctx, l.helios)
		return
	}

	if err := l.retry(ctx, fn); err != nil {
		return err
	}

	l.Log().WithFields(log.Fields{"last_resync": l.lastResyncWithHelios.String(), "last_claimed_eth_height": height}).Infoln("auto resyncing with last claimed event on Helios")

	l.lastObservedEthHeight = height
	l.lastResyncWithHelios = time.Now()

	return nil
}

func (l *oracle) sendEthEventClaim(ctx context.Context, ev event) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Load Failed .env: %v", err)
	}

	switch e := ev.(type) {
	case *deposit:
		ev := hyperionevents.HyperionSendToHeliosEvent(*e)
		return l.helios.SendDepositClaim(ctx, l.cfg.HyperionId, &ev)
	case *valsetUpdate:
		ev := hyperionevents.HyperionValsetUpdatedEvent(*e)
		return l.helios.SendValsetClaim(ctx, l.cfg.HyperionId, &ev)
	case *withdrawal:
		ev := hyperionevents.HyperionTransactionBatchExecutedEvent(*e)
		return l.helios.SendWithdrawalClaim(ctx, l.cfg.HyperionId, &ev)
	case *erc20Deployment:
		ev := hyperionevents.HyperionERC20DeployedEvent(*e)
		return l.helios.SendERC20DeployedClaim(ctx, l.cfg.HyperionId, &ev)
	default:
		panic(errors.Errorf("unknown ev type %T", e))
	}
}

type (
	deposit         hyperionevents.HyperionSendToHeliosEvent
	valsetUpdate    hyperionevents.HyperionValsetUpdatedEvent
	withdrawal      hyperionevents.HyperionTransactionBatchExecutedEvent
	erc20Deployment hyperionevents.HyperionERC20DeployedEvent

	event interface {
		Nonce() uint64
	}
)

func filterEvents(events []event, nonce uint64) (filtered []event) {
	for _, e := range events {
		if e.Nonce() > nonce {
			filtered = append(filtered, e)
		}
	}

	return
}

func (o *deposit) Nonce() uint64 {
	return o.EventNonce.Uint64()
}

func (o *valsetUpdate) Nonce() uint64 {
	return o.EventNonce.Uint64()
}

func (o *withdrawal) Nonce() uint64 {
	return o.EventNonce.Uint64()
}

func (o *erc20Deployment) Nonce() uint64 {
	return o.EventNonce.Uint64()
}
