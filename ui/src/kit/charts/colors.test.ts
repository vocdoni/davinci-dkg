import { describe, expect, it } from 'vitest'
import { CHART_COLORS, heatColor, hexToRgb, mix, rgbToHex, seriesColor, waveColor } from './colors'

describe('hex helpers', () => {
  it('round-trips a colour', () => {
    expect(rgbToHex(hexToRgb('#00d992'))).toBe('#00d992')
  })
  it('expands 3-digit hex', () => {
    expect(hexToRgb('#0f0')).toEqual({ r: 0, g: 255, b: 0 })
  })
  it('mixes endpoints exactly and clamps t', () => {
    expect(mix('#000000', '#ffffff', 0)).toBe('#000000')
    expect(mix('#000000', '#ffffff', 1)).toBe('#ffffff')
    expect(mix('#000000', '#ffffff', 0.5)).toBe('#808080')
    expect(mix('#000000', '#ffffff', 2)).toBe('#ffffff')
    expect(mix('#000000', '#ffffff', -1)).toBe('#000000')
  })
})

describe('seriesColor', () => {
  it('starts on emerald and cycles', () => {
    expect(seriesColor(0)).toBe(CHART_COLORS.emerald)
    expect(seriesColor(6)).toBe(seriesColor(0))
    expect(seriesColor(-1)).toBeTruthy()
  })
})

describe('waveColor', () => {
  it('is pure emerald when everything arrived in one wave', () => {
    expect(waveColor(0, 1)).toBe(CHART_COLORS.emerald)
    expect(waveColor(3, 1)).toBe(CHART_COLORS.emerald)
  })

  it('runs emerald → slate across the waves', () => {
    expect(waveColor(0, 4)).toBe(CHART_COLORS.emerald)
    expect(waveColor(3, 4)).toBe(CHART_COLORS.slate)
  })

  it('cools monotonically: green falls as the wave index rises', () => {
    const greens = [0, 1, 2, 3, 4].map((w) => hexToRgb(waveColor(w, 5)).g)
    for (let i = 1; i < greens.length; i++) {
      expect(greens[i]!).toBeLessThan(greens[i - 1]!)
    }
  })

  it('clamps out-of-range waves', () => {
    expect(waveColor(99, 3)).toBe(waveColor(2, 3))
    expect(waveColor(-5, 3)).toBe(waveColor(0, 3))
  })
})

describe('heatColor', () => {
  it('reads zero as an empty cell, not a faint one', () => {
    expect(heatColor(0, 10)).toBe(CHART_COLORS.onyx)
    expect(heatColor(5, 0)).toBe(CHART_COLORS.onyx)
  })

  it('saturates at the maximum and above', () => {
    expect(heatColor(10, 10)).toBe(CHART_COLORS.emerald)
    expect(heatColor(50, 10)).toBe(CHART_COLORS.emerald)
  })

  it('is monotonic between charcoal and emerald', () => {
    const greens = [1, 3, 6, 9].map((v) => hexToRgb(heatColor(v, 10)).g)
    for (let i = 1; i < greens.length; i++) {
      expect(greens[i]!).toBeGreaterThan(greens[i - 1]!)
    }
  })
})
