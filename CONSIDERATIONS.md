# Churn Restriction Issues

The `ValidatorManager` contract, called by `StakingManager`, includes a [churn limit mechanism](https://github.com/ava-labs/icm-contracts/blob/eb0afb7ff59ad4612fa3d5fd67d8b3e7e135bd06/contracts/validator-manager/ValidatorManager.sol#L728-L773)

In the [POS documentation](https://github.com/ava-labs/icm-contracts/blob/eb0afb7ff59ad4612fa3d5fd67d8b3e7e135bd06/contracts/validator-manager/README.md#pos-1) it is stated that *Unlike with PoA, PoS validators are not able to decrease their weight. This can lead to a scenario in which a PoS validator manager with a high proportion of the L1's weight is not able to exit the validator set due to churn restrictions. Additional validators or delegators will need to first be registered to more evenly distribute weight across the L1's validator set.* These churn limitations depend on the current total validator weight.

For `LicesendStakingManager` implementations the implication of the churn mechanism is higher. 
If L1 operator sets the `licenseToStakeConversionFactor` variable too high it could lead to inability to add new validators unless weight of the existing validator set is increased by delegations.

Over time, if the total validator weight falls below `(licenseToStakeConversionFactor × 100) / (weightToValueFactor × churnPercent)`, no new validators can be added to the system since even one NFT would require a minimum stake amount equal to `licenseToStakeConversionFactor`.

Mitigation of the issue is still possible by registering more delegators until adding a validator is possible again.

Example:

Assume `churnPercent` is `20%`, the current total weight is `1000e6`, `weightToValueFactor` is `1e12`, and `licenseToStakeConversionFactor` is `100e18`.

If the total weight drops below `500e6`, no new validator registrations can occur because staking a single NFT requires `100e18` (plus `1e12` for the ERC-20 token requirement), which exceeds the 20% churn limit of the current weight (`500e6 × 20% = 100e6`, i.e., `100e18` stake amount).

**IMPORTANT**: Consult with AvaCloud Client Services team to ensure configuration variables and initial weights are set correctly to minimize the risk of achieving this edge case.

**IMPORTANT**: Be aware that if the `_maximumStakeMultiplier` is set to 1, it may not allow new delegators for a particular validators.

# Meta Transactions

If the support for meta transactions is added in the future to the staking contracts, it is important that change also includes updating the [SafeERC20TransferFrom](https://github.com/ava-labs/icm-contracts/blob/eb0afb7ff59ad4612fa3d5fd67d8b3e7e135bd06/contracts/utilities/SafeERC20TransferFrom.sol#L31-L33) library in the `icm-contracts` repo as it relies on `msg.sender` and it is used in `ERC20LicensedStakingManager`.