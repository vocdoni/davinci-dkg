import { describe, expect, it } from 'vitest'
import { Base8, mulPointEscalar, subOrder } from '@zk-kit/baby-jubjub'
import type { ElGamalCiphertext } from '@vocdoni/davinci-dkg-sdk'
import {
  DEMO_ACCOUNT,
  demoCombineAt,
  demoParticipant,
  demoPartialsAt,
  demoRecoverPlaintext,
  DEMO_DLOG_LIMIT,
  demoRegister,
  demoReleaseShare,
  demoSubmitCiphertext,
  demoTxHash,
  initialDemoChain,
  sendDemoTx,
  tickDemoChain,
  type DemoDecryptionParams,
} from './demo-chain'

const AID = `0x1c${'ab'.repeat(31)}`
const EPOCH = '0x0102030405060708090a0b0c'

const params = (over: Partial<DemoDecryptionParams> = {}): DemoDecryptionParams => ({
  epochId: EPOCH,
  threshold: 3,
  committeeSize: 6,
  staggerBlocks: 3,
  ciphertextIndex: 1,
  ...over,
})

describe('deterministic identities', () => {
  it('gives the same hash for the same transaction, every time', () => {
    expect(demoTxHash('registerApplication:x', 1)).toBe(demoTxHash('registerApplication:x', 1))
  })

  it('separates transactions by label and by sequence', () => {
    expect(demoTxHash('a', 1)).not.toBe(demoTxHash('b', 1))
    expect(demoTxHash('a', 1)).not.toBe(demoTxHash('a', 2))
  })

  it('produces well-formed addresses and hashes', () => {
    expect(DEMO_ACCOUNT).toMatch(/^0x[0-9a-f]{40}$/)
    expect(demoParticipant(EPOCH, 4)).toMatch(/^0x[0-9a-f]{40}$/)
    expect(demoTxHash('x', 1)).toMatch(/^0x[0-9a-f]{64}$/)
  })
})

describe('sending transactions', () => {
  it('mines one block per transaction and never reuses a hash', () => {
    const start = initialDemoChain(1_000)
    const first = sendDemoTx(start, 'a', 100)
    const second = sendDemoTx(first.state, 'b', 200)
    expect(first.tx.block).toBe(1_001)
    expect(second.tx.block).toBe(1_002)
    expect(second.state.seq).toBe(2)
    expect(first.tx.hash).not.toBe(second.tx.hash)
  })

  it('advances the clock without a transaction', () => {
    expect(tickDemoChain(initialDemoChain(10)).block).toBe(11)
  })

  it('assigns ciphertext indices the way the contract does', () => {
    const registered = demoRegister(initialDemoChain(1), AID)
    const first = demoSubmitCiphertext(registered.state, AID)
    const second = demoSubmitCiphertext(first.state, AID)
    expect(first.ciphertextIndex).toBe(1)
    expect(second.ciphertextIndex).toBe(2)
    // The first ciphertext of an application costs more (fresh storage).
    expect(first.tx.gasUsed).toBeGreaterThan(second.tx.gasUsed)
  })

  it('records the share on the state it returns', () => {
    const state = demoSubmitCiphertext(demoRegister(initialDemoChain(1), AID).state, AID).state
    const released = demoReleaseShare(state, AID, 1)
    expect(released.state.share).toEqual(released.tx)
  })
})

describe('partials', () => {
  it('produces nothing until a block after the ciphertext landed', () => {
    expect(demoPartialsAt(100, 100, params())).toEqual([])
    expect(demoPartialsAt(100, 101, params())).toHaveLength(1)
  })

  it('fills the threshold inside one stagger window, then adds late responders', () => {
    expect(demoPartialsAt(100, 102, params())).toHaveLength(2)
    // t = 3 partials are all in by ciphertextBlock + staggerBlocks…
    expect(demoPartialsAt(100, 103, params())).toHaveLength(3)
    // …and a couple of late members follow one window later.
    expect(demoPartialsAt(100, 500, params())).toHaveLength(5)
  })

  it('numbers waves the way the indexer does', () => {
    const partials = demoPartialsAt(100, 200, params({ threshold: 6, committeeSize: 8 }))
    expect(partials.map((p) => p.wave)).toEqual([0, 0, 0, 0, 1, 1, 1, 1])
    for (const partial of partials) {
      expect(partial.wave).toBe(Math.floor((partial.block - 100) / 3))
    }
  })

  it('answers in a rotation inside the committee, without repeats', () => {
    const partials = demoPartialsAt(100, 200, params({ threshold: 4, committeeSize: 6, ciphertextIndex: 2 }))
    const indices = partials.map((p) => p.participantIndex)
    expect(new Set(indices).size).toBe(indices.length)
    for (const index of indices) {
      expect(index).toBeGreaterThanOrEqual(1)
      expect(index).toBeLessThanOrEqual(6)
    }
    // The rotation start is seeded by the ciphertext index.
    expect(demoPartialsAt(100, 200, params({ ciphertextIndex: 2 }))[0].participantIndex).not.toBe(
      demoPartialsAt(100, 200, params({ ciphertextIndex: 1 }))[0].participantIndex
    )
  })

  it('wraps the rotation past the end of the committee', () => {
    const partials = demoPartialsAt(100, 200, params({ threshold: 3, committeeSize: 4, ciphertextIndex: 3 }))
    expect(partials.map((p) => p.participantIndex)).toEqual([4, 1, 2, 3])
  })

  it('orders partials by block', () => {
    const partials = demoPartialsAt(100, 200, params({ threshold: 6, committeeSize: 8 }))
    const blocks = partials.map((p) => p.block)
    expect([...blocks].sort((a, b) => a - b)).toEqual(blocks)
  })

  it('has nothing to say about an epoch with no committee', () => {
    expect(demoPartialsAt(100, 200, params({ threshold: 0, committeeSize: 0 }))).toEqual([])
  })
})

describe('combine', () => {
  const full = () => demoPartialsAt(100, 200, params())

  it('never lands while the organizer withholds the share', () => {
    expect(demoCombineAt(full(), null, 500, params())).toBeNull()
  })

  it('never lands below the threshold, share or no share', () => {
    expect(demoCombineAt(demoPartialsAt(100, 101, params()), 102, 500, params())).toBeNull()
  })

  it('lands one block after the later of the two halves', () => {
    // The t-th partial lands at block 103; a share released later is the blocker.
    expect(demoCombineAt(full(), 110, 500, params())?.block).toBe(111)
    // …and vice versa. Late responders past t do not hold it up.
    expect(demoCombineAt(full(), 101, 500, params())?.block).toBe(104)
  })

  it('is invisible until the head reaches it', () => {
    expect(demoCombineAt(full(), 101, 103, params())).toBeNull()
    expect(demoCombineAt(full(), 101, 104, params())?.gasUsed).toBeGreaterThan(0)
  })
})

describe('plaintext recovery', () => {
  /** The simulator is the only party that ever holds `sk_aid` in one piece. */
  async function roundTrip(message: bigint, skAid: bigint): Promise<bigint | null> {
    const { buildElGamal } = await import('@vocdoni/davinci-dkg-sdk')
    const elgamal = await buildElGamal()
    const pk = mulPointEscalar(Base8, skAid) as ElGamalCiphertext['c1']
    return demoRecoverPlaintext(elgamal.encrypt(message, pk), skAid)
  }

  it('recovers what was encrypted under sk_aid', async () => {
    expect(await roundTrip(4_242n, 987_654_321n % subOrder)).toBe(4_242n)
  })

  it('recovers zero', async () => {
    expect(await roundTrip(0n, 7n)).toBe(0n)
  })

  it('caps the walk rather than searching forever', () => {
    expect(DEMO_DLOG_LIMIT).toBe(1n << 20n)
  })
})
