// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {ERC20LicensedStakingManager} from "../src/ERC20LicensedStakingManager.sol";

import {IERC20LicensedStakingManager} from "../src/interfaces/IERC20LicensedStakingManager.sol";
import {ILicensedStakingManager} from "../src/interfaces/ILicensedStakingManager.sol";
import {
    MockERC721,
    MockRewardCalculator,
    MockValidatorManager,
    WEIGHT_TO_VALUE_FACTOR
} from "./Common.sol";
import "./Common.sol";
import {ExampleERC20} from "@mocks/ExampleERC20.sol";
import {ERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/ERC721.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {Validator, ValidatorStatus} from "@validator-manager/interfaces/IACP99Manager.sol";
import {ConversionData} from "@validator-manager/interfaces/IACP99Manager.sol";
import {IACP99Manager} from "@validator-manager/interfaces/IACP99Manager.sol";
import {IRewardCalculator} from "@validator-manager/interfaces/IRewardCalculator.sol";
import {IStakingManager} from "@validator-manager/interfaces/IStakingManager.sol";
import {StakingManagerSettings} from "@validator-manager/interfaces/IStakingManager.sol";
import {IValidatorManager} from "@validator-manager/interfaces/IValidatorManager.sol";
import {PChainOwner} from "@validator-manager/interfaces/IValidatorManager.sol";
import {Test} from "forge-std/Test.sol";

contract ERC20LicensedStakingManagerTest is Test {
    ERC20LicensedStakingManager public stakingManager;
    MockERC721 public licenseToken;
    ExampleERC20 public erc20Token;
    MockValidatorManager public validatorManager;
    MockRewardCalculator public rewardCalculator;

    uint256 public constant MAXIMUM_STAKE_AMOUNT = 100 ether;
    uint64 public constant MINIMUM_STAKE_DURATION = 1 days;
    uint16 public constant MINIMUM_DELEGATION_FEE_BIPS = 100;
    uint256 public constant DELEGATION_AMOUNT = 1e13;
    uint8 public constant MAXIMUM_STAKE_MULTIPLIER = 4;

    address constant DEFAULT_VALIDATOR_USER = address(0x1);
    address constant DEFAULT_DELEGATOR_USER = address(0x2);

    uint256[] DEFAULT_VALIDATOR_LICENSE_TOKENS = new uint256[](1);
    uint256[] DEFAULT_DELEGATOR_LICENSE_TOKENS = new uint256[](1);

    PChainOwner DEFAULT_PCHAIN_OWNER = PChainOwner({threshold: 1, addresses: new address[](1)});

    bytes DEFAULT_BLS_PUBLIC_KEY = bytes("BLS-");

    function setUp() public {
        // Deploy mock contracts
        licenseToken = new MockERC721();
        erc20Token = new ExampleERC20();
        validatorManager = new MockValidatorManager();
        rewardCalculator = new MockRewardCalculator();

        // Deploy staking manager
        stakingManager = new ERC20LicensedStakingManager(ICMInitializable.Allowed);

        // Initialize staking manager
        StakingManagerSettings memory settings = StakingManagerSettings({
            manager: IValidatorManager(address(validatorManager)),
            minimumStakeAmount: MINIMUM_STAKE_AMOUNT,
            maximumStakeAmount: MAXIMUM_STAKE_AMOUNT,
            minimumStakeDuration: MINIMUM_STAKE_DURATION,
            minimumDelegationFeeBips: MINIMUM_DELEGATION_FEE_BIPS,
            maximumStakeMultiplier: MAXIMUM_STAKE_MULTIPLIER,
            weightToValueFactor: WEIGHT_TO_VALUE_FACTOR,
            rewardCalculator: IRewardCalculator(address(rewardCalculator)),
            uptimeBlockchainID: bytes32(uint256(1))
        });

        stakingManager.initialize(
            settings, erc20Token, licenseToken, LICENSE_TO_STAKE_CONVERSION_FACTOR
        );
        DEFAULT_PCHAIN_OWNER.addresses[0] = DEFAULT_VALIDATOR_USER;

        // Mint and approve
        // Mint license token to validator user
        licenseToken.mint(DEFAULT_VALIDATOR_USER, 1);
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        licenseToken.approve(address(stakingManager), 1);
        DEFAULT_VALIDATOR_LICENSE_TOKENS[0] = 1;
        vm.stopPrank();

        // Mint license token to delegator user
        licenseToken.mint(DEFAULT_DELEGATOR_USER, 2);
        vm.startPrank(DEFAULT_DELEGATOR_USER);
        licenseToken.approve(address(stakingManager), 2);
        DEFAULT_DELEGATOR_LICENSE_TOKENS[0] = 2;
        vm.stopPrank();

        // Mint ERC20 tokens to users
        erc20Token.mint(DEFAULT_VALIDATOR_USER, MINIMUM_STAKE_AMOUNT);
        erc20Token.mint(DEFAULT_DELEGATOR_USER, DELEGATION_AMOUNT);

        // Approve ERC20 tokens
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        erc20Token.approve(address(stakingManager), MINIMUM_STAKE_AMOUNT);
        vm.stopPrank();

        vm.startPrank(DEFAULT_DELEGATOR_USER);
        erc20Token.approve(address(stakingManager), DELEGATION_AMOUNT);
        vm.stopPrank();
    }

    function test_Initialize() public view {
        assertEq(address(stakingManager.erc721()), address(licenseToken));
        assertEq(address(stakingManager.erc20()), address(erc20Token));
        assertEq(
            stakingManager.licenseToStakeConversionFactor(), LICENSE_TO_STAKE_CONVERSION_FACTOR
        );
    }

    function test_InitiateValidatorRegistration() public {
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        uint256 balanceBefore = erc20Token.balanceOf(DEFAULT_VALIDATOR_USER);
        vm.expectEmit(true, true, true, true);
        emit IStakingManager.InitiatedStakingValidatorRegistration(
            DEFAULT_VALIDATION_ID,
            DEFAULT_VALIDATOR_USER,
            MINIMUM_DELEGATION_FEE_BIPS,
            MINIMUM_STAKE_DURATION,
            DEFAULT_VALIDATOR_USER
        );
        vm.expectEmit(true, true, true, true);
        emit ILicensedStakingManager.ValidatorRegisteredWithLicenses(
            DEFAULT_VALIDATION_ID, DEFAULT_VALIDATOR_LICENSE_TOKENS
        );
        bytes32 validationID = stakingManager.initiateValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_PCHAIN_OWNER,
            DEFAULT_PCHAIN_OWNER,
            MINIMUM_DELEGATION_FEE_BIPS,
            MINIMUM_STAKE_DURATION,
            MINIMUM_STAKE_AMOUNT,
            DEFAULT_VALIDATOR_LICENSE_TOKENS,
            DEFAULT_VALIDATOR_USER
        );
        vm.stopPrank();

        uint256 balanceAfter = erc20Token.balanceOf(DEFAULT_VALIDATOR_USER);
        // Verify balance subtracted
        assertEq(balanceAfter, balanceBefore - MINIMUM_STAKE_AMOUNT);
        // Verify validation ID is set
        assertTrue(validationID == DEFAULT_VALIDATION_ID);
        // Verify license token is transferred
        assertEq(licenseToken.ownerOf(1), address(stakingManager));
        // Verify staking manager state
        assertEq(stakingManager.getValidatorStakedLicenseTokens(validationID).length, 1);
        assertEq(stakingManager.getValidatorStakedLicenseTokens(validationID)[0], 1);
        assertEq(stakingManager.getValidatorDelegations(validationID).length, 0);
        assertEq(stakingManager.getLicenseTokenValidator(1), validationID);
        assertEq(stakingManager.getLicenseTokenDelegator(1), bytes32(0));
    }

    function test_InitiateDelegatorRegistration() public {
        // First set up validator stake
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        stakingManager.initiateValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_PCHAIN_OWNER,
            DEFAULT_PCHAIN_OWNER,
            MINIMUM_DELEGATION_FEE_BIPS,
            MINIMUM_STAKE_DURATION,
            MINIMUM_STAKE_AMOUNT,
            DEFAULT_VALIDATOR_LICENSE_TOKENS,
            DEFAULT_VALIDATOR_USER
        );
        vm.stopPrank();

        uint256 totalDelegationAmount = DELEGATION_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR;
        uint256 newValidatorWeight = (MINIMUM_STAKE_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR)
            / WEIGHT_TO_VALUE_FACTOR + totalDelegationAmount / WEIGHT_TO_VALUE_FACTOR;

        // Then proceed with delegator registration
        vm.startPrank(DEFAULT_DELEGATOR_USER);
        uint256 balanceBefore = erc20Token.balanceOf(DEFAULT_DELEGATOR_USER);
        bytes32 defaultDelegationID = keccak256(abi.encodePacked(DEFAULT_VALIDATION_ID, uint64(0)));
        vm.expectEmit(true, true, true, true);
        emit IStakingManager.InitiatedDelegatorRegistration(
            defaultDelegationID,
            DEFAULT_VALIDATION_ID,
            DEFAULT_DELEGATOR_USER,
            uint64(0),
            uint64(newValidatorWeight),
            uint64(totalDelegationAmount / WEIGHT_TO_VALUE_FACTOR),
            bytes32(0),
            DEFAULT_DELEGATOR_USER
        );
        vm.expectEmit(true, true, true, true);
        emit ILicensedStakingManager.DelegatorRegisteredWithLicenses(
            defaultDelegationID, DEFAULT_VALIDATION_ID, DEFAULT_DELEGATOR_LICENSE_TOKENS
        );
        bytes32 delegationID = stakingManager.initiateDelegatorRegistration(
            DEFAULT_VALIDATION_ID,
            DELEGATION_AMOUNT,
            DEFAULT_DELEGATOR_LICENSE_TOKENS,
            DEFAULT_DELEGATOR_USER
        );
        vm.stopPrank();

        uint256 balanceAfter = erc20Token.balanceOf(DEFAULT_DELEGATOR_USER);
        // Verify balance subtracted
        assertEq(balanceAfter, balanceBefore - DELEGATION_AMOUNT);
        // Verify
        assertTrue(delegationID == defaultDelegationID);
        assertEq(licenseToken.ownerOf(2), address(stakingManager));
        assertEq(stakingManager.getValidatorDelegations(DEFAULT_VALIDATION_ID).length, 1);
        assertEq(
            stakingManager.getValidatorDelegations(DEFAULT_VALIDATION_ID)[0], defaultDelegationID
        );
        assertEq(stakingManager.getLicenseTokenDelegator(2), defaultDelegationID);
        assertEq(stakingManager.getDelegatorStakedLicenseTokens(defaultDelegationID).length, 1);
        assertEq(stakingManager.getDelegatorStakedLicenseTokens(defaultDelegationID)[0], 2);
    }

    function test_ValidatorRemoval() public {
        // First set up validator so staking manager owns tokens
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        stakingManager.initiateValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_PCHAIN_OWNER,
            DEFAULT_PCHAIN_OWNER,
            MINIMUM_DELEGATION_FEE_BIPS,
            MINIMUM_STAKE_DURATION,
            MINIMUM_STAKE_AMOUNT,
            DEFAULT_VALIDATOR_LICENSE_TOKENS,
            DEFAULT_VALIDATOR_USER
        );
        vm.stopPrank();

        uint256 balanceBefore = erc20Token.balanceOf(DEFAULT_VALIDATOR_USER);
        stakingManager.completeValidatorRemoval(0);
        // Verify license token is returned to the user
        assertEq(licenseToken.ownerOf(1), DEFAULT_VALIDATOR_USER);
        // Verify tokens have been unlocked
        assertEq(erc20Token.balanceOf(DEFAULT_VALIDATOR_USER), balanceBefore + MINIMUM_STAKE_AMOUNT);
    }

    function test_ValidatorReward() public {
        // First set up validator so staking manager owns tokens
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        stakingManager.initiateValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_PCHAIN_OWNER,
            DEFAULT_PCHAIN_OWNER,
            MINIMUM_DELEGATION_FEE_BIPS,
            MINIMUM_STAKE_DURATION,
            MINIMUM_STAKE_AMOUNT,
            DEFAULT_VALIDATOR_LICENSE_TOKENS,
            DEFAULT_VALIDATOR_USER
        );
        vm.stopPrank();

        // Skip min stake duration
        vm.warp(block.timestamp + MINIMUM_STAKE_DURATION);
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        stakingManager.initiateValidatorRemoval(DEFAULT_VALIDATION_ID, false, 0);
        vm.stopPrank();

        (address rewardRecipient, uint256 claimableRewards) =
            stakingManager.getValidatorRewardInfo(DEFAULT_VALIDATION_ID);
        assertEq(rewardRecipient, DEFAULT_VALIDATOR_USER);
        assertEq(
            claimableRewards, (MINIMUM_STAKE_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR) * 10 / 100
        );

        Validator memory validator = Validator({
            status: ValidatorStatus.Completed,
            nodeID: DEFAULT_NODE_ID,
            startingWeight: uint64(
                (MINIMUM_STAKE_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR) / WEIGHT_TO_VALUE_FACTOR
            ),
            sentNonce: 1,
            receivedNonce: 1,
            weight: uint64(
                (MINIMUM_STAKE_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR) / WEIGHT_TO_VALUE_FACTOR
            ),
            startTime: uint64(DEFAULT_BLOCK_TIMESTAMP),
            endTime: uint64(DEFAULT_BLOCK_TIMESTAMP + 365 days)
        });
        vm.mockCall(
            address(validatorManager),
            abi.encodeWithSelector(IACP99Manager.getValidator.selector),
            abi.encode(validator)
        );
        uint256 balanceBefore = erc20Token.balanceOf(DEFAULT_VALIDATOR_USER);
        vm.expectEmit(true, true, true, true);
        emit IStakingManager.ValidatorRewardClaimed(
            DEFAULT_VALIDATION_ID,
            DEFAULT_VALIDATOR_USER,
            (MINIMUM_STAKE_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR) * 10 / 100
        );
        stakingManager.completeValidatorRemoval(0);
        // Verify license token is returned to the user
        assertEq(licenseToken.ownerOf(1), DEFAULT_VALIDATOR_USER);
        // Verify tokens have been unlocked
        assertEq(
            erc20Token.balanceOf(DEFAULT_VALIDATOR_USER),
            balanceBefore + MINIMUM_STAKE_AMOUNT
                + (MINIMUM_STAKE_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR) * 10 / 100
        );
    }

    function test_CompleteDelegatorRemoval() public {
        // First set up validator so staking manager owns tokens
        vm.startPrank(DEFAULT_VALIDATOR_USER);
        stakingManager.initiateValidatorRegistration(
            DEFAULT_NODE_ID,
            DEFAULT_BLS_PUBLIC_KEY,
            DEFAULT_PCHAIN_OWNER,
            DEFAULT_PCHAIN_OWNER,
            MINIMUM_DELEGATION_FEE_BIPS,
            MINIMUM_STAKE_DURATION,
            MINIMUM_STAKE_AMOUNT,
            DEFAULT_VALIDATOR_LICENSE_TOKENS,
            DEFAULT_VALIDATOR_USER
        );
        vm.stopPrank();
        // Then proceed with delegator registration
        vm.startPrank(DEFAULT_DELEGATOR_USER);
        bytes32 delegationID = stakingManager.initiateDelegatorRegistration(
            DEFAULT_VALIDATION_ID,
            DELEGATION_AMOUNT,
            DEFAULT_DELEGATOR_LICENSE_TOKENS,
            DEFAULT_DELEGATOR_USER
        );
        vm.stopPrank();

        stakingManager.completeDelegatorRegistration(delegationID, 0);

        // skip min stake duration
        vm.warp(block.timestamp + MINIMUM_STAKE_DURATION);

        vm.startPrank(DEFAULT_DELEGATOR_USER);
        stakingManager.initiateDelegatorRemoval(delegationID, false, 0);

        uint256 balanceBefore = erc20Token.balanceOf(DEFAULT_DELEGATOR_USER);
        uint256 delegationRewards =
            (DELEGATION_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR) * 10 / 100;
        uint256 validatorFee = delegationRewards * MINIMUM_DELEGATION_FEE_BIPS / 10000; // 10000 == BIPS_CONVERSION_FACTOR
        uint256 expectedRewards = delegationRewards - validatorFee;
        vm.expectEmit(true, true, true, true);
        emit IERC20LicensedStakingManager.ERC20TokensUnlocked(
            DEFAULT_DELEGATOR_USER, DELEGATION_AMOUNT
        );
        stakingManager.completeDelegatorRemoval(delegationID, 0);
        assertEq(
            erc20Token.balanceOf(DEFAULT_DELEGATOR_USER),
            balanceBefore + DELEGATION_AMOUNT + expectedRewards
        );
        assertEq(licenseToken.ownerOf(2), DEFAULT_DELEGATOR_USER);
        vm.stopPrank();
    }
}
