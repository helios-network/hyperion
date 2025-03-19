package orchestrator

import (
	"context"
	"os"
	"strconv"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/peggo/orchestrator/loops"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

// runSigner simply signs off on any batches or validator sets provided by the validator
// since these are provided directly by a trusted Helios node they can simply be assumed to be
// valid and signed off on.
func (s *Orchestrator) runSigner(ctx context.Context, hyperionID gethcommon.Hash) error {
	signer := signer{
		Orchestrator: s,
		hyperionID:   hyperionID,
	}

	s.logger.WithField("loop_duration", defaultLoopDur.String()).Debugln("starting Signer...")

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		return signer.sign(ctx)
	})
}

type signer struct {
	*Orchestrator
	hyperionID gethcommon.Hash
}

func (l *signer) Log() log.Logger {
	return l.logger.WithField("loop", "Signer")
}

func (l *signer) sign(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	// if err := l.signValidatorSets(ctx); err != nil {
	// 	return err
	// }

	if err := l.signNewBatch(ctx); err != nil {
		return err
	}

	return nil
}

func (l *signer) signValidatorSets(ctx context.Context) error {
	var valsets []*hyperiontypes.Valset
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Load Failed .env: %v", err)
	}
	hyperionId, _ := strconv.ParseUint(os.Getenv("HYPERION_ID"), 10, 64)
	fn := func() error {
		valsets, _ = l.helios.OldestUnsignedValsets(ctx, l.cfg.CosmosAddr)
		return nil
	}

	if err := l.retry(ctx, fn); err != nil {
		return err
	}

	if len(valsets) == 0 {
		l.Log().Infoln("no validator set to confirm")
		return nil
	}

	for _, vs := range valsets {
		if err := l.retry(ctx, func() error {
			return l.helios.SendValsetConfirm(ctx, hyperionId, l.cfg.EthereumAddr, l.hyperionID, vs)
		}); err != nil {
			return err
		}

		l.Log().WithFields(log.Fields{"valset_nonce": vs.Nonce, "validators": len(vs.Members)}).Infoln("confirmed valset update on Helios")
	}

	return nil
}

func (l *signer) signNewBatch(ctx context.Context) error {
	var oldestUnsignedBatch *hyperiontypes.OutgoingTxBatch
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Load Failed .env: %v", err)
	}
	hyperionId, _ := strconv.ParseUint(os.Getenv("HYPERION_ID"), 10, 64)
	getBatchFn := func() error {
		tmpOldestUnsignedBatch, _ := l.helios.OldestUnsignedTransactionBatch(ctx, l.cfg.CosmosAddr, hyperionId)
		if tmpOldestUnsignedBatch != nil && tmpOldestUnsignedBatch.HyperionId == hyperionId {
			oldestUnsignedBatch = tmpOldestUnsignedBatch
		}
		return nil
	}

	if err := l.retry(ctx, getBatchFn); err != nil {
		return err
	}

	log.Info("oldestUnsignedBatch: ", oldestUnsignedBatch)

	if oldestUnsignedBatch == nil {
		l.Log().Infoln("no token batch to confirm")
		return nil
	}

	if err := l.retry(ctx, func() error {
		return l.helios.SendBatchConfirm(ctx,
			hyperionId,
			l.cfg.EthereumAddr,
			l.hyperionID,
			oldestUnsignedBatch,
		)
	}); err != nil {
		return err
	}

	l.Log().WithFields(log.Fields{"token_contract": oldestUnsignedBatch.TokenContract, "batch_nonce": oldestUnsignedBatch.BatchNonce, "txs": len(oldestUnsignedBatch.Transactions)}).Infoln("confirmed batch on Helios")

	return nil
}
