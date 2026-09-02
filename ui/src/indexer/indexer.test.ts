import { describe, expect, it } from 'vitest'
import type { PublicClient } from 'viem'
import { Indexer } from './indexer'
import { memoryStore } from './persist'
import type { Address, EpochId, Hex } from './types'

const MANAGER = '0x3f9b338706a31f26d49159478015c8aaeab908ad' as Address
const REGISTRY = '0x9a1f2ce4bd0e0b1f5f6d0a4c7e2b3d8f1c5a6e70' as Address
const APP_MANAGER = '0x5c3e7a9d1b8f0426ea5d9c3b7f102a4d6e8b9c11' as Address
const EPOCH = '0x2f1105e90000000000000001' as EpochId
const ALICE = '0x1111111111111111111111111111111111111111' as Address
const BOB = '0x2222222222222222222222222222222222222222' as Address
const FINALIZE_TX = ('0x' + 'fa'.repeat(32)) as Hex

interface FakeLog {
  eventName: string
  args: Record<string, unknown>
  blockNumber: bigint
  transactionHash: Hex
  logIndex: number
}

function log(name: string, block: number, index: number, args: Record<string, unknown>, tx?: Hex): FakeLog {
  return {
    eventName: name,
    args,
    blockNumber: BigInt(block),
    transactionHash: tx ?? (`0x${block.toString(16).padStart(64, '0')}` as Hex),
    logIndex: index,
  }
}

/** A viem `PublicClient` reduced to what the indexer actually calls. */
function fakeChain() {
  const state = {
    head: 250,
    logs: [
      log('NodeRegistered', 101, 0, { operator: ALICE, pubX: 3n, pubY: 4n }),
      log('NodeRegistered', 102, 0, { operator: BOB, pubX: 5n, pubY: 6n }),
      log('EpochCreated', 120, 0, {
        epochId: EPOCH,
        organizer: ALICE,
        startBlock: 120n,
        seedBlock: 121n,
        lotteryThreshold: 1n << 250n,
      }),
      log('SlotClaimed', 121, 0, { epochId: EPOCH, claimer: ALICE, slot: 0 }),
      log('SlotClaimed', 121, 1, { epochId: EPOCH, claimer: BOB, slot: 1 }),
      log('CommitteeFilled', 121, 2, { epochId: EPOCH }),
      log('ContributionSubmitted', 130, 0, {
        epochId: EPOCH,
        contributor: ALICE,
        contributorIndex: 0,
        commitmentsHash: ('0x' + '11'.repeat(32)) as Hex,
        encryptedSharesHash: ('0x' + '22'.repeat(32)) as Hex,
      }),
      log(
        'EpochLive',
        140,
        0,
        {
          epochId: EPOCH,
          aggregateCommitmentsHash: ('0x' + '33'.repeat(32)) as Hex,
          collectivePublicKeyHash: ('0x' + '44'.repeat(32)) as Hex,
          shareCommitmentHash: ('0x' + '55'.repeat(32)) as Hex,
        },
        FINALIZE_TX,
      ),
    ] as FakeLog[],
    ranges: [] as Array<[number, number]>,
    multicalls: 0,
    receipts: 0,
  }

  const answer = (call: { functionName: string; args?: unknown[] }): unknown => {
    switch (call.functionName) {
      case 'REGISTRY':
        return REGISTRY
      case 'appManager':
        return APP_MANAGER
      case 'EPOCH_PREFIX':
        return 0x2f1105e9
      case 'epochDurationBlocks':
        return 300n
      case 'nextEpochStartBlock':
        return 420n
      case 'nodeCount':
        return 2n
      case 'activeCount':
        return 2n
      case 'INACTIVITY_WINDOW':
        return 50_400n
      case 'getEpoch':
        return {
          organizer: ALICE,
          policy: {
            threshold: 2,
            committeeSize: 2,
            minValidContributions: 2,
            lotteryAlphaBps: 15_000,
            committeeSelectionDeadlineBlock: 145n,
            keyAssemblyDeadlineBlock: 170n,
            liveNotBeforeBlock: 175n,
          },
          status: 3,
          nonce: 1n,
          startBlock: 120n,
          seedBlock: 121n,
          seed: ('0x' + 'ab'.repeat(32)) as Hex,
          lotteryThreshold: 1n << 250n,
          claimedCount: 2,
          contributionCount: 1,
          partialDecryptionCount: 0,
          ciphertextCount: 0,
        }
      case 'selectedParticipants':
        return [ALICE, BOB]
      case 'getCollectivePublicKey':
        return { x: 12n, y: 34n }
      case 'getShareCommitmentHash':
        return `0x${String(call.args?.[1] ?? 0).padStart(64, '0')}` as Hex
      case 'getNode':
        return {
          operator: call.args?.[0],
          pubX: 3n,
          pubY: 4n,
          status: 1,
          lastActiveBlock: 130n,
          registeredAtBlock: 101n,
        }
      default:
        throw new Error(`unexpected call ${call.functionName}`)
    }
  }

  const client = {
    async getBlockNumber() {
      return BigInt(state.head)
    },
    async getLogs({ fromBlock, toBlock }: { fromBlock: bigint; toBlock: bigint }) {
      const from = Number(fromBlock)
      const to = Number(toBlock)
      state.ranges.push([from, to])
      return state.logs.filter((entry) => {
        const block = Number(entry.blockNumber)
        return block >= from && block <= to
      })
    },
    async multicall({ contracts }: { contracts: Array<{ functionName: string; args?: unknown[] }> }) {
      state.multicalls += 1
      return contracts.map((call) => {
        try {
          return { status: 'success' as const, result: answer(call) }
        } catch (error) {
          return { status: 'failure' as const, error }
        }
      })
    },
    async getTransactionReceipt({ hash }: { hash: Hex }) {
      state.receipts += 1
      return { from: BOB, gasUsed: 1_112_337n, blockNumber: 140n, status: 'success', transactionHash: hash }
    },
  } as unknown as PublicClient

  return { client, state }
}

function indexerFor(client: PublicClient, kv = memoryStore()) {
  return {
    kv,
    indexer: new Indexer({
      client,
      chainId: 11155111,
      managerAddress: MANAGER,
      deployBlock: 100,
      chunkSize: 100,
      kv,
      pollIntervalMs: 60_000,
    }),
  }
}

describe('Indexer', () => {
  it('backfills, joins contract state and attributes the finalizer', async () => {
    const { client, state } = fakeChain()
    const { indexer, kv } = indexerFor(client)
    await indexer.refresh()

    const { store, status } = indexer.getSnapshot()
    expect(status.phase).toBe('live')
    expect(status.scanning).toBe(false)
    expect(status.progress).toBe(1)
    expect(status.lastBlock).toBe(250)
    expect(status.eventCount).toBe(8)

    // Chunked: 151 blocks at a chunk of 100 (the floor scanRange enforces).
    expect(state.ranges[0]).toEqual([100, 199])
    expect(state.ranges[state.ranges.length - 1][1]).toBe(250)
    expect(state.ranges.length).toBeGreaterThan(1)

    // Sibling addresses came from the manager, not from config.
    expect(store.chain.registryAddress).toBe(REGISTRY)
    expect(store.chain.appManagerAddress).toBe(APP_MANAGER)
    expect(store.chain.epochDurationBlocks).toBe(300)
    expect(store.chain.activeCount).toBe(2)

    const epoch = store.epochs[EPOCH]
    expect(epoch.status).toBe('live')
    expect(epoch.policy?.threshold).toBe(2)
    expect(epoch.committee).toEqual([ALICE, BOB])
    expect(epoch.collectivePublicKey).not.toBeNull()
    expect(epoch.shareCommitmentHashes).toHaveLength(2)
    expect(epoch.finalization?.by).toBe(BOB)
    expect(store.txMeta[FINALIZE_TX].gasUsed).toBe(1_112_337)
    expect(store.operators[ALICE].registeredAtBlock).toBe(101)

    // The store was cached.
    const cached = await kv.get(`dkg-explorer:v${store.version}:11155111:${MANAGER}`)
    expect(cached).toBeTruthy()
  })

  it('resumes from the cache and only scans the new range', async () => {
    const { client, state } = fakeChain()
    const kv = memoryStore()
    const first = indexerFor(client, kv)
    await first.indexer.refresh()
    const eventsAfterFirst = first.indexer.getSnapshot().store.events.length

    state.ranges.length = 0
    state.head = 300
    state.logs.push(
      log('CiphertextSubmitted', 260, 0, {
        epochId: EPOCH,
        aid: ('0x0f' + '00'.repeat(31)) as Hex,
        ciphertextIndex: 1,
        submitter: BOB,
        c1x: 1n,
        c1y: 2n,
        c2x: 3n,
        c2y: 4n,
      }),
    )

    const second = indexerFor(client, kv)
    await second.indexer.refresh()

    // Nothing before block 251 is asked for a second time.
    expect(Math.min(...state.ranges.map(([from]) => from))).toBe(251)
    const { store } = second.indexer.getSnapshot()
    expect(store.events.length).toBe(eventsAfterFirst + 1)
    expect(store.ciphertexts[`${EPOCH}:0x0f${'00'.repeat(31)}:1`]).toBeTruthy()
  })

  it('notifies subscribers with a fresh snapshot identity', async () => {
    const { client } = fakeChain()
    const { indexer } = indexerFor(client)
    let notifications = 0
    const unsubscribe = indexer.subscribe(() => {
      notifications += 1
    })
    const before = indexer.getSnapshot()
    await indexer.refresh()
    const after = indexer.getSnapshot()
    expect(notifications).toBeGreaterThan(0)
    expect(after).not.toBe(before)
    expect(after.store).not.toBe(before.store)
    unsubscribe()
    const count = notifications
    await indexer.refresh()
    expect(notifications).toBe(count)
  })

  it('rewrites the cache only when something changed', async () => {
    const { client } = fakeChain()
    const kv = memoryStore()
    let writes = 0
    const counting = { ...kv, set: async (key: string, value: unknown) => { writes += 1; await kv.set(key, value) } }
    const indexer = new Indexer({
      client,
      chainId: 11155111,
      managerAddress: MANAGER,
      deployBlock: 100,
      chunkSize: 100,
      kv: counting,
      pollIntervalMs: 60_000,
    })
    await indexer.refresh()
    expect(writes).toBeGreaterThan(0)
    const after = writes
    await indexer.refresh()
    expect(writes).toBe(after)
  })

  it('costs one request per poll while the chain is idle', async () => {
    const { client, state } = fakeChain()
    const { indexer } = indexerFor(client)
    await indexer.refresh()
    state.ranges.length = 0
    state.multicalls = 0
    state.receipts = 0
    // Head has not moved and nothing is stale: `eth_blockNumber` and nothing
    // else — no getLogs, no multicall, no receipts.
    await indexer.refresh()
    expect(state.ranges).toHaveLength(0)
    expect(state.multicalls).toBe(0)
    expect(state.receipts).toBe(0)
  })

  it('falls back to sequential reads when multicall is unavailable', async () => {
    const { client, state } = fakeChain()
    let sequential = 0
    const noMulticall = {
      ...client,
      getBlockNumber: client.getBlockNumber,
      getLogs: client.getLogs,
      getTransactionReceipt: client.getTransactionReceipt,
      async multicall() {
        throw new Error('chain does not support contract "multicall3"')
      },
      async readContract(call: { functionName: string; args?: unknown[] }) {
        sequential += 1
        switch (call.functionName) {
          case 'REGISTRY':
            return REGISTRY
          case 'appManager':
            return APP_MANAGER
          case 'nodeCount':
          case 'activeCount':
            return 2n
          case 'EPOCH_PREFIX':
            return 0x2f1105e9
          case 'epochDurationBlocks':
            return 300n
          case 'nextEpochStartBlock':
            return 420n
          case 'INACTIVITY_WINDOW':
            return 50_400n
          case 'selectedParticipants':
            return [ALICE, BOB]
          case 'getCollectivePublicKey':
            return { x: 0n, y: 1n }
          case 'getShareCommitmentHash':
            return ('0x' + '00'.repeat(32)) as Hex
          case 'getNode':
            return { pubX: 1n, pubY: 2n, status: 1, lastActiveBlock: 130n, registeredAtBlock: 101n }
          case 'getEpoch':
            return {
              policy: {
                threshold: 2,
                committeeSize: 2,
                minValidContributions: 2,
                lotteryAlphaBps: 15_000,
                committeeSelectionDeadlineBlock: 145n,
                keyAssemblyDeadlineBlock: 170n,
                liveNotBeforeBlock: 175n,
              },
              status: 3,
              claimedCount: 2,
              contributionCount: 1,
              partialDecryptionCount: 0,
              ciphertextCount: 0,
            }
          default:
            throw new Error(`unexpected ${call.functionName}`)
        }
      },
    } as unknown as PublicClient

    const { indexer } = indexerFor(noMulticall)
    await indexer.refresh()
    expect(sequential).toBeGreaterThan(0)
    expect(indexer.getSnapshot().store.epochs[EPOCH].policy?.committeeSize).toBe(2)
    expect(indexer.getSnapshot().status.errors).toHaveLength(0)
    expect(state.head).toBe(250)
  })

  it('records errors without losing what it already indexed', async () => {
    const { client, state } = fakeChain()
    const { indexer } = indexerFor(client)
    await indexer.refresh()
    const events = indexer.getSnapshot().store.events.length

    state.head = 400
    const broken = indexer as unknown as { config: { client: PublicClient } }
    broken.config.client = {
      ...client,
      async getBlockNumber() {
        return 400n
      },
      async getLogs() {
        throw new Error('connection refused')
      },
    } as unknown as PublicClient

    await indexer.refresh()
    const snapshot = indexer.getSnapshot()
    expect(snapshot.status.phase).toBe('error')
    expect(snapshot.status.errors.at(-1)?.message).toContain('connection refused')
    expect(snapshot.store.events).toHaveLength(events)
  })
})
