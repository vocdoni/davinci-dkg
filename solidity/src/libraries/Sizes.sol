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
//   [0,  REG)            registration: nodes call `claimSlot`           (5%)
//   [REG, REG+CON)       contribution: nodes submit DKG contributions   (15%)
//   [REG+CON, REG+CON+G) finalize gap: anyone can call `finalizeEpoch`  (5%)
//   [REG+CON+G, END)     decryption window: PK_ep is live, organizers   (75%)
//                        register apps, encrypt + decrypt
//
// At the default 100 blocks: 5 / 15 / 5 / 75 — i.e. 5 minutes of MPC and
// 15 minutes of decryption availability. After END blocks, anyone may call
// `createEpoch` again to mint the next epoch (permissionless).
//
// The previous epoch stays in `Finalized` state forever, continuing to
// accept ciphertexts / partials, until `abortEpoch` or it gets evicted
// from the recent-epochs ring buffer.

uint256 constant DEFAULT_EPOCH_DURATION_BLOCKS = 100;
uint256 constant REGISTRATION_BPS              = 500;   // 5 %
uint256 constant CONTRIBUTION_BPS              = 1500;  // 15 %
uint256 constant FINALIZE_GAP_BPS              = 500;   // 5 %
// MPC subtotal = REGISTRATION + CONTRIBUTION + FINALIZE_GAP = 2500 (25 %)
// Decryption window = remaining 7500 (75 %).

uint256 constant SEED_DELAY_BLOCKS = 1;
// `blockhash(startBlock + SEED_DELAY_BLOCKS)` is the lottery seed. The
// constant must be strictly less than the registration window in blocks at
// the deployed `EPOCH_DURATION_BLOCKS`; the constructor enforces this.
