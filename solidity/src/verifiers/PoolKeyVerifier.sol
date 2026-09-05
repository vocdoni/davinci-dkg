// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BasePoolKeyVerifier} from "./poolkey_vkey.sol";

contract PoolKeyVerifier is BasePoolKeyVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"190f9007ad878e9b8d8ef68770c984d3ada5d6ad32e4945393993317cb20b8d5";

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        uint256[8] memory decodedInput = abi.decode(input, (uint256[8]));
        this.verifyProof(proof, decodedInput);
    }
}
