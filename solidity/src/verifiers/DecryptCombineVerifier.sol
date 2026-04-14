// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseDecryptCombineVerifier} from "./decryptcombine_vkey.sol";

contract DecryptCombineVerifier is BaseDecryptCombineVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"43d858ca8532f6b67880bc3a61cd9e3a39b4db1d4d59828d3e84452b7dc76a4f";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[7] memory decodedInput = abi.decode(input, (uint256[7]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
