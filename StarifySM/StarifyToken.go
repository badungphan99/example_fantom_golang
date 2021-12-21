// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StarifySM

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

// StarifySMMetaData contains all meta data concerning the StarifySM contract.
var StarifySMMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051610ae0380380610ae083398181016040528101906100329190610122565b33600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600281905550600254600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505061014f565b600080fd5b6000819050919050565b6100ff816100ec565b811461010a57600080fd5b50565b60008151905061011c816100f6565b92915050565b600060208284031215610138576101376100e7565b5b60006101468482850161010d565b91505092915050565b6109828061015e6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063313ce5671161005b578063313ce567146100ee57806395d89b411461010c578063a0712d681461012a578063a9059cbb1461015a5761007d565b806306fdde031461008257806318160ddd146100a0578063244e0c87146100be575b600080fd5b61008a61018a565b6040516100979190610678565b60405180910390f35b6100a86101c3565b6040516100b591906106b3565b60405180910390f35b6100d860048036038101906100d39190610731565b6101cd565b6040516100e591906106b3565b60405180910390f35b6100f6610215565b604051610103919061077a565b60405180910390f35b61011461021a565b6040516101219190610678565b60405180910390f35b610144600480360381019061013f91906107c1565b610253565b6040516101519190610809565b60405180910390f35b610174600480360381019061016f9190610824565b6103aa565b6040516101819190610809565b60405180910390f35b6040518060400160405280600c81526020017f53746172696679546f6b656e000000000000000000000000000000000000000081525081565b6000600254905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b601281565b6040518060400160405280600581526020017f534361736800000000000000000000000000000000000000000000000000000081525081565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102af57600080fd5b61032282600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461058c90919063ffffffff16565b600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555061039b8260025461058c90919063ffffffff16565b60028190555060019050919050565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548211156103f757600080fd5b610448826000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105b890919063ffffffff16565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506104db826000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461058c90919063ffffffff16565b6000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161057a91906106b3565b60405180910390a36001905092915050565b600080828461059b9190610893565b9050838110156105ae576105ad6108e9565b5b8091505092915050565b6000828211156105cb576105ca6108e9565b5b81836105d79190610918565b905092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106195780820151818401526020810190506105fe565b83811115610628576000848401525b50505050565b6000601f19601f8301169050919050565b600061064a826105df565b61065481856105ea565b93506106648185602086016105fb565b61066d8161062e565b840191505092915050565b60006020820190508181036000830152610692818461063f565b905092915050565b6000819050919050565b6106ad8161069a565b82525050565b60006020820190506106c860008301846106a4565b92915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006106fe826106d3565b9050919050565b61070e816106f3565b811461071957600080fd5b50565b60008135905061072b81610705565b92915050565b600060208284031215610747576107466106ce565b5b60006107558482850161071c565b91505092915050565b600060ff82169050919050565b6107748161075e565b82525050565b600060208201905061078f600083018461076b565b92915050565b61079e8161069a565b81146107a957600080fd5b50565b6000813590506107bb81610795565b92915050565b6000602082840312156107d7576107d66106ce565b5b60006107e5848285016107ac565b91505092915050565b60008115159050919050565b610803816107ee565b82525050565b600060208201905061081e60008301846107fa565b92915050565b6000806040838503121561083b5761083a6106ce565b5b60006108498582860161071c565b925050602061085a858286016107ac565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061089e8261069a565b91506108a98361069a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156108de576108dd610864565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b60006109238261069a565b915061092e8361069a565b92508282101561094157610940610864565b5b82820390509291505056fea26469706673582212207585518490be8c3531be20824ebfe7de6ff7c88dc2c320a6eb7aa71a44c0473e64736f6c634300080a0033",
}

// StarifySMABI is the input ABI used to generate the binding from.
// Deprecated: Use StarifySMMetaData.ABI instead.
var StarifySMABI = StarifySMMetaData.ABI

// StarifySMBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StarifySMMetaData.Bin instead.
var StarifySMBin = StarifySMMetaData.Bin

// DeployStarifySM deploys a new Ethereum contract, binding an instance of StarifySM to it.
func DeployStarifySM(auth *bind.TransactOpts, backend bind.ContractBackend, total *big.Int) (common.Address, *types.Transaction, *StarifySM, error) {
	parsed, err := StarifySMMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StarifySMBin), backend, total)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StarifySM{StarifySMCaller: StarifySMCaller{contract: contract}, StarifySMTransactor: StarifySMTransactor{contract: contract}, StarifySMFilterer: StarifySMFilterer{contract: contract}}, nil
}

// StarifySM is an auto generated Go binding around an Ethereum contract.
type StarifySM struct {
	StarifySMCaller     // Read-only binding to the contract
	StarifySMTransactor // Write-only binding to the contract
	StarifySMFilterer   // Log filterer for contract events
}

// StarifySMCaller is an auto generated read-only Go binding around an Ethereum contract.
type StarifySMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarifySMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StarifySMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarifySMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StarifySMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StarifySMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StarifySMSession struct {
	Contract     *StarifySM        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StarifySMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StarifySMCallerSession struct {
	Contract *StarifySMCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StarifySMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StarifySMTransactorSession struct {
	Contract     *StarifySMTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StarifySMRaw is an auto generated low-level Go binding around an Ethereum contract.
type StarifySMRaw struct {
	Contract *StarifySM // Generic contract binding to access the raw methods on
}

// StarifySMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StarifySMCallerRaw struct {
	Contract *StarifySMCaller // Generic read-only contract binding to access the raw methods on
}

// StarifySMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StarifySMTransactorRaw struct {
	Contract *StarifySMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStarifySM creates a new instance of StarifySM, bound to a specific deployed contract.
func NewStarifySM(address common.Address, backend bind.ContractBackend) (*StarifySM, error) {
	contract, err := bindStarifySM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StarifySM{StarifySMCaller: StarifySMCaller{contract: contract}, StarifySMTransactor: StarifySMTransactor{contract: contract}, StarifySMFilterer: StarifySMFilterer{contract: contract}}, nil
}

// NewStarifySMCaller creates a new read-only instance of StarifySM, bound to a specific deployed contract.
func NewStarifySMCaller(address common.Address, caller bind.ContractCaller) (*StarifySMCaller, error) {
	contract, err := bindStarifySM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StarifySMCaller{contract: contract}, nil
}

// NewStarifySMTransactor creates a new write-only instance of StarifySM, bound to a specific deployed contract.
func NewStarifySMTransactor(address common.Address, transactor bind.ContractTransactor) (*StarifySMTransactor, error) {
	contract, err := bindStarifySM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StarifySMTransactor{contract: contract}, nil
}

// NewStarifySMFilterer creates a new log filterer instance of StarifySM, bound to a specific deployed contract.
func NewStarifySMFilterer(address common.Address, filterer bind.ContractFilterer) (*StarifySMFilterer, error) {
	contract, err := bindStarifySM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StarifySMFilterer{contract: contract}, nil
}

// bindStarifySM binds a generic wrapper to an already deployed contract.
func bindStarifySM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StarifySMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarifySM *StarifySMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarifySM.Contract.StarifySMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarifySM *StarifySMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarifySM.Contract.StarifySMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarifySM *StarifySMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarifySM.Contract.StarifySMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StarifySM *StarifySMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StarifySM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StarifySM *StarifySMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StarifySM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StarifySM *StarifySMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StarifySM.Contract.contract.Transact(opts, method, params...)
}

// BalanceOF is a free data retrieval call binding the contract method 0x244e0c87.
//
// Solidity: function balanceOF(address tokenOwner) view returns(uint256)
func (_StarifySM *StarifySMCaller) BalanceOF(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StarifySM.contract.Call(opts, &out, "balanceOF", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOF is a free data retrieval call binding the contract method 0x244e0c87.
//
// Solidity: function balanceOF(address tokenOwner) view returns(uint256)
func (_StarifySM *StarifySMSession) BalanceOF(tokenOwner common.Address) (*big.Int, error) {
	return _StarifySM.Contract.BalanceOF(&_StarifySM.CallOpts, tokenOwner)
}

// BalanceOF is a free data retrieval call binding the contract method 0x244e0c87.
//
// Solidity: function balanceOF(address tokenOwner) view returns(uint256)
func (_StarifySM *StarifySMCallerSession) BalanceOF(tokenOwner common.Address) (*big.Int, error) {
	return _StarifySM.Contract.BalanceOF(&_StarifySM.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StarifySM *StarifySMCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StarifySM.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StarifySM *StarifySMSession) Decimals() (uint8, error) {
	return _StarifySM.Contract.Decimals(&_StarifySM.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StarifySM *StarifySMCallerSession) Decimals() (uint8, error) {
	return _StarifySM.Contract.Decimals(&_StarifySM.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StarifySM *StarifySMCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StarifySM.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StarifySM *StarifySMSession) Name() (string, error) {
	return _StarifySM.Contract.Name(&_StarifySM.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StarifySM *StarifySMCallerSession) Name() (string, error) {
	return _StarifySM.Contract.Name(&_StarifySM.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StarifySM *StarifySMCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StarifySM.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StarifySM *StarifySMSession) Symbol() (string, error) {
	return _StarifySM.Contract.Symbol(&_StarifySM.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StarifySM *StarifySMCallerSession) Symbol() (string, error) {
	return _StarifySM.Contract.Symbol(&_StarifySM.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StarifySM *StarifySMCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StarifySM.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StarifySM *StarifySMSession) TotalSupply() (*big.Int, error) {
	return _StarifySM.Contract.TotalSupply(&_StarifySM.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StarifySM *StarifySMCallerSession) TotalSupply() (*big.Int, error) {
	return _StarifySM.Contract.TotalSupply(&_StarifySM.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_StarifySM *StarifySMTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _StarifySM.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_StarifySM *StarifySMSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _StarifySM.Contract.Mint(&_StarifySM.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_StarifySM *StarifySMTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _StarifySM.Contract.Mint(&_StarifySM.TransactOpts, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_StarifySM *StarifySMTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StarifySM.contract.Transact(opts, "transfer", receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_StarifySM *StarifySMSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StarifySM.Contract.Transfer(&_StarifySM.TransactOpts, receiver, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 amount) returns(bool)
func (_StarifySM *StarifySMTransactorSession) Transfer(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StarifySM.Contract.Transfer(&_StarifySM.TransactOpts, receiver, amount)
}

// StarifySMTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StarifySM contract.
type StarifySMTransferIterator struct {
	Event *StarifySMTransfer // Event containing the contract specifics and raw log

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
func (it *StarifySMTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StarifySMTransfer)
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
		it.Event = new(StarifySMTransfer)
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
func (it *StarifySMTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StarifySMTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StarifySMTransfer represents a Transfer event raised by the StarifySM contract.
type StarifySMTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokens)
func (_StarifySM *StarifySMFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*StarifySMTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StarifySM.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StarifySMTransferIterator{contract: _StarifySM.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokens)
func (_StarifySM *StarifySMFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StarifySMTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StarifySM.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StarifySMTransfer)
				if err := _StarifySM.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _tokens)
func (_StarifySM *StarifySMFilterer) ParseTransfer(log types.Log) (*StarifySMTransfer, error) {
	event := new(StarifySMTransfer)
	if err := _StarifySM.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
