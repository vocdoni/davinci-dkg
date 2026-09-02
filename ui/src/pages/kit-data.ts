// Deterministic sample data for the /kit showcase. It is *not* the synthetic
// fixture (that is stream B's `src/fixtures`): it exists only so the design
// review has realistic shapes — 300 addresses, a 64 × 16 matrix, 30 epochs of
// activity — without a chain or an indexer.

/** mulberry32: 32-bit PRNG, seeded, so the showcase renders identically twice. */
function rng(seed: number): () => number {
  let a = seed >>> 0
  return () => {
    a = (a + 0x6d2b79f5) >>> 0
    let t = Math.imul(a ^ (a >>> 15), 1 | a)
    t = (t + Math.imul(t ^ (t >>> 7), 61 | t)) ^ t
    return ((t ^ (t >>> 14)) >>> 0) / 4294967296
  }
}

const hex = (random: () => number, bytes: number): string =>
  `0x${Array.from({ length: bytes }, () =>
    Math.floor(random() * 256)
      .toString(16)
      .padStart(2, '0')
  ).join('')}`

export interface DemoOperator {
  address: string
  status: 'active' | 'idle' | 'reaped'
  registeredAt: number
  lastActive: number
  epochs: number
  claims: number
  contributions: number
  partials: number
  trend: number[]
}

export function demoOperators(count = 300): DemoOperator[] {
  const random = rng(0x5eed)
  return Array.from({ length: count }, (_, i) => {
    const claims = Math.floor(random() * 40)
    const contributions = Math.max(0, claims - Math.floor(random() * 4))
    const status = random() < 0.08 ? 'reaped' : random() < 0.2 ? 'idle' : 'active'
    return {
      address: hex(random, 20),
      status,
      registeredAt: 11_500_000 + Math.floor(random() * 80_000),
      lastActive: 11_890_000 + Math.floor(random() * 12_000),
      epochs: Math.max(claims, 1),
      claims,
      contributions,
      partials: contributions * (2 + Math.floor(random() * 6)),
      trend: Array.from({ length: 12 }, () => Math.floor(random() * 10) + i * 0),
    }
  })
}

export interface DemoActivity {
  label: string
  values: Record<string, number>
  note: string
}

export function demoActivity(epochs = 30): DemoActivity[] {
  const random = rng(0xa17c)
  return Array.from({ length: epochs }, (_, i) => ({
    label: String(120 + i),
    note: `blocks ${(11_700_000 + i * 3600).toLocaleString()} → ${(11_700_000 + (i + 1) * 3600).toLocaleString()}`,
    values: {
      claims: 40 + Math.floor(random() * 24),
      contributions: 30 + Math.floor(random() * 20),
      ciphertexts: Math.floor(random() * 12),
      partials: 40 + Math.floor(random() * 90),
    },
  }))
}

export interface DemoMatrix {
  rows: string[]
  columns: string[]
  cells: Array<{ row: number; col: number; wave: number; block: number }>
  waves: number
}

/**
 * A 64-member committee against 16 ciphertexts, with partials arriving in four
 * waves — the exact shape the epoch page's decryption matrix has to survive.
 */
export function demoMatrix(members = 64, ciphertexts = 16, waves = 4): DemoMatrix {
  const random = rng(0xd4c)
  const cells: DemoMatrix['cells'] = []
  for (let col = 0; col < ciphertexts; col++) {
    for (let row = 0; row < members; row++) {
      // ~72% of members answer each ciphertext; the rest are the holes.
      if (random() > 0.72) continue
      const wave = Math.min(waves - 1, Math.floor(random() * waves))
      cells.push({ row, col, wave, block: 11_902_000 + wave * 30 + Math.floor(random() * 20) })
    }
  }
  return {
    rows: Array.from({ length: members }, (_, i) => `#${String(i).padStart(2, '0')} ${hex(rng(i + 1), 2)}`),
    columns: Array.from({ length: ciphertexts }, (_, i) => String(i)),
    cells,
    waves,
  }
}

export interface DemoEpoch {
  id: string
  label: string
  start: number
  end: number
  phase: 'selection' | 'assembly' | 'live' | 'finalize' | 'closed' | 'aborted'
}

export function demoCadence(count = 10): DemoEpoch[] {
  const random = rng(0xcade)
  const phases: DemoEpoch['phase'][] = [
    'closed',
    'closed',
    'closed',
    'aborted',
    'closed',
    'finalize',
    'live',
    'live',
    'assembly',
    'selection',
  ]
  return Array.from({ length: count }, (_, i) => {
    const start = 11_860_000 + i * 4200 + Math.floor(random() * 400)
    return {
      id: `epoch-${i}`,
      label: `epoch ${140 + i}`,
      start,
      end: start + 7200,
      phase: phases[i % phases.length] as DemoEpoch['phase'],
    }
  })
}

export const demoWindows = [
  { label: 'committee selection', from: 11_898_000, to: 11_899_800, tone: 'selection' as const },
  { label: 'key assembly', from: 11_899_800, to: 11_901_600, tone: 'assembly' as const },
  { label: 'live', from: 11_901_600, to: 11_905_200, tone: 'live' as const },
  { label: 'finalize', from: 11_905_200, to: 11_906_000, tone: 'finalize' as const },
]

export const demoSparkline = [3, 7, 5, 9, 12, 8, 14, 11, 15, 13, 18, 16]
