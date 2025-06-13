package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func GetListOutgoingTxs(ctx context.Context, global *global.Global, chainId uint64) (map[string]int, error) {
	network := *global.GetHeliosNetwork()

	batches, err := network.QueryGetListOutgoingTxs(ctx, chainId)
	if err != nil {
		return nil, err
	}

	count := 0
	for _, batch := range batches {
		for _, tx := range batch.Transactions {
			if tx != nil {
				count++
			}
		}
	}

	return map[string]int{
		"total":   count,
		"batches": len(batches),
	}, nil
}
