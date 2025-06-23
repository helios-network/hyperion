package orchestrator

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/util"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

const (
	defaultRelayerLoopDur = 1 * time.Minute
)

func (s *Orchestrator) runRelayer(ctx context.Context) error {
	if noRelay := !s.cfg.RelayValsets && !s.cfg.RelayBatches && !s.cfg.RelayExternalDatas; noRelay {
		return nil
	}

	r := relayer{
		Orchestrator: s,
		logEnabled:   strings.Contains(s.cfg.EnabledLogs, "relayer"),
	}
	s.logger.WithFields(log.Fields{"loop_duration": defaultRelayerLoopDur.String(), "relay_token_batches": r.cfg.RelayBatches, "relay_validator_sets": s.cfg.RelayValsets}).Debugln("starting Relayer...")

	return loops.RunLoop(ctx, s.ethereum, defaultRelayerLoopDur, func() error {
		// if !s.isRegistered() {
		// 	r.Log().Infoln("Orchestrator not registered, skipping...")
		// 	return nil
		// }
		err := r.relay(ctx)
		return err
	})
}

type relayer struct {
	*Orchestrator
	logEnabled bool
}

func (l *relayer) Log() log.Logger {
	return l.logger.WithField("loop", "Relayer")
}

func (l *relayer) relay(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	var pg loops.ParanoidGroup

	if l.logEnabled {
		l.Log().Info("relaying getLatestEthValset")
	}

	if l.cfg.RelayExternalDatas {
		pg.Go(func() error {
			return l.retry(ctx, func() error {
				return l.relayExternalData(ctx)
			})
		})
	}

	ethValset, err := l.getLatestEthValset(ctx)
	if err != nil {
		return err
	}

	if l.logEnabled {
		l.Log().Info("relaying makeCheckpoint")
	}

	heliosCheckpoint, err := l.makeCheckpoint(ctx, ethValset)
	if err != nil {
		return err
	}

	if l.logEnabled {
		l.Log().Info("relaying getLastValsetCheckpoint")
	}

	ethCheckpoint, err := l.ethereum.GetLastValsetCheckpoint(ctx)
	if err != nil {
		return err
	}

	if l.logEnabled {
		l.Log().Info("relaying getLastValsetCheckpoint done")
	}

	if l.logEnabled {
		l.Log().WithFields(log.Fields{
			"Helios": heliosCheckpoint.Hex(),
			"Eth":    ethCheckpoint.Hex(),
			"Synced": heliosCheckpoint.Hex() == ethCheckpoint.Hex(),
		}).Infoln("Relayer: checkpoints")
	}

	if heliosCheckpoint.Hex() != ethCheckpoint.Hex() {
		if l.logEnabled {
			l.Log().Infoln("relayer: checkpoint not synced yet waiting (rpc should be untrustable) ...")
		}

		// json, err := json.Marshal(ethValset)
		// if err == nil {
		// 	os.WriteFile("valset-error.json", json, 0644)
		// }

		return nil
	} else {
		if l.logEnabled {
			l.Log().Info("valset is synced")
		}
	}

	// write valset to file
	// json, err := json.Marshal(ethValset)
	// if err == nil {
	// 	os.WriteFile("valset.json", json, 0644)
	// }

	if l.cfg.RelayValsets {
		pg.Go(func() error {
			return l.retry(ctx, func() error {
				return l.relayValset(ctx, ethValset)
			})
		})
	}

	if l.cfg.RelayBatches {
		pg.Go(func() error {
			return l.retry(ctx, func() error {
				for i := 0; i < 5; i++ { // do 5 batch in same time if possible
					relayed, err := l.relayTokenBatch(ctx, ethValset)
					if err != nil {
						return err
					}
					if relayed {
						time.Sleep(15 * time.Second)
						continue
					}
				}
				return nil
			})
		})
	}

	if pg.Initialized() {
		if err := pg.Wait(); err != nil {
			return err
		}
	}

	return nil

}

func (l *relayer) encodeData(
	hyperionId common.Hash,
	valsetNonce uint64,
	validators []string,
	powers []uint64,
	rewardAmount *big.Int,
	rewardToken string,
) (common.Hash, error) {

	methodName := [32]byte{}
	copy(methodName[:], []byte("checkpoint"))

	// Conversion des validators en common.Address
	validatorsArr := make([]common.Address, len(validators))
	for i, v := range validators {
		validatorsArr[i] = common.HexToAddress(v)
	}

	// Conversion des powers en []*big.Int
	powersArr := make([]*big.Int, len(powers))
	for i, power := range powers {
		powersArr[i] = new(big.Int).SetUint64(power)
	}

	bytes32Ty, _ := abi.NewType("bytes32", "", nil)
	uint256Ty, _ := abi.NewType("uint256", "", nil)
	addressTy, _ := abi.NewType("address", "", nil)
	addressArrayTy, _ := abi.NewType("address[]", "", nil)
	uint256ArrayTy, _ := abi.NewType("uint256[]", "", nil)

	// Préparer les arguments de façon identique à abi.encode() côté Solidity
	arguments := abi.Arguments{
		{Type: bytes32Ty},      // hyperionId
		{Type: bytes32Ty},      // methodName ("checkpoint")
		{Type: uint256Ty},      // valsetNonce
		{Type: addressArrayTy}, // validators
		{Type: uint256ArrayTy}, // powers
		{Type: uint256Ty},      // rewardAmount
		{Type: addressTy},      // rewardToken
	}

	encodedBytes, err := arguments.Pack(
		hyperionId,
		methodName,
		new(big.Int).SetUint64(valsetNonce),
		validatorsArr,
		powersArr,
		rewardAmount,
		common.HexToAddress(rewardToken),
	)

	if err != nil {
		return common.Hash{}, err
	}

	// Enfin, réaliser keccak256 sur les données encodées
	checkpoint := crypto.Keccak256Hash(encodedBytes)

	return checkpoint, nil
}

func (l *relayer) makeCheckpoint(ctx context.Context, valset *hyperiontypes.Valset) (*common.Hash, error) {
	/** function makeCheckpoint(
	      ValsetArgs memory _valsetArgs,
	      bytes32 _hyperionId
	  ) private pure returns (bytes32) {
	      // bytes32 encoding of the string "checkpoint"
	      bytes32 methodName = 0x636865636b706f696e7400000000000000000000000000000000000000000000;

	      bytes32 checkpoint = keccak256(
	          abi.encode(
	              _hyperionId,
	              methodName,
	              _valsetArgs.valsetNonce,
	              _valsetArgs.validators,
	              _valsetArgs.powers,
	              _valsetArgs.rewardAmount,
	              _valsetArgs.rewardToken
	          )
	      );
	      return checkpoint;
	  }
	*/
	validators := []string{}
	powers := []uint64{}

	for _, validator := range valset.Members {
		validators = append(validators, validator.EthereumAddress)
		powers = append(powers, validator.Power)
	}

	hyperionIDHash, err := l.ethereum.GetHyperionID(ctx)
	if err != nil {
		l.Log().WithError(err).Errorln("unable to query hyperion ID from contract")
		return nil, err
	}
	// Encoder les données
	checkpoint, err := l.encodeData(
		hyperionIDHash,
		valset.Nonce,
		validators,
		powers,
		valset.RewardAmount.BigInt(),
		valset.RewardToken,
	)
	if err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func (l *relayer) getLatestEthValset(ctx context.Context) (*hyperiontypes.Valset, error) {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	var latestEthValset *hyperiontypes.Valset
	fn := func() error {
		vs, err := l.findLatestValsetOnEth(ctx)
		if err != nil {
			l.Log().Infoln("findLatestValsetOnEth - 8")
			if strings.Contains(err.Error(), "failed to get") || strings.Contains(err.Error(), "attempting to unmarshall") || strings.Contains(err.Error(), "pruned") {
				l.ethereum.RemoveLastUsedRpc()
				vs, err = l.findLatestValsetOnEth(ctx)
				if err != nil {
					return err
				}
				latestEthValset = vs
				return nil
			}
			return err
		}
		latestEthValset = vs
		return nil
	}

	if err := l.retry(ctx, fn); err != nil {
		return nil, err
	}

	return latestEthValset, nil
}

func (l *relayer) relayValset(ctx context.Context, latestEthValset *hyperiontypes.Valset) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	latestHeliosValsets, err := l.helios.LatestValsets(ctx, l.cfg.HyperionId)
	if err != nil {
		return errors.Wrap(err, "failed to get latest validator set from Helios")
	}

	var (
		latestConfirmedValset *hyperiontypes.Valset
		confirmations         []*hyperiontypes.MsgValsetConfirm
	)

	for _, set := range latestHeliosValsets {
		sigs, err := l.helios.AllValsetConfirms(ctx, l.cfg.HyperionId, set.Nonce)
		if err != nil {
			return errors.Wrapf(err, "failed to get validator set confirmations for nonce %d", set.Nonce)
		}
		if len(sigs) == 0 {
			continue
		}
		confirmations = sigs
		latestConfirmedValset = set
		break
	}

	if latestConfirmedValset == nil {
		l.Log().Infoln("no validator set to relay")
		return nil
	}

	shouldRelay := l.shouldRelayValset(ctx, latestConfirmedValset)

	if l.logEnabled {
		l.Log().WithFields(log.Fields{"eth_nonce": latestEthValset.Nonce, "hls_nonce": latestConfirmedValset.Nonce, "sigs": len(confirmations), "should_relay": shouldRelay, "synched": latestEthValset.Nonce == latestConfirmedValset.Nonce}).Infoln("relayer try relay Valset")
	}

	if !shouldRelay {
		return nil
	}

	txHash, cost, err := l.ethereum.SendEthValsetUpdate(ctx, latestEthValset, latestConfirmedValset, confirmations)
	if err != nil {
		return err
	}

	storage.UpdateFeesFile(latestEthValset.RewardAmount.BigInt(), latestEthValset.RewardToken, cost, txHash.Hex(), latestEthValset.Height, l.cfg.ChainId, "VALSET")

	l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent validator set update to Ethereum")

	return nil
}

func (l *relayer) shouldRelayValset(ctx context.Context, vs *hyperiontypes.Valset) bool {
	latestEthereumValsetNonce, err := l.ethereum.GetValsetNonce(ctx)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to get latest valset nonce from " + l.cfg.ChainName)
		if strings.Contains(err.Error(), "attempting to unmarshall") { // if error is about unmarshalling, remove last used rpc
			l.ethereum.RemoveLastUsedRpc()
		}
		return false
	}

	// Check if other validators already updated the valset
	if vs.Nonce <= latestEthereumValsetNonce.Uint64() {
		l.Log().WithFields(log.Fields{"eth_nonce": latestEthereumValsetNonce, "helios_nonce": vs.Nonce}).Infoln("validator set already updated on " + l.cfg.ChainName)
		return false
	}

	// Check custom time delay offset for determine if we should relay the valset on chain respecting the offset
	block, err := l.helios.GetBlock(ctx, int64(vs.Height))
	if err != nil {
		latestBlockHeight, err := l.helios.GetLatestBlockHeight(ctx)
		if err != nil {
			l.Log().WithError(err).Warningln("unable to get latest block from Helios")
			return false
		}
		block, err = l.helios.GetBlock(ctx, int64(latestBlockHeight))
		if err != nil {
			l.Log().WithError(err).Warningln("unable to get latest block from Helios (tryed block:", int64(vs.Height), ")")
			return false
		}
		if block != nil && latestBlockHeight > int64(vs.Height)+1000 { // should be sufficient to avoid race condition
			l.Log().WithFields(log.Fields{"helios_nonce": vs.Nonce, "eth_nonce": latestEthereumValsetNonce.Uint64()}).Debugln("new valset update")
			return true
		}
		l.Log().WithError(err).Warningln("unable to get latest block of valset from Helios")
		return false
	}

	if timeElapsed := time.Since(block.Block.Time); timeElapsed <= l.cfg.RelayValsetOffsetDur {
		timeRemaining := time.Duration(int64(l.cfg.RelayValsetOffsetDur) - int64(timeElapsed))
		l.Log().WithField("time_remaining", timeRemaining.String()).Infoln("valset relay offset not reached yet")
		return false
	}

	l.Log().WithFields(log.Fields{"helios_nonce": vs.Nonce, "eth_nonce": latestEthereumValsetNonce.Uint64()}).Debugln("new valset update")

	return true
}

func (l *relayer) relayTokenBatch(ctx context.Context, latestEthValset *hyperiontypes.Valset) (bool, error) {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	batches, err := l.helios.LatestTransactionBatches(ctx, l.cfg.HyperionId)
	if l.logEnabled {
		numberOfDifferentTokenBatches := 0
		mapTokenContract := make(map[string]bool)
		for _, batch := range batches {
			if _, ok := mapTokenContract[batch.TokenContract]; !ok {
				mapTokenContract[batch.TokenContract] = true
				numberOfDifferentTokenBatches++
			}
		}
		l.Log().Info("batches: ", len(batches), "numberOfDifferentTokenBatches: ", numberOfDifferentTokenBatches)
	}
	if err != nil {
		l.Log().Info("failed to get latest transaction batches", err)
		return false, err
	}

	latestEthHeight, err := l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		l.Log().Info("failed to get latest "+l.cfg.ChainName+" height", err)
		return false, err
	}

	sort.Slice(batches, func(i, j int) bool {
		return batches[i].BatchNonce < batches[j].BatchNonce
	})

	mapBatchPerTokenContract := make(map[string][]*hyperiontypes.OutgoingTxBatch, 0)

	for _, batch := range batches {
		if _, ok := mapBatchPerTokenContract[batch.TokenContract]; !ok {
			mapBatchPerTokenContract[batch.TokenContract] = make([]*hyperiontypes.OutgoingTxBatch, 0)
		}
		mapBatchPerTokenContract[batch.TokenContract] = append(mapBatchPerTokenContract[batch.TokenContract], batch)
	}

	hasPushedABatch := false
	errorsBatchs := make([]string, 0)
	for _, batches := range mapBatchPerTokenContract {
		err := l.processBatch(ctx, batches, latestEthValset, latestEthHeight)
		if err != nil {
			errorsBatchs = append(errorsBatchs, err.Error())
		} else {
			hasPushedABatch = true
		}
	}

	if !hasPushedABatch {
		return false, errors.New(strings.Join(errorsBatchs, "\n"))
	}

	return true, nil
}

func (l *relayer) SignBatch(ctx context.Context, batch *hyperiontypes.OutgoingTxBatch) error {
	hyperionIdHash := common.HexToHash(strconv.FormatUint(l.cfg.HyperionId, 16))

	if err := l.retry(ctx, func() error {
		return l.helios.SendBatchConfirm(ctx,
			l.cfg.HyperionId,
			l.cfg.EthereumAddr,
			hyperionIdHash,
			batch,
		)
	}); err != nil {
		return err
	}
	return nil
}

type BatchToSign struct {
	Batch  *hyperiontypes.OutgoingTxBatch
	Signed bool
}

func (l *relayer) processBatch(ctx context.Context, batches []*hyperiontypes.OutgoingTxBatch, latestEthValset *hyperiontypes.Valset, latestEthHeight *gethtypes.Header) error {

	var (
		oldestConfirmedBatch *hyperiontypes.OutgoingTxBatch
		confirmations        []*hyperiontypes.MsgConfirmBatch
	)

	mapBatchNonce := make(map[string]uint64, 0)

	sort.Slice(batches, func(i, j int) bool {
		return batches[i].BatchNonce < batches[j].BatchNonce
	})

	for _, batch := range batches {
		l.Log().Info("", batch.BatchNonce)
	}

	batchsTimeout := make([]*hyperiontypes.OutgoingTxBatch, 0)
	batchsNotConfirmed := make([]BatchToSign, 0)
	batchsWithExpiredNonce := make([]*hyperiontypes.OutgoingTxBatch, 0)

	for _, batch := range batches {
		// l.Log().Info("batch details: ", batch)

		if batch.HyperionId != l.cfg.HyperionId {
			continue
		}

		if batch.BatchTimeout <= latestEthHeight.Number.Uint64() {
			l.Log().WithFields(log.Fields{"batch_nonce": batch.BatchNonce, "batch_timeout_height": batch.BatchTimeout, "latest_eth_height": latestEthHeight.Number.Uint64()}).Debugln("skipping timed out batch")
			batchsTimeout = append(batchsTimeout, batch)
			continue
		}

		sigs, err := l.helios.TransactionBatchSignatures(ctx, l.cfg.HyperionId, batch.BatchNonce, gethcommon.HexToAddress(batch.TokenContract))
		// l.Log().Info("sigs", sigs)
		if err != nil {
			return err
		}

		iHaveSigned := false
		for _, sig := range sigs {
			if sig.EthSigner == l.cfg.EthereumAddr.Hex() {
				iHaveSigned = true
				break
			}
		}

		if len(sigs) == 0 {
			batchsNotConfirmed = append(batchsNotConfirmed, BatchToSign{Batch: batch, Signed: iHaveSigned})
			continue
		}

		if _, ok := mapBatchNonce[batch.TokenContract]; ok {
			if batch.BatchNonce <= mapBatchNonce[batch.TokenContract] {
				batchsWithExpiredNonce = append(batchsWithExpiredNonce, batch)
				continue
			}
		} else {
			latestBatchNonce, err := l.ethereum.GetTxBatchNonce(ctx, gethcommon.HexToAddress(batch.TokenContract))
			if err != nil {
				return err
			}
			mapBatchNonce[batch.TokenContract] = latestBatchNonce.Uint64()
			if batch.BatchNonce <= mapBatchNonce[batch.TokenContract] {
				batchsWithExpiredNonce = append(batchsWithExpiredNonce, batch)
				continue
			}
		}

		oldestConfirmedBatch = batch
		confirmations = sigs
		if oldestConfirmedBatch != nil {
			break
		}
	}

	if oldestConfirmedBatch == nil {
		if l.logEnabled {
			l.Log().Infoln("no token batch to relay - batchsTimeout: ", len(batchsTimeout), " - batchsNotConfirmed: ", len(batchsNotConfirmed), " - batchsWithExpiredNonce: ", len(batchsWithExpiredNonce))
		}
		if len(batchsNotConfirmed) > 0 {
			for _, batch := range batchsNotConfirmed {
				if !batch.Signed {
					l.Log().Infoln("signing batch - batch: ", batch.Batch.BatchNonce, " - tokenContract: ", batch.Batch.TokenContract)
					l.SignBatch(ctx, batch.Batch)
				}
			}
		}
		return nil
	} else {
		currentNonceIsWellSuperiorToExpiredNonces := false
		for _, batch := range batchsWithExpiredNonce {
			if batch.BatchNonce > oldestConfirmedBatch.BatchNonce {
				currentNonceIsWellSuperiorToExpiredNonces = true
				break
			}
		}
		l.Log().Infoln("relaying batch - batchsTimeout: ", len(batchsTimeout), " - batchsNotConfirmed: ", len(batchsNotConfirmed), " - batchsWithExpiredNonce: ", len(batchsWithExpiredNonce), " - currentNonceIsWellSuperiorToExpiredNonces: ", currentNonceIsWellSuperiorToExpiredNonces)
	}
	// l.Log().Info("oldestConfirmedBatch", oldestConfirmedBatch)

	// l.Log().Info("shouldRelayBatch", l.shouldRelayBatch(ctx, oldestConfirmedBatch))
	// if !l.shouldRelayBatch(ctx, oldestConfirmedBatch) {
	// 	return nil
	// }

	// if l.logEnabled {
	// l.Log().Infoln("latestEthValset", latestEthValset)
	l.Log().Infoln("Nonce", oldestConfirmedBatch.BatchNonce, "ContractNonce", mapBatchNonce[oldestConfirmedBatch.TokenContract])
	// l.Log().Infoln("confirmations", confirmations)
	// }

	txHash, cost, err := l.ethereum.SendTransactionBatch(ctx, latestEthValset, oldestConfirmedBatch, confirmations)
	if err != nil {
		// Returning an error here triggers retries which don't help much except risk a binary crash
		// Better to warn the user and try again in the next loop interval
		l.Log().WithError(err).Warningln("failed to send outgoing tx batch")
		return err
	}

	feesTaken := oldestConfirmedBatch.GetFees().BigInt()
	// TODO: save fees taken and expenses
	storage.UpdateFeesFile(feesTaken, oldestConfirmedBatch.TokenContract, cost, txHash.Hex(), latestEthHeight.Number.Uint64(), l.cfg.ChainId, "BATCH")
	///////

	l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent outgoing tx batch to " + l.cfg.ChainName)
	return nil
}

// func (l *relayer) testRelayExternalData(ctx context.Context) {

// 	latestEthHeight, err := l.ethereum.GetHeaderByNumber(ctx, nil)
// 	if err != nil {
// 		l.Log().Info("failed to get latest "+l.cfg.ChainName+" height", err)
// 		return
// 	}
// 	//0x3931ab520000000000000000000000000000000000000000000000000000000000000000
// 	data, err := hex.DecodeString("3931ab520000000000000000000000000000000000000000000000000000000000000000")
// 	if err != nil {
// 		l.Log().Info("failed to decode abi call hex", err)
// 		return
// 	}
// 	callData, callErr, rpcUsed, _ := l.ethereum.ExecuteExternalDataTx(ctx, gethcommon.HexToAddress("0x61F2AB7B0C0E10E18a3ed1C3bC7958540374A8DC"), data, latestEthHeight.Number)
// 	l.Log().Info("callData", callData, "callErr", callErr, "rpcUsed", rpcUsed)
// }

func (l *relayer) selectBestClaimFromListOfClaims(claims []*types.MsgExternalDataClaim) *types.MsgExternalDataClaim {
	// If no claims, return nil
	if len(claims) == 0 {
		return nil
	}

	// If only one claim, return it
	if len(claims) == 1 {
		return claims[0]
	}

	// Map to store frequency of each combination
	frequencies := make(map[string]int)
	claimsByKey := make(map[string]*types.MsgExternalDataClaim)

	// Count frequencies of each unique combination
	for _, claim := range claims {
		// Create a unique key combining the relevant fields
		key := fmt.Sprintf("%d|%s|%s",
			claim.TxNonce,
			claim.CallDataResult,
			claim.CallDataResultError,
		)

		frequencies[key]++
		claimsByKey[key] = claim
	}

	// Find the key with highest frequency
	var maxFreq int
	var bestKey string
	for key, freq := range frequencies {
		if freq > maxFreq {
			maxFreq = freq
			bestKey = key
		}
	}

	// Return the claim corresponding to the most frequent combination
	return claimsByKey[bestKey]
}

func (l *relayer) relayExternalData(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	txs, err := l.helios.LatestTransactionExternalCallDataTxs(ctx, l.cfg.HyperionId)
	if l.logEnabled {
		l.Log().Info("txs: ", txs)
	}
	if err != nil {
		l.Log().Info("failed to get latest transaction external call data txs", err)
		return err
	}

	latestEthHeight, err := l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		l.Log().Info("failed to get latest "+l.cfg.ChainName+" height", err)
		return err
	}

	for _, tx := range txs {
		l.Log().Info("tx details: ", tx)

		if tx.HyperionId != l.cfg.HyperionId {
			continue
		}

		targetHeight := latestEthHeight.Number.Uint64()

		if slices.Contains(tx.Votes, l.ethereum.FromAddress().Hex()) {
			l.Log().Info("skipping already claimed tx", tx.Id)
			continue
		}

		bestClaim := l.selectBestClaimFromListOfClaims(tx.Claims)

		if bestClaim != nil {
			targetHeight = bestClaim.BlockHeight
		}

		data, err := hex.DecodeString(strings.TrimPrefix(tx.AbiCallHex, "0x"))
		if err != nil {
			l.Log().Info("failed to decode abi call hex", err, "tx_id", tx.Id)
			continue
		}
		callData, callErr, rpcUsed, err := l.ethereum.ExecuteExternalDataTx(ctx, gethcommon.HexToAddress(tx.ExternalContractAddress), data, big.NewInt(int64(targetHeight)))
		l.Log().Info("callData", callData, "callErr", callErr)

		if err != nil {
			l.Log().Info("failed to execute external data tx with rpc", err, "rpcUsed", rpcUsed)
			continue
		}

		_, err = l.helios.SendExternalDataClaim(ctx, l.cfg.HyperionId, tx.Nonce, latestEthHeight.Number.Uint64(), tx.ExternalContractAddress, callData, callErr, rpcUsed)

		if err != nil {
			l.Log().Info("failed to send external data claim", err)
			continue
		}
	}

	// TODO: get external data batch from helios (maybe edit batch_creator for build special batch who contains only external data)
	// TODO: format abi from tx information then call external data on ethereum
	// TODO: send special claimData to helios

	return nil
}

// FindLatestValset finds the latest valset on the Hyperion contract by looking back through the event
// history and finding the most recent ValsetUpdatedEvent. Most of the time this will be very fast
// as the latest update will be in recent blockchain history and the search moves from the present
// backwards in time. In the case that the validator set has not been updated for a very long time
// this will take longer.
func (l *relayer) findLatestValsetOnEth(ctx context.Context) (*hyperiontypes.Valset, error) {

	lastValsetUpdatedEventHeight, err := l.ethereum.GetLastValsetUpdatedEventHeight(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get last valset updated event")
	}

	latestEthereumValsetNonce, err := l.ethereum.GetValsetNonce(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "attempting to unmarshall") { // if error is about unmarshalling, remove last used rpc
			l.ethereum.RemoveLastUsedRpc()
		}
		return nil, errors.Wrap(err, "failed to get latest valset nonce")
	}

	cosmosValset, err := l.helios.ValsetAt(ctx, l.cfg.HyperionId, latestEthereumValsetNonce.Uint64())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Helios valset")
	}

	if lastValsetUpdatedEventHeight.Uint64() > 0 {
		valsetUpdatedEvents, err := l.ethereum.GetValsetUpdatedEventsAtSpecificBlock(lastValsetUpdatedEventHeight.Uint64())
		if err != nil {
			return nil, errors.Wrap(err, "failed to filter past ValsetUpdated events")
		}

		// by default the lowest found valset goes first, we want the highest
		//
		// TODO(xlab): this follows the original impl, but sort might be skipped there:
		// we could access just the latest element later.
		sort.Sort(sort.Reverse(HyperionValsetUpdatedEvents(valsetUpdatedEvents)))

		if len(valsetUpdatedEvents) == 0 { // return the cosmos valset if no event is found
			return cosmosValset, nil
			// return nil, errors.New("failed to get latest valset (Maybe rpc not answering correctly?)")
		}

		// we take only the first event if we find any at all.
		event := valsetUpdatedEvents[0]

		if l.logEnabled {
			l.Log().Info("found valset at block: ", event.Raw.BlockNumber, " with nonce: ", event.NewValsetNonce.Uint64())
		}

		valset := &hyperiontypes.Valset{
			Nonce:        event.NewValsetNonce.Uint64(),
			Members:      make([]*hyperiontypes.BridgeValidator, 0, len(event.Powers)),
			RewardAmount: sdkmath.NewIntFromBigInt(event.RewardAmount),
			RewardToken:  event.RewardToken.Hex(),
		}

		for idx, p := range event.Powers {
			valset.Members = append(valset.Members, &hyperiontypes.BridgeValidator{
				Power:           p.Uint64(),
				EthereumAddress: event.Validators[idx].Hex(),
			})
		}

		if l.logEnabled {
			checkIfValsetsDiffer(cosmosValset, valset)
		}

		return valset, nil

	}

	return nil, ErrNotFound
}

var ErrNotFound = errors.New("not found")

type HyperionValsetUpdatedEvents []*hyperionevents.HyperionValsetUpdatedEvent

func (a HyperionValsetUpdatedEvents) Len() int { return len(a) }
func (a HyperionValsetUpdatedEvents) Less(i, j int) bool {
	return a[i].NewValsetNonce.Cmp(a[j].NewValsetNonce) < 0
}
func (a HyperionValsetUpdatedEvents) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// This function exists to provide a warning if Cosmos and Ethereum have different validator sets
// for a given nonce. In the mundane version of this warning the validator sets disagree on sorting order
// which can happen if some relayer uses an unstable sort, or in a case of a mild griefing attack.
// The Hyperion contract validates signatures in order of highest to lowest power. That way it can exit
// the loop early once a vote has enough power, if a relayer where to submit things in the reverse order
// they could grief users of the contract into paying more in gas.
// The other (and far worse) way a disagreement here could occur is if validators are colluding to steal
// funds from the Hyperion contract and have submitted a hijacking update. If slashing for off Cosmos chain
// Ethereum signatures is implemented you would put that handler here.
func checkIfValsetsDiffer(cosmosValset, ethereumValset *hyperiontypes.Valset) {
	if cosmosValset == nil && ethereumValset.Nonce == 0 {
		// bootstrapping case
		return
	} else if cosmosValset == nil {
		log.WithField(
			"eth_valset_nonce",
			ethereumValset.Nonce,
		).Errorln("Cosmos does not have a valset for nonce from Ethereum chain. Maybe not synced yet?")
		return
	}

	if cosmosValset.Nonce != ethereumValset.Nonce {
		log.WithFields(log.Fields{
			"cosmos_valset_nonce": cosmosValset.Nonce,
			"eth_valset_nonce":    ethereumValset.Nonce,
		}).Errorln("Cosmos does have a wrong valset nonce, differs from Ethereum chain. Possible bridge hijacking!")
		return
	}

	if len(cosmosValset.Members) != len(ethereumValset.Members) {
		log.WithFields(log.Fields{
			"cosmos_valset": len(cosmosValset.Members),
			"eth_valset":    len(ethereumValset.Members),
		}).Errorln("Cosmos and Ethereum Valsets have different length. Possible bridge hijacking!")
		return
	}

	BridgeValidators(cosmosValset.Members).Sort()
	BridgeValidators(ethereumValset.Members).Sort()

	for idx, member := range cosmosValset.Members {
		if ethereumValset.Members[idx].EthereumAddress != member.EthereumAddress {
			log.Errorln("Valsets are different, a sorting error?")
		}
		if ethereumValset.Members[idx].Power != member.Power {
			log.Errorln("Valsets are different, a sorting error?")
		}
	}
}

type BridgeValidators []*hyperiontypes.BridgeValidator

// Sort sorts the validators by power
func (b BridgeValidators) Sort() {
	sort.Slice(b, func(i, j int) bool {
		if b[i].Power == b[j].Power {
			// Secondary sort on ethereum address in case powers are equal
			return util.EthAddrLessThan(b[i].EthereumAddress, b[j].EthereumAddress)
		}
		return b[i].Power > b[j].Power
	})
}

// HasDuplicates returns true if there are duplicates in the set
func (b BridgeValidators) HasDuplicates() bool {
	m := make(map[string]struct{}, len(b))
	for i := range b {
		m[b[i].EthereumAddress] = struct{}{}
	}
	return len(m) != len(b)
}

// GetPowers returns only the power values for all members
func (b BridgeValidators) GetPowers() []uint64 {
	r := make([]uint64, len(b))
	for i := range b {
		r[i] = b[i].Power
	}
	return r
}
