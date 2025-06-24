package orchestrator

import (
	"context"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

const (
	// Minimum number of confirmations for an Ethereum block to be considered valid
	ethBlockConfirmationDelay uint64 = 4

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
		logEnabled:            strings.Contains(s.cfg.EnabledLogs, "oracle"),
	}

	s.logger.WithField("loop_duration", defaultLoopDur.String()).Debugln("starting Oracle...")

	return loops.RunLoop(ctx, s.ethereum, defaultLoopDur, func() error {

		// if !s.isRegistered() {
		// 	oracle.Log().Infoln("Orchestrator not registered, skipping...")
		// 	return nil
		// }

		err := oracle.observeEthEvents(ctx)
		return err
	})
}

type oracle struct {
	*Orchestrator
	lastResyncWithHelios  time.Time
	lastObservedEthHeight uint64
	logEnabled            bool
}

func (l *oracle) Log() log.Logger {
	return l.logger.WithField("loop", "Oracle").WithField("chain", l.cfg.ChainName)
}

func (l *oracle) observeEthEvents(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	// check if validator is in the active set since claims will fail otherwise
	vs, err := l.helios.CurrentValset(ctx, l.cfg.HyperionId)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to get active validator set on Helios")
		return err
	}

	latestObservedHeight, err := l.helios.QueryGetLastObservedEthereumBlockHeight(ctx, l.cfg.HyperionId)

	if err != nil {
		return errors.Wrap(err, "failed to get latest valsets on Helios")
	}

	if latestObservedHeight.EthereumBlockHeight <= l.cfg.ChainParams.BridgeContractStartHeight && !l.firstTimeSync { // first time total sync needed
		l.lastObservedEthHeight = l.cfg.ChainParams.BridgeContractStartHeight
		l.Log().Info("First Time Hyperion total sync needed BridgeContractStartHeight: ", l.lastObservedEthHeight)
		l.firstTimeSync = true
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
		_, err := l.helios.GetLatestBlockHeight(ctx)
		if err != nil {
			l.Log().WithError(err).Errorln("Connected node is down, cannot make claims...")
			return err
		}
		//l.cfg.CosmosAddr.String()
		validator, err := l.helios.GetValidator(ctx, l.cfg.ValidatorAddress.String())
		if err != nil {
			l.Log().WithError(err).Errorln("failed to get validator on " + l.cfg.ChainName)
			return err
		}
		if validator.Jailed {
			// todo try to unjail
			l.Log().WithFields(log.Fields{"latest_helios_block": vs.Height, "validator": validator.Description.Moniker}).Warningln("validator jailed, cannot make claims...")
			err = l.helios.SendUnjail(ctx, l.cfg.ValidatorAddress.String())
			if err != nil {
				l.Log().WithError(err).Errorln("failed to unjail validator on " + l.cfg.ChainName)
				return err
			}
			//failed to get validator on BSC Testnet
			// chain="BSC Testnet" error="rpc error: code = Unknown desc = codespace sdk code 35:
			/// internal logic error: hrp does not match bech32 prefix: expected 'heliosvaloper' got 'helios'" loop=Oracle
			return nil
		}
		err = l.helios.SendSetOrchestratorAddresses(ctx, uint64(l.cfg.HyperionId), l.cfg.EthereumAddr.String())
		if err != nil {
			return err
		}
		return nil
	}

	latestHeight, err := l.getLatestEthHeight(ctx)
	if err != nil {
		l.Log().WithError(err).Errorln("failed to get latest " + l.cfg.ChainName + " height")
		return err
	}

	targetHeight := latestHeight

	// not enough blocks on ethereum yet
	if targetHeight <= ethBlockConfirmationDelay {
		l.Log().Debugln("not enough blocks on " + l.cfg.ChainName)
		return nil
	}

	// ensure that latest block has minimum confirmations
	targetHeight = targetHeight - ethBlockConfirmationDelay
	if targetHeight <= l.lastObservedEthHeight {
		l.Log().Infoln("Synced", l.lastObservedEthHeight, "to", targetHeight)
		return nil
	}

	targetHeightForSync := targetHeight
	for i := 0; i < 100 && latestHeight > targetHeightForSync; i++ {
		if targetHeightForSync > l.lastObservedEthHeight+defaultBlocksToSearch {
			targetHeightForSync = l.lastObservedEthHeight + defaultBlocksToSearch
		}
		if err := l.syncToTargetHeight(ctx, latestHeight, targetHeightForSync); err != nil {
			return err
		}
		targetHeightForSync = targetHeightForSync + defaultBlocksToSearch
	}

	if time.Since(l.lastResyncWithHelios) >= resyncInterval {
		if err := l.autoResync(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (l *oracle) syncToTargetHeight(ctx context.Context, latestHeight uint64, targetHeight uint64) error {

	l.Orchestrator.SetHeight(l.lastObservedEthHeight)
	l.Orchestrator.SetTargetHeight(latestHeight - ethBlockConfirmationDelay)

	l.Log().Infoln("Sync", strconv.FormatUint(targetHeight-l.lastObservedEthHeight, 10), "Blocks", strconv.FormatUint(l.lastObservedEthHeight, 10)+"("+strconv.FormatUint(((l.lastObservedEthHeight*100)/(latestHeight-ethBlockConfirmationDelay)), 10)+"%) to", strconv.FormatUint(targetHeight, 10)+"("+strconv.FormatUint(((targetHeight*100)/(latestHeight-ethBlockConfirmationDelay)), 10)+"%)")

	events, err := l.getEthEvents(ctx, l.lastObservedEthHeight, targetHeight)
	if err != nil {
		l.Log().WithError(err).Errorln("failed to get events on " + l.cfg.ChainName)
		return err
	}

	lastObservedEventNonce, err := l.helios.QueryGetLastObservedEventNonce(ctx, l.cfg.HyperionId)
	if err != nil {
		l.Log().WithError(err).Errorln("failed to get last observed event nonce on " + l.cfg.ChainName)
		return err
	}

	lastEventNonce, err := l.ethereum.GetLastEventNonce(ctx)
	if err != nil {
		l.Log().WithError(err).Errorln("failed to get last event nonce on " + l.cfg.ChainName)
		return err
	}

	if l.logEnabled {
		l.Log().Infoln("lastObservedEventNonce: ", lastObservedEventNonce, " lastEventNonce: ", lastEventNonce)
	}
	l.Log().Infoln("events Before Filter", events)
	newEvents := filterEvents(events, lastObservedEventNonce)
	if l.logEnabled {
		l.Log().Infoln("newEvents: ", newEvents)
	}

	sort.Slice(newEvents, func(i, j int) bool {
		return newEvents[i].Nonce() < newEvents[j].Nonce()
	})

	if len(newEvents) == 0 {
		l.Log().Infoln("NO EVENTS DETECTED 0", l.ethereum.GetLastUsedRpc())
		if l.logEnabled {
			l.Log().WithFields(log.Fields{"last_observed_event_nonce": lastObservedEventNonce, "eth_block_start": l.lastObservedEthHeight, "eth_block_end": targetHeight}).Infoln("oracle no new events on " + l.cfg.ChainName)
		}
		l.lastObservedEthHeight = targetHeight
		return nil
	}

	if len(newEvents) > 0 {
		l.Log().Infoln("SOME EVENTS DETECTED %d", len(newEvents))
	}

	if l.logEnabled {
		l.Log().WithFields(log.Fields{"event_helios_nonce": lastObservedEventNonce, "event_ethereum_nonce": newEvents[0].Nonce()}).Infoln("try oracle relay to helios")
	}

	if newEvents[0].Nonce() > lastObservedEventNonce+1 {
		// we missed an event
		lastObservedHeight, err := l.helios.QueryGetLastObservedEthereumBlockHeight(ctx, l.cfg.HyperionId)
		if err != nil {
			return err
		}
		l.lastObservedEthHeight = lastObservedHeight.EthereumBlockHeight
		// move back to the last observed event height
		l.Log().WithFields(log.Fields{"current_helios_nonce": lastObservedEventNonce, "wanted_nonce": lastObservedEventNonce + 1, "actual_ethereum_nonce": newEvents[0].Nonce()}).Infoln("orchestrator missed an " + l.cfg.ChainName + " event. Restarting block search from last observed claim...")
		return nil
	}

	if err := l.sendNewEventClaims(ctx, newEvents); err != nil {
		log.Info("err: ", err)
		return err
	}

	if l.logEnabled {
		l.Log().WithFields(log.Fields{"claims": len(newEvents), "eth_block_start": l.lastObservedEthHeight, "eth_block_end": latestHeight}).Infoln("sent new event claims to Helios")
	}
	l.lastObservedEthHeight = targetHeight

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

		var msgs []cosmostypes.Msg
		for _, event := range newEvents {
			msg, err := l.prepareSendEthEventClaim(ctx, event)
			if err != nil {
				return err
			}
			msgs = append(msgs, msg)

			if len(msgs) > 10 {
				log.Infoln("sending bulk of ", len(msgs), "claims messages")
				resp, err := l.helios.SyncBroadcastMsgs(ctx, msgs)
				if err != nil {
					return err
				}
				cost, err := l.helios.GetTxCost(ctx, resp.TxHash)
				if err == nil {
					storage.UpdateFeesFile(big.NewInt(0), "", cost, resp.TxHash, uint64(resp.Height), uint64(42000), "CLAIM")
				}
				msgs = []cosmostypes.Msg{}
				time.Sleep(1100 * time.Millisecond)
			}
		}

		if len(msgs) > 0 {
			log.Infoln("sending bulk of ", len(msgs), "claims messages")
			resp, err := l.helios.SyncBroadcastMsgs(ctx, msgs)
			if err != nil {
				return err
			}
			cost, err := l.helios.GetTxCost(ctx, resp.TxHash)
			if err == nil {
				storage.UpdateFeesFile(big.NewInt(0), "", cost, resp.TxHash, uint64(resp.Height), uint64(42000), "CLAIM")
			}
			time.Sleep(1100 * time.Millisecond)
		}
		// for _, event := range newEvents {
		// 	resp, err := l.sendEthEventClaim(ctx, event)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	cost, err := l.helios.GetTxCost(ctx, resp.TxHash)
		// 	if err == nil {
		// 		storage.UpdateFeesFile(big.NewInt(0), "", cost, resp.TxHash, uint64(resp.Height), uint64(42000), "CLAIM")
		// 	}

		// 	// Considering block time ~1s on Helios chain, adding Sleep to make sure new event is sent
		// 	// only after previous event is executed successfully. Otherwise it will through `non contiguous event nonce` failing CheckTx.
		// 	time.Sleep(1100 * time.Millisecond)
		// }

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

func (l *oracle) prepareSendEthEventClaim(ctx context.Context, ev event) (cosmostypes.Msg, error) {
	rpc := ""
	if !l.IsStaticRpcAnonymous() {
		rpc = l.ethereum.GetLastUsedRpc()
	}
	switch e := ev.(type) {
	case *deposit:
		ev := hyperionevents.HyperionSendToHeliosEvent(*e)
		return l.helios.SendDepositClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
	case *valsetUpdate:
		ev := hyperionevents.HyperionValsetUpdatedEvent(*e)
		return l.helios.SendValsetClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
	case *withdrawal:
		ev := hyperionevents.HyperionTransactionBatchExecutedEvent(*e)
		return l.helios.SendWithdrawalClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
	case *erc20Deployment:
		ev := hyperionevents.HyperionERC20DeployedEvent(*e)
		return l.helios.SendERC20DeployedClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
	default:
		panic(errors.Errorf("unknown ev type %T", e))
	}
}

// func (l *oracle) sendEthEventClaim(ctx context.Context, ev event) (*cosmostypes.TxResponse, error) {
// 	switch e := ev.(type) {
// 	case *deposit:
// 		ev := hyperionevents.HyperionSendToHeliosEvent(*e)
// 		return l.helios.SendDepositClaim(ctx, l.cfg.HyperionId, &ev, l.ethereum.GetLastUsedRpc())
// 	case *valsetUpdate:
// 		ev := hyperionevents.HyperionValsetUpdatedEvent(*e)
// 		return l.helios.SendValsetClaim(ctx, l.cfg.HyperionId, &ev, l.ethereum.GetLastUsedRpc())
// 	case *withdrawal:
// 		ev := hyperionevents.HyperionTransactionBatchExecutedEvent(*e)
// 		return l.helios.SendWithdrawalClaim(ctx, l.cfg.HyperionId, &ev, l.ethereum.GetLastUsedRpc())
// 	case *erc20Deployment:
// 		ev := hyperionevents.HyperionERC20DeployedEvent(*e)
// 		return l.helios.SendERC20DeployedClaim(ctx, l.cfg.HyperionId, &ev, l.ethereum.GetLastUsedRpc())
// 	default:
// 		panic(errors.Errorf("unknown ev type %T", e))
// 	}
// }

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
