// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {Test} from "forge-std/Test.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {IDKGRegistry} from "../src/interfaces/IDKGRegistry.sol";

contract DKGRegistryTest is Test {
    DKGRegistry public registry;

    function setUp() public {
        registry = new DKGRegistry();
    }

    function test_RegisterKey() public view {
        IDKGRegistry.NodeKey memory key = registry.getNode(address(this));

        assertEq(uint256(key.pubX), 0);
        assertEq(uint256(key.pubY), 0);
        assertEq(uint8(key.status), uint8(IDKGRegistry.NodeStatus.NONE));
    }

    function test_RegisterKey_PersistsCoordinates() public {
        registry.registerKey(11, 22);

        IDKGRegistry.NodeKey memory key = registry.getNode(address(this));

        assertEq(key.operator, address(this));
        assertEq(key.pubX, 11);
        assertEq(key.pubY, 22);
        assertEq(uint8(key.status), uint8(IDKGRegistry.NodeStatus.ACTIVE));
    }

    function test_RegisterKey_RejectsDuplicateRegistration() public {
        registry.registerKey(11, 22);

        vm.expectRevert(IDKGRegistry.AlreadyRegistered.selector);
        registry.registerKey(33, 44);
    }

    function test_UpdateKey() public {
        registry.registerKey(11, 22);
        registry.updateKey(33, 44);

        IDKGRegistry.NodeKey memory key = registry.getNode(address(this));

        assertEq(key.pubX, 33);
        assertEq(key.pubY, 44);
    }

    function test_RegisterKey_RejectsZeroCoordinates() public {
        vm.expectRevert(IDKGRegistry.InvalidKey.selector);
        registry.registerKey(0, 22);
    }
}
