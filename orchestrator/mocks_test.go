package orchestrator

import (
	"context"
	"math/big"
	"time"

	cometrpc "github.com/cometbft/cometbft/rpc/core/types"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	log "github.com/xlab/suplog"

	hyperionevents "github.com/Helios-Chain-Labs/hyperion/solidity/wrappers/Hyperion.sol"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/hyperion/types"
)

type MockPriceFeed struct {
	QueryUSDPriceFn func(gethcommon.Address) (float64, error)
}

func (p MockPriceFeed) QueryUSDPrice(address gethcommon.Address) (float64, error) {
	return p.QueryUSDPriceFn(address)
}

type MockCosmosNetwork struct {
	HyperionParamsFn                 func(ctx context.Context) (*hyperiontypes.Params, error)
	LastClaimEventByAddrFn           func(ctx context.Context, hyperionId uint64, address cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error)
	GetValidatorAddressFn            func(ctx context.Context, hyperionId uint64, address gethcommon.Address) (cosmostypes.AccAddress, error)
	CurrentValsetFn                  func(ctx context.Context, hyperionId uint64) (*hyperiontypes.Valset, error)
	ValsetAtFn                       func(ctx context.Context, hyperionId uint64, nonce uint64) (*hyperiontypes.Valset, error)
	OldestUnsignedValsetsFn          func(ctx context.Context, hyperionId uint64, address cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error)
	LatestValsetsFn                  func(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.Valset, error)
	AllValsetConfirmsFn              func(ctx context.Context, hyperionId uint64, uint642 uint64) ([]*hyperiontypes.MsgValsetConfirm, error)
	OldestUnsignedTransactionBatchFn func(ctx context.Context, hyperionId uint64, address cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error)
	LatestTransactionBatchesFn       func(ctx context.Context, hyperionID uint64) ([]*hyperiontypes.OutgoingTxBatch, error)
	UnbatchedTokensWithFeesFn        func(ctx context.Context) ([]*hyperiontypes.BatchFees, error)
	TransactionBatchSignaturesFn     func(ctx context.Context, hyperionID uint64, nonce uint64, address gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error)
	SendValsetConfirmFn              func(ctx context.Context, hyperionID uint64, address gethcommon.Address, hash gethcommon.Hash, valset *hyperiontypes.Valset) error
	SendBatchConfirmFn               func(ctx context.Context, hyperionID2 uint64, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, batch *hyperiontypes.OutgoingTxBatch) error
	SendRequestBatchFn               func(ctx context.Context, hyperionID uint64, denom string) error
	SendToChainFn                    func(ctx context.Context, hyperionId uint64, destination gethcommon.Address, amount, fee cosmostypes.Coin) error
	SendDepositClaimFn               func(ctx context.Context, hyperionID uint64, deposit *hyperionevents.HyperionSendToHeliosEvent) error
	SendWithdrawalClaimFn            func(ctx context.Context, hyperionID uint64, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent) error
	SendValsetClaimFn                func(ctx context.Context, hyperionID uint64, vs *hyperionevents.HyperionValsetUpdatedEvent) error
	SendERC20DeployedClaimFn         func(ctx context.Context, hyperionID uint64, erc20 *hyperionevents.HyperionERC20DeployedEvent) error
	GetBlockFn                       func(ctx context.Context, height int64) (*cometrpc.ResultBlock, error)
	GetLatestBlockHeightFn           func(ctx context.Context) (int64, error)
}

func (n MockCosmosNetwork) HyperionParams(ctx context.Context) (*hyperiontypes.Params, error) {
	return n.HyperionParamsFn(ctx)
}

func (n MockCosmosNetwork) LastClaimEventByAddr(ctx context.Context, hyperionId uint64, validatorAccountAddress cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
	return n.LastClaimEventByAddrFn(ctx, hyperionId, validatorAccountAddress)
}

func (n MockCosmosNetwork) GetValidatorAddress(ctx context.Context, hyperionId uint64, addr gethcommon.Address) (cosmostypes.AccAddress, error) {
	return n.GetValidatorAddressFn(ctx, hyperionId, addr)
}

func (n MockCosmosNetwork) ValsetAt(ctx context.Context, hyperionId uint64, nonce uint64) (*hyperiontypes.Valset, error) {
	return n.ValsetAtFn(ctx, hyperionId, nonce)
}

func (n MockCosmosNetwork) CurrentValset(ctx context.Context, hyperionId uint64) (*hyperiontypes.Valset, error) {
	return n.CurrentValsetFn(ctx, hyperionId)
}

func (n MockCosmosNetwork) OldestUnsignedValsets(ctx context.Context, hyperionId uint64, valAccountAddress cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error) {
	return n.OldestUnsignedValsetsFn(ctx, hyperionId, valAccountAddress)
}

func (n MockCosmosNetwork) LatestValsets(ctx context.Context, hyperionId uint64) ([]*hyperiontypes.Valset, error) {
	return n.LatestValsetsFn(ctx, hyperionId)
}

func (n MockCosmosNetwork) AllValsetConfirms(ctx context.Context, hyperionId uint64, nonce uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
	return n.AllValsetConfirmsFn(ctx, hyperionId, nonce)
}

func (n MockCosmosNetwork) OldestUnsignedTransactionBatch(ctx context.Context, hyperionID uint64, valAccountAddress cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error) {
	return n.OldestUnsignedTransactionBatchFn(ctx, hyperionID, valAccountAddress)
}

func (n MockCosmosNetwork) LatestTransactionBatches(ctx context.Context, hyperionID uint64) ([]*hyperiontypes.OutgoingTxBatch, error) {
	return n.LatestTransactionBatchesFn(ctx, hyperionID)
}

func (n MockCosmosNetwork) UnbatchedTokensWithFees(ctx context.Context) ([]*hyperiontypes.BatchFees, error) {
	return n.UnbatchedTokensWithFeesFn(ctx)
}

func (n MockCosmosNetwork) TransactionBatchSignatures(ctx context.Context, hyperionID uint64, nonce uint64, tokenContract gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
	return n.TransactionBatchSignaturesFn(ctx, hyperionID, nonce, tokenContract)
}

func (n MockCosmosNetwork) SendValsetConfirm(ctx context.Context, hyperionID uint64, ethFrom gethcommon.Address, hash gethcommon.Hash, valset *hyperiontypes.Valset) error {
	return n.SendValsetConfirmFn(ctx, hyperionID, ethFrom, hash, valset)
}

func (n MockCosmosNetwork) SendBatchConfirm(ctx context.Context, hyperionID uint64, ethFrom gethcommon.Address, hash gethcommon.Hash, batch *hyperiontypes.OutgoingTxBatch) error {
	return n.SendBatchConfirmFn(ctx, hyperionID, ethFrom, hash, batch)
}

func (n MockCosmosNetwork) SendRequestBatch(ctx context.Context, hyperionID uint64, denom string) error {
	return n.SendRequestBatchFn(ctx, hyperionID, denom)
}

func (n MockCosmosNetwork) SendToChain(ctx context.Context, hyperionID uint64, destination gethcommon.Address, amount, fee cosmostypes.Coin) error {
	return n.SendToChainFn(ctx, hyperionID, destination, amount, fee)
}

func (n MockCosmosNetwork) SendDepositClaim(ctx context.Context, hyperionID uint64, deposit *hyperionevents.HyperionSendToHeliosEvent) error {
	return n.SendDepositClaimFn(ctx, hyperionID, deposit)
}

func (n MockCosmosNetwork) SendWithdrawalClaim(ctx context.Context, hyperionID uint64, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent) error {
	return n.SendWithdrawalClaimFn(ctx, hyperionID, withdrawal)
}

func (n MockCosmosNetwork) SendValsetClaim(ctx context.Context, hyperionID uint64, vs *hyperionevents.HyperionValsetUpdatedEvent) error {
	return n.SendValsetClaimFn(ctx, hyperionID, vs)
}

func (n MockCosmosNetwork) SendERC20DeployedClaim(ctx context.Context, hyperionID uint64, erc20 *hyperionevents.HyperionERC20DeployedEvent) error {
	return n.SendERC20DeployedClaimFn(ctx, hyperionID, erc20)
}

func (n MockCosmosNetwork) GetBlock(ctx context.Context, height int64) (*cometrpc.ResultBlock, error) {
	return n.GetBlockFn(ctx, height)
}

func (n MockCosmosNetwork) GetLatestBlockHeight(ctx context.Context) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (n MockCosmosNetwork) GetTxs(ctx context.Context, block *cometrpc.ResultBlock) ([]*cometrpc.ResultTx, error) {
	//TODO implement me
	panic("implement me")
}

func (n MockCosmosNetwork) GetValidatorSet(ctx context.Context, height int64) (*cometrpc.ResultValidators, error) {
	//TODO implement me
	panic("implement me")
}

type MockEthereumNetwork struct {
	GetHeaderByNumberFn                 func(ctx context.Context, number *big.Int) (*gethtypes.Header, error)
	GetHyperionIDFn                     func(ctx context.Context) (gethcommon.Hash, error)
	GetSendToHeliosEventsFn             func(startBlock, endBlock uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error)
	GetHyperionERC20DeployedEventsFn    func(startBlock, endBlock uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error)
	GetValsetUpdatedEventsFn            func(startBlock, endBlock uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error)
	GetTransactionBatchExecutedEventsFn func(startBlock, endBlock uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error)
	GetValsetNonceFn                    func(ctx context.Context) (*big.Int, error)
	SendEthValsetUpdateFn               func(ctx context.Context, oldValset *hyperiontypes.Valset, newValset *hyperiontypes.Valset, confirms []*hyperiontypes.MsgValsetConfirm) (*gethcommon.Hash, error)
	GetTxBatchNonceFn                   func(ctx context.Context, erc20ContractAddress gethcommon.Address) (*big.Int, error)
	SendTransactionBatchFn              func(ctx context.Context, currentValset *hyperiontypes.Valset, batch *hyperiontypes.OutgoingTxBatch, confirms []*hyperiontypes.MsgConfirmBatch) (*gethcommon.Hash, error)
	TokenDecimalsFn                     func(ctx context.Context, address gethcommon.Address) (uint8, error)
}

func (n MockEthereumNetwork) GetHeaderByNumber(ctx context.Context, number *big.Int) (*gethtypes.Header, error) {
	return n.GetHeaderByNumberFn(ctx, number)
}

func (n MockEthereumNetwork) TokenDecimals(ctx context.Context, tokenContract gethcommon.Address) (uint8, error) {
	return n.TokenDecimalsFn(ctx, tokenContract)
}

func (n MockEthereumNetwork) GetHyperionID(ctx context.Context) (gethcommon.Hash, error) {
	return n.GetHyperionIDFn(ctx)
}

func (n MockEthereumNetwork) GetSendToHeliosEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionSendToHeliosEvent, error) {
	return n.GetSendToHeliosEventsFn(startBlock, endBlock)
}

func (n MockEthereumNetwork) GetHyperionERC20DeployedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionERC20DeployedEvent, error) {
	return n.GetHyperionERC20DeployedEventsFn(startBlock, endBlock)
}

func (n MockEthereumNetwork) GetValsetUpdatedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionValsetUpdatedEvent, error) {
	return n.GetValsetUpdatedEventsFn(startBlock, endBlock)
}

func (n MockEthereumNetwork) GetTransactionBatchExecutedEvents(startBlock, endBlock uint64) ([]*hyperionevents.HyperionTransactionBatchExecutedEvent, error) {
	return n.GetTransactionBatchExecutedEventsFn(startBlock, endBlock)
}

func (n MockEthereumNetwork) GetValsetNonce(ctx context.Context) (*big.Int, error) {
	return n.GetValsetNonceFn(ctx)
}

func (n MockEthereumNetwork) SendEthValsetUpdate(ctx context.Context, oldValset *hyperiontypes.Valset, newValset *hyperiontypes.Valset, confirms []*hyperiontypes.MsgValsetConfirm) (*gethcommon.Hash, error) {
	return n.SendEthValsetUpdateFn(ctx, oldValset, newValset, confirms)
}

func (n MockEthereumNetwork) GetTxBatchNonce(ctx context.Context, erc20ContractAddress gethcommon.Address) (*big.Int, error) {
	return n.GetTxBatchNonceFn(ctx, erc20ContractAddress)
}

func (n MockEthereumNetwork) SendTransactionBatch(ctx context.Context, currentValset *hyperiontypes.Valset, batch *hyperiontypes.OutgoingTxBatch, confirms []*hyperiontypes.MsgConfirmBatch) (*gethcommon.Hash, error) {
	return n.SendTransactionBatchFn(ctx, currentValset, batch, confirms)
}

var (
	DummyLog = DummyLogger{}
)

type DummyLogger struct{}

func (l DummyLogger) Success(format string, args ...interface{}) {
}

func (l DummyLogger) Warning(format string, args ...interface{}) {
}

func (l DummyLogger) Error(format string, args ...interface{}) {
}

func (l DummyLogger) Debug(format string, args ...interface{}) {
}

func (l DummyLogger) WithField(key string, value interface{}) log.Logger {
	return l
}

func (l DummyLogger) WithFields(fields log.Fields) log.Logger {
	return l
}

func (l DummyLogger) WithError(err error) log.Logger {
	return l
}

func (l DummyLogger) WithContext(ctx context.Context) log.Logger {
	return l
}

func (l DummyLogger) WithTime(t time.Time) log.Logger {
	return l
}

func (l DummyLogger) Logf(level log.Level, format string, args ...interface{}) {
}

func (l DummyLogger) Tracef(format string, args ...interface{}) {
}

func (l DummyLogger) Debugf(format string, args ...interface{}) {
}

func (l DummyLogger) Infof(format string, args ...interface{}) {
}

func (l DummyLogger) Printf(format string, args ...interface{}) {
}

func (l DummyLogger) Warningf(format string, args ...interface{}) {
}

func (l DummyLogger) Errorf(format string, args ...interface{}) {
}

func (l DummyLogger) Fatalf(format string, args ...interface{}) {
}

func (l DummyLogger) Panicf(format string, args ...interface{}) {
}

func (l DummyLogger) Log(level log.Level, args ...interface{}) {
}

func (l DummyLogger) Trace(args ...interface{}) {
}

func (l DummyLogger) Info(args ...interface{}) {
}

func (l DummyLogger) Print(args ...interface{}) {
}

func (l DummyLogger) Fatal(args ...interface{}) {
}

func (l DummyLogger) Panic(args ...interface{}) {
}

func (l DummyLogger) Logln(level log.Level, args ...interface{}) {
}

func (l DummyLogger) Traceln(args ...interface{}) {
}

func (l DummyLogger) Debugln(args ...interface{}) {
}

func (l DummyLogger) Infoln(args ...interface{}) {
}

func (l DummyLogger) Println(args ...interface{}) {
}

func (l DummyLogger) Warningln(args ...interface{}) {
}

func (l DummyLogger) Errorln(args ...interface{}) {
}

func (l DummyLogger) Fatalln(args ...interface{}) {
}

func (l DummyLogger) Panicln(args ...interface{}) {
}
