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
	Bin: "0x60808060405234601557610840908161001a8239f35b5f80fdfe604060808152600480361015610013575f80fd5b5f3560e01c9081630caff28b146103a85781632afbb62e1461034b5781632b939281146102fb5781636071cbd914610237578163715dff7e146100be57816374df6ae4146100c357508063884403e2146100be578063ccf0e74c146100a05763d32e81a514610080575f80fd5b3461009c575f36600319011261009c576020905f549051908152f35b5f80fd5b503461009c575f36600319011261009c576020906003549051908152f35b610597565b90503461009c576100d336610548565b919294805192602093602081019182526918da1958dadc1bda5b9d60b21b8382015260609360608201526060815261010a8161040e565b519020935f945b87518610156102335785610184575b6001600160a01b03610132878a6107e2565b511661013e87846107e2565b51845191878301938452858301528582015284815261015c8161040e565b51902094600181018091116101715794610111565b601187634e487b7160e01b5f525260245ffd5b61018e86836107e2565b515f198701878111610220576101a490846107e2565b51101561012057825162461bcd60e51b8152808801869052604360248201527f56616c696461746f7220706f776572206d757374206e6f74206265206869676860448201527f6572207468616e2070726576696f75732076616c696461746f7220696e2062616064820152620e8c6d60eb1b608482015260a490fd5b601189634e487b7160e01b5f525260245ffd5b5f55005b823461009c5761024636610548565b849294939193516020948582019283526918da1958dadc1bda5b9d60b21b8583015260608201526060815261027a8161040e565b5190209382519182610295868201928784528683019061075d565b03926102a9601f199485810183528261043e565b519020916102d3845191826102c78882019589875288830190610799565b0390810183528261043e565b5190209180519384019485528301526060820152606081526102f48161040e565b5190205f55005b90503461009c57602036600319011261009c57359060025482101561009c5760209160025f527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace01549051908152f35b90503461009c57602036600319011261009c57359060015482101561009c5760015f527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf69091015490516001600160a01b03919091168152602090f35b823461009c576102c76102f46103fb6103c036610548565b96928396929491965196879460208601998a526918da1958dadc1bda5b9d60b21b90860152606085015260a0608085015260c084019061075d565b601f1993848483030160a0850152610799565b6080810190811067ffffffffffffffff82111761042a57604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761042a57604052565b67ffffffffffffffff811161042a5760051b60200190565b9080601f8301121561009c57602090823561049281610460565b936104a0604051958661043e565b81855260208086019260051b82010192831161009c57602001905b8282106104c9575050505090565b81356001600160a01b038116810361009c5781529083019083016104bb565b9080601f8301121561009c57602090823561050281610460565b93610510604051958661043e565b81855260208086019260051b82010192831161009c57602001905b828210610539575050505090565b8135815290830190830161052b565b90608060031983011261009c5767ffffffffffffffff60043581811161009c578361057591600401610478565b9260243591821161009c5761058c916004016104e8565b906044359060643590565b3461009c57606036600319011261009c57600467ffffffffffffffff813581811161009c576105c99036908401610478565b60243582811161009c576105e090369085016104e8565b8151938385116106c657680100000000000000008086116106d95760019560015481600155808210610727575b50602080950160015f52875f5b8381106106ec575050505082519485116106d95784116106c6575060209060025484600255808510610690575b50019060025f525f5b83811061065f57604435600355005b82517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace820155918101918401610650565b6106c090857f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace91820191016107cc565b5f610647565b604190634e487b7160e01b5f525260245ffd5b604182634e487b7160e01b5f525260245ffd5b82516001600160a01b03167fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf68201559187019189910161061a565b61075790827fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf691820191016107cc565b5f61060d565b9081518082526020808093019301915f5b82811061077c575050505090565b83516001600160a01b03168552938101939281019260010161076e565b9081518082526020808093019301915f5b8281106107b8575050505090565b8351855293810193928101926001016107aa565b8181106107d7575050565b5f81556001016107cc565b80518210156107f65760209160051b010190565b634e487b7160e01b5f52603260045260245ffdfea264697066735822122086251addd62c150283091b1c7f801b2128b1e2a36b8eb431da9c5f9bbf34f06d64736f6c63430008190033",
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
