import { describe, expect, it } from 'vitest'
import { render, screen } from '@testing-library/react'
import { Theme } from '~theme/Theme'
import { BarChart } from './BarChart'
import { SeriesContributions, SeriesPartials } from './palette'

const series = [SeriesContributions, SeriesPartials]

function renderChart(groups: Parameters<typeof BarChart>[0]['groups']) {
  return render(
    <Theme>
      <BarChart series={series} groups={groups} ariaLabel='test chart' />
    </Theme>,
  )
}

describe('BarChart', () => {
  it('draws one mark per non-zero value and labels every group', () => {
    const { container } = renderChart([
      { key: 'a', label: 'a', values: { contributions: 3, partials: 1 } },
      { key: 'b', label: 'b', values: { contributions: 2, partials: 0 } },
    ])
    // 3 non-zero values → 3 bar paths (a zero value draws nothing).
    expect(container.querySelectorAll('path')).toHaveLength(3)
    expect(screen.getByLabelText('test chart')).toBeInTheDocument()
    expect(screen.getByText('a')).toBeInTheDocument()
  })

  it('always shows the legend, so identity is never colour-alone', () => {
    renderChart([{ key: 'a', label: 'a', values: { contributions: 1, partials: 1 } }])
    expect(screen.getAllByText('Contributions').length).toBeGreaterThan(0)
    expect(screen.getAllByText('Partial decryptions').length).toBeGreaterThan(0)
  })

  it('falls back to an empty state rather than an empty plot', () => {
    render(
      <Theme>
        <BarChart series={series} groups={[]} ariaLabel='empty' emptyMessage='Nothing yet.' />
      </Theme>,
    )
    expect(screen.getByText('Nothing yet.')).toBeInTheDocument()
  })

  it('renders an all-zero dataset without collapsing the axis', () => {
    const { container } = renderChart([
      { key: 'a', label: 'a', values: { contributions: 0, partials: 0 } },
    ])
    expect(container.querySelectorAll('path')).toHaveLength(0)
    // Axis still spans 0…2 so the plot keeps its shape.
    expect(screen.getByText('2')).toBeInTheDocument()
  })
})
