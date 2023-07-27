// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mb

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

// MTSBridgeMetaData contains all meta data concerning the MTSBridge contract.
var MTSBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockMts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockMts\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"lockMts\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mtsAmountMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subMts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockMts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MTSBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use MTSBridgeMetaData.ABI instead.
var MTSBridgeABI = MTSBridgeMetaData.ABI

// MTSBridge is an auto generated Go binding around an Ethereum contract.
type MTSBridge struct {
	MTSBridgeCaller     // Read-only binding to the contract
	MTSBridgeTransactor // Write-only binding to the contract
	MTSBridgeFilterer   // Log filterer for contract events
}

// MTSBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MTSBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTSBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MTSBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTSBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MTSBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTSBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MTSBridgeSession struct {
	Contract     *MTSBridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MTSBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MTSBridgeCallerSession struct {
	Contract *MTSBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MTSBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MTSBridgeTransactorSession struct {
	Contract     *MTSBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MTSBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MTSBridgeRaw struct {
	Contract *MTSBridge // Generic contract binding to access the raw methods on
}

// MTSBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MTSBridgeCallerRaw struct {
	Contract *MTSBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// MTSBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MTSBridgeTransactorRaw struct {
	Contract *MTSBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMTSBridge creates a new instance of MTSBridge, bound to a specific deployed contract.
func NewMTSBridge(address common.Address, backend bind.ContractBackend) (*MTSBridge, error) {
	contract, err := bindMTSBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MTSBridge{MTSBridgeCaller: MTSBridgeCaller{contract: contract}, MTSBridgeTransactor: MTSBridgeTransactor{contract: contract}, MTSBridgeFilterer: MTSBridgeFilterer{contract: contract}}, nil
}

// NewMTSBridgeCaller creates a new read-only instance of MTSBridge, bound to a specific deployed contract.
func NewMTSBridgeCaller(address common.Address, caller bind.ContractCaller) (*MTSBridgeCaller, error) {
	contract, err := bindMTSBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MTSBridgeCaller{contract: contract}, nil
}

// NewMTSBridgeTransactor creates a new write-only instance of MTSBridge, bound to a specific deployed contract.
func NewMTSBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*MTSBridgeTransactor, error) {
	contract, err := bindMTSBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MTSBridgeTransactor{contract: contract}, nil
}

// NewMTSBridgeFilterer creates a new log filterer instance of MTSBridge, bound to a specific deployed contract.
func NewMTSBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*MTSBridgeFilterer, error) {
	contract, err := bindMTSBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MTSBridgeFilterer{contract: contract}, nil
}

// bindMTSBridge binds a generic wrapper to an already deployed contract.
func bindMTSBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MTSBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MTSBridge *MTSBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MTSBridge.Contract.MTSBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MTSBridge *MTSBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTSBridge.Contract.MTSBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MTSBridge *MTSBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MTSBridge.Contract.MTSBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MTSBridge *MTSBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MTSBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MTSBridge *MTSBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTSBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MTSBridge *MTSBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MTSBridge.Contract.contract.Transact(opts, method, params...)
}

// MtsAmountMap is a free data retrieval call binding the contract method 0x80a25504.
//
// Solidity: function mtsAmountMap(address ) view returns(uint256)
func (_MTSBridge *MTSBridgeCaller) MtsAmountMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MTSBridge.contract.Call(opts, &out, "mtsAmountMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MtsAmountMap is a free data retrieval call binding the contract method 0x80a25504.
//
// Solidity: function mtsAmountMap(address ) view returns(uint256)
func (_MTSBridge *MTSBridgeSession) MtsAmountMap(arg0 common.Address) (*big.Int, error) {
	return _MTSBridge.Contract.MtsAmountMap(&_MTSBridge.CallOpts, arg0)
}

// MtsAmountMap is a free data retrieval call binding the contract method 0x80a25504.
//
// Solidity: function mtsAmountMap(address ) view returns(uint256)
func (_MTSBridge *MTSBridgeCallerSession) MtsAmountMap(arg0 common.Address) (*big.Int, error) {
	return _MTSBridge.Contract.MtsAmountMap(&_MTSBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MTSBridge *MTSBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MTSBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MTSBridge *MTSBridgeSession) Owner() (common.Address, error) {
	return _MTSBridge.Contract.Owner(&_MTSBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MTSBridge *MTSBridgeCallerSession) Owner() (common.Address, error) {
	return _MTSBridge.Contract.Owner(&_MTSBridge.CallOpts)
}

// LockMts is a paid mutator transaction binding the contract method 0x8a944eed.
//
// Solidity: function lockMts() payable returns()
func (_MTSBridge *MTSBridgeTransactor) LockMts(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTSBridge.contract.Transact(opts, "lockMts")
}

// LockMts is a paid mutator transaction binding the contract method 0x8a944eed.
//
// Solidity: function lockMts() payable returns()
func (_MTSBridge *MTSBridgeSession) LockMts() (*types.Transaction, error) {
	return _MTSBridge.Contract.LockMts(&_MTSBridge.TransactOpts)
}

// LockMts is a paid mutator transaction binding the contract method 0x8a944eed.
//
// Solidity: function lockMts() payable returns()
func (_MTSBridge *MTSBridgeTransactorSession) LockMts() (*types.Transaction, error) {
	return _MTSBridge.Contract.LockMts(&_MTSBridge.TransactOpts)
}

// SubMts is a paid mutator transaction binding the contract method 0xc5d56653.
//
// Solidity: function subMts(address account, uint256 amount) returns()
func (_MTSBridge *MTSBridgeTransactor) SubMts(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MTSBridge.contract.Transact(opts, "subMts", account, amount)
}

// SubMts is a paid mutator transaction binding the contract method 0xc5d56653.
//
// Solidity: function subMts(address account, uint256 amount) returns()
func (_MTSBridge *MTSBridgeSession) SubMts(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MTSBridge.Contract.SubMts(&_MTSBridge.TransactOpts, account, amount)
}

// SubMts is a paid mutator transaction binding the contract method 0xc5d56653.
//
// Solidity: function subMts(address account, uint256 amount) returns()
func (_MTSBridge *MTSBridgeTransactorSession) SubMts(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MTSBridge.Contract.SubMts(&_MTSBridge.TransactOpts, account, amount)
}

// UnlockMts is a paid mutator transaction binding the contract method 0x41004b14.
//
// Solidity: function unlockMts(address _to, uint256 _amount) returns()
func (_MTSBridge *MTSBridgeTransactor) UnlockMts(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MTSBridge.contract.Transact(opts, "unlockMts", _to, _amount)
}

// UnlockMts is a paid mutator transaction binding the contract method 0x41004b14.
//
// Solidity: function unlockMts(address _to, uint256 _amount) returns()
func (_MTSBridge *MTSBridgeSession) UnlockMts(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MTSBridge.Contract.UnlockMts(&_MTSBridge.TransactOpts, _to, _amount)
}

// UnlockMts is a paid mutator transaction binding the contract method 0x41004b14.
//
// Solidity: function unlockMts(address _to, uint256 _amount) returns()
func (_MTSBridge *MTSBridgeTransactorSession) UnlockMts(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _MTSBridge.Contract.UnlockMts(&_MTSBridge.TransactOpts, _to, _amount)
}

// MTSBridgeLockMtsIterator is returned from FilterLockMts and is used to iterate over the raw logs and unpacked data for LockMts events raised by the MTSBridge contract.
type MTSBridgeLockMtsIterator struct {
	Event *MTSBridgeLockMts // Event containing the contract specifics and raw log

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
func (it *MTSBridgeLockMtsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTSBridgeLockMts)
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
		it.Event = new(MTSBridgeLockMts)
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
func (it *MTSBridgeLockMtsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTSBridgeLockMtsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTSBridgeLockMts represents a LockMts event raised by the MTSBridge contract.
type MTSBridgeLockMts struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockMts is a free log retrieval operation binding the contract event 0x60745bb8f47da8fa85b2fa96982f4ca228359f2e8ab628004429eb0d67427c25.
//
// Solidity: event LockMts(address indexed account, uint256 amount)
func (_MTSBridge *MTSBridgeFilterer) FilterLockMts(opts *bind.FilterOpts, account []common.Address) (*MTSBridgeLockMtsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MTSBridge.contract.FilterLogs(opts, "LockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return &MTSBridgeLockMtsIterator{contract: _MTSBridge.contract, event: "LockMts", logs: logs, sub: sub}, nil
}

// WatchLockMts is a free log subscription operation binding the contract event 0x60745bb8f47da8fa85b2fa96982f4ca228359f2e8ab628004429eb0d67427c25.
//
// Solidity: event LockMts(address indexed account, uint256 amount)
func (_MTSBridge *MTSBridgeFilterer) WatchLockMts(opts *bind.WatchOpts, sink chan<- *MTSBridgeLockMts, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MTSBridge.contract.WatchLogs(opts, "LockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTSBridgeLockMts)
				if err := _MTSBridge.contract.UnpackLog(event, "LockMts", log); err != nil {
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

// ParseLockMts is a log parse operation binding the contract event 0x60745bb8f47da8fa85b2fa96982f4ca228359f2e8ab628004429eb0d67427c25.
//
// Solidity: event LockMts(address indexed account, uint256 amount)
func (_MTSBridge *MTSBridgeFilterer) ParseLockMts(log types.Log) (*MTSBridgeLockMts, error) {
	event := new(MTSBridgeLockMts)
	if err := _MTSBridge.contract.UnpackLog(event, "LockMts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTSBridgeUnlockMtsIterator is returned from FilterUnlockMts and is used to iterate over the raw logs and unpacked data for UnlockMts events raised by the MTSBridge contract.
type MTSBridgeUnlockMtsIterator struct {
	Event *MTSBridgeUnlockMts // Event containing the contract specifics and raw log

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
func (it *MTSBridgeUnlockMtsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTSBridgeUnlockMts)
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
		it.Event = new(MTSBridgeUnlockMts)
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
func (it *MTSBridgeUnlockMtsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTSBridgeUnlockMtsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTSBridgeUnlockMts represents a UnlockMts event raised by the MTSBridge contract.
type MTSBridgeUnlockMts struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlockMts is a free log retrieval operation binding the contract event 0x8692c5be9f8792a83104866c4465016f4928e73a17f2240e5a7e70d1ab2af2a4.
//
// Solidity: event UnlockMts(address indexed account, uint256 amount)
func (_MTSBridge *MTSBridgeFilterer) FilterUnlockMts(opts *bind.FilterOpts, account []common.Address) (*MTSBridgeUnlockMtsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MTSBridge.contract.FilterLogs(opts, "UnlockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return &MTSBridgeUnlockMtsIterator{contract: _MTSBridge.contract, event: "UnlockMts", logs: logs, sub: sub}, nil
}

// WatchUnlockMts is a free log subscription operation binding the contract event 0x8692c5be9f8792a83104866c4465016f4928e73a17f2240e5a7e70d1ab2af2a4.
//
// Solidity: event UnlockMts(address indexed account, uint256 amount)
func (_MTSBridge *MTSBridgeFilterer) WatchUnlockMts(opts *bind.WatchOpts, sink chan<- *MTSBridgeUnlockMts, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MTSBridge.contract.WatchLogs(opts, "UnlockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTSBridgeUnlockMts)
				if err := _MTSBridge.contract.UnpackLog(event, "UnlockMts", log); err != nil {
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

// ParseUnlockMts is a log parse operation binding the contract event 0x8692c5be9f8792a83104866c4465016f4928e73a17f2240e5a7e70d1ab2af2a4.
//
// Solidity: event UnlockMts(address indexed account, uint256 amount)
func (_MTSBridge *MTSBridgeFilterer) ParseUnlockMts(log types.Log) (*MTSBridgeUnlockMts, error) {
	event := new(MTSBridgeUnlockMts)
	if err := _MTSBridge.contract.UnpackLog(event, "UnlockMts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
