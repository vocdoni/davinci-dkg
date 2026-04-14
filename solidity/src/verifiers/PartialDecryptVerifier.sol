// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IZKVerifier} from "../interfaces/IZKVerifier.sol";
import {Verifier as BasePartialDecryptVerifier} from "./partialdecrypt_vkey.sol";

contract PartialDecryptVerifier is BasePartialDecryptVerifier, IZKVerifier {
    bytes32 internal constant PROVING_KEY_HASH =
        hex"3871aae077eee32665f8ee95220a1ef4ff279fc8c2387bcffedd1869b66466e8";

    error InvalidProofEncoding();
    error InvalidInputEncoding();

    function provingKeyHash() external pure returns (bytes32) {
        return PROVING_KEY_HASH;
    }

    function verifyProof(bytes calldata proof, bytes calldata input) external view {
        if (proof.length == 32 * 8) {
            if (input.length != 32 * 13) revert InvalidInputEncoding();
            _delegateStaticCall(
                abi.encodeWithSelector(
                    BasePartialDecryptVerifier.verifyProof.selector,
                    abi.decode(proof, (uint256[8])),
                    abi.decode(input, (uint256[13]))
                )
            );
            return;
        }
        if (proof.length == 32 * 4) {
            if (input.length != 32 * 13) revert InvalidInputEncoding();
            _delegateStaticCall(
                abi.encodeWithSelector(
                    BasePartialDecryptVerifier.verifyCompressedProof.selector,
                    abi.decode(proof, (uint256[4])),
                    abi.decode(input, (uint256[13]))
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
