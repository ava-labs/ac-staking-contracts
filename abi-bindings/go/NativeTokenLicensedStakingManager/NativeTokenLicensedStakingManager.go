// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nativetokenlicensedstakingmanager

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

// NativeTokenLicensedStakingManagerMetaData contains all meta data concerning the NativeTokenLicensedStakingManager contract.
var NativeTokenLicensedStakingManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BIPS_CONVERSION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LICENSED_STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeDelegatorRewardRecipient\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeValidatorRewardRecipient\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRemoval\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"erc721\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceInitiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDelegatorInfo\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDelegator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startingNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endingNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorRewardInfo\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenDelegator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenValidator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingManagerSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakingManagerSettings\",\"components\":[{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"contractIValidatorManager\"},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"},{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorInfo\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"uptimeSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDelegations\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorRewardInfo\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structStakingManagerSettings\",\"components\":[{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"contractIValidatorManager\"},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"},{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"licenseToken\",\"type\":\"address\",\"internalType\":\"contractIERC721\"},{\"name\":\"licenseToStakeConversionFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"delegationAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRegistration\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"licenseToStakeConversionFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegator\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CompletedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CompletedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardClaimed\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedStakingValidatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unlocked\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRewardRecipient\",\"inputs\":[{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenCount\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidUptimeBlockchainID\",\"inputs\":[{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokenAlreadyStaked\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"TokenNotStaked\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedValidationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWeightToValueFactor\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051614e34380380614e3483398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b614ce78061014d5f395ff3fe608060405260043610610212575f3560e01c80636b1470c01161011e578063af1dd66c116100a8578063c52efd6d1161006d578063c52efd6d146106fe578063caa7187414610711578063e24b268014610732578063f9c4f06414610751578063fb8b11dd1461079b575f80fd5b8063af1dd66c1461063f578063b2c1712e1461067d578063b771b3bc1461069c578063bca6ce64146106b6578063bdddf02f146106df575f80fd5b806393e24598116100ee57806393e24598146105c35780639681d940146105e2578063a0c4b0b514610601578063a3a65e4814610620578063a9778a7a14610423575f80fd5b80636b1470c0146105085780637a63ad85146105525780637b88033a146105855780638ef34c98146105a4575f80fd5b80632aa566381161019f578063383caa731161016f578063383caa731461044b57806353a133381461047e57806360ad7784146104aa57806361df835f146104c957806362065856146104e9575f80fd5b80632aa566381461039b5780632e2194d8146103ba578063329c3e12146103f157806335455ded14610423575f80fd5b80631c39211f116101e55780631c39211f146102b1578063245dafcb146102d257806325e1c776146102f15780632674874b1461031057806327bf60cd1461037c575f80fd5b80630f2cb59714610216578063134096451461024b578063151d30d11461026c5780631667956414610292575b5f80fd5b348015610221575f80fd5b50610235610230366004614062565b6107ba565b6040516102429190614079565b60405180910390f35b348015610256575f80fd5b5061026a6102653660046140cf565b610838565b005b348015610277575f80fd5b50610280600a81565b60405160ff9091168152602001610242565b34801561029d575f80fd5b5061026a6102ac366004614106565b610b78565b6102c46102bf3660046143a7565b610b89565b604051908152602001610242565b3480156102dd575f80fd5b5061026a6102ec366004614062565b610bca565b3480156102fc575f80fd5b5061026a61030b3660046140cf565b610e8e565b34801561031b575f80fd5b5061032f61032a366004614062565b610f6c565b6040805182516001600160a01b0316815260208084015161ffff1690820152828201516001600160401b039081169282019290925260609283015190911691810191909152608001610242565b348015610387575f80fd5b5061026a610396366004614106565b610ff8565b3480156103a6575f80fd5b5061026a6103b5366004614106565b611003565b3480156103c5575f80fd5b506103d96103d4366004614062565b611013565b6040516001600160401b039091168152602001610242565b3480156103fc575f80fd5b5061040b6001600160991b0181565b6040516001600160a01b039091168152602001610242565b34801561042e575f80fd5b5061043861271081565b60405161ffff9091168152602001610242565b348015610456575f80fd5b507f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb01546102c4565b348015610489575f80fd5b5061049d610498366004614062565b611067565b60405161024291906144d8565b3480156104b5575f80fd5b5061026a6104c43660046140cf565b611154565b3480156104d4575f80fd5b506102c45f80516020614c7283398151915281565b3480156104f4575f80fd5b506102c4610503366004614547565b611476565b348015610513575f80fd5b506102c4610522366004614062565b5f9081527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb05602052604090205490565b34801561055d575f80fd5b506102c47fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60081565b348015610590575f80fd5b5061023561059f366004614062565b611496565b3480156105af575f80fd5b5061026a6105be366004614562565b611512565b3480156105ce575f80fd5b5061026a6105dd366004614062565b6115f4565b3480156105ed575f80fd5b506102c46105fc366004614590565b61170d565b34801561060c575f80fd5b5061026a61061b3660046145a9565b6118af565b34801561062b575f80fd5b506102c461063a366004614590565b6119bf565b34801561064a575f80fd5b5061065e610659366004614062565b611a37565b604080516001600160a01b039093168352602083019190915201610242565b348015610688575f80fd5b5061026a610697366004614106565b611a73565b3480156106a7575f80fd5b5061040b6005600160991b0181565b3480156106c1575f80fd5b505f80516020614c72833981519152546001600160a01b031661040b565b3480156106ea575f80fd5b506102356106f9366004614062565b611a7e565b6102c461070c3660046145ef565b611ae7565b34801561071c575f80fd5b50610725611b1e565b6040516102429190614654565b34801561073d575f80fd5b5061065e61074c366004614062565b611bfa565b34801561075c575f80fd5b506102c461076b366004614062565b5f9081527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb04602052604090205490565b3480156107a6575f80fd5b5061026a6107b5366004614562565b611c36565b5f8181527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb03602090815260409182902080548351818402810184019094528084526060939283018282801561082c57602002820191905f5260205f20905b815481526020019060010190808311610818575b50505050509050919050565b610840611cfe565b5f610849611d35565b5f848152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115610882576108826144b0565b6003811115610893576108936144b0565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c090910152905060038151600381111561090c5761090c6144b0565b14610936578051604051633b0d540d60e21b815261092d91906004016146f5565b60405180910390fd5b81546040828101519051636af907fb60e11b815260048101919091525f916001600160a01b03169063d5f20ff6906024015f60405180830381865afa158015610981573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109a89190810190614769565b9050600483546040808501519051636af907fb60e11b81526001600160a01b039092169163d5f20ff6916109e29160040190815260200190565b5f60405180830381865afa1580156109fc573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610a239190810190614769565b516005811115610a3557610a356144b0565b14158015610a5c57508160c001516001600160401b031681608001516001600160401b0316105b15610b5257825460405163338587c560e21b815263ffffffff861660048201525f9182916001600160a01b039091169063ce161f149060240160408051808303815f875af1158015610ab0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ad49190614848565b9150915081846040015114610b0d5781846040015160405163fee3144560e01b815260040161092d929190918252602082015260400190565b806001600160401b03168460c001516001600160401b03161115610b4f57604051632e19bc2d60e11b81526001600160401b038216600482015260240161092d565b50505b610b5b85611d59565b505050610b7460015f80516020614c9283398151915255565b5050565b610b83838383611fa3565b50505050565b5f610b92611cfe565b610ba48b8b8b8b8b8b8b8b8b8b61224a565b9050610bbc60015f80516020614c9283398151915255565b9a9950505050505050505050565b5f610bd3611d35565b5f838152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115610c0c57610c0c6144b0565b6003811115610c1d57610c1d6144b0565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115610c9657610c966144b0565b14158015610cb75750600381516003811115610cb457610cb46144b0565b14155b15610cd8578051604051633b0d540d60e21b815261092d91906004016146f5565b81546040828101519051636af907fb60e11b815260048101919091525f916001600160a01b03169063d5f20ff6906024015f60405180830381865afa158015610d23573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610d4a9190810190614769565b905080606001516001600160401b03165f03610d7c576040516339b894f960e21b81526004810185905260240161092d565b604080830151606083015160a0840151925163854a893f60e01b81526005600160991b019363ee5b48eb9373__$caf7a49f90ce02b15fb10b3b072d8b9489$__9363854a893f93610dea93906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015610e04573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610e2b919081019061486b565b6040518263ffffffff1660e01b8152600401610e4791906148d2565b6020604051808303815f875af1158015610e63573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e8791906148e4565b5050505050565b610e9782612325565b610eb7576040516330efa98b60e01b81526004810183905260240161092d565b5f610ec0611d35565b54604051636af907fb60e11b8152600481018590526001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015610f04573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610f2b9190810190614769565b5190506002816005811115610f4257610f426144b0565b14610f62578060405163170cc93360e21b815260040161092d91906148fb565b610b83838361234e565b604080516080810182525f808252602082018190529181018290526060810191909152610f97611d35565b5f9283526007016020908152604092839020835160808101855281546001600160a01b038116825261ffff600160a01b820416938201939093526001600160401b03600160b01b909304831694810194909452600101541660608301525090565b610b838383836125b8565b61100e8383836129d9565b505050565b5f8061101d611d35565b6004015461102b9084614929565b905080158061104057506001600160401b0381115b156110615760405163222d164360e21b81526004810184905260240161092d565b92915050565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526110a7611d35565b5f838152600891909101602052604090819020815160e081019092528054829060ff1660038111156110db576110db6144b0565b60038111156110ec576110ec6144b0565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015292915050565b5f61115d611d35565b5f848152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115611196576111966144b0565b60038111156111a7576111a76144b0565b8152815461010090046001600160a01b03908116602083015260018301546040808401919091526002909301546001600160401b038082166060850152600160401b820481166080850152600160801b8204811660a0850152600160c01b9091041660c0909201919091528282015185549251636af907fb60e11b815260048101829052939450925f929091169063d5f20ff6906024015f60405180830381865afa158015611258573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261127f9190810190614769565b9050600183516003811115611296576112966144b0565b146112b7578251604051633b0d540d60e21b815261092d91906004016146f5565b6004815160058111156112cc576112cc6144b0565b036112e2576112da86611d59565b505050505050565b8260a001516001600160401b031681608001516001600160401b031610156113ea57835460405163338587c560e21b815263ffffffff871660048201525f9182916001600160a01b039091169063ce161f149060240160408051808303815f875af1158015611353573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113779190614848565b915091508184146113a55760405163fee3144560e01b8152600481018390526024810185905260440161092d565b8460a001516001600160401b0316816001600160401b031610156113e757604051632e19bc2d60e11b81526001600160401b038216600482015260240161092d565b50505b5f868152600885016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b81026fffffffffffffffff000000000000000019909216919091179091559151918252839188917f3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa910160405180910390a3505050505050565b5f61147f611d35565b60040154611061906001600160401b038416614948565b5f8181527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb06602090815260409182902080548351818402810184019094528084526060939283018282801561082c57602002820191905f5260205f20908154815260200190600101908083116108185750505050509050919050565b5f61151b611d35565b90506001600160a01b03821661154f5760405163caa903f960e01b81526001600160a01b038316600482015260240161092d565b5f8381526007820160205260409020546001600160a01b0316331461159557335b604051636e2ccd7560e11b81526001600160a01b03909116600482015260240161092d565b5f838152600c8201602052604080822080546001600160a01b038681166001600160a01b0319831681179093559251921692839287917f28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c3391a450505050565b5f6115fd611d35565b8054604051636af907fb60e11b8152600481018590529192505f916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015611647573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261166e9190810190614769565b5190506004816005811115611685576116856144b0565b146116a5578060405163170cc93360e21b815260040161092d91906148fb565b5f8381526007830160205260409020546001600160a01b031633146116ca5733611570565b5f838152600c830160205260409020546001600160a01b03168061170357505f8381526007830160205260409020546001600160a01b03165b610b838185612a04565b5f611716611cfe565b5f61171f611d35565b805460405163025a076560e61b815263ffffffff861660048201529192505f916001600160a01b0390911690639681d940906024016020604051808303815f875af1158015611770573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061179491906148e4565b8254604051636af907fb60e11b8152600481018390529192505f916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa1580156117de573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526118059190810190614769565b905061181082612325565b61181e575091506118949050565b5f828152600784016020908152604080832054600c8701909252909120546001600160a01b039182169116806118515750805b600483516005811115611866576118666144b0565b03611875576118758185612a04565b61188c826118868560400151611476565b86612a78565b509193505050505b6118aa60015f80516020614c9283398151915255565b919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156118f35750825b90505f826001600160401b0316600114801561190e5750303b155b90508115801561191c575080155b1561193a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561196457845460ff60401b1916600160401b1785555b61196f888888612aec565b83156119b557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f6119c8611d35565b54604051631474cbc960e31b815263ffffffff841660048201526001600160a01b039091169063a3a65e48906024016020604051808303815f875af1158015611a13573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061106191906148e4565b5f805f611a42611d35565b5f948552600a810160209081526040808720546009909301909152909420546001600160a01b039094169492505050565b61100e838383612b07565b5f8181525f80516020614c52833981519152602090815260409182902080548351818402810184019094528084526060939283018282801561082c57602002820191905f5260205f20908154815260200190600101908083116108185750505050509050919050565b5f611af0611cfe565b611afd8686868686612b32565b9050611b1560015f80516020614c9283398151915255565b95945050505050565b60408051610120810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905290611b6e611d35565b604080516101208101825282546001600160a01b0390811682526001840154602083015260028401549282019290925260038301546001600160401b0381166060830152600160401b810461ffff166080830152600160501b900460ff1660a0820152600483015460c0820152600583015490911660e082015260069091015461010082015292915050565b5f805f611c05611d35565b5f948552600c81016020908152604080872054600b909301909152909420546001600160a01b039094169492505050565b6001600160a01b038116611c685760405163caa903f960e01b81526001600160a01b038216600482015260240161092d565b5f611c71611d35565b5f8481526008820160205260409020549091506001600160a01b03610100909104163314611c9f5733611570565b5f838152600a8201602052604080822080546001600160a01b038681166001600160a01b0319831681179093559251921692839287917f6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e00591a450505050565b5f80516020614c92833981519152805460011901611d2f57604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b7fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60090565b5f611d62611d35565b5f838152600882016020526040808220815160e0810190925280549394509192909190829060ff166003811115611d9b57611d9b6144b0565b6003811115611dac57611dac6144b0565b815281546001600160a01b03610100909104811660208084019190915260018401546040808501919091526002909401546001600160401b038082166060860152600160401b820481166080860152600160801b8204811660a0860152600160c01b9091041660c09093019290925283830151865484516304e0efb360e11b8152945195965090949116926309c1df669260048083019391928290030181865afa158015611e5c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e80919061495f565b8260800151611e8f919061497a565b6001600160401b0316421015611ec35760405163fb6ce63f60e01b81526001600160401b034216600482015260240161092d565b5f848152600884016020908152604080832080546001600160a81b031916815560018101849055600201839055600a8601909152902080546001600160a01b031981169091556001600160a01b031680611f1e575060208201515b5f80611f2b838886612c04565b91509150611f4a8560200151611f448760600151611476565b89612a78565b6040805183815260208101839052859189917f5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e910160405180910390a350505050505050565b60015f80516020614c9283398151915255565b5f80611fad611d35565b8054604051635b73516560e11b8152600481018890529192506001600160a01b03169063b6e6a2ca906024015f604051808303815f87803b158015611ff0575f80fd5b505af1158015612002573d5f803e3d5ffd5b50508254604051636af907fb60e11b8152600481018990525f93506001600160a01b03909116915063d5f20ff6906024015f60405180830381865afa15801561204d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526120749190810190614769565b905061207f86612325565b61208e57600192505050612243565b5f8681526007830160205260409020546001600160a01b031633146120b35733611570565b5f86815260078301602052604090205460c08201516120e291600160b01b90046001600160401b03169061497a565b6001600160401b03168160e001516001600160401b031610156121295760e081015160405163fb6ce63f60e01b81526001600160401b03909116600482015260240161092d565b5f85156121415761213a878661234e565b905061215f565b505f8681526007830160205260409020600101546001600160401b03165b600583015460408301515f916001600160a01b031690634f22429f9061218490611476565b60c086015160e0808801516040519185901b6001600160e01b031916825260048201939093526001600160401b0391821660248201819052604482015291811660648301528516608482015260a401602060405180830381865afa1580156121ee573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061221291906148e4565b90508084600b015f8a81526020019081526020015f205f82825461223691906149a1565b9091555050151593505050505b9392505050565b5f82810361226e57604051637a91399f60e11b81526004810184905260240161092d565b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb01545f80516020614c72833981519152905f9087906122ad9087614948565b6122b791906149a1565b90505f6122ca8e8e8e8e8e8e888c612d03565b90506122da81888860015f61302f565b807f0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968888860405161230c9291906149b4565b60405180910390a29d9c50505050505050505050505050565b5f8061232f611d35565b5f938452600701602052505060409020546001600160a01b0316151590565b6040516306f8253560e41b815263ffffffff821660048201525f90819081906005600160991b0190636f825350906024015f60405180830381865afa158015612399573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526123c091908101906149eb565b91509150806123e257604051636b2f19e960e01b815260040160405180910390fd5b5f6123eb611d35565b6006810154845191925014612419578251604051636ba589a560e01b8152600481019190915260240161092d565b60208301516001600160a01b031615612455576020830151604051624de75d60e31b81526001600160a01b03909116600482015260240161092d565b5f8073__$caf7a49f90ce02b15fb10b3b072d8b9489$__63088c246386604001516040518263ffffffff1660e01b815260040161249291906148d2565b6040805180830381865af41580156124ac573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906124d09190614848565b915091508188146124fe5760405163fee3144560e01b8152600481018390526024810189905260440161092d565b5f8881526007840160205260409020600101546001600160401b03908116908216111561258f575f888152600784016020908152604091829020600101805467ffffffffffffffff19166001600160401b038516908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a26125ad565b505f8781526007830160205260409020600101546001600160401b03165b979650505050505050565b5f806125c2611d35565b5f868152600882016020526040808220815160e0810190925280549394509192909190829060ff1660038111156125fb576125fb6144b0565b600381111561260c5761260c6144b0565b8152815461010090046001600160a01b03908116602083015260018301546040808401919091526002909301546001600160401b038082166060850152600160401b820481166080850152600160801b8204811660a0850152600160c01b9091041660c0909201919091528282015185549251636af907fb60e11b815260048101829052939450925f929091169063d5f20ff6906024015f60405180830381865afa1580156126bd573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526126e49190810190614769565b90506002835160038111156126fb576126fb6144b0565b1461271c578251604051633b0d540d60e21b815261092d91906004016146f5565b60208301516001600160a01b031633146127b8575f8281526007850160205260409020546001600160a01b031633146127555733611570565b5f82815260078501602052604090205460c082015161278491600160b01b90046001600160401b03169061497a565b6001600160401b03164210156127b85760405163fb6ce63f60e01b81526001600160401b034216600482015260240161092d565b5f888152600a850160205260409020546001600160a01b03166002825160058111156127e6576127e66144b0565b036129805760038501546080850151612808916001600160401b03169061497a565b6001600160401b031642101561283c5760405163fb6ce63f60e01b81526001600160401b034216600482015260240161092d565b871561284e5761284c838861234e565b505b5f8981526008860160205260409020805460ff191660031790558454606085015160a08401516001600160a01b039092169163661096699186916128929190614a91565b6040516001600160e01b031960e085901b16815260048101929092526001600160401b0316602482015260440160408051808303815f875af11580156128da573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128fe9190614ab1565b505f8a8152600887016020526040812060020180546001600160401b03909316600160c01b026001600160c01b039093169290921790915561294185838c6131c8565b9050838a7f5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf60405160405180910390a315159550612243945050505050565b600482516005811115612995576129956144b0565b036129bd576129a584828b6131c8565b506129af89611d59565b600195505050505050612243565b815160405163170cc93360e21b815261092d91906004016148fb565b6129e48383836125b8565b61100e57604051631036cf9160e11b81526004810184905260240161092d565b5f612a0d611d35565b5f838152600b8201602052604081208054919055909150612a2e8482613402565b836001600160a01b0316837f875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e483604051612a6a91815260200190565b60405180910390a350505050565b5f612a838284613460565b9050612a908484846134d1565b612aa36001600160a01b03851682613518565b836001600160a01b03167f0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d82604051612ade91815260200190565b60405180910390a250505050565b612af46135ab565b612aff8383836135f6565b61100e613611565b612b12838383611fa3565b61100e57604051635bff683f60e11b81526004810184905260240161092d565b5f828103612b5657604051637a91399f60e11b81526004810184905260240161092d565b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb01545f80516020614c72833981519152905f908790612b959087614948565b612b9f91906149a1565b90505f612bae89338488613619565b9050612bbd8188885f8d61302f565b88817fc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb5698989604051612bf09291906149b4565b60405180910390a398975050505050505050565b5f805f612c0f611d35565b5f8681526009820160205260408120549192509081908015612cf5575f88815260098501602090815260408083208390558983526007870190915290205461271090612c6690600160a01b900461ffff1683614948565b612c709190614929565b91508184600b015f8981526020019081526020015f205f828254612c9491906149a1565b90915550612ca490508282614add565b9250612cb08984613402565b886001600160a01b0316887f3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf6385604051612cec91815260200190565b60405180910390a35b509097909650945050505050565b5f80612d0d611d35565b600381015490915061ffff600160401b90910481169087161080612d36575061271061ffff8716115b15612d5a57604051635f12e6c360e11b815261ffff8716600482015260240161092d565b60038101546001600160401b039081169086161015612d96576040516202a06d60e11b81526001600160401b038616600482015260240161092d565b8060010154841080612dab5750806002015484115b15612dcc5760405163222d164360e21b81526004810185905260240161092d565b6001600160a01b038316612dfe5760405163caa903f960e01b81526001600160a01b038416600482015260240161092d565b835f612e0982611013565b90505f835f015f9054906101000a90046001600160a01b03166001600160a01b0316639cb7624e8e8e8e8e876040518663ffffffff1660e01b8152600401612e55959493929190614b56565b6020604051808303815f875af1158015612e71573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e9591906148e4565b90505f33905080856007015f8481526020019081526020015f205f015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555089856007015f8481526020019081526020015f205f0160146101000a81548161ffff021916908361ffff16021790555088856007015f8481526020019081526020015f205f0160166101000a8154816001600160401b0302191690836001600160401b031602179055505f856007015f8481526020019081526020015f206001015f6101000a8154816001600160401b0302191690836001600160401b031602179055508685600c015f8481526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b0316827ff51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f8c8c8b6040516130169392919061ffff9390931683526001600160401b039190911660208301526001600160a01b0316604082015260600190565b60405180910390a3509c9b505050505050505050505050565b5f80516020614c728339815191525f5b848110156131bf575f86868381811061305a5761305a614bbe565b9050602002013590505f801b836004015f8381526020019081526020015f205414158061309557505f81815260058401602052604090205415155b156130b6576040516328f0990960e11b81526004810182905260240161092d565b82546001600160a01b03166323b872dd336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604481018490526064015f604051808303815f87803b158015613112575f80fd5b505af1158015613124573d5f803e3d5ffd5b505050508415613166575f88815260028401602090815260408083208054600181018255908452828420018490558383526004860190915290208890556131b6565b5f8881526003840160209081526040808320805460018181018355918552838520018590558484526005870183528184208c90558784526006870183529083208054918201815583529120018890555b5060010161303f565b50505050505050565b5f806131d2611d35565b80546040808801519051636af907fb60e11b81529293505f926001600160a01b039092169163d5f20ff69161320d9160040190815260200190565b5f60405180830381865afa158015613227573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261324e9190810190614769565b90505f600382516005811115613266576132666144b0565b14806132845750600482516005811115613282576132826144b0565b145b15613294575060e08101516132b1565b6002825160058111156132a9576132a96144b0565b036129bd5750425b86608001516001600160401b0316816001600160401b0316116132d9575f9350505050612243565b600583015460608801515f916001600160a01b031690634f22429f906132fe90611476565b60c086015160808c01516040808e01515f90815260078b0160205281902060010154905160e086901b6001600160e01b031916815260048101949094526001600160401b0392831660248501529082166044840152818716606484015216608482015260a401602060405180830381865afa15801561337f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133a391906148e4565b90506001600160a01b0387166133bb57876020015196505b5f8681526009850160209081526040808320849055600a90960190529390932080546001600160a01b0388166001600160a01b031990911617905550909150509392505050565b6040516327ad555d60e11b81526001600160a01b0383166004820152602481018290526001600160991b0190634f5aaaba906044015f604051808303815f87803b15801561344e575f80fd5b505af11580156112da573d5f803e3d5ffd5b5f8281525f80516020614c5283398151915260205260408120545f80516020614c72833981519152908290156134a657505f8481526002820160205260409020546134b8565b505f8481526003820160205260409020545b60018201546134c79082614948565b611b159085614add565b5f8181525f80516020614c5283398151915260205260409020545f80516020614c72833981519152901561350e576135098285613a60565b610b83565b610b838285613b49565b8047101561353b5760405163cd78605960e01b815230600482015260240161092d565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114613584576040519150601f19603f3d011682016040523d82523d5f602084013e613589565b606091505b505090508061100e57604051630a12f52160e11b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166135f457604051631afcd79f60e31b815260040160405180910390fd5b565b6135fe6135ab565b61360783613cc4565b61100e8282613d42565b6135f46135ab565b5f80613623611d35565b90505f61362f85611013565b905061363a87612325565b61365a576040516330efa98b60e01b81526004810188905260240161092d565b6001600160a01b03841661368c5760405163caa903f960e01b81526001600160a01b038516600482015260240161092d565b8154604051636af907fb60e11b8152600481018990525f9182916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa1580156136d5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526136fc9190810190614769565b9050828160a0015161370e919061497a565b915083600301600a9054906101000a90046001600160401b031681604001516137379190614bd2565b6001600160401b0316826001600160401b0316111561377457604051636d51fe0560e11b81526001600160401b038316600482015260240161092d565b508254604051636610966960e01b8152600481018a90526001600160401b03831660248201525f9182916001600160a01b039091169063661096699060440160408051808303815f875af11580156137ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137f29190614ab1565b915091505f8a8360405160200161382092919091825260c01b6001600160c01b031916602082015260280190565b60408051601f1981840301815291815281516020928301205f81815260088a019093529120805491925060019160ff19168280021790555089866008015f8381526020019081526020015f205f0160016101000a8154816001600160a01b0302191690836001600160a01b031602179055508a866008015f8381526020019081526020015f206001018190555084866008015f8381526020019081526020015f206002015f6101000a8154816001600160401b0302191690836001600160401b031602179055505f866008015f8381526020019081526020015f2060020160086101000a8154816001600160401b0302191690836001600160401b0316021790555082866008015f8381526020019081526020015f2060020160106101000a8154816001600160401b0302191690836001600160401b031602179055505f866008015f8381526020019081526020015f2060020160186101000a8154816001600160401b0302191690836001600160401b031602179055508786600a015f8381526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550896001600160a01b03168b827f77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d7786888a888f604051613a4a9594939291906001600160401b039586168152938516602085015291909316604083015260608201929092526001600160a01b0391909116608082015260a00190565b60405180910390a49a9950505050505050505050565b5f8281525f80516020614c52833981519152602052604081205f80516020614c72833981519152915b8154811015613b30575f828281548110613aa557613aa5614bbe565b5f9182526020909120015484546040516323b872dd60e01b81523060048201526001600160a01b038881166024830152604482018490529293509116906323b872dd906064015f604051808303815f87803b158015613b02575f80fd5b505af1158015613b14573d5f803e3d5ffd5b5050505f91825250600484016020526040812055600101613a89565b505f8481526002830160205260408120610b8391614034565b5f8281527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb03602052604081205f80516020614c72833981519152915b8154811015613c2c575f828281548110613ba157613ba1614bbe565b5f9182526020909120015484546040516323b872dd60e01b81523060048201526001600160a01b038881166024830152604482018490529293509116906323b872dd906064015f604051808303815f87803b158015613bfe575f80fd5b505af1158015613c10573d5f803e3d5ffd5b5050505f91825250600584016020526040812055600101613b85565b505f8481526003830160205260408120613c4591614034565b5f5b5f858152600684016020526040902054811015610e87575f8581526006840160205260409020805486919083908110613c8257613c82614bbe565b905f5260205f20015403613cbc575f8581526006840160205260409020805482908110613cb157613cb1614bbe565b5f9182526020822001555b600101613c47565b613ccc6135ab565b613cd4613da5565b613d3f613ce46020830183614bfd565b60208301356040840135613cfe6080860160608701614547565b613d0e60a0870160808801614c18565b613d1e60c0880160a08901614c31565b60c0880135613d346101008a0160e08b01614bfd565b896101000135613db5565b50565b613d4a6135ab565b5f80516020614c728339815191526001600160a01b038316613d7f5760405163d92e233d60e01b815260040160405180910390fd5b80546001600160a01b0319166001600160a01b0393909316929092178255600190910155565b613dad6135ab565b6135f461402c565b613dbd6135ab565b5f613dc6611d35565b905061ffff86161580613dde575061271061ffff8716115b15613e0257604051635f12e6c360e11b815261ffff8716600482015260240161092d565b87891115613e265760405163222d164360e21b8152600481018a905260240161092d565b60ff85161580613e395750600a60ff8616115b15613e5c5760405163170db35960e31b815260ff8616600482015260240161092d565b6001600160a01b038a16613e835760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038316613eaa5760405163d92e233d60e01b815260040160405180910390fd5b896001600160a01b03166309c1df666040518163ffffffff1660e01b8152600401602060405180830381865afa158015613ee6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613f0a919061495f565b6001600160401b0316876001600160401b03161015613f46576040516202a06d60e11b81526001600160401b038816600482015260240161092d565b835f03613f665760405163a733007160e01b815260040160405180910390fd5b81613f8757604051632f6bd1db60e01b81526004810183905260240161092d565b80546001600160a01b039a8b166001600160a01b031991821617825560018201999099556002810197909755600387018054600160501b60ff9096169590950267ffffffffffffffff60501b1961ffff909716600160401b0269ffffffffffffffffffff199096166001600160401b03909816979097179490941794909416949094179091556004840155600583018054929095169190931617909255600690910155565b611f906135ab565b5080545f8255905f5260205f2090810190613d3f91905b8082111561405e575f815560010161404b565b5090565b5f60208284031215614072575f80fd5b5035919050565b602080825282518282018190525f9190848201906040850190845b818110156140b057835183529284019291840191600101614094565b50909695505050505050565b803563ffffffff811681146118aa575f80fd5b5f80604083850312156140e0575f80fd5b823591506140f0602084016140bc565b90509250929050565b8015158114613d3f575f80fd5b5f805f60608486031215614118575f80fd5b83359250602084013561412a816140f9565b9150614138604085016140bc565b90509250925092565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b038111828210171561417757614177614141565b60405290565b60405161010081016001600160401b038111828210171561417757614177614141565b604051601f8201601f191681016001600160401b03811182821017156141c8576141c8614141565b604052919050565b5f6001600160401b038211156141e8576141e8614141565b50601f01601f191660200190565b5f82601f830112614205575f80fd5b8135614218614213826141d0565b6141a0565b81815284602083860101111561422c575f80fd5b816020850160208301375f918101602001919091529392505050565b6001600160a01b0381168114613d3f575f80fd5b80356118aa81614248565b5f60408284031215614277575f80fd5b61427f614155565b905061428a826140bc565b81526020808301356001600160401b03808211156142a6575f80fd5b818501915085601f8301126142b9575f80fd5b8135818111156142cb576142cb614141565b8060051b91506142dc8483016141a0565b81815291830184019184810190888411156142f5575f80fd5b938501935b8385101561431f578435925061430f83614248565b82825293850193908501906142fa565b808688015250505050505092915050565b803561ffff811681146118aa575f80fd5b6001600160401b0381168114613d3f575f80fd5b80356118aa81614341565b5f8083601f840112614370575f80fd5b5081356001600160401b03811115614386575f80fd5b6020830191508360208260051b85010111156143a0575f80fd5b9250929050565b5f805f805f805f805f806101208b8d0312156143c1575f80fd5b8a356001600160401b03808211156143d7575f80fd5b6143e38e838f016141f6565b9b5060208d01359150808211156143f8575f80fd5b6144048e838f016141f6565b9a5060408d0135915080821115614419575f80fd5b6144258e838f01614267565b995060608d013591508082111561443a575f80fd5b6144468e838f01614267565b985061445460808e01614330565b975061446260a08e01614355565b965060c08d0135955060e08d013591508082111561447e575f80fd5b5061448b8d828e01614360565b909450925061449f90506101008c0161425c565b90509295989b9194979a5092959850565b634e487b7160e01b5f52602160045260245ffd5b600481106144d4576144d46144b0565b9052565b5f60e0820190506144ea8284516144c4565b60018060a01b0360208401511660208301526040830151604083015260608301516001600160401b0380821660608501528060808601511660808501528060a08601511660a08501528060c08601511660c0850152505092915050565b5f60208284031215614557575f80fd5b813561224381614341565b5f8060408385031215614573575f80fd5b82359150602083013561458581614248565b809150509250929050565b5f602082840312156145a0575f80fd5b612243826140bc565b5f805f8385036101608112156145bd575f80fd5b610120808212156145cc575f80fd5b85945084013590506145dd81614248565b92959294505050610140919091013590565b5f805f805f60808688031215614603575f80fd5b853594506020860135935060408601356001600160401b03811115614626575f80fd5b61463288828901614360565b909450925050606086013561464681614248565b809150509295509295909350565b81516001600160a01b031681526020808301519082015260408083015190820152606080830151610120830191614695908401826001600160401b03169052565b5060808301516146ab608084018261ffff169052565b5060a08301516146c060a084018260ff169052565b5060c083015160c083015260e08301516146e560e08401826001600160a01b03169052565b5061010092830151919092015290565b6020810161106182846144c4565b8051600681106118aa575f80fd5b5f82601f830112614720575f80fd5b815161472e614213826141d0565b818152846020838601011115614742575f80fd5b8160208501602083015e5f918101602001919091529392505050565b80516118aa81614341565b5f60208284031215614779575f80fd5b81516001600160401b038082111561478f575f80fd5b9083019061010082860312156147a3575f80fd5b6147ab61417d565b6147b483614703565b81526020830151828111156147c7575f80fd5b6147d387828601614711565b6020830152506147e56040840161475e565b60408201526147f66060840161475e565b60608201526148076080840161475e565b608082015261481860a0840161475e565b60a082015261482960c0840161475e565b60c082015261483a60e0840161475e565b60e082015295945050505050565b5f8060408385031215614859575f80fd5b82519150602083015161458581614341565b5f6020828403121561487b575f80fd5b81516001600160401b03811115614890575f80fd5b61489c84828501614711565b949350505050565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f61224360208301846148a4565b5f602082840312156148f4575f80fd5b5051919050565b602081016006831061490f5761490f6144b0565b91905290565b634e487b7160e01b5f52601160045260245ffd5b5f8261494357634e487b7160e01b5f52601260045260245ffd5b500490565b808202811582820484141761106157611061614915565b5f6020828403121561496f575f80fd5b815161224381614341565b6001600160401b0381811683821601908082111561499a5761499a614915565b5092915050565b8082018082111561106157611061614915565b602080825281018290525f6001600160fb1b038311156149d2575f80fd5b8260051b80856040850137919091016040019392505050565b5f80604083850312156149fc575f80fd5b82516001600160401b0380821115614a12575f80fd5b9084019060608287031215614a25575f80fd5b604051606081018181108382111715614a4057614a40614141565b604052825181526020830151614a5581614248565b6020820152604083015182811115614a6b575f80fd5b614a7788828601614711565b6040830152508094505050506020830151614585816140f9565b6001600160401b0382811682821603908082111561499a5761499a614915565b5f8060408385031215614ac2575f80fd5b8251614acd81614341565b6020939093015192949293505050565b8181038181111561106157611061614915565b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b80831015614b4b5784516001600160a01b03168252938301936001929092019190830190614b22565b509695505050505050565b60a081525f614b6860a08301886148a4565b8281036020840152614b7a81886148a4565b90508281036040840152614b8e8187614af0565b90508281036060840152614ba28186614af0565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603260045260245ffd5b6001600160401b03818116838216028082169190828114614bf557614bf5614915565b505092915050565b5f60208284031215614c0d575f80fd5b813561224381614248565b5f60208284031215614c28575f80fd5b61224382614330565b5f60208284031215614c41575f80fd5b813560ff81168114612243575f80fdfe19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0219dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00a26469706673582212207c08bf3309c6cf940dd413a7461ea5f86d6a119324132fc27b8eefec838ed96864736f6c63430008190033",
}

// NativeTokenLicensedStakingManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeTokenLicensedStakingManagerMetaData.ABI instead.
var NativeTokenLicensedStakingManagerABI = NativeTokenLicensedStakingManagerMetaData.ABI

// NativeTokenLicensedStakingManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeTokenLicensedStakingManagerMetaData.Bin instead.
var NativeTokenLicensedStakingManagerBin = NativeTokenLicensedStakingManagerMetaData.Bin

// DeployNativeTokenLicensedStakingManager deploys a new Ethereum contract, binding an instance of NativeTokenLicensedStakingManager to it.
func DeployNativeTokenLicensedStakingManager(auth *bind.TransactOpts, backend bind.ContractBackend, init uint8) (common.Address, *types.Transaction, *NativeTokenLicensedStakingManager, error) {
	parsed, err := NativeTokenLicensedStakingManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeTokenLicensedStakingManagerBin), backend, init)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeTokenLicensedStakingManager{NativeTokenLicensedStakingManagerCaller: NativeTokenLicensedStakingManagerCaller{contract: contract}, NativeTokenLicensedStakingManagerTransactor: NativeTokenLicensedStakingManagerTransactor{contract: contract}, NativeTokenLicensedStakingManagerFilterer: NativeTokenLicensedStakingManagerFilterer{contract: contract}}, nil
}

// NativeTokenLicensedStakingManager is an auto generated Go binding around an Ethereum contract.
type NativeTokenLicensedStakingManager struct {
	NativeTokenLicensedStakingManagerCaller     // Read-only binding to the contract
	NativeTokenLicensedStakingManagerTransactor // Write-only binding to the contract
	NativeTokenLicensedStakingManagerFilterer   // Log filterer for contract events
}

// NativeTokenLicensedStakingManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeTokenLicensedStakingManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenLicensedStakingManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeTokenLicensedStakingManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenLicensedStakingManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeTokenLicensedStakingManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeTokenLicensedStakingManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeTokenLicensedStakingManagerSession struct {
	Contract     *NativeTokenLicensedStakingManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                      // Call options to use throughout this session
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// NativeTokenLicensedStakingManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeTokenLicensedStakingManagerCallerSession struct {
	Contract *NativeTokenLicensedStakingManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                            // Call options to use throughout this session
}

// NativeTokenLicensedStakingManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeTokenLicensedStakingManagerTransactorSession struct {
	Contract     *NativeTokenLicensedStakingManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                            // Transaction auth options to use throughout this session
}

// NativeTokenLicensedStakingManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeTokenLicensedStakingManagerRaw struct {
	Contract *NativeTokenLicensedStakingManager // Generic contract binding to access the raw methods on
}

// NativeTokenLicensedStakingManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeTokenLicensedStakingManagerCallerRaw struct {
	Contract *NativeTokenLicensedStakingManagerCaller // Generic read-only contract binding to access the raw methods on
}

// NativeTokenLicensedStakingManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeTokenLicensedStakingManagerTransactorRaw struct {
	Contract *NativeTokenLicensedStakingManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeTokenLicensedStakingManager creates a new instance of NativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewNativeTokenLicensedStakingManager(address common.Address, backend bind.ContractBackend) (*NativeTokenLicensedStakingManager, error) {
	contract, err := bindNativeTokenLicensedStakingManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManager{NativeTokenLicensedStakingManagerCaller: NativeTokenLicensedStakingManagerCaller{contract: contract}, NativeTokenLicensedStakingManagerTransactor: NativeTokenLicensedStakingManagerTransactor{contract: contract}, NativeTokenLicensedStakingManagerFilterer: NativeTokenLicensedStakingManagerFilterer{contract: contract}}, nil
}

// NewNativeTokenLicensedStakingManagerCaller creates a new read-only instance of NativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewNativeTokenLicensedStakingManagerCaller(address common.Address, caller bind.ContractCaller) (*NativeTokenLicensedStakingManagerCaller, error) {
	contract, err := bindNativeTokenLicensedStakingManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerCaller{contract: contract}, nil
}

// NewNativeTokenLicensedStakingManagerTransactor creates a new write-only instance of NativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewNativeTokenLicensedStakingManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeTokenLicensedStakingManagerTransactor, error) {
	contract, err := bindNativeTokenLicensedStakingManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerTransactor{contract: contract}, nil
}

// NewNativeTokenLicensedStakingManagerFilterer creates a new log filterer instance of NativeTokenLicensedStakingManager, bound to a specific deployed contract.
func NewNativeTokenLicensedStakingManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeTokenLicensedStakingManagerFilterer, error) {
	contract, err := bindNativeTokenLicensedStakingManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerFilterer{contract: contract}, nil
}

// bindNativeTokenLicensedStakingManager binds a generic wrapper to an already deployed contract.
func bindNativeTokenLicensedStakingManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NativeTokenLicensedStakingManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenLicensedStakingManager.Contract.NativeTokenLicensedStakingManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.NativeTokenLicensedStakingManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.NativeTokenLicensedStakingManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeTokenLicensedStakingManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.contract.Transact(opts, method, params...)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) BIPSCONVERSIONFACTOR(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "BIPS_CONVERSION_FACTOR")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenLicensedStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenLicensedStakingManager.CallOpts)
}

// BIPSCONVERSIONFACTOR is a free data retrieval call binding the contract method 0xa9778a7a.
//
// Solidity: function BIPS_CONVERSION_FACTOR() view returns(uint16)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) BIPSCONVERSIONFACTOR() (uint16, error) {
	return _NativeTokenLicensedStakingManager.Contract.BIPSCONVERSIONFACTOR(&_NativeTokenLicensedStakingManager.CallOpts)
}

// LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x61df835f.
//
// Solidity: function LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) LICENSEDSTAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "LICENSED_STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x61df835f.
//
// Solidity: function LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) LICENSEDSTAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.LICENSEDSTAKINGMANAGERSTORAGELOCATION(&_NativeTokenLicensedStakingManager.CallOpts)
}

// LICENSEDSTAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x61df835f.
//
// Solidity: function LICENSED_STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) LICENSEDSTAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.LICENSEDSTAKINGMANAGERSTORAGELOCATION(&_NativeTokenLicensedStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) MAXIMUMDELEGATIONFEEBIPS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "MAXIMUM_DELEGATION_FEE_BIPS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenLicensedStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenLicensedStakingManager.CallOpts)
}

// MAXIMUMDELEGATIONFEEBIPS is a free data retrieval call binding the contract method 0x35455ded.
//
// Solidity: function MAXIMUM_DELEGATION_FEE_BIPS() view returns(uint16)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) MAXIMUMDELEGATIONFEEBIPS() (uint16, error) {
	return _NativeTokenLicensedStakingManager.Contract.MAXIMUMDELEGATIONFEEBIPS(&_NativeTokenLicensedStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) MAXIMUMSTAKEMULTIPLIERLIMIT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "MAXIMUM_STAKE_MULTIPLIER_LIMIT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenLicensedStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenLicensedStakingManager.CallOpts)
}

// MAXIMUMSTAKEMULTIPLIERLIMIT is a free data retrieval call binding the contract method 0x151d30d1.
//
// Solidity: function MAXIMUM_STAKE_MULTIPLIER_LIMIT() view returns(uint8)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) MAXIMUMSTAKEMULTIPLIERLIMIT() (uint8, error) {
	return _NativeTokenLicensedStakingManager.Contract.MAXIMUMSTAKEMULTIPLIERLIMIT(&_NativeTokenLicensedStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) NATIVEMINTER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "NATIVE_MINTER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenLicensedStakingManager.Contract.NATIVEMINTER(&_NativeTokenLicensedStakingManager.CallOpts)
}

// NATIVEMINTER is a free data retrieval call binding the contract method 0x329c3e12.
//
// Solidity: function NATIVE_MINTER() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) NATIVEMINTER() (common.Address, error) {
	return _NativeTokenLicensedStakingManager.Contract.NATIVEMINTER(&_NativeTokenLicensedStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) STAKINGMANAGERSTORAGELOCATION(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "STAKING_MANAGER_STORAGE_LOCATION")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_NativeTokenLicensedStakingManager.CallOpts)
}

// STAKINGMANAGERSTORAGELOCATION is a free data retrieval call binding the contract method 0x7a63ad85.
//
// Solidity: function STAKING_MANAGER_STORAGE_LOCATION() view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) STAKINGMANAGERSTORAGELOCATION() ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.STAKINGMANAGERSTORAGELOCATION(&_NativeTokenLicensedStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) WARPMESSENGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "WARP_MESSENGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenLicensedStakingManager.Contract.WARPMESSENGER(&_NativeTokenLicensedStakingManager.CallOpts)
}

// WARPMESSENGER is a free data retrieval call binding the contract method 0xb771b3bc.
//
// Solidity: function WARP_MESSENGER() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) WARPMESSENGER() (common.Address, error) {
	return _NativeTokenLicensedStakingManager.Contract.WARPMESSENGER(&_NativeTokenLicensedStakingManager.CallOpts)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) Erc721(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "erc721")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) Erc721() (common.Address, error) {
	return _NativeTokenLicensedStakingManager.Contract.Erc721(&_NativeTokenLicensedStakingManager.CallOpts)
}

// Erc721 is a free data retrieval call binding the contract method 0xbca6ce64.
//
// Solidity: function erc721() view returns(address)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) Erc721() (common.Address, error) {
	return _NativeTokenLicensedStakingManager.Contract.Erc721(&_NativeTokenLicensedStakingManager.CallOpts)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetDelegatorInfo(opts *bind.CallOpts, delegationID [32]byte) (Delegator, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getDelegatorInfo", delegationID)

	if err != nil {
		return *new(Delegator), err
	}

	out0 := *abi.ConvertType(out[0], new(Delegator)).(*Delegator)

	return out0, err

}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetDelegatorInfo(&_NativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorInfo is a free data retrieval call binding the contract method 0x53a13338.
//
// Solidity: function getDelegatorInfo(bytes32 delegationID) view returns((uint8,address,bytes32,uint64,uint64,uint64,uint64))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetDelegatorInfo(delegationID [32]byte) (Delegator, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetDelegatorInfo(&_NativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetDelegatorRewardInfo(opts *bind.CallOpts, delegationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getDelegatorRewardInfo", delegationID)

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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetDelegatorRewardInfo(&_NativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorRewardInfo is a free data retrieval call binding the contract method 0xaf1dd66c.
//
// Solidity: function getDelegatorRewardInfo(bytes32 delegationID) view returns(address, uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetDelegatorRewardInfo(delegationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetDelegatorRewardInfo(&_NativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetDelegatorStakedLicenseTokens(opts *bind.CallOpts, delegationID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getDelegatorStakedLicenseTokens", delegationID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetDelegatorStakedLicenseTokens(delegationID [32]byte) ([]*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetDelegatorStakedLicenseTokens(&_NativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetDelegatorStakedLicenseTokens is a free data retrieval call binding the contract method 0x0f2cb597.
//
// Solidity: function getDelegatorStakedLicenseTokens(bytes32 delegationID) view returns(uint256[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetDelegatorStakedLicenseTokens(delegationID [32]byte) ([]*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetDelegatorStakedLicenseTokens(&_NativeTokenLicensedStakingManager.CallOpts, delegationID)
}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetLicenseTokenDelegator(opts *bind.CallOpts, licenseTokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getLicenseTokenDelegator", licenseTokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetLicenseTokenDelegator(licenseTokenId *big.Int) ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetLicenseTokenDelegator(&_NativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenDelegator is a free data retrieval call binding the contract method 0x6b1470c0.
//
// Solidity: function getLicenseTokenDelegator(uint256 licenseTokenId) view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetLicenseTokenDelegator(licenseTokenId *big.Int) ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetLicenseTokenDelegator(&_NativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetLicenseTokenValidator(opts *bind.CallOpts, licenseTokenId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getLicenseTokenValidator", licenseTokenId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetLicenseTokenValidator(licenseTokenId *big.Int) ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetLicenseTokenValidator(&_NativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetLicenseTokenValidator is a free data retrieval call binding the contract method 0xf9c4f064.
//
// Solidity: function getLicenseTokenValidator(uint256 licenseTokenId) view returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetLicenseTokenValidator(licenseTokenId *big.Int) ([32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetLicenseTokenValidator(&_NativeTokenLicensedStakingManager.CallOpts, licenseTokenId)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetStakingManagerSettings(opts *bind.CallOpts) (StakingManagerSettings, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getStakingManagerSettings")

	if err != nil {
		return *new(StakingManagerSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingManagerSettings)).(*StakingManagerSettings)

	return out0, err

}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetStakingManagerSettings(&_NativeTokenLicensedStakingManager.CallOpts)
}

// GetStakingManagerSettings is a free data retrieval call binding the contract method 0xcaa71874.
//
// Solidity: function getStakingManagerSettings() view returns((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetStakingManagerSettings() (StakingManagerSettings, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetStakingManagerSettings(&_NativeTokenLicensedStakingManager.CallOpts)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetStakingValidator(opts *bind.CallOpts, validationID [32]byte) (PoSValidatorInfo, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getStakingValidator", validationID)

	if err != nil {
		return *new(PoSValidatorInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PoSValidatorInfo)).(*PoSValidatorInfo)

	return out0, err

}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetStakingValidator(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetStakingValidator is a free data retrieval call binding the contract method 0x2674874b.
//
// Solidity: function getStakingValidator(bytes32 validationID) view returns((address,uint16,uint64,uint64))
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetStakingValidator(validationID [32]byte) (PoSValidatorInfo, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetStakingValidator(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetValidatorDelegations(opts *bind.CallOpts, validationID [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getValidatorDelegations", validationID)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetValidatorDelegations(validationID [32]byte) ([][32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetValidatorDelegations(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorDelegations is a free data retrieval call binding the contract method 0x7b88033a.
//
// Solidity: function getValidatorDelegations(bytes32 validationID) view returns(bytes32[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetValidatorDelegations(validationID [32]byte) ([][32]byte, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetValidatorDelegations(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetValidatorRewardInfo(opts *bind.CallOpts, validationID [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getValidatorRewardInfo", validationID)

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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetValidatorRewardInfo(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorRewardInfo is a free data retrieval call binding the contract method 0xe24b2680.
//
// Solidity: function getValidatorRewardInfo(bytes32 validationID) view returns(address, uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetValidatorRewardInfo(validationID [32]byte) (common.Address, *big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetValidatorRewardInfo(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) GetValidatorStakedLicenseTokens(opts *bind.CallOpts, validationID [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "getValidatorStakedLicenseTokens", validationID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) GetValidatorStakedLicenseTokens(validationID [32]byte) ([]*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetValidatorStakedLicenseTokens(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// GetValidatorStakedLicenseTokens is a free data retrieval call binding the contract method 0xbdddf02f.
//
// Solidity: function getValidatorStakedLicenseTokens(bytes32 validationID) view returns(uint256[])
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) GetValidatorStakedLicenseTokens(validationID [32]byte) ([]*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.GetValidatorStakedLicenseTokens(&_NativeTokenLicensedStakingManager.CallOpts, validationID)
}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) LicenseToStakeConversionFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "licenseToStakeConversionFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) LicenseToStakeConversionFactor() (*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.LicenseToStakeConversionFactor(&_NativeTokenLicensedStakingManager.CallOpts)
}

// LicenseToStakeConversionFactor is a free data retrieval call binding the contract method 0x383caa73.
//
// Solidity: function licenseToStakeConversionFactor() view returns(uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) LicenseToStakeConversionFactor() (*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.LicenseToStakeConversionFactor(&_NativeTokenLicensedStakingManager.CallOpts)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) ValueToWeight(opts *bind.CallOpts, value *big.Int) (uint64, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "valueToWeight", value)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenLicensedStakingManager.Contract.ValueToWeight(&_NativeTokenLicensedStakingManager.CallOpts, value)
}

// ValueToWeight is a free data retrieval call binding the contract method 0x2e2194d8.
//
// Solidity: function valueToWeight(uint256 value) view returns(uint64)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) ValueToWeight(value *big.Int) (uint64, error) {
	return _NativeTokenLicensedStakingManager.Contract.ValueToWeight(&_NativeTokenLicensedStakingManager.CallOpts, value)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCaller) WeightToValue(opts *bind.CallOpts, weight uint64) (*big.Int, error) {
	var out []interface{}
	err := _NativeTokenLicensedStakingManager.contract.Call(opts, &out, "weightToValue", weight)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.WeightToValue(&_NativeTokenLicensedStakingManager.CallOpts, weight)
}

// WeightToValue is a free data retrieval call binding the contract method 0x62065856.
//
// Solidity: function weightToValue(uint64 weight) view returns(uint256)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerCallerSession) WeightToValue(weight uint64) (*big.Int, error) {
	return _NativeTokenLicensedStakingManager.Contract.WeightToValue(&_NativeTokenLicensedStakingManager.CallOpts, weight)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) ChangeDelegatorRewardRecipient(opts *bind.TransactOpts, delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "changeDelegatorRewardRecipient", delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ChangeDelegatorRewardRecipient(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeDelegatorRewardRecipient is a paid mutator transaction binding the contract method 0xfb8b11dd.
//
// Solidity: function changeDelegatorRewardRecipient(bytes32 delegationID, address rewardRecipient) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) ChangeDelegatorRewardRecipient(delegationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ChangeDelegatorRewardRecipient(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) ChangeValidatorRewardRecipient(opts *bind.TransactOpts, validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "changeValidatorRewardRecipient", validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ChangeValidatorRewardRecipient(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ChangeValidatorRewardRecipient is a paid mutator transaction binding the contract method 0x8ef34c98.
//
// Solidity: function changeValidatorRewardRecipient(bytes32 validationID, address rewardRecipient) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) ChangeValidatorRewardRecipient(validationID [32]byte, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ChangeValidatorRewardRecipient(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, rewardRecipient)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) ClaimDelegationFees(opts *bind.TransactOpts, validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "claimDelegationFees", validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ClaimDelegationFees(&_NativeTokenLicensedStakingManager.TransactOpts, validationID)
}

// ClaimDelegationFees is a paid mutator transaction binding the contract method 0x93e24598.
//
// Solidity: function claimDelegationFees(bytes32 validationID) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) ClaimDelegationFees(validationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ClaimDelegationFees(&_NativeTokenLicensedStakingManager.TransactOpts, validationID)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) CompleteDelegatorRegistration(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "completeDelegatorRegistration", delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRegistration is a paid mutator transaction binding the contract method 0x60ad7784.
//
// Solidity: function completeDelegatorRegistration(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) CompleteDelegatorRegistration(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteDelegatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) CompleteDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "completeDelegatorRemoval", delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteDelegatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteDelegatorRemoval is a paid mutator transaction binding the contract method 0x13409645.
//
// Solidity: function completeDelegatorRemoval(bytes32 delegationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) CompleteDelegatorRemoval(delegationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteDelegatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) CompleteValidatorRegistration(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "completeValidatorRegistration", messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRegistration is a paid mutator transaction binding the contract method 0xa3a65e48.
//
// Solidity: function completeValidatorRegistration(uint32 messageIndex) returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) CompleteValidatorRegistration(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteValidatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) CompleteValidatorRemoval(opts *bind.TransactOpts, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "completeValidatorRemoval", messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteValidatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// CompleteValidatorRemoval is a paid mutator transaction binding the contract method 0x9681d940.
//
// Solidity: function completeValidatorRemoval(uint32 messageIndex) returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) CompleteValidatorRemoval(messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.CompleteValidatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) ForceInitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "forceInitiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ForceInitiateDelegatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x27bf60cd.
//
// Solidity: function forceInitiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) ForceInitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ForceInitiateDelegatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) ForceInitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "forceInitiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ForceInitiateValidatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ForceInitiateValidatorRemoval is a paid mutator transaction binding the contract method 0x16679564.
//
// Solidity: function forceInitiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) ForceInitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ForceInitiateValidatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0c4b0b5.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address licenseToken, uint256 licenseToStakeConversionFactor) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) Initialize(opts *bind.TransactOpts, settings StakingManagerSettings, licenseToken common.Address, licenseToStakeConversionFactor *big.Int) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "initialize", settings, licenseToken, licenseToStakeConversionFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0c4b0b5.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address licenseToken, uint256 licenseToStakeConversionFactor) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) Initialize(settings StakingManagerSettings, licenseToken common.Address, licenseToStakeConversionFactor *big.Int) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.Initialize(&_NativeTokenLicensedStakingManager.TransactOpts, settings, licenseToken, licenseToStakeConversionFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0xa0c4b0b5.
//
// Solidity: function initialize((address,uint256,uint256,uint64,uint16,uint8,uint256,address,bytes32) settings, address licenseToken, uint256 licenseToStakeConversionFactor) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) Initialize(settings StakingManagerSettings, licenseToken common.Address, licenseToStakeConversionFactor *big.Int) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.Initialize(&_NativeTokenLicensedStakingManager.TransactOpts, settings, licenseToken, licenseToStakeConversionFactor)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) InitiateDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, delegationAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "initiateDelegatorRegistration", validationID, delegationAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateDelegatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, delegationAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0xc52efd6d.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256 delegationAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) InitiateDelegatorRegistration(validationID [32]byte, delegationAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateDelegatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, delegationAmount, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) InitiateDelegatorRemoval(opts *bind.TransactOpts, delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "initiateDelegatorRemoval", delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateDelegatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateDelegatorRemoval is a paid mutator transaction binding the contract method 0x2aa56638.
//
// Solidity: function initiateDelegatorRemoval(bytes32 delegationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) InitiateDelegatorRemoval(delegationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateDelegatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateValidatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x1c39211f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256 stakeAmount, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, stakeAmount *big.Int, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateValidatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, stakeAmount, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) InitiateValidatorRemoval(opts *bind.TransactOpts, validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "initiateValidatorRemoval", validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateValidatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// InitiateValidatorRemoval is a paid mutator transaction binding the contract method 0xb2c1712e.
//
// Solidity: function initiateValidatorRemoval(bytes32 validationID, bool includeUptimeProof, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) InitiateValidatorRemoval(validationID [32]byte, includeUptimeProof bool, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateValidatorRemoval(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, includeUptimeProof, messageIndex)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) ResendUpdateDelegator(opts *bind.TransactOpts, delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "resendUpdateDelegator", delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ResendUpdateDelegator(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID)
}

// ResendUpdateDelegator is a paid mutator transaction binding the contract method 0x245dafcb.
//
// Solidity: function resendUpdateDelegator(bytes32 delegationID) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) ResendUpdateDelegator(delegationID [32]byte) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.ResendUpdateDelegator(&_NativeTokenLicensedStakingManager.TransactOpts, delegationID)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) SubmitUptimeProof(opts *bind.TransactOpts, validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "submitUptimeProof", validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.SubmitUptimeProof(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, messageIndex)
}

// SubmitUptimeProof is a paid mutator transaction binding the contract method 0x25e1c776.
//
// Solidity: function submitUptimeProof(bytes32 validationID, uint32 messageIndex) returns()
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) SubmitUptimeProof(validationID [32]byte, messageIndex uint32) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.SubmitUptimeProof(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, messageIndex)
}

// NativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator is returned from FilterCompletedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRegistration events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator struct {
	Event *NativeTokenLicensedStakingManagerCompletedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
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
		it.Event = new(NativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
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
func (it *NativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerCompletedDelegatorRegistration represents a CompletedDelegatorRegistration event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerCompletedDelegatorRegistration struct {
	DelegationID [32]byte
	ValidationID [32]byte
	StartTime    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRegistration is a free log retrieval operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterCompletedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerCompletedDelegatorRegistrationIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "CompletedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRegistration is a free log subscription operation binding the contract event 0x3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa.
//
// Solidity: event CompletedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 startTime)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchCompletedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerCompletedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRegistration", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseCompletedDelegatorRegistration(log types.Log) (*NativeTokenLicensedStakingManagerCompletedDelegatorRegistration, error) {
	event := new(NativeTokenLicensedStakingManagerCompletedDelegatorRegistration)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator is returned from FilterCompletedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for CompletedDelegatorRemoval events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator struct {
	Event *NativeTokenLicensedStakingManagerCompletedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
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
		it.Event = new(NativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
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
func (it *NativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerCompletedDelegatorRemoval represents a CompletedDelegatorRemoval event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerCompletedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Rewards      *big.Int
	Fees         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompletedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterCompletedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerCompletedDelegatorRemovalIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "CompletedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchCompletedDelegatorRemoval is a free log subscription operation binding the contract event 0x5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e.
//
// Solidity: event CompletedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID, uint256 rewards, uint256 fees)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchCompletedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerCompletedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "CompletedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseCompletedDelegatorRemoval(log types.Log) (*NativeTokenLicensedStakingManagerCompletedDelegatorRemoval, error) {
	event := new(NativeTokenLicensedStakingManagerCompletedDelegatorRemoval)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "CompletedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator is returned from FilterDelegatorRegisteredWithLicenses and is used to iterate over the raw logs and unpacked data for DelegatorRegisteredWithLicenses events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator struct {
	Event *NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
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
		it.Event = new(NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
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
func (it *NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses represents a DelegatorRegisteredWithLicenses event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses struct {
	DelegationID    [32]byte
	ValidationID    [32]byte
	LicenseTokenIds []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRegisteredWithLicenses is a free log retrieval operation binding the contract event 0xc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb569.
//
// Solidity: event DelegatorRegisteredWithLicenses(bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterDelegatorRegisteredWithLicenses(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "DelegatorRegisteredWithLicenses", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicensesIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "DelegatorRegisteredWithLicenses", logs: logs, sub: sub}, nil
}

// WatchDelegatorRegisteredWithLicenses is a free log subscription operation binding the contract event 0xc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb569.
//
// Solidity: event DelegatorRegisteredWithLicenses(bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchDelegatorRegisteredWithLicenses(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "DelegatorRegisteredWithLicenses", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRegisteredWithLicenses", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseDelegatorRegisteredWithLicenses(log types.Log) (*NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses, error) {
	event := new(NativeTokenLicensedStakingManagerDelegatorRegisteredWithLicenses)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRegisteredWithLicenses", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator is returned from FilterDelegatorRewardClaimed and is used to iterate over the raw logs and unpacked data for DelegatorRewardClaimed events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator struct {
	Event *NativeTokenLicensedStakingManagerDelegatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerDelegatorRewardClaimed)
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
		it.Event = new(NativeTokenLicensedStakingManagerDelegatorRewardClaimed)
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
func (it *NativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerDelegatorRewardClaimed represents a DelegatorRewardClaimed event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerDelegatorRewardClaimed struct {
	DelegationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardClaimed is a free log retrieval operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterDelegatorRewardClaimed(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address) (*NativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerDelegatorRewardClaimedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "DelegatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardClaimed is a free log subscription operation binding the contract event 0x3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63.
//
// Solidity: event DelegatorRewardClaimed(bytes32 indexed delegationID, address indexed recipient, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchDelegatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerDelegatorRewardClaimed, delegationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "DelegatorRewardClaimed", delegationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerDelegatorRewardClaimed)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseDelegatorRewardClaimed(log types.Log) (*NativeTokenLicensedStakingManagerDelegatorRewardClaimed, error) {
	event := new(NativeTokenLicensedStakingManagerDelegatorRewardClaimed)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator is returned from FilterDelegatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for DelegatorRewardRecipientChanged events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator struct {
	Event *NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
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
		it.Event = new(NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
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
func (it *NativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged represents a DelegatorRewardRecipientChanged event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged struct {
	DelegationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterDelegatorRewardRecipientChanged(opts *bind.FilterOpts, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*NativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator, error) {

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

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerDelegatorRewardRecipientChangedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "DelegatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchDelegatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e005.
//
// Solidity: event DelegatorRewardRecipientChanged(bytes32 indexed delegationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchDelegatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged, delegationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "DelegatorRewardRecipientChanged", delegationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseDelegatorRewardRecipientChanged(log types.Log) (*NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged, error) {
	event := new(NativeTokenLicensedStakingManagerDelegatorRewardRecipientChanged)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "DelegatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitializedIterator struct {
	Event *NativeTokenLicensedStakingManagerInitialized // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerInitialized)
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
		it.Event = new(NativeTokenLicensedStakingManagerInitialized)
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
func (it *NativeTokenLicensedStakingManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerInitialized represents a Initialized event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*NativeTokenLicensedStakingManagerInitializedIterator, error) {

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerInitializedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerInitialized)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseInitialized(log types.Log) (*NativeTokenLicensedStakingManagerInitialized, error) {
	event := new(NativeTokenLicensedStakingManagerInitialized)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator is returned from FilterInitiatedDelegatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRegistration events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator struct {
	Event *NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
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
		it.Event = new(NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
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
func (it *NativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration represents a InitiatedDelegatorRegistration event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration struct {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterInitiatedDelegatorRegistration(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (*NativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator, error) {

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

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerInitiatedDelegatorRegistrationIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "InitiatedDelegatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRegistration is a free log subscription operation binding the contract event 0x77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d77.
//
// Solidity: event InitiatedDelegatorRegistration(bytes32 indexed delegationID, bytes32 indexed validationID, address indexed delegatorAddress, uint64 nonce, uint64 validatorWeight, uint64 delegatorWeight, bytes32 setWeightMessageID, address rewardRecipient)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchInitiatedDelegatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration, delegationID [][32]byte, validationID [][32]byte, delegatorAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRegistration", delegationIDRule, validationIDRule, delegatorAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseInitiatedDelegatorRegistration(log types.Log) (*NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration, error) {
	event := new(NativeTokenLicensedStakingManagerInitiatedDelegatorRegistration)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator is returned from FilterInitiatedDelegatorRemoval and is used to iterate over the raw logs and unpacked data for InitiatedDelegatorRemoval events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator struct {
	Event *NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
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
		it.Event = new(NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
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
func (it *NativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval represents a InitiatedDelegatorRemoval event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval struct {
	DelegationID [32]byte
	ValidationID [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitiatedDelegatorRemoval is a free log retrieval operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterInitiatedDelegatorRemoval(opts *bind.FilterOpts, delegationID [][32]byte, validationID [][32]byte) (*NativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerInitiatedDelegatorRemovalIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "InitiatedDelegatorRemoval", logs: logs, sub: sub}, nil
}

// WatchInitiatedDelegatorRemoval is a free log subscription operation binding the contract event 0x5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf.
//
// Solidity: event InitiatedDelegatorRemoval(bytes32 indexed delegationID, bytes32 indexed validationID)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchInitiatedDelegatorRemoval(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval, delegationID [][32]byte, validationID [][32]byte) (event.Subscription, error) {

	var delegationIDRule []interface{}
	for _, delegationIDItem := range delegationID {
		delegationIDRule = append(delegationIDRule, delegationIDItem)
	}
	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "InitiatedDelegatorRemoval", delegationIDRule, validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseInitiatedDelegatorRemoval(log types.Log) (*NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval, error) {
	event := new(NativeTokenLicensedStakingManagerInitiatedDelegatorRemoval)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedDelegatorRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator is returned from FilterInitiatedStakingValidatorRegistration and is used to iterate over the raw logs and unpacked data for InitiatedStakingValidatorRegistration events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator struct {
	Event *NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
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
		it.Event = new(NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
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
func (it *NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration represents a InitiatedStakingValidatorRegistration event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration struct {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterInitiatedStakingValidatorRegistration(opts *bind.FilterOpts, validationID [][32]byte, owner []common.Address) (*NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistrationIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "InitiatedStakingValidatorRegistration", logs: logs, sub: sub}, nil
}

// WatchInitiatedStakingValidatorRegistration is a free log subscription operation binding the contract event 0xf51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f.
//
// Solidity: event InitiatedStakingValidatorRegistration(bytes32 indexed validationID, address indexed owner, uint16 delegationFeeBips, uint64 minStakeDuration, address rewardRecipient)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchInitiatedStakingValidatorRegistration(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration, validationID [][32]byte, owner []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "InitiatedStakingValidatorRegistration", validationIDRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseInitiatedStakingValidatorRegistration(log types.Log) (*NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration, error) {
	event := new(NativeTokenLicensedStakingManagerInitiatedStakingValidatorRegistration)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "InitiatedStakingValidatorRegistration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerUnlockedIterator is returned from FilterUnlocked and is used to iterate over the raw logs and unpacked data for Unlocked events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerUnlockedIterator struct {
	Event *NativeTokenLicensedStakingManagerUnlocked // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerUnlocked)
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
		it.Event = new(NativeTokenLicensedStakingManagerUnlocked)
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
func (it *NativeTokenLicensedStakingManagerUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerUnlocked represents a Unlocked event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerUnlocked struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlocked is a free log retrieval operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed to, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterUnlocked(opts *bind.FilterOpts, to []common.Address) (*NativeTokenLicensedStakingManagerUnlockedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "Unlocked", toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerUnlockedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "Unlocked", logs: logs, sub: sub}, nil
}

// WatchUnlocked is a free log subscription operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed to, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchUnlocked(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerUnlocked, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "Unlocked", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerUnlocked)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "Unlocked", log); err != nil {
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

// ParseUnlocked is a log parse operation binding the contract event 0x0f0bc5b519ddefdd8e5f9e6423433aa2b869738de2ae34d58ebc796fc749fa0d.
//
// Solidity: event Unlocked(address indexed to, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseUnlocked(log types.Log) (*NativeTokenLicensedStakingManagerUnlocked, error) {
	event := new(NativeTokenLicensedStakingManagerUnlocked)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "Unlocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerUptimeUpdatedIterator is returned from FilterUptimeUpdated and is used to iterate over the raw logs and unpacked data for UptimeUpdated events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerUptimeUpdatedIterator struct {
	Event *NativeTokenLicensedStakingManagerUptimeUpdated // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerUptimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerUptimeUpdated)
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
		it.Event = new(NativeTokenLicensedStakingManagerUptimeUpdated)
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
func (it *NativeTokenLicensedStakingManagerUptimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerUptimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerUptimeUpdated represents a UptimeUpdated event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerUptimeUpdated struct {
	ValidationID [32]byte
	Uptime       uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUptimeUpdated is a free log retrieval operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterUptimeUpdated(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenLicensedStakingManagerUptimeUpdatedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerUptimeUpdatedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "UptimeUpdated", logs: logs, sub: sub}, nil
}

// WatchUptimeUpdated is a free log subscription operation binding the contract event 0xec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435.
//
// Solidity: event UptimeUpdated(bytes32 indexed validationID, uint64 uptime)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchUptimeUpdated(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerUptimeUpdated, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "UptimeUpdated", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerUptimeUpdated)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseUptimeUpdated(log types.Log) (*NativeTokenLicensedStakingManagerUptimeUpdated, error) {
	event := new(NativeTokenLicensedStakingManagerUptimeUpdated)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "UptimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator is returned from FilterValidatorRegisteredWithLicenses and is used to iterate over the raw logs and unpacked data for ValidatorRegisteredWithLicenses events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator struct {
	Event *NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
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
		it.Event = new(NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
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
func (it *NativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses represents a ValidatorRegisteredWithLicenses event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses struct {
	ValidationID    [32]byte
	LicenseTokenIds []*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegisteredWithLicenses is a free log retrieval operation binding the contract event 0x0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968.
//
// Solidity: event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterValidatorRegisteredWithLicenses(opts *bind.FilterOpts, validationID [][32]byte) (*NativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "ValidatorRegisteredWithLicenses", validationIDRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerValidatorRegisteredWithLicensesIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "ValidatorRegisteredWithLicenses", logs: logs, sub: sub}, nil
}

// WatchValidatorRegisteredWithLicenses is a free log subscription operation binding the contract event 0x0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f6968.
//
// Solidity: event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchValidatorRegisteredWithLicenses(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses, validationID [][32]byte) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "ValidatorRegisteredWithLicenses", validationIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRegisteredWithLicenses", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseValidatorRegisteredWithLicenses(log types.Log) (*NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses, error) {
	event := new(NativeTokenLicensedStakingManagerValidatorRegisteredWithLicenses)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRegisteredWithLicenses", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerValidatorRewardClaimedIterator is returned from FilterValidatorRewardClaimed and is used to iterate over the raw logs and unpacked data for ValidatorRewardClaimed events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerValidatorRewardClaimedIterator struct {
	Event *NativeTokenLicensedStakingManagerValidatorRewardClaimed // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerValidatorRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerValidatorRewardClaimed)
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
		it.Event = new(NativeTokenLicensedStakingManagerValidatorRewardClaimed)
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
func (it *NativeTokenLicensedStakingManagerValidatorRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerValidatorRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerValidatorRewardClaimed represents a ValidatorRewardClaimed event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerValidatorRewardClaimed struct {
	ValidationID [32]byte
	Recipient    common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardClaimed is a free log retrieval operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterValidatorRewardClaimed(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address) (*NativeTokenLicensedStakingManagerValidatorRewardClaimedIterator, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerValidatorRewardClaimedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "ValidatorRewardClaimed", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardClaimed is a free log subscription operation binding the contract event 0x875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e4.
//
// Solidity: event ValidatorRewardClaimed(bytes32 indexed validationID, address indexed recipient, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchValidatorRewardClaimed(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerValidatorRewardClaimed, validationID [][32]byte, recipient []common.Address) (event.Subscription, error) {

	var validationIDRule []interface{}
	for _, validationIDItem := range validationID {
		validationIDRule = append(validationIDRule, validationIDItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "ValidatorRewardClaimed", validationIDRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerValidatorRewardClaimed)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseValidatorRewardClaimed(log types.Log) (*NativeTokenLicensedStakingManagerValidatorRewardClaimed, error) {
	event := new(NativeTokenLicensedStakingManagerValidatorRewardClaimed)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator is returned from FilterValidatorRewardRecipientChanged and is used to iterate over the raw logs and unpacked data for ValidatorRewardRecipientChanged events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator struct {
	Event *NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
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
		it.Event = new(NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
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
func (it *NativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged represents a ValidatorRewardRecipientChanged event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged struct {
	ValidationID [32]byte
	Recipient    common.Address
	OldRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorRewardRecipientChanged is a free log retrieval operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterValidatorRewardRecipientChanged(opts *bind.FilterOpts, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (*NativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator, error) {

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

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerValidatorRewardRecipientChangedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "ValidatorRewardRecipientChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorRewardRecipientChanged is a free log subscription operation binding the contract event 0x28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c33.
//
// Solidity: event ValidatorRewardRecipientChanged(bytes32 indexed validationID, address indexed recipient, address indexed oldRecipient)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchValidatorRewardRecipientChanged(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged, validationID [][32]byte, recipient []common.Address, oldRecipient []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "ValidatorRewardRecipientChanged", validationIDRule, recipientRule, oldRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
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
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseValidatorRewardRecipientChanged(log types.Log) (*NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged, error) {
	event := new(NativeTokenLicensedStakingManagerValidatorRewardRecipientChanged)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "ValidatorRewardRecipientChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
