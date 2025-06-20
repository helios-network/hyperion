package committer

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
)

// EVMCommitter defines an interface for submitting transactions
// into Ethereum, Matic, and other EVM-compatible networks.
type EVMCommitter interface {
	FromAddress() common.Address
	Provider() provider.EVMProvider
	SendTx(
		ctx context.Context,
		recipient common.Address,
		txData []byte,
	) (txHash common.Hash, cost *big.Int, err error)
	GetTransactOpts(ctx context.Context) *bind.TransactOpts
}

type EVMCommitterOption func(o *options) error

type options struct {
	GasPrice    decimal.Decimal
	GasLimit    uint64
	EstimateGas bool
	RPCTimeout  time.Duration
}

func defaultOptions() *options {
	v, _ := decimal.NewFromString("20")
	return &options{
		GasPrice:    v.Shift(9), // 20 gwei
		GasLimit:    1000000,
		EstimateGas: true,
		RPCTimeout:  10 * time.Second,
	}
}

func applyOptions(o *options, opts ...EVMCommitterOption) error {
	for _, oo := range opts {
		if err := oo(o); err != nil {
			err = errors.Wrap(err, "failed to apply option to EVMCommitter")
			return err
		}
	}

	return nil
}

func OptionGasPriceFromString(str string) EVMCommitterOption {
	return func(o *options) error {
		gasPrice, err := decimal.NewFromString(str)
		if err != nil {
			err = errors.Wrap(err, "unable to parse gas price from string to decimal")
			return err
		}

		o.GasPrice = gasPrice
		return nil
	}
}

func OptionEstimateGas(estimateGas bool) EVMCommitterOption {
	return func(o *options) error {
		o.EstimateGas = estimateGas
		return nil
	}
}

func ParseMaxGasPrice(maxGasPriceStr string) int64 {
	maxGasPriceStr = strings.TrimSpace(maxGasPriceStr)
	maxGasPriceStr = strings.ToLower(maxGasPriceStr)

	// If the denom is gwei, convert to wei
	unit := "gwei"
	isGwei := false
	if strings.HasSuffix(maxGasPriceStr, unit) {
		maxGasPriceStr = strings.TrimSuffix(maxGasPriceStr, unit)
		isGwei = true
	}

	// if denom is not present, consider it as wei
	maxGasPriceStr = strings.TrimSpace(maxGasPriceStr)
	maxGasPrice, err := decimal.NewFromString(maxGasPriceStr)
	if err != nil {
		err = errors.Wrap(err, "unable to parse max gas price. max_gas_price should be in gwei")
		panic(err)
	}

	if isGwei {
		maxGasPrice = maxGasPrice.Shift(9) // Gwei to wei
	}

	return maxGasPrice.IntPart()
}

func ParseGasPrice(gasPriceStr string) int64 {
	return ParseMaxGasPrice(gasPriceStr)
}

func OptionGasPriceFromDecimal(gasPrice decimal.Decimal) EVMCommitterOption {
	return func(o *options) error {
		o.GasPrice = gasPrice
		return nil
	}
}

func OptionGasPriceFromBigInt(i *big.Int) EVMCommitterOption {
	return func(o *options) error {
		o.GasPrice = decimal.NewFromBigInt(i, 0)
		return nil
	}
}

func OptionGasLimit(limit uint64) EVMCommitterOption {
	return func(o *options) error {
		o.GasLimit = limit
		return nil
	}
}

func TxBroadcastTimeout(dur time.Duration) EVMCommitterOption {
	return func(o *options) error {
		o.RPCTimeout = dur
		return nil
	}
}
