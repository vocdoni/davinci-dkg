// Block-explorer URL helpers. The explorer base URL ships in the runtime
// config (set by the EXPLORER_URL env var at build time) — this module
// just appends the right path. Etherscan-compatible ("/address/0x…",
// "/tx/0x…", "/block/123") which covers Etherscan, Blockscout, and most
// chain-explorer clones in use today.

function trimTrailingSlash(url: string): string {
  return url.endsWith('/') ? url.slice(0, -1) : url
}

export function explorerAddressUrl(explorerUrl: string | undefined, address: string): string | null {
  if (!explorerUrl) return null
  return `${trimTrailingSlash(explorerUrl)}/address/${address}`
}

export function explorerTxUrl(explorerUrl: string | undefined, hash: string): string | null {
  if (!explorerUrl) return null
  return `${trimTrailingSlash(explorerUrl)}/tx/${hash}`
}

export function explorerBlockUrl(explorerUrl: string | undefined, block: bigint | number): string | null {
  if (!explorerUrl) return null
  return `${trimTrailingSlash(explorerUrl)}/block/${block.toString()}`
}
