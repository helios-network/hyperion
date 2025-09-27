package provider

import (
	"context"
	"errors"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	ethereumrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcs"
	"github.com/Helios-Chain-Labs/metrics"
)

type EVMProvider interface {
	bind.ContractCaller
	bind.ContractFilterer

	GetRpc() *rpcs.Rpc

	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	Balance(ctx context.Context, account common.Address) (*big.Int, error)
}

type EVMProviderWithRet interface {
	EVMProvider

	SendTransactionWithRet(ctx context.Context, tx *types.Transaction) (txHash common.Hash, err error)
	SendTransactionWithRetSync(ctx context.Context, tx *types.Transaction) (txHash common.Hash, err error)
}

type evmProviderWithRet struct {
	rpc       *rpcs.Rpc
	ethClient *ethclient.Client
	rpcClient *rpc.Client
	svcTags   metrics.Tags
}

func NewEVMProvider(rpc *rpcs.Rpc) EVMProviderWithRet {

	client, err := ethereumrpc.Dial(rpc.Url)
	if err != nil {
		// Failed to connect to the RPC, return nil
		return nil
	}
	ethClient := ethclient.NewClient(client)

	return &evmProviderWithRet{
		rpc:       rpc,
		ethClient: ethClient,
		rpcClient: client,
		svcTags: metrics.Tags{
			"svc": string("eth_provider"),
		},
	}
}

func (p *evmProviderWithRet) GetRpc() *rpcs.Rpc {
	return p.rpc
}

func (p *evmProviderWithRet) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	metrics.ReportFuncCall(p.svcTags)
	doneFn := metrics.ReportFuncTiming(p.svcTags)
	defer doneFn()

	logs, err := p.ethClient.FilterLogs(ctx, query)
	if err != nil {
		metrics.ReportFuncError(p.svcTags)
		return nil, err
	}
	return logs, nil
}

// SubscribeFilterLogs creates a background log filtering operation, returning
// a subscription immediately, which can be used to stream the found events.
func (p *evmProviderWithRet) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}

func (p *evmProviderWithRet) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	result, err := p.ethClient.CodeAt(ctx, contract, blockNumber)
	if err != nil && result != nil {
		metrics.ReportFuncError(p.svcTags)
		return nil, err
	}
	return result, nil
}

// CallContract executes an Ethereum contract call with the specified data as the
// input.
func (p *evmProviderWithRet) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	result, err := p.ethClient.CallContract(ctx, call, blockNumber)
	if err != nil && result != nil {
		metrics.ReportFuncError(p.svcTags)
		return nil, err
	}
	return result, nil
}

func (p *evmProviderWithRet) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	err := p.ethClient.SendTransaction(ctx, tx)
	if err != nil {
		metrics.ReportFuncError(p.svcTags)
		return err
	}
	return nil
}

func (p *evmProviderWithRet) SendTransactionWithRet(ctx context.Context, tx *types.Transaction) (txHash common.Hash, err error) {
	metrics.ReportFuncCall(p.svcTags)
	doneFn := metrics.ReportFuncTiming(p.svcTags)
	defer doneFn()

	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		metrics.ReportFuncError(p.svcTags)
		return common.Hash{}, err
	}

	err = p.rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(data))
	if err != nil {
		metrics.ReportFuncError(p.svcTags)
		return common.Hash{}, err
	}

	return txHash, nil
}

func (p *evmProviderWithRet) SendTransactionWithRetSync(ctx context.Context, tx *types.Transaction) (txHash common.Hash, err error) {
	metrics.ReportFuncCall(p.svcTags)
	doneFn := metrics.ReportFuncTiming(p.svcTags)
	defer doneFn()

	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		metrics.ReportFuncError(p.svcTags)
		return common.Hash{}, err
	}

	err = p.rpcClient.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(data))
	if err != nil {
		metrics.ReportFuncError(p.svcTags)
		return common.Hash{}, err
	}

	_, txIsPending, err := p.ethClient.TransactionByHash(ctx, txHash)
	if err != nil || !txIsPending {
		return txHash, nil
	}

	// wait for tx to be mined
	maxRetries := 10
	retryCount := 0
	for retryCount < maxRetries {
		time.Sleep(10 * time.Second)
		_, txIsPending, err = p.ethClient.TransactionByHash(ctx, txHash)
		if err != nil || !txIsPending {
			return txHash, nil
		}
		retryCount++
	}

	return txHash, errors.New("transaction is not mined after " + strconv.Itoa(maxRetries) + " retries")
}

// Implement other methods of EVMProvider using the pool
func (p *evmProviderWithRet) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return p.ethClient.PendingNonceAt(ctx, account)
}

func (p *evmProviderWithRet) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return p.ethClient.PendingCodeAt(ctx, account)
}

func (p *evmProviderWithRet) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return p.ethClient.EstimateGas(ctx, msg)
}

func (p *evmProviderWithRet) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return p.ethClient.SuggestGasTipCap(ctx)
}

func (p *evmProviderWithRet) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return p.ethClient.SuggestGasPrice(ctx)
}

func (p *evmProviderWithRet) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return p.ethClient.TransactionByHash(ctx, hash)
}

func (p *evmProviderWithRet) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return p.ethClient.TransactionReceipt(ctx, txHash)
}

func (p *evmProviderWithRet) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	return p.ethClient.HeaderByNumber(ctx, number)
}

func (p *evmProviderWithRet) Balance(ctx context.Context, account common.Address) (*big.Int, error) {
	return p.ethClient.BalanceAt(ctx, account, nil)
}
