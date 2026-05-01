# DAVINCI DKG

**Non-Interactive Distributed Key Generation using Blockchain and ZK Proofs**

`davinci-dkg` is the Go implementation of the NI-DKG protocol described in the paper
[*NI-DKG: Non-Interactive Distributed Key Generation using Blockchain and ZK Proofs*](https://eprint.iacr.org/2026/552).
It provides the node service, cryptographic primitives, zk-SNARK circuits, and Solidity smart contracts
for threshold key generation and threshold decryption on EVM-compatible chains.

The protocol is designed as the key-generation and threshold-decryption layer for
[DAVINCI](https://davinci.vote) voting system.

## Table of Contents

- [Run a DKG Node](#run-a-dkg-node)
  - [What you need before you start](#what-you-need-before-you-start)
  - [Configure the node](#configure-the-node)
  - [Option A — Docker Compose (recommended)](#option-a--docker-compose-recommended)
  - [Option B — Download a release binary](#option-b--download-a-release-binary)
  - [Option C — Build from source](#option-c--build-from-source)
  - [Start the node](#start-the-node)
  - [Verify you joined the network](#verify-you-joined-the-network)
  - [Operational notes](#operational-notes)
- [Overview](#overview)
- [Mathematical Background](#mathematical-background)
  - [Setting](#setting)
  - [Shamir Secret Sharing](#shamir-secret-sharing)
  - [Feldman Verifiable Secret Sharing](#feldman-verifiable-secret-sharing)
  - [Hashed ElGamal Share Encryption](#hashed-elgamal-share-encryption)
  - [DKG Protocol](#dkg-protocol)
  - [Trustless Lottery Committee Selection](#trustless-lottery-committee-selection)
  - [Threshold Decryption](#threshold-decryption)
  - [Public-Input Compression (BRLC)](#public-input-compression-brlc)
- [ZK-SNARK Circuits](#zk-snark-circuits)
  - [Contribution Circuit](#contribution-circuit-dkg-phase-3)
  - [Finalize Circuit](#finalize-circuit-dkg-phase-4)
  - [PartialDecrypt Circuit](#partialdecrypt-circuit-decryption-phase-2)
  - [DecryptCombine Circuit](#decryptcombine-circuit-decryption-phase-3)
- [Smart Contracts](#smart-contracts)
  - [DKGRegistry](#dkgregistry)
  - [DKGManager](#dkgmanager)
  - [DKGAppManager](#dkgappmanager)
  - [Per-Application Surface](#per-application-surface)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Build](#build)
  - [Run Tests](#run-tests)
  - [Compile Circuits](#compile-circuits)
  - [Deploy Contracts](#deploy-contracts)
- [Gas Profile](#gas-profile)
- [Switching `MaxN`](#switching-maxn)
- [Local Testnet](#local-testnet)
  - [Start the network](#start-the-network)
  - [Run the scenario](#run-the-scenario)
  - [Configuration](#configuration)
- [Web Explorer](#web-explorer)
  - [What it shows](#what-it-shows)
  - [How it reaches the chain](#how-it-reaches-the-chain)
  - [Running it](#running-it)
- [References](#references)

---

## Run a DKG Node

This is the fastest path to joining an existing davinci-dkg network as a
participant. If you just want to understand the protocol, jump to
[Overview](#overview); if you want to stand up your own network instead of
joining one, see [Local Testnet](#local-testnet) and [Deploy Contracts](#deploy-contracts).

### What you need before you start

1. **A funded EVM account.** The node submits signed transactions — key
   registration, slot claim, contribution, partial decryption, combine
   every time an epoch runs. You need enough native gas on the target
   network to cover those. Gas costs are bounded and documented in the
   [Gas Profile](#gas-profile) section; as a rule of thumb, budget a few
   million gas per epoch you expect to participate in. Testnet gas from a
   faucet is usually enough; mainnet deployments should hold a comfortable
   balance.
2. **The target network's JSON-RPC URL.** Any HTTPS or WSS endpoint that
   speaks the standard Ethereum JSON-RPC will work (Infura, Alchemy, your
   own node, a local Anvil instance, etc.).
3. **The target network.** For well-known deployments (currently **Sepolia**),
   the node already knows the contract addresses — just pass `--network sepolia`
   (or `DAVINCI_DKG_NETWORK=sepolia`) and no further address configuration is
   needed. For a custom or private network you will need the `DKGManager`
   address published alongside that network's announcement; deploy your own
   with [`make solidity-deploy`](#deploy-contracts) if you are bootstrapping.
4. **An operator private key** that controls the funded account. It is
   used only for signing, the node never exports or transmits it.

> the DKG epoch cadence and policy are decided by whoever
> creates epochs on `DKGManager`. As a node operator you react to epochs;
> you don't need to run the orchestrator.

### Configure the node

All three install options below share the same configuration surface: a
`.env` file at the repo root (or next to the binary). The node reads its
settings from environment variables — every CLI flag has an env-var
equivalent, e.g. `--web3.rpc` → `DAVINCI_DKG_WEB3_RPC`,
`--poll-interval` → `DAVINCI_DKG_POLL_INTERVAL`.

```bash
cp .env.example .env
$EDITOR .env
```

At minimum, fill in:

```dotenv
DAVINCI_DKG_WEB3_RPC=https://your-rpc-endpoint
DAVINCI_DKG_PRIVKEY=0x<your-64-hex-char-operator-key>

# For Sepolia — contract addresses are built in:
DAVINCI_DKG_NETWORK=sepolia

# For any other network — supply the DKGManager address explicitly:
# DAVINCI_DKG_MANAGER=0x<DKGManager address>
```

See `.env.example` for the full list and `davinci-dkg-node --help` for
defaults. The node binary itself does not serve any HTTP — it only talks
to the chain via `--web3.rpc`. Use the standalone UI image
(`ghcr.io/vocdoni/davinci-dkg-ui`, see [Web Explorer](#web-explorer)
below) if you want a browser-facing explorer alongside the node.

### Option A — Docker Compose (recommended)

The repo ships a ready-to-run `docker-compose.yml` at the root. It pulls
a prebuilt fully-static Go binary on `debian:bookworm-slim` from
`ghcr.io/vocdoni/davinci-dkg`. This is the recommended path for most
operators: one command, no host-side toolchain, automatic restart on
failure, and automatic image upgrades via Watchtower.

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg

# 1. Configure the node (see the previous section).
cp .env.example .env && $EDITOR .env

# 2. Start the node + Watchtower.
docker compose --profile node up -d

# 3. Tail the logs.
docker compose --profile node logs -f node
```

The compose file defines three profiles:

| Profile | Services                       | Use case                                     |
|---------|--------------------------------|----------------------------------------------|
| `node`  | `node`, `watchtower`           | Long-running DKG node with auto-updates      |
| `ui`    | `ui`                           | Standalone explorer SPA (nginx, port 8082)   |
| `test`  | `unit-test`, `integration-test`| Run the Go test suites inside a container    |

The `node` and `ui` services are independent — combine them with
`docker compose --profile node --profile ui up` to run both side-by-side
on a single host. See [Web Explorer](#web-explorer) below for UI
configuration.

Pin `DAVINCI_DKG_TAG=v0.1.0` in `.env` (or remove the `watchtower`
service) if you want to control upgrades manually.

**Build the image yourself** instead of pulling from `ghcr.io`:

```bash
docker build -t davinci-dkg:local .
DAVINCI_DKG_TAG=local docker compose --profile node up -d
```

### Option B — Download a release binary

Every tagged release publishes fully-static `davinci-dkg-node` binaries
for Linux (amd64 + arm64) on the
[**GitHub Releases**](https://github.com/vocdoni/davinci-dkg/releases) page.

```bash
VERSION=v0.1.0
TARGET=linux-amd64
curl -LO "https://github.com/vocdoni/davinci-dkg/releases/download/${VERSION}/davinci-dkg-${VERSION}-${TARGET}.tar.gz"
curl -LO "https://github.com/vocdoni/davinci-dkg/releases/download/${VERSION}/davinci-dkg-${VERSION}-${TARGET}.tar.gz.sha256"

# Verify the checksum before running anything.
sha256sum -c "davinci-dkg-${VERSION}-${TARGET}.tar.gz.sha256"

tar -xzf "davinci-dkg-${VERSION}-${TARGET}.tar.gz"
cd "davinci-dkg-${VERSION}-${TARGET}"
./davinci-dkg-node --help
```

The binaries are self-contained Go executables — no Node, no pnpm, no
extra build toolchain at runtime. The explorer UI is a separate Docker
image (`ghcr.io/vocdoni/davinci-dkg-ui`); the node binary itself is
UI-blind.

### Option C — Build from source

You will need **Go 1.25+** (and optionally **Foundry** if you want to
rebuild contracts).

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg

make build

# The binary is produced at the repo root:
./davinci-dkg-node --version
./davinci-dkg-node --help
```

To also build the UI bundle locally (for `make ui-dev`, the standalone
container, etc.) you'll additionally need **pnpm** — see `make ui-build`.

### Start the node

First-run behaviour (identical for all three install options):

1. The node derives a BabyJubJub encryption key from your Ethereum
   private key, reads its current registry row, and depending on what it
   finds either calls `registerKey` (first time), `updateKey` (key needs
   rotation or previous row was `INACTIVE`, which auto-reactivates), or
   skips the call entirely (already `ACTIVE` with the right key).
2. Immediately after, the node prints a verbose startup banner: local
   config (identity, RPC, contracts, poll interval),
   on-chain state (chain head, epoch prefix + nonce, `nodeCount`,
   `activeCount`, `INACTIVITY_WINDOW`), and its own registry row
   (`status`, `lastActiveBlock`, `blocksSinceActive`, and the remaining
   liveness budget before reap). Grepping the logs for `self:` is usually
   the fastest way to see whether you joined the network correctly.
3. From then on the node polls `DKGManager` for active epochs and
   reacts to every phase it is eligible for (`claimSlot`, then
   `submitContribution`, then `submitPartialDecryption`, and so on).
4. **Liveness is automatic.** Every poll tick the node refreshes its
   own `lastActiveBlock` if it has drifted past 80% of
   `INACTIVITY_WINDOW` (by sending `heartbeat()`), and if the row has
   been reaped out-of-band it immediately calls `reactivate()` to
   rejoin the active set. You never need to touch the registry
   manually unless you want to rotate your BabyJubJub key.
5. Each phase emits a structured log line (`log.level=info` by
   default). Enable `log.level=debug` if you want the full protocol
   trace.

```bash
# With Docker Compose (Option A):
docker compose --profile node up -d
docker compose --profile node logs -f node

# With a binary or source build (Options B and C):
./davinci-dkg-node

# Sepolia — contract addresses are resolved automatically:
./davinci-dkg-node \
  --web3.rpc=https://your-sepolia-rpc \
  --privkey=0x<your-key> \
  --network=sepolia

# Custom network — supply the DKGManager address explicitly:
./davinci-dkg-node \
  --web3.rpc=https://your-rpc-endpoint \
  --privkey=0x<your-key> \
  --manager=0x<DKGManager> \
  --log.level=info
```

### Verify you joined the network

Three independent checks:

1. **Explorer**: open `http://<your-host>:8081/` in a browser and look at
   the **Registry** tab. Your operator address should appear with a green
   "Active" badge and its BabyJubJub public key.
2. **On-chain**: `cast call $REGISTRY "getNode(address)(address,uint256,uint256,uint8)" $YOUR_ADDR`
   against your RPC should return a non-zero public key.
3. **Logs**: on startup the node prints a banner delimited by
   `==================== davinci-dkg-node startup ====================`
   lines. Inside you should see a `self: registry row` entry with
   `status=ACTIVE` and a recent `lastActiveBlock`, plus a `self: liveness
   budget` entry showing `blocksUntilReap > 0`. A healthy node then
   cycles through poll ticks; `liveness: heartbeat` and `liveness:
   reactivate` entries appear only when the mechanism actually kicks in.

When the next epoch is created you will see `claiming slot`, `submitting
contribution`, and so on. Every phase emits a log line with the epoch ID
so you can trace progress against the explorer.

### Operational notes

- **Gas**: see [Gas Profile](#gas-profile) and [BENCHMARKS.md](BENCHMARKS.md)
  for the per-call breakdown. Your node pays gas only on the phases it
  actually executes.
- **Upgrades**: replace the binary and restart the process. All
  per-epoch in-flight state (claimed epochs, own VSS contributions,
  decryption partials) is held in memory. It is rebuilt automatically
  from calldata and on-chain records when the node restarts — a restart
  mid-epoch is safe.
- **Multiple operators on the same host**: each node needs its own
  private key and data directory. Keep one process per operator to avoid
  confused state.
- **Key loss = slot loss**: losing your operator key (or its funded
  balance) means you can no longer participate until you register a new
  address. The chain still advances without you, the DKG is `t`-of-`n`.
- **Observability**: the structured logs are machine-readable
  (`zerolog`), so piping the node into any log aggregator (Loki, Datadog,
  Grafana Cloud…) is straightforward.

If anything goes wrong, check that (a) your RPC endpoint is reachable,
(b) your account has enough gas, (c) the `REGISTRY` and `MANAGER`
addresses match the network you're talking to, and (d) your clock is not
in the past, stale block numbers will make the node wait forever for a
seed that is already resolved.

---

## Overview

The protocol eliminates interactive complaint procedures through ZK proofs.
Every participant proves the correctness of their contribution in a single Groth16 zk-SNARK submitted
alongside their data. The smart contract rejects any invalid submission at transaction time — no
dispute phase exists.

Each participant submits exactly one transaction per phase. Any observer can verify correctness by
inspecting the on-chain record. The secret sharing is `t`-of-`n`: any `t` participants can decrypt
or reconstruct. The cryptographic stack uses Groth16 on BN254 with Poseidon1 hashing, matching the
rest of the DAVINCI system.

---

## Mathematical Background

### Setting

Let `𝔾` be the BabyJubJub twisted Edwards curve over the BN254 scalar field `𝔽_q`, written additively,
with generator `G`. This is the same group used by the rest of the DAVINCI stack.
Scalars are elements of `𝔽_q`.

BabyJubJub parameters:
- Curve equation: `a·x² + y² = 1 + d·x²·y²` (twisted Edwards, reduced form)
- Base field: BN254 scalar field, `q ≈ 2²⁵⁴`
- Subgroup order: `l = 2736030358979909402780800718157159386076813972158567259200215660948447373041`
- Generator `G = (Gx, Gy)` — the standard BabyJubJub generator point

### Shamir Secret Sharing

A `(t, n)` threshold scheme distributes a secret `σ ∈ 𝔽_q` among `n` parties such that any `t` shares
suffice to reconstruct it and any `t−1` shares reveal nothing.

Each dealer `i` generates a random polynomial of degree `t−1`:
```
f_i(x) = Σ_{k=0}^{t-1} a_{i,k} · x^k   (mod q)
```
where `a_{i,0}` is the secret contributed by dealer `i`.
The share sent to participant `j` is `s_i(j) = f_i(j)`.

Reconstruction uses Lagrange interpolation at `x = 0`:
```
F(0) = Σ_{k ∈ [t]} λ_k · d_{x_k}
```
where `λ_k = Π_{u ∈ Q\{x_k}} u/(u − x_k)` are Lagrange coefficients over `{x_1, …, x_t} ⊆ [n]`.

### Feldman Verifiable Secret Sharing

Feldman VSS extends Shamir by publishing group-element commitments to the polynomial coefficients:
```
C_i(k) = a_{i,k} · G   for k ∈ {0, …, t−1}
```

Any participant `j` can verify their share `s_i(j)` against the published commitments:
```
s_i(j) · G  =?=  Σ_{k=0}^{t-1} j^k · C_i(k)
```
This is the **Feldman verification equation** — the core statement proved by the contribution circuit.

### Hashed ElGamal Share Encryption

Shares are published on-chain in encrypted form so only the intended recipient can read them.
The encryption of share `s_i(j)` for recipient `j` with public key `pub_j = sk_j · G` is:

```
R_i(j)    = r_{i,j} · G                                     (ephemeral key)
σ_i(j)    = s_i(j) + H_share(eid, i, j, r_{i,j} · pub_j)   (masked share, mod q)
```

where `H_share` is the Poseidon1 hash of the domain separator, epoch ID, participant indices, and
the shared secret point `r_{i,j} · pub_j = sk_j · R_i(j)`.

Decryption by participant `j`:
```
s_i(j) = σ_i(j) − H_share(eid, i, j, sk_j · R_i(j))   (mod q)
```

### DKG Protocol

The full DKG proceeds in 4 phases, all block-number delimited:

**Phase 1 — Initiation**: Any caller (typically a participating dkg-node, racing
others on a random jitter) creates an epoch specifying `(t, n)` and policy
parameters. A unique 12-byte `epochId` is generated on-chain. The contract
snapshots `registry.activeCount()`, derives a per-epoch **lottery threshold** so
that on average `α × n` nodes are eligible, and pins a `seedBlock = block.number
+ SEED_DELAY_BLOCKS` whose future blockhash will become the epoch seed. Phase
deadline blocks (`registrationDeadlineBlock`, `contributionDeadlineBlock`,
`finalizeNotBeforeBlock`) are derived from the contract's immutable
`EPOCH_DURATION_BLOCKS` and the per-phase BPS constants in
`solidity/src/libraries/Sizes.sol`. `createEpoch` is permissionless and reverts
unless `block.number >= nextEpochStartBlock()` so the cadence is enforced
trustlessly.

**Phase 2 — Trustless committee selection (lottery)**: Once `block.number ≥ seedBlock`,
any registered node calls `claimSlot(epochId)`. The first such call lazily resolves
`seed = blockhash(seedBlock)`. A node is eligible iff
`keccak256(seed ‖ msg.sender) < lotteryThreshold`. Eligible nodes race
**first-come-first-served** until `committeeSize` slots are filled, at which point the
contract snapshots the committee key hash and transitions to Contribution. An
epoch that fails to fill within the registration window is aborted; the next
scheduled epoch then opens automatically once the cadence threshold elapses.

**Phase 3 — Main DKG (contribution)**: Each participant `i` samples random polynomial coefficients
`{a_{i,k}}` and encryption nonces `{r_{i,j}}`, then publishes:
- Commitments: `C_i(k) = a_{i,k} · G` for `k ∈ {0, …, t−1}`
- Encrypted shares: `(R_i(j), σ_i(j))` for all `j ∈ [n]`
- A Groth16 proof `π_i` of correctness (see [Contribution Circuit](#contribution-circuit-dkg-phase-3))

The contract rejects the transaction if the proof is invalid.

**Phase 4 — Finalization**: Once `minValidContributions` contributions are accepted, anyone may call
`finalizeEpoch`. This computes and persists:
- Aggregate commitments: `C̄(k) = Σ_{ℓ ∈ I} C_ℓ(k)` for `k ∈ {0, …, t−1}`
- Collective public key: `PK = C̄(0) = F(0) · G`
- Share commitments: `D_i = Σ_k i^k · C̄(k) = F(i) · G` for each accepted participant `i`

Finalization is also proof-gated (see [Finalize Circuit](#finalize-circuit-dkg-phase-4)).

Each participating node privately computes their secret share:
```
d_i = Σ_{ℓ ∈ I} s_ℓ(i)  =  F(i)
```
by decrypting the encrypted shares they received on-chain.

### Trustless Lottery Committee Selection

Given an epoch with policy parameters `(n, α)` — where `n = committeeSize` and
`α ∈ (0, 1]` (encoded as `lotteryAlphaBps`, basis points out of 10 000) — and
the registry snapshot `R = activeCount()` at the moment of epoch creation, the
contract computes:

```
lotteryThreshold = ⌊ (α · n · 2²⁵⁶) / R ⌋
```

This is the **eligibility threshold**: a pseudo-random 256-bit value uniformly
derived from the future seed and the node's address must fall below
`lotteryThreshold` for that node to claim a slot. By construction, the
expected number of eligible nodes is `E[|eligible|] = α · n`. With `α = 1.0`
the expectation equals the committee size; with `α > 1.0` one oversubscribes
to absorb liveness failures (the testnet default is α = 1.5, configured via
`--epoch-policy.lottery-alpha-bps` on the dkg-node binary). An epoch that
fails to fill its committee within the registration window simply gets
aborted; the next scheduled epoch opens automatically once the cadence
threshold elapses.

**Seeding.** At `createEpoch` the contract pins `seedBlock = block.number +
SEED_DELAY_BLOCKS` but does **not** yet know the seed. Once `block.number ≥
seedBlock`, the first call to `claimSlot` reads `blockhash(seedBlock)` and
stores it as `seed`. Binding the seed to a future blockhash (`SEED_DELAY_BLOCKS
≥ 1`) prevents the proposer from tuning the eligibility set by picking a
favourable `createEpoch` block — they cannot predict the future blockhash.

**Eligibility check.** For each registered node calling `claimSlot`:

```
h = keccak256(seed ‖ msg.sender)           (256-bit big-endian integer)
eligible iff h < lotteryThreshold
```

The keccak hash is a verifiable random function seeded by the blockhash:
every observer can independently recompute `h` for any address and confirm
whether that node was allowed to claim the slot. No trusted coordinator is
involved and no ZK proof is needed — the check is a handful of opcodes in the
contract.

**Race and termination.** Eligible nodes race first-come-first-served until
`committeeSize` slots have been filled, at which point the contract snapshots
`keccak256(indexes ‖ publicKeys)` of the final committee and transitions the
epoch from Registration to Contribution. Any further `claimSlot` calls revert.
The committee snapshot is what later contribution proofs are verified against
— the `contributionVerifier` only accepts a proof whose recipient list keccak
matches this snapshot, so the committee is effectively locked in a single slot
of storage.

**Security properties.**
- *No organizer influence over membership.* The organizer sets `n` and `α` but
  cannot prefer specific nodes: the eligibility set is determined by a
  blockhash published after `createEpoch`.
- *No validator griefing beyond 1-block withholding.* A malicious block
  proposer of `seedBlock` can choose to withhold or reveal their block to
  shift the seed by one candidate. In practice the seed is domain-separated
  and the lottery is a uniform threshold check, so withholding buys
  negligible bias.
- *Bounded bias from registered Sybils.* Because eligibility is uniform in
  the hash, registering `k` Sybil addresses grows the attacker's expected
  slots by `k · α · n / R`. The registry is append-only and nodes must
  publish a valid BabyJubJub key, which is the designed registration cost.
- *Liveness under node failure.* If fewer than `committeeSize` eligible nodes
  claim before `registrationDeadlineBlock`, the epoch is aborted and the next
  scheduled epoch opens automatically once `block.number >=
  nextEpochStartBlock()`. No epoch is stuck waiting for a node that went
  offline.

**Keeping the registry honest.** `DKGRegistry` is append-only at the storage
level, but it tracks an `activeCount` alongside `nodeCount` and a
per-operator `lastActiveBlock` that `DKGManager` refreshes on every
accepted contribution (via a one-shot `setManager` callback). The lottery
uses `activeCount` as the denominator, not `nodeCount`, so stragglers are
automatically excluded the moment they are demoted. Any address can call
`reap(operator)` once
`block.number > lastActiveBlock + INACTIVITY_WINDOW`; the target flips to
`INACTIVE`, `activeCount--`, and a reaped node's subsequent `claimSlot`
calls revert. An operator who is simply unlucky — healthy, but never
selected — can call the cheap `heartbeat()` entry point to refresh their
timestamp; reaped operators rejoin via `reactivate()` (or by rotating
their key with `updateKey`). The per-epoch cost of this mechanism is a
single cross-contract SSTORE on each successful `submitContribution`, and
none of the other phases pay anything.

```bash
# Demote a known-dead operator (permissionless, anyone can call).
cast send $DKG_REGISTRY "reap(address)" 0xDeadOperator --rpc-url $RPC_URL --private-key $KEY

# Self-refresh as an unlucky-but-healthy operator.
cast send $DKG_REGISTRY "heartbeat()" --rpc-url $RPC_URL --private-key $KEY
```

`INACTIVITY_WINDOW` is set once at registry deployment (default: 50 400
blocks ≈ 7 days at 12-second block time, overridable via the
`INACTIVITY_WINDOW` env var on `make solidity-deploy`).

### Threshold Decryption

An ElGamal ciphertext `(C_1, C_2)` encrypts message `M` under public key `PK`:
```
C_1 = r · G
C_2 = M · G + r · PK
```
where the message is embedded as a scalar `m` via `M = m · G`.

Decryption without revealing `sk = F(0)`:

**Partial decryption** by node `i`:
```
δ_i = d_i · C_1
```
accompanied by a Chaum-Pedersen DLEQ proof that `δ_i` and `D_i` share the same discrete log
with respect to `C_1` and `G` respectively.

**Combination** (Lagrange interpolation in the exponent):
```
Δ = sk · C_1 = Σ_{k ∈ [t]} λ_k · δ_{x_k}
M · G = C_2 − Δ
```

The message `m` can be recovered by brute force or BSGS if the plaintext space is small.

#### Plaintext range

Recovering `m` from `m · G` is a discrete-log problem with no general efficient
solution on BabyJubJub. The protocol relies on `m` being bounded so the
committee can search exhaustively. The implementation pins two caps:

| Surface | Algorithm | Cap | First-call cost | Per-call cost (worst case) |
|---|---|---|---|---|
| Go committee node — `cmd/davinci-dkg-node/dlog.go` | Baby-step / giant-step (BSGS) | **2⁵⁰** ≈ 1.13 × 10¹⁵ | ~30–60 s table build, ~1–2 GB heap | ~30–60 s |
| TypeScript SDK — `sdk/src/crypto/elgamal.ts` `decrypt()` | BSGS (in-process) | **2³²** ≈ 4.3 × 10⁹ | ~1–2 s table build, ~16 MB heap | <1 s |

The two caps differ on purpose. The committee runs on operator hardware with
gigabytes of RAM headroom; SDK consumers may run in a browser, where a 2⁵⁰
table (~1 GB) is not practical. The SDK's lower cap only affects users
calling `decrypt()` directly with a private key — for example in tests or
single-key demos. The on-chain threshold-decryption flow always uses the
committee's 2⁵⁰ limit.

Submitting a ciphertext whose plaintext is at or above 2⁵⁰ leaves the epoch
unrecoverable: the committee will fail at the combine step, and retrying will
always produce the same failure. The playground UI rejects such inputs
client-side. To support larger plaintexts (up to ~2⁵⁶) the operator-side
algorithm would need to switch from BSGS to Pollard's kangaroo, which trades
compute time for constant memory.

### Public-Input Compression (BRLC)

Because Groth16 on BN254 costs ~6,650 gas per public input, large transcripts
(commitment vectors, encrypted shares, partial decryptions) are compressed using
**Binding Random Linear Combinations (BRLC)**:

```
C = Σ_{i=1}^{l} ρ^i · v_i
```

The challenge `ρ` is derived from the epoch ID and a domain separator using `keccak256`, making it
unpredictable at the time the inputs are committed (Fiat-Shamir). The in-circuit check recomputes
the BRLC and asserts equality, reducing `l` public inputs to a single scalar.

On-chain cost: ~70 gas per element (vs. ~6,650 gas per Groth16 public input).

---

## ZK-SNARK Circuits

The production circuit set comprises **four** circuits — Contribution,
Finalize, PartialDecrypt, and DecryptCombine. All use **Groth16** on
**BN254**. BabyJubJub curve operations are performed natively (inside the
BN254 scalar field). Hashing uses **Poseidon1**.

Fixed-size circuit arrays use prefix masks derived from the actual
threshold/committee size, so one compiled circuit serves all parameter choices
up to the compile-time maximum. The bound is the single Go constant
`circuits/common.MaxN`, which all four circuit-side aliases
(`MaxCoefficients` / `MaxRecipients` / `MaxParticipants` / `MaxShares`) reference.
The Solidity contract reads the same value from `DKGManager.sol::MAX_N`. The
default build uses `MaxN = 32`; see [`BENCHMARKS.md`](BENCHMARKS.md) for
constraint counts, proving / verifying times, and gas figures.

### Contribution Circuit (DKG Phase 3)

**Package**: `circuits/contribution`
**Public inputs** (8 scalars): `EpochHash`, `Threshold`, `CommitteeSize`, `ContributorIndex`,
`CommitmentHash`, `ShareHash`, `Challenge`, `TranscriptCommitment`

**Private inputs**: polynomial coefficients, encryption nonces, Shamir shares,
mask quotients, share masks, carry bits

**Proves**:
1. Coefficient commitments: `C_i(k) = a_{i,k} · G` for all `k ∈ {0, …, t−1}`
2. Shamir evaluation: `s_i(j) = Σ_k a_{i,k} · j^k` for all `j ∈ [n]`
3. Feldman verification: `s_i(j) · G = Σ_k j^k · C_i(k)` for all `j ∈ [n]`
4. Ephemeral key: `R_i(j) = r_{i,j} · G` for all `j ∈ [n]`
5. Share encryption: `σ_i(j) = s_i(j) + H_share(eid, i, j, r_{i,j} · pub_j) (mod l)`
6. Commitment hash: `CommitmentHash = Poseidon1(EpochHash, ContributorIndex, t, C_i(0), …)`
7. Share hash: `ShareHash = Poseidon1(EpochHash, ContributorIndex, n, idx_1, R_1, σ_1, …)`
8. BRLC transcript: `TranscriptCommitment = BRLC(Challenge, transcript_vector)`

The transcript vector encodes all commitments, recipient indexes, recipient public keys,
ephemeral points, and masked shares. The contributor's individual public key share
`a_{i,0}·G` is the first element of the commitment vector inside the BRLC transcript;
it is verified at finalize time as part of `aggregateCommitments[0]`.

### Finalize Circuit (DKG Phase 4)

**Package**: `circuits/finalize`
**Public inputs** (9 scalars): `EpochHash`, `Threshold`, `CommitteeSize`, `AcceptedCount`,
`AggregateHash`, `CollectivePublicKey`, `ShareCommitmentHash`, `Challenge`, `TranscriptCommitment`

**Private inputs**: participant indexes, per-participant commitment vectors,
aggregate commitments, share commitments

**Proves**:
1. Aggregate commitments: `C̄(k) = Σ_{ℓ ∈ I} C_ℓ(k)` for all `k`
2. Public key hash: `CollectivePublicKey = Poseidon1(EpochHash, C̄(0).X, C̄(0).Y)`
3. Aggregate hash: `AggregateHash = Poseidon1(EpochHash, t, n, |I|, C̄(0), …)`
4. Share commitments: `D_i = Σ_k i^k · C̄(k)` for each accepted `i ∈ I`
5. Share commitment hash: `ShareCommitmentHash = Poseidon1(EpochHash, t, n, |I|, i_1, D_1, …)`
6. BRLC transcript: covers all participant indexes, contribution commitments, aggregate commitments,
   and share commitments

### PartialDecrypt Circuit (Decryption Phase 2)

**Package**: `circuits/partialdecrypt`
**Public inputs** (13 scalars): `EpochHash`, `ParticipantIndex`, `Base.X`, `Base.Y`,
`PublicKey.X`, `PublicKey.Y`, `Delta.X`, `Delta.Y`, `A1.X`, `A1.Y`, `A2.X`, `A2.Y`, `Response`

**Private inputs**: `Secret` (`d_i`), `Nonce` (`r`)

**Proves** a Chaum-Pedersen DLEQ relation:
1. `PublicKey = Secret · G` (commitment to secret: `D_i = d_i · G`)
2. `Delta = Secret · Base` (partial decryption: `δ_i = d_i · C_1`)
3. `A1 = Nonce · G` (nonce commitment)
4. `A2 = Nonce · Base` (nonce commitment on base)
5. Challenge: `e = Poseidon1(domain, D_i, C_1, δ_i, A_1, B_1)` (Fiat-Shamir)
6. Response equations: `Response · G = A1 + e · PublicKey`  and  `Response · Base = A2 + e · Delta`

### DecryptCombine Circuit (Decryption Phase 3)

**Package**: `circuits/decryptcombine`
**Public inputs** (7 scalars): `EpochHash`, `Threshold`, `ShareCount`,
`CombineHash`, `PlaintextHash`, `Challenge`, `TranscriptCommitment`

**Private inputs**: `CiphertextC1`, `CiphertextC2`, `Plaintext`,
participant indexes, partial decryption points, pre-computed Lagrange coefficients

**Proves**:
1. Combine hash: `CombineHash = Poseidon1(EpochHash, t, |Q|, C1, C2, idx_1, δ_1, …)`
2. Plaintext binding: `PlaintextHash = Plaintext` (the scalar `m` is exposed directly)
3. Lagrange combination: `Δ = Σ_{k ∈ [t]} λ_k · δ_{x_k}`
4. ElGamal decryption: `Plaintext · G + Δ = C_2`
5. BRLC transcript: covers ciphertext, participant indexes, partial decryption points

Lagrange coefficients `λ_k` are pre-computed natively in the BabyJubJub scalar field
(`r_bjj`) and passed as private witnesses. Computing them in-circuit via `api.Div` would
use `BN254.Fr` arithmetic, which gives incorrect results for negative coefficients because
`BN254.Fr − 1 ≠ r_bjj − 1` as BJJ scalars. The `Plaintext · G + Δ = C_2` constraint
implicitly validates that the witnesses were used correctly.

---

## Smart Contracts

The Solidity workspace lives in `solidity/` (Foundry, `solc 0.8.28`, EVM Cancun, `via_ir = true`).

The production deployment is three contracts: `DKGRegistry` (operator
identities + liveness), `DKGManager` (epoch lifecycle, contributions,
finalize, ciphertexts, partial / combined decryption), and
`DKGAppManager` (per-application surface: `registerApplication`,
`registerApplicationCoDec`, `submitOrganizerShare`, `getApplication`).
The manager / app-manager split is a deployment concern: it keeps each
contract under the EIP-170 24,576-byte runtime limit. Deploy order is
`DKGRegistry → DKGManager → DKGAppManager`, followed by the one-shot
wiring calls `DKGRegistry.setManager(DKGManager)` and
`DKGManager.setAppManager(DKGAppManager)`. After wiring, both manager
contracts share the same on-chain epoch storage through the manager.

### DKGRegistry

**Source**: `solidity/src/DKGRegistry.sol`
**Interface**: `solidity/src/interfaces/IDKGRegistry.sol`

Stores the share-encryption public keys (BabyJubJub points) of eligible operator nodes.

| Function | Description |
|---|---|
| `registerKey(pubX, pubY)` | Register caller's BabyJubJub public key. Reverts if already registered or coordinates are zero. Increments `nodeCount` and `activeCount`. |
| `updateKey(pubX, pubY)` | Update caller's previously registered key. Auto-reactivates an `INACTIVE` row. |
| `heartbeat()` | Cheap self-refresh that updates `lastActiveBlock` for an unlucky-but-healthy operator. |
| `reactivate()` | Rejoin the active set after being reaped. |
| `reap(operator)` | Permissionless: demote an operator whose `lastActiveBlock` has fallen behind `INACTIVITY_WINDOW`. |
| `getNode(operator)` | Returns the `NodeKey` struct `{operator, pubX, pubY, status}` for the given address. |
| `nodeCount()` / `activeCount()` | Total registered addresses / currently-active addresses. The lottery threshold uses `activeCount`. |
| `INACTIVITY_WINDOW()` | Liveness window in blocks (set at deployment). |

**Events**: `NodeRegistered`, `NodeUpdated`, `NodeMarkedActive`,
`NodeReaped`, `NodeReactivated`, `ManagerSet`.

### DKGManager

**Source**: `solidity/src/DKGManager.sol`
**Interface**: `solidity/src/interfaces/IDKGManager.sol`

Owns the complete epoch lifecycle: creation, trustless lottery-based committee
selection, proof-gated contribution, finalization, and threshold decryption.
Each state-mutating operation that involves cryptographic claims is gated by
a Groth16 verifier.

#### Epoch Lifecycle

```
Created → Registration (lottery) → Contribution → Finalized
                                                ↘ Aborted
```

The `EpochPhase` enum exposes `None`, `Registration`, `Contribution`,
`Finalized`, `Aborted` (plus a reserved `Completed`).

The contract retains a fixed-size ring buffer of the most recent `EPOCH_HISTORY_SIZE`
epoch IDs. When a new epoch is created and the buffer is full, the oldest
live epoch's storage is wiped (`delete epochs[…]`, etc.), keeping long-term
storage bounded. Off-chain consumers reconstruct historical epoch data from
the event log.

#### State-Mutating Functions

| Function | Phase | Access | Description |
|---|---|---|---|
| `createEpoch(threshold, committeeSize, minValidContributions, lotteryAlphaBps, decryptionPolicy)` | Any block ≥ `nextEpochStartBlock()` | Open (permissionless) | Create a new DKG epoch. Snapshots `activeCount` from the registry and derives the per-epoch lottery threshold. Phase deadline blocks (`registrationDeadlineBlock`, `contributionDeadlineBlock`, `finalizeNotBeforeBlock`) and `seedBlock` are derived on-chain from the immutable `EPOCH_DURATION_BLOCKS` plus the per-phase BPS constants in `Sizes.sol`. The cadence guard `block.number >= nextEpochStartBlock()` enforces one full `EPOCH_DURATION_BLOCKS` between epoch starts. `decryptionPolicy` gates the legacy per-epoch `submitCiphertext` path (owner-only, not-before/not-after block and timestamp, max submissions) — all-zero = no constraint. Returns `bytes12 epochId`. |
| `claimSlot(epochId)` | Registration | Any registered eligible node | First-come-first-served self-claim. The first call after `block.number ≥ seedBlock` lazily resolves `seed = blockhash(seedBlock)`. The caller is admitted iff `keccak256(seed ‖ msg.sender) < lotteryThreshold`. The contract stops accepting claims once `committeeSize` slots are filled and immediately advances to Contribution. |
| `submitContribution(epochId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)` | Contribution | Selected participant | Submit polynomial commitments and encrypted shares with a Groth16 proof. The committee membership / pubkey list is verified against a single keccak snapshot taken when the lottery filled (no per-recipient registry calls). The collective public key is captured later by `finalizeEpoch` from `aggregateCommitments[0]`, so contributions don't pay for an on-chain BabyJubJub addition. |
| `finalizeEpoch(epochId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)` | After min contributions, on/after `finalizeNotBeforeBlock` | Open | Aggregate commitments, publish collective public key and share commitments. Advances to Finalized. Reverts with `FinalizeTooEarly` if `block.number < policy.finalizeNotBeforeBlock`. The transcript is read directly from calldata; share commitments are stored as `keccak256(x,y)` (1 slot each). In production, `davinci-dkg-node` instances finalize automatically after their contribution lands, using a deterministic per-epoch stagger derived from the lottery seed so only one node submits at a time (the rest see `AlreadyFinalized` and stop). |
| `submitCiphertext(epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y)` | Finalized | `aid == bytes32(0)` is the legacy per-epoch path gated by the epoch `DecryptionPolicy`; non-zero `aid` is gated by the application's own `AppPolicy`. Write-once per `(epochId, aid, ciphertextIndex)`. | Publish a ciphertext to be threshold-decrypted under either the epoch key (`aid = 0`) or the application-specific key. Stores `keccak256(c1,c2)` and emits `CiphertextSubmitted` carrying the raw coordinates so nodes (and consumers) can read them from the event log. |
| `submitPartialDecryption(epochId, aid, participantIndex, ciphertextIndex, c1x, c1y, c2x, c2y, deltaHash, proof, input)` | Finalized | Selected participant | Submit a partial decryption `δ_i = d_i · C_1` with a Chaum-Pedersen DLEQ proof. Keyed by `(epochId, aid, participant, ciphertextIndex)` to support multiple ciphertexts per epoch and per application. |
| `combineDecryption(epochId, aid, ciphertextIndex, combineHash, plaintext, transcript, proof, input)` | Finalized, ciphertext submitted, ≥t partials | Open | Combine `t` partial decryptions via Lagrange interpolation. Proof is bound to the on-chain ciphertext hash (no substitution possible). For `aid != 0` in mode-1 applications the combine step also folds in the organizer's `Δ_org`. Stores the recovered plaintext `uint256`; readable via `getPlaintext`. |
| `abortEpoch(epochId)` | Any non-terminal | Organizer | Abort the epoch. Advances to Aborted. |

> `registerApplication`, `registerApplicationCoDec`, `submitOrganizerShare`,
> and the read-side `getApplication` live on
> [`DKGAppManager`](#dkgappmanager), not `DKGManager`. They operate on
> the same epoch storage via the wired manager reference.

#### View Functions

| Function | Returns |
|---|---|
| `getEpoch(epochId)` | `Epoch` struct: `organizer, policy, decryptionPolicy, status, nonce, startBlock, seedBlock, seed, lotteryThreshold, claimedCount, contributionCount, partialDecryptionCount, ciphertextCount`. The `policy` field is an `EpochPolicy` struct: `threshold, committeeSize, minValidContributions, lotteryAlphaBps, registrationDeadlineBlock, contributionDeadlineBlock, finalizeNotBeforeBlock` — the deadline blocks are populated by `createEpoch` from the contract's immutable per-phase offsets. |
| `nextEpochStartBlock()` | Earliest block at which the next `createEpoch` may succeed (`lastEpochStartBlock + EPOCH_DURATION_BLOCKS`, or current block before any epoch). |
| `epochDurationBlocks()` | The deploy-time `EPOCH_DURATION_BLOCKS` immutable. |
| `lastEpochStartBlock()` | Block in which the most recent epoch was created. |
| `getCollectivePublicKey(epochId)` | `Point {x, y}` — the collective public key `PK = Σ_i a_{i,0}·G`. Written exactly once at `finalizeEpoch` from `aggregateCommitments[0]`; returns the identity `(0, 1)` before the epoch is finalized. |
| `getDecryptionPolicy(epochId)` | `DecryptionPolicy` struct: `ownerOnly, maxDecryptions, notBeforeBlock, notBeforeTimestamp, notAfterBlock, notAfterTimestamp`. Set at `createEpoch`; gates the legacy `aid == 0` path only. |
| `selectedParticipants(epochId)` | `address[]` — ordered committee in claim order. |
| `getContribution(epochId, contributor)` | `ContributionRecord` (only `contributorIndex`, `commitmentVectorDigest`, `accepted` are persisted; the rest live in `ContributionSubmitted` events). |
| `getPartialDecryption(epochId, aid, participantIndex, ciphertextIndex)` | `PartialDecryptionRecord` — `(participantIndex, ciphertextIndex, deltaHash, accepted)`. The raw δ point is not stored on-chain; subscribers reconstruct it from the `PartialDecryptionSubmitted(epochId, aid, participant, participantIndex, ciphertextIndex, deltaX, deltaY)` event log. |
| `getCombinedDecryption(epochId, aid, ciphertextIndex)` | `CombinedDecryptionRecord`: `ciphertextIndex`, `completed`, `plaintext`. `combineHash` is only in the `DecryptionCombined` event. |
| `getPlaintext(epochId, aid, ciphertextIndex)` | `uint256` — recovered plaintext scalar; `0` if the decryption has not been combined yet (check `getCombinedDecryption(...).completed` to disambiguate). |
| `getCiphertextHash(epochId, aid, ciphertextIndex)` | `bytes32` — `keccak256(abi.encode(c1x, c1y, c2x, c2y))` of the submitted ciphertext; raw coordinates are only in the `CiphertextSubmitted` event. |
| `getShareCommitmentHash(epochId, participantIndex)` | `bytes32` = `keccak256(abi.encode(x, y))`. The pre-image lives in the `EpochFinalized` event. |
| `getContributionVerifierVKeyHash()` | `bytes32` |
| `getPartialDecryptVerifierVKeyHash()` | `bytes32` |
| `getFinalizeVerifierVKeyHash()` | `bytes32` |
| `getDecryptCombineVerifierVKeyHash()` | `bytes32` |

### DKGAppManager

**Source**: `solidity/src/DKGAppManager.sol`

Sibling contract to `DKGManager` that hosts the per-application surface.
Split out of `DKGManager` to keep both contracts under the EIP-170 24,576-byte
runtime-bytecode limit; conceptually the two are one logical "manager" that
shares the same epoch and application storage. The link is established
exactly once by the deployer via `DKGManager.setAppManager(address)`.

| Function | Description |
|---|---|
| `registerApplication(epochId, aid, policy)` | Register a public-derivation (mode-0) application against a finalized epoch. Derives `S = keccak256(epochId ‖ PK_ep ‖ aid) mod q` on-chain and stores it. |
| `registerApplicationCoDec(epochId, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)` | Register an organizer co-decryption (mode-1) application. Verifies the Schnorr proof of knowledge of `sk_org` on-chain (challenge `c = keccak256(domain ‖ epochId ‖ aid ‖ PK_org ‖ A) mod L`). |
| `submitOrganizerShare(epochId, aid, ciphertextIndex, c1x, c1y, c2x, c2y, deltaOrgX, deltaOrgY, dleqProof, dleqInput)` | Publish the organizer's `Δ_org = sk_org · C_1` with a Chaum-Pedersen DLEQ proof. Required before `combineDecryption` in organizer co-decryption mode. |
| `getApplication(epochId, aid)` | `Application` struct: `creator, mode, derivationS, organizerPK, policy, createdAtBlock, exists`. |

### Per-Application Surface

A finalized epoch fixes a single collective public key `PK_ep` shared by the
committee. To support many independent encryption contexts that all reuse the
same committee — without re-running DKG — `DKGAppManager` exposes a per-application
key derivation surface keyed by an arbitrary `bytes32 aid`. Each application
records a `creator`, `mode`, `derivationS`, `organizerPK`, an `AppPolicy`
(`authorizedSubmitter, maxCiphertexts, notBeforeBlock, notAfterBlock`),
`createdAtBlock`, and `exists` flag. The legacy per-epoch path is the special
case `aid = bytes32(0)`.

There are two registration modes:

- **Mode 0 — public derivation**, via
  `registerApplication(eid, aid, policy)`. The contract derives a public
  scalar `S = keccak256(eid || PK_ep || aid) mod q` and stores it on-chain.
  The implicit per-application encryption key is `PK_aid = PK_ep + S·G`, which
  any reader can recompute. The committee handles partial decryptions
  natively: the additive `S·G` term cancels out at the combine step.
- **Mode 1 — organizer co-decryption**, via
  `registerApplicationCoDec(eid, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)`.
  The organizer publishes `PK_org = sk_org · G` together with a Schnorr proof
  of knowledge of `sk_org`. The implicit per-application encryption key is
  `PK_aid = PK_ep + PK_org`, so decryption requires both the threshold
  committee and the organizer to co-operate. The organizer contributes
  `Δ_org = sk_org · C_1` via `submitOrganizerShare`, which the combine step
  folds in alongside the committee's Lagrange interpolation.

All ciphertext-handling entry points (`submitCiphertext`,
`submitPartialDecryption`, `combineDecryption`) take an `aid` argument.
Passing `aid = bytes32(0)` selects the legacy per-epoch key and is gated by
the epoch-level `DecryptionPolicy`; any non-zero `aid` is gated by that
application's own `AppPolicy` instead.

---

## Getting Started

### Prerequisites

- **Go 1.25+**
- **Foundry** (`forge`, `cast`, `anvil`) — [install](https://book.getfoundry.sh/getting-started/installation)
- **Docker** + **Docker Compose** — for integration tests
- `abigen` — Go Ethereum ABI tool (`go install github.com/ethereum/go-ethereum/cmd/abigen@v1.17.1`)
- `jq`

### Build

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg

# Download Go dependencies
go mod download

# Build the node binary
go build ./cmd/davinci-dkg-node/...

# Build the Solidity contracts
cd solidity && forge build
```

### Run Tests

**Unit and circuit tests** (no Docker required — uses cached or freshly compiled artifacts):

```bash
go test ./...
```

**Solidity contract tests**:

```bash
cd solidity && forge test
```

**Integration tests** (requires Docker; spins up Anvil + deployer container):

```bash
RUN_INTEGRATION_TESTS=true go test ./tests/... -count=1 -timeout=10m
```

### Compile Circuits

Circuit artifacts (constraint system, proving key, verifying key) are cached under
`~/.davinci/artifacts/` and keyed by SHA-256 hash. The first time a circuit test or
integration test runs, missing artifacts are compiled from source via a local trusted setup.

To explicitly recompile all circuits and regenerate the Solidity verifier files, use the
provided Makefile pipeline:

```bash
# Full circuit update pipeline (compile → patch Go hashes → Solidity build → Go bindings)
make circuits
```

Or run each step individually:

```bash
# 1. Compile all 4 circuits; write artifacts and update Solidity verifier stubs
#    Output JSON with hashes is saved to /tmp/circuit-artifacts.json
make circuits-compile

# 2. Patch config/circuit_artifacts.go with the new hashes
make circuits-update-hashes

# 3. Rebuild the Solidity workspace
make solidity-build

# 4. Regenerate Go ABI bindings
make solidity-bind
```

The artifact directory defaults to `~/.davinci/artifacts`. Override with:

```bash
ARTIFACTS_DIR=/path/to/artifacts make circuits-compile
```

After any circuit change, commit `config/circuit_artifacts.go`,
`solidity/src/verifiers/*.sol`, and the updated Go bindings in `solidity/bindings/`.

### Deploy Contracts

```bash
# Against a local Anvil instance
anvil &
RPC_URL=http://127.0.0.1:8545 CHAIN_ID=31337 PRIVATE_KEY=0x<key> make solidity-deploy
```

For the Docker-based integration harness, the deployer container handles this automatically
(see [Local Testnet](#local-testnet)).

---

## Gas Profile

Gas costs are bounded per phase. A single committee node pays for a slot
claim, one contribution (~213 k gas, dominated by Groth16 verification
plus calldata), and — when scheduled — one partial decryption per
ciphertext (~99 k). The organizer pays only for `createEpoch` and
optionally `finalizeEpoch` if no node finalizes first. The two
heavyweight entry points are `registerKey` on `DKGRegistry` (~1.27 M
after the keccak-Schnorr swap, paid once per node-key lifetime) and
`submitCiphertext` on `DKGManager` (~2.06 M with the full BabyJubJub
prime-subgroup check on both ciphertext points). The authoritative
per-call breakdown, including how figures shift with `MaxN` and committee
size, lives in [`BENCHMARKS.md`](BENCHMARKS.md).

---

## Switching `MaxN`

Changing the maximum committee size is a **two-line edit**:

```go
// circuits/common/sizes.go
const MaxN = 32   // 16 or 32 (or any other value)
```

```solidity
// solidity/src/DKGManager.sol
uint256 internal constant MAX_N = 32;   // must equal circuits/common.MaxN
```

After editing, run `make circuits` to recompile the circuits, regenerate the
proving keys, patch the artifact hashes in `config/circuit_artifacts.go`, rebuild
the Solidity verifier wrappers, and regenerate the Go ABI bindings.

---

## Local Testnet

The `testnet/` directory contains a self-contained multi-node DKG test network
that demonstrates the full protocol end-to-end.


### Start the network

The Makefile wraps compose with sensible defaults and starts every non-runner
service (including the web explorer) in one shot:

```bash
make testnet-up                                # 3 nodes + anvil + deployer + standalone UI
make testnet-up DKG_NODE_COUNT=8 DKG_THRESHOLD=5

# Expose the browser-side RPC URL if you will access from another host:
UI_PUBLIC_RPC=http://<host-ip>:8545 make testnet-up
```

Once the command returns, open `http://<host-ip>:8081/` in a browser.

The equivalent raw compose invocation:

```bash
cd testnet
DKG_NODE_COUNT=3 DKG_THRESHOLD=2 \
  docker compose up --scale dkg-node=3 --build
```

### Run the scenario

Each `davinci-dkg-node` instance auto-creates new epochs by default
(`--auto-create-epochs=true`, env `DAVINCI_DKG_AUTO_CREATE_EPOCHS`),
racing other nodes on a uniform-random jitter and reverting cheaply when
another node wins. Bring up the fleet and the schedule drives itself:

```bash
make testnet-up                                          # defaults: 3 nodes
make testnet-up DKG_NODE_COUNT=8                         # custom sizing

# Watch the cadence + per-phase progress:
make testnet-logs
```

Each scheduled epoch then runs through:
1. Any node fires `createEpoch` once `block.number >= nextEpochStartBlock()`
2. Lottery: every active node calls `claimSlot` and self-checks eligibility
3. Contribution: each selected participant submits their DKG contribution proof
4. Finalize: one node (deterministic per-epoch stagger) submits the finalize proof
5. The collective public key is now live on-chain; the UI / SDK can read it
6. Anyone can submit a ciphertext to be threshold-decrypted (mode-0 derivation)
7. Each selected participant submits a partial decryption with DLEQ proof
8. Anyone calls `combineDecryption` to recover the plaintext on-chain


### Configuration

| Env variable | Default | Description |
|---|---|---|
| `DKG_NODE_COUNT` | `3` | Number of DKG node replicas to start |
| `DKG_THRESHOLD` | `2` | Decryption threshold (`t`-of-`n`) |
| `ANVIL_PORT` | `8545` | Host port for the Anvil RPC (bound on `0.0.0.0`) |
| `DEPLOYER_PORT` | `8888` | Host port for the deployer HTTP server |
| `UI_PORT` | `8081` | Host port the DKG explorer listens on (bound on `0.0.0.0`) |
| `UI_PUBLIC_RPC` | `http://localhost:8545` | RPC URL advertised to browsers in `/config.json`. Override with the LAN/public IP of the host when accessing the explorer from a remote machine. Can also be changed live in the explorer's Settings page. |

---

## Web Explorer

`ui/` contains a single-page React application that acts as both a
block-explorer and an interactive playground for a live `DKGManager` /
`DKGRegistry` pair. Vite + React + TypeScript + Chakra UI v3 + React
Query + RainbowKit stack, talking to the chain directly via `viem`. Ships
as its own Docker image (`ghcr.io/vocdoni/davinci-dkg-ui`) — completely
decoupled from the `davinci-dkg-node` Go binary, which itself does not
serve any HTTP.

In the testnet the explorer is available at `http://<host>:8081/` as soon
as `make testnet-up` returns — the testnet's `dkg-ui` compose service
runs the standalone UI image and points it at the in-cluster Anvil RPC.

### What it shows

The UI ships two surfaces: a plain-English default view and a **debug
mode** (terminal icon in the header) that auto-expands every "Show
technical details" disclosure. Power users can flip the toggle to see raw
event args, BabyJubJub coordinates, transcript hashes and full hex
addresses inline; everyone else gets short hashes, durations, and status
badges.

- **Overview**: total epochs, active / total nodes, latest block, chain
  ID, and the 5 most recent epochs.
- **Epochs**: filter chips (registration / contribution / finalized /
  aborted) over the ring-buffered on-chain history. Click a row to open
  the epoch detail view.
- **Epoch detail**: status badge, plain-English summary ("Awaiting
  contributions — 1/2 accepted so far"), a four-step phase timeline, a KV
  grid of policy facts (deadlines as durations + block #'s), counters,
  and Participants / Activity tabs. Each on-chain event is summarised in
  English; raw `args` are hidden behind a per-event disclosure.
- **Registry**: stats + node table; key coordinates live behind a
  disclosure so the table stays scannable.
- **Playground**: an interactive end-to-end demo of the full DKG +
  threshold decryption flow in seven steps:
  1. Connect a wallet (RainbowKit — MetaMask, WalletConnect, etc.).
  2. Create a new DKG epoch with configurable epoch + decryption policy.
  3. Watch the epoch progress live through registration, contribution,
     and finalization phases (with an Abort button while non-terminal).
  4. Read the collective public key `(x, y)` from
     `getCollectivePublicKey(epochId)` once the epoch is **Finalized**.
     The encrypt step is gated on Finalized so `submitCiphertext` cannot
     be called against a not-yet-final epoch.
  5. Enter a plaintext integer and ElGamal-encrypt it with the collective
     public key (BabyJubJub, in-browser).
  6. Submit the ciphertext on-chain via
     `DKGManager.submitCiphertext(epochId, aid, ctIdx, c1x, c1y, c2x, c2y)` —
     the contract stores `keccak256(c1x,c1y,c2x,c2y)` and emits a
     `CiphertextSubmitted` event; nodes watch the event and produce their
     partial decryptions, then poll until the combined decryption lands.
  7. Verify that the recovered plaintext (readable via
     `getPlaintext(epochId, aid, ctIdx)`) matches the original input — the UI
     pops a green Alert on match.
- **Settings**: live-editable **RPC endpoint** override (stored in the
  browser's `localStorage`, per-user), debug-mode toggle, plus the chain
  / contract info from `/config.json` and the build version.

Errors anywhere in the UI offer a "Copy error report" button that bundles
the route, chain, wallet, build SHA and stack trace into a markdown blob
ready to paste into a GitHub issue.

### How it reaches the chain

On startup the browser fetches `/config.json` from the UI bundle; that
file is templated at **build time** from `RPC_URL` / `MANAGER_ADDRESS`
/ `CHAIN_ID` / `CHAIN_NAME` (and optional `REGISTRY_ADDRESS` /
`START_BLOCK`) and shipped inside `dist/`. All contract reads after that
go directly from the browser to the RPC via `viem`.

Ciphertext submission is fully on-chain — the SPA calls
`DKGManager.submitCiphertext` directly from the browser wallet; nodes
subscribe to `CiphertextSubmitted` events to pick up work.

### Running it

The Vite bundle is the unit of deployment — there is no runtime config
templating any more. Build once for each deployment with the chain
config baked in, then host the static files anywhere.

#### DigitalOcean App Platform (recommended for public deployments)

The repo ships an App Platform spec at
[`ui/.do/davinci-dkg-ui.yaml`](./ui/.do/davinci-dkg-ui.yaml). Deploy via:

```bash
doctl apps create --spec ui/.do/davinci-dkg-ui.yaml
```

App Platform clones the repo, builds `ui/Dockerfile` from the repo root
on every push to `main`, passes the spec's `BUILD_TIME` envs as
`--build-arg`s, and serves the resulting `dist/` from its edge — no
nginx in the loop. Edit the env values in the spec to retarget the
deployment at a different chain.

#### Compose (self-hosted on your own box)

```bash
# 1. Build the dist with the chain config you want.
make ui-build \
  RPC_URL=https://eth-sepolia.public.blastapi.io \
  MANAGER_ADDRESS=0x6683f889ce518945053f7d01abef7da842283078 \
  CHAIN_ID=11155111 CHAIN_NAME=sepolia

# 2. Serve it via stock nginx:alpine, bind-mounted from ui/dist.
docker compose --profile ui up                       # UI alone, on :8082
docker compose --profile node --profile ui up        # node + UI together
```

#### Plain Docker

```bash
docker build -f ui/Dockerfile \
  --build-arg RPC_URL=https://eth-sepolia.public.blastapi.io \
  --build-arg MANAGER_ADDRESS=0x6683f889ce518945053f7d01abef7da842283078 \
  --build-arg CHAIN_ID=11155111 \
  --build-arg CHAIN_NAME=sepolia \
  -t my-davinci-dkg-ui .

# The image is build-only — extract the dist and serve it yourself:
docker create --name extract my-davinci-dkg-ui
docker cp extract:/usr/share/nginx/html ./dist
docker rm extract
# Now serve ./dist with anything (Caddy, nginx, S3, Cloudflare R2, …).
```

#### Build-time knobs

| Var | Default | Notes |
|---|---|---|
| `RPC_URL` | Sepolia public RPC | JSON-RPC endpoint the browser will hit. |
| `MANAGER_ADDRESS` | Sepolia DKGManager | Required for any non-Sepolia deployment. |
| `CHAIN_ID` | `11155111` | EIP-155 chain id. |
| `CHAIN_NAME` | `sepolia` | Display name in the header. |
| `REGISTRY_ADDRESS` | (auto-derived) | DKGRegistry override. |
| `START_BLOCK` | (none) | Lower bound for getLogs scans. |
| `VITE_BUILD_VERSION` | `dev` | Shown in the Settings page. |
| `VITE_WALLETCONNECT_PROJECT_ID` | (none) | Without this, only injected wallets appear in the picker. |

The same knobs work for `make ui-dev` and `make ui-build` — they
re-render `ui/public/config.json` from the env when `RPC_URL` is set:

```bash
make ui-dev RPC_URL=http://127.0.0.1:8545 MANAGER_ADDRESS=0x... CHAIN_ID=31337 CHAIN_NAME=anvil
```

See [`ui/README.md`](./ui/README.md) for the full local development
reference.

---

## References

- [NI-DKG paper (eprint Paper)](https://eprint.iacr.org/2026/552)
- [DAVINCI voting protocol](https://davinci.vote)
- [Vocdoni Association](https://vocdoni.io)
