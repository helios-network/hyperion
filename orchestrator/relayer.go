package orchestrator

import (
	"context"
	"math/big"
	"os"
	"sort"
	"strconv"
	"time"

	sdkmath "cosmossdk.io/math"

	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/util"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

const (
	defaultRelayerLoopDur    = 5 * time.Minute
	findValsetBlocksToSearch = 2000
)

func (s *Orchestrator) runRelayer(ctx context.Context) error {
	if noRelay := !s.cfg.RelayValsets && !s.cfg.RelayBatches; noRelay {
		return nil
	}

	r := relayer{Orchestrator: s}
	s.logger.WithFields(log.Fields{"loop_duration": defaultRelayerLoopDur.String(), "relay_token_batches": r.cfg.RelayBatches, "relay_validator_sets": s.cfg.RelayValsets}).Debugln("starting Relayer...")

	return loops.RunLoop(ctx, defaultRelayerLoopDur, func() error {
		return r.relay(ctx)
	})
}

func (s *Orchestrator) testReplayTokenBatch(ctx context.Context) {
	r := relayer{Orchestrator: s}
	r.mockRelayTokenBatch(ctx, nil)
}

type relayer struct {
	*Orchestrator
}

func (l *relayer) Log() log.Logger {
	return l.logger.WithField("loop", "Relayer")
}

func (l *relayer) relay(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	ethValset, err := l.getLatestEthValset(ctx)
	if err != nil {
		return err
	}

	var pg loops.ParanoidGroup

	// if l.cfg.RelayValsets {
	// 	pg.Go(func() error {
	// 		return l.retry(ctx, func() error {
	// 			return l.relayValset(ctx, ethValset)
	// 		})
	// 	})
	// }

	if l.cfg.RelayBatches {
		pg.Go(func() error {
			return l.retry(ctx, func() error {
				return l.relayTokenBatch(ctx, ethValset)
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

func (l *relayer) getLatestEthValset(ctx context.Context) (*hyperiontypes.Valset, error) {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	var latestEthValset *hyperiontypes.Valset
	fn := func() error {
		vs, err := l.findLatestValsetOnEth(ctx)
		if err != nil {
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

	if !l.shouldRelayValset(ctx, latestConfirmedValset) {
		return nil
	}

	txHash, err := l.ethereum.SendEthValsetUpdate(ctx, latestEthValset, latestConfirmedValset, confirmations)
	if err != nil {
		return err
	}

	l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent validator set update to Ethereum")

	return nil
}

func (l *relayer) shouldRelayValset(ctx context.Context, vs *hyperiontypes.Valset) bool {
	latestEthereumValsetNonce, err := l.ethereum.GetValsetNonce(ctx)
	if err != nil {
		l.Log().WithError(err).Warningln("failed to get latest valset nonce from Ethereum")
		return false
	}

	// Check if other validators already updated the valset
	if vs.Nonce <= latestEthereumValsetNonce.Uint64() {
		l.Log().WithFields(log.Fields{"eth_nonce": latestEthereumValsetNonce, "helios_nonce": vs.Nonce}).Debugln("validator set already updated on Ethereum")
		return false
	}

	// Check custom time delay offset
	block, err := l.helios.GetBlock(ctx, int64(vs.Height))
	if err != nil {
		l.Log().WithError(err).Warningln("unable to get latest block from Helios")
		return false
	}

	if timeElapsed := time.Since(block.Block.Time); timeElapsed <= l.cfg.RelayValsetOffsetDur {
		timeRemaining := time.Duration(int64(l.cfg.RelayValsetOffsetDur) - int64(timeElapsed))
		l.Log().WithField("time_remaining", timeRemaining.String()).Debugln("valset relay offset not reached yet")
		return false
	}

	l.Log().WithFields(log.Fields{"helios_nonce": vs.Nonce, "eth_nonce": latestEthereumValsetNonce.Uint64()}).Debugln("new valset update")

	return true
}

func (l *relayer) relayTokenBatch(ctx context.Context, latestEthValset *hyperiontypes.Valset) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Load Failed .env: %v", err)
	}
	hyperionId, _ := strconv.ParseUint(os.Getenv("HYPERION_ID"), 10, 64)

	batches, err := l.helios.LatestTransactionBatches(ctx)
	log.Info("batches: ", batches)
	if err != nil {
		log.Info("failed to get latest transaction batches", err)
		return err
	}

	_, err = l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		log.Info("failed to get latest ethereum height", err)
		return err
	}
	// log.Info("latestEthHeight", latestEthHeight)

	var (
		oldestConfirmedBatch *hyperiontypes.OutgoingTxBatch
		confirmations        []*hyperiontypes.MsgConfirmBatch
	)

	for _, batch := range batches {
		log.Info("batch details: ", batch)
		// if batch.BatchTimeout <= latestEthHeight.Number.Uint64() {
		// 	l.Log().WithFields(log.Fields{"batch_nonce": batch.BatchNonce, "batch_timeout_height": batch.BatchTimeout, "latest_eth_height": latestEthHeight.Number.Uint64()}).Debugln("skipping timed out batch")
		// 	continue
		// }

		if batch.HyperionId != hyperionId {
			continue
		}

		sigs, err := l.helios.TransactionBatchSignatures(ctx, batch.BatchNonce, gethcommon.HexToAddress(batch.TokenContract))
		log.Info("sigs", sigs)
		if err != nil {
			return err
		}

		if len(sigs) == 0 {
			continue
		}

		oldestConfirmedBatch = batch
		confirmations = sigs
		if oldestConfirmedBatch != nil {
			break
		}
	}

	if oldestConfirmedBatch == nil {
		l.Log().Infoln("no token batch to relay")
		return nil
	}
	// log.Info("oldestConfirmedBatch", oldestConfirmedBatch)

	// log.Info("shouldRelayBatch", l.shouldRelayBatch(ctx, oldestConfirmedBatch))
	// if !l.shouldRelayBatch(ctx, oldestConfirmedBatch) {
	// 	return nil
	// }

	txHash, err := l.ethereum.SendTransactionBatch(ctx, latestEthValset, oldestConfirmedBatch, confirmations)
	if err != nil {
		// Returning an error here triggers retries which don't help much except risk a binary crash
		// Better to warn the user and try again in the next loop interval
		log.WithError(err).Warningln("failed to send outgoing tx batch to Ethereum")
		return nil
	}

	l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent outgoing tx batch to Ethereum")

	return nil
}

func (l *relayer) mockRelayTokenBatch(ctx context.Context, latestEthValset *hyperiontypes.Valset) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Load Failed .env: %v", err)
	}
	hyperionId, _ := strconv.ParseUint(os.Getenv("HYPERION_ID"), 10, 64)

	batches := []*hyperiontypes.OutgoingTxBatch{
		{
			HyperionId:    1,
			TokenContract: "0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8",
			BatchNonce:    1,
			BatchTimeout:  1,
			Block:         43,
			Transactions: []*hyperiontypes.OutgoingTransferTx{
				{
					Id:          1,
					Sender:      "helios1q0d2nv8xpf9qy22djzgrkgrrcst9frcs34fqra",
					DestAddress: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
					Erc20Token: &hyperiontypes.ERC20Token{
						Contract: "0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8",
						Amount:   sdkmath.NewIntFromBigInt(big.NewInt(3322222)),
					},
					Erc20Fee: &hyperiontypes.ERC20Token{
						Contract: "0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8",
						Amount:   sdkmath.NewIntFromBigInt(big.NewInt(10000)),
					},
				},
				{
					Id:          2,
					Sender:      "helios1q0d2nv8xpf9qy22djzgrkgrrcst9frcs34fqra",
					DestAddress: "0x17267eB1FEC301848d4B5140eDDCFC48945427Ab",
					Erc20Token: &hyperiontypes.ERC20Token{
						Contract: "0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8",
						Amount:   sdkmath.NewIntFromBigInt(big.NewInt(3322222)),
					},
					Erc20Fee: &hyperiontypes.ERC20Token{
						Contract: "0x1ae1cf7d011589e552E26f7F34A7716A4b4B6Ff8",
						Amount:   sdkmath.NewIntFromBigInt(big.NewInt(10000)),
					},
				},
			},
		},
	}

	_, err = l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		log.Info("failed to get latest ethereum height", err)
		return err
	}
	// log.Info("latestEthHeight", latestEthHeight)

	var (
		oldestConfirmedBatch *hyperiontypes.OutgoingTxBatch
		confirmations        []*hyperiontypes.MsgConfirmBatch
	)

	for _, batch := range batches {
		log.Info("batch details: ", batch)
		// if batch.BatchTimeout <= latestEthHeight.Number.Uint64() {
		// 	l.Log().WithFields(log.Fields{"batch_nonce": batch.BatchNonce, "batch_timeout_height": batch.BatchTimeout, "latest_eth_height": latestEthHeight.Number.Uint64()}).Debugln("skipping timed out batch")
		// 	continue
		// }

		if batch.HyperionId != hyperionId {
			log.Info("skipping batch with hyperion id: ", batch.HyperionId)
			continue
		}

		sigs, err := l.helios.TransactionBatchSignatures(ctx, batch.BatchNonce, gethcommon.HexToAddress(batch.TokenContract))
		log.Info("sigs", sigs)
		if err != nil {
			return err
		}

		if len(sigs) == 0 {
			continue
		}

		oldestConfirmedBatch = batch
		confirmations = sigs
		if oldestConfirmedBatch != nil {
			break
		}
	}

	if oldestConfirmedBatch == nil {
		l.Log().Infoln("no token batch to relay")
		return nil
	}
	// log.Info("oldestConfirmedBatch", oldestConfirmedBatch)

	// log.Info("shouldRelayBatch", l.shouldRelayBatch(ctx, oldestConfirmedBatch))
	// if !l.shouldRelayBatch(ctx, oldestConfirmedBatch) {
	// 	return nil
	// }

	txHash, err := l.ethereum.SendTransactionBatch(ctx, latestEthValset, oldestConfirmedBatch, confirmations)
	if err != nil {
		// Returning an error here triggers retries which don't help much except risk a binary crash
		// Better to warn the user and try again in the next loop interval
		log.WithError(err).Warningln("failed to send outgoing tx batch to Ethereum")
		return nil
	}

	l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent outgoing tx batch to Ethereum")

	return nil
}

func (l *relayer) shouldRelayBatch(ctx context.Context, batch *hyperiontypes.OutgoingTxBatch) bool {
	latestEthBatch, err := l.ethereum.GetTxBatchNonce(ctx, gethcommon.HexToAddress(batch.TokenContract))
	if err != nil {
		l.Log().WithError(err).Warningf("unable to get latest batch nonce from Ethereum: token_contract=%s", gethcommon.HexToAddress(batch.TokenContract))
		return false
	}

	// Check if ethereum batch was updated by other validators
	if batch.BatchNonce <= latestEthBatch.Uint64() {
		l.Log().WithFields(log.Fields{"eth_nonce": latestEthBatch.Uint64(), "helios_nonce": batch.BatchNonce}).Debugln("batch already updated on Ethereum")
		return false
	}

	// Check custom time delay offset
	blockTime, err := l.helios.GetBlock(ctx, int64(batch.Block))
	if err != nil {
		l.Log().WithError(err).Warningln("unable to get latest block from Helios")
		return false
	}

	if timeElapsed := time.Since(blockTime.Block.Time); timeElapsed <= l.cfg.RelayBatchOffsetDur {
		timeRemaining := time.Duration(int64(l.cfg.RelayBatchOffsetDur) - int64(timeElapsed))
		l.Log().WithField("time_remaining", timeRemaining.String()).Debugln("batch relay offset not reached yet")
		return false
	}

	l.Log().WithFields(log.Fields{"helios_nonce": batch.BatchNonce, "eth_nonce": latestEthBatch.Uint64()}).Debugln("new batch update")

	return true
}

// FindLatestValset finds the latest valset on the Hyperion contract by looking back through the event
// history and finding the most recent ValsetUpdatedEvent. Most of the time this will be very fast
// as the latest update will be in recent blockchain history and the search moves from the present
// backwards in time. In the case that the validator set has not been updated for a very long time
// this will take longer.
func (l *relayer) findLatestValsetOnEth(ctx context.Context) (*hyperiontypes.Valset, error) {
	latestHeader, err := l.ethereum.GetHeaderByNumber(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get latest ethereum header")
	}

	latestEthereumValsetNonce, err := l.ethereum.GetValsetNonce(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get latest valset nonce on Ethereum")
	}

	cosmosValset, err := l.helios.ValsetAt(ctx, l.cfg.HyperionId, latestEthereumValsetNonce.Uint64())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Helios valset")
	}

	currentBlock := latestHeader.Number.Uint64()

	for currentBlock > 0 {
		var startSearchBlock uint64
		if currentBlock <= findValsetBlocksToSearch {
			startSearchBlock = 0
		} else {
			startSearchBlock = currentBlock - findValsetBlocksToSearch
		}

		valsetUpdatedEvents, err := l.ethereum.GetValsetUpdatedEvents(startSearchBlock, currentBlock)
		if err != nil {
			return nil, errors.Wrap(err, "failed to filter past ValsetUpdated events from Ethereum")
		}

		// by default the lowest found valset goes first, we want the highest
		//
		// TODO(xlab): this follows the original impl, but sort might be skipped there:
		// we could access just the latest element later.
		sort.Sort(sort.Reverse(HyperionValsetUpdatedEvents(valsetUpdatedEvents)))

		if len(valsetUpdatedEvents) == 0 {
			currentBlock = startSearchBlock
			continue
		}

		// we take only the first event if we find any at all.
		event := valsetUpdatedEvents[0]
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

		checkIfValsetsDiffer(cosmosValset, valset)

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
		).Errorln("Cosmos does not have a valset for nonce from Ethereum chain. Possible bridge hijacking!")
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
