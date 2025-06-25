// (c) 2025, Ava Labs, Inc. All rights reserved.
// SPDX-License-Identifier: LicenseRef-AvaLabsProprietary

// This code is licensed under a custom license.
// See the file LICENSE for licensing terms.
pragma solidity 0.8.25;

import {LicensedStakingManager} from "./LicensedStakingManager.sol";
import {IERC20LicensedStakingManager} from "./interfaces/IERC20LicensedStakingManager.sol";
import {Initializable} from
    "@openzeppelin/contracts-upgradeable@5.0.2/proxy/utils/Initializable.sol";
import {SafeERC20} from "@openzeppelin/contracts@5.0.2/token/ERC20/utils/SafeERC20.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {SafeERC20TransferFrom} from "@utilities/SafeERC20TransferFrom.sol";
import {IERC20Mintable} from "@validator-manager/interfaces/IERC20Mintable.sol";
import {StakingManagerSettings} from "@validator-manager/interfaces/IStakingManager.sol";
import {PChainOwner} from "@validator-manager/interfaces/IValidatorManager.sol";

/**
 * @dev Implementation of the {IERC20LicensedStakingManager} interface.
 */
contract ERC20LicensedStakingManager is
    Initializable,
    LicensedStakingManager,
    IERC20LicensedStakingManager
{
    using SafeERC20 for IERC20Mintable;
    using SafeERC20TransferFrom for IERC20Mintable;

    // solhint-disable private-vars-leading-underscore
    /// @custom:storage-location erc7201:avacloud.storage.ERC20LicensedStakingManager
    struct ERC20LicensedStakingManagerStorage {
        IERC20Mintable _token;
    }
    // solhint-enable private-vars-leading-underscore

    // keccak256(abi.encode(uint256(keccak256("avacloud.storage.ERC20LicensedStakingManager")) - 1)) & ~bytes32(uint256(0xff));
    bytes32 public constant ERC20_LICENSED_STAKING_MANAGER_STORAGE_LOCATION =
        0x9d37ef67e865cad1eb988c62f5e45a5866d6dd4ddd905252e31276591c701c00;

    // solhint-disable ordering
    function _getERC20LicensedStakingManagerStorage()
        private
        pure
        returns (ERC20LicensedStakingManagerStorage storage $)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            $.slot := ERC20_LICENSED_STAKING_MANAGER_STORAGE_LOCATION
        }
    }

    constructor(
        ICMInitializable init
    ) {
        if (init == ICMInitializable.Disallowed) {
            _disableInitializers();
        }
    }

    /**
     * @notice Initialize the ERC20 token staking manager
     * @param settings Initial settings for the PoS validator manager
     * @param token The ERC20 token to be staked
     * @param licenseToken The ERC721 token license to be staked
     * @param licenseToStakeConversionFactor The conversion factor from license tokens to stake amount
     */
    function initialize(
        StakingManagerSettings calldata settings,
        IERC20Mintable token,
        IERC721 licenseToken,
        uint256 licenseToStakeConversionFactor
    ) external initializer {
        __ERC20LicensedStakingManager_init(
            settings, token, licenseToken, licenseToStakeConversionFactor
        );
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20LicensedStakingManager_init(
        StakingManagerSettings calldata settings,
        IERC20Mintable token,
        IERC721 licenseToken,
        uint256 licenseToStakeConversionFactor
    ) internal onlyInitializing {
        __LicensedStakingManager_init(settings, licenseToken, licenseToStakeConversionFactor);
        __ERC20LicensedStakingManager_init_unchained(token);
    }

    // solhint-disable-next-line func-name-mixedcase
    function __ERC20LicensedStakingManager_init_unchained(
        IERC20Mintable token
    ) internal onlyInitializing {
        ERC20LicensedStakingManagerStorage storage $ = _getERC20LicensedStakingManagerStorage();
        if (address(token) == address(0)) {
            revert ZeroAddress();
        }
        $._token = token;
    }

    /**
     * @notice See {IERC20LicensedStakingManager-initiateValidatorRegistration}
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
    ) external nonReentrant returns (bytes32) {
        return _initiateLicensedValidatorRegistration({
            nodeID: nodeID,
            blsPublicKey: blsPublicKey,
            remainingBalanceOwner: remainingBalanceOwner,
            disableOwner: disableOwner,
            delegationFeeBips: delegationFeeBips,
            minStakeDuration: minStakeDuration,
            tokenStakeAmount: stakeAmount,
            licenseTokenIds: licenseTokenIds,
            rewardRecipient: rewardRecipient
        });
    }

    /**
     * @notice See {IERC20LicenseStakingManager-initiateDelegatorRegistration}
     */
    function initiateDelegatorRegistration(
        bytes32 validationID,
        uint256 delegationAmount,
        uint256[] calldata licenseTokenIds,
        address rewardRecipient
    ) external nonReentrant returns (bytes32) {
        return _initiateLicensedDelegatorRegistration(
            validationID, delegationAmount, licenseTokenIds, rewardRecipient
        );
    }

    /**
     * @notice Returns the ERC20 token being staked
     * @return The ERC20 token being staked
     */
    function erc20() external view returns (IERC20Mintable) {
        return _getERC20LicensedStakingManagerStorage()._token;
    }

    /**
     * @notice See {LicensedStakingManager-_lockTokens}
     * Note: Must be guarded with reentrancy guard for safe transfer from.
     */
    function _lockTokens(
        uint256 value
    ) internal virtual override returns (uint256) {
        _getERC20LicensedStakingManagerStorage()._token.safeTransferFrom(value);
        return value;
    }

    /**
     * @notice See {LicensedStakingManager-_unlock}
     * Note: Must be guarded with reentrancy guard for safe transfer.
     */
    function _unlock(address to, uint256 value, bytes32 stakeId) internal virtual override {
        uint256 tokenAmount = _tokenAmountFromStakeAmount(stakeId, value);
        LicensedStakingManager._unlock(to, value, stakeId);
        _getERC20LicensedStakingManagerStorage()._token.safeTransfer(to, tokenAmount);
        emit ERC20TokensUnlocked(to, tokenAmount);
    }

    /**
     * @notice See {StakingManager-_reward}
     */
    function _reward(address account, uint256 amount) internal virtual override {
        ERC20LicensedStakingManagerStorage storage $ = _getERC20LicensedStakingManagerStorage();
        $._token.mint(account, amount);
    }
}
