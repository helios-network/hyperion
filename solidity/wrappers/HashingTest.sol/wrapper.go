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

// HashingTestMetaData contains all meta data concerning the HashingTest contract.
var HashingTestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"}],\"name\":\"ConcatHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"}],\"name\":\"ConcatHash2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_hyperionId\",\"type\":\"bytes32\"}],\"name\":\"IterativeHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"}],\"name\":\"JustSaveEverything\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_valsetNonce\",\"type\":\"uint256\"}],\"name\":\"JustSaveEverythingAgain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastCheckpoint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state_nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_powers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"state_validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608080604052346015576107da908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80630caff28b146103b45780632afbb62e146103505780632b939281146102e25780636071cbd914610223578063715dff7e146100b957806374df6ae4146100be578063884403e2146100b9578063ccf0e74c1461009c5763d32e81a51461007c575f80fd5b34610098575f3660031901126100985760205f54604051908152f35b5f80fd5b34610098575f366003190112610098576020600354604051908152f35b610588565b34610098576100cc36610531565b9060405190602082019283526918da1958dadc1bda5b9d60b21b60408301526060820152606081526100ff608082610418565b519020905f915b835183101561021f5782610181575b6001600160a01b036101278486610790565b51166101338484610790565b5160405191602083019384526040830152606082015260608152610158608082610418565b519020916001810180911161016d5791610106565b634e487b7160e01b5f52601160045260245ffd5b61018b8383610790565b515f19840184811161016d576101a19084610790565b5110156101155760405162461bcd60e51b815260206004820152604360248201527f56616c696461746f7220706f776572206d757374206e6f74206265206869676860448201527f6572207468616e2070726576696f75732076616c696461746f7220696e2062616064820152620e8c6d60eb1b608482015260a490fd5b5f55005b346100985761023136610531565b9060405190602082019283526918da1958dadc1bda5b9d60b21b6040830152606082015260608152610264608082610418565b5190209160405161029381610285602082019460208652604083019061070b565b03601f198101835282610418565b519020906040516102b4816102856020820194602086526040830190610747565b519020604051916020830193845260408301526060820152606081526102db608082610418565b5190205f55005b34610098576020366003190112610098576004356002548110156100985760025481101561033c5760025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0154604051908152602090f35b634e487b7160e01b5f52603260045260245ffd5b34610098576020366003190112610098576004356001548110156100985760015481101561033c5760015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601546040516001600160a01b039091168152602090f35b34610098576102db6102856103c836610531565b9492939091610406604051958694602086019889526918da1958dadc1bda5b9d60b21b6040870152606086015260a0608086015260c085019061070b565b838103601f190160a085015290610747565b90601f8019910116810190811067ffffffffffffffff82111761043a57604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff811161043a5760051b60200190565b9080601f830112156100985781359061047e8261044e565b9261048c6040519485610418565b82845260208085019360051b82010191821161009857602001915b8183106104b45750505090565b82356001600160a01b0381168103610098578152602092830192016104a7565b9080601f830112156100985781356104eb8161044e565b926104f96040519485610418565b81845260208085019260051b82010192831161009857602001905b8282106105215750505090565b8135815260209182019101610514565b9060806003198301126100985760043567ffffffffffffffff8111610098578261055d91600401610466565b916024359067ffffffffffffffff82116100985761057d916004016104d4565b906044359060643590565b346100985760603660031901126100985760043567ffffffffffffffff8111610098576105b9903690600401610466565b60243567ffffffffffffffff8111610098576105d99036906004016104d4565b81516044359267ffffffffffffffff821161043a57600160401b821161043a57602090600154836001558084106106ef575b500160015f525f5b8281106106b25750505080519067ffffffffffffffff821161043a57600160401b821161043a5760209060025483600255808410610696575b500160025f525f5b828110610662576003849055005b60019060208351930192817f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace015501610654565b6106ac9060025f5284845f20918201910161077a565b5f61064c565b81516001600160a01b03167fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6820155602090910190600101610613565b6107059060015f5284845f20918201910161077a565b5f61060b565b90602080835192838152019201905f5b8181106107285750505090565b82516001600160a01b031684526020938401939092019160010161071b565b90602080835192838152019201905f5b8181106107645750505090565b8251845260209384019390920191600101610757565b818110610785575050565b5f815560010161077a565b805182101561033c5760209160051b01019056fea2646970667358221220c6409085b0941c71856d5a00e0e7e59c5f049822a61bf835cb5f15f514e8e3a764736f6c634300081c0033",
}

// HashingTestABI is the input ABI used to generate the binding from.
// Deprecated: Use HashingTestMetaData.ABI instead.
var HashingTestABI = HashingTestMetaData.ABI

// HashingTestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashingTestMetaData.Bin instead.
var HashingTestBin = HashingTestMetaData.Bin

// DeployHashingTest deploys a new Ethereum contract, binding an instance of HashingTest to it.
func DeployHashingTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashingTest, error) {
	parsed, err := HashingTestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashingTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashingTest{HashingTestCaller: HashingTestCaller{contract: contract}, HashingTestTransactor: HashingTestTransactor{contract: contract}, HashingTestFilterer: HashingTestFilterer{contract: contract}}, nil
}

// HashingTest is an auto generated Go binding around an Ethereum contract.
type HashingTest struct {
	HashingTestCaller     // Read-only binding to the contract
	HashingTestTransactor // Write-only binding to the contract
	HashingTestFilterer   // Log filterer for contract events
}

// HashingTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashingTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashingTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashingTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashingTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashingTestSession struct {
	Contract     *HashingTest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashingTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashingTestCallerSession struct {
	Contract *HashingTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HashingTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashingTestTransactorSession struct {
	Contract     *HashingTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HashingTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashingTestRaw struct {
	Contract *HashingTest // Generic contract binding to access the raw methods on
}

// HashingTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashingTestCallerRaw struct {
	Contract *HashingTestCaller // Generic read-only contract binding to access the raw methods on
}

// HashingTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashingTestTransactorRaw struct {
	Contract *HashingTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashingTest creates a new instance of HashingTest, bound to a specific deployed contract.
func NewHashingTest(address common.Address, backend bind.ContractBackend) (*HashingTest, error) {
	contract, err := bindHashingTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashingTest{HashingTestCaller: HashingTestCaller{contract: contract}, HashingTestTransactor: HashingTestTransactor{contract: contract}, HashingTestFilterer: HashingTestFilterer{contract: contract}}, nil
}

// NewHashingTestCaller creates a new read-only instance of HashingTest, bound to a specific deployed contract.
func NewHashingTestCaller(address common.Address, caller bind.ContractCaller) (*HashingTestCaller, error) {
	contract, err := bindHashingTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashingTestCaller{contract: contract}, nil
}

// NewHashingTestTransactor creates a new write-only instance of HashingTest, bound to a specific deployed contract.
func NewHashingTestTransactor(address common.Address, transactor bind.ContractTransactor) (*HashingTestTransactor, error) {
	contract, err := bindHashingTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashingTestTransactor{contract: contract}, nil
}

// NewHashingTestFilterer creates a new log filterer instance of HashingTest, bound to a specific deployed contract.
func NewHashingTestFilterer(address common.Address, filterer bind.ContractFilterer) (*HashingTestFilterer, error) {
	contract, err := bindHashingTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashingTestFilterer{contract: contract}, nil
}

// bindHashingTest binds a generic wrapper to an already deployed contract.
func bindHashingTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HashingTestMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashingTest *HashingTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashingTest.Contract.HashingTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashingTest *HashingTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashingTest.Contract.HashingTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashingTest *HashingTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashingTest.Contract.HashingTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashingTest *HashingTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashingTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashingTest *HashingTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashingTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashingTest *HashingTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashingTest.Contract.contract.Transact(opts, method, params...)
}

// LastCheckpoint is a free data retrieval call binding the contract method 0xd32e81a5.
//
// Solidity: function lastCheckpoint() view returns(bytes32)
func (_HashingTest *HashingTestCaller) LastCheckpoint(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "lastCheckpoint")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LastCheckpoint is a free data retrieval call binding the contract method 0xd32e81a5.
//
// Solidity: function lastCheckpoint() view returns(bytes32)
func (_HashingTest *HashingTestSession) LastCheckpoint() ([32]byte, error) {
	return _HashingTest.Contract.LastCheckpoint(&_HashingTest.CallOpts)
}

// LastCheckpoint is a free data retrieval call binding the contract method 0xd32e81a5.
//
// Solidity: function lastCheckpoint() view returns(bytes32)
func (_HashingTest *HashingTestCallerSession) LastCheckpoint() ([32]byte, error) {
	return _HashingTest.Contract.LastCheckpoint(&_HashingTest.CallOpts)
}

// StateNonce is a free data retrieval call binding the contract method 0xccf0e74c.
//
// Solidity: function state_nonce() view returns(uint256)
func (_HashingTest *HashingTestCaller) StateNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "state_nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StateNonce is a free data retrieval call binding the contract method 0xccf0e74c.
//
// Solidity: function state_nonce() view returns(uint256)
func (_HashingTest *HashingTestSession) StateNonce() (*big.Int, error) {
	return _HashingTest.Contract.StateNonce(&_HashingTest.CallOpts)
}

// StateNonce is a free data retrieval call binding the contract method 0xccf0e74c.
//
// Solidity: function state_nonce() view returns(uint256)
func (_HashingTest *HashingTestCallerSession) StateNonce() (*big.Int, error) {
	return _HashingTest.Contract.StateNonce(&_HashingTest.CallOpts)
}

// StatePowers is a free data retrieval call binding the contract method 0x2b939281.
//
// Solidity: function state_powers(uint256 ) view returns(uint256)
func (_HashingTest *HashingTestCaller) StatePowers(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "state_powers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatePowers is a free data retrieval call binding the contract method 0x2b939281.
//
// Solidity: function state_powers(uint256 ) view returns(uint256)
func (_HashingTest *HashingTestSession) StatePowers(arg0 *big.Int) (*big.Int, error) {
	return _HashingTest.Contract.StatePowers(&_HashingTest.CallOpts, arg0)
}

// StatePowers is a free data retrieval call binding the contract method 0x2b939281.
//
// Solidity: function state_powers(uint256 ) view returns(uint256)
func (_HashingTest *HashingTestCallerSession) StatePowers(arg0 *big.Int) (*big.Int, error) {
	return _HashingTest.Contract.StatePowers(&_HashingTest.CallOpts, arg0)
}

// StateValidators is a free data retrieval call binding the contract method 0x2afbb62e.
//
// Solidity: function state_validators(uint256 ) view returns(address)
func (_HashingTest *HashingTestCaller) StateValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _HashingTest.contract.Call(opts, &out, "state_validators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StateValidators is a free data retrieval call binding the contract method 0x2afbb62e.
//
// Solidity: function state_validators(uint256 ) view returns(address)
func (_HashingTest *HashingTestSession) StateValidators(arg0 *big.Int) (common.Address, error) {
	return _HashingTest.Contract.StateValidators(&_HashingTest.CallOpts, arg0)
}

// StateValidators is a free data retrieval call binding the contract method 0x2afbb62e.
//
// Solidity: function state_validators(uint256 ) view returns(address)
func (_HashingTest *HashingTestCallerSession) StateValidators(arg0 *big.Int) (common.Address, error) {
	return _HashingTest.Contract.StateValidators(&_HashingTest.CallOpts, arg0)
}

// ConcatHash is a paid mutator transaction binding the contract method 0x6071cbd9.
//
// Solidity: function ConcatHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestTransactor) ConcatHash(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "ConcatHash", _validators, _powers, _valsetNonce, _hyperionId)
}

// ConcatHash is a paid mutator transaction binding the contract method 0x6071cbd9.
//
// Solidity: function ConcatHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestSession) ConcatHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// ConcatHash is a paid mutator transaction binding the contract method 0x6071cbd9.
//
// Solidity: function ConcatHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestTransactorSession) ConcatHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// ConcatHash2 is a paid mutator transaction binding the contract method 0x0caff28b.
//
// Solidity: function ConcatHash2(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestTransactor) ConcatHash2(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "ConcatHash2", _validators, _powers, _valsetNonce, _hyperionId)
}

// ConcatHash2 is a paid mutator transaction binding the contract method 0x0caff28b.
//
// Solidity: function ConcatHash2(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestSession) ConcatHash2(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash2(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// ConcatHash2 is a paid mutator transaction binding the contract method 0x0caff28b.
//
// Solidity: function ConcatHash2(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestTransactorSession) ConcatHash2(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.ConcatHash2(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// IterativeHash is a paid mutator transaction binding the contract method 0x74df6ae4.
//
// Solidity: function IterativeHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestTransactor) IterativeHash(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "IterativeHash", _validators, _powers, _valsetNonce, _hyperionId)
}

// IterativeHash is a paid mutator transaction binding the contract method 0x74df6ae4.
//
// Solidity: function IterativeHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestSession) IterativeHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.IterativeHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// IterativeHash is a paid mutator transaction binding the contract method 0x74df6ae4.
//
// Solidity: function IterativeHash(address[] _validators, uint256[] _powers, uint256 _valsetNonce, bytes32 _hyperionId) returns()
func (_HashingTest *HashingTestTransactorSession) IterativeHash(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int, _hyperionId [32]byte) (*types.Transaction, error) {
	return _HashingTest.Contract.IterativeHash(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce, _hyperionId)
}

// JustSaveEverything is a paid mutator transaction binding the contract method 0x884403e2.
//
// Solidity: function JustSaveEverything(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactor) JustSaveEverything(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "JustSaveEverything", _validators, _powers, _valsetNonce)
}

// JustSaveEverything is a paid mutator transaction binding the contract method 0x884403e2.
//
// Solidity: function JustSaveEverything(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestSession) JustSaveEverything(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverything(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}

// JustSaveEverything is a paid mutator transaction binding the contract method 0x884403e2.
//
// Solidity: function JustSaveEverything(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactorSession) JustSaveEverything(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverything(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}

// JustSaveEverythingAgain is a paid mutator transaction binding the contract method 0x715dff7e.
//
// Solidity: function JustSaveEverythingAgain(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactor) JustSaveEverythingAgain(opts *bind.TransactOpts, _validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.contract.Transact(opts, "JustSaveEverythingAgain", _validators, _powers, _valsetNonce)
}

// JustSaveEverythingAgain is a paid mutator transaction binding the contract method 0x715dff7e.
//
// Solidity: function JustSaveEverythingAgain(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestSession) JustSaveEverythingAgain(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverythingAgain(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}

// JustSaveEverythingAgain is a paid mutator transaction binding the contract method 0x715dff7e.
//
// Solidity: function JustSaveEverythingAgain(address[] _validators, uint256[] _powers, uint256 _valsetNonce) returns()
func (_HashingTest *HashingTestTransactorSession) JustSaveEverythingAgain(_validators []common.Address, _powers []*big.Int, _valsetNonce *big.Int) (*types.Transaction, error) {
	return _HashingTest.Contract.JustSaveEverythingAgain(&_HashingTest.TransactOpts, _validators, _powers, _valsetNonce)
}
