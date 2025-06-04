// SPDX-License-Indetifier: UNLICENSED
pragma solidity =0.8.25;

contract ERC7201 {
    function getStorageAddress(
        string calldata namespace
    ) public pure returns (bytes32) {
        return keccak256(abi.encode(uint256(keccak256(abi.encodePacked(namespace))) - 1))
            & ~bytes32(uint256(0xff));
    }
}
