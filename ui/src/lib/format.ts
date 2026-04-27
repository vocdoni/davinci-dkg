// Pure presentation helpers shared across components. Keep these tiny and
// side-effect-free so they're trivially testable.

/** Truncate a 0x-prefixed hex string to "0xABCD…1234" form for compact display. */
export function shortHash(value: string | undefined | null, head = 6, tail = 4): string {
  if (!value) return ''
  if (!value.startsWith('0x')) return value
  if (value.length <= head + tail + 2) return value
  return `${value.slice(0, 2 + head)}…${value.slice(-tail)}`
}

/** Truncate an Ethereum address to "0xABCD…1234" form. */
export function shortAddress(addr: string | undefined | null): string {
  return shortHash(addr, 4, 4)
}

/** Render a bigint as an 0x-prefixed lowercase hex string of fixed 64 chars. */
export function bigIntToHex(value: bigint): `0x${string}` {
  return `0x${value.toString(16).padStart(64, '0')}`
}

/** Render a (possibly bigint-typed) block delta as a coarse human duration. */
export function blocksToDuration(blocks: number, secondsPerBlock = 12): string {
  if (blocks <= 0) return 'now'
  const seconds = blocks * secondsPerBlock
  if (seconds < 60) return `~${seconds}s`
  const minutes = Math.round(seconds / 60)
  if (minutes < 60) return `~${minutes} min`
  const hours = Math.round(minutes / 60)
  if (hours < 24) return `~${hours} h`
  const days = Math.round(hours / 24)
  return `~${days} d`
}

/** Compute "blocks remaining until target" given a current head; null if either input missing. */
export function blocksRemaining(current: bigint | null | undefined, target: bigint | number | undefined): number | null {
  if (current == null || target == null) return null
  const t = typeof target === 'bigint' ? target : BigInt(target)
  const delta = t - current
  return delta < 0n ? 0 : Number(delta)
}
