import { describe, expect, it } from 'vitest'
import {
  arcPath,
  areaPath,
  bandScale,
  clamp,
  extent,
  formatCompact,
  formatPercent,
  linePath,
  linearScale,
  niceTicks,
  num,
  polarPoint,
  stackMax,
  stackSeries,
} from './scale'

describe('linearScale', () => {
  it('maps the domain onto the range', () => {
    const scale = linearScale([0, 100], [0, 200])
    expect(scale(0)).toBe(0)
    expect(scale(50)).toBe(100)
    expect(scale(100)).toBe(200)
  })

  it('inverts when the range is inverted (SVG y axis)', () => {
    const y = linearScale([0, 10], [120, 0])
    expect(y(0)).toBe(120)
    expect(y(10)).toBe(0)
    expect(y(5)).toBe(60)
  })

  it('collapses a zero-width domain to the range start', () => {
    const scale = linearScale([5, 5], [0, 100])
    expect(scale(5)).toBe(0)
    expect(scale(9)).toBe(0)
  })
})

describe('niceTicks', () => {
  it('produces 1-2-5-10 steps inside the domain', () => {
    expect(niceTicks(0, 100, 5)).toEqual([0, 20, 40, 60, 80, 100])
    expect(niceTicks(0, 10, 5)).toEqual([0, 2, 4, 6, 8, 10])
  })

  it('never leaves the domain', () => {
    for (const tick of niceTicks(3, 97, 4)) {
      expect(tick).toBeGreaterThanOrEqual(3)
      expect(tick).toBeLessThanOrEqual(97)
    }
  })

  it('handles a degenerate domain', () => {
    expect(niceTicks(7, 7)).toEqual([7])
    expect(niceTicks(Number.NaN, 3)).toEqual([])
  })

  it('handles small fractional domains without float drift', () => {
    expect(niceTicks(0, 1, 5)).toEqual([0, 0.2, 0.4, 0.6, 0.8, 1])
  })
})

describe('bandScale', () => {
  it('spaces bands evenly across the width', () => {
    const band = bandScale(4, 400, 0)
    expect(band.step).toBe(100)
    expect(band.bandWidth).toBe(100)
    expect(band.at(0)).toBe(0)
    expect(band.at(3)).toBe(300)
    expect(band.center(0)).toBe(50)
  })

  it('applies padding inside the step and stays centred', () => {
    const band = bandScale(2, 200, 0.5)
    expect(band.bandWidth).toBe(50)
    expect(band.at(0)).toBe(25)
    expect(band.center(1)).toBe(150)
  })

  it('never returns a zero-width band', () => {
    expect(bandScale(0, 100).bandWidth).toBeGreaterThan(0)
    expect(bandScale(500, 10, 0.9).bandWidth).toBeGreaterThanOrEqual(1)
  })
})

describe('stackSeries', () => {
  const keys = ['claims', 'contributions'] as const

  it('accumulates segments bottom-up in key order', () => {
    const [row] = stackSeries([{ claims: 3, contributions: 2 }], [...keys])
    expect(row).toEqual([
      { key: 'claims', value: 3, y0: 0, y1: 3 },
      { key: 'contributions', value: 2, y0: 3, y1: 5 },
    ])
  })

  it('treats missing and negative values as zero', () => {
    const [row] = stackSeries([{ contributions: -4 }], [...keys])
    expect(row?.map((s) => s.value)).toEqual([0, 0])
    expect(row?.[1]?.y1).toBe(0)
  })

  it('stackMax is the largest total', () => {
    const rows = [{ claims: 1, contributions: 1 }, { claims: 5 }, {}]
    expect(stackMax(rows, [...keys])).toBe(5)
    expect(stackMax([], [...keys])).toBe(0)
  })
})

describe('extent', () => {
  it('returns min and max', () => {
    expect(extent([3, 1, 9, 4])).toEqual({ min: 1, max: 9 })
  })
  it('is zeroed for an empty series', () => {
    expect(extent([])).toEqual({ min: 0, max: 0 })
  })
})

describe('formatters', () => {
  it('formats compact numbers', () => {
    expect(formatCompact(0)).toBe('0')
    expect(formatCompact(999)).toBe('999')
    expect(formatCompact(1200)).toBe('1.2k')
    expect(formatCompact(5_400_000)).toBe('5.4M')
    expect(formatCompact(2_000_000_000)).toBe('2B')
    expect(formatCompact(1000)).toBe('1k')
  })

  it('formats percentages', () => {
    expect(formatPercent(0.4213)).toBe('42.1%')
    expect(formatPercent(1, 0)).toBe('100%')
  })
})

describe('geometry', () => {
  it('puts 0° at 12 o’clock and 90° at 3 o’clock', () => {
    const top = polarPoint(0, 0, 10, 0)
    expect(top.x).toBeCloseTo(0)
    expect(top.y).toBeCloseTo(-10)
    const right = polarPoint(0, 0, 10, 90)
    expect(right.x).toBeCloseTo(10)
    expect(right.y).toBeCloseTo(0)
  })

  it('draws an annular arc, and nothing for an empty sweep', () => {
    expect(arcPath(50, 50, 40, 24, 0, 90)).toMatch(/^M .* A .* L .* A .* Z$/)
    expect(arcPath(50, 50, 40, 24, 30, 30)).toBe('')
  })

  it('splits a full circle into two arcs', () => {
    expect(arcPath(50, 50, 40, 24, 0, 360).match(/Z/g)).toHaveLength(2)
  })

  it('builds line and area paths', () => {
    const points = [
      { x: 0, y: 10 },
      { x: 10, y: 0 },
    ]
    expect(linePath(points)).toBe('M 0 10 L 10 0')
    expect(areaPath(points, 20)).toBe('M 0 10 L 10 0 L 10 20 L 0 20 Z')
    expect(linePath([])).toBe('')
    expect(areaPath([], 5)).toBe('')
  })

  it('narrows bigint block numbers', () => {
    expect(num(11_619_019n)).toBe(11_619_019)
    expect(num(42)).toBe(42)
    expect(num(null, -1)).toBe(-1)
    expect(num(undefined)).toBe(0)
  })

  it('clamps', () => {
    expect(clamp(5, 0, 3)).toBe(3)
    expect(clamp(-1, 0, 3)).toBe(0)
  })
})
