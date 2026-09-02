// Chart palette. These hex values mirror the `--color-*` tokens declared in
// `src/styles/index.css`; they are duplicated here (rather than read as
// `var(--color-…)`) because wave and heat colouring interpolates between them,
// and you cannot interpolate a CSS variable in JavaScript.

export const CHART_COLORS = {
  emerald: '#00d992',
  emeraldSoft: '#10b981',
  /** The two desaturated companions the spec allows for series. */
  teal: '#4f9d8a',
  slate: '#7a8fa6',
  warmGray: '#5c5855',
  pewter: '#b8b3b0',
  charcoal: '#3d3a39',
  onyx: '#1a1a1a',
  carbon: '#101010',
  obsidian: '#050507',
  amber: '#e3b341',
  red: '#f85149',
} as const

/** Series order: emerald first, then the companions, then surface greys. */
export const SERIES_COLORS: string[] = [
  CHART_COLORS.emerald,
  CHART_COLORS.teal,
  CHART_COLORS.slate,
  CHART_COLORS.warmGray,
  CHART_COLORS.emeraldSoft,
  CHART_COLORS.pewter,
]

export function seriesColor(index: number): string {
  return SERIES_COLORS[((index % SERIES_COLORS.length) + SERIES_COLORS.length) % SERIES_COLORS.length] as string
}

interface Rgb {
  r: number
  g: number
  b: number
}

export function hexToRgb(hex: string): Rgb {
  const value = hex.replace('#', '')
  const full =
    value.length === 3
      ? value
          .split('')
          .map((c) => c + c)
          .join('')
      : value
  const int = Number.parseInt(full, 16)
  return { r: (int >> 16) & 255, g: (int >> 8) & 255, b: int & 255 }
}

export function rgbToHex({ r, g, b }: Rgb): string {
  const hex = (n: number) =>
    Math.round(Math.min(Math.max(n, 0), 255))
      .toString(16)
      .padStart(2, '0')
  return `#${hex(r)}${hex(g)}${hex(b)}`
}

/** Linear mix of two hex colours; `t=0` → `a`, `t=1` → `b`. */
export function mix(a: string, b: string, t: number): string {
  const clamped = Math.min(Math.max(t, 0), 1)
  const from = hexToRgb(a)
  const to = hexToRgb(b)
  return rgbToHex({
    r: from.r + (to.r - from.r) * clamped,
    g: from.g + (to.g - from.g) * clamped,
    b: from.b + (to.b - from.b) * clamped,
  })
}

/**
 * Colour for the n-th *wave* of partial decryptions. The first responders are
 * full emerald and later waves cool towards slate, so the decryption matrix
 * shows at a glance who answered first and who trailed. A single-wave epoch is
 * all emerald rather than a lone cold cell.
 */
export function waveColor(wave: number, totalWaves: number): string {
  if (totalWaves <= 1) return CHART_COLORS.emerald
  const t = Math.min(Math.max(wave, 0), totalWaves - 1) / (totalWaves - 1)
  return mix(CHART_COLORS.emerald, CHART_COLORS.slate, t)
}

/**
 * Sequential ramp for heatmap cells: onyx at zero (an empty cell must read as
 * "absent", not "a little"), then charcoal → emerald as the value climbs.
 */
export function heatColor(value: number, max: number): string {
  if (!(max > 0) || value <= 0) return CHART_COLORS.onyx
  const t = Math.min(value / max, 1)
  return mix(CHART_COLORS.charcoal, CHART_COLORS.emerald, t)
}

/** Colours for the four epoch phases; also the BlockTimeline window tones. */
export const PHASE_COLORS = {
  selection: CHART_COLORS.teal,
  assembly: CHART_COLORS.emerald,
  live: CHART_COLORS.emeraldSoft,
  finalize: CHART_COLORS.slate,
  closed: CHART_COLORS.warmGray,
  aborted: CHART_COLORS.red,
  pending: CHART_COLORS.charcoal,
} as const

export type PhaseColorKey = keyof typeof PHASE_COLORS
