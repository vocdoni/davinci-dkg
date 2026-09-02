import {
  parseEventLogs,
  type WalletClient,
  type Address,
  type Hash,
  type TransactionReceipt,
} from 'viem';
import { dkgManagerAbi, dkgRegistryAbi, dkgAppManagerAbi } from './abi.js';
import {
  type CreateEpochParams,
  type CiphertextPoK,
  type ElGamalCiphertext,
  type DKGWriterConfig,
} from './types.js';
import { DKGClient } from './client.js';
import { fromTEtoRTE } from './crypto/babyjub-form.js';
import { proveOperator, verifyCiphertextPoK } from './schnorr.js';

/** Outcome of `DKGWriter.submitCiphertext` (the call waits for the receipt). */
export interface SubmitCiphertextResult {
  hash: Hash;
  receipt: TransactionReceipt;
  /**
   * Index the contract assigned to this ciphertext (1, 2, … per
   * `(epochId, aid)`), read back from the `CiphertextSubmitted` event. Use
   * it for `getPlaintext`, `waitForDecryption`, `getCombinedDecryption`.
   */
  ciphertextIndex: number;
}

/**
 * Write client for the DKG Manager and Registry contracts.
 *
 * Extends the read-only DKGClient with transaction-sending methods.
 * Requires a viem `WalletClient` in addition to a `PublicClient`.
 */
export class DKGWriter extends DKGClient {
  readonly walletClient: WalletClient;
  private _writerAccount: Address;

  constructor(config: DKGWriterConfig) {
    super(config);
    this.walletClient = config.walletClient;
    const account = config.walletClient.account;
    if (!account) throw new Error('DKGWriter: walletClient must have an account set');
    this._writerAccount = account.address;
  }

  // ── DKGManager write functions ─────────────────────────────────────────────

  /**
   * Create a new DKG epoch. Permissionless, but only succeeds once
   * `block.number >= nextEpochStartBlock()` (the cadence guard).
   *
   * @param policy  the four `createEpoch` arguments (threshold, committeeSize,
   *                minValidContributions, lotteryAlphaBps). Phase deadlines are
   *                derived on-chain from `EPOCH_DURATION_BLOCKS`, and the
   *                values must satisfy `getEpochBounds()` plus
   *                `1 ≤ threshold ≤ minValidContributions ≤ committeeSize ≤ MaxN`
   *                and `lotteryAlphaBps ≥ 10000`, else `InvalidPolicy()`.
   *                A full `EpochPolicy` is accepted too; its deadline fields
   *                are ignored.
   * @returns The transaction hash. Once mined, derive the id with
   *          `buildEpochId(prefix, epochNonce())`.
   */
  async createEpoch(policy: CreateEpochParams): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'createEpoch',
      args: [
        policy.threshold,
        policy.committeeSize,
        policy.minValidContributions,
        policy.lotteryAlphaBps,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Claim a lottery slot in a epoch.
   * The caller must be a registered and active DKG node.
   * The seed block (seedBlock from the epoch) must have been mined.
   */
  async claimSlot(epochId: `0x${string}`): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'claimSlot',
      args: [epochId as any],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  // `extendRegistration` was removed in SDK 0.2.0 alongside the Solidity
  // auto-cadence refactor. With `EPOCH_DURATION_BLOCKS` driving the
  // schedule, a stalled registration just gets aborted (`abortEpoch`) and
  // the next scheduled epoch picks up automatically.


  /**
   * Submit a contribution (ZK proof + encrypted shares) for a epoch.
   * Only callable by selected participants.
   */
  async submitContribution(
    epochId: `0x${string}`,
    contributorIndex: number,
    commitmentsHash: `0x${string}`,
    encryptedSharesHash: `0x${string}`,
    transcript: `0x${string}`,
    proof: `0x${string}`,
    input: `0x${string}`,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'submitContribution',
      args: [
        epochId as any,
        contributorIndex,
        commitmentsHash,
        encryptedSharesHash,
        transcript,
        proof,
        input,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Finalize a epoch by submitting the aggregate commitments and collective
   * public key (ZK proof required).
   */
  async finalizeEpoch(
    epochId: `0x${string}`,
    aggregateCommitmentsHash: `0x${string}`,
    collectivePublicKeyHash: `0x${string}`,
    shareCommitmentHash: `0x${string}`,
    transcript: `0x${string}`,
    proof: `0x${string}`,
    input: `0x${string}`,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'finalizeEpoch',
      args: [
        epochId as any,
        aggregateCommitmentsHash,
        collectivePublicKeyHash,
        shareCommitmentHash,
        transcript,
        proof,
        input,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Submit a partial decryption for a ciphertext.
   *
   * `aid` binds the proof transcript to a specific application.
   * Pass `0x00…00` for the legacy per-epoch path.
   *
   * `c1`/`c2` are the on-chain ciphertext coords (RTE form). The
   * contract verifies they match the stored ciphertext hash and binds
   * pi[5..6] to c1. Pass them as TE coords; the
   * writer converts to RTE before sending, matching the convention used
   * by `submitCiphertext`.
   */
  async submitPartialDecryption(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    participantIndex: number,
    ciphertextIndex: number,
    c1x: bigint, c1y: bigint, c2x: bigint, c2y: bigint,
    deltaHash: `0x${string}`,
    proof: `0x${string}`,
    input: `0x${string}`,
  ): Promise<Hash> {
    const [c1xR, c1yR] = fromTEtoRTE(c1x, c1y);
    const [c2xR, c2yR] = fromTEtoRTE(c2x, c2y);
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'submitPartialDecryption',
      args: [epochId as any, aid as any, participantIndex, ciphertextIndex,
        c1xR, c1yR, c2xR, c2yR, deltaHash, proof, input],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Submit a ciphertext to be threshold-decrypted by the committee. The
   * epoch must be Live and, for a registered application, the submission
   * must pass its `AppPolicy` (authorized submitter, block window, cap).
   *
   * The ciphertext index is assigned on-chain (1, 2, … per `(epochId, aid)`)
   * and returned in the result, read back from the `CiphertextSubmitted`
   * event; this method therefore waits for the receipt.
   *
   * `pok` is the Schnorr proof of knowledge of the ElGamal randomness bound
   * to exactly this `(epochId, aid)` — obtain it from `encryptWithProof` or
   * `proveCiphertext`. The contract only records it, but every committee
   * node verifies it before releasing a partial decryption, so a ciphertext
   * with a bad proof is silently never decrypted. The writer verifies the
   * proof locally and throws rather than sending such a ciphertext.
   *
   * `ciphertext` is expected in circomlib TE form (what this SDK's
   * `encrypt` returns). It is converted to gnark RTE form just before
   * sending so the contract's on-curve check accepts it; the proof is
   * already in RTE form. See `crypto/babyjub-form.ts`.
   */
  async submitCiphertext(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    ciphertext: ElGamalCiphertext,
    pok: CiphertextPoK,
  ): Promise<SubmitCiphertextResult> {
    const [c1xR, c1yR] = fromTEtoRTE(ciphertext.c1[0], ciphertext.c1[1]);
    const [c2xR, c2yR] = fromTEtoRTE(ciphertext.c2[0], ciphertext.c2[1]);
    if (!verifyCiphertextPoK(epochId, aid, c1xR, c1yR, c2xR, c2yR, pok)) {
      throw new Error(
        'submitCiphertext: ciphertext proof of knowledge does not verify for this (epochId, aid); ' +
          'build it with encryptWithProof() / proveCiphertext() for exactly this epoch and application',
      );
    }
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'submitCiphertext',
      args: [epochId as any, aid as any, c1xR, c1yR, c2xR, c2yR, pok.ax, pok.ay, pok.z],
      account: this._writerAccount,
    });
    const hash = await this.walletClient.writeContract(request);
    const receipt = await this.publicClient.waitForTransactionReceipt({ hash });
    if (receipt.status !== 'success') {
      throw new Error(`submitCiphertext: transaction ${hash} reverted`);
    }
    const manager = this.managerAddress.toLowerCase();
    const submitted = parseEventLogs({
      abi: dkgManagerAbi,
      eventName: 'CiphertextSubmitted',
      logs: receipt.logs,
    }).find(
      (l) =>
        l.address.toLowerCase() === manager &&
        String(l.args.epochId).toLowerCase() === epochId.toLowerCase() &&
        String(l.args.aid).toLowerCase() === aid.toLowerCase(),
    );
    if (!submitted) {
      throw new Error(`submitCiphertext: no CiphertextSubmitted event in receipt of ${hash}`);
    }
    return { hash, receipt, ciphertextIndex: Number(submitted.args.ciphertextIndex) };
  }

  /**
   * Combine partial decryptions to finalize a decryption. The on-chain
   * `CombinedDecryptionRecord` will hold the recovered `plaintext` and a
   * `DecryptionCombined` event is emitted.
   *
   * `aid` is the per-application identifier. Pass `0x00…00` for the
   * legacy per-epoch path that doesn't use applications.
   */
  async combineDecryption(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    ciphertextIndex: number,
    combineHash: `0x${string}`,
    plaintext: bigint,
    transcript: `0x${string}`,
    proof: `0x${string}`,
    input: `0x${string}`,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'combineDecryption',
      args: [
        epochId as any,
        aid as any,
        ciphertextIndex,
        combineHash,
        plaintext,
        transcript,
        proof,
        input,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  // ── Application lifecycle (P8/P9) ──────────────────────────────────────────

  /**
   * Register a public-derivation (mode 0) application against `(epochId, aid)`.
   * The contract derives `S = keccak256(eid || PK_ep || aid) % L` on-chain;
   * callers can preview the same value via `computeS()` from `~/derive`.
   */
  async registerApplication(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    policy: import('./types.js').AppPolicy,
  ): Promise<Hash> {
    const appManagerAddress = await this._getAppManagerAddress();
    const { request } = await this.publicClient.simulateContract({
      address: appManagerAddress,
      abi: dkgAppManagerAbi,
      functionName: 'registerApplication',
      args: [epochId as any, aid as any, policy as any],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Register an organizer-co-decryption (mode 1) application. The
   * `(pkOrgX, pkOrgY)` pair must arrive in TE form — the writer converts
   * to RTE before signing. The Schnorr proof `(ax, ay, z)` arrives in RTE
   * form (it's computed against the on-chain transcript).
   */
  async registerApplicationCoDec(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    policy: import('./types.js').AppPolicy,
    pkOrgX: bigint,
    pkOrgY: bigint,
    schnorrAx: bigint,
    schnorrAy: bigint,
    schnorrZ: bigint,
  ): Promise<Hash> {
    const [pkXR, pkYR] = fromTEtoRTE(pkOrgX, pkOrgY);
    const appManagerAddress = await this._getAppManagerAddress();
    const { request } = await this.publicClient.simulateContract({
      address: appManagerAddress,
      abi: dkgAppManagerAbi,
      functionName: 'registerApplicationCoDec',
      args: [
        epochId as any,
        aid as any,
        policy as any,
        pkXR,
        pkYR,
        schnorrAx,
        schnorrAy,
        schnorrZ,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Submit the organizer's `Δ_org = sk_org · C_1` share (with Chaum-Pedersen
   * DLEQ binding it to `PK_org`). Only used in mode 1 (organizer co-decryption).
   *
   * Coordinate forms:
   *   - `c1`/`c2` arrive in TE form; the writer converts to RTE so they
   *     match the on-chain ciphertext hash.
   *   - `deltaOrg` arrives in TE form; the writer converts to RTE.
   *
   * The contract checks the converted (c1, c2) against the stored
   * ciphertext hash and binds pi[5..6] to c1.
   */
  async submitOrganizerShare(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    ciphertextIndex: number,
    c1x: bigint, c1y: bigint, c2x: bigint, c2y: bigint,
    deltaOrgX: bigint,
    deltaOrgY: bigint,
    dleqProof: `0x${string}`,
    dleqInput: `0x${string}`,
  ): Promise<Hash> {
    const [c1xR, c1yR] = fromTEtoRTE(c1x, c1y);
    const [c2xR, c2yR] = fromTEtoRTE(c2x, c2y);
    const [dxR, dyR] = fromTEtoRTE(deltaOrgX, deltaOrgY);
    const appManagerAddress = await this._getAppManagerAddress();
    const { request } = await this.publicClient.simulateContract({
      address: appManagerAddress,
      abi: dkgAppManagerAbi,
      functionName: 'submitOrganizerShare',
      args: [
        epochId as any,
        aid as any,
        ciphertextIndex,
        c1xR, c1yR, c2xR, c2yR,
        dxR,
        dyR,
        dleqProof,
        dleqInput,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Abort a dead epoch. Permissionless — anyone may call it — but the
   * contract only accepts it once the epoch can no longer progress: the
   * committee-selection deadline passed without a full committee, or the
   * key-assembly deadline passed with fewer than `minValidContributions`
   * accepted. Any other state reverts `InvalidPhase()`.
   */
  async abortEpoch(epochId: `0x${string}`): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'abortEpoch',
      args: [epochId as any],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  // ── DKGRegistry write functions ────────────────────────────────────────────

  /**
   * Register a new BabyJubJub key in the DKG Registry. The caller proves
   * knowledge of the private key via a Schnorr PoK bound to the wallet
   * address; the SDK builds the proof in-browser from `privateKey`
   * and never sends it on the wire. The derived public key + proof are
   * submitted in RTE form, matching the on-chain transcript.
   *
   * `nonce` exists only for tests that need a deterministic witness;
   * production callers should leave it undefined to draw fresh.
   */
  async registerKey(privateKey: bigint, nonce?: bigint): Promise<Hash> {
    const operator = this._writerAccount;
    const { pubX, pubY, proof } = proveOperator(privateKey, operator, nonce);
    const registryAddress = await this._registryAddressResolved();
    const { request } = await this.publicClient.simulateContract({
      address: registryAddress,
      abi: dkgRegistryAbi,
      functionName: 'registerKey',
      args: [pubX, pubY, proof.ax, proof.ay, proof.z],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Update an existing registered key. Requires a fresh Schnorr PoK over
   * the new key (P4 — `updateKey` and `registerKey` enforce the same check
   * to prevent silent key replacement without proof of knowledge).
   */
  async updateKey(privateKey: bigint, nonce?: bigint): Promise<Hash> {
    const operator = this._writerAccount;
    const { pubX, pubY, proof } = proveOperator(privateKey, operator, nonce);
    const registryAddress = await this._registryAddressResolved();
    const { request } = await this.publicClient.simulateContract({
      address: registryAddress,
      abi: dkgRegistryAbi,
      functionName: 'updateKey',
      args: [pubX, pubY, proof.ax, proof.ay, proof.z],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Send a heartbeat to keep the node active.
   */
  async heartbeat(): Promise<Hash> {
    const registryAddress = await this._registryAddressResolved();
    const { request } = await this.publicClient.simulateContract({
      address: registryAddress,
      abi: dkgRegistryAbi,
      functionName: 'heartbeat',
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Reactivate a node that was previously reaped.
   */
  async reactivate(): Promise<Hash> {
    const registryAddress = await this._registryAddressResolved();
    const { request } = await this.publicClient.simulateContract({
      address: registryAddress,
      abi: dkgRegistryAbi,
      functionName: 'reactivate',
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Permissionlessly reap a stale node that has exceeded the inactivity window.
   */
  async reap(operator: Address): Promise<Hash> {
    const registryAddress = await this._registryAddressResolved();
    const { request } = await this.publicClient.simulateContract({
      address: registryAddress,
      abi: dkgRegistryAbi,
      functionName: 'reap',
      args: [operator],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  // ── Convenience helpers ────────────────────────────────────────────────────

  /**
   * Wait for a transaction to be included and return its receipt.
   */
  async waitForTransaction(hash: Hash) {
    return this.publicClient.waitForTransactionReceipt({ hash });
  }

  /**
   * Create an epoch and wait for the receipt.
   * Returns the transaction receipt (check `status === 'success'`).
   */
  async createRoundAndWait(policy: CreateEpochParams) {
    const hash = await this.createEpoch(policy);
    return this.waitForTransaction(hash);
  }
}
