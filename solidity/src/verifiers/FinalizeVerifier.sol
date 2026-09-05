// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseFinalizeVerifier} from "./finalize_vkey.sol";

/// @dev Wrapper over the generated finalize_vkey.sol; PROVING_KEY_HASH is
///      patched by `make circuits-update-hashes` (cmd/circuit-compile).
contract FinalizeVerifier is BaseFinalizeVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"786fddb4f3190c1807ceb80865f0575c7f2f5ecf4ac1db52c890a37d280cea96";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[7] memory decodedInput = abi.decode(input, (uint256[7]));
        this.verifyProof(proof, decodedInput);
    }
}
