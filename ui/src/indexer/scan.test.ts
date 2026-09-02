import { describe, expect, it, vi } from 'vitest'
import type { PublicClient } from 'viem'
import { estimateRequests, isRangeError, scanRange, DEFAULT_CHUNK } from './scan'
import { ALL_EVENT_ABIS } from './events'
import type { Address } from './types'

const MANAGER = '0x3f9b338706a31f26d49159478015c8aaeab908ad' as Address

interface Call {
  from: number
  to: number
}

/** Minimal fake: records the ranges asked for and can reject wide ones. */
function fakeClient(maxSpan: number, calls: Call[]): PublicClient {
  return {
    async getLogs({ fromBlock, toBlock }: { fromBlock: bigint; toBlock: bigint }) {
      const from = Number(fromBlock)
      const to = Number(toBlock)
      if (to - from + 1 > maxSpan) {
        throw new Error('query returned more than 10000 results, block range too large')
      }
      calls.push({ from, to })
      return [
        {
          eventName: 'SlotClaimed',
          args: { epochId: '0x2f1105e90000000000000001', claimer: MANAGER, slot: 1 },
          blockNumber: BigInt(from),
          transactionHash: '0x' + 'aa'.repeat(32),
          logIndex: 0,
        },
      ]
    },
  } as unknown as PublicClient
}

describe('scanRange', () => {
  it('halves the chunk until the provider accepts it', async () => {
    const calls: Call[] = []
    const result = await scanRange({
      client: fakeClient(1_000, calls),
      addresses: [MANAGER],
      events: ALL_EVENT_ABIS,
      fromBlock: 1,
      toBlock: 2_000,
      chunkSize: 8_000,
      minChunk: 100,
    })
    expect(calls.every((call) => call.to - call.from + 1 <= 1_000)).toBe(true)
    expect(calls[0]).toEqual({ from: 1, to: 1_000 })
    expect(result.events).toHaveLength(calls.length)
    expect(result.chunkSize).toBeLessThanOrEqual(1_000)
  })

  it('grows the chunk again after a run of successes', async () => {
    const calls: Call[] = []
    const result = await scanRange({
      client: fakeClient(1_000_000, calls),
      addresses: [MANAGER],
      events: ALL_EVENT_ABIS,
      fromBlock: 0,
      toBlock: 10_000,
      chunkSize: 1_000,
      maxChunk: 50_000,
    })
    const spans = calls.map((call) => call.to - call.from + 1)
    expect(spans[0]).toBe(1_000)
    expect(Math.max(...spans)).toBeGreaterThan(1_000)
    expect(result.chunkSize).toBeGreaterThan(1_000)
  })

  it('covers the range exactly once and reports progress per chunk', async () => {
    const calls: Call[] = []
    const seen: number[] = []
    await scanRange({
      client: fakeClient(1_000_000, calls),
      addresses: [MANAGER],
      events: ALL_EVENT_ABIS,
      fromBlock: 100,
      toBlock: 349,
      chunkSize: 100,
      maxChunk: 100,
      onChunk: (chunk) => {
        seen.push(chunk.to)
      },
    })
    expect(calls).toEqual([
      { from: 100, to: 199 },
      { from: 200, to: 299 },
      { from: 300, to: 349 },
    ])
    expect(seen).toEqual([199, 299, 349])
  })

  it('propagates errors that are not about the range', async () => {
    const client = {
      getLogs: vi.fn().mockRejectedValue(new Error('connection refused')),
    } as unknown as PublicClient
    await expect(
      scanRange({
        client,
        addresses: [MANAGER],
        events: ALL_EVENT_ABIS,
        fromBlock: 0,
        toBlock: 10,
      }),
    ).rejects.toThrow('connection refused')
  })
})

describe('isRangeError', () => {
  it('recognises the usual provider phrasings', () => {
    for (const message of [
      'query returned more than 10000 results',
      'eth_getLogs is limited to 10000 blocks',
      'block range too large',
      'exceeds the maximum block range',
      'response size exceeded',
    ]) {
      expect(isRangeError(new Error(message))).toBe(true)
    }
    expect(isRangeError(new Error('nonce too low'))).toBe(false)
  })
})

describe('estimateRequests', () => {
  it('counts the chunks a backfill costs', () => {
    expect(estimateRequests(0)).toBe(0)
    expect(estimateRequests(10_000, 5_000)).toBe(2)
    expect(estimateRequests(10_001, 5_000)).toBe(3)
    expect(estimateRequests(216_000, DEFAULT_CHUNK)).toBe(44)
  })
})
