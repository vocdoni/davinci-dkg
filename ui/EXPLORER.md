# Explorer architecture

How the explorer is put together: where the data comes from, what each route shows and the
limits every page has to hold. For configuration, hosting and the development loop see
[README.md](README.md). The protocol is described in [docs/protocol.md](../docs/protocol.md) and
the contract interfaces in [`solidity/src/interfaces`](../solidity/src/interfaces).

The explorer shows every on-chain fact about an epoch, an operator, an application or a
ciphertext, with the block and transaction behind it. The browser talks to a JSON-RPC endpoint
and to nothing else.

## Data layer

One in-browser indexer feeds every page. It scans all events of the three contracts once, from
`deployBlock` in `public/config.json`, in chunks through the SDK; reduces them into an entity
store; persists the store in IndexedDB keyed by chain id and manager address; and then polls from
the last indexed block. Contract state that events do not carry is read with `multicall` in
batches and cached per block: epoch records and the pool cursor, node records, application
records, ciphertext and combine records, plaintexts. Pool keys appear in no event, so the 16 keys
of a `Live` epoch are read once, right after `EpochLive`.

Pages never query logs themselves. They read a snapshot of the store through pure, memoised
selectors; [`src/data/README.md`](src/data/README.md) documents the hooks, and the store shape
lives in `src/indexer/types.ts`.

| Entity | Key | Holds |
|---|---|---|
| `operators` | address | Public key, status, registration block, last activity, event history |
| `epochs` | epoch id | Nonce, creator, start and seed blocks, seed, policy, status, committee in slot order, contributions, finalization, the 16 pool slots (key and claiming application), applications |
| `slots` | epoch, slot | Operator, block, transaction |
| `applications` | epoch, aid | Creator, mode, pool index, `PK_org`, revealed `sk_org` with its block and transaction, policy, ciphertexts |
| `ciphertexts` | epoch, aid, index | Submitter, `C1`, `C2`, partials, combined record and plaintext |
| `txMeta` | transaction | Sender and gas used, fetched lazily |

Every entity that comes from an event carries its block and transaction.

## Demo mode

`src/fixtures/synthetic.ts` generates a deterministic network by pushing a generated event stream
through the real reducers, so the fixture store has the same shape as a live one: 300 operators
including reaped and reactivated ones, 8 epochs with 64-member committees, one aborted epoch and
one still in key assembly, organizer-locked and automatic applications, partials arriving in
waves, one organizer secret still kept. `?demo=1` or a `VITE_DEMO=1` build runs the whole app from
it with no RPC; read the flag from `useRuntimeConfig().demo`. Large tables and charts are
reviewed in demo mode.

## Routes

`src/routes/paths.ts` is the URL table; build links from it.

| Route | Shows |
|---|---|
| `/` | Chain and manager, current block and the next epoch countdown, status cards, activity over the last 30 epochs, the epoch cadence on the block axis, a live event feed and the global search (epoch id, aid, address or transaction hash) |
| `/epochs` | All epochs with phase, `t` of `n`, claim and contribution progress, ciphertexts, creator and finalizer |
| `/epochs/:id` | Lifecycle timeline, the lottery (seed, threshold, `α`, `R`, admission probability, claims), the committee grid, the pool of 16 keys and who claimed them, the applications, a members by ciphertexts matrix of partials coloured by wave, the event log and the raw policy |
| `/operators` | Every operator with status, activity counters and participation, plus charts of work per operator |
| `/operators/:address` | Identity, per-epoch history and every event |
| `/applications` | Applications across epochs: mode, organizer, submission policy, pool key, decryption window and its state, ciphertexts and the reveal state |
| `/applications/:epoch/:aid` | The application record and keys (`P_j`, `PK_org`, `PK_aid`), its ciphertexts and their state, the partial matrix, and the reveal tool for the organizer |
| `/playground` | The organizer's steps against a live epoch: register, encrypt, submit, reveal, watch the committee decrypt and verify locally. Resumable through the URL and session storage |
| `/docs/*` | Protocol, operator and SDK documentation |
| `/kit` | Every component and chart with sample data, linked from the footer |

A ciphertext is `submitted`, `partials`, `awaiting-reveal`, `ready` or `combined`. Every
ciphertext of an organizer-locked application is `awaiting-reveal` until the reveal, since the
contract refuses partials before it. Waves are numbered from 0: wave `w` is the `w`-th stagger
window after the ciphertext became decryptable (its own block, or the reveal block).

## Scale requirements

The fixture is sized to exceed these.

| Surface | Requirement |
|---|---|
| Operators table | 300+ rows virtualised; sorting and search stay instant |
| Committee grid, partial matrix | 64 members legible; cells at least 10 px, with hover detail |
| Epoch list | 200+ epochs, paginated or virtualised |
| First scan | Chunked `getLogs` with adaptive chunk size; the UI stays usable and shows progress |
| Steady state | At most one RPC round per poll when idle |

Any list that can exceed about 50 rows is virtualised. Wide panels scroll inside themselves; the
page never scrolls sideways.
