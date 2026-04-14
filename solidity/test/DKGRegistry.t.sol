// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {Test} from "forge-std/Test.sol";
import {DKGRegistry} from "../src/DKGRegistry.sol";
import {IDKGRegistry} from "../src/interfaces/IDKGRegistry.sol";

contract DKGRegistryTest is Test {
    uint64 internal constant WINDOW = 1_000;

    DKGRegistry public registry;
    address internal alice = address(0xA11CE);
    address internal bob = address(0xB0B);
    address internal fakeManager = address(0xBEEF);

    function setUp() public {
        registry = new DKGRegistry(WINDOW);
    }

    // ── registration ──────────────────────────────────────────────────────

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

    // ── liveness ──────────────────────────────────────────────────────────

    function test_RegisterKey_InitialisesLiveness() public {
        vm.roll(123);
        vm.prank(alice);
        registry.registerKey(11, 22);

        IDKGRegistry.NodeKey memory key = registry.getNode(alice);
        assertEq(uint8(key.status), uint8(IDKGRegistry.NodeStatus.ACTIVE));
        assertEq(key.lastActiveBlock, 123);
        assertEq(registry.activeCount(), 1);
        assertEq(registry.nodeCount(), 1);
        assertTrue(registry.isActive(alice));
    }

    function test_Reap_RejectsFreshNode() public {
        vm.prank(alice);
        registry.registerKey(11, 22);

        vm.expectRevert(IDKGRegistry.StillActive.selector);
        registry.reap(alice);
    }

    function test_Reap_RejectsCooldownBoundary() public {
        vm.roll(100);
        vm.prank(alice);
        registry.registerKey(11, 22);

        // Exactly WINDOW blocks later is still "still active" — the reap
        // condition is strict (block > lastActiveBlock + WINDOW).
        vm.roll(100 + WINDOW);
        vm.expectRevert(IDKGRegistry.StillActive.selector);
        registry.reap(alice);
    }

    function test_Reap_SucceedsAfterWindow() public {
        vm.roll(100);
        vm.prank(alice);
        registry.registerKey(11, 22);

        vm.roll(100 + WINDOW + 1);
        registry.reap(alice);

        IDKGRegistry.NodeKey memory key = registry.getNode(alice);
        assertEq(uint8(key.status), uint8(IDKGRegistry.NodeStatus.INACTIVE));
        assertEq(registry.activeCount(), 0);
        assertEq(registry.nodeCount(), 1);
        assertFalse(registry.isActive(alice));
    }

    function test_Reap_RejectsAlreadyInactive() public {
        vm.prank(alice);
        registry.registerKey(11, 22);
        vm.roll(block.number + WINDOW + 1);
        registry.reap(alice);

        vm.expectRevert(IDKGRegistry.NotActive.selector);
        registry.reap(alice);
    }

    function test_Reap_RejectsUnregisteredOperator() public {
        vm.expectRevert(IDKGRegistry.NotActive.selector);
        registry.reap(alice);
    }

    function test_Reactivate_RestoresActiveRow() public {
        vm.prank(alice);
        registry.registerKey(11, 22);
        vm.roll(block.number + WINDOW + 1);
        registry.reap(alice);

        vm.roll(block.number + 5);
        vm.prank(alice);
        registry.reactivate();

        IDKGRegistry.NodeKey memory key = registry.getNode(alice);
        assertEq(uint8(key.status), uint8(IDKGRegistry.NodeStatus.ACTIVE));
        assertEq(key.lastActiveBlock, uint64(block.number));
        assertEq(registry.activeCount(), 1);
    }

    function test_Reactivate_RejectsActiveCaller() public {
        vm.prank(alice);
        registry.registerKey(11, 22);

        vm.expectRevert(IDKGRegistry.NotInactive.selector);
        vm.prank(alice);
        registry.reactivate();
    }

    function test_UpdateKey_ReactivatesInactive() public {
        vm.prank(alice);
        registry.registerKey(11, 22);
        vm.roll(block.number + WINDOW + 1);
        registry.reap(alice);
        assertEq(registry.activeCount(), 0);

        vm.prank(alice);
        registry.updateKey(33, 44);

        IDKGRegistry.NodeKey memory key = registry.getNode(alice);
        assertEq(uint8(key.status), uint8(IDKGRegistry.NodeStatus.ACTIVE));
        assertEq(key.pubX, 33);
        assertEq(key.lastActiveBlock, uint64(block.number));
        assertEq(registry.activeCount(), 1);
    }

    function test_Heartbeat_UpdatesLastActiveBlock() public {
        vm.roll(100);
        vm.prank(alice);
        registry.registerKey(11, 22);
        assertEq(registry.getNode(alice).lastActiveBlock, 100);

        vm.roll(200);
        vm.prank(alice);
        registry.heartbeat();
        assertEq(registry.getNode(alice).lastActiveBlock, 200);
    }

    function test_Heartbeat_RejectsInactive() public {
        vm.prank(alice);
        registry.registerKey(11, 22);
        vm.roll(block.number + WINDOW + 1);
        registry.reap(alice);

        vm.expectRevert(IDKGRegistry.NotActive.selector);
        vm.prank(alice);
        registry.heartbeat();
    }

    // ── manager wiring ────────────────────────────────────────────────────

    function test_SetManager_LocksAfterFirstCall() public {
        registry.setManager(fakeManager);
        assertEq(registry.manager(), fakeManager);

        vm.expectRevert(IDKGRegistry.ManagerAlreadySet.selector);
        registry.setManager(address(0xDEAD));
    }

    function test_SetManager_RejectsZero() public {
        vm.expectRevert(IDKGRegistry.InvalidAddress.selector);
        registry.setManager(address(0));
    }

    function test_SetManager_RejectsNonDeployer() public {
        // alice did not deploy the registry — she should be rejected.
        vm.expectRevert(IDKGRegistry.Unauthorized.selector);
        vm.prank(alice);
        registry.setManager(fakeManager);
    }

    function test_MarkActive_OnlyManager() public {
        registry.setManager(fakeManager);
        vm.prank(alice);
        registry.registerKey(11, 22);

        vm.expectRevert(IDKGRegistry.NotManager.selector);
        registry.markActive(alice);

        vm.roll(500);
        vm.prank(fakeManager);
        registry.markActive(alice);
        assertEq(registry.getNode(alice).lastActiveBlock, 500);
    }

    function test_MarkActive_RequiresManagerSet() public {
        vm.prank(alice);
        registry.registerKey(11, 22);

        vm.expectRevert(IDKGRegistry.ManagerNotSet.selector);
        vm.prank(fakeManager);
        registry.markActive(alice);
    }

    function test_MarkActive_NoOpOnInactive() public {
        registry.setManager(fakeManager);

        vm.roll(10);
        vm.prank(alice);
        registry.registerKey(11, 22);

        vm.roll(10 + WINDOW + 1);
        registry.reap(alice);
        uint64 frozen = registry.getNode(alice).lastActiveBlock;

        vm.roll(block.number + 5);
        vm.prank(fakeManager);
        registry.markActive(alice);

        // lastActiveBlock must not move while the row is INACTIVE.
        assertEq(registry.getNode(alice).lastActiveBlock, frozen);
        assertEq(
            uint8(registry.getNode(alice).status),
            uint8(IDKGRegistry.NodeStatus.INACTIVE)
        );
    }

    function test_ActiveCount_TracksChurn() public {
        registry.setManager(fakeManager);

        vm.prank(alice);
        registry.registerKey(11, 22);
        vm.prank(bob);
        registry.registerKey(33, 44);
        assertEq(registry.activeCount(), 2);

        vm.roll(block.number + WINDOW + 1);
        registry.reap(alice);
        assertEq(registry.activeCount(), 1);

        vm.prank(alice);
        registry.reactivate();
        assertEq(registry.activeCount(), 2);
    }
}
