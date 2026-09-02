# DAVINCI DKG

**Non-Interactive Distributed Key Generation on EVM chains.**

Reference Go implementation of the protocol described in
[*NI-DKG: Non-Interactive Distributed Key Generation using Blockchain and ZK Proofs*](https://eprint.iacr.org/2026/552).
Built as the threshold-key layer for the [DAVINCI](https://davinci.vote) voting system, but the
protocol is generic — any application that needs a `t`-of-`n` collective public key on an EVM chain
can use it.

The protocol replaces interactive complaint rounds with Groth16 ZK proofs: every contribution,
finalization, and partial decryption is verified at transaction time. There is no dispute phase.

---

## Contents

- [What you get](#what-you-get)
- [Protocol model](#protocol-model)
  - [Epoch lifecycle](#epoch-lifecycle)
  - [Lottery committee selection](#lottery-committee-selection)
  - [Threshold decryption](#threshold-decryption)
  - [Per-application keys](#per-application-keys)
- [On-chain surface](#on-chain-surface)
- [Integrating](#integrating)
  - [Run a node](#run-a-node)
  - [TypeScript SDK](#typescript-sdk)
  - [Encrypting and decrypting](#encrypting-and-decrypting)
- [Deployments](#deployments)
- [Build from source](#build-from-source)
- [References](#references)

---

## What you get

| Component         | Path                          | Purpose                                                          |
|-------------------|-------------------------------|------------------------------------------------------------------|
| `davinci-dkg-node`| `cmd/davinci-dkg-node`        | Node binary — joins the active set and reacts to every epoch     |
| Solidity contracts| `solidity/src`                | `DKGRegistry` + `DKGManager` + `DKGAppManager`                   |
| Circuits          | `circuits/`                   | Groth16 / BN254 — Contribution, Finalize, PartialDecrypt, Combine|
| TypeScript SDK    | `sdk/`                        | `@vocdoni/davinci-dkg-sdk` — read client, writer, encryption     |
| Web explorer / UI | `ui/`                         | React SPA + interactive playground                               |

Crypto primitives: BabyJubJub on the BN254 scalar field; Poseidon1 for in-circuit hashing;
keccak256 for on-chain Fiat–Shamir challenges. ElGamal for share and ciphertext encryption.

---

## Protocol model

### Epoch lifecycle

An **epoch** is one DKG run. It produces a single collective public key `PK_ep` shared by `n`
committee members; any `t` of them can decrypt. Epochs are scheduled at a fixed cadence — every
`EPOCH_DURATION_BLOCKS` blocks (set per-deploy as a `DKGManager` immutable).

Each epoch splits into two top-level phases:

```
   startBlock                                                                 endBlock
   │                                                                          │
   │ ─── Preparation (small, fixed) ────► ◄────── Service (the rest) ────────│
   │                                                                          │
   │ CommitteeSelection │ KeyAssembly │ gap │            Live                 │
   │     ~5 min         │   ~5 min    │~1min│   (whatever is left)            │
   ▼                                                                          ▼
   ├────────────────────┼─────────────┼─────┼─────────────────────────────────┤
   │  claimSlot         │submitContrib│ ... │ registerApplication[CoDec] /    │
   │  (lottery)         │  (Groth16)  │     │ submitCiphertext /              │
   │                    │             │     │ submitPartialDecryption /       │
   │                    │             │     │ combineDecryption               │
   └────────────────────┴─────────────┴─────┴─────────────────────────────────┘
                                       ▲
                             finalizeEpoch (Groth16)
                             flips state from KeyAssembly → Live

       ◄──────────────── EPOCH_DURATION_BLOCKS ─────────────────────────►
```

- **Preparation** — committee is assembled and the collective key `PK_ep` is generated. Three
  contiguous block windows:
  - `CommitteeSelection`: lottery via `claimSlot` picks `n` operators.
  - `KeyAssembly`: each committee member submits a Feldman VSS contribution with a Groth16 proof.
  - Finalize gap: short window before `finalizeEpoch` may run.
- **Service** — `PK_ep` is live for the rest of the epoch.
  - Apps register their derived keys via `registerApplication[CoDec]`.
  - Anyone can `submitCiphertext`; the committee posts partials and any caller `combineDecryption`s
    to land the recovered plaintext on chain.

Each Preparation window is an **absolute** block count, not a fraction of the epoch — the lottery
is one keccak per claimer and the contribution proof is one tx per committee member, so a fixed
budget is the right shape. The four block constants are deploy-time immutables (defaults in
`solidity/src/libraries/Sizes.sol`, overridable via `EPOCH_DURATION_BLOCKS`,
`COMMITTEE_SELECTION_BLOCKS`, `KEY_ASSEMBLY_BLOCKS`, `FINALIZE_GAP_BLOCKS` env vars at deploy
time). Long epochs (multi-day) keep the same short Preparation; the extra time falls into Service.

The epoch stays `Live` for the entire Service window — its key remains usable while the next epoch
bootstraps.

`createEpoch` is **permissionless** but cadence-gated: it reverts unless
`block.number >= nextEpochStartBlock()`. In production, every node races to fire it once the
window opens (random jitter, env-toggleable). Only the first call lands; the others revert
cheaply.

States exposed by `EpochPhase`: `None`, `CommitteeSelection`, `KeyAssembly`, `Live`, `Aborted`.
A reserved `Completed` value exists but is not used in the live state machine.

### Lottery committee selection

Each epoch picks a fresh committee from the registry by trustless lottery — no organizer can
prefer specific operators. Inputs:

- `n` = `committeeSize`, `α` = `lotteryAlphaBps / 10_000` (oversubscription factor)
- `R` = `registry.activeCount()` snapshotted at `createEpoch`
- `seed` = `blockhash(startBlock + SEED_DELAY_BLOCKS)`, resolved on the first `claimSlot` call

A node is eligible iff:

```
keccak256(seed ‖ msg.sender) < (α · n · 2²⁵⁶) / R
```

Eligible nodes race first-come-first-served until `n` slots are filled, at which point the epoch
auto-advances to `KeyAssembly`. Anyone can recompute eligibility by replaying the keccak — no ZK
proof, no trusted coordinator. Only operators registered *before* `createEpoch` may claim (the
registry is snapshotted with `R`, so fresh identities cannot be ground against a revealed seed).
If the committee fails to fill within `CommitteeSelection`, the epoch is dead: anyone may record
the abort and the next scheduled epoch opens automatically. An epoch that can still be finalized
cannot be aborted by anybody.

Whoever wins the `createEpoch` race chooses `(t, n, minValidContributions, α)`, so the deployment
pins floors: `MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE` and a ceiling `MAX_LOTTERY_ALPHA_BPS`
(constructor immutables, `MIN_THRESHOLD`/`MIN_COMMITTEE_SIZE`/`MAX_LOTTERY_ALPHA_BPS` env vars at
deploy time).

### Threshold decryption

Once the epoch is `Live`, the collective public key `PK_ep` is on-chain. To decrypt an ElGamal
ciphertext `(C₁, C₂)` published via `submitCiphertext`:

1. Each selected node `i` publishes its partial `δ_i = d_i · C₁` plus a Chaum–Pedersen DLEQ proof
   binding `δ_i` to its share commitment `D_i` and the ciphertext `C₁`.
2. Once `t` partials are accepted, anyone calls `combineDecryption`. A Groth16 proof attests that
   `Σ λ_k · δ_k` Lagrange-interpolates correctly and that `plaintext · G + Δ = C₂`.
3. The recovered scalar `m` is stored on-chain and readable via `getPlaintext`.

The combine proof discovers `m` via baby-step / giant-step (BSGS) discrete-log inversion. The
committee node uses a 2⁵⁰ cap (~1 GB table); the SDK uses 2³² (~16 MB) so it can run in a browser.
Submitting a plaintext above the relevant cap is unrecoverable.

### Per-application keys

A `Live` epoch can host many independent encryption contexts — one per **application**, keyed
by a 32-byte `aid` chosen by the integrator. `aid` is bound into every decryption proof as a
BN254 scalar-field public input, so it must be non-zero and below the field modulus
(clear the top three bits of a random or hashed id); the contract rejects other values. `DKGAppManager` exposes two registration modes:

- **Mode 0 — public derivation** (`registerApplication`). The contract derives
  `S = keccak256(epochId ‖ PK_ep ‖ aid) mod q`. The implicit per-app key is `PK_aid = PK_ep + S·G`,
  which any reader can recompute. Threshold decryption uses the committee alone.

- **Mode 1 — organizer co-decryption** (`registerApplicationCoDec`). The integrator publishes
  `PK_org = sk_org · G` plus a Schnorr proof of knowledge of `sk_org`. The implicit per-app key is
  `PK_aid = PK_ep + PK_org`. Decryption now requires both the committee and the organizer to
  cooperate; the latter contributes `Δ_org = sk_org · C₁` via `submitOrganizerShare`.

The legacy "epoch-key only" path is the special case `aid = bytes32(0)`: open submission,
no `AppPolicy`.

---

## On-chain surface

Three contracts. Deploy order: `DKGRegistry → DKGManager → DKGAppManager`, then wire with
`DKGRegistry.setManager(...)` and `DKGManager.setAppManager(...)`. The split exists only to keep
each contract under EIP-170; logically `DKGManager` and `DKGAppManager` share one storage.

| Contract        | Owns                                                                                    |
|-----------------|-----------------------------------------------------------------------------------------|
| `DKGRegistry`   | Operator identities (BabyJubJub pub keys), liveness (`heartbeat`, `reactivate`, `reap`) |
| `DKGManager`    | Epoch lifecycle: `createEpoch`, `claimSlot`, `submitContribution`, `finalizeEpoch`, ciphertexts, partial / combined decryption |
| `DKGAppManager` | Per-application registration: `registerApplication[CoDec]`, `submitOrganizerShare`     |

Read the `solidity/src/interfaces/*.sol` files for the full method signatures and event schemas —
they are the integration contract.

A few load-bearing knobs:

| Constant                       | Where                                | Default                        | Notes                                                                |
|--------------------------------|--------------------------------------|--------------------------------|----------------------------------------------------------------------|
| `EPOCH_DURATION_BLOCKS`        | `DKGManager` constructor (immutable) | `100` (~20 min @ 12 s)         | Cadence anchor: next epoch can start `EPOCH_DURATION_BLOCKS` after the previous one |
| `COMMITTEE_SELECTION_BLOCKS`   | `DKGManager` constructor (immutable) | `25` (~5 min @ 12 s)           | Absolute lottery window length                                       |
| `KEY_ASSEMBLY_BLOCKS`          | `DKGManager` constructor (immutable) | `25` (~5 min @ 12 s)           | Absolute window for committee `submitContribution` calls             |
| `FINALIZE_GAP_BLOCKS`          | `DKGManager` constructor (immutable) | `5`  (~1 min @ 12 s)           | Cooldown before `finalizeEpoch` may run                              |
| `MAX_N`                        | `solidity/src/libraries/Sizes.sol`   | `32`                           | Compile-time committee cap; mirrors `circuits/common.MaxN`           |
| `INACTIVITY_WINDOW`            | `DKGRegistry` constructor            | `50_400` blocks (~7 d @ 12 s)  | Heartbeat window before `reap` is permitted                          |
| `SEED_DELAY_BLOCKS`            | `Sizes.sol`                          | `1`                            | Lottery seed = `blockhash(startBlock + this)`                        |

---

## Integrating

### Run a node

Run a node and you become eligible to be selected on every epoch.

**Quickstart (Docker Compose):**

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg
cp .env.example .env && $EDITOR .env       # set DAVINCI_DKG_WEB3_RPC + PRIVKEY + NETWORK
docker compose --profile node up -d
docker compose --profile node logs -f node
```

For Sepolia you only need `DAVINCI_DKG_NETWORK=sepolia`; contract addresses are baked in. For any
other network, supply `DAVINCI_DKG_MANAGER=0x...`. The node binary itself does not serve any HTTP
— pair it with the standalone `ghcr.io/vocdoni/davinci-dkg-ui` image if you want an explorer.

What happens on first start:

1. Derive the BabyJubJub key from your operator EVM key, register or update it in `DKGRegistry`.
2. Print a startup banner with chain head, registry stats, and your `self:` row.
3. Poll `DKGManager`, react to every phase you are eligible for. Heartbeat / reactivate
   automatically before `INACTIVITY_WINDOW` expires.

You only pay gas on the phases you actually participate in. Per-call cost breakdown is in
[`BENCHMARKS.md`](BENCHMARKS.md).

Release binaries and source builds work the same way; see `davinci-dkg-node --help` for flags
(every flag has a `DAVINCI_DKG_…` env-var equivalent).

### Application CLI

`cmd/dkgapp` is the organizer-side companion of the node: register an application (mode 0 or
mode 1), encrypt and submit a ciphertext with its proof of knowledge, post the organizer share of a
mode-1 application, and read the combined plaintext.

```bash
export DAVINCI_DKG_WEB3_RPC=https://ethereum-sepolia-rpc.publicnode.com
export DAVINCI_DKG_NETWORK=sepolia DAVINCI_DKG_PRIVKEY=0x...
go run ./cmd/dkgapp epoch                                   # newest epoch and PK_ep
go run ./cmd/dkgapp register -aid 0x0a…                     # mode 0 (public derivation)
go run ./cmd/dkgapp register -aid 0x0b… -codec              # mode 1; prints the organizer secret
go run ./cmd/dkgapp encrypt  -aid 0x0a… -m 42               # submits; prints the assigned index
go run ./cmd/dkgapp encrypt  -aid 0x0b… -m 7 -org-secret …  # also posts Δ_org with its DLEQ proof
go run ./cmd/dkgapp plaintext -aid 0x0a… -index 1 -wait 5m
```

Application ids must be non-zero and below the BN254 scalar field (clear the top three bits of a
random or hashed id); `-epoch` defaults to the newest epoch.

### TypeScript SDK

```bash
pnpm add @vocdoni/davinci-dkg-sdk
```

```ts
import { DKGClient, DKGWriter, buildElGamal, buildEpochId } from '@vocdoni/davinci-dkg-sdk';

const client = new DKGClient({ publicClient, managerAddress });
const epoch  = await client.getEpoch(epochId);
const pk     = await client.getCollectivePublicKey(epochId);

// ElGamal encrypt + on-chain submit
const eg     = await buildElGamal();
const ct     = eg.encrypt(42n, [pk.x, pk.y]);
const writer = new DKGWriter({ publicClient, walletClient, managerAddress });
const { hash, ciphertextIndex } = await writer.submitCiphertext(epochId, '0x00…', ct); // index is assigned on chain

// Wait for the committee to finish, then read the plaintext
await writer.waitForCombinedDecryption(epochId, '0x00…', ciphertextIndex);
const m = await client.getPlaintext(epochId, '0x00…', 0);
```

Full reference: `sdk/README.md` and the typed entry points under `sdk/src/`.

### Encrypting and decrypting

The protocol stays threshold-secure as long as fewer than `t` committee operators collude. The
honest path:

1. Read `PK_ep` (or `PK_aid` for an application) from chain.
2. ElGamal-encrypt your plaintext scalar `m` (must fit under the BSGS cap — 2⁵⁰ on the committee,
   2³² in the SDK).
3. `DKGManager.submitCiphertext(epochId, aid, c1, c2, pok)` — the contract assigns the next
   `ctIdx` for the application and emits `CiphertextSubmitted`. `pok` is a Schnorr proof of
   knowledge of the encryption randomness `r` (`C₁ = r·G`), bound to `(epochId, aid, C₁, C₂)`.
   Committee nodes verify it before releasing a partial decryption, which is what stops anyone
   from re-submitting another application's `C₁` (or `C₁ + x·G`) under their own application
   as a decryption oracle. The SDK and `dkgapp` produce it for you.
4. Every committee node watches `CiphertextSubmitted` events (for any `aid`), submits its
   partial, and the first node to reach `t` partials calls `combineDecryption`. The recovered
   plaintext is readable on `getPlaintext`. A restarted node re-scans the last
   `--decrypt-lookback-blocks` (default ~7 days) for ciphertexts still awaiting decryption.

For mode-1 applications, the organizer must additionally call `submitOrganizerShare` for each
ciphertext before the combine can land.

---

## Deployments

| Network | DKGManager                                 | Notes |
|---------|--------------------------------------------|-------|
| Sepolia | `0x613589ea717311d5ddf71f1326c42871576c63a5` | Built into the node + SDK; just pass `--network sepolia`. Registry `0x2638cd0c2432f0af1602b3ecf2dd13a11076981f`, app manager `0xd7115537af9e021c055b0d19d11c29df4c79f25e`, deployed at block 11,616,299 with floors `MIN_THRESHOLD=2`, `MIN_COMMITTEE_SIZE=3`, `MAX_LOTTERY_ALPHA_BPS=20000` and windows 100/25/25/5 blocks. |

`DKGRegistry` and `DKGAppManager` are auto-resolved from `DKGManager` on-chain — only the manager
address needs to be configured.

---

## Build from source

Requires **Go 1.25+**, **Foundry**, and (for the UI) **pnpm**.

```bash
make build                  # davinci-dkg-node binary
cd solidity && forge build  # contracts
forge test                  # contract tests
go test ./...               # Go unit + circuit tests

# Integration tests (Anvil + Docker)
RUN_INTEGRATION_TESTS=true go test ./tests/... -timeout=15m
```

Circuit artifacts are cached under `~/.davinci/artifacts`. Recompile + regenerate the Solidity
verifier wrappers + Go bindings with:

```bash
make circuits
```

Switching `MaxN` is a two-line edit (`circuits/common/sizes.go` and `solidity/src/libraries/Sizes.sol`)
followed by `make circuits`. See [`BENCHMARKS.md`](BENCHMARKS.md) for the per-`MaxN` cost
breakdown.

A self-contained multi-node testnet (Anvil + N nodes + UI) lives in `testnet/`:

```bash
make testnet-up                                  # 3 nodes, defaults
make testnet-up DKG_NODE_COUNT=8 DKG_THRESHOLD=5 # custom sizing
```

The web explorer is then available at `http://localhost:8081/`.

---

## References

- NI-DKG paper: https://eprint.iacr.org/2026/552
- DAVINCI voting protocol: https://davinci.vote
- Vocdoni: https://vocdoni.io
