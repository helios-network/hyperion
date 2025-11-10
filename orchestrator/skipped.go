package orchestrator

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
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
		if s.HyperionState.SkippedStatus == "running" {
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

		l.HyperionState.SkippedStatus = "getting events for nonce " + strconv.FormatUint(skippedNonce.Nonce, 10)

		events, err := l.Orchestrator.Oracle.getEthEvents(ctx, skippedNonce.StartHeight, skippedNonce.EndHeight)
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

		if err := l.Orchestrator.Oracle.sendNewEventClaimsWithoutFilter(ctx, eventsToSend, int(maxClaimsMsgPerBulk)); err != nil {
			log.Info("err: ", err)
			return err
		}

		log.Infoln("sent events for nonce " + strconv.FormatUint(skippedNonce.Nonce, 10))
	}

	return nil
}
