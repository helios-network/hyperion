package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/committer"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/hyperion"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

type NetworkConfig struct {
	EthNodeRPCs           []*hyperiontypes.Rpc
	GasPriceAdjustment    float64
	MaxGasPrice           string
	PendingTxWaitDuration string
	ChainID               int
}

// Network is the orchestrator's reference endpoint to the Ethereum network
type Network interface {
	FromAddress() gethcommon.Address
	GetLastUsedRpc() string
	ReduceReputationOfLastRpc()
	RemoveLastUsedRpc()
	RemoveRpc(targetUrl string) bool
	TestRpcs(ctx context.Context) bool
	GetRpcs() []*hyperiontypes.Rpc

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

	GetLastEventNonce(ctx context.Context) (*big.Int, error)
	GetLastValsetCheckpoint(ctx context.Context) (*gethcommon.Hash, error)
	GetLastValsetUpdatedEventHeight(ctx context.Context) (*big.Int, error)
	GetLastEventHeight(ctx context.Context) (*big.Int, error)

	GetTxBatchNonce(ctx context.Context, erc20ContractAddress gethcommon.Address) (*big.Int, error)
	SendTransactionBatch(ctx context.Context,
		currentValset *hyperiontypes.Valset,
		batch *hyperiontypes.OutgoingTxBatch,
		confirms []*hyperiontypes.MsgConfirmBatch,
	) (*gethcommon.Hash, *big.Int, error)

	TokenDecimals(ctx context.Context, tokenContract gethcommon.Address) (uint8, error)
	ExecuteExternalDataTx(ctx context.Context, address gethcommon.Address, txAbi []byte, blockNumber *big.Int) ([]byte, []byte, string, error)
}

type network struct {
	hyperion.HyperionContract

	FromAddr gethcommon.Address
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
		provider.NewEVMProvider(cfg.EthNodeRPCs),
		options...,
	)
	if err != nil {
		return gethcommon.Address{}, 0, err, false
	}
	fmt.Println("Deploying Hyperion contract...", cfg.EthNodeRPCs)
	return hyperion.DeployHyperionContract(context.Background(), ethCommitter)
}

func NewNetwork(
	hyperionContractAddr,
	fromAddr gethcommon.Address,
	signerFn bind.SignerFn,
	cfg NetworkConfig,
	options ...committer.EVMCommitterOption,
) (Network, error) {
	log.Info("fromAddr: ", fromAddr)
	ethCommitter, err := committer.NewEthCommitter(
		fromAddr,
		cfg.GasPriceAdjustment,
		cfg.MaxGasPrice,
		signerFn,
		provider.NewEVMProvider(cfg.EthNodeRPCs),
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

	// If Alchemy Websocket URL is set, then Subscribe to Pending Transaction of Hyperion Contract.
	// disabled for now
	// if cfg.EthNodeAlchemyWS != "" {
	// 	log.WithFields(log.Fields{
	// 		"url": cfg.EthNodeAlchemyWS,
	// 	}).Infoln("subscribing to Alchemy websocket")
	// 	go hyperionContract.SubscribeToPendingTxs(cfg.EthNodeAlchemyWS)
	// }

	//ethCommitter.Provider().CallContract(context.Background(), ethereum.CallMsg{}, nil)

	n := &network{
		HyperionContract: hyperionContract,
		FromAddr:         fromAddr,
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

func (n *network) FromAddress() gethcommon.Address {
	return n.FromAddr
}

func (n *network) GetLastUsedRpc() string {
	return n.Provider().GetLastUsedRpc()
}

func (n *network) ReduceReputationOfLastRpc() {
	n.Provider().ReduceReputationOfLastRpc()
}

func (n *network) RemoveLastUsedRpc() {
	n.Provider().RemoveLastUsedRpc()
}

func (n *network) GetGasPrice(ctx context.Context) (*big.Int, error) {
	return n.Provider().SuggestGasPrice(ctx)
}

func (n *network) TestRpcs(ctx context.Context) bool {
	return n.Provider().TestRpcs(ctx, func(client *ethclient.Client, url string) error {
		_, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return err
		}

		_, err = client.BalanceAt(ctx, n.FromAddr, nil)
		if err != nil {
			return err
		}
		fmt.Println("ok: ", url)
		return nil
	})
}

func (n *network) GetRpcs() []*hyperiontypes.Rpc {
	return n.Provider().GetRpcs()
}

func (n *network) RemoveRpc(targetUrl string) bool {
	return n.Provider().RemoveRpc(targetUrl)
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
		return []byte{}, []byte(err.Error()), n.Provider().GetLastUsedRpc(), err
	}
	return result, []byte{}, n.Provider().GetLastUsedRpc(), nil
}
