package queries

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func GetListRpcs(ctx context.Context, global *global.Global, chainId uint64) (map[string]interface{}, error) {
	orchestrator := global.GetOrchestrator(chainId)
	if orchestrator == nil {
		return nil, fmt.Errorf("orchestrator not found for chainId %d", chainId)
	}

	rpcs := orchestrator.GetEthereum().GetRpcs()

	return map[string]interface{}{
		"rpcs": rpcs,
	}, nil
}
