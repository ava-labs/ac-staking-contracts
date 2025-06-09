// SPDX-License-Identifier: MIT
pragma solidity =0.8.25;

import {ERC20LicensedStakingManager} from "../src/ERC20LicensedStakingManager.sol";
import {IERC721} from "@openzeppelin/contracts@5.0.2/token/ERC721/IERC721.sol";
import {ICMInitializable} from "@utilities/ICMInitializable.sol";
import {IERC20Mintable} from "@validator-manager/interfaces/IERC20Mintable.sol";
import {IRewardCalculator} from "@validator-manager/interfaces/IRewardCalculator.sol";
import {StakingManagerSettings} from "@validator-manager/interfaces/IStakingManager.sol";
import {IValidatorManager} from "@validator-manager/interfaces/IValidatorManager.sol";
import {Script} from "forge-std/Script.sol";

contract ERC20StakingManagerDeployScript is Script {
    uint256 public constant DEFAULT_LICENSE_TO_STAKE_CONVERSION_FACTOR = 1000000000000000000;
    uint256 public constant DEFAULT_MINIMUM_STAKE_AMOUNT = 1e16;
    uint256 public constant DEFAULT_MAXIMUM_STAKE_AMOUNT = 10e18;
    uint64 public constant DEFAULT_MINIMUM_STAKE_DURATION = 1;
    uint16 public constant DEFAULT_MINIMUM_DELEGATION_FEE_BIPS = 1;
    uint8 public constant DEFAULT_MAXIMUM_STAKE_MULTIPLIER = 4;
    uint256 public constant DEFAULT_WEIGHT_TO_VALUE_FACTOR = 1e12;

    string public constant ENV_VALIDATOR_MANAGER_ADDRESS = "validator_manager_address";
    string public constant ENV_REWARD_CALCULATOR_ADDRESS = "reward_calculator_address";
    string public constant ENV_LICENSE_TOKEN_ADDRESS = "license_token_address";
    string public constant ENV_UPTIME_BLOCKCHAIN_ID = "uptime_blockchain_id";
    string public constant ENV_ERC20_TOKEN_ADDRESS = "erc20_token_address";

    function run() public {
        vm.startBroadcast();
        ERC20LicensedStakingManager stakingManager =
            new ERC20LicensedStakingManager(ICMInitializable.Allowed);

        address validatorManager = vm.envAddress(ENV_VALIDATOR_MANAGER_ADDRESS);
        address rewardCalculator = vm.envAddress(ENV_REWARD_CALCULATOR_ADDRESS);
        address licenseToken = vm.envAddress(ENV_LICENSE_TOKEN_ADDRESS);
        address erc20Token = vm.envAddress(ENV_ERC20_TOKEN_ADDRESS);
        bytes32 uptimeBlockchainID = vm.envBytes32(ENV_UPTIME_BLOCKCHAIN_ID);

        StakingManagerSettings memory settings = StakingManagerSettings({
            manager: IValidatorManager(validatorManager),
            minimumStakeAmount: DEFAULT_MINIMUM_STAKE_AMOUNT,
            maximumStakeAmount: DEFAULT_MAXIMUM_STAKE_AMOUNT,
            minimumStakeDuration: DEFAULT_MINIMUM_STAKE_DURATION,
            minimumDelegationFeeBips: DEFAULT_MINIMUM_DELEGATION_FEE_BIPS,
            maximumStakeMultiplier: DEFAULT_MAXIMUM_STAKE_MULTIPLIER,
            weightToValueFactor: DEFAULT_WEIGHT_TO_VALUE_FACTOR,
            rewardCalculator: IRewardCalculator(rewardCalculator),
            uptimeBlockchainID: uptimeBlockchainID
        });

        stakingManager.initialize(
            settings,
            IERC20Mintable(erc20Token),
            IERC721(licenseToken),
            DEFAULT_LICENSE_TO_STAKE_CONVERSION_FACTOR
        );
        vm.stopBroadcast();
    }
}
