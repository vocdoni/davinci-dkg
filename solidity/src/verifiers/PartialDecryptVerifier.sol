// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BasePartialDecryptVerifier} from "./partialdecrypt_vkey.sol";

contract PartialDecryptVerifier is BasePartialDecryptVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"74e759dfc2477a128e4d830e7161f32d02f0cc8582b825d29a115c58201b7c0e";

    error InvalidProofEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length != 32 * 8) revert InvalidProofEncoding();
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[16] memory decodedInput = abi.decode(input, (uint256[16]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
