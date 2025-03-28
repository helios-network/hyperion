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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"ERC20DeployedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"SendToHeliosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchExecutedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"ValsetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_methodName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_args\",\"type\":\"bytes\"}],\"name\":\"callData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"deployERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"deployERC20WithSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnershipExpiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHeliosNativeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwnershipExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"lastBatchNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnershipAfterExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"sendToHelios\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_hyperionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"state_invalidationMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"state_lastBatchNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_newValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234602c5760ff196066541660665560016067555f606b555f606c55613aae90816100318239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c8063011b2174146101405780631ee7a1081461019f5780632b7553c71461019a578063308ff208146101955780634a4e3bd51461019057806351858e271461018b5780635afe97bb146101865780635c975abb14610181578063715018a61461017c57806373b20547146101775780637dfb6f8614610172578063817474181461016d5780638c64865f146101685780638da5cb5b14610163578063a4b52ca21461015e578063a5352f5b14610159578063a6c42b0214610154578063b56561fe1461014f578063c2d0732e1461014a578063c359a21214610145578063df97174b14610140578063e5a2b5d21461013b578063f2b5330714610136578063f2fde38b146101315763f79556371461012c575f80fd5b6111b1565b6110d3565b6110b6565b611099565b6101c9565b610fea565b610ef5565b610ed8565b610e98565b610df9565b610d9e565b610d76565b610d0a565b610c2e565b610a11565b6109f4565b6109c7565b6109a5565b610981565b610916565b610876565b61041f565b61035d565b61020e565b6001600160a01b038116036101b557565b5f80fd5b61010435906101c7826101a4565b565b346101b55760203660031901126101b5576004356101e6816101a4565b60018060a01b03165f526069602052602060405f2054604051908152f35b5f9103126101b557565b346101b5575f3660031901126101b557602061022861131b565b604051908152f35b634e487b7160e01b5f52604160045260245ffd5b6001600160401b03811161025757604052565b610230565b60a081019081106001600160401b0382111761025757604052565b604081019081106001600160401b0382111761025757604052565b90601f801991011681019081106001600160401b0382111761025757604052565b604051906101c78261025c565b6001600160401b03811161025757601f01601f191660200190565b9291926102e7826102c0565b916102f56040519384610292565b8294818452818301116101b5578281602093845f960137010152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b909161034c61035a93604084526040840190610311565b916020818403910152610311565b90565b346101b55760603660031901126101b55760043561037a816101a4565b6001600160401b03906024358281116101b557366023820112156101b5576103ac9036906024816004013591016102db565b6044359283116101b557366023840112156101b5576103d86103de9336906024816004013591016102db565b916113bb565b906103ee60405192839283610335565b0390f35b9181601f840112156101b5578235916001600160401b0383116101b557602083818601950101116101b557565b346101b55760803660031901126101b557600480359061043e826101a4565b60243591604435916064356001600160401b0381116101b55761046490369083016103f2565b9261047460ff6066541615611465565b610483600260675414156114a4565b60026067556001600160a01b0381165f908152606f602052604090206104ab905b5460ff1690565b15610570576001600160a01b031692833b156101b55760408051632770a7eb60e21b81523394810194855260208501879052935f918591829101038183885af192831561056b577f272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b4793610552575b5061052d610528606c546112da565b606c55565b610542606c54916040519384933398856115e8565b0390a45b6105506001606755565b005b8061055f61056592610244565b80610204565b5f610519565b6114ff565b604080516370a0823160e01b8082523086830190815296976001600160a01b0394909416969293602093919291849083908190830103818b5afa91821561056b575f92610853575b506105c59030338a61236a565b83519182523086830190815283908390819060200103818a5afa801561056b576105f6925f91610826575b5061150a565b93610605610528606c546112da565b825163313ce56760e01b815293828583818a5afa94851561056b575f956107f7575b50606c54978451956395d89b4160e01b87525f8785818c5afa96871561056b575f976107db575b505f865180956306fdde0360e01b8252818c5afa93841561056b575f946107b0575b5060ff61067d911661240e565b9181156107a05761068f9136916102db565b915b8451607b60f81b9481019485529586946001016c226d65746164617461223a207b60981b8152600d016a1139bcb6b137b6111d101160a91b8152600b016106d79161132f565b61088b60f21b815260020168113730b6b2911d101160b91b81526009016106fd9161132f565b61088b60f21b81526002016b0113232b1b4b6b0b639911d160a51b8152600c016107269161132f565b611f4b60f21b8152600201670113230ba30911d160c51b815260080161074b9161132f565b607d60f81b815260010103601f19810183526107679083610292565b51918291339561077792846115ac565b037f272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b4791a4610546565b50506107aa61158e565b91610691565b61067d9194506107d360ff913d805f833e6107cb8183610292565b81019061152c565b949150610670565b6107f09197503d805f833e6107cb8183610292565b955f61064e565b610818919550833d851161081f575b6108108183610292565b810190611517565b935f610627565b503d610806565b6108469150843d861161084c575b61083e8183610292565b8101906114f0565b5f6105f0565b503d610834565b6105c591925061086f90853d871161084c5761083e8183610292565b91906105b8565b346101b5575f3660031901126101b55761089b60018060a01b03603354163314611604565b60665460ff8116156108da5760ff19166066557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b60405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606490fd5b346101b5575f3660031901126101b55761093b60018060a01b03603354163314611604565b600160665461094d60ff821615611465565b60ff1916176066557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346101b5575f3660031901126101b557602061099b61131b565b4211604051908152f35b346101b5575f3660031901126101b557602060ff606654166040519015158152f35b346101b5575f3660031901126101b5576109ec60018060a01b03603354163314611604565b6105506124ae565b346101b5575f3660031901126101b5576020606c54604051908152f35b346101b55760203660031901126101b5576004355f52606a602052602060405f2054604051908152f35b6001600160401b0381116102575760051b60200190565b9291610a5d82610a3b565b91610a6b6040519384610292565b829481845260208094019160051b81019283116101b557905b828210610a915750505050565b8380918335610a9f816101a4565b815201910190610a84565b9080601f830112156101b55781602061035a93359101610a52565b9291610ad082610a3b565b91610ade6040519384610292565b829481845260208094019160051b81019283116101b557905b828210610b045750505050565b81358152908301908301610af7565b9080601f830112156101b55781602061035a93359101610ac5565b91909160a0818403126101b55760405190610b488261025c565b81938135916001600160401b03928381116101b55782610b69918301610aaa565b845260208101359283116101b557610b876080939284938301610b13565b60208501526040810135604085015260608101356060850152013591610bac836101a4565b0152565b60ff8116036101b557565b9291610bc682610a3b565b91610bd46040519384610292565b829481845260208094019160051b81019283116101b557905b828210610bfa5750505050565b8380918335610c0881610bb0565b815201910190610bed565b9080601f830112156101b55781602061035a93359101610bbb565b346101b5576101403660031901126101b55760046001600160401b0381358181116101b557610c609036908401610b2e565b906024358181116101b557610c789036908501610c13565b916044358281116101b557610c909036908601610b13565b916064358181116101b557610ca89036908701610b13565b906084358181116101b557610cc09036908801610b13565b60a4358281116101b557610cd79036908901610aaa565b9160c4359081116101b55761055097610cf291369101610b13565b92610cfb6101b9565b95610124359760e4359661164f565b346101b5575f3660031901126101b557610d2261131b565b421115610d31576105506124ae565b60405162461bcd60e51b815260206004820152601960248201527f4f776e657273686970206e6f74207965742065787069726564000000000000006044820152606490fd5b346101b5575f3660031901126101b5576033546040516001600160a01b039091168152602090f35b346101b5575f3660031901126101b5576020606d54604051908152f35b908160a09103126101b55790565b9181601f840112156101b5578235916001600160401b0383116101b5576020808501948460051b0101116101b557565b346101b55760a03660031901126101b55760046001600160401b0381358181116101b557610e2a9036908401610dbb565b6024358281116101b557610e419036908501610dbb565b916044358181116101b557610e599036908601610dc9565b6064929192358281116101b557610e739036908801610dc9565b9390926084359081116101b55761055097610e9091369101610dc9565b969095611cce565b346101b55760203660031901126101b557600435610eb5816101a4565b60018060a01b03165f52606f602052602060ff60405f2054166040519015158152f35b346101b5575f3660031901126101b5576020606b54604051908152f35b346101b55760a03660031901126101b5576001600160401b036004358181116101b557610f269036906004016103f2565b50506024358181116101b557610f409036906004016103f2565b91906044358281116101b557610f5a9036906004016103f2565b60643591610f6783610bb0565b60405195611055948588019688881090881117610257578796610f8e96612a248939612112565b03905ff0801561056b576001600160a01b0316803b156101b5576040516340c10f1960e01b81523360048201526084356024820152905f908290604490829084905af1801561056b57610fdd57005b8061055f61055092610244565b346101b55760803660031901126101b5576001600160401b036044358181116101b55761101b903690600401610dc9565b6064929192359182116101b557611039611069923690600401610dc9565b915f549460ff8660081c1695861580978161108e575b61105890612147565b61107d575b50602435600435612200565b61106f57005b61055061ff00195f54165f55565b61ffff1916610101175f555f61105d565b5060ff82161561104f565b346101b5575f3660031901126101b5576020606e54604051908152f35b346101b5575f3660031901126101b5576020606854604051908152f35b346101b55760203660031901126101b5576004356110f0816101a4565b6033546001600160a01b03908116919061110b338414611604565b8116801561115d57610550927f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a360018060a01b03166bffffffffffffffffffffffff60a01b6033541617603355565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346101b55760803660031901126101b5576001600160401b036004358181116101b5576111e29036906004016103f2565b6024929192358281116101b5576111fd9036906004016103f2565b93906044358481116101b5576112179036906004016103f2565b9190946064359061122782610bb0565b6040519061105590818301908111838210176102575783868a8c88611252958897612a248939612112565b03905ff092831561056b577f82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7966112c19460018060a01b031698895f52606f60205260405f20600160ff198254161790556112b1610528606c546112da565b606c54946040519889988961231d565b0390a2005b634e487b7160e01b5f52601160045260245ffd5b90600182018092116112e857565b6112c6565b906509184e72a00082018092116112e857565b60300190816030116112e857565b919082018092116112e857565b6034546302f4bd0081018091116112e85790565b805191908290602001825e015f815290565b60405190602082018281106001600160401b03821117610257576040525f8252565b3d1561138d573d90611374826102c0565b916113826040519384610292565b82523d5f602084013e565b606090565b6040519061139f82610277565b600d82526c2ab735b737bbb71032b93937b960991b6020830152565b5f929161141961142785946113e96113e3604051956020815191012063ffffffff60e01b1690565b94610244565b604051936020850152600484526113ff84610277565b60405192839161141360208401809761132f565b9061132f565b03601f198101835282610292565b51915afa611433611363565b9015611442579061035a611341565b805161145f5750611451611392565b905b61145b611341565b9190565b90611453565b1561146c57565b60405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606490fd5b156114ab57565b60405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606490fd5b908160209103126101b5575190565b6040513d5f823e3d90fd5b919082039182116112e857565b908160209103126101b5575161035a81610bb0565b6020818303126101b5578051906001600160401b0382116101b5570181601f820112156101b55780519061155f826102c0565b9261156d6040519485610292565b828452602083830101116101b557815f9260208093018386015e8301015290565b6040519061159b82610277565b6002825261111160f11b6020830152565b61035a9392606092825260208201528160408201520190610311565b908060209392818452848401375f828201840152601f01601f1916010190565b61035a94926060928252602082015281604082015201916115c8565b1561160b57565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b928799929a979698949a611668600260675414156114a4565b600260675561167c60ff6066541615611465565b6001600160a01b0386165f9081526069602052604090206116a09089905410611946565b6001600160a01b0386165f9081526069602052604090206116cc906116c590546112ed565b89106119b1565b6116d7844310611a3f565b8b8a888751519761170460208201998a51518114908161193b575b81611930575b81611925575b50611ab1565b606d549061171e61171583836124ef565b60685414611afd565b611734855184518114908161191a575b50611b6f565b519851978a60409d6040519586956020870198611751968a611c2a565b03601f19810182526117639082610292565b51902092606e54946117749661264b565b6001600160a01b0381165f90815260696020526040812087905593845b885186101561189c576001600160a01b0383165f908152606f602052604090206117ba906104a4565b15611867576001600160a01b038316906117e46117d78887611cb5565b516001600160a01b031690565b6117ee888c611cb5565b51833b156101b55787516340c10f1960e01b81526001600160a01b039290921660048301526024820152915f908390604490829084905af190811561056b5760019261184c92611854575b505b611845888a611cb5565b519061130e565b950194611791565b8061055f61186192610244565b5f611839565b61184c6001916118978b6118888a6118826117d7828c611cb5565b92611cb5565b5190858060a01b03881661277a565b61183b565b94509650949350505080611900575b506118ba610528606c546112da565b606c546040519081526001600160a01b0392909216917f02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab70890602090a36101c76001606755565b61191490336001600160a01b03851661277a565b5f6118ab565b90508551145f61172e565b90508751145f6116fe565b8751811491506116f8565b8951811491506116f2565b1561194d57565b60405162461bcd60e51b815260206004820152603660248201527f4e6577206261746368206e6f6e6365206d7573742062652067726561746572206044820152757468616e207468652063757272656e74206e6f6e636560501b6064820152608490fd5b156119b857565b60405162461bcd60e51b815260206004820152605360248201527f4e6577206261746368206e6f6e6365206d757374206265206c6573732074686160448201527f6e2031305f3030305f3030305f3030305f3030302067726561746572207468616064820152726e207468652063757272656e74206e6f6e636560681b608482015260a490fd5b15611a4657565b60405162461bcd60e51b815260206004820152603b60248201527f42617463682074696d656f7574206d757374206265206772656174657220746860448201527f616e207468652063757272656e7420626c6f636b2068656967687400000000006064820152608490fd5b15611ab857565b60405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d65642063757272656e742076616c696461746f7220736574006044820152606490fd5b15611b0457565b60405162461bcd60e51b815260206004820152603f60248201527f537570706c6965642063757272656e742076616c696461746f727320616e642060448201527f706f7765727320646f206e6f74206d6174636820636865636b706f696e742e006064820152608490fd5b15611b7657565b60405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d6564206261746368206f66207472616e73616374696f6e73006044820152606490fd5b9081518082526020808093019301915f5b828110611bda575050505090565b835185529381019392810192600101611bcc565b9081518082526020808093019301915f5b828110611c0d575050505090565b83516001600160a01b031685529381019392810192600101611bff565b939060e09598979693611c6b611c8794611c79936101009089526f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b60208a01528060408a0152880190611bbb565b908682036060880152611bee565b908482036080860152611bbb565b60a08301969096526001600160a01b031660c08201520152565b634e487b7160e01b5f52603260045260245ffd5b8051821015611cc95760209160051b010190565b611ca1565b94919295939095611ce460ff6066541615611465565b604086013597604088013593611cfb858b11611f26565b611d058980611f98565b60208b01969150611d16878c611f98565b9190501480611f12575b80611efe575b80611ee5575b611d3590611ab1565b611d3e906112ed565b8a10611d4990611fcd565b606d5480611d57368c610b2e565b90611d61916124ef565b60685414611d6e90611afd565b611d78368a610b2e565b90611d82916124ef565b968794611d8f8b80611f98565b97611d9a919c611f98565b9890606e549c8d993690611dad92610a52565b993690611db992610ac5565b953690611dc592610bbb565b923690611dd192610ac5565b923690611ddd92610ac5565b92611de79661264b565b611df18280611f98565b90506020830193611e028585611f98565b90611e0c936127b6565b606855611e1883606b55565b60808101906001600160a01b03611e2e8361205c565b1615157f76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a9360609382611ea893611ed9575b611ead575b611e73610528606c546112da565b611e96611e82606c549261205c565b92611e8d8680611f98565b92909187611f98565b949093604051988998013590886120d1565b0390a2565b611ed4611ec8611ebc8361205c565b6001600160a01b031690565b8686013590339061277a565b611e65565b50848401351515611e60565b50611d3583611ef48c80611f98565b9050149050611d2c565b5081611f0a8b80611f98565b905014611d26565b5083611f1e8b80611f98565b905014611d20565b15611f2d57565b60405162461bcd60e51b815260206004820152603760248201527f4e65772076616c736574206e6f6e6365206d757374206265206772656174657260448201527f207468616e207468652063757272656e74206e6f6e63650000000000000000006064820152608490fd5b903590601e19813603018212156101b557018035906001600160401b0382116101b557602001918160051b360383136101b557565b15611fd457565b60405162461bcd60e51b815260206004820152605460248201527f4e65772076616c736574206e6f6e6365206d757374206265206c65737320746860448201527f616e2031305f3030305f3030305f3030305f3030302067726561746572207468606482015273616e207468652063757272656e74206e6f6e636560601b608482015260a490fd5b3561035a816101a4565b9190808252602080920192915f5b828110612082575050505090565b9091929382806001928735612096816101a4565b848060a01b03168152019501910192919092612074565b81835290916001600160fb1b0383116101b55760209260051b809284830137010190565b959391926121049361035a9896928852602088015260018060a01b0316604087015260a0606087015260a0860191612066565b9260808185039101526120ad565b93926040936121326121409360ff959998996060895260608901916115c8565b9186830360208801526115c8565b9416910152565b1561214e57565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b604051906121b78261025c565b5f60808360608152606060208201528260408201528260608201520152565b939161035a95936121049286525f60208701525f604087015260a0606087015260a0860191612066565b946122cb866122c6846122c16122bb8a9b8a61229f7f76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a9d6122958d9e9d61227c611ea89e8e5f5460ff8160081c16908115809281612312575b61226290612147565b612300575b506122ed575b612275612801565b85886127b6565b6122846121aa565b5061228d6102b3565b943691610a52565b83528a3691610ac5565b60208201525f60408201525f60608201525f60808201526124ef565b93606d55565b606e55565b606855565b6122d9610528606c546112da565b606b5495606c5493604051958695866121d6565b6122fb61ff00195f54165f55565b61226d565b61ffff1916610101175f908155612267565b5060ff821615612259565b949161235f9360ff9561234360809994612351949d9c9b9d60a08b5260a08b01916115c8565b9188830360208a01526115c8565b9185830360408701526115c8565b951660608201520152565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526101c7916123ac8261025c565b61290b565b5f1981146112e85760010190565b906123c9826102c0565b6123d66040519182610292565b82815280926123e7601f19916102c0565b0190602036910137565b80156112e8575f190190565b908151811015611cc9570160200190565b80156124905780815f925b61247c575080612428836123bf565b92915b61243457505090565b600a9061246761246161245161244b858506611300565b60ff1690565b60f81b6001600160f81b03191690565b936123f1565b925f1a61247484866123fd565b53048061242b565b91612488600a916123b1565b920480612419565b5060405161249d81610277565b60018152600360fc1b602082015290565b6033545f6001600160a01b0382167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a36001600160a01b031916603355565b61257b604082015191805192602082015191606081015190608060018060a01b0391015116612564612551604051978895602087019a8b526918da1958dadc1bda5b9d60b21b6040880152606087015260e06080870152610100860190611bee565b601f1995868683030160a0870152611bbb565b9160c084015260e083015203908101835282610292565b51902090565b1561258857565b60405162461bcd60e51b815260206004820152602360248201527f56616c696461746f72207369676e617475726520646f6573206e6f74206d617460448201526231b41760e91b6064820152608490fd5b156125e057565b60405162461bcd60e51b815260206004820152603c60248201527f5375626d69747465642076616c696461746f7220736574207369676e6174757260448201527f657320646f206e6f74206861766520656e6f75676820706f7765722e000000006064820152608490fd5b95949390935f935f5b88518110156127695760ff8061266a8388611cb5565b511661267a575b50600101612654565b90956001600160a01b038061268f898d611cb5565b51169261269c8989611cb5565b51166126a88986611cb5565b51906126b48a89611cb5565b5190604092835192602094858501947f19457468657265756d205369676e6564204d6573736167653a0a3332000000008652603c8c81830152815260608101938185106001600160401b03861117610257575f968560c094528251902085526080958683015260a0820152015282805260015afa1561056b576127489261273e915f511614612581565b6118458789611cb5565b94878611612756575f612671565b505050505090506101c792505b116125d9565b505050505090506101c79250612763565b60405163a9059cbb60e01b60208201526001600160a01b039290921660248301526044808301939093529181526101c7916123ac606483610292565b826127c49194939414611ab1565b5f905f5b8481106127dd575b50506101c79250116125d9565b8060051b82013583018093116112e8578383116127fc576001016127c8565b6127d0565b5f5460ff8160081c16908115809281612889575b61281e90612147565b612878575b50336bffffffffffffffffffffffff60a01b603354161760335542603455335f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a361286d57565b61ff00195f54165f55565b61ffff1916610101175f555f612823565b5060ff821615612815565b908160209103126101b5575180151581036101b55790565b156128b357565b60405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b6040516001600160a01b03919091169161292482610277565b6020928383527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656484840152803b1561299a575f82819282876129759796519301915af161296f611363565b906129df565b8051908161298257505050565b826101c793612995938301019101612894565b6128ac565b60405162461bcd60e51b815260048101859052601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b909190156129eb575090565b8151156129fb5750805190602001fd5b60405162461bcd60e51b815260206004820152908190612a1f906024830190610311565b0390fdfe60a06040523461032e576110558038038061001981610332565b928339810160608282031261032e5781516001600160401b039081811161032e5782610046918501610357565b906020928385015182811161032e57604091610063918701610357565b9401519360ff8516850361032e578251828111610245576003918254916001958684811c94168015610324575b88851014610310578190601f948581116102c2575b508890858311600114610264575f92610259575b50505f1982861b1c191690861b1783555b80519384116102455760049586548681811c9116801561023b575b82821014610228578381116101e5575b508092851160011461018057509383949184925f95610175575b50501b925f19911b1c19161790555b600580546001600160a01b0319163390811790915560405191905f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a3608052610cac90816103a98239608051816107090152f35b015193505f8061010f565b92919084601f198116885f52855f20955f905b898383106101cb57505050106101b2575b50505050811b01905561011e565b01519060f8845f19921b161c191690555f8080806101a4565b858701518955909701969485019488935090810190610193565b875f52815f208480880160051c82019284891061021f575b0160051c019087905b8281106102145750506100f5565b5f8155018790610206565b925081926101fd565b602288634e487b7160e01b5f525260245ffd5b90607f16906100e5565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100b9565b90889350601f19831691875f528a5f20925f5b8c8282106102ac5750508411610295575b505050811b0183556100ca565b01515f1983881b60f8161c191690555f8080610288565b8385015186558c97909501949384019301610277565b909150855f52885f208580850160051c8201928b8610610307575b918a91869594930160051c01915b8281106102f95750506100a5565b5f81558594508a91016102eb565b925081926102dd565b634e487b7160e01b5f52602260045260245ffd5b93607f1693610090565b5f80fd5b6040519190601f01601f191682016001600160401b0381118382101761024557604052565b81601f8201121561032e578051906001600160401b03821161024557610386601f8301601f1916602001610332565b928284526020838301011161032e57815f9260208093018386015e830101529056fe608060409080825260049081361015610016575f80fd5b5f3560e01c90816306fdde031461083857508063095ea7b31461080f57806318160ddd146107f157806323b872dd1461072d578063313ce567146106f057806339509351146106ab57806340c10f19146105d657806370a08231146105a0578063715018a6146105445780638da5cb5b1461051c57806395d89b41146103fc5780639dc29fac146102bf578063a457c2d714610213578063a9059cbb146101e3578063dd62ed3e1461019a5763f2fde38b146100d0575f80fd5b34610196576020366003190112610196576100e9610958565b600554916001600160a01b03808416926101043385146109b2565b1693841561014457505082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a36001600160a01b03191617600555005b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b5f80fd5b82346101965780600319360112610196576020906101b6610958565b6101be61096e565b9060018060a01b038091165f5260018452825f2091165f528252805f20549051908152f35b823461019657806003193601126101965760209061020c610202610958565b6024359033610afb565b5160018152f35b503461019657816003193601126101965761022c610958565b9060243590335f526001602052835f2060018060a01b0384165f52602052835f20549082821061026e5760208561020c866102678787610984565b90336109fd565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b50346101965781600319360112610196576102d8610958565b600554602435916001600160a01b03916102f590831633146109b2565b169182156103af57825f525f602052835f2054908282106103615750815f946103417fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef94602094610984565b8587528684528187205561035782600254610984565b60025551908152a3005b608490602086519162461bcd60e51b8352820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152fd5b608490602085519162461bcd60e51b8352820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152fd5b509034610196575f366003190112610196578051905f835460018160011c9060018316928315610512575b60209384841081146104ff578388529081156104e3575060011461048f575b505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b0390f35b604190634e487b7160e01b5f525260245ffd5b5f878152929350837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8385106104cf57505050508301015f8080610446565b8054888601830152930192849082016104b9565b60ff1916878501525050151560051b84010190505f8080610446565b602289634e487b7160e01b5f525260245ffd5b91607f1691610427565b8234610196575f3660031901126101965760055490516001600160a01b039091168152602090f35b34610196575f366003190112610196576005545f6001600160a01b03821661056d3382146109b2565b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a36001600160a01b031916600555005b8234610196576020366003190112610196576020906001600160a01b036105c5610958565b165f525f8252805f20549051908152f35b5090346101965780600319360112610196576105f0610958565b60055460243592916001600160a01b039161060e90831633146109b2565b1692831561066957506020827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926106495f956002546109a5565b6002558585528483528085206106608382546109a5565b905551908152a3005b6020606492519162461bcd60e51b8352820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b823461019657806003193601126101965760209061020c6106ca610958565b335f5260018452825f2060018060a01b0382165f528452610267602435845f20546109a5565b8234610196575f366003190112610196576020905160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461019657606036600319011261019657610747610958565b9061075061096e565b9061075f604435809385610afb565b60018060a01b0383165f526001602052835f20335f52602052835f20549082821061079d5760208561020c866107958787610984565b9033906109fd565b608490602086519162461bcd60e51b8352820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152fd5b8234610196575f366003190112610196576020906002549051908152f35b823461019657806003193601126101965760209061020c61082e610958565b60243590336109fd565b90508234610196575f366003190112610196575f60035460018160011c9060018316928315610924575b60209384841081146104ff5783885290811561090857506001146108b257505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b60035f908152929350837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8385106108f45750505050830101848080610446565b8054888601830152930192849082016108de565b60ff1916878501525050151560051b8401019050848080610446565b91607f1691610862565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361019657565b602435906001600160a01b038216820361019657565b9190820391821161099157565b634e487b7160e01b5f52601160045260245ffd5b9190820180921161099157565b156109b957565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6001600160a01b03908116918215610aaa5716918215610a5a5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b6001600160a01b03908116918215610c235716918215610bd257815f525f60205260405f2054818110610b7e5781610b567fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef93602093610984565b845f525f835260405f2055845f5260405f20610b738282546109a5565b9055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fdfea2646970667358221220418daa12254165c4c9282976d969382b8b711ad2a05322ad1712abe67c4e3ebe64736f6c63430008190033a2646970667358221220ef55bd14af6ffe8b843fea8d7c13e4fb7aa1251ec97526bd7db7bf6682d9c44a64736f6c63430008190033",
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

// HyperionSubgraphMetaData contains all meta data concerning the HyperionSubgraph contract.
var HyperionSubgraphMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"ERC20DeployedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"SendToCosmosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"SendToHeliosEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"}],\"name\":\"TransactionBatchExecutedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_newValsetNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_eventNonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"ValsetUpdatedEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_methodName\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_args\",\"type\":\"bytes\"}],\"name\":\"callData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"err\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_cosmosDenom\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"deployERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"name\":\"deployERC20WithSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyUnpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnershipExpiryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_powerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isHeliosNativeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwnershipExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_erc20Address\",\"type\":\"address\"}],\"name\":\"lastBatchNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnershipAfterExpiry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"sendToCosmos\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_destination\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_data\",\"type\":\"string\"}],\"name\":\"sendToHelios\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_hyperionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"state_invalidationMapping\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"state_lastBatchNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastEventNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_lastValsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_powerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_destinations\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_batchNonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_batchTimeout\",\"type\":\"uint256\"}],\"name\":\"submitBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_newValset\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"}],\"internalType\":\"structValsetArgs\",\"name\":\"_currentValset\",\"type\":\"tuple\"},{\"internalType\":\"uint8[]\",\"name\":\"_v\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"updateValset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60808060405234602c5760ff196066541660665560016067555f606b555f606c55613b6690816100318239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c8063011b21741461014b5780631ee7a108146101af5780631ffbe7f9146101aa5780632b7553c7146101a5578063308ff208146101a05780634a4e3bd51461019b57806351858e27146101965780635afe97bb146101915780635c975abb1461018c578063715018a61461018757806373b20547146101825780637dfb6f861461017d57806381747418146101785780638c64865f146101735780638da5cb5b1461016e578063a4b52ca214610169578063a5352f5b14610164578063a6c42b021461015f578063b56561fe1461015a578063c2d0732e14610155578063c359a21214610150578063df97174b1461014b578063e5a2b5d214610146578063f2b5330714610141578063f2fde38b1461013c5763f795563714610137575f80fd5b61126e565b611190565b611173565b611156565b6101d9565b6110a7565b610fb2565b610f95565b610f55565b610eb6565b610e5b565b610e33565b610dc7565b610ceb565b610ace565b610ab1565b610a84565b610a62565b610a3e565b6109d3565b610933565b6104dc565b61041a565b610240565b61021e565b6001600160a01b038116036101c557565b5f80fd5b61010435906101d7826101b4565b565b346101c55760203660031901126101c5576004356101f6816101b4565b60018060a01b03165f526069602052602060405f2054604051908152f35b5f9103126101c557565b346101c5575f3660031901126101c55760206102386113d3565b604051908152f35b346101c55760603660031901126101c55760043561025d816101b4565b6044359061027060ff60665416156113e7565b61027f60026067541415611426565b60026067556001600160a01b03169061029a81303385612422565b606c54600181018091116102e85780606c556040519182526020820152602435917fd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e460403393a46001606755005b611383565b634e487b7160e01b5f52604160045260245ffd5b6001600160401b03811161031457604052565b6102ed565b60a081019081106001600160401b0382111761031457604052565b604081019081106001600160401b0382111761031457604052565b90601f801991011681019081106001600160401b0382111761031457604052565b604051906101d782610319565b6001600160401b03811161031457601f01601f191660200190565b9291926103a48261037d565b916103b2604051938461034f565b8294818452818301116101c5578281602093845f960137010152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b9091610409610417936040845260408401906103ce565b9160208184039101526103ce565b90565b346101c55760603660031901126101c557600435610437816101b4565b6001600160401b03906024358281116101c557366023820112156101c557610469903690602481600401359101610398565b6044359283116101c557366023840112156101c55761049561049b933690602481600401359101610398565b916114fe565b906104ab604051928392836103f2565b0390f35b9181601f840112156101c5578235916001600160401b0383116101c557602083818601950101116101c557565b346101c55760803660031901126101c55760048035906104fb826101b4565b60243591604435916064356001600160401b0381116101c55761052190369083016104af565b9261053160ff60665416156113e7565b61054060026067541415611426565b60026067556001600160a01b0381165f908152606f60205260409020610568905b5460ff1690565b1561062d576001600160a01b031692833b156101c55760408051632770a7eb60e21b81523394810194855260208501879052935f918591829101038183885af1928315610628577f272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b479361060f575b506105ea6105e5606c54611397565b606c55565b6105ff606c54916040519384933398856116a0565b0390a45b61060d6001606755565b005b8061061c61062292610301565b80610214565b5f6105d6565b6115b7565b604080516370a0823160e01b8082523086830190815296976001600160a01b0394909416969293602093919291849083908190830103818b5afa918215610628575f92610910575b506106829030338a612422565b83519182523086830190815283908390819060200103818a5afa8015610628576106b3925f916108e3575b506115c2565b936106c26105e5606c54611397565b825163313ce56760e01b815293828583818a5afa948515610628575f956108b4575b50606c54978451956395d89b4160e01b87525f8785818c5afa968715610628575f97610898575b505f865180956306fdde0360e01b8252818c5afa938415610628575f9461086d575b5060ff61073a91166124c6565b91811561085d5761074c913691610398565b915b8451607b60f81b9481019485529586946001016c226d65746164617461223a207b60981b8152600d016a1139bcb6b137b6111d101160a91b8152600b0161079491611472565b61088b60f21b815260020168113730b6b2911d101160b91b81526009016107ba91611472565b61088b60f21b81526002016b0113232b1b4b6b0b639911d160a51b8152600c016107e391611472565b611f4b60f21b8152600201670113230ba30911d160c51b815260080161080891611472565b607d60f81b815260010103601f1981018352610824908361034f565b5191829133956108349284611664565b037f272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b4791a4610603565b5050610867611646565b9161074e565b61073a91945061089060ff913d805f833e610888818361034f565b8101906115e4565b94915061072d565b6108ad9197503d805f833e610888818361034f565b955f61070b565b6108d5919550833d85116108dc575b6108cd818361034f565b8101906115cf565b935f6106e4565b503d6108c3565b6109039150843d8611610909575b6108fb818361034f565b8101906115a8565b5f6106ad565b503d6108f1565b61068291925061092c90853d8711610909576108fb818361034f565b9190610675565b346101c5575f3660031901126101c55761095860018060a01b036033541633146116bc565b60665460ff8116156109975760ff19166066557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b60405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606490fd5b346101c5575f3660031901126101c5576109f860018060a01b036033541633146116bc565b6001606654610a0a60ff8216156113e7565b60ff1916176066557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b346101c5575f3660031901126101c5576020610a586113d3565b4211604051908152f35b346101c5575f3660031901126101c557602060ff606654166040519015158152f35b346101c5575f3660031901126101c557610aa960018060a01b036033541633146116bc565b61060d612566565b346101c5575f3660031901126101c5576020606c54604051908152f35b346101c55760203660031901126101c5576004355f52606a602052602060405f2054604051908152f35b6001600160401b0381116103145760051b60200190565b9291610b1a82610af8565b91610b28604051938461034f565b829481845260208094019160051b81019283116101c557905b828210610b4e5750505050565b8380918335610b5c816101b4565b815201910190610b41565b9080601f830112156101c55781602061041793359101610b0f565b9291610b8d82610af8565b91610b9b604051938461034f565b829481845260208094019160051b81019283116101c557905b828210610bc15750505050565b81358152908301908301610bb4565b9080601f830112156101c55781602061041793359101610b82565b91909160a0818403126101c55760405190610c0582610319565b81938135916001600160401b03928381116101c55782610c26918301610b67565b845260208101359283116101c557610c446080939284938301610bd0565b60208501526040810135604085015260608101356060850152013591610c69836101b4565b0152565b60ff8116036101c557565b9291610c8382610af8565b91610c91604051938461034f565b829481845260208094019160051b81019283116101c557905b828210610cb75750505050565b8380918335610cc581610c6d565b815201910190610caa565b9080601f830112156101c55781602061041793359101610c78565b346101c5576101403660031901126101c55760046001600160401b0381358181116101c557610d1d9036908401610beb565b906024358181116101c557610d359036908501610cd0565b916044358281116101c557610d4d9036908601610bd0565b916064358181116101c557610d659036908701610bd0565b906084358181116101c557610d7d9036908801610bd0565b60a4358281116101c557610d949036908901610b67565b9160c4359081116101c55761060d97610daf91369101610bd0565b92610db86101c9565b95610124359760e43596611707565b346101c5575f3660031901126101c557610ddf6113d3565b421115610dee5761060d612566565b60405162461bcd60e51b815260206004820152601960248201527f4f776e657273686970206e6f74207965742065787069726564000000000000006044820152606490fd5b346101c5575f3660031901126101c5576033546040516001600160a01b039091168152602090f35b346101c5575f3660031901126101c5576020606d54604051908152f35b908160a09103126101c55790565b9181601f840112156101c5578235916001600160401b0383116101c5576020808501948460051b0101116101c557565b346101c55760a03660031901126101c55760046001600160401b0381358181116101c557610ee79036908401610e78565b6024358281116101c557610efe9036908501610e78565b916044358181116101c557610f169036908601610e86565b6064929192358281116101c557610f309036908801610e86565b9390926084359081116101c55761060d97610f4d91369101610e86565b969095611d86565b346101c55760203660031901126101c557600435610f72816101b4565b60018060a01b03165f52606f602052602060ff60405f2054166040519015158152f35b346101c5575f3660031901126101c5576020606b54604051908152f35b346101c55760a03660031901126101c5576001600160401b036004358181116101c557610fe39036906004016104af565b50506024358181116101c557610ffd9036906004016104af565b91906044358281116101c5576110179036906004016104af565b6064359161102483610c6d565b6040519561105594858801968888109088111761031457879661104b96612adc89396121ca565b03905ff08015610628576001600160a01b0316803b156101c5576040516340c10f1960e01b81523360048201526084356024820152905f908290604490829084905af180156106285761109a57005b8061061c61060d92610301565b346101c55760803660031901126101c5576001600160401b036044358181116101c5576110d8903690600401610e86565b6064929192359182116101c5576110f6611126923690600401610e86565b915f549460ff8660081c1695861580978161114b575b611115906121ff565b61113a575b506024356004356122b8565b61112c57005b61060d61ff00195f54165f55565b61ffff1916610101175f555f61111a565b5060ff82161561110c565b346101c5575f3660031901126101c5576020606e54604051908152f35b346101c5575f3660031901126101c5576020606854604051908152f35b346101c55760203660031901126101c5576004356111ad816101b4565b6033546001600160a01b0390811691906111c83384146116bc565b8116801561121a5761060d927f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a360018060a01b03166bffffffffffffffffffffffff60a01b6033541617603355565b60405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608490fd5b346101c55760803660031901126101c5576001600160401b036004358181116101c55761129f9036906004016104af565b6024929192358281116101c5576112ba9036906004016104af565b93906044358481116101c5576112d49036906004016104af565b919094606435906112e482610c6d565b6040519061105590818301908111838210176103145783868a8c8861130f958897612adc89396121ca565b03905ff0928315610628577f82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c79661137e9460018060a01b031698895f52606f60205260405f20600160ff1982541617905561136e6105e5606c54611397565b606c5494604051988998896123d5565b0390a2005b634e487b7160e01b5f52601160045260245ffd5b90600182018092116102e857565b906509184e72a00082018092116102e857565b60300190816030116102e857565b919082018092116102e857565b6034546302f4bd0081018091116102e85790565b156113ee57565b60405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606490fd5b1561142d57565b60405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606490fd5b805191908290602001825e015f815290565b60405190602082018281106001600160401b03821117610314576040525f8252565b3d156114d0573d906114b78261037d565b916114c5604051938461034f565b82523d5f602084013e565b606090565b604051906114e282610334565b600d82526c2ab735b737bbb71032b93937b960991b6020830152565b5f929161155c61156a859461152c611526604051956020815191012063ffffffff60e01b1690565b94610301565b6040519360208501526004845261154284610334565b604051928391611556602084018097611472565b90611472565b03601f19810183528261034f565b51915afa6115766114a6565b90156115855790610417611484565b80516115a257506115946114d5565b905b61159e611484565b9190565b90611596565b908160209103126101c5575190565b6040513d5f823e3d90fd5b919082039182116102e857565b908160209103126101c5575161041781610c6d565b6020818303126101c5578051906001600160401b0382116101c5570181601f820112156101c5578051906116178261037d565b92611625604051948561034f565b828452602083830101116101c557815f9260208093018386015e8301015290565b6040519061165382610334565b6002825261111160f11b6020830152565b61041793926060928252602082015281604082015201906103ce565b908060209392818452848401375f828201840152601f01601f1916010190565b6104179492606092825260208201528160408201520191611680565b156116c357565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b928799929a979698949a61172060026067541415611426565b600260675561173460ff60665416156113e7565b6001600160a01b0386165f90815260696020526040902061175890899054106119fe565b6001600160a01b0386165f9081526069602052604090206117849061177d90546113a5565b8910611a69565b61178f844310611af7565b8b8a88875151976117bc60208201998a5151811490816119f3575b816119e8575b816119dd575b50611b69565b606d54906117d66117cd83836125a7565b60685414611bb5565b6117ec85518451811490816119d2575b50611c27565b519851978a60409d6040519586956020870198611809968a611ce2565b03601f198101825261181b908261034f565b51902092606e549461182c96612703565b6001600160a01b0381165f90815260696020526040812087905593845b8851861015611954576001600160a01b0383165f908152606f6020526040902061187290610561565b1561191f576001600160a01b0383169061189c61188f8887611d6d565b516001600160a01b031690565b6118a6888c611d6d565b51833b156101c55787516340c10f1960e01b81526001600160a01b039290921660048301526024820152915f908390604490829084905af1908115610628576001926119049261190c575b505b6118fd888a611d6d565b51906113c6565b950194611849565b8061061c61191992610301565b5f6118f1565b61190460019161194f8b6119408a61193a61188f828c611d6d565b92611d6d565b5190858060a01b038816612832565b6118f3565b945096509493505050806119b8575b506119726105e5606c54611397565b606c546040519081526001600160a01b0392909216917f02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab70890602090a36101d76001606755565b6119cc90336001600160a01b038516612832565b5f611963565b90508551145f6117e6565b90508751145f6117b6565b8751811491506117b0565b8951811491506117aa565b15611a0557565b60405162461bcd60e51b815260206004820152603660248201527f4e6577206261746368206e6f6e6365206d7573742062652067726561746572206044820152757468616e207468652063757272656e74206e6f6e636560501b6064820152608490fd5b15611a7057565b60405162461bcd60e51b815260206004820152605360248201527f4e6577206261746368206e6f6e6365206d757374206265206c6573732074686160448201527f6e2031305f3030305f3030305f3030305f3030302067726561746572207468616064820152726e207468652063757272656e74206e6f6e636560681b608482015260a490fd5b15611afe57565b60405162461bcd60e51b815260206004820152603b60248201527f42617463682074696d656f7574206d757374206265206772656174657220746860448201527f616e207468652063757272656e7420626c6f636b2068656967687400000000006064820152608490fd5b15611b7057565b60405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d65642063757272656e742076616c696461746f7220736574006044820152606490fd5b15611bbc57565b60405162461bcd60e51b815260206004820152603f60248201527f537570706c6965642063757272656e742076616c696461746f727320616e642060448201527f706f7765727320646f206e6f74206d6174636820636865636b706f696e742e006064820152608490fd5b15611c2e57565b60405162461bcd60e51b815260206004820152601f60248201527f4d616c666f726d6564206261746368206f66207472616e73616374696f6e73006044820152606490fd5b9081518082526020808093019301915f5b828110611c92575050505090565b835185529381019392810192600101611c84565b9081518082526020808093019301915f5b828110611cc5575050505090565b83516001600160a01b031685529381019392810192600101611cb7565b939060e09598979693611d23611d3f94611d31936101009089526f0e8e4c2dce6c2c6e8d2dedc84c2e8c6d60831b60208a01528060408a0152880190611c73565b908682036060880152611ca6565b908482036080860152611c73565b60a08301969096526001600160a01b031660c08201520152565b634e487b7160e01b5f52603260045260245ffd5b8051821015611d815760209160051b010190565b611d59565b94919295939095611d9c60ff60665416156113e7565b604086013597604088013593611db3858b11611fde565b611dbd8980612050565b60208b01969150611dce878c612050565b9190501480611fca575b80611fb6575b80611f9d575b611ded90611b69565b611df6906113a5565b8a10611e0190612085565b606d5480611e0f368c610beb565b90611e19916125a7565b60685414611e2690611bb5565b611e30368a610beb565b90611e3a916125a7565b968794611e478b80612050565b97611e52919c612050565b9890606e549c8d993690611e6592610b0f565b993690611e7192610b82565b953690611e7d92610c78565b923690611e8992610b82565b923690611e9592610b82565b92611e9f96612703565b611ea98280612050565b90506020830193611eba8585612050565b90611ec49361286e565b606855611ed083606b55565b60808101906001600160a01b03611ee683612114565b1615157f76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a9360609382611f6093611f91575b611f65575b611f2b6105e5606c54611397565b611f4e611f3a606c5492612114565b92611f458680612050565b92909187612050565b94909360405198899801359088612189565b0390a2565b611f8c611f80611f7483612114565b6001600160a01b031690565b86860135903390612832565b611f1d565b50848401351515611f18565b50611ded83611fac8c80612050565b9050149050611de4565b5081611fc28b80612050565b905014611dde565b5083611fd68b80612050565b905014611dd8565b15611fe557565b60405162461bcd60e51b815260206004820152603760248201527f4e65772076616c736574206e6f6e6365206d757374206265206772656174657260448201527f207468616e207468652063757272656e74206e6f6e63650000000000000000006064820152608490fd5b903590601e19813603018212156101c557018035906001600160401b0382116101c557602001918160051b360383136101c557565b1561208c57565b60405162461bcd60e51b815260206004820152605460248201527f4e65772076616c736574206e6f6e6365206d757374206265206c65737320746860448201527f616e2031305f3030305f3030305f3030305f3030302067726561746572207468606482015273616e207468652063757272656e74206e6f6e636560601b608482015260a490fd5b35610417816101b4565b9190808252602080920192915f5b82811061213a575050505090565b909192938280600192873561214e816101b4565b848060a01b0316815201950191019291909261212c565b81835290916001600160fb1b0383116101c55760209260051b809284830137010190565b959391926121bc936104179896928852602088015260018060a01b0316604087015260a0606087015260a086019161211e565b926080818503910152612165565b93926040936121ea6121f89360ff95999899606089526060890191611680565b918683036020880152611680565b9416910152565b1561220657565b60405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b6040519061226f82610319565b5f60808360608152606060208201528260408201528260608201520152565b939161041795936121bc9286525f60208701525f604087015260a0606087015260a086019161211e565b946123838661237e846123796123738a9b8a6123577f76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a9d61234d8d9e9d612334611f609e8e5f5460ff8160081c169081158092816123ca575b61231a906121ff565b6123b8575b506123a5575b61232d6128b9565b858861286e565b61233c612262565b50612345610370565b943691610b0f565b83528a3691610b82565b60208201525f60408201525f60608201525f60808201526125a7565b93606d55565b606e55565b606855565b6123916105e5606c54611397565b606b5495606c54936040519586958661228e565b6123b361ff00195f54165f55565b612325565b61ffff1916610101175f90815561231f565b5060ff821615612311565b94916124179360ff956123fb60809994612409949d9c9b9d60a08b5260a08b0191611680565b9188830360208a0152611680565b918583036040870152611680565b951660608201520152565b6040516323b872dd60e01b60208201526001600160a01b0392831660248201529290911660448301526064808301939093529181526101d79161246482610319565b6129c3565b5f1981146102e85760010190565b906124818261037d565b61248e604051918261034f565b828152809261249f601f199161037d565b0190602036910137565b80156102e8575f190190565b908151811015611d81570160200190565b80156125485780815f925b6125345750806124e083612477565b92915b6124ec57505090565b600a9061251f6125196125096125038585066113b8565b60ff1690565b60f81b6001600160f81b03191690565b936124a9565b925f1a61252c84866124b5565b5304806124e3565b91612540600a91612469565b9204806124d1565b5060405161255581610334565b60018152600360fc1b602082015290565b6033545f6001600160a01b0382167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a36001600160a01b031916603355565b612633604082015191805192602082015191606081015190608060018060a01b039101511661261c612609604051978895602087019a8b526918da1958dadc1bda5b9d60b21b6040880152606087015260e06080870152610100860190611ca6565b601f1995868683030160a0870152611c73565b9160c084015260e08301520390810183528261034f565b51902090565b1561264057565b60405162461bcd60e51b815260206004820152602360248201527f56616c696461746f72207369676e617475726520646f6573206e6f74206d617460448201526231b41760e91b6064820152608490fd5b1561269857565b60405162461bcd60e51b815260206004820152603c60248201527f5375626d69747465642076616c696461746f7220736574207369676e6174757260448201527f657320646f206e6f74206861766520656e6f75676820706f7765722e000000006064820152608490fd5b95949390935f935f5b88518110156128215760ff806127228388611d6d565b5116612732575b5060010161270c565b90956001600160a01b0380612747898d611d6d565b5116926127548989611d6d565b51166127608986611d6d565b519061276c8a89611d6d565b5190604092835192602094858501947f19457468657265756d205369676e6564204d6573736167653a0a3332000000008652603c8c81830152815260608101938185106001600160401b03861117610314575f968560c094528251902085526080958683015260a0820152015282805260015afa1561062857612800926127f6915f511614612639565b6118fd8789611d6d565b9487861161280e575f612729565b505050505090506101d792505b11612691565b505050505090506101d7925061281b565b60405163a9059cbb60e01b60208201526001600160a01b039290921660248301526044808301939093529181526101d79161246460648361034f565b8261287c9194939414611b69565b5f905f5b848110612895575b50506101d7925011612691565b8060051b82013583018093116102e8578383116128b457600101612880565b612888565b5f5460ff8160081c16908115809281612941575b6128d6906121ff565b612930575b50336bffffffffffffffffffffffff60a01b603354161760335542603455335f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a361292557565b61ff00195f54165f55565b61ffff1916610101175f555f6128db565b5060ff8216156128cd565b908160209103126101c5575180151581036101c55790565b1561296b57565b60405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608490fd5b6040516001600160a01b0391909116916129dc82610334565b6020928383527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656484840152803b15612a52575f8281928287612a2d9796519301915af1612a276114a6565b90612a97565b80519081612a3a57505050565b826101d793612a4d93830101910161294c565b612964565b60405162461bcd60e51b815260048101859052601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b90919015612aa3575090565b815115612ab35750805190602001fd5b60405162461bcd60e51b815260206004820152908190612ad79060248301906103ce565b0390fdfe60a06040523461032e576110558038038061001981610332565b928339810160608282031261032e5781516001600160401b039081811161032e5782610046918501610357565b906020928385015182811161032e57604091610063918701610357565b9401519360ff8516850361032e578251828111610245576003918254916001958684811c94168015610324575b88851014610310578190601f948581116102c2575b508890858311600114610264575f92610259575b50505f1982861b1c191690861b1783555b80519384116102455760049586548681811c9116801561023b575b82821014610228578381116101e5575b508092851160011461018057509383949184925f95610175575b50501b925f19911b1c19161790555b600580546001600160a01b0319163390811790915560405191905f7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a3608052610cac90816103a98239608051816107090152f35b015193505f8061010f565b92919084601f198116885f52855f20955f905b898383106101cb57505050106101b2575b50505050811b01905561011e565b01519060f8845f19921b161c191690555f8080806101a4565b858701518955909701969485019488935090810190610193565b875f52815f208480880160051c82019284891061021f575b0160051c019087905b8281106102145750506100f5565b5f8155018790610206565b925081926101fd565b602288634e487b7160e01b5f525260245ffd5b90607f16906100e5565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100b9565b90889350601f19831691875f528a5f20925f5b8c8282106102ac5750508411610295575b505050811b0183556100ca565b01515f1983881b60f8161c191690555f8080610288565b8385015186558c97909501949384019301610277565b909150855f52885f208580850160051c8201928b8610610307575b918a91869594930160051c01915b8281106102f95750506100a5565b5f81558594508a91016102eb565b925081926102dd565b634e487b7160e01b5f52602260045260245ffd5b93607f1693610090565b5f80fd5b6040519190601f01601f191682016001600160401b0381118382101761024557604052565b81601f8201121561032e578051906001600160401b03821161024557610386601f8301601f1916602001610332565b928284526020838301011161032e57815f9260208093018386015e830101529056fe608060409080825260049081361015610016575f80fd5b5f3560e01c90816306fdde031461083857508063095ea7b31461080f57806318160ddd146107f157806323b872dd1461072d578063313ce567146106f057806339509351146106ab57806340c10f19146105d657806370a08231146105a0578063715018a6146105445780638da5cb5b1461051c57806395d89b41146103fc5780639dc29fac146102bf578063a457c2d714610213578063a9059cbb146101e3578063dd62ed3e1461019a5763f2fde38b146100d0575f80fd5b34610196576020366003190112610196576100e9610958565b600554916001600160a01b03808416926101043385146109b2565b1693841561014457505082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a36001600160a01b03191617600555005b906020608492519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b5f80fd5b82346101965780600319360112610196576020906101b6610958565b6101be61096e565b9060018060a01b038091165f5260018452825f2091165f528252805f20549051908152f35b823461019657806003193601126101965760209061020c610202610958565b6024359033610afb565b5160018152f35b503461019657816003193601126101965761022c610958565b9060243590335f526001602052835f2060018060a01b0384165f52602052835f20549082821061026e5760208561020c866102678787610984565b90336109fd565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b50346101965781600319360112610196576102d8610958565b600554602435916001600160a01b03916102f590831633146109b2565b169182156103af57825f525f602052835f2054908282106103615750815f946103417fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef94602094610984565b8587528684528187205561035782600254610984565b60025551908152a3005b608490602086519162461bcd60e51b8352820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152fd5b608490602085519162461bcd60e51b8352820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152fd5b509034610196575f366003190112610196578051905f835460018160011c9060018316928315610512575b60209384841081146104ff578388529081156104e3575060011461048f575b505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b0390f35b604190634e487b7160e01b5f525260245ffd5b5f878152929350837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8385106104cf57505050508301015f8080610446565b8054888601830152930192849082016104b9565b60ff1916878501525050151560051b84010190505f8080610446565b602289634e487b7160e01b5f525260245ffd5b91607f1691610427565b8234610196575f3660031901126101965760055490516001600160a01b039091168152602090f35b34610196575f366003190112610196576005545f6001600160a01b03821661056d3382146109b2565b7f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a36001600160a01b031916600555005b8234610196576020366003190112610196576020906001600160a01b036105c5610958565b165f525f8252805f20549051908152f35b5090346101965780600319360112610196576105f0610958565b60055460243592916001600160a01b039161060e90831633146109b2565b1692831561066957506020827fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef926106495f956002546109a5565b6002558585528483528085206106608382546109a5565b905551908152a3005b6020606492519162461bcd60e51b8352820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b823461019657806003193601126101965760209061020c6106ca610958565b335f5260018452825f2060018060a01b0382165f528452610267602435845f20546109a5565b8234610196575f366003190112610196576020905160ff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b503461019657606036600319011261019657610747610958565b9061075061096e565b9061075f604435809385610afb565b60018060a01b0383165f526001602052835f20335f52602052835f20549082821061079d5760208561020c866107958787610984565b9033906109fd565b608490602086519162461bcd60e51b8352820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b6064820152fd5b8234610196575f366003190112610196576020906002549051908152f35b823461019657806003193601126101965760209061020c61082e610958565b60243590336109fd565b90508234610196575f366003190112610196575f60035460018160011c9060018316928315610924575b60209384841081146104ff5783885290811561090857506001146108b257505050829003601f01601f191682019267ffffffffffffffff84118385101761047c575082918261047892528261092e565b60035f908152929350837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8385106108f45750505050830101848080610446565b8054888601830152930192849082016108de565b60ff1916878501525050151560051b8401019050848080610446565b91607f1691610862565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361019657565b602435906001600160a01b038216820361019657565b9190820391821161099157565b634e487b7160e01b5f52601160045260245ffd5b9190820180921161099157565b156109b957565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fd5b6001600160a01b03908116918215610aaa5716918215610a5a5760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b6001600160a01b03908116918215610c235716918215610bd257815f525f60205260405f2054818110610b7e5781610b567fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef93602093610984565b845f525f835260405f2055845f5260405f20610b738282546109a5565b9055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fdfea2646970667358221220418daa12254165c4c9282976d969382b8b711ad2a05322ad1712abe67c4e3ebe64736f6c63430008190033a264697066735822122024b2b0fd1bd8cc125c4a9d7dfe3460a4241385109697b7a9721a4beace8d926464736f6c63430008190033",
}

// HyperionSubgraphABI is the input ABI used to generate the binding from.
// Deprecated: Use HyperionSubgraphMetaData.ABI instead.
var HyperionSubgraphABI = HyperionSubgraphMetaData.ABI

// HyperionSubgraphBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HyperionSubgraphMetaData.Bin instead.
var HyperionSubgraphBin = HyperionSubgraphMetaData.Bin

// DeployHyperionSubgraph deploys a new Ethereum contract, binding an instance of HyperionSubgraph to it.
func DeployHyperionSubgraph(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HyperionSubgraph, error) {
	parsed, err := HyperionSubgraphMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HyperionSubgraphBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HyperionSubgraph{HyperionSubgraphCaller: HyperionSubgraphCaller{contract: contract}, HyperionSubgraphTransactor: HyperionSubgraphTransactor{contract: contract}, HyperionSubgraphFilterer: HyperionSubgraphFilterer{contract: contract}}, nil
}

// HyperionSubgraph is an auto generated Go binding around an Ethereum contract.
type HyperionSubgraph struct {
	HyperionSubgraphCaller     // Read-only binding to the contract
	HyperionSubgraphTransactor // Write-only binding to the contract
	HyperionSubgraphFilterer   // Log filterer for contract events
}

// HyperionSubgraphCaller is an auto generated read-only Go binding around an Ethereum contract.
type HyperionSubgraphCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionSubgraphTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HyperionSubgraphTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionSubgraphFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HyperionSubgraphFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HyperionSubgraphSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HyperionSubgraphSession struct {
	Contract     *HyperionSubgraph // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HyperionSubgraphCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HyperionSubgraphCallerSession struct {
	Contract *HyperionSubgraphCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// HyperionSubgraphTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HyperionSubgraphTransactorSession struct {
	Contract     *HyperionSubgraphTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// HyperionSubgraphRaw is an auto generated low-level Go binding around an Ethereum contract.
type HyperionSubgraphRaw struct {
	Contract *HyperionSubgraph // Generic contract binding to access the raw methods on
}

// HyperionSubgraphCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HyperionSubgraphCallerRaw struct {
	Contract *HyperionSubgraphCaller // Generic read-only contract binding to access the raw methods on
}

// HyperionSubgraphTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HyperionSubgraphTransactorRaw struct {
	Contract *HyperionSubgraphTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHyperionSubgraph creates a new instance of HyperionSubgraph, bound to a specific deployed contract.
func NewHyperionSubgraph(address common.Address, backend bind.ContractBackend) (*HyperionSubgraph, error) {
	contract, err := bindHyperionSubgraph(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraph{HyperionSubgraphCaller: HyperionSubgraphCaller{contract: contract}, HyperionSubgraphTransactor: HyperionSubgraphTransactor{contract: contract}, HyperionSubgraphFilterer: HyperionSubgraphFilterer{contract: contract}}, nil
}

// NewHyperionSubgraphCaller creates a new read-only instance of HyperionSubgraph, bound to a specific deployed contract.
func NewHyperionSubgraphCaller(address common.Address, caller bind.ContractCaller) (*HyperionSubgraphCaller, error) {
	contract, err := bindHyperionSubgraph(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphCaller{contract: contract}, nil
}

// NewHyperionSubgraphTransactor creates a new write-only instance of HyperionSubgraph, bound to a specific deployed contract.
func NewHyperionSubgraphTransactor(address common.Address, transactor bind.ContractTransactor) (*HyperionSubgraphTransactor, error) {
	contract, err := bindHyperionSubgraph(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphTransactor{contract: contract}, nil
}

// NewHyperionSubgraphFilterer creates a new log filterer instance of HyperionSubgraph, bound to a specific deployed contract.
func NewHyperionSubgraphFilterer(address common.Address, filterer bind.ContractFilterer) (*HyperionSubgraphFilterer, error) {
	contract, err := bindHyperionSubgraph(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphFilterer{contract: contract}, nil
}

// bindHyperionSubgraph binds a generic wrapper to an already deployed contract.
func bindHyperionSubgraph(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HyperionSubgraphMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperionSubgraph *HyperionSubgraphRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperionSubgraph.Contract.HyperionSubgraphCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperionSubgraph *HyperionSubgraphRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.HyperionSubgraphTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperionSubgraph *HyperionSubgraphRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.HyperionSubgraphTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HyperionSubgraph *HyperionSubgraphCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HyperionSubgraph.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HyperionSubgraph *HyperionSubgraphTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HyperionSubgraph *HyperionSubgraphTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.contract.Transact(opts, method, params...)
}

// CallData is a free data retrieval call binding the contract method 0x2b7553c7.
//
// Solidity: function callData(address _contractAddress, string _methodName, bytes _args) view returns(bytes data, bytes err)
func (_HyperionSubgraph *HyperionSubgraphCaller) CallData(opts *bind.CallOpts, _contractAddress common.Address, _methodName string, _args []byte) (struct {
	Data []byte
	Err  []byte
}, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "callData", _contractAddress, _methodName, _args)

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
func (_HyperionSubgraph *HyperionSubgraphSession) CallData(_contractAddress common.Address, _methodName string, _args []byte) (struct {
	Data []byte
	Err  []byte
}, error) {
	return _HyperionSubgraph.Contract.CallData(&_HyperionSubgraph.CallOpts, _contractAddress, _methodName, _args)
}

// CallData is a free data retrieval call binding the contract method 0x2b7553c7.
//
// Solidity: function callData(address _contractAddress, string _methodName, bytes _args) view returns(bytes data, bytes err)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) CallData(_contractAddress common.Address, _methodName string, _args []byte) (struct {
	Data []byte
	Err  []byte
}, error) {
	return _HyperionSubgraph.Contract.CallData(&_HyperionSubgraph.CallOpts, _contractAddress, _methodName, _args)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCaller) GetOwnershipExpiryTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "getOwnershipExpiryTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _HyperionSubgraph.Contract.GetOwnershipExpiryTimestamp(&_HyperionSubgraph.CallOpts)
}

// GetOwnershipExpiryTimestamp is a free data retrieval call binding the contract method 0x1ee7a108.
//
// Solidity: function getOwnershipExpiryTimestamp() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) GetOwnershipExpiryTimestamp() (*big.Int, error) {
	return _HyperionSubgraph.Contract.GetOwnershipExpiryTimestamp(&_HyperionSubgraph.CallOpts)
}

// IsHeliosNativeToken is a free data retrieval call binding the contract method 0xa6c42b02.
//
// Solidity: function isHeliosNativeToken(address ) view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphCaller) IsHeliosNativeToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "isHeliosNativeToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeliosNativeToken is a free data retrieval call binding the contract method 0xa6c42b02.
//
// Solidity: function isHeliosNativeToken(address ) view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphSession) IsHeliosNativeToken(arg0 common.Address) (bool, error) {
	return _HyperionSubgraph.Contract.IsHeliosNativeToken(&_HyperionSubgraph.CallOpts, arg0)
}

// IsHeliosNativeToken is a free data retrieval call binding the contract method 0xa6c42b02.
//
// Solidity: function isHeliosNativeToken(address ) view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) IsHeliosNativeToken(arg0 common.Address) (bool, error) {
	return _HyperionSubgraph.Contract.IsHeliosNativeToken(&_HyperionSubgraph.CallOpts, arg0)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphCaller) IsOwnershipExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "isOwnershipExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphSession) IsOwnershipExpired() (bool, error) {
	return _HyperionSubgraph.Contract.IsOwnershipExpired(&_HyperionSubgraph.CallOpts)
}

// IsOwnershipExpired is a free data retrieval call binding the contract method 0x5afe97bb.
//
// Solidity: function isOwnershipExpired() view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) IsOwnershipExpired() (bool, error) {
	return _HyperionSubgraph.Contract.IsOwnershipExpired(&_HyperionSubgraph.CallOpts)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCaller) LastBatchNonce(opts *bind.CallOpts, _erc20Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "lastBatchNonce", _erc20Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _HyperionSubgraph.Contract.LastBatchNonce(&_HyperionSubgraph.CallOpts, _erc20Address)
}

// LastBatchNonce is a free data retrieval call binding the contract method 0x011b2174.
//
// Solidity: function lastBatchNonce(address _erc20Address) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) LastBatchNonce(_erc20Address common.Address) (*big.Int, error) {
	return _HyperionSubgraph.Contract.LastBatchNonce(&_HyperionSubgraph.CallOpts, _erc20Address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperionSubgraph *HyperionSubgraphCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperionSubgraph *HyperionSubgraphSession) Owner() (common.Address, error) {
	return _HyperionSubgraph.Contract.Owner(&_HyperionSubgraph.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) Owner() (common.Address, error) {
	return _HyperionSubgraph.Contract.Owner(&_HyperionSubgraph.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphSession) Paused() (bool, error) {
	return _HyperionSubgraph.Contract.Paused(&_HyperionSubgraph.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) Paused() (bool, error) {
	return _HyperionSubgraph.Contract.Paused(&_HyperionSubgraph.CallOpts)
}

// StateHyperionId is a free data retrieval call binding the contract method 0xa4b52ca2.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_HyperionSubgraph *HyperionSubgraphCaller) StateHyperionId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "state_hyperionId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateHyperionId is a free data retrieval call binding the contract method 0xa4b52ca2.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_HyperionSubgraph *HyperionSubgraphSession) StateHyperionId() ([32]byte, error) {
	return _HyperionSubgraph.Contract.StateHyperionId(&_HyperionSubgraph.CallOpts)
}

// StateHyperionId is a free data retrieval call binding the contract method 0xa4b52ca2.
//
// Solidity: function state_hyperionId() view returns(bytes32)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) StateHyperionId() ([32]byte, error) {
	return _HyperionSubgraph.Contract.StateHyperionId(&_HyperionSubgraph.CallOpts)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCaller) StateInvalidationMapping(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "state_invalidationMapping", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphSession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateInvalidationMapping(&_HyperionSubgraph.CallOpts, arg0)
}

// StateInvalidationMapping is a free data retrieval call binding the contract method 0x7dfb6f86.
//
// Solidity: function state_invalidationMapping(bytes32 ) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) StateInvalidationMapping(arg0 [32]byte) (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateInvalidationMapping(&_HyperionSubgraph.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCaller) StateLastBatchNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "state_lastBatchNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateLastBatchNonces(&_HyperionSubgraph.CallOpts, arg0)
}

// StateLastBatchNonces is a free data retrieval call binding the contract method 0xdf97174b.
//
// Solidity: function state_lastBatchNonces(address ) view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) StateLastBatchNonces(arg0 common.Address) (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateLastBatchNonces(&_HyperionSubgraph.CallOpts, arg0)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCaller) StateLastEventNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "state_lastEventNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphSession) StateLastEventNonce() (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateLastEventNonce(&_HyperionSubgraph.CallOpts)
}

// StateLastEventNonce is a free data retrieval call binding the contract method 0x73b20547.
//
// Solidity: function state_lastEventNonce() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) StateLastEventNonce() (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateLastEventNonce(&_HyperionSubgraph.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_HyperionSubgraph *HyperionSubgraphCaller) StateLastValsetCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "state_lastValsetCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_HyperionSubgraph *HyperionSubgraphSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _HyperionSubgraph.Contract.StateLastValsetCheckpoint(&_HyperionSubgraph.CallOpts)
}

// StateLastValsetCheckpoint is a free data retrieval call binding the contract method 0xf2b53307.
//
// Solidity: function state_lastValsetCheckpoint() view returns(bytes32)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) StateLastValsetCheckpoint() ([32]byte, error) {
	return _HyperionSubgraph.Contract.StateLastValsetCheckpoint(&_HyperionSubgraph.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCaller) StateLastValsetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "state_lastValsetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphSession) StateLastValsetNonce() (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateLastValsetNonce(&_HyperionSubgraph.CallOpts)
}

// StateLastValsetNonce is a free data retrieval call binding the contract method 0xb56561fe.
//
// Solidity: function state_lastValsetNonce() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) StateLastValsetNonce() (*big.Int, error) {
	return _HyperionSubgraph.Contract.StateLastValsetNonce(&_HyperionSubgraph.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCaller) StatePowerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HyperionSubgraph.contract.Call(opts, &out, "state_powerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphSession) StatePowerThreshold() (*big.Int, error) {
	return _HyperionSubgraph.Contract.StatePowerThreshold(&_HyperionSubgraph.CallOpts)
}

// StatePowerThreshold is a free data retrieval call binding the contract method 0xe5a2b5d2.
//
// Solidity: function state_powerThreshold() view returns(uint256)
func (_HyperionSubgraph *HyperionSubgraphCallerSession) StatePowerThreshold() (*big.Int, error) {
	return _HyperionSubgraph.Contract.StatePowerThreshold(&_HyperionSubgraph.CallOpts)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) DeployERC20(opts *bind.TransactOpts, _cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "deployERC20", _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.DeployERC20(&_HyperionSubgraph.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20 is a paid mutator transaction binding the contract method 0xf7955637.
//
// Solidity: function deployERC20(string _cosmosDenom, string _name, string _symbol, uint8 _decimals) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) DeployERC20(_cosmosDenom string, _name string, _symbol string, _decimals uint8) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.DeployERC20(&_HyperionSubgraph.TransactOpts, _cosmosDenom, _name, _symbol, _decimals)
}

// DeployERC20WithSupply is a paid mutator transaction binding the contract method 0xc2d0732e.
//
// Solidity: function deployERC20WithSupply(string , string _name, string _symbol, uint8 _decimals, uint256 supply) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) DeployERC20WithSupply(opts *bind.TransactOpts, arg0 string, _name string, _symbol string, _decimals uint8, supply *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "deployERC20WithSupply", arg0, _name, _symbol, _decimals, supply)
}

// DeployERC20WithSupply is a paid mutator transaction binding the contract method 0xc2d0732e.
//
// Solidity: function deployERC20WithSupply(string , string _name, string _symbol, uint8 _decimals, uint256 supply) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) DeployERC20WithSupply(arg0 string, _name string, _symbol string, _decimals uint8, supply *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.DeployERC20WithSupply(&_HyperionSubgraph.TransactOpts, arg0, _name, _symbol, _decimals, supply)
}

// DeployERC20WithSupply is a paid mutator transaction binding the contract method 0xc2d0732e.
//
// Solidity: function deployERC20WithSupply(string , string _name, string _symbol, uint8 _decimals, uint256 supply) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) DeployERC20WithSupply(arg0 string, _name string, _symbol string, _decimals uint8, supply *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.DeployERC20WithSupply(&_HyperionSubgraph.TransactOpts, arg0, _name, _symbol, _decimals, supply)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) EmergencyPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "emergencyPause")
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_HyperionSubgraph *HyperionSubgraphSession) EmergencyPause() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.EmergencyPause(&_HyperionSubgraph.TransactOpts)
}

// EmergencyPause is a paid mutator transaction binding the contract method 0x51858e27.
//
// Solidity: function emergencyPause() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) EmergencyPause() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.EmergencyPause(&_HyperionSubgraph.TransactOpts)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) EmergencyUnpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "emergencyUnpause")
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_HyperionSubgraph *HyperionSubgraphSession) EmergencyUnpause() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.EmergencyUnpause(&_HyperionSubgraph.TransactOpts)
}

// EmergencyUnpause is a paid mutator transaction binding the contract method 0x4a4e3bd5.
//
// Solidity: function emergencyUnpause() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) EmergencyUnpause() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.EmergencyUnpause(&_HyperionSubgraph.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc359a212.
//
// Solidity: function initialize(bytes32 _hyperionId, uint256 _powerThreshold, address[] _validators, uint256[] _powers) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) Initialize(opts *bind.TransactOpts, _hyperionId [32]byte, _powerThreshold *big.Int, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "initialize", _hyperionId, _powerThreshold, _validators, _powers)
}

// Initialize is a paid mutator transaction binding the contract method 0xc359a212.
//
// Solidity: function initialize(bytes32 _hyperionId, uint256 _powerThreshold, address[] _validators, uint256[] _powers) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) Initialize(_hyperionId [32]byte, _powerThreshold *big.Int, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.Initialize(&_HyperionSubgraph.TransactOpts, _hyperionId, _powerThreshold, _validators, _powers)
}

// Initialize is a paid mutator transaction binding the contract method 0xc359a212.
//
// Solidity: function initialize(bytes32 _hyperionId, uint256 _powerThreshold, address[] _validators, uint256[] _powers) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) Initialize(_hyperionId [32]byte, _powerThreshold *big.Int, _validators []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.Initialize(&_HyperionSubgraph.TransactOpts, _hyperionId, _powerThreshold, _validators, _powers)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HyperionSubgraph *HyperionSubgraphSession) RenounceOwnership() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.RenounceOwnership(&_HyperionSubgraph.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.RenounceOwnership(&_HyperionSubgraph.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) RenounceOwnershipAfterExpiry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "renounceOwnershipAfterExpiry")
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_HyperionSubgraph *HyperionSubgraphSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.RenounceOwnershipAfterExpiry(&_HyperionSubgraph.TransactOpts)
}

// RenounceOwnershipAfterExpiry is a paid mutator transaction binding the contract method 0x8c64865f.
//
// Solidity: function renounceOwnershipAfterExpiry() returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) RenounceOwnershipAfterExpiry() (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.RenounceOwnershipAfterExpiry(&_HyperionSubgraph.TransactOpts)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x1ffbe7f9.
//
// Solidity: function sendToCosmos(address _tokenContract, bytes32 _destination, uint256 _amount) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) SendToCosmos(opts *bind.TransactOpts, _tokenContract common.Address, _destination [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "sendToCosmos", _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x1ffbe7f9.
//
// Solidity: function sendToCosmos(address _tokenContract, bytes32 _destination, uint256 _amount) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) SendToCosmos(_tokenContract common.Address, _destination [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.SendToCosmos(&_HyperionSubgraph.TransactOpts, _tokenContract, _destination, _amount)
}

// SendToCosmos is a paid mutator transaction binding the contract method 0x1ffbe7f9.
//
// Solidity: function sendToCosmos(address _tokenContract, bytes32 _destination, uint256 _amount) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) SendToCosmos(_tokenContract common.Address, _destination [32]byte, _amount *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.SendToCosmos(&_HyperionSubgraph.TransactOpts, _tokenContract, _destination, _amount)
}

// SendToHelios is a paid mutator transaction binding the contract method 0x308ff208.
//
// Solidity: function sendToHelios(address _tokenContract, bytes32 _destination, uint256 _amount, string _data) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) SendToHelios(opts *bind.TransactOpts, _tokenContract common.Address, _destination [32]byte, _amount *big.Int, _data string) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "sendToHelios", _tokenContract, _destination, _amount, _data)
}

// SendToHelios is a paid mutator transaction binding the contract method 0x308ff208.
//
// Solidity: function sendToHelios(address _tokenContract, bytes32 _destination, uint256 _amount, string _data) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) SendToHelios(_tokenContract common.Address, _destination [32]byte, _amount *big.Int, _data string) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.SendToHelios(&_HyperionSubgraph.TransactOpts, _tokenContract, _destination, _amount, _data)
}

// SendToHelios is a paid mutator transaction binding the contract method 0x308ff208.
//
// Solidity: function sendToHelios(address _tokenContract, bytes32 _destination, uint256 _amount, string _data) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) SendToHelios(_tokenContract common.Address, _destination [32]byte, _amount *big.Int, _data string) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.SendToHelios(&_HyperionSubgraph.TransactOpts, _tokenContract, _destination, _amount, _data)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x81747418.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) SubmitBatch(opts *bind.TransactOpts, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "submitBatch", _currentValset, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x81747418.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) SubmitBatch(_currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.SubmitBatch(&_HyperionSubgraph.TransactOpts, _currentValset, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// SubmitBatch is a paid mutator transaction binding the contract method 0x81747418.
//
// Solidity: function submitBatch((address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s, uint256[] _amounts, address[] _destinations, uint256[] _fees, uint256 _batchNonce, address _tokenContract, uint256 _batchTimeout) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) SubmitBatch(_currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte, _amounts []*big.Int, _destinations []common.Address, _fees []*big.Int, _batchNonce *big.Int, _tokenContract common.Address, _batchTimeout *big.Int) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.SubmitBatch(&_HyperionSubgraph.TransactOpts, _currentValset, _v, _r, _s, _amounts, _destinations, _fees, _batchNonce, _tokenContract, _batchTimeout)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.TransferOwnership(&_HyperionSubgraph.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.TransferOwnership(&_HyperionSubgraph.TransactOpts, newOwner)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xa5352f5b.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactor) UpdateValset(opts *bind.TransactOpts, _newValset ValsetArgs, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _HyperionSubgraph.contract.Transact(opts, "updateValset", _newValset, _currentValset, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xa5352f5b.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_HyperionSubgraph *HyperionSubgraphSession) UpdateValset(_newValset ValsetArgs, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.UpdateValset(&_HyperionSubgraph.TransactOpts, _newValset, _currentValset, _v, _r, _s)
}

// UpdateValset is a paid mutator transaction binding the contract method 0xa5352f5b.
//
// Solidity: function updateValset((address[],uint256[],uint256,uint256,address) _newValset, (address[],uint256[],uint256,uint256,address) _currentValset, uint8[] _v, bytes32[] _r, bytes32[] _s) returns()
func (_HyperionSubgraph *HyperionSubgraphTransactorSession) UpdateValset(_newValset ValsetArgs, _currentValset ValsetArgs, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _HyperionSubgraph.Contract.UpdateValset(&_HyperionSubgraph.TransactOpts, _newValset, _currentValset, _v, _r, _s)
}

// HyperionSubgraphERC20DeployedEventIterator is returned from FilterERC20DeployedEvent and is used to iterate over the raw logs and unpacked data for ERC20DeployedEvent events raised by the HyperionSubgraph contract.
type HyperionSubgraphERC20DeployedEventIterator struct {
	Event *HyperionSubgraphERC20DeployedEvent // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphERC20DeployedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphERC20DeployedEvent)
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
		it.Event = new(HyperionSubgraphERC20DeployedEvent)
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
func (it *HyperionSubgraphERC20DeployedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphERC20DeployedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphERC20DeployedEvent represents a ERC20DeployedEvent event raised by the HyperionSubgraph contract.
type HyperionSubgraphERC20DeployedEvent struct {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterERC20DeployedEvent(opts *bind.FilterOpts, _tokenContract []common.Address) (*HyperionSubgraphERC20DeployedEventIterator, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphERC20DeployedEventIterator{contract: _HyperionSubgraph.contract, event: "ERC20DeployedEvent", logs: logs, sub: sub}, nil
}

// WatchERC20DeployedEvent is a free log subscription operation binding the contract event 0x82fe3a4fa49c6382d0c085746698ddbbafe6c2bf61285b19410644b5b26287c7.
//
// Solidity: event ERC20DeployedEvent(string _cosmosDenom, address indexed _tokenContract, string _name, string _symbol, uint8 _decimals, uint256 _eventNonce)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchERC20DeployedEvent(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphERC20DeployedEvent, _tokenContract []common.Address) (event.Subscription, error) {

	var _tokenContractRule []interface{}
	for _, _tokenContractItem := range _tokenContract {
		_tokenContractRule = append(_tokenContractRule, _tokenContractItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "ERC20DeployedEvent", _tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphERC20DeployedEvent)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParseERC20DeployedEvent(log types.Log) (*HyperionSubgraphERC20DeployedEvent, error) {
	event := new(HyperionSubgraphERC20DeployedEvent)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "ERC20DeployedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSubgraphOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HyperionSubgraph contract.
type HyperionSubgraphOwnershipTransferredIterator struct {
	Event *HyperionSubgraphOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphOwnershipTransferred)
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
		it.Event = new(HyperionSubgraphOwnershipTransferred)
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
func (it *HyperionSubgraphOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphOwnershipTransferred represents a OwnershipTransferred event raised by the HyperionSubgraph contract.
type HyperionSubgraphOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HyperionSubgraphOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphOwnershipTransferredIterator{contract: _HyperionSubgraph.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphOwnershipTransferred)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParseOwnershipTransferred(log types.Log) (*HyperionSubgraphOwnershipTransferred, error) {
	event := new(HyperionSubgraphOwnershipTransferred)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSubgraphPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the HyperionSubgraph contract.
type HyperionSubgraphPausedIterator struct {
	Event *HyperionSubgraphPaused // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphPaused)
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
		it.Event = new(HyperionSubgraphPaused)
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
func (it *HyperionSubgraphPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphPaused represents a Paused event raised by the HyperionSubgraph contract.
type HyperionSubgraphPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterPaused(opts *bind.FilterOpts) (*HyperionSubgraphPausedIterator, error) {

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphPausedIterator{contract: _HyperionSubgraph.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphPaused) (event.Subscription, error) {

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphPaused)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParsePaused(log types.Log) (*HyperionSubgraphPaused, error) {
	event := new(HyperionSubgraphPaused)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSubgraphSendToCosmosEventIterator is returned from FilterSendToCosmosEvent and is used to iterate over the raw logs and unpacked data for SendToCosmosEvent events raised by the HyperionSubgraph contract.
type HyperionSubgraphSendToCosmosEventIterator struct {
	Event *HyperionSubgraphSendToCosmosEvent // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphSendToCosmosEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphSendToCosmosEvent)
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
		it.Event = new(HyperionSubgraphSendToCosmosEvent)
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
func (it *HyperionSubgraphSendToCosmosEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphSendToCosmosEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphSendToCosmosEvent represents a SendToCosmosEvent event raised by the HyperionSubgraph contract.
type HyperionSubgraphSendToCosmosEvent struct {
	TokenContract common.Address
	Sender        common.Address
	Destination   [32]byte
	Amount        *big.Int
	EventNonce    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSendToCosmosEvent is a free log retrieval operation binding the contract event 0xd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e4.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce)
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterSendToCosmosEvent(opts *bind.FilterOpts, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (*HyperionSubgraphSendToCosmosEventIterator, error) {

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

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphSendToCosmosEventIterator{contract: _HyperionSubgraph.contract, event: "SendToCosmosEvent", logs: logs, sub: sub}, nil
}

// WatchSendToCosmosEvent is a free log subscription operation binding the contract event 0xd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e4.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchSendToCosmosEvent(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphSendToCosmosEvent, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "SendToCosmosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphSendToCosmosEvent)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
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

// ParseSendToCosmosEvent is a log parse operation binding the contract event 0xd7767894d73c589daeca9643f445f03d7be61aad2950c117e7cbff4176fca7e4.
//
// Solidity: event SendToCosmosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce)
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParseSendToCosmosEvent(log types.Log) (*HyperionSubgraphSendToCosmosEvent, error) {
	event := new(HyperionSubgraphSendToCosmosEvent)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "SendToCosmosEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSubgraphSendToHeliosEventIterator is returned from FilterSendToHeliosEvent and is used to iterate over the raw logs and unpacked data for SendToHeliosEvent events raised by the HyperionSubgraph contract.
type HyperionSubgraphSendToHeliosEventIterator struct {
	Event *HyperionSubgraphSendToHeliosEvent // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphSendToHeliosEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphSendToHeliosEvent)
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
		it.Event = new(HyperionSubgraphSendToHeliosEvent)
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
func (it *HyperionSubgraphSendToHeliosEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphSendToHeliosEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphSendToHeliosEvent represents a SendToHeliosEvent event raised by the HyperionSubgraph contract.
type HyperionSubgraphSendToHeliosEvent struct {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterSendToHeliosEvent(opts *bind.FilterOpts, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (*HyperionSubgraphSendToHeliosEventIterator, error) {

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

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "SendToHeliosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphSendToHeliosEventIterator{contract: _HyperionSubgraph.contract, event: "SendToHeliosEvent", logs: logs, sub: sub}, nil
}

// WatchSendToHeliosEvent is a free log subscription operation binding the contract event 0x272cb0695a9182efb214ae0bc3e2c8163469c94b2cef2471499f6237d4ca8b47.
//
// Solidity: event SendToHeliosEvent(address indexed _tokenContract, address indexed _sender, bytes32 indexed _destination, uint256 _amount, uint256 _eventNonce, string _data)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchSendToHeliosEvent(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphSendToHeliosEvent, _tokenContract []common.Address, _sender []common.Address, _destination [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "SendToHeliosEvent", _tokenContractRule, _senderRule, _destinationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphSendToHeliosEvent)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "SendToHeliosEvent", log); err != nil {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParseSendToHeliosEvent(log types.Log) (*HyperionSubgraphSendToHeliosEvent, error) {
	event := new(HyperionSubgraphSendToHeliosEvent)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "SendToHeliosEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSubgraphTransactionBatchExecutedEventIterator is returned from FilterTransactionBatchExecutedEvent and is used to iterate over the raw logs and unpacked data for TransactionBatchExecutedEvent events raised by the HyperionSubgraph contract.
type HyperionSubgraphTransactionBatchExecutedEventIterator struct {
	Event *HyperionSubgraphTransactionBatchExecutedEvent // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphTransactionBatchExecutedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphTransactionBatchExecutedEvent)
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
		it.Event = new(HyperionSubgraphTransactionBatchExecutedEvent)
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
func (it *HyperionSubgraphTransactionBatchExecutedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphTransactionBatchExecutedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphTransactionBatchExecutedEvent represents a TransactionBatchExecutedEvent event raised by the HyperionSubgraph contract.
type HyperionSubgraphTransactionBatchExecutedEvent struct {
	BatchNonce *big.Int
	Token      common.Address
	EventNonce *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionBatchExecutedEvent is a free log retrieval operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterTransactionBatchExecutedEvent(opts *bind.FilterOpts, _batchNonce []*big.Int, _token []common.Address) (*HyperionSubgraphTransactionBatchExecutedEventIterator, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphTransactionBatchExecutedEventIterator{contract: _HyperionSubgraph.contract, event: "TransactionBatchExecutedEvent", logs: logs, sub: sub}, nil
}

// WatchTransactionBatchExecutedEvent is a free log subscription operation binding the contract event 0x02c7e81975f8edb86e2a0c038b7b86a49c744236abf0f6177ff5afc6986ab708.
//
// Solidity: event TransactionBatchExecutedEvent(uint256 indexed _batchNonce, address indexed _token, uint256 _eventNonce)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchTransactionBatchExecutedEvent(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphTransactionBatchExecutedEvent, _batchNonce []*big.Int, _token []common.Address) (event.Subscription, error) {

	var _batchNonceRule []interface{}
	for _, _batchNonceItem := range _batchNonce {
		_batchNonceRule = append(_batchNonceRule, _batchNonceItem)
	}
	var _tokenRule []interface{}
	for _, _tokenItem := range _token {
		_tokenRule = append(_tokenRule, _tokenItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "TransactionBatchExecutedEvent", _batchNonceRule, _tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphTransactionBatchExecutedEvent)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParseTransactionBatchExecutedEvent(log types.Log) (*HyperionSubgraphTransactionBatchExecutedEvent, error) {
	event := new(HyperionSubgraphTransactionBatchExecutedEvent)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "TransactionBatchExecutedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSubgraphUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the HyperionSubgraph contract.
type HyperionSubgraphUnpausedIterator struct {
	Event *HyperionSubgraphUnpaused // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphUnpaused)
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
		it.Event = new(HyperionSubgraphUnpaused)
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
func (it *HyperionSubgraphUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphUnpaused represents a Unpaused event raised by the HyperionSubgraph contract.
type HyperionSubgraphUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterUnpaused(opts *bind.FilterOpts) (*HyperionSubgraphUnpausedIterator, error) {

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphUnpausedIterator{contract: _HyperionSubgraph.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphUnpaused) (event.Subscription, error) {

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphUnpaused)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParseUnpaused(log types.Log) (*HyperionSubgraphUnpaused, error) {
	event := new(HyperionSubgraphUnpaused)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HyperionSubgraphValsetUpdatedEventIterator is returned from FilterValsetUpdatedEvent and is used to iterate over the raw logs and unpacked data for ValsetUpdatedEvent events raised by the HyperionSubgraph contract.
type HyperionSubgraphValsetUpdatedEventIterator struct {
	Event *HyperionSubgraphValsetUpdatedEvent // Event containing the contract specifics and raw log

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
func (it *HyperionSubgraphValsetUpdatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HyperionSubgraphValsetUpdatedEvent)
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
		it.Event = new(HyperionSubgraphValsetUpdatedEvent)
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
func (it *HyperionSubgraphValsetUpdatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HyperionSubgraphValsetUpdatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HyperionSubgraphValsetUpdatedEvent represents a ValsetUpdatedEvent event raised by the HyperionSubgraph contract.
type HyperionSubgraphValsetUpdatedEvent struct {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) FilterValsetUpdatedEvent(opts *bind.FilterOpts, _newValsetNonce []*big.Int) (*HyperionSubgraphValsetUpdatedEventIterator, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.FilterLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return &HyperionSubgraphValsetUpdatedEventIterator{contract: _HyperionSubgraph.contract, event: "ValsetUpdatedEvent", logs: logs, sub: sub}, nil
}

// WatchValsetUpdatedEvent is a free log subscription operation binding the contract event 0x76d08978c024a4bf8cbb30c67fd78fcaa1827cbc533e4e175f36d07e64ccf96a.
//
// Solidity: event ValsetUpdatedEvent(uint256 indexed _newValsetNonce, uint256 _eventNonce, uint256 _rewardAmount, address _rewardToken, address[] _validators, uint256[] _powers)
func (_HyperionSubgraph *HyperionSubgraphFilterer) WatchValsetUpdatedEvent(opts *bind.WatchOpts, sink chan<- *HyperionSubgraphValsetUpdatedEvent, _newValsetNonce []*big.Int) (event.Subscription, error) {

	var _newValsetNonceRule []interface{}
	for _, _newValsetNonceItem := range _newValsetNonce {
		_newValsetNonceRule = append(_newValsetNonceRule, _newValsetNonceItem)
	}

	logs, sub, err := _HyperionSubgraph.contract.WatchLogs(opts, "ValsetUpdatedEvent", _newValsetNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HyperionSubgraphValsetUpdatedEvent)
				if err := _HyperionSubgraph.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
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
func (_HyperionSubgraph *HyperionSubgraphFilterer) ParseValsetUpdatedEvent(log types.Log) (*HyperionSubgraphValsetUpdatedEvent, error) {
	event := new(HyperionSubgraphValsetUpdatedEvent)
	if err := _HyperionSubgraph.contract.UnpackLog(event, "ValsetUpdatedEvent", log); err != nil {
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
