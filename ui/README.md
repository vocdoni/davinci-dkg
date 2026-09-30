# DAVINCI DKG Explorer

Web explorer for a [DAVINCI DKG](../README.md) deployment: epochs, committees, operators,
applications and decryptions, a playground that walks through the organizer's steps, and the
operator and SDK documentation. It is a static single-page app that reads the chain through a
JSON-RPC endpoint; there is no backend.

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](../LICENSE)

## Overview

On first load the explorer indexes every event of the deployment's contracts from the manager's
deployment block, keeps the result in the browser's IndexedDB and then follows the chain head.
Contract state that events do not carry (epoch records, pool keys, plaintexts) is read with
batched calls. A wallet (injected, or WalletConnect when configured) is only needed for the
playground and the organizer's reveal.

Built with Vite, React, TypeScript, Tailwind CSS, wagmi and viem, and the
[SDK](../sdk/README.md) through `link:../sdk`.

## Quick start

From the repository root, with Node.js 20 and pnpm 10:

```bash
make ui-dev          # builds the SDK, installs dependencies and serves on http://localhost:5174
```

The committed `public/config.json` targets the Gnosis Chain deployment. Append `?demo=1` to any
URL to run from a synthetic network with no RPC at all.

## Usage

### Configuration

The explorer loads `/config.json` at startup, so a build can be pointed at another deployment by
replacing that one file. `make ui-dev`, `make ui-build` and `make ui-config` render it from these
variables; a variable left out keeps the value already in the file.

| Variable | `config.json` key | Default |
|---|---|---|
| `RPC_URL` | `rpcUrl` | `https://gnosis-rpc.publicnode.com` |
| `MANAGER_ADDRESS` | `managerAddress` | Gnosis `DKGManager` |
| `CHAIN_ID` | `chainId` | `100` |
| `CHAIN_NAME` | `chainName` | `gnosis` |
| `DEPLOY_BLOCK` | `deployBlock` | `48483860` |
| `EXPLORER_URL` | `explorerUrl` | `https://gnosisscan.io` |

`DEPLOY_BLOCK` is where the historical scan starts. Set it to the manager's deployment block;
with `0` the explorer searches for that block with about 25 `eth_getCode` calls, which needs an
RPC endpoint that serves historical state. Two build-time variables are read by Vite:
`VITE_WALLETCONNECT_PROJECT_ID` enables WalletConnect, and `VITE_DEMO=1` builds a demo-only bundle.

The Sepolia testnet, for example:

```bash
make ui-dev RPC_URL=https://ethereum-sepolia-rpc.publicnode.com \
  MANAGER_ADDRESS=0xc73b7a868eca6ac7e3e647e2665aa16a793cf551 \
  CHAIN_ID=11155111 CHAIN_NAME=sepolia DEPLOY_BLOCK=11668198 EXPLORER_URL=https://sepolia.etherscan.io
```

### Hosting

The build output in `ui/dist` is plain static files; serve it from any web server or CDN with a
fallback to `index.html` for client-side routes.

```bash
make ui-build RPC_URL=... MANAGER_ADDRESS=... CHAIN_ID=... CHAIN_NAME=... DEPLOY_BLOCK=...
docker compose --profile ui up -d    # serves ui/dist with nginx on port 8082 (DAVINCI_DKG_UI_PORT)
```

`ui/Dockerfile` builds the same bundle into an image that only carries the files, at
`/usr/share/nginx/html`; release builds are published as `ghcr.io/vocdoni/davinci-dkg-ui` with
the Gnosis configuration. Build it from the repository root, overriding any configuration
variable with `--build-arg`:

```bash
docker build -f ui/Dockerfile -t davinci-dkg-ui --build-arg RPC_URL=https://... .
docker create --name ui davinci-dkg-ui && docker cp ui:/usr/share/nginx/html ./dist && docker rm ui
```

`ui/.do/davinci-dkg-ui.yaml` is a DigitalOcean App Platform spec that builds this Dockerfile on
every push to `main` and serves the files (`doctl apps create --spec ui/.do/davinci-dkg-ui.yaml`);
edit its `BUILD_TIME` values to target another deployment.

## Documentation

- [EXPLORER.md](EXPLORER.md): routes, data layer, demo fixture and scale requirements.
- [src/data/README.md](src/data/README.md): the hooks pages use to read the indexer.
- `design/`: design tokens and theme.

## Development

```bash
cd ui
pnpm install
pnpm dev             # http://localhost:5174, bound on all interfaces
pnpm lint            # tsc --noEmit and eslint
pnpm test            # vitest
pnpm format          # prettier
pnpm build           # builds the SDK, then the bundle into dist/
```

Source lives in `src/`: `pages/` (one module per route), `kit/` (design-system components and
SVG charts), `indexer/` (the in-browser event indexer), `data/` (hooks over it), `fixtures/` (the
synthetic network for demo mode and tests), `config/`, `app/` (shell and wallet) and `routes/`.
Import through the path aliases in `tsconfig.paths.json` (`~kit`, `~data/*`, ...), and build
links from `src/routes/paths.ts`. The `/kit` route renders every component and chart with sample
data; check it after changing the design system.

Design rules: dark surfaces with hairline borders and no drop shadows, one emerald accent plus
one amber and one red for status, Inter for text and JetBrains Mono for addresses, hashes and
numbers, a 4 px grid. Lists that can exceed about 50 rows are virtualised, and wide panels scroll
inside themselves rather than the page.

## License

[GNU Affero General Public License v3.0](../LICENSE).
