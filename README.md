# DAVINCI DKG

**Non-interactive distributed key generation on EVM chains.**

Reference implementation of a threshold-key layer for smart contracts: a rotating committee of
operators jointly generates `t`-of-`n` public keys with one transaction each, applications encrypt
under those keys, and the committee decrypts on demand, every step verified on chain with a Groth16
proof. There are no interactive complaint rounds and no dispute phase: a contribution, a
finalization, a partial decryption and a combine are each one proof-carrying call, and a call that
verifies is final. Built as the key layer of the [DAVINCI](https://davinci.vote) voting system; the
protocol is generic and any application that needs a collective key on an EVM chain can use it.

The repository holds the Go node and circuits, the Solidity contracts, a TypeScript SDK and a web
explorer. A public testnet runs on Sepolia (see [Deployments](#deployments)).

---

## Contents

- [What you get](#what-you-get)
- [Protocol model](#protocol-model)
  - [Epoch lifecycle](#epoch-lifecycle)
  - [Committee selection](#committee-selection)
  - [Per-application keys](#per-application-keys)
  - [Threshold decryption](#threshold-decryption)
- [Proofs](#proofs)
- [On-chain surface](#on-chain-surface)
- [Running a node](#running-a-node)
- [Application CLI](#application-cli)
- [TypeScript SDK](#typescript-sdk)
- [Encrypting and decrypting](#encrypting-and-decrypting)
- [Deployments](#deployments)
- [Build from source](#build-from-source)
- [Repository layout](#repository-layout)
- [References](#references)

---

## What you get

| Component          | Path                     | Purpose                                                                 |
|--------------------|--------------------------|-------------------------------------------------------------------------|
| `davinci-dkg-node` | `cmd/davinci-dkg-node`   | Operator daemon: registers, claims committee slots, deals, finalizes, decrypts |
| `dkgapp`           | `cmd/dkgapp`             | Organizer CLI: register an application, encrypt, reveal, read plaintexts |
| Solidity contracts | `solidity/src`           | `DKGRegistry`, `DKGManager`, `DKGAppManager` and the four Groth16 verifiers |
| Circuits           | `circuits/`              | Groth16 over BN254: Contribution, Finalize, PartialDecrypt, DecryptCombine |
| TypeScript SDK     | `sdk/`                   | `@vocdoni/davinci-dkg-sdk`: read client, writer, ElGamal encryption      |
| Explorer           | `ui/`                    | React SPA with a playground; image `ghcr.io/vocdoni/davinci-dkg-ui`      |

Cryptography: BabyJubJub over the BN254 scalar field for keys, shares and ciphertexts; Poseidon for
in-circuit hashing; keccak256 for the on-chain Fiat–Shamir challenges; ElGamal for share and
ciphertext encryption; Groth16 for every proof.

---

## Protocol model

### Epoch lifecycle

An **epoch** is one DKG run. Its `n` committee members jointly deal `MAX_K` (16) independent
**pool keys** `P_0 … P_15`, and any `t` of the members can decrypt under any of them. Every
application later claims one pool key for itself, so a committee's partial decryptions are scoped to
one application. Epochs are created at a fixed cadence of `EPOCH_DURATION_BLOCKS` blocks, a
`DKGManager` immutable.

```
   startBlock                                                                 endBlock
   │                                                                          │
   │ ─── Preparation (small, fixed) ────► ◄────── Service (the rest) ────────│
   │                                                                          │
   │ CommitteeSelection │ KeyAssembly │ gap │            Live                 │
   ▼                                                                          ▼
   ├────────────────────┼─────────────┼─────┼─────────────────────────────────┤
   │  claimSlot         │submitContrib│ ... │ registerApplication /           │
   │  (lottery)         │  (Groth16,  │     │ submitCiphertext /              │
   │                    │  deals 16   │     │ submitPartialDecryption /       │
   │                    │  pool keys) │     │ revealOrganizerSecret /         │
   │                    │             │     │ combineDecryption               │
   └────────────────────┴─────────────┴─────┴─────────────────────────────────┘
                                       ▲
                    finalizeEpoch: one proof stores all 16 keys and share roots, sets Live

       ◄──────────────── EPOCH_DURATION_BLOCKS ─────────────────────────►
```

**Preparation** assembles the committee and deals the pool, in three contiguous block windows:

- `CommitteeSelection`: the lottery (`claimSlot`) picks `n` operators.
- `KeyAssembly`: each member submits one Feldman VSS contribution dealing all `MAX_K` polynomials
  at once, a compact `K·(2t+n) + 5n`-word transcript with a single Groth16 proof.
- A short finalize gap, after which `finalizeEpoch` may run.

**Service** is the rest of the epoch. `finalizeEpoch` is permissionless and proof-carrying: one
Groth16 proof over the whole pool reproduces every aggregate key `P_j` and the Merkle root of the
whole committee's share commitments for it, stores all `MAX_K` keys and roots atomically and sets
the epoch `Live`. From then on applications claim keys, submit ciphertexts and get them decrypted;
the epoch stays `Live` until the end of the cadence while the next epoch bootstraps.

The window lengths are absolute block counts, not fractions of the epoch: the lottery is one
keccak per claimer and the contribution is one transaction per member, so a fixed budget is the
right shape, and a multi-day epoch keeps the same short preparation. The four constants are
constructor immutables (`EPOCH_DURATION_BLOCKS`, `COMMITTEE_SELECTION_BLOCKS`,
`KEY_ASSEMBLY_BLOCKS`, `FINALIZE_GAP_BLOCKS`, settable at deploy time).

`createEpoch` is permissionless but cadence-gated: it reverts before `nextEpochStartBlock()`,
except when the newest epoch is `Live` with at most one unclaimed key (`poolNext >= MAX_K − 1`)
or `Aborted`, so a busy deployment never runs dry and a dead one need not wait out its cadence.
Every node races to fire it with a random jitter; the first call lands and the rest revert
cheaply. Whoever creates the epoch chooses `(t, n, minValidContributions, α)` within the
deployment's floors (`MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS`, and
`t ≤ MAX_T`); nodes derive them from the registry: three quarters of the active operators, capped
at 32, with a majority threshold. The contract requires `minValidContributions ≥ threshold`,
because the accepted dealers together know the key and there must be at least `t` of them.

Two limits follow from the fixed pool. An epoch serves at most `MAX_K` applications; once its keys
are claimed, `registerApplication` reverts `PoolExhausted` until the next epoch is `Live`, which
takes one preparation window plus one finalization. And since registration is permissionless,
anyone registering fifteen applications forces the next epoch to open early and every member to
contribute again; the cost is bounded (fifteen registrations against one extra epoch) but it is an
amplification, and a registration fee or allow-list is a deployment's own policy decision.

`EpochPhase` values: `None`, `CommitteeSelection`, `KeyAssembly`, `Live`, `Aborted` (`Completed`
is reserved).

### Committee selection

Each epoch draws a fresh committee from the registry by a trustless lottery:

- `n = committeeSize`, `α = lotteryAlphaBps / 10 000` (oversubscription), `R =
  registry.activeCount()` snapshotted at `createEpoch`;
- `seed = blockhash(startBlock + SEED_DELAY_BLOCKS)`, resolved by the first `claimSlot`;
- a node is eligible iff `keccak256(seed ‖ msg.sender) < α · n · 2²⁵⁶ / R`.

Eligible nodes race first-come-first-served until `n` slots are filled, which snapshots the
committee (positions `1..n`, each member's BabyJubJub key) and moves the epoch to `KeyAssembly`.
Anyone can replay the keccak; there is no coordinator. Only operators registered before
`createEpoch` may claim, so fresh identities cannot be ground against a revealed seed. A committee
that does not fill within the window makes the epoch dead: anyone may `abortEpoch` it once the
deadline has passed, and the nodes create the next epoch immediately. An epoch that can still be
finalized cannot be aborted.

### Per-application keys

A `Live` epoch hosts many independent encryption contexts, one per **application**, keyed by a
32-byte `aid` chosen by whoever registers it. `aid` enters every decryption proof as a BN254
scalar-field public input, so it must be non-zero and below the field modulus (clear the top three
bits of a random or hashed id).

`registerApplication` claims the next unclaimed pool key `P_j` for the application and fixes one of
two **modes** for its life:

- **Organizer-locked** (default). The registration publishes `PK_org = sk_org · G` with a Schnorr
  proof of possession (domain `davinci-dkg:organizer-register:v1`, verified on chain) and the
  organizer keeps `sk_org`. The application key is `PK_aid = P_j + PK_org`: the committee alone
  only ever recovers shares of `P_j`'s secret, and the organizer calls `revealOrganizerSecret`
  once, for the whole application, or never.
- **Automatic**. There is no organizer key: `PK_aid = P_j`, `organizerPK` is stored as the identity
  and `organizerSecret = 0`. Decryption needs nobody but the committee and happens as soon as `t`
  partials land inside the decryption window. Use it for ciphertexts that are meant to be opened,
  such as a tally, not for anything that must stay private.

Because every application has its own pool key, a ciphertext copied from one application into
another decrypts under an unrelated secret and yields garbage: there is no cross-application
decryption oracle, and `submitCiphertext` needs no proof of knowledge of the encryption
randomness, which is what keeps homomorphic aggregation possible (the submitter of an aggregated
tally cannot know its randomness).

Submission is gated by the application's policy: `openSubmission` lets anyone submit; otherwise
`submitters` is an exclusive allow-list of up to 32 addresses; with neither, only the registrant
may submit. `maxCiphertexts` caps the count and `decryptNotBefore` / `decryptNotAfter` (unix
seconds, `0` = unbounded) bound the decryption window. Contradictory policies revert
`InvalidPolicy()`. The committee only answers ciphertexts actually submitted under an application
by an authorised submitter, so an unsubmitted ballot stays private in both modes.

For an organizer-locked application:

- **Losing `sk_org` makes the application permanently undecryptable.** It is not derivable from
  anything on chain; back it up at registration.
- **The organizer's silence keeps the application closed.** Until `revealOrganizerSecret`, the
  contract refuses every partial and every combine, so even `t` colluding members cannot open a
  ciphertext on chain, and the organizer sees results no earlier than anyone else.
- **The reveal is a one-time, whole-application act.** Once `sk_org` is public, every past and
  future ciphertext of the application decrypts as soon as `t` partials and the window are there.
- **Never reuse an organizer secret across applications.** Revealing it for one exposes the other;
  `dkgapp register` draws a fresh one.

For an automatic application nothing is withheld and nobody is accountable, by design and for that
application only: pool keys are independent per application and per epoch, and the combine proof
still attests the interpolation.

### Threshold decryption

An ElGamal ciphertext `(C₁, C₂)` published under `PK_aid` through `submitCiphertext` is decrypted
as follows.

1. Each committee member `i` publishes `δ_i = e_{j,i} · C₁`, `e_{j,i}` being its share of `P_j`,
   with a Groth16 proof of the Chaum–Pedersen relation (`D_i = d_i·G`, `δ_i = d_i·C₁`) and a
   `MERKLE_DEPTH`-long Merkle path proving `D_i` against the share-commitment root `finalizeEpoch`
   stored for that key. The tree covers the whole committee: a member that claimed a slot but did
   not contribute still received a share from every accepted dealer and may post partials, so
   decryption survives `n − t` absent members.
2. Decryption must be **open**: `decryptNotBefore ≤ now ≤ decryptNotAfter`, checked on partials
   and combines (`DecryptionNotOpen()` / `DecryptionClosed()`). Submission is gated separately by
   the block window, the submitter policy and `decryptNotAfter`, so a ciphertext may be submitted
   before decryption opens.
3. An organizer-locked application additionally needs `revealOrganizerSecret`, once; the contract
   checks `sk_org · G == PK_org` and accepts the call a single time. Until then both partials and
   combines revert `OrganizerSecretNotRevealed()`.
4. Once `t` partials are on chain, anyone calls `combineDecryption` with a Groth16 proof that
   `Σ λ_k · δ_k` Lagrange-interpolates the qualifying set correctly, that the caller knows the
   secret behind the application's `PK_org` (zero for automatic), and that
   `m · G + Σ λ_k · δ_k + sk_org · C₁ = C₂`.
5. The plaintext `m` is stored on chain and read through `getPlaintext`.

What the window guarantees, stated honestly: it bounds what the contract accepts and what honest
nodes post. It does not bind `t` colluding members, who hold shares and can compute partials off
chain whenever they like; for an automatic application that is the whole confidentiality
assumption, for a locked one they still lack `sk_org`. Gating the partials as well as the combine
ensures there are no on-chain partials from before the window or the reveal to collect later.

The combine recovers `m` by baby-step giant-step: the node caps plaintexts at 2⁵⁰ with a 256 MB
table, the SDK at 2³² with about 16 MB so it runs in a browser. A plaintext above the cap is
unrecoverable.

---

## Proofs

Four Groth16 circuits over BN254 (gnark), each bound to its calldata by a Fiat–Shamir challenge the
contract derives from keccak256 over the calldata **and** the circuit's own Poseidon digests, and
checked inside the circuit through a random linear combination of every transcript word (BRLC).
No transcript word can differ between what the contract read and what the prover proved.

| Circuit | Proves | Constraints | Public inputs |
|---|---|---:|---:|
| Contribution | `MAX_K` polynomials of degree `< t`: the constant term's commitment `C_{j,0} = a·G`, every other commitment in the prime subgroup (a cofactor preimage `8·Q = C`), every recipient's share `s·G = Σ (i)^m C_{j,m}` with `s < r`, and the hashed-ElGamal encryption of each share under the recipient's key with one ECDH secret per recipient | 1,689,543 | 8 |
| Finalize | for up to 32 accepted dealers, that their commitments hash to the stored contribution digests, the aggregate keys `P_j = Σ C_{j,0}` and the share commitment `D_{j,i}` of every committee position for every key | 2,228,434 | 7 |
| PartialDecrypt | the Chaum–Pedersen relation of one partial against a committed share | 26,179 | 15 |
| DecryptCombine | the Lagrange interpolation over the qualifying set, knowledge of the organizer secret and the ElGamal decryption equation | 255,072 | 9 |

Shares are masked in the BN254 scalar field, `masked = s + H(seed_i, j) mod p` with
`seed_i = H(domain, eid, indexes, S_i)` over the ECDH secret `S_i`, a one-time pad the recipient
removes and range-checks. The compiled circuits, proving keys and verifying keys are pinned by
SHA-256 in `config/circuit_artifacts.go` and published as a GitHub release; a node downloads them
on first start and stream-verifies every file against those hashes, at startup and again whenever
a proving key is loaded for a proof. Nothing is compiled at runtime. The normative encodings live
in [`docs/pool-keys.md`](docs/pool-keys.md); measurements in [`BENCHMARKS.md`](BENCHMARKS.md).

---

## On-chain surface

Three contracts, deployed `DKGRegistry → DKGManager → DKGAppManager` and wired with
`DKGRegistry.setManager` and `DKGManager.setAppManager`. The split exists only for EIP-170;
`DKGManager` and `DKGAppManager` share one logical storage.

| Contract        | Owns |
|-----------------|------|
| `DKGRegistry`   | Operator identities (BabyJubJub public keys, checked on curve and in the prime subgroup), liveness (`heartbeat`, `reactivate`, `reap`) |
| `DKGManager`    | Epoch lifecycle (`createEpoch`, `claimSlot`, `submitContribution`, `finalizeEpoch`, `abortEpoch`), pool-key views (`getPoolKey`, `getPoolStatus`, `getPoolShareRoot`, `getAppPoolIndex`), ciphertexts, partial and combined decryption |
| `DKGAppManager` | `registerApplication` (mode, submission policy, decryption window), `revealOrganizerSecret`, and the `requireCanSubmitCiphertext` / `requireDecryptionOpen` views the manager consults |

`solidity/src/interfaces/*.sol` are the integration contract: full signatures and event schemas.
`submitContribution` and `finalizeEpoch` are direct-call gated (`msg.sender == tx.origin`), since
the node and the finalizer recover contributions from transaction calldata.

| Constant                     | Where                                | Default                 | Notes |
|------------------------------|--------------------------------------|-------------------------|-------|
| `EPOCH_DURATION_BLOCKS`      | `DKGManager` constructor (immutable) | `100`                   | Cadence anchor |
| `COMMITTEE_SELECTION_BLOCKS` | `DKGManager` constructor (immutable) | `25`                    | Lottery window |
| `KEY_ASSEMBLY_BLOCKS`        | `DKGManager` constructor (immutable) | `25`                    | Contribution window |
| `FINALIZE_GAP_BLOCKS`        | `DKGManager` constructor (immutable) | `5`                     | Cooldown before `finalizeEpoch` |
| `MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS` | `DKGManager` constructor | deploy-time | Policy floors for `createEpoch` |
| `MAX_N`                      | `solidity/src/libraries/Sizes.sol`   | `32`                    | Committee cap; mirrors `circuits/common.MaxN`, a power of two |
| `MAX_T`                      | `Sizes.sol`                          | `32`                    | Threshold cap; mirrors `circuits/common.MaxT` |
| `MAX_K`                      | `Sizes.sol`                          | `16`                    | Pool keys per epoch; mirrors `circuits/common.MaxK` |
| `MERKLE_DEPTH`               | `Sizes.sol`                          | `5` (= log2 `MAX_N`)    | Share-commitment tree depth |
| `SEED_DELAY_BLOCKS`          | `Sizes.sol`                          | `1`                     | Lottery seed block offset |
| `INACTIVITY_WINDOW`          | `DKGRegistry` constructor            | `50 400` blocks (~7 d)  | Heartbeat window before `reap` |
| `MAX_SUBMITTERS`             | `DKGAppManager`                      | `32`                    | Allow-list cap |

---

## Running a node

Run a node and you are eligible to be drawn into every epoch created after you register. The
Sepolia deployment is open.

You need an Ethereum key with a little Sepolia ETH (a node spends about 0.02 ETH a day under the
public testnet's load), Docker, and a machine with at least 2 cores and **4 GB of RAM, 8 GB to be
comfortable**. Every node deals shares, takes its turn at the finalization proof and decrypts.
Proving keys are loaded for a proof and released afterwards: a node sits at 0.2–0.7 GB at rest,
peaks at about 2.9 GB during its contribution proof, and starts in under 0.2 GB. `GOMEMLIMIT`
(a Go runtime setting, e.g. `GOMEMLIMIT=2500MiB`) trades some CPU for a tighter peak. More cores
shorten the proofs (0.9 s for a contribution on 32 threads).

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg
cp .env.example .env && $EDITOR .env
docker compose --profile node up -d
docker compose --profile node logs -f node
```

Three entries in `.env` are enough: `DAVINCI_DKG_NETWORK=sepolia`, your operator key in
`DAVINCI_DKG_PRIVKEY`, and at least two RPC endpoints in `DAVINCI_DKG_WEB3_RPC` (comma-separated;
the node rotates off rate-limited or unreachable endpoints, so a single endpoint has no fallback).
For a named network the contract addresses are built into the binary; on any other network set
`DAVINCI_DKG_MANAGER=0x…` and the node resolves the registry and the app manager from it.

On first start the node:

1. derives its BabyJubJub key from the operator key and registers it in `DKGRegistry` (one
   transaction, skipped if already registered and active);
2. downloads the pinned circuit artifacts from the
   [`circuits-v5`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v5) release
   (about 1.1 GB: the contribution proving key is 243 MB, the finalization proving key 436 MB) and
   stream-verifies every file against the hashes built into the binary;
3. prints a startup banner with the chain head, registry statistics and its own registry row,
   then polls `DKGManager` and reacts to every phase it is eligible for.

Once per epoch the node claims a slot if the lottery admits it and submits its contribution during
key assembly. When the epoch qualifies it takes its turn in a seed-derived stagger: one node
reconstructs the accepted contributions from calldata, proves the finalization and submits
`finalizeEpoch`; the others see the epoch go `Live` and stop. For every ciphertext of the epochs it
belongs to, it posts its partial in seed-derived waves of `t` members, so an honest ciphertext
costs `t` partials rather than `n`, and combines when its turn comes. Partials and combines are sent
without waiting for receipts, so one tick serves every pending ciphertext, and the node makes
about five RPC calls per tick.

The node keeps itself active in the registry. Left off for more than the inactivity window, it can
be marked inactive by anyone; starting it again reactivates it. It serves no HTTP; pair it with the
explorer image to host a UI of your own. Every flag has a `DAVINCI_DKG_…` environment equivalent
(`davinci-dkg-node --help`). To run nodes on Railway through its API see
[`railway-deploy.md`](railway-deploy.md).

---

## Application CLI

`cmd/dkgapp` is the organizer-side companion: register an application, encrypt and submit a
ciphertext, reveal the organizer secret, read the plaintext.

```bash
export DAVINCI_DKG_WEB3_RPC=https://ethereum-sepolia-rpc.publicnode.com
export DAVINCI_DKG_NETWORK=sepolia DAVINCI_DKG_PRIVKEY=0x...
go run ./cmd/dkgapp epoch                                    # newest epoch and its pool status
go run ./cmd/dkgapp register  -aid 0x0a…                     # organizer-locked; generates and prints the organizer secret
go run ./cmd/dkgapp register  -aid 0x0b… -org-secret …       # or bring your own
go run ./cmd/dkgapp register  -aid 0x0c… -mode automatic     # no organizer key; committee-only decryption
go run ./cmd/dkgapp register  -aid 0x0d… -submitters 0xA…,0xB… -max 10 -decrypt-from 24h -decrypt-until 48h
go run ./cmd/dkgapp encrypt   -aid 0x0a… -m 42               # submits; prints the assigned index
go run ./cmd/dkgapp reveal    -aid 0x0a… -org-secret …       # opens the whole application, once
go run ./cmd/dkgapp plaintext -aid 0x0c… -index 1 -wait 5m    # automatic: no reveal step
```

`register` takes `-mode locked|automatic` (default `locked`), the submission policy (`-submitters
0xA,0xB` as an exclusive allow-list of up to 32, or `-open`; with neither only the registrant may
submit), `-max N`, and `-decrypt-from` / `-decrypt-until` as RFC 3339 timestamps or Go durations
relative to now. `reveal` publishes `sk_org` once for the whole application and is not reversible.
Store the organizer secret of a locked application: without it every ciphertext of that
application is permanently undecryptable. Application ids must be non-zero and below the BN254
scalar field; `-epoch` defaults to the newest epoch.

---

## TypeScript SDK

```bash
pnpm add @vocdoni/davinci-dkg-sdk
```

```ts
import { DKGClient, DKGWriter, buildElGamal, randomAid, randomOrganizerSecret } from '@vocdoni/davinci-dkg-sdk';

const client = new DKGClient({ publicClient, managerAddress });
const epoch  = await client.getEpoch(epochId);
const pool   = await client.getPoolStatus(epochId);       // next unclaimed key index

// Register an organizer-locked application; keep skOrg, it is the other half of the key.
const aid    = randomAid();                                // non-zero, below the scalar field
const skOrg  = randomOrganizerSecret();
const writer = new DKGWriter({ publicClient, walletClient, managerAddress });
await writer.registerApplication(epochId, aid, policy, skOrg);

// PK_aid = P_j (+ PK_org for a locked application); j is the pool key claimed at registration
const pkAid  = await client.getApplicationKey(epochId, aid);

const eg     = await buildElGamal();
const ct     = eg.encrypt(42n, pkAid);
const { hash, ciphertextIndex } = await writer.submitCiphertext(epochId, aid, ct);

await writer.revealOrganizerSecret(epochId, aid, skOrg);   // once, for the whole application
await writer.waitForCombinedDecryption(epochId, aid, ciphertextIndex);
const m = await client.getPlaintext(epochId, aid, ciphertextIndex);
```

The SDK also decodes contribution and finalization transcripts, recomputes their Poseidon digests
and Merkle roots, and asserts the shared protocol constants against `tests/vectors/*.json`.
Full reference: `sdk/README.md`.

---

## Encrypting and decrypting

The protocol stays threshold-secure as long as fewer than `t` committee members collude. The honest
path:

1. Register the application, or use one already registered, and read its key `PK_aid`.
2. ElGamal-encrypt the plaintext scalar `m` under `PK_aid` (below the BSGS cap: 2⁵⁰ on the
   committee, 2³² in the SDK).
3. `DKGManager.submitCiphertext(epochId, aid, c1x, c1y, c2x, c2y)`: plain calldata, no proof. The
   contract assigns the next index and emits `CiphertextSubmitted`. It checks the points are
   canonical, on the curve and not the identity, and deliberately skips the prime-subgroup check
   (about 0.17 M gas for `C₁`): committee nodes perform it off chain before computing any
   partial, and that check is load-bearing, since a cofactor `C₁` would leak a member's share
   modulo the cofactor.
4. Every committee node watches `CiphertextSubmitted`, and once the window is open and, for a
   locked application, the organizer has revealed, posts its partial with the Merkle path against
   the key's share-commitment root. Until then the slot is parked at no cost.
5. For a locked application the organizer calls `revealOrganizerSecret` once, whenever the
   application as a whole should become decryptable.
6. Once `t` partials are on chain, a node whose turn comes in the seed-derived rotation calls
   `combineDecryption`; the plaintext is readable through `getPlaintext`. A restarted node re-scans
   the last `--decrypt-lookback-blocks` (default about seven days) for ciphertexts still awaiting
   decryption; a slot past its window is dropped; and a ciphertext whose plaintext is out of range
   taints its (application, submitter) pair for the epoch, so one bad submitter cannot silence an
   application for its honest submitters.

---

## Deployments

| Network | DKGManager | Details |
|---------|------------|---------|
| Sepolia | `0xf7826a1bc67438856183833b6fbd1c3a93803e9a` | Public testnet, built into the node and the SDK (`--network sepolia`). Registry `0x20ed76408981ae8bbf3a8658edd254f5ad69c3eb`, app manager `0x69e047134ed5080fe02e93db7d6548f3efe24655`; verifiers contribution `0x6d035f862d47e6019fb558f2236c5dfaa6c3525b`, finalize `0x4549ab46bc45c3806e30153c7034feac4ae3dab4`, partial `0x602bb41d21441044a54eb063eb63fdf4a278c70e`, combine `0xc7642f5fc6d531c8155261f08b33a5f38f711518`; deployed at block 11,663,483 with the [`circuits-v5`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v5) artifacts. Epochs last 7,200 blocks (about 24 h); committee selection 100 blocks, key assembly 150, finalize gap 10; floors `MIN_THRESHOLD=2`, `MIN_COMMITTEE_SIZE=3`, `MAX_LOTTERY_ALPHA_BPS=20000`, `MAX_T=32`; inactivity window 50,400 blocks. |

Only the manager address needs configuring; the registry and the app manager are resolved from it
on chain. The public explorer is at [dkg.davinci.vote](https://dkg.davinci.vote).

---

## Build from source

Requires Go 1.25+, Foundry, pnpm 10 and Docker (for the integration tests).

```bash
make build                                   # cmd/... binaries
make test                                    # Go unit tests (circuit tests cache artifacts under ~/.davinci/artifacts)
cd solidity && forge build && forge test     # contracts
cd sdk && pnpm install && pnpm build && pnpm test
RUN_INTEGRATION_TESTS=true go test ./tests/... -timeout 2h -failfast -count=1   # Anvil in Docker
```

`make circuits` recompiles the four circuits, runs the Groth16 setup, rewrites the Solidity
verifying keys, pins the artifact hashes and regenerates the Go bindings; `make vectors`
regenerates the cross-implementation fixtures. The Groth16 setup is randomized, so a fresh setup
never matches the pinned hashes or the committed verifiers until all three are regenerated together.
Changing `MaxN`, `MaxT` or `MaxK` is an edit to `circuits/common/sizes.go` and
`solidity/src/libraries/Sizes.sol` followed by `make circuits`; `MaxN` must be a power of two. The
circuits require gnark ≥ 0.16.2 (older releases have an unsound scalar-multiplication gadget).

A self-contained multi-node testnet (Anvil, deployer, N nodes) lives in `testnet/`:

```bash
make testnet-up                                  # 3 nodes
make testnet-up DKG_NODE_COUNT=8 DKG_THRESHOLD=5 # custom sizing
DKG_THRESHOLD=16 DKG_COMMITTEE_SIZE=24 DKG_MIN_VALID_CONTRIBUTIONS=20 \
  testnet/remote-nodes.sh up user@other-host 16 16    # 16 more nodes on another host
```

Phase windows and the epoch policy the nodes propose are compose variables; point the explorer at
it with `make ui-dev RPC_URL=http://127.0.0.1:8545 MANAGER_ADDRESS=<from
http://127.0.0.1:8888/addresses.env> CHAIN_ID=1337`. `tests/battery/` drives a running fleet
through load, concurrency and adversarial scenarios and writes a per-transaction report.

---

## Repository layout

| Path | Contents |
|---|---|
| `node/` | The daemon: epoch participation, finalization stagger, decryption scanner, RPC pool |
| `finalizer/` | Reconstructs accepted contributions from calldata and proves `finalizeEpoch` |
| `circuits/` | The four gnark circuits, shared gadgets (`common/`), pinned artifact loader |
| `crypto/` | Off-circuit primitives: Feldman, Shamir, Schnorr, ElGamal, share encryption, group |
| `web3/` | Typed wrappers over the generated bindings, RPC pool, transaction manager |
| `solidity/` | Contracts, interfaces, verifiers, Foundry tests, deploy script, Go bindings |
| `sdk/`, `ui/` | TypeScript SDK and explorer |
| `cmd/` | `davinci-dkg-node`, `dkgapp`, `circuit-compile`, `circuit-profile`, `protocol-vectors` |
| `tests/` | Chain-backed integration tests, fixtures, the battery |
| `docs/pool-keys.md` | Normative encodings and proof statements |
| `BENCHMARKS.md` | Constraints, proving times, memory, gas, RPC load |

---

## References

- Groth16: J. Groth, *On the Size of Pairing-based Non-interactive Arguments*, EUROCRYPT 2016.
- Feldman VSS: P. Feldman, *A Practical Scheme for Non-interactive Verifiable Secret Sharing*, FOCS 1987.
- DAVINCI voting protocol: https://davinci.vote
- Vocdoni: https://vocdoni.io
