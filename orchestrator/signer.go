package orchestrator

import (
	"context"
	"strings"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

// runSigner simply signs off on any batches or validator sets provided by the validator
// since these are provided directly by a trusted Helios node they can simply be assumed to be
// valid and signed off on.
func (s *Orchestrator) runSigner(ctx context.Context, hyperionID gethcommon.Hash) error {
	signer := signer{
		Orchestrator: s,
		hyperionID:   hyperionID,
		logEnabled:   strings.Contains(s.cfg.EnabledLogs, "signer"),
	}

	s.logger.WithField("loop_duration", defaultLoopDur.String()).Debugln("starting Signer...")

	return loops.RunLoop(ctx, s.ethereum, defaultLoopDur, func() error {
		// if !s.isRegistered() {
		// 	signer.Log().Infoln("Orchestrator not registered, skipping...")
		// 	return nil
		// }
		err := signer.sign(ctx)
		return err
	})
}

type signer struct {
	*Orchestrator
	hyperionID gethcommon.Hash
	logEnabled bool
}

func (l *signer) Log() log.Logger {
	return l.logger.WithField("loop", "Signer")
}

func (l *signer) sign(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	l.Log().Debugln("signing")
	if err := l.signValidatorSets(ctx); err != nil {
		return err
	}
	l.Log().Debugln("signing validator sets done")

	for i := 0; i < 5; i++ {
		hasPushedABatch, err := l.signNewBatch(ctx)
		if err != nil {
			return err
		}
		if !hasPushedABatch {
			break
		}
	}
	l.Log().Debugln("signing new batch done")
	return nil
}

func (l *signer) signValidatorSets(ctx context.Context) error {
	var valsets []*hyperiontypes.Valset
	fn := func() error {
		valsets, _ = l.helios.OldestUnsignedValsets(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr)
		return nil
	}

	if err := l.retry(ctx, fn); err != nil {
		return errors.Wrap(err, "getting oldest unsigned valsets failed")
	}

	if len(valsets) == 0 {
		return nil
	}

	for _, vs := range valsets {
		if err := l.retry(ctx, func() error {
			l.Log().Infoln("signing valset", vs.Nonce)
			return l.helios.SendValsetConfirm(ctx, l.cfg.HyperionId, l.cfg.EthereumAddr, l.hyperionID, vs)
		}); err != nil {
			l.Log().Infoln("signing valset failed", err)
			return err
		}

		if l.logEnabled {
			l.Log().WithFields(log.Fields{"valset_nonce": vs.Nonce, "validators": len(vs.Members)}).Infoln("confirmed valset update on Helios")
		}
	}

	return nil
}

func (l *signer) signNewBatch(ctx context.Context) (bool, error) {
	var oldestUnsignedBatch *hyperiontypes.OutgoingTxBatch
	getBatchFn := func() error {
		tmpOldestUnsignedBatch, _ := l.helios.OldestUnsignedTransactionBatch(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr)
		if tmpOldestUnsignedBatch != nil && tmpOldestUnsignedBatch.HyperionId == l.cfg.HyperionId {
			oldestUnsignedBatch = tmpOldestUnsignedBatch
		}
		return nil
	}

	if err := l.retry(ctx, getBatchFn); err != nil {
		return false, err
	}

	if oldestUnsignedBatch == nil {
		if l.logEnabled {
			l.Log().Infoln("no token batch to confirm")
		}
		return false, nil
	}

	if err := l.retry(ctx, func() error {
		return l.helios.SendBatchConfirm(ctx,
			l.cfg.HyperionId,
			l.cfg.EthereumAddr,
			l.hyperionID,
			oldestUnsignedBatch,
		)
	}); err != nil {
		return false, err
	}

	l.Log().WithFields(log.Fields{"token_contract": oldestUnsignedBatch.TokenContract, "batch_nonce": oldestUnsignedBatch.BatchNonce, "txs": len(oldestUnsignedBatch.Transactions)}).Infoln("confirmed batch on Helios")

	return true, nil
}
