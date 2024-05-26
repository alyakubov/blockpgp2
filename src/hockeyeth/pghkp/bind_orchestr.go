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

// PgpOrchestrMetaData contains all meta data concerning the PgpOrchestr contract.
var PgpOrchestrMetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_arrInd\",\"type\":\"uint256\"}],\"name\":\"getSignRevokeInd\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cOrchestrAddr\",\"type\":\"address\"},{\"name\":\"_cOridenAddr\",\"type\":\"address\"}],\"name\":\"InitOrchestr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cFingers\",\"outputs\":[{\"name\":\"cOrchestrAddr\",\"type\":\"address\"},{\"name\":\"cOrchestr\",\"type\":\"address\"},{\"name\":\"cOridenAddr\",\"type\":\"address\"},{\"name\":\"cOriden\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_signedFinger\",\"type\":\"bytes\"}],\"name\":\"findSigntIndex\",\"outputs\":[{\"name\":\"_ind\",\"type\":\"uint256\"},{\"name\":\"_proposedCert\",\"type\":\"bytes\"},{\"name\":\"_err\",\"type\":\"int256\"},{\"name\":\"_errMsg\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_signedFinger\",\"type\":\"bytes\"}],\"name\":\"getEvIndSignId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_email\",\"type\":\"string\"}],\"name\":\"getFingersLen\",\"outputs\":[{\"name\":\"len\",\"type\":\"uint256\"},{\"name\":\"err\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_cOrchestrAddr\",\"type\":\"address\"}],\"name\":\"Init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_proposedCert\",\"type\":\"bytes\"}],\"name\":\"isKsrvCertFound\",\"outputs\":[{\"name\":\"isFound\",\"type\":\"uint256\"},{\"name\":\"_ind\",\"type\":\"uint256\"},{\"name\":\"_err\",\"type\":\"int256\"},{\"name\":\"_errMsg\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_ind\",\"type\":\"uint256\"}],\"name\":\"getSigntRevocation\",\"outputs\":[{\"name\":\"isRevoked\",\"type\":\"bool\"},{\"name\":\"revDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_email\",\"type\":\"string\"},{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"newFinger\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOrchestrAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cKsrv\",\"outputs\":[{\"name\":\"cOrchestrAddr\",\"type\":\"address\"},{\"name\":\"cOrchestr\",\"type\":\"address\"},{\"name\":\"cOridenAddr\",\"type\":\"address\"},{\"name\":\"cOriden\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_certUser\",\"type\":\"address\"}],\"name\":\"checkRights\",\"outputs\":[{\"name\":\"err\",\"type\":\"int256\"},{\"name\":\"errMsg\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_email\",\"type\":\"string\"},{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"isFingersFinger\",\"outputs\":[{\"name\":\"_ind\",\"type\":\"uint256\"},{\"name\":\"_err\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getEvIndCertId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_email\",\"type\":\"string\"},{\"name\":\"_ind\",\"type\":\"uint256\"}],\"name\":\"getFingersItem\",\"outputs\":[{\"name\":\"_res\",\"type\":\"bytes\"},{\"name\":\"_err\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_revocationDate\",\"type\":\"uint256\"}],\"name\":\"revokeCertificate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getIdenAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_armouredCert\",\"type\":\"bytes\"}],\"name\":\"calcCertHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"}],\"name\":\"getCertRevocation\",\"outputs\":[{\"name\":\"isRevoked\",\"type\":\"bool\"},{\"name\":\"revDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cRevoc\",\"outputs\":[{\"name\":\"cOrchestrAddr\",\"type\":\"address\"},{\"name\":\"cOrchestr\",\"type\":\"address\"},{\"name\":\"cOridenAddr\",\"type\":\"address\"},{\"name\":\"cOriden\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_finger\",\"type\":\"bytes\"},{\"name\":\"_signedFinger\",\"type\":\"bytes\"},{\"name\":\"_revocationDate\",\"type\":\"uint256\"}],\"name\":\"revokeSignt\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PgpOrchestrABI is the input ABI used to generate the binding from.
// Deprecated: Use PgpOrchestrMetaData.ABI instead.
var PgpOrchestrABI = PgpOrchestrMetaData.ABI

// PgpOrchestr is an auto generated Go binding around an Ethereum contract.
type PgpOrchestr struct {
	PgpOrchestrCaller     // Read-only binding to the contract
	PgpOrchestrTransactor // Write-only binding to the contract
	PgpOrchestrFilterer   // Log filterer for contract events
}

// PgpOrchestrCaller is an auto generated read-only Go binding around an Ethereum contract.
type PgpOrchestrCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgpOrchestrTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PgpOrchestrTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgpOrchestrFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PgpOrchestrFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PgpOrchestrSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PgpOrchestrSession struct {
	Contract     *PgpOrchestr      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PgpOrchestrCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PgpOrchestrCallerSession struct {
	Contract *PgpOrchestrCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PgpOrchestrTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PgpOrchestrTransactorSession struct {
	Contract     *PgpOrchestrTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PgpOrchestrRaw is an auto generated low-level Go binding around an Ethereum contract.
type PgpOrchestrRaw struct {
	Contract *PgpOrchestr // Generic contract binding to access the raw methods on
}

// PgpOrchestrCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PgpOrchestrCallerRaw struct {
	Contract *PgpOrchestrCaller // Generic read-only contract binding to access the raw methods on
}

// PgpOrchestrTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PgpOrchestrTransactorRaw struct {
	Contract *PgpOrchestrTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPgpOrchestr creates a new instance of PgpOrchestr, bound to a specific deployed contract.
func NewPgpOrchestr(address common.Address, backend bind.ContractBackend) (*PgpOrchestr, error) {
	contract, err := bindPgpOrchestr(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PgpOrchestr{PgpOrchestrCaller: PgpOrchestrCaller{contract: contract}, PgpOrchestrTransactor: PgpOrchestrTransactor{contract: contract}, PgpOrchestrFilterer: PgpOrchestrFilterer{contract: contract}}, nil
}

// NewPgpOrchestrCaller creates a new read-only instance of PgpOrchestr, bound to a specific deployed contract.
func NewPgpOrchestrCaller(address common.Address, caller bind.ContractCaller) (*PgpOrchestrCaller, error) {
	contract, err := bindPgpOrchestr(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PgpOrchestrCaller{contract: contract}, nil
}

// NewPgpOrchestrTransactor creates a new write-only instance of PgpOrchestr, bound to a specific deployed contract.
func NewPgpOrchestrTransactor(address common.Address, transactor bind.ContractTransactor) (*PgpOrchestrTransactor, error) {
	contract, err := bindPgpOrchestr(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PgpOrchestrTransactor{contract: contract}, nil
}

// NewPgpOrchestrFilterer creates a new log filterer instance of PgpOrchestr, bound to a specific deployed contract.
func NewPgpOrchestrFilterer(address common.Address, filterer bind.ContractFilterer) (*PgpOrchestrFilterer, error) {
	contract, err := bindPgpOrchestr(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PgpOrchestrFilterer{contract: contract}, nil
}

// bindPgpOrchestr binds a generic wrapper to an already deployed contract.
func bindPgpOrchestr(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PgpOrchestrABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PgpOrchestr *PgpOrchestrRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PgpOrchestr.Contract.PgpOrchestrCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PgpOrchestr *PgpOrchestrRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.PgpOrchestrTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PgpOrchestr *PgpOrchestrRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.PgpOrchestrTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PgpOrchestr *PgpOrchestrCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PgpOrchestr.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PgpOrchestr *PgpOrchestrTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PgpOrchestr *PgpOrchestrTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.contract.Transact(opts, method, params...)
}

// CFingers is a free data retrieval call binding the contract method 0x0a72beab.
//
// Solidity: function cFingers() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrCaller) CFingers(opts *bind.CallOpts) (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "cFingers")

	outstruct := new(struct {
		COrchestrAddr common.Address
		COrchestr     common.Address
		COridenAddr   common.Address
		COriden       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.COrchestrAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.COrchestr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.COridenAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.COriden = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CFingers is a free data retrieval call binding the contract method 0x0a72beab.
//
// Solidity: function cFingers() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrSession) CFingers() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	return _PgpOrchestr.Contract.CFingers(&_PgpOrchestr.CallOpts)
}

// CFingers is a free data retrieval call binding the contract method 0x0a72beab.
//
// Solidity: function cFingers() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrCallerSession) CFingers() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	return _PgpOrchestr.Contract.CFingers(&_PgpOrchestr.CallOpts)
}

// CKsrv is a free data retrieval call binding the contract method 0x65439dad.
//
// Solidity: function cKsrv() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrCaller) CKsrv(opts *bind.CallOpts) (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "cKsrv")

	outstruct := new(struct {
		COrchestrAddr common.Address
		COrchestr     common.Address
		COridenAddr   common.Address
		COriden       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.COrchestrAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.COrchestr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.COridenAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.COriden = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CKsrv is a free data retrieval call binding the contract method 0x65439dad.
//
// Solidity: function cKsrv() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrSession) CKsrv() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	return _PgpOrchestr.Contract.CKsrv(&_PgpOrchestr.CallOpts)
}

// CKsrv is a free data retrieval call binding the contract method 0x65439dad.
//
// Solidity: function cKsrv() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrCallerSession) CKsrv() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	return _PgpOrchestr.Contract.CKsrv(&_PgpOrchestr.CallOpts)
}

// CRevoc is a free data retrieval call binding the contract method 0xf20a60b1.
//
// Solidity: function cRevoc() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrCaller) CRevoc(opts *bind.CallOpts) (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "cRevoc")

	outstruct := new(struct {
		COrchestrAddr common.Address
		COrchestr     common.Address
		COridenAddr   common.Address
		COriden       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.COrchestrAddr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.COrchestr = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.COridenAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.COriden = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CRevoc is a free data retrieval call binding the contract method 0xf20a60b1.
//
// Solidity: function cRevoc() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrSession) CRevoc() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	return _PgpOrchestr.Contract.CRevoc(&_PgpOrchestr.CallOpts)
}

// CRevoc is a free data retrieval call binding the contract method 0xf20a60b1.
//
// Solidity: function cRevoc() view returns(address cOrchestrAddr, address cOrchestr, address cOridenAddr, address cOriden)
func (_PgpOrchestr *PgpOrchestrCallerSession) CRevoc() (struct {
	COrchestrAddr common.Address
	COrchestr     common.Address
	COridenAddr   common.Address
	COriden       common.Address
}, error) {
	return _PgpOrchestr.Contract.CRevoc(&_PgpOrchestr.CallOpts)
}

// CalcCertHash is a free data retrieval call binding the contract method 0xbfc5de5b.
//
// Solidity: function calcCertHash(bytes _armouredCert) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCaller) CalcCertHash(opts *bind.CallOpts, _armouredCert []byte) (*big.Int, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "calcCertHash", _armouredCert)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcCertHash is a free data retrieval call binding the contract method 0xbfc5de5b.
//
// Solidity: function calcCertHash(bytes _armouredCert) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrSession) CalcCertHash(_armouredCert []byte) (*big.Int, error) {
	return _PgpOrchestr.Contract.CalcCertHash(&_PgpOrchestr.CallOpts, _armouredCert)
}

// CalcCertHash is a free data retrieval call binding the contract method 0xbfc5de5b.
//
// Solidity: function calcCertHash(bytes _armouredCert) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCallerSession) CalcCertHash(_armouredCert []byte) (*big.Int, error) {
	return _PgpOrchestr.Contract.CalcCertHash(&_PgpOrchestr.CallOpts, _armouredCert)
}

// CheckRights is a free data retrieval call binding the contract method 0x66b531d6.
//
// Solidity: function checkRights(bytes _finger, address _certUser) view returns(int256 err, string errMsg)
func (_PgpOrchestr *PgpOrchestrCaller) CheckRights(opts *bind.CallOpts, _finger []byte, _certUser common.Address) (struct {
	Err    *big.Int
	ErrMsg string
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "checkRights", _finger, _certUser)

	outstruct := new(struct {
		Err    *big.Int
		ErrMsg string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Err = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ErrMsg = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// CheckRights is a free data retrieval call binding the contract method 0x66b531d6.
//
// Solidity: function checkRights(bytes _finger, address _certUser) view returns(int256 err, string errMsg)
func (_PgpOrchestr *PgpOrchestrSession) CheckRights(_finger []byte, _certUser common.Address) (struct {
	Err    *big.Int
	ErrMsg string
}, error) {
	return _PgpOrchestr.Contract.CheckRights(&_PgpOrchestr.CallOpts, _finger, _certUser)
}

// CheckRights is a free data retrieval call binding the contract method 0x66b531d6.
//
// Solidity: function checkRights(bytes _finger, address _certUser) view returns(int256 err, string errMsg)
func (_PgpOrchestr *PgpOrchestrCallerSession) CheckRights(_finger []byte, _certUser common.Address) (struct {
	Err    *big.Int
	ErrMsg string
}, error) {
	return _PgpOrchestr.Contract.CheckRights(&_PgpOrchestr.CallOpts, _finger, _certUser)
}

// FindSigntIndex is a free data retrieval call binding the contract method 0x0b4189df.
//
// Solidity: function findSigntIndex(bytes _finger, bytes _signedFinger) view returns(uint256 _ind, bytes _proposedCert, int256 _err, string _errMsg)
func (_PgpOrchestr *PgpOrchestrCaller) FindSigntIndex(opts *bind.CallOpts, _finger []byte, _signedFinger []byte) (struct {
	Ind          *big.Int
	ProposedCert []byte
	Err          *big.Int
	ErrMsg       string
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "findSigntIndex", _finger, _signedFinger)

	outstruct := new(struct {
		Ind          *big.Int
		ProposedCert []byte
		Err          *big.Int
		ErrMsg       string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ind = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProposedCert = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	outstruct.Err = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ErrMsg = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// FindSigntIndex is a free data retrieval call binding the contract method 0x0b4189df.
//
// Solidity: function findSigntIndex(bytes _finger, bytes _signedFinger) view returns(uint256 _ind, bytes _proposedCert, int256 _err, string _errMsg)
func (_PgpOrchestr *PgpOrchestrSession) FindSigntIndex(_finger []byte, _signedFinger []byte) (struct {
	Ind          *big.Int
	ProposedCert []byte
	Err          *big.Int
	ErrMsg       string
}, error) {
	return _PgpOrchestr.Contract.FindSigntIndex(&_PgpOrchestr.CallOpts, _finger, _signedFinger)
}

// FindSigntIndex is a free data retrieval call binding the contract method 0x0b4189df.
//
// Solidity: function findSigntIndex(bytes _finger, bytes _signedFinger) view returns(uint256 _ind, bytes _proposedCert, int256 _err, string _errMsg)
func (_PgpOrchestr *PgpOrchestrCallerSession) FindSigntIndex(_finger []byte, _signedFinger []byte) (struct {
	Ind          *big.Int
	ProposedCert []byte
	Err          *big.Int
	ErrMsg       string
}, error) {
	return _PgpOrchestr.Contract.FindSigntIndex(&_PgpOrchestr.CallOpts, _finger, _signedFinger)
}

// GetCertRevocation is a free data retrieval call binding the contract method 0xdd95208e.
//
// Solidity: function getCertRevocation(bytes _finger) view returns(bool isRevoked, uint256 revDate)
func (_PgpOrchestr *PgpOrchestrCaller) GetCertRevocation(opts *bind.CallOpts, _finger []byte) (struct {
	IsRevoked bool
	RevDate   *big.Int
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getCertRevocation", _finger)

	outstruct := new(struct {
		IsRevoked bool
		RevDate   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRevoked = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RevDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetCertRevocation is a free data retrieval call binding the contract method 0xdd95208e.
//
// Solidity: function getCertRevocation(bytes _finger) view returns(bool isRevoked, uint256 revDate)
func (_PgpOrchestr *PgpOrchestrSession) GetCertRevocation(_finger []byte) (struct {
	IsRevoked bool
	RevDate   *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetCertRevocation(&_PgpOrchestr.CallOpts, _finger)
}

// GetCertRevocation is a free data retrieval call binding the contract method 0xdd95208e.
//
// Solidity: function getCertRevocation(bytes _finger) view returns(bool isRevoked, uint256 revDate)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetCertRevocation(_finger []byte) (struct {
	IsRevoked bool
	RevDate   *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetCertRevocation(&_PgpOrchestr.CallOpts, _finger)
}

// GetEvIndCertId is a free data retrieval call binding the contract method 0x85f8a6f2.
//
// Solidity: function getEvIndCertId(bytes _finger) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCaller) GetEvIndCertId(opts *bind.CallOpts, _finger []byte) (*big.Int, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getEvIndCertId", _finger)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEvIndCertId is a free data retrieval call binding the contract method 0x85f8a6f2.
//
// Solidity: function getEvIndCertId(bytes _finger) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrSession) GetEvIndCertId(_finger []byte) (*big.Int, error) {
	return _PgpOrchestr.Contract.GetEvIndCertId(&_PgpOrchestr.CallOpts, _finger)
}

// GetEvIndCertId is a free data retrieval call binding the contract method 0x85f8a6f2.
//
// Solidity: function getEvIndCertId(bytes _finger) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetEvIndCertId(_finger []byte) (*big.Int, error) {
	return _PgpOrchestr.Contract.GetEvIndCertId(&_PgpOrchestr.CallOpts, _finger)
}

// GetEvIndSignId is a free data retrieval call binding the contract method 0x0bf6e944.
//
// Solidity: function getEvIndSignId(bytes _finger, bytes _signedFinger) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCaller) GetEvIndSignId(opts *bind.CallOpts, _finger []byte, _signedFinger []byte) (*big.Int, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getEvIndSignId", _finger, _signedFinger)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEvIndSignId is a free data retrieval call binding the contract method 0x0bf6e944.
//
// Solidity: function getEvIndSignId(bytes _finger, bytes _signedFinger) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrSession) GetEvIndSignId(_finger []byte, _signedFinger []byte) (*big.Int, error) {
	return _PgpOrchestr.Contract.GetEvIndSignId(&_PgpOrchestr.CallOpts, _finger, _signedFinger)
}

// GetEvIndSignId is a free data retrieval call binding the contract method 0x0bf6e944.
//
// Solidity: function getEvIndSignId(bytes _finger, bytes _signedFinger) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetEvIndSignId(_finger []byte, _signedFinger []byte) (*big.Int, error) {
	return _PgpOrchestr.Contract.GetEvIndSignId(&_PgpOrchestr.CallOpts, _finger, _signedFinger)
}

// GetFingersItem is a free data retrieval call binding the contract method 0xa15e7272.
//
// Solidity: function getFingersItem(string _email, uint256 _ind) view returns(bytes _res, int256 _err)
func (_PgpOrchestr *PgpOrchestrCaller) GetFingersItem(opts *bind.CallOpts, _email string, _ind *big.Int) (struct {
	Res []byte
	Err *big.Int
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getFingersItem", _email, _ind)

	outstruct := new(struct {
		Res []byte
		Err *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Res = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Err = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetFingersItem is a free data retrieval call binding the contract method 0xa15e7272.
//
// Solidity: function getFingersItem(string _email, uint256 _ind) view returns(bytes _res, int256 _err)
func (_PgpOrchestr *PgpOrchestrSession) GetFingersItem(_email string, _ind *big.Int) (struct {
	Res []byte
	Err *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetFingersItem(&_PgpOrchestr.CallOpts, _email, _ind)
}

// GetFingersItem is a free data retrieval call binding the contract method 0xa15e7272.
//
// Solidity: function getFingersItem(string _email, uint256 _ind) view returns(bytes _res, int256 _err)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetFingersItem(_email string, _ind *big.Int) (struct {
	Res []byte
	Err *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetFingersItem(&_PgpOrchestr.CallOpts, _email, _ind)
}

// GetFingersLen is a free data retrieval call binding the contract method 0x4b67b238.
//
// Solidity: function getFingersLen(string _email) view returns(uint256 len, int256 err)
func (_PgpOrchestr *PgpOrchestrCaller) GetFingersLen(opts *bind.CallOpts, _email string) (struct {
	Len *big.Int
	Err *big.Int
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getFingersLen", _email)

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

// GetFingersLen is a free data retrieval call binding the contract method 0x4b67b238.
//
// Solidity: function getFingersLen(string _email) view returns(uint256 len, int256 err)
func (_PgpOrchestr *PgpOrchestrSession) GetFingersLen(_email string) (struct {
	Len *big.Int
	Err *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetFingersLen(&_PgpOrchestr.CallOpts, _email)
}

// GetFingersLen is a free data retrieval call binding the contract method 0x4b67b238.
//
// Solidity: function getFingersLen(string _email) view returns(uint256 len, int256 err)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetFingersLen(_email string) (struct {
	Len *big.Int
	Err *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetFingersLen(&_PgpOrchestr.CallOpts, _email)
}

// GetIdenAddr is a free data retrieval call binding the contract method 0xbda2831d.
//
// Solidity: function getIdenAddr() view returns(address)
func (_PgpOrchestr *PgpOrchestrCaller) GetIdenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getIdenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIdenAddr is a free data retrieval call binding the contract method 0xbda2831d.
//
// Solidity: function getIdenAddr() view returns(address)
func (_PgpOrchestr *PgpOrchestrSession) GetIdenAddr() (common.Address, error) {
	return _PgpOrchestr.Contract.GetIdenAddr(&_PgpOrchestr.CallOpts)
}

// GetIdenAddr is a free data retrieval call binding the contract method 0xbda2831d.
//
// Solidity: function getIdenAddr() view returns(address)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetIdenAddr() (common.Address, error) {
	return _PgpOrchestr.Contract.GetIdenAddr(&_PgpOrchestr.CallOpts)
}

// GetOrchestrAddr is a free data retrieval call binding the contract method 0x6371a0f3.
//
// Solidity: function getOrchestrAddr() view returns(address)
func (_PgpOrchestr *PgpOrchestrCaller) GetOrchestrAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getOrchestrAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOrchestrAddr is a free data retrieval call binding the contract method 0x6371a0f3.
//
// Solidity: function getOrchestrAddr() view returns(address)
func (_PgpOrchestr *PgpOrchestrSession) GetOrchestrAddr() (common.Address, error) {
	return _PgpOrchestr.Contract.GetOrchestrAddr(&_PgpOrchestr.CallOpts)
}

// GetOrchestrAddr is a free data retrieval call binding the contract method 0x6371a0f3.
//
// Solidity: function getOrchestrAddr() view returns(address)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetOrchestrAddr() (common.Address, error) {
	return _PgpOrchestr.Contract.GetOrchestrAddr(&_PgpOrchestr.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PgpOrchestr *PgpOrchestrCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PgpOrchestr *PgpOrchestrSession) GetOwner() (common.Address, error) {
	return _PgpOrchestr.Contract.GetOwner(&_PgpOrchestr.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetOwner() (common.Address, error) {
	return _PgpOrchestr.Contract.GetOwner(&_PgpOrchestr.CallOpts)
}

// GetSignRevokeInd is a free data retrieval call binding the contract method 0x006147ff.
//
// Solidity: function getSignRevokeInd(bytes _finger, uint256 _arrInd) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCaller) GetSignRevokeInd(opts *bind.CallOpts, _finger []byte, _arrInd *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getSignRevokeInd", _finger, _arrInd)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSignRevokeInd is a free data retrieval call binding the contract method 0x006147ff.
//
// Solidity: function getSignRevokeInd(bytes _finger, uint256 _arrInd) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrSession) GetSignRevokeInd(_finger []byte, _arrInd *big.Int) (*big.Int, error) {
	return _PgpOrchestr.Contract.GetSignRevokeInd(&_PgpOrchestr.CallOpts, _finger, _arrInd)
}

// GetSignRevokeInd is a free data retrieval call binding the contract method 0x006147ff.
//
// Solidity: function getSignRevokeInd(bytes _finger, uint256 _arrInd) pure returns(uint256)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetSignRevokeInd(_finger []byte, _arrInd *big.Int) (*big.Int, error) {
	return _PgpOrchestr.Contract.GetSignRevokeInd(&_PgpOrchestr.CallOpts, _finger, _arrInd)
}

// GetSigntRevocation is a free data retrieval call binding the contract method 0x59c524b9.
//
// Solidity: function getSigntRevocation(bytes _finger, uint256 _ind) view returns(bool isRevoked, uint256 revDate)
func (_PgpOrchestr *PgpOrchestrCaller) GetSigntRevocation(opts *bind.CallOpts, _finger []byte, _ind *big.Int) (struct {
	IsRevoked bool
	RevDate   *big.Int
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "getSigntRevocation", _finger, _ind)

	outstruct := new(struct {
		IsRevoked bool
		RevDate   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsRevoked = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RevDate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSigntRevocation is a free data retrieval call binding the contract method 0x59c524b9.
//
// Solidity: function getSigntRevocation(bytes _finger, uint256 _ind) view returns(bool isRevoked, uint256 revDate)
func (_PgpOrchestr *PgpOrchestrSession) GetSigntRevocation(_finger []byte, _ind *big.Int) (struct {
	IsRevoked bool
	RevDate   *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetSigntRevocation(&_PgpOrchestr.CallOpts, _finger, _ind)
}

// GetSigntRevocation is a free data retrieval call binding the contract method 0x59c524b9.
//
// Solidity: function getSigntRevocation(bytes _finger, uint256 _ind) view returns(bool isRevoked, uint256 revDate)
func (_PgpOrchestr *PgpOrchestrCallerSession) GetSigntRevocation(_finger []byte, _ind *big.Int) (struct {
	IsRevoked bool
	RevDate   *big.Int
}, error) {
	return _PgpOrchestr.Contract.GetSigntRevocation(&_PgpOrchestr.CallOpts, _finger, _ind)
}

// IsFingersFinger is a free data retrieval call binding the contract method 0x77428a08.
//
// Solidity: function isFingersFinger(string _email, bytes _finger) view returns(uint256 _ind, int256 _err)
func (_PgpOrchestr *PgpOrchestrCaller) IsFingersFinger(opts *bind.CallOpts, _email string, _finger []byte) (struct {
	Ind *big.Int
	Err *big.Int
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "isFingersFinger", _email, _finger)

	outstruct := new(struct {
		Ind *big.Int
		Err *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ind = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Err = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// IsFingersFinger is a free data retrieval call binding the contract method 0x77428a08.
//
// Solidity: function isFingersFinger(string _email, bytes _finger) view returns(uint256 _ind, int256 _err)
func (_PgpOrchestr *PgpOrchestrSession) IsFingersFinger(_email string, _finger []byte) (struct {
	Ind *big.Int
	Err *big.Int
}, error) {
	return _PgpOrchestr.Contract.IsFingersFinger(&_PgpOrchestr.CallOpts, _email, _finger)
}

// IsFingersFinger is a free data retrieval call binding the contract method 0x77428a08.
//
// Solidity: function isFingersFinger(string _email, bytes _finger) view returns(uint256 _ind, int256 _err)
func (_PgpOrchestr *PgpOrchestrCallerSession) IsFingersFinger(_email string, _finger []byte) (struct {
	Ind *big.Int
	Err *big.Int
}, error) {
	return _PgpOrchestr.Contract.IsFingersFinger(&_PgpOrchestr.CallOpts, _email, _finger)
}

// IsKsrvCertFound is a free data retrieval call binding the contract method 0x57555f95.
//
// Solidity: function isKsrvCertFound(bytes _finger, bytes _proposedCert) view returns(uint256 isFound, uint256 _ind, int256 _err, string _errMsg)
func (_PgpOrchestr *PgpOrchestrCaller) IsKsrvCertFound(opts *bind.CallOpts, _finger []byte, _proposedCert []byte) (struct {
	IsFound *big.Int
	Ind     *big.Int
	Err     *big.Int
	ErrMsg  string
}, error) {
	var out []interface{}
	err := _PgpOrchestr.contract.Call(opts, &out, "isKsrvCertFound", _finger, _proposedCert)

	outstruct := new(struct {
		IsFound *big.Int
		Ind     *big.Int
		Err     *big.Int
		ErrMsg  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsFound = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ind = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Err = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ErrMsg = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// IsKsrvCertFound is a free data retrieval call binding the contract method 0x57555f95.
//
// Solidity: function isKsrvCertFound(bytes _finger, bytes _proposedCert) view returns(uint256 isFound, uint256 _ind, int256 _err, string _errMsg)
func (_PgpOrchestr *PgpOrchestrSession) IsKsrvCertFound(_finger []byte, _proposedCert []byte) (struct {
	IsFound *big.Int
	Ind     *big.Int
	Err     *big.Int
	ErrMsg  string
}, error) {
	return _PgpOrchestr.Contract.IsKsrvCertFound(&_PgpOrchestr.CallOpts, _finger, _proposedCert)
}

// IsKsrvCertFound is a free data retrieval call binding the contract method 0x57555f95.
//
// Solidity: function isKsrvCertFound(bytes _finger, bytes _proposedCert) view returns(uint256 isFound, uint256 _ind, int256 _err, string _errMsg)
func (_PgpOrchestr *PgpOrchestrCallerSession) IsKsrvCertFound(_finger []byte, _proposedCert []byte) (struct {
	IsFound *big.Int
	Ind     *big.Int
	Err     *big.Int
	ErrMsg  string
}, error) {
	return _PgpOrchestr.Contract.IsKsrvCertFound(&_PgpOrchestr.CallOpts, _finger, _proposedCert)
}

// Init is a paid mutator transaction binding the contract method 0x4f8cfde3.
//
// Solidity: function Init(address _cOrchestrAddr) returns()
func (_PgpOrchestr *PgpOrchestrTransactor) Init(opts *bind.TransactOpts, _cOrchestrAddr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.contract.Transact(opts, "Init", _cOrchestrAddr)
}

// Init is a paid mutator transaction binding the contract method 0x4f8cfde3.
//
// Solidity: function Init(address _cOrchestrAddr) returns()
func (_PgpOrchestr *PgpOrchestrSession) Init(_cOrchestrAddr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.Init(&_PgpOrchestr.TransactOpts, _cOrchestrAddr)
}

// Init is a paid mutator transaction binding the contract method 0x4f8cfde3.
//
// Solidity: function Init(address _cOrchestrAddr) returns()
func (_PgpOrchestr *PgpOrchestrTransactorSession) Init(_cOrchestrAddr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.Init(&_PgpOrchestr.TransactOpts, _cOrchestrAddr)
}

// InitOrchestr is a paid mutator transaction binding the contract method 0x0279a503.
//
// Solidity: function InitOrchestr(address _cOrchestrAddr, address _cOridenAddr) returns()
func (_PgpOrchestr *PgpOrchestrTransactor) InitOrchestr(opts *bind.TransactOpts, _cOrchestrAddr common.Address, _cOridenAddr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.contract.Transact(opts, "InitOrchestr", _cOrchestrAddr, _cOridenAddr)
}

// InitOrchestr is a paid mutator transaction binding the contract method 0x0279a503.
//
// Solidity: function InitOrchestr(address _cOrchestrAddr, address _cOridenAddr) returns()
func (_PgpOrchestr *PgpOrchestrSession) InitOrchestr(_cOrchestrAddr common.Address, _cOridenAddr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.InitOrchestr(&_PgpOrchestr.TransactOpts, _cOrchestrAddr, _cOridenAddr)
}

// InitOrchestr is a paid mutator transaction binding the contract method 0x0279a503.
//
// Solidity: function InitOrchestr(address _cOrchestrAddr, address _cOridenAddr) returns()
func (_PgpOrchestr *PgpOrchestrTransactorSession) InitOrchestr(_cOrchestrAddr common.Address, _cOridenAddr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.InitOrchestr(&_PgpOrchestr.TransactOpts, _cOrchestrAddr, _cOridenAddr)
}

// NewFinger is a paid mutator transaction binding the contract method 0x5ab67db1.
//
// Solidity: function newFinger(string _email, bytes _finger) returns()
func (_PgpOrchestr *PgpOrchestrTransactor) NewFinger(opts *bind.TransactOpts, _email string, _finger []byte) (*types.Transaction, error) {
	return _PgpOrchestr.contract.Transact(opts, "newFinger", _email, _finger)
}

// NewFinger is a paid mutator transaction binding the contract method 0x5ab67db1.
//
// Solidity: function newFinger(string _email, bytes _finger) returns()
func (_PgpOrchestr *PgpOrchestrSession) NewFinger(_email string, _finger []byte) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.NewFinger(&_PgpOrchestr.TransactOpts, _email, _finger)
}

// NewFinger is a paid mutator transaction binding the contract method 0x5ab67db1.
//
// Solidity: function newFinger(string _email, bytes _finger) returns()
func (_PgpOrchestr *PgpOrchestrTransactorSession) NewFinger(_email string, _finger []byte) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.NewFinger(&_PgpOrchestr.TransactOpts, _email, _finger)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xa4fd2992.
//
// Solidity: function revokeCertificate(bytes _finger, uint256 _revocationDate) returns()
func (_PgpOrchestr *PgpOrchestrTransactor) RevokeCertificate(opts *bind.TransactOpts, _finger []byte, _revocationDate *big.Int) (*types.Transaction, error) {
	return _PgpOrchestr.contract.Transact(opts, "revokeCertificate", _finger, _revocationDate)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xa4fd2992.
//
// Solidity: function revokeCertificate(bytes _finger, uint256 _revocationDate) returns()
func (_PgpOrchestr *PgpOrchestrSession) RevokeCertificate(_finger []byte, _revocationDate *big.Int) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.RevokeCertificate(&_PgpOrchestr.TransactOpts, _finger, _revocationDate)
}

// RevokeCertificate is a paid mutator transaction binding the contract method 0xa4fd2992.
//
// Solidity: function revokeCertificate(bytes _finger, uint256 _revocationDate) returns()
func (_PgpOrchestr *PgpOrchestrTransactorSession) RevokeCertificate(_finger []byte, _revocationDate *big.Int) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.RevokeCertificate(&_PgpOrchestr.TransactOpts, _finger, _revocationDate)
}

// RevokeSignt is a paid mutator transaction binding the contract method 0xf8761755.
//
// Solidity: function revokeSignt(bytes _finger, bytes _signedFinger, uint256 _revocationDate) returns()
func (_PgpOrchestr *PgpOrchestrTransactor) RevokeSignt(opts *bind.TransactOpts, _finger []byte, _signedFinger []byte, _revocationDate *big.Int) (*types.Transaction, error) {
	return _PgpOrchestr.contract.Transact(opts, "revokeSignt", _finger, _signedFinger, _revocationDate)
}

// RevokeSignt is a paid mutator transaction binding the contract method 0xf8761755.
//
// Solidity: function revokeSignt(bytes _finger, bytes _signedFinger, uint256 _revocationDate) returns()
func (_PgpOrchestr *PgpOrchestrSession) RevokeSignt(_finger []byte, _signedFinger []byte, _revocationDate *big.Int) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.RevokeSignt(&_PgpOrchestr.TransactOpts, _finger, _signedFinger, _revocationDate)
}

// RevokeSignt is a paid mutator transaction binding the contract method 0xf8761755.
//
// Solidity: function revokeSignt(bytes _finger, bytes _signedFinger, uint256 _revocationDate) returns()
func (_PgpOrchestr *PgpOrchestrTransactorSession) RevokeSignt(_finger []byte, _signedFinger []byte, _revocationDate *big.Int) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.RevokeSignt(&_PgpOrchestr.TransactOpts, _finger, _signedFinger, _revocationDate)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_PgpOrchestr *PgpOrchestrTransactor) SetOwner(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.contract.Transact(opts, "setOwner", _addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_PgpOrchestr *PgpOrchestrSession) SetOwner(_addr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.SetOwner(&_PgpOrchestr.TransactOpts, _addr)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _addr) returns()
func (_PgpOrchestr *PgpOrchestrTransactorSession) SetOwner(_addr common.Address) (*types.Transaction, error) {
	return _PgpOrchestr.Contract.SetOwner(&_PgpOrchestr.TransactOpts, _addr)
}
