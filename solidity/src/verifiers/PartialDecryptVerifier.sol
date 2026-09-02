// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BasePartialDecryptVerifier} from "./partialdecrypt_vkey.sol";

contract PartialDecryptVerifier is BasePartialDecryptVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"8d5cb7d575354215d07b8b2721e1a5465424a85f78ca2d0beb56041efd2b0820";

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
