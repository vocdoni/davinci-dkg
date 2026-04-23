// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package golangtypes

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

// IDKGRegistryNodeKey is an auto generated low-level Go binding around an user-defined struct.
type IDKGRegistryNodeKey struct {
	Operator        common.Address
	PubX            *big.Int
	PubY            *big.Int
	Status          uint8
	LastActiveBlock uint64
}

// DKGRegistryMetaData contains all meta data concerning the DKGRegistry contract.
var DKGRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"inactivityWindow\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"INACTIVITY_WINDOW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGRegistry.NodeKey\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDKGRegistry.NodeStatus\"},{\"name\":\"lastActiveBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"heartbeat\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isActive\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"markActive\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nodeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reactivate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reap\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerKey\",\"inputs\":[{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"m\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateKey\",\"inputs\":[{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ManagerSet\",\"inputs\":[{\"name\":\"manager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeMarkedActive\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"atBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeReactivated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeReaped\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"lastActiveBlock\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManagerAlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ManagerNotSet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInactive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StillActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60c03461009357601f610b2538819003918201601f19168301916001600160401b038311848410176100975780849260209460405283398101031261009357516001600160401b0381168082036100935715610084576080523360a052604051610a7990816100ac823960805181818161044901526108fc015260a051816102040152f35b630eda9c3d60e31b5f5260045ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f3560e01c90816303ac08f6146108df575080633defb9621461085e57806341a89c571461070b5780634331ed1f146106e2578063481c6a75146106ba5780636da49b831461069457806384aa2df31461051f5780638af9f493146103e85780639d209048146102fb5780639f8a13d7146102ad578063d0ebdbe7146101dc578063d18611d6146100de5763f06f37bc146100ab575f80fd5b346100da5760203660031901126100da576004356001600160a01b03811681036100da576100d890610969565b005b5f80fd5b346100da575f3660031901126100da57335f525f602052600360405f200160ff81541660038110156101c8576002036101b957805468ffffffffffffffffff191643600881901b68ffffffffffffffff0016919091176001179091556001600160401b031660015467ffffffffffffffff60401b60016001600160401b038360401c160160401b169067ffffffffffffffff60401b19161760015560405190337ff979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe707420525f80a281525f516020610a245f395f51905f5260203392a2005b63442d617b60e11b5f5260045ffd5b634e487b7160e01b5f52602160045260245ffd5b346100da5760203660031901126100da576004356001600160a01b038116908190036100da577f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316330361029f5760025460ff8160a01c16610290578115610281576001600160a81b0319168117600160a01b176002557f60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa695f80a2005b63e6c4247b60e01b5f5260045ffd5b634294267360e11b5f5260045ffd5b6282b42960e81b5f5260045ffd5b346100da5760203660031901126100da576004356001600160a01b038116908190036100da575f525f60205260ff600360405f2001541660038110156101c857602090600160405191148152f35b346100da5760203660031901126100da576004356001600160a01b038116908190036100da575f608061032c610936565b82815282602082015282604082015282606082015201525f525f60205260405f20610355610936565b81546001600160a01b031681526001820154602082019081526002830154604083019081526003938401549360608401929060ff8616908110156101c85783526001600160401b03608085019560081c1685526040519360018060a01b0390511684525160208401525160408301525160038110156101c85760a0926001600160401b0391606084015251166080820152f35b346100da5760203660031901126100da576004356001600160a01b038116908190036100da57805f525f602052600360405f2001805460ff811660038110156101c857600103610510576001600160401b039060081c166001600160401b037f00000000000000000000000000000000000000000000000000000000000000001681018091116104fc574311156104ed5760206001600160401b037f17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f92600260ff1982541617815560015467ffffffffffffffff60401b5f19848360401c160160401b169067ffffffffffffffff60401b1916176001555460081c16604051908152a2005b63785bbc6d60e11b5f5260045ffd5b634e487b7160e01b5f52601160045260245ffd5b634065aaf160e11b5f5260045ffd5b346100da5761052d36610920565b908015801561068c575b61067d57335f525f60205260405f2091600383019260ff84541660038110156101c8571561066e576002818460018594015501556001600160401b0343169261059f84829068ffffffffffffffff0082549160081b169068ffffffffffffffff001916179055565b60ff81541660038110156101c857600214610602575b5060405191825260208201527f1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc60403392a26040519081525f516020610a245f395f51905f5260203392a2005b600160ff1982541617905560015467ffffffffffffffff60401b60016001600160401b038360401c160160401b169067ffffffffffffffff60401b191617600155337ff979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe707420525f80a2836105b5565b63aba4733960e01b5f5260045ffd5b630eda9c3d60e31b5f5260045ffd5b508115610537565b346100da575f3660031901126100da5760206001600160401b0360015416604051908152f35b346100da575f3660031901126100da576002546040516001600160a01b039091168152602090f35b346100da575f3660031901126100da5760206001600160401b0360015460401c16604051908152f35b346100da5761071936610920565b9080158015610856575b61067d57335f525f60205260405f2091600383019260ff84541660038110156101c8576108475760028291336bffffffffffffffffffffffff60a01b8254161781558460018201550155600160ff198454161783556107ab6001600160401b03431680949068ffffffffffffffff0082549160081b169068ffffffffffffffff001916179055565b6001546001600160401b036001818316011667ffffffffffffffff60401b60016001600160401b0383811986161760401c160160401b16916fffffffffffffffffffffffffffffffff1916171760015560405191825260208201527f99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f760403392a26040519081525f516020610a245f395f51905f5260203392a2005b630ea075bf60e21b5f5260045ffd5b508115610723565b346100da575f3660031901126100da57335f525f602052600360405f200160ff81541660038110156101c857600103610510576108c46001600160401b03431680929068ffffffffffffffff0082549160081b169068ffffffffffffffff001916179055565b6040519081525f516020610a245f395f51905f5260203392a2005b346100da575f3660031901126100da576020906001600160401b037f0000000000000000000000000000000000000000000000000000000000000000168152f35b60409060031901126100da576004359060243590565b6040519060a082018281106001600160401b0382111761095557604052565b634e487b7160e01b5f52604160045260245ffd5b6002546001600160a01b03168015610a14573303610a055760018060a01b0316805f525f602052600360405f200180549060ff821660038110156101c857600103610a00576001600160401b034381169260081c168214610a0057805468ffffffffffffffff001916600883901b68ffffffffffffffff00161790555f516020610a245f395f51905f5290602090604051908152a2565b505050565b63607e454560e11b5f5260045ffd5b6321f7ab5360e01b5f5260045ffdfe02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72ca2646970667358221220977d04b2f022b37540e3a5429e267311e37311beac65bebbc2db4abcbb1e1f1064736f6c634300081c0033",
}

// DKGRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGRegistryMetaData.ABI instead.
var DKGRegistryABI = DKGRegistryMetaData.ABI

// DKGRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGRegistryMetaData.Bin instead.
var DKGRegistryBin = DKGRegistryMetaData.Bin

// DeployDKGRegistry deploys a new Ethereum contract, binding an instance of DKGRegistry to it.
func DeployDKGRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, inactivityWindow uint64) (common.Address, *types.Transaction, *DKGRegistry, error) {
	parsed, err := DKGRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGRegistryBin), backend, inactivityWindow)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DKGRegistry{DKGRegistryCaller: DKGRegistryCaller{contract: contract}, DKGRegistryTransactor: DKGRegistryTransactor{contract: contract}, DKGRegistryFilterer: DKGRegistryFilterer{contract: contract}}, nil
}

// DKGRegistry is an auto generated Go binding around an Ethereum contract.
type DKGRegistry struct {
	DKGRegistryCaller     // Read-only binding to the contract
	DKGRegistryTransactor // Write-only binding to the contract
	DKGRegistryFilterer   // Log filterer for contract events
}

// DKGRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DKGRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DKGRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DKGRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DKGRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DKGRegistrySession struct {
	Contract     *DKGRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DKGRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DKGRegistryCallerSession struct {
	Contract *DKGRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DKGRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DKGRegistryTransactorSession struct {
	Contract     *DKGRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DKGRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DKGRegistryRaw struct {
	Contract *DKGRegistry // Generic contract binding to access the raw methods on
}

// DKGRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DKGRegistryCallerRaw struct {
	Contract *DKGRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DKGRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DKGRegistryTransactorRaw struct {
	Contract *DKGRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDKGRegistry creates a new instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistry(address common.Address, backend bind.ContractBackend) (*DKGRegistry, error) {
	contract, err := bindDKGRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DKGRegistry{DKGRegistryCaller: DKGRegistryCaller{contract: contract}, DKGRegistryTransactor: DKGRegistryTransactor{contract: contract}, DKGRegistryFilterer: DKGRegistryFilterer{contract: contract}}, nil
}

// NewDKGRegistryCaller creates a new read-only instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistryCaller(address common.Address, caller bind.ContractCaller) (*DKGRegistryCaller, error) {
	contract, err := bindDKGRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryCaller{contract: contract}, nil
}

// NewDKGRegistryTransactor creates a new write-only instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DKGRegistryTransactor, error) {
	contract, err := bindDKGRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryTransactor{contract: contract}, nil
}

// NewDKGRegistryFilterer creates a new log filterer instance of DKGRegistry, bound to a specific deployed contract.
func NewDKGRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DKGRegistryFilterer, error) {
	contract, err := bindDKGRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryFilterer{contract: contract}, nil
}

// bindDKGRegistry binds a generic wrapper to an already deployed contract.
func bindDKGRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DKGRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGRegistry *DKGRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGRegistry.Contract.DKGRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGRegistry *DKGRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.Contract.DKGRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGRegistry *DKGRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGRegistry.Contract.DKGRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DKGRegistry *DKGRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DKGRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DKGRegistry *DKGRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DKGRegistry *DKGRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DKGRegistry.Contract.contract.Transact(opts, method, params...)
}

// INACTIVITYWINDOW is a free data retrieval call binding the contract method 0x03ac08f6.
//
// Solidity: function INACTIVITY_WINDOW() view returns(uint64)
func (_DKGRegistry *DKGRegistryCaller) INACTIVITYWINDOW(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "INACTIVITY_WINDOW")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// INACTIVITYWINDOW is a free data retrieval call binding the contract method 0x03ac08f6.
//
// Solidity: function INACTIVITY_WINDOW() view returns(uint64)
func (_DKGRegistry *DKGRegistrySession) INACTIVITYWINDOW() (uint64, error) {
	return _DKGRegistry.Contract.INACTIVITYWINDOW(&_DKGRegistry.CallOpts)
}

// INACTIVITYWINDOW is a free data retrieval call binding the contract method 0x03ac08f6.
//
// Solidity: function INACTIVITY_WINDOW() view returns(uint64)
func (_DKGRegistry *DKGRegistryCallerSession) INACTIVITYWINDOW() (uint64, error) {
	return _DKGRegistry.Contract.INACTIVITYWINDOW(&_DKGRegistry.CallOpts)
}

// ActiveCount is a free data retrieval call binding the contract method 0x4331ed1f.
//
// Solidity: function activeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCaller) ActiveCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "activeCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ActiveCount is a free data retrieval call binding the contract method 0x4331ed1f.
//
// Solidity: function activeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistrySession) ActiveCount() (uint64, error) {
	return _DKGRegistry.Contract.ActiveCount(&_DKGRegistry.CallOpts)
}

// ActiveCount is a free data retrieval call binding the contract method 0x4331ed1f.
//
// Solidity: function activeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCallerSession) ActiveCount() (uint64, error) {
	return _DKGRegistry.Contract.ActiveCount(&_DKGRegistry.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8,uint64))
func (_DKGRegistry *DKGRegistryCaller) GetNode(opts *bind.CallOpts, operator common.Address) (IDKGRegistryNodeKey, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "getNode", operator)

	if err != nil {
		return *new(IDKGRegistryNodeKey), err
	}

	out0 := *abi.ConvertType(out[0], new(IDKGRegistryNodeKey)).(*IDKGRegistryNodeKey)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8,uint64))
func (_DKGRegistry *DKGRegistrySession) GetNode(operator common.Address) (IDKGRegistryNodeKey, error) {
	return _DKGRegistry.Contract.GetNode(&_DKGRegistry.CallOpts, operator)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8,uint64))
func (_DKGRegistry *DKGRegistryCallerSession) GetNode(operator common.Address) (IDKGRegistryNodeKey, error) {
	return _DKGRegistry.Contract.GetNode(&_DKGRegistry.CallOpts, operator)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address operator) view returns(bool)
func (_DKGRegistry *DKGRegistryCaller) IsActive(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "isActive", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address operator) view returns(bool)
func (_DKGRegistry *DKGRegistrySession) IsActive(operator common.Address) (bool, error) {
	return _DKGRegistry.Contract.IsActive(&_DKGRegistry.CallOpts, operator)
}

// IsActive is a free data retrieval call binding the contract method 0x9f8a13d7.
//
// Solidity: function isActive(address operator) view returns(bool)
func (_DKGRegistry *DKGRegistryCallerSession) IsActive(operator common.Address) (bool, error) {
	return _DKGRegistry.Contract.IsActive(&_DKGRegistry.CallOpts, operator)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_DKGRegistry *DKGRegistryCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_DKGRegistry *DKGRegistrySession) Manager() (common.Address, error) {
	return _DKGRegistry.Contract.Manager(&_DKGRegistry.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_DKGRegistry *DKGRegistryCallerSession) Manager() (common.Address, error) {
	return _DKGRegistry.Contract.Manager(&_DKGRegistry.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCaller) NodeCount(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _DKGRegistry.contract.Call(opts, &out, "nodeCount")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistrySession) NodeCount() (uint64, error) {
	return _DKGRegistry.Contract.NodeCount(&_DKGRegistry.CallOpts)
}

// NodeCount is a free data retrieval call binding the contract method 0x6da49b83.
//
// Solidity: function nodeCount() view returns(uint64)
func (_DKGRegistry *DKGRegistryCallerSession) NodeCount() (uint64, error) {
	return _DKGRegistry.Contract.NodeCount(&_DKGRegistry.CallOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_DKGRegistry *DKGRegistryTransactor) Heartbeat(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "heartbeat")
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_DKGRegistry *DKGRegistrySession) Heartbeat() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Heartbeat(&_DKGRegistry.TransactOpts)
}

// Heartbeat is a paid mutator transaction binding the contract method 0x3defb962.
//
// Solidity: function heartbeat() returns()
func (_DKGRegistry *DKGRegistryTransactorSession) Heartbeat() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Heartbeat(&_DKGRegistry.TransactOpts)
}

// MarkActive is a paid mutator transaction binding the contract method 0xf06f37bc.
//
// Solidity: function markActive(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactor) MarkActive(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "markActive", operator)
}

// MarkActive is a paid mutator transaction binding the contract method 0xf06f37bc.
//
// Solidity: function markActive(address operator) returns()
func (_DKGRegistry *DKGRegistrySession) MarkActive(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.MarkActive(&_DKGRegistry.TransactOpts, operator)
}

// MarkActive is a paid mutator transaction binding the contract method 0xf06f37bc.
//
// Solidity: function markActive(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) MarkActive(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.MarkActive(&_DKGRegistry.TransactOpts, operator)
}

// Reactivate is a paid mutator transaction binding the contract method 0xd18611d6.
//
// Solidity: function reactivate() returns()
func (_DKGRegistry *DKGRegistryTransactor) Reactivate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "reactivate")
}

// Reactivate is a paid mutator transaction binding the contract method 0xd18611d6.
//
// Solidity: function reactivate() returns()
func (_DKGRegistry *DKGRegistrySession) Reactivate() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reactivate(&_DKGRegistry.TransactOpts)
}

// Reactivate is a paid mutator transaction binding the contract method 0xd18611d6.
//
// Solidity: function reactivate() returns()
func (_DKGRegistry *DKGRegistryTransactorSession) Reactivate() (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reactivate(&_DKGRegistry.TransactOpts)
}

// Reap is a paid mutator transaction binding the contract method 0x8af9f493.
//
// Solidity: function reap(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactor) Reap(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "reap", operator)
}

// Reap is a paid mutator transaction binding the contract method 0x8af9f493.
//
// Solidity: function reap(address operator) returns()
func (_DKGRegistry *DKGRegistrySession) Reap(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reap(&_DKGRegistry.TransactOpts, operator)
}

// Reap is a paid mutator transaction binding the contract method 0x8af9f493.
//
// Solidity: function reap(address operator) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) Reap(operator common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.Reap(&_DKGRegistry.TransactOpts, operator)
}

// RegisterKey is a paid mutator transaction binding the contract method 0x41a89c57.
//
// Solidity: function registerKey(uint256 pubX, uint256 pubY) returns()
func (_DKGRegistry *DKGRegistryTransactor) RegisterKey(opts *bind.TransactOpts, pubX *big.Int, pubY *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "registerKey", pubX, pubY)
}

// RegisterKey is a paid mutator transaction binding the contract method 0x41a89c57.
//
// Solidity: function registerKey(uint256 pubX, uint256 pubY) returns()
func (_DKGRegistry *DKGRegistrySession) RegisterKey(pubX *big.Int, pubY *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.RegisterKey(&_DKGRegistry.TransactOpts, pubX, pubY)
}

// RegisterKey is a paid mutator transaction binding the contract method 0x41a89c57.
//
// Solidity: function registerKey(uint256 pubX, uint256 pubY) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) RegisterKey(pubX *big.Int, pubY *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.RegisterKey(&_DKGRegistry.TransactOpts, pubX, pubY)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address m) returns()
func (_DKGRegistry *DKGRegistryTransactor) SetManager(opts *bind.TransactOpts, m common.Address) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "setManager", m)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address m) returns()
func (_DKGRegistry *DKGRegistrySession) SetManager(m common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.SetManager(&_DKGRegistry.TransactOpts, m)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address m) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) SetManager(m common.Address) (*types.Transaction, error) {
	return _DKGRegistry.Contract.SetManager(&_DKGRegistry.TransactOpts, m)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x84aa2df3.
//
// Solidity: function updateKey(uint256 pubX, uint256 pubY) returns()
func (_DKGRegistry *DKGRegistryTransactor) UpdateKey(opts *bind.TransactOpts, pubX *big.Int, pubY *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.contract.Transact(opts, "updateKey", pubX, pubY)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x84aa2df3.
//
// Solidity: function updateKey(uint256 pubX, uint256 pubY) returns()
func (_DKGRegistry *DKGRegistrySession) UpdateKey(pubX *big.Int, pubY *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.UpdateKey(&_DKGRegistry.TransactOpts, pubX, pubY)
}

// UpdateKey is a paid mutator transaction binding the contract method 0x84aa2df3.
//
// Solidity: function updateKey(uint256 pubX, uint256 pubY) returns()
func (_DKGRegistry *DKGRegistryTransactorSession) UpdateKey(pubX *big.Int, pubY *big.Int) (*types.Transaction, error) {
	return _DKGRegistry.Contract.UpdateKey(&_DKGRegistry.TransactOpts, pubX, pubY)
}

// DKGRegistryManagerSetIterator is returned from FilterManagerSet and is used to iterate over the raw logs and unpacked data for ManagerSet events raised by the DKGRegistry contract.
type DKGRegistryManagerSetIterator struct {
	Event *DKGRegistryManagerSet // Event containing the contract specifics and raw log

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
func (it *DKGRegistryManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryManagerSet)
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
		it.Event = new(DKGRegistryManagerSet)
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
func (it *DKGRegistryManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryManagerSet represents a ManagerSet event raised by the DKGRegistry contract.
type DKGRegistryManagerSet struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerSet is a free log retrieval operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_DKGRegistry *DKGRegistryFilterer) FilterManagerSet(opts *bind.FilterOpts, manager []common.Address) (*DKGRegistryManagerSetIterator, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "ManagerSet", managerRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryManagerSetIterator{contract: _DKGRegistry.contract, event: "ManagerSet", logs: logs, sub: sub}, nil
}

// WatchManagerSet is a free log subscription operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_DKGRegistry *DKGRegistryFilterer) WatchManagerSet(opts *bind.WatchOpts, sink chan<- *DKGRegistryManagerSet, manager []common.Address) (event.Subscription, error) {

	var managerRule []interface{}
	for _, managerItem := range manager {
		managerRule = append(managerRule, managerItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "ManagerSet", managerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryManagerSet)
				if err := _DKGRegistry.contract.UnpackLog(event, "ManagerSet", log); err != nil {
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

// ParseManagerSet is a log parse operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address indexed manager)
func (_DKGRegistry *DKGRegistryFilterer) ParseManagerSet(log types.Log) (*DKGRegistryManagerSet, error) {
	event := new(DKGRegistryManagerSet)
	if err := _DKGRegistry.contract.UnpackLog(event, "ManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeMarkedActiveIterator is returned from FilterNodeMarkedActive and is used to iterate over the raw logs and unpacked data for NodeMarkedActive events raised by the DKGRegistry contract.
type DKGRegistryNodeMarkedActiveIterator struct {
	Event *DKGRegistryNodeMarkedActive // Event containing the contract specifics and raw log

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
func (it *DKGRegistryNodeMarkedActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeMarkedActive)
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
		it.Event = new(DKGRegistryNodeMarkedActive)
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
func (it *DKGRegistryNodeMarkedActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeMarkedActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeMarkedActive represents a NodeMarkedActive event raised by the DKGRegistry contract.
type DKGRegistryNodeMarkedActive struct {
	Operator common.Address
	AtBlock  uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeMarkedActive is a free log retrieval operation binding the contract event 0x02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72c.
//
// Solidity: event NodeMarkedActive(address indexed operator, uint64 atBlock)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeMarkedActive(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeMarkedActiveIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeMarkedActive", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeMarkedActiveIterator{contract: _DKGRegistry.contract, event: "NodeMarkedActive", logs: logs, sub: sub}, nil
}

// WatchNodeMarkedActive is a free log subscription operation binding the contract event 0x02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72c.
//
// Solidity: event NodeMarkedActive(address indexed operator, uint64 atBlock)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeMarkedActive(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeMarkedActive, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeMarkedActive", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeMarkedActive)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeMarkedActive", log); err != nil {
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

// ParseNodeMarkedActive is a log parse operation binding the contract event 0x02c36b03f66c867a89d996a43b2ea1f9c0e5740578642d17ef1b3d259073e72c.
//
// Solidity: event NodeMarkedActive(address indexed operator, uint64 atBlock)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeMarkedActive(log types.Log) (*DKGRegistryNodeMarkedActive, error) {
	event := new(DKGRegistryNodeMarkedActive)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeMarkedActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeReactivatedIterator is returned from FilterNodeReactivated and is used to iterate over the raw logs and unpacked data for NodeReactivated events raised by the DKGRegistry contract.
type DKGRegistryNodeReactivatedIterator struct {
	Event *DKGRegistryNodeReactivated // Event containing the contract specifics and raw log

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
func (it *DKGRegistryNodeReactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeReactivated)
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
		it.Event = new(DKGRegistryNodeReactivated)
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
func (it *DKGRegistryNodeReactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeReactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeReactivated represents a NodeReactivated event raised by the DKGRegistry contract.
type DKGRegistryNodeReactivated struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeReactivated is a free log retrieval operation binding the contract event 0xf979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe70742052.
//
// Solidity: event NodeReactivated(address indexed operator)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeReactivated(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeReactivatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeReactivated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeReactivatedIterator{contract: _DKGRegistry.contract, event: "NodeReactivated", logs: logs, sub: sub}, nil
}

// WatchNodeReactivated is a free log subscription operation binding the contract event 0xf979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe70742052.
//
// Solidity: event NodeReactivated(address indexed operator)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeReactivated(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeReactivated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeReactivated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeReactivated)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeReactivated", log); err != nil {
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

// ParseNodeReactivated is a log parse operation binding the contract event 0xf979d653049f5a10edd541959ecb5c2ced8fd1b0adaefc8fd66744fe70742052.
//
// Solidity: event NodeReactivated(address indexed operator)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeReactivated(log types.Log) (*DKGRegistryNodeReactivated, error) {
	event := new(DKGRegistryNodeReactivated)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeReactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeReapedIterator is returned from FilterNodeReaped and is used to iterate over the raw logs and unpacked data for NodeReaped events raised by the DKGRegistry contract.
type DKGRegistryNodeReapedIterator struct {
	Event *DKGRegistryNodeReaped // Event containing the contract specifics and raw log

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
func (it *DKGRegistryNodeReapedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeReaped)
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
		it.Event = new(DKGRegistryNodeReaped)
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
func (it *DKGRegistryNodeReapedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeReapedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeReaped represents a NodeReaped event raised by the DKGRegistry contract.
type DKGRegistryNodeReaped struct {
	Operator        common.Address
	LastActiveBlock uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNodeReaped is a free log retrieval operation binding the contract event 0x17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f.
//
// Solidity: event NodeReaped(address indexed operator, uint64 lastActiveBlock)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeReaped(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeReapedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeReaped", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeReapedIterator{contract: _DKGRegistry.contract, event: "NodeReaped", logs: logs, sub: sub}, nil
}

// WatchNodeReaped is a free log subscription operation binding the contract event 0x17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f.
//
// Solidity: event NodeReaped(address indexed operator, uint64 lastActiveBlock)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeReaped(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeReaped, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeReaped", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeReaped)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeReaped", log); err != nil {
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

// ParseNodeReaped is a log parse operation binding the contract event 0x17b35aacc7270dcc7c9993688488c4a6267c1ca2e4ab73b83c6411855a54376f.
//
// Solidity: event NodeReaped(address indexed operator, uint64 lastActiveBlock)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeReaped(log types.Log) (*DKGRegistryNodeReaped, error) {
	event := new(DKGRegistryNodeReaped)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeReaped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeRegisteredIterator is returned from FilterNodeRegistered and is used to iterate over the raw logs and unpacked data for NodeRegistered events raised by the DKGRegistry contract.
type DKGRegistryNodeRegisteredIterator struct {
	Event *DKGRegistryNodeRegistered // Event containing the contract specifics and raw log

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
func (it *DKGRegistryNodeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeRegistered)
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
		it.Event = new(DKGRegistryNodeRegistered)
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
func (it *DKGRegistryNodeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeRegistered represents a NodeRegistered event raised by the DKGRegistry contract.
type DKGRegistryNodeRegistered struct {
	Operator common.Address
	PubX     *big.Int
	PubY     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeRegistered is a free log retrieval operation binding the contract event 0x99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f7.
//
// Solidity: event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeRegistered(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeRegisteredIterator{contract: _DKGRegistry.contract, event: "NodeRegistered", logs: logs, sub: sub}, nil
}

// WatchNodeRegistered is a free log subscription operation binding the contract event 0x99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f7.
//
// Solidity: event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeRegistered(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeRegistered)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
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

// ParseNodeRegistered is a log parse operation binding the contract event 0x99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f7.
//
// Solidity: event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeRegistered(log types.Log) (*DKGRegistryNodeRegistered, error) {
	event := new(DKGRegistryNodeRegistered)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DKGRegistryNodeUpdatedIterator is returned from FilterNodeUpdated and is used to iterate over the raw logs and unpacked data for NodeUpdated events raised by the DKGRegistry contract.
type DKGRegistryNodeUpdatedIterator struct {
	Event *DKGRegistryNodeUpdated // Event containing the contract specifics and raw log

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
func (it *DKGRegistryNodeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DKGRegistryNodeUpdated)
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
		it.Event = new(DKGRegistryNodeUpdated)
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
func (it *DKGRegistryNodeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DKGRegistryNodeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DKGRegistryNodeUpdated represents a NodeUpdated event raised by the DKGRegistry contract.
type DKGRegistryNodeUpdated struct {
	Operator common.Address
	PubX     *big.Int
	PubY     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNodeUpdated is a free log retrieval operation binding the contract event 0x1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc.
//
// Solidity: event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) FilterNodeUpdated(opts *bind.FilterOpts, operator []common.Address) (*DKGRegistryNodeUpdatedIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.FilterLogs(opts, "NodeUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return &DKGRegistryNodeUpdatedIterator{contract: _DKGRegistry.contract, event: "NodeUpdated", logs: logs, sub: sub}, nil
}

// WatchNodeUpdated is a free log subscription operation binding the contract event 0x1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc.
//
// Solidity: event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) WatchNodeUpdated(opts *bind.WatchOpts, sink chan<- *DKGRegistryNodeUpdated, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DKGRegistry.contract.WatchLogs(opts, "NodeUpdated", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DKGRegistryNodeUpdated)
				if err := _DKGRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
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

// ParseNodeUpdated is a log parse operation binding the contract event 0x1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc.
//
// Solidity: event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY)
func (_DKGRegistry *DKGRegistryFilterer) ParseNodeUpdated(log types.Log) (*DKGRegistryNodeUpdated, error) {
	event := new(DKGRegistryNodeUpdated)
	if err := _DKGRegistry.contract.UnpackLog(event, "NodeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
