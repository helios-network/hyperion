package orchestrator

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/loops"
	"github.com/Helios-Chain-Labs/metrics"
	"github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	log "github.com/xlab/suplog"
)

const (
	defaultExternalDataLoopDur = 1 * time.Minute
)

func (s *Orchestrator) runExternalData(ctx context.Context) error {
	externalData := externalData{
		Orchestrator: s,
		logEnabled:   strings.Contains(s.cfg.EnabledLogs, "external_data"),
	}
	s.logger.WithField("loop_duration", defaultExternalDataLoopDur.String()).Debugln("starting ExternalData...")

	return loops.RunLoop(ctx, s.ethereum, defaultExternalDataLoopDur, func() error {
		if s.HyperionState.ExternalDataStatus == "running" {
			return nil
		}

		start := time.Now()
		s.HyperionState.ExternalDataStatus = "running"
		err := externalData.Process(ctx)
		s.HyperionState.ExternalDataStatus = "idle"
		s.HyperionState.ExternalDataLastExecutionFinishedTimestamp = uint64(time.Now().Unix())
		s.HyperionState.ExternalDataNextExecutionTimestamp = uint64(start.Add(defaultExternalDataLoopDur).Unix())
		return err
	})
}

type externalData struct {
	*Orchestrator
	logEnabled bool
}

func (l *externalData) Log() log.Logger {
	return l.logger.WithField("loop", "ExternalData")
}

func (l *externalData) Process(ctx context.Context) error {

	var pg loops.ParanoidGroup

	if l.cfg.RelayExternalDatas {
		pg.Go(func() error {
			return l.retry(ctx, func() error {
				return l.relayExternalData(ctx)
			})
		})
	}

	if err := pg.Wait(); err != nil {
		return err
	}

	l.logger.Info("ExternalData processed")
	return nil
}

func (l *externalData) relayExternalData(ctx context.Context) error {
	metrics.ReportFuncCall(l.svcTags)
	doneFn := metrics.ReportFuncTiming(l.svcTags)
	defer doneFn()

	txs, err := l.GetHelios().LatestTransactionExternalCallDataTxs(ctx, l.cfg.HyperionId)
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

	msgsToBroadcast := make([]sdk.Msg, 0)
	maxBulkSize := 10 // todo: make it configurable

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

		msg, err := l.GetHelios().SendExternalDataClaimMsg(ctx, l.cfg.HyperionId, tx.Nonce, latestEthHeight.Number.Uint64(), tx.ExternalContractAddress, callData, callErr, rpcUsed)
		if err != nil {
			l.Log().Info("failed to send external data claim message", err)
			continue
		}

		err = l.GetHelios().SyncBroadcastMsgsSimulate(ctx, []sdk.Msg{msg})
		if err != nil {
			VerifyTxError(ctx, err.Error(), l.Orchestrator)
			l.Log().Info("failed to simulate external data claim message", err)
		}

		if len(msgsToBroadcast) >= maxBulkSize {
			_, err = l.global.SyncBroadcastMsgs(ctx, msgsToBroadcast)
			if err != nil {
				l.Log().Info("failed to broadcast external data claim messages", err)
			}
			l.Orchestrator.HyperionState.ExternalDataCount += len(msgsToBroadcast)
			msgsToBroadcast = make([]sdk.Msg, 0)
		}

		msgsToBroadcast = append(msgsToBroadcast, msg)
	}

	if len(msgsToBroadcast) > 0 {
		_, err = l.global.SyncBroadcastMsgs(ctx, msgsToBroadcast)
		if err != nil {
			l.Log().Info("failed to broadcast external data claim messages", err)
		}
		l.Orchestrator.HyperionState.ExternalDataCount += len(msgsToBroadcast)
	}

	return nil
}

func (l *externalData) selectBestClaimFromListOfClaims(claims []*types.MsgExternalDataClaim) *types.MsgExternalDataClaim {
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
