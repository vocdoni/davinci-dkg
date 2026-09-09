# DAVINCI DKG — Benchmarks

Measurements of the current release: circuits at `MaxN = 32`, `MaxT = 32`, `MaxK = 16`
(gnark v0.16.3, BN254 Groth16), contracts compiled with solc 0.8.28 via IR, node v0.7.
Reproduce constraint counts with `go run ./cmd/circuit-profile <circuit>` (then
`go tool pprof -sample_index=0 -top /tmp/<circuit>.pprof` for the per-gadget split), proving
times and memory with the circuit tests under `/usr/bin/time -v`, gas from the integration
suite's receipts (`RUN_INTEGRATION_TESTS=true go test ./tests/...`), and node figures from the
container cgroups of a running fleet.

## Circuits

| Circuit | Constraints | Public inputs | Proving key | Proof time |
|---|---:|---:|---:|---:|
| Contribution (16 keys, compact transcript) | 1,689,551 | 8 | 243 MB | 1,035 ms |
| Finalize (all 16 keys, up to 32 dealers) | 2,228,441 | 7 | 436 MB | 2,369 ms |
| PartialDecrypt | 26,194 | 15 | 4 MB | 33 ms |
| DecryptCombine | 255,081 | 9 | 35 MB | 146 ms |

Proof times are the mean of five proofs from each circuit package's `BenchmarkProve`
(`go test -run '^$' -bench '^BenchmarkProve$' -benchtime 5x ./circuits/<circuit>` with the pinned
keys in `DAVINCI_ARTIFACTS_DIR`), wall-clock on the 32 threads of an idle AMD Ryzen 9 9950X3D
(64 GiB), witness solving included, key loading excluded; the raw outputs, wrapped in
`/usr/bin/time -v`, are in `docs/benchmarks/prove-2026-09-09/`. Observed on nodes: the Sepolia seed
nodes log about 1.3 s per contribution proof, and the integration suite 3.7 s per finalization
with two nodes proving concurrently. Verification is a constant few milliseconds. The whole
artifact release (`circuits-v6`: compiled circuits, proving and verifying keys) is 1.0 GB.

Where the constraints go (gnark's constraint profiler, `go run ./cmd/circuit-profile <circuit>` then
`go tool pprof -sample_index=0 -top /tmp/<circuit>.pprof`; attributions overlap where gadgets nest,
so they are shares of the total, not an additive breakdown). Contribution: the 512 constant-position
Horner evaluations of the commitment polynomials 29%, the 560 fixed-base products (shares,
ephemerals, constant terms) 29%, the 528 canonical scalar decompositions 20%, the 512 mask hashes
8%, the 32 ECDH products 6%, the digests 5%, the 496 cofactor certificates 0.6%. Finalize: 63% is
the per-key Poseidon digest that re-absorbs every dealer's 2·16·32 commitment words, 22% the
constant-position Horner evaluation of the aggregated polynomials. Profiler attribution per call of
the shared gadgets: `FixedBaseMulBits` 875 (plus the caller's scalar decomposition, 254 bits or a
639-constraint `CanonicalScalarBits`), `ScalarMulVarBits` 3,291, `AssertCofactorPreimage` 21,
`AssertPointOnCurve` 4, Poseidon ≈ 43 per absorbed element. Every circuit also
spends one row per public input (8, 7, 15 and 9) on `ccommon.CertifyPublicInputs`, which puts each
public wire alone on the left-hand side of a constraint so that the QAP meets the hypothesis of
Groth16's weak simulation-extractability theorem (`MissingDedicatedPublicRows` checks the compiled
matrices in each circuit's `qap_test.go`). The circuit size is fixed by the compiled capacities
(`MaxN`, `MaxT`, `MaxK`); the live `(t, n)` size only the calldata. A single-key build (`MaxK = 1`)
of the same contribution circuit has 271,656 constraints.

## Memory

| Figure | Value |
|---|---:|
| Contribution prover, peak resident set of the `BenchmarkProve` process (key and circuit loaded, `/usr/bin/time -v`, which reports KiB) | 3.0 GiB |
| Finalize prover, same measurement | 5.4 GiB |
| Node at rest (cgroup, Sepolia seed nodes) | 0.2–0.7 GB |
| Node peak during its contribution proof | 2.9 GB |
| Node peak at startup (stream-verifies the four circuits, decodes the two decryption runtimes) | 0.17 GB |
| Node average CPU (Sepolia, 24-hour epochs, public bot load) | 0.05 vCPU |

Proving keys are decoded for a proof and released afterwards, so the resting footprint is the two
small decryption circuits plus Go heap headroom. Hardware guidance: 2 cores and 4 GB of RAM
minimum, 8 GB comfortable; `GOMEMLIMIT` tightens the peak at some CPU cost.

## Gas

Measured on the integration harness (Anvil, Cancun) by `TestGasProfilesMultiNode`
(`RUN_INTEGRATION_TESTS=true RUN_BENCHMARKS_MULTI=true go test -run TestGasProfilesMultiNode ./tests/`,
receipts' `gasUsed`; raw sweep in `docs/benchmarks/gas-k16-2026-09-09.txt`). Groth16 verification gas
does not depend on circuit size, so these hold for any circuit build with the same public inputs.

| Call | n = 4, t = 3 | n = 32, t = 22 | Notes |
|---|---:|---:|---|
| `createEpoch` | 133,312 | 133,312 | 150,323 for the first seed |
| `claimSlot` (average) | 137,463 | 121,021 | |
| `submitContribution` | 500,286 | 1,315,171 | 16 keys; `32·(K·(2t+n)+5n)` bytes of transcript calldata |
| `finalizeEpoch` | 2,201,901 | 2,810,569 | verifier, 16 keys and 16 roots stored, 16 Merkle trees |
| `registerApplication` locked / automatic | 619,157 / 232,394 | 621,089 / 232,394 | locked verifies a Schnorr proof of possession and a subgroup check |
| `revealOrganizerSecret` | 222,690 | 225,262 | |
| `submitCiphertext` | 102,770 | 102,770 | no proof; on-curve and canonical checks only |
| `submitPartialDecryption` | 402,263 | 402,271 | verifier plus a 5-level Merkle path |
| `combineDecryption` | 410,626 | 483,931 | |

Observed on Sepolia: `submitContribution` 550,413 at `n = 6`, `finalizeEpoch` 2,170,777 to
2,180,237 at `n = 3`.

An epoch that yields all 16 keys costs about `0.15 + 4 × 0.14 + 4 × 0.50 + 2.20 ≈ 4.9 M` gas at
`n = 4` (0.31 M per key) and about `32 × (0.12 + 1.31) + 2.81 ≈ 48.9 M` at `n = 32` (3.1 M per
key). Decrypting one ciphertext costs one `submitCiphertext`, `t` partials and one combine:
about 1.7 M gas at `t = 3`.

### Single-key baseline (`MaxK = 1`)

The same circuits, contracts and test built with `MaxK = 1` / `MAX_K = 1` (worktree patch and raw
log: `docs/benchmarks/gas-k1-2026-09-09.patch`, `gas-k1-2026-09-09.txt`): the contribution circuit
has 271,656 constraints; `submitContribution` costs 397,392 at
`n = 4` and 529,452 at `n = 32`, `finalizeEpoch` 461,523 and 747,571, `createEpoch` and `claimSlot`
are unchanged. An epoch yielding one key therefore costs about 2.7 M gas at `n = 4` and 21.7 M at
`n = 32`, nine and seven times the pooled per-key figure above, and dealing 16 keys costs 1.26–2.5×
dealing one.

In fiat, at the prices of 6 September 2026 (ETH $2,482, POL $0.098; base fees Ethereum 0.053 gwei,
Arbitrum 0.020, Base 0.006, Polygon 278 gwei, zkSync Era about 0.045 gwei) for a committee of
`n = 8, t = 6`: one key costs about $0.10 on Ethereum, $0.04 on Arbitrum, $0.01 on Base, $0.02 on
Polygon and roughly $0.09 on zkSync Era; one decryption about $0.38, $0.15, $0.04, $0.08 and
$0.33. Ethereum scales linearly with the base fee (×19 at 1 gwei). The zkSync figures are
estimates, since its fee also depends on pubdata.

## RPC load

A node makes about five JSON-RPC calls per 30-second tick at rest: one `eth_getBlockByNumber`,
one combined `eth_getLogs` over both contracts, two to three `eth_call`, plus a balance read every
20 ticks and a liveness read every 50 blocks. Finalization and share recovery batch their reads.
Eight seed nodes together put about 1.3 requests per second on the public Sepolia endpoints.
