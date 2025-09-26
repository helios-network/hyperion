package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

type Rpc struct {
	Url       string `json:"url"`
	IsPrimary bool   `json:"is_primary"`
}

func GetListRpcs(ctx context.Context, global *global.Global, chainId uint64) ([]Rpc, error) {
	orchestrator := global.GetOrchestrator(chainId)
	rpcs := make([]*hyperiontypes.Rpc, 0)
	if orchestrator != nil {
		rpcs = orchestrator.GetEthereum().GetRpcs()
	}

	staticRpcs := make([]storage.Rpc, 0)
	staticRpcsList, err := storage.GetStaticRpcs(chainId)
	if err == nil {
		staticRpcs = staticRpcsList
	}

	allRpcs := make([]Rpc, 0)
	for _, rpc := range rpcs {
		for _, staticRpc := range staticRpcs {
			if staticRpc.Url == rpc.Url && staticRpc.IsPrimary {
				allRpcs = append(allRpcs, Rpc{
					Url:       rpc.Url,
					IsPrimary: staticRpc.IsPrimary,
				})
				continue
			}
		}
		allRpcs = append(allRpcs, Rpc{
			Url:       rpc.Url,
			IsPrimary: false,
		})
	}

	return allRpcs, nil
}
