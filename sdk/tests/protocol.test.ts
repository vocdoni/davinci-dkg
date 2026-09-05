// Cross-impl byte-equality assertions for the protocol constants. These
// vectors are the same as the ones in the Foundry suite
// (solidity/test/DKGProtocol.t.sol) and the same as the canonical
// `tests/vectors/protocol.json` generated from the Go side. Any change
// requires updating all three.

import { describe, it, expect } from 'vitest';
import { keccak256, toHex } from 'viem';
import {
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
  DomainContributionTranscriptV2,
  DomainFinalizeTranscriptV2,
  DomainDecryptCombineTranscriptV1,
  DomainOperatorRegisterV1Str,
  DomainOrganizerRegisterV1Str,
  DomainContributionTranscriptV2Str,
  DomainFinalizeTranscriptV2Str,
  DomainDecryptCombineTranscriptV1Str,
} from '../src/protocol';
import { SUBGROUP_ORDER } from '../src/schnorr';

describe('protocol constants', () => {
  it('DOMAIN_OPERATOR_REGISTER_V1 matches the Go vector', () => {
    expect(DomainOperatorRegisterV1).toBe(
      '0x4599aabb337c91d65fe440ef7e20c6dcc72c2459fd0901c45add50b08b3bb34d',
    );
  });

  it('DOMAIN_ORGANIZER_REGISTER_V1 matches the Go vector', () => {
    expect(DomainOrganizerRegisterV1).toBe(
      '0x41ea6f3fa95eccd1f3b1ce8e05efa11027280aa0c6b4167fd6695db659c30b28',
    );
  });

  it('the v4 BRLC domains are the v2 strings, not the v1 / poolkey ones', () => {
    expect(DomainContributionTranscriptV2Str).toBe('davinci-dkg:contribution:v2');
    expect(DomainFinalizeTranscriptV2Str).toBe('davinci-dkg:finalize:v2');
    expect(DomainDecryptCombineTranscriptV1Str).toBe('davinci-dkg:decrypt-combine:v1');
    expect(DomainContributionTranscriptV2).toBe(
      '0x4b37311b22cd0f09ae11d49f42ab65dce8fccf2600e6e2e7d41f51dc3d44b752',
    );
    expect(DomainFinalizeTranscriptV2).toBe(
      '0xe28959afa6ea38549c61aff75344fc2c9f148f1259fcef44fdd297a1d9a39d0f',
    );
    // A v3.1 contribution or activation transcript must not verify under v4.
    expect(DomainContributionTranscriptV2).not.toBe(keccak256(toHex('davinci-dkg:contribution:v1')));
    expect(DomainFinalizeTranscriptV2).not.toBe(keccak256(toHex('davinci-dkg:poolkey:v1')));
  });

  it('every digest is keccak256 of its documented preimage', () => {
    const pairs: Array<[string, `0x${string}`]> = [
      [DomainOperatorRegisterV1Str, DomainOperatorRegisterV1],
      [DomainOrganizerRegisterV1Str, DomainOrganizerRegisterV1],
      [DomainContributionTranscriptV2Str, DomainContributionTranscriptV2],
      [DomainFinalizeTranscriptV2Str, DomainFinalizeTranscriptV2],
      [DomainDecryptCombineTranscriptV1Str, DomainDecryptCombineTranscriptV1],
    ];
    for (const [preimage, digest] of pairs) {
      expect(keccak256(toHex(preimage))).toBe(digest);
    }
    // All are distinct — that separation is what stops a proof from one
    // transcript being replayed into another.
    expect(new Set(pairs.map(([, d]) => d)).size).toBe(pairs.length);
  });

  it('SUBGROUP_ORDER matches the gnark canonical BabyJubJub subgroup order', () => {
    expect(SUBGROUP_ORDER).toBe(
      2736030358979909402780800718157159386076813972158567259200215660948447373041n,
    );
  });
});
