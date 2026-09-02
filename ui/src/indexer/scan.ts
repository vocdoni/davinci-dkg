// Chunked, adaptive `eth_getLogs` scanning.
//
// One filter covers all three contracts (their event names — and therefore
// topic0s — are disjoint), so a chunk costs one request no matter how many
// event types it returns. The chunk size halves on a "range too large" /
// "too many results" rejection and grows back on a run of successes, which is
// what lets the same code run against a generous private RPC and a 10k-block
// public one without configuration.

import type { AbiEvent, Address, PublicClient } from 'viem'
import { normalizeLog, type RawLog } from './events'
import type { IndexedEvent } from './types'

export const DEFAULT_CHUNK = 5_000
export const MIN_CHUNK = 100
export const MAX_CHUNK = 50_000
/** Successful chunks in a row before the size grows again. */
const GROW_AFTER = 4

export interface ChunkReport {
  from: number
  to: number
  events: IndexedEvent[]
  /** Requests issued so far in this scan. */
  requests: number
  /** Chunk size in force after this chunk. */
  chunkSize: number
}

export interface ScanOptions {
  client: PublicClient
  addresses: Address[]
  events: AbiEvent[]
  fromBlock: number
  toBlock: number
  chunkSize?: number
  minChunk?: number
  maxChunk?: number
  /** Called after every chunk so the UI can render while the scan runs. */
  onChunk?: (report: ChunkReport) => void | Promise<void>
  signal?: AbortSignal
}

export interface ScanResult {
  events: IndexedEvent[]
  requests: number
  /** Chunk size the scan settled on; feed it back into the next scan. */
  chunkSize: number
}

/** Provider rejections that mean "ask for a smaller range". */
export function isRangeError(err: unknown): boolean {
  const message = (err instanceof Error ? err.message : String(err)).toLowerCase()
  return (
    message.includes('range') ||
    message.includes('exceed') ||
    message.includes('too large') ||
    message.includes('too many') ||
    message.includes('limited') ||
    message.includes('more than') ||
    message.includes('query returned') ||
    message.includes('response size') ||
    message.includes('block limit') ||
    message.includes('-32005')
  )
}

/**
 * Scan `[fromBlock, toBlock]` for every event of `events` emitted by any of
 * `addresses`, normalised and in chronological order.
 */
export async function scanRange(options: ScanOptions): Promise<ScanResult> {
  const {
    client,
    addresses,
    events,
    fromBlock,
    toBlock,
    minChunk = MIN_CHUNK,
    maxChunk = MAX_CHUNK,
    onChunk,
    signal,
  } = options

  const all: IndexedEvent[] = []
  let chunkSize = Math.max(minChunk, Math.min(maxChunk, options.chunkSize ?? DEFAULT_CHUNK))
  let requests = 0
  let streak = 0
  let cursor = fromBlock

  while (cursor <= toBlock) {
    if (signal?.aborted) break
    const end = Math.min(cursor + chunkSize - 1, toBlock)
    let logs: RawLog[]
    try {
      requests += 1
      logs = (await client.getLogs({
        address: addresses.length === 1 ? addresses[0] : addresses,
        events,
        fromBlock: BigInt(cursor),
        toBlock: BigInt(end),
        // A log we cannot decode strictly (an event fragment that changed
        // between builds) must not abort the whole chunk.
        strict: false,
        // eslint-disable-next-line @typescript-eslint/no-explicit-any
      } as any)) as unknown as RawLog[]
    } catch (err) {
      if (isRangeError(err) && chunkSize > minChunk) {
        chunkSize = Math.max(minChunk, Math.floor(chunkSize / 2))
        streak = 0
        continue
      }
      throw err
    }

    const decoded: IndexedEvent[] = []
    for (const log of logs) {
      const event = normalizeLog(log)
      if (event) decoded.push(event)
    }
    decoded.sort((a, b) => a.block - b.block || a.logIndex - b.logIndex)
    all.push(...decoded)

    if (onChunk) await onChunk({ from: cursor, to: end, events: decoded, requests, chunkSize })

    cursor = end + 1
    streak += 1
    if (streak >= GROW_AFTER && chunkSize < maxChunk) {
      chunkSize = Math.min(maxChunk, Math.floor(chunkSize * 1.5))
      streak = 0
    }
  }

  return { events: all, requests, chunkSize }
}

/** Number of requests a scan of `span` blocks costs at a given chunk size. */
export function estimateRequests(span: number, chunkSize = DEFAULT_CHUNK): number {
  return Math.max(0, Math.ceil(span / chunkSize))
}
