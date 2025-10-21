package queries

import (
	"context"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

func VoteForProposal(ctx context.Context, global *global.Global, proposalId uint64, voteOption govtypes.VoteOption) (map[string]interface{}, error) {
	err := global.VoteOnProposalWithOption(proposalId, voteOption)
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
