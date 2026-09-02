// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DKGProtocol} from "../src/libraries/DKGProtocol.sol";

/// @title DKGProtocolTest
/// @notice Cross-impl byte-equality assertions against the canonical protocol
///         vectors generated from `internal/protocol/protocol.go`. The hex
///         strings below are copied verbatim from `tests/vectors/protocol.json`;
///         updating the vector is the canonical way to evolve the protocol —
///         all three layers (Solidity, Go, TS) MUST agree on every value.
contract DKGProtocolTest is Test {
    function test_AppMode_Constants() public pure {
        assertEq(uint256(DKGProtocol.MODE_PUBLIC_DERIVATION), 0);
        assertEq(uint256(DKGProtocol.MODE_ORGANIZER_CODEC), 1);
    }

    function test_Role_Constants() public pure {
        assertEq(uint256(DKGProtocol.ROLE_COMMITTEE), 1);
        assertEq(uint256(DKGProtocol.ROLE_ORGANIZER), 2);
    }

    function test_DomainOperatorRegisterV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_OPERATOR_REGISTER_V1),
            uint256(0x4599aabb337c91d65fe440ef7e20c6dcc72c2459fd0901c45add50b08b3bb34d)
        );
    }

    function test_DomainOrganizerRegisterV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_ORGANIZER_REGISTER_V1),
            uint256(0x41ea6f3fa95eccd1f3b1ce8e05efa11027280aa0c6b4167fd6695db659c30b28)
        );
    }

    function test_DomainDLEQV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_DLEQ_V1),
            uint256(0x48fabea26e7a072780483852e403ea60b2f51a07c735c3e4b852ac6bb99b5a91)
        );
    }

    function test_DomainCiphertextPoKV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_CIPHERTEXT_POK_V1),
            uint256(0x459d6189bebb2c0081a105f25a158877a71c77d3d77dca6bdecd5711c0db6d37)
        );
    }
}
