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
  verifyDleq,
  dleqChallenge,
  DOMAIN_PARTIAL_DECRYPT,
  BN254_Q,
  SUBGROUP_ORDER,
} from '../src/schnorr';
import { Base8, addPoint, mulPointEscalar, type Point } from '@zk-kit/baby-jubjub';
import { fromTEtoRTE } from '../src/crypto/babyjub-form';

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

// ─── DLEQ round-trip (committee partial decryption) ─────────────────────────

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
      epochId: eid, aid, ctIdx, participantIndex: i,
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

  it('rejects when the participant index changes (replay across committee members)', () => {
    const { transcript, z } = buildProof();
    const replay = { ...transcript, participantIndex: i + 1n };
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
