/**
 * High-level DKG flow helpers.
 *
 * These functions compose the lower-level DKGClient / DKGWriter / monitor
 * primitives into the common end-to-end flows.
 *
 * Full flow:
 *   1. organizer calls createEpoch()
 *   2. DKG nodes call claimSlot() once the seed block is mined
 *   3. Nodes submit contributions → epoch moves to Contribution phase
 *   4. Nodes finalize epoch → epoch moves to Finalized; collective public key available
 *   5. Anyone can encrypt data using the collective public key (ElGamal)
 *   6. DKG nodes submit partial decryptions for a given ciphertext
 *   7. Any node calls combineDecryption → DecryptionCombined event emitted
 *   8. Caller can verify the plaintext matches the original message
 */

import { type PublicClient, type Hash } from 'viem';
import { DKGClient } from './client.js';
import { type EpochPolicy, type ElGamalCiphertext, type BabyJubPoint } from './types.js';
import { waitForEpochPhase, waitForDecryption } from './monitor.js';
import { buildElGamal } from './crypto/elgamal.js';
import { EpochPhase } from './types.js';

export interface CollectivePublicKey {
  /**
   * The BabyJubJub point that is the collective public key.
   * Equals PK = Σ_i a_{i,0}·G, the sum of each contributor's zeroth Feldman
   * commitment. The contract accumulates this incrementally as contributions
   * are accepted. Retrieve it with `client.getCollectivePublicKey(epochId)`.
   *
   * NOTE: The on-chain `collectivePublicKeyHash` is keccak256(x, y).
   */
  x: bigint;
  y: bigint;
}

/**
 * Wait until a epoch is Finalized, then return the collective public key hash
 * (keccak256 of the key point, emitted in the EpochFinalized event).
 *
 * To get the actual curve point (x, y) call `client.getCollectivePublicKey(epochId)` —
 * a simple view-call that returns the key accumulated on-chain during contribution.
 */
export async function waitForCollectivePublicKeyHash(
  client: DKGClient,
  epochId: `0x${string}`,
  options?: { intervalMs?: number; timeoutMs?: number },
): Promise<`0x${string}`> {
  await waitForEpochPhase(client, epochId, EpochPhase.Finalized, options);
  const events = await client.getEpochFinalizedEvents(epochId);
  if (events.length === 0) {
    throw new Error(`No EpochFinalized event found for epoch ${epochId}`);
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
 * Recovery uses baby-step / giant-step DLOG, capped at **2^32** plaintexts.
 * That cap is the SDK-side limit, intended for tests and direct (non-threshold)
 * use; the real DKG protocol does threshold decryption inside the Go committee
 * and supports plaintexts up to 2^50. See `cmd/davinci-dkg-node/dlog.go`.
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
  epochId: `0x${string}`,
  ciphertextIndex: number,
  options?: { intervalMs?: number; timeoutMs?: number },
) {
  return waitForDecryption(client, epochId, ciphertextIndex, options);
}

/**
 * End-to-end encrypt/decrypt flow for testing and documentation.
 *
 * Assumes the epoch was already created. The function encrypts `plaintext`
 * with `collectivePub`, then waits for the on-chain combined decryption to
 * complete.
 *
 * In production these steps happen across different parties: the data producer
 * encrypts and publishes the ciphertext, DKG nodes submit partial decryptions,
 * and any caller with enough partial decryptions calls combineDecryption.
 *
 * @param client         Read-only DKGClient
 * @param epochId        The epoch ID
 * @param collectivePub  The collective public key point [x, y] (from
 *                       `client.getCollectivePublicKey(epochId)`)
 * @param plaintext      Small integer to encrypt/decrypt
 * @param ciphertextIndex  Index to identify which ciphertext to wait for (1-based)
 */
export async function demonstrateEncryptDecryptFlow(
  client: DKGClient,
  epochId: `0x${string}`,
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
  const record = await waitForCombinedDecryption(client, epochId, ciphertextIndex, {
    timeoutMs: 300_000,
  });

  return {
    ciphertext,
    decryptionCompleted: record.completed,
  };
}
