// Organizer decryption share: Δ = sk_org · C1 plus a Chaum-Pedersen DLEQ.
//
// This is the browser-side half of the v2 decryption path. The organizer holds
// `sk_org` (the secret half of the application key `PK_aid = PK_ep + PK_org`)
// and publishes, for each ciphertext it wants opened, the share `Δ` together
// with a proof that the same `sk_org` relates `(G, PK_org)` and `(C1, Δ)`.
//
// The challenge is a keccak256 — not the Poseidon used by the committee's
// in-circuit DLEQ — for two reasons: the organizer runs in a browser, and the
// contract recomputes exactly this value so the committee's `decryptcombine`
// SNARK can take `e` from the transcript instead of hashing it in-circuit.
//
// Encoding (spec §3), byte for byte identical to `crypto/dleq` on the Go side
// and to `abi.encodePacked` in `DKGAppManager.sol`:
//
//   e = keccak256(
//         DOMAIN_ORGANIZER_SHARE_V1   // bytes32
//       ‖ eid                         // bytes12
//       ‖ aid                         // bytes32
//       ‖ uint256(ctIdx)              // uint256
//       ‖ PK_org.x ‖ PK_org.y         // uint256 ×2
//       ‖ C1.x ‖ C1.y                 // uint256 ×2
//       ‖ Δ.x ‖ Δ.y                   // uint256 ×2
//       ‖ A1.x ‖ A1.y ‖ A2.x ‖ A2.y   // uint256 ×4
//       ) mod q
//   z = w + e·sk_org mod q
//
// Verification: `z·G == A1 + e·PK_org` and `z·C1 == A2 + e·Δ`, with `z < q`,
// every point canonical / on curve / in the prime-order subgroup and `Δ` not
// the identity.
//
// Coordinate forms follow the SDK convention (see `crypto/babyjub-form.ts`):
// every coordinate that crosses the chain boundary — and therefore every
// coordinate hashed into `e` — is in gnark RTE form. `proveOrganizerShare`
// takes `C1` in the TE form this SDK's ElGamal layer produces and returns RTE
// points ready for calldata; `verifyOrganizerShare` takes RTE throughout, so
// event payloads can be checked as they arrive.

import { Base8, addPoint, mulPointEscalar, subOrder, type Point } from '@zk-kit/baby-jubjub';
import { encodePacked, keccak256, type Hex } from 'viem';
import { DomainOrganizerShareV1 } from './protocol.js';
import {
  pointEq,
  randomScalar,
  requireOnCurveSubgroup,
  requireScalarInRange,
  rteToTe,
  teToRte,
} from './schnorr.js';
import type { BabyJubPoint } from './types.js';

/** The organizer's Chaum-Pedersen proof, in on-chain (RTE) coordinates. */
export interface OrganizerShareProof {
  /** A1 = w·G */
  a1: BabyJubPoint;
  /** A2 = w·C1 */
  a2: BabyJubPoint;
  /** z = w + e·sk_org mod q */
  z: bigint;
}

/** `proveOrganizerShare` output: the share itself plus its proof. */
export interface OrganizerShare extends OrganizerShareProof {
  /** Δ = sk_org·C1, in on-chain (RTE) coordinates. */
  delta: BabyJubPoint;
}

/**
 * Re-derive the organizer-share challenge `e`. Every coordinate is the
 * on-chain (RTE) word, exactly as `submitOrganizerShare` sends it and
 * `OrganizerShareSubmitted` reports it.
 *
 * `ctIdx` is the on-chain ciphertext index (uint16) but enters the transcript
 * padded to a full uint256, matching `abi.encodePacked(uint256(ctIdx))`.
 */
export function organizerShareChallenge(
  epochId: Hex, // bytes12
  aid: Hex, // bytes32
  ctIdx: number | bigint,
  pkOrg: BabyJubPoint,
  c1: BabyJubPoint,
  delta: BabyJubPoint,
  a1: BabyJubPoint,
  a2: BabyJubPoint,
): bigint {
  const digest = keccak256(
    encodePacked(
      [
        'bytes32',
        'bytes12',
        'bytes32',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
        'uint256',
      ],
      [
        DomainOrganizerShareV1,
        epochId,
        aid,
        BigInt(ctIdx),
        pkOrg[0],
        pkOrg[1],
        c1[0],
        c1[1],
        delta[0],
        delta[1],
        a1[0],
        a1[1],
        a2[0],
        a2[1],
      ],
    ),
  );
  return BigInt(digest) % subOrder;
}

/**
 * Produce the organizer's decryption share for one ciphertext.
 *
 * @param epochId bytes12 epoch id the ciphertext lives in
 * @param aid     bytes32 application id
 * @param ctIdx   on-chain ciphertext index (1, 2, … per `(epochId, aid)`)
 * @param skOrg   the organizer secret used at registration; losing it makes
 *                every ciphertext of the application permanently undecryptable
 * @param c1      the ciphertext's `C1` in TE form (what `encrypt` returns and
 *                what `watchCiphertextSubmitted` hands back)
 * @param nonce   pins the witness `w`; for tests only — production callers
 *                must leave it undefined so a fresh value is drawn
 *
 * @returns `Δ`, `A1`, `A2` in on-chain (RTE) form plus the response `z`.
 */
export function proveOrganizerShare(
  epochId: Hex,
  aid: Hex,
  ctIdx: number | bigint,
  skOrg: bigint,
  c1: BabyJubPoint,
  nonce?: bigint,
): OrganizerShare {
  requireScalarInRange('organizer secret', skOrg);
  if (skOrg === 0n) throw new Error('organizer secret must be non-zero');

  const c1TE: Point<bigint> = [c1[0], c1[1]];
  requireOnCurveSubgroup('organizer share C1', c1TE);
  if (c1TE[0] === 0n && c1TE[1] === 1n) {
    throw new Error('organizer share C1: point is the identity');
  }

  const deltaTE = mulPointEscalar(c1TE, skOrg);
  if (deltaTE[0] === 0n && deltaTE[1] === 1n) {
    throw new Error('organizer share Δ is the identity — sk_org does not act on C1');
  }

  let w = nonce ?? randomScalar();
  while (w === 0n) w = randomScalar();
  requireScalarInRange('organizer share witness w', w);

  const a1TE = mulPointEscalar(Base8, w);
  const a2TE = mulPointEscalar(c1TE, w);

  // The transcript binds the on-chain (RTE) words.
  const pkOrg = teToRte(mulPointEscalar(Base8, skOrg)) as BabyJubPoint;
  const c1Rte = teToRte(c1TE) as BabyJubPoint;
  const delta = teToRte(deltaTE) as BabyJubPoint;
  const a1 = teToRte(a1TE) as BabyJubPoint;
  const a2 = teToRte(a2TE) as BabyJubPoint;

  const e = organizerShareChallenge(epochId, aid, ctIdx, pkOrg, c1Rte, delta, a1, a2);
  const z = (w + (e * skOrg) % subOrder) % subOrder;
  return { delta, a1, a2, z };
}

/**
 * Verify an organizer share against the application's registered `PK_org`.
 *
 * All points arrive in on-chain (RTE) form — pass the words from an
 * `OrganizerShareSubmitted` / `CiphertextSubmitted` event, or convert with
 * `fromTEtoRTE`. Returns `false` (never throws) for off-curve, off-subgroup,
 * identity or out-of-range inputs.
 *
 * `PK_org` must come from the application record, never from the submitter:
 * the share is only meaningful relative to the key the epoch key was extended
 * with at registration.
 */
export function verifyOrganizerShare(
  epochId: Hex,
  aid: Hex,
  ctIdx: number | bigint,
  pkOrg: BabyJubPoint,
  c1: BabyJubPoint,
  delta: BabyJubPoint,
  proof: OrganizerShareProof,
): boolean {
  let pkTE: Point<bigint>, c1TE: Point<bigint>, deltaTE: Point<bigint>;
  let a1TE: Point<bigint>, a2TE: Point<bigint>;
  try {
    pkTE = rteToTe([pkOrg[0], pkOrg[1]]);
    c1TE = rteToTe([c1[0], c1[1]]);
    deltaTE = rteToTe([delta[0], delta[1]]);
    a1TE = rteToTe([proof.a1[0], proof.a1[1]]);
    a2TE = rteToTe([proof.a2[0], proof.a2[1]]);
    requireOnCurveSubgroup('organizer PK_org', pkTE);
    requireOnCurveSubgroup('organizer share C1', c1TE);
    requireOnCurveSubgroup('organizer share Δ', deltaTE);
    requireOnCurveSubgroup('organizer share A1', a1TE);
    requireOnCurveSubgroup('organizer share A2', a2TE);
    requireScalarInRange('organizer share response z', proof.z);
  } catch {
    return false;
  }
  // A share that is the identity carries no information and would let a
  // malformed submission pass the two verifier equations for sk_org = 0.
  if (deltaTE[0] === 0n && deltaTE[1] === 1n) return false;

  const e = organizerShareChallenge(epochId, aid, ctIdx, pkOrg, c1, delta, proof.a1, proof.a2);

  const lhs1 = mulPointEscalar(Base8, proof.z);
  const rhs1 = addPoint(a1TE, mulPointEscalar(pkTE, e));
  if (!pointEq(lhs1, rhs1)) return false;

  const lhs2 = mulPointEscalar(c1TE, proof.z);
  const rhs2 = addPoint(a2TE, mulPointEscalar(deltaTE, e));
  return pointEq(lhs2, rhs2);
}
