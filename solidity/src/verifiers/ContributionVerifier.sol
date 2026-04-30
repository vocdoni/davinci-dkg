// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BaseContributionVerifier} from "./contribution_vkey.sol";

contract ContributionVerifier is BaseContributionVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"5c77f8f911a3b7933667cc03856a40146de22b1461006e333f6e623b350bb48c";

    error InvalidProofEncoding();
    error InvalidInputEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length == 32 * 8) {
            if (input.length != 32 * 10) revert InvalidInputEncoding();
            _delegateStaticCall(
                abi.encodeWithSelector(
                    BaseContributionVerifier.verifyProof.selector,
                    abi.decode(proof, (uint256[8])),
                    abi.decode(input, (uint256[10]))
                )
            );
            return;
        }
        if (proof.length == 32 * 4) {
            if (input.length != 32 * 10) revert InvalidInputEncoding();
            _delegateStaticCall(
                abi.encodeWithSelector(
                    BaseContributionVerifier.verifyCompressedProof.selector,
                    abi.decode(proof, (uint256[4])),
                    abi.decode(input, (uint256[10]))
                )
            );
            return;
        }
        revert InvalidProofEncoding();
    }

    function _delegateStaticCall(bytes memory payload) internal view {
        (bool ok, bytes memory data) = address(this).staticcall(payload);
        if (!ok) {
            assembly ("memory-safe") {
                revert(add(data, 0x20), mload(data))
            }
        }
    }
}
