// ElGamal / BabyJubJub crypto tests.
// These run without a chain connection — pure in-memory.

import { describe, it, expect, beforeAll } from 'vitest';
import { buildElGamal, type ElGamal } from '../src/index.js';
import { subOrder } from '@zk-kit/baby-jubjub';

describe('ElGamal on BabyJubJub', () => {
  let eg: ElGamal;

  beforeAll(async () => {
    eg = await buildElGamal();
  });

  it('buildElGamal initialises and returns the same singleton on repeated calls', async () => {
    const eg2 = await buildElGamal();
    expect(eg2).toBe(eg);
  });

  it('randomScalar returns a bigint within the subgroup order', () => {
    const s = eg.randomScalar();
    expect(typeof s).toBe('bigint');
    expect(s).toBeGreaterThan(0n);
    expect(s).toBeLessThan(subOrder);
  });

  it('generateKeyPair returns a private key and a curve point', () => {
    const { privKey, pubKey } = eg.generateKeyPair();
    expect(typeof privKey).toBe('bigint');
    expect(privKey).toBeGreaterThan(0n);
    expect(Array.isArray(pubKey)).toBe(true);
    expect(pubKey).toHaveLength(2);
    expect(typeof pubKey[0]).toBe('bigint');
    expect(typeof pubKey[1]).toBe('bigint');
  });

  it('encrypt returns c1 and c2 curve points', () => {
    const { pubKey } = eg.generateKeyPair();
    const ct = eg.encrypt(42n, pubKey);
    expect(ct.c1).toHaveLength(2);
    expect(ct.c2).toHaveLength(2);
    expect(typeof ct.c1[0]).toBe('bigint');
    expect(typeof ct.c2[0]).toBe('bigint');
  });

  it('decrypt(encrypt(m)) === m for a small message', () => {
    const { privKey, pubKey } = eg.generateKeyPair();
    for (const msg of [0n, 1n, 42n, 999n]) {
      const ct = eg.encrypt(msg, pubKey);
      const recovered = eg.decrypt(ct, privKey);
      expect(recovered).toBe(msg);
    }
  });

  // The 2^32 cap is enforced by BSGS (see crypto/elgamal.ts). Values that
  // would have busted the previous 2^20 brute-force ceiling must round-trip;
  // values at or above the new cap must throw.
  it('decrypt handles values well above the old 2^20 brute-force ceiling', () => {
    const { privKey, pubKey } = eg.generateKeyPair();
    for (const msg of [1_048_576n, 4_344_444n, 1_000_000_000n, (1n << 32n) - 1n]) {
      const ct = eg.encrypt(msg, pubKey);
      const recovered = eg.decrypt(ct, privKey);
      expect(recovered).toBe(msg);
    }
  });

  it('decrypt throws when plaintext is at or above 2^32', () => {
    const { privKey, pubKey } = eg.generateKeyPair();
    const ct = eg.encrypt(1n << 32n, pubKey);
    expect(() => eg.decrypt(ct, privKey)).toThrow(/out of range/);
  });

  it('encrypt is non-deterministic when k is omitted', () => {
    const { pubKey } = eg.generateKeyPair();
    const ct1 = eg.encrypt(10n, pubKey);
    const ct2 = eg.encrypt(10n, pubKey);
    // With overwhelming probability the ephemeral keys differ
    const same = ct1.c1[0] === ct2.c1[0] && ct1.c1[1] === ct2.c1[1];
    expect(same).toBe(false);
  });

  it('encrypt with explicit k is deterministic', () => {
    const { pubKey } = eg.generateKeyPair();
    const k = 12345n;
    const ct1 = eg.encrypt(7n, pubKey, k);
    const ct2 = eg.encrypt(7n, pubKey, k);
    expect(ct1.c1[0]).toBe(ct2.c1[0]);
    expect(ct1.c1[1]).toBe(ct2.c1[1]);
    expect(ct1.c2[0]).toBe(ct2.c2[0]);
    expect(ct1.c2[1]).toBe(ct2.c2[1]);
  });

  it('different keys produce different ciphertexts', () => {
    const kp1 = eg.generateKeyPair();
    const kp2 = eg.generateKeyPair();
    const k = eg.randomScalar();
    const ct1 = eg.encrypt(5n, kp1.pubKey, k);
    const ct2 = eg.encrypt(5n, kp2.pubKey, k);
    // c1 is the same (same k, same base point) but c2 differs (different pubkey)
    expect(ct1.c1[0]).toBe(ct2.c1[0]);
    expect(ct1.c2[0]).not.toBe(ct2.c2[0]);
  });

  it('packPoint / unpackPoint roundtrip', () => {
    const { pubKey } = eg.generateKeyPair();
    const packed   = eg.packPoint(pubKey);
    const unpacked = eg.unpackPoint(packed);
    expect(unpacked[0]).toBe(pubKey[0]);
    expect(unpacked[1]).toBe(pubKey[1]);
  });

  it('mulPoint and addPoint produce curve points', () => {
    const { pubKey } = eg.generateKeyPair();
    const doubled = eg.addPoint(pubKey, pubKey);
    const via2    = eg.mulPoint(pubKey, 2n);
    expect(doubled[0]).toBe(via2[0]);
    expect(doubled[1]).toBe(via2[1]);
  });
});
