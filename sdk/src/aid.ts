// Application ids. An id is `salt << 160 | registrant`:
// `DKGAppManager.registerApplication` reverts `InvalidApplication` unless the
// low 160 bits of `aid` are the caller, so nobody can register an id in
// another account's namespace (vocdoni/davinci-dkg#14). Mirrors Go
// `types.ApplicationID` and the contract's `_requireValidAid`.

import { getAddress, type Address } from 'viem';

/**
 * Width of the salt. A salt below `2^92` keeps every id below `2^252`, inside
 * the BN254 scalar field the decryption proofs bind `aid` in.
 */
export const AID_SALT_BITS = 92;

const SALT_BOUND = 1n << BigInt(AID_SALT_BITS);
const ADDRESS_MASK = (1n << 160n) - 1n;

/**
 * The application id `salt << 160 | registrant`, the only shape the contract
 * accepts from `registrant`. Integrators that derive ids from public data
 * (a process id, a hash) put that value in the salt.
 */
export function applicationId(registrant: Address, salt: bigint): `0x${string}` {
  if (salt < 0n || salt >= SALT_BOUND) {
    throw new Error(`applicationId: salt must be in [0, 2^${AID_SALT_BITS})`);
  }
  const value = (salt << 160n) | BigInt(getAddress(registrant));
  return `0x${value.toString(16).padStart(64, '0')}`;
}

/** A fresh application id in `registrant`'s namespace, with a random 92-bit salt. */
export function randomAid(registrant: Address): `0x${string}` {
  const bytes = globalThis.crypto.getRandomValues(new Uint8Array(12));
  let salt = 0n;
  for (const b of bytes) salt = (salt << 8n) | BigInt(b);
  return applicationId(registrant, salt >> 4n); // 96 random bits → 92
}

/** The only account that may register `aid`: its low 160 bits. */
export function aidRegistrant(aid: `0x${string}`): Address {
  return getAddress(`0x${(BigInt(aid) & ADDRESS_MASK).toString(16).padStart(40, '0')}`);
}
