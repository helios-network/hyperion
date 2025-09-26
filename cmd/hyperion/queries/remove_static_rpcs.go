package queries

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func RemoveStaticRpcs(ctx context.Context, global *global.Global, chainId uint64, rpcs string) error {
	rpcs = strings.TrimSpace(rpcs)
	if rpcs == "" {
		return fmt.Errorf("rpcs is empty")
	}
	rpcsToRemoveArray := strings.Split(rpcs, ",")
	rpcsArray, err := storage.GetStaticRpcs(chainId)
	if err != nil {
		return err
	}
	newRpcsArray := make([]storage.Rpc, 0)
	for _, rpc := range rpcsArray {
		if !slices.Contains(rpcsToRemoveArray, rpc.Url) {
			newRpcsArray = append(newRpcsArray, rpc)
		}
	}
	return storage.StoreStaticRpcs(chainId, newRpcsArray)
}
