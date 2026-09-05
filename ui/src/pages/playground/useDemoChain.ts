// The demo chain: the stepper's whole surface, backed by `demo-chain.ts` and
// the synthetic fixture. It imports no wagmi, opens no socket and signs
// nothing, so `/playground?demo=1` walks the complete organizer flow — right
// down to the partials arriving one block apart — on a machine with no chain
// in reach.
//
// Two deliberate substitutions, both surfaced in the UI:
//
//   • the wallet is a fixed address that "connects" instantly;
//   • the pool key `P_j` is a real curve point derived from the epoch id and
//     the key index rather than the fixture's decorative one, so the
//     ciphertext and the local verification are genuine rather than
//     arithmetic on a non-curve pair.
//
// Because the simulator knows that `sk_j`, and the organizer hands it
// `sk_org` when it reveals the secret (an automatic application has none), it
// recovers the plaintext the way the committee would —
// `m·G = C2 − (sk_j + sk_org)·C1`, then a dlog — so the closing "verify
// locally" step is a real comparison, not a tautology. The decryption window
// is not simulated: the demo committee answers whenever the partials are due.
// The reveal gate is: an organizer-locked application gets no partials until
// its secret is revealed, exactly as the contract enforces, so the "watch"
// step shows them arriving after the reveal rather than a finished tally.

import { useCallback, useEffect, useMemo, useRef, useState } from 'react'
import { subOrder } from '@zk-kit/baby-jubjub'
import type { Hex } from 'viem'
import type { BabyJubPoint, ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { useStore } from '~data/hooks'
import { resolveHeadBlock } from './head-block'
import { epochKey as epochStoreKey, POOL_SIZE } from '~indexer/types'
import {
  DEMO_ACCOUNT,
  DEMO_BLOCK_MS,
  demoCombineAt,
  demoPartialsAt,
  demoPoolKey,
  demoPoolSecret,
  demoRecoverPlaintext,
  demoRegister,
  demoRevealSecret,
  demoSubmitCiphertext,
  initialDemoChain,
  tickDemoChain,
  type DemoChainState,
  type DemoDecryptionParams,
} from './demo-chain'
import { applicationPublicKey, organizerPublicKey } from './organizer'
import type {
  ApplicationKeys,
  DecryptionView,
  PlaygroundChain,
  PlaygroundTarget,
  RegisterArgs,
  RegisterResult,
  RevealArgs,
  TxRecord,
} from './types'

export function useDemoChain(target: PlaygroundTarget): PlaygroundChain {
  const store = useStore()
  const [connected, setConnected] = useState(false)
  const [chain, setChain] = useState<DemoChainState>(() => initialDemoChain(resolveHeadBlock(store)))
  const [submitted, setSubmitted] = useState<ElGamalCiphertext | null>(null)
  const [plaintext, setPlaintext] = useState<bigint | null>(null)
  const submittedRef = useRef<ElGamalCiphertext | null>(null)
  /** `PK_org` of the registered application; null when automatic. */
  const organizerPKRef = useRef<BabyJubPoint | null>(null)

  // The simulator runs faster than a real chain: waiting 12 s per partial
  // would make the "watch" step unwatchable.
  useEffect(() => {
    const id = setInterval(() => setChain(tickDemoChain), DEMO_BLOCK_MS)
    return () => clearInterval(id)
  }, [])

  const epoch = target.epochId ? store.epochs[epochStoreKey(target.epochId)] : undefined
  const threshold = epoch?.policy?.threshold ?? 0
  const committeeSize = epoch?.policy?.committeeSize ?? epoch?.committee.length ?? 0
  const staggerBlocks = store.chain.staggerBlocks || 3
  const poolClaimed = epoch?.poolKeys.filter((slot) => slot.claimedBy != null).length ?? 0
  /** The key a registration claims: the fixture's cursor, exactly like `claimPoolKey`. */
  const nextPoolIndex = epoch?.poolNext ?? 0

  // Reads inside the async callbacks below must see the latest simulator
  // state without making every callback change identity each block.
  const latest = useRef(chain)
  latest.current = chain
  const nextPoolIndexRef = useRef(nextPoolIndex)
  nextPoolIndexRef.current = nextPoolIndex

  const register = useCallback(
    async ({ mode, skOrg }: RegisterArgs): Promise<RegisterResult> => {
      const poolIndex = nextPoolIndexRef.current
      if (poolIndex >= POOL_SIZE) throw new Error('PoolExhausted()')
      const { state, tx } = demoRegister(latest.current, target.aid ?? '0x', mode, poolIndex)
      setChain(state)
      organizerPKRef.current = mode === 'organizer-locked' && skOrg != null ? organizerPublicKey(skOrg) : null
      return { tx: { hash: tx.hash, block: tx.block, gasUsed: tx.gasUsed, simulated: true }, poolIndex }
    },
    [target.aid]
  )

  const applicationKeys = useCallback(async (): Promise<ApplicationKeys> => {
    const poolIndex = latest.current.poolIndex
    if (poolIndex == null || !target.epochId) throw new Error('The application is not registered yet')
    const poolKey = demoPoolKey(target.epochId, poolIndex)
    const organizerPK = organizerPKRef.current
    return { poolIndex, poolKey, organizerPK, key: applicationPublicKey(poolKey, organizerPK) }
  }, [target.epochId])

  const submitCiphertext = useCallback(
    async ({ aid, ciphertext }: { aid: Hex; ciphertext: ElGamalCiphertext }) => {
      const { state, tx, ciphertextIndex } = demoSubmitCiphertext(latest.current, aid)
      setChain(state)
      submittedRef.current = ciphertext
      setSubmitted(ciphertext)
      // Automatic: the committee holds the whole of sk_aid from the start, so
      // the plaintext is known as soon as the ciphertext is; it is only shown
      // once the simulated combine lands.
      if (state.mode === 'automatic' && state.poolIndex != null && target.epochId) {
        setPlaintext(demoRecoverPlaintext(ciphertext, demoPoolSecret(target.epochId, state.poolIndex)))
      }
      return {
        tx: { hash: tx.hash, block: tx.block, gasUsed: tx.gasUsed, simulated: true },
        ciphertextIndex,
      }
    },
    [target.epochId]
  )

  const revealSecret = useCallback(
    async ({ aid, skOrg }: RevealArgs): Promise<TxRecord> => {
      const { state, tx } = demoRevealSecret(latest.current, aid)
      setChain(state)
      // With both halves of `sk_aid` in hand the simulator can recover the
      // plaintext exactly as the committee's combine proof does.
      const ciphertext = submittedRef.current
      if (ciphertext && state.poolIndex != null && target.epochId) {
        const skAid = (demoPoolSecret(target.epochId, state.poolIndex) + skOrg) % subOrder
        setPlaintext(demoRecoverPlaintext(ciphertext, skAid))
      }
      return { hash: tx.hash, block: tx.block, gasUsed: tx.gasUsed, simulated: true }
    },
    [target.epochId]
  )

  const decryption = useCallback(
    (ciphertextIndex: number | null): DecryptionView | null => {
      if (ciphertextIndex == null || !chain.ciphertext) return null
      const params: DemoDecryptionParams = {
        epochId: target.epochId ?? '',
        threshold,
        committeeSize,
        staggerBlocks,
        ciphertextIndex,
      }
      const ctBlock = chain.ciphertext.tx.block
      // Automatic: unlocked from the start. Organizer-locked: from the reveal —
      // and, as on chain, no partial exists before it.
      const required = chain.mode === 'organizer-locked'
      const unlockedAt = required ? (chain.reveal?.block ?? null) : 0
      const openBlock = unlockedAt == null ? null : Math.max(ctBlock, unlockedAt)
      const partials = openBlock == null ? [] : demoPartialsAt(openBlock, chain.block, params)
      const combine = demoCombineAt(partials, unlockedAt, chain.block, params)
      return {
        threshold,
        committeeSize,
        staggerBlocks,
        ciphertextBlock: ctBlock,
        partials,
        reveal: {
          required,
          done: unlockedAt != null,
          block: chain.reveal?.block ?? null,
          tx: chain.reveal?.hash ?? null,
        },
        combined: {
          done: combine != null,
          block: combine?.block ?? null,
          tx: combine?.tx ?? null,
          plaintext: combine != null ? plaintext : null,
        },
        onChain: submitted,
      }
    },
    [chain, target.epochId, threshold, committeeSize, staggerBlocks, submitted, plaintext]
  )

  const registered = chain.register != null
  return useMemo<PlaygroundChain>(
    () => ({
      kind: 'demo',
      headBlock: chain.block,
      pool: epoch?.finalization ? { claimed: poolClaimed + (registered ? 1 : 0), size: POOL_SIZE } : null,
      wallet: {
        connected,
        address: connected ? DEMO_ACCOUNT : null,
        label: 'demo wallet',
        connect: () => setConnected(true),
        chainOk: true,
        problem: null,
      },
      register,
      applicationKeys,
      submitCiphertext,
      revealSecret,
      decryption,
    }),
    [
      chain.block,
      epoch?.finalization,
      poolClaimed,
      registered,
      connected,
      decryption,
      register,
      applicationKeys,
      revealSecret,
      submitCiphertext,
    ]
  )
}
