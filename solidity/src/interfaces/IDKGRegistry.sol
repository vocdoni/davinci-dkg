// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

/// @title  IDKGRegistry
/// @notice Append-only registry of operator BabyJubJub encryption keys, with
///         a liveness mechanism that lets inactive nodes age out of the
///         DKGManager lottery without any manual pruning pass.
interface IDKGRegistry {
    /// @notice Lifecycle states for a registered node.
    /// @dev    Transitions:
    ///         NONE → ACTIVE        (registerKey)
    ///         ACTIVE → INACTIVE    (reap, after INACTIVITY_WINDOW elapsed)
    ///         INACTIVE → ACTIVE    (reactivate, updateKey)
    enum NodeStatus {
        NONE,
        ACTIVE,
        INACTIVE
    }

    /// @notice Full registry row for a single operator.
    /// @dev    `status` (1 byte) and `lastActiveBlock` (8 bytes) share the
    ///         same storage slot, so the liveness field is free.
    struct NodeKey {
        address operator;
        uint256 pubX;
        uint256 pubY;
        NodeStatus status;
        uint64 lastActiveBlock;
    }

    // ── events ────────────────────────────────────────────────────────────
    event NodeRegistered(address indexed operator, uint256 pubX, uint256 pubY);
    event NodeUpdated(address indexed operator, uint256 pubX, uint256 pubY);
    /// @notice Emitted whenever an operator's `lastActiveBlock` is refreshed
    ///         — by `markActive` (from DKGManager) or `heartbeat` (self).
    event NodeMarkedActive(address indexed operator, uint64 atBlock);
    /// @notice Emitted when a permissionless reap demotes a stale node.
    event NodeReaped(address indexed operator, uint64 lastActiveBlock);
    /// @notice Emitted when an operator explicitly rejoins after being reaped.
    event NodeReactivated(address indexed operator);
    /// @notice Emitted exactly once, when `setManager` locks in the manager.
    event ManagerSet(address indexed manager);

    // ── errors ────────────────────────────────────────────────────────────
    error InvalidKey();
    error AlreadyRegistered();
    error NotRegistered();
    error NotManager();
    error ManagerAlreadySet();
    error ManagerNotSet();
    error NotActive();
    error StillActive();
    error NotInactive();

    // ── registration ──────────────────────────────────────────────────────
    function registerKey(uint256 pubX, uint256 pubY) external;
    function updateKey(uint256 pubX, uint256 pubY) external;

    // ── liveness ──────────────────────────────────────────────────────────
    /// @notice Refresh the caller's `lastActiveBlock`. Callable only by the
    ///         registered manager (DKGManager.submitContribution).
    function markActive(address operator) external;

    /// @notice Demote a stale node whose `lastActiveBlock + INACTIVITY_WINDOW`
    ///         has passed. Permissionless.
    function reap(address operator) external;

    /// @notice Rejoin the active set after being reaped. Caller must be the
    ///         previously-reaped operator.
    function reactivate() external;

    /// @notice Refresh the caller's `lastActiveBlock` without doing anything
    ///         else. Reverts if the caller is not ACTIVE.
    function heartbeat() external;

    // ── views ─────────────────────────────────────────────────────────────
    function getNode(address operator) external view returns (NodeKey memory);
    function nodeCount() external view returns (uint64);
    function activeCount() external view returns (uint64);
    function manager() external view returns (address);
    function INACTIVITY_WINDOW() external view returns (uint64);
    function isActive(address operator) external view returns (bool);
}
