// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package suyuanContract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SuyuanContractABI is the input ABI used to generate the binding from.
const SuyuanContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"mapWriter\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get_block\",\"outputs\":[{\"name\":\"block_hash\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"qrcode\",\"type\":\"uint256\"},{\"name\":\"fn_name\",\"type\":\"uint256\"}],\"name\":\"get_info\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"infos\",\"type\":\"string\"},{\"name\":\"blocknumber\",\"type\":\"uint256\"},{\"name\":\"images\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"trace_index\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"writer\",\"type\":\"address\"}],\"name\":\"addWriter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"qrcode\",\"type\":\"uint256\"},{\"name\":\"fn_name\",\"type\":\"uint256\"},{\"name\":\"infos\",\"type\":\"string\"},{\"name\":\"images\",\"type\":\"string\"}],\"name\":\"set_info\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"table\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ss\",\"type\":\"string\"}],\"name\":\"updateEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"}],\"name\":\"item\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"string\"}],\"name\":\"list1\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"listname\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"list\",\"type\":\"string\"}],\"name\":\"inserted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SuyuanContract is an auto generated Go binding around an Ethereum contract.
type SuyuanContract struct {
	SuyuanContractCaller     // Read-only binding to the contract
	SuyuanContractTransactor // Write-only binding to the contract
	SuyuanContractFilterer   // Log filterer for contract events
}

// SuyuanContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type SuyuanContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuyuanContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SuyuanContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuyuanContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SuyuanContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SuyuanContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SuyuanContractSession struct {
	Contract     *SuyuanContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SuyuanContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SuyuanContractCallerSession struct {
	Contract *SuyuanContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SuyuanContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SuyuanContractTransactorSession struct {
	Contract     *SuyuanContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SuyuanContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type SuyuanContractRaw struct {
	Contract *SuyuanContract // Generic contract binding to access the raw methods on
}

// SuyuanContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SuyuanContractCallerRaw struct {
	Contract *SuyuanContractCaller // Generic read-only contract binding to access the raw methods on
}

// SuyuanContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SuyuanContractTransactorRaw struct {
	Contract *SuyuanContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSuyuanContract creates a new instance of SuyuanContract, bound to a specific deployed contract.
func NewSuyuanContract(address common.Address, backend bind.ContractBackend) (*SuyuanContract, error) {
	contract, err := bindSuyuanContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SuyuanContract{SuyuanContractCaller: SuyuanContractCaller{contract: contract}, SuyuanContractTransactor: SuyuanContractTransactor{contract: contract}, SuyuanContractFilterer: SuyuanContractFilterer{contract: contract}}, nil
}

// NewSuyuanContractCaller creates a new read-only instance of SuyuanContract, bound to a specific deployed contract.
func NewSuyuanContractCaller(address common.Address, caller bind.ContractCaller) (*SuyuanContractCaller, error) {
	contract, err := bindSuyuanContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SuyuanContractCaller{contract: contract}, nil
}

// NewSuyuanContractTransactor creates a new write-only instance of SuyuanContract, bound to a specific deployed contract.
func NewSuyuanContractTransactor(address common.Address, transactor bind.ContractTransactor) (*SuyuanContractTransactor, error) {
	contract, err := bindSuyuanContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SuyuanContractTransactor{contract: contract}, nil
}

// NewSuyuanContractFilterer creates a new log filterer instance of SuyuanContract, bound to a specific deployed contract.
func NewSuyuanContractFilterer(address common.Address, filterer bind.ContractFilterer) (*SuyuanContractFilterer, error) {
	contract, err := bindSuyuanContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SuyuanContractFilterer{contract: contract}, nil
}

// bindSuyuanContract binds a generic wrapper to an already deployed contract.
func bindSuyuanContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SuyuanContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuyuanContract *SuyuanContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SuyuanContract.Contract.SuyuanContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuyuanContract *SuyuanContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuyuanContract.Contract.SuyuanContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuyuanContract *SuyuanContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuyuanContract.Contract.SuyuanContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SuyuanContract *SuyuanContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SuyuanContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SuyuanContract *SuyuanContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuyuanContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SuyuanContract *SuyuanContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SuyuanContract.Contract.contract.Transact(opts, method, params...)
}

// GetBlock is a free data retrieval call binding the contract method 0x69a1a596.
//
// Solidity: function get_block() constant returns(uint256 block_hash, bytes32, bytes32)
func (_SuyuanContract *SuyuanContractCaller) GetBlock(opts *bind.CallOpts) (*big.Int, [32]byte, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _SuyuanContract.contract.Call(opts, out, "get_block")
	return *ret0, *ret1, *ret2, err
}

// GetBlock is a free data retrieval call binding the contract method 0x69a1a596.
//
// Solidity: function get_block() constant returns(uint256 block_hash, bytes32, bytes32)
func (_SuyuanContract *SuyuanContractSession) GetBlock() (*big.Int, [32]byte, [32]byte, error) {
	return _SuyuanContract.Contract.GetBlock(&_SuyuanContract.CallOpts)
}

// GetBlock is a free data retrieval call binding the contract method 0x69a1a596.
//
// Solidity: function get_block() constant returns(uint256 block_hash, bytes32, bytes32)
func (_SuyuanContract *SuyuanContractCallerSession) GetBlock() (*big.Int, [32]byte, [32]byte, error) {
	return _SuyuanContract.Contract.GetBlock(&_SuyuanContract.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0xb39c4c31.
//
// Solidity: function get_info(uint256 qrcode, uint256 fn_name) constant returns(uint256, uint256, string infos, uint256 blocknumber, string images)
func (_SuyuanContract *SuyuanContractCaller) GetInfo(opts *bind.CallOpts, qrcode *big.Int, fn_name *big.Int) (*big.Int, *big.Int, string, *big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(string)
		ret3 = new(*big.Int)
		ret4 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _SuyuanContract.contract.Call(opts, out, "get_info", qrcode, fn_name)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetInfo is a free data retrieval call binding the contract method 0xb39c4c31.
//
// Solidity: function get_info(uint256 qrcode, uint256 fn_name) constant returns(uint256, uint256, string infos, uint256 blocknumber, string images)
func (_SuyuanContract *SuyuanContractSession) GetInfo(qrcode *big.Int, fn_name *big.Int) (*big.Int, *big.Int, string, *big.Int, string, error) {
	return _SuyuanContract.Contract.GetInfo(&_SuyuanContract.CallOpts, qrcode, fn_name)
}

// GetInfo is a free data retrieval call binding the contract method 0xb39c4c31.
//
// Solidity: function get_info(uint256 qrcode, uint256 fn_name) constant returns(uint256, uint256, string infos, uint256 blocknumber, string images)
func (_SuyuanContract *SuyuanContractCallerSession) GetInfo(qrcode *big.Int, fn_name *big.Int) (*big.Int, *big.Int, string, *big.Int, string, error) {
	return _SuyuanContract.Contract.GetInfo(&_SuyuanContract.CallOpts, qrcode, fn_name)
}

// MapWriter is a free data retrieval call binding the contract method 0x30f3eebc.
//
// Solidity: function mapWriter(address ) constant returns(bool)
func (_SuyuanContract *SuyuanContractCaller) MapWriter(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SuyuanContract.contract.Call(opts, out, "mapWriter", arg0)
	return *ret0, err
}

// MapWriter is a free data retrieval call binding the contract method 0x30f3eebc.
//
// Solidity: function mapWriter(address ) constant returns(bool)
func (_SuyuanContract *SuyuanContractSession) MapWriter(arg0 common.Address) (bool, error) {
	return _SuyuanContract.Contract.MapWriter(&_SuyuanContract.CallOpts, arg0)
}

// MapWriter is a free data retrieval call binding the contract method 0x30f3eebc.
//
// Solidity: function mapWriter(address ) constant returns(bool)
func (_SuyuanContract *SuyuanContractCallerSession) MapWriter(arg0 common.Address) (bool, error) {
	return _SuyuanContract.Contract.MapWriter(&_SuyuanContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SuyuanContract *SuyuanContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SuyuanContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SuyuanContract *SuyuanContractSession) Owner() (common.Address, error) {
	return _SuyuanContract.Contract.Owner(&_SuyuanContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SuyuanContract *SuyuanContractCallerSession) Owner() (common.Address, error) {
	return _SuyuanContract.Contract.Owner(&_SuyuanContract.CallOpts)
}

// TraceIndex is a free data retrieval call binding the contract method 0xd2445621.
//
// Solidity: function trace_index() constant returns(uint256)
func (_SuyuanContract *SuyuanContractCaller) TraceIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SuyuanContract.contract.Call(opts, out, "trace_index")
	return *ret0, err
}

// TraceIndex is a free data retrieval call binding the contract method 0xd2445621.
//
// Solidity: function trace_index() constant returns(uint256)
func (_SuyuanContract *SuyuanContractSession) TraceIndex() (*big.Int, error) {
	return _SuyuanContract.Contract.TraceIndex(&_SuyuanContract.CallOpts)
}

// TraceIndex is a free data retrieval call binding the contract method 0xd2445621.
//
// Solidity: function trace_index() constant returns(uint256)
func (_SuyuanContract *SuyuanContractCallerSession) TraceIndex() (*big.Int, error) {
	return _SuyuanContract.Contract.TraceIndex(&_SuyuanContract.CallOpts)
}

// AddWriter is a paid mutator transaction binding the contract method 0xda2824a8.
//
// Solidity: function addWriter(address writer) returns()
func (_SuyuanContract *SuyuanContractTransactor) AddWriter(opts *bind.TransactOpts, writer common.Address) (*types.Transaction, error) {
	return _SuyuanContract.contract.Transact(opts, "addWriter", writer)
}

// AddWriter is a paid mutator transaction binding the contract method 0xda2824a8.
//
// Solidity: function addWriter(address writer) returns()
func (_SuyuanContract *SuyuanContractSession) AddWriter(writer common.Address) (*types.Transaction, error) {
	return _SuyuanContract.Contract.AddWriter(&_SuyuanContract.TransactOpts, writer)
}

// AddWriter is a paid mutator transaction binding the contract method 0xda2824a8.
//
// Solidity: function addWriter(address writer) returns()
func (_SuyuanContract *SuyuanContractTransactorSession) AddWriter(writer common.Address) (*types.Transaction, error) {
	return _SuyuanContract.Contract.AddWriter(&_SuyuanContract.TransactOpts, writer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SuyuanContract *SuyuanContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SuyuanContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SuyuanContract *SuyuanContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _SuyuanContract.Contract.RenounceOwnership(&_SuyuanContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SuyuanContract *SuyuanContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SuyuanContract.Contract.RenounceOwnership(&_SuyuanContract.TransactOpts)
}

// SetInfo is a paid mutator transaction binding the contract method 0xe41e236c.
//
// Solidity: function set_info(uint256 qrcode, uint256 fn_name, string infos, string images) returns(uint256, uint256)
func (_SuyuanContract *SuyuanContractTransactor) SetInfo(opts *bind.TransactOpts, qrcode *big.Int, fn_name *big.Int, infos string, images string) (*types.Transaction, error) {
	return _SuyuanContract.contract.Transact(opts, "set_info", qrcode, fn_name, infos, images)
}

// SetInfo is a paid mutator transaction binding the contract method 0xe41e236c.
//
// Solidity: function set_info(uint256 qrcode, uint256 fn_name, string infos, string images) returns(uint256, uint256)
func (_SuyuanContract *SuyuanContractSession) SetInfo(qrcode *big.Int, fn_name *big.Int, infos string, images string) (*types.Transaction, error) {
	return _SuyuanContract.Contract.SetInfo(&_SuyuanContract.TransactOpts, qrcode, fn_name, infos, images)
}

// SetInfo is a paid mutator transaction binding the contract method 0xe41e236c.
//
// Solidity: function set_info(uint256 qrcode, uint256 fn_name, string infos, string images) returns(uint256, uint256)
func (_SuyuanContract *SuyuanContractTransactorSession) SetInfo(qrcode *big.Int, fn_name *big.Int, infos string, images string) (*types.Transaction, error) {
	return _SuyuanContract.Contract.SetInfo(&_SuyuanContract.TransactOpts, qrcode, fn_name, infos, images)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_SuyuanContract *SuyuanContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _SuyuanContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_SuyuanContract *SuyuanContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SuyuanContract.Contract.TransferOwnership(&_SuyuanContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_SuyuanContract *SuyuanContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _SuyuanContract.Contract.TransferOwnership(&_SuyuanContract.TransactOpts, _newOwner)
}

// SuyuanContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the SuyuanContract contract.
type SuyuanContractOwnershipRenouncedIterator struct {
	Event *SuyuanContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *SuyuanContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuyuanContractOwnershipRenounced)
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
		it.Event = new(SuyuanContractOwnershipRenounced)
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
func (it *SuyuanContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuyuanContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuyuanContractOwnershipRenounced represents a OwnershipRenounced event raised by the SuyuanContract contract.
type SuyuanContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_SuyuanContract *SuyuanContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*SuyuanContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _SuyuanContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SuyuanContractOwnershipRenouncedIterator{contract: _SuyuanContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_SuyuanContract *SuyuanContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *SuyuanContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _SuyuanContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuyuanContractOwnershipRenounced)
				if err := _SuyuanContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// SuyuanContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SuyuanContract contract.
type SuyuanContractOwnershipTransferredIterator struct {
	Event *SuyuanContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SuyuanContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuyuanContractOwnershipTransferred)
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
		it.Event = new(SuyuanContractOwnershipTransferred)
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
func (it *SuyuanContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuyuanContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuyuanContractOwnershipTransferred represents a OwnershipTransferred event raised by the SuyuanContract contract.
type SuyuanContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SuyuanContract *SuyuanContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SuyuanContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SuyuanContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SuyuanContractOwnershipTransferredIterator{contract: _SuyuanContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SuyuanContract *SuyuanContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SuyuanContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SuyuanContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuyuanContractOwnershipTransferred)
				if err := _SuyuanContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SuyuanContractInsertedIterator is returned from FilterInserted and is used to iterate over the raw logs and unpacked data for Inserted events raised by the SuyuanContract contract.
type SuyuanContractInsertedIterator struct {
	Event *SuyuanContractInserted // Event containing the contract specifics and raw log

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
func (it *SuyuanContractInsertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuyuanContractInserted)
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
		it.Event = new(SuyuanContractInserted)
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
func (it *SuyuanContractInsertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuyuanContractInsertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuyuanContractInserted represents a Inserted event raised by the SuyuanContract contract.
type SuyuanContractInserted struct {
	Listname string
	List     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterInserted is a free log retrieval operation binding the contract event 0xa98ce3f73bb62495b18723690f4da9ca551ebd64be17915dbd279400487f1d1a.
//
// Solidity: event inserted(string listname, string list)
func (_SuyuanContract *SuyuanContractFilterer) FilterInserted(opts *bind.FilterOpts) (*SuyuanContractInsertedIterator, error) {

	logs, sub, err := _SuyuanContract.contract.FilterLogs(opts, "inserted")
	if err != nil {
		return nil, err
	}
	return &SuyuanContractInsertedIterator{contract: _SuyuanContract.contract, event: "inserted", logs: logs, sub: sub}, nil
}

// WatchInserted is a free log subscription operation binding the contract event 0xa98ce3f73bb62495b18723690f4da9ca551ebd64be17915dbd279400487f1d1a.
//
// Solidity: event inserted(string listname, string list)
func (_SuyuanContract *SuyuanContractFilterer) WatchInserted(opts *bind.WatchOpts, sink chan<- *SuyuanContractInserted) (event.Subscription, error) {

	logs, sub, err := _SuyuanContract.contract.WatchLogs(opts, "inserted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuyuanContractInserted)
				if err := _SuyuanContract.contract.UnpackLog(event, "inserted", log); err != nil {
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

// SuyuanContractItemIterator is returned from FilterItem and is used to iterate over the raw logs and unpacked data for Item events raised by the SuyuanContract contract.
type SuyuanContractItemIterator struct {
	Event *SuyuanContractItem // Event containing the contract specifics and raw log

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
func (it *SuyuanContractItemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuyuanContractItem)
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
		it.Event = new(SuyuanContractItem)
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
func (it *SuyuanContractItemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuyuanContractItemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuyuanContractItem represents a Item event raised by the SuyuanContract contract.
type SuyuanContractItem struct {
	Key   string
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItem is a free log retrieval operation binding the contract event 0x7efd2f0fa5ca3bc4bfd4a4da201138210a82e5a8f71ab9a939527d283cdf7725.
//
// Solidity: event item(string key, string value)
func (_SuyuanContract *SuyuanContractFilterer) FilterItem(opts *bind.FilterOpts) (*SuyuanContractItemIterator, error) {

	logs, sub, err := _SuyuanContract.contract.FilterLogs(opts, "item")
	if err != nil {
		return nil, err
	}
	return &SuyuanContractItemIterator{contract: _SuyuanContract.contract, event: "item", logs: logs, sub: sub}, nil
}

// WatchItem is a free log subscription operation binding the contract event 0x7efd2f0fa5ca3bc4bfd4a4da201138210a82e5a8f71ab9a939527d283cdf7725.
//
// Solidity: event item(string key, string value)
func (_SuyuanContract *SuyuanContractFilterer) WatchItem(opts *bind.WatchOpts, sink chan<- *SuyuanContractItem) (event.Subscription, error) {

	logs, sub, err := _SuyuanContract.contract.WatchLogs(opts, "item")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuyuanContractItem)
				if err := _SuyuanContract.contract.UnpackLog(event, "item", log); err != nil {
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

// SuyuanContractList1Iterator is returned from FilterList1 and is used to iterate over the raw logs and unpacked data for List1 events raised by the SuyuanContract contract.
type SuyuanContractList1Iterator struct {
	Event *SuyuanContractList1 // Event containing the contract specifics and raw log

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
func (it *SuyuanContractList1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuyuanContractList1)
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
		it.Event = new(SuyuanContractList1)
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
func (it *SuyuanContractList1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuyuanContractList1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuyuanContractList1 represents a List1 event raised by the SuyuanContract contract.
type SuyuanContractList1 struct {
	Key   string
	Value string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterList1 is a free log retrieval operation binding the contract event 0x2a2f50b6564a21a00f4cd3183ad7391b98c32969997b71d4ca1c40914abb4d48.
//
// Solidity: event list1(string key, string value)
func (_SuyuanContract *SuyuanContractFilterer) FilterList1(opts *bind.FilterOpts) (*SuyuanContractList1Iterator, error) {

	logs, sub, err := _SuyuanContract.contract.FilterLogs(opts, "list1")
	if err != nil {
		return nil, err
	}
	return &SuyuanContractList1Iterator{contract: _SuyuanContract.contract, event: "list1", logs: logs, sub: sub}, nil
}

// WatchList1 is a free log subscription operation binding the contract event 0x2a2f50b6564a21a00f4cd3183ad7391b98c32969997b71d4ca1c40914abb4d48.
//
// Solidity: event list1(string key, string value)
func (_SuyuanContract *SuyuanContractFilterer) WatchList1(opts *bind.WatchOpts, sink chan<- *SuyuanContractList1) (event.Subscription, error) {

	logs, sub, err := _SuyuanContract.contract.WatchLogs(opts, "list1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuyuanContractList1)
				if err := _SuyuanContract.contract.UnpackLog(event, "list1", log); err != nil {
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

// SuyuanContractUpdateEventIterator is returned from FilterUpdateEvent and is used to iterate over the raw logs and unpacked data for UpdateEvent events raised by the SuyuanContract contract.
type SuyuanContractUpdateEventIterator struct {
	Event *SuyuanContractUpdateEvent // Event containing the contract specifics and raw log

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
func (it *SuyuanContractUpdateEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SuyuanContractUpdateEvent)
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
		it.Event = new(SuyuanContractUpdateEvent)
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
func (it *SuyuanContractUpdateEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SuyuanContractUpdateEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SuyuanContractUpdateEvent represents a UpdateEvent event raised by the SuyuanContract contract.
type SuyuanContractUpdateEvent struct {
	Table string
	Index *big.Int
	Ss    string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUpdateEvent is a free log retrieval operation binding the contract event 0xcf4e8e6d2193b96484eebe88da60b88447300cf60f48f80ef59ca914caf943e7.
//
// Solidity: event updateEvent(string table, uint256 index, string ss)
func (_SuyuanContract *SuyuanContractFilterer) FilterUpdateEvent(opts *bind.FilterOpts) (*SuyuanContractUpdateEventIterator, error) {

	logs, sub, err := _SuyuanContract.contract.FilterLogs(opts, "updateEvent")
	if err != nil {
		return nil, err
	}
	return &SuyuanContractUpdateEventIterator{contract: _SuyuanContract.contract, event: "updateEvent", logs: logs, sub: sub}, nil
}

// WatchUpdateEvent is a free log subscription operation binding the contract event 0xcf4e8e6d2193b96484eebe88da60b88447300cf60f48f80ef59ca914caf943e7.
//
// Solidity: event updateEvent(string table, uint256 index, string ss)
func (_SuyuanContract *SuyuanContractFilterer) WatchUpdateEvent(opts *bind.WatchOpts, sink chan<- *SuyuanContractUpdateEvent) (event.Subscription, error) {

	logs, sub, err := _SuyuanContract.contract.WatchLogs(opts, "updateEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SuyuanContractUpdateEvent)
				if err := _SuyuanContract.contract.UnpackLog(event, "updateEvent", log); err != nil {
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
