package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

func VoteForProposal(ctx context.Context, global *global.Global, proposalId uint64, voteOption govtypes.VoteOption) (map[string]interface{}, error) {
	network := *global.GetHeliosNetwork()
	msg, err := network.VoteOnProposalWithOptionMsg(ctx, proposalId, global.GetCosmosAddress(), voteOption)
	if err != nil {
		return map[string]interface{}{
			"proposalId": proposalId,
			"voteOption": voteOption,
			"success":    false,
			"error":      err.Error(),
		}, nil
	}
	err = network.SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
	if err != nil {
		return map[string]interface{}{
			"proposalId": proposalId,
			"voteOption": voteOption,
			"success":    false,
			"error":      err.Error(),
		}, nil
	}
	_, err = global.SyncBroadcastMsgs(ctx, []sdk.Msg{msg})
	if err != nil {
		return map[string]interface{}{
			"proposalId": proposalId,
			"voteOption": voteOption,
			"success":    false,
			"error":      err.Error(),
		}, nil
	}
	return map[string]interface{}{
		"proposalId": proposalId,
		"voteOption": voteOption,
		"success":    true,
		"error":      "",
	}, nil
}
