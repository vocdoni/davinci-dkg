// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IDKGRegistry} from "./interfaces/IDKGRegistry.sol";

/// @title DKGRegistry
/// @notice Append-only registry of operator BabyJubJub encryption keys used
///         by the DKGManager's lottery and contribution phases.
/// @dev    The registry is intentionally minimal — no access control, no key
///         revocation, no admin. Each address claims its slot with a single
///         `registerKey` call and can subsequently rotate its key via
///         `updateKey`. `nodeCount` is append-only so the DKGManager can use
///         it as the denominator of the lottery threshold without worrying
///         about the set shrinking mid-round.
contract DKGRegistry is IDKGRegistry {
    mapping(address operator => NodeKey) internal nodes;

    /// @notice Number of distinct addresses that have ever called `registerKey`.
    /// @dev    DKGManager snapshots this at `createRound` to derive the
    ///         per-round lottery threshold.
    uint64 public override nodeCount;

    /// @notice Register the caller's BabyJubJub encryption key.
    /// @param  pubX The X coordinate of the public key (non-zero).
    /// @param  pubY The Y coordinate of the public key (non-zero).
    /// @dev    Reverts with `InvalidKey` if either coordinate is zero and with
    ///         `AlreadyRegistered` if the caller has already registered. Use
    ///         `updateKey` to rotate.
    function registerKey(uint256 pubX, uint256 pubY) external override {
        if (pubX == 0 || pubY == 0) revert InvalidKey();

        NodeKey storage node = nodes[msg.sender];
        if (node.status != NodeStatus.NONE) revert AlreadyRegistered();

        node.operator = msg.sender;
        node.pubX = pubX;
        node.pubY = pubY;
        node.status = NodeStatus.ACTIVE;
        unchecked { nodeCount += 1; }

        emit NodeRegistered(msg.sender, pubX, pubY);
    }

    /// @notice Rotate the caller's previously registered BabyJubJub key.
    /// @param  pubX New X coordinate (non-zero).
    /// @param  pubY New Y coordinate (non-zero).
    /// @dev    Reverts with `NotRegistered` if the caller has never called
    ///         `registerKey`.
    function updateKey(uint256 pubX, uint256 pubY) external override {
        if (pubX == 0 || pubY == 0) revert InvalidKey();

        NodeKey storage node = nodes[msg.sender];
        if (node.status == NodeStatus.NONE) revert NotRegistered();

        node.pubX = pubX;
        node.pubY = pubY;
        node.status = NodeStatus.ACTIVE;

        emit NodeUpdated(msg.sender, pubX, pubY);
    }

    /// @notice Return the registry record for a given operator.
    /// @param  operator The address whose key is being queried.
    /// @return The full `NodeKey` struct (zeroed if the operator is unknown).
    function getNode(address operator) external view override returns (NodeKey memory) {
        return nodes[operator];
    }
}
