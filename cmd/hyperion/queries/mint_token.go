package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MintToken(ctx context.Context, global *global.Global, chainId uint64, tokenAddress string, decimals uint64, amount float64, receiverAddress string) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()

	amountMath, err := utils.FormatAmount(amount, decimals)
	if err != nil {
		return nil, err
	}
	msg, err := network.MintTokensMsg(ctx, chainId, tokenAddress, amountMath, receiverAddress)
	if err != nil {
		return nil, err
	}
	err = network.SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}
	resp, err := global.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"success": true,
		"tx_hash": resp.TxHash,
	}, nil
}
