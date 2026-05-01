# DAVINCI DKG — Benchmarks

Reference figures for the production **MaxN = 32** build with the four
Groth16 circuits: Contribution, Finalize, PartialDecrypt, DecryptCombine.

Constraint counts come from `<circuit>.Compile()` + `ccs.GetNbConstraints()`.
Proving / verifying times are wall-clock per single proof, single CPU core
(`GOMAXPROCS=1`), measured on AMD Ryzen 7 7840U with 64 GiB RAM. Gas
figures are from `forge test --gas-report` against the local Anvil
backend running the production Solidity contracts with mock verifiers
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

| Circuit         | Constraints |
|-----------------|------------:|
| Contribution    |   1,430,854 |
| Finalize        |   1,021,651 |
| PartialDecrypt  |      20,717 |
| DecryptCombine  |      88,170 |
| **Total**       | **2,561,392** |

The dominant cost in Contribution and Finalize is `CommitmentPolynomialValue`,
which evaluates `Σ_k commitments[k] · x^k` for each recipient (or participant)
index `x ≤ MaxN`. The implementation specialises this to a 2-bit-windowed
left-to-right double-and-add over `api.ToBinary(scalar, nbBits)` with a
caller-supplied bit width, so the per-iteration scalar mul shrinks from a
~252-bit BabyJubJub field op (`scalarMulFakeGLV`) to roughly
`14 · (xMaxBits · k + 1)` constraints. The `k = 0` (`power = 1`) case is
special-cased to a direct point add. Range-check on the recipient /
participant index makes the bit-width bound sound — an oversized scalar
fails the proof rather than silently truncating.

---

## Proof generation time (single CPU core)

Wall-clock per single proof on one CPU core. The numbers include the gnark
constraint-system solver (witness solving) **and** the Groth16 prover, but
not the per-process pk/vk load (~hundreds of ms, amortised across many
proofs in the production node).

| Circuit         | Prove + verify (one shot) |
|-----------------|--------------------------:|
| Contribution    |                    7.99 s |
| Finalize        |                    4.26 s |
| PartialDecrypt  |                    240 ms |
| DecryptCombine  |                    602 ms |

Contribution and Finalize are O(N²) in the committee size; the constraints
are dominated by the `MaxN × MaxN` Feldman commitment loop. PartialDecrypt
and DecryptCombine are essentially constant in `MaxN`.

For a `n = 16` epoch with all nodes proving in parallel on separate hosts,
the proving critical path is one Contribution proof (~8 s) +
one Finalize proof (~4 s) ≈ **12 s of proving wall-clock**. With every
node serialised on one CPU core the total is `n × 8 s + 4 s ≈ 132 s`.

---

## Gas costs (MaxN = 32, Cancun fork)

Median values from `forge test --gas-report`. The min often reflects revert
paths and the max sometimes includes a cold-storage boundary; the median is
the production-relevant number.

### DKGManager — write paths

| Function                              |    Min |     Median |        Max |
|---------------------------------------|-------:|-----------:|-----------:|
| `createEpoch`                         |  24,734 |   226,458 |   246,544 |
| `claimSlot`                           |  28,831 |   159,970 |   207,824 |
| `submitContribution`                  |  72,438 |   238,286 |   238,286 |
| `finalizeEpoch`                       | 317,193 |   727,896 |   727,896 |
| `submitCiphertext`                    |  27,257 |    66,800 |    66,970 |
| `submitPartialDecryption`             |  34,819 |   159,339 |   176,427 |
| `combineDecryption`                   |  47,228 |   101,258 |   155,289 |
| `registerApplication` (mode 0)        |  24,846 |    29,424 |   176,103 |
| `registerApplicationCoDec` (mode 1)   |  27,063 |    28,172 |    29,282 |
| `extendRegistration`                  |  30,959 |    35,997 |    41,036 |
| `abortEpoch`                          |  26,069 |    28,222 |    30,376 |

### DKGRegistry — write paths

| Function           |    Min |       Median |          Max |
|--------------------|-------:|-------------:|-------------:|
| `registerKey`      | 23,516 |  **3,679,160** |  3,765,959 |
| `updateKey`        | 27,593 |  **3,574,420** |  3,579,126 |
| `markActive`       | 23,874 |       25,050 |     30,758 |
| `heartbeat`        | 23,484 |       25,748 |     28,013 |
| `reactivate`       | 23,719 |       34,642 |     34,642 |
| `reap`             | 23,868 |       33,894 |     33,894 |

`registerKey` and `updateKey` carry the on-chain Schnorr proof of knowledge
that the caller controls the secret behind the published BabyJubJub public
key (paper §5.1.1). The cost is dominated by the in-EVM Poseidon hashing
(`PoseidonT6` + `PoseidonT3`) and the BabyJubJub scalar multiplication that
verifies `z·G = A + c·PK`. Paid exactly once per node-key lifecycle event.

### Read paths

| Function                          |     Gas |
|-----------------------------------|--------:|
| `getEpoch`                        |  23,336 |
| `getApplication`                  |  17,141 |
| `getContribution`                 |  12,409 |
| `getPartialDecryption`            |  12,652 |
| `getCombinedDecryption`           |   5,712 |
| `getCiphertextHash`               |   2,814 |
| `getPlaintext`                    |   2,974 |
| `selectedParticipants` (n = 2)    |   8,026 |
| `getContributionVerifierVKeyHash` |   3,392 |
| `getPartialDecryptVerifierVKeyHash` | 3,876 |
| `nodeCount`                       |   2,366 |
| `activeCount`                     |   2,328 |
| `isActive`                        |   2,637 |

### Deployment

| Contract     | Bytecode size | Deploy gas |
|--------------|--------------:|-----------:|
| DKGManager   |        26,491 |  5,611,226 |
| DKGRegistry  |         5,111 |  1,115,813 |
| PoseidonT2   |         7,600 |  1,515,601 |
| PoseidonT3   |        29,345 |  5,870,224 |
| PoseidonT6   |       115,920 | 23,225,077 |

The `PoseidonTN` helper contracts are deployed once per chain by
`script/DeployAll.s.sol`. The DKGManager and DKGRegistry constructors take
their addresses, so deploying both Poseidon helpers + the two DKG contracts
adds up to roughly **37.3 M gas** end-to-end. The four Groth16 verifier
contracts (one per circuit) are deployed separately by the circuit-compile
pipeline; their bytecode is purely the verifying-key dump and the standard
gnark verifier scaffolding.

---

## Whole-epoch totals (MaxN = 32)

Using the medians above, end-to-end gas for one full DKG epoch followed by
one threshold decryption:

| Phase | n = 16 | n = 32 |
|---|---:|---:|
| `createEpoch`                                      |     226,458 |     226,458 |
| n × `claimSlot`                                    |   2,559,520 |   5,119,040 |
| n × `submitContribution`                           |   3,812,576 |   7,625,152 |
| 1 × `finalizeEpoch`                                |     727,896 |     727,896 |
| 1 × `submitCiphertext`                             |      66,800 |      66,800 |
| t × `submitPartialDecryption` (t = ⌈2n/3⌉)         |   1,752,729 (×11) |   3,505,458 (×22) |
| 1 × `combineDecryption`                            |     101,258 |     101,258 |
| **Round total**                                    | **9,247,237** | **17,372,062** |

These are *epoch-only* costs — application registration is paid once per
`(eid, aid)` pair (~29 k for mode 0, ~28 k for mode 1).

The big-ticket cost is node registration via `registerKey` (~3.7 M each),
amortised across every epoch the node participates in.

---

## Switching MaxN

Two-line edit, then one `make` command:

```go
// circuits/common/sizes.go
const MaxN = 16   // ← edit this
```

```solidity
// solidity/src/DKGManager.sol
uint256 internal constant MAX_N = 16;   // ← keep equal to circuits/common.MaxN
```

```bash
make circuits   # compile circuits → patch hashes → rebuild Solidity → regen Go bindings
```

`MaxN = 16` cuts the Contribution and Finalize circuits by roughly 4×
(they're O(N²)), which translates to ~4× faster proving and roughly ~30 %
lower `finalizeEpoch` gas. Per-call gas for `submitContribution` /
`submitPartialDecryption` / `combineDecryption` is essentially independent
of MaxN at the production gas-report level.

---

## How to reproduce

```bash
# 1. (Optional) Switch MaxN in circuits/common/sizes.go AND
#    solidity/src/DKGManager.sol::MAX_N. Both must match — the test
#    `TestSolidityMaxNMatchesGoMaxN` enforces this.

# 2. Refresh circuit constraint counts:
#    Compile each circuit and read ccs.GetNbConstraints() — see
#    circuits/contribution/compile.go for the entry point.

# 3. Refresh proof timings (single CPU): drop a one-line bench file
#    into each circuit package that calls the existing testAssignment +
#    Artifacts.LoadOrSetupForCircuit + ProveAndVerify path inside
#    a Benchmark function, then:
GOMAXPROCS=1 go test -count=1 -bench='^BenchmarkProve$' -benchtime=3x \
  -run='^$' -timeout 1800s \
  ./circuits/contribution/ ./circuits/finalize/ \
  ./circuits/partialdecrypt/ ./circuits/decryptcombine/

# 4. Refresh gas table:
cd solidity && forge test --gas-report --no-match-test '_Heavy|Stress'
```

The canonical inputs are `circuits/{contribution,finalize,partialdecrypt,decryptcombine}/`
and the Foundry suite. CI runs the gas report on every PR via the existing
test target.
