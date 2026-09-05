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
        hex"bff0af27903c29e2e759e5b9653430e743cd8d56d24a2a4556fad311ec650a6a";

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
        uint256[8] memory decodedInput = abi.decode(input, (uint256[8]));
        this.verifyProof(proof, decodedInput);
    }
}
