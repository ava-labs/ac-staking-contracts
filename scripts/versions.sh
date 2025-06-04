#!/usr/bin/env bash
# Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.
# Based on https://github.com/ava-labs/icm-services/blob/main/scripts/versions.sh

set -e
set -o pipefail

REPO_CONTRACTS_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)
# Extract the Solidity version from foundry.toml
SOLIDITY_VERSION=$(awk -F"'" '/^solc_version/ {print $2}' $REPO_CONTRACTS_PATH/foundry.toml)
EVM_VERSION=$(awk -F"'" '/^evm_version/ {print $2}' $REPO_CONTRACTS_PATH/foundry.toml)

echo $SOLIDITY_VERSION
echo $EVM_VERSION