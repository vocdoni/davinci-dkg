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
	Operator common.Address
	PubX     *big.Int
	PubY     *big.Int
	Status   uint8
}

// DKGRegistryMetaData contains all meta data concerning the DKGRegistry contract.
var DKGRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getNode\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIDKGRegistry.NodeKey\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIDKGRegistry.NodeStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerKey\",\"inputs\":[{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateKey\",\"inputs\":[{\"name\":\"pubX\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"NodeRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NodeUpdated\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubX\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"pubY\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyRegistered\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidKey\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotRegistered\",\"inputs\":[]}]",
	Bin: "0x60808060405234601557610386908161001a8239f35b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c806341a89c571461021d5780636da49b83146101f657806384aa2df31461014857639d20904814610045575f80fd5b346101445760203660031901126101445760043573ffffffffffffffffffffffffffffffffffffffff8116809103610144575f606061008261031c565b82815282602082015282604082015201525f525f60205260405f206100a561031c565b9073ffffffffffffffffffffffffffffffffffffffff81541682526001810154906020830191825260ff600360028301549260408601938452015416916060840192600381101561013057835273ffffffffffffffffffffffffffffffffffffffff604051945116845251602084015251604083015251906003821015610130576080916060820152f35b634e487b7160e01b5f52602160045260245ffd5b5f80fd5b346101445761015636610306565b811580156101ee575b6101df57335f525f60205260405f20600381019060ff825416600381101561013057156101d057600281856001869401550155600160ff1982541617905560405191825260208201527f1e2215a8512058e371c99f86c2731c45755267c9d5fb9eb3c911230fa9b55cfc60403392a2005b63aba4733960e01b5f5260045ffd5b630eda9c3d60e31b5f5260045ffd5b50801561015f565b34610144575f36600319011261014457602067ffffffffffffffff60015416604051908152f35b346101445761022b36610306565b811580156102fe575b6101df57335f525f60205260405f20600381019060ff8254166003811015610130576102ef5760028391337fffffffffffffffffffffffff00000000000000000000000000000000000000008254161781558560018201550155600160ff1982541617905560015467ffffffffffffffff600181831601169067ffffffffffffffff19161760015560405191825260208201527f99140a41575033d78b1016979e49f1b8a4943ef274d75edf0dac1bc3ccbce5f760403392a2005b630ea075bf60e21b5f5260045ffd5b508015610234565b6040906003190112610144576004359060243590565b604051906080820182811067ffffffffffffffff82111761033c57604052565b634e487b7160e01b5f52604160045260245ffdfea2646970667358221220cab529a4fca8334da49dabaf58ded9e5a0e8e64375d862c060c224f29b41d3bd64736f6c634300081c0033",
}

// DKGRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DKGRegistryMetaData.ABI instead.
var DKGRegistryABI = DKGRegistryMetaData.ABI

// DKGRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DKGRegistryMetaData.Bin instead.
var DKGRegistryBin = DKGRegistryMetaData.Bin

// DeployDKGRegistry deploys a new Ethereum contract, binding an instance of DKGRegistry to it.
func DeployDKGRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DKGRegistry, error) {
	parsed, err := DKGRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DKGRegistryBin), backend)
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

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8))
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
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8))
func (_DKGRegistry *DKGRegistrySession) GetNode(operator common.Address) (IDKGRegistryNodeKey, error) {
	return _DKGRegistry.Contract.GetNode(&_DKGRegistry.CallOpts, operator)
}

// GetNode is a free data retrieval call binding the contract method 0x9d209048.
//
// Solidity: function getNode(address operator) view returns((address,uint256,uint256,uint8))
func (_DKGRegistry *DKGRegistryCallerSession) GetNode(operator common.Address) (IDKGRegistryNodeKey, error) {
	return _DKGRegistry.Contract.GetNode(&_DKGRegistry.CallOpts, operator)
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
