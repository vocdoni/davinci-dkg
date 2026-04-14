# DAVINCI DKG

**Non-Interactive Distributed Key Generation using Blockchain and ZK Proofs**

`davinci-dkg` is the Go implementation of the NI-DKG protocol described in the paper
[*NI-DKG: Non-Interactive Distributed Key Generation using Blockchain and ZK Proofs*](https://eprint.iacr.org/2026/552).
It provides the node service, cryptographic primitives, zk-SNARK circuits, and Solidity smart contracts
for threshold key generation, threshold decryption, and optional secret-key disclosure on EVM-compatible chains.

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
  - [Secret Key Disclosure](#secret-key-disclosure)
  - [Public-Input Compression (BRLC)](#public-input-compression-brlc)
- [ZK-SNARK Circuits](#zk-snark-circuits)
  - [Contribution Circuit](#contribution-circuit-dkg-phase-4)
  - [Finalize Circuit](#finalize-circuit-dkg-phase-5)
  - [PartialDecrypt Circuit](#partialdecrypt-circuit-decryption-phase-2)
  - [DecryptCombine Circuit](#decryptcombine-circuit-decryption-phase-3)
  - [RevealSubmit Circuit](#revealsubmit-circuit-disclosure-phase-2)
  - [RevealShare Circuit](#revealshare-circuit-disclosure-phase-3)
- [Smart Contracts](#smart-contracts)
  - [DKGRegistry](#dkgregistry)
  - [DKGManager](#dkgmanager)
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
  - [Running it outside the testnet](#running-it-outside-the-testnet)
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
   every time a round runs. You need enough native gas on the target
   network to cover those. Gas costs are bounded and documented in the
   [Gas Profile](#gas-profile) section; as a rule of thumb, budget a few
   million gas per round you expect to participate in. Testnet gas from a
   faucet is usually enough; mainnet deployments should hold a comfortable
   balance.
2. **The target network's JSON-RPC URL.** Any HTTPS or WSS endpoint that
   speaks the standard Ethereum JSON-RPC will work (Infura, Alchemy, your
   own node, a local Anvil instance, etc.).
3. **The deployed contract addresses**: `DKGRegistry` and `DKGManager` on
   the network you want to join. These are per-network and are published
   alongside the network's announcement. If you are bootstrapping your own
   network, deploy them first with [`make solidity-deploy`](#deploy-contracts).
4. **An operator private key** that controls the funded account. It is
   used only for signing, the node never exports or transmits it.

> the DKG round cadence and policy are decided by whoever
> creates rounds on `DKGManager`. As a node operator you react to rounds;
> you don't need to run the orchestrator.

### Configure the node

All three install options below share the same configuration surface: a
`.env` file at the repo root (or next to the binary). The node reads its
settings from environment variables — every CLI flag has an env-var
equivalent, e.g. `--web3.rpc` → `DAVINCI_DKG_WEB3_RPC`,
`--poll-interval` → `DAVINCI_DKG_POLL_INTERVAL`, `--webapp.listen` →
`DAVINCI_DKG_WEBAPP_LISTEN`.

```bash
cp .env.example .env
$EDITOR .env
```

At minimum, fill in:

```dotenv
DAVINCI_DKG_WEB3_RPC=https://your-rpc-endpoint
DAVINCI_DKG_PRIVKEY=0x<your-64-hex-char-operator-key>
DAVINCI_DKG_REGISTRY=0x<DKGRegistry address on this network>
DAVINCI_DKG_MANAGER=0x<DKGManager address on this network>
```

See `.env.example` for the full list and `davinci-dkg-node --help` for
defaults. **The embedded explorer UI** is on by default and listens on
`0.0.0.0:8081`. Set `DAVINCI_DKG_WEBAPP_ENABLED=false` to disable it or
restrict the bind address with `DAVINCI_DKG_WEBAPP_LISTEN=127.0.0.1:8081`
if you only want local access.

### Option A — Docker Compose (recommended)

The repo ships a ready-to-run `docker-compose.yml` at the root. It pulls
a prebuilt multi-stage image (webapp bundle + fully static Go binary on
`debian:bookworm-slim`) from `ghcr.io/vocdoni/davinci-dkg` and exposes the
explorer on `${DAVINCI_DKG_WEBAPP_PORT:-8081}`. This is the recommended
path for most operators: one command, no host-side toolchain, automatic
restart on failure, and automatic image upgrades via Watchtower.

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

The compose file defines two profiles:

| Profile | Services                       | Use case                                     |
|---------|--------------------------------|----------------------------------------------|
| `node`  | `node`, `watchtower`           | Long-running DKG node with auto-updates      |
| `test`  | `unit-test`, `integration-test`| Run the Go test suites inside a container    |

Pin `DAVINCI_DKG_TAG=v0.1.0` in `.env` (or remove the `watchtower`
service) if you want to control upgrades manually.

**Build the image yourself** instead of pulling from `ghcr.io`:

```bash
docker build -t davinci-dkg:local .
DAVINCI_DKG_TAG=local docker compose --profile node up -d
```

### Option B — Download a release binary

Every tagged release publishes fully-static `davinci-dkg-node` and
`dkg-runner` binaries for Linux (amd64 + arm64) and macOS (amd64 + arm64)
on the [**GitHub Releases**](https://github.com/vocdoni/davinci-dkg/releases)
page.

```bash
# Pick the archive that matches your OS/arch from the Releases page.
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

The binaries are self-contained: they embed the DKG explorer webapp
(`//go:embed`) and the `version.Version` string, so you do not need Node,
pnpm, or any build toolchain to run a node.

### Option C — Build from source

You will need **Go 1.25+**, **pnpm** (for the embedded webapp), and
optionally **Foundry** if you also want to rebuild contracts.

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg

# Build the webapp once so the embedded explorer is up-to-date, then build
# all Go binaries. The `make build` target does both steps.
make build

# The binaries are produced at the repo root:
./davinci-dkg-node --version
./dkg-runner --help
```

Or if you only want the node and don't care about the embedded UI:

```bash
go build -o davinci-dkg-node ./cmd/davinci-dkg-node
```

Note: the `webapp/dist/` directory must exist before the Go build because
of the `//go:embed all:dist` directive — `make build` guarantees that. If
you invoke `go build` directly, first run `make webapp-build`.

### Start the node

First-run behaviour (identical for all three install options):

1. The node derives a BabyJubJub encryption key from your Ethereum
   private key, reads its current registry row, and depending on what it
   finds either calls `registerKey` (first time), `updateKey` (key needs
   rotation or previous row was `INACTIVE`, which auto-reactivates), or
   skips the call entirely (already `ACTIVE` with the right key).
2. Immediately after, the node prints a verbose startup banner: local
   config (identity, RPC, contracts, poll interval, webapp settings),
   on-chain state (chain head, round prefix + nonce, `nodeCount`,
   `activeCount`, `INACTIVITY_WINDOW`), and its own registry row
   (`status`, `lastActiveBlock`, `blocksSinceActive`, and the remaining
   liveness budget before reap). Grepping the logs for `self:` is usually
   the fastest way to see whether you joined the network correctly.
3. From then on the node polls `DKGManager` for active rounds and
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

# Equivalent, overriding everything on the command line:
./davinci-dkg-node \
  --web3.rpc=https://your-rpc-endpoint \
  --privkey=0x<your-key> \
  --registry=0x<DKGRegistry> \
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

When the next round is created you will see `claiming slot`, `submitting
contribution`, and so on. Every phase emits a log line with the round ID
so you can trace progress against the explorer.

### Operational notes

- **Gas**: a full n=16 round costs ≈ 15.4 M gas spread across every
  committee member. Your node pays gas only on the phases it actually
  executes; see [Gas Profile](#gas-profile) for the per-call breakdown.
- **Upgrades**: replace the binary, restart the process. State that must
  persist across restarts (claimed rounds, private shares) lives on disk
  under `--datadir` (default `$HOME/.davinci-dkg`).
- **Multiple operators on the same host**: each node needs its own
  `--datadir`, private key, and port for the webapp. Keep one binary per
  operator to avoid confused state.
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
σ_i(j)    = s_i(j) + H_share(rid, i, j, r_{i,j} · pub_j)   (masked share, mod q)
```

where `H_share` is the Poseidon1 hash of the domain separator, round ID, participant indices, and
the shared secret point `r_{i,j} · pub_j = sk_j · R_i(j)`.

Decryption by participant `j`:
```
s_i(j) = σ_i(j) − H_share(rid, i, j, sk_j · R_i(j))   (mod q)
```

### DKG Protocol

The full DKG proceeds in 4 phases, all block-number delimited:

**Phase 1 — Initiation**: The organizer creates a round specifying `(t, n)` and policy
parameters. A unique 12-byte `roundId` is generated on-chain. The contract snapshots
`registry.nodeCount()`, derives a per-round **lottery threshold** so that on average
`α × n` nodes are eligible, and pins a `seedBlock = block.number + seedDelay` whose
future blockhash will become the round seed. The organizer's only job is calling
`createRound`; they do not pick the committee.

**Phase 2 — Trustless committee selection (lottery)**: Once `block.number ≥ seedBlock`,
any registered node calls `claimSlot(roundId)`. The first such call lazily resolves
`seed = blockhash(seedBlock)`. A node is eligible iff
`keccak256(seed ‖ msg.sender) < lotteryThreshold`. Eligible nodes race
**first-come-first-served** until `committeeSize` slots are filled, at which point the
contract snapshots the committee key hash and transitions to Contribution. If the
round stalls past `registrationDeadlineBlock`, anyone can call
`extendRegistration(roundId)` to reroll the seed and reopen the lottery.

**Phase 3 — Main DKG (contribution)**: Each participant `i` samples random polynomial coefficients
`{a_{i,k}}` and encryption nonces `{r_{i,j}}`, then publishes:
- Commitments: `C_i(k) = a_{i,k} · G` for `k ∈ {0, …, t−1}`
- Encrypted shares: `(R_i(j), σ_i(j))` for all `j ∈ [n]`
- A Groth16 proof `π_i` of correctness (see [Contribution Circuit](#contribution-circuit-dkg-phase-4))

The contract rejects the transaction if the proof is invalid.

**Phase 4 — Finalization**: Once `minValidContributions` contributions are accepted, anyone may call
`finalizeRound`. This computes and persists:
- Aggregate commitments: `C̄(k) = Σ_{ℓ ∈ I} C_ℓ(k)` for `k ∈ {0, …, t−1}`
- Collective public key: `PK = C̄(0) = F(0) · G`
- Share commitments: `D_i = Σ_k i^k · C̄(k) = F(i) · G` for each accepted participant `i`

Finalization is also proof-gated (see [Finalize Circuit](#finalize-circuit-dkg-phase-5)).

Each participating node privately computes their secret share:
```
d_i = Σ_{ℓ ∈ I} s_ℓ(i)  =  F(i)
```
by decrypting the encrypted shares they received on-chain.

### Trustless Lottery Committee Selection

Given a round with policy parameters `(n, α)` — where `n = committeeSize` and
`α ∈ (0, 1]` (encoded as `lotteryAlphaBps`, basis points out of 10 000) — and
the registry snapshot `R = nodeCount()` at the moment of round creation, the
contract computes:

```
lotteryThreshold = ⌊ (α · n · 2²⁵⁶) / R ⌋
```

This is the **eligibility threshold**: a pseudo-random 256-bit value uniformly
derived from the future seed and the node's address must fall below
`lotteryThreshold` for that node to claim a slot. By construction, the
expected number of eligible nodes is `E[|eligible|] = α · n`. With `α = 1.0`
(the testnet default) the expectation equals the committee size; with `α > 1.0`
(not currently supported — the contract clamps to 10 000 bps) one would
oversubscribe to absorb liveness failures. In the current build a round that
fails to fill reopens the lottery via `extendRegistration`, which captures a
fresh blockhash and resets the deadline.

**Seeding.** At `createRound(seedDelay)` the contract pins `seedBlock =
block.number + seedDelay` but does **not** yet know the seed. Once
`block.number ≥ seedBlock`, the first call to `claimSlot` reads
`blockhash(seedBlock)` and stores it as `seed`. Binding the seed to a future
blockhash (`seedDelay ≥ 1`) prevents the organizer from tuning the eligibility
set by picking a favourable `createRound` block — they cannot predict the
future blockhash.

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
round from Registration to Contribution. Any further `claimSlot` calls revert.
The committee snapshot is what later contribution proofs are verified against
— the `contributionVerifier` only accepts a proof whose recipient list keccak
matches this snapshot, so the committee is effectively locked in a single slot
of storage.

**Security properties.**
- *No organizer influence over membership.* The organizer sets `n` and `α` but
  cannot prefer specific nodes: the eligibility set is determined by a
  blockhash published after `createRound`.
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
  claim before `registrationDeadlineBlock`, `extendRegistration` reseeds and
  reopens — no round is stuck waiting for a node that went offline.

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
their key with `updateKey`). The per-round cost of this mechanism is a
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

### Secret Key Disclosure

When the round policy allows it, the secret key can be reconstructed openly.
Each participant `i` reveals `d_i` along with a proof that `d_i · G = D_i`.
Any `t` revealed shares suffice to reconstruct `sk = F(0)` via Lagrange interpolation.

### Public-Input Compression (BRLC)

Because Groth16 on BN254 costs ~6,650 gas per public input, large transcripts
(commitment vectors, encrypted shares, partial decryptions) are compressed using
**Binding Random Linear Combinations (BRLC)**:

```
C = Σ_{i=1}^{l} ρ^i · v_i
```

The challenge `ρ` is derived from the round ID and a domain separator using `keccak256`, making it
unpredictable at the time the inputs are committed (Fiat-Shamir). The in-circuit check recomputes
the BRLC and asserts equality, reducing `l` public inputs to a single scalar.

On-chain cost: ~70 gas per element (vs. ~6,650 gas per Groth16 public input).

---

## ZK-SNARK Circuits

All circuits use **Groth16** on **BN254**. BabyJubJub curve operations are performed natively
(inside the BN254 scalar field). Hashing uses **Poseidon1**.

Fixed-size circuit arrays use prefix masks derived from the actual
threshold/committee size, so one compiled circuit serves all parameter choices
up to the compile-time maximum. The bound is the single Go constant
`circuits/common.MaxN`, which all four circuit-side aliases
(`MaxCoefficients` / `MaxRecipients` / `MaxParticipants` / `MaxShares`) reference.
The Solidity contract reads the same value from `DKGManager.sol::MAX_N`. The
default build uses `MaxN = 16`; see `BENCHMARKS.md` for the side-by-side gas
and proof-time comparison against a `MaxN = 32` build.

### Contribution Circuit (DKG Phase 4)

**Package**: `circuits/contribution`  
**Constraints**: ~802 k (MaxN=16) / ~2.84 M (MaxN=32) — `O(MaxN²)`  
**Public inputs** (8 scalars): `RoundHash`, `Threshold`, `CommitteeSize`, `ContributorIndex`,
`CommitmentHash`, `ShareHash`, `Challenge`, `TranscriptCommitment`

**Private inputs** (112 scalars): polynomial coefficients, encryption nonces, Shamir shares,
mask quotients, share masks, carry bits

**Proves**:
1. Coefficient commitments: `C_i(k) = a_{i,k} · G` for all `k ∈ {0, …, t−1}`
2. Shamir evaluation: `s_i(j) = Σ_k a_{i,k} · j^k` for all `j ∈ [n]`
3. Feldman verification: `s_i(j) · G = Σ_k j^k · C_i(k)` for all `j ∈ [n]`
4. Ephemeral key: `R_i(j) = r_{i,j} · G` for all `j ∈ [n]`
5. Share encryption: `σ_i(j) = s_i(j) + H_share(rid, i, j, r_{i,j} · pub_j) (mod l)`
6. Commitment hash: `CommitmentHash = Poseidon1(RoundHash, ContributorIndex, t, C_i(0), …)`
7. Share hash: `ShareHash = Poseidon1(RoundHash, ContributorIndex, n, idx_1, R_1, σ_1, …)`
8. BRLC transcript: `TranscriptCommitment = BRLC(Challenge, transcript_vector)`

The transcript vector encodes all commitments, recipient indexes, recipient public keys,
ephemeral points, and masked shares.

### Finalize Circuit (DKG Phase 5)

**Package**: `circuits/finalize`  
**Constraints**: ~626 k (MaxN=16) / ~2.49 M (MaxN=32) — `O(MaxN²)`  
**Public inputs** (9 scalars): `RoundHash`, `Threshold`, `CommitteeSize`, `AcceptedCount`,
`AggregateHash`, `CollectivePublicKey`, `ShareCommitmentHash`, `Challenge`, `TranscriptCommitment`

**Private inputs** (168 scalars): participant indexes, per-participant commitment vectors,
aggregate commitments, share commitments

**Proves**:
1. Aggregate commitments: `C̄(k) = Σ_{ℓ ∈ I} C_ℓ(k)` for all `k`
2. Public key hash: `CollectivePublicKey = Poseidon1(RoundHash, C̄(0).X, C̄(0).Y)`
3. Aggregate hash: `AggregateHash = Poseidon1(RoundHash, t, n, |I|, C̄(0), …)`
4. Share commitments: `D_i = Σ_k i^k · C̄(k)` for each accepted `i ∈ I`
5. Share commitment hash: `ShareCommitmentHash = Poseidon1(RoundHash, t, n, |I|, i_1, D_1, …)`
6. BRLC transcript: covers all participant indexes, contribution commitments, aggregate commitments,
   and share commitments

### PartialDecrypt Circuit (Decryption Phase 2)

**Package**: `circuits/partialdecrypt`  
**Constraints**: ~20,361 (independent of MaxN)  
**Public inputs** (13 scalars): `RoundHash`, `ParticipantIndex`, `Base.X`, `Base.Y`,
`PublicKey.X`, `PublicKey.Y`, `Delta.X`, `Delta.Y`, `A1.X`, `A1.Y`, `A2.X`, `A2.Y`, `Response`

**Private inputs** (2 scalars): `Secret` (`d_i`), `Nonce` (`r`)

**Proves** a Chaum-Pedersen DLEQ relation:
1. `PublicKey = Secret · G` (commitment to secret: `D_i = d_i · G`)
2. `Delta = Secret · Base` (partial decryption: `δ_i = d_i · C_1`)
3. `A1 = Nonce · G` (nonce commitment)
4. `A2 = Nonce · Base` (nonce commitment on base)
5. Challenge: `e = Poseidon1(domain, D_i, C_1, δ_i, A_1, B_1)` (Fiat-Shamir)
6. Response equations: `Response · G = A1 + e · PublicKey`  and  `Response · Base = A2 + e · Delta`

### DecryptCombine Circuit (Decryption Phase 3)

**Package**: `circuits/decryptcombine`  
**Constraints**: ~43,635 (MaxN=16) / ~84,314 (MaxN=32) — `O(MaxN)`  
**Public inputs** (7 scalars): `RoundHash`, `Threshold`, `ShareCount`,
`CombineHash`, `PlaintextHash`, `Challenge`, `TranscriptCommitment`

**Private inputs** (37 scalars): `CiphertextC1`, `CiphertextC2`, `Plaintext`,
participant indexes, partial decryption points, pre-computed Lagrange coefficients

**Proves**:
1. Combine hash: `CombineHash = Poseidon1(RoundHash, t, |Q|, C1, C2, idx_1, δ_1, …)`
2. Plaintext binding: `PlaintextHash = Plaintext` (the scalar `m` is exposed directly)
3. Lagrange combination: `Δ = Σ_{k ∈ [t]} λ_k · δ_{x_k}`
4. ElGamal decryption: `Plaintext · G + Δ = C_2`
5. BRLC transcript: covers ciphertext, participant indexes, partial decryption points

Lagrange coefficients `λ_k` are pre-computed natively in the BabyJubJub scalar field
(`r_bjj`) and passed as private witnesses. Computing them in-circuit via `api.Div` would
use `BN254.Fr` arithmetic, which gives incorrect results for negative coefficients because
`BN254.Fr − 1 ≠ r_bjj − 1` as BJJ scalars. The `Plaintext · G + Δ = C_2` constraint
implicitly validates that the witnesses were used correctly.

### RevealSubmit Circuit (Disclosure Phase 2)

**Package**: `circuits/revealsubmit`  
**Constraints**: ~2,346 (independent of MaxN)  
**Public inputs** (5 scalars): `RoundHash`, `ParticipantIndex`, `ShareValue`, `ShareCommitment.X`, `ShareCommitment.Y`

**Private inputs**: none

**Proves**: `ShareValue · G = ShareCommitment`, i.e., `d_i · G = D_i`

This is the simplest circuit — it proves knowledge of the discrete log of the published share
commitment without any additional private witnesses.

### RevealShare Circuit (Disclosure Phase 3)

**Package**: `circuits/revealshare`  
**Constraints**: ~1,904 (MaxN=16) / ~3,342 (MaxN=32) — `O(MaxN)`  
**Public inputs** (7 scalars): `RoundHash`, `Threshold`, `ShareCount`,
`DisclosureHash`, `ReconstructedSecretHash`, `Challenge`, `TranscriptCommitment`

**Private inputs** (17 scalars): `ReconstructedSecret`, participant indexes, revealed shares

**Proves**:
1. Disclosure hash: `DisclosureHash = Poseidon1(RoundHash, t, |Q|, idx_1, d_1, …)`
2. Lagrange reconstruction: `ReconstructedSecret = Σ_{k ∈ [t]} λ_k · d_{x_k}  (mod l)`
3. Secret binding: `ReconstructedSecretHash = ReconstructedSecret`
4. BRLC transcript: covers participant indexes and revealed shares

---

## Smart Contracts

The Solidity workspace lives in `solidity/` (Foundry, `solc 0.8.28`, EVM Cancun, `via_ir = true`).

### DKGRegistry

**Source**: `solidity/src/DKGRegistry.sol`  
**Interface**: `solidity/src/interfaces/IDKGRegistry.sol`

Stores the share-encryption public keys (BabyJubJub points) of eligible operator nodes.

| Function | Description |
|---|---|
| `registerKey(pubX, pubY)` | Register caller's BabyJubJub public key. Reverts if already registered or coordinates are zero. Increments `nodeCount`. |
| `updateKey(pubX, pubY)` | Update caller's previously registered key. |
| `getNode(operator)` | Returns the `NodeKey` struct `{operator, pubX, pubY, status}` for the given address. |
| `nodeCount()` | Returns the number of distinct addresses that have ever called `registerKey`. Snapshotted by DKGManager at `createRound` to derive the lottery threshold. |

**Events**: `KeyRegistered(address indexed operator, uint256 pubX, uint256 pubY)`,
`KeyUpdated(address indexed operator, uint256 pubX, uint256 pubY)`

### DKGManager

**Source**: `solidity/src/DKGManager.sol`  
**Interface**: `solidity/src/interfaces/IDKGManager.sol`

Owns the complete round lifecycle: creation, trustless lottery-based committee
selection, proof-gated contribution, finalization, threshold decryption, and
optional secret disclosure. Each state-mutating operation that involves
cryptographic claims is gated by a Groth16 verifier.

#### Round Lifecycle

```
Created → Registration (lottery) → Contribution → Finalized → Completed
                                                            ↘ Aborted
```

The contract retains a fixed-size ring buffer of the most recent `ROUND_HISTORY_SIZE`
(64) round IDs. When a new round is created and the buffer is full, the oldest
live round's storage is wiped (`delete rounds[…]`, etc.), keeping long-term
storage bounded. Off-chain consumers reconstruct historical round data from the
event log.

#### State-Mutating Functions

| Function | Phase | Access | Description |
|---|---|---|---|
| `createRound(threshold, committeeSize, minValidContributions, lotteryAlphaBps, seedDelay, registrationDeadlineBlock, contributionDeadlineBlock, disclosureAllowed)` | Any | Open | Create a new DKG round. Snapshots `nodeCount` from the registry and derives the per-round lottery threshold. Pins `seedBlock = block.number + seedDelay`. Returns `bytes12 roundId`. |
| `claimSlot(roundId)` | Registration | Any registered eligible node | First-come-first-served self-claim. The first call after `block.number ≥ seedBlock` lazily resolves `seed = blockhash(seedBlock)`. The caller is admitted iff `keccak256(seed ‖ msg.sender) < lotteryThreshold`. The contract stops accepting claims once `committeeSize` slots are filled and immediately advances to Contribution. |
| `extendRegistration(roundId)` | Registration, after deadline | Open | Reroll the seed if the round failed to fill in its registration window. Captures a fresh `blockhash` and pushes the deadline forward. |
| `submitContribution(roundId, contributorIndex, commitmentsHash, encryptedSharesHash, transcript, proof, input)` | Contribution | Selected participant | Submit polynomial commitments and encrypted shares with a Groth16 proof. The committee membership / pubkey list is verified against a single keccak snapshot taken when the lottery filled (no per-recipient registry calls). |
| `finalizeRound(roundId, aggregateCommitmentsHash, collectivePublicKeyHash, shareCommitmentHash, transcript, proof, input)` | After min contributions | Open | Aggregate commitments, publish collective public key and share commitments. Advances to Finalized. The transcript is read directly from calldata; share commitments are stored as `keccak256(x,y)` (1 slot each). |
| `submitPartialDecryption(roundId, participantIndex, ciphertextIndex, deltaHash, proof, input)` | Finalized | Selected participant | Submit a partial decryption `δ_i = d_i · C_1` with a Chaum-Pedersen DLEQ proof. Keyed by `(roundId, participant, ciphertextIndex)` to support multiple ciphertexts per round. |
| `combineDecryption(roundId, ciphertextIndex, combineHash, plaintextHash, transcript, proof, input)` | Finalized | Open | Combine `t` partial decryptions via Lagrange interpolation. Emits the recovered plaintext hash. |
| `submitRevealedShare(roundId, participantIndex, shareValue, proof, input)` | Finalized, disclosureAllowed | Selected participant | Reveal secret share `d_i` with a proof that `d_i · G = D_i`. |
| `reconstructSecret(roundId, disclosureHash, reconstructedSecretHash, transcript, proof, input)` | Finalized, disclosureAllowed, ≥t reveals | Open | Reconstruct `sk = F(0)` via Lagrange interpolation. Advances to Completed. |
| `abortRound(roundId)` | Any non-terminal | Organizer | Abort the round. Advances to Aborted. |

#### View Functions

| Function | Returns |
|---|---|
| `getRound(roundId)` | `Round` struct: organizer, policy, status, nonce, seedBlock, seed, lotteryThreshold, claimedCount, contributionCount, partialDecryptionCount, revealedShareCount |
| `selectedParticipants(roundId)` | `address[]` — ordered committee in claim order |
| `getContribution(roundId, contributor)` | `ContributionRecord` (only `contributorIndex`, `commitmentVectorDigest`, `accepted` are persisted; the rest live in `ContributionSubmitted` events) |
| `getPartialDecryption(roundId, participant, ciphertextIndex)` | `PartialDecryptionRecord` (only `participantIndex`, `ciphertextIndex`, `accepted`, `delta` are persisted) |
| `getCombinedDecryption(roundId, ciphertextIndex)` | `CombinedDecryptionRecord` (only `completed` is persisted; hashes live in `DecryptionCombined` events) |
| `getRevealedShare(roundId, participant)` | `RevealedShareRecord` (only `participantIndex`, `shareValue`, `accepted` persisted) |
| `getShareCommitmentHash(roundId, participantIndex)` | `bytes32` = `keccak256(abi.encode(x, y))`. The pre-image lives in the `RoundFinalized` event. |
| `getContributionVerifierVKeyHash()` | `bytes32` |
| `getPartialDecryptVerifierVKeyHash()` | `bytes32` |
| `getFinalizeVerifierVKeyHash()` | `bytes32` |
| `getDecryptCombineVerifierVKeyHash()` | `bytes32` |
| `getRevealSubmitVerifierVKeyHash()` | `bytes32` |
| `getRevealShareVerifierVKeyHash()` | `bytes32` |

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
# 1. Compile all 6 circuits; write artifacts and update Solidity verifier stubs
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

See [BENCHMARKS.md](BENCHMARKS.md).

---

## Switching `MaxN`

Changing the maximum committee size is a **two-line edit**:

```go
// circuits/common/sizes.go
const MaxN = 16   // 16 or 32 (or any other value)
```

```solidity
// solidity/src/DKGManager.sol
uint256 internal constant MAX_N = 16;   // must equal circuits/common.MaxN
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
make testnet-up                                # 3 nodes + anvil + deployer + dkg-webapp
make testnet-up DKG_NODE_COUNT=8 DKG_THRESHOLD=5

# Expose the browser-side RPC URL if you will access from another host:
WEBAPP_PUBLIC_RPC=http://<host-ip>:8545 make testnet-up
```

Once the command returns, open `http://<host-ip>:8081/` in a browser.

The equivalent raw compose invocation:

```bash
cd testnet
DKG_NODE_COUNT=3 DKG_THRESHOLD=2 \
  docker compose up --scale dkg-node=3 --build
```

### Run the scenario

In a second terminal:

```bash
make testnet-run                                         # defaults: 3 nodes, t=2, disclosure off
make testnet-run DKG_NODE_COUNT=8 DKG_THRESHOLD=5        # custom sizing
make testnet-run DKG_DISCLOSURE_ALLOWED=true             # enable reveal-share phase

# or bypass the Makefile:
cd testnet && docker compose run --rm dkg-runner
```

The runner will:
1. Create a DKG round with `--nodes` committee size and `--threshold` decryption threshold
2. Wait until the lottery committee fills (each node self-claims via `claimSlot` once
   it sees the round and verifies it's eligible — no organizer participation needed)
3. Wait until ≥ threshold nodes submit their contributions
4. Submit the finalize proof (aggregates Feldman commitments, publishes the collective public key)
5. Encrypt a random test message `m` as `(C1, C2) = (r·G, m·G + r·PK)`
6. Write the ciphertext to the shared volume so nodes can compute partial decryptions
7. Wait until ≥ threshold nodes submit partial decryptions with DLEQ proofs
8. Combine the partial decryptions via Lagrange interpolation and submit the combine proof
9. Verify that the on-chain recovered `m` matches the original


### Configuration

| Env variable | Default | Description |
|---|---|---|
| `DKG_NODE_COUNT` | `3` | Number of DKG node replicas to start |
| `DKG_THRESHOLD` | `2` | Decryption threshold (`t`-of-`n`) |
| `DKG_DISCLOSURE_ALLOWED` | `false` | Pass `disclosureAllowed=true` to `createRound` so the reveal-share / secret reconstruction phase is enabled |
| `DKG_RUNNER_NODES` | `3` | Committee size seen by the runner (same as `DKG_NODE_COUNT`) |
| `DKG_RUNNER_THRESHOLD` | `2` | Decryption threshold seen by the runner |
| `DKG_RUNNER_DISCLOSURE_ALLOWED` | `false` | Forwarded from `DKG_DISCLOSURE_ALLOWED` to the runner flag |
| `ANVIL_PORT` | `8545` | Host port for the Anvil RPC (bound on `0.0.0.0`) |
| `DEPLOYER_PORT` | `8888` | Host port for the deployer HTTP server |
| `WEBAPP_PORT` | `8081` | Host port the DKG explorer listens on (bound on `0.0.0.0`) |
| `WEBAPP_PUBLIC_RPC` | `http://localhost:8545` | RPC URL advertised to browsers in `/config.json`. Override with the LAN/public IP of the host when accessing the explorer from a remote machine. Can also be changed live in the explorer's Settings page. |

---

## Web Explorer

`webapp/` contains a single-page React application — a read-only
block-explorer for a live `DKGManager` / `DKGRegistry` pair. It is a Vite +
React + TypeScript + Chakra UI + React Query stack talking to the chain
directly via `viem`, and it is **embedded into the `davinci-dkg-node` binary
via `//go:embed`**, so any running node can serve the UI with no external
assets.

In the testnet the explorer is available at `http://<host>:8081/` as soon as
`make testnet-up` returns — the `dkg-webapp` compose service starts
`davinci-dkg-node` in idle mode (no private key, no participation) with
`--webapp.enabled=true --webapp.listen=0.0.0.0:8081` and passes the contract
addresses pulled from `/addresses/addresses.env`.

### What it shows

- **Overview**: total rounds, registered nodes, latest block, chain ID, and
  the 5 most recent rounds.
- **Rounds**: tabular view of the ring-buffered on-chain history (up to 64
  rounds), status, organizer, committee/threshold, claim / contribution
  counters. Click a row to open the round detail view.
- **Round detail**: policy, seed block + seed, lottery threshold, all four
  phase progress bars (claims, contributions, partial decryptions, revealed
  shares), the full list of selected participants, and every decoded event
  touching the round (`RoundCreated`, `SeedResolved`, `SlotClaimed`,
  `ContributionSubmitted`, `RoundFinalized`, `PartialDecryptionSubmitted`,
  `DecryptionCombined`, `RevealedShareSubmitted`, `SecretReconstructed`,
  `RoundEvicted`, `RoundAborted`).
- **Registry**: every registered operator and their BabyJubJub public key
  coordinates.
- **Settings**: live-editable **RPC endpoint** override (stored in the
  browser's `localStorage`, per-user) plus the chain / contract info from
  `/config.json`.

### How it reaches the chain

On startup the browser fetches `/config.json` from whichever node is serving
it; the response carries the RPC URL, manager + registry addresses, chain ID
and chain name. All contract reads after that go **directly** from the
browser to the RPC via `viem`, so the Go node does no proxying and does not
need any contract bindings for the UI to work. Polling cadence is 4 s.

Because the RPC URL is a runtime field (not baked into the bundle), the same
embedded SPA works against local, testnet, or public deployments without any
rebuild — override `WEBAPP_PUBLIC_RPC` for the compose service or flip
the endpoint in the Settings page.

### Running it outside the testnet

Any `davinci-dkg-node` instance serves the webapp by default on `:8081`:

```bash
davinci-dkg-node \
  --web3.rpc=http://127.0.0.1:8545 \
  --manager=0x... --registry=0x... \
  --privkey=0x... \
  --webapp.listen=0.0.0.0:8081 \
  --webapp.public-rpc=http://<your-host>:8545     # URL the browser will use
```

Flags:

| Flag | Env | Default | Description |
|---|---|---|---|
| `--webapp.enabled` | `DAVINCI_DKG_WEBAPP_ENABLED` | `true` | Turn the embedded UI on/off |
| `--webapp.listen` | `DAVINCI_DKG_WEBAPP_LISTEN` | `0.0.0.0:8081` | Address/port the explorer HTTP server binds to |
| `--webapp.public-rpc` | `DAVINCI_DKG_WEBAPP_PUBLIC_RPC` | first `--web3.rpc` | Browser-reachable RPC URL advertised in `/config.json` |

---

## References

- [NI-DKG paper (eprint Paper)](https://eprint.iacr.org/2026/552)
- [DAVINCI voting protocol](https://davinci.vote)
- [Vocdoni Association](https://vocdoni.io)
