package hyperion

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

// Gets the latest transaction batch nonce
func (s *hyperionContract) GetTxBatchNonce(
	ctx context.Context,
	erc20ContractAddress common.Address,
	callerAddress common.Address,
) (*big.Int, error) {

	nonce, err := s.ethHyperion.LastBatchNonce(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	}, erc20ContractAddress)

	if err != nil {
		err = errors.Wrap(err, "LastBatchNonce call failed")
		return nil, err
	}

	return nonce, nil
}

// Gets the latest validator set nonce
func (s *hyperionContract) GetValsetNonce(
	ctx context.Context,
	callerAddress common.Address,
) (*big.Int, error) {

	nonce, err := s.ethHyperion.StateLastValsetNonce(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateLastValsetNonce call failed")
		return nil, err
	}

	return nonce, nil
}

func (s *hyperionContract) GetLastEventNonce(
	ctx context.Context,
	callerAddress common.Address,
) (*big.Int, error) {

	nonce, err := s.ethHyperion.StateLastEventNonce(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateLastEventNonce call failed for hyperion contract "+s.hyperionAddress.Hex())
		return nil, err
	}

	return nonce, nil
}

func (s *hyperionContract) GetLastValsetCheckpoint(
	ctx context.Context,
	callerAddress common.Address,
) (*common.Hash, error) {

	checkpointBytes, err := s.ethHyperion.StateLastValsetCheckpoint(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateLastValsetCheckpoint call failed")
		return nil, err
	}

	// bts := []byte{}
	// for _, b := range checkpointBytes {
	// 	bts = append(bts, b)
	// }

	checkpoint := common.BytesToHash(checkpointBytes[:])
	return &checkpoint, nil
}

func (s *hyperionContract) GetLastValsetUpdatedEventHeight(
	ctx context.Context,
	callerAddress common.Address,
) (*big.Int, error) {

	height, err := s.ethHyperion.StateLastValsetHeight(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateLastValsetHeight call failed")
		return nil, err
	}

	return height, nil
}

func (s *hyperionContract) GetLastEventHeight(
	ctx context.Context,
	callerAddress common.Address,
) (*big.Int, error) {

	height, err := s.ethHyperion.StateLastEventHeight(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateLastEventHeight call failed")
		return nil, err
	}

	return height, nil
}

// Gets the hyperionID
func (s *hyperionContract) GetHyperionID(
	ctx context.Context,
	callerAddress common.Address,
) (common.Hash, error) {

	hyperionID, err := s.ethHyperion.StateHyperionId(&bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	})

	if err != nil {
		err = errors.Wrap(err, "StateHyperionId call failed")
		return common.Hash{}, err
	}

	return hyperionID, nil
}

func (s *hyperionContract) GetERC20Symbol(
	ctx context.Context,
	erc20ContractAddress common.Address,
	callerAddress common.Address,
) (symbol string, err error) {

	erc20Wrapper := bind.NewBoundContract(erc20ContractAddress, erc20ABI, s.ethProvider, nil, nil)

	callOpts := &bind.CallOpts{
		From:    callerAddress,
		Context: ctx,
	}
	var out []interface{}
	err = erc20Wrapper.Call(callOpts, &out, "symbol")
	if err != nil {
		err = errors.Wrap(err, "ERC20 [symbol] call failed")
		return "", err
	}

	symbol = *abi.ConvertType(out[0], new(string)).(*string)

	return symbol, nil
}

func (s *hyperionContract) SendInitializeBlockchainTx(
	ctx context.Context,
	callerAddress common.Address,
	hyperionId [32]byte,
	powerThreshold *big.Int,
	validators []common.Address,
	powers []*big.Int,
) (*gethtypes.Transaction, uint64, error) {

	tx, err := s.ethHyperion.Initialize(&bind.TransactOpts{
		From:    callerAddress,
		Context: ctx,
		Signer:  s.signerFn,
	}, hyperionId, powerThreshold, validators, powers)

	if err != nil {
		err = errors.Wrap(err, "Initialize call failed")
		return nil, 0, err
	}

	maxRetries := 50 // 50 * 5 seconds = 250 seconds = 4 minutes and 10 seconds
	retryCount := 0
	var txHash common.Hash
	fmt.Println("Checking deployment hyperion contract transaction status...", tx.Hash().String())
	for retryCount < maxRetries {
		fmt.Println("Checking deployment hyperion contract transaction status...", retryCount)
		tx, isPending, err := s.EVMCommitter.Provider().TransactionByHash(ctx, tx.Hash())
		if err != nil {
			time.Sleep(5 * time.Second)
			retryCount++
			continue
		}
		if isPending {
			time.Sleep(5 * time.Second)
			retryCount++
			continue
		}
		txHash = tx.Hash()
		break
	}

	if txHash == (common.Hash{}) {
		return nil, 0, errors.New("transaction not found on the blockchain")
	}

	receipt, err := s.EVMCommitter.Provider().TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, 0, err
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return nil, 0, errors.New("transaction failed")
	}

	return tx, receipt.BlockNumber.Uint64(), nil
}

func (s *hyperionContract) WaitForTransaction(ctx context.Context, txHash common.Hash) (*gethtypes.Transaction, uint64, error) {
	tx, isPending, err := s.EVMCommitter.Provider().TransactionByHash(ctx, txHash)
	if err != nil {
		fmt.Println("Error getting transaction by hash:", err)
		return nil, 0, err
	}

	if isPending {
		for isPending {
			time.Sleep(5 * time.Second)
			tx, isPending, err = s.EVMCommitter.Provider().TransactionByHash(ctx, tx.Hash())
			if err != nil {
				fmt.Println("Error getting transaction by hash:", err)
				return nil, 0, err
			}
		}
	}

	receipt, err := s.EVMCommitter.Provider().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, 0, err
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return nil, 0, errors.New("transaction failed")
	}

	return tx, receipt.BlockNumber.Uint64(), nil
}

func (s *hyperionContract) GetTransactionFeesUsedInNetworkNativeCurrency(ctx context.Context, txHash common.Hash) (*big.Int, uint64, error) {
	receipt, err := s.EVMCommitter.Provider().TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, 0, err
	}
	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return nil, 0, errors.New("transaction failed")
	}

	return big.NewInt(0).Mul(big.NewInt(int64(receipt.GasUsed)), receipt.EffectiveGasPrice), receipt.BlockNumber.Uint64(), nil
}

func (s *hyperionContract) DeployERC20(
	ctx context.Context,
	callerAddress common.Address,
	denom string,
	name string,
	symbol string,
	decimals uint8,
) (*gethtypes.Transaction, uint64, error) {

	tx, err := s.ethHyperion.DeployERC20(&bind.TransactOpts{
		From:    callerAddress,
		Context: ctx,
		Signer:  s.signerFn,
	}, denom, name, symbol, decimals)
	if err != nil {
		return nil, 0, err
	}

	time.Sleep(5 * time.Second)

	tx, isPending, err := s.EVMCommitter.Provider().TransactionByHash(ctx, tx.Hash())
	if err != nil {
		return nil, 0, err
	}

	if isPending {
		for isPending {
			time.Sleep(1 * time.Second)
			tx, isPending, err = s.EVMCommitter.Provider().TransactionByHash(ctx, tx.Hash())
			if err != nil {
				fmt.Println("Error getting transaction by hash:", err)
				return nil, 0, err
			}
		}
	}

	receipt, err := s.EVMCommitter.Provider().TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return nil, 0, err
	}

	if receipt.Status != gethtypes.ReceiptStatusSuccessful {
		return nil, 0, errors.New("transaction failed")
	}

	return tx, receipt.BlockNumber.Uint64(), nil
}
