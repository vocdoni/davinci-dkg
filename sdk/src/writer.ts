import {
  getContract,
  type WalletClient,
  type PublicClient,
  type Address,
  type Hash,
} from 'viem';
import { dkgManagerAbi, dkgRegistryAbi } from './abi.js';
import {
  type RoundPolicy,
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
   * Create a new DKG round.
   *
   * @param policy            committee / phase policy for the round.
   * @param decryptionPolicy  gate on `submitCiphertext` (owner-only, time
   *                          windows, submission cap). Defaults to fully open —
   *                          anyone can submit, no caps, no windows. Pair with
   *                          `OpenDecryptionPolicy` for the permissive default.
   * @returns The transaction hash. Use `waitForRoundId` to obtain the round ID
   *          once the tx is mined.
   */
  async createRound(
    policy: RoundPolicy,
    decryptionPolicy: DecryptionPolicy = OpenDecryptionPolicy,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'createRound',
      args: [
        policy.threshold,
        policy.committeeSize,
        policy.minValidContributions,
        policy.lotteryAlphaBps,
        policy.seedDelay,
        policy.registrationDeadlineBlock,
        policy.contributionDeadlineBlock,
        policy.finalizeNotBeforeBlock,
        policy.disclosureAllowed,
        decryptionPolicy,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Claim a lottery slot in a round.
   * The caller must be a registered and active DKG node.
   * The seed block (seedBlock from the round) must have been mined.
   */
  async claimSlot(roundId: `0x${string}`): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'claimSlot',
      args: [roundId as any],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Extend the registration deadline of a round.
   * Only callable by the round organizer.
   */
  async extendRegistration(roundId: `0x${string}`): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'extendRegistration',
      args: [roundId as any],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Submit a contribution (ZK proof + encrypted shares) for a round.
   * Only callable by selected participants.
   */
  async submitContribution(
    roundId: `0x${string}`,
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
        roundId as any,
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
   * Finalize a round by submitting the aggregate commitments and collective
   * public key (ZK proof required).
   */
  async finalizeRound(
    roundId: `0x${string}`,
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
      functionName: 'finalizeRound',
      args: [
        roundId as any,
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
   */
  async submitPartialDecryption(
    roundId: `0x${string}`,
    participantIndex: number,
    ciphertextIndex: number,
    deltaHash: `0x${string}`,
    proof: `0x${string}`,
    input: `0x${string}`,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'submitPartialDecryption',
      args: [roundId as any, participantIndex, ciphertextIndex, deltaHash, proof, input],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Submit a ciphertext to be threshold-decrypted by the committee.
   * The round must be Finalized and the submission must pass the round's
   * DecryptionPolicy (owner-only, time windows, max count).
   *
   * `ciphertextIndex` is caller-chosen and write-once per round.
   *
   * Inputs are expected in circomlib TE form (the form that this SDK's
   * `encrypt` returns and that davinci-sdk also uses). They are converted
   * to gnark RTE form just before sending so the contract's on-curve check
   * (`_isOnBabyJubJub`, in RTE) accepts them. See `crypto/babyjub-form.ts`.
   */
  async submitCiphertext(
    roundId: `0x${string}`,
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
      args: [roundId as any, ciphertextIndex, c1xR, c1yR, c2xR, c2yR],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Combine partial decryptions to finalize a decryption. The on-chain
   * `CombinedDecryptionRecord` will hold the recovered `plaintext` and an
   * `DecryptionCombined` event is emitted.
   */
  async combineDecryption(
    roundId: `0x${string}`,
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
        roundId as any,
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

  /**
   * Submit a revealed share (for disclosure mode rounds).
   */
  async submitRevealedShare(
    roundId: `0x${string}`,
    participantIndex: number,
    shareValue: bigint,
    proof: `0x${string}`,
    input: `0x${string}`,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'submitRevealedShare',
      args: [roundId as any, participantIndex, shareValue, proof, input],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Reconstruct the secret from revealed shares.
   */
  async reconstructSecret(
    roundId: `0x${string}`,
    disclosureHash: `0x${string}`,
    reconstructedSecretHash: `0x${string}`,
    transcript: `0x${string}`,
    proof: `0x${string}`,
    input: `0x${string}`,
  ): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'reconstructSecret',
      args: [
        roundId as any,
        disclosureHash,
        reconstructedSecretHash,
        transcript,
        proof,
        input,
      ],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Abort a round. Only callable by the organizer when the round
   * has not reached the minimum contribution threshold.
   */
  async abortRound(roundId: `0x${string}`): Promise<Hash> {
    const { request } = await this.publicClient.simulateContract({
      address: this.managerAddress,
      abi: dkgManagerAbi,
      functionName: 'abortRound',
      args: [roundId as any],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  // ── DKGRegistry write functions ────────────────────────────────────────────

  /**
   * Register a new BabyJubJub key in the DKG Registry.
   * The caller becomes an active DKG node.
   */
  async registerKey(pubX: bigint, pubY: bigint): Promise<Hash> {
    const registryAddress = await this._registryAddressResolved();
    const { request } = await this.publicClient.simulateContract({
      address: registryAddress,
      abi: dkgRegistryAbi,
      functionName: 'registerKey',
      args: [pubX, pubY],
      account: this._writerAccount,
    });
    return this.walletClient.writeContract(request);
  }

  /**
   * Update an existing registered key.
   */
  async updateKey(pubX: bigint, pubY: bigint): Promise<Hash> {
    const registryAddress = await this._registryAddressResolved();
    const { request } = await this.publicClient.simulateContract({
      address: registryAddress,
      abi: dkgRegistryAbi,
      functionName: 'updateKey',
      args: [pubX, pubY],
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
   * Create a round and wait for the receipt.
   * Returns the transaction receipt (check `status === 'success'`).
   */
  async createRoundAndWait(policy: RoundPolicy) {
    const hash = await this.createRound(policy);
    return this.waitForTransaction(hash);
  }
}
