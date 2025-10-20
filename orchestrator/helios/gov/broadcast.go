package gov

import (
	"context"
	"fmt"
	"strconv"

	"cosmossdk.io/math"
	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/sdk-go/client/chain"

	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govbasetypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

type BClient interface {
	SendProposal(ctx context.Context, title, description, msgContent string, accountAddress sdk.AccAddress, initialDeposit math.Int) (uint64, error)
	VoteOnProposal(ctx context.Context, proposalId uint64, accountAddress sdk.AccAddress) error
	VoteOnProposalWithOption(ctx context.Context, proposalId uint64, accountAddress sdk.AccAddress, voteOption govtypes.VoteOption) error
}

type broadcastClient struct {
	chain.ChainClient

	svcTags metrics.Tags
}

func NewBroadcastClient(client chain.ChainClient) BClient {
	return broadcastClient{
		ChainClient: client,
		svcTags:     metrics.Tags{"svc": "gov_broadcast"},
	}
}

func (c broadcastClient) SendProposal(ctx context.Context, title, description, msgContent string, accountAddress sdk.AccAddress, initialDeposit math.Int) (uint64, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	proposalContent := &hyperiontypes.HyperionProposal{
		Title:       title,
		Description: description,
		Msg:         msgContent,
	}
	authority := authtypes.NewModuleAddress(govbasetypes.ModuleName).String()
	contentMsg, err := govtypes.NewLegacyContent(proposalContent, authority) // todo : recheck here
	if err != nil {
		return 0, fmt.Errorf("error converting legacy content into proposal message: %w", err)
	}

	fmt.Println("contentMsg: ", contentMsg)

	contentAny, err := codectypes.NewAnyWithValue(contentMsg)
	if err != nil {
		return 0, fmt.Errorf("failed to pack content message: %w", err)
	}
	msg := &govtypes.MsgSubmitProposal{
		Messages: []*codectypes.Any{contentAny},
		InitialDeposit: sdk.NewCoins(
			sdk.NewCoin("ahelios", initialDeposit), // todo: change ahelios by default var
		),
		Proposer: accountAddress.String(),
		Metadata: "Optional metadata", // todo update !!
		Title:    title,
		Summary:  description,
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return 0, fmt.Errorf("broadcasting MsgSubmitProposal failed: %w", err)
	}

	if resp.TxResponse.Code == 13 {
		return 0, fmt.Errorf("code 13 - insufficient fee")
	}

	proposalId := func() string {
		for _, e := range resp.TxResponse.Events {
			for _, a := range e.Attributes {
				if string(a.Key) == "proposal_id" {
					return string(a.Value)
				}
			}
		}
		return ""
	}()

	return strconv.ParseUint(proposalId, 10, 64)
}

func (c broadcastClient) VoteOnProposal(ctx context.Context, proposalId uint64, accountAddress sdk.AccAddress) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &govtypes.MsgVote{
		ProposalId: proposalId,
		Voter:      accountAddress.String(),
		Option:     govtypes.OptionYes,
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return fmt.Errorf("broadcasting MsgVote failed: %w", err)
	}

	if resp.TxResponse.Code == 13 {
		return fmt.Errorf("code 13 - insufficient fee")
	}

	return nil
}

func (c broadcastClient) VoteOnProposalWithOption(ctx context.Context, proposalId uint64, accountAddress sdk.AccAddress, voteOption govtypes.VoteOption) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &govtypes.MsgVote{
		ProposalId: proposalId,
		Voter:      accountAddress.String(),
		Option:     voteOption,
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return fmt.Errorf("broadcasting MsgVote failed: %w", err)
	}

	if resp.TxResponse.Code == 13 {
		return fmt.Errorf("code 13 - insufficient fee")
	}

	return nil
}
