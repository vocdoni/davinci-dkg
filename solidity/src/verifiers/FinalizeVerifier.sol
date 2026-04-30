// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseFinalizeVerifier} from "./finalize_vkey.sol";

contract FinalizeVerifier is BaseFinalizeVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"b7d50091bbf5d4629605846b345d0113903db148242238c55d039eb0c88e51d9";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[9] memory decodedInput = abi.decode(input, (uint256[9]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
