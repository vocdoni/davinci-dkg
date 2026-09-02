// The demo chain: the stepper's whole surface, backed by `demo-chain.ts` and
// the synthetic fixture. It imports no wagmi, opens no socket and signs
// nothing, so `/playground?demo=1` walks the complete organizer flow — right
// down to the partials arriving one block apart — on a machine with no chain
// in reach.
//
// Two deliberate substitutions, both surfaced in the UI:
//
//   • the wallet is a fixed address that "connects" instantly;
//   • `PK_ep` is a real curve point derived from the epoch id rather than the
//     fixture's decorative one, so the ciphertext, the DLEQ and the local
//     verification are genuine rather than arithmetic on a non-curve pair.
//
// Because the simulator knows that `sk_ep`, and the organizer hands it
// `sk_org` when it releases the share, it recovers the plaintext the way the
// committee would — `m·G = C2 − (sk_ep + sk_org)·C1`, then a BSGS dlog — so
// the closing "verify locally" step is a real comparison, not a tautology.

import { useCallback, useEffect, useMemo, useRef, useState } from 'react'
import { Base8, mulPointEscalar, subOrder } from '@zk-kit/baby-jubjub'
import { keccak256, stringToHex, type Hex } from 'viem'
import type { BabyJubPoint, ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import { useStore } from '~data/hooks'
import { resolveHeadBlock } from './head-block'
import { epochKey as epochStoreKey } from '~indexer/types'
import {
  DEMO_ACCOUNT,
  DEMO_BLOCK_MS,
  demoCombineAt,
  demoPartialsAt,
  demoRecoverPlaintext,
  demoRegister,
  demoReleaseShare,
  demoSubmitCiphertext,
  initialDemoChain,
  tickDemoChain,
  type DemoChainState,
  type DemoDecryptionParams,
} from './demo-chain'
import type { DecryptionView, PlaygroundChain, PlaygroundTarget, TxRecord } from './types'

/**
 * A deterministic, genuinely-on-curve stand-in for the epoch's collective
 * public key. The fixture's `PK_ep` is decorative bytes; the demo needs a real
 * point so the organizer's crypto is the real crypto.
 */
export function demoEpochSecret(epochId: string): bigint {
  const digest = keccak256(stringToHex(`davinci-dkg:demo-epoch-key:${epochId}`))
  return BigInt(digest) % subOrder || 1n
}

export function demoEpochKey(epochId: string): BabyJubPoint {
  return mulPointEscalar(Base8, demoEpochSecret(epochId)) as BabyJubPoint
}

export function useDemoChain(target: PlaygroundTarget): PlaygroundChain {
  const store = useStore()
  const [connected, setConnected] = useState(false)
  const [chain, setChain] = useState<DemoChainState>(() => initialDemoChain(resolveHeadBlock(store)))
  const [submitted, setSubmitted] = useState<ElGamalCiphertext | null>(null)
  const [plaintext, setPlaintext] = useState<bigint | null>(null)
  const submittedRef = useRef<ElGamalCiphertext | null>(null)

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

  // Reads inside the async callbacks below must see the latest simulator
  // state without making every callback change identity each block.
  const latest = useRef(chain)
  latest.current = chain

  const register = useCallback(async (): Promise<TxRecord> => {
    const { state, tx } = demoRegister(latest.current, target.aid ?? '0x')
    setChain(state)
    return { hash: tx.hash, block: tx.block, gasUsed: tx.gasUsed, simulated: true }
  }, [target.aid])

  const submitCiphertext = useCallback(
    async ({ aid, ciphertext }: { aid: Hex; ciphertext: ElGamalCiphertext }) => {
      const { state, tx, ciphertextIndex } = demoSubmitCiphertext(latest.current, aid)
      setChain(state)
      submittedRef.current = ciphertext
      setSubmitted(ciphertext)
      return {
        tx: { hash: tx.hash, block: tx.block, gasUsed: tx.gasUsed, simulated: true },
        ciphertextIndex,
      }
    },
    []
  )

  const releaseShare = useCallback(
    async ({ aid, ciphertextIndex, skOrg }: { aid: Hex; ciphertextIndex: number; skOrg: bigint }) => {
      const { state, tx } = demoReleaseShare(latest.current, aid, ciphertextIndex)
      setChain(state)
      // With both halves of `sk_aid` in hand the simulator can recover the
      // plaintext exactly as the committee's combine proof does.
      const ciphertext = submittedRef.current
      if (ciphertext && target.epochId) {
        const skAid = (demoEpochSecret(target.epochId) + skOrg) % subOrder
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
      const partials = demoPartialsAt(ctBlock, chain.block, params)
      const shareBlock = chain.share?.block ?? null
      const combine = demoCombineAt(partials, shareBlock, chain.block, params)
      return {
        threshold,
        committeeSize,
        staggerBlocks,
        ciphertextBlock: ctBlock,
        partials,
        share: { present: shareBlock != null, block: shareBlock, tx: chain.share?.hash ?? null },
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

  return useMemo<PlaygroundChain>(
    () => ({
      kind: 'demo',
      headBlock: chain.block,
      epochKey: target.epochId ? demoEpochKey(target.epochId) : null,
      wallet: {
        connected,
        address: connected ? DEMO_ACCOUNT : null,
        label: 'demo wallet',
        connect: () => setConnected(true),
        chainOk: true,
        problem: null,
      },
      register,
      submitCiphertext,
      releaseShare,
      decryption,
    }),
    [chain.block, connected, decryption, register, releaseShare, submitCiphertext, target.epochId]
  )
}
