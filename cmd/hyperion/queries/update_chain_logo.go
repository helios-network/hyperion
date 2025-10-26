package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func UpdateChainLogo(ctx context.Context, global *global.Global, chainId uint64, logo string) map[string]interface{} {
	network := *global.GetHeliosNetwork()
	msg, err := network.UpdateChainLogoMsg(ctx, chainId, logo)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	resp, err := network.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	return map[string]interface{}{
		"tx_hash": resp.TxHash,
	}
}
