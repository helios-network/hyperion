package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func GetListTokens(ctx context.Context, global *global.Global, chainId uint64, page uint64, size uint64) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	// _, err := network.GetCounterpartyChainParamsByChainId(ctx, chainId)
	// if err != nil {
	// 	return nil, err
	// }

	if page == 0 {
		page = 1
	}
	if size == 0 {
		size = 100
	}

	if size > 100 {
		size = 100
	}

	tokens, total, err := network.QueryGetListTokens(ctx, chainId, page, size)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"tokens": tokens,
		"total":  total,
	}, nil
}
