#!/usr/bin/env bash

# Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

# Enable constructor-args
set -e

REPO_BASE_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $REPO_BASE_PATH/.env

forge script script/NativeTokenStakingManagerDeploy.s.sol:NativeTokenStakingManagerDeployScript \
  --rpc-url $rpc_url \
  --private-key $user_private_key \
  --broadcast \
  -vvvv
