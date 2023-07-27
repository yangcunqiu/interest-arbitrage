// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eb

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

// ETHBridgeMetaData contains all meta data concerning the ETHBridge contract.
var ETHBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_mts\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LockMts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"UnlockMts\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lockMts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mts\",\"outputs\":[{\"internalType\":\"contractWMTS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mtsAmountMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"subMts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unlockMts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ETHBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use ETHBridgeMetaData.ABI instead.
var ETHBridgeABI = ETHBridgeMetaData.ABI

// ETHBridge is an auto generated Go binding around an Ethereum contract.
type ETHBridge struct {
	ETHBridgeCaller     // Read-only binding to the contract
	ETHBridgeTransactor // Write-only binding to the contract
	ETHBridgeFilterer   // Log filterer for contract events
}

// ETHBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ETHBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ETHBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ETHBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ETHBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ETHBridgeSession struct {
	Contract     *ETHBridge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ETHBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ETHBridgeCallerSession struct {
	Contract *ETHBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ETHBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ETHBridgeTransactorSession struct {
	Contract     *ETHBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ETHBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ETHBridgeRaw struct {
	Contract *ETHBridge // Generic contract binding to access the raw methods on
}

// ETHBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ETHBridgeCallerRaw struct {
	Contract *ETHBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ETHBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ETHBridgeTransactorRaw struct {
	Contract *ETHBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewETHBridge creates a new instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridge(address common.Address, backend bind.ContractBackend) (*ETHBridge, error) {
	contract, err := bindETHBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ETHBridge{ETHBridgeCaller: ETHBridgeCaller{contract: contract}, ETHBridgeTransactor: ETHBridgeTransactor{contract: contract}, ETHBridgeFilterer: ETHBridgeFilterer{contract: contract}}, nil
}

// NewETHBridgeCaller creates a new read-only instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridgeCaller(address common.Address, caller bind.ContractCaller) (*ETHBridgeCaller, error) {
	contract, err := bindETHBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeCaller{contract: contract}, nil
}

// NewETHBridgeTransactor creates a new write-only instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ETHBridgeTransactor, error) {
	contract, err := bindETHBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeTransactor{contract: contract}, nil
}

// NewETHBridgeFilterer creates a new log filterer instance of ETHBridge, bound to a specific deployed contract.
func NewETHBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ETHBridgeFilterer, error) {
	contract, err := bindETHBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeFilterer{contract: contract}, nil
}

// bindETHBridge binds a generic wrapper to an already deployed contract.
func bindETHBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ETHBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHBridge *ETHBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHBridge.Contract.ETHBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHBridge *ETHBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHBridge.Contract.ETHBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHBridge *ETHBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHBridge.Contract.ETHBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ETHBridge *ETHBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ETHBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ETHBridge *ETHBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ETHBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ETHBridge *ETHBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ETHBridge.Contract.contract.Transact(opts, method, params...)
}

// Mts is a free data retrieval call binding the contract method 0x7ed4f9a1.
//
// Solidity: function mts() view returns(address)
func (_ETHBridge *ETHBridgeCaller) Mts(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "mts")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mts is a free data retrieval call binding the contract method 0x7ed4f9a1.
//
// Solidity: function mts() view returns(address)
func (_ETHBridge *ETHBridgeSession) Mts() (common.Address, error) {
	return _ETHBridge.Contract.Mts(&_ETHBridge.CallOpts)
}

// Mts is a free data retrieval call binding the contract method 0x7ed4f9a1.
//
// Solidity: function mts() view returns(address)
func (_ETHBridge *ETHBridgeCallerSession) Mts() (common.Address, error) {
	return _ETHBridge.Contract.Mts(&_ETHBridge.CallOpts)
}

// MtsAmountMap is a free data retrieval call binding the contract method 0x80a25504.
//
// Solidity: function mtsAmountMap(address ) view returns(uint256)
func (_ETHBridge *ETHBridgeCaller) MtsAmountMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "mtsAmountMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MtsAmountMap is a free data retrieval call binding the contract method 0x80a25504.
//
// Solidity: function mtsAmountMap(address ) view returns(uint256)
func (_ETHBridge *ETHBridgeSession) MtsAmountMap(arg0 common.Address) (*big.Int, error) {
	return _ETHBridge.Contract.MtsAmountMap(&_ETHBridge.CallOpts, arg0)
}

// MtsAmountMap is a free data retrieval call binding the contract method 0x80a25504.
//
// Solidity: function mtsAmountMap(address ) view returns(uint256)
func (_ETHBridge *ETHBridgeCallerSession) MtsAmountMap(arg0 common.Address) (*big.Int, error) {
	return _ETHBridge.Contract.MtsAmountMap(&_ETHBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHBridge *ETHBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ETHBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHBridge *ETHBridgeSession) Owner() (common.Address, error) {
	return _ETHBridge.Contract.Owner(&_ETHBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ETHBridge *ETHBridgeCallerSession) Owner() (common.Address, error) {
	return _ETHBridge.Contract.Owner(&_ETHBridge.CallOpts)
}

// LockMts is a paid mutator transaction binding the contract method 0xb34fd5df.
//
// Solidity: function lockMts(uint256 _amount) returns()
func (_ETHBridge *ETHBridgeTransactor) LockMts(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "lockMts", _amount)
}

// LockMts is a paid mutator transaction binding the contract method 0xb34fd5df.
//
// Solidity: function lockMts(uint256 _amount) returns()
func (_ETHBridge *ETHBridgeSession) LockMts(_amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.LockMts(&_ETHBridge.TransactOpts, _amount)
}

// LockMts is a paid mutator transaction binding the contract method 0xb34fd5df.
//
// Solidity: function lockMts(uint256 _amount) returns()
func (_ETHBridge *ETHBridgeTransactorSession) LockMts(_amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.LockMts(&_ETHBridge.TransactOpts, _amount)
}

// SubMts is a paid mutator transaction binding the contract method 0xc5d56653.
//
// Solidity: function subMts(address account, uint256 amount) returns()
func (_ETHBridge *ETHBridgeTransactor) SubMts(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "subMts", account, amount)
}

// SubMts is a paid mutator transaction binding the contract method 0xc5d56653.
//
// Solidity: function subMts(address account, uint256 amount) returns()
func (_ETHBridge *ETHBridgeSession) SubMts(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.SubMts(&_ETHBridge.TransactOpts, account, amount)
}

// SubMts is a paid mutator transaction binding the contract method 0xc5d56653.
//
// Solidity: function subMts(address account, uint256 amount) returns()
func (_ETHBridge *ETHBridgeTransactorSession) SubMts(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.SubMts(&_ETHBridge.TransactOpts, account, amount)
}

// UnlockMts is a paid mutator transaction binding the contract method 0x41004b14.
//
// Solidity: function unlockMts(address _to, uint256 _amount) returns()
func (_ETHBridge *ETHBridgeTransactor) UnlockMts(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.contract.Transact(opts, "unlockMts", _to, _amount)
}

// UnlockMts is a paid mutator transaction binding the contract method 0x41004b14.
//
// Solidity: function unlockMts(address _to, uint256 _amount) returns()
func (_ETHBridge *ETHBridgeSession) UnlockMts(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.UnlockMts(&_ETHBridge.TransactOpts, _to, _amount)
}

// UnlockMts is a paid mutator transaction binding the contract method 0x41004b14.
//
// Solidity: function unlockMts(address _to, uint256 _amount) returns()
func (_ETHBridge *ETHBridgeTransactorSession) UnlockMts(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ETHBridge.Contract.UnlockMts(&_ETHBridge.TransactOpts, _to, _amount)
}

// ETHBridgeLockMtsIterator is returned from FilterLockMts and is used to iterate over the raw logs and unpacked data for LockMts events raised by the ETHBridge contract.
type ETHBridgeLockMtsIterator struct {
	Event *ETHBridgeLockMts // Event containing the contract specifics and raw log

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
func (it *ETHBridgeLockMtsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeLockMts)
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
		it.Event = new(ETHBridgeLockMts)
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
func (it *ETHBridgeLockMtsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeLockMtsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeLockMts represents a LockMts event raised by the ETHBridge contract.
type ETHBridgeLockMts struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLockMts is a free log retrieval operation binding the contract event 0x60745bb8f47da8fa85b2fa96982f4ca228359f2e8ab628004429eb0d67427c25.
//
// Solidity: event LockMts(address indexed account, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) FilterLockMts(opts *bind.FilterOpts, account []common.Address) (*ETHBridgeLockMtsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "LockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeLockMtsIterator{contract: _ETHBridge.contract, event: "LockMts", logs: logs, sub: sub}, nil
}

// WatchLockMts is a free log subscription operation binding the contract event 0x60745bb8f47da8fa85b2fa96982f4ca228359f2e8ab628004429eb0d67427c25.
//
// Solidity: event LockMts(address indexed account, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) WatchLockMts(opts *bind.WatchOpts, sink chan<- *ETHBridgeLockMts, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "LockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeLockMts)
				if err := _ETHBridge.contract.UnpackLog(event, "LockMts", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseLockMts(log types.Log) (*ETHBridgeLockMts, error) {
	event := new(ETHBridgeLockMts)
	if err := _ETHBridge.contract.UnpackLog(event, "LockMts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ETHBridgeUnlockMtsIterator is returned from FilterUnlockMts and is used to iterate over the raw logs and unpacked data for UnlockMts events raised by the ETHBridge contract.
type ETHBridgeUnlockMtsIterator struct {
	Event *ETHBridgeUnlockMts // Event containing the contract specifics and raw log

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
func (it *ETHBridgeUnlockMtsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ETHBridgeUnlockMts)
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
		it.Event = new(ETHBridgeUnlockMts)
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
func (it *ETHBridgeUnlockMtsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ETHBridgeUnlockMtsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ETHBridgeUnlockMts represents a UnlockMts event raised by the ETHBridge contract.
type ETHBridgeUnlockMts struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlockMts is a free log retrieval operation binding the contract event 0x8692c5be9f8792a83104866c4465016f4928e73a17f2240e5a7e70d1ab2af2a4.
//
// Solidity: event UnlockMts(address indexed account, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) FilterUnlockMts(opts *bind.FilterOpts, account []common.Address) (*ETHBridgeUnlockMtsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ETHBridge.contract.FilterLogs(opts, "UnlockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return &ETHBridgeUnlockMtsIterator{contract: _ETHBridge.contract, event: "UnlockMts", logs: logs, sub: sub}, nil
}

// WatchUnlockMts is a free log subscription operation binding the contract event 0x8692c5be9f8792a83104866c4465016f4928e73a17f2240e5a7e70d1ab2af2a4.
//
// Solidity: event UnlockMts(address indexed account, uint256 amount)
func (_ETHBridge *ETHBridgeFilterer) WatchUnlockMts(opts *bind.WatchOpts, sink chan<- *ETHBridgeUnlockMts, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ETHBridge.contract.WatchLogs(opts, "UnlockMts", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ETHBridgeUnlockMts)
				if err := _ETHBridge.contract.UnpackLog(event, "UnlockMts", log); err != nil {
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
func (_ETHBridge *ETHBridgeFilterer) ParseUnlockMts(log types.Log) (*ETHBridgeUnlockMts, error) {
	event := new(ETHBridgeUnlockMts)
	if err := _ETHBridge.contract.UnpackLog(event, "UnlockMts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
