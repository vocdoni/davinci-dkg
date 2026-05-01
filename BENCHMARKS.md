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
| Contribution    |   1,430,852 |
| Finalize        |   1,021,651 |
| PartialDecrypt  |      20,717 |
| DecryptCombine  |      88,170 |
| **Total**       | **2,561,390** |

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

> **Contract layout.** `DKGManager` and `DKGAppManager` are deployed as a
> sibling pair to fit the EIP-170 24,576-byte contract-size limit. The
> per-application surface (`registerApplication`, `registerApplicationCoDec`,
> `submitOrganizerShare`, `getApplication`) lives on `DKGAppManager` and is
> wired one-shot via `DKGManager.setAppManager()`. The two contracts share
> the same epoch storage through the manager, but bill gas independently.

### DKGManager / DKGAppManager — write paths

| Function                              |    Min |     Median |        Max |
|---------------------------------------|-------:|-----------:|-----------:|
| `createEpoch`                         |  24,694 |   227,676 |   247,774 |
| `claimSlot`                           |  28,831 |   160,153 |   208,087 |
| `submitContribution`                  |  71,736 |   213,420 |   213,432 |
| `finalizeEpoch`                       | 317,228 |   747,825 |   747,825 |
| `submitCiphertext`                    |  27,293 | 2,055,970 | 2,056,140 |
| `submitPartialDecryption`             |  34,798 |    99,045 |   133,233 |
| `combineDecryption`                   |  47,217 |    92,844 |   138,472 |
| `registerApplication` (mode 0)        |  52,025 |    54,571 |   178,521 |
| `registerApplicationCoDec` (mode 1)   |  54,414 |    54,496 |    54,579 |
| `extendRegistration`                  |  31,037 |    36,116 |    41,195 |
| `abortEpoch`                          |  26,149 |    28,302 |    30,456 |

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

`submitCiphertext` runs a full prime-subgroup membership check on both
ciphertext points (`isInPrimeSubgroup` = `[L]·P == identity`). This costs
~1 M gas per point and ~2 M gas total — one full BabyJubJub scalar
multiplication each, dominated by `bigModExp` calls inside `pointAdd`.
Without this check a small-order point would let a permitted submitter
park a ciphertext slot the combine circuit could never accept.

### Read paths

| Function                          |     Gas |
|-----------------------------------|--------:|
| `getEpoch`                        |  23,751 |
| `getApplication`                  |  17,250 |
| `getContribution`                 |  12,366 |
| `getPartialDecryption`            |   6,473 |
| `getCombinedDecryption`           |   5,633 |
| `getCiphertextHash`               |   2,813 |
| `getPlaintext`                    |   2,917 |
| `getCollectivePublicKey`          |   5,022 |
| `selectedParticipants` (n = 2)    |   7,985 |
| `getContributionVerifierVKeyHash` |   3,395 |
| `getPartialDecryptVerifierVKeyHash` | 3,791 |
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

## Whole-epoch totals (MaxN = 32)

Using the medians above, end-to-end gas for one full DKG epoch followed by
one threshold decryption:

| Phase | n = 16 | n = 32 |
|---|---:|---:|
| `createEpoch`                                      |     227,676 |     227,676 |
| n × `claimSlot`                                    |   2,562,448 |   5,124,896 |
| n × `submitContribution`                           |   3,414,720 |   6,829,440 |
| 1 × `finalizeEpoch`                                |     747,825 |     747,825 |
| 1 × `submitCiphertext`                             |   2,055,970 |   2,055,970 |
| t × `submitPartialDecryption` (t = ⌈2n/3⌉)         |   1,089,495 (×11) |   2,178,990 (×22) |
| 1 × `combineDecryption`                            |      92,844 |      92,844 |
| **Round total**                                    | **10,190,978** | **17,257,641** |

These are *epoch-only* costs — application registration is paid once per
`(eid, aid)` pair on `DKGAppManager` (~55 k for mode 0, ~55 k for mode 1).

The big-ticket cost is node registration via `registerKey` (~1.27 M after
the keccak swap), amortised across every epoch the node participates in.

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
