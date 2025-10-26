package orchestrator

import (
	"context"
	"slices"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

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

	// Use a custom loop that waits for completion before starting next iteration
	ticker := time.NewTicker(defaultLoopDur)
	defer ticker.Stop()

	// Run first iteration immediately
	if err := signer.sign(ctx); err != nil {
		s.logger.WithError(err).Errorln("signer function returned an error")
	}

	for {
		select {
		case <-ticker.C:
			if s.HyperionState.SignerStatus == "running" {
				continue
			}

			start := time.Now()
			s.HyperionState.SignerStatus = "running"
			if err := signer.sign(ctx); err != nil {
				s.logger.WithError(err).Errorln("signer function returned an error")
			}
			s.HyperionState.SignerStatus = "idle"
			s.HyperionState.SignerLastExecutionFinishedTimestamp = uint64(time.Now().Unix())
			s.HyperionState.SignerNextExecutionTimestamp = uint64(start.Add(defaultLoopDur).Unix())
		case <-ctx.Done():
			return nil
		}
	}
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

	noncesPushed := []uint64{}
	for i := 0; i < 50; i++ {
		hasPushedABatch, noncePushed, err := l.signNewBatch(ctx, noncesPushed)
		if err != nil {
			return err
		}
		noncesPushed = append(noncesPushed, noncePushed)
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
			return l.helios.SendValsetConfirm(ctx, l.cfg.HyperionId, l.cfg.EthereumAddr, l.hyperionID, l.ethereum.GetPersonalSignFn(), vs)
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

func (l *signer) signNewBatch(ctx context.Context, noncesPushed []uint64) (bool, uint64, error) {
	var oldestUnsignedBatch *hyperiontypes.OutgoingTxBatch
	getBatchFn := func() error {
		tmpOldestUnsignedBatch, _ := l.helios.OldestUnsignedTransactionBatch(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr)
		if tmpOldestUnsignedBatch != nil && tmpOldestUnsignedBatch.HyperionId == l.cfg.HyperionId && !slices.Contains(noncesPushed, tmpOldestUnsignedBatch.BatchNonce) {
			oldestUnsignedBatch = tmpOldestUnsignedBatch
		}
		return nil
	}

	if err := l.retry(ctx, getBatchFn); err != nil {
		return false, 0, err
	}

	if oldestUnsignedBatch == nil {
		if l.logEnabled {
			l.Log().Infoln("no token batch to confirm")
		}
		return false, 0, nil
	}

	symbol, ok := l.Orchestrator.CacheSymbol[gethcommon.HexToAddress(oldestUnsignedBatch.TokenContract)]
	if !ok {
		symbol = oldestUnsignedBatch.TokenContract
	}

	l.Orchestrator.HyperionState.SignerStatus = "signing batch " + strconv.Itoa(int(oldestUnsignedBatch.BatchNonce)) + " " + symbol

	msg, err := l.helios.SendBatchConfirmMsg(ctx, l.cfg.HyperionId, l.cfg.EthereumAddr, l.hyperionID, l.ethereum.GetPersonalSignFn(), oldestUnsignedBatch)
	if err != nil {
		return false, 0, errors.Wrap(err, "failed to send batch confirm message")
	}

	err = l.helios.SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
	if err != nil {
		l.Log().WithError(err).Warningln("failed to simulate batch confirm message")
		return false, 0, err
	}

	_, err = l.helios.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return false, 0, errors.Wrap(err, "failed to broadcast batch confirm message")
	}

	// err := l.helios.SendBatchConfirmSync(ctx,
	// 	l.cfg.HyperionId,
	// 	l.cfg.EthereumAddr,
	// 	l.hyperionID,
	// 	l.ethereum.GetPersonalSignFn(),
	// 	oldestUnsignedBatch,
	// )

	if err != nil {
		l.Orchestrator.HyperionState.SignerStatus = "error signing batch " + strconv.Itoa(int(oldestUnsignedBatch.BatchNonce)) + " " + symbol
		return false, 0, err
	}

	l.Orchestrator.HyperionState.SignerStatus = "batch " + strconv.Itoa(int(oldestUnsignedBatch.BatchNonce)) + " " + symbol + " signed"

	l.Log().WithFields(log.Fields{"token_contract": oldestUnsignedBatch.TokenContract, "batch_nonce": oldestUnsignedBatch.BatchNonce, "txs": len(oldestUnsignedBatch.Transactions)}).Infoln("confirmed batch on Helios")

	return true, oldestUnsignedBatch.BatchNonce, nil
}
