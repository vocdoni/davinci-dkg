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
go test ./circuits/contribution -run TestContributionCircuitProveAndVerify  # heavy: does a trusted setup
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
make testnet-up DKG_NODE_COUNT=3 DKG_THRESHOLD=2   # Anvil + deployer + N nodes + UI on :8081
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
- **Committee size cap `MaxN`**: `circuits/common/sizes.go` and `solidity/src/libraries/Sizes.sol`, then
  `make circuits`.
- **BRLC transcript encoding** (contribution calldata): `circuits/common/brlc.go`, `web3/brlc.go` and
  `solidity/src/libraries/BRLC.sol` must agree bit-for-bit, including the challenge anchor
  `keccak(digests ‖ keccak(transcript))` (`ccommon.ChallengeAnchor`).
- **Organizer-share DLEQ challenge** `e = keccak(DOMAIN_ORGANIZER_SHARE_V1 ‖ eid ‖ aid ‖ ctIdx ‖ PK_org ‖
  C1 ‖ Δ ‖ A1 ‖ A2) mod q`: `crypto/dleq/organizer.go`, `DKGManager._verifyOrganizerWords` and
  `sdk/src/dleq.ts`. The combine circuit consumes `e` as a transcript word and the contract recomputes
  it from calldata, so a one-byte divergence makes every combine revert.
- **Circuit ↔ verifier binding**: `config/circuit_artifacts.go` pins SHA-256 hashes of every circuit's
  ccs/pk/vk; `circuits/artifacts.go` verifies them on load (downloads from the CDN, else falls back to a
  local setup). `make circuits-update-hashes` keeps the file in sync.

### Go packages

- `node/` is the daemon (`cmd/davinci-dkg-node` is a thin main). `node.go` polls the two newest epochs
  and does claimSlot → submitContribution (Groth16) → auto-finalize (via `finalizer/`); `decrypt.go`
  scans `CiphertextSubmitted` events for every aid and checks the prime-subgroup membership of C1/C2
  before computing a partial — the contract deliberately skips that check (~2 M gas), so this one is
  load-bearing, not belt-and-braces. A slot becomes combinable only when `t` partials **and** a
  verifying organizer share are on chain: `latestOrganizerShare` reads the newest
  `OrganizerShareSubmitted` event, verifies it with `dleq.VerifyOrganizerShare` against the
  application's registered `PK_org` (read through `DKGAppManager.getApplication`), and skips the slot
  until the next tick if it does not verify — the organizer may overwrite a malformed share, so this
  must never be terminal. The combine then runs on a seed-derived stagger anchored on whichever of the
  two prerequisites landed last. All secret scalars come from `scalars.go` (`crypto/rand`, never
  deterministic); `dlog.go` is a compact parallel BSGS (2^50 cap, ~256 MB). Every flag has a
  `DAVINCI_DKG_*` env equivalent (`config.go`). `--network sepolia` resolves the manager from
  `config/networks.go`; registry, verifiers and app manager are read from the manager on-chain.
- `cmd/dkgapp` is the application/organizer CLI: `register` (organizer key + Schnorr PoP; generates and
  prints `sk_org` when not given), `encrypt` (ElGamal under `PK_aid = PK_ep + PK_org`, no proof),
  `share` (posts `Δ = sk_org·C1` with its keccak-challenge DLEQ — no SNARK, no circuit artifacts),
  `plaintext`. Losing `sk_org` makes the application permanently undecryptable.
- `circuits/{contribution,finalize,partialdecrypt,decryptcombine}` — gnark circuits (BN254 / BabyJubJub /
  Poseidon). Each has `circuit.go`, `witness.go` (Go-side witness construction), `assignment.go`,
  `artifacts.go`. `circuits/common` has the shared point/hash/Lagrange gadgets.
- `crypto/` — off-circuit primitives (feldman, shamir, schnorr, dleq, shareenc, group, hash).
- `web3/` — thin typed wrappers over the abigen bindings in `solidity/golang-types` plus an RPC pool.
  No business logic. `web3/txmanager` allocates nonces locally inside the signer hook (safe for
  concurrent goroutines), applies gas headroom to estimates, rebroadcasts and fee-bumps stuck txs.
  Callers must `RecordPending` after sending and `WaitTxByHash`.
- `finalizer/` — reconstructs accepted contributions from calldata, proves finalize, submits.
- `prover/` — Groth16 backend wrapper; `GPU_PROVER=true` switches to the icicle backend.
- `types/` — neutral structs shared by web3/finalizer/node to avoid import cycles.

### Contracts (`solidity/src`)

`DKGRegistry` (operators, liveness) → `DKGManager` (epoch state machine, ciphertexts, decryption) →
`DKGAppManager` (per-app registration, organizer shares). Split only for EIP-170; they share one
logical storage. Phase windows and policy floors are constructor immutables (`EPOCH_DURATION_BLOCKS`,
`COMMITTEE_SELECTION_BLOCKS`, `KEY_ASSEMBLY_BLOCKS`, `FINALIZE_GAP_BLOCKS`, `MIN_THRESHOLD`,
`MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS`; all settable via env in `DeployAll.s.sol`).
Invariants worth knowing: only operators registered before `createEpoch` may claim; `abortEpoch` only
works on provably dead epochs; ciphertext indices are assigned on chain; `submitCiphertext` takes no
proof and only accepts a registered `aid` (there is no `aid = 0` path and no open submission — a zero
`authorizedSubmitter` resolves to the registrant); `submitOrganizerShare` is permissionless and
unverified on chain (it only stores `keccak(Δ ‖ A1 ‖ A2 ‖ z)`, and re-submission overwrites until the
plaintext lands, so a malformed share cannot brick a ciphertext) — the DLEQ is verified inside the
combine SNARK against the `e` the contract recomputes; every BRLC challenge is
keccak over the calldata *and* the circuit's digests (Fiat–Shamir on both sides, see BRLC.sol). `interfaces/*.sol` is the
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
  Every decryption test registers an application with `helpers.RegisterApplication` and posts an
  organizer share (`helpers.SubmitOrganizerShareAs` for a fresh one, `helpers.PostOrganizerShare` to
  publish the exact words a combine output already committed to — the DLEQ nonce is fresh per call, so
  re-proving would not match the stored hash).
- `sdk/tests/*-e2e.test.ts` start their own Anvil stack and use `cmd/sdk-test-fixture` to produce
  proofs.
- `solidity/test/TestHelpers.t.sol` + `TestInputs.t.sol` hold canned proofs/inputs for Foundry.

## Conventions

- Formatting is `gofumpt`; lint is `golangci-lint` with revive `enable-all-rules` (see `.golangci.yml`
  for the disabled ones). Max line length 130.
- Generated files: `solidity/golang-types/*.go`, `solidity/src/verifiers/*`, `tests/vectors/*.json`,
  `ui/tests/vectors/*.json`. Regenerate, don't edit.
- New production deployments go in `config/networks.go` (`KnownNetworks`) and `README.md` Deployments.
