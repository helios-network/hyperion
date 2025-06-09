package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func GetListTransactions(ctx context.Context, global *global.Global, page int, size int) (map[string]interface{}, error) {
	fees, err := storage.GetFeesFile()
	if err != nil {
		return map[string]interface{}{
			"transactions": []map[string]interface{}{},
			"total":        0,
		}, nil
	}

	total := len(fees)

	// get page and size
	start := (page - 1) * size
	end := start + size
	if end > len(fees) {
		end = len(fees)
	}
	fees = fees[start:end]

	return map[string]interface{}{
		"transactions": fees,
		"total":        total,
	}, nil
}
