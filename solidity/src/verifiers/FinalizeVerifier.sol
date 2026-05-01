// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseFinalizeVerifier} from "./finalize_vkey.sol";

contract FinalizeVerifier is BaseFinalizeVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"0cf61bd16ef3d481b31b40d655f1e901429f88911b1abdc675cf23dadcf86a09";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[9] memory decodedInput = abi.decode(input, (uint256[9]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
