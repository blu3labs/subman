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

// LibSubSubPayment is an auto generated low-level Go binding around an user-defined struct.
type LibSubSubPayment struct {
	SubPlanId    *big.Int
	Subscriber   common.Address
	StartTime    *big.Int
	EndTime      *big.Int
	Duration     *big.Int
	PaymentToken common.Address
	Price        *big.Int
}

// LibSubSubPlan is an auto generated low-level Go binding around an user-defined struct.
type LibSubSubPlan struct {
	Title           string
	Description     string
	Owner           common.Address
	PaymentReceiver common.Address
	PaymentToken    common.Address
	Price           *big.Int
	Duration        *big.Int
	Deadline        *big.Int
	SubPlanId       *big.Int
	ServiceFee      *big.Int
	Active          bool
}

// LibSubSubscription is an auto generated low-level Go binding around an user-defined struct.
type LibSubSubscription struct {
	SubPlan  LibSubSubPlan
	Deadline *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriber\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLibSub.SubPayment\",\"name\":\"subPayment\",\"type\":\"tuple\"}],\"name\":\"SubPaymentCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriber\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structLibSub.SubPayment\",\"name\":\"subPayment\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newDeadline\",\"type\":\"uint256\"}],\"name\":\"SubPaymentProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structLibSub.SubPlan\",\"name\":\"subPlan\",\"type\":\"tuple\"}],\"name\":\"SubPlanActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structLibSub.SubPlan\",\"name\":\"subPlan\",\"type\":\"tuple\"}],\"name\":\"SubPlanCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structLibSub.SubPlan\",\"name\":\"subPlan\",\"type\":\"tuple\"}],\"name\":\"SubPlanDeactivated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subPlanId\",\"type\":\"uint256\"}],\"name\":\"activateSubPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriber\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structLibSub.SubPayment\",\"name\":\"_subPayment\",\"type\":\"tuple\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_serviceFee\",\"type\":\"uint256\"}],\"name\":\"createSubPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subPlanId\",\"type\":\"uint256\"}],\"name\":\"deactivateSubPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subscriber\",\"type\":\"address\"}],\"name\":\"getActiveSubPlans\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subscriber\",\"type\":\"address\"}],\"name\":\"getActiveSubscriptions\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structLibSub.SubPlan\",\"name\":\"subPlan\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structLibSub.Subscription[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_subPlanId\",\"type\":\"uint256\"}],\"name\":\"getSubPlan\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"internalType\":\"structLibSub.SubPlan\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"getSubPlansByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subscriber\",\"type\":\"address\"}],\"name\":\"getSubedPlans\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_subscriber\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_subPlanId\",\"type\":\"uint256\"}],\"name\":\"getUserSubDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriber\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structLibSub.SubPayment\",\"name\":\"_subPayment\",\"type\":\"tuple\"}],\"name\":\"isCanceled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"subPlanId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"subscriber\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structLibSub.SubPayment\",\"name\":\"_subPayment\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_serviceFeeReceiver\",\"type\":\"address\"}],\"name\":\"processPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"subPlanCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}],",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

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

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contract *ContractCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contract *ContractSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Contract.Contract.Eip712Domain(&_Contract.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Contract *ContractCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Contract.Contract.Eip712Domain(&_Contract.CallOpts)
}

// GetActiveSubPlans is a free data retrieval call binding the contract method 0x17609e31.
//
// Solidity: function getActiveSubPlans(address _subscriber) view returns(uint256[])
func (_Contract *ContractCaller) GetActiveSubPlans(opts *bind.CallOpts, _subscriber common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getActiveSubPlans", _subscriber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveSubPlans is a free data retrieval call binding the contract method 0x17609e31.
//
// Solidity: function getActiveSubPlans(address _subscriber) view returns(uint256[])
func (_Contract *ContractSession) GetActiveSubPlans(_subscriber common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetActiveSubPlans(&_Contract.CallOpts, _subscriber)
}

// GetActiveSubPlans is a free data retrieval call binding the contract method 0x17609e31.
//
// Solidity: function getActiveSubPlans(address _subscriber) view returns(uint256[])
func (_Contract *ContractCallerSession) GetActiveSubPlans(_subscriber common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetActiveSubPlans(&_Contract.CallOpts, _subscriber)
}

// GetActiveSubscriptions is a free data retrieval call binding the contract method 0xc70eba2a.
//
// Solidity: function getActiveSubscriptions(address _subscriber) view returns(((string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool),uint256)[])
func (_Contract *ContractCaller) GetActiveSubscriptions(opts *bind.CallOpts, _subscriber common.Address) ([]LibSubSubscription, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getActiveSubscriptions", _subscriber)

	if err != nil {
		return *new([]LibSubSubscription), err
	}

	out0 := *abi.ConvertType(out[0], new([]LibSubSubscription)).(*[]LibSubSubscription)

	return out0, err

}

// GetActiveSubscriptions is a free data retrieval call binding the contract method 0xc70eba2a.
//
// Solidity: function getActiveSubscriptions(address _subscriber) view returns(((string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool),uint256)[])
func (_Contract *ContractSession) GetActiveSubscriptions(_subscriber common.Address) ([]LibSubSubscription, error) {
	return _Contract.Contract.GetActiveSubscriptions(&_Contract.CallOpts, _subscriber)
}

// GetActiveSubscriptions is a free data retrieval call binding the contract method 0xc70eba2a.
//
// Solidity: function getActiveSubscriptions(address _subscriber) view returns(((string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool),uint256)[])
func (_Contract *ContractCallerSession) GetActiveSubscriptions(_subscriber common.Address) ([]LibSubSubscription, error) {
	return _Contract.Contract.GetActiveSubscriptions(&_Contract.CallOpts, _subscriber)
}

// GetSubPlan is a free data retrieval call binding the contract method 0xc775b19e.
//
// Solidity: function getSubPlan(uint256 _subPlanId) view returns((string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool))
func (_Contract *ContractCaller) GetSubPlan(opts *bind.CallOpts, _subPlanId *big.Int) (LibSubSubPlan, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSubPlan", _subPlanId)

	if err != nil {
		return *new(LibSubSubPlan), err
	}

	out0 := *abi.ConvertType(out[0], new(LibSubSubPlan)).(*LibSubSubPlan)

	return out0, err

}

// GetSubPlan is a free data retrieval call binding the contract method 0xc775b19e.
//
// Solidity: function getSubPlan(uint256 _subPlanId) view returns((string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool))
func (_Contract *ContractSession) GetSubPlan(_subPlanId *big.Int) (LibSubSubPlan, error) {
	return _Contract.Contract.GetSubPlan(&_Contract.CallOpts, _subPlanId)
}

// GetSubPlan is a free data retrieval call binding the contract method 0xc775b19e.
//
// Solidity: function getSubPlan(uint256 _subPlanId) view returns((string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool))
func (_Contract *ContractCallerSession) GetSubPlan(_subPlanId *big.Int) (LibSubSubPlan, error) {
	return _Contract.Contract.GetSubPlan(&_Contract.CallOpts, _subPlanId)
}

// GetSubPlansByOwner is a free data retrieval call binding the contract method 0x8562ecc8.
//
// Solidity: function getSubPlansByOwner(address _owner) view returns(uint256[])
func (_Contract *ContractCaller) GetSubPlansByOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSubPlansByOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSubPlansByOwner is a free data retrieval call binding the contract method 0x8562ecc8.
//
// Solidity: function getSubPlansByOwner(address _owner) view returns(uint256[])
func (_Contract *ContractSession) GetSubPlansByOwner(_owner common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetSubPlansByOwner(&_Contract.CallOpts, _owner)
}

// GetSubPlansByOwner is a free data retrieval call binding the contract method 0x8562ecc8.
//
// Solidity: function getSubPlansByOwner(address _owner) view returns(uint256[])
func (_Contract *ContractCallerSession) GetSubPlansByOwner(_owner common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetSubPlansByOwner(&_Contract.CallOpts, _owner)
}

// GetSubedPlans is a free data retrieval call binding the contract method 0x96e7f554.
//
// Solidity: function getSubedPlans(address _subscriber) view returns(uint256[])
func (_Contract *ContractCaller) GetSubedPlans(opts *bind.CallOpts, _subscriber common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSubedPlans", _subscriber)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetSubedPlans is a free data retrieval call binding the contract method 0x96e7f554.
//
// Solidity: function getSubedPlans(address _subscriber) view returns(uint256[])
func (_Contract *ContractSession) GetSubedPlans(_subscriber common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetSubedPlans(&_Contract.CallOpts, _subscriber)
}

// GetSubedPlans is a free data retrieval call binding the contract method 0x96e7f554.
//
// Solidity: function getSubedPlans(address _subscriber) view returns(uint256[])
func (_Contract *ContractCallerSession) GetSubedPlans(_subscriber common.Address) ([]*big.Int, error) {
	return _Contract.Contract.GetSubedPlans(&_Contract.CallOpts, _subscriber)
}

// GetUserSubDeadline is a free data retrieval call binding the contract method 0xc31d916f.
//
// Solidity: function getUserSubDeadline(address _subscriber, uint256 _subPlanId) view returns(uint256)
func (_Contract *ContractCaller) GetUserSubDeadline(opts *bind.CallOpts, _subscriber common.Address, _subPlanId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUserSubDeadline", _subscriber, _subPlanId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserSubDeadline is a free data retrieval call binding the contract method 0xc31d916f.
//
// Solidity: function getUserSubDeadline(address _subscriber, uint256 _subPlanId) view returns(uint256)
func (_Contract *ContractSession) GetUserSubDeadline(_subscriber common.Address, _subPlanId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUserSubDeadline(&_Contract.CallOpts, _subscriber, _subPlanId)
}

// GetUserSubDeadline is a free data retrieval call binding the contract method 0xc31d916f.
//
// Solidity: function getUserSubDeadline(address _subscriber, uint256 _subPlanId) view returns(uint256)
func (_Contract *ContractCallerSession) GetUserSubDeadline(_subscriber common.Address, _subPlanId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetUserSubDeadline(&_Contract.CallOpts, _subscriber, _subPlanId)
}

// IsCanceled is a free data retrieval call binding the contract method 0xf03413b3.
//
// Solidity: function isCanceled((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment) view returns(bool)
func (_Contract *ContractCaller) IsCanceled(opts *bind.CallOpts, _subPayment LibSubSubPayment) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "isCanceled", _subPayment)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCanceled is a free data retrieval call binding the contract method 0xf03413b3.
//
// Solidity: function isCanceled((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment) view returns(bool)
func (_Contract *ContractSession) IsCanceled(_subPayment LibSubSubPayment) (bool, error) {
	return _Contract.Contract.IsCanceled(&_Contract.CallOpts, _subPayment)
}

// IsCanceled is a free data retrieval call binding the contract method 0xf03413b3.
//
// Solidity: function isCanceled((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment) view returns(bool)
func (_Contract *ContractCallerSession) IsCanceled(_subPayment LibSubSubPayment) (bool, error) {
	return _Contract.Contract.IsCanceled(&_Contract.CallOpts, _subPayment)
}

// SubPlanCount is a free data retrieval call binding the contract method 0x894d2d39.
//
// Solidity: function subPlanCount() view returns(uint256)
func (_Contract *ContractCaller) SubPlanCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "subPlanCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubPlanCount is a free data retrieval call binding the contract method 0x894d2d39.
//
// Solidity: function subPlanCount() view returns(uint256)
func (_Contract *ContractSession) SubPlanCount() (*big.Int, error) {
	return _Contract.Contract.SubPlanCount(&_Contract.CallOpts)
}

// SubPlanCount is a free data retrieval call binding the contract method 0x894d2d39.
//
// Solidity: function subPlanCount() view returns(uint256)
func (_Contract *ContractCallerSession) SubPlanCount() (*big.Int, error) {
	return _Contract.Contract.SubPlanCount(&_Contract.CallOpts)
}

// ActivateSubPlan is a paid mutator transaction binding the contract method 0x9052dcdc.
//
// Solidity: function activateSubPlan(uint256 _subPlanId) returns()
func (_Contract *ContractTransactor) ActivateSubPlan(opts *bind.TransactOpts, _subPlanId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "activateSubPlan", _subPlanId)
}

// ActivateSubPlan is a paid mutator transaction binding the contract method 0x9052dcdc.
//
// Solidity: function activateSubPlan(uint256 _subPlanId) returns()
func (_Contract *ContractSession) ActivateSubPlan(_subPlanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ActivateSubPlan(&_Contract.TransactOpts, _subPlanId)
}

// ActivateSubPlan is a paid mutator transaction binding the contract method 0x9052dcdc.
//
// Solidity: function activateSubPlan(uint256 _subPlanId) returns()
func (_Contract *ContractTransactorSession) ActivateSubPlan(_subPlanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ActivateSubPlan(&_Contract.TransactOpts, _subPlanId)
}

// Cancel is a paid mutator transaction binding the contract method 0x3eb63527.
//
// Solidity: function cancel((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment) returns()
func (_Contract *ContractTransactor) Cancel(opts *bind.TransactOpts, _subPayment LibSubSubPayment) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "cancel", _subPayment)
}

// Cancel is a paid mutator transaction binding the contract method 0x3eb63527.
//
// Solidity: function cancel((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment) returns()
func (_Contract *ContractSession) Cancel(_subPayment LibSubSubPayment) (*types.Transaction, error) {
	return _Contract.Contract.Cancel(&_Contract.TransactOpts, _subPayment)
}

// Cancel is a paid mutator transaction binding the contract method 0x3eb63527.
//
// Solidity: function cancel((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment) returns()
func (_Contract *ContractTransactorSession) Cancel(_subPayment LibSubSubPayment) (*types.Transaction, error) {
	return _Contract.Contract.Cancel(&_Contract.TransactOpts, _subPayment)
}

// CreateSubPlan is a paid mutator transaction binding the contract method 0x09d64b3a.
//
// Solidity: function createSubPlan(string _title, string _description, address _paymentReceiver, address _paymentToken, uint256 _price, uint256 _duration, uint256 _serviceFee) returns()
func (_Contract *ContractTransactor) CreateSubPlan(opts *bind.TransactOpts, _title string, _description string, _paymentReceiver common.Address, _paymentToken common.Address, _price *big.Int, _duration *big.Int, _serviceFee *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createSubPlan", _title, _description, _paymentReceiver, _paymentToken, _price, _duration, _serviceFee)
}

// CreateSubPlan is a paid mutator transaction binding the contract method 0x09d64b3a.
//
// Solidity: function createSubPlan(string _title, string _description, address _paymentReceiver, address _paymentToken, uint256 _price, uint256 _duration, uint256 _serviceFee) returns()
func (_Contract *ContractSession) CreateSubPlan(_title string, _description string, _paymentReceiver common.Address, _paymentToken common.Address, _price *big.Int, _duration *big.Int, _serviceFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateSubPlan(&_Contract.TransactOpts, _title, _description, _paymentReceiver, _paymentToken, _price, _duration, _serviceFee)
}

// CreateSubPlan is a paid mutator transaction binding the contract method 0x09d64b3a.
//
// Solidity: function createSubPlan(string _title, string _description, address _paymentReceiver, address _paymentToken, uint256 _price, uint256 _duration, uint256 _serviceFee) returns()
func (_Contract *ContractTransactorSession) CreateSubPlan(_title string, _description string, _paymentReceiver common.Address, _paymentToken common.Address, _price *big.Int, _duration *big.Int, _serviceFee *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.CreateSubPlan(&_Contract.TransactOpts, _title, _description, _paymentReceiver, _paymentToken, _price, _duration, _serviceFee)
}

// DeactivateSubPlan is a paid mutator transaction binding the contract method 0x84b60745.
//
// Solidity: function deactivateSubPlan(uint256 _subPlanId) returns()
func (_Contract *ContractTransactor) DeactivateSubPlan(opts *bind.TransactOpts, _subPlanId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deactivateSubPlan", _subPlanId)
}

// DeactivateSubPlan is a paid mutator transaction binding the contract method 0x84b60745.
//
// Solidity: function deactivateSubPlan(uint256 _subPlanId) returns()
func (_Contract *ContractSession) DeactivateSubPlan(_subPlanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateSubPlan(&_Contract.TransactOpts, _subPlanId)
}

// DeactivateSubPlan is a paid mutator transaction binding the contract method 0x84b60745.
//
// Solidity: function deactivateSubPlan(uint256 _subPlanId) returns()
func (_Contract *ContractTransactorSession) DeactivateSubPlan(_subPlanId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DeactivateSubPlan(&_Contract.TransactOpts, _subPlanId)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x16db84d2.
//
// Solidity: function processPayment((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment, bytes _signature, address _serviceFeeReceiver) returns()
func (_Contract *ContractTransactor) ProcessPayment(opts *bind.TransactOpts, _subPayment LibSubSubPayment, _signature []byte, _serviceFeeReceiver common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "processPayment", _subPayment, _signature, _serviceFeeReceiver)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x16db84d2.
//
// Solidity: function processPayment((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment, bytes _signature, address _serviceFeeReceiver) returns()
func (_Contract *ContractSession) ProcessPayment(_subPayment LibSubSubPayment, _signature []byte, _serviceFeeReceiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ProcessPayment(&_Contract.TransactOpts, _subPayment, _signature, _serviceFeeReceiver)
}

// ProcessPayment is a paid mutator transaction binding the contract method 0x16db84d2.
//
// Solidity: function processPayment((uint256,address,uint256,uint256,uint256,address,uint256) _subPayment, bytes _signature, address _serviceFeeReceiver) returns()
func (_Contract *ContractTransactorSession) ProcessPayment(_subPayment LibSubSubPayment, _signature []byte, _serviceFeeReceiver common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ProcessPayment(&_Contract.TransactOpts, _subPayment, _signature, _serviceFeeReceiver)
}

// ContractEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Contract contract.
type ContractEIP712DomainChangedIterator struct {
	Event *ContractEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ContractEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEIP712DomainChanged)
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
		it.Event = new(ContractEIP712DomainChanged)
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
func (it *ContractEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEIP712DomainChanged represents a EIP712DomainChanged event raised by the Contract contract.
type ContractEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contract *ContractFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ContractEIP712DomainChangedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ContractEIP712DomainChangedIterator{contract: _Contract.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contract *ContractFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ContractEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEIP712DomainChanged)
				if err := _Contract.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Contract *ContractFilterer) ParseEIP712DomainChanged(log types.Log) (*ContractEIP712DomainChanged, error) {
	event := new(ContractEIP712DomainChanged)
	if err := _Contract.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSubPaymentCanceledIterator is returned from FilterSubPaymentCanceled and is used to iterate over the raw logs and unpacked data for SubPaymentCanceled events raised by the Contract contract.
type ContractSubPaymentCanceledIterator struct {
	Event *ContractSubPaymentCanceled // Event containing the contract specifics and raw log

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
func (it *ContractSubPaymentCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSubPaymentCanceled)
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
		it.Event = new(ContractSubPaymentCanceled)
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
func (it *ContractSubPaymentCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSubPaymentCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSubPaymentCanceled represents a SubPaymentCanceled event raised by the Contract contract.
type ContractSubPaymentCanceled struct {
	SubPayment LibSubSubPayment
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubPaymentCanceled is a free log retrieval operation binding the contract event 0x83da4f2ee417cc6162cbed15aad334689919b6305d9913e464e2d418627afabb.
//
// Solidity: event SubPaymentCanceled((uint256,address,uint256,uint256,uint256,address,uint256) subPayment)
func (_Contract *ContractFilterer) FilterSubPaymentCanceled(opts *bind.FilterOpts) (*ContractSubPaymentCanceledIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SubPaymentCanceled")
	if err != nil {
		return nil, err
	}
	return &ContractSubPaymentCanceledIterator{contract: _Contract.contract, event: "SubPaymentCanceled", logs: logs, sub: sub}, nil
}

// WatchSubPaymentCanceled is a free log subscription operation binding the contract event 0x83da4f2ee417cc6162cbed15aad334689919b6305d9913e464e2d418627afabb.
//
// Solidity: event SubPaymentCanceled((uint256,address,uint256,uint256,uint256,address,uint256) subPayment)
func (_Contract *ContractFilterer) WatchSubPaymentCanceled(opts *bind.WatchOpts, sink chan<- *ContractSubPaymentCanceled) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SubPaymentCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSubPaymentCanceled)
				if err := _Contract.contract.UnpackLog(event, "SubPaymentCanceled", log); err != nil {
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

// ParseSubPaymentCanceled is a log parse operation binding the contract event 0x83da4f2ee417cc6162cbed15aad334689919b6305d9913e464e2d418627afabb.
//
// Solidity: event SubPaymentCanceled((uint256,address,uint256,uint256,uint256,address,uint256) subPayment)
func (_Contract *ContractFilterer) ParseSubPaymentCanceled(log types.Log) (*ContractSubPaymentCanceled, error) {
	event := new(ContractSubPaymentCanceled)
	if err := _Contract.contract.UnpackLog(event, "SubPaymentCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSubPaymentProcessedIterator is returned from FilterSubPaymentProcessed and is used to iterate over the raw logs and unpacked data for SubPaymentProcessed events raised by the Contract contract.
type ContractSubPaymentProcessedIterator struct {
	Event *ContractSubPaymentProcessed // Event containing the contract specifics and raw log

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
func (it *ContractSubPaymentProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSubPaymentProcessed)
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
		it.Event = new(ContractSubPaymentProcessed)
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
func (it *ContractSubPaymentProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSubPaymentProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSubPaymentProcessed represents a SubPaymentProcessed event raised by the Contract contract.
type ContractSubPaymentProcessed struct {
	SubPayment  LibSubSubPayment
	NewDeadline *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubPaymentProcessed is a free log retrieval operation binding the contract event 0x6ce2aeff28f1b230f8d1e87f0a0e0eb9a37cf49ca39ff7105dce41fdeddf7a3d.
//
// Solidity: event SubPaymentProcessed((uint256,address,uint256,uint256,uint256,address,uint256) subPayment, uint256 newDeadline)
func (_Contract *ContractFilterer) FilterSubPaymentProcessed(opts *bind.FilterOpts) (*ContractSubPaymentProcessedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SubPaymentProcessed")
	if err != nil {
		return nil, err
	}
	return &ContractSubPaymentProcessedIterator{contract: _Contract.contract, event: "SubPaymentProcessed", logs: logs, sub: sub}, nil
}

// WatchSubPaymentProcessed is a free log subscription operation binding the contract event 0x6ce2aeff28f1b230f8d1e87f0a0e0eb9a37cf49ca39ff7105dce41fdeddf7a3d.
//
// Solidity: event SubPaymentProcessed((uint256,address,uint256,uint256,uint256,address,uint256) subPayment, uint256 newDeadline)
func (_Contract *ContractFilterer) WatchSubPaymentProcessed(opts *bind.WatchOpts, sink chan<- *ContractSubPaymentProcessed) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SubPaymentProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSubPaymentProcessed)
				if err := _Contract.contract.UnpackLog(event, "SubPaymentProcessed", log); err != nil {
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

// ParseSubPaymentProcessed is a log parse operation binding the contract event 0x6ce2aeff28f1b230f8d1e87f0a0e0eb9a37cf49ca39ff7105dce41fdeddf7a3d.
//
// Solidity: event SubPaymentProcessed((uint256,address,uint256,uint256,uint256,address,uint256) subPayment, uint256 newDeadline)
func (_Contract *ContractFilterer) ParseSubPaymentProcessed(log types.Log) (*ContractSubPaymentProcessed, error) {
	event := new(ContractSubPaymentProcessed)
	if err := _Contract.contract.UnpackLog(event, "SubPaymentProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSubPlanActivatedIterator is returned from FilterSubPlanActivated and is used to iterate over the raw logs and unpacked data for SubPlanActivated events raised by the Contract contract.
type ContractSubPlanActivatedIterator struct {
	Event *ContractSubPlanActivated // Event containing the contract specifics and raw log

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
func (it *ContractSubPlanActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSubPlanActivated)
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
		it.Event = new(ContractSubPlanActivated)
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
func (it *ContractSubPlanActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSubPlanActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSubPlanActivated represents a SubPlanActivated event raised by the Contract contract.
type ContractSubPlanActivated struct {
	SubPlanId *big.Int
	SubPlan   LibSubSubPlan
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubPlanActivated is a free log retrieval operation binding the contract event 0x3ce6aef75e0733c43ab0b4f75edbb60d4a987bfdbc5cf4cac7efc7525992bf31.
//
// Solidity: event SubPlanActivated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) FilterSubPlanActivated(opts *bind.FilterOpts) (*ContractSubPlanActivatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SubPlanActivated")
	if err != nil {
		return nil, err
	}
	return &ContractSubPlanActivatedIterator{contract: _Contract.contract, event: "SubPlanActivated", logs: logs, sub: sub}, nil
}

// WatchSubPlanActivated is a free log subscription operation binding the contract event 0x3ce6aef75e0733c43ab0b4f75edbb60d4a987bfdbc5cf4cac7efc7525992bf31.
//
// Solidity: event SubPlanActivated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) WatchSubPlanActivated(opts *bind.WatchOpts, sink chan<- *ContractSubPlanActivated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SubPlanActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSubPlanActivated)
				if err := _Contract.contract.UnpackLog(event, "SubPlanActivated", log); err != nil {
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

// ParseSubPlanActivated is a log parse operation binding the contract event 0x3ce6aef75e0733c43ab0b4f75edbb60d4a987bfdbc5cf4cac7efc7525992bf31.
//
// Solidity: event SubPlanActivated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) ParseSubPlanActivated(log types.Log) (*ContractSubPlanActivated, error) {
	event := new(ContractSubPlanActivated)
	if err := _Contract.contract.UnpackLog(event, "SubPlanActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSubPlanCreatedIterator is returned from FilterSubPlanCreated and is used to iterate over the raw logs and unpacked data for SubPlanCreated events raised by the Contract contract.
type ContractSubPlanCreatedIterator struct {
	Event *ContractSubPlanCreated // Event containing the contract specifics and raw log

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
func (it *ContractSubPlanCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSubPlanCreated)
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
		it.Event = new(ContractSubPlanCreated)
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
func (it *ContractSubPlanCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSubPlanCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSubPlanCreated represents a SubPlanCreated event raised by the Contract contract.
type ContractSubPlanCreated struct {
	SubPlanId *big.Int
	SubPlan   LibSubSubPlan
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubPlanCreated is a free log retrieval operation binding the contract event 0x2798d7140983ea6a604617e7deeb807f3254ec2f94ccd345bafeb8551599d67a.
//
// Solidity: event SubPlanCreated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) FilterSubPlanCreated(opts *bind.FilterOpts) (*ContractSubPlanCreatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SubPlanCreated")
	if err != nil {
		return nil, err
	}
	return &ContractSubPlanCreatedIterator{contract: _Contract.contract, event: "SubPlanCreated", logs: logs, sub: sub}, nil
}

// WatchSubPlanCreated is a free log subscription operation binding the contract event 0x2798d7140983ea6a604617e7deeb807f3254ec2f94ccd345bafeb8551599d67a.
//
// Solidity: event SubPlanCreated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) WatchSubPlanCreated(opts *bind.WatchOpts, sink chan<- *ContractSubPlanCreated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SubPlanCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSubPlanCreated)
				if err := _Contract.contract.UnpackLog(event, "SubPlanCreated", log); err != nil {
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

// ParseSubPlanCreated is a log parse operation binding the contract event 0x2798d7140983ea6a604617e7deeb807f3254ec2f94ccd345bafeb8551599d67a.
//
// Solidity: event SubPlanCreated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) ParseSubPlanCreated(log types.Log) (*ContractSubPlanCreated, error) {
	event := new(ContractSubPlanCreated)
	if err := _Contract.contract.UnpackLog(event, "SubPlanCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSubPlanDeactivatedIterator is returned from FilterSubPlanDeactivated and is used to iterate over the raw logs and unpacked data for SubPlanDeactivated events raised by the Contract contract.
type ContractSubPlanDeactivatedIterator struct {
	Event *ContractSubPlanDeactivated // Event containing the contract specifics and raw log

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
func (it *ContractSubPlanDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSubPlanDeactivated)
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
		it.Event = new(ContractSubPlanDeactivated)
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
func (it *ContractSubPlanDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSubPlanDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSubPlanDeactivated represents a SubPlanDeactivated event raised by the Contract contract.
type ContractSubPlanDeactivated struct {
	SubPlanId *big.Int
	SubPlan   LibSubSubPlan
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSubPlanDeactivated is a free log retrieval operation binding the contract event 0xc23878edc7a44457406383ee185dc7c2642a3d87506d0cc57aae5132d87684f9.
//
// Solidity: event SubPlanDeactivated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) FilterSubPlanDeactivated(opts *bind.FilterOpts) (*ContractSubPlanDeactivatedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SubPlanDeactivated")
	if err != nil {
		return nil, err
	}
	return &ContractSubPlanDeactivatedIterator{contract: _Contract.contract, event: "SubPlanDeactivated", logs: logs, sub: sub}, nil
}

// WatchSubPlanDeactivated is a free log subscription operation binding the contract event 0xc23878edc7a44457406383ee185dc7c2642a3d87506d0cc57aae5132d87684f9.
//
// Solidity: event SubPlanDeactivated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) WatchSubPlanDeactivated(opts *bind.WatchOpts, sink chan<- *ContractSubPlanDeactivated) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SubPlanDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSubPlanDeactivated)
				if err := _Contract.contract.UnpackLog(event, "SubPlanDeactivated", log); err != nil {
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

// ParseSubPlanDeactivated is a log parse operation binding the contract event 0xc23878edc7a44457406383ee185dc7c2642a3d87506d0cc57aae5132d87684f9.
//
// Solidity: event SubPlanDeactivated(uint256 subPlanId, (string,string,address,address,address,uint256,uint256,uint256,uint256,uint256,bool) subPlan)
func (_Contract *ContractFilterer) ParseSubPlanDeactivated(log types.Log) (*ContractSubPlanDeactivated, error) {
	event := new(ContractSubPlanDeactivated)
	if err := _Contract.contract.UnpackLog(event, "SubPlanDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
