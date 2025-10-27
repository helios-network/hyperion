package queries

import (
	"context"
	"fmt"

	gethcommon "github.com/ethereum/go-ethereum/common"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func AddWhitelistedAddress(ctx context.Context, global *global.Global, chainId uint64, address string) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()

	// check if address is ethereum address
	if !gethcommon.IsHexAddress(address) {
		return nil, fmt.Errorf("address %s is not a valid hex encoded address", address)
	}

	msg, err := network.AddWhitelistedAddressMsg(ctx, chainId, address)
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
		"tx_hash": resp.TxHash,
	}, nil
}
