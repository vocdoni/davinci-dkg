// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

/// @title  Lagrange interpolation over the BabyJubJub scalar field.
/// @notice Helper used by `reconstructSecret` to verify that a claimed
///         `sk = F(0)` is the unique Lagrange interpolation of the disclosed
///         shares. Without this check, the reveal-share circuit alone does
///         not bind the reconstructed secret (see SECURITY.md C-2).
library Lagrange {
    /// @dev BabyJubJub subgroup order r (prime). The DKG polynomial F lives
    /// in F_r, so all Lagrange arithmetic must be performed mod r.
    uint256 internal constant R =
        2736030358979909402780800718157159386076813972158567259200215660948447373041;

    error LagrangeMismatch();
    error InvalidShareCount();
    error DuplicateIndex();
    error ModExpFailed();

    /// @notice Compute sk = sum_i lambda_i * shares[i] mod r and require it to
    ///         equal `expectedSecret`. Reverts otherwise.
    /// @dev Uses Montgomery batch inversion: one `modexp` precompile call
    ///      regardless of `count`. The two scratch buffers `xs`/`shares` are
    ///      expected to hold exactly `count` valid entries.
    function verifyReconstruction(
        uint256[] memory xs,
        uint256[] memory shares,
        uint256 count,
        uint256 expectedSecret
    ) internal view {
        if (count == 0 || count > xs.length || count > shares.length) revert InvalidShareCount();

        uint256[] memory denom = new uint256[](count);
        uint256[] memory numer = new uint256[](count);

        // Build per-i numerator (∏_{j≠i} x_j) and denominator (∏_{j≠i} (x_j − x_i))
        // in F_r. Subtraction wraps via R + a − b to stay in [0, R).
        for (uint256 i = 0; i < count; i++) {
            uint256 xi = xs[i] % R;
            uint256 num = 1;
            uint256 den = 1;
            for (uint256 j = 0; j < count; j++) {
                if (i == j) continue;
                uint256 xj = xs[j] % R;
                if (xi == xj) revert DuplicateIndex();
                num = mulmod(num, xj, R);
                uint256 diff = addmod(xj, R - xi, R);
                den = mulmod(den, diff, R);
            }
            numer[i] = num;
            denom[i] = den;
        }

        // Montgomery batch inversion: one modexp for the joint denominator.
        uint256[] memory prefix = new uint256[](count);
        prefix[0] = denom[0];
        for (uint256 i = 1; i < count; i++) {
            prefix[i] = mulmod(prefix[i - 1], denom[i], R);
        }
        uint256 invJoint = _modInverse(prefix[count - 1]);

        uint256[] memory inv = new uint256[](count);
        // Walk backwards: inv[i] = prefix[i-1] · invJoint(i)
        // where invJoint(i) starts as invJoint(count-1) = (∏ d_j)^-1 and is
        // peeled by multiplying with d_{i+1} as we move left.
        for (uint256 k = count; k > 1; k--) {
            uint256 i = k - 1;
            inv[i] = mulmod(prefix[i - 1], invJoint, R);
            invJoint = mulmod(invJoint, denom[i], R);
        }
        inv[0] = invJoint;

        // sk = Σ_i (numer[i] / denom[i]) · share[i]   (all mod r)
        uint256 sk = 0;
        for (uint256 i = 0; i < count; i++) {
            uint256 lambda = mulmod(numer[i], inv[i], R);
            sk = addmod(sk, mulmod(lambda, shares[i] % R, R), R);
        }
        if (sk != expectedSecret % R) revert LagrangeMismatch();
    }

    /// @dev Compute base^(R-2) mod R via the EVM modexp precompile (0x05).
    function _modInverse(uint256 base) private view returns (uint256 result) {
        uint256 exponent = R - 2;
        uint256 modulus = R;
        assembly ("memory-safe") {
            let p := mload(0x40)
            mstore(p, 0x20)             // baseLen
            mstore(add(p, 0x20), 0x20)  // expLen
            mstore(add(p, 0x40), 0x20)  // modLen
            mstore(add(p, 0x60), base)
            mstore(add(p, 0x80), exponent)
            mstore(add(p, 0xa0), modulus)
            let success := staticcall(gas(), 0x05, p, 0xc0, p, 0x20)
            if iszero(success) {
                mstore(0x00, 0x9e44e6e0) // ModExpFailed()
                revert(0x1c, 0x04)
            }
            result := mload(p)
            mstore(0x40, add(p, 0xc0))
        }
    }
}
