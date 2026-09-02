import { isAddress } from 'viem'
import { paths } from '~routes/paths'

/**
 * Where a search query should send the user.
 *
 * `route` is in-app; `external` is the block explorer; `ambiguous` means the
 * shape matched more than one entity (a 32-byte value is both a transaction
 * hash and an application id) and the caller should offer the choices;
 * `unknown` is a dead end with a reason to show.
 */
export type SearchTarget =
  | { kind: 'route'; path: string; label: string }
  | { kind: 'external'; url: string; label: string }
  | { kind: 'ambiguous'; label: string; options: SearchTarget[] }
  | { kind: 'unknown'; label: string }

export interface SearchContext {
  /** Block-explorer base URL from the runtime config, if any. */
  explorerUrl?: string
}

/** An extra resolver contributed by a later stream (indexer/pages). */
export type SearchResolver = (query: string, ctx: SearchContext) => SearchTarget | null

const EPOCH_ID = /^0x[0-9a-fA-F]{24}$/ // bytes12
const BYTES32 = /^0x[0-9a-fA-F]{64}$/ // tx hash or application id
const DIGITS = /^\d+$/

/**
 * Shape-based routing for the global search box.
 *
 * Deliberately dumb: it only knows what a value *looks* like. Anything that
 * needs the store to disambiguate — which epoch an application id belongs to,
 * whether a 32-byte value is a known aid or a transaction — is contributed by
 * the indexer stream through `registerSearchResolver`, which runs first.
 */
export function resolveSearch(rawQuery: string, ctx: SearchContext = {}, extra: SearchResolver[] = []): SearchTarget {
  const query = rawQuery.trim()
  if (query === '') return { kind: 'unknown', label: 'Type an epoch id, application id, address, block or transaction' }

  for (const resolver of extra) {
    const hit = resolver(query, ctx)
    if (hit) return hit
  }

  if (isAddress(query, { strict: false })) {
    return { kind: 'route', path: paths.operator(query), label: `Operator ${query}` }
  }

  if (EPOCH_ID.test(query)) {
    return { kind: 'route', path: paths.epoch(query), label: `Epoch ${query}` }
  }

  if (BYTES32.test(query)) {
    // Both a transaction hash and an application id are 32 bytes. Without the
    // index we can only offer the explorer; the indexer stream adds the
    // application branch by registering a resolver.
    if (ctx.explorerUrl) {
      return {
        kind: 'external',
        url: `${ctx.explorerUrl.replace(/\/+$/, '')}/tx/${query}`,
        label: 'Transaction on the block explorer',
      }
    }
    return { kind: 'unknown', label: 'Looks like a transaction hash, but no block explorer is configured' }
  }

  if (DIGITS.test(query)) {
    if (ctx.explorerUrl) {
      return {
        kind: 'external',
        url: `${ctx.explorerUrl.replace(/\/+$/, '')}/block/${query}`,
        label: `Block ${query} on the block explorer`,
      }
    }
    return { kind: 'unknown', label: 'Looks like a block number, but no block explorer is configured' }
  }

  return { kind: 'unknown', label: `No epoch, application, operator or transaction matches “${query}”` }
}
