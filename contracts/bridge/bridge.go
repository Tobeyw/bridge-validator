// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// BridgeEventMetaData contains all meta data concerning the BridgeEvent contract.
var BridgeEventMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_withdrawalHash\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"root\",\"type\":\"bytes\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ad\",\"type\":\"address\"}],\"name\":\"test\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeEventABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeEventMetaData.ABI instead.
var BridgeEventABI = BridgeEventMetaData.ABI

// BridgeEvent is an auto generated Go binding around an Ethereum contract.
type BridgeEvent struct {
	BridgeEventCaller     // Read-only binding to the contract
	BridgeEventTransactor // Write-only binding to the contract
	BridgeEventFilterer   // Log filterer for contract events
}

// BridgeEventCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeEventCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEventTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeEventTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEventFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeEventFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeEventSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeEventSession struct {
	Contract     *BridgeEvent      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeEventCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeEventCallerSession struct {
	Contract *BridgeEventCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeEventTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeEventTransactorSession struct {
	Contract     *BridgeEventTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeEventRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeEventRaw struct {
	Contract *BridgeEvent // Generic contract binding to access the raw methods on
}

// BridgeEventCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeEventCallerRaw struct {
	Contract *BridgeEventCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeEventTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeEventTransactorRaw struct {
	Contract *BridgeEventTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeEvent creates a new instance of BridgeEvent, bound to a specific deployed contract.
func NewBridgeEvent(address common.Address, backend bind.ContractBackend) (*BridgeEvent, error) {
	contract, err := bindBridgeEvent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeEvent{BridgeEventCaller: BridgeEventCaller{contract: contract}, BridgeEventTransactor: BridgeEventTransactor{contract: contract}, BridgeEventFilterer: BridgeEventFilterer{contract: contract}}, nil
}

// NewBridgeEventCaller creates a new read-only instance of BridgeEvent, bound to a specific deployed contract.
func NewBridgeEventCaller(address common.Address, caller bind.ContractCaller) (*BridgeEventCaller, error) {
	contract, err := bindBridgeEvent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEventCaller{contract: contract}, nil
}

// NewBridgeEventTransactor creates a new write-only instance of BridgeEvent, bound to a specific deployed contract.
func NewBridgeEventTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeEventTransactor, error) {
	contract, err := bindBridgeEvent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeEventTransactor{contract: contract}, nil
}

// NewBridgeEventFilterer creates a new log filterer instance of BridgeEvent, bound to a specific deployed contract.
func NewBridgeEventFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeEventFilterer, error) {
	contract, err := bindBridgeEvent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeEventFilterer{contract: contract}, nil
}

// bindBridgeEvent binds a generic wrapper to an already deployed contract.
func bindBridgeEvent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeEventMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEvent *BridgeEventRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEvent.Contract.BridgeEventCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEvent *BridgeEventRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEvent.Contract.BridgeEventTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEvent *BridgeEventRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEvent.Contract.BridgeEventTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeEvent *BridgeEventCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeEvent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeEvent *BridgeEventTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeEvent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeEvent *BridgeEventTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeEvent.Contract.contract.Transact(opts, method, params...)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_BridgeEvent *BridgeEventTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _BridgeEvent.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_BridgeEvent *BridgeEventSession) Store(num *big.Int) (*types.Transaction, error) {
	return _BridgeEvent.Contract.Store(&_BridgeEvent.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_BridgeEvent *BridgeEventTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _BridgeEvent.Contract.Store(&_BridgeEvent.TransactOpts, num)
}

// Test is a paid mutator transaction binding the contract method 0xbb29998e.
//
// Solidity: function test(address ad) returns()
func (_BridgeEvent *BridgeEventTransactor) Test(opts *bind.TransactOpts, ad common.Address) (*types.Transaction, error) {
	return _BridgeEvent.contract.Transact(opts, "test", ad)
}

// Test is a paid mutator transaction binding the contract method 0xbb29998e.
//
// Solidity: function test(address ad) returns()
func (_BridgeEvent *BridgeEventSession) Test(ad common.Address) (*types.Transaction, error) {
	return _BridgeEvent.Contract.Test(&_BridgeEvent.TransactOpts, ad)
}

// Test is a paid mutator transaction binding the contract method 0xbb29998e.
//
// Solidity: function test(address ad) returns()
func (_BridgeEvent *BridgeEventTransactorSession) Test(ad common.Address) (*types.Transaction, error) {
	return _BridgeEvent.Contract.Test(&_BridgeEvent.TransactOpts, ad)
}

// BridgeEventDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the BridgeEvent contract.
type BridgeEventDepositIterator struct {
	Event *BridgeEventDeposit // Event containing the contract specifics and raw log

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
func (it *BridgeEventDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEventDeposit)
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
		it.Event = new(BridgeEventDeposit)
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
func (it *BridgeEventDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEventDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEventDeposit represents a Deposit event raised by the BridgeEvent contract.
type BridgeEventDeposit struct {
	Nonce  *big.Int
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 _nonce, address _to, uint256 _amount)
func (_BridgeEvent *BridgeEventFilterer) FilterDeposit(opts *bind.FilterOpts) (*BridgeEventDepositIterator, error) {

	logs, sub, err := _BridgeEvent.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &BridgeEventDepositIterator{contract: _BridgeEvent.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 _nonce, address _to, uint256 _amount)
func (_BridgeEvent *BridgeEventFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeEventDeposit) (event.Subscription, error) {

	logs, sub, err := _BridgeEvent.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEventDeposit)
				if err := _BridgeEvent.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xeaa18152488ce5959073c9c79c88ca90b3d96c00de1f118cfaad664c3dab06b9.
//
// Solidity: event Deposit(uint256 _nonce, address _to, uint256 _amount)
func (_BridgeEvent *BridgeEventFilterer) ParseDeposit(log types.Log) (*BridgeEventDeposit, error) {
	event := new(BridgeEventDeposit)
	if err := _BridgeEvent.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeEventWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the BridgeEvent contract.
type BridgeEventWithdrawalIterator struct {
	Event *BridgeEventWithdrawal // Event containing the contract specifics and raw log

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
func (it *BridgeEventWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeEventWithdrawal)
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
		it.Event = new(BridgeEventWithdrawal)
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
func (it *BridgeEventWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeEventWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeEventWithdrawal represents a Withdrawal event raised by the BridgeEvent contract.
type BridgeEventWithdrawal struct {
	Nonce          *big.Int
	From           common.Address
	To             common.Address
	Amount         *big.Int
	WithdrawalHash []byte
	Root           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xcd767d411218346986e952388bf7dcadd340e99678799eca0e827fa2e08a568c.
//
// Solidity: event Withdrawal(uint256 _nonce, address _from, address _to, uint256 _amount, bytes _withdrawalHash, bytes root)
func (_BridgeEvent *BridgeEventFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*BridgeEventWithdrawalIterator, error) {

	logs, sub, err := _BridgeEvent.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &BridgeEventWithdrawalIterator{contract: _BridgeEvent.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xcd767d411218346986e952388bf7dcadd340e99678799eca0e827fa2e08a568c.
//
// Solidity: event Withdrawal(uint256 _nonce, address _from, address _to, uint256 _amount, bytes _withdrawalHash, bytes root)
func (_BridgeEvent *BridgeEventFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *BridgeEventWithdrawal) (event.Subscription, error) {

	logs, sub, err := _BridgeEvent.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeEventWithdrawal)
				if err := _BridgeEvent.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0xcd767d411218346986e952388bf7dcadd340e99678799eca0e827fa2e08a568c.
//
// Solidity: event Withdrawal(uint256 _nonce, address _from, address _to, uint256 _amount, bytes _withdrawalHash, bytes root)
func (_BridgeEvent *BridgeEventFilterer) ParseWithdrawal(log types.Log) (*BridgeEventWithdrawal, error) {
	event := new(BridgeEventWithdrawal)
	if err := _BridgeEvent.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
