package logos

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Helios-Chain-Labs/metrics"
	logostypes "github.com/Helios-Chain-Labs/sdk-go/chain/logos/types"
)

var ErrNotFound = errors.New("not found")

type QLogosClient interface {
	Logo(ctx context.Context, hash string) (string, error)
}

type queryClient struct {
	logostypes.QueryClient

	svcTags metrics.Tags
}

func NewQueryClient(client logostypes.QueryClient) QLogosClient {
	return queryClient{
		QueryClient: client,
		svcTags:     metrics.Tags{"svc": "logos_query"},
	}
}

func (c queryClient) Logo(ctx context.Context, hash string) (string, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	response, err := c.QueryClient.Logo(ctx, &logostypes.QueryLogoRequest{
		Hash: hash,
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to get logo")
	}
	return response.Logo.Data, nil
}
