// (c) 2025, Ava Labs, Inc. All rights reserved.
// SPDX-License-Identifier: LicenseRef-AvaLabsProprietary

// This code is licensed under a custom license.
// See the file LICENSE for licensing terms.
pragma solidity 0.8.25;

import {ILicensedStakingManager} from "./ILicensedStakingManager.sol";
import {PChainOwner} from "@validator-manager/interfaces/IACP99Manager.sol";

/**
 * Proof of Stake Validator Manager that stakes native tokens
 * and requires ERC721 tokens as license for staking.
 */
interface INativeTokenLicensedStakingManager is ILicensedStakingManager {
    /**
     * @notice Begins the validator registration process. Locks the specified native tokens in the contract as the stake.
     * @param nodeID The ID of the node to add to the L1.
     * @param blsPublicKey The BLS public key of the validator.
     * @param remainingBalanceOwner The remaining balance owner of the validator.
     * @param disableOwner The disable owner of the validator.
     * @param delegationFeeBips The fee that delegators must pay to delegate to this validator.
     * @param minStakeDuration The minimum amount of time this validator must be staked for in seconds.
     * @param stakeAmount The amount of tokens to stake.
     * @param licenseTokenIds The IDs of the license tokens to stake.
     * @param rewardRecipient The address of the reward recipient.
     * @return validationID The ID of the registered validator.
     */
    function initiateValidatorRegistration(
        bytes memory nodeID,
        bytes memory blsPublicKey,
        PChainOwner memory remainingBalanceOwner,
        PChainOwner memory disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256 stakeAmount,
        uint256[] calldata licenseTokenIds,
        address rewardRecipient
    ) external returns (bytes32);

    /**
     * @notice Begins the delegator registration process. Locks the specified native tokens in the contract as the stake.
     * @param validationID The ID of the validator to stake to.
     * @param stakeAmount The amount of tokens to stake.
     * @param licenseTokenIds The IDs of the license tokens to stake.
     * @param rewardRecipient The address of the reward recipient.
     * @return delegationID The ID of the registered delegator.
     */
    function initiateDelegatorRegistration(
        bytes32 validationID,
        uint256 stakeAmount,
        uint256[] calldata licenseTokenIds,
        address rewardRecipient
    ) external returns (bytes32);
}
