// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CreditLetterABI is the input ABI used to generate the binding from.
const CreditLetterABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"status\",\"type\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setStatusPaid\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"status\",\"type\":\"bool\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"setStatusApproved\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInfo\",\"outputs\":[{\"name\":\"issuerVal\",\"type\":\"address\"},{\"name\":\"holderVal\",\"type\":\"address\"},{\"name\":\"acceptorVal\",\"type\":\"address\"},{\"name\":\"issuanceTimeVal\",\"type\":\"uint256\"},{\"name\":\"interestRateVal\",\"type\":\"uint256\"},{\"name\":\"creditVal\",\"type\":\"uint256\"},{\"name\":\"approved\",\"type\":\"bool\"},{\"name\":\"paid\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllTransferLogs\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"resetCreditAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"issuerVal\",\"type\":\"address\"},{\"name\":\"holderVal\",\"type\":\"address\"},{\"name\":\"acceptorVal\",\"type\":\"address\"},{\"name\":\"issuanceTimeVal\",\"type\":\"uint256\"},{\"name\":\"interestRateVal\",\"type\":\"uint256\"},{\"name\":\"creditVal\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"latestStatus\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"HolderChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"original\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CreditChanged\",\"type\":\"event\"}]"

// CreditLetterBin is the compiled bytecode used for deploying new contracts.
var CreditLetterBin = "0x60806040526000600660006101000a81548160ff0219169083151502179055506000600660016101000a81548160ff0219169083151502179055506000600755600160085534801561005057600080fd5b5060405160c0806110778339810180604052810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291905050508573ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16148061010a57508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b15156101a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001807f4f6e6c7920697373756572206f72206163636570746f722063616e206973737581526020017f65206869732f686572206f776e206c657474657200000000000000000000000081525060400191505060405180910390fd5b856000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600381905550816004819055508060058190555060098590806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a839080600181540180825580915050906001820390600052602060002001600090919290919091505550505050505050610d55806103226000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631c9242801461007d57806358007b86146100b65780635a9b0b89146100ef5780635dfd3c98146101d7578063a9059cbb1461028b578063bfe389ed146102d8575b600080fd5b34801561008957600080fd5b506100b46004803603810190808035151590602001909291908035906020019092919050505061030f565b005b3480156100c257600080fd5b506100ed6004803603810190808035151590602001909291908035906020019092919050505061044e565b005b3480156100fb57600080fd5b5061010461060a565b604051808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815260200184815260200183151515158152602001821515151581526020019850505050505050505060405180910390f35b3480156101e357600080fd5b506101ec6106c1565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610233578082015181840152602081019050610218565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561027557808201518184015260208101905061025a565b5050505090500194505050505060405180910390f35b34801561029757600080fd5b506102d6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108ec565b005b3480156102e457600080fd5b5061030d6004803603810190808035906020019092919080359060200190929190505050610b93565b005b60008111151561031e57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4f6e6c79206163636570746f722063616e20706179000000000000000000000081525060200191505060405180910390fd5b81600660016101000a81548160ff0219169083151502179055507f1ecf1b9cee605f7dc61806c2f40429418dbc06abbb7cfe58d1b5480f084af9a060085483836040518084815260200183151515158152602001828152602001935050505060405180910390a15050565b60008111151561045d57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061050557506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561059f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001807f4f6e6c79206163636570746f72206f72206973737565722063616e206170707281526020017f6f7665000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b81600660006101000a81548160ff0219169083151502179055507f1ecf1b9cee605f7dc61806c2f40429418dbc06abbb7cfe58d1b5480f084af9a060075483836040518084815260200183151515158152602001828152602001935050505060405180910390a15050565b6000806000806000806000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600354600454600554600660009054906101000a900460ff16600660019054906101000a900460ff16975097509750975097509750975097509091929394959697565b6060806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061076c5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610806576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f4f6e6c7920697373756572206f72206163636570746f722063616e206368656381526020017f6b207472616e73666572206c6f6773000000000000000000000000000000000081525060400191505060405180910390fd5b6009600a8180548060200260200160405190810160405280929190818152602001828054801561088b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610841575b50505050509150808054806020026020016040519081016040528092919081815260200182805480156108dd57602002820191906000526020600020905b8154815260200190600101908083116108c9575b50505050509050915091509091565b8173ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561094957600080fd5b60008111151561095857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a1d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f4f6e6c7920686f6c6465722063616e207472616e73666572000000000000000081525060200191505060405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060093390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a8190806001815401808255809150509060018203906000526020600020016000909192909190915055507f0b208f30897dd37bbda3db1da3cf917bce605c53cdab4f7be8eaac32c8082f10338383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161480610c3b5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b1515610cd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f4f6e6c7920697373756572206f72206163636570746f722063616e207265736581526020017f742063726564697420616d6f756e74000000000000000000000000000000000081525060400191505060405180910390fd5b816005819055507fbe251dfecb445e45c89b4940942865904cf0efeea8fb1e4aa0f340e8370c03cd600554838360405180848152602001838152602001828152602001935050505060405180910390a150505600a165627a7a723058206cb1decaec00d899522b4d4c6d65da8f648f68771ad00179dfc9ae57dbdfebc90029"

// DeployCreditLetter deploys a new contract, binding an instance of CreditLetter to it.
func DeployCreditLetter(auth *bind.TransactOpts, backend bind.ContractBackend, issuerVal common.Address, holderVal common.Address, acceptorVal common.Address, issuanceTimeVal *big.Int, interestRateVal *big.Int, creditVal *big.Int) (common.Address, *types.Transaction, *CreditLetter, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditLetterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CreditLetterBin), backend, issuerVal, holderVal, acceptorVal, issuanceTimeVal, interestRateVal, creditVal)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreditLetter{CreditLetterCaller: CreditLetterCaller{contract: contract}, CreditLetterTransactor: CreditLetterTransactor{contract: contract}, CreditLetterFilterer: CreditLetterFilterer{contract: contract}}, nil
}

func AsyncDeployCreditLetter(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, issuerVal common.Address, holderVal common.Address, acceptorVal common.Address, issuanceTimeVal *big.Int, interestRateVal *big.Int, creditVal *big.Int) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditLetterABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(CreditLetterBin), backend, issuerVal, holderVal, acceptorVal, issuanceTimeVal, interestRateVal, creditVal)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CreditLetter is an auto generated Go binding around a Solidity contract.
type CreditLetter struct {
	CreditLetterCaller     // Read-only binding to the contract
	CreditLetterTransactor // Write-only binding to the contract
	CreditLetterFilterer   // Log filterer for contract events
}

// CreditLetterCaller is an auto generated read-only Go binding around a Solidity contract.
type CreditLetterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditLetterTransactor is an auto generated write-only Go binding around a Solidity contract.
type CreditLetterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditLetterFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type CreditLetterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditLetterSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type CreditLetterSession struct {
	Contract     *CreditLetter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CreditLetterCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type CreditLetterCallerSession struct {
	Contract *CreditLetterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CreditLetterTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type CreditLetterTransactorSession struct {
	Contract     *CreditLetterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CreditLetterRaw is an auto generated low-level Go binding around a Solidity contract.
type CreditLetterRaw struct {
	Contract *CreditLetter // Generic contract binding to access the raw methods on
}

// CreditLetterCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type CreditLetterCallerRaw struct {
	Contract *CreditLetterCaller // Generic read-only contract binding to access the raw methods on
}

// CreditLetterTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type CreditLetterTransactorRaw struct {
	Contract *CreditLetterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditLetter creates a new instance of CreditLetter, bound to a specific deployed contract.
func NewCreditLetter(address common.Address, backend bind.ContractBackend) (*CreditLetter, error) {
	contract, err := bindCreditLetter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditLetter{CreditLetterCaller: CreditLetterCaller{contract: contract}, CreditLetterTransactor: CreditLetterTransactor{contract: contract}, CreditLetterFilterer: CreditLetterFilterer{contract: contract}}, nil
}

// NewCreditLetterCaller creates a new read-only instance of CreditLetter, bound to a specific deployed contract.
func NewCreditLetterCaller(address common.Address, caller bind.ContractCaller) (*CreditLetterCaller, error) {
	contract, err := bindCreditLetter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditLetterCaller{contract: contract}, nil
}

// NewCreditLetterTransactor creates a new write-only instance of CreditLetter, bound to a specific deployed contract.
func NewCreditLetterTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditLetterTransactor, error) {
	contract, err := bindCreditLetter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditLetterTransactor{contract: contract}, nil
}

// NewCreditLetterFilterer creates a new log filterer instance of CreditLetter, bound to a specific deployed contract.
func NewCreditLetterFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditLetterFilterer, error) {
	contract, err := bindCreditLetter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditLetterFilterer{contract: contract}, nil
}

// bindCreditLetter binds a generic wrapper to an already deployed contract.
func bindCreditLetter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditLetterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditLetter *CreditLetterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreditLetter.Contract.CreditLetterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditLetter *CreditLetterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.CreditLetterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditLetter *CreditLetterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.CreditLetterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditLetter *CreditLetterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreditLetter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditLetter *CreditLetterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditLetter *CreditLetterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.contract.Transact(opts, method, params...)
}

// GetAllTransferLogs is a free data retrieval call binding the contract method 0x5dfd3c98.
//
// Solidity: function getAllTransferLogs() constant returns(address[], uint256[])
func (_CreditLetter *CreditLetterCaller) GetAllTransferLogs(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _CreditLetter.contract.Call(opts, out, "getAllTransferLogs")
	return *ret0, *ret1, err
}

// GetAllTransferLogs is a free data retrieval call binding the contract method 0x5dfd3c98.
//
// Solidity: function getAllTransferLogs() constant returns(address[], uint256[])
func (_CreditLetter *CreditLetterSession) GetAllTransferLogs() ([]common.Address, []*big.Int, error) {
	return _CreditLetter.Contract.GetAllTransferLogs(&_CreditLetter.CallOpts)
}

// GetAllTransferLogs is a free data retrieval call binding the contract method 0x5dfd3c98.
//
// Solidity: function getAllTransferLogs() constant returns(address[], uint256[])
func (_CreditLetter *CreditLetterCallerSession) GetAllTransferLogs() ([]common.Address, []*big.Int, error) {
	return _CreditLetter.Contract.GetAllTransferLogs(&_CreditLetter.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(address issuerVal, address holderVal, address acceptorVal, uint256 issuanceTimeVal, uint256 interestRateVal, uint256 creditVal, bool approved, bool paid)
func (_CreditLetter *CreditLetterCaller) GetInfo(opts *bind.CallOpts) (struct {
	IssuerVal       common.Address
	HolderVal       common.Address
	AcceptorVal     common.Address
	IssuanceTimeVal *big.Int
	InterestRateVal *big.Int
	CreditVal       *big.Int
	Approved        bool
	Paid            bool
}, error) {
	ret := new(struct {
		IssuerVal       common.Address
		HolderVal       common.Address
		AcceptorVal     common.Address
		IssuanceTimeVal *big.Int
		InterestRateVal *big.Int
		CreditVal       *big.Int
		Approved        bool
		Paid            bool
	})
	out := ret
	err := _CreditLetter.contract.Call(opts, out, "getInfo")
	return *ret, err
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(address issuerVal, address holderVal, address acceptorVal, uint256 issuanceTimeVal, uint256 interestRateVal, uint256 creditVal, bool approved, bool paid)
func (_CreditLetter *CreditLetterSession) GetInfo() (struct {
	IssuerVal       common.Address
	HolderVal       common.Address
	AcceptorVal     common.Address
	IssuanceTimeVal *big.Int
	InterestRateVal *big.Int
	CreditVal       *big.Int
	Approved        bool
	Paid            bool
}, error) {
	return _CreditLetter.Contract.GetInfo(&_CreditLetter.CallOpts)
}

// GetInfo is a free data retrieval call binding the contract method 0x5a9b0b89.
//
// Solidity: function getInfo() constant returns(address issuerVal, address holderVal, address acceptorVal, uint256 issuanceTimeVal, uint256 interestRateVal, uint256 creditVal, bool approved, bool paid)
func (_CreditLetter *CreditLetterCallerSession) GetInfo() (struct {
	IssuerVal       common.Address
	HolderVal       common.Address
	AcceptorVal     common.Address
	IssuanceTimeVal *big.Int
	InterestRateVal *big.Int
	CreditVal       *big.Int
	Approved        bool
	Paid            bool
}, error) {
	return _CreditLetter.Contract.GetInfo(&_CreditLetter.CallOpts)
}

// ResetCreditAmount is a paid mutator transaction binding the contract method 0xbfe389ed.
//
// Solidity: function resetCreditAmount(uint256 amount, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactor) ResetCreditAmount(opts *bind.TransactOpts, amount *big.Int, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.contract.Transact(opts, "resetCreditAmount", amount, timestamp)
}

func (_CreditLetter *CreditLetterTransactor) AsyncResetCreditAmount(handler func(*types.Receipt, error), opts *bind.TransactOpts, amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.contract.AsyncTransact(opts, handler, "resetCreditAmount", amount, timestamp)
}

// ResetCreditAmount is a paid mutator transaction binding the contract method 0xbfe389ed.
//
// Solidity: function resetCreditAmount(uint256 amount, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterSession) ResetCreditAmount(amount *big.Int, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.ResetCreditAmount(&_CreditLetter.TransactOpts, amount, timestamp)
}

func (_CreditLetter *CreditLetterSession) AsyncResetCreditAmount(handler func(*types.Receipt, error), amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncResetCreditAmount(handler, &_CreditLetter.TransactOpts, amount, timestamp)
}

// ResetCreditAmount is a paid mutator transaction binding the contract method 0xbfe389ed.
//
// Solidity: function resetCreditAmount(uint256 amount, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactorSession) ResetCreditAmount(amount *big.Int, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.ResetCreditAmount(&_CreditLetter.TransactOpts, amount, timestamp)
}

func (_CreditLetter *CreditLetterTransactorSession) AsyncResetCreditAmount(handler func(*types.Receipt, error), amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncResetCreditAmount(handler, &_CreditLetter.TransactOpts, amount, timestamp)
}

// SetStatusApproved is a paid mutator transaction binding the contract method 0x58007b86.
//
// Solidity: function setStatusApproved(bool status, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactor) SetStatusApproved(opts *bind.TransactOpts, status bool, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.contract.Transact(opts, "setStatusApproved", status, timestamp)
}

func (_CreditLetter *CreditLetterTransactor) AsyncSetStatusApproved(handler func(*types.Receipt, error), opts *bind.TransactOpts, status bool, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.contract.AsyncTransact(opts, handler, "setStatusApproved", status, timestamp)
}

// SetStatusApproved is a paid mutator transaction binding the contract method 0x58007b86.
//
// Solidity: function setStatusApproved(bool status, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterSession) SetStatusApproved(status bool, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.SetStatusApproved(&_CreditLetter.TransactOpts, status, timestamp)
}

func (_CreditLetter *CreditLetterSession) AsyncSetStatusApproved(handler func(*types.Receipt, error), status bool, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncSetStatusApproved(handler, &_CreditLetter.TransactOpts, status, timestamp)
}

// SetStatusApproved is a paid mutator transaction binding the contract method 0x58007b86.
//
// Solidity: function setStatusApproved(bool status, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactorSession) SetStatusApproved(status bool, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.SetStatusApproved(&_CreditLetter.TransactOpts, status, timestamp)
}

func (_CreditLetter *CreditLetterTransactorSession) AsyncSetStatusApproved(handler func(*types.Receipt, error), status bool, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncSetStatusApproved(handler, &_CreditLetter.TransactOpts, status, timestamp)
}

// SetStatusPaid is a paid mutator transaction binding the contract method 0x1c924280.
//
// Solidity: function setStatusPaid(bool status, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactor) SetStatusPaid(opts *bind.TransactOpts, status bool, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.contract.Transact(opts, "setStatusPaid", status, timestamp)
}

func (_CreditLetter *CreditLetterTransactor) AsyncSetStatusPaid(handler func(*types.Receipt, error), opts *bind.TransactOpts, status bool, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.contract.AsyncTransact(opts, handler, "setStatusPaid", status, timestamp)
}

// SetStatusPaid is a paid mutator transaction binding the contract method 0x1c924280.
//
// Solidity: function setStatusPaid(bool status, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterSession) SetStatusPaid(status bool, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.SetStatusPaid(&_CreditLetter.TransactOpts, status, timestamp)
}

func (_CreditLetter *CreditLetterSession) AsyncSetStatusPaid(handler func(*types.Receipt, error), status bool, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncSetStatusPaid(handler, &_CreditLetter.TransactOpts, status, timestamp)
}

// SetStatusPaid is a paid mutator transaction binding the contract method 0x1c924280.
//
// Solidity: function setStatusPaid(bool status, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactorSession) SetStatusPaid(status bool, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.SetStatusPaid(&_CreditLetter.TransactOpts, status, timestamp)
}

func (_CreditLetter *CreditLetterTransactorSession) AsyncSetStatusPaid(handler func(*types.Receipt, error), status bool, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncSetStatusPaid(handler, &_CreditLetter.TransactOpts, status, timestamp)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactor) Transfer(opts *bind.TransactOpts, to common.Address, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.contract.Transact(opts, "transfer", to, timestamp)
}

func (_CreditLetter *CreditLetterTransactor) AsyncTransfer(handler func(*types.Receipt, error), opts *bind.TransactOpts, to common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.contract.AsyncTransact(opts, handler, "transfer", to, timestamp)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterSession) Transfer(to common.Address, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.Transfer(&_CreditLetter.TransactOpts, to, timestamp)
}

func (_CreditLetter *CreditLetterSession) AsyncTransfer(handler func(*types.Receipt, error), to common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncTransfer(handler, &_CreditLetter.TransactOpts, to, timestamp)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 timestamp) returns()
func (_CreditLetter *CreditLetterTransactorSession) Transfer(to common.Address, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetter.Contract.Transfer(&_CreditLetter.TransactOpts, to, timestamp)
}

func (_CreditLetter *CreditLetterTransactorSession) AsyncTransfer(handler func(*types.Receipt, error), to common.Address, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetter.Contract.AsyncTransfer(handler, &_CreditLetter.TransactOpts, to, timestamp)
}

// CreditLetterCreditChangedIterator is returned from FilterCreditChanged and is used to iterate over the raw logs and unpacked data for CreditChanged events raised by the CreditLetter contract.
type CreditLetterCreditChangedIterator struct {
	Event *CreditLetterCreditChanged // Event containing the contract specifics and raw log

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
func (it *CreditLetterCreditChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditLetterCreditChanged)
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
		it.Event = new(CreditLetterCreditChanged)
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
func (it *CreditLetterCreditChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditLetterCreditChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditLetterCreditChanged represents a CreditChanged event raised by the CreditLetter contract.
type CreditLetterCreditChanged struct {
	Original  *big.Int
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreditChanged is a free log retrieval operation binding the contract event 0xbe251dfecb445e45c89b4940942865904cf0efeea8fb1e4aa0f340e8370c03cd.
//
// Solidity: event CreditChanged(uint256 original, uint256 amount, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) FilterCreditChanged(opts *bind.FilterOpts) (*CreditLetterCreditChangedIterator, error) {

	logs, sub, err := _CreditLetter.contract.FilterLogs(opts, "CreditChanged")
	if err != nil {
		return nil, err
	}
	return &CreditLetterCreditChangedIterator{contract: _CreditLetter.contract, event: "CreditChanged", logs: logs, sub: sub}, nil
}

// WatchCreditChanged is a free log subscription operation binding the contract event 0xbe251dfecb445e45c89b4940942865904cf0efeea8fb1e4aa0f340e8370c03cd.
//
// Solidity: event CreditChanged(uint256 original, uint256 amount, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) WatchCreditChanged(opts *bind.WatchOpts, sink chan<- *CreditLetterCreditChanged) (event.Subscription, error) {

	logs, sub, err := _CreditLetter.contract.WatchLogs(opts, "CreditChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditLetterCreditChanged)
				if err := _CreditLetter.contract.UnpackLog(event, "CreditChanged", log); err != nil {
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

// ParseCreditChanged is a log parse operation binding the contract event 0xbe251dfecb445e45c89b4940942865904cf0efeea8fb1e4aa0f340e8370c03cd.
//
// Solidity: event CreditChanged(uint256 original, uint256 amount, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) ParseCreditChanged(log types.Log) (*CreditLetterCreditChanged, error) {
	event := new(CreditLetterCreditChanged)
	if err := _CreditLetter.contract.UnpackLog(event, "CreditChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditLetterHolderChangedIterator is returned from FilterHolderChanged and is used to iterate over the raw logs and unpacked data for HolderChanged events raised by the CreditLetter contract.
type CreditLetterHolderChangedIterator struct {
	Event *CreditLetterHolderChanged // Event containing the contract specifics and raw log

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
func (it *CreditLetterHolderChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditLetterHolderChanged)
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
		it.Event = new(CreditLetterHolderChanged)
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
func (it *CreditLetterHolderChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditLetterHolderChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditLetterHolderChanged represents a HolderChanged event raised by the CreditLetter contract.
type CreditLetterHolderChanged struct {
	From      common.Address
	To        common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHolderChanged is a free log retrieval operation binding the contract event 0x0b208f30897dd37bbda3db1da3cf917bce605c53cdab4f7be8eaac32c8082f10.
//
// Solidity: event HolderChanged(address from, address to, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) FilterHolderChanged(opts *bind.FilterOpts) (*CreditLetterHolderChangedIterator, error) {

	logs, sub, err := _CreditLetter.contract.FilterLogs(opts, "HolderChanged")
	if err != nil {
		return nil, err
	}
	return &CreditLetterHolderChangedIterator{contract: _CreditLetter.contract, event: "HolderChanged", logs: logs, sub: sub}, nil
}

// WatchHolderChanged is a free log subscription operation binding the contract event 0x0b208f30897dd37bbda3db1da3cf917bce605c53cdab4f7be8eaac32c8082f10.
//
// Solidity: event HolderChanged(address from, address to, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) WatchHolderChanged(opts *bind.WatchOpts, sink chan<- *CreditLetterHolderChanged) (event.Subscription, error) {

	logs, sub, err := _CreditLetter.contract.WatchLogs(opts, "HolderChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditLetterHolderChanged)
				if err := _CreditLetter.contract.UnpackLog(event, "HolderChanged", log); err != nil {
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

// ParseHolderChanged is a log parse operation binding the contract event 0x0b208f30897dd37bbda3db1da3cf917bce605c53cdab4f7be8eaac32c8082f10.
//
// Solidity: event HolderChanged(address from, address to, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) ParseHolderChanged(log types.Log) (*CreditLetterHolderChanged, error) {
	event := new(CreditLetterHolderChanged)
	if err := _CreditLetter.contract.UnpackLog(event, "HolderChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// CreditLetterStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the CreditLetter contract.
type CreditLetterStatusChangedIterator struct {
	Event *CreditLetterStatusChanged // Event containing the contract specifics and raw log

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
func (it *CreditLetterStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditLetterStatusChanged)
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
		it.Event = new(CreditLetterStatusChanged)
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
func (it *CreditLetterStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditLetterStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditLetterStatusChanged represents a StatusChanged event raised by the CreditLetter contract.
type CreditLetterStatusChanged struct {
	Id           *big.Int
	LatestStatus bool
	Timestamp    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0x1ecf1b9cee605f7dc61806c2f40429418dbc06abbb7cfe58d1b5480f084af9a0.
//
// Solidity: event StatusChanged(uint256 id, bool latestStatus, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) FilterStatusChanged(opts *bind.FilterOpts) (*CreditLetterStatusChangedIterator, error) {

	logs, sub, err := _CreditLetter.contract.FilterLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return &CreditLetterStatusChangedIterator{contract: _CreditLetter.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0x1ecf1b9cee605f7dc61806c2f40429418dbc06abbb7cfe58d1b5480f084af9a0.
//
// Solidity: event StatusChanged(uint256 id, bool latestStatus, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *CreditLetterStatusChanged) (event.Subscription, error) {

	logs, sub, err := _CreditLetter.contract.WatchLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditLetterStatusChanged)
				if err := _CreditLetter.contract.UnpackLog(event, "StatusChanged", log); err != nil {
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

// ParseStatusChanged is a log parse operation binding the contract event 0x1ecf1b9cee605f7dc61806c2f40429418dbc06abbb7cfe58d1b5480f084af9a0.
//
// Solidity: event StatusChanged(uint256 id, bool latestStatus, uint256 timestamp)
func (_CreditLetter *CreditLetterFilterer) ParseStatusChanged(log types.Log) (*CreditLetterStatusChanged, error) {
	event := new(CreditLetterStatusChanged)
	if err := _CreditLetter.contract.UnpackLog(event, "StatusChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}
