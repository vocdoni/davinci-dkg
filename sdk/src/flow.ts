/**
 * High-level DKG flow helpers.
 *
 * These functions compose the lower-level DKGClient / DKGWriter / monitor
 * primitives into the common end-to-end flows.
 *
 * Full flow:
 *   1. anyone calls createEpoch() once the cadence window allows it
 *   2. DKG nodes call claimSlot() once the seed block is mined
 *   3. Nodes submit compact contributions, each dealing all MaxK pool keys
 *   4. A node finalizes the epoch with one proof over the accepted
 *      contributions (finalizeEpoch) → the epoch goes Live with all MaxK
 *      pool keys P_j and share-commitment Merkle roots stored at once
 *   5. An organizer registers an application (registerApplication), claiming
 *      the next unclaimed pool key P_j. In automatic mode the application
 *      key is PK_aid = P_j; in organizer-locked mode it additionally
 *      registers PK_org and the key is PK_aid = P_j + PK_org
 *   6. A permitted submitter (the registrant, the allow-list or anyone, per
 *      the policy) encrypts under PK_aid and publishes the ciphertext via
 *      DKGWriter.submitCiphertext(), which returns the on-chain-assigned
 *      ciphertext index
 *   7. DKG nodes submit partial decryptions (with a Merkle proof against the
 *      pool key's share root, built from the finalization calldata). For an
 *      organizer-locked application the contract refuses partials and
 *      combines (OrganizerSecretNotRevealed) until its organizer reveals
 *      sk_org once (revealOrganizerSecret); an automatic application needs
 *      no reveal
 *   8. A node calls combineDecryption → DecryptionCombined event emitted
 *   9. Caller can verify the plaintext matches the original message
 */

import { DKGClient } from './client.js';
import {
  EpochPhase,
  type BabyJubPoint,
  type ElGamalCiphertext,
} from './types.js';
import { waitForEpochPhase, waitForDecryption } from './monitor.js';
import { applicationKey, buildElGamal } from './crypto/elgamal.js';

/**
 * Wait until an epoch is Live, then return pool key `keyIndex` (the curve
 * point `P_j`, TE form). `finalizeEpoch` stores every key of the pool, so
 * once the epoch is Live the key is there.
 */
export async function waitForPoolKey(
  client: DKGClient,
  epochId: `0x${string}`,
  keyIndex: number,
  options?: { intervalMs?: number; timeoutMs?: number },
): Promise<BabyJubPoint> {
  await waitForEpochPhase(client, epochId, EpochPhase.Live, options);
  return client.getPoolKey(epochId, keyIndex);
}

/**
 * Encrypt a message under a BabyJubJub public key.
 *
 * @param message    Small integer plaintext (must fit in BabyJubJub scalar)
 * @param pubKey     Public key as a BabyJubJub point [x, y]
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
 * Encrypt a message under an application's key.
 *
 * `PK_aid = P_j` (automatic) or `P_j + PK_org` (organizer-locked) — see
 * `applicationKey`. `poolKey` comes from `client.getPoolKey(epochId,
 * app.poolIndex)` or `client.getApplicationKey(epochId, aid)`; `pkOrg` from
 * `client.getApplication(epochId, aid).organizerPK`, omitted for automatic
 * applications. Both inputs and the result are in TE form.
 *
 * There is no proof of knowledge of the randomness — see
 * `DKGWriter.submitCiphertext` for why — so the result goes straight to the
 * writer.
 *
 * @param message  Small integer plaintext (the committee recovers values < 2^50)
 * @param poolKey  The application's pool key P_j, TE form
 * @param pkOrg    Organizer public key, TE form; omit for automatic applications
 * @param k        Optional randomness; drawn from the CSPRNG when omitted
 */
export async function encryptForApplication(
  message: bigint,
  poolKey: BabyJubPoint,
  pkOrg?: BabyJubPoint,
  k?: bigint,
): Promise<ElGamalCiphertext> {
  const elgamal = await buildElGamal();
  return elgamal.encrypt(message, applicationKey(poolKey, pkOrg), k);
}

/**
 * Decrypt an ElGamal ciphertext given the private key.
 *
 * Recovery uses baby-step / giant-step DLOG, capped at **2^32** plaintexts.
 * That cap is the SDK-side limit, intended for tests and direct (non-threshold)
 * use; the real DKG protocol does threshold decryption inside the Go committee
 * and supports plaintexts up to 2^50. See `node/dlog.go`.
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
  aid: `0x${string}`,
  ciphertextIndex: number,
  options?: { intervalMs?: number; timeoutMs?: number },
) {
  return waitForDecryption(client, epochId, aid, ciphertextIndex, options);
}

/**
 * End-to-end encrypt/decrypt flow for testing and documentation.
 *
 * Assumes the epoch was already created and the application already
 * registered. The function encrypts `plaintext` under `pkAid`, then waits
 * for the on-chain combined decryption to complete.
 *
 * In production these steps happen across different parties: the data
 * producer encrypts and publishes the ciphertext, DKG nodes submit partial
 * decryptions, an organizer-locked application's organizer reveals sk_org
 * once (an automatic application needs no reveal), and any caller with
 * enough partials calls combineDecryption.
 *
 * @param client      Read-only DKGClient
 * @param epochId     The epoch ID
 * @param aid         Application id the ciphertext is submitted under
 * @param pkAid       The application's public key [x, y] (from
 *                    `client.getApplicationKey(epochId, aid)`)
 * @param plaintext   Small integer to encrypt/decrypt
 * @param ciphertextIndex  On-chain-assigned index of the ciphertext to wait for
 *                         (as returned by `DKGWriter.submitCiphertext`)
 */
export async function demonstrateEncryptDecryptFlow(
  client: DKGClient,
  epochId: `0x${string}`,
  aid: `0x${string}`,
  pkAid: BabyJubPoint,
  plaintext: bigint,
  ciphertextIndex: number,
): Promise<{
  ciphertext: ElGamalCiphertext;
  decryptionCompleted: boolean;
}> {
  // 1. Encrypt
  const ciphertext = await encrypt(plaintext, pkAid);

  // 2. Wait for the DKG nodes to decrypt on-chain
  const record = await waitForCombinedDecryption(client, epochId, aid, ciphertextIndex, {
    timeoutMs: 300_000,
  });

  return {
    ciphertext,
    decryptionCompleted: record.completed,
  };
}
