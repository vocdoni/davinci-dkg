// Chart maths. Pure, dependency-free and unit-tested — the SVG components
// below only turn these numbers into elements, so a wrong axis is a failing
// test rather than a squint at a screenshot.

export interface Range {
  min: number
  max: number
}

export const clamp = (value: number, min: number, max: number): number => Math.min(Math.max(value, min), max)

/**
 * Maps `domain` onto `range`. A zero-width domain maps everything to the start
 * of the range (rather than dividing by zero), which is what a single-value
 * series should look like.
 */
export function linearScale(domain: [number, number], range: [number, number]): (value: number) => number {
  const [d0, d1] = domain
  const [r0, r1] = range
  const span = d1 - d0
  if (span === 0) return () => r0
  const k = (r1 - r0) / span
  return (value: number) => r0 + (value - d0) * k
}

/** 1-2-5-10 tick steps covering [min, max] with roughly `count` ticks. */
export function niceTicks(min: number, max: number, count = 5): number[] {
  if (!Number.isFinite(min) || !Number.isFinite(max)) return []
  if (min === max) return [min]
  if (min > max) [min, max] = [max, min]
  const rawStep = (max - min) / Math.max(count, 1)
  const magnitude = 10 ** Math.floor(Math.log10(rawStep))
  const normalized = rawStep / magnitude
  const step = (normalized <= 1 ? 1 : normalized <= 2 ? 2 : normalized <= 5 ? 5 : 10) * magnitude
  const start = Math.ceil(min / step) * step
  const ticks: number[] = []
  // Guard against float drift accumulating across the loop.
  for (let i = 0; start + i * step <= max + step * 1e-9; i++) {
    ticks.push(Number((start + i * step).toPrecision(12)))
    if (ticks.length > 1000) break
  }
  return ticks
}

export interface Band {
  /** Distance between the left edges of consecutive bands. */
  step: number
  /** Drawn width of one band. */
  bandWidth: number
  /** Left edge of band `i`. */
  at: (index: number) => number
  /** Centre of band `i`. */
  center: (index: number) => number
}

/** Categorical x axis: `count` evenly spaced bands across `width`. */
export function bandScale(count: number, width: number, padding = 0.2): Band {
  const n = Math.max(count, 1)
  const step = width / n
  const bandWidth = Math.max(step * (1 - clamp(padding, 0, 0.9)), 1)
  const offset = (step - bandWidth) / 2
  return {
    step,
    bandWidth,
    at: (index: number) => index * step + offset,
    center: (index: number) => index * step + step / 2,
  }
}

export interface StackSegment {
  key: string
  value: number
  y0: number
  y1: number
}

/**
 * Cumulative segments per row, bottom-up in `keys` order. Missing and negative
 * values count as zero: an event count is never negative, and a hole in the
 * data must not punch a hole in the bar.
 */
export function stackSeries<K extends string>(rows: Array<Partial<Record<K, number>>>, keys: K[]): StackSegment[][] {
  return rows.map((row) => {
    let acc = 0
    return keys.map((key) => {
      const value = Math.max(row[key] ?? 0, 0)
      const segment: StackSegment = { key, value, y0: acc, y1: acc + value }
      acc += value
      return segment
    })
  })
}

/** Largest stack total across all rows — the y domain of a stacked bar chart. */
export function stackMax<K extends string>(rows: Array<Partial<Record<K, number>>>, keys: K[]): number {
  return rows.reduce((max, row) => {
    const total = keys.reduce((sum, key) => sum + Math.max(row[key] ?? 0, 0), 0)
    return Math.max(max, total)
  }, 0)
}

/** Min/max of a numeric series, `{min: 0, max: 0}` when empty. */
export function extent(values: number[]): Range {
  if (values.length === 0) return { min: 0, max: 0 }
  let min = values[0] as number
  let max = values[0] as number
  for (const value of values) {
    if (value < min) min = value
    if (value > max) max = value
  }
  return { min, max }
}

/** 1234 → "1.2k", 5_400_000 → "5.4M". Used on axes and legends. */
export function formatCompact(value: number): string {
  const abs = Math.abs(value)
  if (abs < 1000) return Number.isInteger(value) ? String(value) : value.toFixed(1)
  if (abs < 1e6) return `${trimZero(value / 1e3)}k`
  if (abs < 1e9) return `${trimZero(value / 1e6)}M`
  return `${trimZero(value / 1e9)}B`
}

function trimZero(value: number): string {
  const fixed = value.toFixed(1)
  return fixed.endsWith('.0') ? fixed.slice(0, -2) : fixed
}

/** 0.4213 → "42.1%". */
export function formatPercent(fraction: number, digits = 1): string {
  return `${(fraction * 100).toFixed(digits)}%`
}

/** Point on a circle. Angles in degrees, 0° at 12 o'clock, clockwise. */
export function polarPoint(cx: number, cy: number, radius: number, angleDeg: number): { x: number; y: number } {
  const rad = ((angleDeg - 90) * Math.PI) / 180
  return { x: cx + radius * Math.cos(rad), y: cy + radius * Math.sin(rad) }
}

/**
 * SVG path for an annular arc (donut/gauge segment). Sweeps clockwise from
 * `startDeg` to `endDeg`; a full 360° sweep is drawn as two half arcs because
 * a single arc command with identical endpoints draws nothing.
 */
export function arcPath(
  cx: number,
  cy: number,
  outerR: number,
  innerR: number,
  startDeg: number,
  endDeg: number
): string {
  const sweep = endDeg - startDeg
  if (sweep <= 0) return ''
  if (sweep >= 360) {
    return [arcPath(cx, cy, outerR, innerR, 0, 180), arcPath(cx, cy, outerR, innerR, 180, 359.999)].join(' ')
  }
  const largeArc = sweep > 180 ? 1 : 0
  const o0 = polarPoint(cx, cy, outerR, startDeg)
  const o1 = polarPoint(cx, cy, outerR, endDeg)
  const i1 = polarPoint(cx, cy, innerR, endDeg)
  const i0 = polarPoint(cx, cy, innerR, startDeg)
  return [
    `M ${round(o0.x)} ${round(o0.y)}`,
    `A ${round(outerR)} ${round(outerR)} 0 ${largeArc} 1 ${round(o1.x)} ${round(o1.y)}`,
    `L ${round(i1.x)} ${round(i1.y)}`,
    `A ${round(innerR)} ${round(innerR)} 0 ${largeArc} 0 ${round(i0.x)} ${round(i0.y)}`,
    'Z',
  ].join(' ')
}

/** Polyline through points, `''` when there is nothing to draw. */
export function linePath(points: Array<{ x: number; y: number }>): string {
  if (points.length === 0) return ''
  return points.map((p, i) => `${i === 0 ? 'M' : 'L'} ${round(p.x)} ${round(p.y)}`).join(' ')
}

/** Closes a line path down to `baselineY` — the fill under a sparkline. */
export function areaPath(points: Array<{ x: number; y: number }>, baselineY: number): string {
  if (points.length === 0) return ''
  const first = points[0] as { x: number; y: number }
  const last = points[points.length - 1] as { x: number; y: number }
  return `${linePath(points)} L ${round(last.x)} ${round(baselineY)} L ${round(first.x)} ${round(baselineY)} Z`
}

function round(value: number): number {
  return Math.round(value * 100) / 100
}

/**
 * Block numbers arrive as bigints from viem and as numbers from the fixture.
 * Charts need doubles; block heights are far below 2^53 so the narrowing is
 * lossless in practice.
 */
export function num(value: number | bigint | null | undefined, fallback = 0): number {
  if (value == null) return fallback
  return typeof value === 'bigint' ? Number(value) : value
}
