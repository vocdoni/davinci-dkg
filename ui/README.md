# davinci-dkg UI

Standalone Vite + React + TypeScript SPA for the davinci-dkg explorer,
playground, and operator/SDK documentation. Modelled on
[`vocdoni/vocdoni-app`](https://github.com/vocdoni/vocdoni-app)'s structure
and conventions.

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

```sh
# Build & run standalone (against the default Sepolia config)
docker build -f ui/Dockerfile -t davinci-dkg-ui ..
docker run -p 8082:80 davinci-dkg-ui

# Override the chain target at run time:
docker run -p 8082:80 \
  -e DAVINCI_DKG_RPC_URL=http://host.docker.internal:8545 \
  -e DAVINCI_DKG_MANAGER_ADDRESS=0x... \
  -e DAVINCI_DKG_CHAIN_ID=31337 \
  -e DAVINCI_DKG_CHAIN_NAME=anvil \
  davinci-dkg-ui
```

The container's entrypoint renders `/usr/share/nginx/html/config.json`
from environment variables before starting nginx, so a single image
targets any deployment.

## docker-compose

```sh
# UI alone (against external RPC, served on :8082)
docker compose --profile ui up

# Node + UI together (independent services)
docker compose --profile node --profile ui up
```

The UI listens on `${DAVINCI_DKG_UI_PORT:-8082}`. The node service does
not expose any HTTP — see the root `docker-compose.yml` for details.

## Layout

```
src/
├── App.tsx                 provider tree (Theme → Debug → Config → Wagmi → Query → Router)
├── main.tsx                entry
├── router/                 createBrowserRouter + Routes constant + lazy elements
├── elements/               route-level pages (kebab-case)
├── components/             reusable, grouped by domain (PascalCase)
├── providers/              ConfigProvider, DebugModeProvider
├── queries/                react-query hooks + central QueryKeys
├── hooks/                  cross-cutting hooks
├── theme/                  Chakra v3 system + color mode
├── lib/                    pure helpers (format, error-report, debug, wagmi config)
├── constants/              chains, polling cadences, route table
└── types/                  shared TS types
```

Path aliases: `~components/*`, `~elements/*`, `~queries/*`, `~theme/*`,
`~hooks/*`, `~lib/*`, `~providers/*`, `~constants/*`, `~router/*`,
`~types/*` — defined in `tsconfig.paths.json`, resolved by
`vite-tsconfig-paths`. Always prefer aliases over relative imports.

## UX rules

- Every page renders a plain-English summary by default. Hashes are
  truncated, block deltas are durations, status is a badge.
- Technical detail (raw hex, BigInt coords, raw event args) lives inside
  `<DetailDisclosure>` blocks that auto-expand when **debug mode** is on.
- Errors always offer `<ErrorReportButton>` so users can paste a
  ready-made markdown blob into a GitHub issue.
