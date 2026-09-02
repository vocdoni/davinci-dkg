// Runtime configuration: everything the explorer needs to know about *which*
// deployment it is looking at. It is fetched from `/config.json` at boot, not
// baked into the bundle, so one built image can be pointed at any chain by
// rewriting a single file (see `scripts/render-ui-config.sh`).

/** The exact shape `scripts/render-ui-config.sh` writes. */
export interface RuntimeConfigFile {
  rpcUrl: string
  managerAddress: `0x${string}`
  chainId: number
  chainName: string
  explorerUrl?: string
  /**
   * Block the `DKGManager` was deployed at. The indexer starts its historical
   * scan here, so a correct value is the difference between a handful of
   * `eth_getLogs` calls and a walk over the whole chain. `0` means "unknown".
   */
  deployBlock: number
}

/** What the app actually consumes: the file plus the demo switch. */
export interface RuntimeConfig extends RuntimeConfigFile {
  /**
   * True when the app must run entirely off the synthetic fixture with no RPC
   * at all. Set by `?demo=1` in the URL or `VITE_DEMO=1` at build time; the
   * data layer (stream B) reads it from `useRuntimeConfig()` and swaps the
   * indexer for the fixture store.
   */
  demo: boolean
}

export const CONFIG_URL = '/config.json'

/**
 * Config used when demo mode is requested and `/config.json` is missing or
 * unreadable — a demo must never depend on a deployment being configured.
 * The values are cosmetic: nothing in demo mode talks to a chain.
 */
export const DEMO_CONFIG: RuntimeConfig = {
  rpcUrl: 'http://demo.invalid',
  managerAddress: '0x00000000000000000000000000000000000d4691',
  chainId: 11155111,
  chainName: 'sepolia (demo)',
  explorerUrl: 'https://sepolia.etherscan.io',
  deployBlock: 0,
  demo: true,
}

/** `?demo=1` (or `?demo`) anywhere in the URL, or a `VITE_DEMO=1` build. */
export function isDemoRequested(search: string = typeof window === 'undefined' ? '' : window.location.search): boolean {
  if (import.meta.env.VITE_DEMO === '1' || import.meta.env.VITE_DEMO === 'true') return true
  const params = new URLSearchParams(search)
  if (!params.has('demo')) return false
  const value = params.get('demo')
  return value === null || value === '' || value === '1' || value === 'true'
}

/** Narrow an unknown JSON body into a RuntimeConfigFile, or throw. */
export function parseRuntimeConfig(raw: unknown): RuntimeConfigFile {
  if (typeof raw !== 'object' || raw === null) throw new Error('config.json is not an object')
  const obj = raw as Record<string, unknown>
  const str = (key: string): string => {
    const value = obj[key]
    if (typeof value !== 'string' || value === '') throw new Error(`config.json: "${key}" must be a non-empty string`)
    return value
  }
  const num = (key: string): number => {
    const value = obj[key]
    if (typeof value !== 'number' || !Number.isFinite(value)) throw new Error(`config.json: "${key}" must be a number`)
    return value
  }
  const manager = str('managerAddress')
  if (!/^0x[0-9a-fA-F]{40}$/.test(manager)) throw new Error('config.json: "managerAddress" is not an address')
  const explorerUrl = typeof obj.explorerUrl === 'string' && obj.explorerUrl !== '' ? obj.explorerUrl : undefined
  return {
    rpcUrl: str('rpcUrl'),
    managerAddress: manager as `0x${string}`,
    chainId: num('chainId'),
    chainName: str('chainName'),
    explorerUrl: explorerUrl?.replace(/\/+$/, ''),
    deployBlock: typeof obj.deployBlock === 'number' ? obj.deployBlock : 0,
  }
}

/**
 * Fetch and validate `/config.json`. In demo mode a missing or broken file is
 * not fatal — the fixture supplies the data anyway — so we fall back to
 * `DEMO_CONFIG` instead of blocking the boot.
 */
export async function loadRuntimeConfig(demo = isDemoRequested()): Promise<RuntimeConfig> {
  try {
    const res = await fetch(CONFIG_URL, { cache: 'no-store' })
    if (!res.ok) throw new Error(`${CONFIG_URL} HTTP ${res.status}`)
    return { ...parseRuntimeConfig(await res.json()), demo }
  } catch (err) {
    if (demo) return DEMO_CONFIG
    throw err instanceof Error ? err : new Error(String(err))
  }
}
