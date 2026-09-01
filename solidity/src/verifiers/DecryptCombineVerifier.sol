// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseDecryptCombineVerifier} from "./decryptcombine_vkey.sol";

/// @title DecryptCombineVerifier
/// @notice Mode-aware combine verifier wrapper. The public-input vector
///         carries the per-application correction fields: eid, aid, ctIdx,
///         mode, S, DeltaOrg.X, DeltaOrg.Y, threshold, shareCount,
///         combineHash, plaintextHash, challenge, transcriptCommitment.
contract DecryptCombineVerifier is BaseDecryptCombineVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"192977b7b1dd4f4acfe86cb9c81cefa38856e9500b2f2168742e609ab3f9bff7";

    error InvalidProofEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length != 32 * 8) revert InvalidProofEncoding();
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[13] memory decodedInput = abi.decode(input, (uint256[13]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
