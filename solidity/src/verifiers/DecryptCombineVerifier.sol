// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseDecryptCombineVerifier} from "./decryptcombine_vkey.sol";

/// @title DecryptCombineVerifier
/// @notice Combine verifier wrapper. The 9-word public-input vector is
///         eid, aid, ctIdx, threshold, shareCount, combineHash, plaintext,
///         challenge, transcriptCommitment; the ciphertext, the organizer key
///         and the partials travel in the BRLC-bound transcript.
contract DecryptCombineVerifier is BaseDecryptCombineVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"d70de162a56ac4801077f857f9491015916bb238db7fb9b0cfe5061eceae5305";

    error InvalidProofEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length != 32 * 8) revert InvalidProofEncoding();
        uint256[9] memory decodedInput = abi.decode(input, (uint256[9]));
        this.verifyProof(proof, decodedInput);
    }
}
