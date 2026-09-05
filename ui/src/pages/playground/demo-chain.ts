// Deterministic transaction simulator for `?demo=1`.
//
// Demo mode has no chain and no wallet, but the playground still has to show a
// transaction hash, a block and a gas figure for every step, and partial
// decryptions that arrive over time. This module invents all of that from a
// counter: same inputs, same hashes, same blocks — so a screenshot taken today
// matches one taken next month, and the whole thing is unit-testable without a
// clock.
//
// The cryptography around it is *not* simulated: the organizer secret, the
// pool key, the ElGamal ciphertext and the plaintext recovery are computed by
// the real SDK and the real curve in both modes. Only the transport is fake.

import { addPoint, Base8, Fr, mulPointEscalar, subOrder, type Point } from '@zk-kit/baby-jubjub'
import { keccak256, stringToHex, type Address, type Hex } from 'viem'
import type { BabyJubPoint, ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { GAS } from '~fixtures/synthetic'
import type { AppModeName } from '~indexer/types'
import type { PartialView } from './types'

/** The address the demo wallet claims to be. Fixed, so links stay stable. */
export const DEMO_ACCOUNT = '0x00d99200a17e5c1a9b6d0e0c58cbe3d6c6a3d001' as Address

/** Wall-clock ms per simulated block. Faster than a real chain on purpose. */
export const DEMO_BLOCK_MS = 1_200

export interface DemoTx {
  hash: Hex
  block: number
  gasUsed: number
}

export interface DemoChainState {
  /** Simulator head. Advances on its own clock and on every transaction. */
  block: number
  /** Transactions sent so far — the only entropy the hashes have. */
  seq: number
  register: DemoTx | null
  /** Mode of the registered application; decides whether a reveal is needed. */
  mode: AppModeName
  /** Pool key the registration claimed. */
  poolIndex: number | null
  ciphertext: { tx: DemoTx; index: number } | null
  reveal: DemoTx | null
}

export function initialDemoChain(headBlock: number): DemoChainState {
  return {
    block: headBlock,
    seq: 0,
    register: null,
    mode: 'organizer-locked',
    poolIndex: null,
    ciphertext: null,
    reveal: null,
  }
}

/**
 * A transaction hash that depends only on what the transaction *is*. Real
 * enough to be copied into a search box and obviously synthetic once you look
 * at where it points.
 */
export function demoTxHash(label: string, seq: number): Hex {
  return keccak256(stringToHex(`davinci-dkg:demo-tx:${seq}:${label}`))
}

/** Deterministic committee member address for slot `i` of a demo epoch. */
export function demoParticipant(epochId: string, participantIndex: number): Address {
  const digest = keccak256(stringToHex(`davinci-dkg:demo-operator:${epochId}:${participantIndex}`))
  return `0x${digest.slice(26)}` as Address
}

/**
 * The secret behind pool key `j` of a demo epoch. The fixture's pool keys are
 * decorative bytes; the demo needs real points so the organizer's crypto is
 * the real crypto, and this is the one place that knows the scalar.
 */
export function demoPoolSecret(epochId: string, keyIndex: number): bigint {
  const digest = keccak256(stringToHex(`davinci-dkg:demo-pool-key:${epochId}:${keyIndex}`))
  return BigInt(digest) % subOrder || 1n
}

/** `P_j = sk_j·G` for the demo epoch, TE form. */
export function demoPoolKey(epochId: string, keyIndex: number): BabyJubPoint {
  return mulPointEscalar(Base8, demoPoolSecret(epochId, keyIndex)) as BabyJubPoint
}

export interface DemoSendResult {
  state: DemoChainState
  tx: DemoTx
}

/** Mine one block and put `label`'s transaction in it. */
export function sendDemoTx(state: DemoChainState, label: string, gasUsed: number): DemoSendResult {
  const seq = state.seq + 1
  const block = state.block + 1
  const tx: DemoTx = { hash: demoTxHash(label, seq), block, gasUsed }
  return { state: { ...state, seq, block }, tx }
}

/** Register `aid`, claiming pool key `poolIndex` — the epoch's next free one. */
export function demoRegister(
  state: DemoChainState,
  aid: string,
  mode: AppModeName = 'organizer-locked',
  poolIndex = 0
): DemoSendResult {
  const sent = sendDemoTx(state, `registerApplication:${aid}`, GAS.registerApplication)
  return { state: { ...sent.state, register: sent.tx, mode, poolIndex }, tx: sent.tx }
}

export function demoSubmitCiphertext(
  state: DemoChainState,
  aid: string
): DemoSendResult & { ciphertextIndex: number } {
  const index = (state.ciphertext?.index ?? 0) + 1
  const gas = index === 1 ? GAS.submitCiphertextFirst : GAS.submitCiphertext
  const sent = sendDemoTx(state, `submitCiphertext:${aid}:${index}`, gas)
  return {
    state: { ...sent.state, ciphertext: { tx: sent.tx, index } },
    tx: sent.tx,
    ciphertextIndex: index,
  }
}

/** Reveal the organizer secret — once; a second call is the contract's `AlreadyRevealed()`. */
export function demoRevealSecret(state: DemoChainState, aid: string): DemoSendResult {
  if (state.reveal) throw new Error('AlreadyRevealed()')
  const sent = sendDemoTx(state, `revealOrganizerSecret:${aid}`, GAS.revealOrganizerSecret)
  return { state: { ...sent.state, reveal: sent.tx }, tx: sent.tx }
}

/** Advance the simulator by one block without sending anything. */
export function tickDemoChain(state: DemoChainState): DemoChainState {
  return { ...state, block: state.block + 1 }
}

// ── decryption ───────────────────────────────────────────────────────────────

export interface DemoDecryptionParams {
  epochId: string
  threshold: number
  committeeSize: number
  /** The node's per-slot decryption delay; `wave = ⌊Δblocks / stagger⌋`. */
  staggerBlocks: number
  /** Ciphertext index, which seeds the rotation the committee answers in. */
  ciphertextIndex: number
}

/**
 * Partials visible at `head`, counting from `openBlock` — the block the
 * ciphertext became decryptable: its own block for an automatic application,
 * the later of that and the reveal for an organizer-locked one. The contract
 * refuses every partial of a locked application before the reveal, so the
 * caller passes nothing (and gets nothing) until the secret is out.
 *
 * Modelled on what the fixture generates, which is what the node actually
 * does: the seed-derived rotation picks a starting slot, its first `t` members
 * answer inside one stagger window (wave 0, spread over `staggerBlocks`
 * blocks), and a handful of late members follow one window later (wave 1). The
 * wave number is then exactly the indexer's
 * `⌊(partial block − opened block) / staggerBlocks⌋`.
 */
export function demoPartialsAt(openBlock: number, head: number, params: DemoDecryptionParams): PartialView[] {
  const { threshold, committeeSize, staggerBlocks, epochId, ciphertextIndex } = params
  if (threshold <= 0 || committeeSize <= 0) return []
  const stagger = Math.max(1, staggerBlocks)
  const rotation = ciphertextIndex % committeeSize
  const late = Math.min(Math.max(0, committeeSize - threshold), ciphertextIndex % 3 === 0 ? 5 : 2)
  const out: PartialView[] = []
  for (let k = 0; k < threshold + late; k++) {
    const isLate = k >= threshold
    const offset = isLate ? k - threshold : k
    const block = openBlock + 1 + (isLate ? stagger : 0) + (offset % stagger)
    if (block > head) continue
    const participantIndex = ((rotation + k) % committeeSize) + 1
    out.push({
      participantIndex,
      participant: demoParticipant(epochId, participantIndex),
      block,
      wave: Math.floor((block - openBlock) / stagger),
      tx: demoTxHash(`submitPartialDecryption:${ciphertextIndex}:${participantIndex}`, participantIndex),
    })
  }
  return out.sort((a, b) => a.block - b.block || a.participantIndex - b.participantIndex)
}

export interface DemoCombine {
  block: number
  tx: Hex
  gasUsed: number
}

/**
 * When the combine lands: one block after the application is unlocked *and*
 * the `t`-th partial is in. `unlockedAt` is 0 for an automatic application
 * (unlocked from the start), the reveal block for an organizer-locked one,
 * and null while the organizer keeps its secret — which stops the combine
 * forever, exactly the point the "keep" branch is making. Late responders
 * past `t` do not hold it up.
 */
export function demoCombineAt(
  partials: PartialView[],
  unlockedAt: number | null,
  head: number,
  params: DemoDecryptionParams
): DemoCombine | null {
  if (unlockedAt == null) return null
  if (params.threshold <= 0 || partials.length < params.threshold) return null
  const thresholdBlock = partials[params.threshold - 1].block
  const block = Math.max(thresholdBlock, unlockedAt) + 1
  if (block > head) return null
  return {
    block,
    tx: demoTxHash(`combineDecryption:${params.ciphertextIndex}`, block),
    gasUsed: GAS.combineDecryption,
  }
}

/**
 * Recover the plaintext the way the committee's combine proof does — from the
 * full application secret `sk_aid = sk_j + sk_org` (or just `sk_j` for an
 * automatic application), which only the simulator ever holds in one place:
 *
 *   m·G = C2 − sk_aid·C1,  then m = dlog(m·G)
 *
 * The dlog is a plain linear walk rather than the SDK's baby-step/giant-step,
 * because BSGS pays a 2^16-entry table up front and a demo recovers a number
 * someone just typed into a box. Values past `DEMO_DLOG_LIMIT` return null and
 * the step says the plaintext could not be recovered locally.
 */
export const DEMO_DLOG_LIMIT = 1n << 20n

export function demoRecoverPlaintext(ciphertext: ElGamalCiphertext, skAid: bigint): bigint | null {
  const shared = mulPointEscalar(ciphertext.c1 as Point<bigint>, skAid)
  const target = addPoint(ciphertext.c2 as Point<bigint>, [Fr.neg(shared[0]), shared[1]])
  let current: Point<bigint> = [0n, 1n]
  for (let m = 0n; m <= DEMO_DLOG_LIMIT; m++) {
    if (current[0] === target[0] && current[1] === target[1]) return m
    current = addPoint(current, Base8)
  }
  return null
}
