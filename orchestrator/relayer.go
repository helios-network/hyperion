package orchestrator

import (
	"context"
	"sort"
	"strconv"
	"strings"
	"time"

	gethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/provider"
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

	// Sélectionner le meilleur RPC basé sur la réputation
	bestRpcURL := l.ethereum.SelectBestRatedRpcInRpcPool()
	if bestRpcURL != "" {
		l.Log().WithField("selected_rpc", bestRpcURL).Debug("Selected best rated RPC for relay")
		// Ajouter le meilleur RPC au contexte
		ctx = provider.WithRPCURL(ctx, bestRpcURL)
	}

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

	if ethValset == nil {
		l.Log().Infoln("valset not found, skipping relay")
		return nil
	}

	if l.cfg.RelayBatches {
		pg.Go(func() error {
			return l.retry(ctx, func() error {
				_, err := l.relayBatchsOptimised(ctx, ethValset)
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

type BatchAndSigs struct {
	Batch *hyperiontypes.OutgoingTxBatch
	Sigs  []*hyperiontypes.MsgConfirmBatch
}

func (l *relayer) relayBatchsOptimised(ctx context.Context, latestEthValset *hyperiontypes.Valset) (bool, error) {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	latestEthHeight, err := l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		l.Log().Info("failed to get latest "+l.cfg.ChainName+" height", err)
		usedRpc := provider.GetCurrentRPCURL(ctx)
		if usedRpc != "" {
			l.ethereum.PenalizeRpc(usedRpc, 1) // Pénalité de 1 point
			l.Log().WithField("rpc", usedRpc).Debug("Penalized RPC for failed to get latest " + l.cfg.ChainName + " height")
		}
		return false, err
	}

	maxHeightTimeout := uint64(latestEthHeight.Number.Uint64() + 10)

	batchesInHelios, err := l.helios.LatestTransactionBatchesWithOptions(ctx, l.cfg.HyperionId, l.cfg.CosmosAddr.String(), 0, maxHeightTimeout, "", true)

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
		l.Log().Info("batch ", batch.BatchNonce, " - txs: ", len(batch.Transactions), " - batchTimeout: ", batch.BatchTimeout, " - latestEthHeight: ", latestEthHeight.Number.Uint64()+uint64(10))
		if batch.BatchTimeout > maxHeightTimeout {
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
			} else {
				l.Log().Info("batch refused ", batch.BatchNonce, " for tokenContract: ", batch.TokenContract, " Nonce less than latestBatchNonce: ", latestBatchNonce)
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
			packSigned = append(packSigned, &BatchAndSigs{Batch: batch, Sigs: nil})
		}

		if len(packSigned) == 0 {
			continue
		}

		grpOfSignedBatchs[tokenContrat] = packSigned
	}

	for tokenContract, batchs := range grpOfSignedBatchs {
		symbol, ok := l.CacheSymbol[gethcommon.HexToAddress(tokenContract)]
		if !ok {
			symbol = "unknown"
		}
		l.Log().Info("tokenContract: ", tokenContract, " - batchs: ", len(batchs), " - symbol: "+symbol)
	}

	if len(grpOfSignedBatchs) == 0 {
		l.Log().Infoln("no signed batches to relay ", "actualBatchNonce On Contract: ", grpOfSignedBatchs)
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
			sigs, err := l.helios.TransactionBatchSignatures(ctx, l.cfg.HyperionId, batchAndSig.Batch.BatchNonce, gethcommon.HexToAddress(batchAndSig.Batch.TokenContract))
			if err != nil {
				l.Log().WithError(err).Warningln("failed to get transaction batch signatures")
				continue
			}
			batchAndSig.Sigs = sigs
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
				usedRpc := provider.GetCurrentRPCURL(ctx)
				if usedRpc != "" {
					l.ethereum.PenalizeRpc(usedRpc, 2) // Pénalité de 2 points
					l.Log().WithField("rpc", usedRpc).Debug("Penalized RPC for failed transaction")
				}
				time.Sleep(2 * time.Second)
				continue
			}
			time.Sleep(5 * time.Second) // wait for transaction to in pool on multiple nodes
			_, blockNumber, err := l.ethereum.WaitForTransaction(ctx, *txHash)
			if err != nil {
				stopGrp[batchAndSig.Batch.TokenContract] = true
				l.Orchestrator.HyperionState.RelayerStatus = "error waiting for transaction " + symbol
				l.Log().WithError(err).Warningln("failed to wait for transaction")
				usedRpc := provider.GetCurrentRPCURL(ctx)
				// Pénaliser le RPC utilisé pour cet échec
				if usedRpc != "" {
					l.ethereum.PenalizeRpc(usedRpc, 1) // Pénalité de 1 point
					l.Log().WithField("rpc", usedRpc).Debug("Penalized RPC for transaction not found")
				}
				time.Sleep(2 * time.Second)
				continue
			}
			l.Orchestrator.HyperionState.RelayerStatus = "batch sent " + strconv.Itoa(int(batchAndSig.Batch.BatchNonce)) + " - " + symbol + " - block number: " + strconv.Itoa(int(blockNumber))
			l.HyperionState.OutBridgedTxCount += len(batchAndSig.Batch.Transactions)
			feesTaken := batchAndSig.Batch.GetFees().BigInt()
			storage.UpdateFeesFile(feesTaken, batchAndSig.Batch.TokenContract, cost, txHash.Hex(), latestEthHeight.Number.Uint64(), l.cfg.ChainId, "BATCH")
			///////

			usedRpc := provider.GetCurrentRPCURL(ctx)
			// Féliciter le RPC utilisé pour ce succès
			if usedRpc != "" {
				l.ethereum.PraiseRpc(usedRpc, 3) // Récompense de 3 points
				l.Log().WithField("rpc", usedRpc).Debug("Praised RPC for successful transaction")
			}

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

func (l *relayer) relayBatchs(ctx context.Context, latestEthValset *hyperiontypes.Valset) (bool, error) {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	latestEthHeight, err := l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		l.Log().Info("failed to get latest "+l.cfg.ChainName+" height", err)
		usedRpc := provider.GetCurrentRPCURL(ctx)
		if usedRpc != "" {
			l.ethereum.PenalizeRpc(usedRpc, 1) // Pénalité de 1 point
			l.Log().WithField("rpc", usedRpc).Debug("Penalized RPC for failed to get latest " + l.cfg.ChainName + " height")
		}
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
		l.Log().Info("batch ", batch.BatchNonce, " - txs: ", len(batch.Transactions), " - batchTimeout: ", batch.BatchTimeout, " - latestEthHeight: ", latestEthHeight.Number.Uint64()+uint64(10))
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
			} else {
				l.Log().Info("batch refused ", batch.BatchNonce, " for tokenContract: ", batch.TokenContract, " Nonce less than latestBatchNonce: ", latestBatchNonce)
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
	}

	for tokenContract, batchs := range grpOfSignedBatchs {
		symbol, ok := l.CacheSymbol[gethcommon.HexToAddress(tokenContract)]
		if !ok {
			symbol = "unknown"
		}
		l.Log().Info("tokenContract: ", tokenContract, " - batchs: ", len(batchs), " - symbol: "+symbol)
	}

	if len(grpOfSignedBatchs) == 0 {
		l.Log().Infoln("no signed batches to relay ", "actualBatchNonce On Contract: ", grpOfSignedBatchs)
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
				usedRpc := provider.GetCurrentRPCURL(ctx)
				if usedRpc != "" {
					l.ethereum.PenalizeRpc(usedRpc, 2) // Pénalité de 2 points
					l.Log().WithField("rpc", usedRpc).Debug("Penalized RPC for failed transaction")
				}
				time.Sleep(2 * time.Second)
				continue
			}
			time.Sleep(5 * time.Second) // wait for transaction to in pool on multiple nodes
			_, blockNumber, err := l.ethereum.WaitForTransaction(ctx, *txHash)
			if err != nil {
				stopGrp[batchAndSig.Batch.TokenContract] = true
				l.Orchestrator.HyperionState.RelayerStatus = "error waiting for transaction " + symbol
				l.Log().WithError(err).Warningln("failed to wait for transaction")
				usedRpc := provider.GetCurrentRPCURL(ctx)
				// Pénaliser le RPC utilisé pour cet échec
				if usedRpc != "" {
					l.ethereum.PenalizeRpc(usedRpc, 1) // Pénalité de 1 point
					l.Log().WithField("rpc", usedRpc).Debug("Penalized RPC for transaction not found")
				}
				time.Sleep(2 * time.Second)
				continue
			}
			l.Orchestrator.HyperionState.RelayerStatus = "batch sent " + strconv.Itoa(int(batchAndSig.Batch.BatchNonce)) + " - " + symbol + " - block number: " + strconv.Itoa(int(blockNumber))
			l.HyperionState.OutBridgedTxCount += len(batchAndSig.Batch.Transactions)
			feesTaken := batchAndSig.Batch.GetFees().BigInt()
			storage.UpdateFeesFile(feesTaken, batchAndSig.Batch.TokenContract, cost, txHash.Hex(), latestEthHeight.Number.Uint64(), l.cfg.ChainId, "BATCH")
			///////

			usedRpc := provider.GetCurrentRPCURL(ctx)
			// Féliciter le RPC utilisé pour ce succès
			if usedRpc != "" {
				l.ethereum.PraiseRpc(usedRpc, 3) // Récompense de 3 points
				l.Log().WithField("rpc", usedRpc).Debug("Praised RPC for successful transaction")
			}

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

type BatchToSign struct {
	Batch  *hyperiontypes.OutgoingTxBatch
	Signed bool
}

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
