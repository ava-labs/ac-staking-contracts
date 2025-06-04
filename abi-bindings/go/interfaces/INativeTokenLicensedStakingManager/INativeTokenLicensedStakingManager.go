// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package inativetokenlicensedstakingmanager

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ava-labs/subnet-evm/accounts/abi"
	"github.com/ava-labs/subnet-evm/accounts/abi/bind"
	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/interfaces"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = interfaces.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// INativeTokenLicensedStakingManagerMetaData contains all meta data concerning the INativeTokenLicensedStakingManager contract.
var INativeTokenLicensedStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"changeDelegatorRewardRecipient\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeValidatorRewardRecipient\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRemoval\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"erc721\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceInitiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDelegatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenDelegator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenValidator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDelegations\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRegistration\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"licenseToStakeConversionFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegator\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CompletedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CompletedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardClaimed\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedStakingValidatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]",
}

// INativeTokenLicensedStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use INativeTokenLicensedStakingManagerMetaData.ABI instead.
var INativeTokenLicensedStakingManagerABI = INativeTokenLicensedStakingManagerMetaData.ABI

// INativeTokenLicensedStakingManager is an auto generated Go binding around an Ethereum contract.
type INativeTokenLicensedStakingManager struct {
	INativeTokenLicensedStakingManagerCaller     // Read-only binding to the contract
	INativeTokenLicensedStakingManagerTransactor // Write-only binding to the contract
	INativeTokenLicensedStakingManagerFilterer   // Log filterer for contract events
}

// INativeTokenLicensedStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type INativeTokenLicensedStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeTokenLicensedStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type INativeTokenLicensedStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeTokenLicensedStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type INativeTokenLicensedStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// INativeTokenLicensedStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type INativeTokenLicensedStakingManagerSession struct {
	Contract     *INativeTokenLicensedStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// INativeTokenLicensedStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type INativeTokenLicensedStakingManagerCallerSession struct {
	Contract *INativeTokenLicensedStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// INativeTokenLicensedStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type INativeTokenLicensedStakingManagerTransactorSession struct {
	Contract     *INativeTokenLicensedStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// INativeTokenLicensedStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type INativeTokenLicensedStakingManagerRaw struct {
	Contract *INativeTokenLicensedStakingManager // Generic contract binding to access the raw methods on
}

// INativeTokenLicensedStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type INativeTokenLicensedStakingManagerCallerRaw struct {
	Contract *INativeTokenLicensedStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// INativeTokenLicensedStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type INativeTokenLicensedStakingManagerTransactorRaw struct {
	Contract *INativeTokenLicensedStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewINativeTokenLicensedStakingManager creates a new instance of INativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewINativeTokenLicensedStakingManager(address common.Address, backend bind.ContractBackend) (*INativeTokenLicensedStakingManager, error) {
	contract, err := bindINativeTokenLicensedStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManager{INativeTokenLicensedStakingManagerCaller: INativeTokenLicensedStakingManagerCaller{contract: contract}, INativeTokenLicensedStakingManagerTransactor: INativeTokenLicensedStakingManagerTransactor{contract: contract}, INativeTokenLicensedStakingManagerFilterer: INativeTokenLicensedStakingManagerFilterer{contract: contract}}, nil
}

// NewINativeTokenLicensedStakingManagerCaller creates a new read-only instance of INativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewINativeTokenLicensedStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*INativeTokenLicensedStakingManagerCaller, error) {
	contract, err := bindINativeTokenLicensedStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerCaller{contract: contract}, nil
}

// NewINativeTokenLicensedStakingManagerTransactor creates a new write-only instance of INativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewINativeTokenLicensedStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*INativeTokenLicensedStakingManagerTransactor, error) {
	contract, err := bindINativeTokenLicensedStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerTransactor{contract: contract}, nil
}

// NewINativeTokenLicensedStakingManagerFilterer creates a new log filterer instance of INativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewINativeTokenLicensedStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*INativeTokenLicensedStakingManagerFilterer, error) {
	contract, err := bindINativeTokenLicensedStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerFilterer{contract: contract}, nil
}

// bindINativeTokenLicensedStakingManager binds a generic wrapper to an already deployed contract.
func bindINativeTokenLicensedStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := INativeTokenLicensedStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeTokenLicensedStakingManager.Contract.INativeTokenLicensedStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.INativeTokenLicensedStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.INativeTokenLicensedStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _INativeTokenLicensedStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.contract.Transact(opts, method, params...)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCaller) Erc721(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _INativeTokenLicensedStakingManager.contract.Call(opts, &out, "erc721")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) Erc721() (common.Address, error) {
	return _INativeTokenLicensedStakingManager.Contract.Erc721(&_INativeTokenLicensedStakingManager.CallOpts)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerSession) Erc721() (common.Address, error) {
	return _INativeTokenLicensedStakingManager.Contract.Erc721(&_INativeTokenLicensedStakingManager.CallOpts)
}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCaller) GetDelegatorStakedLicenseTokens(opts *bind.CallOpts, delegationID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _INativeTokenLicensedStakingManager.contract.Call(opts, &out, "getDelegatorStakedLicenseTokens", delegationID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) GetDelegatorStakedLicenseTokens(delegationID [32]byte) ([]*big.Int, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetDelegatorStakedLicenseTokens(&_INativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerSession) GetDelegatorStakedLicenseTokens(delegationID [32]byte) ([]*big.Int, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetDelegatorStakedLicenseTokens(&_INativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCaller) GetLicenseTokenDelegator(opts *bind.CallOpts, licenseTokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _INativeTokenLicensedStakingManager.contract.Call(opts, &out, "getLicenseTokenDelegator", licenseTokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) GetLicenseTokenDelegator(licenseTokenId *big.Int) ([32]byte, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetLicenseTokenDelegator(&_INativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerSession) GetLicenseTokenDelegator(licenseTokenId *big.Int) ([32]byte, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetLicenseTokenDelegator(&_INativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCaller) GetLicenseTokenValidator(opts *bind.CallOpts, licenseTokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _INativeTokenLicensedStakingManager.contract.Call(opts, &out, "getLicenseTokenValidator", licenseTokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) GetLicenseTokenValidator(licenseTokenId *big.Int) ([32]byte, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetLicenseTokenValidator(&_INativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerSession) GetLicenseTokenValidator(licenseTokenId *big.Int) ([32]byte, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetLicenseTokenValidator(&_INativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCaller) GetValidatorDelegations(opts *bind.CallOpts, validationID [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _INativeTokenLicensedStakingManager.contract.Call(opts, &out, "getValidatorDelegations", validationID)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) GetValidatorDelegations(validationID [32]byte) ([][32]byte, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetValidatorDelegations(&_INativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerSession) GetValidatorDelegations(validationID [32]byte) ([][32]byte, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetValidatorDelegations(&_INativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCaller) GetValidatorStakedLicenseTokens(opts *bind.CallOpts, validationID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _INativeTokenLicensedStakingManager.contract.Call(opts, &out, "getValidatorStakedLicenseTokens", validationID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) GetValidatorStakedLicenseTokens(validationID [32]byte) ([]*big.Int, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetValidatorStakedLicenseTokens(&_INativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerSession) GetValidatorStakedLicenseTokens(validationID [32]byte) ([]*big.Int, error) {
	return _INativeTokenLicensedStakingManager.Contract.GetValidatorStakedLicenseTokens(&_INativeTokenLicensedStakingManager.CallOpts, validationID)
}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCaller) LicenseToStakeConversionFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _INativeTokenLicensedStakingManager.contract.Call(opts, &out, "licenseToStakeConversionFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) LicenseToStakeConversionFactor() (*big.Int, error) {
	return _INativeTokenLicensedStakingManager.Contract.LicenseToStakeConversionFactor(&_INativeTokenLicensedStakingManager.CallOpts)
}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerCallerSession) LicenseToStakeConversionFactor() (*big.Int, error) {
	return _INativeTokenLicensedStakingManager.Contract.LicenseToStakeConversionFactor(&_INativeTokenLicensedStakingManager.CallOpts)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address recipient) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, recipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address recipient) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ChangeDelegatorRewardRecipient(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, recipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address recipient) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ChangeDelegatorRewardRecipient(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, recipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address recipient) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, recipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address recipient) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ChangeValidatorRewardRecipient(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, recipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address recipient) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, recipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ChangeValidatorRewardRecipient(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, recipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ClaimDelegationFees(&_INativeTokenLicensedStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ClaimDelegationFees(&_INativeTokenLicensedStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteDelegatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteDelegatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) CompleteDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "completeDelegatorRemoval", delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteDelegatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteDelegatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteValidatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteValidatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteValidatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.CompleteValidatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) ForceInitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "forceInitiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ForceInitiateDelegatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ForceInitiateDelegatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) ForceInitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "forceInitiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ForceInitiateValidatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ForceInitiateValidatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) InitiateDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "initiateDelegatorRegistration", validationID, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) InitiateDelegatorRegistration(validationID [32]byte, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateDelegatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) InitiateDelegatorRegistration(validationID [32]byte, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateDelegatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) InitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "initiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateDelegatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateDelegatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateValidatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateValidatorRegistration(&_INativeTokenLicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "initiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateValidatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.InitiateValidatorRemoval(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) ResendUpdateDelegator(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "resendUpdateDelegator", delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ResendUpdateDelegator(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.ResendUpdateDelegator(&_INativeTokenLicensedStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.SubmitUptimeProof(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _INativeTokenLicensedStakingManager.Contract.SubmitUptimeProof(&_INativeTokenLicensedStakingManager.TransactOpts, validationID, messageIndex)
}

// INativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator is returned from FilterCompletedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRegistration events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator struct {
	Event *INativeTokenLicensedStakingManagerCompletedDelegatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
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
		it.Event = new(INativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
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
func (it *INativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerCompletedDelegatorRegistration represents a CompletedDelegatorRegistration event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerCompletedDelegatorRegistration struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRegistration is a free log retrieval operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterCompletedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*INativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "CompletedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRegistration is a free log subscription operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchCompletedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerCompletedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
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

// ParseCompletedDelegatorRegistration is a log parse operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseCompletedDelegatorRegistration(log types.Log) (*INativeTokenLicensedStakingManagerCompletedDelegatorRegistration, error) {
	event := new(INativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator is returned from FilterCompletedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRemoval events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator struct {
	Event *INativeTokenLicensedStakingManagerCompletedDelegatorRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
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
		it.Event = new(INativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
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
func (it *INativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerCompletedDelegatorRemoval represents a CompletedDelegatorRemoval event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerCompletedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterCompletedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*INativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "CompletedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRemoval is a free log subscription operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchCompletedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerCompletedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
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

// ParseCompletedDelegatorRemoval is a log parse operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseCompletedDelegatorRemoval(log types.Log) (*INativeTokenLicensedStakingManagerCompletedDelegatorRemoval, error) {
	event := new(INativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator is returned from FilterDelegatorRegisteredWithLicenses and is used to iterate over the raw logs and unpacked data for DelegatorRegisteredWithLicenses events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator struct {
	Event *INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
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
		it.Event = new(INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
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
func (it *INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses represents a DelegatorRegisteredWithLicenses event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses struct {
	DelegationID    [32]byte
	ValidationID    [32]byte
	LicenseTokenIds []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegisteredWithLicenses is a free log retrieval operation binding the contract event 0xc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb569.
//
// Solidity: event DelegatorRegisteredWithLicenses(bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterDelegatorRegisteredWithLicenses(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "DelegatorRegisteredWithLicenses", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "DelegatorRegisteredWithLicenses", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegisteredWithLicenses is a free log subscription operation binding the contract event 0xc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb569.
//
// Solidity: event DelegatorRegisteredWithLicenses(bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchDelegatorRegisteredWithLicenses(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "DelegatorRegisteredWithLicenses", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRegisteredWithLicenses", log); err != nil {
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

// ParseDelegatorRegisteredWithLicenses is a log parse operation binding the contract event 0xc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb569.
//
// Solidity: event DelegatorRegisteredWithLicenses(bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseDelegatorRegisteredWithLicenses(log types.Log) (*INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses, error) {
	event := new(INativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRegisteredWithLicenses", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator is returned from FilterDelegatorRewardClaimed and is used to iterate over the raw logs and unpacked data for DelegatorRewardClaimed events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator struct {
	Event *INativeTokenLicensedStakingManagerDelegatorRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerDelegatorRewardClaimed)
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
		it.Event = new(INativeTokenLicensedStakingManagerDelegatorRewardClaimed)
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
func (it *INativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerDelegatorRewardClaimed represents a DelegatorRewardClaimed event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerDelegatorRewardClaimed struct {
	DelegationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardClaimed is a free log retrieval operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterDelegatorRewardClaimed(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address) (*INativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "DelegatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardClaimed is a free log subscription operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchDelegatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerDelegatorRewardClaimed, delegationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerDelegatorRewardClaimed)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
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

// ParseDelegatorRewardClaimed is a log parse operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseDelegatorRewardClaimed(log types.Log) (*INativeTokenLicensedStakingManagerDelegatorRewardClaimed, error) {
	event := new(INativeTokenLicensedStakingManagerDelegatorRewardClaimed)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator is returned from FilterDelegatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for DelegatorRewardRecipientChanged events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator struct {
	Event *INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
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
		it.Event = new(INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
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
func (it *INativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged represents a DelegatorRewardRecipientChanged event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged struct {
	DelegationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterDelegatorRewardRecipientChanged(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*INativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "DelegatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchDelegatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
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

// ParseDelegatorRewardRecipientChanged is a log parse operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseDelegatorRewardRecipientChanged(log types.Log) (*INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged, error) {
	event := new(INativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator is returned from FilterInitiatedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRegistration events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator struct {
	Event *INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
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
		it.Event = new(INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
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
func (it *INativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration represents a InitiatedDelegatorRegistration event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration struct {
	DelegationID       [32]byte
	ValidationID       [32]byte
	DelegatorAddress   common.Address
	Nonce              uint64
	ValidatorWeight    uint64
	DelegatorWeight    uint64
	SetWeightMessageID [32]byte
	RewardRecipient    common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRegistration is a free log retrieval operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterInitiatedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*INativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "InitiatedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRegistration is a free log subscription operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchInitiatedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var delegatorAddressRule []interface{}
	for _, delegatorAddressItem := range delegatorAddress {
		delegatorAddressRule = append(delegatorAddressRule, delegatorAddressItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
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

// ParseInitiatedDelegatorRegistration is a log parse operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseInitiatedDelegatorRegistration(log types.Log) (*INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration, error) {
	event := new(INativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator is returned from FilterInitiatedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRemoval events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator struct {
	Event *INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
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
		it.Event = new(INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
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
func (it *INativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval represents a InitiatedDelegatorRemoval event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterInitiatedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*INativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "InitiatedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRemoval is a free log subscription operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchInitiatedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
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

// ParseInitiatedDelegatorRemoval is a log parse operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseInitiatedDelegatorRemoval(log types.Log) (*INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval, error) {
	event := new(INativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator is returned from FilterInitiatedStakingValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedStakingValidatorRegistration events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator struct {
	Event *INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
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
		it.Event = new(INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
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
func (it *INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration represents a InitiatedStakingValidatorRegistration event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration struct {
	ValidationID      [32]byte
	Owner             common.Address
	DelegationFeeBips uint16
	MinStakeDuration  uint64
	RewardRecipient   common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterInitiatedStakingValidatorRegistration is a free log retrieval operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterInitiatedStakingValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, owner []common.Address) (*INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "InitiatedStakingValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedStakingValidatorRegistration is a free log subscription operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchInitiatedStakingValidatorRegistration(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration, validationID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
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

// ParseInitiatedStakingValidatorRegistration is a log parse operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseInitiatedStakingValidatorRegistration(log types.Log) (*INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration, error) {
	event := new(INativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerUptimeUpdatedIterator struct {
	Event *INativeTokenLicensedStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerUptimeUpdated)
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
		it.Event = new(INativeTokenLicensedStakingManagerUptimeUpdated)
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
func (it *INativeTokenLicensedStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*INativeTokenLicensedStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerUptimeUpdatedIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerUptimeUpdated)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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

// ParseUptimeUpdated is a log parse operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*INativeTokenLicensedStakingManagerUptimeUpdated, error) {
	event := new(INativeTokenLicensedStakingManagerUptimeUpdated)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator is returned from FilterValidatorRegisteredWithLicenses and is used to iterate over the raw logs and unpacked data for ValidatorRegisteredWithLicenses events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator struct {
	Event *INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
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
		it.Event = new(INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
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
func (it *INativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses represents a ValidatorRegisteredWithLicenses event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses struct {
	ValidationID    [32]byte
	LicenseTokenIds []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegisteredWithLicenses is a free log retrieval operation binding the contract event 0x0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968.
//
// Solidity: event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterValidatorRegisteredWithLicenses(opts *bind.FilterOpts, validationID [][32]byte) (*INativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "ValidatorRegisteredWithLicenses", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "ValidatorRegisteredWithLicenses", logs: logs, sub: sub}, nil
}

// WatchValidatorRegisteredWithLicenses is a free log subscription operation binding the contract event 0x0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968.
//
// Solidity: event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchValidatorRegisteredWithLicenses(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "ValidatorRegisteredWithLicenses", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRegisteredWithLicenses", log); err != nil {
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

// ParseValidatorRegisteredWithLicenses is a log parse operation binding the contract event 0x0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968.
//
// Solidity: event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseValidatorRegisteredWithLicenses(log types.Log) (*INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses, error) {
	event := new(INativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRegisteredWithLicenses", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerValidatorRewardClaimedIterator is returned from FilterValidatorRewardClaimed and is used to iterate over the raw logs and unpacked data for ValidatorRewardClaimed events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerValidatorRewardClaimedIterator struct {
	Event *INativeTokenLicensedStakingManagerValidatorRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerValidatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerValidatorRewardClaimed)
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
		it.Event = new(INativeTokenLicensedStakingManagerValidatorRewardClaimed)
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
func (it *INativeTokenLicensedStakingManagerValidatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerValidatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerValidatorRewardClaimed represents a ValidatorRewardClaimed event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerValidatorRewardClaimed struct {
	ValidationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardClaimed is a free log retrieval operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterValidatorRewardClaimed(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address) (*INativeTokenLicensedStakingManagerValidatorRewardClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerValidatorRewardClaimedIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "ValidatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardClaimed is a free log subscription operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchValidatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerValidatorRewardClaimed, validationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerValidatorRewardClaimed)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
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

// ParseValidatorRewardClaimed is a log parse operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseValidatorRewardClaimed(log types.Log) (*INativeTokenLicensedStakingManagerValidatorRewardClaimed, error) {
	event := new(INativeTokenLicensedStakingManagerValidatorRewardClaimed)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// INativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator is returned from FilterValidatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for ValidatorRewardRecipientChanged events raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator struct {
	Event *INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log          // Log channel receiving the found contract events
	sub  interfaces.Subscription // Subscription for errors, completion and termination
	done bool                    // Whether the subscription completed delivering logs
	fail error                   // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *INativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
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
		it.Event = new(INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
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
func (it *INativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *INativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged represents a ValidatorRewardRecipientChanged event raised by the INativeTokenLicensedStakingManager contract.
type INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged struct {
	ValidationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) FilterValidatorRewardRecipientChanged(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*INativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.FilterLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &INativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator{contract: _INativeTokenLicensedStakingManager.contract, event: "ValidatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) WatchValidatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}

	logs, sub, err := _INativeTokenLicensedStakingManager.contract.WatchLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
				if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
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

// ParseValidatorRewardRecipientChanged is a log parse operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_INativeTokenLicensedStakingManager *INativeTokenLicensedStakingManagerFilterer) ParseValidatorRewardRecipientChanged(log types.Log) (*INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged, error) {
	event := new(INativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
	if err := _INativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
