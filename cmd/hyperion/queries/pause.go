package queries

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func PauseOrUnpauseWithdrawal(ctx context.Context, global *global.Global, chainId uint64, pause bool) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	msg, err := network.PauseOrUnpauseHyperionWithdrawalMsg(ctx, chainId, pause)
	if err != nil {
		return nil, err
	}
	err = network.SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}
	resp, err := network.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"success": true,
		"tx_hash": resp.TxHash,
	}, nil
}

func PauseOrUnpauseDeposit(ctx context.Context, global *global.Global, chainId uint64, pause bool) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	counterpartyChainParams, err := network.GetCounterpartyChainParamsByChainId(ctx, chainId)
	if err != nil {
		return nil, err
	}

	heliosNetwork := global.GetHeliosNetwork()
	if heliosNetwork == nil {
		return nil, fmt.Errorf("helios network not initialized")
	}

	targetNetworks, err := global.InitTargetNetworks(counterpartyChainParams)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to initialize target network for chain %d", counterpartyChainParams.BridgeChainId))
	}

	if len(targetNetworks) == 0 {
		return nil, fmt.Errorf("no target networks found for chain %d", counterpartyChainParams.BridgeChainId)
	}
	targetNetwork := targetNetworks[0]

	hash, err := (*targetNetwork).PauseOrUnpauseDeposit(ctx, pause)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"success": true,
		"tx_hash": hash.Hex(),
	}, nil
}
