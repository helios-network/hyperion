package orchestrator

import (
	"context"
	"strings"
	"time"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

func (s *Orchestrator) runBatchCreator(ctx context.Context) (err error) {
	bc := batchCreator{
		Orchestrator: s,
		logEnabled:   strings.Contains(s.cfg.EnabledLogs, "batch-creator"),
	}
	s.logger.WithField("loop_duration", defaultLoopDur.String()).Debugln("starting BatchCreator...")

	return loops.RunLoop(ctx, s.ethereum, defaultLoopDur, func() error {
		if s.HyperionState.BatchCreatorStatus == "running" {
			return nil
		}

		start := time.Now()
		s.HyperionState.BatchCreatorStatus = "running"
		err := bc.requestTokenBatches(ctx)
		s.HyperionState.BatchCreatorStatus = "idle"
		s.HyperionState.BatchCreatorLastExecutionFinishedTimestamp = uint64(time.Now().Unix())
		// elapsed := time.Since(start)
		s.HyperionState.BatchCreatorNextExecutionTimestamp = uint64(start.Add(defaultLoopDur).Unix())
		return err
	})
}

type batchCreator struct {
	*Orchestrator
	logEnabled bool
}

func (l *batchCreator) Log() log.Logger {
	return l.logger.WithField("loop", "BatchCreator")
}

func (l *batchCreator) requestTokenBatches(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	fees, err := l.getUnbatchedTokenFees(ctx)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to get withdrawal fees")
		return nil
	}

	if len(fees) == 0 {
		l.Log().Infoln("no withdrawals to batch")
		return nil
	}

	msgs := make([]cosmostypes.Msg, 0, len(fees))
	for _, fee := range fees {
		msg, err := l.requestTokenBatch(ctx, fee)
		if err != nil {
			l.Log().WithError(err).Warningln("failed to request token batch")
			return err
		}
		msgs = append(msgs, msg)
	}

	if len(msgs) > 0 { // bulk send messages
		paquetOfTenMsgs := make([]cosmostypes.Msg, 0)
		for _, msg := range msgs {
			paquetOfTenMsgs = append(paquetOfTenMsgs, msg)
			if len(paquetOfTenMsgs) == 10 {
				l.Log().Info("broadcasting token batches request ", "nb_msgs ", len(paquetOfTenMsgs))
				_, err := l.helios.SyncBroadcastMsgs(ctx, paquetOfTenMsgs)
				if err != nil {
					l.Log().WithError(err).Warningln("failed to broadcast token batches")
					return err
				}
				paquetOfTenMsgs = make([]cosmostypes.Msg, 0)
			}
		}
		if len(paquetOfTenMsgs) > 0 {
			l.Log().Info("broadcasting token batches request ", "nb_msgs ", len(paquetOfTenMsgs))
			_, err := l.helios.SyncBroadcastMsgs(ctx, paquetOfTenMsgs)
			if err != nil {
				l.Log().WithError(err).Warningln("failed to broadcast token batches")
				return err
			}
		}
	}

	return nil
}

func (l *batchCreator) getUnbatchedTokenFees(ctx context.Context) ([]*hyperiontypes.BatchFees, error) {
	var fees []*hyperiontypes.BatchFees
	fn := func() (err error) {
		fees, err = l.helios.UnbatchedTokensWithFees(ctx, l.cfg.HyperionId)
		return
	}

	if err := l.retry(ctx, fn); err != nil {
		return nil, err
	}

	return fees, nil
}

func (l *batchCreator) requestTokenBatch(ctx context.Context, fee *hyperiontypes.BatchFees) (cosmostypes.Msg, error) {
	tokenAddress := gethcommon.HexToAddress(fee.Token)
	tokenDenom, _, err := l.helios.QueryTokenAddressToDenom(ctx, l.cfg.HyperionId, tokenAddress)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to query token address to denom")
		return nil, err
	}

	tokenDecimals, err := l.ethereum.TokenDecimals(ctx, tokenAddress)
	if err != nil {
		l.Log().WithError(err).Warningln("is token address valid?")
		return nil, err
	}

	if _, ok := l.CacheSymbol[tokenAddress]; !ok {
		tokenSymbol, err := l.ethereum.TokenSymbol(ctx, tokenAddress)
		if err == nil {
			l.CacheSymbol[tokenAddress] = tokenSymbol
			l.Orchestrator.HyperionState.BatchCreatorStatus = "batching " + tokenSymbol
		} else {
			l.Orchestrator.HyperionState.BatchCreatorStatus = "batching " + tokenAddress.String()
		}
	} else {
		l.Orchestrator.HyperionState.BatchCreatorStatus = "batching " + l.CacheSymbol[tokenAddress]
	}

	if !l.checkMinBatchFee(fee, tokenAddress, tokenDecimals) {
		return nil, err
	}

	if l.logEnabled {
		l.Log().WithFields(log.Fields{"token_denom": tokenDenom, "token_addr": tokenAddress.String()}).Infoln("requesting token batch on Helios")
	}

	msg, err := l.helios.SendRequestBatchMsg(ctx, l.cfg.HyperionId, tokenDenom)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to send request batch msg")
		return nil, err
	}
	return msg, nil
}

// func (l *batchCreator) getTokenDenom(hyperionId uint64, tokenAddr gethcommon.Address) string {
// 	l.helios.QueryTokenAddressToDenom(ctx, hyperionId, tokenAddr)
// 	return hyperiontypes.HyperionDenomString(hyperionId, tokenAddr)
// }

func (l *batchCreator) checkMinBatchFee(fee *hyperiontypes.BatchFees, tokenAddress gethcommon.Address, tokenDecimals uint8) bool {
	if l.cfg.MinBatchFeeUSD == 0 {
		return true
	}

	tokenPriceUSDFloat, err := l.priceFeed.QueryUSDPrice(tokenAddress)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to query price feed", "token_addr", tokenAddress.String())
		return true
	}

	var (
		minFeeUSD     = decimal.NewFromFloat(l.cfg.MinBatchFeeUSD)
		tokenPriceUSD = decimal.NewFromFloat(tokenPriceUSDFloat)
		totalFeeUSD   = decimal.NewFromBigInt(fee.TotalFees.BigInt(), -1*int32(tokenDecimals)).Mul(tokenPriceUSD)
	)

	if l.logEnabled {
		l.Log().WithFields(log.Fields{
			"token_addr": fee.Token,
			"total_fee":  totalFeeUSD.String() + "USD",
			"min_fee":    minFeeUSD.String() + "USD",
		}).Debugln("checking total batch fees")
	}

	return !totalFeeUSD.LessThan(minFeeUSD)
}
