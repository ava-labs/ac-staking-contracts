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
 * This contract allows ERC721 tokens to server as license for staking.
 */
abstract contract LicensedStakingManager is
    Initializable,
    StakingManager,
    ILicensedStakingManager
{
    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:ac-staking.storage.LicensedStakingManager
    struct LicensedStakingManagerStorage {
        // License token address
        IERC721 _token;
        // Conversion factor for license tokens to stake amount
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

    // keccak256(abi.encode(uint256(keccak256("ac-staking.storage.LicensedStakingManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant LICENSED_STAKING_MANAGER_STORAGE_LOCATION =
        0x7e5bdfcce15e53c3406ea67bfce37dcd26f5152d5492824e43fd5e3c8ac5ab01; // TODO: placeholder fix it

    error TokenAlreadyStaked(uint256 tokenId);
    error TokenNotStaked(uint256 tokenId);
    error InvalidTokenCount(uint256 count);

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
     * @param token The ERC721 token to be staked
     * @param licenseToStakeConversionFactor_ The conversion factor for license tokens to stake amount
     */
    // solhint-disable-next-line func-name-mixedcase
    function __LicensedStakingManager_init(
        StakingManagerSettings calldata settings,
        IERC721 token,
        uint256 licenseToStakeConversionFactor_
    ) internal onlyInitializing {
        __StakingManager_init(settings);
        __LicensedStakingManager_init_unchained(token, licenseToStakeConversionFactor_);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __LicensedStakingManager_init_unchained(
        IERC721 token,
        uint256 licenseToStakeConversionFactor_
    ) internal onlyInitializing {
        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();
        if (address(token) == address(0)) {
            revert ZeroAddress();
        }
        $._token = token;
        $._licenseToStakeConversionFactor = licenseToStakeConversionFactor_;
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
            revert InvalidTokenCount(licenseTokenIds.length);
        }

        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();
        // Calculate total stake value based on number of tokens
        uint256 stakeAmount =
            licenseTokenIds.length * $._licenseToStakeConversionFactor + tokenStakeAmount;

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
            revert InvalidTokenCount(licenseTokenIds.length);
        }

        LicensedStakingManagerStorage storage $ = _getLicensedStakingManagerStorage();

        // Calculate total delegation value based on number of tokens
        uint256 fullDelegationAmount =
            licenseTokenIds.length * $._licenseToStakeConversionFactor + delegationAmount;

        bytes32 delegationID = _initiateDelegatorRegistration(
            validationID, _msgSender(), fullDelegationAmount, rewardRecipient
        );

        // Lock ERC721 tokens
        _lockERC721s(delegationID, licenseTokenIds, false, validationID);
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
                revert TokenAlreadyStaked(tokenId);
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
     * TODO: Add documentation
     */
    function _lockTokens(
        uint256 value
    ) internal virtual returns (uint256);

    /**
     * @notice See {StakingManager-_unlock}
     * @dev This function is not used for ERC721 tokens
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
     * TODO: Add documentation
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
        // TODO: Think about this
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
