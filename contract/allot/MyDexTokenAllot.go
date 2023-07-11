// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package allot

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

// AllotMetaData contains all meta data concerning the Allot contract.
var AllotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_myDexToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_mdtCount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"acount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"acount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"acount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"UserInfoMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"debtAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isUpdate\",\"type\":\"bool\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mdtCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myDexToken\",\"outputs\":[{\"internalType\":\"contractMyDexToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"}],\"name\":\"pendingMdt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"poolInfos\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allocPoint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRewardBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accShareAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAllocPoint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AllotABI is the input ABI used to generate the binding from.
// Deprecated: Use AllotMetaData.ABI instead.
var AllotABI = AllotMetaData.ABI

// Allot is an auto generated Go binding around an Ethereum contract.
type Allot struct {
	AllotCaller     // Read-only binding to the contract
	AllotTransactor // Write-only binding to the contract
	AllotFilterer   // Log filterer for contract events
}

// AllotCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllotSession struct {
	Contract     *Allot            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllotCallerSession struct {
	Contract *AllotCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AllotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllotTransactorSession struct {
	Contract     *AllotTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllotRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllotRaw struct {
	Contract *Allot // Generic contract binding to access the raw methods on
}

// AllotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllotCallerRaw struct {
	Contract *AllotCaller // Generic read-only contract binding to access the raw methods on
}

// AllotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllotTransactorRaw struct {
	Contract *AllotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllot creates a new instance of Allot, bound to a specific deployed contract.
func NewAllot(address common.Address, backend bind.ContractBackend) (*Allot, error) {
	contract, err := bindAllot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Allot{AllotCaller: AllotCaller{contract: contract}, AllotTransactor: AllotTransactor{contract: contract}, AllotFilterer: AllotFilterer{contract: contract}}, nil
}

// NewAllotCaller creates a new read-only instance of Allot, bound to a specific deployed contract.
func NewAllotCaller(address common.Address, caller bind.ContractCaller) (*AllotCaller, error) {
	contract, err := bindAllot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllotCaller{contract: contract}, nil
}

// NewAllotTransactor creates a new write-only instance of Allot, bound to a specific deployed contract.
func NewAllotTransactor(address common.Address, transactor bind.ContractTransactor) (*AllotTransactor, error) {
	contract, err := bindAllot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllotTransactor{contract: contract}, nil
}

// NewAllotFilterer creates a new log filterer instance of Allot, bound to a specific deployed contract.
func NewAllotFilterer(address common.Address, filterer bind.ContractFilterer) (*AllotFilterer, error) {
	contract, err := bindAllot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllotFilterer{contract: contract}, nil
}

// bindAllot binds a generic wrapper to an already deployed contract.
func bindAllot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AllotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Allot *AllotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Allot.Contract.AllotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Allot *AllotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Allot.Contract.AllotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Allot *AllotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Allot.Contract.AllotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Allot *AllotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Allot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Allot *AllotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Allot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Allot *AllotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Allot.Contract.contract.Transact(opts, method, params...)
}

// UserInfoMap is a free data retrieval call binding the contract method 0x3bea633a.
//
// Solidity: function UserInfoMap(uint256 , address ) view returns(address userAddress, uint256 lpAmount, uint256 debtAmount)
func (_Allot *AllotCaller) UserInfoMap(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	UserAddress common.Address
	LpAmount    *big.Int
	DebtAmount  *big.Int
}, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "UserInfoMap", arg0, arg1)

	outstruct := new(struct {
		UserAddress common.Address
		LpAmount    *big.Int
		DebtAmount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LpAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DebtAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfoMap is a free data retrieval call binding the contract method 0x3bea633a.
//
// Solidity: function UserInfoMap(uint256 , address ) view returns(address userAddress, uint256 lpAmount, uint256 debtAmount)
func (_Allot *AllotSession) UserInfoMap(arg0 *big.Int, arg1 common.Address) (struct {
	UserAddress common.Address
	LpAmount    *big.Int
	DebtAmount  *big.Int
}, error) {
	return _Allot.Contract.UserInfoMap(&_Allot.CallOpts, arg0, arg1)
}

// UserInfoMap is a free data retrieval call binding the contract method 0x3bea633a.
//
// Solidity: function UserInfoMap(uint256 , address ) view returns(address userAddress, uint256 lpAmount, uint256 debtAmount)
func (_Allot *AllotCallerSession) UserInfoMap(arg0 *big.Int, arg1 common.Address) (struct {
	UserAddress common.Address
	LpAmount    *big.Int
	DebtAmount  *big.Int
}, error) {
	return _Allot.Contract.UserInfoMap(&_Allot.CallOpts, arg0, arg1)
}

// MdtCount is a free data retrieval call binding the contract method 0x2e56791e.
//
// Solidity: function mdtCount() view returns(uint256)
func (_Allot *AllotCaller) MdtCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "mdtCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MdtCount is a free data retrieval call binding the contract method 0x2e56791e.
//
// Solidity: function mdtCount() view returns(uint256)
func (_Allot *AllotSession) MdtCount() (*big.Int, error) {
	return _Allot.Contract.MdtCount(&_Allot.CallOpts)
}

// MdtCount is a free data retrieval call binding the contract method 0x2e56791e.
//
// Solidity: function mdtCount() view returns(uint256)
func (_Allot *AllotCallerSession) MdtCount() (*big.Int, error) {
	return _Allot.Contract.MdtCount(&_Allot.CallOpts)
}

// MyDexToken is a free data retrieval call binding the contract method 0x9a967d47.
//
// Solidity: function myDexToken() view returns(address)
func (_Allot *AllotCaller) MyDexToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "myDexToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MyDexToken is a free data retrieval call binding the contract method 0x9a967d47.
//
// Solidity: function myDexToken() view returns(address)
func (_Allot *AllotSession) MyDexToken() (common.Address, error) {
	return _Allot.Contract.MyDexToken(&_Allot.CallOpts)
}

// MyDexToken is a free data retrieval call binding the contract method 0x9a967d47.
//
// Solidity: function myDexToken() view returns(address)
func (_Allot *AllotCallerSession) MyDexToken() (common.Address, error) {
	return _Allot.Contract.MyDexToken(&_Allot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Allot *AllotCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Allot *AllotSession) Owner() (common.Address, error) {
	return _Allot.Contract.Owner(&_Allot.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Allot *AllotCallerSession) Owner() (common.Address, error) {
	return _Allot.Contract.Owner(&_Allot.CallOpts)
}

// PendingMdt is a free data retrieval call binding the contract method 0x31731309.
//
// Solidity: function pendingMdt(address _account, uint256 _pid) view returns(uint256 pendingAmount)
func (_Allot *AllotCaller) PendingMdt(opts *bind.CallOpts, _account common.Address, _pid *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "pendingMdt", _account, _pid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingMdt is a free data retrieval call binding the contract method 0x31731309.
//
// Solidity: function pendingMdt(address _account, uint256 _pid) view returns(uint256 pendingAmount)
func (_Allot *AllotSession) PendingMdt(_account common.Address, _pid *big.Int) (*big.Int, error) {
	return _Allot.Contract.PendingMdt(&_Allot.CallOpts, _account, _pid)
}

// PendingMdt is a free data retrieval call binding the contract method 0x31731309.
//
// Solidity: function pendingMdt(address _account, uint256 _pid) view returns(uint256 pendingAmount)
func (_Allot *AllotCallerSession) PendingMdt(_account common.Address, _pid *big.Int) (*big.Int, error) {
	return _Allot.Contract.PendingMdt(&_Allot.CallOpts, _account, _pid)
}

// PoolInfos is a free data retrieval call binding the contract method 0x689d84e4.
//
// Solidity: function poolInfos(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accShareAmount)
func (_Allot *AllotCaller) PoolInfos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccShareAmount  *big.Int
}, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "poolInfos", arg0)

	outstruct := new(struct {
		LpToken         common.Address
		AllocPoint      *big.Int
		LastRewardBlock *big.Int
		AccShareAmount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.AllocPoint = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastRewardBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AccShareAmount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PoolInfos is a free data retrieval call binding the contract method 0x689d84e4.
//
// Solidity: function poolInfos(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accShareAmount)
func (_Allot *AllotSession) PoolInfos(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccShareAmount  *big.Int
}, error) {
	return _Allot.Contract.PoolInfos(&_Allot.CallOpts, arg0)
}

// PoolInfos is a free data retrieval call binding the contract method 0x689d84e4.
//
// Solidity: function poolInfos(uint256 ) view returns(address lpToken, uint256 allocPoint, uint256 lastRewardBlock, uint256 accShareAmount)
func (_Allot *AllotCallerSession) PoolInfos(arg0 *big.Int) (struct {
	LpToken         common.Address
	AllocPoint      *big.Int
	LastRewardBlock *big.Int
	AccShareAmount  *big.Int
}, error) {
	return _Allot.Contract.PoolInfos(&_Allot.CallOpts, arg0)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Allot *AllotCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Allot *AllotSession) StartBlock() (*big.Int, error) {
	return _Allot.Contract.StartBlock(&_Allot.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256)
func (_Allot *AllotCallerSession) StartBlock() (*big.Int, error) {
	return _Allot.Contract.StartBlock(&_Allot.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Allot *AllotCaller) TotalAllocPoint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Allot.contract.Call(opts, &out, "totalAllocPoint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Allot *AllotSession) TotalAllocPoint() (*big.Int, error) {
	return _Allot.Contract.TotalAllocPoint(&_Allot.CallOpts)
}

// TotalAllocPoint is a free data retrieval call binding the contract method 0x17caf6f1.
//
// Solidity: function totalAllocPoint() view returns(uint256)
func (_Allot *AllotCallerSession) TotalAllocPoint() (*big.Int, error) {
	return _Allot.Contract.TotalAllocPoint(&_Allot.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x0dec2312.
//
// Solidity: function add(address _lpToken, uint256 _allocPoint, bool _isUpdate) returns()
func (_Allot *AllotTransactor) Add(opts *bind.TransactOpts, _lpToken common.Address, _allocPoint *big.Int, _isUpdate bool) (*types.Transaction, error) {
	return _Allot.contract.Transact(opts, "add", _lpToken, _allocPoint, _isUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x0dec2312.
//
// Solidity: function add(address _lpToken, uint256 _allocPoint, bool _isUpdate) returns()
func (_Allot *AllotSession) Add(_lpToken common.Address, _allocPoint *big.Int, _isUpdate bool) (*types.Transaction, error) {
	return _Allot.Contract.Add(&_Allot.TransactOpts, _lpToken, _allocPoint, _isUpdate)
}

// Add is a paid mutator transaction binding the contract method 0x0dec2312.
//
// Solidity: function add(address _lpToken, uint256 _allocPoint, bool _isUpdate) returns()
func (_Allot *AllotTransactorSession) Add(_lpToken common.Address, _allocPoint *big.Int, _isUpdate bool) (*types.Transaction, error) {
	return _Allot.Contract.Add(&_Allot.TransactOpts, _lpToken, _allocPoint, _isUpdate)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Allot *AllotTransactor) Deposit(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Allot.contract.Transact(opts, "deposit", _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Allot *AllotSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Allot.Contract.Deposit(&_Allot.TransactOpts, _pid, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xe2bbb158.
//
// Solidity: function deposit(uint256 _pid, uint256 _amount) returns()
func (_Allot *AllotTransactorSession) Deposit(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Allot.Contract.Deposit(&_Allot.TransactOpts, _pid, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Allot *AllotTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Allot.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Allot *AllotSession) RenounceOwnership() (*types.Transaction, error) {
	return _Allot.Contract.RenounceOwnership(&_Allot.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Allot *AllotTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Allot.Contract.RenounceOwnership(&_Allot.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Allot *AllotTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Allot.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Allot *AllotSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Allot.Contract.TransferOwnership(&_Allot.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Allot *AllotTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Allot.Contract.TransferOwnership(&_Allot.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Allot *AllotTransactor) Withdraw(opts *bind.TransactOpts, _pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Allot.contract.Transact(opts, "withdraw", _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Allot *AllotSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Allot.Contract.Withdraw(&_Allot.TransactOpts, _pid, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x441a3e70.
//
// Solidity: function withdraw(uint256 _pid, uint256 _amount) returns()
func (_Allot *AllotTransactorSession) Withdraw(_pid *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Allot.Contract.Withdraw(&_Allot.TransactOpts, _pid, _amount)
}

// AllotAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the Allot contract.
type AllotAddIterator struct {
	Event *AllotAdd // Event containing the contract specifics and raw log

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
func (it *AllotAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllotAdd)
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
		it.Event = new(AllotAdd)
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
func (it *AllotAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllotAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllotAdd represents a Add event raised by the Allot contract.
type AllotAdd struct {
	Acount     common.Address
	LpToken    common.Address
	AllocPoint *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0x5f43b6be860b00b28d44981331005a0503e6b566155ca1423af4299a6bde1f46.
//
// Solidity: event Add(address indexed acount, address indexed lpToken, uint256 allocPoint)
func (_Allot *AllotFilterer) FilterAdd(opts *bind.FilterOpts, acount []common.Address, lpToken []common.Address) (*AllotAddIterator, error) {

	var acountRule []interface{}
	for _, acountItem := range acount {
		acountRule = append(acountRule, acountItem)
	}
	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _Allot.contract.FilterLogs(opts, "Add", acountRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return &AllotAddIterator{contract: _Allot.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0x5f43b6be860b00b28d44981331005a0503e6b566155ca1423af4299a6bde1f46.
//
// Solidity: event Add(address indexed acount, address indexed lpToken, uint256 allocPoint)
func (_Allot *AllotFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *AllotAdd, acount []common.Address, lpToken []common.Address) (event.Subscription, error) {

	var acountRule []interface{}
	for _, acountItem := range acount {
		acountRule = append(acountRule, acountItem)
	}
	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _Allot.contract.WatchLogs(opts, "Add", acountRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllotAdd)
				if err := _Allot.contract.UnpackLog(event, "Add", log); err != nil {
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

// ParseAdd is a log parse operation binding the contract event 0x5f43b6be860b00b28d44981331005a0503e6b566155ca1423af4299a6bde1f46.
//
// Solidity: event Add(address indexed acount, address indexed lpToken, uint256 allocPoint)
func (_Allot *AllotFilterer) ParseAdd(log types.Log) (*AllotAdd, error) {
	event := new(AllotAdd)
	if err := _Allot.contract.UnpackLog(event, "Add", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllotDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Allot contract.
type AllotDepositIterator struct {
	Event *AllotDeposit // Event containing the contract specifics and raw log

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
func (it *AllotDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllotDeposit)
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
		it.Event = new(AllotDeposit)
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
func (it *AllotDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllotDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllotDeposit represents a Deposit event raised by the Allot contract.
type AllotDeposit struct {
	Acount  common.Address
	LpToken common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed acount, address indexed lpToken, uint256 amount)
func (_Allot *AllotFilterer) FilterDeposit(opts *bind.FilterOpts, acount []common.Address, lpToken []common.Address) (*AllotDepositIterator, error) {

	var acountRule []interface{}
	for _, acountItem := range acount {
		acountRule = append(acountRule, acountItem)
	}
	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _Allot.contract.FilterLogs(opts, "Deposit", acountRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return &AllotDepositIterator{contract: _Allot.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed acount, address indexed lpToken, uint256 amount)
func (_Allot *AllotFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *AllotDeposit, acount []common.Address, lpToken []common.Address) (event.Subscription, error) {

	var acountRule []interface{}
	for _, acountItem := range acount {
		acountRule = append(acountRule, acountItem)
	}
	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _Allot.contract.WatchLogs(opts, "Deposit", acountRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllotDeposit)
				if err := _Allot.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5548c837ab068cf56a2c2479df0882a4922fd203edb7517321831d95078c5f62.
//
// Solidity: event Deposit(address indexed acount, address indexed lpToken, uint256 amount)
func (_Allot *AllotFilterer) ParseDeposit(log types.Log) (*AllotDeposit, error) {
	event := new(AllotDeposit)
	if err := _Allot.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllotOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Allot contract.
type AllotOwnershipTransferredIterator struct {
	Event *AllotOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AllotOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllotOwnershipTransferred)
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
		it.Event = new(AllotOwnershipTransferred)
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
func (it *AllotOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllotOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllotOwnershipTransferred represents a OwnershipTransferred event raised by the Allot contract.
type AllotOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Allot *AllotFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AllotOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Allot.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AllotOwnershipTransferredIterator{contract: _Allot.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Allot *AllotFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AllotOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Allot.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllotOwnershipTransferred)
				if err := _Allot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Allot *AllotFilterer) ParseOwnershipTransferred(log types.Log) (*AllotOwnershipTransferred, error) {
	event := new(AllotOwnershipTransferred)
	if err := _Allot.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AllotWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Allot contract.
type AllotWithdrawIterator struct {
	Event *AllotWithdraw // Event containing the contract specifics and raw log

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
func (it *AllotWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllotWithdraw)
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
		it.Event = new(AllotWithdraw)
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
func (it *AllotWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllotWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllotWithdraw represents a Withdraw event raised by the Allot contract.
type AllotWithdraw struct {
	Acount  common.Address
	LpToken common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed acount, address indexed lpToken, uint256 amount)
func (_Allot *AllotFilterer) FilterWithdraw(opts *bind.FilterOpts, acount []common.Address, lpToken []common.Address) (*AllotWithdrawIterator, error) {

	var acountRule []interface{}
	for _, acountItem := range acount {
		acountRule = append(acountRule, acountItem)
	}
	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _Allot.contract.FilterLogs(opts, "Withdraw", acountRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return &AllotWithdrawIterator{contract: _Allot.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed acount, address indexed lpToken, uint256 amount)
func (_Allot *AllotFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AllotWithdraw, acount []common.Address, lpToken []common.Address) (event.Subscription, error) {

	var acountRule []interface{}
	for _, acountItem := range acount {
		acountRule = append(acountRule, acountItem)
	}
	var lpTokenRule []interface{}
	for _, lpTokenItem := range lpToken {
		lpTokenRule = append(lpTokenRule, lpTokenItem)
	}

	logs, sub, err := _Allot.contract.WatchLogs(opts, "Withdraw", acountRule, lpTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllotWithdraw)
				if err := _Allot.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address indexed acount, address indexed lpToken, uint256 amount)
func (_Allot *AllotFilterer) ParseWithdraw(log types.Log) (*AllotWithdraw, error) {
	event := new(AllotWithdraw)
	if err := _Allot.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
