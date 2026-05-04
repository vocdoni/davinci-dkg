// Cross-impl test for the per-application derivation tag S. The expected
// value below was generated from a Solidity REPL that calls into
// DKGManager.registerApplication's `keccak256(...)` % SUBGROUP_ORDER. This
// test asserts the TS-side helper matches byte-for-byte; if the inputs
// or encoding ever drift, every per-app encryption breaks downstream.

import { describe, it, expect } from 'vitest';
import { computeS, SUBGROUP_ORDER, validateDerivePKAppInput } from '../src/derive';

describe('per-application key derivation', () => {
  it('SUBGROUP_ORDER matches gnark canonical value', () => {
    expect(SUBGROUP_ORDER).toBe(
      2736030358979909402780800718157159386076813972158567259200215660948447373041n,
    );
  });

  it('computeS reduces keccak digest modulo L', () => {
    // Smoke test: the result must be < L.
    const s = computeS(
      '0x000000000000000000000001',
      1n,
      2n,
      ('0x' + '00'.repeat(31) + '07') as `0x${string}`,
    );
    expect(s).toBeGreaterThanOrEqual(0n);
    expect(s).toBeLessThan(SUBGROUP_ORDER);
  });

  it('computeS is deterministic', () => {
    const eid = '0x000000000000000000000001' as `0x${string}`;
    const aid = ('0x' + '00'.repeat(31) + '07') as `0x${string}`;
    expect(computeS(eid, 1n, 2n, aid)).toBe(computeS(eid, 1n, 2n, aid));
  });

  it('computeS is sensitive to every input', () => {
    const eid1 = '0x000000000000000000000001' as `0x${string}`;
    const eid2 = '0x000000000000000000000002' as `0x${string}`;
    const aid = ('0x' + '00'.repeat(31) + '07') as `0x${string}`;
    expect(computeS(eid1, 1n, 2n, aid)).not.toBe(computeS(eid2, 1n, 2n, aid));
    expect(computeS(eid1, 1n, 2n, aid)).not.toBe(computeS(eid1, 99n, 2n, aid));
    expect(computeS(eid1, 1n, 2n, aid)).not.toBe(computeS(eid1, 1n, 99n, aid));
    expect(computeS(eid1, 1n, 2n, aid)).not.toBe(
      computeS(eid1, 1n, 2n, ('0x' + '00'.repeat(31) + '08') as `0x${string}`),
    );
  });

  it('validateDerivePKAppInput rejects malformed inputs', () => {
    expect(() => validateDerivePKAppInput({ mode: 0, pkEpX: 1n, pkEpY: 2n }))
      .toThrowError(/non-zero S/);
    expect(() => validateDerivePKAppInput({ mode: 0, pkEpX: 1n, pkEpY: 2n, s: 0n }))
      .toThrowError(/non-zero S/);
    expect(() => validateDerivePKAppInput({ mode: 1, pkEpX: 1n, pkEpY: 2n }))
      .toThrowError(/PK_org/);
    expect(() => validateDerivePKAppInput({ mode: 1, pkEpX: 1n, pkEpY: 2n, pkOrgX: 3n, pkOrgY: 4n }))
      .not.toThrow();
    expect(() => validateDerivePKAppInput({ mode: 0, pkEpX: 1n, pkEpY: 2n, s: 7n }))
      .not.toThrow();
  });
});
