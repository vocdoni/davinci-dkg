// Schnorr proofs of knowledge + Chaum-Pedersen DLEQ verification.
//
// Mirrors `crypto/schnorr/{operator,organizer}.go` and the in-circuit DLEQ
// transcript at `circuits/partialdecrypt/circuit.go`. Cross-impl byte
// equality is the load-bearing property; tests live at sdk/tests/schnorr.test.ts
// and consume vectors emitted by `cmd/operator-schnorr-vectors`.
//
// Verifiers only — proof generation lives in the Go node, since the only
// JS/TS callers are organizers (who use a wallet to sign) and the SDK never
// holds a long-lived Schnorr secret. `proveOrganizer` is provided as a
// convenience for tests; production organizers should use the Go path.

import {
  Base8,
  addPoint,
  inCurve,
  mulPointEscalar,
  Fr,
  subOrder,
  type Point,
} from '@zk-kit/baby-jubjub';
import { poseidon2, poseidon3, poseidon4, poseidon5, poseidon6 } from 'poseidon-lite';
import { keccak256, toHex, type Hex } from 'viem';
import {
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
  Role,
  type RoleValue,
} from './protocol.js';
import { fromRTEtoTE, fromTEtoRTE } from './crypto/babyjub-form.js';

// BN254 scalar field prime — matches `BabyJubJub.Q` in Solidity and the
// modulus the on-chain `_*SchnorrChallenge` reduces the keccak digest into.
export const BN254_Q =
  21888242871839275222246405745257275088548364400416034343698204186575808495617n;

// BabyJubJub subgroup order. Re-exported here so callers don't have to
// reach into derive.ts to range-check a Schnorr response.
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

function negate(p: Point<bigint>): Point<bigint> {
  // In twisted Edwards (a = -1), -(x, y) = (-x, y).
  return [Fr.neg(p[0]), p[1]];
}

function pointEq(a: Point<bigint>, b: Point<bigint>): boolean {
  return a[0] === b[0] && a[1] === b[1];
}

function requireOnCurveSubgroup(label: string, p: Point<bigint>): void {
  if (!inCurve(p)) throw new Error(`${label}: point not on BabyJubJub`);
  // Subgroup membership: r·P must equal identity (0, 1).
  const rp = mulPointEscalar(p, subOrder);
  if (!(rp[0] === 0n && rp[1] === 1n)) {
    throw new Error(`${label}: point not in prime-order subgroup`);
  }
}

// On-chain points (registry, application, partial-decrypt events) live in
// the gnark RTE form (a = -1). The Schnorr / DLEQ challenges are computed
// in-circuit and on-chain over THOSE coordinates, but @zk-kit/baby-jubjub
// implements curve arithmetic in the iden3/circomlib TE form (a = 168700).
// The two are isomorphic — convert the X coordinate at the API boundary.
function rteToTe(p: Point<bigint>): Point<bigint> {
  const [x, y] = fromRTEtoTE(p[0], p[1]);
  return [x, y];
}

function requireScalarInRange(label: string, s: bigint): void {
  if (s < 0n || s >= subOrder) {
    throw new Error(`${label}: scalar out of range [0, L)`);
  }
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
 *   inner = poseidon5(domainField, op, pubX, pubY, ax)
 *   c     = poseidon2(inner, ay)
 */
export function operatorSchnorrChallenge(
  operator: Hex,
  pubX: bigint,
  pubY: bigint,
  ax: bigint,
  ay: bigint,
): bigint {
  const domainField = reduceHexToBn254(DomainOperatorRegisterV1);
  const opField = BigInt(operator);
  const inner = poseidon5([domainField, opField, pubX, pubY, ax]);
  return poseidon2([inner, ay]);
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

// ─── organizer Schnorr (registerApplicationCoDec) ───────────────────────────

export interface OrganizerSchnorrProof {
  ax: bigint;
  ay: bigint;
  z: bigint;
}

/**
 * Re-derive the organizer Schnorr challenge. Mirrors
 * `_organizerSchnorrChallenge` in DKGManager.sol:
 *   inner = poseidon4(domainField, eidField, pkOrgX, pkOrgY)
 *   c     = poseidon4(inner, aidField, ax, ay)
 *
 * `eidField` is the bytes12 epoch id read as an unsigned integer (no
 * mod reduction — bytes12 always fits in bn254). `aidField` is the
 * bytes32 application id reduced mod bn254.
 */
export function organizerSchnorrChallenge(
  epochId: Hex,
  aid: Hex,
  pkOrgX: bigint,
  pkOrgY: bigint,
  ax: bigint,
  ay: bigint,
): bigint {
  const domainField = reduceHexToBn254(DomainOrganizerRegisterV1);
  const eidField = BigInt(epochId);
  const aidField = reduceHexToBn254(aid);
  const inner = poseidon4([domainField, eidField, pkOrgX, pkOrgY]);
  return poseidon4([inner, aidField, ax, ay]);
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
 * Convenience prover — production organizers should use the Go path. This
 * exists for tests and demos that need to round-trip a proof through the
 * SDK without spinning up the node. `nonce` is exposed so tests can pin
 * the witness; production callers must pass undefined to draw fresh.
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
  // the Poseidon transcript) is in RTE form — convert before hashing.
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

// ─── Chaum-Pedersen DLEQ (committee + organizer partial decryptions) ────────

export interface DleqPoints {
  base: Point<bigint>;      // C_1
  publicKey: Point<bigint>; // D_i (committee) or PK_org (organizer)
  delta: Point<bigint>;     // δ_i (committee) or Δ_org (organizer)
  a1: Point<bigint>;        // w·G
  a2: Point<bigint>;        // w·C_1
}

export interface DleqTranscriptInputs {
  epochId: Hex;             // bytes12
  aid: Hex;                 // bytes32
  ctIdx: number | bigint;
  role: RoleValue;
  participantIndex: number | bigint; // i (0 for organizer)
  points: DleqPoints;
}

/**
 * Re-derive the DLEQ Fiat-Shamir challenge. Mirrors the in-circuit
 * transcript at `circuits/partialdecrypt/circuit.go::Define`:
 *
 *   state = poseidon6(domain, eid, aid, ctIdx, role, i)
 *   c     = HashPointTuple(state, PK, C_1, δ, A1, A2)
 *
 * where `HashPointTuple` folds each point as
 *   current = poseidon3(current, p.x, p.y).
 */
export function dleqChallenge(t: DleqTranscriptInputs): bigint {
  const eidField = BigInt(t.epochId);
  const aidField = reduceHexToBn254(t.aid);
  const state = poseidon6([
    DOMAIN_PARTIAL_DECRYPT,
    eidField,
    aidField,
    BigInt(t.ctIdx),
    BigInt(t.role),
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
 * (G, PK) and (C_1, δ). Returns `true` only if every point sits on the
 * prime-order subgroup, the response is in range, and the two verifier
 * equations hold.
 *
 * Used by the SDK monitor (audit mode) and by tests; the on-chain combine
 * verifier rejects malformed shares before this would ever fail in a
 * production flow, but offline auditors may want to re-check.
 */
export function verifyDleq(
  t: DleqTranscriptInputs,
  z: bigint,
): boolean {
  if (t.role !== Role.Committee && t.role !== Role.Organizer) return false;
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

// ─── small RNG for the convenience prover ───────────────────────────────────

function randomScalar(): bigint {
  const bytes = globalThis.crypto.getRandomValues(new Uint8Array(32));
  let bi = 0n;
  for (let i = 0; i < bytes.length; i++) bi += BigInt(bytes[i]) << BigInt(8 * i);
  return bi % subOrder;
}

// Touch keccak256/toHex so they remain importable without unused warnings
// when the future organizer-side ABI helpers move into this module.
void keccak256;
void toHex;
