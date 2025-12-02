package queries

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func MintToken(ctx context.Context, global *global.Global, chainId uint64, tokenAddress string, decimals uint64, amount float64, receiverAddress string) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()

	dec, err := sdkmath.LegacyNewDecFromStr(fmt.Sprintf("%g", amount))
	if err != nil {
		return nil, err
	}

	tenPow := sdkmath.LegacyNewDec(1)
	for i := 0; i < int(decimals); i++ {
		tenPow = tenPow.MulInt64(10)
	}

	amountDec := dec.Mul(tenPow)
	amountMath := amountDec.TruncateInt()
	if amountMath.IsNegative() {
		return nil, fmt.Errorf("amount is negative")
	}
	msg, err := network.MintTokensMsg(ctx, chainId, tokenAddress, amountMath, receiverAddress)
	if err != nil {
		return nil, err
	}
	fmt.Println("msg: ", msg)
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
