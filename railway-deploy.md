# Deploying davinci-dkg nodes on Railway

Railway (https://railway.com) runs a node as a container from the public image
`ghcr.io/vocdoni/davinci-dkg:latest`, bills per second on the memory and CPU
the container actually uses (about $11 a month per v0.7 node, see
`BENCHMARKS.md`), and is fully driven through its GraphQL API, which is what
`scripts/railway-deploy-node.sh` uses. This is how the Sepolia fleet's nodes 9
and 10 were deployed on 2026-09-08 and how to add more.

## What you need

- A Railway **workspace token** (Workspace settings → Tokens) in a file that
  never enters git. The repository ignores `/railway-api-key`; keep it there.
  A workspace token cannot answer `me` (that is expected) but can create and
  manage projects in its workspace.
- An operator key per node in the fleet format
  `[{"address": "0x…", "private_key": "0x…"}]`, e.g. generated with
  `cast wallet new --json` and stored under `~/.davinci-dkg-sepolia/nodeN.json`
  (mode 0600). Fund the address with a little Sepolia ETH before the first
  start: registration is one transaction, a node spends about 0.02 ETH a day
  under the public bot's load.
- `curl` and `python3`.

The scripts read the token from `RAILWAY_TOKEN_FILE` and send it only in the
`Authorization` header; the operator key is read from the key file into the
request body through a 0600 temp file, never passed as an argument.

## Plan requirement

The workspace must be on the **Hobby** plan or higher ($5 a month, offset by
$5 of included usage; 48 GB of memory and 48 vCPU per service). A Trial or
Free workspace caps every service at 2 vCPU and 1 GB, and a node needs about
2.9 GB while it proves a contribution (startup itself is light since v0.7.1:
artifacts are stream-verified, nothing is compiled, 0.17 GB measured). Under the cap a v0.7.0
container was killed during `building constraint builder` and restarted every
few seconds; a v0.7.1 container would start, register and claim a slot and
then be killed at its first proof, which hurts the committee, so keep the
deployments stopped (`deploymentStop`) until the plan allows the memory. The plan is upgraded in
the Railway dashboard (a card is required); it cannot be changed through the
API. `serviceInstanceLimits(serviceId, environmentId)` shows the effective
`memoryBytes`; after an upgrade run `serviceInstanceRedeploy` on each node so
the new limits apply. **Volumes keep the size cap of the plan they were
created under**: a volume made on Trial is 500 MB for good, too small for the
1.1 GB of artifacts (`no space left on device` while downloading), and the API
cannot grow it. Delete it (`volumeDelete`), create a new one (5 GB on Hobby)
and redeploy; `volumeCreate` refuses while the old one is still attached, so
delete first and check the project's volumes if a create fails half-way.

## One-time: create the project

```bash
export RAILWAY_TOKEN_FILE=railway-api-key
curl -sS https://backboard.railway.com/graphql/v2 \
  -H "Authorization: Bearer $(tr -d '[:space:]' < $RAILWAY_TOKEN_FILE)" \
  -H 'Content-Type: application/json' \
  --data '{"query":"mutation { projectCreate(input: { name: \"davinci-dkg-sepolia\" }) { id environments { edges { node { id name } } } } }"}'
```

Keep the project id and the `production` environment id. The Sepolia fleet's
project is `a2242bc7-5e58-4fda-a326-86464d4670cd`, environment
`f1a9f7a4-ed36-4bca-9369-06e290fdd0b0`.

## Per node: deploy

```bash
export RAILWAY_TOKEN_FILE=railway-api-key
export RAILWAY_PROJECT_ID=a2242bc7-5e58-4fda-a326-86464d4670cd
export RAILWAY_ENVIRONMENT_ID=f1a9f7a4-ed36-4bca-9369-06e290fdd0b0
NODE_NAME=dkg-node11 NODE_KEY_FILE=~/.davinci-dkg-sepolia/node11.json scripts/railway-deploy-node.sh
```

The script does four API calls, in this order, and prints the ids it gets back:

1. `serviceCreate` with `source.image = ghcr.io/vocdoni/davinci-dkg:latest` and
   the node's variables, so the very first start already has them:
   `DAVINCI_DKG_NETWORK=sepolia`, `DAVINCI_DKG_PRIVKEY`, `DAVINCI_DKG_WEB3_RPC`
   (three public endpoints; the node rotates off rate-limited ones),
   `DAVINCI_DKG_POLL_INTERVAL=30s`, `DAVINCI_DKG_DATADIR=/app/run/data`,
   `DAVINCI_DKG_ARTIFACTS_DIR=/app/run/artifacts`.
2. `volumeCreate` mounted at `/app/run`: the pinned circuit artifacts
   (`circuits-v6`, about 1.0 GB, downloaded from the GitHub release and
   hash-checked on first start) and the node's caches survive redeploys.
3. `serviceInstanceUpdate` with `restartPolicyType: ALWAYS`.
4. `serviceInstanceDeployV2` to deploy with all of the above in place.

Overridable through the environment: `IMAGE`, `NETWORK`, `RPC`,
`POLL_INTERVAL`, `MOUNT_PATH`. On another network set `NETWORK` to its preset
name, or add `DAVINCI_DKG_MANAGER` to the variables in the script.

Each node is its own service. Do not use Railway replicas for nodes: replicas
share variables, so they would share the operator key.

## Verify

```bash
scripts/railway-node-status.sh <serviceId> 40
```

prints the latest deployment's status (`BUILDING`, `DEPLOYING`, `SUCCESS`,
`CRASHED`, …) and its last log lines. A healthy first start shows the four
`circuit artifacts loaded` lines after the download, `self: registry row` once
the registration transaction is mined, and then `node running`. On chain, the
registry's `activeCount()` goes up by one per node. A node that logs
`starting davinci-dkg-node` every 10–20 s without ever reaching `node
running`, or restarting right after `contribution assignment`, is being killed
by the memory cap (see the plan requirement above).

## Operate

- **New image**: Railway watches the image tag and offers an update; to force
  it, `serviceInstanceRedeploy(serviceId, environmentId)` pulls `:latest`
  again. Pin `IMAGE` to a version tag (`ghcr.io/vocdoni/davinci-dkg:v0.7.0`)
  for a fleet that must not move on its own.
- **Logs and status**: the status script, or the deployments/deploymentLogs
  queries it wraps.
- **Remove a node**: `serviceDelete(id)` then `volumeDelete(volumeId)`; a
  stopped service costs nothing but its volume does until deleted.
- **Cost**: usage-billed, memory $10 per GB-month, CPU $20 per vCPU-month,
  volume $0.15 per GB-month. A v0.7 node sits at 0.2–0.7 GB at rest, peaks at
  2.9 GB while proving, and averages 0.05 vCPU, about $11 a month; the Hobby
  plan's $5 fee is offset by its $5 of included usage and allows 50 services.
- **API limits**: 1,000 requests per hour and 10 per second on Hobby (10,000
  and 50 on Pro). The scripts make four calls per node.

## Fleet record

| Node | Service id | Volume id | Operator |
|---|---|---|---|
| dkg-node1 | `7f2d0623-f369-4158-b568-5f4aa3d9c4e7` | `dacac9f5-a70c-4705-88b7-92a389467a16` (5 GB) | `0x6aFbc1474dA64a25cD12842C0B2410053a245DD5` |
| dkg-node2 | `5f1bcff1-e82c-402b-8a60-8ddc7f2035fb` | `9dc37e56-175c-4b01-9281-2b57e73488e3` (5 GB) | `0x9D417cB81380B09546e8EbA8d1d215D8C98F81FF` |
| dkg-node3 | `8c94159e-70fe-4cb0-876d-69a4a3d6775d` | `4eb4e660-dfa7-4ae2-8f72-d782f5a15f33` (5 GB) | `0x51084b010395C62b5294599e06617bAc22d55F36` |
| dkg-node4 | `c8aa6050-0dde-4ce8-a395-b238321cbea0` | `81be8d2f-9728-40df-812b-d440d922b3cf` (5 GB) | `0xb8860186663Aa9b282A3aa873BAccb438Defa76B` |
| dkg-node5 | `d1acc7d3-5227-4e73-bcf8-8dba4af03258` | `fb8af313-d013-4ed7-bac2-8b57ed9159f4` (5 GB) | `0x7736a6c997DFef24463B42FB29234d6539e83cCC` |
| dkg-node6 | `4c71c556-bc19-4237-b86f-c96d3ed2ca0d` | `56f4d802-44fc-4792-b5c7-1a066a3369cf` (5 GB) | `0xE0D26f0C10a8091EAFed58c8c642714aDa8bCDfA` |
| dkg-node7 | `c8a78a6e-fc51-4a5b-b133-6c6f6cfec266` | `6694e6ac-0085-451b-aa3a-b9159a545737` (5 GB) | `0x8EDd1838db9fc9A833B759Ee0C78242273dA0459` |
| dkg-node8 | `66e1f124-465f-43d7-b485-de96454335d8` | `65a055be-d258-4fde-9c5b-8a2acc93ce1d` (5 GB) | `0x647F35C2f7F308c75c52b2c64b5e7615068a4a5b` |
| dkg-node9 | `b7c719f1-8a87-4de2-8033-b62bfea71c15` | `386e62ee-6a80-40f6-a3f5-cc665955cd1a` (5 GB) | `0xf4FE1328f6a9C3391bFf4024d0348DdF188eD35c` |
| dkg-node10 | `9bc7a731-1902-4bcf-b2fa-14052fd07cc0` | `c8413cf6-83e2-44b0-a713-4d4be9853bc0` (5 GB) | `0x31e162dD9c3Bc94a903a90E710d8A3577F14203b` |

Operator keys: `~/.davinci-dkg-sepolia/node1.json` … `node10.json` on the
workstation (never in git).

Since 2026-09-09 13:27 UTC the whole public Sepolia fleet (nodes 1–10, v0.8.0 image via `:latest`,
Sepolia v6 deployment at manager `0xc73b7a868eca6ac7e3e647e2665aa16a793cf551`) runs on Railway; the
former workstation/z5/z7 compose fleets are stopped. Fresh volumes download the 1.0 GB of `circuits-v6`
artifacts on first start; with eight nodes starting at once GitHub answered a few downloads with a
non-200 status, the node exited (`unexpected http status`) and the ALWAYS restart policy retried until
it succeeded, within a minute.
