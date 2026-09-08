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
| Contribution (16 keys, compact transcript) | 1,689,543 | 8 | 243 MB | 1.3 s |
| Finalize (all 16 keys, up to 32 dealers) | 2,228,434 | 7 | 436 MB | ≈ 2 s |
| PartialDecrypt | 26,179 | 15 | 4 MB | < 0.1 s |
| DecryptCombine | 255,072 | 9 | 35 MB | ≈ 0.2 s |

Proof times are wall-clock on a 32-thread AMD Ryzen 9 9950X3D with the proving key already
decoded; the contribution figure is what the Sepolia seed nodes log, the finalize figure was 3.7 s
in the integration suite with two nodes proving concurrently. Verification is a constant few
milliseconds. The whole artifact release (`circuits-v5`: compiled circuits, proving and verifying
keys) is 1.1 GB.

Where the constraints go, contribution circuit: 512 recipient shares at about 3.3 k each (a
canonical-range decomposition, the constant-index Horner evaluation of the commitment polynomial
at 1.16 k on average, the fixed-base product `s·G` at 1.39 k, one Poseidon for the mask), 32 ECDH
products at 3.8 k, 16 constant-term proofs, 496 cofactor certificates at 21, and the digests and
BRLC fold. Finalize: 61% re-derives every dealer's commitment digest from witness points (32
dealers × 16 keys × 64 field elements at 43 constraints each), 26% is the constant-index Horner
evaluation of the aggregated polynomials. Unit costs of the shared gadgets: `FixedBaseMulBits`
1,393, `ScalarMulVarBits` 3,800, `CanonicalScalarBits` ≈ 500, `AssertCofactorPreimage` 21,
`AssertPointOnCurve` 4, Poseidon ≈ 43 per absorbed element.

## Memory

| Figure | Value |
|---|---:|
| Isolated contribution prover, peak resident set (`/usr/bin/time -v` around the prove-and-verify test, compilation included) | 2.3 GB |
| Node at rest (cgroup, Sepolia seed nodes) | 0.2–0.7 GB |
| Node peak during its contribution proof | 2.9 GB |
| Node peak at startup (stream-verifies the four circuits, decodes the two decryption runtimes) | 0.17 GB |
| Node average CPU (Sepolia, 24-hour epochs, public bot load) | 0.05 vCPU |

Proving keys are decoded for a proof and released afterwards, so the resting footprint is the two
small decryption circuits plus Go heap headroom. Hardware guidance: 2 cores and 4 GB of RAM
minimum, 8 GB comfortable; `GOMEMLIMIT` tightens the peak at some CPU cost.

## Gas

Measured on the integration harness (Anvil, Cancun) at two policies; Groth16 verification gas does
not depend on circuit size, so these hold for any circuit build with the same public inputs.

| Call | n = 4, t = 3 | n = 32, t = 22 | Notes |
|---|---:|---:|---|
| `createEpoch` | 133,269 | 133,269 | 150,280 for the first seed |
| `claimSlot` (average) | 137,463 | 121,021 | |
| `submitContribution` | 500,070 | 1,314,499 | 16 keys; `32·(K·(2t+n)+5n)` bytes of transcript calldata |
| `finalizeEpoch` | 2,201,937 | 2,810,545 | verifier, 16 keys and 16 roots stored, 16 Merkle trees |
| `registerApplication` locked / automatic | 619,074 / 232,394 | 620,994 / 233,000 | locked verifies a Schnorr proof of possession and a subgroup check |
| `revealOrganizerSecret` | 225,902 | 223,750 | |
| `submitCiphertext` | 102,770 | 102,770 | no proof; on-curve and canonical checks only |
| `submitPartialDecryption` | 402,251 | 402,223 | verifier plus a 5-level Merkle path |
| `combineDecryption` | 410,614 | 483,955 | |

Observed on Sepolia: `submitContribution` 550,413 at `n = 6`, `finalizeEpoch` 2,170,777 to
2,180,237 at `n = 3`.

An epoch that yields all 16 keys costs about `0.15 + 4 × 0.14 + 4 × 0.50 + 2.20 ≈ 4.9 M` gas at
`n = 4` (0.31 M per key) and about `32 × (0.12 + 1.31) + 2.81 ≈ 48.9 M` at `n = 32` (3.1 M per
key). Decrypting one ciphertext costs one `submitCiphertext`, `t` partials and one combine:
about 1.7 M gas at `t = 3`.

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
