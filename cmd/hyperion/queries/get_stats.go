package queries

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func GetStats(ctx context.Context, global *global.Global) (map[string]interface{}, error) {
	// network := *global.GetHeliosNetwork()

	orchestrators := global.GetOrchestrators()
	stats := make(map[string]interface{})
	for chainId, orchestrator := range orchestrators {
		// batches, err := network.QueryGetListOutgoingTxs(ctx, chainId)
		// if err != nil {
		// 	return nil, err
		// }
		// txInBatches, txUnbatched, err := network.QueryGetAllPendingSendToChain(ctx, chainId)
		// if err != nil {
		// 	return nil, err
		// }

		stats[fmt.Sprintf("%d", chainId)] = map[string]interface{}{
			"totalTxs":          orchestrator.HyperionState.TxCount,
			"batches":           orchestrator.HyperionState.BatchCount,
			"outBridgedTxCount": orchestrator.HyperionState.OutBridgedTxCount,
			"inBridgedTxCount":  orchestrator.HyperionState.InBridgedTxCount,
			"height":            orchestrator.GetHeight(),
			"targetHeight":      orchestrator.GetTargetHeight(),
			"hyperionState":     orchestrator.HyperionState,
		}
	}

	return stats, nil
}
