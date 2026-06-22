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

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/rpcerrors"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

const (
	// minBlocksToSearch is the smallest window we'll ever ask a provider to
	// scan. Going below this buys nothing and usually means the provider is
	// misbehaving, not the window.
	minBlocksToSearch float64 = 50

	// defaultMaxBlocksToSearch caps how far we'll re-grow the scan window
	// after consecutive successes. Providers that handled larger ranges in
	// the past can still bite us under load.
	defaultMaxBlocksToSearch float64 = 10000

	// blocksGrowAfterSuccesses controls how many successful chunks we must
	// accumulate before widening the scan window again.
	blocksGrowAfterSuccesses = 10

	// blocksGrowFactor is how aggressively the scan window grows back.
	blocksGrowFactor = 1.5
)

const (
// Minimum number of confirmations for an Ethereum block to be considered valid
// ethBlockConfirmationDelay uint64 = 4

// Maximum block range for Ethereum event query. If the orchestrator has been offline for a long time,
// the oracle loop can potentially run longer than defaultLoopDur due to a surge of events. This usually happens
// when there are more than ~50 events to claim in a single run.
// defaultBlocksToSearch uint64 = 2000

// Auto re-sync to catch up the validator's last observed event nonce. Reasons why event nonce fall behind:
//  1. It takes some time for events to be indexed on Ethereum. So if hyperion queried events immediately as block produced, there is a chance the event is missed.
//     We need to re-scan this block to ensure events are not missed due to indexing delay.
//  2. if validator was in UnBonding state, the claims broadcasted in last iteration are failed.
//  3. if infura call failed while filtering events, the hyperion missed to broadcast claim events occured in last iteration.
//
// resyncInterval = 24 * time.Hour
)

// runOracle is responsible for making sure that Ethereum events are retrieved from the Ethereum blockchain
// and ferried over to Cosmos where they will be used to issue tokens or process batches.
func (s *Orchestrator) runOracle(ctx context.Context, lastObservedBlock uint64) error {
	oracle := oracle{
		Orchestrator:            s,
		lastObservedEthHeight:   lastObservedBlock,
		lastResyncWithHelios:    time.Now(),
		logEnabled:              strings.Contains(s.cfg.EnabledLogs, "oracle"),
		missedEventsBlockHeight: 0,
	}

	s.Oracle = &oracle

	s.logger.WithField("loop_duration", defaultLoopDur.String()).Debugln("starting Oracle...")

	ticker := time.NewTicker(defaultLoopDur)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if s.HyperionState.OracleStatus == "running" {
				continue
			}

			start := time.Now()
			s.HyperionState.OracleStatus = "running"
			if err := oracle.observeEthEvents(ctx); err != nil {
				s.logger.WithError(err).Errorln("oracle function returned an error")
			}
			s.HyperionState.OracleStatus = "idle"
			s.HyperionState.OracleLastExecutionFinishedTimestamp = uint64(start.Unix())
			s.HyperionState.OracleNextExecutionTimestamp = uint64(start.Add(defaultLoopDur).Unix())
		case <-ctx.Done():
			return nil
		}
	}
}

type oracle struct {
	*Orchestrator
	lastResyncWithHelios    time.Time
	lastObservedEthHeight   uint64
	logEnabled              bool
	missedEventsBlockHeight uint64
	consecutiveSuccesses    int
}

func (l *oracle) Log() log.Logger {
	return l.logger.WithField("loop", "Oracle").WithField("chain", l.cfg.ChainName)
}

func (l *oracle) observeEthEvents(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	settings, err := storage.GetChainSettings(l.cfg.ChainId)
	if err != nil {
		return errors.Wrap(err, "failed to get chain settings")
	}
	defaultBlocksToSearch, ok := settings["oracle_eth_default_blocks_to_search"].(float64)

	if !ok {
		l.Log().Infoln("oracle_eth_default_blocks_to_search not found in chain settings, using default value 2000")
		defaultBlocksToSearch = 2000
	} else {
		l.Log().Infoln("oracle_eth_default_blocks_to_search found in chain settings, using value", defaultBlocksToSearch)
	}

	// Self-heal when a past rate-limit incident shrank the window below the
	// usable floor. Previously this required manual intervention.
	if defaultBlocksToSearch < minBlocksToSearch {
		l.Log().Warningln("oracle_eth_default_blocks_to_search below floor, resetting to", minBlocksToSearch)
		defaultBlocksToSearch = minBlocksToSearch
		settings["oracle_eth_default_blocks_to_search"] = defaultBlocksToSearch
		_ = storage.SetChainSettings(l.cfg.ChainId, settings)
	}
	if l.Orchestrator.HyperionState.ErrorStatus == "oracle_eth_default_blocks_to_search is less than 10, please increase the value" {
		l.Orchestrator.HyperionState.ErrorStatus = "okay"
	}

	maxBlocksToSearch, ok := settings["oracle_eth_max_blocks_to_search"].(float64)
	if !ok || maxBlocksToSearch < defaultBlocksToSearch {
		maxBlocksToSearch = defaultMaxBlocksToSearch
	}

	// oracle_skip_pruned_history: when every configured RPC rejects a chunk
	// with a "history pruned / archive required" error, there is no way to
	// scan those blocks — any events in them are permanently unreachable.
	// With this flag the oracle advances past the unreadable range instead
	// of looping forever. Default false: data-safety over liveness.
	skipPrunedHistory, _ := settings["oracle_skip_pruned_history"].(bool)

	ethBlockConfirmationDelay, ok := settings["oracle_block_confirmation_delay"].(float64)
	if !ok {
		l.Log().Infoln("oracle_block_confirmation_delay not found in chain settings, using default value 4")
		ethBlockConfirmationDelay = 4
	} else {
		l.Log().Infoln("oracle_block_confirmation_delay found in chain settings, using value", ethBlockConfirmationDelay)
	}

	maxClaimsMsgPerBulk, ok := settings["oracle_max_claims_msg_per_bulk"].(float64)
	if !ok {
		l.Log().Infoln("oracle_max_claims_msg_per_bulk not found in chain settings, using default value 50")
		maxClaimsMsgPerBulk = 50
	} else {
		l.Log().Infoln("oracle_max_claims_msg_per_bulk found in chain settings, using value", maxClaimsMsgPerBulk)
	}

	l.HyperionState.OracleStatus = "getting CurrentValset on Helios"
	// check if validator is in the active set since claims will fail otherwise
	vs, err := l.GetHelios().CurrentValset(ctx, l.cfg.HyperionId)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to get active validator set on Helios")
		return err
	}

	l.HyperionState.OracleStatus = "getting LatestObservedHeight on Helios"
	latestObservedHeight, err := l.GetHelios().QueryGetLastObservedEthereumBlockHeight(ctx, l.cfg.HyperionId)
	if err != nil {
		return errors.Wrap(err, "failed to get latest valsets on Helios")
	}

	// state
	l.HyperionState.LastObservedHeight = latestObservedHeight.EthereumBlockHeight

	l.HyperionState.OracleStatus = "checking if first time sync is needed"
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
		l.HyperionState.OracleStatus = "validator not in active set, cannot make claims..."
		l.Log().WithFields(log.Fields{"latest_helios_block": vs.Height}).Warningln("validator not in active set, cannot make claims...")
		_, err := l.GetHelios().GetLatestBlockHeight(ctx)
		if err != nil {
			l.HyperionState.OracleStatus = "Connected node is down, cannot make claims..."
			l.Log().WithError(err).Errorln("Connected node is down, cannot make claims...")
			return err
		}
		//l.cfg.CosmosAddr.String()
		validator, err := l.GetHelios().GetValidator(ctx, l.cfg.ValidatorAddress.String())
		if err != nil {
			l.Log().WithError(err).Errorln("failed to get validator on " + l.cfg.ChainName)
			return err
		}
		if validator.Jailed {
			// todo try to unjail
			l.Log().WithFields(log.Fields{"latest_helios_block": vs.Height, "validator": validator.Description.Moniker}).Warningln("validator jailed, cannot make claims...")
			msg, err := l.GetHelios().SendUnjailMsg(ctx, l.cfg.ValidatorAddress.String())
			if err != nil {
				l.Log().WithError(err).Errorln("failed to send unjail message on " + l.cfg.ChainName)
				return err
			}
			err = l.GetHelios().SyncBroadcastMsgsSimulate(ctx, []cosmostypes.Msg{msg})
			if err != nil {
				VerifyTxError(ctx, err.Error(), l.Orchestrator)
				l.Log().WithError(err).Errorln("failed to simulate unjail message on " + l.cfg.ChainName)
				return err
			}
			_, err = l.global.SyncBroadcastMsgs(ctx, []cosmostypes.Msg{msg})
			if err != nil {
				l.Log().WithError(err).Errorln("failed to broadcast unjail message on " + l.cfg.ChainName)
				return err
			}
			return nil
		}
		msg, err := l.GetHelios().SendSetOrchestratorAddressesMsg(ctx, uint64(l.cfg.HyperionId), l.cfg.EthereumAddr.String())
		if err != nil {
			l.Log().WithError(err).Errorln("failed to send set orchestrator addresses message on " + l.cfg.ChainName)
			return err
		}
		err = l.GetHelios().SyncBroadcastMsgsSimulate(ctx, []cosmostypes.Msg{msg})
		if err != nil {
			VerifyTxError(ctx, err.Error(), l.Orchestrator)
			l.Log().WithError(err).Errorln("failed to simulate set orchestrator addresses message on " + l.cfg.ChainName)
			return err
		}
		_, err = l.global.SyncBroadcastMsgs(ctx, []cosmostypes.Msg{msg})
		if err != nil {
			l.Log().WithError(err).Errorln("failed to broadcast set orchestrator addresses message on " + l.cfg.ChainName)
			return err
		}
		return nil
	}

	l.HyperionState.OracleStatus = "getting latest ethereum height"
	latestHeight, err := l.getLatestEthHeight(ctx)
	if err != nil {
		l.Log().WithError(err).Errorln("failed to get latest " + l.cfg.ChainName + " height")
		return err
	}

	// state
	l.HyperionState.Height = latestHeight
	l.HyperionState.TargetHeight = latestHeight - uint64(ethBlockConfirmationDelay)

	targetHeight := latestHeight

	// not enough blocks on ethereum yet
	if targetHeight <= uint64(ethBlockConfirmationDelay) {
		l.Log().Debugln("not enough blocks on " + l.cfg.ChainName)
		return nil
	}

	// ensure that latest block has minimum confirmations
	targetHeight = targetHeight - uint64(ethBlockConfirmationDelay)
	if targetHeight <= l.lastObservedEthHeight {
		l.Log().Infoln("Synced", l.lastObservedEthHeight, "to", targetHeight)
		return nil
	}

	// One-shot nonce snapshot for this tick. The contract's state_lastEventNonce
	// is a global, monotonically increasing counter (incremented on every bridge
	// event); lastObservedEventNonce is what Helios has already processed.
	// Fetching them once here instead of once per chunk in syncToTargetHeight
	// removes ~2 RPC calls per 2000-block chunk.
	lastObservedEventNonce, err := l.GetHelios().QueryGetLastObservedEventNonce(ctx, l.cfg.HyperionId)
	if err != nil {
		l.Log().WithError(err).Errorln("failed to get last observed event nonce on " + l.cfg.ChainName)
		return err
	}

	lastEventNonceBig, err := l.ethereum.GetLastEventNonce(ctx)
	if err != nil {
		kind := rpcerrors.Classify(err)
		if rpcerrors.ShouldRotate(kind) || strings.Contains(err.Error(), "no contract code at given address") {
			l.Orchestrator.RotateRpc()
		}
		l.Log().WithError(err).WithField("kind", kind.String()).Errorln("failed to get last event nonce on " + l.cfg.ChainName)
		return err
	}
	lastEventNonce := lastEventNonceBig.Uint64()

	// Fast-path: Helios has already observed every event the contract ever
	// emitted, so there is provably nothing to claim between here and the target.
	// Advance the height pointer in a single jump instead of grinding chunk by
	// chunk — which previously cost ~2 RPC calls per 2000 blocks for nothing.
	if l.missedEventsBlockHeight == 0 && lastObservedEventNonce == lastEventNonce {
		l.lastObservedEthHeight = targetHeight
		// Keep the UI/state in sync. The chunk loop used to publish these every
		// iteration via syncToTargetHeight; since we skip the loop entirely we
		// must publish them here, otherwise GetHeight() freezes and the UI shows
		// a stale "out of sync" height even though scanning is fully caught up.
		l.Orchestrator.SetHeight(l.lastObservedEthHeight)
		l.Orchestrator.SetTargetHeight(latestHeight - uint64(ethBlockConfirmationDelay))
		l.Log().Infoln("nonces in sync (", lastObservedEventNonce, "), fast-forwarded observed height to", targetHeight)
		return nil
	}

	// Catch-up loop. We iterate up to 100 chunks per tick. Previously a single
	// failing chunk tore down the whole loop; now we classify the failure and
	// respond in kind so that transient RPC issues don't stall the sync.
	for i := 0; i < 100; i++ {
		if l.lastObservedEthHeight >= targetHeight {
			break
		}

		chunkEnd := l.lastObservedEthHeight + uint64(defaultBlocksToSearch)
		if chunkEnd > targetHeight {
			chunkEnd = targetHeight
		}

		l.HyperionState.OracleStatus = "scanning " + strconv.FormatUint(l.lastObservedEthHeight, 10) +
			" -> " + strconv.FormatUint(chunkEnd, 10) +
			" (" + strconv.FormatUint(targetHeight-l.lastObservedEthHeight, 10) + " left)"

		sentClaims, err := l.syncToTargetHeight(ctx, latestHeight, chunkEnd, uint64(ethBlockConfirmationDelay), int(maxClaimsMsgPerBulk), lastObservedEventNonce, lastEventNonce)
		if err == nil {
			if sentClaims {
				// Claims advanced Helios's observed nonce (and new events may have
				// landed on-chain meanwhile); refresh the snapshot so the next
				// chunk evaluates against current state, not stale nonces.
				if n, e := l.GetHelios().QueryGetLastObservedEventNonce(ctx, l.cfg.HyperionId); e == nil {
					lastObservedEventNonce = n
				}
				if n, e := l.ethereum.GetLastEventNonce(ctx); e == nil {
					lastEventNonce = n.Uint64()
				}
			}
			l.consecutiveSuccesses++
			// After a streak of clean chunks, tentatively widen the window.
			// Providers recover; we shouldn't stay shrunk forever.
			if l.consecutiveSuccesses >= blocksGrowAfterSuccesses && defaultBlocksToSearch < maxBlocksToSearch {
				newSize := defaultBlocksToSearch * blocksGrowFactor
				if newSize > maxBlocksToSearch {
					newSize = maxBlocksToSearch
				}
				defaultBlocksToSearch = newSize
				l.consecutiveSuccesses = 0
				if s, sErr := storage.GetChainSettings(l.cfg.ChainId); sErr == nil {
					s["oracle_eth_default_blocks_to_search"] = defaultBlocksToSearch
					_ = storage.SetChainSettings(l.cfg.ChainId, s)
				}
				l.Log().Infoln("scan window widened to", defaultBlocksToSearch)
			}
			continue
		}

		kind := rpcerrors.Classify(err)
		l.consecutiveSuccesses = 0

		switch kind {
		case rpcerrors.RangeTooLarge:
			// The provider rejected the range size itself; shrink and retry
			// the same chunk. This is the one case where the retry layer
			// can't help us — only a smaller window can.
			newSize := defaultBlocksToSearch / 2
			if newSize < minBlocksToSearch {
				newSize = minBlocksToSearch
			}
			if newSize == defaultBlocksToSearch {
				// Already at floor; rotating is the only remaining lever.
				l.Log().Warningln("range error at minimum window, rotating rpc")
				l.Orchestrator.RotateRpc()
				return nil
			}
			defaultBlocksToSearch = newSize
			if s, sErr := storage.GetChainSettings(l.cfg.ChainId); sErr == nil {
				s["oracle_eth_default_blocks_to_search"] = defaultBlocksToSearch
				_ = storage.SetChainSettings(l.cfg.ChainId, s)
			}
			l.Log().Infoln("scan window halved to", defaultBlocksToSearch, "(range_too_large)")
			// Retry the same chunk on the next loop iteration.
			continue

		case rpcerrors.RateLimit, rpcerrors.Timeout, rpcerrors.Network, rpcerrors.NodeIssue:
			// retry() already exhausted its attempts across multiple RPCs;
			// continuing would just burn through the remaining iterations
			// pointlessly. Bail out and let the next tick pick up where we
			// left off.
			l.Log().WithError(err).
				WithField("kind", kind.String()).
				Warningln("transient RPC error, deferring remaining chunks to next tick")
			return nil

		case rpcerrors.HistoryPruned:
			// All configured RPCs reject this block range as pruned. Two
			// paths: (a) user has opted into skipping — we advance past the
			// range and flag the gap in state; (b) default — bail with a
			// clear actionable message. Retrying the same chunk indefinitely
			// is never the right answer.
			if skipPrunedHistory {
				l.Log().WithField("from", l.lastObservedEthHeight).
					WithField("to", chunkEnd).
					Warningln("history pruned on all rpcs, skipping unreachable range (oracle_skip_pruned_history=true)")
				l.Orchestrator.HyperionState.ErrorStatus = "skipped pruned range " +
					strconv.FormatUint(l.lastObservedEthHeight, 10) + "-" +
					strconv.FormatUint(chunkEnd, 10)
				l.lastObservedEthHeight = chunkEnd
				continue
			}
			l.Orchestrator.HyperionState.ErrorStatus = "history pruned on all rpcs at block " +
				strconv.FormatUint(l.lastObservedEthHeight, 10) +
				" — add an archive rpc or set oracle_skip_pruned_history=true"
			l.Log().WithError(err).
				WithField("block", l.lastObservedEthHeight).
				Errorln("history pruned on all rpcs; add archive rpc or enable oracle_skip_pruned_history")
			return nil

		case rpcerrors.UnknownBlock:
			// Reorg or a node lagging behind; skip this chunk's scan and
			// advance — the target window will shift next tick anyway.
			l.Log().WithError(err).Warningln("unknown block, advancing past chunk")
			if chunkEnd > l.lastObservedEthHeight {
				l.lastObservedEthHeight = chunkEnd
			}
			continue

		default:
			// Missed-event or genuinely unknown — propagate so the outer loop
			// logs and the runOracle tick ends cleanly.
			return err
		}
	}

	// TODO delete normaly useless
	// if time.Since(l.lastResyncWithHelios) >= resyncInterval {
	// 	if err := l.autoResync(ctx); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

// syncToTargetHeight scans one chunk and broadcasts any new claims. The nonce
// pair (lastObservedEventNonce / lastEventNonce) is captured once per tick by
// the caller and threaded in, so this no longer issues its own nonce RPC calls.
// It returns true when it actually broadcast claims, so the caller knows to
// refresh the snapshot.
func (l *oracle) syncToTargetHeight(ctx context.Context, latestHeight uint64, targetHeight uint64, ethBlockConfirmationDelay uint64, maxClaimsMsgPerBulk int, lastObservedEventNonce uint64, lastEventNonce uint64) (bool, error) {

	l.Orchestrator.SetHeight(l.lastObservedEthHeight)
	l.Orchestrator.SetTargetHeight(latestHeight - uint64(ethBlockConfirmationDelay))

	if targetHeight-l.lastObservedEthHeight == 0 {
		l.Log().Infoln("No blocks to sync", "last_observed_eth_height", l.lastObservedEthHeight, "latest_height", latestHeight, "target_height", targetHeight)
		return false, nil
	}

	if l.missedEventsBlockHeight == 0 {
		if lastObservedEventNonce == lastEventNonce {
			l.Log().Infoln("lastObservedEventNonce is equal to lastEventNonce, no new events to process")
			l.lastObservedEthHeight = targetHeight
			// permit to reduce the number of calls to the ethereum rpc
			return false, nil
		} else { // special case to reduce the number of calls to the ethereum rpc cause we can miss events if we don't rewind few minutes
			// blockTimeOnTheChain is in milliseconds
			blockTimeOnTheChain := l.cfg.ChainParams.AverageCounterpartyBlockTime // 12000ms (12s)

			// Guard against misconfigured chain params: a zero block time
			// would panic on division below.
			if blockTimeOnTheChain == 0 {
				l.Log().Warningln("AverageCounterpartyBlockTime is 0, skipping rewind")
			} else {
				// compute number of blocks in 2 minutes
				const twoMinutesMs = 2 * 60 * 1000 // 120000 ms

				nbBlocksToRewind := twoMinutesMs / blockTimeOnTheChain

				l.Log().Infoln("rewinding the last observed height by ", nbBlocksToRewind, " blocks")
				// rewind the last observed height (clamped to 0)
				if nbBlocksToRewind > l.lastObservedEthHeight {
					l.lastObservedEthHeight = 0
				} else {
					l.lastObservedEthHeight = l.lastObservedEthHeight - nbBlocksToRewind
				}
			}
		}
	}

	events, err := l.getEthEvents(ctx, l.lastObservedEthHeight, targetHeight)
	if err != nil {
		l.Log().WithError(err).Errorln("failed to get events on " + l.cfg.ChainName)
		return false, err
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
		// l.Log().Infoln("NO EVENTS DETECTED 0", l.ethereum.GetLastUsedRpc())
		if l.logEnabled {
			l.Log().WithFields(log.Fields{"last_observed_event_nonce": lastObservedEventNonce, "eth_block_start": l.lastObservedEthHeight, "eth_block_end": targetHeight}).Infoln("oracle no new events on " + l.cfg.ChainName)
		}
		l.lastObservedEthHeight = targetHeight
		return false, nil
	}

	if len(newEvents) > 0 {
		l.Log().Infoln("SOME EVENTS DETECTED %d", len(newEvents))
	}

	if l.logEnabled {
		l.Log().WithFields(log.Fields{"event_helios_nonce": lastObservedEventNonce, "event_ethereum_nonce": newEvents[0].Nonce()}).Infoln("try oracle relay to helios")
	}

	if newEvents[0].Nonce() > lastObservedEventNonce+1 {
		// we missed an event
		lastObservedHeight, err := l.GetHelios().QueryGetLastObservedEthereumBlockHeight(ctx, l.cfg.HyperionId)
		if err != nil {
			return false, err
		}

		// if we missed an event, we need to rewind the last observed height by 5 minutes and continue from there
		if l.missedEventsBlockHeight == 0 {
			l.missedEventsBlockHeight = lastObservedHeight.EthereumBlockHeight
		} else {
			// blockTimeOnTheChain is in milliseconds
			blockTimeOnTheChain := l.cfg.ChainParams.AverageCounterpartyBlockTime // 12000ms (12s)

			if blockTimeOnTheChain == 0 {
				l.Log().Warningln("AverageCounterpartyBlockTime is 0, skipping missed-event rewind")
			} else {
				// compute number of blocks in 5 minutes
				const fiveMinutesMs = 5 * 60 * 1000 // 300000 ms

				nbBlocksToRewind := fiveMinutesMs / blockTimeOnTheChain

				if nbBlocksToRewind > l.missedEventsBlockHeight {
					l.missedEventsBlockHeight = 0
				} else {
					l.missedEventsBlockHeight = l.missedEventsBlockHeight - nbBlocksToRewind
				}
			}
		}
		l.lastObservedEthHeight = l.missedEventsBlockHeight
		// move back to the last observed event height
		l.Log().WithFields(log.Fields{"current_helios_nonce": lastObservedEventNonce, "wanted_nonce": lastObservedEventNonce + 1, "actual_ethereum_nonce": newEvents[0].Nonce()}).Infoln("orchestrator missed an " + l.cfg.ChainName + " event. Restarting block search from last observed claim...")
		return false, errors.New("missed an event")
	}

	l.missedEventsBlockHeight = 0

	if err := l.sendNewEventClaims(ctx, newEvents, maxClaimsMsgPerBulk); err != nil {
		log.Info("err: ", err)
		return false, err
	}

	if l.logEnabled {
		l.Log().WithFields(log.Fields{"claims": len(newEvents), "eth_block_start": l.lastObservedEthHeight, "eth_block_end": latestHeight}).Infoln("sent new event claims to Helios")
	}
	l.lastObservedEthHeight = targetHeight

	return true, nil
}

func (l *oracle) getEthEvents(ctx context.Context, startBlock, endBlock uint64) ([]event, error) {
	var events []event
	scanEthEventsFn := func() error {
		events = nil // clear previous result in case a retry occurred

		// Single eth_getLogs round-trip for all four event types instead of one
		// call per type (see network.GetAllHyperionEvents).
		bundle, err := l.ethereum.GetAllHyperionEvents(startBlock, endBlock)
		if err != nil {
			if strings.Contains(err.Error(), "limit exceeded") {
				return errors.Wrap(err, "failed to get Hyperion events - limit exceeded")
			}
			return errors.Wrap(err, "failed to get Hyperion events")
		}

		for _, e := range bundle.Deposits {
			ev := deposit(*e)
			events = append(events, &ev)
		}
		for _, e := range bundle.Withdrawals {
			ev := withdrawal(*e)
			events = append(events, &ev)
		}
		for _, e := range bundle.ValsetUpdates {
			ev := valsetUpdate(*e)
			events = append(events, &ev)
		}
		for _, e := range bundle.ERC20Deployments {
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
	attempt := 0
	fn := func() error {
		attempt++
		// Keep the UI status in sync while retries are in flight. Previously
		// this stayed stuck on "getting latest ethereum height" for the whole
		// retry budget, which made it look like the orchestrator was frozen.
		if attempt > 1 {
			l.HyperionState.OracleStatus = "getting latest height (retry #" +
				strconv.Itoa(attempt) + " on " + l.ethereum.GetRpc().Url + ")"
		}
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
		claim, err = l.GetHelios().LastClaimEventByAddr(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr)
		return
	}

	if err := l.retry(ctx, fn); err != nil {
		return nil, err
	}

	return claim, nil
}

func (l *oracle) sendNewEventClaims(ctx context.Context, events []event, maxClaimsMsgPerBulk int) error {
	sendEventsFn := func() error {
		// in case sending one of more claims fails, we reload the latest claimed nonce to filter processed events
		lastClaim, err := l.GetHelios().LastClaimEventByAddr(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr)
		if err != nil {
			return err
		}

		newEvents := filterEvents(events, lastClaim.EthereumEventNonce)
		if len(newEvents) == 0 {
			log.Infoln("No new events to send lastClaimNonce on Helios: ", lastClaim.EthereumEventNonce)
			return nil
		}

		var msgs []cosmostypes.Msg
		for _, event := range newEvents {
			msg, err := l.prepareSendEthEventClaim(ctx, event)
			if err != nil {
				return err
			}
			msgs = append(msgs, msg)

			if len(msgs) >= maxClaimsMsgPerBulk {
				log.Infoln("sending bulk of ", len(msgs), "claims messages")
				l.Orchestrator.HyperionState.OracleStatus = "sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"

				err = l.GetHelios().SyncBroadcastMsgsSimulate(ctx, msgs)
				if err != nil {
					VerifyTxError(ctx, err.Error(), l.Orchestrator)
					l.Log().WithError(err).Warningln("failed to simulate bulk of claims messages")
					return err
				}
				resp, err := l.global.SyncBroadcastMsgs(ctx, msgs)
				if err != nil {
					l.Orchestrator.HyperionState.OracleStatus = "error sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"
					log.Errorln("error sending bulk of ", len(msgs), "claims messages", err)
					return err
				}
				l.Orchestrator.HyperionState.OracleStatus = "bulk of " + strconv.Itoa(len(msgs)) + " claims messages sent"
				cost, err := l.GetHelios().GetTxCost(ctx, resp.TxHash)
				if err == nil {
					storage.UpdateFeesFile(big.NewInt(0), "", cost, resp.TxHash, uint64(resp.Height), uint64(42000), "CLAIM")
				}
				msgs = []cosmostypes.Msg{}
				time.Sleep(1100 * time.Millisecond)
			}
		}

		if len(msgs) > 0 {
			log.Infoln("sending bulk of ", len(msgs), "claims messages")
			l.Orchestrator.HyperionState.OracleStatus = "sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"
			err = l.GetHelios().SyncBroadcastMsgsSimulate(ctx, msgs)
			if err != nil {
				VerifyTxError(ctx, err.Error(), l.Orchestrator)
				l.Log().WithError(err).Warningln("failed to simulate bulk of claims messages")
				return err
			} else {
				l.Orchestrator.HyperionState.ErrorStatus = "okay"
			}
			resp, err := l.global.SyncBroadcastMsgs(ctx, msgs)
			if err != nil {
				l.Orchestrator.HyperionState.OracleStatus = "error sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"
				log.Errorln("error sending bulk of ", len(msgs), "claims messages", err)
				return err
			}
			cost, err := l.GetHelios().GetTxCost(ctx, resp.TxHash)
			if err == nil {
				l.Orchestrator.HyperionState.OracleStatus = "bulk of " + strconv.Itoa(len(msgs)) + " claims messages sent"
				storage.UpdateFeesFile(big.NewInt(0), "", cost, resp.TxHash, uint64(resp.Height), uint64(42000), "CLAIM")
			}
			l.Orchestrator.HyperionState.OracleStatus = "bulk of " + strconv.Itoa(len(msgs)) + " claims messages sent"
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
		height, err = l.getLastClaimBlockHeight(ctx, l.GetHelios())
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
	switch e := ev.(type) {
	case *deposit:
		ev := hyperionevents.HyperionSendToHeliosEvent(*e)
		l.HyperionState.InBridgedTxCount++
		return l.GetHelios().SendDepositClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
	case *valsetUpdate:
		l.HyperionState.ValsetUpdateCount++
		ev := hyperionevents.HyperionValsetUpdatedEvent(*e)
		return l.GetHelios().SendValsetClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
	case *withdrawal:
		ev := hyperionevents.HyperionTransactionBatchExecutedEvent(*e)
		return l.GetHelios().SendWithdrawalClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
	case *erc20Deployment:
		l.HyperionState.ERC20DeploymentCount++
		ev := hyperionevents.HyperionERC20DeployedEvent(*e)
		return l.GetHelios().SendERC20DeployedClaimMsg(ctx, l.cfg.HyperionId, &ev, rpc)
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
