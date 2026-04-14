// ElGamal encryption on the BabyJubJub curve.
//
// Uses @zk-kit/baby-jubjub which is pure TypeScript, synchronous, and
// browser-compatible without any Node.js polyfills.

import {
  mulPointEscalar,
  addPoint,
  packPoint as _packPoint,
  unpackPoint as _unpackPoint,
  Base8,
  subOrder,
  Fr,
} from '@zk-kit/baby-jubjub';
import type { BabyJubPoint, ElGamalCiphertext } from '../types.js';

export interface ElGamal {
  encrypt: (msg: bigint, pubKey: BabyJubPoint, k?: bigint) => ElGamalCiphertext;
  decrypt: (ciphertext: ElGamalCiphertext, privKey: bigint) => bigint;
  generateKeyPair: () => { privKey: bigint; pubKey: BabyJubPoint };
  randomScalar: () => bigint;
  /** Encode a point as a single compressed bigint */
  packPoint: (p: BabyJubPoint) => bigint;
  /** Decode a compressed bigint back to a point */
  unpackPoint: (packed: bigint) => BabyJubPoint;
  /** Multiply a curve point by a scalar */
  mulPoint: (point: BabyJubPoint, scalar: bigint) => BabyJubPoint;
  /** Add two curve points */
  addPoint: (a: BabyJubPoint, b: BabyJubPoint) => BabyJubPoint;
}

// Singleton — the library is sync so construction is instant.
let _elgamal: ElGamal | undefined;

/**
 * Return the ElGamal instance backed by @zk-kit/baby-jubjub.
 *
 * The function is kept async for API compatibility with existing callers
 * (tests, flow.ts) — it resolves immediately.
 */
export async function buildElGamal(): Promise<ElGamal> {
  if (_elgamal) return _elgamal;

  // @zk-kit/baby-jubjub works directly with [bigint, bigint] points, so
  // BabyJubPoint ≡ Point<bigint> — no conversion layer needed.

  function randomScalar(): bigint {
    const bytes = globalThis.crypto.getRandomValues(new Uint8Array(32));
    let bi = 0n;
    for (let i = 0; i < bytes.length; i++) bi += BigInt(bytes[i]) << BigInt(8 * i);
    return bi % subOrder;
  }

  function generateKeyPair(): { privKey: bigint; pubKey: BabyJubPoint } {
    const privKey = randomScalar();
    const pubKey = mulPointEscalar(Base8, privKey) as BabyJubPoint;
    return { privKey, pubKey };
  }

  /**
   * ElGamal encrypt.
   *
   * c1 = k * G
   * c2 = m*G + k*PubKey
   *
   * @param k  Optional ephemeral scalar; random when omitted.
   */
  function encrypt(msg: bigint, pubKey: BabyJubPoint, k?: bigint): ElGamalCiphertext {
    const kVal = k ?? randomScalar();
    const c1 = mulPointEscalar(Base8, kVal) as BabyJubPoint;
    const s  = mulPointEscalar(pubKey, kVal) as BabyJubPoint;
    const mPoint = mulPointEscalar(Base8, msg) as BabyJubPoint;
    const c2 = addPoint(mPoint, s) as BabyJubPoint;
    return { c1, c2 };
  }

  /**
   * ElGamal decrypt via brute-force discrete log.
   *
   * s      = privKey * c1
   * mPoint = c2 - s
   *
   * Only suitable for small messages (< 2^20).
   */
  function decrypt(ciphertext: ElGamalCiphertext, privKey: bigint): bigint {
    const s = mulPointEscalar(ciphertext.c1, privKey) as BabyJubPoint;
    // Negate s: in twisted Edwards, -(x, y) = (-x, y)
    const negS: BabyJubPoint = [Fr.neg(s[0]), s[1]];
    const mPoint = addPoint(ciphertext.c2, negS) as BabyJubPoint;

    let candidate: BabyJubPoint = [0n, 1n]; // identity point
    for (let i = 0n; i < 1_048_576n; i++) {
      if (candidate[0] === mPoint[0] && candidate[1] === mPoint[1]) return i;
      candidate = addPoint(candidate, Base8) as BabyJubPoint;
    }
    throw new Error('decrypt: message out of brute-force range (> 2^20)');
  }

  function mulPoint(point: BabyJubPoint, scalar: bigint): BabyJubPoint {
    return mulPointEscalar(point, scalar) as BabyJubPoint;
  }

  function packPoint(p: BabyJubPoint): bigint {
    return _packPoint(p);
  }

  function unpackPoint(packed: bigint): BabyJubPoint {
    const p = _unpackPoint(packed);
    if (!p) throw new Error('unpackPoint: invalid packed point');
    return p as BabyJubPoint;
  }

  _elgamal = {
    randomScalar, generateKeyPair,
    mulPoint, addPoint: (a, b) => addPoint(a, b) as BabyJubPoint,
    encrypt, decrypt,
    packPoint, unpackPoint,
  };
  return _elgamal;
}
