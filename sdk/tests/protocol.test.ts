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
  DomainDLEQV1,
  DomainOrganizerShareV1,
  DomainOperatorRegisterV1Str,
  DomainOrganizerRegisterV1Str,
  DomainDLEQV1Str,
  DomainOrganizerShareV1Str,
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

  it('DOMAIN_DLEQ_V1 matches the Go vector', () => {
    expect(DomainDLEQV1).toBe(
      '0x48fabea26e7a072780483852e403ea60b2f51a07c735c3e4b852ac6bb99b5a91',
    );
  });

  it('DOMAIN_ORGANIZER_SHARE_V1 matches the Go vector', () => {
    expect(DomainOrganizerShareV1).toBe(
      '0x1608b6df1dd60f54655f6e7cf082d648cc3ca53756f1527d1f112085c2ddad2d',
    );
  });

  it('every digest is keccak256 of its documented preimage', () => {
    const pairs: Array<[string, `0x${string}`]> = [
      [DomainOperatorRegisterV1Str, DomainOperatorRegisterV1],
      [DomainOrganizerRegisterV1Str, DomainOrganizerRegisterV1],
      [DomainDLEQV1Str, DomainDLEQV1],
      [DomainOrganizerShareV1Str, DomainOrganizerShareV1],
    ];
    for (const [preimage, digest] of pairs) {
      expect(keccak256(toHex(preimage))).toBe(digest);
    }
    // All four are distinct — that separation is what stops a proof from one
    // transcript being replayed into another.
    expect(new Set(pairs.map(([, d]) => d)).size).toBe(pairs.length);
  });

  it('SUBGROUP_ORDER matches the gnark canonical BabyJubJub subgroup order', () => {
    expect(SUBGROUP_ORDER).toBe(
      2736030358979909402780800718157159386076813972158567259200215660948447373041n,
    );
  });
});
