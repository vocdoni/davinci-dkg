// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseRevealSubmitVerifier} from "./revealsubmit_vkey.sol";

contract RevealSubmitVerifier is BaseRevealSubmitVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"413a4850abadf1894fc17fc886857c9f6df7395e37db93c75310b7cdc3860186";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[5] memory decodedInput = abi.decode(input, (uint256[5]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
