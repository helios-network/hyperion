package queries

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func SetPrimaryRpc(ctx context.Context, global *global.Global, chainId uint64, rpcUrl string) error {
	if rpcUrl == "" {
		return fmt.Errorf("rpcUrl is empty")
	}

	existingRpcs, _, err := storage.GetRpcsFromStorge(chainId)
	if err != nil {
		return err
	}

	found := false
	for _, rpc := range existingRpcs {
		if rpc.Url == rpcUrl {
			rpc.IsPrimary = true
			found = true
		} else {
			rpc.IsPrimary = false
		}
	}

	if !found {
		return fmt.Errorf("rpc %s not found for chain %d", rpcUrl, chainId)
	}

	return storage.UpdateRpcsToStorge(chainId, existingRpcs)
}
