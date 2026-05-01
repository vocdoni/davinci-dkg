import {
  getContract,
  type WalletClient,
  type PublicClient,
  type Address,
  type Hash,
} from 'viem';
import { dkgManagerAbi, dkgRegistryAbi } from './abi.js';
import {
  type EpochPolicy,
  type DecryptionPolicy,
  type DKGWriterConfig,
  OpenDecryptionPolicy,
} from './types.js';
import { DKGClient } from './client.js';
import { fromTEtoRTE } from './crypto/babyjub-form.js';

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
   * Create a new DKG epoch.
   *
   * @param policy            committee / phase policy for the epoch.
   * @param decryptionPolicy  gate on `submitCiphertext` (owner-only, time
   *                          windows, submission cap). Defaults to fully open —
   *                          anyone can submit, no caps, no windows. Pair with
   *                          `OpenDecryptionPolicy` for the permissive default.
   * @returns The transaction hash. Use `waitForRoundId` to obtain the epoch ID
   *          once the tx is mined.
   */
  async createEpoch(
    policy: EpochPolicy,
    decryptionPolicy: DecryptionPolicy = OpenDecryptionPolicy,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'createEpoch',
      args: [
        policy.threshold,
        policy.committeeSize,
        policy.minValidContributions,
        policy.lotteryAlphaBps,
        policy.seedDelay,
        policy.registrationDeadlineBlock,
        policy.contributionDeadlineBlock,
        policy.finalizeNotBeforeBlock,
        decryptionPolicy,
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

  /**
   * Extend the registration deadline of a epoch.
   * Only callable by the epoch organizer.
   */
  async extendRegistration(epochId: `0x${string}`): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'extendRegistration',
      args: [epochId as any],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Submit a contribution (ZK proof + encrypted shares) for a epoch.
   * Only callable by selected participants.
   */
  async submitContribution(
    epochId: `0x${string}`,
    contributorIndex: number,
    commitmentsHash: `0x${string}`,
    encryptedSharesHash: `0x${string}`,
    commitment0X: bigint,
    commitment0Y: bigint,
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
        commitment0X,
        commitment0Y,
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
   * `aid` binds the proof transcript to a specific application (P9).
   * Pass `0x00…00` for the legacy per-epoch path.
   *
   * `c1`/`c2` are the on-chain ciphertext coords (RTE form). The
   * contract verifies they match the stored ciphertext hash and binds
   * pi[5..6] to c1 (CIRCUITS_AUDIT #2). Pass them as TE coords; the
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
   * Submit a ciphertext to be threshold-decrypted by the committee.
   * The epoch must be Finalized and the submission must pass the epoch's
   * DecryptionPolicy (owner-only, time windows, max count).
   *
   * `ciphertextIndex` is caller-chosen and write-once per epoch.
   *
   * Inputs are expected in circomlib TE form (the form that this SDK's
   * `encrypt` returns and that davinci-sdk also uses). They are converted
   * to gnark RTE form just before sending so the contract's on-curve check
   * (`_isOnBabyJubJub`, in RTE) accepts them. See `crypto/babyjub-form.ts`.
   */
  async submitCiphertext(
    epochId: `0x${string}`,
    aid: `0x${string}`,
    ciphertextIndex: number,
    c1x: bigint,
    c1y: bigint,
    c2x: bigint,
    c2y: bigint,
  ): Promise<Hash> {
    const [c1xR, c1yR] = fromTEtoRTE(c1x, c1y);
    const [c2xR, c2yR] = fromTEtoRTE(c2x, c2y);
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'submitCiphertext',
      args: [epochId as any, aid as any, ciphertextIndex, c1xR, c1yR, c2xR, c2yR],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Combine partial decryptions to finalize a decryption. The on-chain
   * `CombinedDecryptionRecord` will hold the recovered `plaintext` and a
   * `DecryptionCombined` event is emitted.
   *
   * `aid` is the per-application identifier (P9). Pass `0x00…00` for the
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
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
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
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
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
   * Coordinate forms (CIRCUITS_AUDIT #1):
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
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
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
   * Abort a epoch. Only callable by the organizer when the epoch
   * has not reached the minimum contribution threshold.
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
   * address (P4); the SDK builds the proof in-browser from `privateKey`
   * and never sends it on the wire. The derived public key + proof are
   * submitted in RTE form, matching the on-chain transcript.
   *
   * `nonce` exists only for tests that need a deterministic witness;
   * production callers should leave it undefined to draw fresh.
   */
  async registerKey(privateKey: bigint, nonce?: bigint): Promise<Hash> {
    const operator = this._writerAccount;
    const { proveOperator } = await import('./schnorr.js');
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
    const { proveOperator } = await import('./schnorr.js');
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
   * Create a epoch and wait for the receipt.
   * Returns the transaction receipt (check `status === 'success'`).
   */
  async createRoundAndWait(policy: EpochPolicy) {
    const hash = await this.createEpoch(policy);
    return this.waitForTransaction(hash);
  }
}
