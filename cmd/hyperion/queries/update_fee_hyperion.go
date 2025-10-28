package queries

import (
	"context"
	"fmt"
	"slices"

	sdk "github.com/cosmos/cosmos-sdk/types"

	sdkmath "cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/helios"
)

func UpdateFeeHyperion(ctx context.Context, global *global.Global, minTxFeeHLS float64, minBatchFeeHLS float64, chainId uint64) error {
	network := *global.GetHeliosNetwork()

	registeredNetworks, _ := helios.GetListOfNetworksWhereRegistered(*global.GetHeliosNetwork(), global.GetAddress())
	if slices.Contains(registeredNetworks, chainId) {
		return fmt.Errorf("hyperion already registered for chain %d", chainId)
	}

	minTxFeeHLSMath := sdkmath.NewInt(0)
	minBatchFeeHLSMath := sdkmath.NewInt(0)

	if minTxFeeHLS != 0.0 {
		minTxFeeHLSMath = sdkmath.NewInt(int64(minTxFeeHLS * 1000000000000000000))
	}
	if minBatchFeeHLS != 0.0 {
		minBatchFeeHLSMath = sdkmath.NewInt(int64(minBatchFeeHLS * 1000000000000000000))
	}

	msg, err := network.SendUpdateOrchestratorAddressesFeeMsg(ctx, chainId, minTxFeeHLSMath, minBatchFeeHLSMath)
	if err != nil {
		return err
	}
	err = network.SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
	if err != nil {
		return err
	}
	_, err = network.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return err
	}
	return nil
}
