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
  skOrg: bigint
  /** 0 means "the registering address". */
  authorizedSubmitter: Address
  /** 0 means unlimited. */
  maxCiphertexts: number
  /** Schnorr witness, pinned by the caller so the printed words are the sent words. */
  nonce: bigint
}

export interface SubmitArgs {
  aid: Hex
  ciphertext: ElGamalCiphertext
}

export interface ShareArgs extends SubmitArgs {
  ciphertextIndex: number
  skOrg: bigint
  /** DLEQ witness, pinned by the caller for the same reason. */
  nonce: bigint
}

export interface PartialView {
  participantIndex: number
  participant: Address | null
  block: number
  /** `floor((partialBlock − ciphertextBlock) / staggerBlocks)`. */
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
  share: { present: boolean; block: number | null; tx: Hex | null }
  combined: { done: boolean; block: number | null; tx: Hex | null; plaintext: bigint | null }
  /** The `(C1, C2)` the chain actually stores, for the local verification. */
  onChain: ElGamalCiphertext | null
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
  /** `PK_ep` of the target epoch, TE form; null until the epoch is Live. */
  epochKey: BabyJubPoint | null
  register(args: RegisterArgs): Promise<TxRecord>
  submitCiphertext(args: SubmitArgs): Promise<{ tx: TxRecord; ciphertextIndex: number }>
  releaseShare(args: ShareArgs): Promise<TxRecord>
  /** Null while the ciphertext is not visible on chain yet. */
  decryption(ciphertextIndex: number | null): DecryptionView | null
}
