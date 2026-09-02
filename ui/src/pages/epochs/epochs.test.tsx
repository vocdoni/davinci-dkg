// @vitest-environment jsdom
import { afterAll, beforeAll, describe, expect, it } from 'vitest'
import { screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { DataSourceProvider } from '~data/context'
import { createDemoDataSource } from '~fixtures/demo'
import { renderWithProviders } from '../../test-utils'
import { EpochsPage } from '../epochs'

// jsdom gives every element a zero rect, so the virtualiser would window down
// to nothing and the charts would only ever render skeletons. Pin a viewport
// for the suite — that is what makes the scale assertion meaningful.
const realRect = HTMLElement.prototype.getBoundingClientRect
const realObserver = globalThis.ResizeObserver

// The virtualiser sizes its window from a ResizeObserver callback, and falls
// back to `offsetHeight` — which jsdom pins at 0 — when the entry carries no
// border box. The global stub in vitest.setup never fires at all, so the table
// would mount zero rows. This one reports the pinned box immediately.
class ImmediateResizeObserver implements ResizeObserver {
  constructor(private readonly callback: ResizeObserverCallback) {}
  observe(target: Element) {
    const rect = target.getBoundingClientRect()
    const box = [{ inlineSize: rect.width, blockSize: rect.height }]
    this.callback(
      [{ target, contentRect: rect, borderBoxSize: box, contentBoxSize: box } as unknown as ResizeObserverEntry],
      this
    )
  }
  unobserve() {}
  disconnect() {}
}

beforeAll(() => {
  HTMLElement.prototype.getBoundingClientRect = function rect() {
    return { width: 1200, height: 620, top: 0, left: 0, right: 1200, bottom: 620, x: 0, y: 0, toJSON: () => ({}) }
  }
  globalThis.ResizeObserver = ImmediateResizeObserver
})
afterAll(() => {
  HTMLElement.prototype.getBoundingClientRect = realRect
  globalThis.ResizeObserver = realObserver
})

function renderEpochs(epochs: number) {
  // A source of our own: `renderWithProviders` supplies one, but the page needs
  // a fixture of a specific size to exercise the virtualisation threshold.
  const source = createDemoDataSource({
    operators: 24,
    epochs,
    committeeSize: 6,
    threshold: 3,
    minValidContributions: 4,
    applicationsPerEpoch: 1,
    ciphertextsPerApplication: 2,
    blockIntervalMs: 0,
  })
  const result = renderWithProviders(
    <DataSourceProvider source={source}>
      <EpochsPage />
    </DataSourceProvider>,
    { route: '/epochs' }
  )
  return { ...result, source }
}

describe('EpochsPage', () => {
  it('lists every epoch with its phase, policy and progress', () => {
    renderEpochs(4)
    expect(screen.getByRole('heading', { name: 'Epochs' })).toBeInTheDocument()
    expect(screen.getByText('4 epochs')).toBeInTheDocument()
    // t of n comes from the epoch policy, not from the committee length.
    expect(screen.getAllByText('3 / 6').length).toBeGreaterThan(0)
    // Two progress bars per row: claims and contributions.
    expect(screen.getAllByRole('progressbar').length).toBe(8)
  })

  it('filters by phase from the chips', async () => {
    const user = userEvent.setup()
    const { source } = renderEpochs(4)
    const store = source.getSnapshot().store
    const live = store.epochOrder.filter((id) => store.epochs[id].status === 'live').length

    await user.click(screen.getByRole('button', { name: /^live/ }))
    expect(screen.getByText(`${live} epoch${live === 1 ? '' : 's'}`)).toBeInTheDocument()
  })

  it('narrows to a single epoch from the search box', async () => {
    const user = userEvent.setup()
    const { source } = renderEpochs(4)
    const id = source.getSnapshot().store.epochOrder[1]

    await user.type(screen.getByLabelText('Search epochs'), id)
    expect(screen.getByText('1 epoch')).toBeInTheDocument()
  })

  it('stays windowed at 200 epochs', async () => {
    const { container } = renderEpochs(200)
    expect(screen.getByText('200 epochs')).toBeInTheDocument()
    // The plain table is not rendered at all above the threshold.
    expect(container.querySelectorAll('tbody tr')).toHaveLength(0)
    // The windowed body tags what it actually mounted; a 620 px viewport of
    // 44 px rows plus overscan is far short of 200. (It measures on an
    // animation frame, hence the wait.)
    await waitFor(() => expect(container.querySelectorAll('[data-index]').length).toBeGreaterThan(0))
    expect(container.querySelectorAll('[data-index]').length).toBeLessThan(80)
  })
})
