# DAVINCI DKG — Benchmarks

Reference figures for the production **MaxN = 32** build, with side-by-side
comparison columns for **MaxN = 16** and **MaxN = 48** so operators can size
proving infrastructure for alternative committee bounds. The four Groth16
circuits are: Contribution, Finalize, PartialDecrypt, DecryptCombine.

Constraint counts come from `<circuit>.Compile()` + `ccs.GetNbConstraints()`.
Proving / verifying times are wall-clock per single proof on the full host
CPU (gnark parallelises internally over all logical cores), measured on
**Intel Core i9-14900K (32 logical cores) with 64 GiB RAM**. Gas figures
are from `forge test --gas-report` against the local Anvil backend running
the production Solidity contracts with mock verifiers
(`MockContributionVerifier` etc.) — proof verification itself adds a
constant ~280 k of pairing-check gas to the Groth16-bound write paths,
which the production deployment includes.

> **Caveat — single-party local trusted setup.** All proving / verifying
> figures use a single-party Groth16 setup. A multi-party trusted-setup
> ceremony will produce fresh pk/vk values; the *constraint counts and
> call-shape gas costs* are unchanged by the ceremony, but the verifier
> contract bytecode may shift by a few hundred bytes worth of vkey
> constants.

---

## Circuit constraint counts

| Circuit         |  MaxN = 16  |   MaxN = 32 |   MaxN = 48 |
|-----------------|------------:|------------:|------------:|
| Contribution    |   251,792 |     536,262 |   859,076 |
| Finalize        |     65,343 |     211,600 |     444,799 |
| PartialDecrypt  |      22,061 |      22,061 |      22,061 |
| DecryptCombine  |     104,328 |     238,207 |     409,969 |
| **Total**       | **443,524** | **1,008,130** | **1,735,905** |

The dominant cost in Contribution and Finalize is `CommitmentPolynomialValue`,
which evaluates `Σ_k commitments[k] · x^k` for each recipient (or participant)
index `x ≤ MaxN`. It uses Horner's rule, `acc ← x·acc + C_k` from the top
coefficient down, so every step is one 6-bit `ScalarMulSmallScalar` plus one
point addition (≈ 75 constraints) regardless of `k`; the previous form
multiplied each commitment by a growing power `x^k` and cost ~4× more.
`ScalarMulSmallScalar` range-checks `x` via `api.ToBinary`, so an oversized
index fails the proof rather than wrapping.

Since the Fiat–Shamir hardening (challenge over calldata + digests), Finalize
also digests the contribution rows in circuit (one Poseidon per row, one over
the row digests, ~40 k constraints) and Contribution absorbs the recipient
keys into its share digest.

The empirical scaling roughly tracks O(N²) for Contribution and Finalize
(constraint count grows ~4× from N=16→32 and ~2.8× from N=32→48), is flat
in N for PartialDecrypt (per-share, not per-committee), and roughly linear
in N for DecryptCombine (qualifying-set Lagrange plus the `MaxN × t` short
scalar multiplications of the Vandermonde check that pins the Lagrange
vector; that check is ~60% of the combine circuit).

---

## Proof generation time (full host, all cores)

Wall-clock per single proof, with gnark parallelising internally over all
32 logical cores of the i9-14900K. The numbers include the gnark
constraint-system solver (witness solving) **and** the Groth16 prover, but
not the per-process pk/vk load (~hundreds of ms, amortised across many
proofs in the production node). Each cell is `time` from
`go test -bench='^BenchmarkProve$' -benchtime=3x`.

| Circuit         |  MaxN = 16 |  MaxN = 32 |  MaxN = 48 |
|-----------------|-----------:|-----------:|-----------:|
| Contribution    |     145 ms |     441 ms |     PT48C |
| Finalize        |     66 ms |     193 ms |     364 ms |
| PartialDecrypt  |     26 ms |      27 ms |     PT48P |
| DecryptCombine  |      70 ms |     133 ms |     226 ms |

Contribution and Finalize are the only N-scaling circuits at the proving
level, mirroring the constraint scaling above. PartialDecrypt is per-share
work and is essentially constant; DecryptCombine grows linearly in N.

For a `n = 16` epoch on hardware comparable to the i9-14900K with all
nodes proving in parallel on separate hosts, the proving critical path is
one Contribution proof (~1.0 s) + one Finalize proof (~0.6 s) ≈ **1.6 s
of wall-clock proving** at MaxN = 32. Serialised on a single host the
total is `n × 1.0 s + 0.6 s ≈ 16.6 s`. At MaxN = 48 the same critical
path is ~4.0 s parallelised, ~32.4 s serialised.

---

## Gas costs (Cancun fork)

Median values from `forge test --gas-report`. The min often reflects revert
paths and the max sometimes includes a cold-storage boundary; the median is
the production-relevant number.

> **Contract layout.** `DKGManager` and `DKGAppManager` are deployed as a
> sibling pair to fit the EIP-170 24,576-byte contract-size limit. The
> per-application surface (`registerApplication`, `registerApplicationCoDec`,
> `submitOrganizerShare`, `getApplication`) lives on `DKGAppManager` and is
> wired one-shot via `DKGManager.setAppManager()`. The two contracts share
> the same epoch storage through the manager, but bill gas independently.

### DKGManager / DKGAppManager — write paths (median gas, three N values)

| Function                              | MaxN = 16 | MaxN = 32 | MaxN = 48 |
|---------------------------------------|----------:|----------:|----------:|
| `createEpoch`                         |   223,800 |   223,800 |   223,800 |
| `claimSlot`                           |   154,083 |   159,184 |   159,184 |
| `submitContribution`                  |   175,605 |   212,116 |   248,679 |
| `finalizeEpoch`                       |   303,569 |   745,957 | 1,467,876 |
| `submitCiphertext`                    |    65,905 |    65,905 |    65,905 |
| `submitPartialDecryption`             |    97,939 |    97,939 |    97,939 |
| `combineDecryption`                   |    80,656 |    91,960 |   103,264 |
| `registerApplication` (mode 0)        |    53,481 |    53,481 |    53,481 |
| `registerApplicationCoDec` (mode 1)   |    53,406 |    53,406 |    53,406 |
| `abortEpoch`                          |    28,115 |    28,115 |    28,115 |

The N-dependent write paths are:

* `finalizeEpoch` — O(N²) transcript verification, the dominant scaling
  cost on chain. ~5× from N=16 to N=48.
* `submitContribution` — linear in N via per-recipient calldata + storage.
  ~1.4× from N=16 to N=48.
* `combineDecryption` — linear in N via the Lagrange qualifying-set loop.
  ~1.3× from N=16 to N=48.
* `claimSlot` — marginally N-dependent (~3% spread N=16 vs N=32).

Everything else (registry, ciphertext submission, partial decryption,
application registration, lifecycle controls) is constant in N at the EVM
level. Min/max columns match closely across all three N values; we report
medians here for compactness — the per-N raw output from
`forge test --gas-report` is regenerable in seconds via the procedure in
the "How to reproduce" section below.

**Methodology.** A single file-level constant `MAX_N` lives in
`solidity/src/libraries/Sizes.sol` and is imported by both `DKGManager.sol`
and `test/TestHelpers.t.sol`, so switching N for a gas sweep is a
one-line edit + `forge test --gas-report`. All 114 Foundry tests pass at
each of N = 16, 32, 48.

`registerApplication`, `registerApplicationCoDec` (and the read-side
`getApplication`) live on `DKGAppManager`; everything else above is
`DKGManager`. Both contracts share the same Foundry gas-report run.

### DKGRegistry — write paths

| Function           |    Min |       Median |          Max |
|--------------------|-------:|-------------:|-------------:|
| `registerKey`      | 23,513 |  **1,269,784** |  1,278,511 |
| `updateKey`        | 27,693 |  **1,165,147** |  1,180,691 |
| `markActive`       | 23,874 |       25,063 |     30,959 |
| `heartbeat`        | 23,484 |       25,772 |     28,061 |
| `reactivate`       | 23,719 |       35,036 |     35,036 |
| `reap`             | 23,895 |       34,053 |     34,053 |

`registerKey` and `updateKey` carry the on-chain Schnorr proof of knowledge
that the caller controls the secret behind the published BabyJubJub public
key (paper §5.1.1). The Fiat-Shamir challenge is now `c = keccak256(domain ‖
… ‖ PK ‖ A) mod L` — the legacy `PoseidonT5` / `PoseidonT6` helper
contracts (which exceeded EIP-170 anyway) are no longer deployed. After the
keccak swap the dominant cost is the in-EVM BabyJubJub double scalar
multiplication that verifies `z·G == A + c·PK`. The verifier uses a width-2
windowed Strauss–Shamir double-and-add over the basis `(G, -PK)` with
multiples of `G` precomputed as Solidity constants, so the per-window cost
is two doublings plus at most one conditional add against a 16-entry
`i·G + j·(-PK)` lookup table. Twisted-Edwards `pointAdd` uses the
single-inverse trick: both denominators are inverted via one `bigModExp`
precompile call instead of two. Paid exactly once per node-key lifecycle
event.

`submitCiphertext` enforces only the cheap membership checks on each
ciphertext point: canonical encoding (X, Y < Q), on-curve, and non-identity.
The expensive prime-order subgroup test (`[L]·P == identity`, ~1 M gas per
point, ~2 M gas total) is intentionally **not** performed on chain — that
check lives in the off-chain DKG-node software
(`crypto/group/validation.go::ValidateCiphertext`), which refuses to compute
a partial decryption for any toxic ciphertext. See SECURITY.md §O-1 for the
threat model. Skipping the on-chain test cuts `submitCiphertext` from
~2.06 M gas to ~66 k.

### Read paths

| Function                          |     Gas |
|-----------------------------------|--------:|
| `getEpoch`                        |  22,661 |
| `getApplication`                  |  17,250 |
| `getContribution`                 |  12,233 |
| `getPartialDecryption`            |   6,126 |
| `getCombinedDecryption`           |   5,542 |
| `getCiphertextHash`               |   2,766 |
| `getPlaintext`                    |   2,882 |
| `getCollectivePublicKey`          |   5,034 |
| `selectedParticipants` (n = 2)    |   7,989 |
| `getContributionVerifierVKeyHash` |   3,368 |
| `nodeCount`                       |   2,372 |
| `activeCount`                     |   2,334 |
| `isActive`                        |   2,637 |

### Deployment

| Contract       | Runtime bytecode | Deploy gas |
|----------------|-----------------:|-----------:|
| DKGManager     |           20,818 |  4,560,386 |
| DKGAppManager  |            8,216 |  1,832,236 |
| DKGRegistry    |            4,931 |  1,120,982 |

The three production contracts are deployed by `script/DeployAll.s.sol` in
the order `DKGRegistry → DKGManager → DKGAppManager`, followed by the
one-shot wiring calls `DKGRegistry.setManager(...)` and
`DKGManager.setAppManager(...)`. Total core deployment cost is roughly
**7.5 M gas**. The `DKGManager` runtime is 20,818 bytes, leaving a
3,758-byte margin against the 24,576-byte EIP-170 limit; splitting the
per-application surface into `DKGAppManager` is what created that margin.
Because the on-chain Schnorr challenge moved from Poseidon to keccak256,
the legacy `PoseidonT2 / PoseidonT3 / PoseidonT5 / PoseidonT6` helper
contracts are no longer required for the production deploy. The four
Groth16 verifier contracts (one per circuit) are deployed separately by
the circuit-compile pipeline; their bytecode is purely the verifying-key
dump and the standard gnark verifier scaffolding.

---

## Whole-epoch totals

Using the medians above, end-to-end gas for one full DKG epoch followed by
one threshold decryption, evaluated at the natural pairing of (committee
size n, MaxN bound):

Each row uses the gas figures from the matching MaxN column above (so the
finalize/contribution costs reflect the actual on-chain transcript scaling
at that bound).

| Phase | n=16 (MaxN=16) | n=32 (MaxN=32) | n=48 (MaxN=48) |
|---|---:|---:|---:|
| `createEpoch`                                  |     223,800 |     223,800 |      223,800 |
| n × `claimSlot`                                |   2,465,328 |   5,093,888 |    7,640,832 |
| n × `submitContribution`                       |   2,809,680 |   6,787,712 |   11,936,592 |
| 1 × `finalizeEpoch`                            |     303,569 |     745,957 |    1,467,876 |
| 1 × `submitCiphertext`                         |      65,905 |      65,905 |       65,905 |
| t × `submitPartialDecryption` (t = ⌈2n/3⌉)     |   1,077,329 (×11) | 2,154,658 (×22) | 3,231,987 (×33) |
| 1 × `combineDecryption`                        |      80,656 |      91,960 |      103,264 |
| **Epoch total**                                | **7,026,267** | **15,163,880** | **24,670,256** |

These are *epoch-only* costs — application registration is paid once per
`(eid, aid)` pair on `DKGAppManager` (~55 k for mode 0, ~55 k for mode 1).

The big-ticket cost is node registration via `registerKey` (~1.27 M after
the keccak swap), amortised across every epoch the node participates in.

---

## Switching MaxN

Two-line edit, then one `make` command:

```go
// circuits/common/sizes.go
const MaxN = 16   // ← edit this (Go side)
```

```solidity
// solidity/src/libraries/Sizes.sol
uint256 constant MAX_N = 16;   // ← edit this (Solidity side, must match)
```

```bash
make circuits   # compile circuits → patch hashes → rebuild Solidity → regen Go bindings
```

The single `Sizes.sol` constant is imported by both the production
`DKGManager` contract and the Foundry test helpers
(`test/TestHelpers.t.sol`), so a `MaxN` change does not require any other
Solidity edits — `forge test --gas-report` will produce a complete gas
table for the new N immediately.

Empirically (see the constraint-count and proving-time tables above):

* MaxN = 16 cuts Contribution + Finalize constraints by ~4× vs MaxN = 32
  (~5.7× total work), and proving wall-clock by ~3.5×.
* MaxN = 48 inflates Contribution + Finalize constraints by ~3× vs
  MaxN = 32, and proving wall-clock by ~2× for Contribution and ~3.2× for
  Finalize.
* PartialDecrypt is fully independent of MaxN.
* DecryptCombine grows roughly linearly in MaxN.
* On-chain, only `submitContribution` and `finalizeEpoch` are meaningfully
  N-scaling; the rest of the write paths (registry, partial-decrypt,
  combine, ciphertext submission) are constant-gas.

---

## How to reproduce

```bash
# 1. (Optional) Switch MaxN in circuits/common/sizes.go AND
#    solidity/src/DKGManager.sol::MAX_N. Both must match — the test
#    `TestSolidityMaxNMatchesGoMaxN` enforces this.

# 2. Refresh circuit constraint counts:
go run ./cmd/constraints/

# 3. Refresh proof timings (full host, all cores). Each circuit package has
#    a one-line BenchmarkProve in bench_test.go that wraps testAssignment +
#    Artifacts.LoadOrSetupForCircuit + ProveAndVerify. Do NOT pin
#    GOMAXPROCS — gnark parallelises internally and the production node
#    proves with all cores available.
go test -count=1 -bench='^BenchmarkProve$' -benchtime=3x \
  -run='^$' -timeout 1800s \
  ./circuits/contribution/ ./circuits/finalize/ \
  ./circuits/partialdecrypt/ ./circuits/decryptcombine/

# 4. Refresh gas table:
cd solidity && forge test --gas-report --no-match-test '_Heavy|Stress'
```

The canonical inputs are `circuits/{contribution,finalize,partialdecrypt,decryptcombine}/`
and the Foundry suite. CI runs the gas report on every PR via the existing
test target.
