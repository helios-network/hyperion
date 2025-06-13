package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func CancelAllPendingOutTx(ctx context.Context, global *global.Global, chainId uint64) error {
	network := *global.GetHeliosNetwork()

	// cancel all pending outgoing txs
	err := network.SendCancelAllPendingOutTx(ctx, chainId)
	if err != nil {
		return err
	}
	return nil
}
