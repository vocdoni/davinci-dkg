// @vitest-environment jsdom
import { describe, expect, it, vi } from 'vitest'
import { act, renderHook } from '@testing-library/react'
import type { ReactNode } from 'react'
import { DataSourceProvider } from './context'
import { createDemoDataSource } from '../fixtures/demo'
import {
  useActivity,
  useApplication,
  useApplications,
  useEpoch,
  useEpochs,
  useEventFeed,
  useIndexer,
  useNetworkStats,
  useOperator,
  useOperators,
  useIndexerSearchResolver,
  usePartialMatrix,
  useSearch,
} from './hooks'
import type { DataSource } from './source'

const FIXTURE = {
  operators: 24,
  epochs: 4,
  committeeSize: 6,
  threshold: 3,
  minValidContributions: 4,
  applicationsPerEpoch: 1,
  ciphertextsPerApplication: 4,
  blockIntervalMs: 0,
}

function harness(): { source: DataSource; wrapper: (props: { children: ReactNode }) => JSX.Element } {
  const source = createDemoDataSource(FIXTURE)
  const wrapper = ({ children }: { children: ReactNode }): JSX.Element => (
    <DataSourceProvider source={source}>{children}</DataSourceProvider>
  )
  return { source, wrapper }
}

describe('data hooks', () => {
  it('exposes indexer status from the demo source', () => {
    const { wrapper } = harness()
    const { result } = renderHook(() => useIndexer(), { wrapper })
    expect(result.current.kind).toBe('demo')
    expect(result.current.status.phase).toBe('live')
    expect(result.current.scanning).toBe(false)
    expect(result.current.progress).toBe(1)
    expect(result.current.headBlock).toBeGreaterThan(0)
  })

  it('re-renders when the source publishes a new block', async () => {
    const { source, wrapper } = harness()
    const { result } = renderHook(() => useIndexer(), { wrapper })
    const before = result.current.headBlock
    await act(async () => {
      await source.refresh()
    })
    expect(result.current.headBlock).toBe(before + 1)
  })

  it('serves network stats, epochs and operators', () => {
    const { wrapper } = harness()
    const stats = renderHook(() => useNetworkStats(), { wrapper })
    expect(stats.result.current.operatorsRegistered).toBe(24)
    expect(stats.result.current.epochs).toBe(4)

    const epochs = renderHook(() => useEpochs(), { wrapper })
    expect(epochs.result.current).toHaveLength(4)
    expect(epochs.result.current[0].nonce).toBe(4)

    const live = renderHook(() => useEpochs({ phase: 'live' }), { wrapper })
    expect(live.result.current.every((row) => row.phase === 'live')).toBe(true)

    const operators = renderHook(() => useOperators(), { wrapper })
    expect(operators.result.current).toHaveLength(24)
  })

  it('serves one epoch, one operator and one application in detail', () => {
    const { source, wrapper } = harness()
    const store = source.getSnapshot().store
    const liveId = store.epochOrder.find((id) => store.epochs[id].status === 'live')!

    const epoch = renderHook(() => useEpoch(liveId), { wrapper })
    expect(epoch.result.current?.committee).toHaveLength(6)
    expect(epoch.result.current?.lottery.claims).toHaveLength(6)

    const address = store.epochs[liveId].committee[0]
    const operator = renderHook(() => useOperator(address), { wrapper })
    expect(operator.result.current?.row.address).toBe(address)
    expect(operator.result.current?.history.length).toBeGreaterThan(0)

    const apps = renderHook(() => useApplications(), { wrapper })
    expect(apps.result.current.length).toBeGreaterThan(0)
    const first = apps.result.current[0]
    const application = renderHook(() => useApplication(first.epoch, first.aid), { wrapper })
    expect(application.result.current?.ciphertexts).toHaveLength(4)

    const matrix = renderHook(() => usePartialMatrix(first.epoch, first.aid), { wrapper })
    expect(matrix.result.current?.rows).toHaveLength(6)
    expect(matrix.result.current?.columns).toHaveLength(4)
  })

  it('serves activity, the event feed and search', () => {
    const { source, wrapper } = harness()
    const activity = renderHook(() => useActivity(10), { wrapper })
    expect(activity.result.current).toHaveLength(4)

    const feed = renderHook(() => useEventFeed(5), { wrapper })
    expect(feed.result.current).toHaveLength(5)

    const store = source.getSnapshot().store
    const epochId = store.epochOrder[0]
    const search = renderHook(() => useSearch(epochId), { wrapper })
    expect(search.result.current[0]).toMatchObject({ kind: 'epoch', id: epochId })

    const empty = renderHook(() => useSearch(''), { wrapper })
    expect(empty.result.current).toHaveLength(0)
  })

  it('resolves the shell search box against the store', () => {
    const { source, wrapper } = harness()
    const store = source.getSnapshot().store
    const { result } = renderHook(() => useIndexerSearchResolver(), { wrapper })
    const resolver = result.current

    expect(resolver(store.epochOrder[0])).toEqual({
      kind: 'route',
      path: `/epochs/${store.epochOrder[0]}`,
      label: expect.stringContaining('Epoch'),
    })

    const tx = store.events.find((event) => event.tx)!.tx as string
    expect(resolver(tx)).toMatchObject({ kind: 'external', url: expect.stringContaining('/tx/') })

    expect(resolver('definitely-not-a-thing')).toBeNull()
  })

  it('returns null for unknown ids instead of throwing', () => {
    const { wrapper } = harness()
    const epoch = renderHook(() => useEpoch('0x2f1105e9000000000000ffff'), { wrapper })
    expect(epoch.result.current).toBeNull()
    const operator = renderHook(() => useOperator(undefined), { wrapper })
    expect(operator.result.current).toBeNull()
  })

  it('throws a useful error outside the provider', () => {
    // React logs the thrown render error; the assertion is the point, not the
    // stack trace it dumps into the test output.
    const consoleError = vi.spyOn(console, 'error').mockImplementation(() => {})
    try {
      expect(() => renderHook(() => useNetworkStats())).toThrow(/DataSourceProvider/)
    } finally {
      consoleError.mockRestore()
    }
  })
})
