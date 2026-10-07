// Application-id namespace (vocdoni/davinci-dkg#14): `aid = salt << 160 |
// registrant`. The layout literal below is the one Go's
// `types.TestApplicationIDLayout` asserts.

import { describe, it, expect } from 'vitest';
import { createPublicClient, createWalletClient, getAddress, http } from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import {
  AID_SALT_BITS,
  AppMode,
  DKGWriter,
  aidRegistrant,
  applicationId,
  randomAid,
} from '../src/index.js';

const BN254_R = 21888242871839275222246405745257275088548364400416034343698204186575808495617n;

describe('application ids', () => {
  it('lays out salt << 160 | registrant like the Go side', () => {
    const registrant = '0x42fC000000000000000000000000000000Ab589F';
    const aid = applicationId(registrant, 0x1234n);
    expect(aid).toBe('0x000000000000000000001234' + '42fc000000000000000000000000000000ab589f');
    expect(aidRegistrant(aid)).toBe(getAddress(registrant));
  });

  it('stays inside the BN254 scalar field for every salt', () => {
    const max = (1n << BigInt(AID_SALT_BITS)) - 1n;
    const aid = applicationId('0xffffffffffffffffffffffffffffffffffffffff', max);
    expect(BigInt(aid) < BN254_R).toBe(true);
    expect(() => applicationId('0xffffffffffffffffffffffffffffffffffffffff', max + 1n)).toThrow();
    expect(() => applicationId('0xffffffffffffffffffffffffffffffffffffffff', -1n)).toThrow();
  });

  it('randomAid stays in the registrant namespace', () => {
    const registrant = '0x00000000000000000000000000000000000A11CE';
    const a = randomAid(registrant);
    const b = randomAid(registrant);
    expect(a).not.toBe(b);
    for (const aid of [a, b]) {
      expect(aid).toMatch(/^0x[0-9a-f]{64}$/);
      expect(aidRegistrant(aid)).toBe(getAddress(registrant));
      expect(BigInt(aid) >> BigInt(160 + AID_SALT_BITS)).toBe(0n);
    }
  });

  it('DKGWriter refuses to register an id in another account namespace', async () => {
    const account = privateKeyToAccount(
      '0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d',
    );
    // Unroutable endpoint: the check must fire before any RPC call.
    const transport = http('http://127.0.0.1:1');
    const writer = new DKGWriter({
      publicClient: createPublicClient({ transport }),
      walletClient: createWalletClient({ account, transport }),
      managerAddress: '0x0000000000000000000000000000000000000001',
    });
    const foreign = randomAid('0x000000000000000000000000000000000000BEEF');
    await expect(
      writer.registerApplication('0x000000000000000000000001', foreign, { mode: AppMode.Automatic }),
    ).rejects.toThrow(/belongs to 0x000000000000000000000000000000000000bEEF/i);
  });
});
