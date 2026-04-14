// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity ^0.8.28;

import {IDKGRegistry} from "./interfaces/IDKGRegistry.sol";

/// @title DKGRegistry
/// @notice Append-only registry of operator BabyJubJub encryption keys used
///         by the DKGManager's lottery and contribution phases.
/// @dev    The registry is intentionally minimal — no access control, no
///         key revocation, no admin. Each address claims its slot with a
///         single `registerKey` call and may rotate its key via
///         `updateKey`.
///
///         Liveness is tracked through `lastActiveBlock` on each row,
///         refreshed by DKGManager on every accepted contribution (via
///         `markActive`) or voluntarily via `heartbeat`. Stale rows can be
///         demoted to INACTIVE by anyone through the permissionless `reap`
///         function once `INACTIVITY_WINDOW` blocks have passed with no
///         activity; the `activeCount` counter tracks the ACTIVE set and
///         is the denominator the DKGManager lottery uses.
contract DKGRegistry is IDKGRegistry {
    mapping(address operator => NodeKey) internal nodes;

    /// @notice Total number of distinct addresses that have ever called
    ///         `registerKey`. Monotonically increasing, never decremented.
    uint64 public override nodeCount;

    /// @notice Number of nodes currently in the ACTIVE state. Decremented
    ///         on `reap`, incremented on `registerKey`, `reactivate` and
    ///         the auto-reactivation path in `updateKey`. DKGManager reads
    ///         this as the denominator of the lottery threshold.
    uint64 public override activeCount;

    /// @notice Number of blocks a node may remain silent before it becomes
    ///         eligible for `reap`. Set at construction time.
    uint64 public immutable override INACTIVITY_WINDOW;

    /// @notice The DKGManager contract that is allowed to call `markActive`.
    ///         Set exactly once via `setManager` by the deployer, after the
    ///         DKGManager is itself deployed (the manager's constructor
    ///         takes the registry address, so the link is one-shot in the
    ///         other direction).
    address public override manager;
    bool private _managerSet;

    constructor(uint64 inactivityWindow) {
        if (inactivityWindow == 0) revert InvalidKey();
        INACTIVITY_WINDOW = inactivityWindow;
    }

    /// @notice Pin the DKGManager address that will be authorised to call
    ///         `markActive`. May only be called once; subsequent calls
    ///         revert with `ManagerAlreadySet`.
    /// @param  m The deployed DKGManager contract address.
    function setManager(address m) external {
        if (_managerSet) revert ManagerAlreadySet();
        if (m == address(0)) revert InvalidKey();
        manager = m;
        _managerSet = true;
        emit ManagerSet(m);
    }

    /// @notice Register the caller's BabyJubJub encryption key.
    /// @dev    Reverts if either coordinate is zero or if the caller has
    ///         already registered. Initialises liveness with the current
    ///         block number and increments `activeCount`.
    function registerKey(uint256 pubX, uint256 pubY) external override {
        if (pubX == 0 || pubY == 0) revert InvalidKey();

        NodeKey storage node = nodes[msg.sender];
        if (node.status != NodeStatus.NONE) revert AlreadyRegistered();

        node.operator = msg.sender;
        node.pubX = pubX;
        node.pubY = pubY;
        node.status = NodeStatus.ACTIVE;
        node.lastActiveBlock = uint64(block.number);
        unchecked {
            nodeCount += 1;
            activeCount += 1;
        }

        emit NodeRegistered(msg.sender, pubX, pubY);
        emit NodeMarkedActive(msg.sender, uint64(block.number));
    }

    /// @notice Rotate the caller's previously registered BabyJubJub key.
    /// @dev    If the caller was previously reaped (status == INACTIVE),
    ///         this call implicitly reactivates them — rotating a key is a
    ///         strong signal that the operator is alive.
    function updateKey(uint256 pubX, uint256 pubY) external override {
        if (pubX == 0 || pubY == 0) revert InvalidKey();

        NodeKey storage node = nodes[msg.sender];
        if (node.status == NodeStatus.NONE) revert NotRegistered();

        node.pubX = pubX;
        node.pubY = pubY;
        node.lastActiveBlock = uint64(block.number);
        if (node.status == NodeStatus.INACTIVE) {
            node.status = NodeStatus.ACTIVE;
            unchecked {
                activeCount += 1;
            }
            emit NodeReactivated(msg.sender);
        }

        emit NodeUpdated(msg.sender, pubX, pubY);
        emit NodeMarkedActive(msg.sender, uint64(block.number));
    }

    /// @notice Refresh an operator's `lastActiveBlock` after a successful
    ///         contribution. Only the configured DKGManager may call this.
    /// @dev    Silently no-ops for unregistered or inactive nodes so the
    ///         manager never reverts mid-round on a stale registry row.
    ///         Skips the SSTORE when the row was already refreshed at the
    ///         same block (cheap hot path).
    function markActive(address operator) external override {
        if (manager == address(0)) revert ManagerNotSet();
        if (msg.sender != manager) revert NotManager();

        NodeKey storage node = nodes[operator];
        if (node.status != NodeStatus.ACTIVE) return;
        uint64 nowBlock = uint64(block.number);
        if (node.lastActiveBlock == nowBlock) return;

        node.lastActiveBlock = nowBlock;
        emit NodeMarkedActive(operator, nowBlock);
    }

    /// @notice Demote a stale node that has not produced a contribution or
    ///         heartbeat within `INACTIVITY_WINDOW` blocks. Permissionless.
    /// @dev    Reverts `NotActive` if the node is already inactive (or was
    ///         never registered), and `StillActive` if the cooldown has
    ///         not elapsed.
    function reap(address operator) external override {
        NodeKey storage node = nodes[operator];
        if (node.status != NodeStatus.ACTIVE) revert NotActive();

        uint256 deadline = uint256(node.lastActiveBlock) + uint256(INACTIVITY_WINDOW);
        if (block.number <= deadline) revert StillActive();

        node.status = NodeStatus.INACTIVE;
        unchecked {
            activeCount -= 1;
        }
        emit NodeReaped(operator, node.lastActiveBlock);
    }

    /// @notice Rejoin the active set after being reaped.
    /// @dev    Reverts `NotInactive` if the caller's row is not INACTIVE.
    ///         Resets `lastActiveBlock` to the current block so the new
    ///         grace period starts from now.
    function reactivate() external override {
        NodeKey storage node = nodes[msg.sender];
        if (node.status != NodeStatus.INACTIVE) revert NotInactive();

        node.status = NodeStatus.ACTIVE;
        node.lastActiveBlock = uint64(block.number);
        unchecked {
            activeCount += 1;
        }
        emit NodeReactivated(msg.sender);
        emit NodeMarkedActive(msg.sender, uint64(block.number));
    }

    /// @notice Refresh the caller's `lastActiveBlock` without touching the
    ///         key or participating in a round. The escape valve for
    ///         healthy operators that the lottery never selects.
    /// @dev    Reverts `NotActive` if the caller is not currently ACTIVE
    ///         (use `reactivate` first in that case).
    function heartbeat() external override {
        NodeKey storage node = nodes[msg.sender];
        if (node.status != NodeStatus.ACTIVE) revert NotActive();

        node.lastActiveBlock = uint64(block.number);
        emit NodeMarkedActive(msg.sender, uint64(block.number));
    }

    /// @notice Return the registry record for a given operator.
    /// @param  operator The address whose key is being queried.
    /// @return The full `NodeKey` struct (zeroed if the operator is unknown).
    function getNode(address operator) external view override returns (NodeKey memory) {
        return nodes[operator];
    }

    /// @notice Shorthand for `getNode(operator).status == ACTIVE`.
    function isActive(address operator) external view override returns (bool) {
        return nodes[operator].status == NodeStatus.ACTIVE;
    }
}
