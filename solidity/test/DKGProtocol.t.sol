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

    function test_DomainContributionTranscriptV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_CONTRIBUTION_TRANSCRIPT_V1),
            uint256(0x29aa19fbd94aef15994e2f585c00bbd3e7aa5aefc9372efb2ce55433ca0c6a72)
        );
    }

    function test_DomainPoolKeyTranscriptV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_POOLKEY_TRANSCRIPT_V1),
            uint256(0xae031fc261aed61242596185b006e57bdcba774d6ea39d3348a9a570b38d9ff4)
        );
    }

    function test_DomainDecryptCombineTranscriptV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_DECRYPT_COMBINE_TRANSCRIPT_V1),
            uint256(0xb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb)
        );
    }
}
