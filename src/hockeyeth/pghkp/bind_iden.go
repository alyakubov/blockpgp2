// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pghkp

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

// PgpIdenMetaData contains all meta data concerning the PgpIden contract.
var PgpIdenMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getSigntLen\",\"outputs\":[{\"name\":\"len\",\"type\":\"uint256\"},{\"name\":\"err\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_introducingFinger\",\"type\":\"bytes\"},{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"acceptProposedCert\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_signedFinger\",\"type\":\"bytes\"},{\"name\":\"_proposedCert\",\"type\":\"bytes\"}],\"name\":\"newSignt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cOrchestrAddr\",\"type\":\"address\"}],\"name\":\"Init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrchestrAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getIdenName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getIdenLoadDate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_ind\",\"type\":\"uint256\"}],\"name\":\"getSignedFinger\",\"outputs\":[{\"name\":\"_signedFinger\",\"type\":\"bytes\"},{\"name\":\"_proposedCert\",\"type\":\"bytes\"},{\"name\":\"_revocationDate\",\"type\":\"uint256\"},{\"name\":\"_err\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_email\",\"type\":\"string\"},{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_ownCert\",\"type\":\"bytes\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_certUser\",\"type\":\"address\"}],\"name\":\"newCertificate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getIdenCertUser\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getIdenOwnCert\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cIden\",\"outputs\":[{\"name\":\"cOrchestrAddr\",\"type\":\"address\"},{\"name\":\"cOrchestr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getIdenEmail\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_finger\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_armoured\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_code\",\"type\":\"uint256\"}],\"name\":\"evNewCertificateAnnounce\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_ind\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_err\",\"type\":\"int256\"},{\"indexed\":false,\"name\":\"_errMsg\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_hSha2\",\"type\":\"uint256\"}],\"name\":\"evNewCertificateReturn\",\"type\":\"event\"}]",
}

// PgpIdenABI is the input ABI used to generate the binding from.
// Deprecated: Use PgpIdenMetaData.ABI instead.
var PgpIdenABI = PgpIdenMetaData.ABI

// PgpIden is an auto generated Go binding around an Ethereum contract.
type PgpIden struct {
	PgpIdenCaller     // Read-only binding to the contract
	PgpIdenTransactor // Write-only binding to the contract
	PgpIdenFilterer   // Log filterer for contract events
}

// PgpIdenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PgpIdenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgpIdenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PgpIdenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgpIdenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PgpIdenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgpIdenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PgpIdenSession struct {
	Contract     *PgpIden          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PgpIdenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PgpIdenCallerSession struct {
	Contract *PgpIdenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PgpIdenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PgpIdenTransactorSession struct {
	Contract     *PgpIdenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PgpIdenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PgpIdenRaw struct {
	Contract *PgpIden // Generic contract binding to access the raw methods on
}

// PgpIdenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PgpIdenCallerRaw struct {
	Contract *PgpIdenCaller // Generic read-only contract binding to access the raw methods on
}

// PgpIdenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PgpIdenTransactorRaw struct {
	Contract *PgpIdenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPgpIden creates a new instance of PgpIden, bound to a specific deployed contract.
func NewPgpIden(address common.Address, backend bind.ContractBackend) (*PgpIden, error) {
	contract, err := bindPgpIden(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PgpIden{PgpIdenCaller: PgpIdenCaller{contract: contract}, PgpIdenTransactor: PgpIdenTransactor{contract: contract}, PgpIdenFilterer: PgpIdenFilterer{contract: contract}}, nil
}

// NewPgpIdenCaller creates a new read-only instance of PgpIden, bound to a specific deployed contract.
func NewPgpIdenCaller(address common.Address, caller bind.ContractCaller) (*PgpIdenCaller, error) {
	contract, err := bindPgpIden(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PgpIdenCaller{contract: contract}, nil
}

// NewPgpIdenTransactor creates a new write-only instance of PgpIden, bound to a specific deployed contract.
func NewPgpIdenTransactor(address common.Address, transactor bind.ContractTransactor) (*PgpIdenTransactor, error) {
	contract, err := bindPgpIden(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PgpIdenTransactor{contract: contract}, nil
}

// NewPgpIdenFilterer creates a new log filterer instance of PgpIden, bound to a specific deployed contract.
func NewPgpIdenFilterer(address common.Address, filterer bind.ContractFilterer) (*PgpIdenFilterer, error) {
	contract, err := bindPgpIden(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PgpIdenFilterer{contract: contract}, nil
}

// bindPgpIden binds a generic wrapper to an already deployed contract.
func bindPgpIden(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PgpIdenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PgpIden *PgpIdenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PgpIden.Contract.PgpIdenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PgpIden *PgpIdenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PgpIden.Contract.PgpIdenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PgpIden *PgpIdenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PgpIden.Contract.PgpIdenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PgpIden *PgpIdenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PgpIden.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PgpIden *PgpIdenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PgpIden.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PgpIden *PgpIdenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PgpIden.Contract.contract.Transact(opts, method, params...)
}

// CIden is a free data retrieval call binding the contract method 0xef2df7be.
//
// Solidity: function cIden() view returns(address cOrchestrAddr, address cOrchestr)
func (_PgpIden *PgpIdenCaller) CIden(opts *bind.CallOpts) (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
}, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "cIden")

	outstruct := new(struct {
		COrchestrAddr common.Address
		COrchestr     common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.COrchestrAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.COrchestr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CIden is a free data retrieval call binding the contract method 0xef2df7be.
//
// Solidity: function cIden() view returns(address cOrchestrAddr, address cOrchestr)
func (_PgpIden *PgpIdenSession) CIden() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
}, error) {
	return _PgpIden.Contract.CIden(&_PgpIden.CallOpts)
}

// CIden is a free data retrieval call binding the contract method 0xef2df7be.
//
// Solidity: function cIden() view returns(address cOrchestrAddr, address cOrchestr)
func (_PgpIden *PgpIdenCallerSession) CIden() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
}, error) {
	return _PgpIden.Contract.CIden(&_PgpIden.CallOpts)
}

// GetIdenCertUser is a free data retrieval call binding the contract method 0xd95b4e06.
//
// Solidity: function getIdenCertUser(bytes _finger) view returns(address)
func (_PgpIden *PgpIdenCaller) GetIdenCertUser(opts *bind.CallOpts, _finger []byte) (common.Address, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getIdenCertUser", _finger)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIdenCertUser is a free data retrieval call binding the contract method 0xd95b4e06.
//
// Solidity: function getIdenCertUser(bytes _finger) view returns(address)
func (_PgpIden *PgpIdenSession) GetIdenCertUser(_finger []byte) (common.Address, error) {
	return _PgpIden.Contract.GetIdenCertUser(&_PgpIden.CallOpts, _finger)
}

// GetIdenCertUser is a free data retrieval call binding the contract method 0xd95b4e06.
//
// Solidity: function getIdenCertUser(bytes _finger) view returns(address)
func (_PgpIden *PgpIdenCallerSession) GetIdenCertUser(_finger []byte) (common.Address, error) {
	return _PgpIden.Contract.GetIdenCertUser(&_PgpIden.CallOpts, _finger)
}

// GetIdenEmail is a free data retrieval call binding the contract method 0xf127ae9d.
//
// Solidity: function getIdenEmail(bytes _finger) view returns(string)
func (_PgpIden *PgpIdenCaller) GetIdenEmail(opts *bind.CallOpts, _finger []byte) (string, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getIdenEmail", _finger)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIdenEmail is a free data retrieval call binding the contract method 0xf127ae9d.
//
// Solidity: function getIdenEmail(bytes _finger) view returns(string)
func (_PgpIden *PgpIdenSession) GetIdenEmail(_finger []byte) (string, error) {
	return _PgpIden.Contract.GetIdenEmail(&_PgpIden.CallOpts, _finger)
}

// GetIdenEmail is a free data retrieval call binding the contract method 0xf127ae9d.
//
// Solidity: function getIdenEmail(bytes _finger) view returns(string)
func (_PgpIden *PgpIdenCallerSession) GetIdenEmail(_finger []byte) (string, error) {
	return _PgpIden.Contract.GetIdenEmail(&_PgpIden.CallOpts, _finger)
}

// GetIdenLoadDate is a free data retrieval call binding the contract method 0x8f448152.
//
// Solidity: function getIdenLoadDate(bytes _finger) view returns(uint256)
func (_PgpIden *PgpIdenCaller) GetIdenLoadDate(opts *bind.CallOpts, _finger []byte) (*big.Int, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getIdenLoadDate", _finger)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdenLoadDate is a free data retrieval call binding the contract method 0x8f448152.
//
// Solidity: function getIdenLoadDate(bytes _finger) view returns(uint256)
func (_PgpIden *PgpIdenSession) GetIdenLoadDate(_finger []byte) (*big.Int, error) {
	return _PgpIden.Contract.GetIdenLoadDate(&_PgpIden.CallOpts, _finger)
}

// GetIdenLoadDate is a free data retrieval call binding the contract method 0x8f448152.
//
// Solidity: function getIdenLoadDate(bytes _finger) view returns(uint256)
func (_PgpIden *PgpIdenCallerSession) GetIdenLoadDate(_finger []byte) (*big.Int, error) {
	return _PgpIden.Contract.GetIdenLoadDate(&_PgpIden.CallOpts, _finger)
}

// GetIdenName is a free data retrieval call binding the contract method 0x766b2c0d.
//
// Solidity: function getIdenName(bytes _finger) view returns(string)
func (_PgpIden *PgpIdenCaller) GetIdenName(opts *bind.CallOpts, _finger []byte) (string, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getIdenName", _finger)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetIdenName is a free data retrieval call binding the contract method 0x766b2c0d.
//
// Solidity: function getIdenName(bytes _finger) view returns(string)
func (_PgpIden *PgpIdenSession) GetIdenName(_finger []byte) (string, error) {
	return _PgpIden.Contract.GetIdenName(&_PgpIden.CallOpts, _finger)
}

// GetIdenName is a free data retrieval call binding the contract method 0x766b2c0d.
//
// Solidity: function getIdenName(bytes _finger) view returns(string)
func (_PgpIden *PgpIdenCallerSession) GetIdenName(_finger []byte) (string, error) {
	return _PgpIden.Contract.GetIdenName(&_PgpIden.CallOpts, _finger)
}

// GetIdenOwnCert is a free data retrieval call binding the contract method 0xe17ac131.
//
// Solidity: function getIdenOwnCert(bytes _finger) view returns(bytes)
func (_PgpIden *PgpIdenCaller) GetIdenOwnCert(opts *bind.CallOpts, _finger []byte) ([]byte, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getIdenOwnCert", _finger)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetIdenOwnCert is a free data retrieval call binding the contract method 0xe17ac131.
//
// Solidity: function getIdenOwnCert(bytes _finger) view returns(bytes)
func (_PgpIden *PgpIdenSession) GetIdenOwnCert(_finger []byte) ([]byte, error) {
	return _PgpIden.Contract.GetIdenOwnCert(&_PgpIden.CallOpts, _finger)
}

// GetIdenOwnCert is a free data retrieval call binding the contract method 0xe17ac131.
//
// Solidity: function getIdenOwnCert(bytes _finger) view returns(bytes)
func (_PgpIden *PgpIdenCallerSession) GetIdenOwnCert(_finger []byte) ([]byte, error) {
	return _PgpIden.Contract.GetIdenOwnCert(&_PgpIden.CallOpts, _finger)
}

// GetOrchestrAddr is a free data retrieval call binding the contract method 0x6371a0f3.
//
// Solidity: function getOrchestrAddr() view returns(address)
func (_PgpIden *PgpIdenCaller) GetOrchestrAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getOrchestrAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrchestrAddr is a free data retrieval call binding the contract method 0x6371a0f3.
//
// Solidity: function getOrchestrAddr() view returns(address)
func (_PgpIden *PgpIdenSession) GetOrchestrAddr() (common.Address, error) {
	return _PgpIden.Contract.GetOrchestrAddr(&_PgpIden.CallOpts)
}

// GetOrchestrAddr is a free data retrieval call binding the contract method 0x6371a0f3.
//
// Solidity: function getOrchestrAddr() view returns(address)
func (_PgpIden *PgpIdenCallerSession) GetOrchestrAddr() (common.Address, error) {
	return _PgpIden.Contract.GetOrchestrAddr(&_PgpIden.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PgpIden *PgpIdenCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PgpIden *PgpIdenSession) GetOwner() (common.Address, error) {
	return _PgpIden.Contract.GetOwner(&_PgpIden.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PgpIden *PgpIdenCallerSession) GetOwner() (common.Address, error) {
	return _PgpIden.Contract.GetOwner(&_PgpIden.CallOpts)
}

// GetSignedFinger is a free data retrieval call binding the contract method 0x918f4909.
//
// Solidity: function getSignedFinger(bytes _finger, uint256 _ind) view returns(bytes _signedFinger, bytes _proposedCert, uint256 _revocationDate, int256 _err)
func (_PgpIden *PgpIdenCaller) GetSignedFinger(opts *bind.CallOpts, _finger []byte, _ind *big.Int) (struct {
	SignedFinger   []byte
	ProposedCert   []byte
	RevocationDate *big.Int
	Err            *big.Int
}, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getSignedFinger", _finger, _ind)

	outstruct := new(struct {
		SignedFinger   []byte
		ProposedCert   []byte
		RevocationDate *big.Int
		Err            *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SignedFinger = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.ProposedCert = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.RevocationDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Err = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSignedFinger is a free data retrieval call binding the contract method 0x918f4909.
//
// Solidity: function getSignedFinger(bytes _finger, uint256 _ind) view returns(bytes _signedFinger, bytes _proposedCert, uint256 _revocationDate, int256 _err)
func (_PgpIden *PgpIdenSession) GetSignedFinger(_finger []byte, _ind *big.Int) (struct {
	SignedFinger   []byte
	ProposedCert   []byte
	RevocationDate *big.Int
	Err            *big.Int
}, error) {
	return _PgpIden.Contract.GetSignedFinger(&_PgpIden.CallOpts, _finger, _ind)
}

// GetSignedFinger is a free data retrieval call binding the contract method 0x918f4909.
//
// Solidity: function getSignedFinger(bytes _finger, uint256 _ind) view returns(bytes _signedFinger, bytes _proposedCert, uint256 _revocationDate, int256 _err)
func (_PgpIden *PgpIdenCallerSession) GetSignedFinger(_finger []byte, _ind *big.Int) (struct {
	SignedFinger   []byte
	ProposedCert   []byte
	RevocationDate *big.Int
	Err            *big.Int
}, error) {
	return _PgpIden.Contract.GetSignedFinger(&_PgpIden.CallOpts, _finger, _ind)
}

// GetSigntLen is a free data retrieval call binding the contract method 0x04dc9712.
//
// Solidity: function getSigntLen(bytes _finger) view returns(uint256 len, int256 err)
func (_PgpIden *PgpIdenCaller) GetSigntLen(opts *bind.CallOpts, _finger []byte) (struct {
	Len *big.Int
	Err *big.Int
}, error) {
	var out []interface{}
	err := _PgpIden.contract.Call(opts, &out, "getSigntLen", _finger)

	outstruct := new(struct {
		Len *big.Int
		Err *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Len = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Err = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSigntLen is a free data retrieval call binding the contract method 0x04dc9712.
//
// Solidity: function getSigntLen(bytes _finger) view returns(uint256 len, int256 err)
func (_PgpIden *PgpIdenSession) GetSigntLen(_finger []byte) (struct {
	Len *big.Int
	Err *big.Int
}, error) {
	return _PgpIden.Contract.GetSigntLen(&_PgpIden.CallOpts, _finger)
}

// GetSigntLen is a free data retrieval call binding the contract method 0x04dc9712.
//
// Solidity: function getSigntLen(bytes _finger) view returns(uint256 len, int256 err)
func (_PgpIden *PgpIdenCallerSession) GetSigntLen(_finger []byte) (struct {
	Len *big.Int
	Err *big.Int
}, error) {
	return _PgpIden.Contract.GetSigntLen(&_PgpIden.CallOpts, _finger)
}

// Init is a paid mutator transaction binding the contract method 0x4f8cfde3.
//
// Solidity: function Init(address _cOrchestrAddr) returns()
func (_PgpIden *PgpIdenTransactor) Init(opts *bind.TransactOpts, _cOrchestrAddr common.Address) (*types.Transaction, error) {
	return _PgpIden.contract.Transact(opts, "Init", _cOrchestrAddr)
}

// Init is a paid mutator transaction binding the contract method 0x4f8cfde3.
//
// Solidity: function Init(address _cOrchestrAddr) returns()
func (_PgpIden *PgpIdenSession) Init(_cOrchestrAddr common.Address) (*types.Transaction, error) {
	return _PgpIden.Contract.Init(&_PgpIden.TransactOpts, _cOrchestrAddr)
}

// Init is a paid mutator transaction binding the contract method 0x4f8cfde3.
//
// Solidity: function Init(address _cOrchestrAddr) returns()
func (_PgpIden *PgpIdenTransactorSession) Init(_cOrchestrAddr common.Address) (*types.Transaction, error) {
	return _PgpIden.Contract.Init(&_PgpIden.TransactOpts, _cOrchestrAddr)
}

// AcceptProposedCert is a paid mutator transaction binding the contract method 0x13dde595.
//
// Solidity: function acceptProposedCert(bytes _introducingFinger, bytes _finger) returns()
func (_PgpIden *PgpIdenTransactor) AcceptProposedCert(opts *bind.TransactOpts, _introducingFinger []byte, _finger []byte) (*types.Transaction, error) {
	return _PgpIden.contract.Transact(opts, "acceptProposedCert", _introducingFinger, _finger)
}

// AcceptProposedCert is a paid mutator transaction binding the contract method 0x13dde595.
//
// Solidity: function acceptProposedCert(bytes _introducingFinger, bytes _finger) returns()
func (_PgpIden *PgpIdenSession) AcceptProposedCert(_introducingFinger []byte, _finger []byte) (*types.Transaction, error) {
	return _PgpIden.Contract.AcceptProposedCert(&_PgpIden.TransactOpts, _introducingFinger, _finger)
}

// AcceptProposedCert is a paid mutator transaction binding the contract method 0x13dde595.
//
// Solidity: function acceptProposedCert(bytes _introducingFinger, bytes _finger) returns()
func (_PgpIden *PgpIdenTransactorSession) AcceptProposedCert(_introducingFinger []byte, _finger []byte) (*types.Transaction, error) {
	return _PgpIden.Contract.AcceptProposedCert(&_PgpIden.TransactOpts, _introducingFinger, _finger)
}

// NewCertificate is a paid mutator transaction binding the contract method 0xb8d36c69.
//
// Solidity: function newCertificate(string _email, bytes _finger, bytes _ownCert, string _name, address _certUser) returns()
func (_PgpIden *PgpIdenTransactor) NewCertificate(opts *bind.TransactOpts, _email string, _finger []byte, _ownCert []byte, _name string, _certUser common.Address) (*types.Transaction, error) {
	return _PgpIden.contract.Transact(opts, "newCertificate", _email, _finger, _ownCert, _name, _certUser)
}

// NewCertificate is a paid mutator transaction binding the contract method 0xb8d36c69.
//
// Solidity: function newCertificate(string _email, bytes _finger, bytes _ownCert, string _name, address _certUser) returns()
func (_PgpIden *PgpIdenSession) NewCertificate(_email string, _finger []byte, _ownCert []byte, _name string, _certUser common.Address) (*types.Transaction, error) {
	return _PgpIden.Contract.NewCertificate(&_PgpIden.TransactOpts, _email, _finger, _ownCert, _name, _certUser)
}

// NewCertificate is a paid mutator transaction binding the contract method 0xb8d36c69.
//
// Solidity: function newCertificate(string _email, bytes _finger, bytes _ownCert, string _name, address _certUser) returns()
func (_PgpIden *PgpIdenTransactorSession) NewCertificate(_email string, _finger []byte, _ownCert []byte, _name string, _certUser common.Address) (*types.Transaction, error) {
	return _PgpIden.Contract.NewCertificate(&_PgpIden.TransactOpts, _email, _finger, _ownCert, _name, _certUser)
}

// NewSignt is a paid mutator transaction binding the contract method 0x458e6a72.
//
// Solidity: function newSignt(bytes _finger, bytes _signedFinger, bytes _proposedCert) returns()
func (_PgpIden *PgpIdenTransactor) NewSignt(opts *bind.TransactOpts, _finger []byte, _signedFinger []byte, _proposedCert []byte) (*types.Transaction, error) {
	return _PgpIden.contract.Transact(opts, "newSignt", _finger, _signedFinger, _proposedCert)
}

// NewSignt is a paid mutator transaction binding the contract method 0x458e6a72.
//
// Solidity: function newSignt(bytes _finger, bytes _signedFinger, bytes _proposedCert) returns()
func (_PgpIden *PgpIdenSession) NewSignt(_finger []byte, _signedFinger []byte, _proposedCert []byte) (*types.Transaction, error) {
	return _PgpIden.Contract.NewSignt(&_PgpIden.TransactOpts, _finger, _signedFinger, _proposedCert)
}

// NewSignt is a paid mutator transaction binding the contract method 0x458e6a72.
//
// Solidity: function newSignt(bytes _finger, bytes _signedFinger, bytes _proposedCert) returns()
func (_PgpIden *PgpIdenTransactorSession) NewSignt(_finger []byte, _signedFinger []byte, _proposedCert []byte) (*types.Transaction, error) {
	return _PgpIden.Contract.NewSignt(&_PgpIden.TransactOpts, _finger, _signedFinger, _proposedCert)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_PgpIden *PgpIdenTransactor) SetOwner(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _PgpIden.contract.Transact(opts, "setOwner", _addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_PgpIden *PgpIdenSession) SetOwner(_addr common.Address) (*types.Transaction, error) {
	return _PgpIden.Contract.SetOwner(&_PgpIden.TransactOpts, _addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_PgpIden *PgpIdenTransactorSession) SetOwner(_addr common.Address) (*types.Transaction, error) {
	return _PgpIden.Contract.SetOwner(&_PgpIden.TransactOpts, _addr)
}

// PgpIdenEvNewCertificateAnnounceIterator is returned from FilterEvNewCertificateAnnounce and is used to iterate over the raw logs and unpacked data for EvNewCertificateAnnounce events raised by the PgpIden contract.
type PgpIdenEvNewCertificateAnnounceIterator struct {
	Event *PgpIdenEvNewCertificateAnnounce // Event containing the contract specifics and raw log

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
func (it *PgpIdenEvNewCertificateAnnounceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PgpIdenEvNewCertificateAnnounce)
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
		it.Event = new(PgpIdenEvNewCertificateAnnounce)
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
func (it *PgpIdenEvNewCertificateAnnounceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PgpIdenEvNewCertificateAnnounceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PgpIdenEvNewCertificateAnnounce represents a EvNewCertificateAnnounce event raised by the PgpIden contract.
type PgpIdenEvNewCertificateAnnounce struct {
	Finger   []byte
	Armoured []byte
	Code     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvNewCertificateAnnounce is a free log retrieval operation binding the contract event 0x10d6588a05abaa2ef9b4c79dd215f59bd0d0c51846b6677b05c57efb7f5166a1.
//
// Solidity: event evNewCertificateAnnounce(bytes _finger, bytes _armoured, uint256 _code)
func (_PgpIden *PgpIdenFilterer) FilterEvNewCertificateAnnounce(opts *bind.FilterOpts) (*PgpIdenEvNewCertificateAnnounceIterator, error) {

	logs, sub, err := _PgpIden.contract.FilterLogs(opts, "evNewCertificateAnnounce")
	if err != nil {
		return nil, err
	}
	return &PgpIdenEvNewCertificateAnnounceIterator{contract: _PgpIden.contract, event: "evNewCertificateAnnounce", logs: logs, sub: sub}, nil
}

// WatchEvNewCertificateAnnounce is a free log subscription operation binding the contract event 0x10d6588a05abaa2ef9b4c79dd215f59bd0d0c51846b6677b05c57efb7f5166a1.
//
// Solidity: event evNewCertificateAnnounce(bytes _finger, bytes _armoured, uint256 _code)
func (_PgpIden *PgpIdenFilterer) WatchEvNewCertificateAnnounce(opts *bind.WatchOpts, sink chan<- *PgpIdenEvNewCertificateAnnounce) (event.Subscription, error) {

	logs, sub, err := _PgpIden.contract.WatchLogs(opts, "evNewCertificateAnnounce")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PgpIdenEvNewCertificateAnnounce)
				if err := _PgpIden.contract.UnpackLog(event, "evNewCertificateAnnounce", log); err != nil {
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

// ParseEvNewCertificateAnnounce is a log parse operation binding the contract event 0x10d6588a05abaa2ef9b4c79dd215f59bd0d0c51846b6677b05c57efb7f5166a1.
//
// Solidity: event evNewCertificateAnnounce(bytes _finger, bytes _armoured, uint256 _code)
func (_PgpIden *PgpIdenFilterer) ParseEvNewCertificateAnnounce(log types.Log) (*PgpIdenEvNewCertificateAnnounce, error) {
	event := new(PgpIdenEvNewCertificateAnnounce)
	if err := _PgpIden.contract.UnpackLog(event, "evNewCertificateAnnounce", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PgpIdenEvNewCertificateReturnIterator is returned from FilterEvNewCertificateReturn and is used to iterate over the raw logs and unpacked data for EvNewCertificateReturn events raised by the PgpIden contract.
type PgpIdenEvNewCertificateReturnIterator struct {
	Event *PgpIdenEvNewCertificateReturn // Event containing the contract specifics and raw log

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
func (it *PgpIdenEvNewCertificateReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PgpIdenEvNewCertificateReturn)
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
		it.Event = new(PgpIdenEvNewCertificateReturn)
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
func (it *PgpIdenEvNewCertificateReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PgpIdenEvNewCertificateReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PgpIdenEvNewCertificateReturn represents a EvNewCertificateReturn event raised by the PgpIden contract.
type PgpIdenEvNewCertificateReturn struct {
	Ind    *big.Int
	Err    *big.Int
	ErrMsg string
	HSha2  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEvNewCertificateReturn is a free log retrieval operation binding the contract event 0x23eb6b1bea389065ac60a5eabcdf5de5bf39835c4f98b15144a0fb37ae4d68c5.
//
// Solidity: event evNewCertificateReturn(uint256 indexed _ind, int256 _err, string _errMsg, uint256 _hSha2)
func (_PgpIden *PgpIdenFilterer) FilterEvNewCertificateReturn(opts *bind.FilterOpts, _ind []*big.Int) (*PgpIdenEvNewCertificateReturnIterator, error) {

	var _indRule []interface{}
	for _, _indItem := range _ind {
		_indRule = append(_indRule, _indItem)
	}

	logs, sub, err := _PgpIden.contract.FilterLogs(opts, "evNewCertificateReturn", _indRule)
	if err != nil {
		return nil, err
	}
	return &PgpIdenEvNewCertificateReturnIterator{contract: _PgpIden.contract, event: "evNewCertificateReturn", logs: logs, sub: sub}, nil
}

// WatchEvNewCertificateReturn is a free log subscription operation binding the contract event 0x23eb6b1bea389065ac60a5eabcdf5de5bf39835c4f98b15144a0fb37ae4d68c5.
//
// Solidity: event evNewCertificateReturn(uint256 indexed _ind, int256 _err, string _errMsg, uint256 _hSha2)
func (_PgpIden *PgpIdenFilterer) WatchEvNewCertificateReturn(opts *bind.WatchOpts, sink chan<- *PgpIdenEvNewCertificateReturn, _ind []*big.Int) (event.Subscription, error) {

	var _indRule []interface{}
	for _, _indItem := range _ind {
		_indRule = append(_indRule, _indItem)
	}

	logs, sub, err := _PgpIden.contract.WatchLogs(opts, "evNewCertificateReturn", _indRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PgpIdenEvNewCertificateReturn)
				if err := _PgpIden.contract.UnpackLog(event, "evNewCertificateReturn", log); err != nil {
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

// ParseEvNewCertificateReturn is a log parse operation binding the contract event 0x23eb6b1bea389065ac60a5eabcdf5de5bf39835c4f98b15144a0fb37ae4d68c5.
//
// Solidity: event evNewCertificateReturn(uint256 indexed _ind, int256 _err, string _errMsg, uint256 _hSha2)
func (_PgpIden *PgpIdenFilterer) ParseEvNewCertificateReturn(log types.Log) (*PgpIdenEvNewCertificateReturn, error) {
	event := new(PgpIdenEvNewCertificateReturn)
	if err := _PgpIden.contract.UnpackLog(event, "evNewCertificateReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
