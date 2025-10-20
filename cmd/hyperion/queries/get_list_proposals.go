package queries

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/global"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

func GetListProposals(ctx context.Context, global *global.Global, page int, size int) (map[string]interface{}, error) {
	fmt.Println("GetListProposals", page, size)
	network := *global.GetHeliosNetwork()
	proposals, total, err := network.GetProposalsByPageAndSize(ctx, page, size)
	if err != nil {
		return nil, err
	}

	proposalsResult := make([]map[string]interface{}, 0)

	for _, proposal := range proposals {
		formattedProposal, err := ParseProposal(proposal)
		if err != nil {
			continue
		}
		proposalsResult = append(proposalsResult, formattedProposal)
	}

	return map[string]interface{}{
		"proposals": proposalsResult,
		"total":     total,
	}, nil
}

func ParseProposal(proposal *govtypes.Proposal) (map[string]interface{}, error) {
	statusTypes := map[govtypes.ProposalStatus]interface{}{
		govtypes.ProposalStatus_PROPOSAL_STATUS_UNSPECIFIED:    "UNSPECIFIED",
		govtypes.ProposalStatus_PROPOSAL_STATUS_DEPOSIT_PERIOD: "DEPOSIT_PERIOD",
		govtypes.ProposalStatus_PROPOSAL_STATUS_VOTING_PERIOD:  "VOTING_PERIOD",
		govtypes.ProposalStatus_PROPOSAL_STATUS_PASSED:         "PASSED",
		govtypes.ProposalStatus_PROPOSAL_STATUS_REJECTED:       "REJECTED",
		govtypes.ProposalStatus_PROPOSAL_STATUS_FAILED:         "FAILED",
	}

	proposerAddr, err := sdk.AccAddressFromBech32(proposal.Proposer)
	if err != nil {
		return nil, err
	}
	details := make([]map[string]interface{}, 0)

	return map[string]interface{}{
		"id":         proposal.Id,
		"statusCode": proposal.Status,
		"status":     statusTypes[proposal.Status],
		"proposer":   common.BytesToAddress(proposerAddr.Bytes()).String(),
		"title":      proposal.Title,
		"metadata":   proposal.Metadata,
		"summary":    proposal.Summary,
		"details":    details,
		"options": []*govtypes.WeightedVoteOption{
			{Option: govtypes.OptionYes, Weight: "Yes"},
			{Option: govtypes.OptionAbstain, Weight: "Abstain"},
			{Option: govtypes.OptionNo, Weight: "No"},
			{Option: govtypes.OptionNoWithVeto, Weight: "No With Veto"},
		},
		"votingStartTime":    proposal.VotingStartTime,
		"votingEndTime":      proposal.VotingEndTime,
		"submitTime":         proposal.SubmitTime,
		"totalDeposit":       proposal.TotalDeposit,
		"minDeposit":         0,
		"finalTallyResult":   proposal.FinalTallyResult,
		"currentTallyResult": proposal.CurrentTallyResult,
	}, nil
}
