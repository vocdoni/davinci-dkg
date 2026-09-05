// The keccak Merkle tree over one pool key's per-member share commitments.
// `finalizeEpoch` stores its root (`poolShareRoots[eid][j]`) and
// `submitPartialDecryption` proves the member's leaf against it; this module
// mirrors `circuits/common/merkle.go` and the tree the contract folds so the
// SDK can build and check that path. Tagged encoding, so no leaf can be
// replayed as an internal node and an absent member is distinguishable from
// any point:
//
//   leaf[p-1] = keccak256(0x00 ‖ D_p.x ‖ D_p.y)      for committee member p
//   leaf[i]   = keccak256("davinci-dkg:merkle-empty:v1")   for i ≥ committeeSize
//   node      = keccak256(0x01 ‖ left ‖ right)
//
// Coordinates are 32-byte big-endian words (the on-chain RTE form); the tree
// is a fixed `MERKLE_DEPTH` levels over `MAX_N` leaves, indexed by
// participant index − 1.

import { keccak256, toHex, type Hex } from 'viem';
import { MAX_N, MERKLE_DEPTH } from './sizes.js';
import { bytes32, type TranscriptPoint } from './transcript.js';

/** Leaf of a committee slot that does not exist (`i ≥ committeeSize`). */
export const MERKLE_EMPTY_LEAF: Hex = keccak256(toHex('davinci-dkg:merkle-empty:v1'));

const LEAF_TAG = '00';
const NODE_TAG = '01';

/** `keccak256(0x00 ‖ D.x ‖ D.y)`: one member's share commitment as its leaf. */
export function shareCommitmentLeaf(commitment: TranscriptPoint): Hex {
  return keccak256(('0x' + LEAF_TAG + bytes32(commitment.x).slice(2) + bytes32(commitment.y).slice(2)) as Hex);
}

/** `keccak256(0x01 ‖ left ‖ right)`: two children into their parent. */
export function merkleNode(left: Hex, right: Hex): Hex {
  return keccak256(('0x' + NODE_TAG + left.slice(2) + right.slice(2)) as Hex);
}

/**
 * Lay one pool key's share commitments out as the `MAX_N` leaves: slot `i`
 * (committee member `i + 1`) is `shareCommitmentLeaf(D_{i+1})` for
 * `i < committeeSize` and the empty leaf beyond. `shareCommitments` is in
 * committee order and may be the full padded row from the finalize
 * transcript; only the first `committeeSize` entries are read.
 */
export function shareCommitmentLeaves(
  shareCommitments: readonly TranscriptPoint[],
  committeeSize: number,
): Hex[] {
  if (!Number.isInteger(committeeSize) || committeeSize < 0 || committeeSize > MAX_N) {
    throw new Error(`merkle: committee size ${committeeSize} out of range [0, ${MAX_N}]`);
  }
  if (shareCommitments.length < committeeSize) {
    throw new Error(`merkle: got ${shareCommitments.length} share commitments for a committee of ${committeeSize}`);
  }
  const leaves: Hex[] = new Array<Hex>(MAX_N).fill(MERKLE_EMPTY_LEAF);
  for (let i = 0; i < committeeSize; i++) leaves[i] = shareCommitmentLeaf(shareCommitments[i]);
  return leaves;
}

function merkleLevel(nodes: readonly Hex[]): Hex[] {
  const next: Hex[] = [];
  for (let i = 0; i < nodes.length; i += 2) next.push(merkleNode(nodes[i], nodes[i + 1]));
  return next;
}

function requireLeaves(leaves: readonly Hex[]): void {
  if (leaves.length !== MAX_N) throw new Error(`merkle: expected ${MAX_N} leaves, got ${leaves.length}`);
}

/** Fold the `MAX_N` leaves into the root over `MERKLE_DEPTH` levels. */
export function merkleRoot(leaves: readonly Hex[]): Hex {
  requireLeaves(leaves);
  let level: Hex[] = [...leaves];
  for (let depth = 0; depth < MERKLE_DEPTH; depth++) level = merkleLevel(level);
  return level[0];
}

/**
 * The siblings of leaf `index` bottom-up — the `shareProof` argument of
 * `submitPartialDecryption` for participant `index + 1`.
 */
export function merklePath(leaves: readonly Hex[], index: number): Hex[] {
  requireLeaves(leaves);
  if (!Number.isInteger(index) || index < 0 || index >= MAX_N) {
    throw new Error(`merkle: leaf index ${index} out of range [0, ${MAX_N})`);
  }
  const path: Hex[] = [];
  let level: Hex[] = [...leaves];
  let cursor = index;
  for (let depth = 0; depth < MERKLE_DEPTH; depth++) {
    path.push(level[cursor ^ 1]);
    cursor >>= 1;
    level = merkleLevel(level);
  }
  return path;
}

/**
 * Recompute the root a `MERKLE_DEPTH`-long path claims for `leaf` at `index`
 * — exactly the fold `submitPartialDecryption` runs before comparing with
 * `poolShareRoots`.
 */
export function merkleRootFromPath(leaf: Hex, index: number, path: readonly Hex[]): Hex {
  if (path.length !== MERKLE_DEPTH) throw new Error(`merkle: expected a ${MERKLE_DEPTH}-long path, got ${path.length}`);
  if (!Number.isInteger(index) || index < 0 || index >= MAX_N) {
    throw new Error(`merkle: leaf index ${index} out of range [0, ${MAX_N})`);
  }
  let node = leaf;
  let cursor = index;
  for (const sibling of path) {
    node = cursor & 1 ? merkleNode(sibling, node) : merkleNode(node, sibling);
    cursor >>= 1;
  }
  return node;
}

/** True when `path` proves `leaf` at `index` under `root`. */
export function verifyMerklePath(root: Hex, leaf: Hex, index: number, path: readonly Hex[]): boolean {
  try {
    return merkleRootFromPath(leaf, index, path).toLowerCase() === root.toLowerCase();
  } catch {
    return false;
  }
}

/**
 * The `shareProof` of committee member `participantIndex` (1-based) under one
 * pool key, from that key's share commitments in committee order (e.g.
 * `FinalizeTranscript.shareCommitments[j]`), together with the leaf and the
 * root the contract must hold.
 */
export function shareProof(
  shareCommitments: readonly TranscriptPoint[],
  committeeSize: number,
  participantIndex: number,
): { leaf: Hex; path: Hex[]; root: Hex } {
  if (!Number.isInteger(participantIndex) || participantIndex < 1 || participantIndex > committeeSize) {
    throw new Error(`merkle: participant index ${participantIndex} out of range [1, ${committeeSize}]`);
  }
  const leaves = shareCommitmentLeaves(shareCommitments, committeeSize);
  const index = participantIndex - 1;
  return { leaf: leaves[index], path: merklePath(leaves, index), root: merkleRoot(leaves) };
}
