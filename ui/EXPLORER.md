# davinci-dkg explorer — architecture and content plan

This is the specification for the explorer rebuild. Every stream implements
against it; when something is unclear, the protocol description in
`../README.md` and the contract interfaces in `../solidity/src/interfaces`
are the source of truth.

## 1. Goals

1. **A real explorer.** Every on-chain fact about an epoch, an operator, an
   application or a ciphertext is reachable, with the block and transaction
   behind it. Nothing is summarised away.
2. **Scale.** Committees of 64 members and registries of several hundred
   operators must stay readable and fast: virtualised tables, pagination,
   dense matrices, no per-page log scans.
3. **Visibility of the protocol.** Lottery, committee assembly, key assembly,
   decryption pipeline and operator behaviour over time are drawn, not only
   listed.
4. **The organizer's playground.** A resumable walkthrough of the organizer
   role against a live epoch, and the tools an organizer needs later
   (release a share, inspect its application).
5. **The design system in `ui/design/`** (Voltagent-inspired): obsidian canvas,
   one emerald accent, charcoal hairlines, Inter + JetBrains Mono, 4 px
   spacing, 4–8 px radii, borders over shadows, four surface levels.

## 2. Stack

- Vite + React 18 + TypeScript (strict), react-router v6, TanStack Query,
  wagmi/viem + RainbowKit for the wallet, `@vocdoni/davinci-dkg-sdk`
  (`link:../sdk`) for every contract read/write and all cryptography.
- **Tailwind CSS v4** with `ui/design/theme.css` as the `@theme`; global
  CSS variables from `ui/design/variables.css`. No Chakra. Headless
  primitives from `@radix-ui/react-*` only where behaviour is non-trivial
  (dialog, tooltip, popover, tabs). Fonts self-hosted:
  `@fontsource-variable/inter`, `@fontsource/jetbrains-mono`.
- Tables: `@tanstack/react-table` + `@tanstack/react-virtual` for any list that
  can exceed ~50 rows.
- Charts: hand-rolled SVG in `src/kit/charts` (no chart library).
- Persistence: `idb-keyval` (IndexedDB) for the indexer cache.
- Tests: vitest + testing-library (unit, with the synthetic fixture);
  Playwright (run by the integrator) against Sepolia and against demo mode.

## 3. Data layer (`src/indexer`)

One in-browser indexer feeds every page. It scans, once, all events of the
three contracts from `deployBlock` (from `public/config.json`) in chunks
through the SDK's chunked `getLogs`, normalises them into a store, persists
the store in IndexedDB keyed by `chainId:managerAddress`, and keeps it
current by polling from the last indexed block (every ~12 s, or on demand).
Contract state that events do not carry is fetched through viem `multicall`
in batches and cached per block height: epoch structs, node records,
application records, ciphertext/combine records, plaintexts.

Entities (all keyed, all with `block` and `tx` where they come from an event):

- `operators`: address → { pubKey, status, registeredAt, lastActive, events[] }
- `epochs`: id → { nonce, creator, startBlock, seedBlock, seed, threshold τ,
  policy {t, n, mMin, alphaBps, windows}, status, committee[] (slot order),
  contributions[] (index, contributor, block, tx, gasUsed), finalization
  {by, block, tx, gasUsed}, PK_ep, shareCommitmentHashes[], applications[] }
- `slots`: (epoch, slot) → { operator, block, tx }
- `applications`: (epoch, aid) → { creator, organizerPK, policy, createdAt,
  ciphertexts[] }
- `ciphertexts`: (epoch, aid, idx) → { submitter, c1, c2, block, tx,
  partials[] (participant, submitter, block, tx), organizerShare {block, tx,
  hash, overwrites}, combined {by, block, tx, plaintext} }
- `txMeta`: tx → { from, gasUsed, blockNumber } (lazy, fetched for
  finalizations, combines and any row the user opens)

Selectors are pure functions over the store, unit-tested with the fixture:
network stats, per-epoch progress, per-operator history and participation,
per-application pipeline state, activity per epoch, wave analysis of
partials (which members answered a ciphertext and in which block).

**Fixture and demo mode.** `src/fixtures/synthetic.ts` generates a
deterministic network: 300 registered operators (some inactive, some
reaped), 8 epochs with 64-member committees (t = 33), one aborted epoch,
applications with 8 ciphertexts each, partials arriving in waves of t, some
shares withheld, some plaintexts combined. `?demo=1` (or `/demo/...`) makes
the whole app run from the fixture with no RPC, so large tables and charts
can be reviewed and screenshotted without a chain.

## 4. Routes and content

`/` **Overview**
- Header strip: chain, manager, current block, next epoch countdown.
- Status cards: newest epoch (phase + block countdown), live epochs,
  operators active / registered, committee size and threshold in force,
  ciphertexts decrypted (all time).
- Activity chart: last 30 epochs, stacked bars (claims, contributions,
  ciphertexts, partials).
- Epoch cadence strip: epochs on the block axis with phase colouring.
- Latest events feed (20 rows, every event type, tx links), auto-refresh.
- Global search box (epoch id, application id, address, tx hash) → routes.

`/epochs` **Epochs**: virtualised table (id, nonce, phase, t of n, claims
progress bar, contributions progress bar, ciphertexts, live since, creator,
finalizer), phase filter, search; row → epoch.

`/epochs/:id` **Epoch**
- Header: id, nonce, phase badge, creator, block-axis lifecycle timeline
  with the current block marker and the four windows.
- Lottery panel: seed block and hash, τ as a fraction of the hash space,
  α, N snapshotted, admissible probability, claims in order (slot, operator,
  block, tx), abort reason when aborted.
- Committee panel: grid of n members (index, operator, claim block,
  contribution block/tx/gas, D_i hash) that stays readable at 64; summary
  bar of contributions t/n/m_min.
- Key panel: PK_ep coordinates (copyable), finalizer, finalization tx and
  gas, transcript size.
- Applications table (aid, organizer, submitter, policy window, cap,
  ciphertexts, decrypted) → application page.
- Decryption matrix: members × ciphertexts heatmap of partials (colour by
  wave/block), share status row, combined row.
- Event log: every event of the epoch with block and tx.
- Raw: policy fields, ABI-level struct.

`/operators` **Operators**: virtualised, searchable, sortable table for
hundreds of rows: address, status, registered block, last active, epochs
served, claims, contributions, partials, finalizations, combines,
participation (contributions/claims, "—" when no claims), key (expand).
Header cards: active/registered, inactivity window, committee of the
newest epoch. Chart: work per operator (top 32, rest grouped), and a
status donut.

`/operators/:address` **Operator**: identity (address, key, status,
registered, last active), history timeline across epochs (claimed /
contributed / finalized / partials / combines per epoch), participation
sparkline, table of every event with tx.

`/applications` **Applications**: table across epochs (epoch, aid, organizer,
submitter, window, cap, ciphertexts, decrypted, share status) with search.

`/applications/:epoch/:aid` **Application**: record (organizer, PK_org,
PK_aid derived and shown, policy), ciphertext table (idx, submitted block,
partials t/n, share, combined, plaintext, tx links), partial matrix for this
application, organizer tools: release a share (index + secret, computed in
the browser), copy PK_aid, resume the playground here.

`/playground` **Playground (organizer)**: steps connect → choose a live
epoch (newest by default, any Live selectable) → register (secret
generate/paste, aid, cap, submitter) → encrypt → submit → release or
withhold → watch (partials with wave/block, share, combine, plaintext) →
verify locally. State in URL (`?epoch=&aid=`) and session storage so it is
resumable; every step shows the transaction hash and gas; an "advanced"
toggle prints transcripts (PoP, DLEQ words, e, z). Demo mode drives the
same steps from the fixture with a fake wallet.

`/docs/protocol`, `/docs/run-a-node`, `/docs/sdk`: the existing texts,
restructured under one section with the design system's typography.

`/kit`: showcase of every primitive and chart with fixture data (for visual
review; linked from the footer only).

## 5. Design rules (from `ui/design/preview.html`)

Canvas `#050507`/`#101010` (obsidian, carbon), cards on carbon with
charcoal `#3d3a39` 1 px borders and 8 px radius, hover → onyx `#1a1a1a`
surface with warm-gray border; emerald `#00d992` only for the primary
action, active states, brand accents and the "live/ok" semantic; text
ghost `#ffffff` / silver `#bdbdbd` / pewter `#8b949e`; labels in 13 px
uppercase with 0.1 em tracking; JetBrains Mono for addresses, hashes,
numbers in tables and code; buttons primary (emerald fill), ghost
(emerald text/border), secondary (charcoal border); inputs transparent
with emerald focus glow; pill badges. Semantic colours beyond emerald are
limited to one warning amber and one danger red, used sparingly (phase
badges, errors). Charts use the four surface levels plus emerald and two
desaturated companions for series; every chart has a legend, a skeleton
and a tooltip.

## 6. Scale requirements (tested with the fixture)

- Operators table: 300+ rows virtualised; sort and search stay instant.
- Committee grid and partial matrix: 64 members legible; matrix cells ≥ 10 px
  with hover detail; horizontal scroll inside the panel, never the page.
- Epoch list: 200+ epochs paginated/virtualised.
- Indexer: initial scan bounded by chunked getLogs with adaptive chunk size;
  UI usable while scanning (progress indicator); incremental sync ≤ 1 RPC
  round per poll when idle.

## 7. Streams

- **A Foundation**: new scaffold on Tailwind v4 with the tokens, app shell,
  primitives in `src/kit`, charts in `src/kit/charts`, `/kit` showcase,
  routing skeleton with placeholder pages, wallet provider, config loader
  (`deployBlock`, `rpcUrl`, `managerAddress`, `chainId`, `explorerUrl`),
  demo-mode switch plumbing.
- **B Indexer**: `src/indexer`, `src/fixtures`, hooks in `src/data`, tests.
- **C Pages I**: overview, epochs, epoch detail.
- **D Pages II**: operators, operator detail, applications, application
  detail, global search.
- **E Playground + docs**.
- **F Review**: Playwright over every route on Sepolia and in demo mode;
  screenshots inspected by the integrator.
