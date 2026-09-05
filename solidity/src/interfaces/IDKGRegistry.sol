// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

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
        /// @dev Block at which the node last entered the active set (registerKey
        ///      or reactivation); a node only enters lotteries of epochs created
        ///      after it, so neither fresh nor dormant identities can be ground
        ///      against a revealed seed.
        uint64 registeredAtBlock;
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
    error InvalidAddress();
    error AlreadyRegistered();
    error NotRegistered();
    error NotManager();
    error ManagerAlreadySet();
    error ManagerNotSet();
    error NotActive();
    error StillActive();
    error NotInactive();
    error Unauthorized();
    error InvalidSchnorrProof();
    error PointNotCanonical();
    error PointNotOnCurve();
    error PointIsIdentity();
    /// @dev The registered encryption key is not in the curve's prime-order
    ///      subgroup. The Schnorr PoK does not imply membership: small-order
    ///      (cofactor) points satisfy its verification equation too, and
    ///      shares encrypted to such a key can never be decrypted.
    error PointNotInSubgroup();

    // ── registration ──────────────────────────────────────────────────────
    /// @notice Register the caller's BabyJubJub encryption key together
    ///         with a Schnorr proof of knowledge of the secret. Implements
    ///         paper §5.1.1 (line ~747).
    /// @param  pubX     X coordinate of the operator's encryption key.
    /// @param  pubY     Y coordinate of the operator's encryption key.
    /// @param  schnorrAx X coordinate of the Schnorr nonce point A = w·G.
    /// @param  schnorrAy Y coordinate of the Schnorr nonce point.
    /// @param  schnorrZ  Schnorr response: z = w + c · sk_op  (mod L), where
    ///                   c = Poseidon(domain || op || pub || A) computed via
    ///                   `_operatorSchnorrChallenge`.
    function registerKey(
        uint256 pubX,
        uint256 pubY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external;
    function updateKey(
        uint256 pubX,
        uint256 pubY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external;

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
