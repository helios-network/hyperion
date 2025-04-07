package provider

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/Helios-Chain-Labs/metrics"
)

type EVMProvider interface {
	bind.ContractCaller
	bind.ContractFilterer

	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error)
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

type EVMProviderWithRet interface {
	EVMProvider

	SendTransactionWithRet(ctx context.Context, tx *types.Transaction) (txHash common.Hash, err error)
}

type evmProviderWithRet struct {
	pool    *EVMProviders
	svcTags metrics.Tags
}

func NewEVMProvider(rpcUrls string) EVMProviderWithRet {
	// split the RPC URLs by comma
	rpcUrls = strings.TrimSpace(rpcUrls)
	// split by comma to rpcUrlsTuple
	rpcUrlsTuple := strings.Split(rpcUrls, ",")
	pool := NewEVMProviders(rpcUrlsTuple)
	return &evmProviderWithRet{
		pool: pool,
		svcTags: metrics.Tags{
			"svc": string("eth_provider"),
		},
	}
}

func (p *evmProviderWithRet) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	metrics.ReportFuncCall(p.svcTags)
	doneFn := metrics.ReportFuncTiming(p.svcTags)
	defer doneFn()

	var logs []types.Log
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		logs, err = client.FilterLogs(ctx, query)
		return err
	})
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
	result := make([]byte, 0)
	var err error
	p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		result, err = client.CodeAt(ctx, contract, blockNumber)
		return nil
	})
	if err != nil && result != nil {
		metrics.ReportFuncError(p.svcTags)
		return nil, err
	}
	return result, nil
}

// CallContract executes an Ethereum contract call with the specified data as the
// input.
func (p *evmProviderWithRet) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	result := make([]byte, 0)
	var err error
	p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		result, err = client.CallContract(ctx, call, blockNumber)
		return nil
	})
	if err != nil && result != nil {
		metrics.ReportFuncError(p.svcTags)
		return nil, err
	}
	return result, nil
}

func (p *evmProviderWithRet) SendTransaction(ctx context.Context, tx *types.Transaction) error {
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

	err = p.pool.CallRpcClientWithRetry(ctx, func(client *rpc.Client) error {
		return client.CallContext(ctx, &txHash, "eth_sendRawTransaction", hexutil.Encode(data))
	})
	if err != nil {
		metrics.ReportFuncError(p.svcTags)
		return common.Hash{}, err
	}

	return txHash, nil
}

// Implement other methods of EVMProvider using the pool
func (p *evmProviderWithRet) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	var nonce uint64
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		nonce, err = client.PendingNonceAt(ctx, account)
		return err
	})
	return nonce, err
}

func (p *evmProviderWithRet) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	var code []byte
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		code, err = client.PendingCodeAt(ctx, account)
		return err
	})
	return code, err
}

func (p *evmProviderWithRet) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	var gas uint64
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		gas, err = client.EstimateGas(ctx, msg)
		return err
	})
	return gas, err
}

func (p *evmProviderWithRet) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	var tipCap *big.Int
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		tipCap, err = client.SuggestGasTipCap(ctx)
		return err
	})
	return tipCap, err
}

func (p *evmProviderWithRet) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	var gasPrice *big.Int
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		gasPrice, err = client.SuggestGasPrice(ctx)
		return err
	})
	return gasPrice, err
}

func (p *evmProviderWithRet) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	var tx *types.Transaction
	var isPending bool
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		tx, isPending, err = client.TransactionByHash(ctx, hash)
		return err
	})
	return tx, isPending, err
}

func (p *evmProviderWithRet) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		receipt, err = client.TransactionReceipt(ctx, txHash)
		return err
	})
	return receipt, err
}

func (p *evmProviderWithRet) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	var header *types.Header
	err := p.pool.CallEthClientWithRetry(ctx, func(client *ethclient.Client) error {
		var err error
		header, err = client.HeaderByNumber(ctx, number)
		return err
	})
	return header, err
}
