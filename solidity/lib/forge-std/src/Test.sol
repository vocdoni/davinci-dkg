// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface Vm {
    function expectRevert(bytes4 revertData) external;
    function prank(address msgSender) external;
    function roll(uint256 newHeight) external;
}

abstract contract Test {
    Vm internal constant vm = Vm(address(uint160(uint256(keccak256("hevm cheat code")))));

    function assertEq(uint256 left, uint256 right) internal pure {
        require(left == right, "assertEq(uint256)");
    }

    function assertEq(address left, address right) internal pure {
        require(left == right, "assertEq(address)");
    }
}
