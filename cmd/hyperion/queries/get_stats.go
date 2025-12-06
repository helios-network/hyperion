package queries

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
)

func GetStats(ctx context.Context, global *global.Global) (map[string]interface{}, error) {
	orchestrators := global.GetOrchestrators()
	stats := make(map[string]interface{})
	for chainId, orchestrator := range orchestrators {
		stats[fmt.Sprintf("%d", chainId)] = map[string]interface{}{
			"totalTxs":             orchestrator.HyperionState.TxCount,
			"batches":              orchestrator.HyperionState.BatchCount,
			"outBridgedTxCount":    orchestrator.HyperionState.OutBridgedTxCount,
			"inBridgedTxCount":     orchestrator.HyperionState.InBridgedTxCount,
			"valsetUpdateCount":    orchestrator.HyperionState.ValsetUpdateCount,
			"erc20DeploymentCount": orchestrator.HyperionState.ERC20DeploymentCount,
			"skippedRetriedCount":  orchestrator.HyperionState.SkippedRetriedCount,
			"externalDataCount":    orchestrator.HyperionState.ExternalDataCount,
			"height":               orchestrator.GetHeight(),
			"targetHeight":         orchestrator.GetTargetHeight(),
			"hyperionState":        orchestrator.HyperionState,
			"depositPaused":        orchestrator.HyperionState.IsDepositPaused,
			"withdrawalPaused":     orchestrator.HyperionState.IsWithdrawalPaused,
			"gasPrice":             orchestrator.HyperionState.GasPrice,
		}
	}

	return stats, nil
}
