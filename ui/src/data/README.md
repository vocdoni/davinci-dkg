# `src/data` — how a page gets its data

Every page reads the chain through one in-browser indexer. It scans the three
contracts once, keeps an entity store in memory (and in IndexedDB), and
publishes a snapshot; the hooks here are `useSyncExternalStore` over that
snapshot plus a pure selector. There is no per-page `getLogs`, no TanStack
Query around contract reads, and no derived state in components.

```
viem PublicClient ─┐
                   ├─► Indexer ─► IndexerStore ─► selectors ─► hooks ─► pages
synthetic fixture ─┘   (src/indexer)  (src/indexer/types.ts)   (src/data)
```

## Wiring it up (once, in the app shell)

```tsx
import { createPublicClient, http } from 'viem'
import { DataSourceProvider, createDataSource } from '~/data'

const client = useMemo(
  () => createPublicClient({ transport: http(config.rpcUrl) }),
  [config.rpcUrl],
)

const source = useMemo(
  () => createDataSource({ demo, client, config }),   // config = public/config.json
  [demo, client, config],
)

<DataSourceProvider source={source}>
  <App />
</DataSourceProvider>
```

`config` is the shell's `RuntimeConfig` as-is: `chainId`, `managerAddress`,
`deployBlock`, `chainName`, `explorerUrl` and `demo`. `registryAddress` /
`appManagerAddress` are optional — the indexer reads them off the manager.
`demo` (from `?demo=1`) swaps in the synthetic fixture and makes zero RPC
calls; passing `demo` explicitly to `createDataSource` overrides the config.

The provider starts the source on mount and stops it on unmount. Mount it
**once**, above the router.

## Hooks

| Hook | Returns | Notes |
|---|---|---|
| `useIndexer()` | `{ status, kind, scanning, headBlock, lastBlock, progress, refresh, source }` | Scan progress, chain head, errors, manual refresh |
| `useNetworkStats()` | `NetworkStats` | Overview header and status cards |
| `useEpochs(filter?)` | `EpochRow[]` | Newest first; `filter = { phase, query, limit }` |
| `useEpoch(id)` | `EpochDetail \| null` | Lottery, committee, windows, applications, finalization, event log |
| `useOperators()` | `OperatorRow[]` | Registry members only, with work counters |
| `useOperator(address)` | `OperatorDetail \| null` | Row + per-epoch history + every event |
| `useApplications()` | `ApplicationRow[]` | Across all epochs, newest first |
| `useApplication(epoch, aid)` | `ApplicationDetail \| null` | Ciphertext pipeline, per index |
| `useActivity(n?)` | `ActivityBucket[]` | Newest `n` epochs, **oldest first** (chart order) |
| `useEventFeed(n?)` | `FeedEntry[]` | Newest `n` events, newest first |
| `usePartialMatrix(epoch, aid?)` | `PartialMatrix \| null` | Members × ciphertexts, with wave per cell |
| `useSearch(query, limit?)` | `SearchResult[]` | Epoch id/nonce, aid, address, tx hash; each carries an `href` (aliased as `useStoreSearch`, because the shell exports a `useSearch()` of its own) |
| `useIndexerSearchResolver()` | `SearchResolver` | Feed the shell's global search box: `useRegisterSearchResolver(useIndexerSearchResolver())` |
| `useTxMeta(hashes)` | `void` | Ask for `from`/`gasUsed` of rows the page shows |
| `useSnapshot()` / `useStore()` | raw snapshot / store | Escape hatch; prefer the above |

All of them are safe to call with `undefined` and return `null` / `[]` while
the first scan is still running — render a skeleton off
`useIndexer().scanning`, not off empty data.

```tsx
function EpochPage() {
  const { id } = useParams()
  const epoch = useEpoch(id)
  const matrix = usePartialMatrix(id)
  const { scanning, progress } = useIndexer()

  if (!epoch) return scanning ? <Scanning value={progress} /> : <NotFound />
  return <EpochHeader row={epoch.row} lottery={epoch.lottery} />
}
```

## Conventions worth knowing

- **Blocks are `number`s** everywhere in the store; curve coordinates,
  plaintexts and the lottery threshold τ stay `bigint`.
- **Slots are 0-based, participant indices are 1-based.** `SlotClaimed.slot`
  counts from 0; `contributorIndex` and `participantIndex` count from 1
  (`epochParticipants[i - 1]`; a partial's Merkle leaf is `index - 1`). Selectors
  expose both (`CommitteeRow.slot` / `.participantIndex`) and the partial
  matrix is addressed by **slot**. `ciphertextIndex` is 1-based too.
- **τ as a fraction**: `epochDetail(...).lottery.thresholdFraction` is
  `τ / 2²⁵⁶`, i.e. the share of the hash space that wins a slot;
  `registrySnapshot` is `R` recovered from τ, and `admissibleProbability` is
  `min(1, α·n/R)`.
- **Participation is `contributions / claims`**, and `null` when the operator
  never claimed — render `formatParticipation(row.participation)` to get the
  `—`.
- **Waves**: a partial's `wave` is
  `floor((partialBlock − ciphertextBlock) / staggerBlocks)`, with
  `staggerBlocks = 3` (the node's per-slot decryption delay). Wave 0 is the
  first `t` responders.
- **Attribution**: `EpochLive` and `DecryptionCombined` name no submitter, so
  `finalizer` / `combined.by` come from the transaction sender and are `null`
  until the indexer has fetched that receipt. It fetches them automatically
  (25 per poll); `useTxMeta` moves a row to the front of that queue.
- **Selectors are memoised on store identity** — calling `useEpochs()` in ten
  components costs one computation per publish.

## Demo mode

`createDemoDataSource()` serves `src/fixtures/synthetic.ts`: 300 operators,
8 epochs of 64 members (t = 33, m_min = 40), one aborted, one in KeyAssembly,
the whole pool of 16 keys stored at finalization, 2 applications × 8 ciphertexts
(one organizer-locked, one automatic), partials in waves, one organizer secret
still kept, gas figures from `BENCHMARKS.md`. Its head block
advances every 12 s. Pass options through `createDataSource({ demoOptions })`
to shrink it (tests use 24 operators / 4 epochs / committee 6).

## Cache

The store is persisted in IndexedDB under
`dkg-explorer:v<STORE_VERSION>:<chainId>:<manager>` and resumed on the next
visit; a schema bump, another chain or another manager address invalidates it
silently. Pass `kv: null` to `createDataSource` to disable persistence, or any
object with `get`/`set`/`del` (e.g. `idb-keyval` itself) to replace it.
