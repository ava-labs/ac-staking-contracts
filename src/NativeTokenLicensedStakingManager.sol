// (c) 2025, Ava Labs, Inc. All rights reserved.
// SPDX-License-Identifier: LicenseRef-AvaLabsProprietary

// This code is licensed under a custom license.
// See the file LICENSE for licensing terms.
pragma solidity 0.8.25;

import {LicensedStakingManager} from "./LicensedStakingManager.sol";
import {INativeTokenLicensedStakingManager} from
    "./interfaces/INativeTokenLicensedStakingManager.sol";
import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/INativeMinter.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {Address} from "@openzeppelin/contracts@5.0.2/utils/Address.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {StakingManagerSettings} from "@validator-manager/interfaces/IStakingManager.sol";
import {PChainOwner} from "@validator-manager/interfaces/IValidatorManager.sol";

/**
 * @dev Implementation of the {INativeTokenLicensedStakingManager} interface.
 */
contract NativeTokenLicensedStakingManager is
    Initializable,
    LicensedStakingManager,
    INativeTokenLicensedStakingManager
{
    using Address for address payable;

    INativeMinter public constant NATIVE_MINTER =
        INativeMinter(0x0200000000000000000000000000000000000001);

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initialize the native token licensed staking manager
     * @param settings Initial settings for the PoS validator manager
     * @param licenseToken The ERC721 token license to be staked
     * @param licenseToStakeConversionFactor The conversion factor from license tokens to stake amount
     */
    //solhint-disable ordering
    function initialize(
        StakingManagerSettings calldata settings,
        IERC721 licenseToken,
        uint256 licenseToStakeConversionFactor
    ) external initializer {
        __NativeTokenLicensedStakingManager_init(
            settings, licenseToken, licenseToStakeConversionFactor
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __NativeTokenLicensedStakingManager_init(
        StakingManagerSettings calldata settings,
        IERC721 licenseToken,
        uint256 licenseToStakeConversionFactor
    ) internal onlyInitializing {
        __LicensedStakingManager_init(settings, licenseToken, licenseToStakeConversionFactor);
        __NativeTokenLicensedStakingManager_init_unchained();
    }

    // solhint-disable-next-line func-name-mixedcase, no-empty-blocks
    function __NativeTokenLicensedStakingManager_init_unchained() internal onlyInitializing {}

    /**
     * @notice See {INativeTokenLicensedStakingManager-initiateValidatorRegistration}
     */
    function initiateValidatorRegistration(
        bytes calldata nodeID,
        bytes calldata blsPublicKey,
        PChainOwner calldata remainingBalanceOwner,
        PChainOwner calldata disableOwner,
        uint16 delegationFeeBips,
        uint64 minStakeDuration,
        uint256[] calldata licenseTokenIds,
        address rewardRecipient
    ) external payable nonReentrant returns (bytes32) {
        return _initiateLicensedValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            tokenStakeAmount: msg.value,
            licenseTokenIds: licenseTokenIds,
            rewardRecipient: rewardRecipient
        });
    }

    /**
     * @notice See {INativeTokenLicensedStakingManager-initiateDelegatorRegistration}
     */
    function initiateDelegatorRegistration(
        bytes32 validationID,
        uint256[] calldata licenseTokenIds,
        address rewardRecipient
    ) external payable nonReentrant returns (bytes32) {
        return _initiateLicensedDelegatorRegistration(
            validationID, msg.value, licenseTokenIds, rewardRecipient
        );
    }

    /**
     * @notice see {LicensedStakingManager-_lockTokens}
     */
    function _lockTokens(
        uint256 value
    ) internal virtual override returns (uint256) {
        return value;
    }

    /**
     * @notice See {LicensedStakingManager-_unlock}
     * Note: Must be guarded with reentrancy guard for safe transfer.
     */
    function _unlock(address to, uint256 value, bytes32 stakeId) internal virtual override {
        uint256 tokenAmount = _tokenAmountFromStakeAmount(stakeId, value);
        LicensedStakingManager._unlock(to, value, stakeId);
        payable(to).sendValue(tokenAmount);
        emit NativeTokensUnlocked(to, tokenAmount);
    }

    /**
     * @notice See {StakingManager-_reward}
     */
    function _reward(address account, uint256 amount) internal virtual override {
        NATIVE_MINTER.mintNativeCoin(account, amount);
    }
}
