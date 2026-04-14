// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

library BRLC {
    uint256 internal constant FR_MODULUS =
        21888242871839275222246405745257275088548364400416034343698204186575808495617;

    function deriveChallenge(bytes12 roundId, bytes32 domain, bytes32 anchor) internal pure returns (uint256) {
        return uint256(keccak256(abi.encodePacked(roundId, domain, anchor))) % FR_MODULUS;
    }

    function commit(uint256 challenge, uint256[] memory values) internal pure returns (uint256 acc) {
        uint256 rho = challenge % FR_MODULUS;
        uint256 power = rho;

        for (uint256 i = 0; i < values.length; i++) {
            acc = addmod(acc, mulmod(power, values[i], FR_MODULUS), FR_MODULUS);
            power = mulmod(power, rho, FR_MODULUS);
        }
    }

    /// @dev Streams the BRLC commitment over a contiguous calldata region of `count`
    /// 32-byte words starting at `dataOffset` (a calldata byte offset). Uses calldataload,
    /// so no memory is allocated and no abi.decode copy is needed.
    function commitCalldata(uint256 challenge, uint256 dataOffset, uint256 count)
        internal
        pure
        returns (uint256 acc)
    {
        uint256 m = FR_MODULUS;
        uint256 rho = challenge % m;
        assembly ("memory-safe") {
            let power := rho
            let end := add(dataOffset, mul(count, 0x20))
            for { let p := dataOffset } lt(p, end) { p := add(p, 0x20) } {
                acc := addmod(acc, mulmod(power, calldataload(p), m), m)
                power := mulmod(power, rho, m)
            }
        }
    }

    /// @dev Streams the BRLC commitment over four contiguous memory regions in sequence
    /// without copying them into a dynamic vector. `lenN` is the number of 32-byte words.
    function commit4(
        uint256 challenge,
        uint256 ptr0, uint256 len0,
        uint256 ptr1, uint256 len1,
        uint256 ptr2, uint256 len2,
        uint256 ptr3, uint256 len3
    ) internal pure returns (uint256 acc) {
        uint256 m = FR_MODULUS;
        uint256 rho = challenge % m;
        assembly ("memory-safe") {
            let power := rho
            function streamRegion(a, p, ptr, count, modulus, r) -> newAcc, newPower {
                newAcc := a
                newPower := p
                let end := add(ptr, mul(count, 0x20))
                for { let q := ptr } lt(q, end) { q := add(q, 0x20) } {
                    newAcc := addmod(newAcc, mulmod(newPower, mload(q), modulus), modulus)
                    newPower := mulmod(newPower, r, modulus)
                }
            }
            acc, power := streamRegion(acc, power, ptr0, len0, m, rho)
            acc, power := streamRegion(acc, power, ptr1, len1, m, rho)
            acc, power := streamRegion(acc, power, ptr2, len2, m, rho)
            acc, power := streamRegion(acc, power, ptr3, len3, m, rho)
        }
    }
}
