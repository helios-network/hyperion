package hyperion

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

func (s *hyperionContract) SendTransactionBatch(
	ctx context.Context,
	currentValset *types.Valset,
	batch *types.OutgoingTxBatch,
	confirms []*types.MsgConfirmBatch,
) (*common.Hash, *big.Int, error) {
	metrics.ReportFuncCall(s.svcTags)
	doneFn := metrics.ReportFuncTiming(s.svcTags)
	defer doneFn()

	// log.Info("batch.Transactions", batch.Transactions)
	log.WithFields(log.Fields{
		"token_contract": batch.TokenContract,
		"batch_nonce":    batch.BatchNonce,
		"txs":            len(batch.Transactions),
		"confirmations":  len(confirms),
	}).Infoln("checking signatures and submitting batch")

	validators, powers, sigV, sigR, sigS, err := checkBatchSigsAndRepack(currentValset, confirms)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		err = errors.Wrap(err, "submit_batch.go confirmations check failed")
		return nil, big.NewInt(0), err
	}

	amounts, destinations, fees := getBatchCheckpointValues(batch)
	currentValsetNonce := new(big.Int).SetUint64(currentValset.Nonce)
	batchNonce := new(big.Int).SetUint64(batch.BatchNonce)
	batchTimeout := new(big.Int).SetUint64(batch.BatchTimeout)

	// Solidity function signature
	// function submitBatch(
	// 		// The validators that approve the batch and new valset
	// 		address[] memory _currentValidators,
	// 		uint256[] memory _currentPowers,
	// 		uint256 _currentValsetNonce,
	//
	// 		// These are arrays of the parts of the validators signatures
	// 		uint8[] memory _v,
	// 		bytes32[] memory _r,
	// 		bytes32[] memory _s,
	//
	// 		// The batch of transactions
	// 		uint256[] memory _amounts,
	// 		address[] memory _destinations,
	// 		uint256[] memory _fees,
	// 		uint256 _batchNonce,
	// 		address _tokenContract
	// )

	currentValsetArs := ValsetArgs{
		Validators:   validators,
		Powers:       powers,
		ValsetNonce:  currentValsetNonce,
		RewardAmount: currentValset.RewardAmount.BigInt(),
		RewardToken:  common.HexToAddress(currentValset.RewardToken),
	}

	// log.Info("currentValsetArs", currentValsetArs)
	// log.Info("sigV", sigV)
	// log.Info("sigR", sigR)
	// log.Info("sigS", sigS)
	// log.Info("amounts", amounts)
	// log.Info("destinations", destinations)
	// log.Info("fees", fees)
	// log.Info("batchNonce", batchNonce)
	// log.Info("batchTimeout", batchTimeout)
	// log.Info("batch.TokenContract", common.HexToAddress(batch.TokenContract))

	txData, err := hyperionABI.Pack("submitBatch",
		currentValsetArs,
		sigV, sigR, sigS,
		amounts,
		destinations,
		fees,
		batchNonce,
		common.HexToAddress(batch.TokenContract),
		batchTimeout,
	)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		log.WithError(err).Errorln("ABI Pack (Hyperion submitBatch) method")
		return nil, big.NewInt(0), err
	}

	// Checking in pending txs(mempool) if tx with same input is already submitted
	if s.pendingTxInputList.IsPendingTxInput(txData, s.pendingTxWaitDuration) {
		return nil, big.NewInt(0), errors.New("Transaction with same batch input data is already present in mempool")
	}

	txHash, cost, err := s.SendTx(ctx, s.hyperionAddress, txData)
	if err != nil {
		metrics.ReportFuncError(s.svcTags)
		return nil, big.NewInt(0), err
	}

	//     let before_nonce = get_tx_batch_nonce(
	//         hyperion_contract_address,
	//         batch.token_contract,
	//         eth_address,
	//         &web3,
	//     )
	//     .await?;
	//     if before_nonce >= new_batch_nonce {
	//         info!(
	//             "Someone else updated the batch to {}, exiting early",
	//             before_nonce
	//         );
	//         return Ok(());
	//     }

	//     let tx = web3
	//         .send_transaction(
	//             hyperion_contract_address,
	//             payload,
	//             0u32.into(),
	//             eth_address,
	//             our_eth_key,
	//             vec![SendTxOption::GasLimit(1_000_000u32.into())],
	//         )
	//         .await?;
	//     info!("Sent batch update with txid {:#066x}", tx);

	//     // TODO this segment of code works around the race condition for submitting batches mostly
	//     // by not caring if our own submission reverts and only checking if the valset has been updated
	//     // period not if our update succeeded in particular. This will require some further consideration
	//     // in the future as many independent relayers racing to update the same thing will hopefully
	//     // be the common case.
	//     web3.wait_for_transaction(tx, timeout, None).await?;

	//     let last_nonce = get_tx_batch_nonce(
	//         hyperion_contract_address,
	//         batch.token_contract,
	//         eth_address,
	//         &web3,
	//     )
	//     .await?;
	//     if last_nonce != new_batch_nonce {
	//         error!(
	//             "Current nonce is {} expected to update to nonce {}",
	//             last_nonce, new_batch_nonce
	//         );
	//     } else {
	//         info!("Successfully updated Batch with new Nonce {:?}", last_nonce);
	//     }
	//     Ok(())

	return &txHash, cost, nil
}

func getBatchCheckpointValues(batch *types.OutgoingTxBatch) (amounts []*big.Int, destinations []common.Address, fees []*big.Int) {
	amounts = make([]*big.Int, len(batch.Transactions))
	destinations = make([]common.Address, len(batch.Transactions))
	fees = make([]*big.Int, len(batch.Transactions))

	for i, tx := range batch.Transactions {
		amounts[i] = tx.Token.Amount.BigInt()
		destinations[i] = common.HexToAddress(tx.DestAddress)
		fees[i] = tx.Fee.Amount.BigInt()
	}

	return
}

func checkBatchSigsAndRepack(
	valset *types.Valset,
	confirms []*types.MsgConfirmBatch,
) (
	validators []common.Address,
	powers []*big.Int,
	v []uint8,
	r []common.Hash,
	s []common.Hash,
	err error,
) {
	if len(confirms) == 0 {
		err = errors.New("no signatures in batch confirmation")
		return
	}

	signerToSig := make(map[string]*types.MsgConfirmBatch, len(confirms))
	for _, sig := range confirms {
		signerToSig[sig.EthSigner] = sig
	}

	powerOfGoodSigs := new(big.Int)

	for _, m := range valset.Members {
		mPower := big.NewInt(0).SetUint64(m.Power)
		if sig, ok := signerToSig[m.EthereumAddress]; ok && sig.EthSigner == m.EthereumAddress {
			powerOfGoodSigs.Add(powerOfGoodSigs, mPower)

			validators = append(validators, common.HexToAddress(m.EthereumAddress))
			powers = append(powers, mPower)

			sigV, sigR, sigS := sigToVRS(sig.Signature)
			v = append(v, sigV)
			r = append(r, sigR)
			s = append(s, sigS)
		} else {
			validators = append(validators, common.HexToAddress(m.EthereumAddress))
			powers = append(powers, mPower)
			v = append(v, 0)
			r = append(r, [32]byte{})
			s = append(s, [32]byte{})
		}
	}
	// Vérifier si une seule signature est présente et que le signataire est l'administrateur
	if len(validators) == 1 {
		return
	}

	if hyperionPowerToPercent(powerOfGoodSigs) < 66 {
		err = errors.New(fmt.Sprintf("insufficient voting power %f", hyperionPowerToPercent(powerOfGoodSigs)))
		return
	}

	return
}
