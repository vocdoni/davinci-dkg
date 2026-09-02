// @vitest-environment jsdom
import { afterAll, beforeAll, describe, expect, it } from 'vitest'
import { screen, within } from '@testing-library/react'
import { DataSourceProvider } from '~data/context'
import { createDemoDataSource } from '~fixtures/demo'
import { renderWithProviders } from '../../test-utils'
import { OverviewPage } from '../overview'

const realRect = HTMLElement.prototype.getBoundingClientRect
beforeAll(() => {
  HTMLElement.prototype.getBoundingClientRect = function rect() {
    return { width: 1200, height: 400, top: 0, left: 0, right: 1200, bottom: 400, x: 0, y: 0, toJSON: () => ({}) }
  }
})
afterAll(() => {
  HTMLElement.prototype.getBoundingClientRect = realRect
})

function renderOverview() {
  const source = createDemoDataSource({
    operators: 24,
    epochs: 4,
    committeeSize: 6,
    threshold: 3,
    minValidContributions: 4,
    applicationsPerEpoch: 1,
    ciphertextsPerApplication: 4,
    blockIntervalMs: 0,
  })
  const result = renderWithProviders(
    <DataSourceProvider source={source}>
      <OverviewPage />
    </DataSourceProvider>,
    { route: '/' }
  )
  return { ...result, source }
}

describe('OverviewPage', () => {
  it('names the deployment in the header strip', () => {
    const { source } = renderOverview()
    const chain = source.getSnapshot().store.chain
    expect(screen.getByText(chain.chainName)).toBeInTheDocument()
    expect(screen.getByText(`#${chain.chainId}`)).toBeInTheDocument()
    expect(screen.getByText('Next epoch')).toBeInTheDocument()
  })

  it('summarises the network in the status cards', () => {
    const { source } = renderOverview()
    const store = source.getSnapshot().store
    const live = store.epochOrder.filter((id) => store.epochs[id].status === 'live').length
    const newest = store.epochs[store.epochOrder[store.epochOrder.length - 1]]

    expect(screen.getByText('Newest epoch')).toBeInTheDocument()
    expect(screen.getByRole('link', { name: `#${newest.nonce}` })).toHaveAttribute('href', `/epochs/${newest.id}`)
    const aborted = store.epochOrder.filter((id) => store.epochs[id].status === 'aborted').length
    expect(screen.getByText(`of ${store.epochOrder.length} created · ${aborted} aborted`)).toBeInTheDocument()
    expect(screen.getByText('Live epochs').parentElement?.parentElement).toHaveTextContent(String(live))
    expect(screen.getByText('threshold t of n in force')).toBeInTheDocument()
  })

  it('draws the activity chart and the cadence strip once measured', () => {
    renderOverview()
    expect(screen.getByLabelText('Stacked activity chart')).toBeInTheDocument()
    expect(screen.getByLabelText('Epoch cadence')).toBeInTheDocument()
  })

  it('shows the twenty newest events with their block and transaction', () => {
    const { source, container } = renderOverview()
    const feed = source.getSnapshot().store.events.slice(-20)
    const panel = screen.getByText('Latest events').closest('div[class*="rounded-md"]') as HTMLElement

    expect(within(panel).getAllByText(feed[feed.length - 1].name).length).toBeGreaterThan(0)
    // 20 rows plus the header row.
    expect(container.querySelectorAll('tbody tr').length).toBe(20)
  })

  it('says nothing about the indexer when it is caught up', () => {
    renderOverview()
    expect(screen.queryByText('Indexing history')).toBeNull()
  })
})
