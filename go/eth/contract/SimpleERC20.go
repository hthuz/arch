// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SimpleERC20MetaData contains all meta data concerning the SimpleERC20 contract.
var SimpleERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516107c03803806107c0833981810160405281019061003191906100b7565b8060808181525050805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550506100e2565b5f80fd5b5f819050919050565b61009681610084565b81146100a0575f80fd5b50565b5f815190506100b18161008d565b92915050565b5f602082840312156100cc576100cb610080565b5b5f6100d9848285016100a3565b91505092915050565b6080516106c66100fa5f395f61017701526106c65ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c806306fdde031461006457806318160ddd14610082578063313ce567146100a057806370a08231146100be57806395d89b41146100ee578063a9059cbb1461010c575b5f80fd5b61006c61013c565b6040516100799190610454565b60405180910390f35b61008a610175565b604051610097919061048c565b60405180910390f35b6100a8610199565b6040516100b591906104c0565b60405180910390f35b6100d860048036038101906100d39190610537565b61019e565b6040516100e5919061048c565b60405180910390f35b6100f66101e3565b6040516101039190610454565b60405180910390f35b6101266004803603810190610121919061058c565b61021c565b60405161013391906105e4565b60405180910390f35b6040518060400160405280600b81526020017f53696d706c65455243323000000000000000000000000000000000000000000081525081565b7f000000000000000000000000000000000000000000000000000000000000000081565b600481565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6040518060400160405280600481526020017f534552430000000000000000000000000000000000000000000000000000000081525081565b5f815f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20541015610265575f80fd5b815f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20546102ad919061062a565b5f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550815f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054610335919061065d565b5f808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516103d2919061048c565b60405180910390a36001905092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610426826103e4565b61043081856103ee565b93506104408185602086016103fe565b6104498161040c565b840191505092915050565b5f6020820190508181035f83015261046c818461041c565b905092915050565b5f819050919050565b61048681610474565b82525050565b5f60208201905061049f5f83018461047d565b92915050565b5f60ff82169050919050565b6104ba816104a5565b82525050565b5f6020820190506104d35f8301846104b1565b92915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610506826104dd565b9050919050565b610516816104fc565b8114610520575f80fd5b50565b5f813590506105318161050d565b92915050565b5f6020828403121561054c5761054b6104d9565b5b5f61055984828501610523565b91505092915050565b61056b81610474565b8114610575575f80fd5b50565b5f8135905061058681610562565b92915050565b5f80604083850312156105a2576105a16104d9565b5b5f6105af85828601610523565b92505060206105c085828601610578565b9150509250929050565b5f8115159050919050565b6105de816105ca565b82525050565b5f6020820190506105f75f8301846105d5565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61063482610474565b915061063f83610474565b9250828203905081811115610657576106566105fd565b5b92915050565b5f61066782610474565b915061067283610474565b925082820190508082111561068a576106896105fd565b5b9291505056fea26469706673582212209bba7af2d3df7610f048399d91c9198f370d84940d70eb538a071c3abe80a93264736f6c634300081a0033",
}

// SimpleERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleERC20MetaData.ABI instead.
var SimpleERC20ABI = SimpleERC20MetaData.ABI

// SimpleERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleERC20MetaData.Bin instead.
var SimpleERC20Bin = SimpleERC20MetaData.Bin

// DeploySimpleERC20 deploys a new Ethereum contract, binding an instance of SimpleERC20 to it.
func DeploySimpleERC20(auth *bind.TransactOpts, backend bind.ContractBackend, total *big.Int) (common.Address, *types.Transaction, *SimpleERC20, error) {
	parsed, err := SimpleERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleERC20Bin), backend, total)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleERC20{SimpleERC20Caller: SimpleERC20Caller{contract: contract}, SimpleERC20Transactor: SimpleERC20Transactor{contract: contract}, SimpleERC20Filterer: SimpleERC20Filterer{contract: contract}}, nil
}

// SimpleERC20 is an auto generated Go binding around an Ethereum contract.
type SimpleERC20 struct {
	SimpleERC20Caller     // Read-only binding to the contract
	SimpleERC20Transactor // Write-only binding to the contract
	SimpleERC20Filterer   // Log filterer for contract events
}

// SimpleERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleERC20Session struct {
	Contract     *SimpleERC20      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleERC20CallerSession struct {
	Contract *SimpleERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SimpleERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleERC20TransactorSession struct {
	Contract     *SimpleERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SimpleERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleERC20Raw struct {
	Contract *SimpleERC20 // Generic contract binding to access the raw methods on
}

// SimpleERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleERC20CallerRaw struct {
	Contract *SimpleERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SimpleERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleERC20TransactorRaw struct {
	Contract *SimpleERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleERC20 creates a new instance of SimpleERC20, bound to a specific deployed contract.
func NewSimpleERC20(address common.Address, backend bind.ContractBackend) (*SimpleERC20, error) {
	contract, err := bindSimpleERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleERC20{SimpleERC20Caller: SimpleERC20Caller{contract: contract}, SimpleERC20Transactor: SimpleERC20Transactor{contract: contract}, SimpleERC20Filterer: SimpleERC20Filterer{contract: contract}}, nil
}

// NewSimpleERC20Caller creates a new read-only instance of SimpleERC20, bound to a specific deployed contract.
func NewSimpleERC20Caller(address common.Address, caller bind.ContractCaller) (*SimpleERC20Caller, error) {
	contract, err := bindSimpleERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleERC20Caller{contract: contract}, nil
}

// NewSimpleERC20Transactor creates a new write-only instance of SimpleERC20, bound to a specific deployed contract.
func NewSimpleERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SimpleERC20Transactor, error) {
	contract, err := bindSimpleERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleERC20Transactor{contract: contract}, nil
}

// NewSimpleERC20Filterer creates a new log filterer instance of SimpleERC20, bound to a specific deployed contract.
func NewSimpleERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SimpleERC20Filterer, error) {
	contract, err := bindSimpleERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleERC20Filterer{contract: contract}, nil
}

// bindSimpleERC20 binds a generic wrapper to an already deployed contract.
func bindSimpleERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleERC20 *SimpleERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleERC20.Contract.SimpleERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleERC20 *SimpleERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleERC20.Contract.SimpleERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleERC20 *SimpleERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleERC20.Contract.SimpleERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleERC20 *SimpleERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleERC20 *SimpleERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleERC20 *SimpleERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleERC20.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_SimpleERC20 *SimpleERC20Caller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleERC20.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_SimpleERC20 *SimpleERC20Session) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _SimpleERC20.Contract.BalanceOf(&_SimpleERC20.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_SimpleERC20 *SimpleERC20CallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _SimpleERC20.Contract.BalanceOf(&_SimpleERC20.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SimpleERC20 *SimpleERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SimpleERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SimpleERC20 *SimpleERC20Session) Decimals() (uint8, error) {
	return _SimpleERC20.Contract.Decimals(&_SimpleERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SimpleERC20 *SimpleERC20CallerSession) Decimals() (uint8, error) {
	return _SimpleERC20.Contract.Decimals(&_SimpleERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SimpleERC20 *SimpleERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SimpleERC20 *SimpleERC20Session) Name() (string, error) {
	return _SimpleERC20.Contract.Name(&_SimpleERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SimpleERC20 *SimpleERC20CallerSession) Name() (string, error) {
	return _SimpleERC20.Contract.Name(&_SimpleERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SimpleERC20 *SimpleERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SimpleERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SimpleERC20 *SimpleERC20Session) Symbol() (string, error) {
	return _SimpleERC20.Contract.Symbol(&_SimpleERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SimpleERC20 *SimpleERC20CallerSession) Symbol() (string, error) {
	return _SimpleERC20.Contract.Symbol(&_SimpleERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SimpleERC20 *SimpleERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SimpleERC20 *SimpleERC20Session) TotalSupply() (*big.Int, error) {
	return _SimpleERC20.Contract.TotalSupply(&_SimpleERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SimpleERC20 *SimpleERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _SimpleERC20.Contract.TotalSupply(&_SimpleERC20.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_SimpleERC20 *SimpleERC20Transactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _SimpleERC20.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_SimpleERC20 *SimpleERC20Session) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _SimpleERC20.Contract.Transfer(&_SimpleERC20.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_SimpleERC20 *SimpleERC20TransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _SimpleERC20.Contract.Transfer(&_SimpleERC20.TransactOpts, receiver, numTokens)
}

// SimpleERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SimpleERC20 contract.
type SimpleERC20TransferIterator struct {
	Event *SimpleERC20Transfer // Event containing the contract specifics and raw log

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
func (it *SimpleERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleERC20Transfer)
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
		it.Event = new(SimpleERC20Transfer)
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
func (it *SimpleERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleERC20Transfer represents a Transfer event raised by the SimpleERC20 contract.
type SimpleERC20Transfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_SimpleERC20 *SimpleERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SimpleERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SimpleERC20TransferIterator{contract: _SimpleERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_SimpleERC20 *SimpleERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SimpleERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SimpleERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleERC20Transfer)
				if err := _SimpleERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_SimpleERC20 *SimpleERC20Filterer) ParseTransfer(log types.Log) (*SimpleERC20Transfer, error) {
	event := new(SimpleERC20Transfer)
	if err := _SimpleERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
