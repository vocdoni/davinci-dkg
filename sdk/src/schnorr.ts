// Schnorr proofs of knowledge + Chaum-Pedersen DLEQ verification.
//
// Mirrors `crypto/schnorr/{operator,organizer}.go` and the in-circuit DLEQ
// transcript at `circuits/partialdecrypt/circuit.go`. Cross-impl byte
// equality is the load-bearing property; tests live at sdk/tests/schnorr.test.ts
// and consume vectors emitted by `cmd/protocol-vectors`.
//
// Provers and verifiers for both registration roles. `proveOperator` /
// `proveOrganizer` need the corresponding secret scalar; the committee DLEQ
// (`dleqChallenge`, `verifyDleq`) is verify-only here because committee
// partial decryptions are produced by the Go node together with their
// Groth16 wrapper.
//
// The organizer's decryption-share DLEQ lives in `dleq.ts`: it is keccak-based
// (not Poseidon) and the SDK both proves and verifies it, because the
// organizer runs in a browser.

import {
  Base8,
  addPoint,
  inCurve,
  mulPointEscalar,
  subOrder,
  type Point,
} from '@zk-kit/baby-jubjub';
import { poseidon3, poseidon5 } from 'poseidon-lite';
import { keccak256, type Hex } from 'viem';
import {
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
} from './protocol.js';
import { fromRTEtoTE, fromTEtoRTE } from './crypto/babyjub-form.js';

// BN254 scalar field prime — matches `BabyJubJub.Q` in Solidity and the
// modulus the on-chain `_*SchnorrChallenge` reduces the keccak digest into.
export const BN254_Q =
  21888242871839275222246405745257275088548364400416034343698204186575808495617n;

// BabyJubJub prime-order subgroup order `q`. Every Schnorr / DLEQ response is
// reduced into `[0, q)` and every point the protocol accepts must live in this
// subgroup.
export const SUBGROUP_ORDER = subOrder;

// `davinci-dkg/partial-decrypt/v1` reduced into bn254 scalar field. Mirrors
// `hash.DomainValue(hash.DomainPartialDecrypt)` in the Go circuit witness
// builder. Computed once at module load so verification is allocation-free.
const PARTIAL_DECRYPT_DOMAIN_BYTES = new TextEncoder().encode(
  'davinci-dkg/partial-decrypt/v1',
);
export const DOMAIN_PARTIAL_DECRYPT = (() => {
  let acc = 0n;
  for (const b of PARTIAL_DECRYPT_DOMAIN_BYTES) acc = (acc << 8n) | BigInt(b);
  return acc % BN254_Q;
})();

// ─── shared helpers ─────────────────────────────────────────────────────────
//
// Exported for `dleq.ts`; not re-exported from the package entry points.

/** @internal Curve-point equality on affine coordinates. */
export function pointEq(a: Point<bigint>, b: Point<bigint>): boolean {
  return a[0] === b[0] && a[1] === b[1];
}

/** @internal Throws unless `p` is on BabyJubJub and in the prime-order subgroup. */
export function requireOnCurveSubgroup(label: string, p: Point<bigint>): void {
  if (!inCurve(p)) throw new Error(`${label}: point not on BabyJubJub`);
  // Subgroup membership: q·P must equal identity (0, 1).
  const qp = mulPointEscalar(p, subOrder);
  if (!(qp[0] === 0n && qp[1] === 1n)) {
    throw new Error(`${label}: point not in prime-order subgroup`);
  }
}

// On-chain points (registry, application, partial-decrypt events) live in
// the gnark RTE form (a = -1). The Schnorr / DLEQ challenges are computed
// in-circuit and on-chain over THOSE coordinates, but @zk-kit/baby-jubjub
// implements curve arithmetic in the iden3/circomlib TE form (a = 168700).
// The two are isomorphic — convert the X coordinate at the API boundary.
/** @internal Convert an on-chain (RTE) point into the TE form used for arithmetic. */
export function rteToTe(p: Point<bigint>): Point<bigint> {
  const [x, y] = fromRTEtoTE(p[0], p[1]);
  return [x, y];
}

/** @internal Convert a TE point into the on-chain (RTE) form. */
export function teToRte(p: Point<bigint>): Point<bigint> {
  const [x, y] = fromTEtoRTE(p[0], p[1]);
  return [x, y];
}

/** @internal Throws unless `s` is a canonical scalar in `[0, q)`. */
export function requireScalarInRange(label: string, s: bigint): void {
  if (s < 0n || s >= subOrder) {
    throw new Error(`${label}: scalar out of range [0, q)`);
  }
}

/** @internal Uniform scalar in `[0, q)` drawn from the platform CSPRNG. */
export function randomScalar(): bigint {
  const bytes = globalThis.crypto.getRandomValues(new Uint8Array(32));
  let bi = 0n;
  for (let i = 0; i < bytes.length; i++) bi += BigInt(bytes[i]) << BigInt(8 * i);
  return bi % subOrder;
}

/** @internal big-endian 32-byte encoding of a uint256. */
export function uint256ToBytes32(v: bigint): Uint8Array {
  const out = new Uint8Array(32);
  let x = v;
  for (let i = 31; i >= 0; i--) {
    out[i] = Number(x & 0xffn);
    x >>= 8n;
  }
  return out;
}

/** @internal Strict-length hex → bytes (left-padded if shorter than requested). */
export function hexToBytes(h: Hex, expectedLen: number): Uint8Array {
  let s = h.startsWith('0x') ? h.slice(2) : h;
  if (s.length > expectedLen * 2) {
    throw new Error(`hexToBytes: input ${h} longer than ${expectedLen} bytes`);
  }
  if (s.length < expectedLen * 2) s = s.padStart(expectedLen * 2, '0');
  const out = new Uint8Array(expectedLen);
  for (let i = 0; i < expectedLen; i++) {
    out[i] = parseInt(s.slice(i * 2, i * 2 + 2), 16);
  }
  return out;
}

function reduceHexToBn254(h: Hex): bigint {
  return BigInt(h) % BN254_Q;
}

// ─── operator Schnorr (registerKey / updateKey) ─────────────────────────────

export interface OperatorSchnorrProof {
  ax: bigint;
  ay: bigint;
  z: bigint;
}

/**
 * Re-derive the operator Schnorr challenge. Mirrors
 * `_operatorSchnorrChallenge` in DKGRegistry.sol:
 *   c = keccak256(domain || operator || pubX || pubY || ax || ay) mod q
 *
 * keccak256 is used in place of Poseidon — the challenge is only verified
 * on-chain, so SNARK-native compatibility is unnecessary.
 */
export function operatorSchnorrChallenge(
  operator: Hex,
  pubX: bigint,
  pubY: bigint,
  ax: bigint,
  ay: bigint,
): bigint {
  const buf = new Uint8Array(32 + 20 + 32 * 4);
  // domain (bytes32)
  buf.set(hexToBytes(DomainOperatorRegisterV1, 32), 0);
  // operator (address, 20 bytes)
  buf.set(hexToBytes(operator, 20), 32);
  // pubX, pubY, ax, ay (each uint256 big-endian)
  let off = 32 + 20;
  for (const v of [pubX, pubY, ax, ay]) {
    buf.set(uint256ToBytes32(v), off);
    off += 32;
  }
  const h = BigInt(keccak256(buf));
  return h % subOrder;
}

/**
 * Verify an operator Schnorr proof of knowledge of the discrete log of
 * `(pubX, pubY)` on BabyJubJub, bound to `operator`. Returns `true` only
 * if the public key sits on the prime-order subgroup, the witness `(ax, ay)`
 * sits on the curve, and `z·G == A + c·PK`.
 */
export function verifyOperatorSchnorr(
  operator: Hex,
  pubX: bigint,
  pubY: bigint,
  proof: OperatorSchnorrProof,
): boolean {
  // (pubX, pubY) and (ax, ay) arrive as RTE coords (the on-chain form);
  // convert into TE for curve math.
  let PK: Point<bigint>, A: Point<bigint>;
  try {
    PK = rteToTe([pubX, pubY]);
    A = rteToTe([proof.ax, proof.ay]);
    requireOnCurveSubgroup('operator pubkey', PK);
    requireOnCurveSubgroup('operator witness A', A);
    requireScalarInRange('operator response z', proof.z);
  } catch {
    return false;
  }
  // Challenge is computed over the RTE (on-chain) coordinates.
  const c = operatorSchnorrChallenge(operator, pubX, pubY, proof.ax, proof.ay);
  const lhs = mulPointEscalar(Base8, proof.z);
  const rhs = addPoint(A, mulPointEscalar(PK, c % subOrder));
  return pointEq(lhs, rhs);
}

/**
 * Build an operator Schnorr proof of knowledge of `privateKey` such that
 * `PK = privateKey · G` on BabyJubJub, bound to `operator` (the EVM
 * address that will own the registry slot). Mirrors Go-side
 * `crypto/schnorr.ProveOperatorRegister`.
 *
 * Returns `(pkX, pkY, proof)` with all coordinates in RTE form so callers
 * can pass them directly to `registerKey` / `updateKey`. The witness `w`
 * is drawn from a CSPRNG; the optional `nonce` parameter exists only so
 * tests can pin the value for reproducibility.
 */
export function proveOperator(
  privateKey: bigint,
  operator: Hex,
  nonce?: bigint,
): { pubX: bigint; pubY: bigint; proof: OperatorSchnorrProof } {
  requireScalarInRange('operator private key', privateKey);
  if (privateKey === 0n) throw new Error('operator private key must be non-zero');

  const PK_TE = mulPointEscalar(Base8, privateKey);
  const w = nonce ?? randomScalar();
  if (w === 0n) throw new Error('zero witness drawn — caller must retry');
  const A_TE = mulPointEscalar(Base8, w);

  // Challenge transcript reads the on-chain (RTE) coordinates.
  const [pkX, pkY] = fromTEtoRTE(PK_TE[0], PK_TE[1]);
  const [aX, aY] = fromTEtoRTE(A_TE[0], A_TE[1]);

  const c = operatorSchnorrChallenge(operator, pkX, pkY, aX, aY);
  const z = (w + ((c % subOrder) * privateKey) % subOrder) % subOrder;
  return { pubX: pkX, pubY: pkY, proof: { ax: aX, ay: aY, z } };
}

// ─── organizer Schnorr (registerApplication proof of possession) ────────────

export interface OrganizerSchnorrProof {
  ax: bigint;
  ay: bigint;
  z: bigint;
}

/**
 * Re-derive the organizer Schnorr challenge. Mirrors
 * `_organizerSchnorrChallenge` in DKGAppManager.sol:
 *   c = keccak256(domain || epochId || aid || pkOrgX || pkOrgY || ax || ay) mod q
 *
 * keccak256 instead of Poseidon — see operatorSchnorrChallenge for rationale.
 */
export function organizerSchnorrChallenge(
  epochId: Hex,
  aid: Hex,
  pkOrgX: bigint,
  pkOrgY: bigint,
  ax: bigint,
  ay: bigint,
): bigint {
  const buf = new Uint8Array(32 + 12 + 32 * 5);
  buf.set(hexToBytes(DomainOrganizerRegisterV1, 32), 0);
  buf.set(hexToBytes(epochId, 12), 32);
  let off = 32 + 12;
  buf.set(hexToBytes(aid, 32), off);
  off += 32;
  for (const v of [pkOrgX, pkOrgY, ax, ay]) {
    buf.set(uint256ToBytes32(v), off);
    off += 32;
  }
  const h = BigInt(keccak256(buf));
  return h % subOrder;
}

export function verifyOrganizerSchnorr(
  epochId: Hex,
  aid: Hex,
  pkOrgX: bigint,
  pkOrgY: bigint,
  proof: OrganizerSchnorrProof,
): boolean {
  let PK: Point<bigint>, A: Point<bigint>;
  try {
    PK = rteToTe([pkOrgX, pkOrgY]);
    A = rteToTe([proof.ax, proof.ay]);
    requireOnCurveSubgroup('organizer pubkey', PK);
    requireOnCurveSubgroup('organizer witness A', A);
    requireScalarInRange('organizer response z', proof.z);
  } catch {
    return false;
  }
  const c = organizerSchnorrChallenge(epochId, aid, pkOrgX, pkOrgY, proof.ax, proof.ay);
  const lhs = mulPointEscalar(Base8, proof.z);
  const rhs = addPoint(A, mulPointEscalar(PK, c % subOrder));
  return pointEq(lhs, rhs);
}

/**
 * Build the organizer's proof of possession of `skOrg`, bound to
 * `(epochId, aid)`. `DKGWriter.registerApplication` calls this for you; it is
 * exported so integrators can pre-compute a registration payload offline.
 *
 * Returns `PK_org` and the proof in RTE (on-chain) form. `nonce` is exposed so
 * tests can pin the witness; production callers must leave it undefined so a
 * fresh value is drawn.
 */
export function proveOrganizer(
  skOrg: bigint,
  epochId: Hex,
  aid: Hex,
  nonce?: bigint,
): { pkOrgX: bigint; pkOrgY: bigint; proof: OrganizerSchnorrProof } {
  requireScalarInRange('organizer secret', skOrg);
  if (skOrg === 0n) throw new Error('organizer secret must be non-zero');

  // Curve math is done in TE form, but the on-chain artifact (and therefore
  // the keccak transcript) is in RTE form — convert before hashing.
  const PK_TE = mulPointEscalar(Base8, skOrg);
  const w = nonce ?? randomScalar();
  if (w === 0n) throw new Error('zero witness drawn — caller must retry');
  const A_TE = mulPointEscalar(Base8, w);

  const [pkX, pkY] = fromTEtoRTE(PK_TE[0], PK_TE[1]);
  const [aX, aY] = fromTEtoRTE(A_TE[0], A_TE[1]);

  const c = organizerSchnorrChallenge(epochId, aid, pkX, pkY, aX, aY);
  const z = (w + ((c % subOrder) * skOrg) % subOrder) % subOrder;
  return { pkOrgX: pkX, pkOrgY: pkY, proof: { ax: aX, ay: aY, z } };
}

// ─── Chaum-Pedersen DLEQ (committee partial decryptions) ────────────────────

export interface DleqPoints {
  base: Point<bigint>;      // C_1
  publicKey: Point<bigint>; // D_i
  delta: Point<bigint>;     // δ_i
  a1: Point<bigint>;        // w·G
  a2: Point<bigint>;        // w·C_1
}

export interface DleqTranscriptInputs {
  epochId: Hex;             // bytes12
  aid: Hex;                 // bytes32
  ctIdx: number | bigint;
  participantIndex: number | bigint; // i
  points: DleqPoints;
}

/**
 * Re-derive the committee DLEQ Fiat-Shamir challenge. Mirrors the in-circuit
 * transcript at `circuits/partialdecrypt/circuit.go::Define`:
 *
 *   state = poseidon5(domain, eid, aid, ctIdx, i)
 *   c     = HashPointTuple(state, D_i, C_1, δ, A1, A2)
 *
 * where `HashPointTuple` folds each point as
 *   current = poseidon3(current, p.x, p.y).
 */
export function dleqChallenge(t: DleqTranscriptInputs): bigint {
  const eidField = BigInt(t.epochId);
  const aidField = reduceHexToBn254(t.aid);
  const state = poseidon5([
    DOMAIN_PARTIAL_DECRYPT,
    eidField,
    aidField,
    BigInt(t.ctIdx),
    BigInt(t.participantIndex),
  ]);
  let cur = state;
  for (const p of [
    t.points.publicKey,
    t.points.base,
    t.points.delta,
    t.points.a1,
    t.points.a2,
  ]) {
    cur = poseidon3([cur, p[0], p[1]]);
  }
  return cur;
}

/**
 * Verify a Chaum-Pedersen DLEQ partial-decryption proof. The `response z`
 * proves knowledge of a single discrete log relating two pairs of points
 * (G, D_i) and (C_1, δ_i). Returns `true` only if every point sits on the
 * prime-order subgroup, the response is in range, and the two verifier
 * equations hold.
 *
 * Used by the SDK monitor (audit mode) and by tests; the on-chain
 * partial-decrypt verifier rejects malformed shares before this would ever
 * fail in a production flow, but offline auditors may want to re-check.
 */
export function verifyDleq(
  t: DleqTranscriptInputs,
  z: bigint,
): boolean {
  // Inputs come in RTE form (matching on-chain). Convert each to TE for
  // curve arithmetic; the Poseidon transcript (dleqChallenge) consumes the
  // RTE coords as-is.
  let base: Point<bigint>, publicKey: Point<bigint>, delta: Point<bigint>;
  let a1: Point<bigint>, a2: Point<bigint>;
  try {
    base = rteToTe(t.points.base);
    publicKey = rteToTe(t.points.publicKey);
    delta = rteToTe(t.points.delta);
    a1 = rteToTe(t.points.a1);
    a2 = rteToTe(t.points.a2);
    requireOnCurveSubgroup('DLEQ base C_1', base);
    requireOnCurveSubgroup('DLEQ pubkey', publicKey);
    requireOnCurveSubgroup('DLEQ delta', delta);
    requireOnCurveSubgroup('DLEQ A1', a1);
    requireOnCurveSubgroup('DLEQ A2', a2);
    requireScalarInRange('DLEQ response z', z);
  } catch {
    return false;
  }

  const c = dleqChallenge(t) % subOrder;
  const lhs1 = mulPointEscalar(Base8, z);
  const rhs1 = addPoint(a1, mulPointEscalar(publicKey, c));
  if (!pointEq(lhs1, rhs1)) return false;

  const lhs2 = mulPointEscalar(base, z);
  const rhs2 = addPoint(a2, mulPointEscalar(delta, c));
  return pointEq(lhs2, rhs2);
}
