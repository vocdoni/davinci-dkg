// SPDX-License-Identifier: AGPL-3.0-or-later
pragma solidity 0.8.28;

/// @dev Maximum DKG committee size. SINGLE source of truth for the on-chain
///      side — must equal `circuits/common.MaxN` on the Go side. Update both
///      and run `make circuits` to regenerate verifier wrappers and bindings.
uint256 constant MAX_N = 32;

// ─── Derived transcript sizes ────────────────────────────────────────────────
//
// Word count (1 word = 32 bytes) of the `combineDecryption` transcript:
//
//   w[0..3]    C1.x C1.y C2.x C2.y
//   w[4..5]    PK_org.x PK_org.y
//   w[6..7]    A1.x A1.y            (organizer Chaum-Pedersen nonce)
//   w[8..9]    A2.x A2.y
//   w[10]      z                    (organizer DLEQ response)
//   w[11]      e                    (organizer DLEQ challenge, recomputed on chain)
//   w[12 .. 12+N)        participant indexes x_k (0 in inactive slots)
//   w[12+N .. 12+3N)     partial decryptions δ_k as (x, y) (identity when inactive)
//
// Must equal `len(decryptcombine.TranscriptScalars())` on the Go side.
uint256 constant COMBINE_TRANSCRIPT_WORDS = 12 + 3 * MAX_N;

// ─── Epoch scheduling ────────────────────────────────────────────────────────
//
// Every duration is expressed in BLOCKS (chain-agnostic). The wall-clock time
// depends on the deployment chain's block time and is estimated off-chain by
// the SDK / UI from sampled block timestamps.
//
// All four block constants below are set per-deploy as `DKGManager`
// constructor immutables; the defaults match the values in
// `script/DeployAll.s.sol` when no env override is provided. They use ~12 s
// blocks (Sepolia / mainnet) — for chains with different block times scale
// the *_BLOCKS constants accordingly so wall-time matches.
//
// Lifecycle (Preparation = first three; Service = the rest):
//
//   [0,  CSL)               CommitteeSelection : claimSlot accepted
//   [CSL, CSL+KA)           KeyAssembly        : submitContribution accepted
//   [CSL+KA, +GAP)          finalize gap       : finalizeEpoch may run
//   [CSL+KA+GAP, END)       Live               : PK_ep usable; apps register,
//                                                 ciphertexts decrypt
//
// CommitteeSelection and KeyAssembly are absolute, NOT proportional. The
// lottery is one keccak per claimer and the contribution proof is one tx
// per committee member, so they need a fixed budget — not a fraction of the
// epoch. Long epochs (multi-day) keep the same short Preparation; the extra
// time falls into Service.
//
// After END blocks anyone may call `createEpoch` again to mint the next
// epoch (permissionless). The previous epoch stays `Live` for the duration
// of its Service window, so its key remains usable while the next epoch
// bootstraps.

uint256 constant DEFAULT_EPOCH_DURATION_BLOCKS       = 100;  // ~20 min @ 12 s
uint256 constant DEFAULT_COMMITTEE_SELECTION_BLOCKS  = 25;   // ~5 min  @ 12 s
uint256 constant DEFAULT_KEY_ASSEMBLY_BLOCKS         = 25;   // ~5 min  @ 12 s
uint256 constant DEFAULT_FINALIZE_GAP_BLOCKS         = 5;    // ~1 min  @ 12 s

uint256 constant SEED_DELAY_BLOCKS = 1;
// `blockhash(startBlock + SEED_DELAY_BLOCKS)` is the lottery seed. The
// constant must be strictly less than `COMMITTEE_SELECTION_BLOCKS` so
// claimers get at least one block to call after the seed resolves; the
// constructor enforces this.
