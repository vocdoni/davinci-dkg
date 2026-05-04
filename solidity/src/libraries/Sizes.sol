// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

/// @dev Maximum DKG committee size. SINGLE source of truth for the on-chain
///      side — must equal `circuits/common.MaxN` on the Go side. Update both
///      and run `make circuits` to regenerate verifier wrappers and bindings.
uint256 constant MAX_N = 32;

// ─── Epoch scheduling ────────────────────────────────────────────────────────
//
// Epoch length is expressed in BLOCKS (chain-agnostic). The wall-clock
// duration depends on the deployment chain's block time and is estimated
// off-chain by the SDK / UI from sampled block timestamps.
//
// `EPOCH_DURATION_BLOCKS` is set per-deploy as an immutable in the
// `DKGManager` constructor; the default below is what `script/DeployAll.s.sol`
// uses when no override is provided. At 12-second block time (Sepolia /
// mainnet) the default is 20 minutes per epoch.
//
// The phase fractions below (BPS = basis points, 10000 = 100%) carve up
// every epoch's lifetime as:
//
//   ─── Preparation (25 %) ─────────────────────────  ─── Service (75 %) ──
//   [0,  CSL)         CommitteeSelection: claimSlot         (5 %)
//   [CSL, CSL+KA)     KeyAssembly: submitContribution      (15 %)
//   [CSL+KA, +GAP)    finalize gap: finalizeEpoch can run   (5 %)
//   [+GAP, END)       Live: PK_ep is usable; apps register, decryptions land (75 %)
//
// At the default 100 blocks: 5 / 15 / 5 / 75 — ~5 minutes of MPC and
// ~15 minutes of decryption availability at 12 s blocks. After END blocks,
// anyone may call `createEpoch` again to mint the next epoch (permissionless).
//
// The epoch stays `Live` forever once finalized, continuing to accept
// ciphertexts / partials, until `abortEpoch` or until it gets evicted
// from the recent-epochs ring buffer.

uint256 constant DEFAULT_EPOCH_DURATION_BLOCKS = 100;
uint256 constant COMMITTEE_SELECTION_BPS       = 500;   // 5 %  (Preparation: lottery)
uint256 constant KEY_ASSEMBLY_BPS              = 1500;  // 15 % (Preparation: VSS contributions)
uint256 constant FINALIZE_GAP_BPS              = 500;   // 5 %  (Preparation: gap before finalizeEpoch)
// Preparation subtotal = COMMITTEE_SELECTION + KEY_ASSEMBLY + FINALIZE_GAP = 2500 (25 %).
// Service window (PK_ep is Live)               = remaining 7500 (75 %).

uint256 constant SEED_DELAY_BLOCKS = 1;
// `blockhash(startBlock + SEED_DELAY_BLOCKS)` is the lottery seed. The
// constant must be strictly less than the CommitteeSelection window in
// blocks at the deployed `EPOCH_DURATION_BLOCKS`; the constructor enforces
// this.
