// Per-application key derivation helpers (PLAN.md §7.2 / paper §4.3).
//
// These mirror the on-chain derivation in DKGManager.registerApplication
// and the off-chain derivation in internal/protocol/protocol.go. Cross-impl
// agreement is the load-bearing property — any drift here breaks every
// per-application encryption.

import { keccak256, encodePacked, type Hex } from 'viem';

/**
 * BabyJubJub subgroup (prime-order) cardinality. Matches `BabyJubJub.SUBGROUP_ORDER`
 * in the Solidity library and `twistededwards.GetEdwardsCurve().Order` in
 * gnark-crypto.
 */
export const SUBGROUP_ORDER =
  2736030358979909402780800718157159386076813972158567259200215660948447373041n;

/**
 * Compute the per-application derivation tag
 *
 *   S = keccak256(eid || PK_ep.x || PK_ep.y || aid) mod L
 *
 * where L is the BabyJubJub subgroup order. Matches the on-chain derivation
 * in DKGManager.registerApplication (P8). Both inputs are encoded packed
 * with EIP-191 / abi.encodePacked semantics — `eid` is bytes12, `PK_ep`'s
 * coordinates are uint256, `aid` is bytes32.
 *
 * **Coordinate form:** `pkEpX` / `pkEpY` MUST be the on-chain (gnark RTE)
 * representation, not the iden3/circomlib TE form returned by
 * `DKGClient.getCollectivePublicKey`. The contract hashes the stored RTE
 * coordinates; passing TE coordinates here produces a different `S` and
 * the resulting application registration will refer to a different key.
 * Use `crypto/babyjub-form.fromTEtoRTE` to convert at the SDK boundary.
 */
export function computeS(
  epochId: Hex,        // bytes12
  pkEpX: bigint,       // RTE form (on-chain)
  pkEpY: bigint,       // RTE form (on-chain)
  aid: Hex,            // bytes32
): bigint {
  const digest = keccak256(
    encodePacked(
      ['bytes12', 'uint256', 'uint256', 'bytes32'],
      [epochId, pkEpX, pkEpY, aid],
    ),
  );
  return BigInt(digest) % SUBGROUP_ORDER;
}

/**
 * Derive the per-application public key from the epoch public key and the
 * application's mode-specific correction factor.
 *
 *  Mode 0 (public derivation):    PK_aid = PK_ep + S · G
 *  Mode 1 (organizer co-decryption): PK_aid = PK_ep + PK_org
 *
 * Returns the (x, y) BabyJubJub coordinates as bigints. Callers must
 * supply `mode = 0` with a non-zero `S` derived via `computeS`, or
 * `mode = 1` with the on-chain organizer public key.
 *
 * Note: this helper requires a BabyJubJub point-add and (in mode 0) a
 * scalar mul on the SDK side; we delegate to crypto/elgamal which
 * already includes the relevant arithmetic, but the surface here lets
 * callers stay agnostic to those primitives.
 */
export interface DerivePKAppInput {
  mode: 0 | 1;
  pkEpX: bigint;
  pkEpY: bigint;
  s?: bigint;        // required when mode === 0
  pkOrgX?: bigint;   // required when mode === 1
  pkOrgY?: bigint;   // required when mode === 1
}

/**
 * Validate the input shape for `derivePKApp`. The actual point arithmetic
 * lives in `crypto/elgamal.ts` to avoid duplicating BabyJubJub code; this
 * helper only enforces the mode contract.
 */
export function validateDerivePKAppInput(input: DerivePKAppInput): void {
  if (input.mode === 0) {
    if (input.s === undefined || input.s === 0n) {
      throw new Error('mode 0 (public derivation) requires non-zero S');
    }
  } else if (input.mode === 1) {
    if (input.pkOrgX === undefined || input.pkOrgY === undefined) {
      throw new Error('mode 1 (organizer co-decryption) requires PK_org');
    }
  } else {
    throw new Error(`unknown application mode: ${input.mode}`);
  }
}
