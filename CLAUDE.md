# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

Reference Go implementation of NI-DKG (non-interactive distributed key generation on EVM via Groth16
proofs), plus the Solidity contracts, a TypeScript SDK and a React explorer UI. `README.md` explains
the protocol model (epoch lifecycle, lottery, threshold decryption, per-app keys) — read it before
touching protocol logic. Requires Go 1.25+, Foundry, pnpm 10, Docker (for integration tests).

## Commands

### Go

```bash
make build                                   # go build ./cmd/...
make test                                    # unit tests, excludes ./tests (integration)
go test ./crypto/schnorr -run TestName       # single test
go test ./circuits/contribution -run TestContributionCircuitProveAndVerify -timeout 120m  # heavy: ~10 min setup, 762 MB v4 contribution proving key
go vet ./... && gofumpt -l . && golangci-lint run   # what CI enforces (golangci-lint v2.5, config in .golangci.yml)
go mod tidy                                  # CI fails if this produces a diff
```

Integration tests need Docker and are opt-in:

```bash
RUN_INTEGRATION_TESTS=true go test ./tests/... -timeout=2h -failfast -count=1   # == make test-integration
```

The harness (`tests/helpers/service.go`) brings up `tests/docker/docker-compose.yml` (Anvil + a deployer
container) via testcontainers. **Gotcha:** Groth16 setup is randomized, so the committed
`config/circuit_artifacts.go` hashes and `solidity/src/verifiers/*_vkey.sol` never match artifacts on a
fresh machine. If integration tests fail with `ProofInvalid()` (selector `0x7fcdd1f4`), run
`make circuits-compile circuits-update-hashes solidity-build` first so the Go proving key, pinned hashes
and Solidity verifier come from the same setup run. Don't commit the resulting churn unless you mean to
cut a new circuit release.

### Solidity (`solidity/`)

```bash
cd solidity && forge build      # via_ir, optimizer_runs=1, solc 0.8.28
forge test                       # forge test --match-test test_Name for one test
make solidity-bind               # regen Go bindings in solidity/golang-types via abigen (needs forge build first)
make solidity-deploy             # solidity/deploy_all.sh; reads RPC_URL/CHAIN_ID/PRIVATE_KEY from .env
```

### Circuits

```bash
make circuits          # compile all → ~/.davinci/artifacts, rewrite Solidity verifiers, patch hashes, forge build, abigen
make vectors           # regenerate tests/vectors/*.json (cross-impl fixtures); CI runs `make vectors-check`
```

### SDK (`sdk/`) and UI (`ui/`)

```bash
cd sdk && pnpm install && pnpm build && pnpm test          # vitest; pnpm test:integration needs Docker + circuit artifacts
cd sdk && pnpm check                                        # tsc for src + tests
make ui-dev / ui-build / ui-test                            # builds sdk first; UI depends on sdk via link:../sdk
cd ui && pnpm lint                                          # tsc --noEmit && eslint (CI)
```

`ui/public/config.json` picks the chain (defaults to Sepolia). `make ui-dev RPC_URL=... MANAGER_ADDRESS=...`
templates it via `scripts/render-ui-config.sh`.

### Local multi-node testnet

```bash
make testnet-up DKG_NODE_COUNT=3 DKG_THRESHOLD=2   # Anvil + deployer + N nodes (no UI service; browse via `make ui-dev`)
make testnet-logs / testnet-down
```

## Architecture

### Four implementations that must stay byte-compatible

The protocol is implemented in Go (node + circuits), Solidity (contracts), TypeScript (SDK) and the UI.
Anything that touches encodings, hashes or constants has to be changed in all of them:

- **Protocol constants** (Fiat–Shamir domain prefixes): source of truth is
  `internal/protocol/protocol.go`; mirrors in `solidity/src/libraries/DKGProtocol.sol` and
  `sdk/src/protocol.ts`. `cmd/protocol-vectors` emits `tests/vectors/*.json` from the Go side; the SDK
  (`sdk/tests/vectors.test.ts`) and Foundry (`DKGProtocol.t.sol`) assert against them.
- **Committee size cap `MaxN`, pool size `MaxK`, Merkle depth**: `circuits/common/sizes.go` and
  `solidity/src/libraries/Sizes.sol` (`MAX_N=32`, `MAX_K=16`, `MERKLE_DEPTH=5=log2(MaxN)`; `MaxN` must be a power of two, so 16, 32 or 64), then
  `make circuits`. See `docs/pool-keys.md` for the normative pool-key spec.
- **BRLC transcript encoding**: `circuits/common/brlc.go`, `web3/brlc.go` and
  `solidity/src/libraries/BRLC.sol` must agree bit-for-bit. The v4 contribution transcript
  (`CompactContributionWords(t,n) = MaxK·(2t+n) + 5n`) is **compact** — no padding travels in
  calldata; each commitment coordinate is gated by `b=[m<t]`, every recipient-index, public-key,
  ephemeral and masked-share entry by `b=[i<n]`, and the BRLC fold neither counts nor advances
  the exponent for inactive entries. Its challenge anchor is unchanged: `keccak(commitmentsHash ‖
  encryptedSharesHash ‖ keccak(transcript))`. The batched finalization (`circuits/finalize`,
  domain `davinci-dkg:finalize:v2`, `FINALIZE_TRANSCRIPT_WORDS = 2·MaxN + MaxK·(2+2·MaxN)` = 1,120
  words at N=32, K=16) has 7 public inputs `[eid, threshold, committeeSize, acceptedCount,
  transcriptDigest, challenge, transcriptCommitment]`, where
  `transcriptDigest = Poseidon MultiHash(eid, t, n, a, K, L_F, R, B_0…B_{K−1})` over the masked
  transcript words, and the anchor `keccak(transcriptDigest ‖ keccak(transcript))` — the same
  discipline as contribution and combine. Do not drop the digest again: with
  `keccak(keccak(transcript))` alone the challenge never depends on the witness words, the
  calldata aggregate / share-commitment region (where the contract reads `P_j` from) is bound to
  the proof only by the BRLC linear relation, and a permissionless finalizer can grind a calldata
  transcript carrying a forged `P_j` that still verifies (a generalized-birthday search, not
  `2^128` work) — see `docs/pool-keys.md`. Combine
  (`davinci-dkg:decrypt-combine:v1`, `COMBINE_TRANSCRIPT_WORDS = 6 + 3·MaxN`) is unchanged in shape but
  now carries `PK_org` and proves knowledge of `OrganizerSecret` instead of a DLEQ, and pins the
  Lagrange coefficients to the canonical vector of the qualifying set (masked by `shareCount`). On
  every proof-carrying call the BRLC calldata commitment refuses a non-canonical word (`>= p`), so a
  transcript has exactly one encoding.
- **Share-commitment Merkle tree**: the finalization transcript's per-key region holds `D_j,i`
  (share commitment for committee position `i+1`) for every `j < MaxK`, with `D_j,i = (0,1)` for
  `i >= committeeSize` — every member, not only the contributors, because a member that did not
  contribute still received a share from every accepted dealer (so decryption liveness is `n − t`,
  not `m − t`). Leaves `keccak256(0x00 ‖ D.x ‖ D.y)`, empty leaves
  `keccak256("davinci-dkg:merkle-empty:v1")`, internal nodes
  `keccak256(0x01 ‖ left ‖ right)`, `MERKLE_DEPTH = 5` levels over `MAX_N = 32` leaves — computed by
  `finalizeEpoch` from the finalization transcript calldata (stored as `poolShareRoots[eid][j]`),
  checked again by `submitPartialDecryption`'s trailing `shareProof` (`MERKLE_DEPTH` siblings,
  bottom-up, leaf index `participantIndex − 1`) and by the node and the SDK when they build that
  proof. A one-bit divergence in leaf order, prefix byte or hash order makes every partial revert.
- **Circuit ↔ verifier binding**: `config/circuit_artifacts.go` pins SHA-256 hashes of every circuit's
  ccs/pk/vk; `circuits/artifacts.go` verifies them on load (downloads from the CDN, else falls back to a
  local setup). `make circuits-update-hashes` keeps the file in sync.
- **gnark pin** (`go.mod` invariant): gnark v0.16.3 / gnark-crypto v0.21.0. **Never downgrade gnark
  below v0.16.2.** Every gnark release up to and including v0.15.0, and the snapshot this repo pinned
  before (`v0.14.1-0.20260126…`), has an unsound variable-base twisted-Edwards `ScalarMul`
  (`std/algebra/native/twistededwards`, `scalarMulFakeGLV`): the fake-GLV decomposition check
  `s1 + s2·s = k·order` is evaluated in the native field with the quotient `k` a free hint output, so it
  holds for any `(s1, s2)` and a malicious prover can make `ScalarMul(P, s)` return any point. Reproduced
  with verifying Groth16 proofs on v0.14.0, v0.15.0 and the snapshot (repro, outputs and the upstream
  report draft live in `~/davinci-dkg-gnark-repro`, outside the repo). Upstream fixed it in PR #1765
  (gnark v0.16.0: emulated-field decomposition check plus subgroup binding) without an advisory —
  GHSA-3mvx-pp85-pm65 does not list it; the related but far weaker cofactor-torsion offset is IACR ePrint
  2026/1776. Every circuit is compiled and re-pinned against v0.16.3 and none uses the hinted gadget any
  more (`ccommon.ScalarMulVar`); any gnark bump changes the compiled R1CS and therefore means
  `make circuits` plus a circuit release.

### Go packages

- `node/` is the daemon (`cmd/davinci-dkg-node` is a thin main). `node.go` polls the two newest epochs
  and does claimSlot → submitContribution (Groth16, dealing all `MaxK` pool-key polynomials at once,
  compact `MaxK·(2t+n) + 5n`-word transcript) → one proof-carrying `finalizeEpoch` once the epoch
  qualifies. There is no per-key activation any more: the finalization proof stores every pool key
  and share root atomically, so a `Live` epoch is fully usable. The node runs **one** finalization
  attempt per epoch in the existing seed-derived stagger (slot `mySlot`, anchored at
  `liveNotBeforeBlock`); the first proof to land wins and the rest see `AlreadyLive` and stop. A
  transient failure is retried with an exponential per-epoch backoff (the proof is expensive, so a
  persistent fault must not cost one per tick). A qualifying epoch that was never finalized stays
  discoverable across cadence changes and restarts. It also creates the next epoch early — bypassing
  the normal cadence gate — when the newest epoch has at most one unclaimed key
  (`poolNext >= MAX_K - 1`, i.e. `poolNext >= 15`) or is `Aborted`.
  `decrypt.go` scans `CiphertextSubmitted` events for every aid and checks the prime-subgroup
  membership of C1/C2 before computing a partial — the contract deliberately skips that check (about
  0.17 M gas for `C1`; skipped to keep submission cheap, not because it is prohibitive), so this one is
  load-bearing, not belt-and-braces. Partials are posted in seed-derived waves of
  `t` members (`staggerSlot / t`); a later wave only posts if, `staggerBlocks` later, fewer than `t`
  partials are on chain **and** the earlier waves have stopped landing partials (`laterWaveDue`), so an
  honest ciphertext costs `t` partials, not `n`, even under load; each partial now carries a
  `MERKLE_DEPTH`-long Merkle path proving its share commitment against the claimed pool key's
  `poolShareRoots` entry. Partial and combine transactions are sent without waiting for the receipt
  (`inflight`, settled next tick) so one tick serves every pending ciphertext. A slot becomes
  combinable once `t` partials are on chain, the application's decryption window is open
  (`decryptNotBefore <= now <= decryptNotAfter`, else `DecryptionNotOpen()` / `DecryptionClosed()`),
  and — for an organizer-locked application — the organizer has called `revealOrganizerSecret` (a
  one-time, whole-application act; automatic applications need no reveal at all, since `OrganizerSecret`
  is `0` from registration). The contract enforces the same gates on partials (`requireDecryptionOpen`
  reverts `OrganizerSecretNotRevealed()` until the reveal), so the node parks a locked slot *before*
  posting its partial, and parks a slot whose window has not opened; a parked slot costs nothing per
  tick until the relevant block or an `OrganizerSecretRevealed` event wakes it (the wake rescans the
  application's ciphertexts from its registration block), because epochs stay Live on chain forever. The combine (dlog search, proof, send) runs in a per-slot goroutine, one at a time
  per node (`combineSem`), yielding to an in-progress contribution or finalization (`critical`). A
  ciphertext whose plaintext is out of range taints its source for the epoch (`taints`, persisted
  in `<datadir>/tainted-apps.json`): always the offending (application, submitter) pair, so an
  attacker pays one search per submitter address and cannot silence an application for its
  honest submitters. All
  secret scalars come from `scalars.go` (`crypto/rand`, never deterministic); `dlog.go` is a compact
  parallel BSGS (2^50 cap, ~256 MB). Every flag has a `DAVINCI_DKG_*` env equivalent (`config.go`).
  `--network sepolia` resolves the manager from `config/networks.go`; registry, verifiers and app
  manager are read from the manager on-chain.
- `cmd/dkgapp` is the application/organizer CLI: `register` (`-mode locked|automatic`, default locked:
  locked = organizer key + Schnorr PoP, automatic = no organizer key at all; submission policy
  `-submitters 0xA,0xB` (exclusive allow-list, ≤ 32) / `-open` / neither = registrant only; `-max`;
  `-decrypt-from` / `-decrypt-until <RFC3339 | Go duration such as 48h>` set the decryption window),
  `reveal -aid -org-secret` (locked apps only, permissionless, once: posts `revealOrganizerSecret` and
  from then on the committee combines by itself, application-wide, not per-ciphertext), `encrypt`
  (ElGamal under `PK_aid = P_j` or `P_j + PK_org`, no proof), `plaintext`, `epoch` (shows pool status:
  which keys are claimed vs. free — `getPoolStatus` returns the `poolNext` cursor; there is no
  activation bitmap any more). Losing `sk_org` makes a locked application permanently
  undecryptable; never reuse an organizer secret across two locked applications (revealing it for one
  opens the other).
- `circuits/{contribution,finalize,partialdecrypt,decryptcombine}` — gnark circuits (BN254 / BabyJubJub /
  Poseidon). Each has `circuit.go`, `witness.go` (Go-side witness construction), `assignment.go`,
  `artifacts.go`. `circuits/common` has the shared point/hash/Lagrange gadgets.
- `crypto/` — off-circuit primitives (feldman, shamir, schnorr, dleq, shareenc, group, hash).
- `web3/` — thin typed wrappers over the abigen bindings in `solidity/golang-types` plus an RPC pool.
  No business logic. `web3/txmanager` allocates nonces locally inside the signer hook (safe for
  concurrent goroutines), applies gas headroom to estimates, rebroadcasts and fee-bumps stuck txs.
  Callers must `RecordPending` after sending and `WaitTxByHash`.
- `finalizer/` — one batched finalization operation: it reconstructs every accepted contribution's
  calldata at a single canonical block after the contribution deadline (recovered from the event log
  via the `CalldataCache`), proves the `circuits/finalize` Groth16 statement, re-checks the epoch
  state right before sending (a finalization that landed meanwhile is reported as `ErrAlreadyLive`, a
  reorg divergence is an error the caller retries), and submits `finalizeEpoch`, which atomically
  stores all `MaxK` keys and share roots and flips the epoch `Live`.
- `prover/` — Groth16 backend wrapper; `GPU_PROVER=true` switches to the icicle backend.
- `types/` — neutral structs shared by web3/finalizer/node to avoid import cycles.

### Contracts (`solidity/src`)

`DKGRegistry` (operators, liveness) → `DKGManager` (epoch state machine, ciphertexts, pool keys,
decryption) → `DKGAppManager` (per-app registration, organizer secret reveal). Split only for EIP-170;
they share one logical storage. Phase windows and policy floors are constructor immutables
(`EPOCH_DURATION_BLOCKS`, `COMMITTEE_SELECTION_BLOCKS`, `KEY_ASSEMBLY_BLOCKS`, `FINALIZE_GAP_BLOCKS`,
`MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS`; all settable via env in
`DeployAll.s.sol`).
Invariants worth knowing: only operators registered before `createEpoch` may claim; `createEpoch` may
also run before the normal cadence only when the newest epoch is `Live` with at most one unclaimed key
(`poolNext >= MAX_K - 1`, i.e. `poolNext >= 15`) or `Aborted` — so an epoch serves at most `MAX_K = 16`
applications before registrations revert `PoolExhausted` until the next epoch has gone through its
preparation window (pool exhaustion), and anyone registering fifteen automatic applications forces a new
epoch (registration-driven amplification; a registration fee or allow-list is future work); `abortEpoch`
only works on provably dead epochs; ciphertext indices are
assigned on chain; `submitCiphertext` takes no proof and only accepts a registered `aid` (there is no
`aid = 0` path; who may submit is the app's policy — `openSubmission`, else an exclusive `submitters`
allow-list of ≤ 32, else the registrant only; contradictory policies revert `InvalidPolicy()`);
`finalizeEpoch(eid, transcriptDigest, transcript, proof, input)` is permissionless (direct-call
gated: `msg.sender == tx.origin`), one proof-carrying call per epoch while `KeyAssembly` (after the
positive finalize gap, `contributionCount >= minValidContributions`) — it decodes the 7 public inputs
`[eid, threshold, committeeSize, acceptedCount, transcriptDigest, challenge, transcriptCommitment]`,
binds positions 0…4 to state and the digest argument, derives the challenge from
`keccak(transcriptDigest ‖ keccak(transcript))` under the `davinci-dkg:finalize:v2` domain,
canonical-streams the BRLC commitment over the `FINALIZE_TRANSCRIPT_WORDS` (= 1,120 at N=32, K=16)
calldata words, validates every active dealer row against the stored `commitmentsHash`, requires the
inactive share slots to hold the identity, verifies one `circuits/finalize` proof, then — atomically,
with no partial state on revert — stores all `MAX_K` keys (`poolKeys[eid][j]`) and the Merkle root of
the whole committee's share commitments for each key (`poolShareRoots[eid][j]`), flips the epoch to
`Live` (freezing the accepted contributor set) and emits `EpochLive(eid, acceptedCount)`;
`claimPoolKey(eid, aid)` (callable only by the app manager) assigns the next unclaimed pool key,
reverts `PoolExhausted` when all `MAX_K` keys are taken, records the index-plus-one
`appPoolIndex[eid][aid]` marker and emits `PoolKeyClaimed`;
`registerApplication` claims the next unclaimed key via `claimPoolKey` (reverts `PoolExhausted` when
all `MAX_K` keys are taken; there is no activation state — `finalizeEpoch` proved and stored the whole
pool atomically, so every unclaimed key of a `Live` epoch is usable) and has two modes
(`DKGTypes.AppMode`): organizer-locked verifies a Schnorr PoP of `sk_org` and sets `PK_aid = P_j +
PK_org`, automatic stores the identity `(0, 1)` as `organizerPK` and `0` as `organizerSecret` and sets
`PK_aid = P_j` directly — there is no organizer key at all in automatic mode; `revealOrganizerSecret(eid,
aid, sk)` is permissionless, locked-apps-only, checks `sk·G == organizerPK`, may run exactly once
(`AlreadyRevealed()` after) and is application-wide, not per-ciphertext — once revealed the committee
combines by itself for every past and future ciphertext of that application; `submitCiphertext` is
gated by the block window (`notBeforeBlock` / `notAfterBlock`), the submitter policy, `maxCiphertexts`
and `decryptNotAfter` (`requireCanSubmitCiphertext`) — a ciphertext may land before decryption opens;
`submitPartialDecryption` and `combineDecryption` go through `requireDecryptionOpen`, which reverts
`DecryptionNotOpen()` before `policy.decryptNotBefore`, `DecryptionClosed()` after
`policy.decryptNotAfter` (unix seconds, 0 = unbounded) and `OrganizerSecretNotRevealed()` for a locked
application until the reveal — so no partial and no combine of a locked application exists before the
organizer reveals: the organizer learns results together with everyone and decides *when*, never
*which*, contract-enforced. State the window's guarantee honestly: it bounds what the contract accepts
and what honest nodes post; it does not bind `t` colluding members, who hold shares and can compute
partials off chain at any time (for a locked application they still lack `sk_org`), and because nothing
is accepted before the window or the reveal there are no on-chain partials from before it either;
`submitPartialDecryption` additionally verifies its trailing `shareProof` (`MERKLE_DEPTH` siblings,
leaf `keccak256(0x00 ‖ D.x ‖ D.y)` at index `participantIndex − 1`) against
`poolShareRoots[eid][appPoolIndex[eid][aid]]`; every BRLC challenge is keccak over the calldata *and*
the circuit's digests (Fiat–Shamir on both sides, see BRLC.sol), pool-key finalization included via
`transcriptDigest`. Because
every application claims a distinct pool key `P_j`, a ciphertext copied from one application into
another decrypts under an unrelated key and yields garbage, not the original plaintext — the old
cross-application decryption oracle (shared `PK_ep`) no longer exists. `interfaces/*.sol` is the
integration contract for the SDK/UI. Verifier wrappers in `src/verifiers` are generated by
`cmd/circuit-compile` — don't hand-edit.

### Test layers

- Go unit tests run without a chain. Circuit tests under `circuits/` compile + set up the circuit
  (minutes, cached under `~/.davinci/artifacts`). The node overrides the cache dir with
  `DAVINCI_DKG_ARTIFACTS_DIR`; the `circuits` package itself (and therefore tests) reads `DAVINCI_ARTIFACTS_DIR`.
- `tests/` = chain-backed Go integration (Anvil in Docker). `tests/helpers/round_flow.go` drives full
  epochs; `tests/helpers/proofs.go` builds the proofs. The harness registers only 6 operators and uses
  α=6.5535 so every actor passes the (real) lottery; `tests/node_service_test.go` runs three real
  `node.Node` instances end to end. Application ids in tests must be < the BN254 scalar field.
  Every decryption test registers an application with `helpers.RegisterApplication` (claiming the next
  unclaimed pool key — every key of a Live epoch is usable since `finalizeEpoch` stored the whole pool
  atomically) and, for organizer-locked applications, reveals the organizer
  secret once before combining; automatic applications need no such step since they have no organizer
  key.
- `sdk/tests/*-e2e.test.ts` start their own Anvil stack and use `cmd/sdk-test-fixture` to produce
  proofs.
- `solidity/test/TestHelpers.t.sol` + `TestInputs.t.sol` hold canned proofs/inputs for Foundry.

## Conventions

- Formatting is `gofumpt`; lint is `golangci-lint` with revive `enable-all-rules` (see `.golangci.yml`
  for the disabled ones). Max line length 130.
- Generated files: `solidity/golang-types/*.go`, `solidity/src/verifiers/*`, `tests/vectors/*.json`,
  `ui/tests/vectors/*.json`. Regenerate, don't edit.
- New production deployments go in `config/networks.go` (`KnownNetworks`) and `README.md` Deployments.
