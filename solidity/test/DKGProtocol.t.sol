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
///
///         v4 note: the contribution domain moved to `:v2` (compact
///         transcript) and the poolkey domain was replaced by
///         `finalize:v2`. The v2 hashes below were computed with
///         `cast keccak` and will be cross-checked against the regenerated
///         `protocol.json` once the circuit-side vectors land.
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

    function test_DomainContributionTranscriptV2_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_CONTRIBUTION_TRANSCRIPT_V2),
            uint256(0x4b37311b22cd0f09ae11d49f42ab65dce8fccf2600e6e2e7d41f51dc3d44b752)
        );
    }

    function test_DomainFinalizeTranscriptV2_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_FINALIZE_TRANSCRIPT_V2),
            uint256(0xe28959afa6ea38549c61aff75344fc2c9f148f1259fcef44fdd297a1d9a39d0f)
        );
    }

    function test_DomainDecryptCombineTranscriptV1_MatchesGoVector() public pure {
        assertEq(
            uint256(DKGProtocol.DOMAIN_DECRYPT_COMBINE_TRANSCRIPT_V1),
            uint256(0xb22315ced73b8ff8bb301780e4a47d6c7771b0e8a551a02a7c0df167eca08dcb)
        );
    }
}
