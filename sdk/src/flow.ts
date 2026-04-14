/**
 * High-level DKG flow helpers.
 *
 * These functions compose the lower-level DKGClient / DKGWriter / monitor
 * primitives into the common end-to-end flows.
 *
 * Full flow:
 *   1. organizer calls createRound()
 *   2. DKG nodes call claimSlot() once the seed block is mined
 *   3. Nodes submit contributions → round moves to Contribution phase
 *   4. Nodes finalize round → round moves to Finalized; collective public key available
 *   5. Anyone can encrypt data using the collective public key (ElGamal)
 *   6. DKG nodes submit partial decryptions for a given ciphertext
 *   7. Any node calls combineDecryption → DecryptionCombined event emitted
 *   8. Caller can verify the plaintext matches the original message
 */

import { type PublicClient, type Hash } from 'viem';
import { DKGClient } from './client.js';
import { type RoundPolicy, type ElGamalCiphertext, type BabyJubPoint } from './types.js';
import { waitForRoundStatus, waitForDecryption } from './monitor.js';
import { buildElGamal } from './crypto/elgamal.js';
import { RoundStatus } from './types.js';

export interface CollectivePublicKey {
  /**
   * The BabyJubJub point that is the collective public key.
   * Encoded as the AggregateCommitments[0] point from the finalize proof.
   * x and y are the field-element coordinates.
   *
   * NOTE: The on-chain `collectivePublicKeyHash` is keccak256(x, y).
   * To obtain the actual (x, y) coordinates you must decode the calldata of
   * the `finalizeRound` transaction that triggered the RoundFinalized event.
   * The `transcript` argument of that call contains the proof's public inputs.
   */
  x: bigint;
  y: bigint;
}

/**
 * Wait until a round is Finalized, then return the collective public key hash
 * (the value stored on-chain in the RoundFinalized event).
 *
 * To get the actual curve point (x, y) you need to decode the `finalizeRound`
 * calldata. This helper returns the on-chain hash so callers can verify a
 * candidate public key against it.
 */
export async function waitForCollectivePublicKeyHash(
  client: DKGClient,
  roundId: `0x${string}`,
  options?: { intervalMs?: number; timeoutMs?: number },
): Promise<`0x${string}`> {
  await waitForRoundStatus(client, roundId, RoundStatus.Finalized, options);
  const events = await client.getRoundFinalizedEvents(roundId);
  if (events.length === 0) {
    throw new Error(`No RoundFinalized event found for round ${roundId}`);
  }
  return events[events.length - 1].collectivePublicKeyHash;
}

/**
 * Encrypt a message using the DKG collective public key.
 *
 * @param message    Small integer plaintext (must fit in BabyJubJub scalar)
 * @param pubKey     Collective public key as a BabyJubJub point [x, y]
 * @param k          Optional ephemeral key; a random scalar is used when omitted
 * @returns          ElGamal ciphertext {c1, c2}
 */
export async function encrypt(
  message: bigint,
  pubKey: BabyJubPoint,
  k?: bigint,
): Promise<ElGamalCiphertext> {
  const elgamal = await buildElGamal();
  return elgamal.encrypt(message, pubKey, k);
}

/**
 * Decrypt an ElGamal ciphertext given the private key.
 *
 * Only usable for small plaintexts (< 2^20) because recovery is brute-force DLOG.
 * In the real DKG protocol the decryption is performed threshold-style on-chain
 * by the committee; this helper is for testing and verification only.
 */
export async function decrypt(
  ciphertext: ElGamalCiphertext,
  privKey: bigint,
): Promise<bigint> {
  const elgamal = await buildElGamal();
  return elgamal.decrypt(ciphertext, privKey);
}

/**
 * Wait for the on-chain combined decryption of a ciphertext to complete.
 *
 * @returns The completed CombinedDecryptionRecord (check record.completed === true).
 */
export async function waitForCombinedDecryption(
  client: DKGClient,
  roundId: `0x${string}`,
  ciphertextIndex: number,
  options?: { intervalMs?: number; timeoutMs?: number },
) {
  return waitForDecryption(client, roundId, ciphertextIndex, options);
}

/**
 * Full demonstration flow (for testing / documentation purposes).
 *
 * This function illustrates the complete end-to-end DKG encryption flow:
 *
 *   1. Wait for round to be Finalized (assuming it was already created)
 *   2. Encrypt a plaintext with the collective public key
 *   3. Wait for the on-chain combined decryption to complete
 *
 * In production, steps 2–3 are done by different parties:
 *  - The data producer encrypts and publishes the ciphertext (off-chain / on-chain).
 *  - DKG nodes submit partial decryptions for the ciphertext.
 *  - Any party that has the threshold of partial decryptions calls combineDecryption.
 *
 * @param client         Read-only DKGClient
 * @param roundId        The round ID
 * @param collectivePub  The collective public key point [x, y] derived from the
 *                       finalizeRound calldata
 * @param plaintext      Small integer to encrypt/decrypt
 * @param ciphertextIndex  Index to identify which ciphertext to wait for (1-based)
 */
export async function demonstrateEncryptDecryptFlow(
  client: DKGClient,
  roundId: `0x${string}`,
  collectivePub: BabyJubPoint,
  plaintext: bigint,
  ciphertextIndex: number,
): Promise<{
  ciphertext: ElGamalCiphertext;
  decryptionCompleted: boolean;
}> {
  // 1. Encrypt
  const ciphertext = await encrypt(plaintext, collectivePub);

  // 2. Wait for the DKG nodes to decrypt on-chain
  const record = await waitForCombinedDecryption(client, roundId, ciphertextIndex, {
    timeoutMs: 300_000,
  });

  return {
    ciphertext,
    decryptionCompleted: record.completed,
  };
}
