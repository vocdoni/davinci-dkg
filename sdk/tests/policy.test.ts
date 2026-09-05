// AppPolicy defaults — what `registerApplication` sends when a field is left out.

import { describe, it, expect } from 'vitest';
import { AppMode, appModeLabel, normalizeAppPolicy } from '../src/index.js';

describe('normalizeAppPolicy', () => {
  it('defaults to organizer-locked, registrant-only, uncapped, no deadline', () => {
    expect(normalizeAppPolicy()).toEqual({
      mode: AppMode.OrganizerLocked,
      openSubmission: false,
      submitters: [],
      maxCiphertexts: 0,
      notBeforeBlock: 0n,
      notAfterBlock: 0n,
      decryptNotBefore: 0n,
      decryptNotAfter: 0n,
    });
  });

  it('keeps every field that is given', () => {
    const submitters = ['0x1111111111111111111111111111111111111111'] as const;
    const policy = normalizeAppPolicy({
      mode: AppMode.Automatic,
      submitters: [...submitters],
      maxCiphertexts: 4,
      notAfterBlock: 99n,
      decryptNotAfter: 1_800_000_000n,
    });
    expect(policy.mode).toBe(AppMode.Automatic);
    expect(policy.openSubmission).toBe(false);
    expect(policy.submitters).toEqual([...submitters]);
    expect(policy.maxCiphertexts).toBe(4);
    expect(policy.notBeforeBlock).toBe(0n);
    expect(policy.notAfterBlock).toBe(99n);
    expect(policy.decryptNotAfter).toBe(1_800_000_000n);
  });

  it('labels both modes', () => {
    expect(appModeLabel(AppMode.OrganizerLocked)).toBe('organizer-locked');
    expect(appModeLabel(AppMode.Automatic)).toBe('automatic');
    expect(appModeLabel(7)).toBe('Unknown(7)');
  });
});
