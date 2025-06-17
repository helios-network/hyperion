package staking

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Helios-Chain-Labs/metrics"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var ErrNotFound = errors.New("not found")

type StakingQueryClient interface {
	GetValidator(ctx context.Context, validatorAddress string) (*stakingtypes.Validator, error)
}

type queryClient struct {
	stakingtypes.QueryClient

	svcTags metrics.Tags
}

func NewQueryClient(client stakingtypes.QueryClient) StakingQueryClient {
	return queryClient{
		QueryClient: client,
		svcTags:     metrics.Tags{"svc": "hyperion_query"},
	}
}

func (c queryClient) GetValidator(ctx context.Context, validatorAddress string) (*stakingtypes.Validator, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	response, err := c.QueryClient.Validator(ctx, &stakingtypes.QueryValidatorRequest{
		ValidatorAddr: validatorAddress,
	})
	if err != nil {
		return nil, err
	}
	return &response.Validator, nil
}
