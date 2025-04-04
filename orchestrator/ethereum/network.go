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
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/committer"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/hyperion"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

type NetworkConfig struct {
	EthNodeRPC            string
	GasPriceAdjustment    float64
	MaxGasPrice           string
	PendingTxWaitDuration string
	EthNodeAlchemyWS      string
}

// Network is the orchestrator's reference endpoint to the Ethereum network
type Network interface {
	GetHeaderByNumber(ctx context.Context, number *big.Int) (*gethtypes.Header, error)
	GetHyperionID(ctx context.Context) (gethcommon.Hash, error)

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
	) (*gethcommon.Hash, error)

	GetLastEventNonce(ctx context.Context) (*big.Int, error)
	GetLastValsetCheckpoint(ctx context.Context) (*gethcommon.Hash, error)

	GetTxBatchNonce(ctx context.Context, erc20ContractAddress gethcommon.Address) (*big.Int, error)
	SendTransactionBatch(ctx context.Context,
		currentValset *hyperiontypes.Valset,
		batch *hyperiontypes.OutgoingTxBatch,
		confirms []*hyperiontypes.MsgConfirmBatch,
	) (*gethcommon.Hash, error)

	TokenDecimals(ctx context.Context, tokenContract gethcommon.Address) (uint8, error)
}

type network struct {
	hyperion.HyperionContract

	FromAddr gethcommon.Address
}

func NewNetwork(
	hyperionContractAddr,
	fromAddr gethcommon.Address,
	signerFn bind.SignerFn,
	cfg NetworkConfig,
) (Network, error) {
	evmRPC, err := rpc.Dial(cfg.EthNodeRPC)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to connect to ethereum RPC: %s", cfg.EthNodeRPC)
	}

	log.Info("fromAddr: ", fromAddr)
	ethCommitter, err := committer.NewEthCommitter(
		fromAddr,
		cfg.GasPriceAdjustment,
		cfg.MaxGasPrice,
		signerFn,
		provider.NewEVMProvider(evmRPC),
	)
	if err != nil {
		return nil, err
	}

	pendingTxDuration, err := time.ParseDuration(cfg.PendingTxWaitDuration)
	if err != nil {
		return nil, err
	}

	hyperionContract, err := hyperion.NewHyperionContract(ethCommitter, hyperionContractAddr, hyperion.PendingTxInputList{}, pendingTxDuration)
	if err != nil {
		return nil, err
	}

	// If Alchemy Websocket URL is set, then Subscribe to Pending Transaction of Hyperion Contract.
	if cfg.EthNodeAlchemyWS != "" {
		log.WithFields(log.Fields{
			"url": cfg.EthNodeAlchemyWS,
		}).Infoln("subscribing to Alchemy websocket")
		go hyperionContract.SubscribeToPendingTxs(cfg.EthNodeAlchemyWS)
	}

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

func (n *network) GetHeaderByNumber(ctx context.Context, number *big.Int) (*gethtypes.Header, error) {
	return n.Provider().HeaderByNumber(ctx, number)
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

func (n *network) GetTxBatchNonce(ctx context.Context, erc20ContractAddress gethcommon.Address) (*big.Int, error) {
	return n.HyperionContract.GetTxBatchNonce(ctx, erc20ContractAddress, n.FromAddr)
}

func (n *network) GetSendToHeliosEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
	hyperionFilterer, err := hyperionevents.NewHyperionFilterer(n.Address(), n.Provider())
	if err != nil {
		return nil, errors.Wrap(err, "failed to init Hyperion events filterer")
	}

	log.Infof("GetSendToHeliosEvents %d %d", startBlock, endBlock)

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
		log.Infof("NEW EVENTS OF SEND TO HELIOS FILTERED")
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
	fmt.Println("ALALALALLALALALLALAL-1")
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
		fmt.Println("ALALALALLALALALLALAL")
		valsetUpdatedEvents = append(valsetUpdatedEvents, iter.Event)
	}

	fmt.Println("ALALALALLALALALLALAL0")
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
