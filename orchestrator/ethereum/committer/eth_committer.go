package committer

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/metrics"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/util"
)

// NewEthCommitter returns an instance of EVMCommitter, which
// can be used to submit txns into Ethereum, Matic, and other EVM-compatible networks.
func NewEthCommitter(
	fromAddress common.Address,
	ethGasPriceAdjustment float64,
	ethMaxGasPrice string,
	fromSigner bind.SignerFn,
	evmProvider provider.EVMProviderWithRet,
	committerOpts ...EVMCommitterOption,
) (EVMCommitter, error) {
	committer := &ethCommitter{
		committerOpts: defaultOptions(),
		svcTags: metrics.Tags{
			"module": "eth_committer",
		},

		ethGasPriceAdjustment: ethGasPriceAdjustment,
		ethMaxGasPrice:        ParseMaxGasPrice(ethMaxGasPrice),
		fromAddress:           fromAddress,
		fromSigner:            fromSigner,
		evmProvider:           evmProvider,
		nonceCache:            util.NewNonceCache(),
	}

	if err := applyOptions(committer.committerOpts, committerOpts...); err != nil {
		return nil, err
	}

	committer.nonceCache.Sync(fromAddress, func() (uint64, error) {
		nonce, err := evmProvider.PendingNonceAt(context.TODO(), fromAddress)
		return nonce, err
	})

	return committer, nil
}

type ethCommitter struct {
	committerOpts *options

	fromAddress common.Address
	fromSigner  bind.SignerFn

	ethGasPriceAdjustment float64
	ethMaxGasPrice        int64
	evmProvider           provider.EVMProviderWithRet
	nonceCache            util.NonceCache

	svcTags metrics.Tags
}

func (e *ethCommitter) FromAddress() common.Address {
	return e.fromAddress
}

func (e *ethCommitter) Provider() provider.EVMProvider {
	return e.evmProvider
}

func (e *ethCommitter) SendTx(
	ctx context.Context,
	recipient common.Address,
	txData []byte,
) (txHash common.Hash, err error) {
	metrics.ReportFuncCall(e.svcTags)
	doneFn := metrics.ReportFuncTiming(e.svcTags)
	defer doneFn()

	opts := &bind.TransactOpts{
		From:   e.fromAddress,
		Signer: e.fromSigner,

		GasPrice: e.committerOpts.GasPrice.BigInt(),
		GasLimit: e.committerOpts.GasLimit,
		Context:  ctx, // with RPC timeout
	}

	// Figure out the gas price values
	suggestedGasPrice, err := e.evmProvider.SuggestGasPrice(opts.Context)
	if err != nil {
		metrics.ReportFuncError(e.svcTags)
		return common.Hash{}, errors.Errorf("failed to suggest gas price: %v", err)
	}

	// Suggested gas price is not accurate. Increment by multiplying with gasprice adjustment factor
	incrementedPrice := big.NewFloat(0).Mul(new(big.Float).SetInt(suggestedGasPrice), big.NewFloat(e.ethGasPriceAdjustment))

	// set gasprice to incremented gas price.
	gasPrice := new(big.Int)
	incrementedPrice.Int(gasPrice)

	opts.GasPrice = gasPrice

	//The gas price should be less than max gas price
	maxGasPrice := big.NewInt(int64(e.ethMaxGasPrice))
	if opts.GasPrice.Cmp(maxGasPrice) > 0 {
		return common.Hash{}, errors.Errorf("Suggested gas price %v is greater than max gas price %v", opts.GasPrice.Int64(), maxGasPrice.Int64())
	}

	// estimate gas limit
	msg := ethereum.CallMsg{
		From:     opts.From,
		To:       &recipient,
		GasPrice: gasPrice,
		Value:    new(big.Int),
		Data:     txData,
	}

	gasLimit, err := e.evmProvider.EstimateGas(opts.Context, msg)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "failed to estimate gas")
	}

	opts.GasLimit = gasLimit

	resyncNonces := func(from common.Address) error {
		var nonce uint64
		var err error

		done := make(chan struct{})
		go func() {
			nonce, err = e.evmProvider.PendingNonceAt(context.Background(), from)
			if err != nil {
				log.WithError(err).Warningln("unable to acquire nonce")
			} else {
				e.nonceCache.Set(from, int64(nonce))
			}
			close(done)
		}()

		select {
		case <-done:
			return err
		case <-time.After(10 * time.Second):
			return errors.New("timeout while syncing nonce")
		}
	}

	if err := e.nonceCache.Serialize(e.fromAddress, func() (err error) {
		nonce, _ := e.nonceCache.Get(e.fromAddress)
		var resyncUsed bool

		for {
			opts.Nonce = big.NewInt(nonce)
			ctxTimed, cancelFn := context.WithTimeout(ctx, e.committerOpts.RPCTimeout)
			defer cancelFn()
			opts.Context = ctxTimed

			tx := types.NewTransaction(opts.Nonce.Uint64(), recipient, nil, opts.GasLimit, opts.GasPrice, txData)
			// log.Info("e.fromAddress: ", e.fromAddress)
			// log.Info("opts.From: ", opts.From)
			// log.Info("opts.Signer: ", opts.Signer)
			// txDataHex := hexutil.Encode(txData)
			// log.Info("txDataHex: ", txDataHex)
			signedTx, err := opts.Signer(opts.From, tx)
			if err != nil {
				err := errors.Wrap(err, "failed to sign transaction")
				return err
			}

			txHash = signedTx.Hash()

			txHashRet, err := e.evmProvider.SendTransactionWithRet(opts.Context, signedTx)
			if err == nil {
				// override with a real hash from node resp
				txHash = txHashRet
				e.nonceCache.Incr(e.fromAddress)
				return nil
			} else {
				log.WithFields(log.Fields{
					"tx_hash": txHash.Hex(),
				}).WithError(err).Warningln("failed to send tx")
			}

			switch {
			case strings.Contains(err.Error(), "invalid sender"):
				log.Info("err when sending tx", err)
				err := errors.New("failed to sign transaction")
				e.nonceCache.Incr(e.fromAddress)
				return err
			case strings.Contains(err.Error(), "nonce too low"),
				strings.Contains(err.Error(), "nonce too high"),
				strings.Contains(err.Error(), "the tx doesn't have the correct nonce"):
				log.Info("warning when sending tx", err)
				if resyncUsed {
					log.Errorf("nonces synced, but still wrong nonce for %s: %d", e.fromAddress, nonce)
					err = errors.Wrapf(err, "nonce %d mismatch", nonce)
					return err
				}

				nonceCh := make(chan struct{})
				go func() {
					resyncNonces(e.fromAddress)
					close(nonceCh)
				}()

				select {
				case <-nonceCh:
					log.Info("Nonce resynchronized successfully.")
				case <-time.After(10 * time.Second):
					log.Error("Nonce resynchronization timed out.")
					return errors.New("nonce resynchronization timed out")
				}

				resyncUsed = true
				// try again with updated nonce
				nonce, _ = e.nonceCache.Get(e.fromAddress)
				opts.Nonce = big.NewInt(nonce)

				continue

			default:
				log.Info("err when sending tx", err)
				if strings.Contains(err.Error(), "known transaction") {
					// skip one nonce step, try to send again
					nonce := e.nonceCache.Incr(e.fromAddress)
					opts.Nonce = big.NewInt(nonce)
					continue
				}

				if strings.Contains(err.Error(), "VM Exception") {
					// a VM execution consumes gas and nonce is increasing
					e.nonceCache.Incr(e.fromAddress)
					return err
				}

				return err
			}
		}
	}); err != nil {
		metrics.ReportFuncError(e.svcTags)

		log.WithError(err).Errorln("SendTx serialize failed")

		return common.Hash{}, err
	}

	return txHash, nil
}
