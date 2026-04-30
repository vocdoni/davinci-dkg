# davinci-dkg UI rewrite — implementation plan (HISTORICAL)

A clean-slate rewrite of the legacy embedded React explorer into `ui/`,
mirroring the structure and coding conventions of
[`vocdoni/vocdoni-app`](https://github.com/vocdoni/vocdoni-app).

> **Status: complete.** The legacy `webapp/` directory has been removed and
> the Go binary is now UI-blind — the explorer ships only as a standalone
> Docker image (`ghcr.io/vocdoni/davinci-dkg-ui`). This document is kept as
> an architecture record; the up-to-date entry points are `ui/README.md`
> and the root `README.md`'s "Web Explorer" section.

---

## 1. Goals & non-goals

### Goals
- Feature parity with `webapp/` (Home, Rounds, EpochDetail, Registry, Settings, Playground).
- A noticeably less technical default UX. Plain English labels, sensible
  defaults in forms, no raw `0x…` walls of text in the primary view.
- A first-class **debug mode** toggle. When on, every page exposes the same
  raw-data inspector views the current webapp shows by default (event arg
  dumps, hex coordinates, polling traces). Off by default.
- A first-class **error reporting** affordance. Any error renders a plain
  English summary plus a "Copy error report" button that bundles
  `{ error.message, stack, route, chainId, address?, roundId?, blockNumber?,
  buildVersion, userAgent }` to the clipboard for issue filing.
- Code structure / conventions copied from vocdoni-app so anyone already
  working in that codebase can navigate this one without re-learning.
- Docker: standalone static-site image, runtime-templatable config.
- CI: build & publish that image alongside the existing node image.

### Non-goals (this rewrite)
- Replacing the embedded webapp inside the Go node. The new UI runs as its
  own service first; Go embed switchover is a follow-up.
- i18n: stub the `i18n/` folder for future use, but ship en-only.
- Light theme: dark-first, system-preference aware. Toggle UI ships but only
  dark+system are supported initially.
- Mobile design polish beyond Chakra's responsive defaults.
- Any new on-chain capability. Read/write surface mirrors what the SDK
  already exposes.

---

## 2. Tech stack

Picked to mirror `vocdoni-app` so the structure transfers cleanly. Where
vocdoni-app carries dead weight from its CRA→Vite migration or its own
ecosystem, we drop it.

| Concern | Choice | Notes |
|---|---|---|
| Build | Vite 7 + `@vitejs/plugin-react` + `vite-tsconfig-paths` + `vite-plugin-svgr` | Mirrors vocdoni-app. |
| Language | TypeScript 5.x, **`strict: true`** | vocdoni-app runs `strict: false`; we don't inherit that. |
| Runtime | React 18 with `StrictMode` | |
| UI | Chakra UI v3 (`@chakra-ui/react` ^3) + `next-themes` | Same major as vocdoni-app. New API (`createSystem`, slot recipes). |
| Icons | `react-icons` (lucide subset) | Match vocdoni-app's `lu*` usage. |
| Wallet | wagmi v2 + viem 2 + RainbowKit 2 | Env-gated WalletConnect projectId. |
| Data | `@tanstack/react-query` v5 | Sole data layer. |
| Routing | `react-router-dom` v6 with `createBrowserRouter`, lazy elements | Central `Routes` constant per vocdoni-app. |
| Forms | `react-hook-form` v7 | Inline validation rules. |
| i18n | `i18next` + `react-i18next` (en stub only) | Structure ready; one locale shipped. |
| Test | Vitest + Testing Library | Co-located `*.test.tsx`. |
| Lint | flat `eslint.config.js` + `typescript-eslint` + `react-hooks` + `react-refresh` + Prettier | Drop vocdoni-app's `react-app` preset; pick davinci-ui's flat config. |
| Format | Prettier — single quotes, no semicolons, `printWidth: 120`, `trailingComma: 'es5'` | Match vocdoni-app `.prettierrc` exactly. |
| Pkg mgr | pnpm 10 | Matches root repo. |

**Deliberately omitted** vs vocdoni-app: ethers v5 adapter (our SDK is
viem-native), `@vocdoni/react-components`, `@vocdoni/sdk`, Stripe, Crisp,
Memberbase, the wide-net 16-chain whitelist, the committed WalletConnect
projectId.

---

## 3. Directory layout

```
ui/
├── Dockerfile                  # multi-stage build → nginx static serve
├── docker/
│   ├── nginx.conf
│   └── entrypoint.sh           # templates /usr/share/nginx/html/config.json from env
├── .dockerignore
├── package.json
├── pnpm-lock.yaml
├── vite.config.ts
├── tsconfig.json
├── tsconfig.paths.json
├── tsconfig.node.json
├── eslint.config.js
├── .prettierrc
├── index.html
├── public/
│   └── config.json             # dev default; replaced at container start in prod
├── src/
│   ├── main.tsx                # entry; mounts <App/>
│   ├── App.tsx                 # provider tree (Chakra → wagmi → QueryClient → Router)
│   ├── router/
│   │   ├── Router.tsx          # RouterProvider + ScrollRestoration
│   │   ├── routes/
│   │   │   ├── index.ts        # Routes = { home: '/', rounds: '/rounds', round: '/rounds/:id', ... }
│   │   │   └── root.tsx        # createBrowserRouter table (lazy elements)
│   │   └── SuspenseLoader.tsx
│   ├── elements/               # route-level pages, kebab-case files
│   │   ├── layout.tsx          # public shell (Header/Footer/Outlet)
│   │   ├── home.tsx
│   │   ├── rounds/
│   │   │   ├── index.tsx       # rounds list
│   │   │   └── view.tsx        # round detail
│   │   ├── registry.tsx
│   │   ├── playground.tsx
│   │   ├── settings.tsx
│   │   └── error.tsx           # router errorElement
│   ├── components/             # PascalCase, grouped by domain
│   │   ├── Layout/
│   │   │   ├── Header.tsx
│   │   │   ├── Footer.tsx
│   │   │   ├── ConnectButton.tsx
│   │   │   ├── NetworkBadge.tsx
│   │   │   ├── DebugModeToggle.tsx
│   │   │   ├── QueryDataLayout.tsx
│   │   │   └── ConnectionToast.tsx
│   │   ├── Round/
│   │   │   ├── StatusBadge.tsx
│   │   │   ├── PhaseTimeline.tsx
│   │   │   ├── PhaseProgress.tsx
│   │   │   ├── RoundCard.tsx
│   │   │   ├── EpochList.tsx
│   │   │   ├── PolicyForm.tsx
│   │   │   ├── DecryptionPolicyForm.tsx
│   │   │   ├── EventLog.tsx
│   │   │   ├── ParticipantList.tsx
│   │   │   └── AbortButton.tsx
│   │   ├── Registry/
│   │   │   ├── NodeTable.tsx
│   │   │   └── RegistryStats.tsx
│   │   ├── Playground/
│   │   │   ├── PlaygroundShell.tsx     # step orchestration
│   │   │   ├── StepCard.tsx
│   │   │   ├── steps/
│   │   │   │   ├── ConnectStep.tsx
│   │   │   │   ├── CreateEpochStep.tsx
│   │   │   │   ├── WatchProgressStep.tsx
│   │   │   │   ├── KeyAvailableStep.tsx
│   │   │   │   ├── EncryptStep.tsx
│   │   │   │   ├── SubmitCiphertextStep.tsx
│   │   │   │   └── VerifyDecryptionStep.tsx
│   │   │   └── ActivityLog.tsx
│   │   ├── Form/
│   │   │   ├── InputBasic.tsx
│   │   │   ├── NumberInputBasic.tsx
│   │   │   ├── BlockOffsetInput.tsx    # "in N blocks (~30s)"
│   │   │   └── DurationDisplay.tsx
│   │   ├── Debug/
│   │   │   ├── DetailDisclosure.tsx    # collapsible "Show technical details"
│   │   │   ├── RawJson.tsx
│   │   │   └── ErrorReportButton.tsx
│   │   └── ui/
│   │       ├── HashCell.tsx
│   │       ├── AddressBadge.tsx
│   │       ├── ExternalLink.tsx
│   │       ├── EmptyState.tsx
│   │       ├── LoadingState.tsx
│   │       └── KeyValueGrid.tsx
│   ├── providers/
│   │   ├── ConfigProvider.tsx          # loads /config.json once
│   │   ├── DebugModeProvider.tsx
│   │   └── ThemeProvider.tsx
│   ├── queries/
│   │   ├── keys.ts                     # central QueryKeys registry
│   │   ├── chain.ts                    # block number, chain id
│   │   ├── rounds.ts                   # list, detail, events
│   │   ├── registry.ts                 # nodes, stats
│   │   └── playground.ts               # post-create poll, decryption watch
│   ├── hooks/
│   │   ├── use-dkg-client.ts
│   │   ├── use-dkg-writer.ts
│   │   ├── use-current-block.ts
│   │   ├── use-debug-mode.ts
│   │   └── use-clipboard-copy.ts
│   ├── theme/
│   │   ├── system.ts                   # createSystem
│   │   ├── tokens.ts
│   │   ├── semantic.ts
│   │   ├── color-mode.tsx
│   │   ├── recipes/
│   │   │   ├── button.ts
│   │   │   ├── card.ts
│   │   │   └── badge.ts
│   │   └── Theme.tsx
│   ├── lib/
│   │   ├── wagmi.ts                    # createConfig(chains, transports, connectors)
│   │   ├── format.ts                   # short(addr), shortHash(), formatBlocks(), formatBigInt()
│   │   ├── round-utils.ts              # phase derivation, deadline-remaining helpers
│   │   ├── error-report.ts             # buildErrorReport({error, ctx}) → string
│   │   └── debug.ts                    # localStorage-backed debug-mode helpers
│   ├── constants/
│   │   ├── chains.ts                   # supported chains for wagmi
│   │   ├── polling.ts                  # one place for refetchIntervals
│   │   └── routes.ts                   # re-export Routes for non-router consumers
│   ├── types/
│   │   └── index.ts
│   └── i18n/
│       ├── index.ts
│       └── locales/en/common.json
└── README.md
```

**Path aliases** (`tsconfig.paths.json` consumed by `vite-tsconfig-paths`):
`~components/*`, `~elements/*`, `~queries/*`, `~theme/*`, `~hooks/*`,
`~lib/*`, `~providers/*`, `~constants/*`, `~router`, `~types/*`.

---

## 4. UX philosophy: defaults vs debug mode

The current webapp's biggest UX problem is leaky abstraction — it dumps
bigints, hex coords, raw event arg keys, and protocol jargon directly into
the primary UI. The rewrite makes a hard cut:

- Every page renders **two layers**: a plain summary (always visible), and
  technical details (collapsed by default, auto-expanded when debug mode is
  on).
- One shared component, `<DetailDisclosure title="Show technical details">`,
  enforces this contract.
- A `DebugModeProvider` exposes `useDebugMode()` returning
  `{ enabled, setEnabled }`. Persisted to `localStorage:dkg-ui:debug`. A
  toggle in the header (gear icon) flips it globally and re-renders all
  disclosures expanded.
- When debug mode is **off**, the UI shows: short addresses (`0x1234…abcd`),
  block deltas as durations (`~3 min`), event names, status badges, and
  short SHA-style hashes only on hover/copy.
- When debug mode is **on**, it additionally shows: full hex hashes inline,
  raw event args dumped as JSON, polling status traces, BabyJubJub
  coordinates, calldata transcripts, gas usage.

**Error UX:**
- Failed mutations / queries surface as Chakra toasts and an inline
  `<ErrorBanner>` with: a one-line plain message (the SDK error mapped to
  human English), a "Copy error report" button, and a "Show details"
  disclosure containing the full error + stack. The "Copy error report"
  bundles enough context (route, chain, round if applicable, build SHA, UA)
  that an issue filer doesn't need to ask follow-ups.
- Top-level `errorElement` (`elements/error.tsx`) catches router-level
  errors and offers the same affordance.
- A persistent `<ConnectionToast>` (copied from vocdoni-app's
  `ConnectionToastProvider` pattern) tells the user when RPC reads start
  failing and when they recover.

---

## 5. Page plan

### `home.tsx` — `/`
- Four KPI cards: total rounds, active nodes, latest block, chain.
- "Recent rounds" table (5 items) with status badge + "View" link.
- All data via React Query; `<QueryDataLayout>` wrapper for loading/error.

### `rounds/index.tsx` — `/rounds`
- Filter bar: status (all / live / finalized / completed / aborted).
- Paginated table (client-side; ring-buffer cap of 64 from contract).
- Columns: round id (short), status, organizer (short), `t / n`, last activity.
- Click row → `/rounds/:id`.

### `rounds/view.tsx` — `/rounds/:id`
- **Header card**: status badge, plain English phase summary
  ("Waiting for 2 more contributions"), short id with copy.
- **Phase timeline**: visual horizontal timeline with the 5 phases
  (Registration → Contribution → Finalized → Decrypting → Completed) and
  current position. Replaces the current 5 progress bars.
- **Policy summary**: human-readable ("Threshold 2 of 3, registration
  closes in ~4 min", etc.); raw block numbers behind a disclosure.
- **Participants**: list of selected committee members with copy.
- **Activity**: condensed event log — one line per event with plain English
  ("Alice submitted contribution 1 of 3"); raw event args behind disclosure.
- **Debug-mode-only**: full EpochFinalized event hashes, transcript hashes,
  collective pubkey coordinates, decryption policy raw fields.

### `registry.tsx` — `/registry`
- Stats: active / total / inactivity window.
- Node table: address, status pill (Active/Inactive), last-active relative
  ("3 min ago"), public key hidden behind disclosure.
- Sticky table header.

### `playground.tsx` — `/playground`
A reorganised version of the current 7-step flow with:
- Better explanatory copy at every step (one sentence of why this happens).
- Form fields with sensible defaults so an "I just want to see it work"
  user can click through with default values.
- Each step has its own component file under `components/Playground/steps/`.
- The "Activity log" stays — useful for both researchers and bug reports —
  but lives in a collapsible right-rail (open by default for debug mode).
- `submitCiphertext` is gated on the round actually being Finalized (already
  fixed in the current webapp; rewrite preserves the gate).

### `settings.tsx` — `/settings`
- RPC override input (localStorage-backed; same key as current webapp so the
  setting carries over for users using the same domain).
- Chain info readout.
- Debug mode toggle (also accessible from header).
- "About" panel: build SHA + version, link to GitHub.

---

## 6. SDK integration

- A single `<ConfigProvider>` loads `/config.json` once at boot, exposes
  `{ rpcUrl, managerAddress, chainId, chainName, registryAddress?,
  startBlock? }` via context.
- One `useDkgClient()` hook constructs and memoises a `DKGClient` against
  the active config. Re-keys on RPC override changes.
- One `useDkgWriter(walletClient)` hook constructs a `DKGWriter` once a
  wallet is connected.
- All read calls live in `queries/*.ts` as `useQuery` hooks with keys from
  the central `QueryKeys` registry (`queries/keys.ts`):
  ```ts
  export const QueryKeys = {
    chain: ['chain'] as const,
    blockNumber: ['chain', 'blockNumber'] as const,
    roundsRecent: (limit: number) => ['rounds', 'recent', limit] as const,
    round: (id: `0x${string}`) => ['rounds', id] as const,
    roundEvents: (id: `0x${string}`, fromBlock?: bigint) =>
      ['rounds', id, 'events', fromBlock?.toString()] as const,
    registryNodes: ['registry', 'nodes'] as const,
    decryption: (id: `0x${string}`, ix: number) =>
      ['rounds', id, 'decryption', ix] as const,
  } as const
  ```
- All write calls (createRound, abortRound, submitCiphertext) live in their
  respective step components as `useMutation`s.

---

## 7. Docker plan

### `ui/Dockerfile`

Three stages:

1. **`sdk-builder`** — `node:20-bookworm-slim`, copies `sdk/`, runs
   `pnpm install && pnpm build`. Mirrors the existing root `Dockerfile`'s
   sdk-builder stage.
2. **`ui-builder`** — `node:20-bookworm-slim`, copies the prebuilt SDK from
   stage 1, then copies `ui/`, runs `pnpm install && pnpm build`. Output:
   `/root/ui/dist`.
3. **`runtime`** — `nginx:alpine`. Copies `dist/` into
   `/usr/share/nginx/html/`, copies `docker/nginx.conf` and
   `docker/entrypoint.sh`. The entrypoint:
   - Reads env vars `DAVINCI_DKG_RPC_URL`, `DAVINCI_DKG_MANAGER_ADDRESS`,
     `DAVINCI_DKG_CHAIN_ID`, `DAVINCI_DKG_CHAIN_NAME`,
     `DAVINCI_DKG_REGISTRY_ADDRESS` (optional),
     `DAVINCI_DKG_START_BLOCK` (optional).
   - Writes a `config.json` to `/usr/share/nginx/html/config.json` matching
     the `RuntimeConfig` shape the SDK expects.
   - `exec nginx -g 'daemon off;'`.
- Image exposes `:80`. Single-binary container (no Node at runtime).

### `docker-compose.yml` changes

Add a `ui` service. The existing `node` service is untouched (still serves
the embedded webapp on `:8081`). Profiles let users compose them:

```yaml
ui:
  profiles: ["ui", "node"]
  image: "ghcr.io/vocdoni/davinci-dkg-ui:${DAVINCI_DKG_UI_TAG:-latest}"
  build:
    context: ./
    dockerfile: ui/Dockerfile
  env_file: .env
  environment:
    DAVINCI_DKG_RPC_URL: "${DAVINCI_DKG_RPC_URL:-https://eth-sepolia.public.blastapi.io}"
    DAVINCI_DKG_MANAGER_ADDRESS: "${DAVINCI_DKG_MANAGER_ADDRESS:-0x01ee71fdce1705c8823f9f8b2f312100165fdd70}"
    DAVINCI_DKG_CHAIN_ID: "${DAVINCI_DKG_CHAIN_ID:-11155111}"
    DAVINCI_DKG_CHAIN_NAME: "${DAVINCI_DKG_CHAIN_NAME:-sepolia}"
  ports:
    - "${DAVINCI_DKG_UI_PORT:-8082}:80"
  restart: ${RESTART:-unless-stopped}
  labels:
    - "com.centurylinklabs.watchtower.enable=true"
```

The dual `["ui", "node"]` profile means:
- `docker compose --profile ui up` → just the UI (against an external RPC).
- `docker compose --profile node up` → node + UI together.
- `docker compose --profile ui --profile node up` → equivalent.

The existing node profile already exposes the embedded old webapp at
`:8081`. The new UI lands at `:8082` so the two coexist during transition.

### CI changes

Add `.github/workflows/ui-docker-build.yml` — a reusable workflow modelled
on the existing `docker-build.yml`. Builds `ui/Dockerfile` and pushes
`vocdoni/davinci-dkg-ui:<tag>` and `ghcr.io/vocdoni/davinci-dkg-ui:<tag>`,
with the same semver-→`latest` aliasing rule.

In `.github/workflows/main.yml`:
1. Add a `job_ui_build` (modelled on the existing `job_webapp`) that runs
   pnpm install + lint + build + uploads `ui-dist` artifact. Required so we
   catch UI build failures on PRs without going through Docker.
2. Add a `call-ui-docker-build` job, modelled on the existing
   `call-docker-build`, that runs after `job_ui_build` succeeds, on the
   same push/PR conditions. Pushes only on branch pushes.

---

## 8. Implementation phases

A full execution is ~5–10k lines of new code. Phased so each phase ships
something runnable.

### Phase 1 — Scaffold & infrastructure ✅ done
- `ui/` directory created with package.json, vite/ts/eslint/prettier config,
  Chakra v3 theme scaffold, provider tree, router skeleton, `Routes` constant,
  one rendered route (Home placeholder).
- `<DebugModeProvider>` + `<DetailDisclosure>` + `<ErrorReportButton>` wired
  up.
- `lib/wagmi.ts` configured (Sepolia + localhost), `<ConnectButton>` in
  header.
- `<ConfigProvider>` loading `/config.json`.
- `useDkgClient()` returning a memoised SDK client.
- `ui/Dockerfile` + `docker/entrypoint.sh` + `docker/nginx.conf`.
- `docker-compose.yml` updated with `ui` service.
- `.github/workflows/ui-docker-build.yml` + `main.yml` updates.

### Phase 2 — Read-only pages ✅ done
- `home.tsx` with 4-card KPI row + recent-rounds table.
- `rounds/index.tsx` list with filter chips.
- `rounds/view.tsx` detail with phase timeline, plain-English summary,
  policy KV grid, counters, and Participants/Activity tabs.
- `registry.tsx` with stats and node table (key coords behind disclosure).
- `settings.tsx` with RPC override (localStorage), debug toggle, chain
  info, build version.

### Phase 3 — Playground ✅ done
- 7 step components under `components/Playground/steps/` plus a sticky
  Activity Log right-rail.
- Wallet step uses RainbowKit's connect button (no manual EIP-1193
  handling).
- Create-round step exercises `PolicyForm` + `DecryptionPolicyForm` and
  fires `writer.createRound(...)`.
- Watch-progress step mirrors live status into the activity log.
- Encrypt step uses `buildElGamal()` and gates on Finalized.
- Verify step polls `getCombinedDecryption` and asserts the recovered
  plaintext matches what we sent.

### Phase 4 — Polish, tests, parity check ✅ done
- Vitest configured (`vitest.config.ts` + `vitest.setup.ts`).
- 27 unit tests covering `lib/format.ts`, `lib/round-utils.ts`, and
  `lib/error-report.ts`. All passing.
- `pnpm run test` wired into the `job_ui` CI job.
- README in `ui/` documenting dev workflow + Docker usage.

### Phase 5 — Cutover ✅ partial
- `ui/embed.go` created; `cmd/davinci-dkg-node/webapp.go` now imports
  `github.com/vocdoni/davinci-dkg/ui` and serves `ui/dist`.
- Root `Dockerfile` builds `ui/` instead of `webapp/`.
- `.github/workflows/main.yml` `job_webapp` replaced by `job_ui` (also
  runs `pnpm test`); `release.yml` `webapp` job replaced by `ui`.
- `docker-compose.yml`: standalone `ui` service is now opt-in only
  (`profiles: ["ui"]`); the `node` profile relies on the embedded UI on
  `:8081`. Standalone container stays on `:8082` to allow side-by-side
  hosting when both profiles are enabled.
- ⚠️  `webapp/` is **kept in tree** as a rollback hatch with a
  `webapp/DEPRECATED.md` note. Delete it in a follow-up commit once the
  new UI has been browser-verified against the production deployment(s).

---

## 9. Risks & open questions

- **Chakra v3 maturity in our context.** We adopt the same major as
  vocdoni-app, but several v3 APIs (slot recipes, system tokens) have
  shifted recently. Mitigation: keep the recipe layer thin in Phase 1.
- **WalletConnect projectId.** Required for RainbowKit. Will be sourced
  from `VITE_WALLETCONNECT_PROJECT_ID` env, with a clear "wallet connect
  unavailable" fallback if missing rather than a hard error.
- **Bundle size.** Chakra v3 + RainbowKit + wagmi is heavy (~500 kB
  minified). Acceptable for an explorer; we don't ship to mobile constrained
  environments.
- **Old webapp during transition.** Both UIs accessible (`:8081` Go embed,
  `:8082` standalone). Documentation in `ui/README.md` clarifies which is
  which until cutover.
