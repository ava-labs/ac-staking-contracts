#!/usr/bin/env bash

set -e
REPO_CONTRACTS_PATH=$(
    cd "$(dirname "${BASH_SOURCE[0]}")"
    cd .. && pwd
)

source $REPO_CONTRACTS_PATH/scripts/constants.sh
source $REPO_CONTRACTS_PATH/.env


user_address=$(cast wallet address --private-key $user_private_key)
erc7201_address=$(cast compute-address $user_address --rpc-url $rpc_url)
erc7201_address=${erc7201_address:18}


echo "Deploying ERC7201 contract to $erc7201_address"
forge create src/utils/ERC7201.sol:ERC7201 --rpc-url $rpc_url --private-key $user_private_key --broadcast
echo "Deployed ERC7201 contract to $erc7201_address"

echo "Storage location for ERC20LicensedStakingManagerStorage: '$ERC20_LICENSED_STAKING_MANAGER_STORAGE'"
erc20_licensed_staking_manager_storage_location=$(cast call $erc7201_address "getStorageAddress(string)(bytes32)" $ERC20_LICENSED_STAKING_MANAGER_STORAGE --rpc-url $rpc_url)
echo "Storage location for ERC20LicensedStakingManagerStorage: '$erc20_licensed_staking_manager_storage_location'"

echo "Storage location for LicensedStakingManagerStorage: '$LICENSED_STAKING_MANAGER_STORAGE'"
licensed_staking_manager_storage_location=$(cast call $erc7201_address "getStorageAddress(string)(bytes32)" $LICENSED_STAKING_MANAGER_STORAGE --rpc-url $rpc_url)
echo "Storage location for LicensedStakingManagerStorage: '$licensed_staking_manager_storage_location'"