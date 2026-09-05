// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface Vm {
    struct Log {
        bytes32[] topics;
        bytes data;
        address emitter;
    }

    function expectRevert(bytes4 revertData) external;
    function prank(address msgSender) external;
    /// @dev Two-arg prank: sets both msg.sender and tx.origin for the next
    ///      call, needed by tests of entry points that require a direct
    ///      EOA caller (`msg.sender == tx.origin`).
    function prank(address msgSender, address txOrigin) external;
    function roll(uint256 newHeight) external;
    function warp(uint256 newTimestamp) external;
    function recordLogs() external;
    function getRecordedLogs() external returns (Log[] memory logs);
}

abstract contract Test {
    Vm internal constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function assertEq(uint256 left, uint256 right) internal pure {
        require(left == right, "assertEq(uint256)");
    }

    function assertEq(address left, address right) internal pure {
        require(left == right, "assertEq(address)");
    }

    function assertTrue(bool condition) internal pure {
        require(condition, "assertTrue");
    }

    function assertFalse(bool condition) internal pure {
        require(!condition, "assertFalse");
    }
}
