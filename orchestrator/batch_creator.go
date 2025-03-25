package orchestrator

import (
	"context"
	"os"
	"strconv"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/peggo/orchestrator/loops"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

func (s *Orchestrator) runBatchCreator(ctx context.Context, hyperionId uint64) (err error) {
	bc := batchCreator{Orchestrator: s}
	s.logger.WithField("loop_duration", defaultLoopDur.String()).Debugln("starting BatchCreator...")

	return loops.RunLoop(ctx, defaultLoopDur, func() error {
		return bc.requestTokenBatches(ctx, hyperionId)
	})
}

type batchCreator struct {
	*Orchestrator
}

func (l *batchCreator) Log() log.Logger {
	return l.logger.WithField("loop", "BatchCreator")
}

func (l *batchCreator) requestTokenBatches(ctx context.Context, hyperionId uint64) error {
	log.Infoln("requesting token batches")
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	fees, err := l.getUnbatchedTokenFees(ctx, hyperionId)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to get withdrawal fees")
		return nil
	}

	if len(fees) == 0 {
		l.Log().Infoln("no withdrawals to batch")
		return nil
	}

	for _, fee := range fees {
		l.requestTokenBatch(ctx, fee)
	}

	return nil
}

func (l *batchCreator) getUnbatchedTokenFees(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.BatchFees, error) {
	var fees []*hyperiontypes.BatchFees
	fn := func() (err error) {
		fees, err = l.helios.UnbatchedTokensWithFees(ctx, hyperionId)
		return
	}

	if err := l.retry(ctx, fn); err != nil {
		return nil, err
	}

	return fees, nil
}

func (l *batchCreator) requestTokenBatch(ctx context.Context, fee *hyperiontypes.BatchFees) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Load Failed .env: %v", err)
	}
	hyperionId, _ := strconv.ParseUint(os.Getenv("HYPERION_ID"), 10, 64)
	tokenAddress := gethcommon.HexToAddress(fee.Token)
	tokenDenom := l.getTokenDenom(tokenAddress)

	tokenDecimals, err := l.ethereum.TokenDecimals(ctx, tokenAddress)
	if err != nil {
		l.Log().WithError(err).Warningln("is token address valid?")
		return
	}

	if !l.checkMinBatchFee(fee, tokenAddress, tokenDecimals) {
		return
	}

	l.Log().WithFields(log.Fields{"token_denom": tokenDenom, "token_addr": tokenAddress.String()}).Infoln("requesting token batch on Helios")

	_ = l.helios.SendRequestBatch(ctx, hyperionId, tokenDenom)
}

func (l *batchCreator) getTokenDenom(tokenAddr gethcommon.Address) string {
	if cosmosDenom, ok := l.cfg.ERC20ContractMapping[tokenAddr]; ok {
		return cosmosDenom
	}

	return hyperiontypes.HyperionDenomString(tokenAddr)
}

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

	l.Log().WithFields(log.Fields{
		"token_addr": fee.Token,
		"total_fee":  totalFeeUSD.String() + "USD",
		"min_fee":    minFeeUSD.String() + "USD",
	}).Debugln("checking total batch fees")

	if totalFeeUSD.LessThan(minFeeUSD) {
		return false
	}

	return true
}
