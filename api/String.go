// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// StringMetaData contains all meta data concerning the String contract.
var StringMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"a\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"b\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"Test\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"A\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req2\",\"type\":\"string\"}],\"name\":\"A1\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"req\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"req2\",\"type\":\"address\"}],\"name\":\"A2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"B\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"}],\"name\":\"C\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req2\",\"type\":\"string\"}],\"name\":\"D\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"req2\",\"type\":\"address\"}],\"name\":\"E\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"req\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req2\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"req3\",\"type\":\"string\"}],\"name\":\"F\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506106d68061001c5f395ff3fe608060405234801561000f575f80fd5b5060043610610085575f3560e01c8063561882871161005857806356188287146101085780637076fab51461011b578063d35cb0d81461012e578063df7fc1431461014c575f80fd5b8063075d36fa14610089578063133b2f5c146100b6578063142906eb146100cb578063341f2919146100de575b5f80fd5b61009f61009736600461025a565b509192909150565b6040516100ad9291906102c6565b60405180910390f35b6100c96100c43660046102f4565b61015f565b005b6100c96100d936600461025a565b610174565b6100f06100ec3660046103ae565b5090565b6040516001600160a01b0390911681526020016100ad565b6100c96101163660046103df565b610187565b6100c961012936600461042f565b610199565b61013f61013c366004610482565b90565b6040516100ad9190610535565b61009f61015a36600461042f565b610212565b5f61016b8688836105e6565b50505050505050565b5f6101808486836105e6565b5050505050565b5f6101938385836105e6565b50505050565b60408051606080825260029082015261686960f01b608082015260a060208201819052600490820152636b796c6560e01b60c082015260648183015290517fffc9941d5e6439cb3286033e3f6715488ea73b0fe4246b2a8b021e77bca297649181900360e00190a15f61020d8284836105e6565b505050565b81815b9250929050565b5f8083601f84011261022c575f80fd5b50813567ffffffffffffffff811115610243575f80fd5b602083019150836020828501011115610215575f80fd5b5f805f806040858703121561026d575f80fd5b843567ffffffffffffffff811115610283575f80fd5b61028f8782880161021c565b909550935050602085013567ffffffffffffffff8111156102ae575f80fd5b6102ba8782880161021c565b95989497509550505050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b5f805f805f8060608789031215610309575f80fd5b863567ffffffffffffffff81111561031f575f80fd5b61032b89828a0161021c565b909750955050602087013567ffffffffffffffff81111561034a575f80fd5b61035689828a0161021c565b909550935050604087013567ffffffffffffffff811115610375575f80fd5b61038189828a0161021c565b979a9699509497509295939492505050565b80356001600160a01b03811681146103a9575f80fd5b919050565b5f80604083850312156103bf575f80fd5b6103c883610393565b91506103d660208401610393565b90509250929050565b5f805f604084860312156103f1575f80fd5b833567ffffffffffffffff811115610407575f80fd5b6104138682870161021c565b9094509250610426905060208501610393565b90509250925092565b5f8060208385031215610440575f80fd5b823567ffffffffffffffff811115610456575f80fd5b6104628582860161021c565b90969095509350505050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610492575f80fd5b813567ffffffffffffffff8111156104a8575f80fd5b8201601f810184136104b8575f80fd5b803567ffffffffffffffff8111156104d2576104d261046e565b604051601f8201601f19908116603f0116810167ffffffffffffffff811182821017156105015761050161046e565b604052818152828201602001861015610518575f80fd5b816020840160208301375f91810160200191909152949350505050565b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b600181811c9082168061057e57607f821691505b60208210810361059c57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111561020d57805f5260205f20601f840160051c810160208510156105c75750805b601f840160051c820191505b81811015610180575f81556001016105d3565b67ffffffffffffffff8311156105fe576105fe61046e565b6106128361060c835461056a565b836105a2565b5f601f841160018114610643575f851561062c5750838201355b5f19600387901b1c1916600186901b178355610180565b5f83815260208120601f198716915b828110156106725786850135825560209485019460019092019101610652565b508682101561068e575f1960f88860031b161c19848701351681555b505060018560011b018355505050505056fea264697066735822122061d82728ff33755e8825a5c847488426d34f64e42f7fb6b129a8cea8ee37f24764736f6c634300081a0033",
}

// StringABI is the input ABI used to generate the binding from.
// Deprecated: Use StringMetaData.ABI instead.
var StringABI = StringMetaData.ABI

// StringBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringMetaData.Bin instead.
var StringBin = StringMetaData.Bin

// DeployString deploys a new Ethereum contract, binding an instance of String to it.
func DeployString(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *String, error) {
	parsed, err := StringMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &String{StringCaller: StringCaller{contract: contract}, StringTransactor: StringTransactor{contract: contract}, StringFilterer: StringFilterer{contract: contract}}, nil
}

// String is an auto generated Go binding around an Ethereum contract.
type String struct {
	StringCaller     // Read-only binding to the contract
	StringTransactor // Write-only binding to the contract
	StringFilterer   // Log filterer for contract events
}

// StringCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringSession struct {
	Contract     *String           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringCallerSession struct {
	Contract *StringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringTransactorSession struct {
	Contract     *StringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringRaw struct {
	Contract *String // Generic contract binding to access the raw methods on
}

// StringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringCallerRaw struct {
	Contract *StringCaller // Generic read-only contract binding to access the raw methods on
}

// StringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringTransactorRaw struct {
	Contract *StringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewString creates a new instance of String, bound to a specific deployed contract.
func NewString(address common.Address, backend bind.ContractBackend) (*String, error) {
	contract, err := bindString(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &String{StringCaller: StringCaller{contract: contract}, StringTransactor: StringTransactor{contract: contract}, StringFilterer: StringFilterer{contract: contract}}, nil
}

// NewStringCaller creates a new read-only instance of String, bound to a specific deployed contract.
func NewStringCaller(address common.Address, caller bind.ContractCaller) (*StringCaller, error) {
	contract, err := bindString(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringCaller{contract: contract}, nil
}

// NewStringTransactor creates a new write-only instance of String, bound to a specific deployed contract.
func NewStringTransactor(address common.Address, transactor bind.ContractTransactor) (*StringTransactor, error) {
	contract, err := bindString(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringTransactor{contract: contract}, nil
}

// NewStringFilterer creates a new log filterer instance of String, bound to a specific deployed contract.
func NewStringFilterer(address common.Address, filterer bind.ContractFilterer) (*StringFilterer, error) {
	contract, err := bindString(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringFilterer{contract: contract}, nil
}

// bindString binds a generic wrapper to an already deployed contract.
func bindString(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StringMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_String *StringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _String.Contract.StringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_String *StringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _String.Contract.StringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_String *StringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _String.Contract.StringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_String *StringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _String.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_String *StringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _String.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_String *StringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _String.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0xdf7fc143.
//
// Solidity: function A(string req) pure returns(string)
func (_String *StringCaller) A(opts *bind.CallOpts, req string) (string, error) {
	var out []interface{}
	err := _String.contract.Call(opts, &out, "A", req)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// A is a free data retrieval call binding the contract method 0xdf7fc143.
//
// Solidity: function A(string req) pure returns(string)
func (_String *StringSession) A(req string) (string, error) {
	return _String.Contract.A(&_String.CallOpts, req)
}

// A is a free data retrieval call binding the contract method 0xdf7fc143.
//
// Solidity: function A(string req) pure returns(string)
func (_String *StringCallerSession) A(req string) (string, error) {
	return _String.Contract.A(&_String.CallOpts, req)
}

// A1 is a free data retrieval call binding the contract method 0x075d36fa.
//
// Solidity: function A1(string req, string req2) pure returns(string)
func (_String *StringCaller) A1(opts *bind.CallOpts, req string, req2 string) (string, error) {
	var out []interface{}
	err := _String.contract.Call(opts, &out, "A1", req, req2)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// A1 is a free data retrieval call binding the contract method 0x075d36fa.
//
// Solidity: function A1(string req, string req2) pure returns(string)
func (_String *StringSession) A1(req string, req2 string) (string, error) {
	return _String.Contract.A1(&_String.CallOpts, req, req2)
}

// A1 is a free data retrieval call binding the contract method 0x075d36fa.
//
// Solidity: function A1(string req, string req2) pure returns(string)
func (_String *StringCallerSession) A1(req string, req2 string) (string, error) {
	return _String.Contract.A1(&_String.CallOpts, req, req2)
}

// A2 is a free data retrieval call binding the contract method 0x341f2919.
//
// Solidity: function A2(address req, address req2) pure returns(address)
func (_String *StringCaller) A2(opts *bind.CallOpts, req common.Address, req2 common.Address) (common.Address, error) {
	var out []interface{}
	err := _String.contract.Call(opts, &out, "A2", req, req2)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// A2 is a free data retrieval call binding the contract method 0x341f2919.
//
// Solidity: function A2(address req, address req2) pure returns(address)
func (_String *StringSession) A2(req common.Address, req2 common.Address) (common.Address, error) {
	return _String.Contract.A2(&_String.CallOpts, req, req2)
}

// A2 is a free data retrieval call binding the contract method 0x341f2919.
//
// Solidity: function A2(address req, address req2) pure returns(address)
func (_String *StringCallerSession) A2(req common.Address, req2 common.Address) (common.Address, error) {
	return _String.Contract.A2(&_String.CallOpts, req, req2)
}

// B is a free data retrieval call binding the contract method 0xd35cb0d8.
//
// Solidity: function B(string req) pure returns(string)
func (_String *StringCaller) B(opts *bind.CallOpts, req string) (string, error) {
	var out []interface{}
	err := _String.contract.Call(opts, &out, "B", req)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// B is a free data retrieval call binding the contract method 0xd35cb0d8.
//
// Solidity: function B(string req) pure returns(string)
func (_String *StringSession) B(req string) (string, error) {
	return _String.Contract.B(&_String.CallOpts, req)
}

// B is a free data retrieval call binding the contract method 0xd35cb0d8.
//
// Solidity: function B(string req) pure returns(string)
func (_String *StringCallerSession) B(req string) (string, error) {
	return _String.Contract.B(&_String.CallOpts, req)
}

// C is a paid mutator transaction binding the contract method 0x7076fab5.
//
// Solidity: function C(string req) returns()
func (_String *StringTransactor) C(opts *bind.TransactOpts, req string) (*types.Transaction, error) {
	return _String.contract.Transact(opts, "C", req)
}

// C is a paid mutator transaction binding the contract method 0x7076fab5.
//
// Solidity: function C(string req) returns()
func (_String *StringSession) C(req string) (*types.Transaction, error) {
	return _String.Contract.C(&_String.TransactOpts, req)
}

// C is a paid mutator transaction binding the contract method 0x7076fab5.
//
// Solidity: function C(string req) returns()
func (_String *StringTransactorSession) C(req string) (*types.Transaction, error) {
	return _String.Contract.C(&_String.TransactOpts, req)
}

// D is a paid mutator transaction binding the contract method 0x142906eb.
//
// Solidity: function D(string req, string req2) returns()
func (_String *StringTransactor) D(opts *bind.TransactOpts, req string, req2 string) (*types.Transaction, error) {
	return _String.contract.Transact(opts, "D", req, req2)
}

// D is a paid mutator transaction binding the contract method 0x142906eb.
//
// Solidity: function D(string req, string req2) returns()
func (_String *StringSession) D(req string, req2 string) (*types.Transaction, error) {
	return _String.Contract.D(&_String.TransactOpts, req, req2)
}

// D is a paid mutator transaction binding the contract method 0x142906eb.
//
// Solidity: function D(string req, string req2) returns()
func (_String *StringTransactorSession) D(req string, req2 string) (*types.Transaction, error) {
	return _String.Contract.D(&_String.TransactOpts, req, req2)
}

// E is a paid mutator transaction binding the contract method 0x56188287.
//
// Solidity: function E(string req, address req2) returns()
func (_String *StringTransactor) E(opts *bind.TransactOpts, req string, req2 common.Address) (*types.Transaction, error) {
	return _String.contract.Transact(opts, "E", req, req2)
}

// E is a paid mutator transaction binding the contract method 0x56188287.
//
// Solidity: function E(string req, address req2) returns()
func (_String *StringSession) E(req string, req2 common.Address) (*types.Transaction, error) {
	return _String.Contract.E(&_String.TransactOpts, req, req2)
}

// E is a paid mutator transaction binding the contract method 0x56188287.
//
// Solidity: function E(string req, address req2) returns()
func (_String *StringTransactorSession) E(req string, req2 common.Address) (*types.Transaction, error) {
	return _String.Contract.E(&_String.TransactOpts, req, req2)
}

// F is a paid mutator transaction binding the contract method 0x133b2f5c.
//
// Solidity: function F(string req, string req2, string req3) returns()
func (_String *StringTransactor) F(opts *bind.TransactOpts, req string, req2 string, req3 string) (*types.Transaction, error) {
	return _String.contract.Transact(opts, "F", req, req2, req3)
}

// F is a paid mutator transaction binding the contract method 0x133b2f5c.
//
// Solidity: function F(string req, string req2, string req3) returns()
func (_String *StringSession) F(req string, req2 string, req3 string) (*types.Transaction, error) {
	return _String.Contract.F(&_String.TransactOpts, req, req2, req3)
}

// F is a paid mutator transaction binding the contract method 0x133b2f5c.
//
// Solidity: function F(string req, string req2, string req3) returns()
func (_String *StringTransactorSession) F(req string, req2 string, req3 string) (*types.Transaction, error) {
	return _String.Contract.F(&_String.TransactOpts, req, req2, req3)
}

// StringTestIterator is returned from FilterTest and is used to iterate over the raw logs and unpacked data for Test events raised by the String contract.
type StringTestIterator struct {
	Event *StringTest // Event containing the contract specifics and raw log

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
func (it *StringTestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StringTest)
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
		it.Event = new(StringTest)
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
func (it *StringTestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StringTestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StringTest represents a Test event raised by the String contract.
type StringTest struct {
	A   string
	B   string
	C   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTest is a free log retrieval operation binding the contract event 0xffc9941d5e6439cb3286033e3f6715488ea73b0fe4246b2a8b021e77bca29764.
//
// Solidity: event Test(string a, string b, uint256 c)
func (_String *StringFilterer) FilterTest(opts *bind.FilterOpts) (*StringTestIterator, error) {

	logs, sub, err := _String.contract.FilterLogs(opts, "Test")
	if err != nil {
		return nil, err
	}
	return &StringTestIterator{contract: _String.contract, event: "Test", logs: logs, sub: sub}, nil
}

// WatchTest is a free log subscription operation binding the contract event 0xffc9941d5e6439cb3286033e3f6715488ea73b0fe4246b2a8b021e77bca29764.
//
// Solidity: event Test(string a, string b, uint256 c)
func (_String *StringFilterer) WatchTest(opts *bind.WatchOpts, sink chan<- *StringTest) (event.Subscription, error) {

	logs, sub, err := _String.contract.WatchLogs(opts, "Test")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StringTest)
				if err := _String.contract.UnpackLog(event, "Test", log); err != nil {
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

// ParseTest is a log parse operation binding the contract event 0xffc9941d5e6439cb3286033e3f6715488ea73b0fe4246b2a8b021e77bca29764.
//
// Solidity: event Test(string a, string b, uint256 c)
func (_String *StringFilterer) ParseTest(log types.Log) (*StringTest, error) {
	event := new(StringTest)
	if err := _String.contract.UnpackLog(event, "Test", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
