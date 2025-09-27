package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcs"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func GetListRpcs(ctx context.Context, global *global.Global, chainId uint64) ([]*rpcs.Rpc, error) {

	// later get used rpcs and stats
	// orchestrator := global.GetOrchestrator(chainId)

	rpcs, _, err := storage.GetRpcsFromStorge(chainId)
	if err != nil {
		return nil, err
	}

	return rpcs, nil
}
