// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseRevealShareVerifier} from "./revealshare_vkey.sol";

contract RevealShareVerifier is BaseRevealShareVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"4e1ae67d1bb1e750c48fa01e0f40783926027cd66180d9e72f8d9eed6d99cd36";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[8] memory decodedProof = abi.decode(proof, (uint256[8]));
        uint256[7] memory decodedInput = abi.decode(input, (uint256[7]));
        this.verifyProof(decodedProof, decodedInput);
    }
}
