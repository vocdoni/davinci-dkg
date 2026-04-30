# DAVINCI DKG — Benchmarks

Measured on AMD Ryzen 7 7840U with 64 GiB RAM. Single CPU, no GPU, no
icicle/CUDA acceleration. The numbers below are for the production
**MaxN = 32** build with the post-rewrite circuit set (Contribution,
Finalize, PartialDecrypt, DecryptCombine — no disclosure circuits).

Last refresh: 2026-04-30 (after the per-Application rewrite + DEEPSEEK
hardening). Re-measure with `go run /tmp/circuit_stats.go` and `forge
test --gas-report` after any circuit / contract change.

> **Caveat — dev setup, not production.** All proving / verification
> figures use a single-party local Groth16 setup. The S2 multi-party
> trusted-setup ceremony will produce fresh pk/vk values; the
> *constraint counts and call-shape gas costs* are unchanged by the
> ceremony, but the verifier deployment cost may shift by a few hundred
> bytes worth of vkey constants.

---

## Circuit Constraint Counts

| Circuit         | Constraints | Public wires | Secret wires | Scaling |
|-----------------|------------:|-------------:|-------------:|---------|
| Contribution    |   2,900,192 |    2,827,483 |          448 | O(N²)   |
| Finalize        |   2,490,861 |    2,489,577 |        2,208 | O(N²)   |
| PartialDecrypt  |      20,715 |       19,471 |            2 | O(1)    |
| DecryptCombine  |      86,890 |       86,407 |          133 | O(N)    |
| **Total**       | **5,498,648** |          —   |            — |         |

Notable since the last benchmark:

- **Contribution** +2.3 %: in-circuit range checks on coefficients and
  per-recipient shares (paper §5.1, fix for the C-1 unbounded-share
  vulnerability).
- **PartialDecrypt** +1.7 %: three new public inputs (`aid`, `ctIdx`,
  `role`) bound into the Fiat-Shamir transcript for cross-application
  / cross-epoch replay defence (paper §4.4 / DEEPSEEK §1.3).
- **DecryptCombine** +3.1 %: mode-bit booleanity check, mode-conditional
  `Select` for the per-application correction term `T = mode==0 ? S·C₁ :
  Δ_org`, plus the per-app fields in the transcript (`aid`, `ctIdx`,
  `mode`, `S`, `DeltaOrg`).
- **Total** is down vs. the pre-rewrite total because the
  `RevealShare` (~3,342) and `RevealSubmit` (~2,346) circuits were
  removed in P1 (paper §3 — disclosure path retired).

---

## Proof Generation Time

Wall-clock per single proof on one CPU core. The numbers include the
gnark constraint-system solver (witness solving) **and** the Groth16
prover, but not the per-process pk/vk load (~hundreds of ms, amortised
across many proofs in the production node).

| Circuit         | Prove + verify (one shot) |
|-----------------|--------------------------:|
| Contribution    |                    1.91 s |
| Finalize        |                    1.88 s |
| PartialDecrypt  |                      69 ms |
| DecryptCombine  |                     137 ms |

A full DKG epoch at n = 16 takes roughly **16 × 1.9 s ≈ 30 s** of
contribution-circuit proving (one proof per node, embarrassingly
parallel) + 1.9 s of finalize-circuit proving (single party, the
caller of `finalizeEpoch`) ≈ 32 s of proving wall-clock if all nodes
run on one CPU core. With nodes on separate hosts, the critical path
is one contribution proof (1.9 s) + one finalize proof (1.9 s) ≈ 4 s.

---

## Gas Costs (MaxN = 32, Cancun fork)

Measured via `forge test --gas-report` on local Anvil. Min / median / max
are across the entire test suite — the **median** is the production-relevant
number since the min often reflects revert paths and the max sometimes
includes a cold-storage boundary.

### DKGManager — write paths

| Function                      |    Min |     Median |        Max |
|-------------------------------|-------:|-----------:|-----------:|
| `createEpoch`                 |   24,712 |   226,436 |   246,522 |
| `claimSlot`                   |   28,809 |   159,948 |   207,802 |
| `submitContribution`          |   72,416 |   238,264 |   238,264 |
| `finalizeEpoch`               |  317,156 |   723,200 |   723,200 |
| `submitCiphertext`            |   27,538 |    37,734 |    66,790 |
| `submitPartialDecryption`     |   31,828 |   161,565 |   170,109 |
| `combineDecryption`           |   47,172 |    49,490 |   154,011 |
| `registerApplication` (mode 0) |  24,802 |    29,380 |   131,421 |
| `registerApplicationCoDec` (mode 1) | 27,063 | 28,172 | 29,282 |
| `extendRegistration`          |   30,959 |    35,997 |    41,036 |
| `abortEpoch`                  |   26,069 |    28,222 |    30,376 |

### DKGRegistry — write paths

| Function           |    Min |       Median |          Max |
|--------------------|-------:|-------------:|-------------:|
| `registerKey`      | 23,516 |  **3,679,160** | 3,765,959 |
| `updateKey`        | 27,593 |  **3,574,420** | 3,579,126 |
| `markActive`       | 23,874 |     25,050    |    30,758 |
| `heartbeat`        | 23,484 |     25,748    |    28,013 |
| `reactivate`       | 23,719 |     34,642    |    34,642 |
| `reap`             | 23,868 |     33,894    |    33,894 |

The big number on `registerKey` / `updateKey` is **the Schnorr proof
of knowledge** added in P4 (paper §6.1, addressing the C-2 unverified-DLP
issue). The cost is dominated by the in-EVM Poseidon hashing
(`PoseidonT6` + `PoseidonT3`) and the BabyJubJub scalar multiplication
required to verify `z·G = A + c·PK`. Compared to the pre-P4 path
(~50 k for `registerKey`), this is **~70× more expensive** but is paid
exactly once per node-key lifecycle event.

### Read paths (unchanged from pre-rewrite, included for completeness)

| Function                          |  Gas |
|-----------------------------------|-----:|
| `getEpoch`                        | 23,292 |
| `getApplication`                  | 17,097 |
| `getContribution`                 | 12,387 |
| `getPartialDecryption`            | 12,582 |
| `getCombinedDecryption`           |  5,831 |
| `getCiphertextHash`               |  2,760 |
| `getPlaintext`                    |  2,898 |
| `selectedParticipants` (n = 4)    |  8,004 |
| `getContributionVerifierVKeyHash` |  3,392 |
| `getPartialDecryptVerifierVKeyHash` | 3,832 |

### Deployment

| Contract     | Bytecode size | Deploy gas |
|--------------|--------------:|-----------:|
| DKGManager   |        25,143 |  5,319,675 |
| DKGRegistry  |         5,111 |  1,115,813 |

(The `PoseidonTN` precompile-style helper contracts are deployed
separately by `script/DeployAll.s.sol` and account for an additional
~6.5 M gas in total.)

---

## Whole-epoch totals (MaxN = 32)

Using the median values above, end-to-end gas for one DKG epoch +
one threshold decryption:

| Phase | n = 16 | n = 32 |
|---|---:|---:|
| `createEpoch`                |     226,436 |     226,436 |
| n × `claimSlot` (≈160 k each) |   2,560,000 |   5,120,000 |
| n × `submitContribution`     |   3,812,224 |   7,624,448 |
| 1 × `finalizeEpoch`          |     723,200 |     723,200 |
| 1 × `submitCiphertext`       |      37,734 |      37,734 |
| t × `submitPartialDecryption` (median t = ⌈2n/3⌉) | ~1,777,215 (×11) | ~3,554,430 (×22) |
| 1 × `combineDecryption`      |      49,490 |      49,490 |
| **Round total**              | **9,186,299** | **17,335,738** |

These are *epoch-only* costs — application registration is paid once
per `(eid, aid)` pair when the application is added (~29 k for mode 0,
~28 k for mode 1).

The big-ticket cost remains node registration via `registerKey` (~3.7 M
each), which is amortised across every epoch the node participates in.

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

`MaxN = 16` cuts the Contribution and Finalize circuits by roughly
4× (they're O(N²)), which translates to ~4× faster proving and
roughly ~30 % lower `finalizeEpoch` gas. Per-call gas for
`submitContribution` / `submitPartialDecryption` /
`combineDecryption` is essentially independent of MaxN at the
production gas-report level.

---

## How to Reproduce

```bash
# 1. (Optional) Switch MaxN in circuits/common/sizes.go AND
#    solidity/src/DKGManager.sol::MAX_N. Both must match — the test
#    `TestSolidityMaxNMatchesGoMaxN` enforces this.

# 2. Refresh circuit constraint counts:
go run /tmp/circuit_stats.go    # if you kept the helper, otherwise:
# Compile each circuit and read ccs.GetNbConstraints() — see
# circuits/contribution/compile.go for the entry point.

# 3. Refresh proof timings (single CPU):
go run /tmp/circuit_bench.go    # see git history for the helper

# 4. Refresh gas table:
cd solidity && forge test --gas-report --no-match-test '_Heavy|Stress'
```

The bench / stats helpers are throwaway scripts; the canonical inputs
are `circuits/{contribution,finalize,partialdecrypt,decryptcombine}/`
and the Foundry suite. CI runs the gas report on every PR via the
existing test target.
