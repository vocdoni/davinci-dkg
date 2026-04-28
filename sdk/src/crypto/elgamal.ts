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
   * ElGamal decrypt via baby-step / giant-step (BSGS) discrete log.
   *
   *   s      = privKey * c1
   *   mPoint = c2 - s
   *   m      = dlog(mPoint)
   *
   * The production threshold-decryption path runs in the Go node and supports
   * plaintexts up to 2^50; this client-side single-key recovery is only
   * intended for tests, demos, and direct (non-threshold) use, so the cap is
   * the more browser-friendly **2^32** (~4.3 billion). Lookup table is built
   * lazily on first call (~16 MB heap, ~1–2 s in the browser) and cached.
   *
   * For the on-chain protocol the matching limit is `MAX_DLOG_PLAINTEXT` in
   * `cmd/davinci-dkg-node/dlog.go` (2^50). The two values are deliberately
   * different — the threshold combine runs on a server with more headroom.
   */
  function decrypt(ciphertext: ElGamalCiphertext, privKey: bigint): bigint {
    const s = mulPointEscalar(ciphertext.c1, privKey) as BabyJubPoint;
    // Negate s: in twisted Edwards, -(x, y) = (-x, y)
    const negS: BabyJubPoint = [Fr.neg(s[0]), s[1]];
    const mPoint = addPoint(ciphertext.c2, negS) as BabyJubPoint;
    return dlogBSGS(mPoint);
  }

  /**
   * Largest plaintext (exclusive) the client-side `decrypt` can recover.
   * Mirrors the SDK's documented contract; values at or above this throw.
   */
  const MAX_CLIENT_DLOG_PLAINTEXT = 1n << 32n;
  const BSGS_M = 1n << 16n; // ⌈√MAX_CLIENT_DLOG_PLAINTEXT⌉ = 2^16 = 65,536.

  // Lazy BSGS table. Keys are point.x serialised as decimal strings — BJJ's
  // x-coordinate uniquely identifies a point on the prime-order subgroup we
  // operate on, so we don't need to mix in y. Decimal stringify is fine
  // here: BigInt → string is fast and the alternative (toString(16) padded)
  // adds nothing because we never expose the keys.
  let bsgsTable: Map<string, number> | undefined;
  let bsgsNegM: BabyJubPoint | undefined;

  function initBsgsTable(): void {
    if (bsgsTable !== undefined) return;
    const t = new Map<string, number>();
    let cur: BabyJubPoint = [0n, 1n]; // identity
    for (let i = 0; BigInt(i) < BSGS_M; i++) {
      t.set(cur[0].toString(), i);
      cur = addPoint(cur, Base8) as BabyJubPoint;
    }
    bsgsTable = t;
    // Giant step: M = m·G, then negate so the search loop adds rather
    // than subtracts on every iteration.
    const M = mulPointEscalar(Base8, BSGS_M) as BabyJubPoint;
    bsgsNegM = [Fr.neg(M[0]), M[1]];
  }

  function dlogBSGS(target: BabyJubPoint): bigint {
    initBsgsTable();
    const table = bsgsTable!;
    const negM = bsgsNegM!;
    let cur = target;
    for (let j = 0n; j < BSGS_M; j++) {
      const hit = table.get(cur[0].toString());
      if (hit !== undefined) return j * BSGS_M + BigInt(hit);
      cur = addPoint(cur, negM) as BabyJubPoint;
    }
    throw new Error(
      `decrypt: plaintext out of range (>= 2^32 ≈ ${MAX_CLIENT_DLOG_PLAINTEXT.toString()}). ` +
        'Threshold decryption on the Go committee handles up to 2^50.',
    );
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
