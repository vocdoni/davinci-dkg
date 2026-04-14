// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

library DKGIdLib {
    function getPrefix(uint32 chainId, address manager) internal pure returns (uint32) {
        return uint32(uint256(keccak256(abi.encodePacked(chainId, manager))));
    }

    function computeRoundId(uint32 prefix, uint64 nonce) internal pure returns (bytes12) {
        return bytes12((uint96(prefix) << 64) | uint96(nonce));
    }
}
