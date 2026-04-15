export function shortHash(hex: string | undefined | null, head = 6, tail = 4): string {
  if (!hex) return '—';
  const clean = hex.startsWith('0x') ? hex : `0x${hex}`;
  if (clean.length <= head + tail + 2) return clean;
  return `${clean.slice(0, 2 + head)}…${clean.slice(-tail)}`;
}

export function shortAddr(addr: string | undefined | null): string {
  return shortHash(addr, 6, 4);
}

export function formatBigInt(n: bigint | number | undefined | null): string {
  if (n === undefined || n === null) return '—';
  return BigInt(n).toString();
}

export function formatBlock(n: bigint | number | undefined | null): string {
  if (n === undefined || n === null) return '—';
  return `#${BigInt(n).toString()}`;
}

export function formatRoundId(id: string | undefined | null): string {
  if (!id) return '—';
  return id.startsWith('0x') ? id : `0x${id}`;
}

export function shortRoundId(id: string | undefined | null): string {
  return shortHash(id, 8, 6);
}

export function copyToClipboard(text: string): Promise<void> {
  if (navigator.clipboard) {
    return navigator.clipboard.writeText(text);
  }
  return Promise.resolve();
}

const ZERO_HASH = '0x0000000000000000000000000000000000000000000000000000000000000000';

export function isZeroHash(hex: string | undefined | null): boolean {
  if (!hex) return true;
  return hex.toLowerCase() === ZERO_HASH;
}

/**
 * Compute blocks remaining until a deadline.
 * Returns a positive number when deadline is in the future, negative when past.
 * Returns null when either argument is undefined.
 */
export function blocksRemaining(
  current: bigint | undefined | null,
  deadline: bigint | number | undefined | null,
): number | null {
  if (current === undefined || current === null) return null;
  if (deadline === undefined || deadline === null) return null;
  return Number(BigInt(deadline) - current);
}

/**
 * Format a blocks-remaining value as a human-readable string.
 * e.g. "+47 blocks" / "0 blocks" / "−12 blocks ago"
 */
export function formatBlocksRemaining(delta: number | null): string {
  if (delta === null) return '—';
  if (delta > 0) return `+${delta} block${delta === 1 ? '' : 's'}`;
  if (delta === 0) return 'closing now';
  return `${delta} blocks ago`;
}

export function formatTimestamp(ts: bigint | number | undefined | null): string {
  if (ts === undefined || ts === null) return '—';
  const n = Number(ts);
  if (!Number.isFinite(n) || n <= 0) return '—';
  return new Date(n * 1000).toLocaleString();
}
