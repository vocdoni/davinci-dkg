// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

import {IDKGRegistry} from "./interfaces/IDKGRegistry.sol";
import {BabyJubJub} from "./libraries/BabyJubJub.sol";
import {DKGProtocol} from "./libraries/DKGProtocol.sol";
import {PoseidonT3} from "poseidon-solidity/PoseidonT3.sol";
import {PoseidonT6} from "poseidon-solidity/PoseidonT6.sol";

/// @title DKGRegistry
/// @notice Append-only registry of operator BabyJubJub encryption keys used
///         by the DKGManager's lottery and contribution phases.
/// @dev    The registry is intentionally minimal — no key revocation, no
///         rotating admin. Each address claims its slot with a single
///         `registerKey` call and may rotate its key via `updateKey`.
///
///         Liveness is tracked through `lastActiveBlock` on each row,
///         refreshed by DKGManager on every accepted contribution (via
///         `markActive`) or voluntarily via `heartbeat`. Stale rows can be
///         demoted to INACTIVE by anyone through the permissionless `reap`
///         function once `INACTIVITY_WINDOW` blocks have passed with no
///         activity; the `activeCount` counter tracks the ACTIVE set and
///         is the denominator the DKGManager lottery uses.
///
///         `setManager` is restricted to the deployer address captured at
///         construction time — this prevents a front-running attack where a
///         third party locks in a malicious manager before the deployer's
///         wiring transaction lands.
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

    /// @dev Address that deployed this contract; the only address allowed to
    ///      call `setManager`. Immutable after construction.
    address private immutable _deployer;

    constructor(uint64 inactivityWindow) {
        if (inactivityWindow == 0) revert InvalidKey();
        INACTIVITY_WINDOW = inactivityWindow;
        _deployer = msg.sender;
    }

    /// @notice Pin the DKGManager address that will be authorised to call
    ///         `markActive`. May only be called once by the deployer;
    ///         subsequent calls revert with `ManagerAlreadySet`.
    /// @param  m The deployed DKGManager contract address.
    function setManager(address m) external {
        if (msg.sender != _deployer) revert Unauthorized();
        if (_managerSet) revert ManagerAlreadySet();
        if (m == address(0)) revert InvalidAddress();
        manager = m;
        _managerSet = true;
        emit ManagerSet(m);
    }

    /// @notice Register the caller's BabyJubJub encryption key together
    ///         with a Schnorr proof of knowledge of the secret. Implements
    ///         paper §5.1.1 (line ~747) and PLAN.md A14.
    /// @dev    Validates the key is a canonical, on-curve, non-identity
    ///         point. Then verifies the Schnorr PoK natively in BabyJubJub
    ///         arithmetic (no SNARK). Initialises liveness and increments
    ///         `activeCount`.
    function registerKey(
        uint256 pubX,
        uint256 pubY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external override {
        _requireValidEncryptionPoint(pubX, pubY);
        _requireValidEncryptionPoint(schnorrAx, schnorrAy);

        NodeKey storage node = nodes[msg.sender];
        if (node.status != NodeStatus.NONE) revert AlreadyRegistered();

        if (!_verifyOperatorSchnorr(msg.sender, pubX, pubY, schnorrAx, schnorrAy, schnorrZ)) {
            revert InvalidSchnorrProof();
        }

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
    /// @dev    Requires a fresh Schnorr proof of knowledge over the new key.
    ///         If the caller was previously reaped (status == INACTIVE),
    ///         this call implicitly reactivates them — rotating a key is a
    ///         strong signal that the operator is alive.
    function updateKey(
        uint256 pubX,
        uint256 pubY,
        uint256 schnorrAx,
        uint256 schnorrAy,
        uint256 schnorrZ
    ) external override {
        _requireValidEncryptionPoint(pubX, pubY);
        _requireValidEncryptionPoint(schnorrAx, schnorrAy);

        NodeKey storage node = nodes[msg.sender];
        if (node.status == NodeStatus.NONE) revert NotRegistered();

        if (!_verifyOperatorSchnorr(msg.sender, pubX, pubY, schnorrAx, schnorrAy, schnorrZ)) {
            revert InvalidSchnorrProof();
        }

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

    // ─── Schnorr proof of knowledge (paper §6.2 with operator identifiers) ──

    /// @dev Compute the Fiat-Shamir challenge `c = Poseidon(...)` over the
    ///      operator Schnorr transcript. Two-pass because the Poseidon
    ///      library only supports up to 5 inputs (T6); we hash the first 5
    ///      fields then mix in `A_y` with T3.
    ///
    ///      Transcript layout (paper §6.2 with operator identifiers
    ///      substituted for organizer identifiers, per paper line 747):
    ///
    ///        challenge = Poseidon(
    ///          inner = T6(domain, operator, pubX, pubY, A_x),
    ///          A_y
    ///        )
    ///
    ///      The shared `DOMAIN_OPERATOR_REGISTER_V1` digest namespaces the
    ///      proof so it cannot be replayed as an organizer Schnorr proof
    ///      (cross-protocol replay safety, PLAN §2.8).
    function _operatorSchnorrChallenge(
        address op,
        uint256 pubX,
        uint256 pubY,
        uint256 ax,
        uint256 ay
    ) internal pure returns (uint256) {
        // Reduce the bytes32 domain digest into BN254 scalar field so it is
        // a valid Poseidon input.
        uint256 domainField = uint256(DKGProtocol.DOMAIN_OPERATOR_REGISTER_V1) % BabyJubJub.Q;
        uint256[5] memory in1;
        in1[0] = domainField;
        in1[1] = uint256(uint160(op));
        in1[2] = pubX;
        in1[3] = pubY;
        in1[4] = ax;
        uint256 inner = PoseidonT6.hash(in1);
        uint256[2] memory in2;
        in2[0] = inner;
        in2[1] = ay;
        return PoseidonT3.hash(in2);
    }

    /// @dev Verify the Schnorr PoK: `z·G == A + c·PK_op` on BabyJubJub.
    ///      Returns true on success; the caller reverts with
    ///      `InvalidSchnorrProof` on false.
    function _verifyOperatorSchnorr(
        address op,
        uint256 pubX,
        uint256 pubY,
        uint256 ax,
        uint256 ay,
        uint256 z
    ) internal view returns (bool) {
        uint256 c = _operatorSchnorrChallenge(op, pubX, pubY, ax, ay);
        (uint256 zGx, uint256 zGy) = BabyJubJub.scalarMulBase(z);
        (uint256 cPKx, uint256 cPKy) = BabyJubJub.scalarMul(c, pubX, pubY);
        (uint256 rhsX, uint256 rhsY) = BabyJubJub.pointAdd(ax, ay, cPKx, cPKy);
        return zGx == rhsX && zGy == rhsY;
    }

    /// @dev Validate that a BabyJubJub point received as encryption key or
    ///      Schnorr nonce is canonical, on-curve, and non-identity. Prime-
    ///      subgroup membership is implicit in the Schnorr PoK (any scalar
    ///      multiple of the generator G lies in the prime subgroup), so we
    ///      do not pay for the explicit `isInPrimeSubgroup` check here.
    function _requireValidEncryptionPoint(uint256 x, uint256 y) internal pure {
        if (x >= BabyJubJub.Q || y >= BabyJubJub.Q) revert PointNotCanonical();
        if (x == 0 && y == 1) revert PointIsIdentity();
        if (!BabyJubJub.isOnCurve(x, y)) revert PointNotOnCurve();
    }

    /// @notice Refresh an operator's `lastActiveBlock` after a successful
    ///         contribution. Only the configured DKGManager may call this.
    /// @dev    Silently no-ops for unregistered or inactive nodes so the
    ///         manager never reverts mid-epoch on a stale registry row.
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
    ///         key or participating in a epoch. The escape valve for
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
