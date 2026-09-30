# Running a node

A node registers its operator in `DKGRegistry`, enters the lottery of every epoch created after
its registration, and when drawn deals its contribution, takes its turn at finalization and
decrypts the committee's ciphertexts. With no network settings it joins the Gnosis Chain
deployment, which anyone may join.

## Requirements

- 2 CPU cores and 4 GB of RAM; 8 GB is comfortable. Proving keys are loaded for a proof and
  released afterwards, so a node sits well below 1 GB at rest and peaks at about 3 GB while it
  proves its contribution. More cores shorten the proofs. `GOMEMLIMIT` (for example
  `GOMEMLIMIT=2500MiB`) trades some CPU for a lower peak.
- About 1.5 GB of disk for the circuit artifacts and the node state.
- An operator key funded with the chain's native currency (xDAI on Gnosis Chain) for
  registration, contributions, finalizations and decryptions. Gas per call is in
  [BENCHMARKS.md](../BENCHMARKS.md#gas).
- At least two JSON-RPC endpoints for a long-lived node. The Gnosis preset has three public
  endpoints built in; the node rotates off endpoints that are rate-limited or unreachable.

## Installing

**Docker Compose** (recommended). The `node` profile runs the published image
`ghcr.io/vocdoni/davinci-dkg` and a Watchtower container that updates it:

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg
cp .env.example .env        # set DAVINCI_DKG_PRIVKEY
docker compose --profile node up -d
docker compose --profile node logs -f node
```

State lives in the `run` volume and the artifacts in `~/.davinci/artifacts` on the host.

**Binary.** Every [release](https://github.com/vocdoni/davinci-dkg/releases) has static Linux
builds for amd64 and arm64:

```bash
DAVINCI_DKG_PRIVKEY=0x... ./davinci-dkg-node
```

**From source** (Go 1.25+): `go build -o davinci-dkg-node ./cmd/davinci-dkg-node`.

## Configuration

Every flag has an environment variable: prefix `DAVINCI_DKG_`, upper case, dots and dashes
replaced by underscores (`--web3.rpc` is `DAVINCI_DKG_WEB3_RPC`). The compose file reads them from
`.env`.

| Flag | Default | Description |
|---|---|---|
| `--privkey` | | Operator key (hex). Without it the node starts idle |
| `--network` | `gnosis` | Built-in deployment: `gnosis`, or `sepolia` (alias `sep`), which has no public endpoints built in |
| `--manager` | | `DKGManager` of any other deployment. Overrides the network and skips the chain id check |
| `--web3.rpc` | the network's public endpoints | Comma-separated JSON-RPC endpoints; `http://127.0.0.1:8545` with `--manager` |
| `--web3.network` | `localhost` | Display name of a custom deployment in the logs |
| `--web3.gasMultiplier` | `1.2` | Headroom applied to gas estimates |
| `--datadir` | `~/.davinci-dkg` | State directory |
| `--poll-interval` | `5s` | How often the node polls the chain |
| `--auto-create-epochs` | `true` | Create the next epoch when it is due and abort a dead newest epoch |
| `--auto-create-jitter` | `12s` | Maximum random delay before an automatic create or abort |
| `--decrypt-lookback-blocks` | `50400` | On start, how far back to scan for ciphertexts still awaiting decryption |
| `--epoch-policy.committee-size` | `0` | Committee size when this node creates an epoch; `0` derives it from the registry |
| `--epoch-policy.threshold` | `0` | Threshold for a fixed committee size |
| `--epoch-policy.min-valid-contributions` | `0` | Minimum accepted contributions for a fixed committee size |
| `--epoch-policy.lottery-alpha-bps` | `15000` | Lottery oversubscription in basis points (10000 = 1.0) |
| `--log.level` | `info` | `debug`, `info`, `warn` or `error` |
| `--log.output` | `stdout` | `stdout`, `stderr` or a file path |

`DAVINCI_DKG_ARTIFACTS_DIR` (environment only, default `~/.davinci/artifacts`) sets where the
circuit artifacts are cached.

With the derived policy, an epoch created by a node has a committee of three quarters of the
active operators (at most 32), a majority threshold and at least two thirds of the committee as
minimum valid contributions, all raised to the deployment's floors.

### Networks

- **Default**: the Gnosis Chain deployment, its public endpoints and chain id 100.
- `--network sepolia`: the Sepolia testnet. Set `--web3.rpc` as well.
- `--manager 0x... --web3.rpc ...`: any other deployment, on any chain.

Only the manager address is configured; the registry, the app manager and the verifiers are read
from it. The node refuses endpoints that serve a different chain than the selected network.

## What the node does

On first start it:

1. derives its BabyJubJub key from the operator key and registers it in `DKGRegistry` (skipped
   when already registered and active);
2. downloads the pinned circuit artifacts from the GitHub release named in
   `config/circuit_artifacts.go` (about 1 GB) and verifies every file against the hashes built
   into the binary;
3. logs a startup banner with the chain head, the registry statistics and its own registry row,
   then polls the manager.

Then, for every epoch:

- it claims a committee slot if the lottery admits it and submits its contribution during key
  assembly;
- when the epoch can be finalized, the nodes take turns in a seed-derived order: the first one
  proves and submits `finalizeEpoch`, the others see the epoch go `Live` and stop;
- for every ciphertext of an epoch it belongs to, it posts its partial decryption in waves of
  `t` members and combines when its turn comes;
- with `--auto-create-epochs`, it creates the next epoch when the cadence allows it or the
  newest pool is spent, and aborts a newest epoch that can no longer progress.

The node sends a heartbeat to stay active in the registry. A node that stays off longer than the
registry's inactivity window can be marked inactive by anyone; starting it again reactivates it.
It serves no HTTP. Watch it through its logs, or through the [explorer](../ui/README.md).

## Data directory

State is kept under `<datadir>/<chainid>-<manager>`, one directory per deployment: the cache of
contribution calldata and the list of tainted applications. A node moved to another deployment
starts from an empty directory and finds its old state again when moved back.

## Upgrades

The compose file runs the `latest` image under Watchtower. `latest` moves on stable releases only
(`vX.Y.Z`, never a release candidate). When a release moves the Gnosis preset to a new
deployment, a node without `--network` or `--manager` follows it on update and starts from a
fresh state directory; a node started with `--manager` stays where it is. To upgrade by hand, pin
`DAVINCI_DKG_TAG` to a version in `.env` or remove the `watchtower` service.

## Running on Railway

[Railway](https://railway.com) can run a node from the public image.
`scripts/railway-deploy-node.sh` creates the service through Railway's GraphQL API and
`scripts/railway-node-status.sh` shows its deployment status and log tail. They need `curl` and
`python3`.

The workspace must be on a plan that allows at least 4 GB of memory per service (Hobby or
higher): under a 1 GB cap the node is killed at its first proof. Volumes keep the size limit of
the plan they were created under, so create the volume after upgrading; the artifacts need more
than 1 GB.

1. Create a workspace token and store it in a file outside the repository (`/railway-api-key`
   is git-ignored).
2. Store the operator key as `[{"address": "0x...", "private_key": "0x..."}]` in a file with
   mode `0600`, for example the output of `cast wallet new --json`, and fund the address.
3. Create a project in the Railway dashboard, or through the API with `projectCreate`, and note
   the project id and its `production` environment id.
4. Deploy each node as its own service. Do not use Railway replicas: they share variables, so
   they would share the operator key.

```bash
export RAILWAY_TOKEN_FILE=railway-api-key
export RAILWAY_PROJECT_ID=<project id> RAILWAY_ENVIRONMENT_ID=<environment id>
NODE_NAME=dkg-node1 NODE_KEY_FILE=~/keys/node1.json scripts/railway-deploy-node.sh
scripts/railway-node-status.sh <service id> 40
```

The script creates the service with the node's variables, a volume mounted at `/app/run` for the
artifacts and the state, and an always-restart policy, then deploys it. `IMAGE`, `NETWORK`, `RPC`,
`POLL_INTERVAL` and `MOUNT_PATH` override its defaults; Sepolia needs `NETWORK=sepolia` and `RPC`.
The key is sent only in the request body, never as an argument.

A healthy first start logs `circuit artifacts verified` for the four circuits, `self: registry row`
once the registration is mined, and `node running`. A node that restarts every few seconds, or
right after `contribution assignment`, is being killed by the memory limit.
