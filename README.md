# AC Staking Contracts

A collection of smart contracts for managing staking operations in the AvaCloud ecosystem, supporting both ERC20 and native token staking with license-based validation.

## Overview

This repository contains smart contracts that implement a licensed staking system for validators and delegators. The system requires both staking tokens (either ERC20 or native tokens) and license tokens (ERC721) to participate in the network's proof-of-stake mechanism.

## Key Features

- **Dual Token System**: Combines staking tokens (ERC20/native) with license tokens (ERC721)
- **Validator Management**: Register and manage validators with configurable parameters
- **Delegation System**: Allow token holders to delegate to validators
- **Reward Distribution**: Automated reward calculation and distribution
- **Configurable Parameters**: Flexible settings for stake amounts, durations, and fees

## Contract Architecture

### Main Contracts

1. **LicensedStakingManager**: Base contract implementing the core staking logic with license requirements
2. **ERC20LicensedStakingManager**: Implementation for ERC20 token staking
3. **NativeTokenLicensedStakingManager**: Implementation for native token staking

### Key Components

- **Validator Registration**: Process for registering validators with required stake and licenses
- **Delegator Registration**: Process for delegators to stake with validators
- **Reward System**: Calculation and distribution of staking rewards
- **Token Management**: Handling of both staking tokens and license tokens

### Modifications of existing contracts

Check [description](src/README.md) for scope of changes to original implementation of `StakingManager.sol`

## Configuration Parameters

- `LICENSE_TO_STAKE_CONVERSION_FACTOR`: Conversion factor for license tokens to stake amount

## Development

Utilize `Makefile` to build, test, lint...

## License

This code is licensed under a custom license. See the `LICENSE` file for details.

## Contact

For questions and support, please contact the AvaCloud team.