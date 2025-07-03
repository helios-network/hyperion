package hyperion

import (
	"context"
	"math/big"

	"github.com/Helios-Chain-Labs/metrics"
	"github.com/ethereum/go-ethereum/common"
)

func (s *hyperionContract) SendPreparedTx(ctx context.Context, txData []byte) (*common.Hash, *big.Int, error) {
	txHash, cost, err := s.SendTx(ctx, s.hyperionAddress, txData)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, big.NewInt(0), err
	}
	return &txHash, cost, nil
}

func (s *hyperionContract) SendPreparedTxSync(ctx context.Context, txData []byte) (*common.Hash, *big.Int, error) {
	txHash, cost, err := s.SendTxSync(ctx, s.hyperionAddress, txData)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, big.NewInt(0), err
	}
	return &txHash, cost, nil
}
