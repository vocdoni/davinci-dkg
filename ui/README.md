# davinci-dkg UI

Standalone Vite + React + TypeScript SPA for the davinci-dkg explorer,
playground, and operator/SDK documentation. `EXPLORER.md` is the
specification it is built against; `design/` is the design package (tokens,
theme, preview) that defines every colour, radius and type step.

Stack: Vite 5 + React 18 + TypeScript (strict), Tailwind CSS v4 (with
`design/theme.css` as the `@theme`), react-router v6, TanStack Query /
Table / Virtual, wagmi + viem + RainbowKit, `@vocdoni/davinci-dkg-sdk`
(`link:../sdk`), Radix primitives for dialog/tooltip/popover/tabs, and
hand-rolled SVG charts (no chart library).

The UI is fully decoupled from the `davinci-dkg-node` Go binary. It talks
to the chain directly via JSON-RPC and ships as its own Docker image
(`ghcr.io/vocdoni/davinci-dkg-ui`).

## Local development

```sh
cd ui
pnpm install
pnpm dev          # http://localhost:5174 (also bound on 0.0.0.0)
```

The dev server binds on `0.0.0.0:5174` so it's reachable from containers,
VMs, and other hosts on the LAN. Override with
`pnpm dev --host 127.0.0.1` if you need a tighter bind.

The dev server reads `public/config.json` for chain + manager address.
Edit it directly, or have `make` template it from env vars:

```sh
make ui-dev \
  RPC_URL=http://127.0.0.1:8545 \
  MANAGER_ADDRESS=0xabc... \
  CHAIN_ID=31337 CHAIN_NAME=anvil
```

Recognised vars (all optional; defaults match the bundled
`public/config.json`):

| Var | Default |
|---|---|
| `RPC_URL` | Sepolia public RPC |
| `MANAGER_ADDRESS` | Sepolia DKGManager |
| `CHAIN_ID` | 11155111 |
| `CHAIN_NAME` | sepolia |
| `REGISTRY_ADDRESS` | (auto-derived) |
| `START_BLOCK` | (none) |
| `DEPLOY_BLOCK` | `0` |

`DEPLOY_BLOCK` is the block the `DKGManager` was deployed at, and it is the
floor of every historical log scan the explorer runs — the operator
leaderboard, the per-epoch decryption pipeline, the registry roster. Set it
and those views cost a handful of `eth_getLogs` calls; leave it at `0` and the
SDK falls back to scanning a recent-block window instead, which is fast but
only sees recent history.

## Demo mode

Append `?demo=1` to any URL (or build with `VITE_DEMO=1`) and the whole app
runs from the synthetic fixture with no RPC at all: 300 operators, 64-member
committees, applications with partials arriving in waves. `useRuntimeConfig()`
exposes it as `config.demo`; pages must read it from there rather than parsing
the URL themselves. Demo mode also survives a missing `/config.json`.

## Build

```sh
pnpm build        # → ./dist
pnpm preview      # serves dist on a local port for sanity-checking
```

## Quality checks

```sh
pnpm lint         # tsc --noEmit + eslint
pnpm format       # prettier --write
pnpm test         # vitest
```

## Docker

`ui/Dockerfile` is a build-only target — the final image carries just
the static `dist/` at `/usr/share/nginx/html`, no nginx runtime. Chain
config is **baked in at build time** via `--build-arg`s.

```sh
# Build a Sepolia-targeted bundle (defaults).
docker build -f ui/Dockerfile -t davinci-dkg-ui:sepolia ..

# Build a different deployment by overriding the args.
docker build -f ui/Dockerfile \
  --build-arg RPC_URL=http://host.docker.internal:8545 \
  --build-arg MANAGER_ADDRESS=0x... \
  --build-arg CHAIN_ID=31337 \
  --build-arg CHAIN_NAME=anvil \
  -t davinci-dkg-ui:anvil ..
```

The image is meant to be consumed by a static-site host (DigitalOcean
App Platform spec at `.do/davinci-dkg-ui.yaml`) or extracted and served
manually:

```sh
docker create --name extract davinci-dkg-ui:sepolia
docker cp extract:/usr/share/nginx/html ./dist
docker rm extract
# Then serve ./dist with anything (nginx, Caddy, S3, Cloudflare R2, …).
```

## docker-compose (self-hosted serve)

The compose `ui` service skips the custom image entirely — it bind-mounts
the locally-built `ui/dist` into stock `nginx:alpine`. Build the dist
once with the chain args you want, then bring up the service:

```sh
make ui-build \
  RPC_URL=http://127.0.0.1:8545 \
  MANAGER_ADDRESS=0x... CHAIN_ID=31337 CHAIN_NAME=anvil

docker compose --profile ui up                       # UI alone, on :8082
docker compose --profile node --profile ui up        # node + UI together
```

The UI listens on `${DAVINCI_DKG_UI_PORT:-8082}`. The node service does
not expose any HTTP — see the root `docker-compose.yml` for details.

## DigitalOcean App Platform

```sh
doctl apps create --spec ui/.do/davinci-dkg-ui.yaml
```

Edit the `BUILD_TIME` env values in the spec to retarget a different
chain. App Platform builds the Dockerfile per push, reads the static
files out of `output_dir: /usr/share/nginx/html`, and serves them from
its edge.

## Layout

```
src/
├── main.tsx                entry: fonts, styles, <App/>
├── App.tsx                 provider tree (Config → Wagmi → Query → RainbowKit → Tooltip → Router)
├── styles/index.css        Tailwind v4 entry; imports design/theme.css + design/variables.css
├── config/                 runtime config loader, <ConfigProvider>, useRuntimeConfig()
├── app/                    shell: TopBar, ChainPill, WalletButton, GlobalSearch, Footer, wagmi config
├── routes/                 paths.ts (URL table) + router.tsx (lazy route elements)
├── pages/                  one file per route
├── kit/                    design-system primitives (Button, Card, Table, …)
│   └── charts/             hand-rolled SVG charts + their tested maths
├── hooks/                  cross-cutting hooks (copy, latest block, measured width)
├── lib/                    pure helpers (format, explorer URLs, operator stats, cn)
├── indexer/                in-browser event indexer + selectors
├── fixtures/               synthetic network for demo mode and tests
└── data/                   data-source hooks the pages consume
```

Path aliases: `~app/*`, `~config/*`, `~data/*`, `~fixtures/*`, `~hooks/*`,
`~indexer/*`, `~kit` (and `~kit/*`), `~lib/*`, `~pages/*`, `~routes/*` —
defined in `tsconfig.paths.json`, resolved by `vite-tsconfig-paths`. Always
prefer aliases over relative imports.

Pages import primitives from `~kit` only, never from a component file
directly. `/kit` renders every primitive and chart with sample data — open it
after any change to the design system.

## Design rules

The full statement is `design/preview.html`; the short version:

- Obsidian `#050507` canvas, carbon `#101010` cards, onyx `#1a1a1a` hover,
  charcoal `#3d3a39` hairlines. Borders and glows, never drop shadows.
- Emerald `#00d992` is the *only* accent: primary action, active state, brand,
  and the live/ok semantic. Beyond it there is exactly one amber and one red.
- Inter for text, JetBrains Mono for every address, hash and number in a
  table. Labels are 13 px uppercase with 0.1 em tracking (`.label-caps`).
- 4 px spacing grid, 4–8 px radii, dense rows.
- Any list that can exceed ~50 rows is virtualised; wide panels scroll inside
  themselves, the page never scrolls sideways.
- Every chart has a legend, a skeleton and a tooltip, and its maths lives in
  `kit/charts/scale.ts` / `colors.ts` with unit tests.
