import { afterAll, beforeAll, describe, expect, it } from 'vitest'
import { render, screen } from '@testing-library/react'
import { BlockTimeline, CadenceStrip, Donut, Gauge, Matrix, Sparkline, StackedBars, waveColor } from './index'

// jsdom gives every element a zero rect, so ChartFrame would only ever render
// its skeleton and none of the drawing code would run. Pin a width for the
// suite: that is what makes these smoke tests worth anything.
const realRect = HTMLElement.prototype.getBoundingClientRect
beforeAll(() => {
  HTMLElement.prototype.getBoundingClientRect = function rect() {
    return { width: 800, height: 240, top: 0, left: 0, right: 800, bottom: 240, x: 0, y: 0, toJSON: () => ({}) }
  }
})
afterAll(() => {
  HTMLElement.prototype.getBoundingClientRect = realRect
})

/** No chart may ever emit a NaN into an SVG attribute. */
function expectNoNaN(container: HTMLElement) {
  expect(container.innerHTML).not.toContain('NaN')
  expect(container.innerHTML).not.toContain('Infinity')
}

const activity = [
  { label: '1', values: { claims: 4, contributions: 3 } },
  { label: '2', values: { claims: 6, contributions: 5 } },
  { label: '3', values: { claims: 0, contributions: 0 } },
]
const series = [
  { key: 'claims', label: 'claims' },
  { key: 'contributions', label: 'contributions' },
]

describe('StackedBars', () => {
  it('draws a bar per datum and a legend per series', () => {
    const { container } = render(<StackedBars data={activity} series={series} />)
    expect(screen.getByRole('img', { name: /stacked activity/i })).toBeInTheDocument()
    expect(screen.getByText('claims')).toBeInTheDocument()
    // Two non-zero segments in each of the first two columns.
    expect(container.querySelectorAll('rect[fill="#00d992"]').length).toBe(2)
    expectNoNaN(container)
  })

  it('renders an empty state rather than an empty axis', () => {
    render(<StackedBars data={[]} series={series} />)
    expect(screen.getByText('No activity in this range')).toBeInTheDocument()
  })

  it('renders a skeleton while loading', () => {
    const { container } = render(<StackedBars data={activity} series={series} loading />)
    expect(container.querySelector('.animate-skeleton')).toBeTruthy()
  })
})

describe('BlockTimeline', () => {
  it('draws every window and the current-block marker', () => {
    const { container } = render(
      <BlockTimeline
        windows={[
          { label: 'selection', from: 100, to: 200, tone: 'selection' },
          { label: 'assembly', from: 200, to: 300, tone: 'assembly' },
        ]}
        current={250}
      />
    )
    expect(container.querySelector('circle')).toBeTruthy()
    expect(screen.getAllByText('selection').length).toBeGreaterThan(0)
    expectNoNaN(container)
  })

  it('survives a zero-width span', () => {
    const { container } = render(<BlockTimeline windows={[{ label: 'x', from: 10, to: 10 }]} current={10} />)
    expectNoNaN(container)
  })
})

describe('Matrix', () => {
  it('renders 64 × 16 with a background cell per position', () => {
    const cells = Array.from({ length: 200 }, (_, i) => ({ row: i % 64, col: i % 16, value: (i % 4) + 1 }))
    const { container } = render(
      <Matrix
        rows={Array.from({ length: 64 }, (_, i) => `#${i}`)}
        columns={Array.from({ length: 16 }, (_, i) => String(i))}
        cells={cells}
      />
    )
    expect(container.querySelectorAll('rect').length).toBe(64 * 16 + cells.length)
    expectNoNaN(container)
  })

  it('colours cells by wave when a colour is supplied', () => {
    const { container } = render(
      <Matrix rows={['a']} columns={['0']} cells={[{ row: 0, col: 0, color: waveColor(0, 3) }]} />
    )
    expect(container.querySelector('rect[fill="#00d992"]')).toBeTruthy()
  })
})

describe('Donut and Gauge', () => {
  it('draws one path per slice plus the centre value', () => {
    const { container } = render(
      <Donut
        slices={[
          { label: 'a', value: 3 },
          { label: 'b', value: 1 },
        ]}
        centerValue={4}
        centerLabel='total'
      />
    )
    expect(container.querySelectorAll('path').length).toBe(2)
    expect(screen.getByText('total')).toBeInTheDocument()
    expectNoNaN(container)
  })

  it('is empty rather than a zero-radius ring', () => {
    render(<Donut slices={[{ label: 'a', value: 0 }]} />)
    expect(screen.getByText('Nothing to break down')).toBeInTheDocument()
  })

  it('formats tiny gauge fractions with more precision', () => {
    const { container } = render(<Gauge value={0.0001} label='τ' />)
    expect(screen.getByText('0.010%')).toBeInTheDocument()
    expectNoNaN(container)
  })

  it('clamps an out-of-range gauge value', () => {
    render(<Gauge value={4} />)
    expect(screen.getByText('100.0%')).toBeInTheDocument()
  })
})

describe('CadenceStrip', () => {
  it('lanes overlapping epochs instead of hiding them', () => {
    const { container } = render(
      <CadenceStrip
        epochs={[
          { id: 'a', start: 0, end: 100, phase: 'closed' },
          { id: 'b', start: 50, end: 150, phase: 'live' },
          { id: 'c', start: 160, end: 220, phase: 'selection' },
        ]}
        current={120}
      />
    )
    const rects = container.querySelectorAll('rect')
    expect(rects.length).toBe(3)
    // a and c share a lane; b is pushed onto its own.
    const ys = Array.from(rects).map((r) => r.getAttribute('y'))
    expect(new Set(ys).size).toBe(2)
    expectNoNaN(container)
  })
})

describe('Sparkline', () => {
  it('draws a line and a last-point dot', () => {
    const { container } = render(<Sparkline values={[1, 4, 2, 8]} />)
    expect(container.querySelectorAll('path').length).toBe(2)
    expect(container.querySelector('circle')).toBeTruthy()
    expectNoNaN(container)
  })

  it('renders a dash for an empty series', () => {
    render(<Sparkline values={[]} />)
    expect(screen.getByText('—')).toBeInTheDocument()
  })

  it('survives a flat series', () => {
    const { container } = render(<Sparkline values={[5, 5, 5]} />)
    expectNoNaN(container)
  })
})
