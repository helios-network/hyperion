package hyperion

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	sdkmath "cosmossdk.io/math"

	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	log "github.com/xlab/suplog"

	"github.com/Helios-Chain-Labs/metrics"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
	"github.com/Helios-Chain-Labs/sdk-go/client/chain"

	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/hyperion"
	"github.com/Helios-Chain-Labs/hyperion/orchestrator/ethereum/keystore"
	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
)

type BroadcastClient interface {
	GetTxCost(ctx context.Context, txHash string) (*big.Int, error)
	SendValsetConfirm(ctx context.Context, hyperionId uint64, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, signFn keystore.PersonalSignFn, valset *hyperiontypes.Valset) error
	SendBatchConfirm(ctx context.Context, hyperionId uint64, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, signFn keystore.PersonalSignFn, batch *hyperiontypes.OutgoingTxBatch) error
	SendRequestBatch(ctx context.Context, hyperionId uint64, denom string) error
	SendToChain(ctx context.Context, chainId uint64, destination gethcommon.Address, amount, fee cosmostypes.Coin) error

	SendDepositClaim(ctx context.Context, hyperionId uint64, deposit *hyperionevents.HyperionSendToHeliosEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error)
	SendWithdrawalClaim(ctx context.Context, hyperionId uint64, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error)
	SendExternalDataClaim(ctx context.Context, hyperionId uint64, nonce uint64, blockHeight uint64, externalContractAddress string, callData []byte, callErr []byte, rpcUsedForObservation string) (*cosmostypes.TxResponse, error)
	SendValsetClaim(ctx context.Context, hyperionId uint64, vs *hyperionevents.HyperionValsetUpdatedEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error)
	SendERC20DeployedClaim(ctx context.Context, hyperionId uint64, erc20 *hyperionevents.HyperionERC20DeployedEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error)

	SendDepositClaimMsg(ctx context.Context, hyperionId uint64, deposit *hyperionevents.HyperionSendToHeliosEvent, rpcUsedForObservation string) (cosmostypes.Msg, error)
	SendWithdrawalClaimMsg(ctx context.Context, hyperionId uint64, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent, rpcUsedForObservation string) (cosmostypes.Msg, error)
	SendExternalDataClaimMsg(ctx context.Context, hyperionId uint64, nonce uint64, blockHeight uint64, externalContractAddress string, callData []byte, callErr []byte, rpcUsedForObservation string) (cosmostypes.Msg, error)
	SendValsetClaimMsg(ctx context.Context, hyperionId uint64, vs *hyperionevents.HyperionValsetUpdatedEvent, rpcUsedForObservation string) (cosmostypes.Msg, error)
	SendERC20DeployedClaimMsg(ctx context.Context, hyperionId uint64, erc20 *hyperionevents.HyperionERC20DeployedEvent, rpcUsedForObservation string) (cosmostypes.Msg, error)

	SyncBroadcastMsgs(ctx context.Context, msgs []cosmostypes.Msg) (*cosmostypes.TxResponse, error)

	SendSetOrchestratorAddresses(ctx context.Context, hyperionId uint64, ethAddress string) error
	SendUnSetOrchestratorAddresses(ctx context.Context, hyperionId uint64, ethAddress string) error
	SendForceSetValsetAndLastObservedEventNonce(ctx context.Context, hyperionId uint64, nonce uint64, blockHeight uint64, valset *hyperiontypes.Valset) error
	SendCancelAllPendingOutTx(ctx context.Context, chainId uint64) error
	SendCancelPendingOutTxs(ctx context.Context, chainId uint64, count uint64) error
}

type broadcastClient struct {
	chain.ChainClient

	svcTags metrics.Tags
}

func NewBroadcastClient(client chain.ChainClient) BroadcastClient {
	return broadcastClient{
		ChainClient: client,
		svcTags:     metrics.Tags{"svc": "hyperion_broadcast"},
	}
}

func (c broadcastClient) SendValsetConfirm(_ context.Context, hyperionId uint64, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, signerFn keystore.PersonalSignFn, valset *hyperiontypes.Valset) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.Infoln("sending valset confirm")

	confirmHash := hyperion.EncodeValsetConfirm(hyperionID, valset)
	signature, err := signerFn(ethFrom, confirmHash.Bytes())
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.New("failed to sign validator address")
	}

	// MsgValsetConfirm
	// this is the message sent by the validators when they wish to submit their
	// signatures over the validator set at a given block height. A validator must
	// first call MsgSetEthAddress to set their Ethereum address to be used for
	// signing. Then someone (anyone) must make a ValsetRequest the request is
	// essentially a messaging mechanism to determine which block all validators
	// should submit signatures over. Finally validators sign the validator set,
	// powers, and Ethereum addresses of the entire validator set at the height of a
	// ValsetRequest and submit that signature with this message.
	//
	// If a sufficient number of validators (66% of voting power) (A) have set
	// Ethereum addresses and (B) submit ValsetConfirm messages with their
	// signatures it is then possible for anyone to view these signatures in the
	// chain store and submit them to Ethereum to update the validator set
	// -------------
	msg := &hyperiontypes.MsgValsetConfirm{
		HyperionId:   hyperionId,
		Orchestrator: c.FromAddress().String(),
		EthAddress:   ethFrom.Hex(),
		Nonce:        valset.Nonce,
		Signature:    gethcommon.Bytes2Hex(signature),
	}

	if err = c.ChainClient.QueueBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgValsetConfirm failed")
	}

	return nil
}

func sigToVRS(sigHex string) (v uint8, r, s gethcommon.Hash) {
	signatureBytes := gethcommon.FromHex(sigHex)
	vParam := signatureBytes[64]
	if vParam == byte(0) {
		vParam = byte(27)
	} else if vParam == byte(1) {
		vParam = byte(28)
	}

	v = vParam
	r = gethcommon.BytesToHash(signatureBytes[0:32])
	s = gethcommon.BytesToHash(signatureBytes[32:64])

	return
}

func (c broadcastClient) SendBatchConfirm(_ context.Context, hyperionId uint64, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, signerFn keystore.PersonalSignFn, batch *hyperiontypes.OutgoingTxBatch) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	confirmHash := hyperion.EncodeTxBatchConfirm(hyperionID, batch)
	// log.Info("confirmHash: ", confirmHash, "batch: ", batch, "hyperionID: ", hyperionID, "ethFrom: ", ethFrom.Hex())
	// log.Info("confirmHashLength: ", len(confirmHash.Bytes()))
	signature, err := signerFn(ethFrom, confirmHash.Bytes())
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.New("failed to sign validator address")
	}

	// sigV, sigR, sigS := sigToVRS(gethcommon.Bytes2Hex(signature))
	// log.Info("sigV: ", sigV, "sigR: ", sigR, "sigS: ", sigS)

	// MsgConfirmBatch
	// When validators observe a MsgRequestBatch they form a batch by ordering
	// transactions currently in the txqueue in order of highest to lowest fee,
	// cutting off when the batch either reaches a hardcoded maximum size (to be
	// decided, probably around 100) or when transactions stop being profitable
	// (TODO determine this without nondeterminism) This message includes the batch
	// as well as an Ethereum signature over this batch by the validator
	// -------------
	msg := &hyperiontypes.MsgConfirmBatch{
		HyperionId:    hyperionId,
		Orchestrator:  c.FromAddress().String(),
		Nonce:         batch.BatchNonce,
		Signature:     gethcommon.Bytes2Hex(signature),
		EthSigner:     ethFrom.Hex(),
		TokenContract: batch.TokenContract,
	}
	log.Info("start confirm batch, msg", msg)

	if err = c.ChainClient.QueueBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgConfirmBatch failed")
	}

	return nil
}

func (c broadcastClient) SendToChain(ctx context.Context, chainId uint64, destination gethcommon.Address, amount, fee cosmostypes.Coin) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	// MsgSendToChain
	// This is the message that a user calls when they want to bridge an asset
	// it will later be removed when it is included in a batch and successfully
	// submitted tokens are removed from the users balance immediately
	// -------------
	// AMOUNT:
	// the coin to send across the bridge, note the restriction that this is a
	// single coin not a set of coins that is normal in other Cosmos messages
	// FEE:
	// the fee paid for the bridge, distinct from the fee paid to the chain to
	// actually send this message in the first place. So a successful send has
	// two layers of fees for the user
	// -------------
	msg := &hyperiontypes.MsgSendToChain{
		Sender:      c.FromAddress().String(),
		DestChainId: chainId,
		Dest:        destination.Hex(),
		Amount:      amount,
		BridgeFee:   fee, // TODO: use exactly that fee for transaction
	}

	if err := c.ChainClient.QueueBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgSendToChain failed")
	}

	return nil
}

func (c broadcastClient) SendRequestBatch(ctx context.Context, hyperionId uint64, denom string) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	// MsgRequestBatch
	// this is a message anyone can send that requests a batch of transactions to
	// send across the bridge be created for whatever block height this message is
	// included in. This acts as a coordination point, the handler for this message
	// looks at the AddToOutgoingPool tx's in the store and generates a batch, also
	// available in the store tied to this message. The validators then grab this
	// batch, sign it, submit the signatures with a MsgConfirmBatch before a relayer
	// can finally submit the batch
	// -------------
	msg := &hyperiontypes.MsgRequestBatch{
		HyperionId:   hyperionId,
		Denom:        denom,
		Orchestrator: c.FromAddress().String(),
	}
	if err := c.ChainClient.QueueBroadcastMsg(msg); err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgRequestBatch failed")
	}

	return nil
}

func (c broadcastClient) GetTxCost(ctx context.Context, txHash string) (*big.Int, error) {
	tx, err := c.ChainClient.GetTx(ctx, txHash)
	if err != nil {
		return nil, err
	}
	return tx.Tx.AuthInfo.Fee.Amount[0].Amount.BigInt(), nil
}

func (c broadcastClient) SendSetOrchestratorAddresses(ctx context.Context, hyperionId uint64, ethAddress string) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	// MsgSetOrchestratorAddresses
	// Permit to set the orchestrator address on the hyperion module
	// -------------
	msg := &hyperiontypes.MsgSetOrchestratorAddresses{
		Sender:       c.FromAddress().String(),
		HyperionId:   hyperionId,
		EthAddress:   ethAddress,
		Orchestrator: c.FromAddress().String(),
	}
	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgSetOrchestratorAddresses failed")
	}

	if resp.TxResponse.Code == 13 {
		log.WithFields(log.Fields{
			"tx_hash": resp.TxResponse.TxHash,
			"code":    resp.TxResponse.Code,
			"Error":   "insufficient fee",
		}).Infoln("EthOracle sent MsgSetOrchestratorAddresses")
		return errors.Wrap(errors.New("code 13 - insufficient fee"), "broadcasting MsgSetOrchestratorAddresses failed")
	}

	// TODO: wait for the tx to be included in a block

	time.Sleep(10 * time.Second)

	return nil
}

func (c broadcastClient) SendUnSetOrchestratorAddresses(ctx context.Context, hyperionId uint64, ethAddress string) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	// MsgUnSetOrchestratorAddresses
	// Permit to unset the orchestrator address on the hyperion module
	// -------------
	msg := &hyperiontypes.MsgUnSetOrchestratorAddresses{
		Sender:     c.FromAddress().String(),
		HyperionId: hyperionId,
		EthAddress: ethAddress,
	}
	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgUnSetOrchestratorAddresses failed")
	}

	if resp.TxResponse.Code == 13 {
		log.WithFields(log.Fields{
			"tx_hash": resp.TxResponse.TxHash,
			"code":    resp.TxResponse.Code,
			"Error":   "insufficient fee",
		}).Infoln(hyperionId, " - sent MsgUnSetOrchestratorAddresses")
		return errors.Wrap(errors.New("code 13 - insufficient fee"), "broadcasting MsgUnSetOrchestratorAddresses failed")
	}

	time.Sleep(10 * time.Second)

	return nil
}

func (c broadcastClient) SendDepositClaim(ctx context.Context, hyperionId uint64, deposit *hyperionevents.HyperionSendToHeliosEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error) {
	// EthereumBridgeDepositClaim
	// When more than 66% of the active validator set has
	// claimed to have seen the deposit enter the ethereum blockchain coins are
	// issued to the Cosmos address in question
	// -------------
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"sender":         deposit.Sender.Hex(),
		"destination":    cosmostypes.AccAddress(deposit.Destination[12:32]).String(),
		"amount":         deposit.Amount.String(),
		"data":           deposit.Data,
		"token_contract": deposit.TokenContract.Hex(),
	}).Debugln(hyperionId, " - observed SendToHeliosEvent")

	// check if data is valid json
	if !json.Valid([]byte(deposit.Data)) {
		deposit.Data = ""
	}

	msg := &hyperiontypes.MsgDepositClaim{
		HyperionId:     hyperionId,
		EventNonce:     deposit.EventNonce.Uint64(),
		BlockHeight:    deposit.Raw.BlockNumber,
		TokenContract:  deposit.TokenContract.Hex(),
		Amount:         sdkmath.NewIntFromBigInt(deposit.Amount),
		EthereumSender: deposit.Sender.Hex(),
		CosmosReceiver: cosmostypes.AccAddress(deposit.Destination[12:32]).String(),
		Orchestrator:   c.ChainClient.FromAddress().String(),
		Data:           deposit.Data,
		TxHash:         deposit.Raw.TxHash.Hex(),
		RpcUsed:        rpcUsedForObservation,
	}

	log.WithFields(log.Fields{
		"event_nonce":  msg.EventNonce,
		"event_height": msg.BlockHeight,
		"data":         deposit.Data,
	}).Infoln(hyperionId, " - sending MsgDepositClaim")

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "broadcasting MsgDepositClaim failed")
	}

	// cost := big.NewInt(0)
	// gasFee, err := c.ChainClient.GetGasFee()//client.HeaderByNumber(ctx, nil)
	// if err == nil {
	// 	gasUsed := resp.TxResponse.GasUsed
	// 	// calculate the cost of the transaction
	// 	gasFeeBig := new(big.Int)
	// 	gasFeeBig, _ = gasFeeBig.SetString(gasFee, 10)
	// 	cost = cost.Mul(gasFeeBig, big.NewInt(gasUsed))
	// }

	if resp.TxResponse.Code == 13 {
		log.WithFields(log.Fields{
			"event_nonce":  msg.EventNonce,
			"event_height": msg.BlockHeight,
			"tx_hash":      resp.TxResponse.TxHash,
			"code":         resp.TxResponse.Code,
			"Error":        "insufficient fee",
		}).Infoln(hyperionId, " - sent MsgDepositClaim")
		return nil, errors.Wrap(errors.New("code 13 - insufficient fee"), "broadcasting MsgDepositClaim failed")
	}

	log.WithFields(log.Fields{
		"event_nonce":  msg.EventNonce,
		"event_height": msg.BlockHeight,
		"tx_hash":      resp.TxResponse.TxHash,
		"code":         resp.TxResponse.Code,
		"GasUsed":      resp.TxResponse.GasUsed,
	}).Infoln(hyperionId, " - sent MsgDepositClaim")

	return resp.TxResponse, nil
}

func (c broadcastClient) SendWithdrawalClaim(_ context.Context, hyperionId uint64, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"batch_nonce":    withdrawal.BatchNonce.String(),
		"token_contract": withdrawal.Token.Hex(),
	}).Debugln(hyperionId, " - observed TransactionBatchExecutedEvent")

	// WithdrawClaim claims that a batch of withdrawal
	// operations on the bridge contract was executed.
	// -------------
	msg := &hyperiontypes.MsgWithdrawClaim{
		HyperionId:    hyperionId,
		EventNonce:    withdrawal.EventNonce.Uint64(),
		BatchNonce:    withdrawal.BatchNonce.Uint64(),
		BlockHeight:   withdrawal.Raw.BlockNumber,
		TokenContract: withdrawal.Token.Hex(),
		Orchestrator:  c.FromAddress().String(),
		TxHash:        withdrawal.Raw.TxHash.Hex(),
		RpcUsed:       rpcUsedForObservation,
	}

	log.WithFields(log.Fields{
		"event_height": msg.BlockHeight,
		"event_nonce":  msg.EventNonce,
	}).Infoln(hyperionId, " - sending MsgWithdrawClaim")

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "broadcasting MsgWithdrawClaim failed")
	}

	log.WithFields(log.Fields{
		"event_height": msg.BlockHeight,
		"event_nonce":  msg.EventNonce,
		"tx_hash":      resp.TxResponse.TxHash,
	}).Infoln(hyperionId, " - sent MsgWithdrawClaim")

	return resp.TxResponse, nil
}

func (c broadcastClient) SendExternalDataClaim(ctx context.Context, hyperionId uint64, nonce uint64, blockHeight uint64, externalContractAddress string, callData []byte, callErr []byte, rpcUsedForObservation string) (*cosmostypes.TxResponse, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"tx_nonce":          nonce,
		"tx_height":         blockHeight,
		"external_contract": externalContractAddress,
	}).Debugln(hyperionId, " - observed ExternalDataClaim")

	// MsgExternalDataClaim claims that a batch of external data
	// was executed.
	// -------------
	msg := &hyperiontypes.MsgExternalDataClaim{
		HyperionId:              hyperionId,
		TxNonce:                 nonce,
		BlockHeight:             blockHeight,
		ExternalContractAddress: externalContractAddress,
		Orchestrator:            c.FromAddress().String(),
		CallDataResult:          hex.EncodeToString(callData),
		CallDataResultError:     hex.EncodeToString(callErr),
		RpcUsed:                 rpcUsedForObservation,
	}

	log.WithFields(log.Fields{
		"tx_height": msg.BlockHeight,
		"tx_nonce":  msg.TxNonce,
		"call_data": msg.CallDataResult,
	}).Infoln(hyperionId, " - sending MsgExternalDataClaim")

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "broadcasting MsgExternalDataClaim failed")
	}

	log.WithFields(log.Fields{
		"tx_height": msg.BlockHeight,
		"tx_nonce":  msg.TxNonce,
		"tx_hash":   resp.TxResponse.TxHash,
	}).Infoln(hyperionId, " - sent MsgExternalDataClaim")

	return resp.TxResponse, nil
}

func (c broadcastClient) SendValsetClaim(ctx context.Context, hyperionId uint64, vs *hyperionevents.HyperionValsetUpdatedEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"valset_nonce":  vs.NewValsetNonce.Uint64(),
		"validators":    vs.Validators,
		"powers":        vs.Powers,
		"reward_amount": vs.RewardAmount,
		"reward_token":  vs.RewardToken.Hex(),
	}).Debugln("observed ValsetUpdatedEvent")

	members := make([]*hyperiontypes.BridgeValidator, len(vs.Validators))
	for i, val := range vs.Validators {
		members[i] = &hyperiontypes.BridgeValidator{
			EthereumAddress: val.Hex(),
			Power:           vs.Powers[i].Uint64(),
		}
	}

	// MsgValsetUpdatedClaim this message permit to share to
	// hyperion module the valset was updated on source blockchain
	// this permit to insure the power is well share on both side
	// -------------
	msg := &hyperiontypes.MsgValsetUpdatedClaim{
		HyperionId:   hyperionId,
		EventNonce:   vs.EventNonce.Uint64(),
		ValsetNonce:  vs.NewValsetNonce.Uint64(),
		BlockHeight:  vs.Raw.BlockNumber,
		RewardAmount: sdkmath.NewIntFromBigInt(vs.RewardAmount),
		RewardToken:  vs.RewardToken.Hex(),
		Members:      members,
		Orchestrator: c.FromAddress().String(),
		RpcUsed:      rpcUsedForObservation,
	}

	log.WithFields(log.Fields{
		"event_nonce":  msg.EventNonce,
		"event_height": msg.BlockHeight,
		"claim_hash":   gethcommon.Bytes2Hex(msg.ClaimHash()),
	}).Infoln(hyperionId, " - sending MsgValsetUpdatedClaim")

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "broadcasting MsgValsetUpdatedClaim failed")
	}

	log.WithFields(log.Fields{
		"event_nonce":  msg.EventNonce,
		"event_height": msg.BlockHeight,
		"tx_hash":      resp.TxResponse.TxHash,
		"claim_hash":   gethcommon.Bytes2Hex(msg.ClaimHash()),
	}).Infoln(hyperionId, " - sent MsgValsetUpdatedClaim")

	// // Attendre que l'attestation soit observée avec un timeout de 5 minutes
	// if err := c.WaitForAttestation(ctx, msg.EventNonce, msg.ClaimHash(), 5*time.Minute); err != nil {
	// 	return errors.Wrap(err, "waiting for attestation to be observed")
	// }

	// log.WithFields(log.Fields{
	// 	"event_nonce":  msg.EventNonce,
	// 	"event_height": msg.BlockHeight,
	// 	"tx_hash":      resp.TxResponse.TxHash,
	// 	"claim_hash":   gethcommon.Bytes2Hex(msg.ClaimHash()),
	// }).Infoln("Oracle Comfirmed MsgValsetUpdatedClaim")

	return resp.TxResponse, nil
}

func (c broadcastClient) SendERC20DeployedClaim(_ context.Context, hyperionId uint64, erc20 *hyperionevents.HyperionERC20DeployedEvent, rpcUsedForObservation string) (*cosmostypes.TxResponse, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"helios_denom":   erc20.HeliosDenom,
		"token_contract": erc20.TokenContract.Hex(),
		"name":           erc20.Name,
		"symbol":         erc20.Symbol,
		"decimals":       erc20.Decimals,
	}).Debugln("observed ERC20DeployedEvent")

	// MsgERC20DeployedClaim claims that new erc20 token
	// was deployed on the source blockchain and will be linked
	// as ERC20 to cosmosDenom in hyperion Module on HyperionId
	// ----------
	msg := &hyperiontypes.MsgERC20DeployedClaim{
		HyperionId:    hyperionId,
		EventNonce:    erc20.EventNonce.Uint64(),
		BlockHeight:   erc20.Raw.BlockNumber,
		CosmosDenom:   erc20.HeliosDenom,
		TokenContract: erc20.TokenContract.Hex(),
		Name:          erc20.Name,
		Symbol:        erc20.Symbol,
		Decimals:      uint64(erc20.Decimals),
		Orchestrator:  c.FromAddress().String(),
		RpcUsed:       rpcUsedForObservation,
	}

	log.WithFields(log.Fields{
		"event_nonce":  msg.EventNonce,
		"event_height": msg.BlockHeight,
	}).Infoln(hyperionId, " - sending MsgERC20DeployedClaim")

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "broadcasting MsgERC20DeployedClaim failed")
	}

	log.WithFields(log.Fields{
		"event_nonce":  msg.EventNonce,
		"event_height": msg.BlockHeight,
		"tx_hash":      resp.TxResponse.TxHash,
	}).Infoln(hyperionId, " - sent MsgERC20DeployedClaim")

	return resp.TxResponse, nil
}

func (c broadcastClient) SendForceSetValsetAndLastObservedEventNonce(ctx context.Context, hyperionId uint64, nonce uint64, blockHeight uint64, valset *hyperiontypes.Valset) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &hyperiontypes.MsgForceSetValsetAndLastObservedEventNonce{
		HyperionId:                      hyperionId,
		Valset:                          valset,
		Signer:                          c.FromAddress().String(),
		LastObservedEventNonce:          nonce,
		LastObservedEthereumBlockHeight: blockHeight,
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgForceSetValset failed")
	}

	log.WithFields(log.Fields{
		"tx_hash": resp.TxResponse.TxHash,
	}).Infoln("Oracle sent MsgForceSetValset")

	time.Sleep(10 * time.Second)

	return nil
}

func (c broadcastClient) SendCancelAllPendingOutTx(ctx context.Context, chainId uint64) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &hyperiontypes.MsgCancelAllPendingOutgoingTxs{
		ChainId: chainId,
		Signer:  c.FromAddress().String(),
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgCancelAllPendingOutTx failed")
	}

	log.WithFields(log.Fields{
		"tx_hash": resp.TxResponse.TxHash,
	}).Infoln("Oracle sent MsgCancelAllPendingOutTx")

	return nil
}

func (c broadcastClient) SendCancelPendingOutTxs(ctx context.Context, chainId uint64, count uint64) error {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	msg := &hyperiontypes.MsgCancelPendingOutgoingTxs{
		ChainId: chainId,
		Signer:  c.FromAddress().String(),
		Count:   count,
	}

	resp, err := c.ChainClient.SyncBroadcastMsg(msg)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return errors.Wrap(err, "broadcasting MsgCancelAllPendingOutTx failed")
	}

	log.WithFields(log.Fields{
		"tx_hash": resp.TxResponse.TxHash,
	}).Infoln("Oracle sent MsgCancelAllPendingOutTx")

	return nil
}

// / Potential use for wait observed state of specifical claim
func (c broadcastClient) WaitForAttestation(ctx context.Context, eventNonce uint64, claimHash []byte, timeout time.Duration) error {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	deadline := time.Now().Add(timeout)

	for {
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for attestation to be observed (nonce: %d)", eventNonce)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			// Vérifier le statut de l'attestation
			att, err := c.ChainClient.GetAttestation(ctx, eventNonce, claimHash)
			if err != nil {
				log.WithError(err).Debugf("failed to get attestation status for nonce %d", eventNonce)
				continue
			}

			if att.Attestation.Observed {
				log.WithFields(log.Fields{
					"event_nonce": eventNonce,
				}).Infoln("Attestation has been observed")
				return nil
			}
		}
	}
}

func (c broadcastClient) SendDepositClaimMsg(ctx context.Context, hyperionId uint64, deposit *hyperionevents.HyperionSendToHeliosEvent, rpcUsedForObservation string) (cosmostypes.Msg, error) {
	// EthereumBridgeDepositClaim
	// When more than 66% of the active validator set has
	// claimed to have seen the deposit enter the ethereum blockchain coins are
	// issued to the Cosmos address in question
	// -------------
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"sender":         deposit.Sender.Hex(),
		"destination":    cosmostypes.AccAddress(deposit.Destination[12:32]).String(),
		"amount":         deposit.Amount.String(),
		"data":           deposit.Data,
		"token_contract": deposit.TokenContract.Hex(),
	}).Debugln(hyperionId, " - observed SendToHeliosEvent")

	// check if data is valid json
	if !json.Valid([]byte(deposit.Data)) {
		deposit.Data = ""
	}

	msg := &hyperiontypes.MsgDepositClaim{
		HyperionId:     hyperionId,
		EventNonce:     deposit.EventNonce.Uint64(),
		BlockHeight:    deposit.Raw.BlockNumber,
		TokenContract:  deposit.TokenContract.Hex(),
		Amount:         sdkmath.NewIntFromBigInt(deposit.Amount),
		EthereumSender: deposit.Sender.Hex(),
		CosmosReceiver: cosmostypes.AccAddress(deposit.Destination[12:32]).String(),
		Orchestrator:   c.ChainClient.FromAddress().String(),
		Data:           deposit.Data,
		TxHash:         deposit.Raw.TxHash.Hex(),
		RpcUsed:        rpcUsedForObservation,
	}

	return msg, nil
}

func (c broadcastClient) SendWithdrawalClaimMsg(_ context.Context, hyperionId uint64, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent, rpcUsedForObservation string) (cosmostypes.Msg, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"batch_nonce":    withdrawal.BatchNonce.String(),
		"token_contract": withdrawal.Token.Hex(),
	}).Debugln(hyperionId, " - observed TransactionBatchExecutedEvent")

	// WithdrawClaim claims that a batch of withdrawal
	// operations on the bridge contract was executed.
	// -------------
	msg := &hyperiontypes.MsgWithdrawClaim{
		HyperionId:    hyperionId,
		EventNonce:    withdrawal.EventNonce.Uint64(),
		BatchNonce:    withdrawal.BatchNonce.Uint64(),
		BlockHeight:   withdrawal.Raw.BlockNumber,
		TokenContract: withdrawal.Token.Hex(),
		Orchestrator:  c.FromAddress().String(),
		TxHash:        withdrawal.Raw.TxHash.Hex(),
		RpcUsed:       rpcUsedForObservation,
	}

	return msg, nil
}

func (c broadcastClient) SendExternalDataClaimMsg(ctx context.Context, hyperionId uint64, nonce uint64, blockHeight uint64, externalContractAddress string, callData []byte, callErr []byte, rpcUsedForObservation string) (cosmostypes.Msg, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"tx_nonce":          nonce,
		"tx_height":         blockHeight,
		"external_contract": externalContractAddress,
	}).Debugln(hyperionId, " - observed ExternalDataClaim")

	// MsgExternalDataClaim claims that a batch of external data
	// was executed.
	// -------------
	msg := &hyperiontypes.MsgExternalDataClaim{
		HyperionId:              hyperionId,
		TxNonce:                 nonce,
		BlockHeight:             blockHeight,
		ExternalContractAddress: externalContractAddress,
		Orchestrator:            c.FromAddress().String(),
		CallDataResult:          hex.EncodeToString(callData),
		CallDataResultError:     hex.EncodeToString(callErr),
		RpcUsed:                 rpcUsedForObservation,
	}

	return msg, nil
}

func (c broadcastClient) SendValsetClaimMsg(ctx context.Context, hyperionId uint64, vs *hyperionevents.HyperionValsetUpdatedEvent, rpcUsedForObservation string) (cosmostypes.Msg, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"valset_nonce":  vs.NewValsetNonce.Uint64(),
		"validators":    vs.Validators,
		"powers":        vs.Powers,
		"reward_amount": vs.RewardAmount,
		"reward_token":  vs.RewardToken.Hex(),
	}).Debugln("observed ValsetUpdatedEvent")

	members := make([]*hyperiontypes.BridgeValidator, len(vs.Validators))
	for i, val := range vs.Validators {
		members[i] = &hyperiontypes.BridgeValidator{
			EthereumAddress: val.Hex(),
			Power:           vs.Powers[i].Uint64(),
		}
	}

	// MsgValsetUpdatedClaim this message permit to share to
	// hyperion module the valset was updated on source blockchain
	// this permit to insure the power is well share on both side
	// -------------
	msg := &hyperiontypes.MsgValsetUpdatedClaim{
		HyperionId:   hyperionId,
		EventNonce:   vs.EventNonce.Uint64(),
		ValsetNonce:  vs.NewValsetNonce.Uint64(),
		BlockHeight:  vs.Raw.BlockNumber,
		RewardAmount: sdkmath.NewIntFromBigInt(vs.RewardAmount),
		RewardToken:  vs.RewardToken.Hex(),
		Members:      members,
		Orchestrator: c.FromAddress().String(),
		RpcUsed:      rpcUsedForObservation,
	}
	return msg, nil
}

func (c broadcastClient) SendERC20DeployedClaimMsg(_ context.Context, hyperionId uint64, erc20 *hyperionevents.HyperionERC20DeployedEvent, rpcUsedForObservation string) (cosmostypes.Msg, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	log.WithFields(log.Fields{
		"helios_denom":   erc20.HeliosDenom,
		"token_contract": erc20.TokenContract.Hex(),
		"name":           erc20.Name,
		"symbol":         erc20.Symbol,
		"decimals":       erc20.Decimals,
	}).Debugln("observed ERC20DeployedEvent")

	// MsgERC20DeployedClaim claims that new erc20 token
	// was deployed on the source blockchain and will be linked
	// as ERC20 to cosmosDenom in hyperion Module on HyperionId
	// ----------
	msg := &hyperiontypes.MsgERC20DeployedClaim{
		HyperionId:    hyperionId,
		EventNonce:    erc20.EventNonce.Uint64(),
		BlockHeight:   erc20.Raw.BlockNumber,
		CosmosDenom:   erc20.HeliosDenom,
		TokenContract: erc20.TokenContract.Hex(),
		Name:          erc20.Name,
		Symbol:        erc20.Symbol,
		Decimals:      uint64(erc20.Decimals),
		Orchestrator:  c.FromAddress().String(),
		RpcUsed:       rpcUsedForObservation,
	}

	return msg, nil
}

func (c broadcastClient) SyncBroadcastMsgs(ctx context.Context, msgs []cosmostypes.Msg) (*cosmostypes.TxResponse, error) {
	metrics.ReportFuncCall(c.svcTags)
	doneFn := metrics.ReportFuncTiming(c.svcTags)
	defer doneFn()

	resp, err := c.ChainClient.SyncBroadcastMsg(msgs...)
	if err != nil {
		metrics.ReportFuncError(c.svcTags)
		return nil, errors.Wrap(err, "broadcasting Msgs failed")
	}

	return resp.TxResponse, nil
}
