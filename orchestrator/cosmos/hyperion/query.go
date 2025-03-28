package hyperion

import (
	"context"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

var ErrNotFound = errors.New("not found")

type QueryClient interface {
	HyperionParams(ctx context.Context) (*hyperiontypes.Params, error)
	LastClaimEventByAddr(ctx context.Context, hyperionId uint64, validatorAccountAddress cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error)
	GetValidatorAddress(ctx context.Context, hyperionId uint64, addr gethcommon.Address) (cosmostypes.AccAddress, error)

	ValsetAt(ctx context.Context, hyperionId uint64, nonce uint64) (*hyperiontypes.Valset, error)
	CurrentValset(ctx context.Context, hyperionId uint64) (*hyperiontypes.Valset, error)
	OldestUnsignedValsets(ctx context.Context, hyperionId uint64, valAccountAddress cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error)
	LatestValsets(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.Valset, error)
	AllValsetConfirms(ctx context.Context, hyperionId uint64, nonce uint64) ([]*hyperiontypes.MsgValsetConfirm, error)

	OldestUnsignedTransactionBatch(ctx context.Context, hyperionId uint64, valAccountAddress cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error)
	LatestTransactionBatches(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.OutgoingTxBatch, error)
	UnbatchedTokensWithFees(ctx context.Context) ([]*hyperiontypes.BatchFees, error)
	TransactionBatchSignatures(ctx context.Context, hyperionId uint64, nonce uint64, tokenContract gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error)
}

type queryClient struct {
	hyperiontypes.QueryClient

	svcTags metrics.Tags
}

func NewQueryClient(client hyperiontypes.QueryClient) QueryClient {
	return queryClient{
		QueryClient: client,
		svcTags:     metrics.Tags{"svc": "hyperion_query"},
	}
}

func (c queryClient) ValsetAt(ctx context.Context, hyperionId uint64, nonce uint64) (*hyperiontypes.Valset, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryValsetRequestRequest{
		HyperionId: hyperionId,
		Nonce:      nonce,
	}

	resp, err := c.QueryClient.ValsetRequest(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query ValsetRequest from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Valset, nil
}

func (c queryClient) CurrentValset(ctx context.Context, hyperionId uint64) (*hyperiontypes.Valset, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.CurrentValset(ctx, &hyperiontypes.QueryCurrentValsetRequest{
		HyperionId: hyperionId,
	})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query CurrentValset from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Valset, nil
}

func (c queryClient) OldestUnsignedValsets(ctx context.Context, hyperionId uint64, valAccountAddress cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryLastPendingValsetRequestByAddrRequest{
		HyperionId: hyperionId,
		Address:    valAccountAddress.String(),
	}

	resp, err := c.QueryClient.LastPendingValsetRequestByAddr(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query LastPendingValsetRequestByAddr from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Valsets, nil
}

func (c queryClient) LatestValsets(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.Valset, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.LastValsetRequests(ctx, &hyperiontypes.QueryLastValsetRequestsRequest{
		HyperionId: hyperionId,
	})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query LastValsetRequests from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Valsets, nil
}

func (c queryClient) AllValsetConfirms(ctx context.Context, hyperionId uint64, nonce uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.ValsetConfirmsByNonce(ctx, &hyperiontypes.QueryValsetConfirmsByNonceRequest{
		HyperionId: hyperionId,
		Nonce:      nonce,
	})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query ValsetConfirmsByNonce from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Confirms, nil
}

func (c queryClient) OldestUnsignedTransactionBatch(ctx context.Context, hyperionId uint64, valAccountAddress cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryLastPendingBatchRequestByAddrRequest{
		HyperionId: hyperionId,
		Address:    valAccountAddress.String(),
	}

	resp, err := c.QueryClient.LastPendingBatchRequestByAddr(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query LastPendingBatchRequestByAddr from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Batch, nil
}

func (c queryClient) LatestTransactionBatches(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.OutgoingTxBatches(ctx, &hyperiontypes.QueryOutgoingTxBatchesRequest{
		HyperionId: hyperionId,
	})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query OutgoingTxBatches from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Batches, nil
}

func (c queryClient) UnbatchedTokensWithFees(ctx context.Context) ([]*hyperiontypes.BatchFees, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.BatchFees(ctx, &hyperiontypes.QueryBatchFeeRequest{})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query BatchFees from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.BatchFees, nil
}

func (c queryClient) TransactionBatchSignatures(ctx context.Context, hyperionId uint64, nonce uint64, tokenContract gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryBatchConfirmsRequest{
		HyperionId:      hyperionId,
		Nonce:           nonce,
		ContractAddress: tokenContract.String(),
	}

	resp, err := c.QueryClient.BatchConfirms(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query BatchConfirms from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Confirms, nil
}

func (c queryClient) LastClaimEventByAddr(ctx context.Context, hyperionId uint64, validatorAccountAddress cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
	log.Info("LastClaimEventByAddr", validatorAccountAddress)
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryLastEventByAddrRequest{
		Address:    validatorAccountAddress.String(),
		HyperionId: hyperionId,
	}

	resp, err := c.QueryClient.LastEventByAddr(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query LastEventByAddr from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.LastClaimEvent, nil
}

func (c queryClient) HyperionParams(ctx context.Context) (*hyperiontypes.Params, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.Params(ctx, &hyperiontypes.QueryParamsRequest{})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query HyperionParams from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return &resp.Params, nil
}

func (c queryClient) GetValidatorAddress(ctx context.Context, hyperionId uint64, addr gethcommon.Address) (cosmostypes.AccAddress, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryDelegateKeysByEthAddress{
		HyperionId: hyperionId,
		EthAddress: addr.Hex(),
	}

	resp, err := c.QueryClient.GetDelegateKeyByEth(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query GetDelegateKeyByEth from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	valAddr, err := cosmostypes.AccAddressFromBech32(resp.ValidatorAddress)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode validator address: %v", resp.ValidatorAddress)
	}

	return valAddr, nil
}
