package orchestrator

import (
	"context"
	"sort"
	"strconv"
	"strings"
	"time"

	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/util"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

const (
	defaultRelayerLoopDur = 30 * time.Second
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

	// Use a custom loop that waits for completion before starting next iteration
	ticker := time.NewTicker(defaultRelayerLoopDur)
	defer ticker.Stop()

	// Run first iteration immediately
	if err := r.relay(ctx); err != nil {
		s.logger.WithError(err).Errorln("relay function returned an error")
	}

	for {
		select {
		case <-ticker.C:
			if s.HyperionState.RelayerStatus == "running" {
				continue
			}

			start := time.Now()
			s.HyperionState.RelayerStatus = "running"
			if err := r.relay(ctx); err != nil {
				s.logger.WithError(err).Errorln("relay function returned an error")
			}
			s.HyperionState.RelayerStatus = "idle"
			s.HyperionState.RelayerLastExecutionFinishedTimestamp = uint64(time.Now().Unix())
			s.HyperionState.RelayerNextExecutionTimestamp = uint64(start.Add(defaultRelayerLoopDur).Unix())
		case <-ctx.Done():
			return nil
		}
	}
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

	valsetManager := l.GetValsetManager()

	if valsetManager == nil || !valsetManager.IsValsetSynced(ctx) {
		l.Log().Infoln("valset is not synced, skipping relay", " synced ", valsetManager.synced, " consideredSynced ", valsetManager.consideredSynced)
		return nil
	}

	l.Log().Info("relaying batches")

	ethValset := valsetManager.GetEthValset(ctx)

	if l.cfg.RelayBatches {
		pg.Go(func() error {
			return l.retry(ctx, func() error {
				_, err := l.relayBatchs(ctx, ethValset)
				if err != nil {
					return err
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

// func (l *relayer) encodeData(
// 	hyperionId common.Hash,
// 	valsetNonce uint64,
// 	validators []string,
// 	powers []uint64,
// 	rewardAmount *big.Int,
// 	rewardToken string,
// ) (common.Hash, error) {

// 	methodName := [32]byte{}
// 	copy(methodName[:], []byte("checkpoint"))

// 	// Conversion des validators en common.Address
// 	validatorsArr := make([]common.Address, len(validators))
// 	for i, v := range validators {
// 		validatorsArr[i] = common.HexToAddress(v)
// 	}

// 	// Conversion des powers en []*big.Int
// 	powersArr := make([]*big.Int, len(powers))
// 	for i, power := range powers {
// 		powersArr[i] = new(big.Int).SetUint64(power)
// 	}

// 	bytes32Ty, _ := abi.NewType("bytes32", "", nil)
// 	uint256Ty, _ := abi.NewType("uint256", "", nil)
// 	addressTy, _ := abi.NewType("address", "", nil)
// 	addressArrayTy, _ := abi.NewType("address[]", "", nil)
// 	uint256ArrayTy, _ := abi.NewType("uint256[]", "", nil)

// 	// Préparer les arguments de façon identique à abi.encode() côté Solidity
// 	arguments := abi.Arguments{
// 		{Type: bytes32Ty},      // hyperionId
// 		{Type: bytes32Ty},      // methodName ("checkpoint")
// 		{Type: uint256Ty},      // valsetNonce
// 		{Type: addressArrayTy}, // validators
// 		{Type: uint256ArrayTy}, // powers
// 		{Type: uint256Ty},      // rewardAmount
// 		{Type: addressTy},      // rewardToken
// 	}

// 	encodedBytes, err := arguments.Pack(
// 		hyperionId,
// 		methodName,
// 		new(big.Int).SetUint64(valsetNonce),
// 		validatorsArr,
// 		powersArr,
// 		rewardAmount,
// 		common.HexToAddress(rewardToken),
// 	)

// 	if err != nil {
// 		return common.Hash{}, err
// 	}

// 	// Enfin, réaliser keccak256 sur les données encodées
// 	checkpoint := crypto.Keccak256Hash(encodedBytes)

// 	return checkpoint, nil
// }

// func (l *relayer) makeCheckpoint(ctx context.Context, valset *hyperiontypes.Valset) (*common.Hash, error) {
// 	/** function makeCheckpoint(
// 	      ValsetArgs memory _valsetArgs,
// 	      bytes32 _hyperionId
// 	  ) private pure returns (bytes32) {
// 	      // bytes32 encoding of the string "checkpoint"
// 	      bytes32 methodName = 0x636865636b706f696e7400000000000000000000000000000000000000000000;

// 	      bytes32 checkpoint = keccak256(
// 	          abi.encode(
// 	              _hyperionId,
// 	              methodName,
// 	              _valsetArgs.valsetNonce,
// 	              _valsetArgs.validators,
// 	              _valsetArgs.powers,
// 	              _valsetArgs.rewardAmount,
// 	              _valsetArgs.rewardToken
// 	          )
// 	      );
// 	      return checkpoint;
// 	  }
// 	*/
// 	validators := []string{}
// 	powers := []uint64{}

// 	for _, validator := range valset.Members {
// 		validators = append(validators, validator.EthereumAddress)
// 		powers = append(powers, validator.Power)
// 	}

// 	hyperionIDHash, err := l.ethereum.GetHyperionID(ctx)
// 	if err != nil {
// 		l.Log().WithError(err).Errorln("unable to query hyperion ID from contract")
// 		return nil, err
// 	}
// 	// Encoder les données
// 	checkpoint, err := l.encodeData(
// 		hyperionIDHash,
// 		valset.Nonce,
// 		validators,
// 		powers,
// 		valset.RewardAmount.BigInt(),
// 		valset.RewardToken,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &checkpoint, nil
// }

// func (l *relayer) getLatestEthValset(ctx context.Context) (*hyperiontypes.Valset, error) {
// 	metrics.ReportFuncCall(l.svcTags)
// 	doneFn := metrics.ReportFuncTiming(l.svcTags)
// 	defer doneFn()

// 	var latestEthValset *hyperiontypes.Valset
// 	fn := func() error {
// 		vs, err := l.findLatestValsetOnEth(ctx)
// 		if err != nil {
// 			l.Log().Infoln("findLatestValsetOnEth - 8")
// 			if strings.Contains(err.Error(), "failed to get") || strings.Contains(err.Error(), "attempting to unmarshall") || strings.Contains(err.Error(), "pruned") {
// 				l.ethereum.RemoveLastUsedRpc()
// 				vs, err = l.findLatestValsetOnEth(ctx)
// 				if err != nil {
// 					return err
// 				}
// 				latestEthValset = vs
// 				return nil
// 			}
// 			return err
// 		}
// 		latestEthValset = vs
// 		return nil
// 	}

// 	if err := l.retry(ctx, fn); err != nil {
// 		return nil, err
// 	}

// 	return latestEthValset, nil
// }

// func (l *relayer) relayValset(ctx context.Context, latestEthValset *hyperiontypes.Valset) error {
// 	metrics.ReportFuncCall(l.svcTags)
// 	doneFn := metrics.ReportFuncTiming(l.svcTags)
// 	defer doneFn()

// 	latestHeliosValsets, err := l.helios.LatestValsets(ctx, l.cfg.HyperionId)
// 	if err != nil {
// 		return errors.Wrap(err, "failed to get latest validator set from Helios")
// 	}

// 	var (
// 		latestConfirmedValset *hyperiontypes.Valset
// 		confirmations         []*hyperiontypes.MsgValsetConfirm
// 	)

// 	for _, set := range latestHeliosValsets {
// 		sigs, err := l.helios.AllValsetConfirms(ctx, l.cfg.HyperionId, set.Nonce)
// 		if err != nil {
// 			return errors.Wrapf(err, "failed to get validator set confirmations for nonce %d", set.Nonce)
// 		}
// 		if len(sigs) == 0 {
// 			continue
// 		}
// 		confirmations = sigs
// 		latestConfirmedValset = set
// 		break
// 	}

// 	if latestConfirmedValset == nil {
// 		l.Log().Infoln("no validator set to relay")
// 		return nil
// 	}

// 	shouldRelay := l.shouldRelayValset(ctx, latestConfirmedValset)

// 	if l.logEnabled {
// 		l.Log().WithFields(log.Fields{"eth_nonce": latestEthValset.Nonce, "hls_nonce": latestConfirmedValset.Nonce, "sigs": len(confirmations), "should_relay": shouldRelay, "synched": latestEthValset.Nonce == latestConfirmedValset.Nonce}).Infoln("relayer try relay Valset")
// 	}

// 	if !shouldRelay {
// 		return nil
// 	}

// 	txHash, cost, err := l.ethereum.SendEthValsetUpdate(ctx, latestEthValset, latestConfirmedValset, confirmations)
// 	if err != nil {
// 		return err
// 	}

// 	storage.UpdateFeesFile(latestEthValset.RewardAmount.BigInt(), latestEthValset.RewardToken, cost, txHash.Hex(), latestEthValset.Height, l.cfg.ChainId, "VALSET")

// 	l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent validator set update to Ethereum")

// 	return nil
// }

// func (l *relayer) shouldRelayValset(ctx context.Context, vs *hyperiontypes.Valset) bool {
// 	latestEthereumValsetNonce, err := l.ethereum.GetValsetNonce(ctx)
// 	if err != nil {
// 		l.Log().WithError(err).Warningln("failed to get latest valset nonce from " + l.cfg.ChainName)
// 		if strings.Contains(err.Error(), "attempting to unmarshall") { // if error is about unmarshalling, remove last used rpc
// 			l.ethereum.RemoveLastUsedRpc()
// 		}
// 		return false
// 	}

// 	// Check if other validators already updated the valset
// 	if vs.Nonce <= latestEthereumValsetNonce.Uint64() {
// 		l.Log().WithFields(log.Fields{"eth_nonce": latestEthereumValsetNonce, "helios_nonce": vs.Nonce}).Infoln("validator set already updated on " + l.cfg.ChainName)
// 		return false
// 	}

// 	// Check custom time delay offset for determine if we should relay the valset on chain respecting the offset
// 	block, err := l.helios.GetBlock(ctx, int64(vs.Height))
// 	if err != nil {
// 		latestBlockHeight, err := l.helios.GetLatestBlockHeight(ctx)
// 		if err != nil {
// 			l.Log().WithError(err).Warningln("unable to get latest block from Helios")
// 			return false
// 		}
// 		block, err = l.helios.GetBlock(ctx, int64(latestBlockHeight))
// 		if err != nil {
// 			l.Log().WithError(err).Warningln("unable to get latest block from Helios (tryed block:", int64(vs.Height), ")")
// 			return false
// 		}
// 		if block != nil && latestBlockHeight > int64(vs.Height)+1000 { // should be sufficient to avoid race condition
// 			l.Log().WithFields(log.Fields{"helios_nonce": vs.Nonce, "eth_nonce": latestEthereumValsetNonce.Uint64()}).Debugln("new valset update")
// 			return true
// 		}
// 		l.Log().WithError(err).Warningln("unable to get latest block of valset from Helios")
// 		return false
// 	}

// 	if timeElapsed := time.Since(block.Block.Time); timeElapsed <= l.cfg.RelayValsetOffsetDur {
// 		timeRemaining := time.Duration(int64(l.cfg.RelayValsetOffsetDur) - int64(timeElapsed))
// 		l.Log().WithField("time_remaining", timeRemaining.String()).Infoln("valset relay offset not reached yet")
// 		return false
// 	}

// 	l.Log().WithFields(log.Fields{"helios_nonce": vs.Nonce, "eth_nonce": latestEthereumValsetNonce.Uint64()}).Debugln("new valset update")

// 	return true
// }

type BatchAndSigs struct {
	Batch *hyperiontypes.OutgoingTxBatch
	Sigs  []*hyperiontypes.MsgConfirmBatch
}

func (l *relayer) relayBatchs(ctx context.Context, latestEthValset *hyperiontypes.Valset) (bool, error) {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	latestEthHeight, err := l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		l.Log().Info("failed to get latest "+l.cfg.ChainName+" height", err)
		return false, err
	}

	batchesInHelios, err := l.helios.LatestTransactionBatches(ctx, l.cfg.HyperionId)

	if err != nil {
		l.Log().Info("failed to get latest transaction batches", err)
		return false, err
	}

	l.HyperionState.BatchCount = len(batchesInHelios)

	if len(batchesInHelios) == 0 {
		l.Log().Infoln("no transaction batches to relay")
		return false, nil
	}

	sort.Slice(batchesInHelios, func(i, j int) bool {
		return batchesInHelios[i].BatchNonce < batchesInHelios[j].BatchNonce
	})

	batchsAbleToRelay := make([]*hyperiontypes.OutgoingTxBatch, 0)

	l.HyperionState.TxCount = 0
	for _, batch := range batchesInHelios {
		l.HyperionState.TxCount += len(batch.Transactions)
		if batch.BatchTimeout > uint64(latestEthHeight.Number.Uint64()+uint64(10)) {
			batchsAbleToRelay = append(batchsAbleToRelay, batch)
		}
	}

	if len(batchsAbleToRelay) == 0 {
		l.Log().Infoln("no transaction batches to relay")
		return false, nil
	}

	mapBatchPerTokenContract := make(map[string][]*hyperiontypes.OutgoingTxBatch, 0)

	for _, batch := range batchsAbleToRelay {
		if _, ok := mapBatchPerTokenContract[batch.TokenContract]; !ok {
			mapBatchPerTokenContract[batch.TokenContract] = make([]*hyperiontypes.OutgoingTxBatch, 0)
		}
		mapBatchPerTokenContract[batch.TokenContract] = append(mapBatchPerTokenContract[batch.TokenContract], batch)
	}

	if len(mapBatchPerTokenContract) == 0 {
		return false, nil
	}

	mapLatestBatchNonce := make(map[string]uint64, 0)
	for tokenContract, _ := range mapBatchPerTokenContract {
		if _, ok := mapLatestBatchNonce[tokenContract]; ok {
			continue
		}
		latestBatchNonce, err := l.ethereum.GetTxBatchNonce(ctx, gethcommon.HexToAddress(tokenContract))
		if err != nil {
			continue
		}
		mapLatestBatchNonce[tokenContract] = latestBatchNonce.Uint64()
	}

	paquetsOfTokenContractBatchs := make(map[string][]*hyperiontypes.OutgoingTxBatch, 0)
	// for each token contract, get the latest batch nonce
	for tokenContract, batches := range mapBatchPerTokenContract {

		if _, ok := mapLatestBatchNonce[tokenContract]; !ok {
			continue
		}
		latestBatchNonce := mapLatestBatchNonce[tokenContract]

		for _, batch := range batches {
			if batch.BatchNonce > latestBatchNonce {
				paquetsOfTokenContractBatchs[tokenContract] = append(paquetsOfTokenContractBatchs[tokenContract], batch)
				mapLatestBatchNonce[tokenContract] = batch.BatchNonce
			}
		}
	}

	grpOfSignedBatchs := make(map[string][]*BatchAndSigs, 0)

	for tokenContrat, batchs := range paquetsOfTokenContractBatchs {

		sort.Slice(batchs, func(i, j int) bool {
			return batchs[i].BatchNonce < batchs[j].BatchNonce
		})

		packSigned := make([]*BatchAndSigs, 0)

		for _, batch := range batchs {
			sigs, err := l.helios.TransactionBatchSignatures(ctx, l.cfg.HyperionId, batch.BatchNonce, gethcommon.HexToAddress(batch.TokenContract))
			if err != nil {
				break // should be relayed later with signature
			}

			iHaveSigned := false
			for _, sig := range sigs {
				if sig.EthSigner == l.cfg.EthereumAddr.Hex() {
					iHaveSigned = true
					break
				}
			}

			symbol, ok := l.CacheSymbol[gethcommon.HexToAddress(batch.TokenContract)]
			if !ok {
				symbol = batch.TokenContract
			}
			l.Log().Info("batch ", batch.BatchNonce, " - tokenContract: ", batch.TokenContract, " - sigs: ", len(sigs), " - iHaveSigned: ", iHaveSigned, " - symbol: ", symbol)

			if iHaveSigned {
				packSigned = append(packSigned, &BatchAndSigs{Batch: batch, Sigs: sigs})
			}
		}

		if len(packSigned) == 0 {
			continue
		}

		grpOfSignedBatchs[tokenContrat] = packSigned

		// sigs, err := l.helios.TransactionBatchSignatures(ctx, l.cfg.HyperionId, batch.BatchNonce, gethcommon.HexToAddress(batch.TokenContract))
		// // l.Log().Info("sigs", sigs)
		// if err != nil {
		// 	return false, err
		// }

		// iHaveSigned := false
		// for _, sig := range sigs {
		// 	if sig.EthSigner == l.cfg.EthereumAddr.Hex() {
		// 		iHaveSigned = true
		// 		break
		// 	}
		// }

		// if len(sigs) == 0 || !iHaveSigned {
		// 	// should be relayed later with my signature
		// 	return false, nil
		// }

		// l.Log().Info("batch", batch.BatchNonce, " - tokenContract: ", batch.TokenContract)

		// txData, err := l.ethereum.PrepareTransactionBatch(ctx, latestEthValset, batch, sigs)
		// if err != nil {
		// 	// Returning an error here triggers retries which don't help much except risk a binary crash
		// 	// Better to warn the user and try again in the next loop interval
		// 	l.Log().WithError(err).Warningln("failed to prepare outgoing tx batch")
		// 	return false, err
		// }

		// txHash, cost, err := l.ethereum.SendPreparedTxSync(ctx, txData)
		// if err != nil {
		// 	l.Log().WithError(err).Warningln("failed to send outgoing tx batch - Please check your RPC's")
		// 	return false, err
		// }

		// feesTaken := batch.GetFees().BigInt()
		// // TODO: save fees taken and expenses
		// storage.UpdateFeesFile(feesTaken, batch.TokenContract, cost, txHash.Hex(), latestEthHeight.Number.Uint64(), l.cfg.ChainId, "BATCH")
		// ///////

		// l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent outgoing tx batch to " + l.cfg.ChainName)
	}

	for tokenContract, batchs := range grpOfSignedBatchs {
		symbol, ok := l.CacheSymbol[gethcommon.HexToAddress(tokenContract)]
		if !ok {
			symbol = "unknown"
		}
		l.Log().Info("tokenContract: ", tokenContract, " - batchs: ", len(batchs), " - symbol: "+symbol)
	}

	if len(grpOfSignedBatchs) == 0 {
		return false, nil
	}

	// process by taking 1 of each grp then send tx and iterate like this
	stopGrp := make(map[string]bool, 0)
	retryGrp := make(map[string]bool, 0)
	for {
		// take 1 of each grp
		batchToRelay := make([]*BatchAndSigs, 0)

		for tokenContract, batchAndSigs := range grpOfSignedBatchs {
			if len(batchAndSigs) == 0 || stopGrp[tokenContract] {
				continue
			}
			batchAndSig := batchAndSigs[0]
			batchToRelay = append(batchToRelay, batchAndSig)
		}

		if len(batchToRelay) == 0 || len(grpOfSignedBatchs) == 0 {
			break
		}

		// send tx
		for _, batchAndSig := range batchToRelay {

			symbol, ok := l.CacheSymbol[gethcommon.HexToAddress(batchAndSig.Batch.TokenContract)]
			if !ok {
				symbol = batchAndSig.Batch.TokenContract
			}

			l.Orchestrator.HyperionState.RelayerStatus = "sending batch " + strconv.Itoa(int(batchAndSig.Batch.BatchNonce)) + " - " + symbol
			txData, err := l.ethereum.PrepareTransactionBatch(ctx, latestEthValset, batchAndSig.Batch, batchAndSig.Sigs)
			if err != nil {
				stopGrp[batchAndSig.Batch.TokenContract] = true
				l.Orchestrator.HyperionState.RelayerStatus = "error preparing batch " + symbol
				l.Log().WithError(err).Warningln("failed to prepare outgoing tx batch")
				time.Sleep(2 * time.Second)
				continue
			}
			txHash, cost, err := l.ethereum.SendPreparedTx(ctx, txData)
			if err != nil {
				stopGrp[batchAndSig.Batch.TokenContract] = true
				l.Orchestrator.HyperionState.RelayerStatus = "error sending batch " + symbol
				l.Log().WithError(err).Warningln("failed to send outgoing tx batch")
				if strings.Contains(err.Error(), "failed to estimate gas") {
					l.ethereum.RemoveRpc(l.ethereum.GetLastUsedRpc())
					l.Log().Infoln("rpc don't work, removing rpc ", l.ethereum.GetLastUsedRpc())
					retryGrp[batchAndSig.Batch.TokenContract] = true
					stopGrp[batchAndSig.Batch.TokenContract] = false
				}
				time.Sleep(2 * time.Second)
				continue
			}
			lastUsedRpc := l.ethereum.GetLastUsedRpc()
			time.Sleep(5 * time.Second) // wait for transaction to in pool on multiple nodes
			_, blockNumber, err := l.ethereum.WaitForTransaction(ctx, *txHash)
			if err != nil {
				stopGrp[batchAndSig.Batch.TokenContract] = true
				l.Orchestrator.HyperionState.RelayerStatus = "error waiting for transaction " + symbol
				l.Log().WithError(err).Warningln("failed to wait for transaction")
				if strings.Contains(err.Error(), "not found") {
					l.ethereum.RemoveRpc(lastUsedRpc)
					l.Log().Infoln("rpc don't work, removing rpc ", lastUsedRpc)
					retryGrp[batchAndSig.Batch.TokenContract] = true
					stopGrp[batchAndSig.Batch.TokenContract] = false
				}
				time.Sleep(2 * time.Second)
				continue
			}
			l.Orchestrator.HyperionState.RelayerStatus = "batch sent " + strconv.Itoa(int(batchAndSig.Batch.BatchNonce)) + " - " + symbol + " - block number: " + strconv.Itoa(int(blockNumber))
			l.HyperionState.OutBridgedTxCount += len(batchAndSig.Batch.Transactions)
			feesTaken := batchAndSig.Batch.GetFees().BigInt()
			storage.UpdateFeesFile(feesTaken, batchAndSig.Batch.TokenContract, cost, txHash.Hex(), latestEthHeight.Number.Uint64(), l.cfg.ChainId, "BATCH")
			///////

			l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent outgoing tx batch to " + l.cfg.ChainName)
		}

		for tokenContract, batchAndSigs := range grpOfSignedBatchs {
			if len(batchAndSigs) == 0 || stopGrp[tokenContract] {
				stopGrp[tokenContract] = true
				continue
			}
			if retryGrp[tokenContract] {
				retryGrp[tokenContract] = false
				continue
			}
			grpOfSignedBatchs[tokenContract] = batchAndSigs[1:]
		}
		// wait 1 second
		// l.Orchestrator.HyperionState.RelayerStatus = "waiting 20sec before next batch"
		time.Sleep(1 * time.Second)
	}

	return true, nil

}

// func (l *relayer) relayTokenBatch(ctx context.Context, latestEthValset *hyperiontypes.Valset) (bool, error) {
// 	metrics.ReportFuncCall(l.svcTags)
// 	doneFn := metrics.ReportFuncTiming(l.svcTags)
// 	defer doneFn()

// 	batches, err := l.helios.LatestTransactionBatches(ctx, l.cfg.HyperionId)
// 	if l.logEnabled {
// 		numberOfDifferentTokenBatches := 0
// 		mapTokenContract := make(map[string]bool)
// 		for _, batch := range batches {
// 			if _, ok := mapTokenContract[batch.TokenContract]; !ok {
// 				mapTokenContract[batch.TokenContract] = true
// 				numberOfDifferentTokenBatches++
// 			}
// 		}
// 		l.Log().Info("batches: ", len(batches), "numberOfDifferentTokenBatches: ", numberOfDifferentTokenBatches)
// 	}
// 	if err != nil {
// 		l.Log().Info("failed to get latest transaction batches", err)
// 		return false, err
// 	}

// 	latestEthHeight, err := l.ethereum.GetHeaderByNumber(ctx, nil)
// 	if err != nil {
// 		l.Log().Info("failed to get latest "+l.cfg.ChainName+" height", err)
// 		return false, err
// 	}

// 	sort.Slice(batches, func(i, j int) bool {
// 		return batches[i].BatchNonce < batches[j].BatchNonce
// 	})

// 	mapBatchPerTokenContract := make(map[string][]*hyperiontypes.OutgoingTxBatch, 0)

// 	for _, batch := range batches {
// 		if _, ok := mapBatchPerTokenContract[batch.TokenContract]; !ok {
// 			mapBatchPerTokenContract[batch.TokenContract] = make([]*hyperiontypes.OutgoingTxBatch, 0)
// 		}
// 		mapBatchPerTokenContract[batch.TokenContract] = append(mapBatchPerTokenContract[batch.TokenContract], batch)
// 	}

// 	if len(mapBatchPerTokenContract) == 0 {
// 		return false, nil
// 	}

// 	hasPushedABatch := false
// 	errorsBatchs := make([]string, 0)
// 	for _, batches := range mapBatchPerTokenContract {
// 		err := l.processBatch(ctx, batches, latestEthValset, latestEthHeight)
// 		if err != nil {
// 			errorsBatchs = append(errorsBatchs, err.Error())
// 		} else {
// 			hasPushedABatch = true
// 		}
// 	}

// 	if !hasPushedABatch {
// 		return false, errors.New(strings.Join(errorsBatchs, "\n"))
// 	}

// 	return true, nil
// }

// func (l *relayer) SignBatch(ctx context.Context, batch *hyperiontypes.OutgoingTxBatch) error {
// 	hyperionIdHash := common.HexToHash(strconv.FormatUint(l.cfg.HyperionId, 16))

// 	if err := l.retry(ctx, func() error {
// 		return l.helios.SendBatchConfirm(ctx,
// 			l.cfg.HyperionId,
// 			l.cfg.EthereumAddr,
// 			hyperionIdHash,
// 			l.ethereum.GetPersonalSignFn(),
// 			batch,
// 		)
// 	}); err != nil {
// 		return err
// 	}
// 	return nil
// }

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
					// l.SignBatch(ctx, batch.Batch)
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
	l.Log().Infoln("Nonce", oldestConfirmedBatch.BatchNonce, "ContractNonce", mapBatchNonce[oldestConfirmedBatch.TokenContract])
	txData, err := l.ethereum.PrepareTransactionBatch(ctx, latestEthValset, oldestConfirmedBatch, confirmations)
	if err != nil {
		// Returning an error here triggers retries which don't help much except risk a binary crash
		// Better to warn the user and try again in the next loop interval
		l.Log().WithError(err).Warningln("failed to prepare outgoing tx batch")
		return err
	}

	txHash, cost, err := l.ethereum.SendPreparedTx(ctx, txData)
	if err != nil {
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

// FindLatestValset finds the latest valset on the Hyperion contract by looking back through the event
// history and finding the most recent ValsetUpdatedEvent. Most of the time this will be very fast
// as the latest update will be in recent blockchain history and the search moves from the present
// backwards in time. In the case that the validator set has not been updated for a very long time
// this will take longer.
// func (l *relayer) findLatestValsetOnEth(ctx context.Context) (*hyperiontypes.Valset, error) {

// 	lastValsetUpdatedEventHeight, err := l.ethereum.GetLastValsetUpdatedEventHeight(ctx)
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to get last valset updated event")
// 	}

// 	latestEthereumValsetNonce, err := l.ethereum.GetValsetNonce(ctx)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "attempting to unmarshall") { // if error is about unmarshalling, remove last used rpc
// 			l.ethereum.RemoveLastUsedRpc()
// 		}
// 		return nil, errors.Wrap(err, "failed to get latest valset nonce")
// 	}

// 	cosmosValset, err := l.helios.ValsetAt(ctx, l.cfg.HyperionId, latestEthereumValsetNonce.Uint64())
// 	if err != nil {
// 		return nil, errors.Wrap(err, "failed to get Helios valset")
// 	}

// 	if lastValsetUpdatedEventHeight.Uint64() > 0 {
// 		valsetUpdatedEvents, err := l.ethereum.GetValsetUpdatedEventsAtSpecificBlock(lastValsetUpdatedEventHeight.Uint64())
// 		if err != nil {
// 			return nil, errors.Wrap(err, "failed to filter past ValsetUpdated events")
// 		}

// 		// by default the lowest found valset goes first, we want the highest
// 		//
// 		// TODO(xlab): this follows the original impl, but sort might be skipped there:
// 		// we could access just the latest element later.
// 		sort.Sort(sort.Reverse(HyperionValsetUpdatedEvents(valsetUpdatedEvents)))

// 		if len(valsetUpdatedEvents) == 0 { // return the cosmos valset if no event is found
// 			return cosmosValset, nil
// 			// return nil, errors.New("failed to get latest valset (Maybe rpc not answering correctly?)")
// 		}

// 		// we take only the first event if we find any at all.
// 		event := valsetUpdatedEvents[0]

// 		if l.logEnabled {
// 			l.Log().Info("found valset at block: ", event.Raw.BlockNumber, " with nonce: ", event.NewValsetNonce.Uint64())
// 		}

// 		valset := &hyperiontypes.Valset{
// 			Nonce:        event.NewValsetNonce.Uint64(),
// 			Members:      make([]*hyperiontypes.BridgeValidator, 0, len(event.Powers)),
// 			RewardAmount: sdkmath.NewIntFromBigInt(event.RewardAmount),
// 			RewardToken:  event.RewardToken.Hex(),
// 		}

// 		for idx, p := range event.Powers {
// 			valset.Members = append(valset.Members, &hyperiontypes.BridgeValidator{
// 				Power:           p.Uint64(),
// 				EthereumAddress: event.Validators[idx].Hex(),
// 			})
// 		}

// 		if l.logEnabled {
// 			checkIfValsetsDiffer(cosmosValset, valset)
// 		}

// 		return valset, nil

// 	}

// 	return nil, ErrNotFound
// }

// var ErrNotFound = errors.New("not found")

// type HyperionValsetUpdatedEvents []*hyperionevents.HyperionValsetUpdatedEvent

// func (a HyperionValsetUpdatedEvents) Len() int { return len(a) }
// func (a HyperionValsetUpdatedEvents) Less(i, j int) bool {
// 	return a[i].NewValsetNonce.Cmp(a[j].NewValsetNonce) < 0
// }
// func (a HyperionValsetUpdatedEvents) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

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
