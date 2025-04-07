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

// ValsetArgs is an auto generated low-level Go binding around an user-defined struct.
type ValsetArgs struct {
	Validators   []common.Address
	Powers       []*big.Int
	ValsetNonce  *big.Int
	RewardAmount *big.Int
	RewardToken  common.Address
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea26469706673582212207a0ae3ab2cc6eb7725a4e95c11639c7d83ea3d2d23b312edf53696ad7dfd29b464736f6c63430008190033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// CosmosERC20MetaData contains all meta data concerning the CosmosERC20 contract.
var CosmosERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523461032e576110558038038061001981610332565b928339810160608282031261032e5781516001600160401b039081811161032e5782610046918501610357565b906020928385015182811161032e57604091610063918701610357565b9401519360ff8516850361032e578251828111610245576003918254916001958684811c94168015610324575b88851014610310578190601f948581116102c2575b508890858311600114610264575f92610259575b50505f1982861b1c191690861b1783555b80519384116102455760049586548681811c9116801561023b575b82821014610228578381116101e5575b508092851160011461018057509383949184925f95610175575b50501b925f19911b1c19161790555b600580546001600160a01b0319163390811790915560405191905f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a3608052610cac90816103a98239608051816107090152f35b015193505f8061010f565b92919084601f198116885f52855f20955f905b898383106101cb57505050106101b2575b50505050811b01905561011e565b01519060f8845f19921b161c191690555f8080806101a4565b858701518955909701969485019488935090810190610193565b875f52815f208480880160051c82019284891061021f575b0160051c019087905b8281106102145750506100f5565b5f8155018790610206565b925081926101fd565b602288634e487b7160e01b5f525260245ffd5b90607f16906100e5565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100b9565b90889350601f19831691875f528a5f20925f5b8c8282106102ac5750508411610295575b505050811b0183556100ca565b01515f1983881b60f8161c191690555f8080610288565b8385015186558c97909501949384019301610277565b909150855f52885f208580850160051c8201928b8610610307575b918a91869594930160051c01915b8281106102f95750506100a5565b5f81558594508a91016102eb565b925081926102dd565b634e487b7160e01b5f52602260045260245ffd5b93607f1693610090565b5f80fd5b6040519190601f01601f191682016001600160401b0381118382101761024557604052565b81601f8201121561032e578051906001600160401b03821161024557610386601f8301601f1916602001610332565b928284526020838301011161032e57815f9260208093018386015e830101529056fe608060409080825260049081361015610016575f80fd5b5f3560e01c90816306fdde031461083857508063095ea7b31461080f57806318160ddd146107f157806323b872dd1461072d578063313ce567146106f057806339509351146106ab57806340c10f19146105d657806370a08231146105a0578063715018a6146105445780638da5cb5b1461051c57806395d89b41146103fc5780639dc29fac146102bf578063a457c2d714610213578063a9059cbb146101e3578063dd62ed3e1461019a5763f2fde38b146100d0575f80fd5b34610196576020366003190112610196576100e9610958565b600554916001600160a01b03808416926101043385146109b2565b1693841561014457505082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a36001600160a01b03191617600555005b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b5f80fd5b82346101965780600319360112610196576020906101b6610958565b6101be61096e565b9060018060a01b038091165f5260018452825f2091165f528252805f20549051908152f35b823461019657806003193601126101965760209061020c610202610958565b6024359033610afb565b5160018152f35b503461019657816003193601126101965761022c610958565b9060243590335f526001602052835f2060018060a01b0384165f52602052835f20549082821061026e5760208561020c866102678787610984565b90336109fd565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b50346101965781600319360112610196576102d8610958565b600554602435916001600160a01b03916102f590831633146109b2565b169182156103af57825f525f602052835f2054908282106103615750815f946103417fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef94602094610984565b8587528684528187205561035782600254610984565b60025551908152a3005b608490602086519162461bcd60e51b8352820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152fd5b608490602085519162461bcd60e51b8352820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152fd5b509034610196575f366003190112610196578051905f835460018160011c9060018316928315610512575b60209384841081146104ff578388529081156104e3575060011461048f575b505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b0390f35b604190634e487b7160e01b5f525260245ffd5b5f878152929350837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8385106104cf57505050508301015f8080610446565b8054888601830152930192849082016104b9565b60ff1916878501525050151560051b84010190505f8080610446565b602289634e487b7160e01b5f525260245ffd5b91607f1691610427565b8234610196575f3660031901126101965760055490516001600160a01b039091168152602090f35b34610196575f366003190112610196576005545f6001600160a01b03821661056d3382146109b2565b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a36001600160a01b031916600555005b8234610196576020366003190112610196576020906001600160a01b036105c5610958565b165f525f8252805f20549051908152f35b5090346101965780600319360112610196576105f0610958565b60055460243592916001600160a01b039161060e90831633146109b2565b1692831561066957506020827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926106495f956002546109a5565b6002558585528483528085206106608382546109a5565b905551908152a3005b6020606492519162461bcd60e51b8352820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b823461019657806003193601126101965760209061020c6106ca610958565b335f5260018452825f2060018060a01b0382165f528452610267602435845f20546109a5565b8234610196575f366003190112610196576020905160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461019657606036600319011261019657610747610958565b9061075061096e565b9061075f604435809385610afb565b60018060a01b0383165f526001602052835f20335f52602052835f20549082821061079d5760208561020c866107958787610984565b9033906109fd565b608490602086519162461bcd60e51b8352820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152fd5b8234610196575f366003190112610196576020906002549051908152f35b823461019657806003193601126101965760209061020c61082e610958565b60243590336109fd565b90508234610196575f366003190112610196575f60035460018160011c9060018316928315610924575b60209384841081146104ff5783885290811561090857506001146108b257505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b60035f908152929350837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8385106108f45750505050830101848080610446565b8054888601830152930192849082016108de565b60ff1916878501525050151560051b8401019050848080610446565b91607f1691610862565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361019657565b602435906001600160a01b038216820361019657565b9190820391821161099157565b634e487b7160e01b5f52601160045260245ffd5b9190820180921161099157565b156109b957565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6001600160a01b03908116918215610aaa5716918215610a5a5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b6001600160a01b03908116918215610c235716918215610bd257815f525f60205260405f2054818110610b7e5781610b567fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef93602093610984565b845f525f835260405f2055845f5260405f20610b738282546109a5565b9055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fdfea2646970667358221220418daa12254165c4c9282976d969382b8b711ad2a05322ad1712abe67c4e3ebe64736f6c63430008190033",
}

// CosmosERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use CosmosERC20MetaData.ABI instead.
var CosmosERC20ABI = CosmosERC20MetaData.ABI

// CosmosERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CosmosERC20MetaData.Bin instead.
var CosmosERC20Bin = CosmosERC20MetaData.Bin

// DeployCosmosERC20 deploys a new Ethereum contract, binding an instance of CosmosERC20 to it.
func DeployCosmosERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, decimals_ uint8) (common.Address, *types.Transaction, *CosmosERC20, error) {
	parsed, err := CosmosERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CosmosERC20Bin), backend, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CosmosERC20{CosmosERC20Caller: CosmosERC20Caller{contract: contract}, CosmosERC20Transactor: CosmosERC20Transactor{contract: contract}, CosmosERC20Filterer: CosmosERC20Filterer{contract: contract}}, nil
}

// CosmosERC20 is an auto generated Go binding around an Ethereum contract.
type CosmosERC20 struct {
	CosmosERC20Caller     // Read-only binding to the contract
	CosmosERC20Transactor // Write-only binding to the contract
	CosmosERC20Filterer   // Log filterer for contract events
}

// CosmosERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CosmosERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CosmosERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CosmosERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CosmosERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CosmosERC20Session struct {
	Contract     *CosmosERC20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CosmosERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CosmosERC20CallerSession struct {
	Contract *CosmosERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CosmosERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CosmosERC20TransactorSession struct {
	Contract     *CosmosERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CosmosERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CosmosERC20Raw struct {
	Contract *CosmosERC20 // Generic contract binding to access the raw methods on
}

// CosmosERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CosmosERC20CallerRaw struct {
	Contract *CosmosERC20Caller // Generic read-only contract binding to access the raw methods on
}

// CosmosERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CosmosERC20TransactorRaw struct {
	Contract *CosmosERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCosmosERC20 creates a new instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20(address common.Address, backend bind.ContractBackend) (*CosmosERC20, error) {
	contract, err := bindCosmosERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20{CosmosERC20Caller: CosmosERC20Caller{contract: contract}, CosmosERC20Transactor: CosmosERC20Transactor{contract: contract}, CosmosERC20Filterer: CosmosERC20Filterer{contract: contract}}, nil
}

// NewCosmosERC20Caller creates a new read-only instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20Caller(address common.Address, caller bind.ContractCaller) (*CosmosERC20Caller, error) {
	contract, err := bindCosmosERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20Caller{contract: contract}, nil
}

// NewCosmosERC20Transactor creates a new write-only instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*CosmosERC20Transactor, error) {
	contract, err := bindCosmosERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20Transactor{contract: contract}, nil
}

// NewCosmosERC20Filterer creates a new log filterer instance of CosmosERC20, bound to a specific deployed contract.
func NewCosmosERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*CosmosERC20Filterer, error) {
	contract, err := bindCosmosERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20Filterer{contract: contract}, nil
}

// bindCosmosERC20 binds a generic wrapper to an already deployed contract.
func bindCosmosERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CosmosERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosERC20 *CosmosERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosERC20.Contract.CosmosERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosERC20 *CosmosERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosERC20.Contract.CosmosERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosERC20 *CosmosERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosERC20.Contract.CosmosERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CosmosERC20 *CosmosERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CosmosERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CosmosERC20 *CosmosERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CosmosERC20 *CosmosERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CosmosERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.Allowance(&_CosmosERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CosmosERC20 *CosmosERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.Allowance(&_CosmosERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosmosERC20 *CosmosERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.BalanceOf(&_CosmosERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_CosmosERC20 *CosmosERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _CosmosERC20.Contract.BalanceOf(&_CosmosERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosmosERC20 *CosmosERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosmosERC20 *CosmosERC20Session) Decimals() (uint8, error) {
	return _CosmosERC20.Contract.Decimals(&_CosmosERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CosmosERC20 *CosmosERC20CallerSession) Decimals() (uint8, error) {
	return _CosmosERC20.Contract.Decimals(&_CosmosERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmosERC20 *CosmosERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmosERC20 *CosmosERC20Session) Name() (string, error) {
	return _CosmosERC20.Contract.Name(&_CosmosERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CosmosERC20 *CosmosERC20CallerSession) Name() (string, error) {
	return _CosmosERC20.Contract.Name(&_CosmosERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmosERC20 *CosmosERC20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmosERC20 *CosmosERC20Session) Owner() (common.Address, error) {
	return _CosmosERC20.Contract.Owner(&_CosmosERC20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CosmosERC20 *CosmosERC20CallerSession) Owner() (common.Address, error) {
	return _CosmosERC20.Contract.Owner(&_CosmosERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmosERC20 *CosmosERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmosERC20 *CosmosERC20Session) Symbol() (string, error) {
	return _CosmosERC20.Contract.Symbol(&_CosmosERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CosmosERC20 *CosmosERC20CallerSession) Symbol() (string, error) {
	return _CosmosERC20.Contract.Symbol(&_CosmosERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmosERC20 *CosmosERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CosmosERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmosERC20 *CosmosERC20Session) TotalSupply() (*big.Int, error) {
	return _CosmosERC20.Contract.TotalSupply(&_CosmosERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CosmosERC20 *CosmosERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _CosmosERC20.Contract.TotalSupply(&_CosmosERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Approve(&_CosmosERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Approve(&_CosmosERC20.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_CosmosERC20 *CosmosERC20Transactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_CosmosERC20 *CosmosERC20Session) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Burn(&_CosmosERC20.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_CosmosERC20 *CosmosERC20TransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Burn(&_CosmosERC20.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.DecreaseAllowance(&_CosmosERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.DecreaseAllowance(&_CosmosERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.IncreaseAllowance(&_CosmosERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.IncreaseAllowance(&_CosmosERC20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_CosmosERC20 *CosmosERC20Transactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_CosmosERC20 *CosmosERC20Session) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Mint(&_CosmosERC20.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_CosmosERC20 *CosmosERC20TransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Mint(&_CosmosERC20.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmosERC20 *CosmosERC20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmosERC20 *CosmosERC20Session) RenounceOwnership() (*types.Transaction, error) {
	return _CosmosERC20.Contract.RenounceOwnership(&_CosmosERC20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CosmosERC20 *CosmosERC20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CosmosERC20.Contract.RenounceOwnership(&_CosmosERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Transfer(&_CosmosERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.Transfer(&_CosmosERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.TransferFrom(&_CosmosERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_CosmosERC20 *CosmosERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CosmosERC20.Contract.TransferFrom(&_CosmosERC20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmosERC20 *CosmosERC20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CosmosERC20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmosERC20 *CosmosERC20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmosERC20.Contract.TransferOwnership(&_CosmosERC20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CosmosERC20 *CosmosERC20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CosmosERC20.Contract.TransferOwnership(&_CosmosERC20.TransactOpts, newOwner)
}

// CosmosERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CosmosERC20 contract.
type CosmosERC20ApprovalIterator struct {
	Event *CosmosERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosmosERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosmosERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosmosERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosERC20Approval represents a Approval event raised by the CosmosERC20 contract.
type CosmosERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CosmosERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CosmosERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20ApprovalIterator{contract: _CosmosERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CosmosERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CosmosERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosERC20Approval)
				if err := _CosmosERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) ParseApproval(log types.Log) (*CosmosERC20Approval, error) {
	event := new(CosmosERC20Approval)
	if err := _CosmosERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmosERC20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CosmosERC20 contract.
type CosmosERC20OwnershipTransferredIterator struct {
	Event *CosmosERC20OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosmosERC20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosERC20OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosmosERC20OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosmosERC20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosERC20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosERC20OwnershipTransferred represents a OwnershipTransferred event raised by the CosmosERC20 contract.
type CosmosERC20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmosERC20 *CosmosERC20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CosmosERC20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmosERC20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20OwnershipTransferredIterator{contract: _CosmosERC20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmosERC20 *CosmosERC20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CosmosERC20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CosmosERC20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosERC20OwnershipTransferred)
				if err := _CosmosERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CosmosERC20 *CosmosERC20Filterer) ParseOwnershipTransferred(log types.Log) (*CosmosERC20OwnershipTransferred, error) {
	event := new(CosmosERC20OwnershipTransferred)
	if err := _CosmosERC20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CosmosERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CosmosERC20 contract.
type CosmosERC20TransferIterator struct {
	Event *CosmosERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CosmosERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CosmosERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CosmosERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CosmosERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CosmosERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CosmosERC20Transfer represents a Transfer event raised by the CosmosERC20 contract.
type CosmosERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CosmosERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CosmosERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CosmosERC20TransferIterator{contract: _CosmosERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CosmosERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CosmosERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CosmosERC20Transfer)
				if err := _CosmosERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CosmosERC20 *CosmosERC20Filterer) ParseTransfer(log types.Log) (*CosmosERC20Transfer, error) {
	event := new(CosmosERC20Transfer)
	if err := _CosmosERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052346102d557610bde80380380610019816102d9565b9283398101906040818303126102d55780516001600160401b03908181116102d557836100479184016102fe565b91602093848201518381116102d55761006092016102fe565b82518281116101ec576003918254916001958684811c941680156102cb575b888510146102b7578190601f94858111610269575b50889085831160011461020b575f92610200575b50505f1982861b1c191690861b1783555b80519384116101ec5760049586548681811c911680156101e2575b828210146101cf5783811161018c575b508092851160011461012757509383949184925f9561011c575b50501b925f19911b1c19161790555b60405161088e90816103508239f35b015193505f806100fe565b92919084601f198116885f52855f20955f905b898383106101725750505010610159575b50505050811b01905561010d565b01519060f8845f19921b161c191690555f80808061014b565b85870151895590970196948501948893509081019061013a565b875f52815f208480880160051c8201928489106101c6575b0160051c019087905b8281106101bb5750506100e4565b5f81550187906101ad565b925081926101a4565b602288634e487b7160e01b5f525260245ffd5b90607f16906100d4565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100a8565b90889350601f19831691875f528a5f20925f5b8c828210610253575050841161023c575b505050811b0183556100b9565b01515f1983881b60f8161c191690555f808061022f565b8385015186558c9790950194938401930161021e565b909150855f52885f208580850160051c8201928b86106102ae575b918a91869594930160051c01915b8281106102a0575050610094565b5f81558594508a9101610292565b92508192610284565b634e487b7160e01b5f52602260045260245ffd5b93607f169361007f565b5f80fd5b6040519190601f01601f191682016001600160401b038111838210176101ec57604052565b81601f820112156102d5578051906001600160401b0382116101ec5761032d601f8301601f19166020016102d9565b92828452602083830101116102d557815f9260208093018386015e830101529056fe6080604090808252600480361015610015575f80fd5b5f3560e01c91826306fdde031461048357508163095ea7b31461045a57816318160ddd1461043c57816323b872dd14610377578163313ce5671461035c578163395093511461031757816370a08231146102e157816395d89b41146101c2578163a457c2d71461011557508063a9059cbb146100e55763dd62ed3e14610099575f80fd5b346100e157806003193601126100e1576020906100b4610585565b6100bc61059b565b9060018060a01b038091165f5260018452825f2091165f528252805f20549051908152f35b5f80fd5b50346100e157806003193601126100e15760209061010e610104610585565b60243590336106dd565b5160018152f35b9050346100e157816003193601126100e15761012f610585565b9060243590335f526001602052835f2060018060a01b0384165f52602052835f2054908282106101715760208561010e8661016a87876105b1565b90336105df565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b82346100e1575f3660031901126100e1578051905f835460018160011c90600183169283156102d7575b60209384841081146102c4578388529081156102a85750600114610254575b505050829003601f01601f191682019267ffffffffffffffff841183851017610241575082918261023d92528261055b565b0390f35b604190634e487b7160e01b5f525260245ffd5b5f878152929350837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b838510610294575050505083010184808061020b565b80548886018301529301928490820161027e565b60ff1916878501525050151560051b840101905084808061020b565b602289634e487b7160e01b5f525260245ffd5b91607f16916101ec565b82346100e15760203660031901126100e1576020906001600160a01b03610306610585565b165f525f8252805f20549051908152f35b82346100e157806003193601126100e15760209061010e610336610585565b335f5260018452825f2060018060a01b0382165f52845261016a602435845f20546105d2565b82346100e1575f3660031901126100e1576020905160128152f35b9050346100e15760603660031901126100e157610392610585565b9061039b61059b565b906103aa6044358093856106dd565b60018060a01b0383165f526001602052835f20335f52602052835f2054908282106103e85760208561010e866103e087876105b1565b9033906105df565b608490602086519162461bcd60e51b8352820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152fd5b82346100e1575f3660031901126100e1576020906002549051908152f35b82346100e157806003193601126100e15760209061010e610479610585565b60243590336105df565b83346100e1575f3660031901126100e1575f60035460018160011c9060018316928315610551575b60209384841081146102c4578388529081156102a857506001146104fb57505050829003601f01601f191682019267ffffffffffffffff841183851017610241575082918261023d92528261055b565b60035f908152929350837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b83851061053d575050505083010184808061020b565b805488860183015293019284908201610527565b91607f16916104ab565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b03821682036100e157565b602435906001600160a01b03821682036100e157565b919082039182116105be57565b634e487b7160e01b5f52601160045260245ffd5b919082018092116105be57565b6001600160a01b0390811691821561068c571691821561063c5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b6001600160a01b0390811691821561080557169182156107b457815f525f60205260405f205481811061076057816107387fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef936020936105b1565b845f525f835260405f2055845f5260405f206107558282546105d2565b9055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fdfea264697066735822122053efe65efc45df18b9134db6c2628174f86bc585399e8070664b218cb274aba264736f6c63430008190033",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// ERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MetaData.Bin instead.
var ERC20Bin = ERC20MetaData.Bin

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionMetaData contains all meta data concerning the Hyperion contract.
var HyperionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"ERC20DeployedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"SendToHeliosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchExecutedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"ValsetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_methodName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_args\",\"type\":\"bytes\"}],\"name\":\"callData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"deployERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"deployERC20WithSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnershipExpiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHeliosNativeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwnershipExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"lastBatchNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnershipAfterExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"sendToHelios\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_hyperionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"state_invalidationMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"state_lastBatchNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_newValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080806040523460345760ff196066541660665560016067555f606b555f606c555f606d555f606e55613b4690816100398239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c8063011b2174146101565780631ee7a108146101bf5780632b7553c7146101ba578063308ff208146101b55780634a4e3bd5146101b057806351858e27146101ab5780635afe97bb146101a65780635c975abb146101a1578063715018a61461019c57806373b20547146101975780637dfb6f8614610192578063817474181461018d5780638c64865f146101885780638da5cb5b14610183578063962ab5b11461017e578063a4b52ca214610179578063a5352f5b14610174578063a6c42b021461016f578063b56561fe1461016a578063c2d0732e14610165578063c359a21214610160578063d03c19b01461015b578063df97174b14610156578063e5a2b5d214610151578063f2b533071461014c578063f2fde38b146101475763f795563714610142575f80fd5b61121d565b61113f565b611122565b611105565b6101e9565b6110e8565b611039565b610f44565b610f27565b610ee7565b610e48565b610ded565b610dd0565b610da8565b610d3c565b610c60565b610a43565b610a26565b6109f9565b6109d7565b6109b3565b610948565b6108a8565b61043f565b61037d565b61022e565b6001600160a01b038116036101d557565b5f80fd5b61010435906101e7826101c4565b565b346101d55760203660031901126101d557600435610206816101c4565b60018060a01b03165f526069602052602060405f2054604051908152f35b5f9103126101d557565b346101d5575f3660031901126101d5576020610248611390565b604051908152f35b634e487b7160e01b5f52604160045260245ffd5b6001600160401b03811161027757604052565b610250565b60a081019081106001600160401b0382111761027757604052565b604081019081106001600160401b0382111761027757604052565b90601f801991011681019081106001600160401b0382111761027757604052565b604051906101e78261027c565b6001600160401b03811161027757601f01601f191660200190565b929192610307826102e0565b9161031560405193846102b2565b8294818452818301116101d5578281602093845f960137010152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b909161036c61037a93604084526040840190610331565b916020818403910152610331565b90565b346101d55760603660031901126101d55760043561039a816101c4565b6001600160401b03906024358281116101d557366023820112156101d5576103cc9036906024816004013591016102fb565b6044359283116101d557366023840112156101d5576103f86103fe9336906024816004013591016102fb565b91611430565b9061040e60405192839283610355565b0390f35b9181601f840112156101d5578235916001600160401b0383116101d557602083818601950101116101d557565b346101d55760803660031901126101d557600480359061045e826101c4565b60243591604435916064356001600160401b0381116101d5576104849036908301610412565b9261049460ff60665416156114da565b6104a360026067541415611519565b60026067556001600160a01b0381165f9081526071602052604090206104cb905b5460ff1690565b15610599576001600160a01b031692833b156101d55760408051632770a7eb60e21b81523394810194855260208501879052935f918591829101038183885af1928315610594577f272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b479361057b575b5061054d610548606d5461134f565b606d55565b61055643606e55565b61056b606d549160405193849333988561165d565b0390a45b6105796001606755565b005b8061058861058e92610264565b80610224565b5f610539565b611574565b604080516370a0823160e01b8082523086830190815296976001600160a01b0394909416969293602093919291849083908190830103818b5afa918215610594575f92610885575b506105ee9030338a612402565b83519182523086830190815283908390819060200103818a5afa80156105945761061f925f91610858575b5061157f565b9361062e610548606d5461134f565b61063743606e55565b825163313ce56760e01b815293828583818a5afa948515610594575f95610829575b50606d54978451956395d89b4160e01b87525f8785818c5afa968715610594575f9761080d575b505f865180956306fdde0360e01b8252818c5afa938415610594575f946107e2575b5060ff6106af91166124a6565b9181156107d2576106c19136916102fb565b915b8451607b60f81b9481019485529586946001016c226d65746164617461223a207b60981b8152600d016a1139bcb6b137b6111d101160a91b8152600b01610709916113a4565b61088b60f21b815260020168113730b6b2911d101160b91b815260090161072f916113a4565b61088b60f21b81526002016b0113232b1b4b6b0b639911d160a51b8152600c01610758916113a4565b611f4b60f21b8152600201670113230ba30911d160c51b815260080161077d916113a4565b607d60f81b815260010103601f198101835261079990836102b2565b5191829133956107a99284611621565b037f272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b4791a461056f565b50506107dc611603565b916106c3565b6106af91945061080560ff913d805f833e6107fd81836102b2565b8101906115a1565b9491506106a2565b6108229197503d805f833e6107fd81836102b2565b955f610680565b61084a919550833d8511610851575b61084281836102b2565b81019061158c565b935f610659565b503d610838565b6108789150843d861161087e575b61087081836102b2565b810190611565565b5f610619565b503d610866565b6105ee9192506108a190853d871161087e5761087081836102b2565b91906105e1565b346101d5575f3660031901126101d5576108cd60018060a01b03603354163314611679565b60665460ff81161561090c5760ff19166066557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b60405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606490fd5b346101d5575f3660031901126101d55761096d60018060a01b03603354163314611679565b600160665461097f60ff8216156114da565b60ff1916176066557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346101d5575f3660031901126101d55760206109cd611390565b4211604051908152f35b346101d5575f3660031901126101d557602060ff606654166040519015158152f35b346101d5575f3660031901126101d557610a1e60018060a01b03603354163314611679565b610579612546565b346101d5575f3660031901126101d5576020606d54604051908152f35b346101d55760203660031901126101d5576004355f52606a602052602060405f2054604051908152f35b6001600160401b0381116102775760051b60200190565b9291610a8f82610a6d565b91610a9d60405193846102b2565b829481845260208094019160051b81019283116101d557905b828210610ac35750505050565b8380918335610ad1816101c4565b815201910190610ab6565b9080601f830112156101d55781602061037a93359101610a84565b9291610b0282610a6d565b91610b1060405193846102b2565b829481845260208094019160051b81019283116101d557905b828210610b365750505050565b81358152908301908301610b29565b9080601f830112156101d55781602061037a93359101610af7565b91909160a0818403126101d55760405190610b7a8261027c565b81938135916001600160401b03928381116101d55782610b9b918301610adc565b845260208101359283116101d557610bb96080939284938301610b45565b60208501526040810135604085015260608101356060850152013591610bde836101c4565b0152565b60ff8116036101d557565b9291610bf882610a6d565b91610c0660405193846102b2565b829481845260208094019160051b81019283116101d557905b828210610c2c5750505050565b8380918335610c3a81610be2565b815201910190610c1f565b9080601f830112156101d55781602061037a93359101610bed565b346101d5576101403660031901126101d55760046001600160401b0381358181116101d557610c929036908401610b60565b906024358181116101d557610caa9036908501610c45565b916044358281116101d557610cc29036908601610b45565b916064358181116101d557610cda9036908701610b45565b906084358181116101d557610cf29036908801610b45565b60a4358281116101d557610d099036908901610adc565b9160c4359081116101d55761057997610d2491369101610b45565b92610d2d6101d9565b95610124359760e435966116c4565b346101d5575f3660031901126101d557610d54611390565b421115610d6357610579612546565b60405162461bcd60e51b815260206004820152601960248201527f4f776e657273686970206e6f74207965742065787069726564000000000000006044820152606490fd5b346101d5575f3660031901126101d5576033546040516001600160a01b039091168152602090f35b346101d5575f3660031901126101d5576020606e54604051908152f35b346101d5575f3660031901126101d5576020606f54604051908152f35b908160a09103126101d55790565b9181601f840112156101d5578235916001600160401b0383116101d5576020808501948460051b0101116101d557565b346101d55760a03660031901126101d55760046001600160401b0381358181116101d557610e799036908401610e0a565b6024358281116101d557610e909036908501610e0a565b916044358181116101d557610ea89036908601610e18565b6064929192358281116101d557610ec29036908801610e18565b9390926084359081116101d55761057997610edf91369101610e18565b969095611d4c565b346101d55760203660031901126101d557600435610f04816101c4565b60018060a01b03165f526071602052602060ff60405f2054166040519015158152f35b346101d5575f3660031901126101d5576020606b54604051908152f35b346101d55760a03660031901126101d5576001600160401b036004358181116101d557610f75903690600401610412565b50506024358181116101d557610f8f903690600401610412565b91906044358281116101d557610fa9903690600401610412565b60643591610fb683610be2565b60405195611055948588019688881090881117610277578796610fdd96612abc893961219d565b03905ff08015610594576001600160a01b0316803b156101d5576040516340c10f1960e01b81523360048201526084356024820152905f908290604490829084905af180156105945761102c57005b8061058861057992610264565b346101d55760803660031901126101d5576001600160401b036044358181116101d55761106a903690600401610e18565b6064929192359182116101d5576110886110b8923690600401610e18565b915f549460ff8660081c169586158097816110dd575b6110a7906121d2565b6110cc575b5060243560043561228b565b6110be57005b61057961ff00195f54165f55565b61ffff1916610101175f555f6110ac565b5060ff82161561109e565b346101d5575f3660031901126101d5576020606c54604051908152f35b346101d5575f3660031901126101d5576020607054604051908152f35b346101d5575f3660031901126101d5576020606854604051908152f35b346101d55760203660031901126101d55760043561115c816101c4565b6033546001600160a01b039081169190611177338414611679565b811680156111c957610579927f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a360018060a01b03166bffffffffffffffffffffffff60a01b6033541617603355565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346101d55760803660031901126101d5576001600160401b036004358181116101d55761124e903690600401610412565b6024929192358281116101d557611269903690600401610412565b93906044358481116101d557611283903690600401610412565b9190946064359061129382610be2565b6040519061105590818301908111838210176102775783868a8c886112be958897612abc893961219d565b03905ff0928315610594577f82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7966113369460018060a01b031698895f52607160205260405f20600160ff1982541617905561131d610548606d5461134f565b61132643606e55565b606d5494604051988998896123b5565b0390a2005b634e487b7160e01b5f52601160045260245ffd5b906001820180921161135d57565b61133b565b906509184e72a000820180921161135d57565b603001908160301161135d57565b9190820180921161135d57565b6034546302f4bd00810180911161135d5790565b805191908290602001825e015f815290565b60405190602082018281106001600160401b03821117610277576040525f8252565b3d15611402573d906113e9826102e0565b916113f760405193846102b2565b82523d5f602084013e565b606090565b6040519061141482610297565b600d82526c2ab735b737bbb71032b93937b960991b6020830152565b5f929161148e61149c859461145e611458604051956020815191012063ffffffff60e01b1690565b94610264565b6040519360208501526004845261147484610297565b6040519283916114886020840180976113a4565b906113a4565b03601f1981018352826102b2565b51915afa6114a86113d8565b90156114b7579061037a6113b6565b80516114d457506114c6611407565b905b6114d06113b6565b9190565b906114c8565b156114e157565b60405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606490fd5b1561152057565b60405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606490fd5b908160209103126101d5575190565b6040513d5f823e3d90fd5b9190820391821161135d57565b908160209103126101d5575161037a81610be2565b6020818303126101d5578051906001600160401b0382116101d5570181601f820112156101d5578051906115d4826102e0565b926115e260405194856102b2565b828452602083830101116101d557815f9260208093018386015e8301015290565b6040519061161082610297565b6002825261111160f11b6020830152565b61037a9392606092825260208201528160408201520190610331565b908060209392818452848401375f828201840152601f01601f1916010190565b61037a949260609282526020820152816040820152019161163d565b1561168057565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b928799929a979698949a6116dd60026067541415611519565b60026067556116f160ff60665416156114da565b6001600160a01b0386165f90815260696020526040902061171590899054106119c4565b6001600160a01b0386165f9081526069602052604090206117419061173a9054611362565b8910611a2f565b61174c844310611abd565b8b8a888751519761177960208201998a5151811490816119b9575b816119ae575b816119a3575b50611b2f565b606f549061179361178a8383612587565b60685414611b7b565b6117a98551845181149081611998575b50611bed565b519851978a60409d60405195869560208701986117c6968a611ca8565b03601f19810182526117d890826102b2565b51902092607054946117e9966126e3565b6001600160a01b0381165f90815260696020526040812087905593845b8851861015611911576001600160a01b0383165f90815260716020526040902061182f906104c4565b156118dc576001600160a01b0383169061185961184c8887611d33565b516001600160a01b031690565b611863888c611d33565b51833b156101d55787516340c10f1960e01b81526001600160a01b039290921660048301526024820152915f908390604490829084905af1908115610594576001926118c1926118c9575b505b6118ba888a611d33565b5190611383565b950194611806565b806105886118d692610264565b5f6118ae565b6118c160019161190c8b6118fd8a6118f761184c828c611d33565b92611d33565b5190858060a01b038816612812565b6118b0565b9450965094935050508061197e575b5061192f610548606d5461134f565b61193843606e55565b606d546040519081526001600160a01b0392909216917f02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab70890602090a36101e76001606755565b61199290336001600160a01b038516612812565b5f611920565b90508551145f6117a3565b90508751145f611773565b87518114915061176d565b895181149150611767565b156119cb57565b60405162461bcd60e51b815260206004820152603660248201527f4e6577206261746368206e6f6e6365206d7573742062652067726561746572206044820152757468616e207468652063757272656e74206e6f6e636560501b6064820152608490fd5b15611a3657565b60405162461bcd60e51b815260206004820152605360248201527f4e6577206261746368206e6f6e6365206d757374206265206c6573732074686160448201527f6e2031305f3030305f3030305f3030305f3030302067726561746572207468616064820152726e207468652063757272656e74206e6f6e636560681b608482015260a490fd5b15611ac457565b60405162461bcd60e51b815260206004820152603b60248201527f42617463682074696d656f7574206d757374206265206772656174657220746860448201527f616e207468652063757272656e7420626c6f636b2068656967687400000000006064820152608490fd5b15611b3657565b60405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d65642063757272656e742076616c696461746f7220736574006044820152606490fd5b15611b8257565b60405162461bcd60e51b815260206004820152603f60248201527f537570706c6965642063757272656e742076616c696461746f727320616e642060448201527f706f7765727320646f206e6f74206d6174636820636865636b706f696e742e006064820152608490fd5b15611bf457565b60405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d6564206261746368206f66207472616e73616374696f6e73006044820152606490fd5b9081518082526020808093019301915f5b828110611c58575050505090565b835185529381019392810192600101611c4a565b9081518082526020808093019301915f5b828110611c8b575050505090565b83516001600160a01b031685529381019392810192600101611c7d565b939060e09598979693611ce9611d0594611cf7936101009089526f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b60208a01528060408a0152880190611c39565b908682036060880152611c6c565b908482036080860152611c39565b60a08301969096526001600160a01b031660c08201520152565b634e487b7160e01b5f52603260045260245ffd5b8051821015611d475760209160051b010190565b611d1f565b94919295939095611d6260ff60665416156114da565b604086013597604088013593611d79858b11611fb1565b611d838980612023565b60208b01969150611d94878c612023565b9190501480611f9d575b80611f89575b80611f70575b611db390611b2f565b611dbc90611362565b8a10611dc790612058565b606f5480611dd5368c610b60565b90611ddf91612587565b60685414611dec90611b7b565b611df6368a610b60565b90611e0091612587565b968794611e0d8b80612023565b97611e18919c612023565b98906070549c8d993690611e2b92610a84565b993690611e3792610af7565b953690611e4392610bed565b923690611e4f92610af7565b923690611e5b92610af7565b92611e65966126e3565b611e6f8280612023565b90506020830193611e808585612023565b90611e8a9361284e565b606855611e9683606b55565b43606c5560808101906001600160a01b03611eb0836120e7565b1615157f76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a9360609382611f3393611f64575b611f38575b611ef5610548606d5461134f565b611efe43606e55565b611f21611f0d606d54926120e7565b92611f188680612023565b92909187612023565b9490936040519889980135908861215c565b0390a2565b611f5f611f53611f47836120e7565b6001600160a01b031690565b86860135903390612812565b611ee7565b50848401351515611ee2565b50611db383611f7f8c80612023565b9050149050611daa565b5081611f958b80612023565b905014611da4565b5083611fa98b80612023565b905014611d9e565b15611fb857565b60405162461bcd60e51b815260206004820152603760248201527f4e65772076616c736574206e6f6e6365206d757374206265206772656174657260448201527f207468616e207468652063757272656e74206e6f6e63650000000000000000006064820152608490fd5b903590601e19813603018212156101d557018035906001600160401b0382116101d557602001918160051b360383136101d557565b1561205f57565b60405162461bcd60e51b815260206004820152605460248201527f4e65772076616c736574206e6f6e6365206d757374206265206c65737320746860448201527f616e2031305f3030305f3030305f3030305f3030302067726561746572207468606482015273616e207468652063757272656e74206e6f6e636560601b608482015260a490fd5b3561037a816101c4565b9190808252602080920192915f5b82811061210d575050505090565b9091929382806001928735612121816101c4565b848060a01b031681520195019101929190926120ff565b81835290916001600160fb1b0383116101d55760209260051b809284830137010190565b9593919261218f9361037a9896928852602088015260018060a01b0316604087015260a0606087015260a08601916120f1565b926080818503910152612138565b93926040936121bd6121cb9360ff9599989960608952606089019161163d565b91868303602088015261163d565b9416910152565b156121d957565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b604051906122428261027c565b5f60808360608152606060208201528260408201528260608201520152565b939161037a959361218f9286525f60208701525f604087015260a0606087015260a08601916120f1565b94612356866123518461234c6123468a9b8a61232a7f76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a9d6123208d9e9d612307611f339e8e5f5460ff8160081c169081158092816123aa575b6122ed906121d2565b612398575b5061238a575b612300612899565b858861284e565b61230f612235565b506123186102d3565b943691610a84565b83528a3691610af7565b60208201525f60408201525f60608201525f6080820152612587565b93606f55565b607055565b606855565b612364610548606d5461134f565b61236d43606c55565b61237643606e55565b606b5495606d549360405195869586612261565b61ff00195f54165f556122f8565b61ffff1916610101175f9081556122f2565b5060ff8216156122e4565b94916123f79360ff956123db608099946123e9949d9c9b9d60a08b5260a08b019161163d565b9188830360208a015261163d565b91858303604087015261163d565b951660608201520152565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526101e7916124448261027c565b6129a3565b5f19811461135d5760010190565b90612461826102e0565b61246e60405191826102b2565b828152809261247f601f19916102e0565b0190602036910137565b801561135d575f190190565b908151811015611d47570160200190565b80156125285780815f925b6125145750806124c083612457565b92915b6124cc57505090565b600a906124ff6124f96124e96124e3858506611375565b60ff1690565b60f81b6001600160f81b03191690565b93612489565b925f1a61250c8486612495565b5304806124c3565b91612520600a91612449565b9204806124b1565b5060405161253581610297565b60018152600360fc1b602082015290565b6033545f6001600160a01b0382167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a36001600160a01b031916603355565b612613604082015191805192602082015191606081015190608060018060a01b03910151166125fc6125e9604051978895602087019a8b526918da1958dadc1bda5b9d60b21b6040880152606087015260e06080870152610100860190611c6c565b601f1995868683030160a0870152611c39565b9160c084015260e0830152039081018352826102b2565b51902090565b1561262057565b60405162461bcd60e51b815260206004820152602360248201527f56616c696461746f72207369676e617475726520646f6573206e6f74206d617460448201526231b41760e91b6064820152608490fd5b1561267857565b60405162461bcd60e51b815260206004820152603c60248201527f5375626d69747465642076616c696461746f7220736574207369676e6174757260448201527f657320646f206e6f74206861766520656e6f75676820706f7765722e000000006064820152608490fd5b95949390935f935f5b88518110156128015760ff806127028388611d33565b5116612712575b506001016126ec565b90956001600160a01b0380612727898d611d33565b5116926127348989611d33565b51166127408986611d33565b519061274c8a89611d33565b5190604092835192602094858501947f19457468657265756d205369676e6564204d6573736167653a0a3332000000008652603c8c81830152815260608101938185106001600160401b03861117610277575f968560c094528251902085526080958683015260a0820152015282805260015afa15610594576127e0926127d6915f511614612619565b6118ba8789611d33565b948786116127ee575f612709565b505050505090506101e792505b11612671565b505050505090506101e792506127fb565b60405163a9059cbb60e01b60208201526001600160a01b039290921660248301526044808301939093529181526101e7916124446064836102b2565b8261285c9194939414611b2f565b5f905f5b848110612875575b50506101e7925011612671565b8060051b820135830180931161135d5783831161289457600101612860565b612868565b5f5460ff8160081c16908115809281612921575b6128b6906121d2565b612910575b50336bffffffffffffffffffffffff60a01b603354161760335542603455335f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a361290557565b61ff00195f54165f55565b61ffff1916610101175f555f6128bb565b5060ff8216156128ad565b908160209103126101d5575180151581036101d55790565b1561294b57565b60405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b6040516001600160a01b0391909116916129bc82610297565b6020928383527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656484840152803b15612a32575f8281928287612a0d9796519301915af1612a076113d8565b90612a77565b80519081612a1a57505050565b826101e793612a2d93830101910161292c565b612944565b60405162461bcd60e51b815260048101859052601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b90919015612a83575090565b815115612a935750805190602001fd5b60405162461bcd60e51b815260206004820152908190612ab7906024830190610331565b0390fdfe60a06040523461032e576110558038038061001981610332565b928339810160608282031261032e5781516001600160401b039081811161032e5782610046918501610357565b906020928385015182811161032e57604091610063918701610357565b9401519360ff8516850361032e578251828111610245576003918254916001958684811c94168015610324575b88851014610310578190601f948581116102c2575b508890858311600114610264575f92610259575b50505f1982861b1c191690861b1783555b80519384116102455760049586548681811c9116801561023b575b82821014610228578381116101e5575b508092851160011461018057509383949184925f95610175575b50501b925f19911b1c19161790555b600580546001600160a01b0319163390811790915560405191905f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a3608052610cac90816103a98239608051816107090152f35b015193505f8061010f565b92919084601f198116885f52855f20955f905b898383106101cb57505050106101b2575b50505050811b01905561011e565b01519060f8845f19921b161c191690555f8080806101a4565b858701518955909701969485019488935090810190610193565b875f52815f208480880160051c82019284891061021f575b0160051c019087905b8281106102145750506100f5565b5f8155018790610206565b925081926101fd565b602288634e487b7160e01b5f525260245ffd5b90607f16906100e5565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100b9565b90889350601f19831691875f528a5f20925f5b8c8282106102ac5750508411610295575b505050811b0183556100ca565b01515f1983881b60f8161c191690555f8080610288565b8385015186558c97909501949384019301610277565b909150855f52885f208580850160051c8201928b8610610307575b918a91869594930160051c01915b8281106102f95750506100a5565b5f81558594508a91016102eb565b925081926102dd565b634e487b7160e01b5f52602260045260245ffd5b93607f1693610090565b5f80fd5b6040519190601f01601f191682016001600160401b0381118382101761024557604052565b81601f8201121561032e578051906001600160401b03821161024557610386601f8301601f1916602001610332565b928284526020838301011161032e57815f9260208093018386015e830101529056fe608060409080825260049081361015610016575f80fd5b5f3560e01c90816306fdde031461083857508063095ea7b31461080f57806318160ddd146107f157806323b872dd1461072d578063313ce567146106f057806339509351146106ab57806340c10f19146105d657806370a08231146105a0578063715018a6146105445780638da5cb5b1461051c57806395d89b41146103fc5780639dc29fac146102bf578063a457c2d714610213578063a9059cbb146101e3578063dd62ed3e1461019a5763f2fde38b146100d0575f80fd5b34610196576020366003190112610196576100e9610958565b600554916001600160a01b03808416926101043385146109b2565b1693841561014457505082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a36001600160a01b03191617600555005b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b5f80fd5b82346101965780600319360112610196576020906101b6610958565b6101be61096e565b9060018060a01b038091165f5260018452825f2091165f528252805f20549051908152f35b823461019657806003193601126101965760209061020c610202610958565b6024359033610afb565b5160018152f35b503461019657816003193601126101965761022c610958565b9060243590335f526001602052835f2060018060a01b0384165f52602052835f20549082821061026e5760208561020c866102678787610984565b90336109fd565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b50346101965781600319360112610196576102d8610958565b600554602435916001600160a01b03916102f590831633146109b2565b169182156103af57825f525f602052835f2054908282106103615750815f946103417fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef94602094610984565b8587528684528187205561035782600254610984565b60025551908152a3005b608490602086519162461bcd60e51b8352820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152fd5b608490602085519162461bcd60e51b8352820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152fd5b509034610196575f366003190112610196578051905f835460018160011c9060018316928315610512575b60209384841081146104ff578388529081156104e3575060011461048f575b505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b0390f35b604190634e487b7160e01b5f525260245ffd5b5f878152929350837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8385106104cf57505050508301015f8080610446565b8054888601830152930192849082016104b9565b60ff1916878501525050151560051b84010190505f8080610446565b602289634e487b7160e01b5f525260245ffd5b91607f1691610427565b8234610196575f3660031901126101965760055490516001600160a01b039091168152602090f35b34610196575f366003190112610196576005545f6001600160a01b03821661056d3382146109b2565b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a36001600160a01b031916600555005b8234610196576020366003190112610196576020906001600160a01b036105c5610958565b165f525f8252805f20549051908152f35b5090346101965780600319360112610196576105f0610958565b60055460243592916001600160a01b039161060e90831633146109b2565b1692831561066957506020827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926106495f956002546109a5565b6002558585528483528085206106608382546109a5565b905551908152a3005b6020606492519162461bcd60e51b8352820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b823461019657806003193601126101965760209061020c6106ca610958565b335f5260018452825f2060018060a01b0382165f528452610267602435845f20546109a5565b8234610196575f366003190112610196576020905160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461019657606036600319011261019657610747610958565b9061075061096e565b9061075f604435809385610afb565b60018060a01b0383165f526001602052835f20335f52602052835f20549082821061079d5760208561020c866107958787610984565b9033906109fd565b608490602086519162461bcd60e51b8352820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152fd5b8234610196575f366003190112610196576020906002549051908152f35b823461019657806003193601126101965760209061020c61082e610958565b60243590336109fd565b90508234610196575f366003190112610196575f60035460018160011c9060018316928315610924575b60209384841081146104ff5783885290811561090857506001146108b257505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b60035f908152929350837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8385106108f45750505050830101848080610446565b8054888601830152930192849082016108de565b60ff1916878501525050151560051b8401019050848080610446565b91607f1691610862565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361019657565b602435906001600160a01b038216820361019657565b9190820391821161099157565b634e487b7160e01b5f52601160045260245ffd5b9190820180921161099157565b156109b957565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6001600160a01b03908116918215610aaa5716918215610a5a5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b6001600160a01b03908116918215610c235716918215610bd257815f525f60205260405f2054818110610b7e5781610b567fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef93602093610984565b845f525f835260405f2055845f5260405f20610b738282546109a5565b9055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fdfea2646970667358221220418daa12254165c4c9282976d969382b8b711ad2a05322ad1712abe67c4e3ebe64736f6c63430008190033a26469706673582212208dc85ee4d64605793932d76602049330e5b8ec75e791007b65ca3722d0efad8b64736f6c63430008190033",
}

// HyperionABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperionMetaData.ABI instead.
var HyperionABI = HyperionMetaData.ABI

// HyperionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HyperionMetaData.Bin instead.
var HyperionBin = HyperionMetaData.Bin

// DeployHyperion deploys a new Ethereum contract, binding an instance of Hyperion to it.
func DeployHyperion(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Hyperion, error) {
	parsed, err := HyperionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HyperionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hyperion{HyperionCaller: HyperionCaller{contract: contract}, HyperionTransactor: HyperionTransactor{contract: contract}, HyperionFilterer: HyperionFilterer{contract: contract}}, nil
}

// Hyperion is an auto generated Go binding around an Ethereum contract.
type Hyperion struct {
	HyperionCaller     // Read-only binding to the contract
	HyperionTransactor // Write-only binding to the contract
	HyperionFilterer   // Log filterer for contract events
}

// HyperionCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperionSession struct {
	Contract     *Hyperion         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HyperionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperionCallerSession struct {
	Contract *HyperionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// HyperionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperionTransactorSession struct {
	Contract     *HyperionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// HyperionRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperionRaw struct {
	Contract *Hyperion // Generic contract binding to access the raw methods on
}

// HyperionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperionCallerRaw struct {
	Contract *HyperionCaller // Generic read-only contract binding to access the raw methods on
}

// HyperionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperionTransactorRaw struct {
	Contract *HyperionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperion creates a new instance of Hyperion, bound to a specific deployed contract.
func NewHyperion(address common.Address, backend bind.ContractBackend) (*Hyperion, error) {
	contract, err := bindHyperion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hyperion{HyperionCaller: HyperionCaller{contract: contract}, HyperionTransactor: HyperionTransactor{contract: contract}, HyperionFilterer: HyperionFilterer{contract: contract}}, nil
}

// NewHyperionCaller creates a new read-only instance of Hyperion, bound to a specific deployed contract.
func NewHyperionCaller(address common.Address, caller bind.ContractCaller) (*HyperionCaller, error) {
	contract, err := bindHyperion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionCaller{contract: contract}, nil
}

// NewHyperionTransactor creates a new write-only instance of Hyperion, bound to a specific deployed contract.
func NewHyperionTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperionTransactor, error) {
	contract, err := bindHyperion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionTransactor{contract: contract}, nil
}

// NewHyperionFilterer creates a new log filterer instance of Hyperion, bound to a specific deployed contract.
func NewHyperionFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperionFilterer, error) {
	contract, err := bindHyperion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperionFilterer{contract: contract}, nil
}

// bindHyperion binds a generic wrapper to an already deployed contract.
func bindHyperion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HyperionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyperion *HyperionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyperion.Contract.HyperionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyperion *HyperionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.Contract.HyperionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyperion *HyperionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyperion.Contract.HyperionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hyperion *HyperionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Hyperion.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hyperion *HyperionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hyperion *HyperionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hyperion.Contract.contract.Transact(opts, method, params...)
}

// CallData is a free data retrieval call binding the contract method 0x2b7553c7.
//
// Solidity: function callData(address _contractAddress, string _methodName, bytes _args) view returns(bytes data, bytes err)
func (_Hyperion *HyperionCaller) CallData(opts *bind.CallOpts, _contractAddress common.Address, _methodName string, _args []byte) (struct {
	Data []byte
	Err  []byte
}, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "callData", _contractAddress, _methodName, _args)

	outstruct := new(struct {
		Data []byte
		Err  []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Err = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// CallData is a free data retrieval call binding the contract method 0x2b7553c7.
//
// Solidity: function callData(address _contractAddress, string _methodName, bytes _args) view returns(bytes data, bytes err)
func (_Hyperion *HyperionSession) CallData(_contractAddress common.Address, _methodName string, _args []byte) (struct {
	Data []byte
	Err  []byte
}, error) {
	return _Hyperion.Contract.CallData(&_Hyperion.CallOpts, _contractAddress, _methodName, _args)
}

// CallData is a free data retrieval call binding the contract method 0x2b7553c7.
//
// Solidity: function callData(address _contractAddress, string _methodName, bytes _args) view returns(bytes data, bytes err)
func (_Hyperion *HyperionCallerSession) CallData(_contractAddress common.Address, _methodName string, _args []byte) (struct {
	Data []byte
	Err  []byte
}, error) {
	return _Hyperion.Contract.CallData(&_Hyperion.CallOpts, _contractAddress, _methodName, _args)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_Hyperion *HyperionCaller) GetOwnershipExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "getOwnershipExpiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_Hyperion *HyperionSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _Hyperion.Contract.GetOwnershipExpiryTimestamp(&_Hyperion.CallOpts)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_Hyperion *HyperionCallerSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _Hyperion.Contract.GetOwnershipExpiryTimestamp(&_Hyperion.CallOpts)
}

// IsHeliosNativeToken is a free data retrieval call binding the contract method 0xa6c42b02.
//
// Solidity: function isHeliosNativeToken(address ) view returns(bool)
func (_Hyperion *HyperionCaller) IsHeliosNativeToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "isHeliosNativeToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeliosNativeToken is a free data retrieval call binding the contract method 0xa6c42b02.
//
// Solidity: function isHeliosNativeToken(address ) view returns(bool)
func (_Hyperion *HyperionSession) IsHeliosNativeToken(arg0 common.Address) (bool, error) {
	return _Hyperion.Contract.IsHeliosNativeToken(&_Hyperion.CallOpts, arg0)
}

// IsHeliosNativeToken is a free data retrieval call binding the contract method 0xa6c42b02.
//
// Solidity: function isHeliosNativeToken(address ) view returns(bool)
func (_Hyperion *HyperionCallerSession) IsHeliosNativeToken(arg0 common.Address) (bool, error) {
	return _Hyperion.Contract.IsHeliosNativeToken(&_Hyperion.CallOpts, arg0)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_Hyperion *HyperionCaller) IsOwnershipExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "isOwnershipExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_Hyperion *HyperionSession) IsOwnershipExpired() (bool, error) {
	return _Hyperion.Contract.IsOwnershipExpired(&_Hyperion.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_Hyperion *HyperionCallerSession) IsOwnershipExpired() (bool, error) {
	return _Hyperion.Contract.IsOwnershipExpired(&_Hyperion.CallOpts)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Hyperion *HyperionCaller) LastBatchNonce(opts *bind.CallOpts, _erc20Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "lastBatchNonce", _erc20Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Hyperion *HyperionSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Hyperion.Contract.LastBatchNonce(&_Hyperion.CallOpts, _erc20Address)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_Hyperion *HyperionCallerSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _Hyperion.Contract.LastBatchNonce(&_Hyperion.CallOpts, _erc20Address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hyperion *HyperionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hyperion *HyperionSession) Owner() (common.Address, error) {
	return _Hyperion.Contract.Owner(&_Hyperion.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Hyperion *HyperionCallerSession) Owner() (common.Address, error) {
	return _Hyperion.Contract.Owner(&_Hyperion.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hyperion *HyperionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hyperion *HyperionSession) Paused() (bool, error) {
	return _Hyperion.Contract.Paused(&_Hyperion.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Hyperion *HyperionCallerSession) Paused() (bool, error) {
	return _Hyperion.Contract.Paused(&_Hyperion.CallOpts)
}

// StateHyperionId is a free data retrieval call binding the contract method 0xa4b52ca2.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_Hyperion *HyperionCaller) StateHyperionId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_hyperionId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateHyperionId is a free data retrieval call binding the contract method 0xa4b52ca2.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_Hyperion *HyperionSession) StateHyperionId() ([32]byte, error) {
	return _Hyperion.Contract.StateHyperionId(&_Hyperion.CallOpts)
}

// StateHyperionId is a free data retrieval call binding the contract method 0xa4b52ca2.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_Hyperion *HyperionCallerSession) StateHyperionId() ([32]byte, error) {
	return _Hyperion.Contract.StateHyperionId(&_Hyperion.CallOpts)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Hyperion *HyperionCaller) StateInvalidationMapping(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_invalidationMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Hyperion *HyperionSession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _Hyperion.Contract.StateInvalidationMapping(&_Hyperion.CallOpts, arg0)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _Hyperion.Contract.StateInvalidationMapping(&_Hyperion.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastBatchNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastBatchNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Hyperion *HyperionSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Hyperion.Contract.StateLastBatchNonces(&_Hyperion.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _Hyperion.Contract.StateLastBatchNonces(&_Hyperion.CallOpts, arg0)
}

// StateLastEventHeight is a free data retrieval call binding the contract method 0x962ab5b1.
//
// Solidity: function state_lastEventHeight() view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastEventHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastEventHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastEventHeight is a free data retrieval call binding the contract method 0x962ab5b1.
//
// Solidity: function state_lastEventHeight() view returns(uint256)
func (_Hyperion *HyperionSession) StateLastEventHeight() (*big.Int, error) {
	return _Hyperion.Contract.StateLastEventHeight(&_Hyperion.CallOpts)
}

// StateLastEventHeight is a free data retrieval call binding the contract method 0x962ab5b1.
//
// Solidity: function state_lastEventHeight() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastEventHeight() (*big.Int, error) {
	return _Hyperion.Contract.StateLastEventHeight(&_Hyperion.CallOpts)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastEventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Hyperion *HyperionSession) StateLastEventNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastEventNonce(&_Hyperion.CallOpts)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastEventNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastEventNonce(&_Hyperion.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Hyperion *HyperionCaller) StateLastValsetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastValsetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Hyperion *HyperionSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Hyperion.Contract.StateLastValsetCheckpoint(&_Hyperion.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_Hyperion *HyperionCallerSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _Hyperion.Contract.StateLastValsetCheckpoint(&_Hyperion.CallOpts)
}

// StateLastValsetHeight is a free data retrieval call binding the contract method 0xd03c19b0.
//
// Solidity: function state_lastValsetHeight() view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastValsetHeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastValsetHeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastValsetHeight is a free data retrieval call binding the contract method 0xd03c19b0.
//
// Solidity: function state_lastValsetHeight() view returns(uint256)
func (_Hyperion *HyperionSession) StateLastValsetHeight() (*big.Int, error) {
	return _Hyperion.Contract.StateLastValsetHeight(&_Hyperion.CallOpts)
}

// StateLastValsetHeight is a free data retrieval call binding the contract method 0xd03c19b0.
//
// Solidity: function state_lastValsetHeight() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastValsetHeight() (*big.Int, error) {
	return _Hyperion.Contract.StateLastValsetHeight(&_Hyperion.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Hyperion *HyperionCaller) StateLastValsetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_lastValsetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Hyperion *HyperionSession) StateLastValsetNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastValsetNonce(&_Hyperion.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StateLastValsetNonce() (*big.Int, error) {
	return _Hyperion.Contract.StateLastValsetNonce(&_Hyperion.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Hyperion *HyperionCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Hyperion.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Hyperion *HyperionSession) StatePowerThreshold() (*big.Int, error) {
	return _Hyperion.Contract.StatePowerThreshold(&_Hyperion.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_Hyperion *HyperionCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _Hyperion.Contract.StatePowerThreshold(&_Hyperion.CallOpts)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Hyperion *HyperionTransactor) DeployERC20(opts *bind.TransactOpts, _cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "deployERC20", _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Hyperion *HyperionSession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Hyperion.Contract.DeployERC20(&_Hyperion.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_Hyperion *HyperionTransactorSession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _Hyperion.Contract.DeployERC20(&_Hyperion.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20WithSupply is a paid mutator transaction binding the contract method 0xc2d0732e.
//
// Solidity: function deployERC20WithSupply(string , string _name, string _symbol, uint8 _decimals, uint256 supply) returns()
func (_Hyperion *HyperionTransactor) DeployERC20WithSupply(opts *bind.TransactOpts, arg0 string, _name string, _symbol string, _decimals uint8, supply *big.Int) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "deployERC20WithSupply", arg0, _name, _symbol, _decimals, supply)
}

// DeployERC20WithSupply is a paid mutator transaction binding the contract method 0xc2d0732e.
//
// Solidity: function deployERC20WithSupply(string , string _name, string _symbol, uint8 _decimals, uint256 supply) returns()
func (_Hyperion *HyperionSession) DeployERC20WithSupply(arg0 string, _name string, _symbol string, _decimals uint8, supply *big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.DeployERC20WithSupply(&_Hyperion.TransactOpts, arg0, _name, _symbol, _decimals, supply)
}

// DeployERC20WithSupply is a paid mutator transaction binding the contract method 0xc2d0732e.
//
// Solidity: function deployERC20WithSupply(string , string _name, string _symbol, uint8 _decimals, uint256 supply) returns()
func (_Hyperion *HyperionTransactorSession) DeployERC20WithSupply(arg0 string, _name string, _symbol string, _decimals uint8, supply *big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.DeployERC20WithSupply(&_Hyperion.TransactOpts, arg0, _name, _symbol, _decimals, supply)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_Hyperion *HyperionTransactor) EmergencyPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "emergencyPause")
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_Hyperion *HyperionSession) EmergencyPause() (*types.Transaction, error) {
	return _Hyperion.Contract.EmergencyPause(&_Hyperion.TransactOpts)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_Hyperion *HyperionTransactorSession) EmergencyPause() (*types.Transaction, error) {
	return _Hyperion.Contract.EmergencyPause(&_Hyperion.TransactOpts)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_Hyperion *HyperionTransactor) EmergencyUnpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "emergencyUnpause")
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_Hyperion *HyperionSession) EmergencyUnpause() (*types.Transaction, error) {
	return _Hyperion.Contract.EmergencyUnpause(&_Hyperion.TransactOpts)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_Hyperion *HyperionTransactorSession) EmergencyUnpause() (*types.Transaction, error) {
	return _Hyperion.Contract.EmergencyUnpause(&_Hyperion.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc359a212.
//
// Solidity: function initialize(bytes32 _hyperionId, uint256 _powerThreshold, address[] _validators, uint256[] _powers) returns()
func (_Hyperion *HyperionTransactor) Initialize(opts *bind.TransactOpts, _hyperionId [32]byte, _powerThreshold *big.Int, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "initialize", _hyperionId, _powerThreshold, _validators, _powers)
}

// Initialize is a paid mutator transaction binding the contract method 0xc359a212.
//
// Solidity: function initialize(bytes32 _hyperionId, uint256 _powerThreshold, address[] _validators, uint256[] _powers) returns()
func (_Hyperion *HyperionSession) Initialize(_hyperionId [32]byte, _powerThreshold *big.Int, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.Initialize(&_Hyperion.TransactOpts, _hyperionId, _powerThreshold, _validators, _powers)
}

// Initialize is a paid mutator transaction binding the contract method 0xc359a212.
//
// Solidity: function initialize(bytes32 _hyperionId, uint256 _powerThreshold, address[] _validators, uint256[] _powers) returns()
func (_Hyperion *HyperionTransactorSession) Initialize(_hyperionId [32]byte, _powerThreshold *big.Int, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.Initialize(&_Hyperion.TransactOpts, _hyperionId, _powerThreshold, _validators, _powers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hyperion *HyperionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hyperion *HyperionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hyperion.Contract.RenounceOwnership(&_Hyperion.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Hyperion *HyperionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Hyperion.Contract.RenounceOwnership(&_Hyperion.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_Hyperion *HyperionTransactor) RenounceOwnershipAfterExpiry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "renounceOwnershipAfterExpiry")
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_Hyperion *HyperionSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _Hyperion.Contract.RenounceOwnershipAfterExpiry(&_Hyperion.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_Hyperion *HyperionTransactorSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _Hyperion.Contract.RenounceOwnershipAfterExpiry(&_Hyperion.TransactOpts)
}

// SendToHelios is a paid mutator transaction binding the contract method 0x308ff208.
//
// Solidity: function sendToHelios(address _tokenContract, bytes32 _destination, uint256 _amount, string _data) returns()
func (_Hyperion *HyperionTransactor) SendToHelios(opts *bind.TransactOpts, _tokenContract common.Address, _destination [32]byte, _amount *big.Int, _data string) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "sendToHelios", _tokenContract, _destination, _amount, _data)
}

// SendToHelios is a paid mutator transaction binding the contract method 0x308ff208.
//
// Solidity: function sendToHelios(address _tokenContract, bytes32 _destination, uint256 _amount, string _data) returns()
func (_Hyperion *HyperionSession) SendToHelios(_tokenContract common.Address, _destination [32]byte, _amount *big.Int, _data string) (*types.Transaction, error) {
	return _Hyperion.Contract.SendToHelios(&_Hyperion.TransactOpts, _tokenContract, _destination, _amount, _data)
}

// SendToHelios is a paid mutator transaction binding the contract method 0x308ff208.
//
// Solidity: function sendToHelios(address _tokenContract, bytes32 _destination, uint256 _amount, string _data) returns()
func (_Hyperion *HyperionTransactorSession) SendToHelios(_tokenContract common.Address, _destination [32]byte, _amount *big.Int, _data string) (*types.Transaction, error) {
	return _Hyperion.Contract.SendToHelios(&_Hyperion.TransactOpts, _tokenContract, _destination, _amount, _data)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x81747418.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Hyperion *HyperionTransactor) SubmitBatch(opts *bind.TransactOpts, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "submitBatch", _currentValset, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x81747418.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Hyperion *HyperionSession) SubmitBatch(_currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.SubmitBatch(&_Hyperion.TransactOpts, _currentValset, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x81747418.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_Hyperion *HyperionTransactorSession) SubmitBatch(_currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _Hyperion.Contract.SubmitBatch(&_Hyperion.TransactOpts, _currentValset, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hyperion *HyperionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hyperion *HyperionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hyperion.Contract.TransferOwnership(&_Hyperion.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Hyperion *HyperionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Hyperion.Contract.TransferOwnership(&_Hyperion.TransactOpts, newOwner)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xa5352f5b.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Hyperion *HyperionTransactor) UpdateValset(opts *bind.TransactOpts, _newValset ValsetArgs, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Hyperion.contract.Transact(opts, "updateValset", _newValset, _currentValset, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xa5352f5b.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Hyperion *HyperionSession) UpdateValset(_newValset ValsetArgs, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Hyperion.Contract.UpdateValset(&_Hyperion.TransactOpts, _newValset, _currentValset, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xa5352f5b.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_Hyperion *HyperionTransactorSession) UpdateValset(_newValset ValsetArgs, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _Hyperion.Contract.UpdateValset(&_Hyperion.TransactOpts, _newValset, _currentValset, _v, _r, _s)
}

// HyperionERC20DeployedEventIterator is returned from FilterERC20DeployedEvent and is used to iterate over the raw logs and unpacked data for ERC20DeployedEvent events raised by the Hyperion contract.
type HyperionERC20DeployedEventIterator struct {
	Event *HyperionERC20DeployedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HyperionERC20DeployedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionERC20DeployedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HyperionERC20DeployedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HyperionERC20DeployedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionERC20DeployedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionERC20DeployedEvent represents a ERC20DeployedEvent event raised by the Hyperion contract.
type HyperionERC20DeployedEvent struct {
	CosmosDenom   string
	TokenContract common.Address
	Name          string
	Symbol        string
	Decimals      uint8
	EventNonce    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterERC20DeployedEvent is a free log retrieval operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) FilterERC20DeployedEvent(opts *bind.FilterOpts, _tokenContract []common.Address) (*HyperionERC20DeployedEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &HyperionERC20DeployedEventIterator{contract: _Hyperion.contract, event: "ERC20DeployedEvent", logs: logs, sub: sub}, nil
}

// WatchERC20DeployedEvent is a free log subscription operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) WatchERC20DeployedEvent(opts *bind.WatchOpts, sink chan<- *HyperionERC20DeployedEvent, _tokenContract []common.Address) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionERC20DeployedEvent)
				if err := _Hyperion.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseERC20DeployedEvent is a log parse operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) ParseERC20DeployedEvent(log types.Log) (*HyperionERC20DeployedEvent, error) {
	event := new(HyperionERC20DeployedEvent)
	if err := _Hyperion.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Hyperion contract.
type HyperionOwnershipTransferredIterator struct {
	Event *HyperionOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HyperionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HyperionOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HyperionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionOwnershipTransferred represents a OwnershipTransferred event raised by the Hyperion contract.
type HyperionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hyperion *HyperionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HyperionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HyperionOwnershipTransferredIterator{contract: _Hyperion.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hyperion *HyperionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HyperionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionOwnershipTransferred)
				if err := _Hyperion.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Hyperion *HyperionFilterer) ParseOwnershipTransferred(log types.Log) (*HyperionOwnershipTransferred, error) {
	event := new(HyperionOwnershipTransferred)
	if err := _Hyperion.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Hyperion contract.
type HyperionPausedIterator struct {
	Event *HyperionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HyperionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HyperionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HyperionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionPaused represents a Paused event raised by the Hyperion contract.
type HyperionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hyperion *HyperionFilterer) FilterPaused(opts *bind.FilterOpts) (*HyperionPausedIterator, error) {

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HyperionPausedIterator{contract: _Hyperion.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hyperion *HyperionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HyperionPaused) (event.Subscription, error) {

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionPaused)
				if err := _Hyperion.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Hyperion *HyperionFilterer) ParsePaused(log types.Log) (*HyperionPaused, error) {
	event := new(HyperionPaused)
	if err := _Hyperion.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSendToHeliosEventIterator is returned from FilterSendToHeliosEvent and is used to iterate over the raw logs and unpacked data for SendToHeliosEvent events raised by the Hyperion contract.
type HyperionSendToHeliosEventIterator struct {
	Event *HyperionSendToHeliosEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HyperionSendToHeliosEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSendToHeliosEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HyperionSendToHeliosEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HyperionSendToHeliosEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSendToHeliosEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSendToHeliosEvent represents a SendToHeliosEvent event raised by the Hyperion contract.
type HyperionSendToHeliosEvent struct {
	TokenContract common.Address
	Sender        common.Address
	Destination   [32]byte
	Amount        *big.Int
	EventNonce    *big.Int
	Data          string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendToHeliosEvent is a free log retrieval operation binding the contract event 0x272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b47.
//
// Solidity: event SendToHeliosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce, string _data)
func (_Hyperion *HyperionFilterer) FilterSendToHeliosEvent(opts *bind.FilterOpts, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (*HyperionSendToHeliosEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "SendToHeliosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSendToHeliosEventIterator{contract: _Hyperion.contract, event: "SendToHeliosEvent", logs: logs, sub: sub}, nil
}

// WatchSendToHeliosEvent is a free log subscription operation binding the contract event 0x272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b47.
//
// Solidity: event SendToHeliosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce, string _data)
func (_Hyperion *HyperionFilterer) WatchSendToHeliosEvent(opts *bind.WatchOpts, sink chan<- *HyperionSendToHeliosEvent, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}
	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _destinationRule []interface{}
	for _, _destinationItem := range _destination {
		_destinationRule = append(_destinationRule, _destinationItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "SendToHeliosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSendToHeliosEvent)
				if err := _Hyperion.contract.UnpackLog(event, "SendToHeliosEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSendToHeliosEvent is a log parse operation binding the contract event 0x272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b47.
//
// Solidity: event SendToHeliosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce, string _data)
func (_Hyperion *HyperionFilterer) ParseSendToHeliosEvent(log types.Log) (*HyperionSendToHeliosEvent, error) {
	event := new(HyperionSendToHeliosEvent)
	if err := _Hyperion.contract.UnpackLog(event, "SendToHeliosEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionTransactionBatchExecutedEventIterator is returned from FilterTransactionBatchExecutedEvent and is used to iterate over the raw logs and unpacked data for TransactionBatchExecutedEvent events raised by the Hyperion contract.
type HyperionTransactionBatchExecutedEventIterator struct {
	Event *HyperionTransactionBatchExecutedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HyperionTransactionBatchExecutedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionTransactionBatchExecutedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HyperionTransactionBatchExecutedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HyperionTransactionBatchExecutedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionTransactionBatchExecutedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionTransactionBatchExecutedEvent represents a TransactionBatchExecutedEvent event raised by the Hyperion contract.
type HyperionTransactionBatchExecutedEvent struct {
	BatchNonce *big.Int
	Token      common.Address
	EventNonce *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchExecutedEvent is a free log retrieval operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) FilterTransactionBatchExecutedEvent(opts *bind.FilterOpts, _batchNonce []*big.Int, _token []common.Address) (*HyperionTransactionBatchExecutedEventIterator, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &HyperionTransactionBatchExecutedEventIterator{contract: _Hyperion.contract, event: "TransactionBatchExecutedEvent", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchExecutedEvent is a free log subscription operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) WatchTransactionBatchExecutedEvent(opts *bind.WatchOpts, sink chan<- *HyperionTransactionBatchExecutedEvent, _batchNonce []*big.Int, _token []common.Address) (event.Subscription, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionTransactionBatchExecutedEvent)
				if err := _Hyperion.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransactionBatchExecutedEvent is a log parse operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_Hyperion *HyperionFilterer) ParseTransactionBatchExecutedEvent(log types.Log) (*HyperionTransactionBatchExecutedEvent, error) {
	event := new(HyperionTransactionBatchExecutedEvent)
	if err := _Hyperion.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Hyperion contract.
type HyperionUnpausedIterator struct {
	Event *HyperionUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HyperionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HyperionUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HyperionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionUnpaused represents a Unpaused event raised by the Hyperion contract.
type HyperionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hyperion *HyperionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HyperionUnpausedIterator, error) {

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HyperionUnpausedIterator{contract: _Hyperion.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hyperion *HyperionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HyperionUnpaused) (event.Subscription, error) {

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionUnpaused)
				if err := _Hyperion.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Hyperion *HyperionFilterer) ParseUnpaused(log types.Log) (*HyperionUnpaused, error) {
	event := new(HyperionUnpaused)
	if err := _Hyperion.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionValsetUpdatedEventIterator is returned from FilterValsetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValsetUpdatedEvent events raised by the Hyperion contract.
type HyperionValsetUpdatedEventIterator struct {
	Event *HyperionValsetUpdatedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HyperionValsetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionValsetUpdatedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HyperionValsetUpdatedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HyperionValsetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionValsetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionValsetUpdatedEvent represents a ValsetUpdatedEvent event raised by the Hyperion contract.
type HyperionValsetUpdatedEvent struct {
	NewValsetNonce *big.Int
	EventNonce     *big.Int
	RewardAmount   *big.Int
	RewardToken    common.Address
	Validators     []common.Address
	Powers         []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValsetUpdatedEvent is a free log retrieval operation binding the contract event 0x76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, uint256 _eventNonce, uint256 _rewardAmount, address _rewardToken, address[] _validators, uint256[] _powers)
func (_Hyperion *HyperionFilterer) FilterValsetUpdatedEvent(opts *bind.FilterOpts, _newValsetNonce []*big.Int) (*HyperionValsetUpdatedEventIterator, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Hyperion.contract.FilterLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return &HyperionValsetUpdatedEventIterator{contract: _Hyperion.contract, event: "ValsetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValsetUpdatedEvent is a free log subscription operation binding the contract event 0x76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, uint256 _eventNonce, uint256 _rewardAmount, address _rewardToken, address[] _validators, uint256[] _powers)
func (_Hyperion *HyperionFilterer) WatchValsetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *HyperionValsetUpdatedEvent, _newValsetNonce []*big.Int) (event.Subscription, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _Hyperion.contract.WatchLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionValsetUpdatedEvent)
				if err := _Hyperion.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValsetUpdatedEvent is a log parse operation binding the contract event 0x76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, uint256 _eventNonce, uint256 _rewardAmount, address _rewardToken, address[] _validators, uint256[] _powers)
func (_Hyperion *HyperionFilterer) ParseValsetUpdatedEvent(log types.Log) (*HyperionValsetUpdatedEvent, error) {
	event := new(HyperionValsetUpdatedEvent)
	if err := _Hyperion.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataMetaData contains all meta data concerning the IERC20Metadata contract.
var IERC20MetadataMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20MetadataABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetadataMetaData.ABI instead.
var IERC20MetadataABI = IERC20MetadataMetaData.ABI

// IERC20Metadata is an auto generated Go binding around an Ethereum contract.
type IERC20Metadata struct {
	IERC20MetadataCaller     // Read-only binding to the contract
	IERC20MetadataTransactor // Write-only binding to the contract
	IERC20MetadataFilterer   // Log filterer for contract events
}

// IERC20MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataSession struct {
	Contract     *IERC20Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataCallerSession struct {
	Contract *IERC20MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataTransactorSession struct {
	Contract     *IERC20MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataRaw struct {
	Contract *IERC20Metadata // Generic contract binding to access the raw methods on
}

// IERC20MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataCallerRaw struct {
	Contract *IERC20MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataTransactorRaw struct {
	Contract *IERC20MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Metadata creates a new instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20Metadata(address common.Address, backend bind.ContractBackend) (*IERC20Metadata, error) {
	contract, err := bindIERC20Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Metadata{IERC20MetadataCaller: IERC20MetadataCaller{contract: contract}, IERC20MetadataTransactor: IERC20MetadataTransactor{contract: contract}, IERC20MetadataFilterer: IERC20MetadataFilterer{contract: contract}}, nil
}

// NewIERC20MetadataCaller creates a new read-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataCaller, error) {
	contract, err := bindIERC20Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataCaller{contract: contract}, nil
}

// NewIERC20MetadataTransactor creates a new write-only instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataTransactor, error) {
	contract, err := bindIERC20Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransactor{contract: contract}, nil
}

// NewIERC20MetadataFilterer creates a new log filterer instance of IERC20Metadata, bound to a specific deployed contract.
func NewIERC20MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataFilterer, error) {
	contract, err := bindIERC20Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataFilterer{contract: contract}, nil
}

// bindIERC20Metadata binds a generic wrapper to an already deployed contract.
func bindIERC20Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetadataMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.IERC20MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.IERC20MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Metadata *IERC20MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Metadata *IERC20MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.Allowance(&_IERC20Metadata.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Metadata.Contract.BalanceOf(&_IERC20Metadata.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Metadata *IERC20MetadataCallerSession) Decimals() (uint8, error) {
	return _IERC20Metadata.Contract.Decimals(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Name() (string, error) {
	return _IERC20Metadata.Contract.Name(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Metadata *IERC20MetadataCallerSession) Symbol() (string, error) {
	return _IERC20Metadata.Contract.Symbol(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Metadata.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Metadata *IERC20MetadataCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Metadata.Contract.TotalSupply(&_IERC20Metadata.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Approve(&_IERC20Metadata.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.Transfer(&_IERC20Metadata.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Metadata *IERC20MetadataTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Metadata.Contract.TransferFrom(&_IERC20Metadata.TransactOpts, sender, recipient, amount)
}

// IERC20MetadataApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Metadata contract.
type IERC20MetadataApprovalIterator struct {
	Event *IERC20MetadataApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20MetadataApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20MetadataApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20MetadataApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataApproval represents a Approval event raised by the IERC20Metadata contract.
type IERC20MetadataApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataApprovalIterator{contract: _IERC20Metadata.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataApproval)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) ParseApproval(log types.Log) (*IERC20MetadataApproval, error) {
	event := new(IERC20MetadataApproval)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Metadata contract.
type IERC20MetadataTransferIterator struct {
	Event *IERC20MetadataTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20MetadataTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20MetadataTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20MetadataTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataTransfer represents a Transfer event raised by the IERC20Metadata contract.
type IERC20MetadataTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataTransferIterator{contract: _IERC20Metadata.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Metadata.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataTransfer)
				if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Metadata *IERC20MetadataFilterer) ParseTransfer(log types.Log) (*IERC20MetadataTransfer, error) {
	event := new(IERC20MetadataTransfer)
	if err := _IERC20Metadata.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableUpgradeableWithExpiryMetaData contains all meta data concerning the OwnableUpgradeableWithExpiry contract.
var OwnableUpgradeableWithExpiryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getOwnershipExpiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwnershipExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnershipAfterExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableUpgradeableWithExpiryABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableUpgradeableWithExpiryMetaData.ABI instead.
var OwnableUpgradeableWithExpiryABI = OwnableUpgradeableWithExpiryMetaData.ABI

// OwnableUpgradeableWithExpiry is an auto generated Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiry struct {
	OwnableUpgradeableWithExpiryCaller     // Read-only binding to the contract
	OwnableUpgradeableWithExpiryTransactor // Write-only binding to the contract
	OwnableUpgradeableWithExpiryFilterer   // Log filterer for contract events
}

// OwnableUpgradeableWithExpiryCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableWithExpiryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableWithExpiryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableUpgradeableWithExpiryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableWithExpirySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableUpgradeableWithExpirySession struct {
	Contract     *OwnableUpgradeableWithExpiry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableWithExpiryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableUpgradeableWithExpiryCallerSession struct {
	Contract *OwnableUpgradeableWithExpiryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// OwnableUpgradeableWithExpiryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableUpgradeableWithExpiryTransactorSession struct {
	Contract     *OwnableUpgradeableWithExpiryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// OwnableUpgradeableWithExpiryRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryRaw struct {
	Contract *OwnableUpgradeableWithExpiry // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableWithExpiryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryCallerRaw struct {
	Contract *OwnableUpgradeableWithExpiryCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableWithExpiryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableUpgradeableWithExpiryTransactorRaw struct {
	Contract *OwnableUpgradeableWithExpiryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeableWithExpiry creates a new instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiry(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeableWithExpiry, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiry{OwnableUpgradeableWithExpiryCaller: OwnableUpgradeableWithExpiryCaller{contract: contract}, OwnableUpgradeableWithExpiryTransactor: OwnableUpgradeableWithExpiryTransactor{contract: contract}, OwnableUpgradeableWithExpiryFilterer: OwnableUpgradeableWithExpiryFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableWithExpiryCaller creates a new read-only instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiryCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableWithExpiryCaller, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryCaller{contract: contract}, nil
}

// NewOwnableUpgradeableWithExpiryTransactor creates a new write-only instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiryTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableWithExpiryTransactor, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableWithExpiryFilterer creates a new log filterer instance of OwnableUpgradeableWithExpiry, bound to a specific deployed contract.
func NewOwnableUpgradeableWithExpiryFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableWithExpiryFilterer, error) {
	contract, err := bindOwnableUpgradeableWithExpiry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryFilterer{contract: contract}, nil
}

// bindOwnableUpgradeableWithExpiry binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeableWithExpiry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableUpgradeableWithExpiryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeableWithExpiry.Contract.OwnableUpgradeableWithExpiryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.OwnableUpgradeableWithExpiryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.OwnableUpgradeableWithExpiryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeableWithExpiry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.contract.Transact(opts, method, params...)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCaller) GetOwnershipExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OwnableUpgradeableWithExpiry.contract.Call(opts, &out, "getOwnershipExpiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _OwnableUpgradeableWithExpiry.Contract.GetOwnershipExpiryTimestamp(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _OwnableUpgradeableWithExpiry.Contract.GetOwnershipExpiryTimestamp(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCaller) IsOwnershipExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OwnableUpgradeableWithExpiry.contract.Call(opts, &out, "isOwnershipExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) IsOwnershipExpired() (bool, error) {
	return _OwnableUpgradeableWithExpiry.Contract.IsOwnershipExpired(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerSession) IsOwnershipExpired() (bool, error) {
	return _OwnableUpgradeableWithExpiry.Contract.IsOwnershipExpired(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableUpgradeableWithExpiry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) Owner() (common.Address, error) {
	return _OwnableUpgradeableWithExpiry.Contract.Owner(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeableWithExpiry.Contract.Owner(&_OwnableUpgradeableWithExpiry.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactor) RenounceOwnershipAfterExpiry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.contract.Transact(opts, "renounceOwnershipAfterExpiry")
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnershipAfterExpiry(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.RenounceOwnershipAfterExpiry(&_OwnableUpgradeableWithExpiry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpirySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.TransferOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeableWithExpiry.Contract.TransferOwnership(&_OwnableUpgradeableWithExpiry.TransactOpts, newOwner)
}

// OwnableUpgradeableWithExpiryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeableWithExpiry contract.
type OwnableUpgradeableWithExpiryOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableWithExpiryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableUpgradeableWithExpiryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableWithExpiryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableUpgradeableWithExpiryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableUpgradeableWithExpiryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableWithExpiryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableWithExpiryOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeableWithExpiry contract.
type OwnableUpgradeableWithExpiryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableWithExpiryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeableWithExpiry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableWithExpiryOwnershipTransferredIterator{contract: _OwnableUpgradeableWithExpiry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableWithExpiryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeableWithExpiry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableWithExpiryOwnershipTransferred)
				if err := _OwnableUpgradeableWithExpiry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeableWithExpiry *OwnableUpgradeableWithExpiryFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableWithExpiryOwnershipTransferred, error) {
	event := new(OwnableUpgradeableWithExpiryOwnershipTransferred)
	if err := _OwnableUpgradeableWithExpiry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableMetaData contains all meta data concerning the Pausable contract.
var PausableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PausableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableMetaData.ABI instead.
var PausableABI = PausableMetaData.ABI

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PausableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// PausablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pausable contract.
type PausablePausedIterator struct {
	Event *PausablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePaused represents a Paused event raised by the Pausable contract.
type PausablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausablePausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausablePausedIterator{contract: _Pausable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausablePaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePaused)
				if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pausable *PausableFilterer) ParsePaused(log types.Log) (*PausablePaused, error) {
	event := new(PausablePaused)
	if err := _Pausable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pausable contract.
type PausableUnpausedIterator struct {
	Event *PausableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpaused represents a Unpaused event raised by the Pausable contract.
type PausableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUnpausedIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUnpausedIterator{contract: _Pausable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpaused)
				if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pausable *PausableFilterer) ParseUnpaused(log types.Log) (*PausableUnpaused, error) {
	event := new(PausableUnpaused)
	if err := _Pausable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardMetaData contains all meta data concerning the ReentrancyGuard contract.
var ReentrancyGuardMetaData = &bind.MetaData{
	ABI: "[]",
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardMetaData.ABI instead.
var ReentrancyGuardABI = ReentrancyGuardMetaData.ABI

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReentrancyGuardMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080806040523460175760399081601c823930815050f35b5f80fdfe5f80fdfea264697066735822122018ece9d28f133d3aaaf53fc3000f907c89bb87746580b95a205438f7187a2df664736f6c63430008190033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}
