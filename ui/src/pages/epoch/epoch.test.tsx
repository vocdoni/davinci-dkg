// @vitest-environment jsdom
import { afterAll, beforeAll, describe, expect, it } from 'vitest'
import { screen, within } from '@testing-library/react'
import { Route, Routes } from 'react-router-dom'
import { DataSourceProvider } from '~data/context'
import { createDemoDataSource } from '~fixtures/demo'
import type { DataSource } from '~data/source'
import { renderWithProviders } from '../../test-utils'
import { EpochPage } from '../epoch'

const realRect = HTMLElement.prototype.getBoundingClientRect
beforeAll(() => {
  HTMLElement.prototype.getBoundingClientRect = function rect() {
    return { width: 1100, height: 520, top: 0, left: 0, right: 1100, bottom: 520, x: 0, y: 0, toJSON: () => ({}) }
  }
})
afterAll(() => {
  HTMLElement.prototype.getBoundingClientRect = realRect
})

function makeSource(): DataSource {
  return createDemoDataSource({
    operators: 24,
    epochs: 4,
    committeeSize: 6,
    threshold: 3,
    minValidContributions: 4,
    applicationsPerEpoch: 1,
    ciphertextsPerApplication: 4,
    blockIntervalMs: 0,
  })
}

function renderEpoch(source: DataSource, id: string) {
  return renderWithProviders(
    <DataSourceProvider source={source}>
      <Routes>
        <Route path='/epochs/:id' element={<EpochPage />} />
      </Routes>
    </DataSourceProvider>,
    { route: `/epochs/${id}` }
  )
}

function liveEpoch(source: DataSource) {
  const store = source.getSnapshot().store
  const id = store.epochOrder.find((key) => store.epochs[key].status === 'live')
  if (!id) throw new Error('the fixture produced no live epoch')
  return store.epochs[id]
}

describe('EpochPage', () => {
  it('heads the page with the epoch, its phase and its policy', () => {
    const source = makeSource()
    const epoch = liveEpoch(source)
    renderEpoch(source, epoch.id)

    expect(screen.getByRole('heading', { name: `Epoch #${epoch.nonce}` })).toBeInTheDocument()
    expect(screen.getAllByText('live').length).toBeGreaterThan(0)
    // Shown twice on purpose: in the epoch record and on the committee panel.
    expect(screen.getAllByText('t = 3 · m_min = 4 · n = 6').length).toBeGreaterThan(0)
    expect(screen.getByLabelText('Epoch lifecycle on the block axis')).toBeInTheDocument()
  })

  it('numbers committee slots from 0 and participant indexes from 1', () => {
    const source = makeSource()
    const epoch = liveEpoch(source)
    renderEpoch(source, epoch.id)

    const table = screen.getByText('D_i').closest('table') as HTMLTableElement
    const rows = within(table).getAllByRole('row').slice(1)
    expect(rows).toHaveLength(6)

    const first = within(rows[0]).getAllByRole('cell')
    expect(first[0]).toHaveTextContent('0')
    expect(first[1]).toHaveTextContent('1')
    const last = within(rows[5]).getAllByRole('cell')
    expect(last[0]).toHaveTextContent('5')
    expect(last[1]).toHaveTextContent('6')
  })

  it('shows the lottery inputs anyone would need to replay it', () => {
    const source = makeSource()
    const epoch = liveEpoch(source)
    renderEpoch(source, epoch.id)

    expect(screen.getByText('Committee selection')).toBeInTheDocument()
    expect(screen.getByText('seed block')).toBeInTheDocument()
    expect(screen.getByText('τ')).toBeInTheDocument()
    expect(screen.getByText('N snapshotted')).toBeInTheDocument()
    expect(screen.getByText('Claims in slot order')).toBeInTheDocument()
    // The gauge draws τ as a share of the hash space.
    expect(screen.getByLabelText('τ / 2²⁵⁶')).toBeInTheDocument()
  })

  it('shows the collective key, the finalizer and the transcript size', () => {
    const source = makeSource()
    const epoch = liveEpoch(source)
    renderEpoch(source, epoch.id)

    expect(screen.getByText('PK_ep.x')).toBeInTheDocument()
    expect(screen.getByText('finalizer')).toBeInTheDocument()
    // 2·n² + 5·n at n = 6.
    expect(screen.getByText('102 words')).toBeInTheDocument()
  })

  it('plots the decryption matrix with the share and combined rows', () => {
    const source = makeSource()
    const epoch = liveEpoch(source)
    renderEpoch(source, epoch.id)

    expect(screen.getByLabelText('Partial decryption matrix')).toBeInTheDocument()
    // Once as a matrix row label, once in the legend.
    expect(screen.getAllByText('organizer share')).toHaveLength(2)
    expect(screen.getAllByText('combined')).toHaveLength(2)
  })

  it('lists the applications and the full event log', () => {
    const source = makeSource()
    const epoch = liveEpoch(source)
    renderEpoch(source, epoch.id)

    expect(screen.getByText('1 registered')).toBeInTheDocument()
    expect(screen.getByText(`${epoch.events.length} events`)).toBeInTheDocument()
    expect(screen.getByText('Epoch record')).toBeInTheDocument()
  })

  it('calls out an epoch id the manager never created', () => {
    const source = makeSource()
    renderEpoch(source, '0x2f1105e9000000000000ffff')

    expect(screen.getByRole('heading', { name: 'Unknown epoch' })).toBeInTheDocument()
    expect(screen.getByText('0x2f1105e9000000000000ffff')).toBeInTheDocument()
  })
})
