// The contract between the stepper UI and whatever is executing its steps.
//
// Two implementations exist and the UI never branches on which one it has:
//
//   • `useLiveChain` — a wagmi wallet plus the SDK's `DKGWriter`, writing to a
//     real deployment. It is the only file in the playground that imports
//     wagmi.
//   • `useDemoChain` — the deterministic simulator in `demo-chain.ts`, driven
//     by the synthetic fixture. It touches neither wagmi nor the network, so
//     the whole walkthrough can be screenshotted (and unit-tested) offline.

import type { Address, Hex } from 'viem'
import type { BabyJubPoint, ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import type { AppModeName } from '~indexer/types'

/** What a step shows about the transaction it sent. */
export interface TxRecord {
  hash: Hex
  /** Null until the receipt (or the indexer) resolves it. */
  block: number | null
  gasUsed: number | null
  /** True for a transaction the demo simulator invented. */
  simulated: boolean
}

export interface PlaygroundWallet {
  connected: boolean
  address: Address | null
  /** Human label for the connected identity ("demo wallet", "0x3F9B…08Ad"). */
  label: string
  /** Opens the wallet modal (live) or connects instantly (demo). */
  connect: () => void
  /** False when the wallet sits on a different chain than `config.chainId`. */
  chainOk: boolean
  /** Why the wallet cannot write right now, if it cannot. */
  problem: string | null
}

export interface RegisterArgs {
  aid: Hex
  /**
   * `sk_org` of an organizer-locked application — `PK_org` and the Schnorr
   * proof are derived from it and it stays here. Null for an automatic
   * application, which has no organizer key.
   */
  skOrg: bigint | null
  mode: AppModeName
  /** Submission allow-list; empty means "the registering address only". */
  submitters: Address[]
  /** 0 means unlimited. */
  maxCiphertexts: number
  /** Decryption window, unix seconds; 0 = unbounded on that side. */
  decryptNotBefore: number
  decryptNotAfter: number
  /**
   * Schnorr witness, pinned by the caller so the printed words are the sent
   * words. Ignored in automatic mode, which sends no proof.
   */
  nonce: bigint
}

export interface RegisterResult {
  tx: TxRecord
  /** The pool key the registration claimed. */
  poolIndex: number
}

/** `PK_aid` and what it is made of, as `getApplicationKey` resolves it. */
export interface ApplicationKeys {
  poolIndex: number
  /** `P_j`, TE form. */
  poolKey: BabyJubPoint
  /** `PK_org`, TE form; null for an automatic application. */
  organizerPK: BabyJubPoint | null
  /** `P_j` or `P_j + PK_org` — the key to encrypt under. */
  key: BabyJubPoint
}

export interface SubmitArgs {
  aid: Hex
  ciphertext: ElGamalCiphertext
}

export interface RevealArgs {
  aid: Hex
  skOrg: bigint
}

export interface PartialView {
  participantIndex: number
  participant: Address | null
  block: number
  /**
   * `floor((partialBlock − openedBlock) / staggerBlocks)`, counted from 0,
   * where the ciphertext opened at its own block or — organizer-locked — at
   * the reveal, whichever came later.
   */
  wave: number
  tx: Hex | null
}

/** Everything the "watch decryption" step draws, from either source. */
export interface DecryptionView {
  threshold: number
  committeeSize: number
  staggerBlocks: number
  ciphertextBlock: number | null
  partials: PartialView[]
  /** The organizer secret: whether the application needs one, and whether it is out. */
  reveal: { required: boolean; done: boolean; block: number | null; tx: Hex | null }
  combined: { done: boolean; block: number | null; tx: Hex | null; plaintext: bigint | null }
  /** The `(C1, C2)` the chain actually stores, for the local verification. */
  onChain: ElGamalCiphertext | null
}

/** How far the target epoch's pool has been dealt. */
export interface PoolView {
  activated: number
  claimed: number
  size: number
}

/**
 * A chain scoped to one `(epoch, aid)` pair — the pair the stepper is working
 * on. Both implementations are hooks, so the target is fixed at construction
 * and every method below reads whatever the current render knows.
 */
export interface PlaygroundTarget {
  epochId: Hex | null
  aid: Hex | null
}

export interface PlaygroundChain {
  kind: 'live' | 'demo'
  wallet: PlaygroundWallet
  /** Chain head (or the simulator's clock). */
  headBlock: number
  /** Pool of the target epoch; null until the epoch is Live. */
  pool: PoolView | null
  register(args: RegisterArgs): Promise<RegisterResult>
  /**
   * `PK_aid` of the registered application: the SDK's `getApplicationKey` on
   * a live chain, the simulator's own key on the demo one. Stable identity,
   * so the controller may depend on it.
   */
  applicationKeys(aid: Hex): Promise<ApplicationKeys>
  submitCiphertext(args: SubmitArgs): Promise<{ tx: TxRecord; ciphertextIndex: number }>
  revealSecret(args: RevealArgs): Promise<TxRecord>
  /** Null while the ciphertext is not visible on chain yet. */
  decryption(ciphertextIndex: number | null): DecryptionView | null
}
