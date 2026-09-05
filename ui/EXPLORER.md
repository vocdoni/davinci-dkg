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
height: epoch structs and pool status, node records, application records,
ciphertext and combine records, plaintexts. Pool keys themselves arrive with
`PoolKeyActivated`, which carries the coordinates.

Pages never call `getLogs` themselves. They read a snapshot of the store
through a pure selector. [`src/data/README.md`](src/data/README.md) is the
reference for that path; the store shape lives in
`src/indexer/types.ts`.

The entities, all keyed, all carrying `block` and `tx` where they come from an
event:

| Entity | Key | Holds |
|---|---|---|
| `operators` | address | public key, status, registered block, last active, event history |
| `epochs` | epoch id | nonce, creator, start and seed blocks, seed, policy, status, committee in slot order, contributions, finalization (block, tx, frozen contribution count), the pool (`POOL_SIZE = 8` slots: `P_j` once activated, the aid that claimed it), next pool index, applications |
| `slots` | epoch, slot | operator, block, tx |
| `applications` | epoch, aid | creator, mode, pool index, `PK_org` (the identity when automatic), `sk_org` once revealed with the reveal block and tx, policy (submission, block window, cap, decryption window), created block, ciphertexts |
| `ciphertexts` | epoch, aid, index | submitter, `C1`, `C2`, partials, combined record and plaintext |
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
still in key assembly, four pool keys activated per live epoch (two claimed,
two ahead), 2 applications per live epoch with 8 ciphertexts each (one
organizer-locked with a one-address allow-list and a decryption window that
opened in 2023 — closed in 2025 on even epochs — and one automatic with open
submission and a deadline in 2033), partials arriving in waves of `t`, every
organizer's secret revealed right after registration except the newest live
epoch's — whose ciphertexts have no partials at all, since the contract refuses
partials and combines before the reveal, and sit at `awaiting-reveal` — most
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
and contribution progress bars, ciphertexts, live since, creator, finalizer
(the proof-less `finalizeEpoch` sender).
Filter by phase, search, click through to the epoch.

**`/epochs/:id` Epoch.** The header carries the id, nonce, phase badge,
creator and a block-axis lifecycle timeline with the current block marker and
the four windows. The lottery panel shows the seed block and hash, the
threshold as a fraction of the hash space, `α`, the snapshotted `R`, the
admission probability and the claims in slot order; when the epoch aborted it
shows why. The committee panel is a grid of `n` members that stays readable at
64, with claim block, contribution block, transaction and gas, over a
contributions summary bar. The pool panel has the finalizer, the finalization
transaction and gas, the frozen contribution count, and the eight pool slots —
each activated (copyable `P_j` coordinates, block, transaction, gas, activator),
claimed by which application ("activated by" the operator, "claimed by" the
application), or not activated yet — with the activation transcript size.
Below that: the applications table (the same compact columns as
`/applications`, with the pool key and the reveal state per row), a members by
ciphertexts heatmap of partials coloured by wave with a combined row, the
epoch's full event log, and the raw policy struct. The header's countdown reads
"service window · ended" once a live epoch's window has passed — the epoch
itself stays Live on chain.

Waves are numbered from 0 everywhere (this heatmap, the application page's
matrix, the playground): wave `w` is the `w`-th stagger window after the
ciphertext became decryptable — its own block, or the reveal block for an
organizer-locked application whose ciphertexts were waiting on it.

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

**`/applications` Applications.** Across epochs: epoch, aid, mode badge
(organizer-locked or automatic), organizer, submission policy (registrant
only, allow-list or open), block window ("any block" when unbounded), pool key
index with `P_j.x` under it, decryption window as dates ("any time" / "no
deadline" for an unbounded side) with its open / not yet open / closed state,
ciphertexts over the cap (`8 / ∞`), decrypted, organizer secret (revealed,
kept, or none), with search and mode and epoch filters. The header counts the
organizer-locked applications whose secret is still kept. Both application
tables fit a 1440 px viewport without a horizontal scroll: every column is a
fixed track except the decryption window, which takes the slack, and the pool
key and the two windows stack on two lines rather than widen.

**`/applications/:epoch/:aid` Application.** The record with the mode, the
pool key index and `P_j`, the organizer, `PK_org` and the reveal state (block
and transaction) for an organizer-locked application, the derived `PK_aid`
(`P_j` or `P_j + PK_org`), the submission policy with its addresses, the
decryption window as dates with its open / not yet open / closed state (the
coordinate pairs `P_j`, `PK_org`, `PK_aid` always sit side by side in the
two-column record; the reveal state closes the list); the ciphertext table
with index, state (`submitted`, `partials`, `awaiting-reveal`, `ready`,
`combined`), submission block, partials `t/n`, combined, plaintext and
transaction links; the partial matrix for this application; and the organizer
panel: for an organizer-locked application whose secret is kept, the reveal
tool (the secret is checked against `PK_org` in the browser before the one
irreversible transaction); once revealed, the block, transaction and the
public secret; for an automatic application, a note that there is no organizer
key. The header copies `PK_aid` and resumes the playground here.

`awaiting-reveal` is every ciphertext of an organizer-locked application until
its organizer reveals: the contract refuses partials and combines before that
(`OrganizerSecretNotRevealed`), so there is nothing else such a ciphertext can
be, and the page says so in the state help, the stat card and the organizer
panel.

**`/playground` Playground.** The organizer walkthrough: connect, choose a
live epoch (each shows its pool as "2 activated free · 2 claimed · 4 not
activated"), pick a mode, a cap, a submitter and a decryption window and
register — claiming the next pool key — with a generated or pasted secret
(none for an automatic application), encrypt under the key
`getApplicationKey` returns, submit, reveal the organizer secret or keep it
for now (the rail marks this step *skipped* for an automatic application),
watch the partials arrive with their waves and blocks until the combine
lands, then verify locally. For an organizer-locked application the partials
only start after the reveal — the contract refuses them before it — and the
watch step says so while the secret is kept. State lives in the URL
(`?epoch=&aid=`) and in session storage, so the walkthrough is resumable.
Every step shows its transaction hash and gas; the activity log records every
transaction this tab sent and the combine once the committee lands it. An
advanced toggle prints the transcripts: the proof of possession words, the
pool key, `PK_org` and `PK_aid`, and the two words of the reveal — as
decimals, the same form the encrypt step shows `PK_aid.x` in. Demo mode drives
the same steps from the fixture with a fake wallet and real curve points for
the pool keys; it does not simulate the decryption window, but it does enforce
the reveal gate, so a locked application's partials arrive over the seconds
after the reveal rather than being there already.

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
