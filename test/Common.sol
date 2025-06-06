// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {INativeMinter} from
    "@avalabs/subnet-evm-contracts@1.2.2/contracts/interfaces/INativeMinter.sol";
import {ERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/ERC721.sol";
import {Validator, ValidatorStatus} from "@validator-manager/interfaces/IACP99Manager.sol";
import {ConversionData} from "@validator-manager/interfaces/IACP99Manager.sol";
import {IValidatorManager} from "@validator-manager/interfaces/IValidatorManager.sol";
import {PChainOwner} from "@validator-manager/interfaces/IValidatorManager.sol";
import {Test} from "forge-std/Test.sol";

uint64 constant DEFAULT_BLOCK_TIMESTAMP = 1717660800;
uint256 constant MINIMUM_STAKE_AMOUNT = 1 ether;
bytes constant DEFAULT_NODE_ID = bytes("NodeID-");
uint256 constant WEIGHT_TO_VALUE_FACTOR = 1e12;
uint256 constant LICENSE_TO_STAKE_CONVERSION_FACTOR = 1 ether;

contract MockERC721 is ERC721 {
    constructor() ERC721("Mock License Token", "MLT") {}

    function mint(address to, uint256 tokenId) external {
        _safeMint(to, tokenId);
    }
}

event MockNativeCoinMinted(address indexed to, uint256 amount);

bytes32 constant DEFAULT_VALIDATION_ID = bytes32(uint256(100));

contract MockNativeMinter is INativeMinter, Test {
    function mintNativeCoin(address to, uint256 amount) external override {
        deal(to, amount);
        emit MockNativeCoinMinted(to, amount);
    }

    function readAllowList(
        address
    ) external pure returns (uint256) {
        return 0;
    }

    function setAdmin(
        address addr
    ) external {}

    function setEnabled(
        address addr
    ) external {}

    function setManager(
        address addr
    ) external {}

    function setNone(
        address addr
    ) external {}
}

contract MockValidatorManager is IValidatorManager {
    function getValidator(
        bytes32
    ) external pure override returns (Validator memory) {
        uint256 weight =
            (MINIMUM_STAKE_AMOUNT + LICENSE_TO_STAKE_CONVERSION_FACTOR) / WEIGHT_TO_VALUE_FACTOR;
        return Validator({
            status: ValidatorStatus.Active,
            nodeID: DEFAULT_NODE_ID,
            startingWeight: uint64(weight),
            sentNonce: 1,
            receivedNonce: 1,
            weight: uint64(weight),
            startTime: uint64(DEFAULT_BLOCK_TIMESTAMP),
            endTime: uint64(DEFAULT_BLOCK_TIMESTAMP + 365 days)
        });
    }

    function getChurnPeriodSeconds() external pure override returns (uint64) {
        return 60 seconds;
    }

    // Required interface functions with minimal implementations for testing
    function initializeValidatorSet(ConversionData calldata, uint32) external pure override {}

    function completeValidatorRegistration(
        uint32
    ) external pure override returns (bytes32) {
        return bytes32(0);
    }

    function completeValidatorRemoval(
        uint32
    ) external pure override returns (bytes32) {
        return DEFAULT_VALIDATION_ID;
    }

    function completeValidatorWeightUpdate(
        uint32
    ) external pure override returns (bytes32, uint64) {
        return (bytes32(0), 0);
    }

    function subnetID() external pure override returns (bytes32) {
        return bytes32(0);
    }

    function l1TotalWeight() external pure override returns (uint64) {
        return 0;
    }

    function migrateFromV1(bytes32, uint32) external pure override {}

    function initiateValidatorRegistration(
        bytes memory,
        bytes memory,
        PChainOwner memory,
        PChainOwner memory,
        uint64
    ) external pure override returns (bytes32) {
        return DEFAULT_VALIDATION_ID;
    }

    function resendRegisterValidatorMessage(
        bytes32
    ) external pure override {}
    function initiateValidatorRemoval(
        bytes32
    ) external pure override {}
    function resendValidatorRemovalMessage(
        bytes32
    ) external pure override {}

    function initiateValidatorWeightUpdate(
        bytes32,
        uint64
    ) external pure override returns (uint64, bytes32) {
        return (0, bytes32(0));
    }

    function registeredValidators(
        bytes calldata
    ) external pure override returns (bytes32) {
        return bytes32(0);
    }
}

contract MockRewardCalculator {
    function calculateReward(
        uint256 stakeAmount,
        uint64,
        uint64,
        uint64,
        uint64
    ) external pure returns (uint256) {
        return stakeAmount * 10 / 100; // 10% reward
    }
}
