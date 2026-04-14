import type { Hex } from 'viem';

/**
 * Build a bytes12 round ID from its two parts.
 *
 * Layout (big-endian on chain):
 *   [0..3]  uint32 ROUND_PREFIX (= chain ID)
 *   [4..11] uint64 nonce
 *
 * @param prefix  The ROUND_PREFIX constant from the DKGManager contract
 * @param nonce   The nonce returned by roundNonce() at round creation time
 */
export function buildRoundId(prefix: number | bigint, nonce: bigint): Hex {
  const p = BigInt(prefix);
  return `0x${p.toString(16).padStart(8, '0')}${nonce.toString(16).padStart(16, '0')}` as Hex;
}

/**
 * Parse a bytes12 round ID back into its components.
 */
export function parseRoundId(roundId: Hex): { prefix: number; nonce: bigint } {
  const hex = roundId.startsWith('0x') ? roundId.slice(2) : roundId;
  if (hex.length !== 24) throw new Error(`Invalid roundId length: ${roundId}`);
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
