// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseDecryptCombineVerifier} from "./decryptcombine_vkey.sol";

/// @title DecryptCombineVerifier
/// @notice Combine verifier wrapper. The 11-word public-input vector is
///         eid, aid, ctIdx, DeltaOrg.X, DeltaOrg.Y, threshold, shareCount,
///         combineHash, plaintext, challenge, transcriptCommitment; the
///         organizer's DLEQ words travel in the BRLC-bound transcript.
contract DecryptCombineVerifier is BaseDecryptCombineVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"23b50690255b6580e58c4f76addf9359f378856e884fec3bd3cc1e2c2960ecb3";

    error InvalidProofEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length != 32 * 8) revert InvalidProofEncoding();
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[11] memory decodedInput = abi.decode(input, (uint256[11]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
