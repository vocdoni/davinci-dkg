import type { Hex } from 'viem';

/**
 * Build a bytes12 epoch ID from its two parts.
 *
 * Layout (big-endian on chain):
 *   [0..3]  uint32 EPOCH_PREFIX = uint32(keccak256(chainId, manager))
 *           (see solidity/src/libraries/DKGIdLib.sol::getPrefix)
 *   [4..11] uint64 nonce
 *
 * @param prefix  The EPOCH_PREFIX value read from DKGManager (NOT the chain ID)
 * @param nonce   The nonce returned by epochNonce() at epoch creation time
 */
export function buildEpochId(prefix: number | bigint, nonce: bigint): Hex {
  const p = BigInt(prefix);
  return `0x${p.toString(16).padStart(8, '0')}${nonce.toString(16).padStart(16, '0')}` as Hex;
}

/**
 * Parse a bytes12 epoch ID back into its components.
 */
export function parseEpochId(epochId: Hex): { prefix: number; nonce: bigint } {
  const hex = epochId.startsWith('0x') ? epochId.slice(2) : epochId;
  if (hex.length !== 24) throw new Error(`Invalid epochId length: ${epochId}`);
  const prefix = parseInt(hex.slice(0, 8), 16);
  const nonce = BigInt('0x' + hex.slice(8, 24));
  return { prefix, nonce };
}

/**
 * Pad a bigint to a 32-byte hex string (no 0x prefix).
 */
export function padBigInt(n: bigint, bytes = 32): string {
  return n.toString(16).padStart(bytes * 2, '0');
}

/**
 * Sleep for `ms` milliseconds.
 */
export function sleep(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}
