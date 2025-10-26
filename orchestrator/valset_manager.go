package orchestrator

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/storage"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"
)

const (
	defaultValsetManagerLoopDur = 5 * time.Minute
)

func (s *Orchestrator) runValsetManager(ctx context.Context) error {
	valsetManager := valsetManager{
		Orchestrator: s,
		logEnabled:   strings.Contains(s.cfg.EnabledLogs, "updater"),
	}
	s.logger.WithField("loop_duration", defaultValsetManagerLoopDur.String()).Debugln("starting ValsetManager...")
	s.SetValsetManager(&valsetManager)
	return loops.RunLoop(ctx, s.ethereum, defaultValsetManagerLoopDur, func() error {
		if s.HyperionState.ValsetManagerStatus == "running" {
			return nil
		}

		start := time.Now()
		s.HyperionState.ValsetManagerStatus = "running"
		err := valsetManager.Process(ctx)
		s.HyperionState.ValsetManagerStatus = "idle"
		s.HyperionState.ValsetManagerLastExecutionFinishedTimestamp = uint64(time.Now().Unix())
		s.HyperionState.ValsetManagerNextExecutionTimestamp = uint64(start.Add(defaultValsetManagerLoopDur).Unix())

		if err != nil {
			s.logger.WithError(err).Errorln("valset manager function returned an error")
		}

		return err
	})
}

type valsetManager struct {
	*Orchestrator
	logEnabled       bool
	synced           bool
	consideredSynced bool
	ethValset        *hyperiontypes.Valset
}

func (l *valsetManager) Log() log.Logger {
	return l.logger.WithField("loop", "ValsetManager")
}

func (l *valsetManager) IsValsetSynced(ctx context.Context) bool {
	if l.consideredSynced {
		return true
	}
	return l.synced
}

func (l *valsetManager) GetEthValset(ctx context.Context) *hyperiontypes.Valset {
	return l.ethValset
}

func (l *valsetManager) Process(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	// bestRpcURL := l.ethereum.SelectBestRatedRpcInRpcPool()
	// if bestRpcURL != "" {
	// 	ctx = provider.WithRPCURL(ctx, bestRpcURL)
	// }

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

	l.Log().WithFields(log.Fields{
		"Helios": heliosCheckpoint.Hex(),
		"Eth":    ethCheckpoint.Hex(),
		"Synced": heliosCheckpoint.Hex() == ethCheckpoint.Hex(),
	}).Infoln("Relayer: checkpoints")

	if heliosCheckpoint.Hex() != ethCheckpoint.Hex() {
		if l.logEnabled {
			l.Log().Infoln("relayer: checkpoint not synced yet waiting (rpc should be untrustable) ...")
		}
		l.synced = false
		l.ethValset = nil
	} else {
		if l.logEnabled {
			l.Log().Info("valset is synced")
		}
		l.synced = true
		l.ethValset = ethValset
	}

	// write valset to file
	// json, err := json.Marshal(ethValset)
	// if err == nil {
	// 	os.WriteFile("valset.json", json, 0644)
	// }

	var pg loops.ParanoidGroup

	pg.Go(func() error {
		return l.retry(ctx, func() error {
			return l.relayValset(ctx, ethValset)
		})
	})

	if err := pg.Wait(); err != nil {
		return err
	}

	l.Orchestrator.SetValsetManager(l)

	l.logger.Info("ValsetManager processed", " consideredSynced ", l.consideredSynced, " synced ", l.synced)
	return nil
}

func (l *valsetManager) getLatestEthValset(ctx context.Context) (*hyperiontypes.Valset, error) {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	var latestEthValset *hyperiontypes.Valset
	fn := func() error {
		vs, err := l.findLatestValsetOnEth(ctx)
		if err != nil {
			if strings.Contains(err.Error(), "failed to get") || strings.Contains(err.Error(), "attempting to unmarshall") || strings.Contains(err.Error(), "pruned") {
				l.Orchestrator.RotateRpc()
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

func (l *valsetManager) relayValset(ctx context.Context, latestEthValset *hyperiontypes.Valset) error {
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
		l.Log().Infoln("valset is not ready to relay")
		return nil
	}

	fmt.Println("Sending valset update to Ethereum", latestEthValset, latestConfirmedValset, confirmations)

	txHash, cost, err := l.ethereum.SendEthValsetUpdate(ctx, latestEthValset, latestConfirmedValset, confirmations)
	if err != nil {

		if strings.Contains(err.Error(), "insuffficient funds for gas") {
			l.Orchestrator.HyperionState.ValsetManagerStatus = "insufficient funds for gas"
			return err
		}
		return err
	}

	storage.UpdateFeesFile(latestEthValset.RewardAmount.BigInt(), latestEthValset.RewardToken, cost, txHash.Hex(), latestEthValset.Height, l.cfg.ChainId, "VALSET")

	l.Log().WithField("tx_hash", txHash.Hex()).Infoln("sent validator set update to Ethereum")

	return nil
}

func (l *valsetManager) shouldRelayValset(ctx context.Context, vs *hyperiontypes.Valset) bool {

	var latestEthereumValsetNonce *big.Int

	fn := func() error {
		nonce, err := l.ethereum.GetValsetNonce(ctx)
		if err != nil {
			return err
		}
		latestEthereumValsetNonce = nonce
		return nil
	}
	if err := l.retry(ctx, fn); err != nil {
		l.Log().WithError(err).Warningln("failed to get latest valset nonce from " + l.cfg.ChainName)
		return false
	}

	if latestEthereumValsetNonce == nil {
		l.Log().Warningln("failed to get latest valset nonce from " + l.cfg.ChainName)
		return false
	}

	// Check if other validators already updated the valset
	if vs.Nonce <= latestEthereumValsetNonce.Uint64() {
		l.Log().WithFields(log.Fields{"eth_nonce": latestEthereumValsetNonce, "helios_nonce": vs.Nonce}).Infoln("validator set already updated on " + l.cfg.ChainName)
		l.consideredSynced = true
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
		l.consideredSynced = true
		return false
	}

	l.consideredSynced = false

	l.Log().WithFields(log.Fields{"helios_nonce": vs.Nonce, "eth_nonce": latestEthereumValsetNonce.Uint64()}).Debugln("new valset update")

	return true
}

func (l *valsetManager) findLatestValsetOnEth(ctx context.Context) (*hyperiontypes.Valset, error) {

	lastValsetUpdatedEventHeight, err := l.ethereum.GetLastValsetUpdatedEventHeight(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get last valset updated event")
	}

	latestEthereumValsetNonce, err := l.ethereum.GetValsetNonce(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "attempting to unmarshall") { // if error is about unmarshalling, remove last used rpc
			usedRpc := l.ethereum.GetRpc().Url
			if usedRpc != "" {
				// l.ethereum.PenalizeRpc(usedRpc, 1)
				l.Log().WithField("rpc", usedRpc).Debug("Penalized RPC for valset manager")
			}
		}
		return nil, errors.Wrap(err, "failed to get latest valset nonce")
	}

	fmt.Println("latestEthereumValsetNonce", latestEthereumValsetNonce.Uint64())

	cosmosValset, err := l.helios.ValsetAt(ctx, l.cfg.HyperionId, latestEthereumValsetNonce.Uint64())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get Helios valset")
	}

	if lastValsetUpdatedEventHeight.Uint64() > 0 {
		valsetUpdatedEvents, err := l.ethereum.GetValsetUpdatedEventsAtSpecificBlock(lastValsetUpdatedEventHeight.Uint64())
		if err != nil {
			return nil, errors.Wrap(err, "failed to filter past ValsetUpdated events")
		}

		// manage case where the blockchain not manage correctly the block.number (like arbitrum)
		if len(valsetUpdatedEvents) == 0 {
			latestBlockHeader, err := l.ethereum.GetHeaderByNumber(ctx, nil)
			if err != nil {
				return nil, errors.Wrap(err, "failed to get latest block height")
			}

			valsetUpdatedEvents, err = l.ethereum.GetValsetUpdatedEventsWithIndexedNonce(latestEthereumValsetNonce.Uint64(), l.Orchestrator.cfg.ChainParams.BridgeContractStartHeight, latestBlockHeader.Number.Uint64())
			if err != nil {
				return nil, errors.Wrap(err, "failed to filter past ValsetUpdated events")
			}
			fmt.Println("[Warning] - GetValsetUpdatedEventsWithIndexedNonce : ", len(valsetUpdatedEvents))
		}

		// by default the lowest found valset goes first, we want the highest
		//
		// TODO(xlab): this follows the original impl, but sort might be skipped there:
		// we could access just the latest element later.
		sort.Sort(sort.Reverse(HyperionValsetUpdatedEvents(valsetUpdatedEvents)))

		if len(valsetUpdatedEvents) == 0 { // return the cosmos valset if no event is found
			fmt.Println("No valset on "+l.cfg.ChainName+" updated events found, returning cosmos valset", cosmosValset)
			return cosmosValset, nil
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

func (l *valsetManager) encodeData(
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

func (l *valsetManager) makeCheckpoint(ctx context.Context, valset *hyperiontypes.Valset) (*common.Hash, error) {
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

	if valset == nil {
		return nil, errors.New("valset is nil")
	}

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
