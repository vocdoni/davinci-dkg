// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseContributionVerifier} from "./contribution_vkey.sol";

/// @notice IZKVerifier wrapper for the gnark-generated Groth16 contribution
///         verifier. Forwards uncompressed (8-word) proofs only; compressed
///         proofs are not used by the protocol's prover. The wrapper inherits
///         the base verifier so the dispatched call lands on the same code
///         instance (no extra address hop).
contract ContributionVerifier is BaseContributionVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"d747a935a6680c8b6446389ab994a7d939aa12b0def66c921f85fc984b1e69d9";

    error InvalidProofEncoding();
    error InvalidInputEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    /// @notice Verify a Groth16 proof against the contribution circuit's
    ///         public inputs. Reverts on mismatch or invalid proof.
    /// @dev    The base verifier's input arity (`uint256[N] input`) is
    ///         circuit-defined. The wrapper passes through whatever calldata
    ///         the caller supplies; length mismatches surface as the base
    ///         verifier's own revert.
    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length != 32 * 8) revert InvalidProofEncoding();
        if (input.length != 32 * 8) revert InvalidInputEncoding();
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[8] memory decodedInput = abi.decode(input, (uint256[8]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
