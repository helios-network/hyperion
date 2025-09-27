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

func AddRpcs(ctx context.Context, global *global.Global, chainId uint64, rpcsList string, isPrimary bool) error {
	rpcsList = strings.TrimSpace(rpcsList)
	if rpcsList == "" {
		return fmt.Errorf("rpcsList is empty")
	}
	rpcsArray := make([]*rpcs.Rpc, 0)
	for _, rpc := range strings.Split(rpcsList, ",") {
		rpc = strings.TrimSpace(rpc)
		if rpc == "" {
			continue
		}
		if !strings.HasPrefix(rpc, "http://") && !strings.HasPrefix(rpc, "https://") {
			return fmt.Errorf("rpc %s is not a valid URL", rpc)
		}
		rpcsArray = append(rpcsArray, &rpcs.Rpc{
			Url:       rpc,
			IsPrimary: isPrimary,
		})
		if len(rpcsArray) > 100 {
			return fmt.Errorf("too many rpcs, max is 100")
		}
	}
	if len(rpcsArray) == 0 {
		return fmt.Errorf("no rpcs provided")
	}
	existingRpcs, _, err := storage.GetRpcsFromStorge(chainId)
	if err != nil {
		return err
	}
	for _, rpc := range existingRpcs {
		for _, rpcArray := range rpcsArray {
			if rpcArray.Url == rpc.Url {
				return fmt.Errorf("rpc %s already exists", rpc.Url)
			}
		}
	}

	finalRpcs := slices.Concat(existingRpcs, rpcsArray)
	// check if has primary rpc
	hasPrimary := false
	for _, rpc := range finalRpcs {
		if rpc.IsPrimary {
			hasPrimary = true
			break
		}
	}
	if !hasPrimary { // set the first rpc as primary
		finalRpcs[0].IsPrimary = true
	}
	finalRpcs = storage.OrderRpcsByPrimaryFirst(finalRpcs)
	return storage.UpdateRpcsToStorge(chainId, finalRpcs)
}
