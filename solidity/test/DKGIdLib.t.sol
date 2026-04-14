// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DKGIdLib} from "../src/libraries/DKGIdLib.sol";

contract DKGIdLibTest is Test {
    function test_GetPrefix_IsStableForSameInputs() public pure {
        uint32 prefixA = DKGIdLib.getPrefix(31337, address(0x1234));
        uint32 prefixB = DKGIdLib.getPrefix(31337, address(0x1234));

        assertEq(prefixA, prefixB);
    }

    function test_ComputeRoundId_EmbedsPrefixAndNonce() public pure {
        uint32 prefix = 0xAABBCCDD;
        bytes12 roundId = DKGIdLib.computeRoundId(prefix, 7);

        // forge-lint: disable-next-line(unsafe-typecast)
        assertEq(uint32(bytes4(roundId)), prefix);
        // forge-lint: disable-next-line(unsafe-typecast)
        assertEq(uint64(bytes8(roundId << 32)), 7);
    }
}
