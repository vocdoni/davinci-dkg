// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseFinalizeVerifier} from "./finalize_vkey.sol";

contract FinalizeVerifier is BaseFinalizeVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"252e7ff136ac4d02e9caac4aff0d8cc88d9b0d51aa53ee80cf7d857f52dfa8da";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[10] memory decodedInput = abi.decode(input, (uint256[10]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
