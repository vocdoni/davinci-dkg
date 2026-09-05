# DAVINCI DKG

**Non-Interactive Distributed Key Generation on EVM chains.**

Reference Go implementation of the protocol described in
[*NI-DKG: Non-Interactive Distributed Key Generation using Blockchain and ZK Proofs*](https://eprint.iacr.org/2026/552).
Built as the threshold-key layer for the [DAVINCI](https://davinci.vote) voting system, but the
protocol is generic — any application that needs a `t`-of-`n` collective public key on an EVM chain
can use it.

The protocol replaces interactive complaint rounds with Groth16 ZK proofs: every contribution,
batched finalization, partial decryption and combine is verified at transaction time (the v4
finalization carries a proof and stores the whole pool in one call — there is no separate
activation step any more). There is no dispute phase.

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
  - [Application CLI](#application-cli)
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
| Circuits          | `circuits/`                   | Groth16 / BN254 — Contribution, Finalize, PartialDecrypt, DecryptCombine  |
| TypeScript SDK    | `sdk/`                        | `@vocdoni/davinci-dkg-sdk` — read client, writer, encryption     |
| Web explorer / UI | `ui/`                         | React SPA + interactive playground                               |

Crypto primitives: BabyJubJub on the BN254 scalar field; Poseidon1 for in-circuit hashing;
keccak256 for on-chain Fiat–Shamir challenges. ElGamal for share and ciphertext encryption.

---

## Protocol model

### Epoch lifecycle

An **epoch** is one DKG run. Its `n` committee members jointly deal `MAX_K` (16) independent **pool
keys** `P_0 … P_15` in one shot; any `t` of the members can help decrypt under any one of them.
There is no single epoch key any more — every application claims one pool key for itself, so the
committee's partials are scoped to that application (see "Per-application keys"). Epochs are
scheduled at a fixed cadence — every `EPOCH_DURATION_BLOCKS` blocks (set per-deploy as a
`DKGManager` immutable).

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
   │  claimSlot         │submitContrib│ ... │ registerApplication /           │
   │  (lottery)         │  (Groth16,  │     │ submitCiphertext /              │
   │                    │  deals 16   │     │ submitPartialDecryption /       │
   │                    │  pool keys) │     │ revealOrganizerSecret /         │
   │                    │             │     │ combineDecryption               │
   └────────────────────┴─────────────┴─────┴─────────────────────────────────┘
                                       ▲
                    finalizeEpoch (proof-carrying — stores all 16 keys and share roots, sets Live)
                    KeyAssembly → Live

       ◄──────────────── EPOCH_DURATION_BLOCKS ─────────────────────────►
```

- **Preparation** — committee is assembled and its `MAX_K` pool keys are dealt. Three contiguous
  block windows:
  - `CommitteeSelection`: lottery via `claimSlot` picks `n` operators.
  - `KeyAssembly`: each committee member submits one Feldman VSS contribution — dealing all
    `MAX_K` pool keys at once (compact `K·(2t+n) + 5n`-word transcript) — with a single Groth16
    proof.
  - Finalize gap: short window before `finalizeEpoch` may run.
- **Service** — pool keys are claimed for the rest of the epoch.
  - `finalizeEpoch` is **proof-carrying and batched**: one Groth16 proof over the whole pool — it
    reproduces every key's aggregate `P_j` and the Merkle root of the whole committee's share
    commitments for it, stores all `MAX_K` keys and roots on chain, and sets the epoch `Live`. Its
    Fiat–Shamir challenge is anchored on the proof's own Poseidon transcript digest as well as the
    calldata, like a contribution's.
  - Apps claim the next unclaimed key via `registerApplication`, either organizer-locked
    (`PK_aid = P_j + PK_org`, the organizer keeps `sk_org`) or automatic (`PK_aid = P_j`, no
    organizer key at all).
  - A submitter the application's policy admits calls `submitCiphertext`; the committee posts
    partials scoped to `P_j`; an organizer-locked application is opened once its organizer calls
    `revealOrganizerSecret` (never, if it should stay closed); once `t` partials are on chain, the
    decryption window is open, and — for a locked app — the secret is revealed, any caller
    `combineDecryption`s to land the recovered plaintext on chain.

Each Preparation window is an **absolute** block count, not a fraction of the epoch — the lottery
is one keccak per claimer and the contribution proof is one tx per committee member, so a fixed
budget is the right shape. The four block constants are deploy-time immutables (defaults in
`solidity/src/libraries/Sizes.sol`, overridable via `EPOCH_DURATION_BLOCKS`,
`COMMITTEE_SELECTION_BLOCKS`, `KEY_ASSEMBLY_BLOCKS`, `FINALIZE_GAP_BLOCKS` env vars at deploy
time). Long epochs (multi-day) keep the same short Preparation; the extra time falls into Service.

The epoch stays `Live` for the entire Service window — its pool keys remain claimable and usable
while the next epoch bootstraps.

`createEpoch` is **permissionless** but cadence-gated: it reverts unless
`block.number >= nextEpochStartBlock()`, except that it is also allowed early when the newest
epoch is `Live` with `poolNext >= 15` (at most one unclaimed key left), or `Aborted` — so a busy
deployment never runs dry and a dead one does not have to wait out its cadence. In production,
every node races to fire it once the window opens (random jitter, env-toggleable). Only the first
call lands; the others revert cheaply. `finalizeEpoch` stores every key of the pool at once, so a
`Live` epoch has nothing left to activate (the former `--activate-ahead` flag and
`DAVINCI_DKG_ACTIVATE_AHEAD` are gone); nodes create the next epoch early as the pool runs low.

Two limits follow from the fixed pool. **Pool exhaustion:** an epoch serves at most `MAX_K` (16)
applications; once its keys are claimed, `registerApplication` reverts `PoolExhausted` until
the next epoch is `Live` — which takes a full preparation window (committee selection, key assembly,
finalize gap) plus one finalization proof. **Registration-driven epoch amplification:** registration
is permissionless, so anyone registering fifteen automatic applications forces the next epoch to open
early and every committee member to contribute again. The attacker pays fifteen registrations' gas and
the committee pays one extra epoch, so the cost is bounded, but it is an amplification; a
registration fee or an allow-list on `registerApplication` is future work.

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

Once an application has claimed a pool key `P_j` (see "Per-application keys"), decrypting an
ElGamal ciphertext `(C₁, C₂)` published under it via `submitCiphertext` goes:

1. Each committee member `i` publishes its partial `δ_i = e_{j,i} · C₁`, `e_{j,i}` being its share
   of `P_j`, plus a Groth16 proof of the Chaum–Pedersen relation (`D_i = d_i·G`, `δ_i = d_i·C₁`,
   `A_i = w·G`, `B_i = w·C₁`). `submitPartialDecryption` also carries a 5-word Merkle path proving
   `D_i` against the key's share-commitment root that `finalizeEpoch` stored. That tree covers
   the **whole committee**: a member that claimed a slot but never contributed still received a
   share from every accepted dealer and may post partials, so decryption survives `n − t` absent
   members, not `m − t` (`m` = accepted contributions).
2. Decryption must be **open**: `decryptNotBefore ≤ now ≤ decryptNotAfter` (both unix seconds,
   `0` = unbounded), checked on both the partial and the combine (`DecryptionNotOpen()` /
   `DecryptionClosed()`). Submission is gated separately, by the block window (`notBeforeBlock` /
   `notAfterBlock`), the submitter policy and `decryptNotAfter` only — a ciphertext may be
   submitted before decryption opens.
3. An organizer-locked application additionally needs its organizer to call
   `revealOrganizerSecret` — **once, for the whole application**, whenever they choose (or never).
   The contract checks `sk_org · G == PK_org` before storing it and accepts the call only once.
   Until then `submitPartialDecryption` and `combineDecryption` revert
   `OrganizerSecretNotRevealed()`, so **no partial and no combine of a locked application exists
   before the reveal**: the organizer learns every result together with everyone else, and it
   decides *when* the application opens, never *which* ciphertexts — enforced by the contract,
   not by node policy. From the reveal on, every ciphertext of the application, past or future,
   can be combined by the committee alone; there is no per-ciphertext organizer step. An
   automatic application has no organizer key at all, so this step never applies to it.
4. Once `t` partials are on chain, the window is open, and — for a locked application — `sk_org`
   has been revealed, anyone calls `combineDecryption`. A Groth16 proof attests that
   `Σ λ_k · δ_k` Lagrange-interpolates correctly, that the caller knows the secret matching the
   application's registered `PK_org` (the zero scalar, for automatic), and that
   `m · G + Σ λ_k · δ_k + sk_org · C₁ = C₂`.
5. The recovered scalar `m` is stored on-chain and readable via `getPlaintext`.

What the window guarantees, honestly: it bounds what the contract accepts and what honest nodes
post. It does not bind `t` colluding committee members — they hold shares and can compute
partials off chain whenever they like; for an automatic application that is the whole
confidentiality assumption, for a locked one they still lack `sk_org` until the reveal. Partials
are gated on chain as well as the combine, not only because `t` partials alone let anyone finish
the combine off chain, but so that there are no on-chain partials from before the window or the
reveal to collect later. Nodes park a slot whose window has not opened or whose organizer has
not revealed, and drop one whose window has closed.

The combine proof discovers `m` by baby-step giant-step (BSGS) discrete-log inversion. The
committee node caps at 2⁵⁰ and builds a 256 MB table once per process. The SDK caps at 2³², so
its table stays around 16 MB and runs in a browser. Submitting a plaintext above the relevant
cap is unrecoverable.

### Per-application keys

A `Live` epoch deals `MAX_K` (16) independent pool keys `P_0 … P_15` and hosts many independent
encryption contexts — one per **application**, keyed by a 32-byte `aid` chosen by whoever
registers it. `aid` is bound into every decryption proof as a BN254 scalar-field public input, so
it must be non-zero and below the field modulus (clear the top three bits of a random or hashed
id); the contract rejects other values.

Every application registers through `registerApplication`, which claims the next **unclaimed**
pool key on its behalf (reverting `PoolExhausted` when all 16 keys are taken; there is no activation
state — `finalizeEpoch` proved and stored the whole pool atomically, so every unclaimed key of a
`Live` epoch is usable) and fixes one of two **modes** (`policy.mode`) for the life of the
application:

- **Organizer-locked** (the default). The registration publishes `PK_org = sk_org · G` together
  with a Schnorr proof of possession of `sk_org` (domain `davinci-dkg:organizer-register:v1`,
  verified on chain), and `sk_org` stays with the organizer. The application key is
  `PK_aid = P_j + PK_org`, so **decryption needs both the committee and the organizer**: the
  committee alone only ever recovers shares of `P_j`'s secret, and the organizer calls
  `revealOrganizerSecret` **once, for the whole application** — or never.
- **Automatic**. There is no organizer key at all: `PK_aid = P_j` directly, `organizerPK` is
  stored as the identity point and `organizerSecret = 0`. **Decryption needs nobody but the
  committee** and happens as soon as `t` partials land and the decryption window is open. Use it
  for ciphertexts that are meant to be opened — a tally — not for anything that must stay private.

Because every application gets its own pool key, **the cross-application decryption oracle that
used to exist is closed**: a ciphertext `(C₁, C₂)` copied out of one application and re-submitted
under another decrypts under that other application's `P_j` — an unrelated secret — so the result
is garbage, not the original plaintext. Previously every application shared one epoch key
`PK_ep`, so copying a ciphertext into a freshly registered automatic application let anyone learn
`sk_ep · C₁` for a `C₁` they never had a right to; that path no longer exists.

Submission is gated by the application's policy: `openSubmission` lets anyone call
`submitCiphertext`; otherwise `submitters` is an exclusive allow-list of up to 32 addresses (the
registrant is not implicitly on it); when both are empty only the registrant may submit (the
default). Contradictory policies — open submission with a non-empty list, a zero address on the
list, a decryption window that isn't in the future — revert with `InvalidPolicy()`. Every
ciphertext belongs to a registered application, and the committee only ever answers ciphertexts
actually submitted under it by an authorised submitter: an unsubmitted ballot stays private in
both modes, automatic included.

Consequences worth internalising, for an organizer-locked application:

- **Losing `sk_org` makes the application permanently undecryptable.** It is not derivable from
  anything on chain. Back it up at registration time.
- **The organizer's silence is what keeps the application closed** — even `t` colluding
  committee members cannot decrypt before `revealOrganizerSecret` has been called, and the
  contract refuses every partial and combine until then, so the organizer sees results no
  earlier than anyone else. Reveal it only once every ciphertext of the application is meant to
  become openable.
- **The reveal is a one-time, whole-application act, not a per-ciphertext release.** Once
  `sk_org` is public, every past and future ciphertext of that application decrypts as soon as
  `t` partials and the window are there — the organizer stops being a per-ciphertext gate the
  moment it reveals.
- **Never reuse an organizer secret across two locked applications.** `sk_org` is user-chosen at
  registration; revealing it for one application exposes every ciphertext of any other
  application registered with the same secret. Draw a fresh one per application (`dkgapp
  register` does).

And for an automatic application:

- **Nothing is withheld and nobody is accountable — by design, for that application only.** Other
  applications and every other epoch secret are untouched: pool keys are independent per
  application and per epoch, and partials already reveal each member's share of `P_j` in both
  modes. Integrity is unchanged — the combine SNARK still proves the interpolation.

Cross-application replay used to be stopped by the organizer key; now the pool key itself does
that job, so `submitCiphertext` still needs no proof of knowledge of the encryption randomness —
which is what makes homomorphic aggregation possible (the submitter of an aggregated tally cannot
know its randomness). The organizer key, where present, adds a second factor on top: even within
the right application, decryption also needs the organizer's reveal.

---

## On-chain surface

Three contracts. Deploy order: `DKGRegistry → DKGManager → DKGAppManager`, then wire with
`DKGRegistry.setManager(...)` and `DKGManager.setAppManager(...)`. The split exists only to keep
each contract under EIP-170; logically `DKGManager` and `DKGAppManager` share one storage.

| Contract        | Owns                                                                                    |
|-----------------|-----------------------------------------------------------------------------------------|
| `DKGRegistry`   | Operator identities (BabyJubJub pub keys), liveness (`heartbeat`, `reactivate`, `reap`) |
| `DKGManager`    | Epoch lifecycle: `createEpoch`, `claimSlot`, `submitContribution`, `finalizeEpoch` (proof-carrying, batched — stores the whole pool), pool-key views (`getPoolKey`, `getPoolStatus`, `getPoolShareRoot`, `getAppPoolIndex`), ciphertexts, partial / combined decryption |
| `DKGAppManager` | Per-application registration: `registerApplication` (organizer-locked or automatic, submission policy, decryption window), `revealOrganizerSecret`, and the `requireCanSubmitCiphertext` / `requireDecryptionOpen` views the manager consults |

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
| `MAX_K`                        | `solidity/src/libraries/Sizes.sol`   | `16`                           | Pool keys dealt per epoch; mirrors `circuits/common.MaxK`            |
| `MERKLE_DEPTH`                 | `solidity/src/libraries/Sizes.sol`   | `5` (= log2 `MAX_N`)           | Depth of each pool key's share-commitment Merkle tree                |
| `INACTIVITY_WINDOW`            | `DKGRegistry` constructor            | `50_400` blocks (~7 d @ 12 s)  | Heartbeat window before `reap` is permitted                          |
| `SEED_DELAY_BLOCKS`            | `Sizes.sol`                          | `1`                            | Lottery seed = `blockhash(startBlock + this)`                        |
| `MAX_SUBMITTERS`               | `DKGAppManager`                      | `32`                           | Cap on an application's `policy.submitters` allow-list               |

---

## Integrating

### Run a node

Run a node and you become eligible to be drawn on every epoch created after you register. The
Sepolia deployment below is open, so anyone can join the committee.

You need an Ethereum key with a little Sepolia ETH (about 0.05 ETH covers weeks of
participation, and any Sepolia faucet works), Docker, and a machine sized for the proofs:
at least 4 cores and 16 GB of RAM minimum (more is safer). The v0.5.0 node keeps all four
circuits and their proving keys preloaded: about 9 GB at rest (8.5–9.6 GB across the seed
fleet, against 3.1 GB for v0.4) and up to about 10 GB while proving — the contribution proof
(5.9 M constraints at `MaxK = 16`) peaks at 9.3 GB, a running seed node observed ~9.8 GB, and
the finalization proof takes about 2.3 s and 5.5 GB. More cores shorten the proofs; less RAM is
not an option (see [`BENCHMARKS.md`](BENCHMARKS.md)).

The node's RPC list should hold at least two endpoints: the node classifies rate-limited or
unreachable endpoints and rotates off them, so a single-endpoint config has no fallback when a
provider rate-limits you.

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg
cp .env.example .env && $EDITOR .env
docker compose --profile node up -d
docker compose --profile node logs -f node
```

Three entries in `.env` are enough: `DAVINCI_DKG_NETWORK=sepolia`, your operator key in
`DAVINCI_DKG_PRIVKEY`, and at least two Sepolia endpoints in `DAVINCI_DKG_WEB3_RPC` (comma-separated;
the node rotates off rate-limited endpoints, so keep a fallback). For a named network
the contract addresses are built into the binary. On any other network, set
`DAVINCI_DKG_MANAGER=0x...` instead; the node resolves the registry and the app manager from
the manager on chain.

What happens on first start:

1. The node derives its BabyJubJub key from your operator EVM key and registers it in
   `DKGRegistry`. That is one transaction, skipped if you are already registered and active.
2. Before its first proof it downloads the pinned circuit artifacts from the release built into
   the binary — the [`circuits-v4`
   release](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v4), about 1.7 GB, of which
   the contribution proving key is 762 MB and the finalization proving key 425 MB — and checks every
   file against the hashes built into the binary.
3. It prints a startup banner with the chain head, registry statistics and its own `self:` row,
   then polls `DKGManager` and reacts to every phase it is eligible for.

Epochs on Sepolia last about 24 hours. Once per epoch the node claims a slot if the lottery
admits it and submits its contribution during key assembly, which is one Groth16 proof and a
few seconds of CPU. When the epoch qualifies it takes its turn in the seed-derived finalize
stagger — one node reconstructs the accepted contributions, proves the batched finalization and
submits `finalizeEpoch`, which stores all 16 pool keys and share roots at once; the rest answer
decryption requests for the epochs they belong to.

Committee size follows the registry: three quarters of the active operators, capped at 32, with
a majority threshold. Joining therefore takes effect on the next epoch, and nobody has to change
a setting for it.

The node keeps itself active in the registry. Leaving it off for more than seven days
(50,400 blocks) lets anyone mark it inactive; starting it again reactivates it.

You pay gas only for the phases you take part in. The per-call breakdown is in
[`BENCHMARKS.md`](BENCHMARKS.md).

The node binary serves no HTTP. Pair it with the standalone `ghcr.io/vocdoni/davinci-dkg-ui`
image to host an explorer of your own. The public explorer at
[dkg.davinci.vote](https://dkg.davinci.vote) shows every epoch, committee, operator and
decryption, and its playground lets you register an application and decrypt a value against the
live key from the browser.

Release binaries and source builds are configured the same way. Run `davinci-dkg-node --help`
for the full flag list; every flag has a `DAVINCI_DKG_…` environment equivalent.

### Application CLI

`cmd/dkgapp` is the organizer-side companion of the node: register an application, encrypt and
submit a ciphertext, reveal the organizer secret that opens decryption, and read the combined
plaintext.

```bash
export DAVINCI_DKG_WEB3_RPC=https://ethereum-sepolia-rpc.publicnode.com
export DAVINCI_DKG_NETWORK=sepolia DAVINCI_DKG_PRIVKEY=0x...
go run ./cmd/dkgapp epoch                                    # newest epoch and its pool status
go run ./cmd/dkgapp register  -aid 0x0a…                     # organizer-locked; generates + prints the organizer secret
go run ./cmd/dkgapp register  -aid 0x0b… -org-secret …       # or bring your own
go run ./cmd/dkgapp register  -aid 0x0c… -mode automatic     # no organizer key at all; committee-only decryption
go run ./cmd/dkgapp register  -aid 0x0d… -submitters 0xA…,0xB… -max 10 -decrypt-from 24h -decrypt-until 48h
go run ./cmd/dkgapp encrypt   -aid 0x0a… -m 42               # submits; prints the assigned index
go run ./cmd/dkgapp reveal    -aid 0x0a… -org-secret …       # opens the whole application, once, for good
go run ./cmd/dkgapp plaintext -aid 0x0c… -index 1 -wait 5m    # automatic: no reveal step needed
```

`register` takes `-mode locked|automatic` (default `locked`) and the submission policy:
`-submitters 0xA,0xB` (exclusive allow-list, up to 32 — add yourself if you mean to submit) or
`-open` (anyone may submit); with neither, only the registrant may. `-max N` caps the ciphertext
count, and `-decrypt-from` / `-decrypt-until` set the decryption window as RFC 3339 timestamps or
Go durations such as `48h`, relative to now — before the window opens nobody, organizer included,
can decrypt. `encrypt` and `plaintext` work the same in both modes; automatic mode takes no
organizer flags at all, since there is no organizer key to begin with. `epoch` shows the newest
epoch and its pool status: which of the 16 keys are claimed and which are free — there is no
activation bitmap any more (`getPoolStatus` returns the `poolNext` cursor only).

`reveal` publishes `sk_org` of an organizer-locked application **once, for the whole
application** — not per ciphertext, and not reversible. From then on every ciphertext of that
application, past or future, is combinable by the committee alone.

**Store the organizer secret of an organizer-locked application.** `register` prints it once when
it generates one; without it every ciphertext of that application is permanently undecryptable.
Never reuse it for another locked application: revealing it for one exposes the other.

Application ids must be non-zero and below the BN254 scalar field (clear the top three bits of a
random or hashed id); `-epoch` defaults to the newest epoch.

### TypeScript SDK

```bash
pnpm add @vocdoni/davinci-dkg-sdk
```

```ts
import { DKGClient, DKGWriter, buildElGamal, randomAid, randomOrganizerSecret } from '@vocdoni/davinci-dkg-sdk';

const client = new DKGClient({ publicClient, managerAddress });
const epoch  = await client.getEpoch(epochId);
const pool   = await client.getPoolStatus(epochId);       // next unclaimed key index (no activation bitmap)

// Register an organizer-locked application; keep skOrg — it is the other half of the key.
// aid must be non-zero and below the BabyJubJub scalar field: randomAid() does that.
const aid    = randomAid();
const skOrg  = randomOrganizerSecret();
const writer = new DKGWriter({ publicClient, walletClient, managerAddress });
await writer.registerApplication(epochId, aid, policy, skOrg);

// PK_aid = P_j (+ PK_org for a locked application); j is the pool key claimed at registration
const pkAid  = await client.getApplicationKey(epochId, aid);

// ElGamal encrypt under PK_aid, then submit
const eg     = await buildElGamal();
const ct     = eg.encrypt(42n, pkAid);
const { hash, ciphertextIndex } = await writer.submitCiphertext(epochId, aid, ct); // index is assigned on chain

// Open the whole application once, then wait for the committee and read the plaintext
await writer.revealOrganizerSecret(epochId, aid, skOrg);
await writer.waitForCombinedDecryption(epochId, aid, ciphertextIndex);
const m = await client.getPlaintext(epochId, aid, ciphertextIndex);
```

Full reference: `sdk/README.md` and the typed entry points under `sdk/src/`.

### Encrypting and decrypting

The protocol stays threshold-secure as long as fewer than `t` committee operators collude. The
honest path:

1. Register the application (or use one already registered) and read its key: `PK_aid = P_j`
   (automatic) or `PK_aid = P_j + PK_org` (organizer-locked) — see "Per-application keys".
2. ElGamal-encrypt your plaintext scalar `m` under `PK_aid` (must fit under the BSGS cap — 2⁵⁰ on
   the committee, 2³² in the SDK).
3. `DKGManager.submitCiphertext(epochId, aid, c1x, c1y, c2x, c2y)` — plain calldata, no proof.
   The contract assigns the next `ctIdx` for the application and emits `CiphertextSubmitted`. It
   checks the points are canonical, on-curve and non-identity, but deliberately **skips the
   prime-subgroup check** (about 0.17 M gas for `C₁` — skipped to keep submission cheap, not
   because it is prohibitive). Committee nodes perform it off chain before computing any
   partial, and that off-chain check is the load-bearing defence: a cofactor `C₁` would leak
   a member's share mod `h` from any node that skipped it. Cross-application replay is stopped by
   the per-application pool key, not by a proof of knowledge — see "Per-application keys".
4. Every committee node watches `CiphertextSubmitted` events and submits its partial — with a
   Merkle path against the key's share-commitment root — once the decryption window is open
   and, for a locked application, the organizer has revealed; until then the slot is parked and
   nothing is posted.
5. For an organizer-locked application, the organizer calls `revealOrganizerSecret` **once**, not
   per ciphertext, whenever the application as a whole should become decryptable. An automatic
   application has no such step: there is no organizer key to reveal.
6. Once `t` partials are on chain, the window is open, and (for a locked application) `sk_org`
   has been revealed, any caller — typically the committee node whose turn comes first in the
   seed-derived rotation — calls `combineDecryption`. The recovered plaintext is readable on
   `getPlaintext`. A restarted node re-scans the last `--decrypt-lookback-blocks` (default ~7
   days) for ciphertexts still awaiting decryption; slots waiting on the decryption window to
   open or, for a locked application, on the reveal, are parked at no cost until the relevant
   block or event arrives (a reveal rescans the application's ciphertexts from its registration
   block); slots past their application's decryption window are dropped; and an undecryptable
   ciphertext taints its source for the epoch (`<datadir>/tainted-apps.json`): always the
   offending (application, submitter) pair, so one bad submitter cannot silence an application
   for its honest submitters.

---

## Deployments

| Network | DKGManager                                 | Notes |
|---------|--------------------------------------------|-------|
| Sepolia | `0xf4fc804388211949b56b166281b2b86879b6278e` | Public v4 testnet (batched finalization, sixteen keys per epoch; contracts and [`circuits-v4`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v4) artifacts of this release), built into the node and SDK: pass `--network sepolia`. Registry `0x7f35800c5f81fd55799e5b6ebb78c0fc8f86bb5c`, app manager `0x735496bb75a4ec91faf9a1cf61a2c2325168e8cd`, deployed at block 11,642,464 (2026-09-05). Epochs last 7,200 blocks (about 24 h); committee selection 100 blocks, key assembly 150 blocks, finalize gap 10 blocks; floors `MIN_THRESHOLD=2`, `MIN_COMMITTEE_SIZE=3`, `MAX_LOTTERY_ALPHA_BPS=20000`; inactivity window 50,400 blocks. Earlier deployments (`0x6dd442e9…` v3.1 with `circuits-v3`, `0xd38af14c…` single-key with `circuits-v2`) are retired. |

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
followed by `make circuits`; `MaxN` must be a power of two (16, 32 or 64), since the share-commitment
Merkle tree has `MaxN` leaves. The circuits build against gnark v0.16.3 — never an older gnark: the
snapshot pinned before had an unsound twisted-Edwards scalar multiplication (see `CLAUDE.md`). See
[`BENCHMARKS.md`](BENCHMARKS.md) for the per-`MaxN` cost breakdown.

A self-contained multi-node testnet (Anvil + deployer + N nodes) lives in `testnet/`:

```bash
make testnet-up                                  # 3 nodes, defaults
make testnet-up DKG_NODE_COUNT=8 DKG_THRESHOLD=5 # custom sizing
# spread the fleet: 16 more nodes on another host, keys 17-32, against this host's Anvil
DKG_THRESHOLD=16 DKG_COMMITTEE_SIZE=24 DKG_MIN_VALID_CONTRIBUTIONS=20 \
  testnet/remote-nodes.sh up user@other-host 16 16
```

Phase windows (`EPOCH_DURATION_BLOCKS`, …) and the epoch policy the nodes propose
(`DKG_THRESHOLD`, `DKG_COMMITTEE_SIZE`, `DKG_MIN_VALID_CONTRIBUTIONS`, `DKG_ALPHA_BPS`) are
compose variables. Point the explorer at it with `make ui-dev RPC_URL=http://127.0.0.1:8545
MANAGER_ADDRESS=<from http://127.0.0.1:8888/addresses.env> CHAIN_ID=1337`.

`tests/battery/` drives a running fleet through load, concurrency and adversarial scenarios
(organizer swarm, tampered and replayed shares, poisoned ciphertexts, snapshot-rule and
duplicate-claim attacks, a lazy committee member) and writes a per-transaction report:

```bash
DAVINCI_DKG_BATTERY=1 DAVINCI_DKG_TEST_RPC_URL=http://127.0.0.1:8545 \
DAVINCI_DKG_TEST_ADDRESSES=/tmp/addresses.env DAVINCI_ARTIFACTS_DIR=~/.davinci/artifacts \
  go test ./tests/battery -run TestOrganizerSwarm -v -count=1 -timeout 40m
```

---

## References

- NI-DKG paper: https://eprint.iacr.org/2026/552
- DAVINCI voting protocol: https://davinci.vote
- Vocdoni: https://vocdoni.io
