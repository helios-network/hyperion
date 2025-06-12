package queries

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
)

func AddStaticRpcs(ctx context.Context, global *global.Global, chainId uint64, rpcs string) error {
	rpcs = strings.TrimSpace(rpcs)
	if rpcs == "" {
		return fmt.Errorf("rpcs is empty")
	}
	rpcsArray := make([]string, 0)
	for _, rpc := range strings.Split(rpcs, ",") {
		rpc = strings.TrimSpace(rpc)
		if rpc == "" {
			continue
		}
		if !strings.HasPrefix(rpc, "http://") && !strings.HasPrefix(rpc, "https://") {
			return fmt.Errorf("rpc %s is not a valid URL", rpc)
		}
		rpcsArray = append(rpcsArray, rpc)
		if len(rpcsArray) > 100 {
			return fmt.Errorf("too many rpcs, max is 100")
		}
	}
	if len(rpcsArray) == 0 {
		return fmt.Errorf("no rpcs provided")
	}
	existingRpcs, err := storage.GetStaticRpcs(chainId)
	if err != nil {
		return err
	}
	for _, rpc := range existingRpcs {
		if slices.Contains(rpcsArray, rpc) {
			return fmt.Errorf("rpc %s already exists", rpc)
		}
	}
	return storage.StoreStaticRpcs(chainId, slices.Concat(existingRpcs, rpcsArray))
}
