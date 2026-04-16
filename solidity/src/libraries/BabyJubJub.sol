// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

/// @title BabyJubJub
/// @notice Twisted Edwards elliptic curve arithmetic over the BN254 scalar field.
///         Adapted from https://github.com/yondonfu/sol-baby-jubjub (MIT)
///         and updated to Solidity 0.8.x.
///
/// Curve equation: a·x² + y² = 1 + d·x²·y²
///   a = 168700 (= 0x292FC)
///   d = 168696 (= 0x292F8)
///   Q = 21888242871839275222246405745257275088548364400416034343698204186575808495617
///       (BN254 scalar field prime)
///
/// Identity element: (0, 1)
library BabyJubJub {
    /// @dev Curve parameter a = 168700
    uint256 internal constant A = 168700;
    /// @dev Curve parameter d = 168696
    uint256 internal constant D = 168696;
    /// @dev BN254 scalar field prime Q
    uint256 internal constant Q =
        21888242871839275222246405745257275088548364400416034343698204186575808495617;

    /// @notice Add two points on the Baby JubJub curve.
    ///         Implements the unified twisted Edwards addition formula:
    ///           x3 = (x1·y2 + y1·x2) / (1 + d·x1·x2·y1·y2)
    ///           y3 = (y1·y2 - a·x1·x2) / (1 - d·x1·x2·y1·y2)
    /// @dev The formula is complete (works for all points including the identity).
    ///      The identity element is (0, 1).
    /// @param _x1 x-coordinate of point P1
    /// @param _y1 y-coordinate of point P1
    /// @param _x2 x-coordinate of point P2
    /// @param _y2 y-coordinate of point P2
    /// @return x3 x-coordinate of P1 + P2
    /// @return y3 y-coordinate of P1 + P2
    function pointAdd(
        uint256 _x1,
        uint256 _y1,
        uint256 _x2,
        uint256 _y2
    ) internal view returns (uint256 x3, uint256 y3) {
        // Identity shortcuts: (0,1) is the neutral element
        if (_x1 == 0 && _y1 == 1) {
            return (_x2, _y2);
        }
        if (_x2 == 0 && _y2 == 1) {
            return (_x1, _y1);
        }

        uint256 x1x2 = mulmod(_x1, _x2, Q);
        uint256 y1y2 = mulmod(_y1, _y2, Q);
        uint256 dx1x2y1y2 = mulmod(D, mulmod(x1x2, y1y2, Q), Q);

        uint256 x3Num = addmod(mulmod(_x1, _y2, Q), mulmod(_y1, _x2, Q), Q);
        uint256 y3Num = _submod(y1y2, mulmod(A, x1x2, Q));

        x3 = mulmod(x3Num, _inverse(addmod(1, dx1x2y1y2, Q)), Q);
        y3 = mulmod(y3Num, _inverse(_submod(1, dx1x2y1y2)), Q);
    }

    // ─── Private helpers ───────────────────────────────────────────────────────

    /// @dev Modular subtraction: (a - b) mod Q, handling underflow.
    function _submod(uint256 _a, uint256 _b) private pure returns (uint256) {
        return addmod(_a, Q - (_b % Q), Q);
    }

    /// @dev Q - 2, used as the Fermat exponent for modular inverse (a^(Q-2) mod Q).
    ///      Stored as a constant so the assembly block can reference it via mload.
    uint256 private constant Q_MINUS_2 =
        21888242871839275222246405745257275088548364400416034343698204186575808495615;

    /// @dev Modular multiplicative inverse via Fermat's little theorem.
    ///      Returns a^(Q-2) mod Q using the bigModExp precompile (0x05).
    function _inverse(uint256 _a) private view returns (uint256 o) {
        uint256 exponent = Q_MINUS_2;
        uint256 modulus = Q;
        assembly {
            let memPtr := mload(0x40)
            mstore(memPtr, 0x20)               // length of base
            mstore(add(memPtr, 0x20), 0x20)    // length of exponent
            mstore(add(memPtr, 0x40), 0x20)    // length of modulus
            mstore(add(memPtr, 0x60), _a)      // base
            mstore(add(memPtr, 0x80), exponent) // exponent = Q - 2
            mstore(add(memPtr, 0xa0), modulus)  // modulus = Q
            // bigModExp precompile at 0x05
            let success := staticcall(gas(), 0x05, memPtr, 0xc0, memPtr, 0x20)
            if iszero(success) { revert(0, 0) }
            o := mload(memPtr)
        }
    }
}
