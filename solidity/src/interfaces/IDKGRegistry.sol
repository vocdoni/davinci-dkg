// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

interface IDKGRegistry {
    enum NodeStatus {
        NONE,
        ACTIVE,
        INACTIVE
    }

    struct NodeKey {
        address operator;
        uint256 pubX;
        uint256 pubY;
        NodeStatus status;
    }

    event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY);
    event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY);

    error InvalidKey();
    error AlreadyRegistered();
    error NotRegistered();

    function registerKey(uint256 pubX, uint256 pubY) external;
    function updateKey(uint256 pubX, uint256 pubY) external;
    function getNode(address operator) external view returns (NodeKey memory);
    function nodeCount() external view returns (uint64);
}
