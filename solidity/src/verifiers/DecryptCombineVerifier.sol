// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseDecryptCombineVerifier} from "./decryptcombine_vkey.sol";

/// @title DecryptCombineVerifier
/// @notice Mode-aware combine verifier wrapper. Public-input width bumped to
///         13 in P5/P6 to carry the per-application correction fields:
///         eid, aid, ctIdx, mode, S, DeltaOrg.X, DeltaOrg.Y, threshold,
///         shareCount, combineHash, plaintextHash, challenge, transcriptCommitment.
contract DecryptCombineVerifier is BaseDecryptCombineVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"d1cd294544102bb194baa6a3255de8e23899df4c1f489c4d45c23b84bb44af43";

    error InvalidProofEncoding();
    error InvalidInputEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length == 32 * 8) {
            if (input.length != 32 * 13) revert InvalidInputEncoding();
            _delegateStaticCall(
                abi.encodeWithSelector(
                    BaseDecryptCombineVerifier.verifyProof.selector,
                    abi.decode(proof, (uint256[8])),
                    abi.decode(input, (uint256[13]))
                )
            );
            return;
        }
        if (proof.length == 32 * 4) {
            if (input.length != 32 * 13) revert InvalidInputEncoding();
            _delegateStaticCall(
                abi.encodeWithSelector(
                    BaseDecryptCombineVerifier.verifyCompressedProof.selector,
                    abi.decode(proof, (uint256[4])),
                    abi.decode(input, (uint256[13]))
                )
            );
            return;
        }
        revert InvalidProofEncoding();
    }

    function _delegateStaticCall(bytes memory payload) internal view {
        (bool ok, bytes memory data) = address(this).staticcall(payload);
        if (!ok) {
            assembly ("memory-safe") {
                revert(add(data, 0x20), mload(data))
            }
        }
    }
}
