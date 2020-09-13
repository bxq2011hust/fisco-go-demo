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

// CreditLetterFactoryABI is the input ABI used to generate the binding from.
const CreditLetterFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"issuer\",\"type\":\"address\"},{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"acceptor\",\"type\":\"address\"},{\"name\":\"issuanceTime\",\"type\":\"uint256\"},{\"name\":\"interestRate\",\"type\":\"uint256\"},{\"name\":\"credit\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"originalAddress\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"split\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"CreateLog\",\"type\":\"event\"}]"

// CreditLetterFactoryBin is the compiled bytecode used for deploying new contracts.
var CreditLetterFactoryBin = "0x608060405234801561001057600080fd5b506116c4806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680637f368f65146100515780638afbc1ed14610132575b600080fd5b34801561005d57600080fd5b506100f0600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190803590602001909291905050506101c9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561013e57600080fd5b50610187600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919080359060200190929190505050610323565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000808787878787876101da610611565b808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020018281526020019650505050505050604051809103906000f0801580156102a7573d6000803e3d6000fd5b5090507f4481db8423b8eae544e932e436f55884c6487e4d3ca295a3ca992ecf84ecc6958186604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1809150509695505050505050565b6000806000806000806000808a96508673ffffffffffffffffffffffffffffffffffffffff16635a9b0b896040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161010060405180830381600087803b15801561039757600080fd5b505af11580156103ab573d6000803e3d6000fd5b505050506040513d6101008110156103c257600080fd5b810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190505050505096509650509550955095508673ffffffffffffffffffffffffffffffffffffffff1663bfe389ed8b84038b6040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182815260200192505050600060405180830381600087803b15801561049f57600080fd5b505af11580156104b3573d6000803e3d6000fd5b505050508585858b868e6104c5610611565b808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018381526020018281526020019650505050505050604051809103906000f080158015610592573d6000803e3d6000fd5b5090507f4481db8423b8eae544e932e436f55884c6487e4d3ca295a3ca992ecf84ecc695818a604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019250505060405180910390a1809750505050505050509392505050565b6040516110778061062283390190560060806040526000600660006101000a81548160ff0219169083151502179055506000600660016101000a81548160ff0219169083151502179055506000600755600160085534801561005057600080fd5b5060405160c0806110778339810180604052810190808051906020019092919080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291905050508573ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16148061010a57508373ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b15156101a4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260348152602001807f4f6e6c7920697373756572206f72206163636570746f722063616e206973737581526020017f65206869732f686572206f776e206c657474657200000000000000000000000081525060400191505060405180910390fd5b856000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600381905550816004819055508060058190555060098590806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a839080600181540180825580915050906001820390600052602060002001600090919290919091505550505050505050610d55806103226000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680631c9242801461007d57806358007b86146100b65780635a9b0b89146100ef5780635dfd3c98146101d7578063a9059cbb1461028b578063bfe389ed146102d8575b600080fd5b34801561008957600080fd5b506100b46004803603810190808035151590602001909291908035906020019092919050505061030f565b005b3480156100c257600080fd5b506100ed6004803603810190808035151590602001909291908035906020019092919050505061044e565b005b3480156100fb57600080fd5b5061010461060a565b604051808973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200186815260200185815260200184815260200183151515158152602001821515151581526020019850505050505050505060405180910390f35b3480156101e357600080fd5b506101ec6106c1565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b83811015610233578082015181840152602081019050610218565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561027557808201518184015260208101905061025a565b5050505090500194505050505060405180910390f35b34801561029757600080fd5b506102d6600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506108ec565b005b3480156102e457600080fd5b5061030d6004803603810190808035906020019092919080359060200190929190505050610b93565b005b60008111151561031e57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156103e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260158152602001807f4f6e6c79206163636570746f722063616e20706179000000000000000000000081525060200191505060405180910390fd5b81600660016101000a81548160ff0219169083151502179055507f1ecf1b9cee605f7dc61806c2f40429418dbc06abbb7cfe58d1b5480f084af9a060085483836040518084815260200183151515158152602001828152602001935050505060405180910390a15050565b60008111151561045d57600080fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061050557506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b151561059f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260238152602001807f4f6e6c79206163636570746f72206f72206973737565722063616e206170707281526020017f6f7665000000000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b81600660006101000a81548160ff0219169083151502179055507f1ecf1b9cee605f7dc61806c2f40429418dbc06abbb7cfe58d1b5480f084af9a060075483836040518084815260200183151515158152602001828152602001935050505060405180910390a15050565b6000806000806000806000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600354600454600554600660009054906101000a900460ff16600660019054906101000a900460ff16975097509750975097509750975097509091929394959697565b6060806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061076c5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515610806576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f4f6e6c7920697373756572206f72206163636570746f722063616e206368656381526020017f6b207472616e73666572206c6f6773000000000000000000000000000000000081525060400191505060405180910390fd5b6009600a8180548060200260200160405190810160405280929190818152602001828054801561088b57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610841575b50505050509150808054806020026020016040519081016040528092919081815260200182805480156108dd57602002820191906000526020600020905b8154815260200190600101908083116108c9575b50505050509050915091509091565b8173ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561094957600080fd5b60008111151561095857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a1d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f4f6e6c7920686f6c6465722063616e207472616e73666572000000000000000081525060200191505060405180910390fd5b81600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060093390806001815401808255809150509060018203906000526020600020016000909192909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050600a8190806001815401808255809150509060018203906000526020600020016000909192909190915055507f0b208f30897dd37bbda3db1da3cf917bce605c53cdab4f7be8eaac32c8082f10338383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a15050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff161480610c3b5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163273ffffffffffffffffffffffffffffffffffffffff16145b1515610cd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f4f6e6c7920697373756572206f72206163636570746f722063616e207265736581526020017f742063726564697420616d6f756e74000000000000000000000000000000000081525060400191505060405180910390fd5b816005819055507fbe251dfecb445e45c89b4940942865904cf0efeea8fb1e4aa0f340e8370c03cd600554838360405180848152602001838152602001828152602001935050505060405180910390a150505600a165627a7a723058206cb1decaec00d899522b4d4c6d65da8f648f68771ad00179dfc9ae57dbdfebc90029a165627a7a72305820e2ed5b580cfe9dd6da8968ebaba67ab79a372c34404d4f02d52094a596e4f8900029"

// DeployCreditLetterFactory deploys a new contract, binding an instance of CreditLetterFactory to it.
func DeployCreditLetterFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CreditLetterFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditLetterFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CreditLetterFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CreditLetterFactory{CreditLetterFactoryCaller: CreditLetterFactoryCaller{contract: contract}, CreditLetterFactoryTransactor: CreditLetterFactoryTransactor{contract: contract}, CreditLetterFactoryFilterer: CreditLetterFactoryFilterer{contract: contract}}, nil
}

func AsyncDeployCreditLetterFactory(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditLetterFactoryABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(CreditLetterFactoryBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// CreditLetterFactory is an auto generated Go binding around a Solidity contract.
type CreditLetterFactory struct {
	CreditLetterFactoryCaller     // Read-only binding to the contract
	CreditLetterFactoryTransactor // Write-only binding to the contract
	CreditLetterFactoryFilterer   // Log filterer for contract events
}

// CreditLetterFactoryCaller is an auto generated read-only Go binding around a Solidity contract.
type CreditLetterFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditLetterFactoryTransactor is an auto generated write-only Go binding around a Solidity contract.
type CreditLetterFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditLetterFactoryFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type CreditLetterFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CreditLetterFactorySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type CreditLetterFactorySession struct {
	Contract     *CreditLetterFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CreditLetterFactoryCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type CreditLetterFactoryCallerSession struct {
	Contract *CreditLetterFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// CreditLetterFactoryTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type CreditLetterFactoryTransactorSession struct {
	Contract     *CreditLetterFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// CreditLetterFactoryRaw is an auto generated low-level Go binding around a Solidity contract.
type CreditLetterFactoryRaw struct {
	Contract *CreditLetterFactory // Generic contract binding to access the raw methods on
}

// CreditLetterFactoryCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type CreditLetterFactoryCallerRaw struct {
	Contract *CreditLetterFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CreditLetterFactoryTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type CreditLetterFactoryTransactorRaw struct {
	Contract *CreditLetterFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCreditLetterFactory creates a new instance of CreditLetterFactory, bound to a specific deployed contract.
func NewCreditLetterFactory(address common.Address, backend bind.ContractBackend) (*CreditLetterFactory, error) {
	contract, err := bindCreditLetterFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CreditLetterFactory{CreditLetterFactoryCaller: CreditLetterFactoryCaller{contract: contract}, CreditLetterFactoryTransactor: CreditLetterFactoryTransactor{contract: contract}, CreditLetterFactoryFilterer: CreditLetterFactoryFilterer{contract: contract}}, nil
}

// NewCreditLetterFactoryCaller creates a new read-only instance of CreditLetterFactory, bound to a specific deployed contract.
func NewCreditLetterFactoryCaller(address common.Address, caller bind.ContractCaller) (*CreditLetterFactoryCaller, error) {
	contract, err := bindCreditLetterFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CreditLetterFactoryCaller{contract: contract}, nil
}

// NewCreditLetterFactoryTransactor creates a new write-only instance of CreditLetterFactory, bound to a specific deployed contract.
func NewCreditLetterFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CreditLetterFactoryTransactor, error) {
	contract, err := bindCreditLetterFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CreditLetterFactoryTransactor{contract: contract}, nil
}

// NewCreditLetterFactoryFilterer creates a new log filterer instance of CreditLetterFactory, bound to a specific deployed contract.
func NewCreditLetterFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CreditLetterFactoryFilterer, error) {
	contract, err := bindCreditLetterFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CreditLetterFactoryFilterer{contract: contract}, nil
}

// bindCreditLetterFactory binds a generic wrapper to an already deployed contract.
func bindCreditLetterFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CreditLetterFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditLetterFactory *CreditLetterFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreditLetterFactory.Contract.CreditLetterFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditLetterFactory *CreditLetterFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.CreditLetterFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditLetterFactory *CreditLetterFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.CreditLetterFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CreditLetterFactory *CreditLetterFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CreditLetterFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CreditLetterFactory *CreditLetterFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CreditLetterFactory *CreditLetterFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.contract.Transact(opts, method, params...)
}

// Create is a paid mutator transaction binding the contract method 0x7f368f65.
//
// Solidity: function create(address issuer, address holder, address acceptor, uint256 issuanceTime, uint256 interestRate, uint256 credit) returns(address)
func (_CreditLetterFactory *CreditLetterFactoryTransactor) Create(opts *bind.TransactOpts, issuer common.Address, holder common.Address, acceptor common.Address, issuanceTime *big.Int, interestRate *big.Int, credit *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.contract.Transact(opts, "create", issuer, holder, acceptor, issuanceTime, interestRate, credit)
}

func (_CreditLetterFactory *CreditLetterFactoryTransactor) AsyncCreate(handler func(*types.Receipt, error), opts *bind.TransactOpts, issuer common.Address, holder common.Address, acceptor common.Address, issuanceTime *big.Int, interestRate *big.Int, credit *big.Int) (*types.Transaction, error) {
	return _CreditLetterFactory.contract.AsyncTransact(opts, handler, "create", issuer, holder, acceptor, issuanceTime, interestRate, credit)
}

// Create is a paid mutator transaction binding the contract method 0x7f368f65.
//
// Solidity: function create(address issuer, address holder, address acceptor, uint256 issuanceTime, uint256 interestRate, uint256 credit) returns(address)
func (_CreditLetterFactory *CreditLetterFactorySession) Create(issuer common.Address, holder common.Address, acceptor common.Address, issuanceTime *big.Int, interestRate *big.Int, credit *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.Create(&_CreditLetterFactory.TransactOpts, issuer, holder, acceptor, issuanceTime, interestRate, credit)
}

func (_CreditLetterFactory *CreditLetterFactorySession) AsyncCreate(handler func(*types.Receipt, error), issuer common.Address, holder common.Address, acceptor common.Address, issuanceTime *big.Int, interestRate *big.Int, credit *big.Int) (*types.Transaction, error) {
	return _CreditLetterFactory.Contract.AsyncCreate(handler, &_CreditLetterFactory.TransactOpts, issuer, holder, acceptor, issuanceTime, interestRate, credit)
}

// Create is a paid mutator transaction binding the contract method 0x7f368f65.
//
// Solidity: function create(address issuer, address holder, address acceptor, uint256 issuanceTime, uint256 interestRate, uint256 credit) returns(address)
func (_CreditLetterFactory *CreditLetterFactoryTransactorSession) Create(issuer common.Address, holder common.Address, acceptor common.Address, issuanceTime *big.Int, interestRate *big.Int, credit *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.Create(&_CreditLetterFactory.TransactOpts, issuer, holder, acceptor, issuanceTime, interestRate, credit)
}

func (_CreditLetterFactory *CreditLetterFactoryTransactorSession) AsyncCreate(handler func(*types.Receipt, error), issuer common.Address, holder common.Address, acceptor common.Address, issuanceTime *big.Int, interestRate *big.Int, credit *big.Int) (*types.Transaction, error) {
	return _CreditLetterFactory.Contract.AsyncCreate(handler, &_CreditLetterFactory.TransactOpts, issuer, holder, acceptor, issuanceTime, interestRate, credit)
}

// Split is a paid mutator transaction binding the contract method 0x8afbc1ed.
//
// Solidity: function split(address originalAddress, uint256 amount, uint256 timestamp) returns(address)
func (_CreditLetterFactory *CreditLetterFactoryTransactor) Split(opts *bind.TransactOpts, originalAddress common.Address, amount *big.Int, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.contract.Transact(opts, "split", originalAddress, amount, timestamp)
}

func (_CreditLetterFactory *CreditLetterFactoryTransactor) AsyncSplit(handler func(*types.Receipt, error), opts *bind.TransactOpts, originalAddress common.Address, amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetterFactory.contract.AsyncTransact(opts, handler, "split", originalAddress, amount, timestamp)
}

// Split is a paid mutator transaction binding the contract method 0x8afbc1ed.
//
// Solidity: function split(address originalAddress, uint256 amount, uint256 timestamp) returns(address)
func (_CreditLetterFactory *CreditLetterFactorySession) Split(originalAddress common.Address, amount *big.Int, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.Split(&_CreditLetterFactory.TransactOpts, originalAddress, amount, timestamp)
}

func (_CreditLetterFactory *CreditLetterFactorySession) AsyncSplit(handler func(*types.Receipt, error), originalAddress common.Address, amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetterFactory.Contract.AsyncSplit(handler, &_CreditLetterFactory.TransactOpts, originalAddress, amount, timestamp)
}

// Split is a paid mutator transaction binding the contract method 0x8afbc1ed.
//
// Solidity: function split(address originalAddress, uint256 amount, uint256 timestamp) returns(address)
func (_CreditLetterFactory *CreditLetterFactoryTransactorSession) Split(originalAddress common.Address, amount *big.Int, timestamp *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _CreditLetterFactory.Contract.Split(&_CreditLetterFactory.TransactOpts, originalAddress, amount, timestamp)
}

func (_CreditLetterFactory *CreditLetterFactoryTransactorSession) AsyncSplit(handler func(*types.Receipt, error), originalAddress common.Address, amount *big.Int, timestamp *big.Int) (*types.Transaction, error) {
	return _CreditLetterFactory.Contract.AsyncSplit(handler, &_CreditLetterFactory.TransactOpts, originalAddress, amount, timestamp)
}

// CreditLetterFactoryCreateLogIterator is returned from FilterCreateLog and is used to iterate over the raw logs and unpacked data for CreateLog events raised by the CreditLetterFactory contract.
type CreditLetterFactoryCreateLogIterator struct {
	Event *CreditLetterFactoryCreateLog // Event containing the contract specifics and raw log

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
func (it *CreditLetterFactoryCreateLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CreditLetterFactoryCreateLog)
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
		it.Event = new(CreditLetterFactoryCreateLog)
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
func (it *CreditLetterFactoryCreateLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CreditLetterFactoryCreateLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CreditLetterFactoryCreateLog represents a CreateLog event raised by the CreditLetterFactory contract.
type CreditLetterFactoryCreateLog struct {
	Addr      common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateLog is a free log retrieval operation binding the contract event 0x4481db8423b8eae544e932e436f55884c6487e4d3ca295a3ca992ecf84ecc695.
//
// Solidity: event CreateLog(address addr, uint256 timestamp)
func (_CreditLetterFactory *CreditLetterFactoryFilterer) FilterCreateLog(opts *bind.FilterOpts) (*CreditLetterFactoryCreateLogIterator, error) {

	logs, sub, err := _CreditLetterFactory.contract.FilterLogs(opts, "CreateLog")
	if err != nil {
		return nil, err
	}
	return &CreditLetterFactoryCreateLogIterator{contract: _CreditLetterFactory.contract, event: "CreateLog", logs: logs, sub: sub}, nil
}

// WatchCreateLog is a free log subscription operation binding the contract event 0x4481db8423b8eae544e932e436f55884c6487e4d3ca295a3ca992ecf84ecc695.
//
// Solidity: event CreateLog(address addr, uint256 timestamp)
func (_CreditLetterFactory *CreditLetterFactoryFilterer) WatchCreateLog(opts *bind.WatchOpts, sink chan<- *CreditLetterFactoryCreateLog) (event.Subscription, error) {

	logs, sub, err := _CreditLetterFactory.contract.WatchLogs(opts, "CreateLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CreditLetterFactoryCreateLog)
				if err := _CreditLetterFactory.contract.UnpackLog(event, "CreateLog", log); err != nil {
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

// ParseCreateLog is a log parse operation binding the contract event 0x4481db8423b8eae544e932e436f55884c6487e4d3ca295a3ca992ecf84ecc695.
//
// Solidity: event CreateLog(address addr, uint256 timestamp)
func (_CreditLetterFactory *CreditLetterFactoryFilterer) ParseCreateLog(log types.Log) (*CreditLetterFactoryCreateLog, error) {
	event := new(CreditLetterFactoryCreateLog)
	if err := _CreditLetterFactory.contract.UnpackLog(event, "CreateLog", log); err != nil {
		return nil, err
	}
	return event, nil
}
