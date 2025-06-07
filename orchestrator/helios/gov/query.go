package gov

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Helios-Chain-Labs/metrics"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

var ErrNotFound = errors.New("not found")

type QClient interface {
	GetProposal(ctx context.Context, proposalId uint64) (*govtypes.Proposal, error)
}

type queryClient struct {
	govtypes.QueryClient

	svcTags metrics.Tags
}

func NewQueryClient(client govtypes.QueryClient) QClient {
	return queryClient{
		QueryClient: client,
		svcTags:     metrics.Tags{"svc": "hyperion_query"},
	}
}

func (c queryClient) GetProposal(ctx context.Context, proposalId uint64) (*govtypes.Proposal, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	response, err := c.QueryClient.Proposal(ctx, &govtypes.QueryProposalRequest{
		ProposalId: proposalId,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Proposal:", response.Proposal)
	return response.Proposal, nil
}
