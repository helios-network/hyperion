package queries

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MintToken(ctx context.Context, global *global.Global, chainId uint64, tokenAddress string, amount float64, receiverAddress string) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()

	amountMath := sdkmath.NewInt(int64(amount * 1000000000000000000)) // 10^18
	if amountMath.IsNegative() {
		return nil, fmt.Errorf("amount is negative")
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
