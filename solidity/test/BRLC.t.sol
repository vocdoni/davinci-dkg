// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {BRLC} from "../src/libraries/BRLC.sol";

contract BRLCHarness {
    function deriveChallenge(bytes12 roundId, bytes32 domain, bytes32 anchor) external pure returns (uint256) {
        return BRLC.deriveChallenge(roundId, domain, anchor);
    }

    function commit(uint256 challenge, uint256[] memory values) external pure returns (uint256) {
        return BRLC.commit(challenge, values);
    }
}

contract BRLCTest is Test {
    BRLCHarness internal harness;

    function setUp() public {
        harness = new BRLCHarness();
    }

    function test_DeriveChallenge_IsDeterministic() public view {
        bytes12 roundId = bytes12("round-000001");
        bytes32 domain = keccak256("contribution");
        bytes32 anchor = bytes32(uint256(0x1234));

        uint256 first = harness.deriveChallenge(roundId, domain, anchor);
        uint256 second = harness.deriveChallenge(roundId, domain, anchor);

        assertEq(first, second);
    }

    function test_Commit_ComputesExpectedCombination() public view {
        uint256[] memory values = new uint256[](3);
        values[0] = 2;
        values[1] = 3;
        values[2] = 7;

        uint256 commitment = harness.commit(5, values);

        assertEq(commitment, 960);
    }
}
