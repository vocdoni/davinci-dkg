// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

interface VmSafe {
    function envUint(string calldata name) external returns (uint256);
    function startBroadcast(uint256 privateKey) external;
    function stopBroadcast() external;
}

abstract contract Script {
    VmSafe internal constant vm = VmSafe(address(uint160(uint256(keccak256("hevm cheat code")))));
}

library console {
    function log(string memory, address) internal pure {}
}
