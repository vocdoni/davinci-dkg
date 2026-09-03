# Explorer architecture

How the explorer is put together: what each route shows, where its data comes
from, and the constraints every page has to hold. For the stack, the local
dev loop and the design rules, see [`README.md`](README.md). When this note
and the code disagree, the code wins; the protocol itself is described in
[`../README.md`](../README.md) and the contract interfaces in
[`../solidity/src/interfaces`](../solidity/src/interfaces).

## What it is for

The explorer surfaces every on-chain fact about an epoch, an operator, an
application or a ciphertext, with the block and transaction behind it. Nothing
is summarised away, and nothing is served by a backend: the browser talks to a
JSON-RPC endpoint and to nothing else.

Three things follow from that:

- **The protocol has to be visible, not just listed.** The lottery, committee
  assembly, key assembly, the decryption pipeline and operator behaviour over
  time are drawn.
- **Scale is a correctness requirement.** Committees of 64 members and
  registries of several hundred operators stay readable and fast.
- **The organizer needs a place to work.** The playground walks the organizer
  role against a live epoch and keeps the tools an organizer needs afterwards.

## Data layer

One in-browser indexer feeds every page. It scans all events of the three
contracts once, from `deployBlock` in `public/config.json`, in chunks through
the SDK's chunked `getLogs`; normalises them into an entity store; persists
that store in IndexedDB keyed by `chainId:managerAddress`; and keeps it
current by polling from the last indexed block. Contract state that events do
not carry is fetched with viem `multicall` in batches and cached per block
height: epoch structs, node records, application records, ciphertext and
combine records, plaintexts.

Pages never call `getLogs` themselves. They read a snapshot of the store
through a pure selector. [`src/data/README.md`](src/data/README.md) is the
reference for that path; the store shape lives in
`src/indexer/types.ts`.

The entities, all keyed, all carrying `block` and `tx` where they come from an
event:

| Entity | Key | Holds |
|---|---|---|
| `operators` | address | public key, status, registered block, last active, event history |
| `epochs` | epoch id | nonce, creator, start and seed blocks, seed, policy, status, committee in slot order, contributions, finalization, `PK_ep`, share commitment hashes, applications |
| `slots` | epoch, slot | operator, block, tx |
| `applications` | epoch, aid | creator, `PK_org`, policy, created block, ciphertexts |
| `ciphertexts` | epoch, aid, index | submitter, `C1`, `C2`, partials, organizer share, combined record and plaintext |
| `txMeta` | tx | sender, gas used, block number, fetched lazily |

Selectors are pure functions over the store and are unit-tested against the
fixture: network statistics, per-epoch progress, per-operator history and
participation, per-application pipeline state, activity per epoch, and wave
analysis of partials (which members answered a ciphertext, and in which
block).

## Fixture and demo mode

`src/fixtures/synthetic.ts` generates a deterministic network by pushing a
generated event stream through the real reducers, so the fixture store has the
same shape as a live one: 300 operators including reaped and reactivated rows,
8 epochs with 64-member committees at `t = 33`, one aborted epoch and one
still in key assembly, 2 applications per live epoch with 8 ciphertexts each,
partials arriving in waves of `t`, some organizer shares withheld, most
plaintexts combined.

Append `?demo=1` to any URL, or build with `VITE_DEMO=1`, and the whole app
runs from the fixture with no RPC. Read it from `useRuntimeConfig().demo`
rather than parsing the URL. Demo mode is how large tables and charts get
reviewed and screenshotted without a chain.

## Routes

`src/routes/paths.ts` is the single URL table. Build links from it, never from
string literals.

**`/` Overview.** Header strip with chain, manager, current block and the next
epoch countdown. Status cards for the newest epoch, live epochs, operators
active and registered, committee size and threshold in force, and ciphertexts
decrypted all time. An activity chart over the last 30 epochs, stacked by
claims, contributions, ciphertexts and partials. An epoch cadence strip on the
block axis with phase colouring. A live feed of the latest events with
transaction links, and the global search box that routes an epoch id,
application id, address or transaction hash to its page.

**`/epochs` Epochs.** Virtualised table: id, nonce, phase, `t` of `n`, claim
and contribution progress bars, ciphertexts, live since, creator, finalizer.
Filter by phase, search, click through to the epoch.

**`/epochs/:id` Epoch.** The header carries the id, nonce, phase badge,
creator and a block-axis lifecycle timeline with the current block marker and
the four windows. The lottery panel shows the seed block and hash, the
threshold as a fraction of the hash space, `α`, the snapshotted `R`, the
admission probability and the claims in slot order; when the epoch aborted it
shows why. The committee panel is a grid of `n` members that stays readable at
64, with claim block, contribution block, transaction, gas and `D_i` hash, over
a contributions summary bar. The key panel has the copyable `PK_ep`
coordinates, the finalizer, the finalization transaction and gas, and the
transcript size. Below that: the applications table, a members by ciphertexts
heatmap of partials coloured by wave, the epoch's full event log, and the raw
policy struct.

**`/operators` Operators.** Virtualised, searchable, sortable over hundreds of
rows: address, status, registered block, last active, epochs served, claims,
contributions, partials, finalizations, combines, participation, and the
expandable key. Participation reads `contributions/claims` and shows a dash
when there are no claims. Header cards cover active against registered, the
inactivity window and the newest epoch's committee. Charts show work per
operator (top 32, the rest grouped) and a status donut.

**`/operators/:address` Operator.** Identity, a history timeline across epochs
(claimed, contributed, finalized, partials, combines per epoch), a
participation sparkline, and every event with its transaction.

**`/applications` Applications.** Across epochs: epoch, aid, organizer,
submitter, window, cap, ciphertexts, decrypted, share status, with search.

**`/applications/:epoch/:aid` Application.** The record with the organizer,
`PK_org`, the derived `PK_aid` and the policy; the ciphertext table with
index, submission block, partials `t/n`, share, combined, plaintext and
transaction links; the partial matrix for this application; and the organizer
tools, which release a share from an index and a secret computed in the
browser, copy `PK_aid`, and resume the playground here.

**`/playground` Playground.** The organizer walkthrough: connect, choose a
live epoch, register with a generated or pasted secret, encrypt, submit,
release or withhold, watch the partials arrive with their waves and blocks
until the share and the combine land, then verify locally. State lives in the
URL (`?epoch=&aid=`) and in session storage, so the walkthrough is resumable.
Every step shows its transaction hash and gas, and an advanced toggle prints
the transcripts: proof of possession, discrete-logarithm equality (DLEQ)
words, `e` and `z`. Demo mode drives the same steps from the fixture with a
fake wallet.

**`/docs/protocol`, `/docs/run-a-node`, `/docs/sdk`.** The operator and
application documentation, under the design system's typography.

**`/kit`.** Every primitive and chart rendered with fixture data, for visual
review. Linked from the footer only. Open it after any change to the design
system.

## Scale requirements

These are checked against the fixture, which is sized to exceed them.

| Surface | Requirement |
|---|---|
| Operators table | 300+ rows virtualised; sort and search stay instant |
| Committee grid, partial matrix | 64 members legible; cells at least 10 px with hover detail |
| Epoch list | 200+ epochs paginated or virtualised |
| Indexer, first scan | chunked `getLogs` with adaptive chunk size; the UI stays usable while it runs, with a progress indicator |
| Indexer, steady state | at most one RPC round per poll when idle |

Any list that can exceed roughly 50 rows is virtualised. Wide panels scroll
inside themselves; the page never scrolls sideways.
