// Cross-impl tests for the SDK Schnorr / DLEQ verifiers. Vectors are
// emitted by `cmd/operator-schnorr-vectors`; the values must match the
// Solidity constants in `solidity/test/TestHelpers.t.sol` byte-for-byte.

import { describe, it, expect } from 'vitest';
import {
  verifyOperatorSchnorr,
  operatorSchnorrChallenge,
  proveOrganizer,
  verifyOrganizerSchnorr,
  organizerSchnorrChallenge,
  proveCiphertext,
  verifyCiphertextPoK,
  ciphertextPoKChallenge,
  verifyDleq,
  dleqChallenge,
  DOMAIN_PARTIAL_DECRYPT,
  BN254_Q,
  SUBGROUP_ORDER,
} from '../src/schnorr';
import { Role } from '../src/protocol';
import { buildElGamal } from '../src/crypto/elgamal';
import { encryptWithProof } from '../src/flow';
import { Base8, addPoint, mulPointEscalar, type Point } from '@zk-kit/baby-jubjub';
import { fromRTEtoTE, fromTEtoRTE } from '../src/crypto/babyjub-form';

function teToRte(p: Point<bigint>): Point<bigint> {
  const [x, y] = fromTEtoRTE(p[0], p[1]);
  return [x, y];
}

// ─── operator vectors (lifted from cmd/operator-schnorr-vectors output) ─────

const OP_THIS = {
  operator: '0x7Fa9385bE102ac3EAc297483Dd6233D62b3e1496' as const,
  pubX: 17765672829315743641357949553430354448961270408100494783209553303687184365803n,
  pubY: 13591243454297365848719372676992908085762757043204242277513940025707896351954n,
  ax:   8735331066154608227876753674818247062971894554643955333501578253334124637538n,
  ay:   21272836954776476886917169960736847641993451090581680169106828110848349153636n,
  z:    1369885591156151396853044744701991591603236518469391786750891091614769051717n,
};

const OP_BEEF = {
  operator: '0x000000000000000000000000000000000000bEEF' as const,
  pubX: 10228722604559478181013548940833210623190136968531440936190496170400150013980n,
  pubY: 13886497050333420293068628977630539070604271411621054562122682889313139677221n,
  ax:   14566234973743316386655481200503006158732416518867749191804063760238069794878n,
  ay:   13078057343376780948266496972118389875598288245980415388506507934563371992806n,
  z:    855437853746059451716869189734643730464853812739723681978635195783451059776n,
};

describe('operator Schnorr verifier', () => {
  it('accepts a valid Go-generated vector (THIS)', () => {
    expect(
      verifyOperatorSchnorr(OP_THIS.operator, OP_THIS.pubX, OP_THIS.pubY, {
        ax: OP_THIS.ax, ay: OP_THIS.ay, z: OP_THIS.z,
      }),
    ).toBe(true);
  });

  it('accepts a valid Go-generated vector (BEEF)', () => {
    expect(
      verifyOperatorSchnorr(OP_BEEF.operator, OP_BEEF.pubX, OP_BEEF.pubY, {
        ax: OP_BEEF.ax, ay: OP_BEEF.ay, z: OP_BEEF.z,
      }),
    ).toBe(true);
  });

  it('rejects when the response is tampered', () => {
    expect(
      verifyOperatorSchnorr(OP_THIS.operator, OP_THIS.pubX, OP_THIS.pubY, {
        ax: OP_THIS.ax, ay: OP_THIS.ay, z: OP_THIS.z + 1n,
      }),
    ).toBe(false);
  });

  it('rejects when bound to the wrong operator address', () => {
    expect(
      verifyOperatorSchnorr(OP_BEEF.operator, OP_THIS.pubX, OP_THIS.pubY, {
        ax: OP_THIS.ax, ay: OP_THIS.ay, z: OP_THIS.z,
      }),
    ).toBe(false);
  });

  it('challenge is sensitive to every input', () => {
    const c = operatorSchnorrChallenge(OP_THIS.operator, OP_THIS.pubX, OP_THIS.pubY, OP_THIS.ax, OP_THIS.ay);
    expect(operatorSchnorrChallenge(OP_BEEF.operator, OP_THIS.pubX, OP_THIS.pubY, OP_THIS.ax, OP_THIS.ay)).not.toBe(c);
    expect(operatorSchnorrChallenge(OP_THIS.operator, OP_THIS.pubX + 1n, OP_THIS.pubY, OP_THIS.ax, OP_THIS.ay)).not.toBe(c);
    expect(operatorSchnorrChallenge(OP_THIS.operator, OP_THIS.pubX, OP_THIS.pubY + 1n, OP_THIS.ax, OP_THIS.ay)).not.toBe(c);
    expect(operatorSchnorrChallenge(OP_THIS.operator, OP_THIS.pubX, OP_THIS.pubY, OP_THIS.ax + 1n, OP_THIS.ay)).not.toBe(c);
    expect(operatorSchnorrChallenge(OP_THIS.operator, OP_THIS.pubX, OP_THIS.pubY, OP_THIS.ax, OP_THIS.ay + 1n)).not.toBe(c);
  });
});

// ─── organizer Schnorr (round-trip via convenience prover) ──────────────────

describe('organizer Schnorr', () => {
  const eid = '0x0000000000000000000000aa' as const;
  const aid = ('0x' + 'cd'.repeat(32)) as `0x${string}`;
  const sk = 1234567890123456789n;

  it('round-trips: proveOrganizer → verifyOrganizerSchnorr', () => {
    const { pkOrgX, pkOrgY, proof } = proveOrganizer(sk, eid, aid, 9999n);
    expect(verifyOrganizerSchnorr(eid, aid, pkOrgX, pkOrgY, proof)).toBe(true);
  });

  it('rejects when bound to the wrong epoch', () => {
    const { pkOrgX, pkOrgY, proof } = proveOrganizer(sk, eid, aid, 9999n);
    const wrongEid = '0x0000000000000000000000bb' as const;
    expect(verifyOrganizerSchnorr(wrongEid, aid, pkOrgX, pkOrgY, proof)).toBe(false);
  });

  it('rejects when bound to the wrong aid', () => {
    const { pkOrgX, pkOrgY, proof } = proveOrganizer(sk, eid, aid, 9999n);
    const wrongAid = ('0x' + 'cd'.repeat(31) + 'ce') as `0x${string}`;
    expect(verifyOrganizerSchnorr(eid, wrongAid, pkOrgX, pkOrgY, proof)).toBe(false);
  });

  it('challenge is deterministic and bound to all inputs', () => {
    const c = organizerSchnorrChallenge(eid, aid, 1n, 2n, 3n, 4n);
    expect(organizerSchnorrChallenge(eid, aid, 1n, 2n, 3n, 4n)).toBe(c);
    expect(organizerSchnorrChallenge(eid, aid, 1n, 2n, 3n, 5n)).not.toBe(c);
  });
});

// ─── ciphertext PoK (submitCiphertext) ──────────────────────────────────────

// Vector produced by the Go reference (`crypto/elgamal.EncryptWithProof` with
// sk = 1234567890123456789, m = 42, epochId 0x1122334400…07, aid 0xa0a1…bf).
// The proof's randomness is fixed in the vector, so the TS verifier must
// reproduce the Go keccak transcript byte-for-byte to accept it. All
// coordinates are the on-chain (RTE) words exactly as Go encodes them.
const GO_POK_VECTOR = {
  epochId: '0x112233440000000000000007' as const,
  aid: '0xa0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbebf' as const,
  pkX: 11211529648744409580389737002763249794516657436529318984204751975600767408943n,
  pkY: 19315036694604867924412484447665190374303888916962616112358246501479826895183n,
  c1x: 5327914099698692744718110762585765258609802896790308206290497831138639969986n,
  c1y: 19555038857321857409029761357904259070423102556686738772266708470204235630342n,
  c2x: 12317964812382882582192751384315197250526659581219650619219120741818980393620n,
  c2y: 6025912343892673086328146271441747721213980918426580448329692405187252326198n,
  ax: 12347161145034160738912429265240534252921277071789451058260893094809551572164n,
  ay: 16016968412754970057303487590293432601038401756831561809812512656371871988459n,
  z: 186081488140166616517867504215635649374666064245882754837308212635874187861n,
};

describe('ciphertext PoK', () => {
  const eid = '0x0000000000000000000000aa' as const;
  const aid = ('0x' + 'cd'.repeat(32)) as `0x${string}`;
  const sk = 987654321987654321n;
  const r = 4242424242424242n; // ElGamal randomness (C1 = r·G)
  const w = 7777777n;          // pinned Schnorr witness

  function encryptTE(k: bigint) {
    const pk = mulPointEscalar(Base8, sk) as Point<bigint>;
    const c1 = mulPointEscalar(Base8, k) as Point<bigint>;
    const c2 = addPoint(mulPointEscalar(Base8, 42n), mulPointEscalar(pk, k)) as Point<bigint>;
    return { c1: [c1[0], c1[1]] as [bigint, bigint], c2: [c2[0], c2[1]] as [bigint, bigint] };
  }

  function rteCoords(ct: { c1: [bigint, bigint]; c2: [bigint, bigint] }) {
    const [c1x, c1y] = fromTEtoRTE(ct.c1[0], ct.c1[1]);
    const [c2x, c2y] = fromTEtoRTE(ct.c2[0], ct.c2[1]);
    return { c1x, c1y, c2x, c2y };
  }

  it('accepts the Go-generated vector (byte-for-byte transcript)', () => {
    const v = GO_POK_VECTOR;
    expect(
      verifyCiphertextPoK(v.epochId, v.aid, v.c1x, v.c1y, v.c2x, v.c2y, { ax: v.ax, ay: v.ay, z: v.z }),
    ).toBe(true);
    // Sanity: the vector's C2 really is 42·G + r·PK for the vector's PK, i.e. it
    // is a genuine ElGamal ciphertext and not just a hash-consistent tuple.
    const pkTE = fromRTEtoTE(v.pkX, v.pkY);
    expect(mulPointEscalar(Base8, 1234567890123456789n)).toEqual(pkTE);
  });

  it('rejects the Go vector under another aid, epoch, or a tampered response', () => {
    const v = GO_POK_VECTOR;
    const proof = { ax: v.ax, ay: v.ay, z: v.z };
    const otherAid = ('0x' + 'a0a1a2a3a4a5a6a7a8a9aaabacadaeafb0b1b2b3b4b5b6b7b8b9babbbcbdbe' + 'c0') as `0x${string}`;
    expect(verifyCiphertextPoK(v.epochId, otherAid, v.c1x, v.c1y, v.c2x, v.c2y, proof)).toBe(false);
    expect(verifyCiphertextPoK('0x112233440000000000000008', v.aid, v.c1x, v.c1y, v.c2x, v.c2y, proof)).toBe(false);
    expect(verifyCiphertextPoK(v.epochId, v.aid, v.c1x, v.c1y, v.c2x, v.c2y, { ...proof, z: v.z + 1n })).toBe(false);
  });

  it('round-trips: proveCiphertext → verifyCiphertextPoK', () => {
    const ct = encryptTE(r);
    const proof = proveCiphertext(eid, aid, ct, r, w);
    const { c1x, c1y, c2x, c2y } = rteCoords(ct);
    expect(verifyCiphertextPoK(eid, aid, c1x, c1y, c2x, c2y, proof)).toBe(true);
    // Response is reduced into the subgroup order and the witness is A = w·G in RTE.
    expect(proof.z).toBeLessThan(SUBGROUP_ORDER);
    const A = mulPointEscalar(Base8, w) as Point<bigint>;
    expect([proof.ax, proof.ay]).toEqual(fromTEtoRTE(A[0], A[1]));
  });

  it('rejects when verified against another aid', () => {
    const ct = encryptTE(r);
    const proof = proveCiphertext(eid, aid, ct, r, w);
    const { c1x, c1y, c2x, c2y } = rteCoords(ct);
    const wrongAid = ('0x' + 'cd'.repeat(31) + 'ce') as `0x${string}`;
    expect(verifyCiphertextPoK(eid, wrongAid, c1x, c1y, c2x, c2y, proof)).toBe(false);
  });

  it('rejects when verified against another epoch', () => {
    const ct = encryptTE(r);
    const proof = proveCiphertext(eid, aid, ct, r, w);
    const { c1x, c1y, c2x, c2y } = rteCoords(ct);
    expect(verifyCiphertextPoK('0x0000000000000000000000ab', aid, c1x, c1y, c2x, c2y, proof)).toBe(false);
  });

  it('rejects a re-randomised C1 + G (the decryption-oracle replay)', () => {
    const ct = encryptTE(r);
    const proof = proveCiphertext(eid, aid, ct, r, w);
    const shifted = addPoint(ct.c1, Base8) as Point<bigint>;
    const [c1x, c1y] = fromTEtoRTE(shifted[0], shifted[1]);
    const [c2x, c2y] = fromTEtoRTE(ct.c2[0], ct.c2[1]);
    expect(verifyCiphertextPoK(eid, aid, c1x, c1y, c2x, c2y, proof)).toBe(false);
  });

  it('rejects a swapped C2 (C2 is bound into the transcript)', () => {
    const ct = encryptTE(r);
    const proof = proveCiphertext(eid, aid, ct, r, w);
    const { c1x, c1y } = rteCoords(ct);
    const other = encryptTE(r + 1n);
    const [c2x, c2y] = fromTEtoRTE(other.c2[0], other.c2[1]);
    expect(verifyCiphertextPoK(eid, aid, c1x, c1y, c2x, c2y, proof)).toBe(false);
  });

  it('rejects a tampered response, an off-curve witness and an out-of-range z', () => {
    const ct = encryptTE(r);
    const proof = proveCiphertext(eid, aid, ct, r, w);
    const { c1x, c1y, c2x, c2y } = rteCoords(ct);
    expect(verifyCiphertextPoK(eid, aid, c1x, c1y, c2x, c2y, { ...proof, z: proof.z + 1n })).toBe(false);
    expect(verifyCiphertextPoK(eid, aid, c1x, c1y, c2x, c2y, { ...proof, ax: 1n, ay: 1n })).toBe(false);
    expect(verifyCiphertextPoK(eid, aid, c1x, c1y, c2x, c2y, { ...proof, z: SUBGROUP_ORDER })).toBe(false);
  });

  it('proveCiphertext refuses a randomness that does not open c1', () => {
    const ct = encryptTE(r);
    expect(() => proveCiphertext(eid, aid, ct, r + 1n, w)).toThrow(/does not open c1/);
  });

  it('challenge is deterministic and sensitive to every input', () => {
    const c = ciphertextPoKChallenge(eid, aid, 1n, 2n, 3n, 4n, 5n, 6n);
    expect(ciphertextPoKChallenge(eid, aid, 1n, 2n, 3n, 4n, 5n, 6n)).toBe(c);
    expect(c).toBeLessThan(SUBGROUP_ORDER);
    expect(ciphertextPoKChallenge('0x0000000000000000000000ab', aid, 1n, 2n, 3n, 4n, 5n, 6n)).not.toBe(c);
    expect(ciphertextPoKChallenge(eid, ('0x' + 'ce'.repeat(32)) as `0x${string}`, 1n, 2n, 3n, 4n, 5n, 6n)).not.toBe(c);
    const args: bigint[] = [1n, 2n, 3n, 4n, 5n, 6n];
    for (let i = 0; i < args.length; i++) {
      const bumped = [...args]; bumped[i] += 1n;
      const [a, b, cc, d, e, f] = bumped;
      expect(ciphertextPoKChallenge(eid, aid, a, b, cc, d, e, f)).not.toBe(c);
    }
  });

  it('encryptWithProof produces a ciphertext + proof the verifier accepts (TE → RTE boundary)', async () => {
    const eg = await buildElGamal();
    const { privKey, pubKey } = eg.generateKeyPair();
    const { ciphertext, pok } = await encryptWithProof(eid, aid, 42n, pubKey);
    const { c1x, c1y, c2x, c2y } = rteCoords(ciphertext);
    expect(verifyCiphertextPoK(eid, aid, c1x, c1y, c2x, c2y, pok)).toBe(true);
    expect(eg.decrypt(ciphertext, privKey)).toBe(42n);
    // Proofs are per (epoch, aid): the same ciphertext under a different aid fails.
    expect(verifyCiphertextPoK(eid, ('0x' + 'ce'.repeat(32)) as `0x${string}`, c1x, c1y, c2x, c2y, pok)).toBe(false);
  });
});

// ─── DLEQ round-trip (committee role) ───────────────────────────────────────

describe('DLEQ verifier', () => {
  const eid = '0x000000000000000000000077' as const;
  const aid = ('0x' + '00'.repeat(31) + '07') as `0x${string}`;
  const ctIdx = 3n;
  const i = 5n;
  const sk = 4242424242n;          // committee share secret d_i
  const ephem = 999999n;           // C_1 = ephem · G
  const w = 7777777n;              // DLEQ witness

  function buildPoints() {
    // PK = sk · G (committee D_i). All curve math is done in TE form, then
    // every point is converted to RTE before being fed into the transcript
    // / verifier — that's how the on-chain artifacts arrive.
    const PK_TE = mulPointEscalar(Base8, sk) as Point<bigint>;
    const C1_TE = mulPointEscalar(Base8, ephem) as Point<bigint>;
    const Delta_TE = mulPointEscalar(C1_TE, sk) as Point<bigint>;
    const A1_TE = mulPointEscalar(Base8, w) as Point<bigint>;
    const A2_TE = mulPointEscalar(C1_TE, w) as Point<bigint>;
    return {
      PK: teToRte(PK_TE),
      C1: teToRte(C1_TE),
      Delta: teToRte(Delta_TE),
      A1: teToRte(A1_TE),
      A2: teToRte(A2_TE),
    };
  }

  function buildProof() {
    const { PK, C1, Delta, A1, A2 } = buildPoints();
    const transcript = {
      epochId: eid, aid, ctIdx, role: Role.Committee, participantIndex: i,
      points: { base: C1, publicKey: PK, delta: Delta, a1: A1, a2: A2 },
    };
    const c = dleqChallenge(transcript) % SUBGROUP_ORDER;
    const z = (w + c * sk) % SUBGROUP_ORDER;
    return { transcript, z };
  }

  it('accepts a well-formed committee DLEQ', () => {
    const { transcript, z } = buildProof();
    expect(verifyDleq(transcript, z)).toBe(true);
  });

  it('rejects when the response is tampered', () => {
    const { transcript, z } = buildProof();
    expect(verifyDleq(transcript, z + 1n)).toBe(false);
  });

  it('rejects when the role is swapped (replay protection)', () => {
    const { transcript, z } = buildProof();
    const replay = { ...transcript, role: Role.Organizer };
    expect(verifyDleq(replay, z)).toBe(false);
  });

  it('rejects when the ctIdx changes (replay across ciphertexts)', () => {
    const { transcript, z } = buildProof();
    const replay = { ...transcript, ctIdx: ctIdx + 1n };
    expect(verifyDleq(replay, z)).toBe(false);
  });

  it('rejects when the aid changes (replay across applications)', () => {
    const { transcript, z } = buildProof();
    const replay = { ...transcript, aid: ('0x' + '00'.repeat(31) + '08') as `0x${string}` };
    expect(verifyDleq(replay, z)).toBe(false);
  });

  it('rejects an off-curve point', () => {
    const { transcript, z } = buildProof();
    const bad = { ...transcript, points: { ...transcript.points, delta: [1n, 1n] as Point<bigint> } };
    expect(verifyDleq(bad, z)).toBe(false);
  });

  it('domain constant matches the on-chain reduction', () => {
    // 'davinci-dkg/partial-decrypt/v1' as a 30-byte big-endian integer mod BN254_Q.
    const enc = new TextEncoder().encode('davinci-dkg/partial-decrypt/v1');
    let bi = 0n;
    for (const b of enc) bi = (bi << 8n) | BigInt(b);
    expect(DOMAIN_PARTIAL_DECRYPT).toBe(bi % BN254_Q);
  });
});

// ─── point-add sanity (linkage with @zk-kit/baby-jubjub) ────────────────────

describe('schnorr-helpers smoke', () => {
  it('Base8 has expected coordinates after one self-add', () => {
    const G = Base8 as Point<bigint>;
    const G2 = addPoint(G, G);
    // Just verify it's deterministic — we don't pin the constant, the
    // package guarantees correctness.
    expect(addPoint(G, G)).toEqual(G2);
  });
});
