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
		return updater.Update()
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
	l.UpdateRpcs()
	l.logger.Info("Updater updated")
	return nil
}
