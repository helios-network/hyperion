package queries

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcs"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func RemoveRpcs(ctx context.Context, global *global.Global, chainId uint64, rpcsList string) error {
	rpcsList = strings.TrimSpace(rpcsList)
	if rpcsList == "" {
		return fmt.Errorf("rpcsList is empty")
	}
	rpcsToRemoveArray := strings.Split(rpcsList, ",")
	rpcsArray, _, err := storage.GetRpcsFromStorge(chainId)
	if err != nil {
		return err
	}
	newRpcsArray := make([]*rpcs.Rpc, 0)
	for _, rpc := range rpcsArray {
		if !slices.Contains(rpcsToRemoveArray, rpc.Url) {
			newRpcsArray = append(newRpcsArray, rpc)
		}
	}
	return storage.UpdateRpcsToStorge(chainId, newRpcsArray)
}
