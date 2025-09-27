package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/committer"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/hyperion"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/keystore"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/rpcs"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

type NetworkConfig struct {
	EthNodeRPC            *rpcs.Rpc
	GasPriceAdjustment    float64
	MaxGasPrice           string
	PendingTxWaitDuration string
	ChainID               int
}

// Network is the orchestrator's reference endpoint to the Ethereum network
type Network interface {
	FromAddress() gethcommon.Address

	GetRpc() *rpcs.Rpc
	TestRpc(ctx context.Context) bool

	GetHeaderByNumber(ctx context.Context, number *big.Int) (*gethtypes.Header, error)
	GetNativeBalance(ctx context.Context) (*big.Int, error)
	GetHyperionID(ctx context.Context) (gethcommon.Hash, error)
	GetGasPrice(ctx context.Context) (*big.Int, error)

	GetSendToHeliosEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error)
	GetHyperionERC20DeployedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error)
	GetValsetUpdatedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error)
	GetValsetUpdatedEventsAtSpecificBlock(block uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error)
	GetTransactionBatchExecutedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error)

	GetValsetNonce(ctx context.Context) (*big.Int, error)
	SendEthValsetUpdate(ctx context.Context,
		oldValset *hyperiontypes.Valset,
		newValset *hyperiontypes.Valset,
		confirms []*hyperiontypes.MsgValsetConfirm,
	) (*gethcommon.Hash, *big.Int, error)

	SendInitializeBlockchainTx(
		ctx context.Context,
		callerAddress gethcommon.Address,
		hyperionId [32]byte,
		powerThreshold *big.Int,
		validators []gethcommon.Address,
		powers []*big.Int,
	) (*gethtypes.Transaction, uint64, error)

	DeployERC20(
		ctx context.Context,
		callerAddress gethcommon.Address,
		denom string,
		name string,
		symbol string,
		decimals uint8,
	) (*gethtypes.Transaction, uint64, error)

	GetLastEventNonce(ctx context.Context) (*big.Int, error)
	GetLastValsetCheckpoint(ctx context.Context) (*gethcommon.Hash, error)
	GetLastValsetUpdatedEventHeight(ctx context.Context) (*big.Int, error)
	GetLastEventHeight(ctx context.Context) (*big.Int, error)

	GetTxBatchNonce(ctx context.Context, erc20ContractAddress gethcommon.Address) (*big.Int, error)
	PrepareTransactionBatch(ctx context.Context,
		currentValset *hyperiontypes.Valset,
		batch *hyperiontypes.OutgoingTxBatch,
		confirms []*hyperiontypes.MsgConfirmBatch,
	) ([]byte, error)

	SendPreparedTx(ctx context.Context,
		txData []byte,
	) (*gethcommon.Hash, *big.Int, error)

	SendPreparedTxSync(ctx context.Context,
		txData []byte,
	) (*gethcommon.Hash, *big.Int, error)

	TokenDecimals(ctx context.Context, tokenContract gethcommon.Address) (uint8, error)
	TokenSymbol(ctx context.Context, tokenContract gethcommon.Address) (string, error)
	ExecuteExternalDataTx(ctx context.Context, address gethcommon.Address, txAbi []byte, blockNumber *big.Int) ([]byte, []byte, string, error)
	GetSignerFn() bind.SignerFn
	GetPersonalSignFn() keystore.PersonalSignFn

	WaitForTransaction(ctx context.Context, txHash gethcommon.Hash) (*gethtypes.Transaction, uint64, error)
	GetTransactionFeesUsedInNetworkNativeCurrency(ctx context.Context, txHash gethcommon.Hash) (*big.Int, uint64, error)
	SendClaimTokensOfOldContract(ctx context.Context, hyperionId uint64, tokenContract string, amountInSdkMath *big.Int, ethFrom common.Address, signerFn keystore.PersonalSignFn) error
}

type network struct {
	hyperion.HyperionContract

	FromAddr       gethcommon.Address
	SignerFn       bind.SignerFn
	PersonalSignFn keystore.PersonalSignFn
}

func DeployNewHyperionContract(
	fromAddr gethcommon.Address,
	signerFn bind.SignerFn,
	cfg NetworkConfig,
	options ...committer.EVMCommitterOption,
) (gethcommon.Address, uint64, error, bool) {
	ethCommitter, err := committer.NewEthCommitter(
		fromAddr,
		cfg.GasPriceAdjustment,
		cfg.MaxGasPrice,
		signerFn,
		provider.NewEVMProvider(cfg.EthNodeRPC),
		options...,
	)
	if err != nil {
		return gethcommon.Address{}, 0, err, false
	}
	fmt.Println("Deploying Hyperion contract...", cfg.EthNodeRPC)
	return hyperion.DeployHyperionContract(context.Background(), ethCommitter)
}

func NewNetwork(
	hyperionContractAddr,
	fromAddr gethcommon.Address,
	signerFn bind.SignerFn,
	personalSignFn keystore.PersonalSignFn,
	cfg NetworkConfig,
	options ...committer.EVMCommitterOption,
) (Network, error) {
	ethCommitter, err := committer.NewEthCommitter(
		fromAddr,
		cfg.GasPriceAdjustment,
		cfg.MaxGasPrice,
		signerFn,
		provider.NewEVMProvider(cfg.EthNodeRPC),
		options...,
	)
	if err != nil {
		return nil, err
	}

	pendingTxDuration, err := time.ParseDuration(cfg.PendingTxWaitDuration)
	if err != nil {
		return nil, err
	}

	hyperionContract, err := hyperion.NewHyperionContract(context.Background(), ethCommitter, hyperionContractAddr, hyperion.PendingTxInputList{}, pendingTxDuration, signerFn)
	if err != nil {
		return nil, err
	}

	n := &network{
		HyperionContract: hyperionContract,
		FromAddr:         fromAddr,
		SignerFn:         signerFn,
		PersonalSignFn:   personalSignFn,
	}

	return n, nil
}

func (n *network) TokenDecimals(ctx context.Context, tokenContract gethcommon.Address) (uint8, error) {
	msg := ethereum.CallMsg{
		To:   &tokenContract,
		Data: gethcommon.Hex2Bytes("313ce567"), // decimals() method signature
	}

	res, err := n.Provider().CallContract(ctx, msg, nil)
	if err != nil {
		return 0, err
	}

	if len(res) == 0 {
		return 0, errors.Errorf("no decimals found for token contract %s", tokenContract.Hex())
	}

	return uint8(big.NewInt(0).SetBytes(res).Uint64()), nil
}

func (n *network) TokenSymbol(ctx context.Context, tokenContract gethcommon.Address) (string, error) {
	msg := ethereum.CallMsg{
		To:   &tokenContract,
		Data: gethcommon.Hex2Bytes("95d89b41"), // symbol() method signature
	}

	res, err := n.Provider().CallContract(ctx, msg, nil)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(res), "\x00"), nil
}

func (n *network) GetSignerFn() bind.SignerFn {
	return n.SignerFn
}

func (n *network) GetPersonalSignFn() keystore.PersonalSignFn {
	return n.PersonalSignFn
}

func (n *network) FromAddress() gethcommon.Address {
	return n.FromAddr
}

func (n *network) GetGasPrice(ctx context.Context) (*big.Int, error) {
	return n.Provider().SuggestGasPrice(ctx)
}

func (n *network) GetHeaderByNumber(ctx context.Context, number *big.Int) (*gethtypes.Header, error) {
	return n.Provider().HeaderByNumber(ctx, number)
}

func (n *network) GetNativeBalance(ctx context.Context) (*big.Int, error) {
	return n.Provider().Balance(ctx, n.FromAddr)
}

func (n *network) GetHyperionID(ctx context.Context) (gethcommon.Hash, error) {
	return n.HyperionContract.GetHyperionID(ctx, n.FromAddr)
}

func (n *network) GetValsetNonce(ctx context.Context) (*big.Int, error) {
	return n.HyperionContract.GetValsetNonce(ctx, n.FromAddr)
}

func (n *network) GetLastEventNonce(ctx context.Context) (*big.Int, error) {
	return n.HyperionContract.GetLastEventNonce(ctx, n.FromAddr)
}

func (n *network) GetLastValsetCheckpoint(ctx context.Context) (*gethcommon.Hash, error) {
	return n.HyperionContract.GetLastValsetCheckpoint(ctx, n.FromAddr)
}

func (n *network) GetLastValsetUpdatedEventHeight(ctx context.Context) (*big.Int, error) {
	return n.HyperionContract.GetLastValsetUpdatedEventHeight(ctx, n.FromAddr)
}

func (n *network) GetLastEventHeight(ctx context.Context) (*big.Int, error) {
	return n.HyperionContract.GetLastEventHeight(ctx, n.FromAddr)
}

func (n *network) GetTxBatchNonce(ctx context.Context, erc20ContractAddress gethcommon.Address) (*big.Int, error) {
	return n.HyperionContract.GetTxBatchNonce(ctx, erc20ContractAddress, n.FromAddr)
}

func (n *network) GetSendToHeliosEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
	hyperionFilterer, err := hyperionevents.NewHyperionFilterer(n.Address(), n.Provider())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init Hyperion events filterer")
	}

	iter, err := hyperionFilterer.FilterSendToHeliosEvent(&bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}, nil, nil, nil)
	if err != nil {
		if !isUnknownBlockErr(err) {
			return nil, errors.Wrap(err, "failed to scan past SendToHelios events from Ethereum")
		} else if iter == nil {
			return nil, errors.New("no iterator returned")
		}
	}

	defer iter.Close()

	var sendToHeliosEvents []*hyperionevents.HyperionSendToHeliosEvent
	for iter.Next() {
		sendToHeliosEvents = append(sendToHeliosEvents, iter.Event)
	}

	return sendToHeliosEvents, nil
}

func (n *network) GetHyperionERC20DeployedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error) {
	hyperionFilterer, err := hyperionevents.NewHyperionFilterer(n.Address(), n.Provider())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init Hyperion events filterer")
	}

	iter, err := hyperionFilterer.FilterERC20DeployedEvent(&bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}, nil)
	if err != nil {
		if !isUnknownBlockErr(err) {
			return nil, errors.Wrap(err, "failed to scan past TransactionBatchExecuted events from Ethereum")
		} else if iter == nil {
			return nil, errors.New("no iterator returned")
		}
	}

	defer iter.Close()

	var transactionBatchExecutedEvents []*hyperionevents.HyperionERC20DeployedEvent
	for iter.Next() {
		transactionBatchExecutedEvents = append(transactionBatchExecutedEvents, iter.Event)
	}

	return transactionBatchExecutedEvents, nil
}

func (n *network) GetValsetUpdatedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
	hyperionFilterer, err := hyperionevents.NewHyperionFilterer(n.Address(), n.Provider())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init Hyperion events filterer")
	}

	iter, err := hyperionFilterer.FilterValsetUpdatedEvent(&bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}, nil)
	if err != nil {
		if !isUnknownBlockErr(err) {
			return nil, errors.Wrap(err, "failed to scan past ValsetUpdatedEvent events from Ethereum")
		} else if iter == nil {
			return nil, errors.New("no iterator returned")
		}
	}

	defer iter.Close()

	var valsetUpdatedEvents []*hyperionevents.HyperionValsetUpdatedEvent
	for iter.Next() {
		valsetUpdatedEvents = append(valsetUpdatedEvents, iter.Event)
	}

	return valsetUpdatedEvents, nil
}

func (n *network) GetValsetUpdatedEventsAtSpecificBlock(block uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
	hyperionFilterer, err := hyperionevents.NewHyperionFilterer(n.Address(), n.Provider())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init Hyperion events filterer")
	}

	iter, err := hyperionFilterer.FilterValsetUpdatedEvent(&bind.FilterOpts{
		Start: block,
		End:   &block,
	}, nil)
	if err != nil {
		if !isUnknownBlockErr(err) {
			return nil, errors.Wrap(err, "failed to scan past ValsetUpdatedEvent events from Ethereum")
		} else if iter == nil {
			return nil, errors.New("no iterator returned")
		}
	}

	defer iter.Close()

	var valsetUpdatedEvents []*hyperionevents.HyperionValsetUpdatedEvent
	for iter.Next() {
		valsetUpdatedEvents = append(valsetUpdatedEvents, iter.Event)
	}

	return valsetUpdatedEvents, nil
}

func (n *network) GetTransactionBatchExecutedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error) {
	hyperionFilterer, err := hyperionevents.NewHyperionFilterer(n.Address(), n.Provider())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init Hyperion events filterer")
	}

	iter, err := hyperionFilterer.FilterTransactionBatchExecutedEvent(&bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}, nil, nil)
	if err != nil {
		if !isUnknownBlockErr(err) {
			return nil, errors.Wrap(err, "failed to scan past TransactionBatchExecuted events from Ethereum")
		} else if iter == nil {
			return nil, errors.New("no iterator returned")
		}
	}

	defer iter.Close()

	var transactionBatchExecutedEvents []*hyperionevents.HyperionTransactionBatchExecutedEvent
	for iter.Next() {
		transactionBatchExecutedEvents = append(transactionBatchExecutedEvents, iter.Event)
	}

	return transactionBatchExecutedEvents, nil
}

func isUnknownBlockErr(err error) bool {
	// Geth error
	if strings.Contains(err.Error(), "unknown block") {
		return true
	}

	// Parity error
	if strings.Contains(err.Error(), "One of the blocks specified in filter") {
		return true
	}

	return false
}

func (n *network) ExecuteExternalDataTx(ctx context.Context, address gethcommon.Address, txAbi []byte, blockNumber *big.Int) ([]byte, []byte, string, error) {
	msg := ethereum.CallMsg{
		To:   &address,
		Data: txAbi,
	}
	result, err := n.Provider().CallContract(ctx, msg, blockNumber)
	if err != nil {
		return []byte{}, []byte(err.Error()), n.Provider().GetRpc().Url, err
	}
	return result, []byte{}, n.Provider().GetRpc().Url, nil
}

func (n *network) TestRpc(ctx context.Context) bool {
	_, err := n.Provider().HeaderByNumber(ctx, nil)
	if err != nil {
		return false
	}
	_, err = n.Provider().Balance(ctx, n.FromAddr)
	if err != nil {
		return false
	}
	return true
}

func (n *network) GetRpc() *rpcs.Rpc {
	return n.Provider().GetRpc()
}
