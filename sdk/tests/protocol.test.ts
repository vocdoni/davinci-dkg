// Cross-impl byte-equality assertions for the protocol constants. These
// vectors are the same as the ones in the Foundry suite
// (solidity/test/DKGProtocol.t.sol) and the same as the canonical
// `tests/vectors/protocol.json` generated from the Go side. Any change
// requires updating all three.

import { describe, it, expect } from 'vitest';
import {
  AppMode,
  Role,
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
  DomainDLEQV1,
} from '../src/protocol';

describe('protocol constants', () => {
  it('AppMode matches the canonical numeric encoding', () => {
    expect(AppMode.PublicDerivation).toBe(0);
    expect(AppMode.OrganizerCoDec).toBe(1);
  });

  it('Role matches the canonical numeric encoding', () => {
    expect(Role.Committee).toBe(1);
    expect(Role.Organizer).toBe(2);
  });

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

  // (Removed DOMAIN_DERIVATION_V1 — DEEPSEEK §2.3. The actual S
  // derivation has no domain prefix.)
});
