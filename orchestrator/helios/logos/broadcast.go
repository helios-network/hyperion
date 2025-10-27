package logos

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Helios-Chain-Labs/metrics"
	logostypes "github.com/Helios-Chain-Labs/sdk-go/chain/logos/types"
	"github.com/Helios-Chain-Labs/sdk-go/client/chain"
)

type BLogosClient interface {
	StoreLogo(ctx context.Context, logo string) (string, error)
	StoreLogoMsg(ctx context.Context, logo string) (sdk.Msg, error)
}

type broadcastClient struct {
	chain.ChainClient

	svcTags metrics.Tags
}

func NewBroadcastClient(client chain.ChainClient) BLogosClient {
	return broadcastClient{
		ChainClient: client,
		svcTags:     metrics.Tags{"svc": "logos_broadcast"},
	}
}

func (c broadcastClient) StoreLogo(ctx context.Context, logo string) (string, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &logostypes.MsgStoreLogoRequest{
		Creator: c.FromAddress().String(),
		Data:    logo,
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return "", fmt.Errorf("broadcasting MsgStoreLogo failed: %w", err)
	}

	if resp.TxResponse.Code == 13 {
		return "", fmt.Errorf("code 13 - insufficient fee")
	}

	return resp.TxResponse.TxHash, nil
}

func (c broadcastClient) StoreLogoMsg(ctx context.Context, logo string) (sdk.Msg, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &logostypes.MsgStoreLogoRequest{
		Creator: c.FromAddress().String(),
		Data:    logo,
	}
	return msg, nil
}
