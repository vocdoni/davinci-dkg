// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../src/interfaces/IZKVerifier.sol";

/// @title  MockVerifier
/// @notice Stand-in for a Groth16 verifier wrapper while the real v4 proofs
///         do not exist yet: `verifyProof` accepts everything by default and
///         reverts `ProofRejected` once toggled off, so tests can exercise
///         both the happy path and the last-word verifier gate. Injected
///         through the DKGManager constructor in the test helpers.
contract MockVerifier is IZKVerifier {
    error ProofRejected();

    /// @dev Toggle: `true` (default) accepts every proof, `false` rejects.
    bool public accept = true;

    bytes32 public constant PROVING_KEY_HASH = keccak256("mock-verifier-proving-key");

    function setAccept(bool value) external {
        accept = value;
    }

    function verifyProof(bytes calldata proof, bytes calldata) external view override {
        if (!accept || proof.length == 0) revert ProofRejected();
    }

    function provingKeyHash() external pure override returns (bytes32) {
        return PROVING_KEY_HASH;
    }
}
