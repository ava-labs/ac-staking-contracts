#!/usr/bin/env bash
# Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

REPO_BASE_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

# Pass in the full name of the dependency.
# Parses go.mod for a matching entry and extracts the version number.
function getDepVersion() {
    grep -m1 "^\s*$1" $REPO_BASE_PATH/go.mod | cut -d ' ' -f2
}

function extract_commit() {
  local version=$1

  # Regex for a commit hash (assumed to be a 12+ character hex string)
  commit_hash_regex="-([0-9a-f]{12,})$"

  if [[ "$version" =~ $commit_hash_regex ]]; then
      # Extract the substring after the last '-'
      version=${BASH_REMATCH[1]}
  fi
  echo "$version"
}

# Don't export them as they're used in the context of other calls
SUBNET_EVM_VERSION=${SUBNET_EVM_VERSION:-$(extract_commit "$(getDepVersion github.com/ava-labs/subnet-evm)")}

export ARCH=$(uname -m)
[ $ARCH = x86_64 ] && ARCH=amd64
echo "ARCH set to $ARCH"

DEFAULT_CONTRACT_LIST="" # TODO: Add default contract list

CONTRACT_LIST=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -c | --contract) CONTRACT_LIST=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/abi_bindings.sh [OPTIONS]"
    echo "Build contracts and generate Go bindings"
    echo ""
    echo "Options:"
    echo "  -c, --contract <contract_name>          Generate Go bindings for the contract. If empty, generate Go bindings for a default list of contracts"
    echo "  -c, --contract "contract1 contract2"    Generate Go bindings for multiple contracts"
    echo "  -h, --help                              Print this help message"
    exit 0
fi

# Install abigen
echo "Building subnet-evm abigen"
go install github.com/ava-labs/subnet-evm/cmd/abigen@${SUBNET_EVM_VERSION}

# Solc does not recursively expand remappings, so we must construct them manually
remappings=$(cat $REPO_BASE_PATH/remappings.txt)

# Recursively search for all remappings.txt files in the lib directory.
# For each file, prepend the remapping with the relative path to the file.
while read -r filepath; do
    relative_path="${filepath#$REPO_BASE_PATH/}"
    dir_path=$(dirname "$relative_path")
    echo $dir_path
  
    # Use sed to transform each line with the relative path if it matches the @token=remapping pattern,
    # so that each remapping is of the form @token=lib/path/to/remapping
    transformed_lines=$(sed -n "s|^\(@[^=]*=\)\(.*\)|\1$dir_path/\2|p" "$filepath")
    remappings+=" $transformed_lines "
done < <(find "$REPO_BASE_PATH/lib" -type f -name "remappings.txt" )

function convertToLower() {
    if [ "$ARCH" = 'arm64' ]; then
        echo $1 | perl -ne 'print lc'
    else
        echo $1 | sed -e 's/\(.*\)/\L\1/'
    fi
}

# Removes a matching string from a comma-separated list
remove_matching_string() {
    input_list="$1"
    match="$2"
    # Split the input list by commas
    IFS=',' read -ra elements <<< "$input_list"
    
    # Initialize an empty result array
    result=()

    # Iterate over each element
    for element in "${elements[@]}"; do
        # Check if the part after the colon matches the given string
        if [[ "${element#*:}" != "$match" ]]; then
        # If it doesn't match, add the element to the result array
        result+=("$element")
        fi
    done

    # Join the result array with commas and print
    (IFS=','; echo "${result[*]}")
}

function generate_bindings() {
    local contract_names=("$@")
    for contract_name in "${contract_names[@]}"
    do
        path=$(find . -name $contract_name.sol)
        dir=$(dirname $path)
        dir="${dir#./}"

        mkdir -p $REPO_BASE_PATH/out/$contract_name.sol
        
        output_path=$REPO_BASE_PATH/out/$contract_name.sol
        forge_output=$output_path/$contract_name.json
        jq .abi "$forge_output" > "$output_path/$contract_name.abi"
        jq -r .bytecode.object "$forge_output" > "$output_path/$contract_name.bin"
        
        gen_path=$REPO_BASE_PATH/abi-bindings/go/$dir/$contract_name
        mkdir -p $gen_path
        echo "Generating Go bindings for $contract_name..."
        
        if [ -z "$filtered_contracts" ]; then
            $GOPATH/bin/abigen --pkg $(convertToLower $contract_name) \
                            --abi "$output_path/$contract_name.abi" \
                            --bin "$output_path/$contract_name.bin" \
                            --type $contract_name \
                            --out $gen_path/$contract_name.go
        else
            # Filter out external libraries
            for lib in $EXTERNAL_LIBS; do
                filtered_contracts=$(remove_matching_string $filtered_contracts $lib)
            done

            $GOPATH/bin/abigen --pkg $(convertToLower $contract_name) \
                            --abi=<(jq .abi $combined_json) \
                            --bin=<(jq -r .bytecode $combined_json) \
                            --type $contract_name \
                            --out $gen_path/$contract_name.go \
                            --exc $filtered_contracts
        fi
        
        echo "Done generating Go bindings for $contract_name."
    done
}

contract_names=($CONTRACT_LIST)

# If CONTRACT_LIST is empty, use DEFAULT_CONTRACT_LIST
if [[ -z "${CONTRACT_LIST}" ]]; then
    contract_names=($DEFAULT_CONTRACT_LIST)
fi

cd $REPO_BASE_PATH/src
generate_bindings "${contract_names[@]}"

exit 0