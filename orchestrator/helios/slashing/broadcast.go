package slashing

import (
	"context"
	"fmt"

	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/sdk-go/client/chain"

	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
)

type SlashingBroadcastClient interface {
	SendUnjail(ctx context.Context, validatorAddress string) error
}

type broadcastClient struct {
	chain.ChainClient

	svcTags metrics.Tags
}

func NewBroadcastClient(client chain.ChainClient) SlashingBroadcastClient {
	return broadcastClient{
		ChainClient: client,
		svcTags:     metrics.Tags{"svc": "gov_broadcast"},
	}
}

func (c broadcastClient) SendUnjail(ctx context.Context, validatorAddress string) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &slashingtypes.MsgUnjail{
		ValidatorAddr: validatorAddress,
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return fmt.Errorf("broadcasting MsgUnjail failed: %w", err)
	}

	if resp.TxResponse.Code == 13 {
		return fmt.Errorf("code 13 - insufficient fee")
	}

	return nil
}
