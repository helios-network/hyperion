package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

func GetListRpcs(ctx context.Context, global *global.Global, chainId uint64) (map[string]interface{}, error) {
	orchestrator := global.GetOrchestrator(chainId)
	rpcs := make([]*hyperiontypes.Rpc, 0)
	if orchestrator != nil {
		rpcs = orchestrator.GetEthereum().GetRpcs()
	}

	staticRpcs := make([]string, 0)
	staticRpcsList, err := storage.GetStaticRpcs(chainId)
	if err == nil {
		staticRpcs = staticRpcsList
	}

	return map[string]interface{}{
		"rpcs":       rpcs,
		"staticRpcs": staticRpcs,
	}, nil
}
