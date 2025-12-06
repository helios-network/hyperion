package queries

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/utils"
)

func GetNetworkGasPrice(ctx context.Context, global *global.Global, chainId uint64) (string, error) {
	gasPrice := global.GetGasPriceBigInt(chainId)
	if gasPrice == nil {
		return "0", fmt.Errorf("gas price is nil")
	}
	fmt.Println("gasPrice", gasPrice.String())

	// BSC EIP1559 is capped at 1gwei minimum ...
	return utils.FormatBigStringToFloat64(gasPrice.String(), 18), nil
}
