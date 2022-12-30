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
)

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"NewTransactionExchange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"NewTransactionReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"receiveToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"sendTokenToContract\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610aea806100606000396000f3fe6080604052600436106100295760003560e01c8063a53ec8e11461002e578063c722c7031461004a575b600080fd5b610048600480360381019061004391906105ce565b610066565b005b610064600480360381019061005f919061063d565b6101a0565b005b60011515610075838584610227565b1515146100b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ae906106c7565b60405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff16826040516100dd90610718565b60006040518083038185875af1925050503d806000811461011a576040519150601f19603f3d011682016040523d82523d6000602084013e61011f565b606091505b5050905080610163576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161015a90610779565b60405180910390fd5b7f4f0222c6a15da55907b401d3d25463ff0c3b38e50ca641d12c43aa4a85c9b9c58360405161019291906107a8565b60405180910390a150505050565b6509184e72a0003410156101e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e090610835565b60405180910390fd5b7f3434032ea91078551b2d5e559bd385096c32e1d79b46cb703ff913f5bf10657433823460405161021c93929190610864565b60405180910390a150565b60008061023485846102a6565b90506000610241826102d9565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166102848287610309565b73ffffffffffffffffffffffffffffffffffffffff1614925050509392505050565b600082826040516020016102bb929190610904565b60405160208183030381529060405280519060200120905092915050565b6000816040516020016102ec91906109b2565b604051602081830303815290604052805190602001209050919050565b60008060008061031885610378565b925092509250600186828585604051600081526020016040526040516103419493929190610a03565b6020604051602081039080840390855afa158015610363573d6000803e3d6000fd5b50505060206040510351935050505092915050565b600080600060418451146103c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b890610a94565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610447826103fe565b810181811067ffffffffffffffff821117156104665761046561040f565b5b80604052505050565b60006104796103e0565b9050610485828261043e565b919050565b600067ffffffffffffffff8211156104a5576104a461040f565b5b6104ae826103fe565b9050602081019050919050565b82818337600083830152505050565b60006104dd6104d88461048a565b61046f565b9050828152602081018484840111156104f9576104f86103f9565b5b6105048482856104bb565b509392505050565b600082601f830112610521576105206103f4565b5b81356105318482602086016104ca565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006105658261053a565b9050919050565b6105758161055a565b811461058057600080fd5b50565b6000813590506105928161056c565b92915050565b6000819050919050565b6105ab81610598565b81146105b657600080fd5b50565b6000813590506105c8816105a2565b92915050565b6000806000606084860312156105e7576105e66103ea565b5b600084013567ffffffffffffffff811115610605576106046103ef565b5b6106118682870161050c565b935050602061062286828701610583565b9250506040610633868287016105b9565b9150509250925092565b600060208284031215610653576106526103ea565b5b600061066184828501610583565b91505092915050565b600082825260208201905092915050565b7f41757468656e7469636174696f6e206661696c65640000000000000000000000600082015250565b60006106b160158361066a565b91506106bc8261067b565b602082019050919050565b600060208201905081810360008301526106e0816106a4565b9050919050565b600081905092915050565b50565b60006107026000836106e7565b915061070d826106f2565b600082019050919050565b6000610723826106f5565b9150819050919050565b7f4661696c656420746f2073656e64000000000000000000000000000000000000600082015250565b6000610763600e8361066a565b915061076e8261072d565b602082019050919050565b6000602082019050818103600083015261079281610756565b9050919050565b6107a28161055a565b82525050565b60006020820190506107bd6000830184610799565b92915050565b7f4661696c656420746f2073656e642c20726571756972652076616c7565203e3d60008201527f20302e3030303031000000000000000000000000000000000000000000000000602082015250565b600061081f60288361066a565b915061082a826107c3565b604082019050919050565b6000602082019050818103600083015261084e81610812565b9050919050565b61085e81610598565b82525050565b60006060820190506108796000830186610799565b6108866020830185610799565b6108936040830184610855565b949350505050565b60008160601b9050919050565b60006108b38261089b565b9050919050565b60006108c5826108a8565b9050919050565b6108dd6108d88261055a565b6108ba565b82525050565b6000819050919050565b6108fe6108f982610598565b6108e3565b82525050565b600061091082856108cc565b60148201915061092082846108ed565b6020820191508190509392505050565b600081905092915050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b6000610971601c83610930565b915061097c8261093b565b601c82019050919050565b6000819050919050565b6000819050919050565b6109ac6109a782610987565b610991565b82525050565b60006109bd82610964565b91506109c9828461099b565b60208201915081905092915050565b6109e181610987565b82525050565b600060ff82169050919050565b6109fd816109e7565b82525050565b6000608082019050610a1860008301876109d8565b610a2560208301866109f4565b610a3260408301856109d8565b610a3f60608301846109d8565b95945050505050565b7f696e76616c6964207369676e6174757265206c656e6774680000000000000000600082015250565b6000610a7e60188361066a565b9150610a8982610a48565b602082019050919050565b60006020820190508181036000830152610aad81610a71565b905091905056fea264697066735822122034b84156c135024f95c7a1dc0d6c457bb14724f5d37bd9c5b979283194a8e07064736f6c63430008110033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0xa53ec8e1.
//
// Solidity: function receiveToken(bytes _signature, address _address, uint256 _value) payable returns()
func (_Contract *ContractTransactor) ReceiveToken(opts *bind.TransactOpts, _signature []byte, _address common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "receiveToken", _signature, _address, _value)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0xa53ec8e1.
//
// Solidity: function receiveToken(bytes _signature, address _address, uint256 _value) payable returns()
func (_Contract *ContractSession) ReceiveToken(_signature []byte, _address common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReceiveToken(&_Contract.TransactOpts, _signature, _address, _value)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0xa53ec8e1.
//
// Solidity: function receiveToken(bytes _signature, address _address, uint256 _value) payable returns()
func (_Contract *ContractTransactorSession) ReceiveToken(_signature []byte, _address common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReceiveToken(&_Contract.TransactOpts, _signature, _address, _value)
}

// SendTokenToContract is a paid mutator transaction binding the contract method 0xc722c703.
//
// Solidity: function sendTokenToContract(address _address) payable returns()
func (_Contract *ContractTransactor) SendTokenToContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sendTokenToContract", _address)
}

// SendTokenToContract is a paid mutator transaction binding the contract method 0xc722c703.
//
// Solidity: function sendTokenToContract(address _address) payable returns()
func (_Contract *ContractSession) SendTokenToContract(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SendTokenToContract(&_Contract.TransactOpts, _address)
}

// SendTokenToContract is a paid mutator transaction binding the contract method 0xc722c703.
//
// Solidity: function sendTokenToContract(address _address) payable returns()
func (_Contract *ContractTransactorSession) SendTokenToContract(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SendTokenToContract(&_Contract.TransactOpts, _address)
}

// ContractNewTransactionExchangeIterator is returned from FilterNewTransactionExchange and is used to iterate over the raw logs and unpacked data for NewTransactionExchange events raised by the Contract contract.
type ContractNewTransactionExchangeIterator struct {
	Event *ContractNewTransactionExchange // Event containing the contract specifics and raw log

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
func (it *ContractNewTransactionExchangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewTransactionExchange)
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
		it.Event = new(ContractNewTransactionExchange)
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
func (it *ContractNewTransactionExchangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewTransactionExchangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewTransactionExchange represents a NewTransactionExchange event raised by the Contract contract.
type ContractNewTransactionExchange struct {
	Arg0 common.Address
	Arg1 common.Address
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTransactionExchange is a free log retrieval operation binding the contract event 0x3434032ea91078551b2d5e559bd385096c32e1d79b46cb703ff913f5bf106574.
//
// Solidity: event NewTransactionExchange(address arg0, address arg1, uint256 arg2)
func (_Contract *ContractFilterer) FilterNewTransactionExchange(opts *bind.FilterOpts) (*ContractNewTransactionExchangeIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewTransactionExchange")
	if err != nil {
		return nil, err
	}
	return &ContractNewTransactionExchangeIterator{contract: _Contract.contract, event: "NewTransactionExchange", logs: logs, sub: sub}, nil
}

// WatchNewTransactionExchange is a free log subscription operation binding the contract event 0x3434032ea91078551b2d5e559bd385096c32e1d79b46cb703ff913f5bf106574.
//
// Solidity: event NewTransactionExchange(address arg0, address arg1, uint256 arg2)
func (_Contract *ContractFilterer) WatchNewTransactionExchange(opts *bind.WatchOpts, sink chan<- *ContractNewTransactionExchange) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewTransactionExchange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewTransactionExchange)
				if err := _Contract.contract.UnpackLog(event, "NewTransactionExchange", log); err != nil {
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

// ParseNewTransactionExchange is a log parse operation binding the contract event 0x3434032ea91078551b2d5e559bd385096c32e1d79b46cb703ff913f5bf106574.
//
// Solidity: event NewTransactionExchange(address arg0, address arg1, uint256 arg2)
func (_Contract *ContractFilterer) ParseNewTransactionExchange(log types.Log) (*ContractNewTransactionExchange, error) {
	event := new(ContractNewTransactionExchange)
	if err := _Contract.contract.UnpackLog(event, "NewTransactionExchange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractNewTransactionReceivedIterator is returned from FilterNewTransactionReceived and is used to iterate over the raw logs and unpacked data for NewTransactionReceived events raised by the Contract contract.
type ContractNewTransactionReceivedIterator struct {
	Event *ContractNewTransactionReceived // Event containing the contract specifics and raw log

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
func (it *ContractNewTransactionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractNewTransactionReceived)
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
		it.Event = new(ContractNewTransactionReceived)
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
func (it *ContractNewTransactionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractNewTransactionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractNewTransactionReceived represents a NewTransactionReceived event raised by the Contract contract.
type ContractNewTransactionReceived struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNewTransactionReceived is a free log retrieval operation binding the contract event 0x4f0222c6a15da55907b401d3d25463ff0c3b38e50ca641d12c43aa4a85c9b9c5.
//
// Solidity: event NewTransactionReceived(address arg0)
func (_Contract *ContractFilterer) FilterNewTransactionReceived(opts *bind.FilterOpts) (*ContractNewTransactionReceivedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "NewTransactionReceived")
	if err != nil {
		return nil, err
	}
	return &ContractNewTransactionReceivedIterator{contract: _Contract.contract, event: "NewTransactionReceived", logs: logs, sub: sub}, nil
}

// WatchNewTransactionReceived is a free log subscription operation binding the contract event 0x4f0222c6a15da55907b401d3d25463ff0c3b38e50ca641d12c43aa4a85c9b9c5.
//
// Solidity: event NewTransactionReceived(address arg0)
func (_Contract *ContractFilterer) WatchNewTransactionReceived(opts *bind.WatchOpts, sink chan<- *ContractNewTransactionReceived) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "NewTransactionReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractNewTransactionReceived)
				if err := _Contract.contract.UnpackLog(event, "NewTransactionReceived", log); err != nil {
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

// ParseNewTransactionReceived is a log parse operation binding the contract event 0x4f0222c6a15da55907b401d3d25463ff0c3b38e50ca641d12c43aa4a85c9b9c5.
//
// Solidity: event NewTransactionReceived(address arg0)
func (_Contract *ContractFilterer) ParseNewTransactionReceived(log types.Log) (*ContractNewTransactionReceived, error) {
	event := new(ContractNewTransactionReceived)
	if err := _Contract.contract.UnpackLog(event, "NewTransactionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
