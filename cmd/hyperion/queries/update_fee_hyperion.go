package queries

import (
	"context"
	"fmt"
	"slices"

	sdkmath "cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/hyperion/cmd/hyperion/queries/utils"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func UpdateFeeHyperion(ctx context.Context, global *global.Global, minTxFeeHLS float64, minBatchFeeHLS float64, chainId uint64) error {
	network := *global.GetHeliosNetwork()

	registeredNetworks, _ := helios.GetListOfNetworksWhereRegistered(*global.GetHeliosNetwork(), global.GetAddress())
	if !slices.Contains(registeredNetworks, chainId) {
		return nil // no need to update
	}

	minTxFeeHLSMath := sdkmath.NewInt(0)
	minBatchFeeHLSMath := sdkmath.NewInt(0)

	if minTxFeeHLS != 0.0 {
		minTxFeeHLSMathv, err := utils.FormatAmount(minTxFeeHLS, 18)
		if err != nil {
			fmt.Println("Error formatting min tx fee hls: ", err)
			return err
		}
		minTxFeeHLSMath = minTxFeeHLSMathv
	}
	if minBatchFeeHLS != 0.0 {
		minBatchFeeHLSMathv, err := utils.FormatAmount(minBatchFeeHLS, 18)
		if err != nil {
			fmt.Println("Error formatting min batch fee hls: ", err)
			return err
		}
		minBatchFeeHLSMath = minBatchFeeHLSMathv
	}

	// check if different of the current min tx fee and min batch fee
	validatorHyperionData, err := network.GetValidatorHyperionData(ctx, global.GetAddress())
	if err != nil {
		return nil // no need to update
	}

	for _, orchestratorHyperionData := range validatorHyperionData.OrchestratorHyperionData {
		if orchestratorHyperionData.HyperionId == chainId {
			if orchestratorHyperionData.MinimumTxFee.String() == minTxFeeHLSMath.String() && orchestratorHyperionData.MinimumBatchFee.String() == minBatchFeeHLSMath.String() {
				fmt.Println("No need to update min tx fee and min batch fee for chain ", chainId)
				return nil // no need to update
			}
			break
		}
	}

	msg, err := network.SendUpdateOrchestratorAddressesFeeMsg(ctx, chainId, minTxFeeHLSMath, minBatchFeeHLSMath)
	if err != nil {
		return err
	}
	err = network.SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
	if err != nil {
		return err
	}
	_, err = global.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return err
	}
	return nil
}
