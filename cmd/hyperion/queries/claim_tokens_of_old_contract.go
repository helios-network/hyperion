package queries

import (
	"context"
	"fmt"

	sdkmath "cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

func ClaimTokensOfOldContract(ctx context.Context, global *global.Global, hyperionId uint64, oldContract string, tokenContract string, amountInt int64) map[string]interface{} {

	amountInSdkMath := sdkmath.NewInt(1000000000000000000).Mul(sdkmath.NewInt(amountInt))

	targetNetworks, err := global.InitTargetNetworks(&hyperiontypes.CounterpartyChainParams{
		HyperionId:                hyperionId,
		BridgeChainId:             hyperionId,
		BridgeChainName:           "Hyperion",
		BridgeChainLogo:           "",
		BridgeCounterpartyAddress: oldContract,
	})
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	if len(targetNetworks) == 0 {
		return map[string]interface{}{
			"error": fmt.Errorf("no target networks found for chain %d", hyperionId),
		}
	}
	targetNetwork := targetNetworks[0]

	err = (*targetNetwork).SendClaimTokensOfOldContract(ctx, hyperionId, tokenContract, amountInSdkMath.BigInt(), (*targetNetwork).FromAddress(), (*targetNetwork).GetPersonalSignFn())
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	return map[string]interface{}{
		"success": true,
	}
}
