// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20licensedstakingmanager

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

// Delegator is an auto generated low-level Go binding around an user-defined struct.
type Delegator struct {
	Status        uint8
	Owner         common.Address
	ValidationID  [32]byte
	Weight        uint64
	StartTime     uint64
	StartingNonce uint64
	EndingNonce   uint64
}

// PChainOwner is an auto generated low-level Go binding around an user-defined struct.
type PChainOwner struct {
	Threshold uint32
	Addresses []common.Address
}

// PoSValidatorInfo is an auto generated low-level Go binding around an user-defined struct.
type PoSValidatorInfo struct {
	Owner             common.Address
	DelegationFeeBips uint16
	MinStakeDuration  uint64
	UptimeSeconds     uint64
}

// StakingManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type StakingManagerSettings struct {
	Manager                  common.Address
	MinimumStakeAmount       *big.Int
	MaximumStakeAmount       *big.Int
	MinimumStakeDuration     uint64
	MinimumDelegationFeeBips uint16
	MaximumStakeMultiplier   uint8
	WeightToValueFactor      *big.Int
	RewardCalculator         common.Address
	UptimeBlockchainID       [32]byte
}

// ERC20LicensedStakingManagerMetaData contains all meta data concerning the ERC20LicensedStakingManager contract.
var ERC20LicensedStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BIPS_CONVERSION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ERC20_LICENSED_STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LICENSED_STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeDelegatorRewardRecipient\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeValidatorRewardRecipient\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRemoval\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"erc20\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20Mintable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"erc721\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceInitiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDelegatorInfo\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDelegator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startingNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endingNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorRewardInfo\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenDelegator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenValidator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingManagerSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakingManagerSettings\",\"components\":[{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"contractIValidatorManager\"},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"},{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorInfo\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"uptimeSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDelegations\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorRewardInfo\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structStakingManagerSettings\",\"components\":[{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"contractIValidatorManager\"},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"},{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"contractIERC20Mintable\"},{\"name\":\"licenseToken\",\"type\":\"address\",\"internalType\":\"contractIERC721\"},{\"name\":\"licenseToStakeConversionFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delegationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRegistration\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"licenseToStakeConversionFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegator\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CompletedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CompletedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardClaimed\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ERC20TokensUnlocked\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedStakingValidatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLicenseToStakeConversionFactor\",\"inputs\":[{\"name\":\"factor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidLicenseTokenCount\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRewardRecipient\",\"inputs\":[{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidUptimeBlockchainID\",\"inputs\":[{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LicenseTokenAlreadyStaked\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedValidationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWeightToValueFactor\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516151b43803806151b483398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6150678061014d5f395ff3fe608060405234801561000f575f80fd5b5060043610610229575f3560e01c8063785e9e861161012a578063b2c1712e116100b4578063c52efd6d11610079578063c52efd6d146105a6578063caa71874146105b9578063e24b2680146105ce578063f9c4f064146105e1578063fb8b11dd1461061f575f80fd5b8063b2c1712e14610556578063b771b3bc14610569578063b82f87df14610577578063bca6ce641461058b578063bdddf02f14610593575f80fd5b806393e24598116100fa57806393e24598146104eb5780639681d940146104fe578063a3a65e4814610511578063a9778a7a14610390578063af1dd66c14610524575f80fd5b8063785e9e861461047e5780637a63ad851461049e5780637b88033a146104c55780638ef34c98146104d8575f80fd5b80632aa56638116101b657806353a133381161017b57806353a13338146103e657806360ad77841461040657806361df835f14610419578063620658561461042d5780636b1470c014610440575f80fd5b80632aa56638146103525780632e2194d81461036557806335455ded14610390578063383caa73146103ac578063401e58ac146103d3575f80fd5b80631c39211f116101fc5780631c39211f14610298578063245dafcb146102b957806325e1c776146102cc5780632674874b146102df57806327bf60cd1461033f575f80fd5b80630f2cb5971461022d5780631340964514610256578063151d30d11461026b5780631667956414610285575b5f80fd5b61024061023b366004614348565b610632565b60405161024d919061435f565b60405180910390f35b6102696102643660046143b5565b6106b0565b005b610273600a81565b60405160ff909116815260200161024d565b6102696102933660046143ec565b6109f0565b6102ab6102a636600461468d565b610a01565b60405190815260200161024d565b6102696102c7366004614348565b610a42565b6102696102da3660046143b5565b610d06565b6102f26102ed366004614348565b610de4565b6040805182516001600160a01b0316815260208084015161ffff1690820152828201516001600160401b03908116928201929092526060928301519091169181019190915260800161024d565b61026961034d3660046143ec565b610e70565b6102696103603660046143ec565b610e7b565b610378610373366004614348565b610e8b565b6040516001600160401b03909116815260200161024d565b61039961271081565b60405161ffff909116815260200161024d565b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb01546102ab565b6102696103e1366004614796565b610edf565b6103f96103f4366004614348565b610ff1565b60405161024d9190614815565b6102696104143660046143b5565b6110de565b6102ab5f80516020614ff283398151915281565b6102ab61043b366004614884565b611400565b6102ab61044e366004614348565b5f9081527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb05602052604090205490565b610486611420565b6040516001600160a01b03909116815260200161024d565b6102ab7fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60081565b6102406104d3366004614348565b61143f565b6102696104e636600461489f565b6114bb565b6102696104f9366004614348565b61159d565b6102ab61050c3660046148cd565b6116b6565b6102ab61051f3660046148cd565b611858565b610537610532366004614348565b6118d0565b604080516001600160a01b03909316835260208301919091520161024d565b6102696105643660046143ec565b61190c565b6104866005600160991b0181565b6102ab5f80516020614fb283398151915281565b610486611917565b6102406105a1366004614348565b61192b565b6102ab6105b43660046148e6565b611994565b6105c16119cb565b60405161024d919061494b565b6105376105dc366004614348565b611aa7565b6102ab6105ef366004614348565b5f9081527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb04602052604090205490565b61026961062d36600461489f565b611ae3565b5f8181527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0360209081526040918290208054835181840281018401909452808452606093928301828280156106a457602002820191905f5260205f20905b815481526020019060010190808311610690575b50505050509050919050565b6106b8611bab565b5f6106c1611be2565b5f848152600882016020526040808220815160e0810190925280549394509192909190829060ff1660038111156106fa576106fa6147ed565b600381111561070b5761070b6147ed565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101529050600381516003811115610784576107846147ed565b146107ae578051604051633b0d540d60e21b81526107a591906004016149ec565b60405180910390fd5b81546040828101519051636af907fb60e11b815260048101919091525f916001600160a01b03169063d5f20ff6906024015f60405180830381865afa1580156107f9573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526108209190810190614a60565b9050600483546040808501519051636af907fb60e11b81526001600160a01b039092169163d5f20ff69161085a9160040190815260200190565b5f60405180830381865afa158015610874573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261089b9190810190614a60565b5160058111156108ad576108ad6147ed565b141580156108d457508160c001516001600160401b031681608001516001600160401b0316105b156109ca57825460405163338587c560e21b815263ffffffff861660048201525f9182916001600160a01b039091169063ce161f149060240160408051808303815f875af1158015610928573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094c9190614b3f565b91509150818460400151146109855781846040015160405163fee3144560e01b81526004016107a5929190918252602082015260400190565b806001600160401b03168460c001516001600160401b031611156109c757604051632e19bc2d60e11b81526001600160401b03821660048201526024016107a5565b50505b6109d385611c06565b5050506109ec60015f8051602061501283398151915255565b5050565b6109fb838383611e54565b50505050565b5f610a0a611bab565b610a1c8b8b8b8b8b8b8b8b8b8b6120fb565b9050610a3460015f8051602061501283398151915255565b9a9950505050505050505050565b5f610a4b611be2565b5f838152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115610a8457610a846147ed565b6003811115610a9557610a956147ed565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115610b0e57610b0e6147ed565b14158015610b2f5750600381516003811115610b2c57610b2c6147ed565b14155b15610b50578051604051633b0d540d60e21b81526107a591906004016149ec565b81546040828101519051636af907fb60e11b815260048101919091525f916001600160a01b03169063d5f20ff6906024015f60405180830381865afa158015610b9b573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610bc29190810190614a60565b905080606001516001600160401b03165f03610bf4576040516339b894f960e21b8152600481018590526024016107a5565b604080830151606083015160a0840151925163854a893f60e01b81526005600160991b019363ee5b48eb9373__$caf7a49f90ce02b15fb10b3b072d8b9489$__9363854a893f93610c6293906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015610c7c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610ca39190810190614b62565b6040518263ffffffff1660e01b8152600401610cbf9190614bc9565b6020604051808303815f875af1158015610cdb573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cff9190614bdb565b5050505050565b610d0f826121ab565b610d2f576040516330efa98b60e01b8152600481018390526024016107a5565b5f610d38611be2565b54604051636af907fb60e11b8152600481018590526001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015610d7c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610da39190810190614a60565b5190506002816005811115610dba57610dba6147ed565b14610dda578060405163170cc93360e21b81526004016107a59190614bf2565b6109fb83836121d4565b604080516080810182525f808252602082018190529181018290526060810191909152610e0f611be2565b5f9283526007016020908152604092839020835160808101855281546001600160a01b038116825261ffff600160a01b820416938201939093526001600160401b03600160b01b909304831694810194909452600101541660608301525090565b6109fb83838361243e565b610e8683838361285f565b505050565b5f80610e95611be2565b60040154610ea39084614c20565b9050801580610eb857506001600160401b0381115b15610ed95760405163222d164360e21b8152600481018490526024016107a5565b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015610f235750825b90505f826001600160401b03166001148015610f3e5750303b155b905081158015610f4c575080155b15610f6a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610f9457845460ff60401b1916600160401b1785555b610fa08989898961288a565b8315610fe657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810191909152611031611be2565b5f838152600891909101602052604090819020815160e081019092528054829060ff166003811115611065576110656147ed565b6003811115611076576110766147ed565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015292915050565b5f6110e7611be2565b5f848152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115611120576111206147ed565b6003811115611131576111316147ed565b8152815461010090046001600160a01b03908116602083015260018301546040808401919091526002909301546001600160401b038082166060850152600160401b820481166080850152600160801b8204811660a0850152600160c01b9091041660c0909201919091528282015185549251636af907fb60e11b815260048101829052939450925f929091169063d5f20ff6906024015f60405180830381865afa1580156111e2573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526112099190810190614a60565b9050600183516003811115611220576112206147ed565b14611241578251604051633b0d540d60e21b81526107a591906004016149ec565b600481516005811115611256576112566147ed565b0361126c5761126486611c06565b505050505050565b8260a001516001600160401b031681608001516001600160401b0316101561137457835460405163338587c560e21b815263ffffffff871660048201525f9182916001600160a01b039091169063ce161f149060240160408051808303815f875af11580156112dd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113019190614b3f565b9150915081841461132f5760405163fee3144560e01b815260048101839052602481018590526044016107a5565b8460a001516001600160401b0316816001600160401b0316101561137157604051632e19bc2d60e11b81526001600160401b03821660048201526024016107a5565b50505b5f868152600885016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b81026fffffffffffffffff000000000000000019909216919091179091559151918252839188917f3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa910160405180910390a3505050505050565b5f611409611be2565b60040154610ed9906001600160401b038416614c3f565b5f5f80516020614fb28339815191525b546001600160a01b0316919050565b5f8181527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0660209081526040918290208054835181840281018401909452808452606093928301828280156106a457602002820191905f5260205f20908154815260200190600101908083116106905750505050509050919050565b5f6114c4611be2565b90506001600160a01b0382166114f85760405163caa903f960e01b81526001600160a01b03831660048201526024016107a5565b5f8381526007820160205260409020546001600160a01b0316331461153e57335b604051636e2ccd7560e11b81526001600160a01b0390911660048201526024016107a5565b5f838152600c8201602052604080822080546001600160a01b038681166001600160a01b0319831681179093559251921692839287917f28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c3391a450505050565b5f6115a6611be2565b8054604051636af907fb60e11b8152600481018590529192505f916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa1580156115f0573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526116179190810190614a60565b519050600481600581111561162e5761162e6147ed565b1461164e578060405163170cc93360e21b81526004016107a59190614bf2565b5f8381526007830160205260409020546001600160a01b031633146116735733611519565b5f838152600c830160205260409020546001600160a01b0316806116ac57505f8381526007830160205260409020546001600160a01b03165b6109fb81856128a6565b5f6116bf611bab565b5f6116c8611be2565b805460405163025a076560e61b815263ffffffff861660048201529192505f916001600160a01b0390911690639681d940906024016020604051808303815f875af1158015611719573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061173d9190614bdb565b8254604051636af907fb60e11b8152600481018390529192505f916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015611787573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526117ae9190810190614a60565b90506117b9826121ab565b6117c75750915061183d9050565b5f828152600784016020908152604080832054600c8701909252909120546001600160a01b039182169116806117fa5750805b60048351600581111561180f5761180f6147ed565b0361181e5761181e81856128a6565b6118358261182f8560400151611400565b8661291a565b509193505050505b61185360015f8051602061501283398151915255565b919050565b5f611861611be2565b54604051631474cbc960e31b815263ffffffff841660048201526001600160a01b039091169063a3a65e48906024016020604051808303815f875af11580156118ac573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ed99190614bdb565b5f805f6118db611be2565b5f948552600a810160209081526040808720546009909301909152909420546001600160a01b039094169492505050565b610e8683838361299e565b5f5f80516020614ff2833981519152611430565b5f8181525f80516020614fd283398151915260209081526040918290208054835181840281018401909452808452606093928301828280156106a457602002820191905f5260205f20908154815260200190600101908083116106905750505050509050919050565b5f61199d611bab565b6119aa86868686866129c9565b90506119c260015f8051602061501283398151915255565b95945050505050565b60408051610120810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905290611a1b611be2565b604080516101208101825282546001600160a01b0390811682526001840154602083015260028401549282019290925260038301546001600160401b0381166060830152600160401b810461ffff166080830152600160501b900460ff1660a0820152600483015460c0820152600583015490911660e082015260069091015461010082015292915050565b5f805f611ab2611be2565b5f948552600c81016020908152604080872054600b909301909152909420546001600160a01b039094169492505050565b6001600160a01b038116611b155760405163caa903f960e01b81526001600160a01b03821660048201526024016107a5565b5f611b1e611be2565b5f8481526008820160205260409020549091506001600160a01b03610100909104163314611b4c5733611519565b5f838152600a8201602052604080822080546001600160a01b038681166001600160a01b0319831681179093559251921692839287917f6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e00591a450505050565b5f80516020615012833981519152805460011901611bdc57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b7fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60090565b5f611c0f611be2565b5f838152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115611c4857611c486147ed565b6003811115611c5957611c596147ed565b815281546001600160a01b03610100909104811660208084019190915260018401546040808501919091526002909401546001600160401b038082166060860152600160401b820481166080860152600160801b8204811660a0860152600160c01b9091041660c09093019290925283830151865484516304e0efb360e11b8152945195965090949116926309c1df669260048083019391928290030181865afa158015611d09573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611d2d9190614c56565b8260800151611d3c9190614c71565b6001600160401b0316421015611d705760405163fb6ce63f60e01b81526001600160401b03421660048201526024016107a5565b5f848152600a84016020526040902080546001600160a01b031981169091556001600160a01b031680611da4575060208201515b5f80611db1838886612a4d565b91509150611dd08560200151611dca8760600151611400565b8961291a565b5f878152600887016020908152604080832080546001600160a81b031916815560018101849055600201929092558151848152908101839052859189917f5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e910160405180910390a350505050505050565b60015f8051602061501283398151915255565b5f80611e5e611be2565b8054604051635b73516560e11b8152600481018890529192506001600160a01b03169063b6e6a2ca906024015f604051808303815f87803b158015611ea1575f80fd5b505af1158015611eb3573d5f803e3d5ffd5b50508254604051636af907fb60e11b8152600481018990525f93506001600160a01b03909116915063d5f20ff6906024015f60405180830381865afa158015611efe573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611f259190810190614a60565b9050611f30866121ab565b611f3f576001925050506120f4565b5f8681526007830160205260409020546001600160a01b03163314611f645733611519565b5f86815260078301602052604090205460c0820151611f9391600160b01b90046001600160401b031690614c71565b6001600160401b03168160e001516001600160401b03161015611fda5760e081015160405163fb6ce63f60e01b81526001600160401b0390911660048201526024016107a5565b5f8515611ff257611feb87866121d4565b9050612010565b505f8681526007830160205260409020600101546001600160401b03165b600583015460408301515f916001600160a01b031690634f22429f9061203590611400565b60c086015160e0808801516040519185901b6001600160e01b031916825260048201939093526001600160401b0391821660248201819052604482015291811660648301528516608482015260a401602060405180830381865afa15801561209f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120c39190614bdb565b90508084600b015f8a81526020019081526020015f205f8282546120e79190614c98565b9091555050151593505050505b9392505050565b5f82810361211f57604051639c8cf5c960e01b8152600481018490526024016107a5565b61212885612b4c565b5f612134868686612b7f565b90505f6121478d8d8d8d8d8d888b612bc6565b905061215781878760015f612ef2565b6121608761307f565b50807f0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f69688787604051612193929190614cab565b60405180910390a29c9b505050505050505050505050565b5f806121b5611be2565b5f938452600701602052505060409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa15801561221f573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526122469190810190614ce2565b915091508061226857604051636b2f19e960e01b815260040160405180910390fd5b5f612271611be2565b600681015484519192501461229f578251604051636ba589a560e01b815260048101919091526024016107a5565b60208301516001600160a01b0316156122db576020830151604051624de75d60e31b81526001600160a01b0390911660048201526024016107a5565b5f8073__$caf7a49f90ce02b15fb10b3b072d8b9489$__63088c246386604001516040518263ffffffff1660e01b81526004016123189190614bc9565b6040805180830381865af4158015612332573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123569190614b3f565b915091508188146123845760405163fee3144560e01b815260048101839052602481018990526044016107a5565b5f8881526007840160205260409020600101546001600160401b039081169082161115612415575f888152600784016020908152604091829020600101805467ffffffffffffffff19166001600160401b038516908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a2612433565b505f8781526007830160205260409020600101546001600160401b03165b979650505050505050565b5f80612448611be2565b5f868152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115612481576124816147ed565b6003811115612492576124926147ed565b8152815461010090046001600160a01b03908116602083015260018301546040808401919091526002909301546001600160401b038082166060850152600160401b820481166080850152600160801b8204811660a0850152600160c01b9091041660c0909201919091528282015185549251636af907fb60e11b815260048101829052939450925f929091169063d5f20ff6906024015f60405180830381865afa158015612543573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261256a9190810190614a60565b9050600283516003811115612581576125816147ed565b146125a2578251604051633b0d540d60e21b81526107a591906004016149ec565b60208301516001600160a01b0316331461263e575f8281526007850160205260409020546001600160a01b031633146125db5733611519565b5f82815260078501602052604090205460c082015161260a91600160b01b90046001600160401b031690614c71565b6001600160401b031642101561263e5760405163fb6ce63f60e01b81526001600160401b03421660048201526024016107a5565b5f888152600a850160205260409020546001600160a01b031660028251600581111561266c5761266c6147ed565b03612806576003850154608085015161268e916001600160401b031690614c71565b6001600160401b03164210156126c25760405163fb6ce63f60e01b81526001600160401b03421660048201526024016107a5565b87156126d4576126d283886121d4565b505b5f8981526008860160205260409020805460ff191660031790558454606085015160a08401516001600160a01b039092169163661096699186916127189190614d88565b6040516001600160e01b031960e085901b16815260048101929092526001600160401b0316602482015260440160408051808303815f875af1158015612760573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127849190614da8565b505f8a8152600887016020526040812060020180546001600160401b03909316600160c01b026001600160c01b03909316929092179091556127c785838c6130a9565b9050838a7f5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf60405160405180910390a3151595506120f4945050505050565b60048251600581111561281b5761281b6147ed565b036128435761282b84828b6130a9565b5061283589611c06565b6001955050505050506120f4565b815160405163170cc93360e21b81526107a59190600401614bf2565b61286a83838361243e565b610e8657604051631036cf9160e11b8152600481018490526024016107a5565b6128926132e3565b61289d84838361332e565b6109fb8361337c565b5f6128af611be2565b5f838152600b82016020526040812080549190559091506128d084826133d9565b836001600160a01b0316837f875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e48360405161290c91815260200190565b60405180910390a350505050565b5f612925828461344f565b90506129328484846134c0565b5f80516020614fb283398151915254612955906001600160a01b03168583613507565b836001600160a01b03167fb96e58ee5d4ea44c7f5b0ec141397ea1d47f915b0d70d6b2355a6ced826ce3498260405161299091815260200190565b60405180910390a250505050565b6129a9838383611e54565b610e8657604051635bff683f60e11b8152600481018490526024016107a5565b5f6129d385612b4c565b5f6129df868686612b7f565b90505f6129ee88338487613566565b90506129fd8187875f8c612ef2565b612a068761307f565b5087817fc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb5698888604051612a3a929190614cab565b60405180910390a3979650505050505050565b5f805f612a58611be2565b5f8681526009820160205260408120549192509081908015612b3e575f88815260098501602090815260408083208390558983526007870190915290205461271090612aaf90600160a01b900461ffff1683614c3f565b612ab99190614c20565b91508184600b015f8981526020019081526020015f205f828254612add9190614c98565b90915550612aed90508282614dd4565b9250612af989846133d9565b886001600160a01b0316887f3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf6385604051612b3591815260200190565b60405180910390a35b509097909650945050505050565b5f612b55611be2565b905080600401548210156109ec5760405163514205d560e11b8152600481018390526024016107a5565b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb01545f905f80516020614ff283398151915290612bbc9084614c3f565b6119c29086614c98565b5f80612bd0611be2565b600381015490915061ffff600160401b90910481169087161080612bf9575061271061ffff8716115b15612c1d57604051635f12e6c360e11b815261ffff871660048201526024016107a5565b60038101546001600160401b039081169086161015612c59576040516202a06d60e11b81526001600160401b03861660048201526024016107a5565b8060010154841080612c6e5750806002015484115b15612c8f5760405163222d164360e21b8152600481018590526024016107a5565b6001600160a01b038316612cc15760405163caa903f960e01b81526001600160a01b03841660048201526024016107a5565b835f612ccc82610e8b565b90505f835f015f9054906101000a90046001600160a01b03166001600160a01b0316639cb7624e8e8e8e8e876040518663ffffffff1660e01b8152600401612d18959493929190614e4d565b6020604051808303815f875af1158015612d34573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d589190614bdb565b90505f33905080856007015f8481526020019081526020015f205f015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555089856007015f8481526020019081526020015f205f0160146101000a81548161ffff021916908361ffff16021790555088856007015f8481526020019081526020015f205f0160166101000a8154816001600160401b0302191690836001600160401b031602179055505f856007015f8481526020019081526020015f206001015f6101000a8154816001600160401b0302191690836001600160401b031602179055508685600c015f8481526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b0316827ff51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f8c8c8b604051612ed99392919061ffff9390931683526001600160401b039190911660208301526001600160a01b0316604082015260600190565b60405180910390a3509c9b505050505050505050505050565b5f80516020614ff28339815191525f5b84811015613052575f868683818110612f1d57612f1d614eb5565b9050602002013590505f801b836004015f8381526020019081526020015f2054141580612f5857505f81815260058401602052604090205415155b15612f7957604051633989863160e21b8152600481018290526024016107a5565b82546001600160a01b03166323b872dd3330846040518463ffffffff1660e01b8152600401612faa93929190614ec9565b5f604051808303815f87803b158015612fc1575f80fd5b505af1158015612fd3573d5f803e3d5ffd5b505050508415613015575f8881526002840160209081526040808320805460018101825590845282842001849055838352600486019091529020889055613049565b5f88815260038401602090815260408083208054600181018255908452828420018490558383526005860190915290208890555b50600101612f02565b5082611264575f918252600601602090815260408220805460018101825590835291200193909355505050565b5f6130a2825f80516020614fb2833981519152546001600160a01b0316906139ad565b5090919050565b5f806130b3611be2565b80546040808801519051636af907fb60e11b81529293505f926001600160a01b039092169163d5f20ff6916130ee9160040190815260200190565b5f60405180830381865afa158015613108573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261312f9190810190614a60565b90505f600382516005811115613147576131476147ed565b14806131655750600482516005811115613163576131636147ed565b145b15613175575060e0810151613192565b60028251600581111561318a5761318a6147ed565b036128435750425b86608001516001600160401b0316816001600160401b0316116131ba575f93505050506120f4565b600583015460608801515f916001600160a01b031690634f22429f906131df90611400565b60c086015160808c01516040808e01515f90815260078b0160205281902060010154905160e086901b6001600160e01b031916815260048101949094526001600160401b0392831660248501529082166044840152818716606484015216608482015260a401602060405180830381865afa158015613260573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132849190614bdb565b90506001600160a01b03871661329c57876020015196505b5f8681526009850160209081526040808320849055600a90960190529390932080546001600160a01b0388166001600160a01b031990911617905550909150509392505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661332c57604051631afcd79f60e31b815260040160405180910390fd5b565b6133366132e3565b5f8111801561334857508260c0013581105b1561336957604051634c5a73bf60e01b8152600481018290526024016107a5565b613372836139b9565b610e868282613a37565b6133846132e3565b5f80516020614fb28339815191526001600160a01b0382166133b95760405163d92e233d60e01b815260040160405180910390fd5b80546001600160a01b0319166001600160a01b0392909216919091179055565b5f5f80516020614fb283398151915280546040516340c10f1960e01b81526001600160a01b038681166004830152602482018690529293509116906340c10f19906044015f604051808303815f87803b158015613434575f80fd5b505af1158015613446573d5f803e3d5ffd5b50505050505050565b5f8281525f80516020614fd283398151915260205260408120545f80516020614ff28339815191529082901561349557505f8481526002820160205260409020546134a7565b505f8481526003820160205260409020545b60018201546134b69082614c3f565b6119c29085614dd4565b5f8181525f80516020614fd283398151915260205260409020545f80516020614ff283398151915290156134fd576134f88285613a9a565b6109fb565b6109fb8285613b7c565b6040516001600160a01b03838116602483015260448201839052610e8691859182169063a9059cbb906064015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050613d83565b5f80613570611be2565b90505f61357c85610e8b565b9050613587876121ab565b6135a7576040516330efa98b60e01b8152600481018890526024016107a5565b6001600160a01b0384166135d95760405163caa903f960e01b81526001600160a01b03851660048201526024016107a5565b8154604051636af907fb60e11b8152600481018990525f9182916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015613622573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526136499190810190614a60565b9050828160a0015161365b9190614c71565b915083600301600a9054906101000a90046001600160401b031681604001516136849190614eed565b6001600160401b0316826001600160401b031611156136c157604051636d51fe0560e11b81526001600160401b03831660048201526024016107a5565b508254604051636610966960e01b8152600481018a90526001600160401b03831660248201525f9182916001600160a01b039091169063661096699060440160408051808303815f875af115801561371b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061373f9190614da8565b915091505f8a8360405160200161376d92919091825260c01b6001600160c01b031916602082015260280190565b60408051601f1981840301815291815281516020928301205f81815260088a019093529120805491925060019160ff19168280021790555089866008015f8381526020019081526020015f205f0160016101000a8154816001600160a01b0302191690836001600160a01b031602179055508a866008015f8381526020019081526020015f206001018190555084866008015f8381526020019081526020015f206002015f6101000a8154816001600160401b0302191690836001600160401b031602179055505f866008015f8381526020019081526020015f2060020160086101000a8154816001600160401b0302191690836001600160401b0316021790555082866008015f8381526020019081526020015f2060020160106101000a8154816001600160401b0302191690836001600160401b031602179055505f866008015f8381526020019081526020015f2060020160186101000a8154816001600160401b0302191690836001600160401b031602179055508786600a015f8381526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550896001600160a01b03168b827f77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d7786888a888f6040516139979594939291906001600160401b039586168152938516602085015291909316604083015260608201929092526001600160a01b0391909116608082015260a00190565b60405180910390a49a9950505050505050505050565b5f6120f4833384613de4565b6139c16132e3565b6139c9613f47565b613a346139d96020830183614f18565b602083013560408401356139f36080860160608701614884565b613a0360a0870160808801614f33565b613a1360c0880160a08901614f4c565b60c0880135613a296101008a0160e08b01614f18565b896101000135613f57565b50565b613a3f6132e3565b5f80516020614ff28339815191526001600160a01b038316613a745760405163d92e233d60e01b815260040160405180910390fd5b80546001600160a01b0319166001600160a01b0393909316929092178255600190910155565b5f8281525f80516020614fd2833981519152602052604081205f80516020614ff2833981519152915b8154811015613b63575f828281548110613adf57613adf614eb5565b5f9182526020909120015484546040516323b872dd60e01b81529192506001600160a01b0316906323b872dd90613b1e90309089908690600401614ec9565b5f604051808303815f87803b158015613b35575f80fd5b505af1158015613b47573d5f803e3d5ffd5b5050505f91825250600484016020526040812055600101613ac3565b505f84815260028301602052604081206109fb9161431a565b5f8281527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb03602052604081205f80516020614ff2833981519152915b8154811015613c58575f828281548110613bd457613bd4614eb5565b5f9182526020909120015484546040516323b872dd60e01b81529192506001600160a01b0316906323b872dd90613c1390309089908690600401614ec9565b5f604051808303815f87803b158015613c2a575f80fd5b505af1158015613c3c573d5f803e3d5ffd5b5050505f91825250600584016020526040812055600101613bb8565b505f8481526003830160205260408120613c719161431a565b5f613c7a611be2565b5f868152600891909101602090815260408083206001015480845260068701909252822054909250905b81811015613446575f8381526006860160205260409020805488919083908110613cd057613cd0614eb5565b905f5260205f20015403613d7b575f8381526006860160205260409020613cf8600184614dd4565b81548110613d0857613d08614eb5565b905f5260205f200154856006015f8581526020019081526020015f208281548110613d3557613d35614eb5565b905f5260205f200181905550846006015f8481526020019081526020015f20805480613d6357613d63614f6c565b600190038181905f5260205f20015f90559055613446565b600101613ca4565b5f613d976001600160a01b038416836141ce565b905080515f14158015613dbb575080806020019051810190613db99190614f80565b155b15610e8657604051635274afe760e01b81526001600160a01b03841660048201526024016107a5565b6040516370a0823160e01b81523060048201525f9081906001600160a01b038616906370a0823190602401602060405180830381865afa158015613e2a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e4e9190614bdb565b9050613e656001600160a01b0386168530866141db565b6040516370a0823160e01b81523060048201525f906001600160a01b038716906370a0823190602401602060405180830381865afa158015613ea9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ecd9190614bdb565b9050818111613f335760405162461bcd60e51b815260206004820152602c60248201527f5361666545524332305472616e7366657246726f6d3a2062616c616e6365206e60448201526b1bdd081a5b98dc99585cd95960a21b60648201526084016107a5565b613f3d8282614dd4565b9695505050505050565b613f4f6132e3565b61332c614203565b613f5f6132e3565b5f613f68611be2565b905061ffff86161580613f80575061271061ffff8716115b15613fa457604051635f12e6c360e11b815261ffff871660048201526024016107a5565b87891115613fc85760405163222d164360e21b8152600481018a90526024016107a5565b60ff85161580613fdb5750600a60ff8616115b15613ffe5760405163170db35960e31b815260ff861660048201526024016107a5565b6001600160a01b038a166140255760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b03831661404c5760405163d92e233d60e01b815260040160405180910390fd5b896001600160a01b03166309c1df666040518163ffffffff1660e01b8152600401602060405180830381865afa158015614088573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906140ac9190614c56565b6001600160401b0316876001600160401b031610156140e8576040516202a06d60e11b81526001600160401b03881660048201526024016107a5565b835f036141085760405163a733007160e01b815260040160405180910390fd5b8161412957604051632f6bd1db60e01b8152600481018390526024016107a5565b80546001600160a01b039a8b166001600160a01b031991821617825560018201999099556002810197909755600387018054600160501b60ff9096169590950267ffffffffffffffff60501b1961ffff909716600160401b0269ffffffffffffffffffff199096166001600160401b03909816979097179490941794909416949094179091556004840155600583018054929095169190931617909255600690910155565b60606120f483835f61420b565b6109fb84856001600160a01b03166323b872dd86868660405160240161353493929190614ec9565b611e416132e3565b6060814710156142305760405163cd78605960e01b81523060048201526024016107a5565b5f80856001600160a01b0316848660405161424b9190614f9b565b5f6040518083038185875af1925050503d805f8114614285576040519150601f19603f3d011682016040523d82523d5f602084013e61428a565b606091505b5091509150613f3d8683836060826142aa576142a5826142f1565b6120f4565b81511580156142c157506001600160a01b0384163b155b156142ea57604051639996b31560e01b81526001600160a01b03851660048201526024016107a5565b50806120f4565b8051156143015780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b5080545f8255905f5260205f2090810190613a3491905b80821115614344575f8155600101614331565b5090565b5f60208284031215614358575f80fd5b5035919050565b602080825282518282018190525f9190848201906040850190845b818110156143965783518352928401929184019160010161437a565b50909695505050505050565b803563ffffffff81168114611853575f80fd5b5f80604083850312156143c6575f80fd5b823591506143d6602084016143a2565b90509250929050565b8015158114613a34575f80fd5b5f805f606084860312156143fe575f80fd5b833592506020840135614410816143df565b915061441e604085016143a2565b90509250925092565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b038111828210171561445d5761445d614427565b60405290565b60405161010081016001600160401b038111828210171561445d5761445d614427565b604051601f8201601f191681016001600160401b03811182821017156144ae576144ae614427565b604052919050565b5f6001600160401b038211156144ce576144ce614427565b50601f01601f191660200190565b5f82601f8301126144eb575f80fd5b81356144fe6144f9826144b6565b614486565b818152846020838601011115614512575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160a01b0381168114613a34575f80fd5b80356118538161452e565b5f6040828403121561455d575f80fd5b61456561443b565b9050614570826143a2565b81526020808301356001600160401b038082111561458c575f80fd5b818501915085601f83011261459f575f80fd5b8135818111156145b1576145b1614427565b8060051b91506145c2848301614486565b81815291830184019184810190888411156145db575f80fd5b938501935b8385101561460557843592506145f58361452e565b82825293850193908501906145e0565b808688015250505050505092915050565b803561ffff81168114611853575f80fd5b6001600160401b0381168114613a34575f80fd5b803561185381614627565b5f8083601f840112614656575f80fd5b5081356001600160401b0381111561466c575f80fd5b6020830191508360208260051b8501011115614686575f80fd5b9250929050565b5f805f805f805f805f806101208b8d0312156146a7575f80fd5b8a356001600160401b03808211156146bd575f80fd5b6146c98e838f016144dc565b9b5060208d01359150808211156146de575f80fd5b6146ea8e838f016144dc565b9a5060408d01359150808211156146ff575f80fd5b61470b8e838f0161454d565b995060608d0135915080821115614720575f80fd5b61472c8e838f0161454d565b985061473a60808e01614616565b975061474860a08e0161463b565b965060c08d0135955060e08d0135915080821115614764575f80fd5b506147718d828e01614646565b909450925061478590506101008c01614542565b90509295989b9194979a5092959850565b5f805f808486036101808112156147ab575f80fd5b610120808212156147ba575f80fd5b86955085013590506147cb8161452e565b92506101408501356147dc8161452e565b939692955092936101600135925050565b634e487b7160e01b5f52602160045260245ffd5b60048110614811576148116147ed565b9052565b5f60e082019050614827828451614801565b60018060a01b0360208401511660208301526040830151604083015260608301516001600160401b0380821660608501528060808601511660808501528060a08601511660a08501528060c08601511660c0850152505092915050565b5f60208284031215614894575f80fd5b81356120f481614627565b5f80604083850312156148b0575f80fd5b8235915060208301356148c28161452e565b809150509250929050565b5f602082840312156148dd575f80fd5b6120f4826143a2565b5f805f805f608086880312156148fa575f80fd5b853594506020860135935060408601356001600160401b0381111561491d575f80fd5b61492988828901614646565b909450925050606086013561493d8161452e565b809150509295509295909350565b81516001600160a01b03168152602080830151908201526040808301519082015260608083015161012083019161498c908401826001600160401b03169052565b5060808301516149a2608084018261ffff169052565b5060a08301516149b760a084018260ff169052565b5060c083015160c083015260e08301516149dc60e08401826001600160a01b03169052565b5061010092830151919092015290565b60208101610ed98284614801565b805160068110611853575f80fd5b5f82601f830112614a17575f80fd5b8151614a256144f9826144b6565b818152846020838601011115614a39575f80fd5b8160208501602083015e5f918101602001919091529392505050565b805161185381614627565b5f60208284031215614a70575f80fd5b81516001600160401b0380821115614a86575f80fd5b908301906101008286031215614a9a575f80fd5b614aa2614463565b614aab836149fa565b8152602083015182811115614abe575f80fd5b614aca87828601614a08565b602083015250614adc60408401614a55565b6040820152614aed60608401614a55565b6060820152614afe60808401614a55565b6080820152614b0f60a08401614a55565b60a0820152614b2060c08401614a55565b60c0820152614b3160e08401614a55565b60e082015295945050505050565b5f8060408385031215614b50575f80fd5b8251915060208301516148c281614627565b5f60208284031215614b72575f80fd5b81516001600160401b03811115614b87575f80fd5b614b9384828501614a08565b949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6120f46020830184614b9b565b5f60208284031215614beb575f80fd5b5051919050565b6020810160068310614c0657614c066147ed565b91905290565b634e487b7160e01b5f52601160045260245ffd5b5f82614c3a57634e487b7160e01b5f52601260045260245ffd5b500490565b8082028115828204841417610ed957610ed9614c0c565b5f60208284031215614c66575f80fd5b81516120f481614627565b6001600160401b03818116838216019080821115614c9157614c91614c0c565b5092915050565b80820180821115610ed957610ed9614c0c565b602080825281018290525f6001600160fb1b03831115614cc9575f80fd5b8260051b80856040850137919091016040019392505050565b5f8060408385031215614cf3575f80fd5b82516001600160401b0380821115614d09575f80fd5b9084019060608287031215614d1c575f80fd5b604051606081018181108382111715614d3757614d37614427565b604052825181526020830151614d4c8161452e565b6020820152604083015182811115614d62575f80fd5b614d6e88828601614a08565b60408301525080945050505060208301516148c2816143df565b6001600160401b03828116828216039080821115614c9157614c91614c0c565b5f8060408385031215614db9575f80fd5b8251614dc481614627565b6020939093015192949293505050565b81810381811115610ed957610ed9614c0c565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015614e425784516001600160a01b03168252938301936001929092019190830190614e19565b509695505050505050565b60a081525f614e5f60a0830188614b9b565b8281036020840152614e718188614b9b565b90508281036040840152614e858187614de7565b90508281036060840152614e998186614de7565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603260045260245ffd5b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160401b03818116838216028082169190828114614f1057614f10614c0c565b505092915050565b5f60208284031215614f28575f80fd5b81356120f48161452e565b5f60208284031215614f43575f80fd5b6120f482614616565b5f60208284031215614f5c575f80fd5b813560ff811681146120f4575f80fd5b634e487b7160e01b5f52603160045260245ffd5b5f60208284031215614f90575f80fd5b81516120f4816143df565b5f82518060208501845e5f92019182525091905056fe9d37ef67e865cad1eb988c62f5e45a5866d6dd4ddd905252e31276591c701c0019dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0219dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a264697066735822122004b45233373ecfdb7ac9b1951b8ea110397d24bc39da0798e1b911d79cd8014d64736f6c63430008190033",
}

// ERC20LicensedStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20LicensedStakingManagerMetaData.ABI instead.
var ERC20LicensedStakingManagerABI = ERC20LicensedStakingManagerMetaData.ABI

// ERC20LicensedStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20LicensedStakingManagerMetaData.Bin instead.
var ERC20LicensedStakingManagerBin = ERC20LicensedStakingManagerMetaData.Bin

// DeployERC20LicensedStakingManager deploys a new Ethereum contract, binding an instance of ERC20LicensedStakingManager to it.
func DeployERC20LicensedStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *ERC20LicensedStakingManager, error) {
	parsed, err := ERC20LicensedStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20LicensedStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20LicensedStakingManager{ERC20LicensedStakingManagerCaller: ERC20LicensedStakingManagerCaller{contract: contract}, ERC20LicensedStakingManagerTransactor: ERC20LicensedStakingManagerTransactor{contract: contract}, ERC20LicensedStakingManagerFilterer: ERC20LicensedStakingManagerFilterer{contract: contract}}, nil
}

// ERC20LicensedStakingManager is an auto generated Go binding around an Ethereum contract.
type ERC20LicensedStakingManager struct {
	ERC20LicensedStakingManagerCaller     // Read-only binding to the contract
	ERC20LicensedStakingManagerTransactor // Write-only binding to the contract
	ERC20LicensedStakingManagerFilterer   // Log filterer for contract events
}

// ERC20LicensedStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20LicensedStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20LicensedStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20LicensedStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20LicensedStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20LicensedStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20LicensedStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20LicensedStakingManagerSession struct {
	Contract     *ERC20LicensedStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC20LicensedStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20LicensedStakingManagerCallerSession struct {
	Contract *ERC20LicensedStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// ERC20LicensedStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20LicensedStakingManagerTransactorSession struct {
	Contract     *ERC20LicensedStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// ERC20LicensedStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20LicensedStakingManagerRaw struct {
	Contract *ERC20LicensedStakingManager // Generic contract binding to access the raw methods on
}

// ERC20LicensedStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20LicensedStakingManagerCallerRaw struct {
	Contract *ERC20LicensedStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20LicensedStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20LicensedStakingManagerTransactorRaw struct {
	Contract *ERC20LicensedStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20LicensedStakingManager creates a new instance of ERC20LicensedStakingManager, bound to a specific deployed contract.
func NewERC20LicensedStakingManager(address common.Address, backend bind.ContractBackend) (*ERC20LicensedStakingManager, error) {
	contract, err := bindERC20LicensedStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManager{ERC20LicensedStakingManagerCaller: ERC20LicensedStakingManagerCaller{contract: contract}, ERC20LicensedStakingManagerTransactor: ERC20LicensedStakingManagerTransactor{contract: contract}, ERC20LicensedStakingManagerFilterer: ERC20LicensedStakingManagerFilterer{contract: contract}}, nil
}

// NewERC20LicensedStakingManagerCaller creates a new read-only instance of ERC20LicensedStakingManager, bound to a specific deployed contract.
func NewERC20LicensedStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*ERC20LicensedStakingManagerCaller, error) {
	contract, err := bindERC20LicensedStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerCaller{contract: contract}, nil
}

// NewERC20LicensedStakingManagerTransactor creates a new write-only instance of ERC20LicensedStakingManager, bound to a specific deployed contract.
func NewERC20LicensedStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20LicensedStakingManagerTransactor, error) {
	contract, err := bindERC20LicensedStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerTransactor{contract: contract}, nil
}

// NewERC20LicensedStakingManagerFilterer creates a new log filterer instance of ERC20LicensedStakingManager, bound to a specific deployed contract.
func NewERC20LicensedStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20LicensedStakingManagerFilterer, error) {
	contract, err := bindERC20LicensedStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerFilterer{contract: contract}, nil
}

// bindERC20LicensedStakingManager binds a generic wrapper to an already deployed contract.
func bindERC20LicensedStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20LicensedStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20LicensedStakingManager.Contract.ERC20LicensedStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ERC20LicensedStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ERC20LicensedStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20LicensedStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.contract.Transact(opts, method, params...)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20LicensedStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20LicensedStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _ERC20LicensedStakingManager.Contract.BIPSCONVERSIONFACTOR(&_ERC20LicensedStakingManager.CallOpts)
}

// ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xb82f87df.
//
// Solidity: function ERC20_LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "ERC20_LICENSED_STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xb82f87df.
//
// Solidity: function ERC20_LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION(&_ERC20LicensedStakingManager.CallOpts)
}

// ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0xb82f87df.
//
// Solidity: function ERC20_LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.ERC20LICENSEDSTAKINGMANAGERSTORAGELOCATION(&_ERC20LicensedStakingManager.CallOpts)
}

// LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x61df835f.
//
// Solidity: function LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) LICENSEDSTAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "LICENSED_STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x61df835f.
//
// Solidity: function LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) LICENSEDSTAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.LICENSEDSTAKINGMANAGERSTORAGELOCATION(&_ERC20LicensedStakingManager.CallOpts)
}

// LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x61df835f.
//
// Solidity: function LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) LICENSEDSTAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.LICENSEDSTAKINGMANAGERSTORAGELOCATION(&_ERC20LicensedStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20LicensedStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20LicensedStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _ERC20LicensedStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_ERC20LicensedStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20LicensedStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20LicensedStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _ERC20LicensedStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_ERC20LicensedStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_ERC20LicensedStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_ERC20LicensedStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20LicensedStakingManager.Contract.WARPMESSENGER(&_ERC20LicensedStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _ERC20LicensedStakingManager.Contract.WARPMESSENGER(&_ERC20LicensedStakingManager.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) Erc20(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "erc20")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) Erc20() (common.Address, error) {
	return _ERC20LicensedStakingManager.Contract.Erc20(&_ERC20LicensedStakingManager.CallOpts)
}

// Erc20 is a free data retrieval call binding the contract method 0x785e9e86.
//
// Solidity: function erc20() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) Erc20() (common.Address, error) {
	return _ERC20LicensedStakingManager.Contract.Erc20(&_ERC20LicensedStakingManager.CallOpts)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) Erc721(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "erc721")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) Erc721() (common.Address, error) {
	return _ERC20LicensedStakingManager.Contract.Erc721(&_ERC20LicensedStakingManager.CallOpts)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) Erc721() (common.Address, error) {
	return _ERC20LicensedStakingManager.Contract.Erc721(&_ERC20LicensedStakingManager.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetDelegatorInfo(opts *bind.CallOpts, delegationID [32]byte) (Delegator, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getDelegatorInfo", delegationID)

	if err != nil {
		return *new(Delegator), err
	}

	out0 := *abi.ConvertType(out[0], new(Delegator)).(*Delegator)

	return out0, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _ERC20LicensedStakingManager.Contract.GetDelegatorInfo(&_ERC20LicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _ERC20LicensedStakingManager.Contract.GetDelegatorInfo(&_ERC20LicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetDelegatorRewardInfo(opts *bind.CallOpts, delegationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getDelegatorRewardInfo", delegationID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetDelegatorRewardInfo(&_ERC20LicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetDelegatorRewardInfo(&_ERC20LicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetDelegatorStakedLicenseTokens(opts *bind.CallOpts, delegationID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getDelegatorStakedLicenseTokens", delegationID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetDelegatorStakedLicenseTokens(delegationID [32]byte) ([]*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetDelegatorStakedLicenseTokens(&_ERC20LicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetDelegatorStakedLicenseTokens(delegationID [32]byte) ([]*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetDelegatorStakedLicenseTokens(&_ERC20LicensedStakingManager.CallOpts, delegationID)
}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetLicenseTokenDelegator(opts *bind.CallOpts, licenseTokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getLicenseTokenDelegator", licenseTokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetLicenseTokenDelegator(licenseTokenId *big.Int) ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.GetLicenseTokenDelegator(&_ERC20LicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetLicenseTokenDelegator(licenseTokenId *big.Int) ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.GetLicenseTokenDelegator(&_ERC20LicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetLicenseTokenValidator(opts *bind.CallOpts, licenseTokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getLicenseTokenValidator", licenseTokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetLicenseTokenValidator(licenseTokenId *big.Int) ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.GetLicenseTokenValidator(&_ERC20LicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetLicenseTokenValidator(licenseTokenId *big.Int) ([32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.GetLicenseTokenValidator(&_ERC20LicensedStakingManager.CallOpts, licenseTokenId)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetStakingManagerSettings(opts *bind.CallOpts) (StakingManagerSettings, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getStakingManagerSettings")

	if err != nil {
		return *new(StakingManagerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingManagerSettings)).(*StakingManagerSettings)

	return out0, err

}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _ERC20LicensedStakingManager.Contract.GetStakingManagerSettings(&_ERC20LicensedStakingManager.CallOpts)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _ERC20LicensedStakingManager.Contract.GetStakingManagerSettings(&_ERC20LicensedStakingManager.CallOpts)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetStakingValidator(opts *bind.CallOpts, validationID [32]byte) (PoSValidatorInfo, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getStakingValidator", validationID)

	if err != nil {
		return *new(PoSValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PoSValidatorInfo)).(*PoSValidatorInfo)

	return out0, err

}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _ERC20LicensedStakingManager.Contract.GetStakingValidator(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _ERC20LicensedStakingManager.Contract.GetStakingValidator(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetValidatorDelegations(opts *bind.CallOpts, validationID [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getValidatorDelegations", validationID)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetValidatorDelegations(validationID [32]byte) ([][32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.GetValidatorDelegations(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetValidatorDelegations(validationID [32]byte) ([][32]byte, error) {
	return _ERC20LicensedStakingManager.Contract.GetValidatorDelegations(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetValidatorRewardInfo(opts *bind.CallOpts, validationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getValidatorRewardInfo", validationID)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetValidatorRewardInfo(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetValidatorRewardInfo(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) GetValidatorStakedLicenseTokens(opts *bind.CallOpts, validationID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "getValidatorStakedLicenseTokens", validationID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) GetValidatorStakedLicenseTokens(validationID [32]byte) ([]*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetValidatorStakedLicenseTokens(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) GetValidatorStakedLicenseTokens(validationID [32]byte) ([]*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.GetValidatorStakedLicenseTokens(&_ERC20LicensedStakingManager.CallOpts, validationID)
}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) LicenseToStakeConversionFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "licenseToStakeConversionFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) LicenseToStakeConversionFactor() (*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.LicenseToStakeConversionFactor(&_ERC20LicensedStakingManager.CallOpts)
}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) LicenseToStakeConversionFactor() (*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.LicenseToStakeConversionFactor(&_ERC20LicensedStakingManager.CallOpts)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20LicensedStakingManager.Contract.ValueToWeight(&_ERC20LicensedStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _ERC20LicensedStakingManager.Contract.ValueToWeight(&_ERC20LicensedStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _ERC20LicensedStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.WeightToValue(&_ERC20LicensedStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _ERC20LicensedStakingManager.Contract.WeightToValue(&_ERC20LicensedStakingManager.CallOpts, weight)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20LicensedStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ChangeDelegatorRewardRecipient(&_ERC20LicensedStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20LicensedStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ChangeValidatorRewardRecipient(&_ERC20LicensedStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ClaimDelegationFees(&_ERC20LicensedStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ClaimDelegationFees(&_ERC20LicensedStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteDelegatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) CompleteDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "completeDelegatorRemoval", delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteDelegatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteDelegatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteValidatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteValidatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteValidatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.CompleteValidatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) ForceInitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "forceInitiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ForceInitiateDelegatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ForceInitiateDelegatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) ForceInitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "forceInitiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ForceInitiateValidatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ForceInitiateValidatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0x401e58ac.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token, address licenseToken, uint256 licenseToStakeConversionFactor) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings StakingManagerSettings, token common.Address, licenseToken common.Address, licenseToStakeConversionFactor *big.Int) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "initialize", settings, token, licenseToken, licenseToStakeConversionFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x401e58ac.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token, address licenseToken, uint256 licenseToStakeConversionFactor) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) Initialize(settings StakingManagerSettings, token common.Address, licenseToken common.Address, licenseToStakeConversionFactor *big.Int) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.Initialize(&_ERC20LicensedStakingManager.TransactOpts, settings, token, licenseToken, licenseToStakeConversionFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0x401e58ac.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address token, address licenseToken, uint256 licenseToStakeConversionFactor) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) Initialize(settings StakingManagerSettings, token common.Address, licenseToken common.Address, licenseToStakeConversionFactor *big.Int) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.Initialize(&_ERC20LicensedStakingManager.TransactOpts, settings, token, licenseToken, licenseToStakeConversionFactor)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, uint256[] licenseTokenIds, address rewardRecipient) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) InitiateDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, delegationAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "initiateDelegatorRegistration", validationID, delegationAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, uint256[] licenseTokenIds, address rewardRecipient) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateDelegatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, validationID, delegationAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, uint256[] licenseTokenIds, address rewardRecipient) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateDelegatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, validationID, delegationAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) InitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "initiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateDelegatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateDelegatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateValidatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) returns(bytes32)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateValidatorRegistration(&_ERC20LicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "initiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateValidatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.InitiateValidatorRemoval(&_ERC20LicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) ResendUpdateDelegator(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "resendUpdateDelegator", delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ResendUpdateDelegator(&_ERC20LicensedStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.ResendUpdateDelegator(&_ERC20LicensedStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.SubmitUptimeProof(&_ERC20LicensedStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _ERC20LicensedStakingManager.Contract.SubmitUptimeProof(&_ERC20LicensedStakingManager.TransactOpts, validationID, messageIndex)
}

// ERC20LicensedStakingManagerCompletedDelegatorRegistrationIterator is returned from FilterCompletedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRegistration events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerCompletedDelegatorRegistrationIterator struct {
	Event *ERC20LicensedStakingManagerCompletedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerCompletedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerCompletedDelegatorRegistration)
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
		it.Event = new(ERC20LicensedStakingManagerCompletedDelegatorRegistration)
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
func (it *ERC20LicensedStakingManagerCompletedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerCompletedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerCompletedDelegatorRegistration represents a CompletedDelegatorRegistration event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerCompletedDelegatorRegistration struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRegistration is a free log retrieval operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterCompletedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20LicensedStakingManagerCompletedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerCompletedDelegatorRegistrationIterator{contract: _ERC20LicensedStakingManager.contract, event: "CompletedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRegistration is a free log subscription operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchCompletedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerCompletedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerCompletedDelegatorRegistration)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseCompletedDelegatorRegistration(log types.Log) (*ERC20LicensedStakingManagerCompletedDelegatorRegistration, error) {
	event := new(ERC20LicensedStakingManagerCompletedDelegatorRegistration)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerCompletedDelegatorRemovalIterator is returned from FilterCompletedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRemoval events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerCompletedDelegatorRemovalIterator struct {
	Event *ERC20LicensedStakingManagerCompletedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerCompletedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerCompletedDelegatorRemoval)
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
		it.Event = new(ERC20LicensedStakingManagerCompletedDelegatorRemoval)
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
func (it *ERC20LicensedStakingManagerCompletedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerCompletedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerCompletedDelegatorRemoval represents a CompletedDelegatorRemoval event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerCompletedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterCompletedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20LicensedStakingManagerCompletedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerCompletedDelegatorRemovalIterator{contract: _ERC20LicensedStakingManager.contract, event: "CompletedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRemoval is a free log subscription operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchCompletedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerCompletedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerCompletedDelegatorRemoval)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseCompletedDelegatorRemoval(log types.Log) (*ERC20LicensedStakingManagerCompletedDelegatorRemoval, error) {
	event := new(ERC20LicensedStakingManagerCompletedDelegatorRemoval)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerDelegatorRegisteredWithLicensesIterator is returned from FilterDelegatorRegisteredWithLicenses and is used to iterate over the raw logs and unpacked data for DelegatorRegisteredWithLicenses events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerDelegatorRegisteredWithLicensesIterator struct {
	Event *ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses)
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
		it.Event = new(ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses)
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
func (it *ERC20LicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses represents a DelegatorRegisteredWithLicenses event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses struct {
	DelegationID    [32]byte
	ValidationID    [32]byte
	LicenseTokenIds []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegisteredWithLicenses is a free log retrieval operation binding the contract event 0xc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb569.
//
// Solidity: event DelegatorRegisteredWithLicenses(bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterDelegatorRegisteredWithLicenses(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20LicensedStakingManagerDelegatorRegisteredWithLicensesIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "DelegatorRegisteredWithLicenses", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerDelegatorRegisteredWithLicensesIterator{contract: _ERC20LicensedStakingManager.contract, event: "DelegatorRegisteredWithLicenses", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegisteredWithLicenses is a free log subscription operation binding the contract event 0xc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb569.
//
// Solidity: event DelegatorRegisteredWithLicenses(bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchDelegatorRegisteredWithLicenses(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "DelegatorRegisteredWithLicenses", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "DelegatorRegisteredWithLicenses", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseDelegatorRegisteredWithLicenses(log types.Log) (*ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses, error) {
	event := new(ERC20LicensedStakingManagerDelegatorRegisteredWithLicenses)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "DelegatorRegisteredWithLicenses", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerDelegatorRewardClaimedIterator is returned from FilterDelegatorRewardClaimed and is used to iterate over the raw logs and unpacked data for DelegatorRewardClaimed events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerDelegatorRewardClaimedIterator struct {
	Event *ERC20LicensedStakingManagerDelegatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerDelegatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerDelegatorRewardClaimed)
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
		it.Event = new(ERC20LicensedStakingManagerDelegatorRewardClaimed)
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
func (it *ERC20LicensedStakingManagerDelegatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerDelegatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerDelegatorRewardClaimed represents a DelegatorRewardClaimed event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerDelegatorRewardClaimed struct {
	DelegationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardClaimed is a free log retrieval operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterDelegatorRewardClaimed(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address) (*ERC20LicensedStakingManagerDelegatorRewardClaimedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerDelegatorRewardClaimedIterator{contract: _ERC20LicensedStakingManager.contract, event: "DelegatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardClaimed is a free log subscription operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchDelegatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerDelegatorRewardClaimed, delegationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerDelegatorRewardClaimed)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseDelegatorRewardClaimed(log types.Log) (*ERC20LicensedStakingManagerDelegatorRewardClaimed, error) {
	event := new(ERC20LicensedStakingManagerDelegatorRewardClaimed)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerDelegatorRewardRecipientChangedIterator is returned from FilterDelegatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for DelegatorRewardRecipientChanged events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerDelegatorRewardRecipientChangedIterator struct {
	Event *ERC20LicensedStakingManagerDelegatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerDelegatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerDelegatorRewardRecipientChanged)
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
		it.Event = new(ERC20LicensedStakingManagerDelegatorRewardRecipientChanged)
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
func (it *ERC20LicensedStakingManagerDelegatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerDelegatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerDelegatorRewardRecipientChanged represents a DelegatorRewardRecipientChanged event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerDelegatorRewardRecipientChanged struct {
	DelegationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterDelegatorRewardRecipientChanged(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*ERC20LicensedStakingManagerDelegatorRewardRecipientChangedIterator, error) {

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

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerDelegatorRewardRecipientChangedIterator{contract: _ERC20LicensedStakingManager.contract, event: "DelegatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchDelegatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerDelegatorRewardRecipientChanged, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerDelegatorRewardRecipientChanged)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseDelegatorRewardRecipientChanged(log types.Log) (*ERC20LicensedStakingManagerDelegatorRewardRecipientChanged, error) {
	event := new(ERC20LicensedStakingManagerDelegatorRewardRecipientChanged)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerERC20TokensUnlockedIterator is returned from FilterERC20TokensUnlocked and is used to iterate over the raw logs and unpacked data for ERC20TokensUnlocked events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerERC20TokensUnlockedIterator struct {
	Event *ERC20LicensedStakingManagerERC20TokensUnlocked // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerERC20TokensUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerERC20TokensUnlocked)
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
		it.Event = new(ERC20LicensedStakingManagerERC20TokensUnlocked)
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
func (it *ERC20LicensedStakingManagerERC20TokensUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerERC20TokensUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerERC20TokensUnlocked represents a ERC20TokensUnlocked event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerERC20TokensUnlocked struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterERC20TokensUnlocked is a free log retrieval operation binding the contract event 0xb96e58ee5d4ea44c7f5b0ec141397ea1d47f915b0d70d6b2355a6ced826ce349.
//
// Solidity: event ERC20TokensUnlocked(address indexed to, uint256 amount)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterERC20TokensUnlocked(opts *bind.FilterOpts, to []common.Address) (*ERC20LicensedStakingManagerERC20TokensUnlockedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "ERC20TokensUnlocked", toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerERC20TokensUnlockedIterator{contract: _ERC20LicensedStakingManager.contract, event: "ERC20TokensUnlocked", logs: logs, sub: sub}, nil
}

// WatchERC20TokensUnlocked is a free log subscription operation binding the contract event 0xb96e58ee5d4ea44c7f5b0ec141397ea1d47f915b0d70d6b2355a6ced826ce349.
//
// Solidity: event ERC20TokensUnlocked(address indexed to, uint256 amount)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchERC20TokensUnlocked(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerERC20TokensUnlocked, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "ERC20TokensUnlocked", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerERC20TokensUnlocked)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ERC20TokensUnlocked", log); err != nil {
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

// ParseERC20TokensUnlocked is a log parse operation binding the contract event 0xb96e58ee5d4ea44c7f5b0ec141397ea1d47f915b0d70d6b2355a6ced826ce349.
//
// Solidity: event ERC20TokensUnlocked(address indexed to, uint256 amount)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseERC20TokensUnlocked(log types.Log) (*ERC20LicensedStakingManagerERC20TokensUnlocked, error) {
	event := new(ERC20LicensedStakingManagerERC20TokensUnlocked)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ERC20TokensUnlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitializedIterator struct {
	Event *ERC20LicensedStakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerInitialized)
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
		it.Event = new(ERC20LicensedStakingManagerInitialized)
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
func (it *ERC20LicensedStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerInitialized represents a Initialized event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC20LicensedStakingManagerInitializedIterator, error) {

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerInitializedIterator{contract: _ERC20LicensedStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerInitialized)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseInitialized(log types.Log) (*ERC20LicensedStakingManagerInitialized, error) {
	event := new(ERC20LicensedStakingManagerInitialized)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerInitiatedDelegatorRegistrationIterator is returned from FilterInitiatedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRegistration events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitiatedDelegatorRegistrationIterator struct {
	Event *ERC20LicensedStakingManagerInitiatedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerInitiatedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerInitiatedDelegatorRegistration)
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
		it.Event = new(ERC20LicensedStakingManagerInitiatedDelegatorRegistration)
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
func (it *ERC20LicensedStakingManagerInitiatedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerInitiatedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerInitiatedDelegatorRegistration represents a InitiatedDelegatorRegistration event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitiatedDelegatorRegistration struct {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterInitiatedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*ERC20LicensedStakingManagerInitiatedDelegatorRegistrationIterator, error) {

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

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerInitiatedDelegatorRegistrationIterator{contract: _ERC20LicensedStakingManager.contract, event: "InitiatedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRegistration is a free log subscription operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchInitiatedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerInitiatedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerInitiatedDelegatorRegistration)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseInitiatedDelegatorRegistration(log types.Log) (*ERC20LicensedStakingManagerInitiatedDelegatorRegistration, error) {
	event := new(ERC20LicensedStakingManagerInitiatedDelegatorRegistration)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerInitiatedDelegatorRemovalIterator is returned from FilterInitiatedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRemoval events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitiatedDelegatorRemovalIterator struct {
	Event *ERC20LicensedStakingManagerInitiatedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerInitiatedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerInitiatedDelegatorRemoval)
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
		it.Event = new(ERC20LicensedStakingManagerInitiatedDelegatorRemoval)
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
func (it *ERC20LicensedStakingManagerInitiatedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerInitiatedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerInitiatedDelegatorRemoval represents a InitiatedDelegatorRemoval event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitiatedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterInitiatedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*ERC20LicensedStakingManagerInitiatedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerInitiatedDelegatorRemovalIterator{contract: _ERC20LicensedStakingManager.contract, event: "InitiatedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRemoval is a free log subscription operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchInitiatedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerInitiatedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerInitiatedDelegatorRemoval)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseInitiatedDelegatorRemoval(log types.Log) (*ERC20LicensedStakingManagerInitiatedDelegatorRemoval, error) {
	event := new(ERC20LicensedStakingManagerInitiatedDelegatorRemoval)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerInitiatedStakingValidatorRegistrationIterator is returned from FilterInitiatedStakingValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedStakingValidatorRegistration events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitiatedStakingValidatorRegistrationIterator struct {
	Event *ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration)
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
		it.Event = new(ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration)
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
func (it *ERC20LicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration represents a InitiatedStakingValidatorRegistration event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration struct {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterInitiatedStakingValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, owner []common.Address) (*ERC20LicensedStakingManagerInitiatedStakingValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerInitiatedStakingValidatorRegistrationIterator{contract: _ERC20LicensedStakingManager.contract, event: "InitiatedStakingValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedStakingValidatorRegistration is a free log subscription operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchInitiatedStakingValidatorRegistration(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration, validationID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseInitiatedStakingValidatorRegistration(log types.Log) (*ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration, error) {
	event := new(ERC20LicensedStakingManagerInitiatedStakingValidatorRegistration)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerUptimeUpdatedIterator struct {
	Event *ERC20LicensedStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerUptimeUpdated)
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
		it.Event = new(ERC20LicensedStakingManagerUptimeUpdated)
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
func (it *ERC20LicensedStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20LicensedStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerUptimeUpdatedIterator{contract: _ERC20LicensedStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerUptimeUpdated)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*ERC20LicensedStakingManagerUptimeUpdated, error) {
	event := new(ERC20LicensedStakingManagerUptimeUpdated)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerValidatorRegisteredWithLicensesIterator is returned from FilterValidatorRegisteredWithLicenses and is used to iterate over the raw logs and unpacked data for ValidatorRegisteredWithLicenses events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerValidatorRegisteredWithLicensesIterator struct {
	Event *ERC20LicensedStakingManagerValidatorRegisteredWithLicenses // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerValidatorRegisteredWithLicensesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerValidatorRegisteredWithLicenses)
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
		it.Event = new(ERC20LicensedStakingManagerValidatorRegisteredWithLicenses)
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
func (it *ERC20LicensedStakingManagerValidatorRegisteredWithLicensesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerValidatorRegisteredWithLicensesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerValidatorRegisteredWithLicenses represents a ValidatorRegisteredWithLicenses event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerValidatorRegisteredWithLicenses struct {
	ValidationID    [32]byte
	LicenseTokenIds []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegisteredWithLicenses is a free log retrieval operation binding the contract event 0x0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968.
//
// Solidity: event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterValidatorRegisteredWithLicenses(opts *bind.FilterOpts, validationID [][32]byte) (*ERC20LicensedStakingManagerValidatorRegisteredWithLicensesIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "ValidatorRegisteredWithLicenses", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerValidatorRegisteredWithLicensesIterator{contract: _ERC20LicensedStakingManager.contract, event: "ValidatorRegisteredWithLicenses", logs: logs, sub: sub}, nil
}

// WatchValidatorRegisteredWithLicenses is a free log subscription operation binding the contract event 0x0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968.
//
// Solidity: event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchValidatorRegisteredWithLicenses(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerValidatorRegisteredWithLicenses, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "ValidatorRegisteredWithLicenses", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerValidatorRegisteredWithLicenses)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ValidatorRegisteredWithLicenses", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseValidatorRegisteredWithLicenses(log types.Log) (*ERC20LicensedStakingManagerValidatorRegisteredWithLicenses, error) {
	event := new(ERC20LicensedStakingManagerValidatorRegisteredWithLicenses)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ValidatorRegisteredWithLicenses", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerValidatorRewardClaimedIterator is returned from FilterValidatorRewardClaimed and is used to iterate over the raw logs and unpacked data for ValidatorRewardClaimed events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerValidatorRewardClaimedIterator struct {
	Event *ERC20LicensedStakingManagerValidatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerValidatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerValidatorRewardClaimed)
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
		it.Event = new(ERC20LicensedStakingManagerValidatorRewardClaimed)
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
func (it *ERC20LicensedStakingManagerValidatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerValidatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerValidatorRewardClaimed represents a ValidatorRewardClaimed event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerValidatorRewardClaimed struct {
	ValidationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardClaimed is a free log retrieval operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterValidatorRewardClaimed(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address) (*ERC20LicensedStakingManagerValidatorRewardClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerValidatorRewardClaimedIterator{contract: _ERC20LicensedStakingManager.contract, event: "ValidatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardClaimed is a free log subscription operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchValidatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerValidatorRewardClaimed, validationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerValidatorRewardClaimed)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseValidatorRewardClaimed(log types.Log) (*ERC20LicensedStakingManagerValidatorRewardClaimed, error) {
	event := new(ERC20LicensedStakingManagerValidatorRewardClaimed)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20LicensedStakingManagerValidatorRewardRecipientChangedIterator is returned from FilterValidatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for ValidatorRewardRecipientChanged events raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerValidatorRewardRecipientChangedIterator struct {
	Event *ERC20LicensedStakingManagerValidatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *ERC20LicensedStakingManagerValidatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20LicensedStakingManagerValidatorRewardRecipientChanged)
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
		it.Event = new(ERC20LicensedStakingManagerValidatorRewardRecipientChanged)
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
func (it *ERC20LicensedStakingManagerValidatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20LicensedStakingManagerValidatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20LicensedStakingManagerValidatorRewardRecipientChanged represents a ValidatorRewardRecipientChanged event raised by the ERC20LicensedStakingManager contract.
type ERC20LicensedStakingManagerValidatorRewardRecipientChanged struct {
	ValidationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) FilterValidatorRewardRecipientChanged(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*ERC20LicensedStakingManagerValidatorRewardRecipientChangedIterator, error) {

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

	logs, sub, err := _ERC20LicensedStakingManager.contract.FilterLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ERC20LicensedStakingManagerValidatorRewardRecipientChangedIterator{contract: _ERC20LicensedStakingManager.contract, event: "ValidatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) WatchValidatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *ERC20LicensedStakingManagerValidatorRewardRecipientChanged, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _ERC20LicensedStakingManager.contract.WatchLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20LicensedStakingManagerValidatorRewardRecipientChanged)
				if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
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
func (_ERC20LicensedStakingManager *ERC20LicensedStakingManagerFilterer) ParseValidatorRewardRecipientChanged(log types.Log) (*ERC20LicensedStakingManagerValidatorRewardRecipientChanged, error) {
	event := new(ERC20LicensedStakingManagerValidatorRewardRecipientChanged)
	if err := _ERC20LicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
