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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"init\",\"type\":\"uint8\",\"internalType\":\"enumICMInitializable\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BIPS_CONVERSION_FACTOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"LICENSED_STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_DELEGATION_FEE_BIPS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAXIMUM_STAKE_MULTIPLIER_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NATIVE_MINTER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINativeMinter\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_MANAGER_STORAGE_LOCATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WARP_MESSENGER\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIWarpMessenger\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeDelegatorRewardRecipient\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeValidatorRewardRecipient\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimDelegationFees\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRegistration\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeValidatorRemoval\",\"inputs\":[{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"erc721\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC721\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forceInitiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceInitiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDelegatorInfo\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDelegator\",\"components\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"startingNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"endingNonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorRewardInfo\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDelegatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenDelegator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLicenseTokenValidator\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingManagerSettings\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structStakingManagerSettings\",\"components\":[{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"contractIValidatorManager\"},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"},{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakingValidator\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPoSValidatorInfo\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"uptimeSeconds\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorDelegations\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorRewardInfo\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getValidatorStakedLicenseTokens\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"settings\",\"type\":\"tuple\",\"internalType\":\"structStakingManagerSettings\",\"components\":[{\"name\":\"manager\",\"type\":\"address\",\"internalType\":\"contractIValidatorManager\"},{\"name\":\"minimumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maximumStakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minimumStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minimumDelegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"weightToValueFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"rewardCalculator\",\"type\":\"address\",\"internalType\":\"contractIRewardCalculator\"},{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"licenseToken\",\"type\":\"address\",\"internalType\":\"contractIERC721\"},{\"name\":\"licenseToStakeConversionFactor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initiateDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRegistration\",\"inputs\":[{\"name\":\"nodeID\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"remainingBalanceOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"disableOwner\",\"type\":\"tuple\",\"internalType\":\"structPChainOwner\",\"components\":[{\"name\":\"threshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"initiateValidatorRemoval\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"includeUptimeProof\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"licenseToStakeConversionFactor\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resendUpdateDelegator\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"submitUptimeProof\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messageIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valueToWeight\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"weightToValue\",\"inputs\":[{\"name\":\"weight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"CompletedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"startTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CompletedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"rewards\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardClaimed\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DelegatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRegistration\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"delegatorAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"validatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"delegatorWeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"setWeightMessageID\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedDelegatorRemoval\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InitiatedStakingValidatorRegistration\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"rewardRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeTokensUnlocked\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UptimeUpdated\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"uptime\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRegisteredWithLicenses\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"licenseTokenIds\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardClaimed\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorRewardRecipientChanged\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldRecipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"DelegatorIneligibleForRewards\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelegationFee\",\"inputs\":[{\"name\":\"delegationFeeBips\",\"type\":\"uint16\",\"internalType\":\"uint16\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegationID\",\"inputs\":[{\"name\":\"delegationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidDelegatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumDelegatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLicenseToStakeConversionFactor\",\"inputs\":[{\"name\":\"factor\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidLicenseTokenCount\",\"inputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidMinStakeDuration\",\"inputs\":[{\"name\":\"minStakeDuration\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"InvalidRewardRecipient\",\"inputs\":[{\"name\":\"rewardRecipient\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeAmount\",\"inputs\":[{\"name\":\"stakeAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidStakeMultiplier\",\"inputs\":[{\"name\":\"maximumStakeMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InvalidTokenStakeAmount\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InvalidUptimeBlockchainID\",\"inputs\":[{\"name\":\"uptimeBlockchainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidValidatorStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumValidatorStatus\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpMessage\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidWarpOriginSenderAddress\",\"inputs\":[{\"name\":\"senderAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidWarpSourceChainID\",\"inputs\":[{\"name\":\"sourceChainID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"LicenseTokenAlreadyStaked\",\"inputs\":[{\"name\":\"licenseTokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"MaxWeightExceeded\",\"inputs\":[{\"name\":\"newValidatorWeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"MinStakeDurationNotPassed\",\"inputs\":[{\"name\":\"endTime\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnauthorizedOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UnexpectedValidationID\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"expectedValidationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorIneligibleForRewards\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ValidatorNotPoS\",\"inputs\":[{\"name\":\"validationID\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroWeightToValueFactor\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516156cd3803806156cd83398101604081905261002e91610107565b60018160018111156100425761004261012c565b0361004f5761004f610055565b50610140565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a55760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101045780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b5f60208284031215610117575f80fd5b815160028110610125575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b6155808061014d5f395ff3fe60806040526004361061025d575f3560e01c80636b1470c01161014b578063a9778a7a116100c6578063bdddf02f1161007c578063e24b268011610062578063e24b268014610985578063f9c4f06414610a05578063fb8b11dd14610a4f575f80fd5b8063bdddf02f14610861578063caa7187414610880575f80fd5b8063b2c1712e116100ac578063b2c1712e146107df578063b771b3bc146107fe578063bca6ce6414610825575f80fd5b8063a9778a7a146104f4578063af1dd66c14610740575f80fd5b80638ef34c981161011b5780639681d940116101015780639681d940146106e3578063a0c4b0b514610702578063a3a65e4814610721575f80fd5b80638ef34c98146106a557806393e24598146106c4575f80fd5b80636b1470c0146106095780636c60f49f146106535780637a63ad85146106665780637b88033a14610686575f80fd5b80632e2194d8116101db5780633d2e15e5116101ab57806360ad77841161019157806360ad77841461059857806361df835f146105b757806362065856146105ea575f80fd5b80633d2e15e51461055957806353a133381461056c575f80fd5b80632e2194d81461047e578063329c3e12146104b557806335455ded146104f4578063383caa731461051c575f80fd5b8063245dafcb116102305780632674874b116102165780632674874b1461033a57806327bf60cd146104405780632aa566381461045f575f80fd5b8063245dafcb146102fc57806325e1c7761461031b575f80fd5b80630f2cb597146102615780631340964514610296578063151d30d1146102b757806316679564146102dd575b5f80fd5b34801561026c575f80fd5b5061028061027b366004614925565b610a6e565b60405161028d919061493c565b60405180910390f35b3480156102a1575f80fd5b506102b56102b0366004614992565b610aec565b005b3480156102c2575f80fd5b506102cb600a81565b60405160ff909116815260200161028d565b3480156102e8575f80fd5b506102b56102f73660046149c9565b610e5d565b348015610307575f80fd5b506102b5610316366004614925565b610e6e565b348015610326575f80fd5b506102b5610335366004614992565b61118f565b348015610345575f80fd5b506103f3610354366004614925565b60408051608080820183525f808352602080840182905283850182905260609384018290529481527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60785528390208351918201845280546001600160a01b0381168352600160a01b810461ffff1695830195909552600160b01b9094046001600160401b03908116938201939093526001909301549091169082015290565b6040805182516001600160a01b0316815260208084015161ffff1690820152828201516001600160401b03908116928201929092526060928301519091169181019190915260800161028d565b34801561044b575f80fd5b506102b561045a3660046149c9565b6112a0565b34801561046a575f80fd5b506102b56104793660046149c9565b6112ab565b348015610489575f80fd5b5061049d610498366004614925565b6112bb565b6040516001600160401b03909116815260200161028d565b3480156104c0575f80fd5b506104dc73020000000000000000000000000000000000000181565b6040516001600160a01b03909116815260200161028d565b3480156104ff575f80fd5b5061050961271081565b60405161ffff909116815260200161028d565b348015610527575f80fd5b507f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb01545b60405190815260200161028d565b61054b610567366004614a6a565b611315565b348015610577575f80fd5b5061058b610586366004614925565b61135e565b60405161028d9190614aec565b3480156105a3575f80fd5b506102b56105b2366004614992565b61145e565b3480156105c2575f80fd5b5061054b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0081565b3480156105f5575f80fd5b5061054b610604366004614b7b565b6117ad565b348015610614575f80fd5b5061054b610623366004614925565b5f9081527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb05602052604090205490565b61054b610661366004614bfa565b6117e4565b348015610671575f80fd5b5061054b5f8051602061552b83398151915281565b348015610691575f80fd5b506102806106a0366004614925565b6118c3565b3480156106b0575f80fd5b506102b56106bf366004614d05565b61193f565b3480156106cf575f80fd5b506102b56106de366004614925565b611a3d565b3480156106ee575f80fd5b5061054b6106fd366004614d33565b611b5c565b34801561070d575f80fd5b506102b561071c366004614d4c565b611d5f565b34801561072c575f80fd5b5061054b61073b366004614d33565b611e92565b34801561074b575f80fd5b506107c061075a366004614925565b5f9081527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60a60209081526040808320547fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab609909252909120546001600160a01b0390911691565b604080516001600160a01b03909316835260208301919091520161028d565b3480156107ea575f80fd5b506102b56107f93660046149c9565b611f29565b348015610809575f80fd5b506104dc73020000000000000000000000000000000000000581565b348015610830575f80fd5b507f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb00546001600160a01b03166104dc565b34801561086c575f80fd5b5061028061087b366004614925565b611f34565b34801561088b575f80fd5b5061097860408051610120810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e081018290526101008101829052905f8051602061552b833981519152604080516101208101825282546001600160a01b0390811682526001840154602083015260028401549282019290925260038301546001600160401b0381166060830152600160401b810461ffff1660808301526a0100000000000000000000900460ff1660a0820152600483015460c0820152600583015490911660e082015260069091015461010082015292915050565b60405161028d9190614d92565b348015610990575f80fd5b506107c061099f366004614925565b5f9081527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60c60209081526040808320547fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60b909252909120546001600160a01b0390911691565b348015610a10575f80fd5b5061054b610a1f366004614925565b5f9081527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb04602052604090205490565b348015610a5a575f80fd5b506102b5610a69366004614d05565b611fb0565b5f8181527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb036020908152604091829020805483518184028101840190945280845260609392830182828015610ae057602002820191905f5260205f20905b815481526020019060010190808311610acc575b50505050509050919050565b610af4612098565b5f8281527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6086020526040808220815160e0810190925280545f8051602061552b83398151915293929190829060ff166003811115610b5457610b54614ac4565b6003811115610b6557610b65614ac4565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c0909101529050600381516003811115610bde57610bde614ac4565b14610c08578051604051633b0d540d60e21b8152610bff9190600401614e35565b60405180910390fd5b81546040828101519051636af907fb60e11b815260048101919091525f916001600160a01b03169063d5f20ff6906024015f60405180830381865afa158015610c53573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610c7a9190810190614f56565b9050600483546040808501519051636af907fb60e11b81526001600160a01b039092169163d5f20ff691610cb49160040190815260200190565b5f60405180830381865afa158015610cce573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610cf59190810190614f56565b516005811115610d0757610d07614ac4565b14158015610d2e57508160c001516001600160401b031681608001516001600160401b0316105b15610e2457825460405163338587c560e21b815263ffffffff861660048201525f9182916001600160a01b039091169063ce161f149060240160408051808303815f875af1158015610d82573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610da69190615035565b9150915081846040015114610ddf5781846040015160405163fee3144560e01b8152600401610bff929190918252602082015260400190565b806001600160401b03168460c001516001600160401b03161115610e2157604051632e19bc2d60e11b81526001600160401b0382166004820152602401610bff565b50505b610e2d856120fb565b505050610e5960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5050565b610e688383836123a8565b50505050565b5f8181527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6086020526040808220815160e0810190925280545f8051602061552b83398151915293929190829060ff166003811115610ece57610ece614ac4565b6003811115610edf57610edf614ac4565b8152815461010090046001600160a01b0316602082015260018083015460408301526002909201546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015290915081516003811115610f5857610f58614ac4565b14158015610f795750600381516003811115610f7657610f76614ac4565b14155b15610f9a578051604051633b0d540d60e21b8152610bff9190600401614e35565b81546040828101519051636af907fb60e11b815260048101919091525f916001600160a01b03169063d5f20ff6906024015f60405180830381865afa158015610fe5573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261100c9190810190614f56565b905080606001516001600160401b03165f03611057576040517fe6e253e400000000000000000000000000000000000000000000000000000000815260048101859052602401610bff565b604080830151606083015160a084015192517f854a893f0000000000000000000000000000000000000000000000000000000081527302000000000000000000000000000000000000059363ee5b48eb9373__$caf7a49f90ce02b15fb10b3b072d8b9489$__9363854a893f936110eb93906004019283526001600160401b03918216602084015216604082015260600190565b5f60405180830381865af4158015611105573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261112c9190810190615058565b6040518263ffffffff1660e01b815260040161114891906150b7565b6020604051808303815f875af1158015611164573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061118891906150c9565b5050505050565b5f8281527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60760205260409020546001600160a01b03166111e5576040516330efa98b60e01b815260048101839052602401610bff565b5f5f8051602061552b83398151915254604051636af907fb60e11b8152600481018590526001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015611238573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261125f9190810190614f56565b519050600281600581111561127657611276614ac4565b14611296578060405163170cc93360e21b8152600401610bff91906150e0565b610e68838361269c565b610e68838383612998565b6112b6838383612de8565b505050565b5f805f8051602061552b833981519152600401546112d9908461510e565b90508015806112ee57506001600160401b0381115b1561130f5760405163222d164360e21b815260048101849052602401610bff565b92915050565b5f61131e612098565b61132b8534868686612e2c565b905061135660017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b949350505050565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091525f8281527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab608602052604090819020815160e081019092528054829060ff1660038111156113e5576113e5614ac4565b60038111156113f6576113f6614ac4565b8152815461010090046001600160a01b03166020820152600182015460408201526002909101546001600160401b038082166060840152600160401b820481166080840152600160801b8204811660a0840152600160c01b9091041660c09091015292915050565b5f8281527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6086020526040808220815160e0810190925280545f8051602061552b83398151915293929190829060ff1660038111156114be576114be614ac4565b60038111156114cf576114cf614ac4565b8152815461010090046001600160a01b03908116602083015260018301546040808401919091526002909301546001600160401b038082166060850152600160401b820481166080850152600160801b8204811660a0850152600160c01b9091041660c0909201919091528282015185549251636af907fb60e11b815260048101829052939450925f929091169063d5f20ff6906024015f60405180830381865afa158015611580573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526115a79190810190614f56565b90506001835160038111156115be576115be614ac4565b146115df578251604051633b0d540d60e21b8152610bff9190600401614e35565b6004815160058111156115f4576115f4614ac4565b0361160a57611602866120fb565b505050505050565b8260a001516001600160401b031681608001516001600160401b0316101561171257835460405163338587c560e21b815263ffffffff871660048201525f9182916001600160a01b039091169063ce161f149060240160408051808303815f875af115801561167b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061169f9190615035565b915091508184146116cd5760405163fee3144560e01b81526004810183905260248101859052604401610bff565b8460a001516001600160401b0316816001600160401b0316101561170f57604051632e19bc2d60e11b81526001600160401b0382166004820152602401610bff565b50505b5f868152600885016020908152604091829020805460ff1916600290811782550180546001600160401b034216600160401b81027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff909216919091179091559151918252839188917f3886b7389bccb22cac62838dee3f400cf8b22289295283e01a2c7093f93dd5aa910160405180910390a3505050505050565b7fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab604545f9061130f906001600160401b03841661512d565b5f6117ed612098565b6118898c8c8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525061187592508d91506151449050565b61187e8b615144565b8a8a348b8b8b612ea6565b90506118b460017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b9b9a5050505050505050505050565b5f8181527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb066020908152604091829020805483518184028101840190945280845260609392830182828015610ae057602002820191905f5260205f2090815481526020019060010190808311610acc5750505050509050919050565b5f8051602061552b8339815191526001600160a01b03821661197f5760405163caa903f960e01b81526001600160a01b0383166004820152602401610bff565b5f8381526007820160205260409020546001600160a01b031633146119de57335b6040517fdc599aea0000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610bff565b5f838152600c8201602052604080822080546001600160a01b038681166001600160a01b0319831681179093559251921692839287917f28c6fc4db51556a07b41aa23b91cedb22c02a7560c431a31255c03ca6ad61c3391a450505050565b5f5f8051602061552b8339815191528054604051636af907fb60e11b8152600481018590529192505f916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015611a96573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611abd9190810190614f56565b5190506004816005811115611ad457611ad4614ac4565b14611af4578060405163170cc93360e21b8152600401610bff91906150e0565b5f8381526007830160205260409020546001600160a01b03163314611b1957336119a0565b5f838152600c830160205260409020546001600160a01b031680611b5257505f8381526007830160205260409020546001600160a01b03165b610e688185612f65565b5f611b65612098565b5f5f8051602061552b83398151915280546040517f9681d94000000000000000000000000000000000000000000000000000000000815263ffffffff861660048201529192505f916001600160a01b0390911690639681d940906024016020604051808303815f875af1158015611bde573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c0291906150c9565b8254604051636af907fb60e11b8152600481018390529192505f916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa158015611c4c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052611c739190810190614f56565b5f8381527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60760205260409020549091506001600160a01b0316611cba57509150611d319050565b5f828152600784016020908152604080832054600c8701909252909120546001600160a01b03918216911680611ced5750805b600483516005811115611d0257611d02614ac4565b03611d1157611d118185612f65565b611d2982611d2285604001516117ad565b865f612ff9565b509193505050505b611d5a60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015611da35750825b90505f826001600160401b03166001148015611dbe5750303b155b905081158015611dcc575080155b15611e03576040517ff92ee8a900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b845467ffffffffffffffff191660011785558315611e3257845468ff00000000000000001916600160401b1785555b611e3d88888861306f565b8315611e8857845468ff000000000000000019168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050565b5f5f8051602061552b833981519152546040517fa3a65e4800000000000000000000000000000000000000000000000000000000815263ffffffff841660048201526001600160a01b039091169063a3a65e48906024016020604051808303815f875af1158015611f05573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061130f91906150c9565b6112b683838361308a565b5f8181527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb026020908152604091829020805483518184028101840190945280845260609392830182828015610ae057602002820191905f5260205f2090815481526020019060010190808311610acc5750505050509050919050565b6001600160a01b038116611fe25760405163caa903f960e01b81526001600160a01b0382166004820152602401610bff565b5f8281527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60860205260409020545f8051602061552b833981519152906001600160a01b0361010090910416331461203957336119a0565b5f838152600a8201602052604080822080546001600160a01b038681166001600160a01b0319831681179093559251921692839287917f6b30f219ab3cc1c43b394679707f3856ff2f3c6f1f6c97f383c6b16687a1e00591a450505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f008054600119016120f5576040517f3ee5aeb500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60029055565b5f8181527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6086020526040808220815160e0810190925280545f8051602061552b83398151915293929190829060ff16600381111561215b5761215b614ac4565b600381111561216c5761216c614ac4565b815281546001600160a01b03610100909104811660208084019190915260018401546040808501919091526002909401546001600160401b038082166060860152600160401b820481166080860152600160801b8204811660a0860152600160c01b9091041660c09093019290925283830151865484517f09c1df66000000000000000000000000000000000000000000000000000000008152945195965090949116926309c1df669260048083019391928290030181865afa158015612235573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612259919061520b565b82608001516122689190615226565b6001600160401b031642101561229c5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610bff565b5f848152600884016020908152604080832080547fffffffffffffffffffffff00000000000000000000000000000000000000000016815560018101849055600201839055600a8601909152902080546001600160a01b031981169091556001600160a01b03168061230f575060208201515b5f8061231c8388866130ce565b9150915061233c856020015161233587606001516117ad565b8987612ff9565b6040805183815260208101839052859189917f5ecc5b26a9265302cf871229b3d983e5ca57dbb1448966c6c58b2d3c68bc7f7e910160405180910390a350505050505050565b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b5f805f8051602061552b83398151915280546040517fb6e6a2ca000000000000000000000000000000000000000000000000000000008152600481018890529192506001600160a01b03169063b6e6a2ca906024015f604051808303815f87803b158015612414575f80fd5b505af1158015612426573d5f803e3d5ffd5b50508254604051636af907fb60e11b8152600481018990525f93506001600160a01b03909116915063d5f20ff6906024015f60405180830381865afa158015612471573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526124989190810190614f56565b5f8781527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60760205260409020549091506001600160a01b03166124e057600192505050612695565b5f8681526007830160205260409020546001600160a01b0316331461250557336119a0565b5f86815260078301602052604090205460c082015161253491600160b01b90046001600160401b031690615226565b6001600160401b03168160e001516001600160401b0316101561257b5760e081015160405163fb6ce63f60e01b81526001600160401b039091166004820152602401610bff565b5f85156125935761258c878661269c565b90506125b1565b505f8681526007830160205260409020600101546001600160401b03165b600583015460408301515f916001600160a01b031690634f22429f906125d6906117ad565b60c086015160e0808801516040519185901b6001600160e01b031916825260048201939093526001600160401b0391821660248201819052604482015291811660648301528516608482015260a401602060405180830381865afa158015612640573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061266491906150c9565b90508084600b015f8a81526020019081526020015f205f828254612688919061524d565b9091555050151593505050505b9392505050565b6040517f6f82535000000000000000000000000000000000000000000000000000000000815263ffffffff821660048201525f908190819073020000000000000000000000000000000000000590636f825350906024015f60405180830381865afa15801561270d573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526127349190810190615260565b915091508061276f576040517f6b2f19e900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6065482515f8051602061552b83398151915291146127df5782516040517f6ba589a50000000000000000000000000000000000000000000000000000000081526004810191909152602401610bff565b60208301516001600160a01b0316156128355760208301516040517f026f3ae80000000000000000000000000000000000000000000000000000000081526001600160a01b039091166004820152602401610bff565b5f8073__$caf7a49f90ce02b15fb10b3b072d8b9489$__63088c246386604001516040518263ffffffff1660e01b815260040161287291906150b7565b6040805180830381865af415801561288c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128b09190615035565b915091508188146128de5760405163fee3144560e01b81526004810183905260248101899052604401610bff565b5f8881526007840160205260409020600101546001600160401b03908116908216111561296f575f888152600784016020908152604091829020600101805467ffffffffffffffff19166001600160401b038516908117909155915191825289917fec44148e8ff271f2d0bacef1142154abacb0abb3a29eb3eb50e2ca97e86d0435910160405180910390a261298d565b505f8781526007830160205260409020600101546001600160401b03165b979650505050505050565b5f8381527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab6086020526040808220815160e0810190925280545f8051602061552b8339815191529284929091829060ff1660038111156129f9576129f9614ac4565b6003811115612a0a57612a0a614ac4565b8152815461010090046001600160a01b03908116602083015260018301546040808401919091526002909301546001600160401b038082166060850152600160401b820481166080850152600160801b8204811660a0850152600160c01b9091041660c0909201919091528282015185549251636af907fb60e11b815260048101829052939450925f929091169063d5f20ff6906024015f60405180830381865afa158015612abb573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612ae29190810190614f56565b9050600283516003811115612af957612af9614ac4565b14612b1a578251604051633b0d540d60e21b8152610bff9190600401614e35565b60208301516001600160a01b03163314612bb6575f8281526007850160205260409020546001600160a01b03163314612b5357336119a0565b5f82815260078501602052604090205460c0820151612b8291600160b01b90046001600160401b031690615226565b6001600160401b0316421015612bb65760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610bff565b5f888152600a850160205260409020546001600160a01b0316600282516005811115612be457612be4614ac4565b03612d8f5760038501546080850151612c06916001600160401b031690615226565b6001600160401b0316421015612c3a5760405163fb6ce63f60e01b81526001600160401b0342166004820152602401610bff565b8715612c4c57612c4a838861269c565b505b5f8981526008860160205260409020805460ff191660031790558454606085015160a08401516001600160a01b03909216916366109669918691612c909190615306565b6040516001600160e01b031960e085901b16815260048101929092526001600160401b0316602482015260440160408051808303815f875af1158015612cd8573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612cfc9190615326565b505f8a8152600887016020526040812060020180546001600160401b03909316600160c01b0277ffffffffffffffffffffffffffffffffffffffffffffffff90931692909217909155612d5085838c6131ee565b9050838a7f5abe543af12bb7f76f6fa9daaa9d95d181c5e90566df58d3c012216b6245eeaf60405160405180910390a315159550612695945050505050565b600482516005811115612da457612da4614ac4565b03612dcc57612db484828b6131ee565b50612dbe896120fb565b600195505050505050612695565b815160405163170cc93360e21b8152610bff91906004016150e0565b612df3838383612998565b6112b6576040517f206d9f2200000000000000000000000000000000000000000000000000000000815260048101849052602401610bff565b5f612e3685613427565b5f612e42868686613494565b90505f612e51883384876134f7565b9050612e608187875f8c6139b9565b87817fc2934a1fcda9c92015cdfb9c1d5f732c9a0fa18fef9ef1e7fb5664a62feeb5698888604051612e93929190615352565b60405180910390a3979650505050505050565b5f828103612ee3576040517f9c8cf5c900000000000000000000000000000000000000000000000000000000815260048101849052602401610bff565b612eec85613427565b5f612ef8868686613494565b90505f612f0b8d8d8d8d8d8d888b613b8c565b9050612f1b81878760015f6139b9565b807f0cd286cc10c4b16cd455b574101a9370644b5f27786e92e214912742544f69688787604051612f4d929190615352565b60405180910390a29c9b505050505050505050505050565b5f8181527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60b6020526040812080549190555f8051602061552b83398151915290612faf8482613ed9565b836001600160a01b0316837f875feb58aa30eeee040d55b00249c5c8c5af4f27c95cd29d64180ad67400c6e483604051612feb91815260200190565b60405180910390a350505050565b5f6130048385613f5d565b905061301285858585613ff4565b6130256001600160a01b03861682614062565b846001600160a01b03167f198c7c2c66ba3bfc16ee51a2dc44bb82ae02f41cfc9a076e9f822f4c619461c48260405161306091815260200190565b60405180910390a25050505050565b613077614127565b61308283838361418b565b6112b66141f2565b6130958383836123a8565b6112b6576040517fb7fed07e00000000000000000000000000000000000000000000000000000000815260048101849052602401610bff565b5f8281527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab609602052604081205481905f8051602061552b833981519152908290819080156131e0575f8881526009850160209081526040808320839055898352600787019091529020546127109061315190600160a01b900461ffff168361512d565b61315b919061510e565b91508184600b015f8981526020019081526020015f205f82825461317f919061524d565b9091555061318f905082826153a2565b925061319b8984613ed9565b886001600160a01b0316887f3ffc31181aadb250503101bd718e5fce8c27650af8d3525b9f60996756efaf63856040516131d791815260200190565b60405180910390a35b509097909650945050505050565b5f8051602061552b83398151915280546040858101519051636af907fb60e11b815260048101919091525f929183916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa15801561324c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526132739190810190614f56565b90505f60038251600581111561328b5761328b614ac4565b14806132a957506004825160058111156132a7576132a7614ac4565b145b156132b9575060e08101516132d6565b6002825160058111156132ce576132ce614ac4565b03612dcc5750425b86608001516001600160401b0316816001600160401b0316116132fe575f9350505050612695565b600583015460608801515f916001600160a01b031690634f22429f90613323906117ad565b60c086015160808c01516040808e01515f90815260078b0160205281902060010154905160e086901b6001600160e01b031916815260048101949094526001600160401b0392831660248501529082166044840152818716606484015216608482015260a401602060405180830381865afa1580156133a4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133c891906150c9565b90506001600160a01b0387166133e057876020015196505b5f8681526009850160209081526040808320849055600a90960190529390932080546001600160a01b0388166001600160a01b031990911617905550909150509392505050565b7fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab604545f8051602061552b83398151915290821015610e59576040517fa2840baa00000000000000000000000000000000000000000000000000000000815260048101839052602401610bff565b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb01545f907f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb00906134e4908461512d565b6134ee908661524d565b95945050505050565b5f5f8051602061552b83398151915281613510856112bb565b5f8881527fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab60760205260409020549091506001600160a01b0316613569576040516330efa98b60e01b815260048101889052602401610bff565b6001600160a01b03841661359b5760405163caa903f960e01b81526001600160a01b0385166004820152602401610bff565b8154604051636af907fb60e11b8152600481018990525f9182916001600160a01b039091169063d5f20ff6906024015f60405180830381865afa1580156135e4573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261360b9190810190614f56565b9050828160a0015161361d9190615226565b915083600301600a9054906101000a90046001600160401b0316816040015161364691906153b5565b6001600160401b0316826001600160401b0316111561369c576040517fdaa3fc0a0000000000000000000000000000000000000000000000000000000081526001600160401b0383166004820152602401610bff565b5082546040517f66109669000000000000000000000000000000000000000000000000000000008152600481018a90526001600160401b03831660248201525f9182916001600160a01b039091169063661096699060440160408051808303815f875af115801561370f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137339190615326565b915091505f8a8360405160200161377992919091825260c01b7fffffffffffffffff00000000000000000000000000000000000000000000000016602082015260280190565b60408051601f1981840301815291815281516020928301205f81815260088a019093529120805491925060019160ff19168280021790555089866008015f8381526020019081526020015f205f0160016101000a8154816001600160a01b0302191690836001600160a01b031602179055508a866008015f8381526020019081526020015f206001018190555084866008015f8381526020019081526020015f206002015f6101000a8154816001600160401b0302191690836001600160401b031602179055505f866008015f8381526020019081526020015f2060020160086101000a8154816001600160401b0302191690836001600160401b0316021790555082866008015f8381526020019081526020015f2060020160106101000a8154816001600160401b0302191690836001600160401b031602179055505f866008015f8381526020019081526020015f2060020160186101000a8154816001600160401b0302191690836001600160401b031602179055508786600a015f8381526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550896001600160a01b03168b827f77499a5603260ef2b34698d88b31f7b1acf28c7b134ad4e3fa636501e6064d7786888a888f6040516139a39594939291906001600160401b039586168152938516602085015291909316604083015260608201929092526001600160a01b0391909116608082015260a00190565b60405180910390a49a9950505050505050505050565b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb00835f5b81811015613b5a575f8787838181106139f8576139f86153e0565b9050602002013590505f801b846004015f8381526020019081526020015f2054141580613a3357505f81815260058501602052604090205415155b15613a6d576040517fe62618c400000000000000000000000000000000000000000000000000000000815260048101829052602401610bff565b83546001600160a01b03166323b872dd336040516001600160e01b031960e084901b1681526001600160a01b039091166004820152306024820152604481018490526064015f604051808303815f87803b158015613ac9575f80fd5b505af1158015613adb573d5f803e3d5ffd5b505050508515613b1d575f8981526002850160209081526040808320805460018101825590845282842001849055838352600487019091529020899055613b51565b5f89815260038501602090815260408083208054600181018255908452828420018490558383526005870190915290208990555b506001016139dd565b5083613b83575f8381526006830160209081526040822080546001810182559083529120018790555b50505050505050565b7fafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab603545f905f8051602061552b83398151915290600160401b900461ffff9081169087161080613be0575061271061ffff8716115b15613c0457604051635f12e6c360e11b815261ffff87166004820152602401610bff565b60038101546001600160401b039081169086161015613c40576040516202a06d60e11b81526001600160401b0386166004820152602401610bff565b8060010154841080613c555750806002015484115b15613c765760405163222d164360e21b815260048101859052602401610bff565b6001600160a01b038316613ca85760405163caa903f960e01b81526001600160a01b0384166004820152602401610bff565b835f613cb3826112bb565b90505f835f015f9054906101000a90046001600160a01b03166001600160a01b0316639cb7624e8e8e8e8e876040518663ffffffff1660e01b8152600401613cff95949392919061545a565b6020604051808303815f875af1158015613d1b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613d3f91906150c9565b90505f33905080856007015f8481526020019081526020015f205f015f6101000a8154816001600160a01b0302191690836001600160a01b0316021790555089856007015f8481526020019081526020015f205f0160146101000a81548161ffff021916908361ffff16021790555088856007015f8481526020019081526020015f205f0160166101000a8154816001600160401b0302191690836001600160401b031602179055505f856007015f8481526020019081526020015f206001015f6101000a8154816001600160401b0302191690836001600160401b031602179055508685600c015f8481526020019081526020015f205f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550806001600160a01b0316827ff51ab9b5253693af2f675b23c4042ccac671873d5f188e405b30019f4c159b7f8c8c8b604051613ec09392919061ffff9390931683526001600160401b039190911660208301526001600160a01b0316604082015260600190565b60405180910390a3509c9b505050505050505050505050565b6040517f4f5aaaba0000000000000000000000000000000000000000000000000000000081526001600160a01b03831660048201526024810182905273020000000000000000000000000000000000000190634f5aaaba906044015f604051808303815f87803b158015613f4b575f80fd5b505af1158015611602573d5f803e3d5ffd5b5f8281527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0260205260408120547f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0090829015613fc957505f848152600282016020526040902054613fdb565b505f8481526003820160205260409020545b6001820154613fea908261512d565b6134ee90856153a2565b5f8281527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0260205260409020547f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb0090156140575761405283866141fa565b611188565b61118883868461430a565b8047101561409e576040517fcd786059000000000000000000000000000000000000000000000000000000008152306004820152602401610bff565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f81146140e7576040519150601f19603f3d011682016040523d82523d5f602084013e6140ec565b606091505b50509050806112b6576040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16614189576040517fd7e6bcf800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b614193614127565b5f811180156141a557508260c0013581105b156141df576040517f4c5a73bf00000000000000000000000000000000000000000000000000000000815260048101829052602401610bff565b6141e88361450b565b6112b68282614589565b614189614127565b5f8281527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb026020526040812080547f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb00925b818110156142f1575f838281548110614266576142666153e0565b5f9182526020909120015485546040516323b872dd60e01b81523060048201526001600160a01b038981166024830152604482018490529293509116906323b872dd906064015f604051808303815f87803b1580156142c3575f80fd5b505af11580156142d5573d5f803e3d5ffd5b5050505f9182525060048501602052604081205560010161424b565b505f8581526002840160205260408120611188916148f7565b5f8381527f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb036020526040812080547f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb00925b81811015614401575f838281548110614376576143766153e0565b5f9182526020909120015485546040516323b872dd60e01b81523060048201526001600160a01b038a81166024830152604482018490529293509116906323b872dd906064015f604051808303815f87803b1580156143d3575f80fd5b505af11580156143e5573d5f803e3d5ffd5b5050505f9182525060058501602052604081205560010161435b565b505f868152600384016020526040812061441a916148f7565b5f848152600684016020526040812054905b81811015611e88575f8681526006860160205260409020805489919083908110614458576144586153e0565b905f5260205f20015403614503575f86815260068601602052604090206144806001846153a2565b81548110614490576144906153e0565b905f5260205f200154856006015f8881526020019081526020015f2082815481106144bd576144bd6153e0565b905f5260205f200181905550846006015f8781526020019081526020015f208054806144eb576144eb6154c2565b600190038181905f5260205f20015f90559055611e88565b60010161442c565b614513614127565b61451b6145ff565b61458661452b60208301836154d6565b602083013560408401356145456080860160608701614b7b565b61455560a08701608088016154f1565b61456560c0880160a0890161550a565b60c088013561457b6101008a0160e08b016154d6565b89610100013561460f565b50565b614591614127565b7f19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb006001600160a01b0383166145d95760405163d92e233d60e01b815260040160405180910390fd5b80546001600160a01b0319166001600160a01b0393909316929092178255600190910155565b614607614127565b6141896148ef565b614617614127565b5f8051602061552b83398151915261ffff8616158061463b575061271061ffff8716115b1561465f57604051635f12e6c360e11b815261ffff87166004820152602401610bff565b878911156146835760405163222d164360e21b8152600481018a9052602401610bff565b60ff851615806146965750600a60ff8616115b156146d2576040517fb86d9ac800000000000000000000000000000000000000000000000000000000815260ff86166004820152602401610bff565b6001600160a01b038a166146f95760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0383166147205760405163d92e233d60e01b815260040160405180910390fd5b896001600160a01b03166309c1df666040518163ffffffff1660e01b8152600401602060405180830381865afa15801561475c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614780919061520b565b6001600160401b0316876001600160401b031610156147bc576040516202a06d60e11b81526001600160401b0388166004820152602401610bff565b835f036147f5576040517fa733007100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8161482f576040517f2f6bd1db00000000000000000000000000000000000000000000000000000000815260048101839052602401610bff565b80546001600160a01b039a8b166001600160a01b0319918216178255600182019990995560028101979097556003870180546a010000000000000000000060ff909616959095027fffffffffffffffffffffffffffff0000000000000000ffffffffffffffffffff61ffff909716600160401b0269ffffffffffffffffffff199096166001600160401b03909816979097179490941794909416949094179091556004840155600583018054929095169190931617909255600690910155565b612382614127565b5080545f8255905f5260205f209081019061458691905b80821115614921575f815560010161490e565b5090565b5f60208284031215614935575f80fd5b5035919050565b602080825282518282018190525f9190848201906040850190845b8181101561497357835183529284019291840191600101614957565b50909695505050505050565b803563ffffffff81168114611d5a575f80fd5b5f80604083850312156149a3575f80fd5b823591506149b36020840161497f565b90509250929050565b8015158114614586575f80fd5b5f805f606084860312156149db575f80fd5b8335925060208401356149ed816149bc565b91506149fb6040850161497f565b90509250925092565b5f8083601f840112614a14575f80fd5b5081356001600160401b03811115614a2a575f80fd5b6020830191508360208260051b8501011115614a44575f80fd5b9250929050565b6001600160a01b0381168114614586575f80fd5b8035611d5a81614a4b565b5f805f8060608587031215614a7d575f80fd5b8435935060208501356001600160401b03811115614a99575f80fd5b614aa587828801614a04565b9094509250506040850135614ab981614a4b565b939692955090935050565b634e487b7160e01b5f52602160045260245ffd5b60048110614ae857614ae8614ac4565b9052565b5f60e082019050614afe828451614ad8565b6001600160a01b0360208401511660208301526040830151604083015260608301516001600160401b0380821660608501528060808601511660808501528060a08601511660a08501528060c08601511660c0850152505092915050565b6001600160401b0381168114614586575f80fd5b8035611d5a81614b5c565b5f60208284031215614b8b575f80fd5b813561269581614b5c565b5f8083601f840112614ba6575f80fd5b5081356001600160401b03811115614bbc575f80fd5b602083019150836020828501011115614a44575f80fd5b5f60408284031215614be3575f80fd5b50919050565b803561ffff81168114611d5a575f80fd5b5f805f805f805f805f805f6101008c8e031215614c15575f80fd5b6001600160401b03808d351115614c2a575f80fd5b614c378e8e358f01614b96565b909c509a5060208d0135811015614c4c575f80fd5b614c5c8e60208f01358f01614b96565b909a50985060408d0135811015614c71575f80fd5b614c818e60408f01358f01614bd3565b97508060608e01351115614c93575f80fd5b614ca38e60608f01358f01614bd3565b9650614cb160808e01614be9565b9550614cbf60a08e01614b70565b94508060c08e01351115614cd1575f80fd5b50614ce28d60c08e01358e01614a04565b9093509150614cf360e08d01614a5f565b90509295989b509295989b9093969950565b5f8060408385031215614d16575f80fd5b823591506020830135614d2881614a4b565b809150509250929050565b5f60208284031215614d43575f80fd5b6126958261497f565b5f805f838503610160811215614d60575f80fd5b61012080821215614d6f575f80fd5b8594508401359050614d8081614a4b565b92959294505050610140919091013590565b5f610120820190506001600160a01b03835116825260208301516020830152604083015160408301526060830151614dd560608401826001600160401b03169052565b506080830151614deb608084018261ffff169052565b5060a0830151614e0060a084018260ff169052565b5060c083015160c083015260e0830151614e2560e08401826001600160a01b03169052565b5061010092830151919092015290565b6020810161130f8284614ad8565b634e487b7160e01b5f52604160045260245ffd5b60405161010081016001600160401b0381118282101715614e7a57614e7a614e43565b60405290565b604080519081016001600160401b0381118282101715614e7a57614e7a614e43565b604051601f8201601f191681016001600160401b0381118282101715614eca57614eca614e43565b604052919050565b805160068110611d5a575f80fd5b5f82601f830112614eef575f80fd5b81516001600160401b03811115614f0857614f08614e43565b614f1b601f8201601f1916602001614ea2565b818152846020838601011115614f2f575f80fd5b8160208501602083015e5f918101602001919091529392505050565b8051611d5a81614b5c565b5f60208284031215614f66575f80fd5b81516001600160401b0380821115614f7c575f80fd5b908301906101008286031215614f90575f80fd5b614f98614e57565b614fa183614ed2565b8152602083015182811115614fb4575f80fd5b614fc087828601614ee0565b602083015250614fd260408401614f4b565b6040820152614fe360608401614f4b565b6060820152614ff460808401614f4b565b608082015261500560a08401614f4b565b60a082015261501660c08401614f4b565b60c082015261502760e08401614f4b565b60e082015295945050505050565b5f8060408385031215615046575f80fd5b825191506020830151614d2881614b5c565b5f60208284031215615068575f80fd5b81516001600160401b0381111561507d575f80fd5b61135684828501614ee0565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b602081525f6126956020830184615089565b5f602082840312156150d9575f80fd5b5051919050565b60208101600683106150f4576150f4614ac4565b91905290565b634e487b7160e01b5f52601160045260245ffd5b5f8261512857634e487b7160e01b5f52601260045260245ffd5b500490565b808202811582820484141761130f5761130f6150fa565b5f60408236031215615154575f80fd5b61515c614e80565b6151658361497f565b81526020808401356001600160401b0380821115615181575f80fd5b9085019036601f830112615193575f80fd5b8135818111156151a5576151a5614e43565b8060051b91506151b6848301614ea2565b81815291830184019184810190368411156151cf575f80fd5b938501935b838510156151f957843592506151e983614a4b565b82825293850193908501906151d4565b94860194909452509295945050505050565b5f6020828403121561521b575f80fd5b815161269581614b5c565b6001600160401b03818116838216019080821115615246576152466150fa565b5092915050565b8082018082111561130f5761130f6150fa565b5f8060408385031215615271575f80fd5b82516001600160401b0380821115615287575f80fd5b908401906060828703121561529a575f80fd5b6040516060810181811083821117156152b5576152b5614e43565b6040528251815260208301516152ca81614a4b565b60208201526040830151828111156152e0575f80fd5b6152ec88828601614ee0565b6040830152508094505050506020830151614d28816149bc565b6001600160401b03828116828216039080821115615246576152466150fa565b5f8060408385031215615337575f80fd5b825161534281614b5c565b6020939093015192949293505050565b602081528160208201525f7f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115615389575f80fd5b8260051b80856040850137919091016040019392505050565b8181038181111561130f5761130f6150fa565b6001600160401b038181168382160280821691908281146153d8576153d86150fa565b505092915050565b634e487b7160e01b5f52603260045260245ffd5b5f6040830163ffffffff8351168452602080840151604060208701528281518085526060880191506020830194505f92505b8083101561544f5784516001600160a01b03168252938301936001929092019190830190615426565b509695505050505050565b60a081525f61546c60a0830188615089565b828103602084015261547e8188615089565b9050828103604084015261549281876153f4565b905082810360608401526154a681866153f4565b9150506001600160401b03831660808301529695505050505050565b634e487b7160e01b5f52603160045260245ffd5b5f602082840312156154e6575f80fd5b813561269581614a4b565b5f60208284031215615501575f80fd5b61269582614be9565b5f6020828403121561551a575f80fd5b813560ff81168114612695575f80fdfeafe6c4731b852fc2be89a0896ae43d22d8b24989064d841b2a1586b4d39ab600a26469706673582212204be8b64e509ea556847e7bf989d6f058765a9d0faf10e8d524d5150eabe1bf1964736f6c63430008190033",
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

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x3d2e15e5.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) InitiateDelegatorRegistration(opts *bind.TransactOpts, validationID [32]byte, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "initiateDelegatorRegistration", validationID, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x3d2e15e5.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) InitiateDelegatorRegistration(validationID [32]byte, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateDelegatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, licenseTokenIds, rewardRecipient)
}

// InitiateDelegatorRegistration is a paid mutator transaction binding the contract method 0x3d2e15e5.
//
// Solidity: function initiateDelegatorRegistration(bytes32 validationID, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) InitiateDelegatorRegistration(validationID [32]byte, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateDelegatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, validationID, licenseTokenIds, rewardRecipient)
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

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x6c60f49f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactor) InitiateValidatorRegistration(opts *bind.TransactOpts, nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.contract.Transact(opts, "initiateValidatorRegistration", nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x6c60f49f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateValidatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, licenseTokenIds, rewardRecipient)
}

// InitiateValidatorRegistration is a paid mutator transaction binding the contract method 0x6c60f49f.
//
// Solidity: function initiateValidatorRegistration(bytes nodeID, bytes blsPublicKey, (uint32,address[]) remainingBalanceOwner, (uint32,address[]) disableOwner, uint16 delegationFeeBips, uint64 minStakeDuration, uint256[] licenseTokenIds, address rewardRecipient) payable returns(bytes32)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerTransactorSession) InitiateValidatorRegistration(nodeID []byte, blsPublicKey []byte, remainingBalanceOwner PChainOwner, disableOwner PChainOwner, delegationFeeBips uint16, minStakeDuration uint64, licenseTokenIds []*big.Int, rewardRecipient common.Address) (*types.Transaction, error) {
	return _NativeTokenLicensedStakingManager.Contract.InitiateValidatorRegistration(&_NativeTokenLicensedStakingManager.TransactOpts, nodeID, blsPublicKey, remainingBalanceOwner, disableOwner, delegationFeeBips, minStakeDuration, licenseTokenIds, rewardRecipient)
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

// NativeTokenLicensedStakingManagerNativeTokensUnlockedIterator is returned from FilterNativeTokensUnlocked and is used to iterate over the raw logs and unpacked data for NativeTokensUnlocked events raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerNativeTokensUnlockedIterator struct {
	Event *NativeTokenLicensedStakingManagerNativeTokensUnlocked // Event containing the contract specifics and raw log

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
func (it *NativeTokenLicensedStakingManagerNativeTokensUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeTokenLicensedStakingManagerNativeTokensUnlocked)
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
		it.Event = new(NativeTokenLicensedStakingManagerNativeTokensUnlocked)
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
func (it *NativeTokenLicensedStakingManagerNativeTokensUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeTokenLicensedStakingManagerNativeTokensUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeTokenLicensedStakingManagerNativeTokensUnlocked represents a NativeTokensUnlocked event raised by the NativeTokenLicensedStakingManager contract.
type NativeTokenLicensedStakingManagerNativeTokensUnlocked struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeTokensUnlocked is a free log retrieval operation binding the contract event 0x198c7c2c66ba3bfc16ee51a2dc44bb82ae02f41cfc9a076e9f822f4c619461c4.
//
// Solidity: event NativeTokensUnlocked(address indexed to, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) FilterNativeTokensUnlocked(opts *bind.FilterOpts, to []common.Address) (*NativeTokenLicensedStakingManagerNativeTokensUnlockedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.FilterLogs(opts, "NativeTokensUnlocked", toRule)
	if err != nil {
		return nil, err
	}
	return &NativeTokenLicensedStakingManagerNativeTokensUnlockedIterator{contract: _NativeTokenLicensedStakingManager.contract, event: "NativeTokensUnlocked", logs: logs, sub: sub}, nil
}

// WatchNativeTokensUnlocked is a free log subscription operation binding the contract event 0x198c7c2c66ba3bfc16ee51a2dc44bb82ae02f41cfc9a076e9f822f4c619461c4.
//
// Solidity: event NativeTokensUnlocked(address indexed to, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) WatchNativeTokensUnlocked(opts *bind.WatchOpts, sink chan<- *NativeTokenLicensedStakingManagerNativeTokensUnlocked, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NativeTokenLicensedStakingManager.contract.WatchLogs(opts, "NativeTokensUnlocked", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeTokenLicensedStakingManagerNativeTokensUnlocked)
				if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "NativeTokensUnlocked", log); err != nil {
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

// ParseNativeTokensUnlocked is a log parse operation binding the contract event 0x198c7c2c66ba3bfc16ee51a2dc44bb82ae02f41cfc9a076e9f822f4c619461c4.
//
// Solidity: event NativeTokensUnlocked(address indexed to, uint256 amount)
func (_NativeTokenLicensedStakingManager *NativeTokenLicensedStakingManagerFilterer) ParseNativeTokensUnlocked(log types.Log) (*NativeTokenLicensedStakingManagerNativeTokensUnlocked, error) {
	event := new(NativeTokenLicensedStakingManagerNativeTokensUnlocked)
	if err := _NativeTokenLicensedStakingManager.contract.UnpackLog(event, "NativeTokensUnlocked", log); err != nil {
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
