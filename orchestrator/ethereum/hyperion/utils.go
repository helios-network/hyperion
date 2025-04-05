package hyperion

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
		err = errors.Wrap(err, "StateLastEventNonce call failed")
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
		err = errors.Wrap(err, "StateLastEventNonce call failed")
		return nil, err
	}

	bts := []byte{}
	for _, b := range checkpointBytes {
		bts = append(bts, b)
	}

	fmt.Println("checkpointBytes: ", hexutil.Encode(bts))

	checkpoint := common.BytesToHash(checkpointBytes[:])
	return &checkpoint, nil
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
