// (c) 2025, Ava Labs, Inc. All rights reserved.
// SPDX-License-Identifier: LicenseRef-AvaLabsProprietary

// This code is licensed under a custom license.
// See the file LICENSE for licensing terms.
pragma solidity 0.8.25;

import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {PChainOwner} from "@validator-manager/interfaces/IACP99Manager.sol";
import {
    IStakingManager,
    StakingManagerSettings
} from "@validator-manager/interfaces/IStakingManager.sol";

/**
 * @dev Interface for the LicensedStakingManager contract.
 * This contract requires ERC721 tokens as license for staking.
 */
interface ILicensedStakingManager is IStakingManager {
    /**
     * @notice Emitted when a validator registers with ERC721 tokens
     * @param validationID The validation ID of the validator
     * @param licenseTokenIds Array of license token IDs that were staked
     */
    event ValidatorRegisteredWithLicenses(bytes32 indexed validationID, uint256[] licenseTokenIds);

    /**
     * @notice Emitted when a delegator registers with ERC721 tokens
     * @param delegationID The delegation ID
     * @param validationID The validation ID of the validator being delegated to
     * @param licenseTokenIds Array of license token IDs that were staked
     */
    event DelegatorRegisteredWithLicenses(
        bytes32 indexed delegationID, bytes32 indexed validationID, uint256[] licenseTokenIds
    );

    /**
     * @notice Get the ERC721 license token contract being used for staking
     * @return The ERC721 license token contract
     */
    function erc721() external view returns (IERC721);

    /**
     * @notice Get all staked license tokens for a validator
     * @param validationID The validation ID to query
     * @return Array of staked license token IDs
     */
    function getValidatorStakedLicenseTokens(
        bytes32 validationID
    ) external view returns (uint256[] memory);

    /**
     * @notice Get all staked license tokens for a delegator
     * @param delegationID The delegation ID to query
     * @return Array of staked license token IDs
     */
    function getDelegatorStakedLicenseTokens(
        bytes32 delegationID
    ) external view returns (uint256[] memory);

    /**
     * @notice Get all delegations for a validator
     * @param validationID The validation ID to query
     * @return Array of delegation IDs
     */
    function getValidatorDelegations(
        bytes32 validationID
    ) external view returns (bytes32[] memory);

    /**
     * @notice Get the validator that staked a specific license token
     * @param licenseTokenId The license token ID to query
     * @return The validation ID of the validator that staked the license token
     */
    function getLicenseTokenValidator(
        uint256 licenseTokenId
    ) external view returns (bytes32);

    /**
     * @notice Get the delegator that staked a specific license token
     * @param licenseTokenId The license token ID to query
     * @return The delegation ID of the delegator that staked the license token
     */
    function getLicenseTokenDelegator(
        uint256 licenseTokenId
    ) external view returns (bytes32);
}
