package queries

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func StopHyperion(ctx context.Context, global *global.Global, chainId uint64) error {
	runner := global.GetRunner(chainId)
	if runner == nil {
		return fmt.Errorf("runner for chainId %d not found", chainId)
	}
	global.CancelRunner(chainId)
	return nil
}
