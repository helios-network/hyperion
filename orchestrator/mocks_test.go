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

	hyperionevents "github.com/Helios-Chain-Labs/peggo/solidity/wrappers/Hyperion.sol"
	hyperionsubgraphevents "github.com/Helios-Chain-Labs/peggo/solidity/wrappers/HyperionSubgraph.sol"
	hyperiontypes "github.com/Helios-Chain-Labs/sdk-go/chain/peggy/types"
)

type MockPriceFeed struct {
	QueryUSDPriceFn func(gethcommon.Address) (float64, error)
}

func (p MockPriceFeed) QueryUSDPrice(address gethcommon.Address) (float64, error) {
	return p.QueryUSDPriceFn(address)
}

type MockCosmosNetwork struct {
	HyperionParamsFn                      func(ctx context.Context) (*hyperiontypes.Params, error)
	LastClaimEventByAddrFn                func(ctx context.Context, address cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error)
	GetValidatorAddressFn                 func(ctx context.Context, address gethcommon.Address) (cosmostypes.AccAddress, error)
	CurrentValsetFn                       func(ctx context.Context) (*hyperiontypes.Valset, error)
	ValsetAtFn                            func(ctx context.Context, uint642 uint64) (*hyperiontypes.Valset, error)
	OldestUnsignedValsetsFn               func(ctx context.Context, address cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error)
	LatestValsetsFn                       func(ctx context.Context) ([]*hyperiontypes.Valset, error)
	AllValsetConfirmsFn                   func(ctx context.Context, uint642 uint64) ([]*hyperiontypes.MsgValsetConfirm, error)
	OldestUnsignedTransactionBatchFn      func(ctx context.Context, address cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error)
	LatestTransactionBatchesFn            func(ctx context.Context) ([]*hyperiontypes.OutgoingTxBatch, error)
	UnbatchedTokensWithFeesFn             func(ctx context.Context) ([]*hyperiontypes.BatchFees, error)
	TransactionBatchSignaturesFn          func(ctx context.Context, uint642 uint64, address gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error)
	UpdateHyperionOrchestratorAddressesFn func(ctx context.Context, address gethcommon.Address, address2 cosmostypes.Address) error
	SendValsetConfirmFn                   func(ctx context.Context, address gethcommon.Address, hash gethcommon.Hash, valset *hyperiontypes.Valset) error
	SendBatchConfirmFn                    func(ctx context.Context, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, batch *hyperiontypes.OutgoingTxBatch) error
	SendRequestBatchFn                    func(ctx context.Context, denom string) error
	SendToEthFn                           func(ctx context.Context, destination gethcommon.Address, amount, fee cosmostypes.Coin) error
	SendOldDepositClaimFn                 func(ctx context.Context, deposit *hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent) error
	SendDepositClaimFn                    func(ctx context.Context, deposit *hyperionevents.HyperionSendToHeliosEvent) error
	SendWithdrawalClaimFn                 func(ctx context.Context, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent) error
	SendValsetClaimFn                     func(ctx context.Context, vs *hyperionevents.HyperionValsetUpdatedEvent) error
	SendERC20DeployedClaimFn              func(ctx context.Context, erc20 *hyperionevents.HyperionERC20DeployedEvent) error
	GetBlockFn                            func(ctx context.Context, height int64) (*cometrpc.ResultBlock, error)
	GetLatestBlockHeightFn                func(ctx context.Context) (int64, error)
}

func (n MockCosmosNetwork) HyperionParams(ctx context.Context) (*hyperiontypes.Params, error) {
	return n.HyperionParamsFn(ctx)
}

func (n MockCosmosNetwork) LastClaimEventByAddr(ctx context.Context, validatorAccountAddress cosmostypes.AccAddress) (*hyperiontypes.LastClaimEvent, error) {
	return n.LastClaimEventByAddrFn(ctx, validatorAccountAddress)
}

func (n MockCosmosNetwork) GetValidatorAddress(ctx context.Context, addr gethcommon.Address) (cosmostypes.AccAddress, error) {
	return n.GetValidatorAddressFn(ctx, addr)
}

func (n MockCosmosNetwork) ValsetAt(ctx context.Context, nonce uint64) (*hyperiontypes.Valset, error) {
	return n.ValsetAtFn(ctx, nonce)
}

func (n MockCosmosNetwork) CurrentValset(ctx context.Context) (*hyperiontypes.Valset, error) {
	return n.CurrentValsetFn(ctx)
}

func (n MockCosmosNetwork) OldestUnsignedValsets(ctx context.Context, valAccountAddress cosmostypes.AccAddress) ([]*hyperiontypes.Valset, error) {
	return n.OldestUnsignedValsetsFn(ctx, valAccountAddress)
}

func (n MockCosmosNetwork) LatestValsets(ctx context.Context) ([]*hyperiontypes.Valset, error) {
	return n.LatestValsetsFn(ctx)
}

func (n MockCosmosNetwork) AllValsetConfirms(ctx context.Context, nonce uint64) ([]*hyperiontypes.MsgValsetConfirm, error) {
	return n.AllValsetConfirmsFn(ctx, nonce)
}

func (n MockCosmosNetwork) OldestUnsignedTransactionBatch(ctx context.Context, valAccountAddress cosmostypes.AccAddress) (*hyperiontypes.OutgoingTxBatch, error) {
	return n.OldestUnsignedTransactionBatchFn(ctx, valAccountAddress)
}

func (n MockCosmosNetwork) LatestTransactionBatches(ctx context.Context) ([]*hyperiontypes.OutgoingTxBatch, error) {
	return n.LatestTransactionBatchesFn(ctx)
}

func (n MockCosmosNetwork) UnbatchedTokensWithFees(ctx context.Context) ([]*hyperiontypes.BatchFees, error) {
	return n.UnbatchedTokensWithFeesFn(ctx)
}

func (n MockCosmosNetwork) TransactionBatchSignatures(ctx context.Context, nonce uint64, tokenContract gethcommon.Address) ([]*hyperiontypes.MsgConfirmBatch, error) {
	return n.TransactionBatchSignaturesFn(ctx, nonce, tokenContract)
}

func (n MockCosmosNetwork) UpdateHyperionOrchestratorAddresses(ctx context.Context, ethFrom gethcommon.Address, orchAddr cosmostypes.AccAddress) error {
	return n.UpdateHyperionOrchestratorAddressesFn(ctx, ethFrom, orchAddr)
}

func (n MockCosmosNetwork) SendValsetConfirm(ctx context.Context, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, valset *hyperiontypes.Valset) error {
	return n.SendValsetConfirmFn(ctx, ethFrom, hyperionID, valset)
}

func (n MockCosmosNetwork) SendBatchConfirm(ctx context.Context, ethFrom gethcommon.Address, hyperionID gethcommon.Hash, batch *hyperiontypes.OutgoingTxBatch) error {
	return n.SendBatchConfirmFn(ctx, ethFrom, hyperionID, batch)
}

func (n MockCosmosNetwork) SendRequestBatch(ctx context.Context, denom string) error {
	return n.SendRequestBatchFn(ctx, denom)
}

func (n MockCosmosNetwork) SendToEth(ctx context.Context, destination gethcommon.Address, amount, fee cosmostypes.Coin) error {
	return n.SendToEthFn(ctx, destination, amount, fee)
}

func (n MockCosmosNetwork) SendOldDepositClaim(ctx context.Context, deposit *hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent) error {
	return n.SendOldDepositClaimFn(ctx, deposit)
}

func (n MockCosmosNetwork) SendDepositClaim(ctx context.Context, deposit *hyperionevents.HyperionSendToHeliosEvent) error {
	return n.SendDepositClaimFn(ctx, deposit)
}

func (n MockCosmosNetwork) SendWithdrawalClaim(ctx context.Context, withdrawal *hyperionevents.HyperionTransactionBatchExecutedEvent) error {
	return n.SendWithdrawalClaimFn(ctx, withdrawal)
}

func (n MockCosmosNetwork) SendValsetClaim(ctx context.Context, vs *hyperionevents.HyperionValsetUpdatedEvent) error {
	return n.SendValsetClaimFn(ctx, vs)
}

func (n MockCosmosNetwork) SendERC20DeployedClaim(ctx context.Context, erc20 *hyperionevents.HyperionERC20DeployedEvent) error {
	return n.SendERC20DeployedClaimFn(ctx, erc20)
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
	GetSendToCosmosEventsFn             func(startBlock, endBlock uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error)
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

func (n MockEthereumNetwork) GetSendToCosmosEvents(startBlock, endBlock uint64) ([]*hyperionsubgraphevents.HyperionSubgraphSendToCosmosEvent, error) {
	return n.GetSendToCosmosEventsFn(startBlock, endBlock)
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
