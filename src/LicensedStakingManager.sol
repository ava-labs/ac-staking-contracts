// (c) 2025, Ava Labs, Inc. All rights reserved.
// SPDX-License-Identifier: LicenseRef-AvaLabsProprietary

// This code is licensed under a custom license.
// See the file LICENSE for licensing terms.
pragma solidity 0.8.25;

import {StakingManager} from "./StakingManager.sol";
import {ILicensedStakingManager} from "./interfaces/ILicensedStakingManager.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {PChainOwner} from "@validator-manager/interfaces/IACP99Manager.sol";
import {StakingManagerSettings} from "@validator-manager/interfaces/IStakingManager.sol";

/**
 * @dev Implementation of the {ILicensedStakingManager} interface.
 * This contract adds additional requirement for staking and delegation
 * in shape of ERC721 tokens.
 * ERC721 tokens are licenses that enable users to stake and delegate.
 */
abstract contract LicensedStakingManager is
    Initializable,
    StakingManager,
    ILicensedStakingManager
{
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avacloud.storage.LicensedStakingManager
    struct LicensedStakingManagerStorage {
        // License token address
        IERC721 _token;
        // Conversion factor for license tokens to stake amount
        // See `_totalStakeAmount` function on how the total stake amount is calculated
        uint256 _licenseToStakeConversionFactor;
        // Maps validationID to array of staked token IDs for validators
        mapping(bytes32 => uint256[]) _validatorStakedTokens;
        // Maps delegationID to array of staked token IDs for delegators
        mapping(bytes32 => uint256[]) _delegatorStakedTokens;
        // Maps token ID to validationID to track which validator staked which token
        mapping(uint256 => bytes32) _tokenToValidator;
        // Maps token ID to delegationID to track which delegator staked which token
        mapping(uint256 => bytes32) _tokenToDelegation;
        // Maps validationID to array of delegationIDs for that validator
        mapping(bytes32 => bytes32[]) _validatorDelegations;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avacloud.storage.LicensedStakingManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant LICENSED_STAKING_MANAGER_STORAGE_LOCATION =
        0x19dadeb42f0d6e5189a2ab22a0f2b6d0770581cc39ab2fd1633a273ba24bdb00;

    error LicenseTokenAlreadyStaked(uint256 licenseTokenId);
    error InvalidLicenseTokenCount(uint256 count);
    error InvalidTokenStakeAmount(uint256 amount);
    error InvalidLicenseToStakeConversionFactor(uint256 factor);

    // solhint-disable ordering
    function _getLicensedStakingManagerStorage()
        private
        pure
        returns (LicensedStakingManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := LICENSED_STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    /**
     * @notice Initialize the ERC721 token staking manager
     * @param settings Initial settings for the PoS validator manager
     * @param licenseToken The ERC721 token used as a staking license
     * @param licenseToStakeConversionFactor_ The conversion factor for license tokens to stake amount
     */
    // solhint-disable-next-line func-name-mixedcase
    function __LicensedStakingManager_init(
        StakingManagerSettings calldata settings,
        IERC721 licenseToken,
        uint256 licenseToStakeConversionFactor_
    ) internal onlyInitializing {
        // License to stake conversion factor can be zero, but if it's not zero, it must be greater than weight to value factor
        // to avoid loss of token funds
        if (
            licenseToStakeConversionFactor_ > 0
                && licenseToStakeConversionFactor_ < settings.weightToValueFactor
        ) {
            revert InvalidLicenseToStakeConversionFactor(licenseToStakeConversionFactor_);
        }
        __StakingManager_init(settings);
        __LicensedStakingManager_init_unchained(licenseToken, licenseToStakeConversionFactor_);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __LicensedStakingManager_init_unchained(
        IERC721 licenseToken,
        uint256 licenseToStakeConversionFactor_
    ) internal onlyInitializing {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();
        if (address(licenseToken) == address(0)) {
            revert ZeroAddress();
        }
        $._token = licenseToken;
        $._licenseToStakeConversionFactor = licenseToStakeConversionFactor_;
    }

    /**
     * @notice Validates the amount of tokens intended to be staked
     * @dev This check is performed in `_initiateLicensedValidatorRegistration` and `_initiateLicensedDelegatorRegistration`
     * It is required because otherwise it would be possible for tokens to stay forever locked in the contract
     * during `weightToValue` and `valueToWeight` conversions.
     * Example:
     * - delegation token amount is 100, but weightToValueFactor is 1000.
     * - _licenseToStakeConversionFactor is 1000 and 1 license token is staked
     * - total stake amount is 100 + 1 * 1000 = 1100 will be converted to weight 1
     * - when unlocking weight 1 will be coverted to 1000, so when license stake amount is deducted, 0 tokens will be returned to user.
     * @param tokenStakeAmount The amount of tokens to stake
     */
    function _validateTokenStakeAmount(
        uint256 tokenStakeAmount
    ) internal view {
        StakingManagerStorage storage $ = _getStakingManagerStorage();
        if (tokenStakeAmount < $._weightToValueFactor) {
            revert InvalidTokenStakeAmount(tokenStakeAmount);
        }
    }

    /**
     * @notice Calculates the total stake amount
     * @param tokenStakeAmount The amount of tokens to stake
     * @param licenseTokenIds The amount of license tokens to stake
     * @return The total stake amount
     */
    function _totalStakeAmount(
        uint256 tokenStakeAmount,
        uint256[] calldata licenseTokenIds
    ) internal view returns (uint256) {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();
        return tokenStakeAmount + licenseTokenIds.length * $._licenseToStakeConversionFactor;
    }

    /**
     * @notice See {ILicensedStakingManager-initiateLicensedValidatorRegistration}
     */
    function _initiateLicensedValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 tokenStakeAmount,
        uint256[] calldata licenseTokenIds,
        address rewardRecipient
    ) internal returns (bytes32) {
        if (licenseTokenIds.length == 0) {
            revert InvalidLicenseTokenCount(licenseTokenIds.length);
        }
        _validateTokenStakeAmount(tokenStakeAmount);
        // Calculate total stake value based on number of tokens
        uint256 stakeAmount = _totalStakeAmount(tokenStakeAmount, licenseTokenIds);

        bytes32 validationID = _initiateValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            stakeAmount: stakeAmount,
            rewardRecipient: rewardRecipient
        });

        // Lock ERC721 tokens
        _lockERC721s(validationID, licenseTokenIds, true, bytes32(0));
        // Lock stake tokens
        _lockTokens(tokenStakeAmount);

        emit ValidatorRegisteredWithLicenses(validationID, licenseTokenIds);

        return validationID;
    }

    /**
     * @notice See {ILicensedStakingManager-initiateDelegatorRegistration}
     */
    function _initiateLicensedDelegatorRegistration(
        bytes32 validationID,
        uint256 delegationAmount,
        uint256[] calldata licenseTokenIds,
        address rewardRecipient
    ) internal returns (bytes32) {
        if (licenseTokenIds.length == 0) {
            revert InvalidLicenseTokenCount(licenseTokenIds.length);
        }
        _validateTokenStakeAmount(delegationAmount);
        // Calculate total delegation value based on number of tokens
        uint256 totalDelegationAmount = _totalStakeAmount(delegationAmount, licenseTokenIds);

        bytes32 delegationID = _initiateDelegatorRegistration(
            validationID, _msgSender(), totalDelegationAmount, rewardRecipient
        );

        // Lock ERC721 tokens
        _lockERC721s(delegationID, licenseTokenIds, false, validationID);
        // Lock stake tokens
        _lockTokens(delegationAmount);

        emit DelegatorRegisteredWithLicenses(delegationID, validationID, licenseTokenIds);

        return delegationID;
    }

    /**
     * @notice Locks ERC721 tokens for staking
     * @param id The validation ID or delegation ID
     * @param licenseTokenIds Array of token IDs to stake
     * @param isValidator Whether the tokens are being staked by a validator or delegator
     * @param validationID The validation ID (only needed for delegators)
     */
    function _lockERC721s(
        bytes32 id,
        uint256[] calldata licenseTokenIds,
        bool isValidator,
        bytes32 validationID
    ) internal {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();

        for (uint256 i = 0; i < licenseTokenIds.length; i++) {
            uint256 tokenId = licenseTokenIds[i];

            // Check if token is already staked
            if (
                $._tokenToValidator[tokenId] != bytes32(0)
                    || $._tokenToDelegation[tokenId] != bytes32(0)
            ) {
                revert LicenseTokenAlreadyStaked(tokenId);
            }

            // Transfer token from user to contract
            $._token.transferFrom(_msgSender(), address(this), tokenId);

            // Store token information
            if (isValidator) {
                $._validatorStakedTokens[id].push(tokenId);
                $._tokenToValidator[tokenId] = id;
            } else {
                $._delegatorStakedTokens[id].push(tokenId);
                $._tokenToDelegation[tokenId] = id;
                // Track this delegation for the validator
                $._validatorDelegations[validationID].push(id);
            }
        }
    }

    /**
     * @notice See {StakingManager-_lock}
     * @dev This function is not transferring any tokens, it's just returning the value.
     * License tokens are locked in the contract when the `_lockERC721s` function is called.
     * Stake tokens are locked in the contract when the `_lockTokens` function is called.
     */
    function _lock(
        uint256 value
    ) internal pure override returns (uint256) {
        return value;
    }

    /**
     * @notice Locks state tokens in the contract
     * @dev Since the `_lock` function from staking manager contract only accepts full stake amount,
     * we wouldn't be able to determine the amount of stake tokens to lock. Therefore `_lockTokens`
     * is used to lock stake tokens and called from `_initiateLicensedValidatorRegistration` and
     * `_initiateLicensedDelegatorRegistration` functions where the exact amounts are known.
     * @param value The amount of tokens to lock
     */
    function _lockTokens(
        uint256 value
    ) internal virtual returns (uint256);

    /**
     * @notice See {StakingManager-_unlock}
     * @dev Used to unlock license tokens.
     * IMPORTANT: Value parameter is not used and stake tokens should be unlocked in the implementation of the contract.
     */
    function _unlock(address to, uint256, bytes32 stakeId) internal virtual override {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();
        if ($._validatorStakedTokens[stakeId].length > 0) {
            _returnValidatorTokens(stakeId, to);
        } else {
            _returnDelegatorTokens(stakeId, to);
        }
    }

    /**
     * @notice Calculates the amount of stake tokens from the total stake amount
     * @dev This function is inverted to `_totalStakeAmount` function.
     * @param stakeId The validation ID or delegation ID
     * @param stakeAmount The total stake amount
     * @return The amount of stake tokens
     */
    function _tokenAmountFromStakeAmount(
        bytes32 stakeId,
        uint256 stakeAmount
    ) internal view returns (uint256) {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();
        uint256 licensesAmount = 0;
        if ($._validatorStakedTokens[stakeId].length > 0) {
            licensesAmount = $._validatorStakedTokens[stakeId].length;
        } else {
            licensesAmount = $._delegatorStakedTokens[stakeId].length;
        }
        return stakeAmount - licensesAmount * $._licenseToStakeConversionFactor;
    }

    /**
     * @notice Returns ERC721 tokens to the validator
     * @param validationID The validation ID
     * @param owner The address to return tokens to
     */
    function _returnValidatorTokens(bytes32 validationID, address owner) internal {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();

        // Return validator's own tokens
        uint256[] storage validatorTokens = $._validatorStakedTokens[validationID];
        for (uint256 i = 0; i < validatorTokens.length; i++) {
            uint256 tokenId = validatorTokens[i];
            $._token.transferFrom(address(this), owner, tokenId);
            delete $._tokenToValidator[tokenId];
        }
        delete $._validatorStakedTokens[validationID];
    }

    /**
     * @notice Returns ERC721 tokens to the delegator
     * @param delegationID The delegation ID
     * @param owner The address to return tokens to
     */
    function _returnDelegatorTokens(bytes32 delegationID, address owner) internal {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();

        uint256[] storage delegatorTokens = $._delegatorStakedTokens[delegationID];
        for (uint256 i = 0; i < delegatorTokens.length; i++) {
            uint256 tokenId = delegatorTokens[i];
            $._token.transferFrom(address(this), owner, tokenId);
            delete $._tokenToDelegation[tokenId];
        }
        delete $._delegatorStakedTokens[delegationID];
        // Remove delegationID from validatorDelegations
        for (uint256 i = 0; i < $._validatorDelegations[delegationID].length; i++) {
            if ($._validatorDelegations[delegationID][i] == delegationID) {
                delete $._validatorDelegations[delegationID][i];
            }
        }
    }

    /**
     * @notice See {ILicensedStakingManager-erc721}
     */
    function erc721() external view returns (IERC721) {
        return _getLicensedStakingManagerStorage()._token;
    }

    /**
     * @notice See {ILicensedStakingManager-getValidatorStakedLicenseTokens}
     */
    function getValidatorStakedLicenseTokens(
        bytes32 validationID
    ) external view returns (uint256[] memory) {
        return _getLicensedStakingManagerStorage()._validatorStakedTokens[validationID];
    }

    /**
     * @notice See {ILicensedStakingManager-getDelegatorStakedLicenseTokens}
     */
    function getDelegatorStakedLicenseTokens(
        bytes32 delegationID
    ) external view returns (uint256[] memory) {
        return _getLicensedStakingManagerStorage()._delegatorStakedTokens[delegationID];
    }

    /**
     * @notice See {ILicensedStakingManager-getValidatorDelegations}
     */
    function getValidatorDelegations(
        bytes32 validationID
    ) external view returns (bytes32[] memory) {
        return _getLicensedStakingManagerStorage()._validatorDelegations[validationID];
    }

    /**
     * @notice See {ILicensedStakingManager-getLicenseTokenValidator}
     */
    function getLicenseTokenValidator(
        uint256 licenseTokenId
    ) external view returns (bytes32) {
        return _getLicensedStakingManagerStorage()._tokenToValidator[licenseTokenId];
    }

    /**
     * @notice See {ILicensedStakingManager-getLicenseTokenDelegator}
     */
    function getLicenseTokenDelegator(
        uint256 licenseTokenId
    ) external view returns (bytes32) {
        return _getLicensedStakingManagerStorage()._tokenToDelegation[licenseTokenId];
    }

    /**
     * @notice See {ILicensedStakingManager-licenseToStakeConversionFactor}
     */
    function licenseToStakeConversionFactor() external view returns (uint256) {
        return _getLicensedStakingManagerStorage()._licenseToStakeConversionFactor;
    }
}
