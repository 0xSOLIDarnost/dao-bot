// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisigwallet

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
)

// MultisigwalletMetaData contains all meta data concerning the Multisigwallet contract.
var MultisigwalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dailyLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcMaxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"calculateData2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data2\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dailyLimit\",\"type\":\"uint256\"}],\"name\":\"changeDailyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dailyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"isTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"spentToday\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MultisigwalletABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisigwalletMetaData.ABI instead.
var MultisigwalletABI = MultisigwalletMetaData.ABI

// Multisigwallet is an auto generated Go binding around an Ethereum contract.
type Multisigwallet struct {
	MultisigwalletCaller     // Read-only binding to the contract
	MultisigwalletTransactor // Write-only binding to the contract
	MultisigwalletFilterer   // Log filterer for contract events
}

// MultisigwalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisigwalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigwalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisigwalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigwalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisigwalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisigwalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisigwalletSession struct {
	Contract     *Multisigwallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisigwalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisigwalletCallerSession struct {
	Contract *MultisigwalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MultisigwalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisigwalletTransactorSession struct {
	Contract     *MultisigwalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MultisigwalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisigwalletRaw struct {
	Contract *Multisigwallet // Generic contract binding to access the raw methods on
}

// MultisigwalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisigwalletCallerRaw struct {
	Contract *MultisigwalletCaller // Generic read-only contract binding to access the raw methods on
}

// MultisigwalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisigwalletTransactorRaw struct {
	Contract *MultisigwalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisigwallet creates a new instance of Multisigwallet, bound to a specific deployed contract.
func NewMultisigwallet(address common.Address, backend bind.ContractBackend) (*Multisigwallet, error) {
	contract, err := bindMultisigwallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisigwallet{MultisigwalletCaller: MultisigwalletCaller{contract: contract}, MultisigwalletTransactor: MultisigwalletTransactor{contract: contract}, MultisigwalletFilterer: MultisigwalletFilterer{contract: contract}}, nil
}

// NewMultisigwalletCaller creates a new read-only instance of Multisigwallet, bound to a specific deployed contract.
func NewMultisigwalletCaller(address common.Address, caller bind.ContractCaller) (*MultisigwalletCaller, error) {
	contract, err := bindMultisigwallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigwalletCaller{contract: contract}, nil
}

// NewMultisigwalletTransactor creates a new write-only instance of Multisigwallet, bound to a specific deployed contract.
func NewMultisigwalletTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisigwalletTransactor, error) {
	contract, err := bindMultisigwallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisigwalletTransactor{contract: contract}, nil
}

// NewMultisigwalletFilterer creates a new log filterer instance of Multisigwallet, bound to a specific deployed contract.
func NewMultisigwalletFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisigwalletFilterer, error) {
	contract, err := bindMultisigwallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisigwalletFilterer{contract: contract}, nil
}

// bindMultisigwallet binds a generic wrapper to an already deployed contract.
func bindMultisigwallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisigwalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisigwallet *MultisigwalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisigwallet.Contract.MultisigwalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisigwallet *MultisigwalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisigwallet.Contract.MultisigwalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisigwallet *MultisigwalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisigwallet.Contract.MultisigwalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisigwallet *MultisigwalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisigwallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisigwallet *MultisigwalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisigwallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisigwallet *MultisigwalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisigwallet.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_Multisigwallet *MultisigwalletCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "MAX_OWNER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_Multisigwallet *MultisigwalletSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _Multisigwallet.Contract.MAXOWNERCOUNT(&_Multisigwallet.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_Multisigwallet *MultisigwalletCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _Multisigwallet.Contract.MAXOWNERCOUNT(&_Multisigwallet.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_Multisigwallet *MultisigwalletCaller) CalcMaxWithdraw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "calcMaxWithdraw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_Multisigwallet *MultisigwalletSession) CalcMaxWithdraw() (*big.Int, error) {
	return _Multisigwallet.Contract.CalcMaxWithdraw(&_Multisigwallet.CallOpts)
}

// CalcMaxWithdraw is a free data retrieval call binding the contract method 0x4bc9fdc2.
//
// Solidity: function calcMaxWithdraw() view returns(uint256)
func (_Multisigwallet *MultisigwalletCallerSession) CalcMaxWithdraw() (*big.Int, error) {
	return _Multisigwallet.Contract.CalcMaxWithdraw(&_Multisigwallet.CallOpts)
}

// CalculateData2 is a free data retrieval call binding the contract method 0x1a2f947d.
//
// Solidity: function calculateData2(bytes data) pure returns(bytes data2, uint256 _fee)
func (_Multisigwallet *MultisigwalletCaller) CalculateData2(opts *bind.CallOpts, data []byte) (struct {
	Data2 []byte
	Fee   *big.Int
}, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "calculateData2", data)

	outstruct := new(struct {
		Data2 []byte
		Fee   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Data2 = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Fee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateData2 is a free data retrieval call binding the contract method 0x1a2f947d.
//
// Solidity: function calculateData2(bytes data) pure returns(bytes data2, uint256 _fee)
func (_Multisigwallet *MultisigwalletSession) CalculateData2(data []byte) (struct {
	Data2 []byte
	Fee   *big.Int
}, error) {
	return _Multisigwallet.Contract.CalculateData2(&_Multisigwallet.CallOpts, data)
}

// CalculateData2 is a free data retrieval call binding the contract method 0x1a2f947d.
//
// Solidity: function calculateData2(bytes data) pure returns(bytes data2, uint256 _fee)
func (_Multisigwallet *MultisigwalletCallerSession) CalculateData2(data []byte) (struct {
	Data2 []byte
	Fee   *big.Int
}, error) {
	return _Multisigwallet.Contract.CalculateData2(&_Multisigwallet.CallOpts, data)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_Multisigwallet *MultisigwalletCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "confirmations", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_Multisigwallet *MultisigwalletSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Multisigwallet.Contract.Confirmations(&_Multisigwallet.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_Multisigwallet *MultisigwalletCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _Multisigwallet.Contract.Confirmations(&_Multisigwallet.CallOpts, arg0, arg1)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_Multisigwallet *MultisigwalletCaller) DailyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "dailyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_Multisigwallet *MultisigwalletSession) DailyLimit() (*big.Int, error) {
	return _Multisigwallet.Contract.DailyLimit(&_Multisigwallet.CallOpts)
}

// DailyLimit is a free data retrieval call binding the contract method 0x67eeba0c.
//
// Solidity: function dailyLimit() view returns(uint256)
func (_Multisigwallet *MultisigwalletCallerSession) DailyLimit() (*big.Int, error) {
	return _Multisigwallet.Contract.DailyLimit(&_Multisigwallet.CallOpts)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_Multisigwallet *MultisigwalletCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "getConfirmationCount", transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_Multisigwallet *MultisigwalletSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _Multisigwallet.Contract.GetConfirmationCount(&_Multisigwallet.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_Multisigwallet *MultisigwalletCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _Multisigwallet.Contract.GetConfirmationCount(&_Multisigwallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_Multisigwallet *MultisigwalletCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "getConfirmations", transactionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_Multisigwallet *MultisigwalletSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _Multisigwallet.Contract.GetConfirmations(&_Multisigwallet.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_Multisigwallet *MultisigwalletCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _Multisigwallet.Contract.GetConfirmations(&_Multisigwallet.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisigwallet *MultisigwalletCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisigwallet *MultisigwalletSession) GetOwners() ([]common.Address, error) {
	return _Multisigwallet.Contract.GetOwners(&_Multisigwallet.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_Multisigwallet *MultisigwalletCallerSession) GetOwners() ([]common.Address, error) {
	return _Multisigwallet.Contract.GetOwners(&_Multisigwallet.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_Multisigwallet *MultisigwalletCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "getTransactionCount", pending, executed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_Multisigwallet *MultisigwalletSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _Multisigwallet.Contract.GetTransactionCount(&_Multisigwallet.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_Multisigwallet *MultisigwalletCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _Multisigwallet.Contract.GetTransactionCount(&_Multisigwallet.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_Multisigwallet *MultisigwalletCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "getTransactionIds", from, to, pending, executed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_Multisigwallet *MultisigwalletSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _Multisigwallet.Contract.GetTransactionIds(&_Multisigwallet.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_Multisigwallet *MultisigwalletCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _Multisigwallet.Contract.GetTransactionIds(&_Multisigwallet.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_Multisigwallet *MultisigwalletCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "isConfirmed", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_Multisigwallet *MultisigwalletSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _Multisigwallet.Contract.IsConfirmed(&_Multisigwallet.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_Multisigwallet *MultisigwalletCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _Multisigwallet.Contract.IsConfirmed(&_Multisigwallet.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisigwallet *MultisigwalletCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisigwallet *MultisigwalletSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Multisigwallet.Contract.IsOwner(&_Multisigwallet.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_Multisigwallet *MultisigwalletCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _Multisigwallet.Contract.IsOwner(&_Multisigwallet.CallOpts, arg0)
}

// IsTransfer is a free data retrieval call binding the contract method 0xaf8a08d9.
//
// Solidity: function isTransfer(bytes data) pure returns(bool)
func (_Multisigwallet *MultisigwalletCaller) IsTransfer(opts *bind.CallOpts, data []byte) (bool, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "isTransfer", data)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTransfer is a free data retrieval call binding the contract method 0xaf8a08d9.
//
// Solidity: function isTransfer(bytes data) pure returns(bool)
func (_Multisigwallet *MultisigwalletSession) IsTransfer(data []byte) (bool, error) {
	return _Multisigwallet.Contract.IsTransfer(&_Multisigwallet.CallOpts, data)
}

// IsTransfer is a free data retrieval call binding the contract method 0xaf8a08d9.
//
// Solidity: function isTransfer(bytes data) pure returns(bool)
func (_Multisigwallet *MultisigwalletCallerSession) IsTransfer(data []byte) (bool, error) {
	return _Multisigwallet.Contract.IsTransfer(&_Multisigwallet.CallOpts, data)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_Multisigwallet *MultisigwalletCaller) LastDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "lastDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_Multisigwallet *MultisigwalletSession) LastDay() (*big.Int, error) {
	return _Multisigwallet.Contract.LastDay(&_Multisigwallet.CallOpts)
}

// LastDay is a free data retrieval call binding the contract method 0x6b0c932d.
//
// Solidity: function lastDay() view returns(uint256)
func (_Multisigwallet *MultisigwalletCallerSession) LastDay() (*big.Int, error) {
	return _Multisigwallet.Contract.LastDay(&_Multisigwallet.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisigwallet *MultisigwalletCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisigwallet *MultisigwalletSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Multisigwallet.Contract.Owners(&_Multisigwallet.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_Multisigwallet *MultisigwalletCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _Multisigwallet.Contract.Owners(&_Multisigwallet.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_Multisigwallet *MultisigwalletCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "required")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_Multisigwallet *MultisigwalletSession) Required() (*big.Int, error) {
	return _Multisigwallet.Contract.Required(&_Multisigwallet.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_Multisigwallet *MultisigwalletCallerSession) Required() (*big.Int, error) {
	return _Multisigwallet.Contract.Required(&_Multisigwallet.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_Multisigwallet *MultisigwalletCaller) SpentToday(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "spentToday")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_Multisigwallet *MultisigwalletSession) SpentToday() (*big.Int, error) {
	return _Multisigwallet.Contract.SpentToday(&_Multisigwallet.CallOpts)
}

// SpentToday is a free data retrieval call binding the contract method 0xf059cf2b.
//
// Solidity: function spentToday() view returns(uint256)
func (_Multisigwallet *MultisigwalletCallerSession) SpentToday() (*big.Int, error) {
	return _Multisigwallet.Contract.SpentToday(&_Multisigwallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_Multisigwallet *MultisigwalletCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_Multisigwallet *MultisigwalletSession) TransactionCount() (*big.Int, error) {
	return _Multisigwallet.Contract.TransactionCount(&_Multisigwallet.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_Multisigwallet *MultisigwalletCallerSession) TransactionCount() (*big.Int, error) {
	return _Multisigwallet.Contract.TransactionCount(&_Multisigwallet.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_Multisigwallet *MultisigwalletCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	var out []interface{}
	err := _Multisigwallet.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Destination = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_Multisigwallet *MultisigwalletSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _Multisigwallet.Contract.Transactions(&_Multisigwallet.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_Multisigwallet *MultisigwalletCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _Multisigwallet.Contract.Transactions(&_Multisigwallet.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_Multisigwallet *MultisigwalletTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_Multisigwallet *MultisigwalletSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.AddOwner(&_Multisigwallet.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.AddOwner(&_Multisigwallet.TransactOpts, owner)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Multisigwallet *MultisigwalletTransactor) ChangeAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "changeAdmin", newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Multisigwallet *MultisigwalletSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ChangeAdmin(&_Multisigwallet.TransactOpts, newAdmin)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(address newAdmin) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) ChangeAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ChangeAdmin(&_Multisigwallet.TransactOpts, newAdmin)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_Multisigwallet *MultisigwalletTransactor) ChangeDailyLimit(opts *bind.TransactOpts, _dailyLimit *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "changeDailyLimit", _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_Multisigwallet *MultisigwalletSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ChangeDailyLimit(&_Multisigwallet.TransactOpts, _dailyLimit)
}

// ChangeDailyLimit is a paid mutator transaction binding the contract method 0xcea08621.
//
// Solidity: function changeDailyLimit(uint256 _dailyLimit) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) ChangeDailyLimit(_dailyLimit *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ChangeDailyLimit(&_Multisigwallet.TransactOpts, _dailyLimit)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_Multisigwallet *MultisigwalletTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_Multisigwallet *MultisigwalletSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ChangeRequirement(&_Multisigwallet.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ChangeRequirement(&_Multisigwallet.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ConfirmTransaction(&_Multisigwallet.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ConfirmTransaction(&_Multisigwallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ExecuteTransaction(&_Multisigwallet.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ExecuteTransaction(&_Multisigwallet.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_Multisigwallet *MultisigwalletTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_Multisigwallet *MultisigwalletSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.RemoveOwner(&_Multisigwallet.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.RemoveOwner(&_Multisigwallet.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_Multisigwallet *MultisigwalletTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_Multisigwallet *MultisigwalletSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ReplaceOwner(&_Multisigwallet.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _Multisigwallet.Contract.ReplaceOwner(&_Multisigwallet.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.RevokeConfirmation(&_Multisigwallet.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_Multisigwallet *MultisigwalletTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _Multisigwallet.Contract.RevokeConfirmation(&_Multisigwallet.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_Multisigwallet *MultisigwalletTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Multisigwallet.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_Multisigwallet *MultisigwalletSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Multisigwallet.Contract.SubmitTransaction(&_Multisigwallet.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_Multisigwallet *MultisigwalletTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Multisigwallet.Contract.SubmitTransaction(&_Multisigwallet.TransactOpts, destination, value, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisigwallet *MultisigwalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisigwallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisigwallet *MultisigwalletSession) Receive() (*types.Transaction, error) {
	return _Multisigwallet.Contract.Receive(&_Multisigwallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Multisigwallet *MultisigwalletTransactorSession) Receive() (*types.Transaction, error) {
	return _Multisigwallet.Contract.Receive(&_Multisigwallet.TransactOpts)
}

// MultisigwalletConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the Multisigwallet contract.
type MultisigwalletConfirmationIterator struct {
	Event *MultisigwalletConfirmation // Event containing the contract specifics and raw log

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
func (it *MultisigwalletConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletConfirmation)
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
		it.Event = new(MultisigwalletConfirmation)
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
func (it *MultisigwalletConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletConfirmation represents a Confirmation event raised by the Multisigwallet contract.
type MultisigwalletConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address sender, uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) FilterConfirmation(opts *bind.FilterOpts) (*MultisigwalletConfirmationIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletConfirmationIterator{contract: _Multisigwallet.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address sender, uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultisigwalletConfirmation) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "Confirmation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletConfirmation)
				if err := _Multisigwallet.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// ParseConfirmation is a log parse operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address sender, uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) ParseConfirmation(log types.Log) (*MultisigwalletConfirmation, error) {
	event := new(MultisigwalletConfirmation)
	if err := _Multisigwallet.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletDailyLimitChangeIterator is returned from FilterDailyLimitChange and is used to iterate over the raw logs and unpacked data for DailyLimitChange events raised by the Multisigwallet contract.
type MultisigwalletDailyLimitChangeIterator struct {
	Event *MultisigwalletDailyLimitChange // Event containing the contract specifics and raw log

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
func (it *MultisigwalletDailyLimitChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletDailyLimitChange)
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
		it.Event = new(MultisigwalletDailyLimitChange)
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
func (it *MultisigwalletDailyLimitChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletDailyLimitChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletDailyLimitChange represents a DailyLimitChange event raised by the Multisigwallet contract.
type MultisigwalletDailyLimitChange struct {
	DailyLimit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitChange is a free log retrieval operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_Multisigwallet *MultisigwalletFilterer) FilterDailyLimitChange(opts *bind.FilterOpts) (*MultisigwalletDailyLimitChangeIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletDailyLimitChangeIterator{contract: _Multisigwallet.contract, event: "DailyLimitChange", logs: logs, sub: sub}, nil
}

// WatchDailyLimitChange is a free log subscription operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_Multisigwallet *MultisigwalletFilterer) WatchDailyLimitChange(opts *bind.WatchOpts, sink chan<- *MultisigwalletDailyLimitChange) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "DailyLimitChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletDailyLimitChange)
				if err := _Multisigwallet.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
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

// ParseDailyLimitChange is a log parse operation binding the contract event 0xc71bdc6afaf9b1aa90a7078191d4fc1adf3bf680fca3183697df6b0dc226bca2.
//
// Solidity: event DailyLimitChange(uint256 dailyLimit)
func (_Multisigwallet *MultisigwalletFilterer) ParseDailyLimitChange(log types.Log) (*MultisigwalletDailyLimitChange, error) {
	event := new(MultisigwalletDailyLimitChange)
	if err := _Multisigwallet.contract.UnpackLog(event, "DailyLimitChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Multisigwallet contract.
type MultisigwalletDepositIterator struct {
	Event *MultisigwalletDeposit // Event containing the contract specifics and raw log

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
func (it *MultisigwalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletDeposit)
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
		it.Event = new(MultisigwalletDeposit)
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
func (it *MultisigwalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletDeposit represents a Deposit event raised by the Multisigwallet contract.
type MultisigwalletDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 value)
func (_Multisigwallet *MultisigwalletFilterer) FilterDeposit(opts *bind.FilterOpts) (*MultisigwalletDepositIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletDepositIterator{contract: _Multisigwallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 value)
func (_Multisigwallet *MultisigwalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultisigwalletDeposit) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletDeposit)
				if err := _Multisigwallet.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address sender, uint256 value)
func (_Multisigwallet *MultisigwalletFilterer) ParseDeposit(log types.Log) (*MultisigwalletDeposit, error) {
	event := new(MultisigwalletDeposit)
	if err := _Multisigwallet.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the Multisigwallet contract.
type MultisigwalletExecutionIterator struct {
	Event *MultisigwalletExecution // Event containing the contract specifics and raw log

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
func (it *MultisigwalletExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletExecution)
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
		it.Event = new(MultisigwalletExecution)
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
func (it *MultisigwalletExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletExecution represents a Execution event raised by the Multisigwallet contract.
type MultisigwalletExecution struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) FilterExecution(opts *bind.FilterOpts) (*MultisigwalletExecutionIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "Execution")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletExecutionIterator{contract: _Multisigwallet.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultisigwalletExecution) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "Execution")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletExecution)
				if err := _Multisigwallet.contract.UnpackLog(event, "Execution", log); err != nil {
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

// ParseExecution is a log parse operation binding the contract event 0x33e13ecb54c3076d8e8bb8c2881800a4d972b792045ffae98fdf46df365fed75.
//
// Solidity: event Execution(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) ParseExecution(log types.Log) (*MultisigwalletExecution, error) {
	event := new(MultisigwalletExecution)
	if err := _Multisigwallet.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the Multisigwallet contract.
type MultisigwalletExecutionFailureIterator struct {
	Event *MultisigwalletExecutionFailure // Event containing the contract specifics and raw log

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
func (it *MultisigwalletExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletExecutionFailure)
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
		it.Event = new(MultisigwalletExecutionFailure)
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
func (it *MultisigwalletExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletExecutionFailure represents a ExecutionFailure event raised by the Multisigwallet contract.
type MultisigwalletExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) FilterExecutionFailure(opts *bind.FilterOpts) (*MultisigwalletExecutionFailureIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletExecutionFailureIterator{contract: _Multisigwallet.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MultisigwalletExecutionFailure) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "ExecutionFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletExecutionFailure)
				if err := _Multisigwallet.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
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

// ParseExecutionFailure is a log parse operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) ParseExecutionFailure(log types.Log) (*MultisigwalletExecutionFailure, error) {
	event := new(MultisigwalletExecutionFailure)
	if err := _Multisigwallet.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the Multisigwallet contract.
type MultisigwalletOwnerAdditionIterator struct {
	Event *MultisigwalletOwnerAddition // Event containing the contract specifics and raw log

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
func (it *MultisigwalletOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletOwnerAddition)
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
		it.Event = new(MultisigwalletOwnerAddition)
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
func (it *MultisigwalletOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletOwnerAddition represents a OwnerAddition event raised by the Multisigwallet contract.
type MultisigwalletOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address owner)
func (_Multisigwallet *MultisigwalletFilterer) FilterOwnerAddition(opts *bind.FilterOpts) (*MultisigwalletOwnerAdditionIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "OwnerAddition")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletOwnerAdditionIterator{contract: _Multisigwallet.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address owner)
func (_Multisigwallet *MultisigwalletFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultisigwalletOwnerAddition) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "OwnerAddition")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletOwnerAddition)
				if err := _Multisigwallet.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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

// ParseOwnerAddition is a log parse operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address owner)
func (_Multisigwallet *MultisigwalletFilterer) ParseOwnerAddition(log types.Log) (*MultisigwalletOwnerAddition, error) {
	event := new(MultisigwalletOwnerAddition)
	if err := _Multisigwallet.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the Multisigwallet contract.
type MultisigwalletOwnerRemovalIterator struct {
	Event *MultisigwalletOwnerRemoval // Event containing the contract specifics and raw log

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
func (it *MultisigwalletOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletOwnerRemoval)
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
		it.Event = new(MultisigwalletOwnerRemoval)
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
func (it *MultisigwalletOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletOwnerRemoval represents a OwnerRemoval event raised by the Multisigwallet contract.
type MultisigwalletOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address owner)
func (_Multisigwallet *MultisigwalletFilterer) FilterOwnerRemoval(opts *bind.FilterOpts) (*MultisigwalletOwnerRemovalIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "OwnerRemoval")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletOwnerRemovalIterator{contract: _Multisigwallet.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address owner)
func (_Multisigwallet *MultisigwalletFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultisigwalletOwnerRemoval) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "OwnerRemoval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletOwnerRemoval)
				if err := _Multisigwallet.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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

// ParseOwnerRemoval is a log parse operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address owner)
func (_Multisigwallet *MultisigwalletFilterer) ParseOwnerRemoval(log types.Log) (*MultisigwalletOwnerRemoval, error) {
	event := new(MultisigwalletOwnerRemoval)
	if err := _Multisigwallet.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the Multisigwallet contract.
type MultisigwalletRequirementChangeIterator struct {
	Event *MultisigwalletRequirementChange // Event containing the contract specifics and raw log

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
func (it *MultisigwalletRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletRequirementChange)
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
		it.Event = new(MultisigwalletRequirementChange)
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
func (it *MultisigwalletRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletRequirementChange represents a RequirementChange event raised by the Multisigwallet contract.
type MultisigwalletRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_Multisigwallet *MultisigwalletFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultisigwalletRequirementChangeIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletRequirementChangeIterator{contract: _Multisigwallet.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_Multisigwallet *MultisigwalletFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultisigwalletRequirementChange) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletRequirementChange)
				if err := _Multisigwallet.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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

// ParseRequirementChange is a log parse operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_Multisigwallet *MultisigwalletFilterer) ParseRequirementChange(log types.Log) (*MultisigwalletRequirementChange, error) {
	event := new(MultisigwalletRequirementChange)
	if err := _Multisigwallet.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the Multisigwallet contract.
type MultisigwalletRevocationIterator struct {
	Event *MultisigwalletRevocation // Event containing the contract specifics and raw log

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
func (it *MultisigwalletRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletRevocation)
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
		it.Event = new(MultisigwalletRevocation)
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
func (it *MultisigwalletRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletRevocation represents a Revocation event raised by the Multisigwallet contract.
type MultisigwalletRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address sender, uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) FilterRevocation(opts *bind.FilterOpts) (*MultisigwalletRevocationIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "Revocation")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletRevocationIterator{contract: _Multisigwallet.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address sender, uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultisigwalletRevocation) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "Revocation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletRevocation)
				if err := _Multisigwallet.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// ParseRevocation is a log parse operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address sender, uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) ParseRevocation(log types.Log) (*MultisigwalletRevocation, error) {
	event := new(MultisigwalletRevocation)
	if err := _Multisigwallet.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultisigwalletSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the Multisigwallet contract.
type MultisigwalletSubmissionIterator struct {
	Event *MultisigwalletSubmission // Event containing the contract specifics and raw log

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
func (it *MultisigwalletSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultisigwalletSubmission)
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
		it.Event = new(MultisigwalletSubmission)
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
func (it *MultisigwalletSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultisigwalletSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultisigwalletSubmission represents a Submission event raised by the Multisigwallet contract.
type MultisigwalletSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) FilterSubmission(opts *bind.FilterOpts) (*MultisigwalletSubmissionIterator, error) {

	logs, sub, err := _Multisigwallet.contract.FilterLogs(opts, "Submission")
	if err != nil {
		return nil, err
	}
	return &MultisigwalletSubmissionIterator{contract: _Multisigwallet.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultisigwalletSubmission) (event.Subscription, error) {

	logs, sub, err := _Multisigwallet.contract.WatchLogs(opts, "Submission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultisigwalletSubmission)
				if err := _Multisigwallet.contract.UnpackLog(event, "Submission", log); err != nil {
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

// ParseSubmission is a log parse operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 transactionId)
func (_Multisigwallet *MultisigwalletFilterer) ParseSubmission(log types.Log) (*MultisigwalletSubmission, error) {
	event := new(MultisigwalletSubmission)
	if err := _Multisigwallet.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
