# Contributing

Issues and pull requests are welcome. For anything larger than a fix, open an issue first so the
design can be agreed before the work starts. Protocol changes usually touch Go, Solidity and
TypeScript at once; read [docs/protocol.md](docs/protocol.md) and
[docs/pool-keys.md](docs/pool-keys.md) before changing an encoding or a check.

## Prerequisites

- Go 1.25+, [gofumpt](https://github.com/mvdan/gofumpt) and
  [golangci-lint](https://golangci-lint.run) v2.5
- [Foundry](https://getfoundry.sh) (solc 0.8.28 is fetched by `forge`)
- Node.js 20 and pnpm 10 for the SDK and the explorer
- Docker, for the integration tests and the local testnet

## Repository layout

| Path | Contents |
|---|---|
| `cmd/` | `davinci-dkg-node`, `dkgapp`, and the tooling: `circuit-compile`, `circuit-profile`, `constraints`, `protocol-vectors`, `operator-schnorr-vectors`, `sdk-test-fixture` |
| `node/` | The daemon: epoch participation, finalization turns, decryption scanner |
| `finalizer/` | Rebuilds the accepted contributions from calldata and proves `finalizeEpoch` |
| `circuits/` | The four gnark circuits, shared gadgets (`common/`) and the pinned artifact loader |
| `crypto/` | Off-circuit primitives: Feldman, Shamir, Schnorr, ElGamal, share encryption, Poseidon |
| `web3/` | Typed wrappers over the generated bindings, RPC pool, transaction manager |
| `config/` | Network presets and pinned circuit artifact hashes |
| `solidity/` | Contracts, interfaces, verifiers, Foundry tests, deploy scripts, Go bindings |
| `sdk/`, `ui/` | TypeScript SDK and explorer |
| `tests/` | Chain-backed integration tests, cross-implementation vectors, the battery |
| `testnet/` | Local multi-node testnet (Anvil, deployer, nodes) |

## Building and testing

```bash
make build                                  # go build ./cmd/...
go test $(go list ./... | grep -vE 'davinci-dkg/(tests|circuits)') -failfast   # unit tests, as CI runs them
go test ./crypto/schnorr -run TestName      # a single test
make test                                   # unit tests including the circuits (slow, see below)
```

The circuit packages compile and set up their circuits; the contribution setup alone takes about
ten minutes, and the results are cached under `~/.davinci/artifacts` (`DAVINCI_ARTIFACTS_DIR`
overrides it). CI runs them in a separate workflow when `circuits/`, `crypto/` or the pinned
hashes change.

Integration tests run the contracts on Anvil in Docker, started by the harness from
`tests/docker/docker-compose.yml`:

```bash
make circuits-compile circuits-update-hashes solidity-build   # see the note below
RUN_INTEGRATION_TESTS=true go test ./tests/... -timeout 2h -failfast -count=1
```

The Groth16 setup is randomized, so a locally compiled circuit never matches the hashes in
`config/circuit_artifacts.go` or the verifying keys in `solidity/src/verifiers/`. Regenerate all
three together before the integration tests, as above, as CI does; a mismatch makes every proof
revert with `ProofInvalid()` (selector `0x7fcdd1f4`). Do not commit the resulting changes unless
you are cutting a circuit release.

Contracts, SDK and explorer:

```bash
(cd solidity && forge build && forge test)
(cd sdk && pnpm install && pnpm build && pnpm check && pnpm test)   # pnpm test:integration needs Docker and the artifacts
(cd ui && pnpm install && pnpm lint && pnpm test)                   # or: make ui-test
```

### What CI checks

- `go mod tidy` leaves no diff, `go vet ./...`, `gofumpt -l .` is empty, `golangci-lint run`
  (configuration in `.golangci.yml`, lines up to 130 characters).
- Go unit tests, circuit tests, and the integration tests including the SDK's.
- `forge build` and `forge test`.
- The explorer: `pnpm lint`, `pnpm test`, `pnpm build`, and `scripts/render-ui-config.test.sh`.
- The battery against a four-node testnet ([tests/battery](tests/battery/README.md)).

Run `make vectors-check` as well when you touch an encoding: it regenerates the
cross-implementation vectors and fails if they changed.

## Local testnet

```bash
make testnet-up                                   # Anvil, deployer and 3 nodes
make testnet-up DKG_NODE_COUNT=8 DKG_THRESHOLD=5  # other sizes, up to 32 nodes
make testnet-logs
make testnet-down                                 # stops everything and wipes the volumes
```

The deployer publishes the contract addresses at `http://127.0.0.1:8888/addresses.env`. Point the
explorer at the testnet with:

```bash
make ui-dev RPC_URL=http://127.0.0.1:8545 MANAGER_ADDRESS=<manager> CHAIN_ID=1337 CHAIN_NAME=anvil DEPLOY_BLOCK=0
```

`testnet/remote-nodes.sh` runs more nodes on another host against the same chain. The
[battery](tests/battery/README.md) drives a running testnet through load and adversarial
scenarios.

## Keeping implementations in sync

The protocol is implemented in Go (node and circuits), Solidity and TypeScript (SDK and
explorer). A change to any of the following must land in every place listed, in one pull request:

| What | Where |
|---|---|
| Fiat–Shamir domain strings | `internal/protocol/protocol.go`, `solidity/src/libraries/DKGProtocol.sol`, `sdk/src/protocol.ts` |
| `MaxN`, `MaxT`, `MaxK`, Merkle depth | `circuits/common/sizes.go`, `solidity/src/libraries/Sizes.sol`, `sdk/src/sizes.ts`; then `make circuits`. `MaxN` must be a power of two |
| Transcript layouts and BRLC | `circuits/common/brlc.go`, the circuits' `witness.go`, `solidity/src/libraries/BRLC.sol`, `sdk/src/transcript.ts` |
| Share-commitment Merkle tree | `circuits/common/merkle.go`, `DKGManager.sol`, `sdk/src/merkle.ts` |
| Network presets | `config/networks.go`, `sdk/src/networks.ts`, `docs/deployments.md` and the README; moving the default network also means `ui/public/config.json`, the fallbacks in `scripts/render-ui-config.sh` and `ui/.do/davinci-dkg-ui.yaml` |

`make vectors` regenerates `tests/vectors/*.json` from the Go code and copies them to
`ui/tests/vectors/`; the SDK, Foundry and explorer tests assert against them. A one-bit
divergence in a transcript or Merkle leaf makes every proof or partial revert.

## Changing the circuits

```bash
make circuits    # compile, set up, rewrite the verifiers, pin the hashes, rebuild, regenerate bindings
```

A circuit change needs a new artifact release: the `Publish Circuits` workflow (label a pull
request `trigger-upload-circuits`) compiles the circuits, pins the hashes and stages the files for
a `circuits-vN` GitHub release; bump `DefaultArtifactsRelease` in `config/circuit_artifacts.go` to
match, and upload the same files to the CDN under `dkg/circuits-vN/` (bucket `davinci-assets`,
fra1, public-read), which nodes try before the GitHub release. `go run ./cmd/circuit-profile <circuit>` shows where the constraints go.

gnark releases through v0.15.0 have an unsound variable-base twisted-Edwards scalar
multiplication that a prover can satisfy for any output point (fixed in v0.16.0). The project
requires gnark v0.16.2 or later and never downgrades it. Any gnark upgrade changes the compiled
constraint systems and therefore needs a circuit release.

## Generated files

Regenerate these instead of editing them:

- `solidity/golang-types/*.go`: `make solidity-bind`
- `solidity/src/verifiers/*`: `make circuits-compile`
- `tests/vectors/*.json`, `ui/tests/vectors/*.json`: `make vectors`

## Commits and releases

Commit messages follow [Conventional Commits](https://www.conventionalcommits.org):
`type(scope): summary`, for example `fix(node): retry a failed finalization`. Pushing a `vX.Y.Z`
tag builds the Linux binaries, publishes the node and explorer images and creates the GitHub
release; `latest` moves on stable tags only, and release candidates (`vX.Y.Z-rc1`) publish their
own tag.
