// BabyJubJub affine-chart conversion between gnark's Reduced Twisted Edwards
// (RTE, A=-1) and Iden3/circomlib's Twisted Edwards (TE, A=168700).
//
// The two forms are isomorphic; only the X coordinate is rescaled. The
// constants and formulas here are vendored verbatim from
//   github.com/vocdoni/davinci-sdk  (src/crypto/BallotBuilder.ts)
// so this SDK stays compatible with davinci-sdk's wire format. The Go
// reference is davinci-node/crypto/ecc/format/twistededwards.go.
//
// Why this matters: davinci-dkg's on-chain validators (DKGManager._isOnBabyJubJub
// and the ZK circuits) operate in RTE. The collective public key returned by
// the contract is in RTE. This SDK's ElGamal layer (and the @zk-kit
// implementation it wraps) operates in TE. Mixing forms produces points that
// pass neither curve equation, hence InvalidCiphertext() reverts.
//
// Convention used by this SDK:
//   - Anything coming OUT of the SDK to the user (getCollectivePublicKey,
//     decrypted points, etc.) is in TE so it composes with circomlibjs/zk-kit.
//   - Anything going TO the chain (submitCiphertext) is converted back to RTE
//     just before sending so it satisfies the contract's curve check.

import type { BabyJubPoint } from '../types.js';

/** BN254 scalar field modulus. */
export const FIELD_MODULUS =
  21888242871839275222246405745257275088548364400416034343698204186575808495617n;

/**
 * Scaling factor used to translate the X coordinate between RTE and TE.
 * Matches davinci-sdk's `SCALING_FACTOR` exactly.
 */
export const SCALING_FACTOR =
  6360561867910373094066688120553762416144456282423235903351243436111059670888n;

function mod(a: bigint, m: bigint): bigint {
  const r = a % m;
  return r < 0n ? r + m : r;
}

function modInverse(a: bigint, m: bigint): bigint {
  // Extended Euclidean algorithm over bigints.
  let [oldR, r] = [mod(a, m), m];
  let [oldS, s] = [1n, 0n];
  while (r !== 0n) {
    const q = oldR / r;
    [oldR, r] = [r, oldR - q * r];
    [oldS, s] = [s, oldS - q * s];
  }
  if (oldR !== 1n) throw new Error('modInverse: value is not invertible');
  return mod(oldS, m);
}

/**
 * Convert a point from gnark RTE form (A=-1) to circomlib TE form (A=168700).
 *
 *   x_TE = x_RTE / (-f)
 *   y_TE = y_RTE
 */
export function fromRTEtoTE(x: bigint, y: bigint): [bigint, bigint] {
  const negF = mod(-SCALING_FACTOR, FIELD_MODULUS);
  const negFInv = modInverse(negF, FIELD_MODULUS);
  return [mod(x * negFInv, FIELD_MODULUS), y];
}

/**
 * Convert a point from circomlib TE form (A=168700) to gnark RTE form (A=-1).
 *
 *   x_RTE = x_TE * (-f)
 *   y_RTE = y_TE
 */
export function fromTEtoRTE(x: bigint, y: bigint): [bigint, bigint] {
  const negF = mod(-SCALING_FACTOR, FIELD_MODULUS);
  return [mod(x * negF, FIELD_MODULUS), y];
}

/** {@link fromRTEtoTE} accepting/returning the SDK's BabyJubPoint tuple. */
export function pointFromRTEtoTE(p: BabyJubPoint): BabyJubPoint {
  return fromRTEtoTE(p[0], p[1]);
}

/** {@link fromTEtoRTE} accepting/returning the SDK's BabyJubPoint tuple. */
export function pointFromTEtoRTE(p: BabyJubPoint): BabyJubPoint {
  return fromTEtoRTE(p[0], p[1]);
}
