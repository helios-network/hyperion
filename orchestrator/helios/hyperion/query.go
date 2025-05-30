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
	GetCounterpartyChainParamsByChainId(ctx context.Context, chainId uint64) (*hyperiontypes.CounterpartyChainParams, error)
	LastClaimEventByAddr(ctx context.Context, hyperionId uint64, validatorAccountAddress cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error)
	GetValidatorAddress(ctx context.Context, hyperionId uint64, addr gethcommon.Address) (cosmostypes.AccAddress, error)
	GetListOfNetworksWhereRegistered(ctx context.Context, addr gethcommon.Address) ([]uint64, error)

	ValsetAt(ctx context.Context, hyperionId uint64, nonce uint64) (*hyperiontypes.Valset, error)
	CurrentValset(ctx context.Context, hyperionId uint64) (*hyperiontypes.Valset, error)
	OldestUnsignedValsets(ctx context.Context, hyperionId uint64, valAccountAddress cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error)
	LatestValsets(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.Valset, error)
	AllValsetConfirms(ctx context.Context, hyperionId uint64, nonce uint64) ([]*hyperiontypes.MsgValsetConfirm, error)

	QueryGetLastObservedEthereumBlockHeight(ctx context.Context, hyperionId uint64) (*hyperiontypes.LastObservedEthereumBlockHeight, error)
	QueryGetLastObservedEventNonce(ctx context.Context, hyperionId uint64) (uint64, error)

	OldestUnsignedTransactionBatch(ctx context.Context, hyperionId uint64, valAccountAddress cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error)
	LatestTransactionBatches(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.OutgoingTxBatch, error)
	UnbatchedTokensWithFees(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.BatchFees, error)
	TransactionBatchSignatures(ctx context.Context, hyperionId uint64, nonce uint64, tokenContract gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error)

	QueryTokenAddressToDenom(ctx context.Context, hyperionId uint64, tokenAddress gethcommon.Address) (string, bool, error)
	QueryDenomToTokenAddress(ctx context.Context, hyperionId uint64, denom string) (gethcommon.Address, bool, error)

	LatestTransactionExternalCallDataTxs(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.OutgoingExternalDataTx, error)
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

func (c queryClient) LatestTransactionExternalCallDataTxs(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.OutgoingExternalDataTx, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.OutgoingExternalDataTxs(ctx, &hyperiontypes.QueryOutgoingExternalDataTxsRequest{
		HyperionId: hyperionId,
	})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query OutgoingExternalDataTxs from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.Txs, nil
}

func (c queryClient) UnbatchedTokensWithFees(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.BatchFees, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.QueryClient.BatchFees(ctx, &hyperiontypes.QueryBatchFeeRequest{
		HyperionId: hyperionId,
	})
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

	log.Println("Try to get HyperionParams")

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

func (c queryClient) GetCounterpartyChainParamsByChainId(ctx context.Context, chainId uint64) (*hyperiontypes.CounterpartyChainParams, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.Println("Try to get GetCounterpartyChainParamsByChainId")

	resp, err := c.QueryClient.QueryGetCounterpartyChainParamsByChainId(ctx, &hyperiontypes.QueryGetCounterpartyChainParamsByChainIdRequest{
		ChainId: chainId,
	})
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query GetCounterpartyChainParamsByChainId from daemon")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.CounterpartyChainParams, nil
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

func (c queryClient) GetListOfNetworksWhereRegistered(ctx context.Context, addr gethcommon.Address) ([]uint64, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryGetDelegateKeysByAddressRequest{
		EthAddress: addr.Hex(),
	}

	log.Info("req: ", req)

	resp, err := c.QueryClient.QueryGetDelegateKeysByAddress(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query GetDelegateKeyByEth from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.ChainIds, nil
}

func (c queryClient) QueryGetLastObservedEthereumBlockHeight(ctx context.Context, hyperionId uint64) (*hyperiontypes.LastObservedEthereumBlockHeight, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryGetLastObservedEthereumBlockHeightRequest{
		HyperionId: hyperionId,
	}

	resp, err := c.QueryClient.QueryGetLastObservedEthereumBlockHeight(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "failed to query GetDelegateKeyByEth from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, ErrNotFound
	}

	return resp.LastObservedHeight, nil
}

func (c queryClient) QueryGetLastObservedEventNonce(ctx context.Context, hyperionId uint64) (uint64, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryGetLastObservedEventNonceRequest{
		HyperionId: hyperionId,
	}

	resp, err := c.QueryClient.QueryGetLastObservedEventNonce(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return 0, errors.Wrap(err, "failed to query GetDelegateKeyByEth from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return 0, ErrNotFound
	}

	return resp.LastObservedEventNonce, nil
}

func (c queryClient) QueryDenomToTokenAddress(ctx context.Context, hyperionId uint64, denom string) (gethcommon.Address, bool, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryDenomToTokenAddressRequest{
		HyperionId: hyperionId,
		Denom:      denom,
	}

	resp, err := c.QueryClient.DenomToTokenAddress(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return gethcommon.Address{}, false, errors.Wrap(err, "failed to query GetDelegateKeyByEth from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return gethcommon.Address{}, false, ErrNotFound
	}

	return gethcommon.HexToAddress(resp.TokenAddress), resp.CosmosOriginated, nil
}

func (c queryClient) QueryTokenAddressToDenom(ctx context.Context, hyperionId uint64, tokenAddress gethcommon.Address) (string, bool, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	req := &hyperiontypes.QueryTokenAddressToDenomRequest{
		HyperionId:   hyperionId,
		TokenAddress: tokenAddress.Hex(),
	}

	resp, err := c.QueryClient.TokenAddressToDenom(ctx, req)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return "", false, errors.Wrap(err, "failed to query TokenAddressToDenom from client")
	}

	if resp == nil {
		metrics.ReportFuncError(c.svcTags)
		return "", false, ErrNotFound
	}

	return resp.Denom, resp.CosmosOriginated, nil
}
