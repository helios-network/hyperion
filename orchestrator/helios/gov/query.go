package gov

import (
	"context"
	"fmt"

	"github.com/pkg/errors"

	"github.com/Helios-Chain-Labs/metrics"
	query "github.com/cosmos/cosmos-sdk/types/query"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

var ErrNotFound = errors.New("not found")

type QClient interface {
	GetProposal(ctx context.Context, proposalId uint64) (*govtypes.Proposal, error)
	GetProposalsByPageAndSize(ctx context.Context, page int, size int) ([]*govtypes.Proposal, uint64, error)
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

func (c queryClient) GetProposalsByPageAndSize(ctx context.Context, page int, size int) ([]*govtypes.Proposal, uint64, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	response, err := c.QueryClient.Proposals(ctx, &govtypes.QueryProposalsRequest{
		Pagination: &query.PageRequest{
			Offset:     (uint64(page) - 1) * uint64(size),
			Limit:      uint64(size),
			Reverse:    true,
			CountTotal: true,
		},
	})
	if err != nil {
		return nil, 0, err
	}
	return response.Proposals, response.Pagination.Total, nil
}
