// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

interface IZKVerifier {
    function verifyProof(bytes calldata proof, bytes calldata input) external view;
    function provingKeyHash() external pure returns (bytes32);
}
