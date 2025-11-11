package orchestrator

import (
	"context"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
)

const (
	defaultSkippedLoopDur = 60 * time.Second
)

func (s *Orchestrator) runSkipped(ctx context.Context) error {
	skipped := skipped{
		Orchestrator: s,
		logEnabled:   strings.Contains(s.cfg.EnabledLogs, "skipped"),
	}
	s.logger.WithField("loop_duration", defaultSkippedLoopDur.String()).Debugln("starting Skipped...")

	return loops.RunLoop(ctx, s.ethereum, defaultSkippedLoopDur, func() error {
		if s.HyperionState.SkippedStatus != "idle" {
			return nil
		}

		start := time.Now()
		s.HyperionState.SkippedStatus = "running"
		err := skipped.Run(ctx)
		s.HyperionState.SkippedStatus = "idle"
		s.HyperionState.SkippedLastExecutionFinishedTimestamp = uint64(time.Now().Unix())
		s.HyperionState.SkippedNextExecutionTimestamp = uint64(start.Add(defaultSkippedLoopDur).Unix())
		return err
	})
}

type skipped struct {
	*Orchestrator
	logEnabled bool
}

func (l *updater) Log() log.Logger {
	return l.logger.WithField("loop", "Updater")
}

func (l *skipped) Run(ctx context.Context) error {
	l.logger.Info("Running Skipped...")

	if l.Orchestrator.Oracle == nil {
		return errors.New("oracle is not initialized")
	}

	settings, err := storage.GetChainSettings(l.cfg.ChainId)
	if err != nil {
		return errors.Wrap(err, "failed to get chain settings")
	}

	maxClaimsMsgPerBulk, ok := settings["oracle_max_claims_msg_per_bulk"].(float64)
	if !ok {
		log.Infoln("oracle_max_claims_msg_per_bulk not found in chain settings, using default value 50")
		maxClaimsMsgPerBulk = 50
	} else {
		log.Infoln("oracle_max_claims_msg_per_bulk found in chain settings, using value", maxClaimsMsgPerBulk)
	}

	skippedNonces, err := l.Orchestrator.GetHelios().QueryGetAllSkippedTxs(ctx, l.cfg.ChainId)
	if err != nil {
		return errors.Wrap(err, "unable to get skipped nonces")
	}

	if len(skippedNonces) == 0 {
		log.Infoln("no skipped nonces found")
		time.Sleep(5 * time.Second)
		return nil
	}

	nonceSent := make(map[uint64]bool)

	for _, skippedNonce := range skippedNonces {

		if nonceSent[skippedNonce.Nonce] {
			continue
		}

		if skippedNonce.Nonce == 0 {
			continue
		}

		l.HyperionState.SkippedStatus = "getting events for nonce " + strconv.FormatUint(skippedNonce.Nonce, 10)

		events, err := l.Orchestrator.Oracle.getEthEvents(ctx, skippedNonce.StartHeight, skippedNonce.EndHeight, []uint64{skippedNonce.Nonce})
		if err != nil {
			log.WithError(err).Errorln("failed to get events on " + l.cfg.ChainName)
			return err
		}
		var eventsToSend []event
		for _, event := range events {
			for _, skippedNonceCheck := range skippedNonces {
				if event.Nonce() == skippedNonceCheck.Nonce { // check if event nonce is in skipped nonces
					eventsToSend = append(eventsToSend, event)
					nonceSent[event.Nonce()] = true
				}
			}
		}
		if len(eventsToSend) == 0 {
			log.Errorln("event not found for nonce " + strconv.FormatUint(skippedNonce.Nonce, 10))
			continue
		}

		l.HyperionState.SkippedStatus = "sending events for nonce " + strconv.FormatUint(skippedNonce.Nonce, 10)

		if err := l.sendNewEventClaimsWithoutFilter(ctx, eventsToSend, int(maxClaimsMsgPerBulk)); err != nil {
			log.Info("err: ", err)
			return err
		}

		log.Infoln("sent events for nonce " + strconv.FormatUint(skippedNonce.Nonce, 10))
	}

	return nil
}

func (l *skipped) sendNewEventClaimsWithoutFilter(ctx context.Context, newEvents []event, maxClaimsMsgPerBulk int) error {
	sendEventsFn := func() error {

		if len(newEvents) == 0 {
			log.Infoln("No new events to send")
			return nil
		}

		var msgs []cosmostypes.Msg
		for _, event := range newEvents {
			msg, err := l.Orchestrator.Oracle.prepareSendEthEventClaim(ctx, event)
			if err != nil {
				return err
			}
			msgs = append(msgs, msg)

			if len(msgs) >= maxClaimsMsgPerBulk {
				log.Infoln("sending bulk of ", len(msgs), "claims messages")
				l.Orchestrator.HyperionState.SkippedStatus = "sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"

				err = l.GetHelios().SyncBroadcastMsgsSimulate(ctx, msgs)
				if err != nil {
					VerifyTxError(ctx, err.Error(), l.Orchestrator)
					log.WithError(err).Warningln("failed to simulate bulk of claims messages")
					return err
				}
				resp, err := l.global.SyncBroadcastMsgs(ctx, msgs)
				if err != nil {
					l.Orchestrator.HyperionState.SkippedStatus = "error sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"
					log.Errorln("error sending bulk of ", len(msgs), "claims messages", err)
					return err
				}
				l.Orchestrator.HyperionState.SkippedRetriedCount += len(msgs)
				l.Orchestrator.HyperionState.SkippedStatus = "bulk of " + strconv.Itoa(len(msgs)) + " claims messages sent"
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
			l.Orchestrator.HyperionState.SkippedStatus = "sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"
			err := l.GetHelios().SyncBroadcastMsgsSimulate(ctx, msgs)
			if err != nil {
				VerifyTxError(ctx, err.Error(), l.Orchestrator)
				log.WithError(err).Warningln("failed to simulate bulk of claims messages")
				return err
			}
			resp, err := l.global.SyncBroadcastMsgs(ctx, msgs)
			if err != nil {
				l.Orchestrator.HyperionState.SkippedStatus = "error sending bulk of " + strconv.Itoa(len(msgs)) + " claims messages"
				log.Errorln("error sending bulk of ", len(msgs), "claims messages", err)
				return err
			}
			cost, err := l.GetHelios().GetTxCost(ctx, resp.TxHash)
			if err == nil {
				l.Orchestrator.HyperionState.SkippedStatus = "bulk of " + strconv.Itoa(len(msgs)) + " claims messages sent"
				storage.UpdateFeesFile(big.NewInt(0), "", cost, resp.TxHash, uint64(resp.Height), uint64(42000), "CLAIM")
			}
			l.Orchestrator.HyperionState.SkippedStatus = "bulk of " + strconv.Itoa(len(msgs)) + " claims messages sent"
			time.Sleep(1100 * time.Millisecond)
		}

		return nil
	}

	if err := l.retry(ctx, sendEventsFn); err != nil {
		return err
	}

	return nil
}
