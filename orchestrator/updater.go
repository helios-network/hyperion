package orchestrator

import (
	"context"
	"strings"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	log "github.com/xlab/suplog"
)

const (
	defaultUpdaterLoopDur = 20 * time.Minute
)

func (s *Orchestrator) runUpdater(ctx context.Context) error {
	updater := updater{
		Orchestrator: s,
		logEnabled:   strings.Contains(s.cfg.EnabledLogs, "updater"),
	}
	s.logger.WithField("loop_duration", defaultUpdaterLoopDur.String()).Debugln("starting Updater...")

	return loops.RunLoop(ctx, s.ethereum, defaultUpdaterLoopDur, func() error {
		if s.HyperionState.UpdaterStatus == "running" {
			return nil
		}

		start := time.Now()
		s.HyperionState.UpdaterStatus = "running"
		err := updater.Update()
		s.HyperionState.UpdaterStatus = "idle"
		s.HyperionState.UpdaterLastExecutionFinishedTimestamp = uint64(time.Now().Unix())
		s.HyperionState.UpdaterNextExecutionTimestamp = uint64(start.Add(defaultUpdaterLoopDur).Unix())
		return err
	})
}

type updater struct {
	*Orchestrator
	logEnabled bool
}

func (l *updater) Log() log.Logger {
	return l.logger.WithField("loop", "Updater")
}

func (l *updater) Update() error {
	l.logger.Info("Updater updated")
	return nil
}
