// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wrappers

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// HyperionPrecompileMetaData contains all meta data concerning the HyperionPrecompile contract.
var HyperionPrecompileMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"hyperionId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"eventNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockHeight\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"tokenContract\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ethereumSender\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cosmosReceiver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orchestrator\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"txHash\",\"type\":\"string\"}],\"name\":\"depositClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"orchestratorAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"hyperionId\",\"type\":\"uint64\"}],\"name\":\"setOrchestratorAddresses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460155761026d908161001a8239f35b5f80fdfe60808060405260049081361015610014575f80fd5b5f3560e01c9081635c1ea75714610144575063b998366a14610034575f80fd5b34610140576101403660031901126101405767ffffffffffffffff9080358281160361014057610062610178565b50604435828116036101405760643582811161014057610085903690830161018f565b5060a4358281116101405761009d903690830161018f565b5060c435828111610140576100b5903690830161018f565b5060e435828111610140576100cd903690830161018f565b5061010435828111610140576100e6903690830161018f565b5061012435918211610140576100fe9136910161018f565b50610107610205565b50610110610205565b50610119610205565b50610122610205565b5061012b610205565b50610134610205565b50602060405160018152f35b5f80fd5b823461014057604036600319011261014057356001600160a01b0381160361014057602090610171610178565b5060018152f35b6024359067ffffffffffffffff8216820361014057565b81601f820112156101405780359067ffffffffffffffff928383116101f15760405193601f8401601f19908116603f01168501908111858210176101f1576040528284526020838301011161014057815f926020809301838601378301015290565b634e487b7160e01b5f52604160045260245ffd5b604051906040820182811067ffffffffffffffff8211176101f1576040526005825264307831323360d81b602083015256fea26469706673582212206967269123dcacbd2004e17ecb7a4f620c818090170bd7be61bd00753aca92c864736f6c63430008190033",
}

// HyperionPrecompileABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperionPrecompileMetaData.ABI instead.
var HyperionPrecompileABI = HyperionPrecompileMetaData.ABI

// HyperionPrecompileBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HyperionPrecompileMetaData.Bin instead.
var HyperionPrecompileBin = HyperionPrecompileMetaData.Bin

// DeployHyperionPrecompile deploys a new Ethereum contract, binding an instance of HyperionPrecompile to it.
func DeployHyperionPrecompile(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HyperionPrecompile, error) {
	parsed, err := HyperionPrecompileMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HyperionPrecompileBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HyperionPrecompile{HyperionPrecompileCaller: HyperionPrecompileCaller{contract: contract}, HyperionPrecompileTransactor: HyperionPrecompileTransactor{contract: contract}, HyperionPrecompileFilterer: HyperionPrecompileFilterer{contract: contract}}, nil
}

// HyperionPrecompile is an auto generated Go binding around an Ethereum contract.
type HyperionPrecompile struct {
	HyperionPrecompileCaller     // Read-only binding to the contract
	HyperionPrecompileTransactor // Write-only binding to the contract
	HyperionPrecompileFilterer   // Log filterer for contract events
}

// HyperionPrecompileCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperionPrecompileCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionPrecompileTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperionPrecompileTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionPrecompileFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperionPrecompileFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionPrecompileSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperionPrecompileSession struct {
	Contract     *HyperionPrecompile // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HyperionPrecompileCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperionPrecompileCallerSession struct {
	Contract *HyperionPrecompileCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// HyperionPrecompileTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperionPrecompileTransactorSession struct {
	Contract     *HyperionPrecompileTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// HyperionPrecompileRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperionPrecompileRaw struct {
	Contract *HyperionPrecompile // Generic contract binding to access the raw methods on
}

// HyperionPrecompileCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperionPrecompileCallerRaw struct {
	Contract *HyperionPrecompileCaller // Generic read-only contract binding to access the raw methods on
}

// HyperionPrecompileTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperionPrecompileTransactorRaw struct {
	Contract *HyperionPrecompileTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperionPrecompile creates a new instance of HyperionPrecompile, bound to a specific deployed contract.
func NewHyperionPrecompile(address common.Address, backend bind.ContractBackend) (*HyperionPrecompile, error) {
	contract, err := bindHyperionPrecompile(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HyperionPrecompile{HyperionPrecompileCaller: HyperionPrecompileCaller{contract: contract}, HyperionPrecompileTransactor: HyperionPrecompileTransactor{contract: contract}, HyperionPrecompileFilterer: HyperionPrecompileFilterer{contract: contract}}, nil
}

// NewHyperionPrecompileCaller creates a new read-only instance of HyperionPrecompile, bound to a specific deployed contract.
func NewHyperionPrecompileCaller(address common.Address, caller bind.ContractCaller) (*HyperionPrecompileCaller, error) {
	contract, err := bindHyperionPrecompile(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionPrecompileCaller{contract: contract}, nil
}

// NewHyperionPrecompileTransactor creates a new write-only instance of HyperionPrecompile, bound to a specific deployed contract.
func NewHyperionPrecompileTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperionPrecompileTransactor, error) {
	contract, err := bindHyperionPrecompile(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionPrecompileTransactor{contract: contract}, nil
}

// NewHyperionPrecompileFilterer creates a new log filterer instance of HyperionPrecompile, bound to a specific deployed contract.
func NewHyperionPrecompileFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperionPrecompileFilterer, error) {
	contract, err := bindHyperionPrecompile(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperionPrecompileFilterer{contract: contract}, nil
}

// bindHyperionPrecompile binds a generic wrapper to an already deployed contract.
func bindHyperionPrecompile(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HyperionPrecompileMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperionPrecompile *HyperionPrecompileRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperionPrecompile.Contract.HyperionPrecompileCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperionPrecompile *HyperionPrecompileRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.HyperionPrecompileTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperionPrecompile *HyperionPrecompileRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.HyperionPrecompileTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperionPrecompile *HyperionPrecompileCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperionPrecompile.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperionPrecompile *HyperionPrecompileTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperionPrecompile *HyperionPrecompileTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.contract.Transact(opts, method, params...)
}

// DepositClaim is a paid mutator transaction binding the contract method 0xb998366a.
//
// Solidity: function depositClaim(uint64 hyperionId, uint64 eventNonce, uint64 blockHeight, string tokenContract, uint256 amount, string ethereumSender, string cosmosReceiver, string orchestrator, string data, string txHash) returns(bool success)
func (_HyperionPrecompile *HyperionPrecompileTransactor) DepositClaim(opts *bind.TransactOpts, hyperionId uint64, eventNonce uint64, blockHeight uint64, tokenContract string, amount *big.Int, ethereumSender string, cosmosReceiver string, orchestrator string, data string, txHash string) (*types.Transaction, error) {
	return _HyperionPrecompile.contract.Transact(opts, "depositClaim", hyperionId, eventNonce, blockHeight, tokenContract, amount, ethereumSender, cosmosReceiver, orchestrator, data, txHash)
}

// DepositClaim is a paid mutator transaction binding the contract method 0xb998366a.
//
// Solidity: function depositClaim(uint64 hyperionId, uint64 eventNonce, uint64 blockHeight, string tokenContract, uint256 amount, string ethereumSender, string cosmosReceiver, string orchestrator, string data, string txHash) returns(bool success)
func (_HyperionPrecompile *HyperionPrecompileSession) DepositClaim(hyperionId uint64, eventNonce uint64, blockHeight uint64, tokenContract string, amount *big.Int, ethereumSender string, cosmosReceiver string, orchestrator string, data string, txHash string) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.DepositClaim(&_HyperionPrecompile.TransactOpts, hyperionId, eventNonce, blockHeight, tokenContract, amount, ethereumSender, cosmosReceiver, orchestrator, data, txHash)
}

// DepositClaim is a paid mutator transaction binding the contract method 0xb998366a.
//
// Solidity: function depositClaim(uint64 hyperionId, uint64 eventNonce, uint64 blockHeight, string tokenContract, uint256 amount, string ethereumSender, string cosmosReceiver, string orchestrator, string data, string txHash) returns(bool success)
func (_HyperionPrecompile *HyperionPrecompileTransactorSession) DepositClaim(hyperionId uint64, eventNonce uint64, blockHeight uint64, tokenContract string, amount *big.Int, ethereumSender string, cosmosReceiver string, orchestrator string, data string, txHash string) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.DepositClaim(&_HyperionPrecompile.TransactOpts, hyperionId, eventNonce, blockHeight, tokenContract, amount, ethereumSender, cosmosReceiver, orchestrator, data, txHash)
}

// SetOrchestratorAddresses is a paid mutator transaction binding the contract method 0x5c1ea757.
//
// Solidity: function setOrchestratorAddresses(address orchestratorAddress, uint64 hyperionId) returns(bool success)
func (_HyperionPrecompile *HyperionPrecompileTransactor) SetOrchestratorAddresses(opts *bind.TransactOpts, orchestratorAddress common.Address, hyperionId uint64) (*types.Transaction, error) {
	return _HyperionPrecompile.contract.Transact(opts, "setOrchestratorAddresses", orchestratorAddress, hyperionId)
}

// SetOrchestratorAddresses is a paid mutator transaction binding the contract method 0x5c1ea757.
//
// Solidity: function setOrchestratorAddresses(address orchestratorAddress, uint64 hyperionId) returns(bool success)
func (_HyperionPrecompile *HyperionPrecompileSession) SetOrchestratorAddresses(orchestratorAddress common.Address, hyperionId uint64) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.SetOrchestratorAddresses(&_HyperionPrecompile.TransactOpts, orchestratorAddress, hyperionId)
}

// SetOrchestratorAddresses is a paid mutator transaction binding the contract method 0x5c1ea757.
//
// Solidity: function setOrchestratorAddresses(address orchestratorAddress, uint64 hyperionId) returns(bool success)
func (_HyperionPrecompile *HyperionPrecompileTransactorSession) SetOrchestratorAddresses(orchestratorAddress common.Address, hyperionId uint64) (*types.Transaction, error) {
	return _HyperionPrecompile.Contract.SetOrchestratorAddresses(&_HyperionPrecompile.TransactOpts, orchestratorAddress, hyperionId)
}
