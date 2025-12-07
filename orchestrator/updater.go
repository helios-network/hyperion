package orchestrator

import (
	"context"
	"strings"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/utils"
	"github.com/pkg/errors"
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
		err := updater.Update(ctx)
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

func (l *updater) Update(ctx context.Context) error {
	l.logger.Info("Updating Updater...")
	l.HyperionState.UpdaterStatus = "updating chain params"
	// update params of cfg
	counterpartyChainParams, err := l.GetHelios().GetCounterpartyChainParamsByChainId(ctx, l.cfg.ChainId)
	if err != nil {
		return errors.Wrap(err, "unable to get counterparty chain params")
	}
	l.cfg.ChainParams = counterpartyChainParams

	l.HyperionState.UpdaterStatus = "updating deposit pause status"
	// update is deposit paused
	isDepositPaused, err := l.ethereum.IsDepositPaused(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to get deposit pause status")
	}
	l.HyperionState.IsDepositPaused = isDepositPaused

	l.HyperionState.UpdaterStatus = "updating withdrawal pause status"
	// update is withdrawal paused
	l.HyperionState.IsWithdrawalPaused = l.cfg.ChainParams.Paused

	// update native balance
	l.HyperionState.UpdaterStatus = "updating native balance"
	err = l.Orchestrator.UpdateNativeBalance(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to update native balance")
	}

	l.HyperionState.UpdaterStatus = "updating gas price"
	gasPrice, err := l.ethereum.GetGasPrice(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to get gas price")
	}
	l.HyperionState.GasPrice = utils.FormatBigStringToFloat64(gasPrice.String(), 9) + " gwei"
	l.HyperionState.UpdaterStatus = "idle"

	l.logger.Info("Updater updated")
	return nil
}
