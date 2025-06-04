## StakingManager

* This contract is a modified version of the Avalanche ICM StakingManager contract.
* Version that was modified is from release validator-manager-v2.0.0
* Original code is from https://github.com/ava-labs/icm-contracts/blob/validator-manager-v2.0.0/contracts/validator-manager/StakingManager.sol

### Modifications

* Function signature changed:

```solidity
// Previous signature
function _unlock(address to, uint256 value) internal virtual;
// New Signature
function _unlock(address to, uint256 value, bytes32 stakeId) internal virtual;
```

* invocations of _unlock changed to include stakeID parameter
```solidity
// in `completeValidatorRemoval`
 _unlock(owner, weightToValue(validator.startingWeight), validationID);
// in `completeDelegatorRemoval`
_unlock(delegator.owner, weightToValue(delegator.weight), delegationID);
```