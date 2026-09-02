// Organizer decryption share (spec §3): prove / verify round trip, the exact
// keccak encoding of the challenge, and every tamper the contract and the
// combine circuit rely on being caught.
//
// The encoding assertion is the load-bearing one: `e` is recomputed by the
// Solidity app manager and consumed as a public transcript word by the
// `decryptcombine` circuit, so a drift here silently breaks decryption for
// every application.

import { describe, it, expect } from 'vitest';
import { keccak256, type Hex } from 'viem';
import { Base8, addPoint, mulPointEscalar, subOrder, type Point } from '@zk-kit/baby-jubjub';
import {
  organizerShareChallenge,
  proveOrganizerShare,
  verifyOrganizerShare,
} from '../src/dleq';
import { DomainOrganizerShareV1 } from '../src/protocol';
import { fromTEtoRTE } from '../src/crypto/babyjub-form';
import type { BabyJubPoint } from '../src/types';

const eid = '0x112233440000000000000007' as Hex;
const aid = '0xa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebf' as Hex;
const ctIdx = 3;

const skOrg = 1234567890123456789n;
const ephem = 999999n; // C1 = ephem·G
const w = 7777777n;    // pinned DLEQ witness

function teToRte(p: Point<bigint>): BabyJubPoint {
  const [x, y] = fromTEtoRTE(p[0], p[1]);
  return [x, y];
}

/** C1 in TE form — what the SDK's ElGamal layer produces. */
const c1TE: BabyJubPoint = mulPointEscalar(Base8, ephem) as BabyJubPoint;
const pkOrgRTE = teToRte(mulPointEscalar(Base8, skOrg));
const c1RTE = teToRte(c1TE);

// ─── challenge encoding ─────────────────────────────────────────────────────

/**
 * Second, independent implementation of the §3 preimage: a hand-packed
 * byte buffer rather than viem's `encodePacked`. If the two agree we know the
 * type list handed to `encodePacked` really lays out
 * `32 + 12 + 32 + 32 + 10·32` bytes in the documented order.
 */
function manualChallenge(
  epochId: Hex,
  aidHex: Hex,
  idx: number,
  pkOrg: BabyJubPoint,
  c1: BabyJubPoint,
  delta: BabyJubPoint,
  a1: BabyJubPoint,
  a2: BabyJubPoint,
): bigint {
  const words = [
    BigInt(idx),
    pkOrg[0], pkOrg[1],
    c1[0], c1[1],
    delta[0], delta[1],
    a1[0], a1[1],
    a2[0], a2[1],
  ];
  const buf = new Uint8Array(32 + 12 + 32 + words.length * 32);
  const hexBytes = (h: Hex, len: number) => {
    const clean = h.slice(2).padStart(len * 2, '0');
    const out = new Uint8Array(len);
    for (let i = 0; i < len; i++) out[i] = parseInt(clean.slice(i * 2, i * 2 + 2), 16);
    return out;
  };
  buf.set(hexBytes(DomainOrganizerShareV1, 32), 0);
  buf.set(hexBytes(epochId, 12), 32);
  buf.set(hexBytes(aidHex, 32), 44);
  let off = 76;
  for (const v of words) {
    let x = v;
    for (let i = 31; i >= 0; i--) {
      buf[off + i] = Number(x & 0xffn);
      x >>= 8n;
    }
    off += 32;
  }
  return BigInt(keccak256(buf)) % subOrder;
}

describe('organizer share challenge', () => {
  const delta = teToRte(mulPointEscalar(c1TE, skOrg));
  const a1 = teToRte(mulPointEscalar(Base8, w));
  const a2 = teToRte(mulPointEscalar(c1TE, w));

  it('matches a hand-packed keccak preimage byte for byte', () => {
    expect(
      organizerShareChallenge(eid, aid, ctIdx, pkOrgRTE, c1RTE, delta, a1, a2),
    ).toBe(manualChallenge(eid, aid, ctIdx, pkOrgRTE, c1RTE, delta, a1, a2));
  });

  it('is reduced into the BabyJubJub subgroup order', () => {
    const e = organizerShareChallenge(eid, aid, ctIdx, pkOrgRTE, c1RTE, delta, a1, a2);
    expect(e).toBeGreaterThanOrEqual(0n);
    expect(e).toBeLessThan(subOrder);
  });

  it('is deterministic and sensitive to every input', () => {
    const base = organizerShareChallenge(eid, aid, ctIdx, pkOrgRTE, c1RTE, delta, a1, a2);
    expect(organizerShareChallenge(eid, aid, ctIdx, pkOrgRTE, c1RTE, delta, a1, a2)).toBe(base);

    expect(
      organizerShareChallenge('0x112233440000000000000008', aid, ctIdx, pkOrgRTE, c1RTE, delta, a1, a2),
    ).not.toBe(base);
    expect(
      organizerShareChallenge(eid, ('0x' + 'cd'.repeat(32)) as Hex, ctIdx, pkOrgRTE, c1RTE, delta, a1, a2),
    ).not.toBe(base);
    expect(
      organizerShareChallenge(eid, aid, ctIdx + 1, pkOrgRTE, c1RTE, delta, a1, a2),
    ).not.toBe(base);

    const points: BabyJubPoint[] = [pkOrgRTE, c1RTE, delta, a1, a2];
    for (let i = 0; i < points.length; i++) {
      for (const coord of [0, 1] as const) {
        const bumped = points.map((p, j) =>
          j === i ? ((coord === 0 ? [p[0] + 1n, p[1]] : [p[0], p[1] + 1n]) as BabyJubPoint) : p,
        );
        expect(
          organizerShareChallenge(eid, aid, ctIdx, bumped[0], bumped[1], bumped[2], bumped[3], bumped[4]),
        ).not.toBe(base);
      }
    }
  });
});

// ─── prove / verify ─────────────────────────────────────────────────────────

describe('organizer share prove/verify', () => {
  it('round-trips: proveOrganizerShare → verifyOrganizerShare', () => {
    const share = proveOrganizerShare(eid, aid, ctIdx, skOrg, c1TE, w);
    expect(verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, share)).toBe(true);
  });

  it('emits Δ = sk_org·C1, A1 = w·G, A2 = w·C1 in on-chain (RTE) form', () => {
    const share = proveOrganizerShare(eid, aid, ctIdx, skOrg, c1TE, w);
    expect(share.delta).toEqual(teToRte(mulPointEscalar(c1TE, skOrg)));
    expect(share.a1).toEqual(teToRte(mulPointEscalar(Base8, w)));
    expect(share.a2).toEqual(teToRte(mulPointEscalar(c1TE, w)));
  });

  it('produces z = w + e·sk_org mod q, in range', () => {
    const share = proveOrganizerShare(eid, aid, ctIdx, skOrg, c1TE, w);
    const e = organizerShareChallenge(
      eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, share.a1, share.a2,
    );
    expect(share.z).toBe((w + (e * skOrg) % subOrder) % subOrder);
    expect(share.z).toBeLessThan(subOrder);
  });

  it('draws a fresh witness when no nonce is pinned', () => {
    const a = proveOrganizerShare(eid, aid, ctIdx, skOrg, c1TE);
    const b = proveOrganizerShare(eid, aid, ctIdx, skOrg, c1TE);
    expect(a.a1).not.toEqual(b.a1);
    expect(a.z).not.toBe(b.z);
    // Δ is deterministic — only the proof randomises.
    expect(a.delta).toEqual(b.delta);
    expect(verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, a.delta, a)).toBe(true);
    expect(verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, b.delta, b)).toBe(true);
  });

  it('rejects a zero or out-of-range organizer secret', () => {
    expect(() => proveOrganizerShare(eid, aid, ctIdx, 0n, c1TE, w)).toThrow(/non-zero/);
    expect(() => proveOrganizerShare(eid, aid, ctIdx, subOrder, c1TE, w)).toThrow(/out of range/);
  });

  it('refuses an off-curve C1', () => {
    expect(() => proveOrganizerShare(eid, aid, ctIdx, skOrg, [1n, 1n], w)).toThrow(/not on BabyJubJub/);
  });
});

// ─── tampering ──────────────────────────────────────────────────────────────

describe('organizer share tamper resistance', () => {
  const share = proveOrganizerShare(eid, aid, ctIdx, skOrg, c1TE, w);

  it('rejects a tampered response z', () => {
    expect(
      verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, { ...share, z: share.z + 1n }),
    ).toBe(false);
  });

  it('rejects an out-of-range response z', () => {
    expect(
      verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, { ...share, z: subOrder }),
    ).toBe(false);
    expect(
      verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, { ...share, z: -1n }),
    ).toBe(false);
  });

  it('rejects the wrong epoch, aid or ciphertext index (replay protection)', () => {
    expect(
      verifyOrganizerShare('0x112233440000000000000008', aid, ctIdx, pkOrgRTE, c1RTE, share.delta, share),
    ).toBe(false);
    expect(
      verifyOrganizerShare(eid, ('0x' + 'cd'.repeat(32)) as Hex, ctIdx, pkOrgRTE, c1RTE, share.delta, share),
    ).toBe(false);
    expect(
      verifyOrganizerShare(eid, aid, ctIdx + 1, pkOrgRTE, c1RTE, share.delta, share),
    ).toBe(false);
  });

  it('rejects a share verified against another organizer key', () => {
    const otherPk = teToRte(mulPointEscalar(Base8, skOrg + 1n));
    expect(verifyOrganizerShare(eid, aid, ctIdx, otherPk, c1RTE, share.delta, share)).toBe(false);
  });

  it('rejects a share verified against another ciphertext', () => {
    const otherC1 = teToRte(mulPointEscalar(Base8, ephem + 1n));
    expect(verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, otherC1, share.delta, share)).toBe(false);
  });

  it('rejects a substituted Δ', () => {
    const wrongDelta = teToRte(addPoint(mulPointEscalar(c1TE, skOrg), Base8));
    expect(verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, wrongDelta, share)).toBe(false);
  });

  it('rejects an identity Δ (the sk_org = 0 forgery)', () => {
    const identity = teToRte([0n, 1n]);
    const forged = proveOrganizerShareForIdentity();
    expect(verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, identity, forged)).toBe(false);
  });

  it('rejects tampered or off-curve witness points', () => {
    expect(
      verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, {
        ...share,
        a1: [share.a1[0] + 1n, share.a1[1]],
      }),
    ).toBe(false);
    expect(
      verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, { ...share, a2: [1n, 1n] }),
    ).toBe(false);
  });

  it('rejects a share whose A1/A2 are swapped', () => {
    expect(
      verifyOrganizerShare(eid, aid, ctIdx, pkOrgRTE, c1RTE, share.delta, {
        ...share,
        a1: share.a2,
        a2: share.a1,
      }),
    ).toBe(false);
  });
});

/**
 * Build a syntactically valid proof for Δ = identity — the shape an attacker
 * would submit to make the combine subtract nothing. `verifyOrganizerShare`
 * must reject it on the identity check even though the second verifier
 * equation is satisfiable.
 */
function proveOrganizerShareForIdentity() {
  const a1 = teToRte(mulPointEscalar(Base8, w));
  const a2 = teToRte(mulPointEscalar(c1TE, w));
  return { a1, a2, z: w };
}
