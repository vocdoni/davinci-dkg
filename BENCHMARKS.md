# DAVINCI DKG — Benchmarks

Two compile-time builds are compared:

- **MaxN = 32** — committees up to 32 participants per round.
- **MaxN = 16** — committees up to 16 participants per round (current default).

Tested on AMD Ryzen 7 7840U with 64 GiB RAM.

---

## Circuit Constraint Counts

| Circuit | MaxN = 32 | MaxN = 16 | Δ | Scaling |
|---|---|---|---|---|
| Contribution    | 2,835,328 | 802,476 | **−71.7%** | O(N²) |
| Finalize        | 2,490,861 | 625,522 | **−74.9%** | O(N²) |
| DecryptCombine  |    84,314 |  43,635 | −48.3%     | O(N)  |
| RevealShare     |     3,342 |   1,904 | −43.0%     | O(N)  |
| PartialDecrypt  |    20,361 |  20,361 | —          | O(1)  |
| RevealSubmit    |     2,346 |   2,346 | —          | O(1)  |
| **Total**       | **5,436,552** | **1,496,244** | **−72.5%** | |

---

## Proof Generation Time

Wall-clock per proof on a single CPU (`gnark` Groth16 BN254, no GPU). Per-node
proving time is constant in actual `n` for a given build — the circuit pays its
full compile-time size regardless of how many slots are used.

| Circuit | MaxN = 32 | MaxN = 16 |
|---|---|---|
| Contribution (per node) | ~3.4 s  | ~0.9 s  |
| Finalize                | ~1.5 s  | ~0.4 s  |
| DecryptCombine          | ~100 ms | ~70 ms  |
| RevealShare             | ~10 ms  | ~5 ms   |
| PartialDecrypt          | ~60 ms  | ~60 ms  |
| RevealSubmit            | ~30 ms  | ~30 ms  |

For a full house at n=16: total proving wall-clock drops from ~109 s
(32 × 3.4 s + 1.5 s) at MaxN=32 to ~14 s (16 × 0.9 s + 0.4 s) at MaxN=16 —
**~8× faster end-to-end** when the committee fits.

---

## Gas Costs by Committee Size

Measured on local Anvil (block gas limit 30M, 2-second blocks, EIP-1559) with
the lottery-based committee selection. Each row is one complete DKG +
threshold decryption run.

### MaxN = 16

| n | t | submitContribution | finalizeRound | submitPartialDecryption | combineDecryption |
|---|---|---|---|---|---|
| 4  | 3  | 454,701 |   610,016 | 427,826 | 358,595 |
| 8  | 6  | 464,469 |   774,651 | 427,802 | 395,672 |
| 12 | 8  | 473,625 |   947,599 | 427,826 | 420,426 |
| 16 | 11 | 483,525 | 1,145,456 | 427,814 | 457,539 |

### MaxN = 32

| n | t | submitContribution | finalizeRound | submitPartialDecryption | combineDecryption |
|---|---|---|---|---|---|
| 4  | 3  | 491,224 | 1,062,448 | 427,802 | 374,675 |
| 8  | 6  | 501,028 | 1,228,943 | 427,802 | 411,788 |
| 12 | 8  | 510,148 | 1,404,019 | 427,802 | 436,518 |
| 16 | 11 | 520,000 | 1,604,080 | 427,838 | 473,607 |
| 20 | 14 | 529,852 | 1,822,634 | 427,814 | 510,756 |
| 24 | 16 | 539,008 | 2,040,793 | 427,802 | 535,402 |
| 28 | 19 | 548,836 | 2,292,901 | 427,838 | 572,575 |
| 32 | 22 | 558,676 | 2,563,382 | 427,802 | 609,652 |

### Side-by-side at the sizes both builds support (n = 4, 8, 12, 16)

| n  | Call | MaxN=32 | MaxN=16 | Δ |
|---|---|---|---|---|
| 4  | submitContribution | 491,224   |   454,701 | −7.4%      |
| 4  | finalizeRound      | 1,062,448 |   610,016 | **−42.6%** |
| 4  | combineDecryption  | 374,675   |   358,595 | −4.3%      |
| 8  | submitContribution | 501,028   |   464,469 | −7.3%      |
| 8  | finalizeRound      | 1,228,943 |   774,651 | **−37.0%** |
| 8  | combineDecryption  | 411,788   |   395,672 | −4.1%      |
| 12 | submitContribution | 510,148   |   473,625 | −7.2%      |
| 12 | finalizeRound      | 1,404,019 |   947,599 | **−32.5%** |
| 12 | combineDecryption  | 436,518   |   420,426 | −3.7%      |
| 16 | submitContribution | 520,000   |   483,525 | **−7.0%**  |
| 16 | finalizeRound      | 1,604,080 | 1,145,456 | **−28.6%** |
| 16 | combineDecryption  | 473,607   |   457,539 | −3.4%      |

`submitPartialDecryption` is essentially identical between builds (~425–427 k);
its dominant cost is the Groth16 verifier base which doesn't scale with N.

### Setup overhead (createRound + n × claimSlot)

`createRound` ≈ 144–196 k in both builds (varies by round nonce / storage state).
`claimSlot` ≈ 122–137 k average per node (first claimer pays the seed-resolve,
last claimer pays the committee snapshot). Setup overhead scales linearly with `n`
and is essentially identical between the two builds — none of it depends on MaxN.

### Whole-round totals at n = 16

Using measured gas figures (MaxN=16: n=16 row; MaxN=32: n=16 row):

| Phase | MaxN = 32 | MaxN = 16 | Saved |
|---|---|---|---|
| Setup (create + 16× claim)  |  2,105,263 |  2,096,647 |       8,616 |
| 16× submitContribution      |  8,320,000 |  7,736,400 |     583,600 |
| 1×  finalizeRound           |  1,604,080 |  1,145,456 |     458,624 |
| 11× submitPartialDecryption |  4,706,218 |  4,705,954 |        −264 |
| 1×  combineDecryption       |    473,607 |    457,539 |      16,068 |
| **Round total at n=16**     | **17,209,168** | **16,141,996** | **1,067,172 (−6.2%)** |

### Whole-round totals at n = 32 (MaxN = 32 only)

| Phase | Gas |
|---|---|
| Setup (create + 32× claim)  |  3,976,227 |
| 32× submitContribution      | 17,877,632 |
| 1×  finalizeRound           |  2,563,382 |
| 22× submitPartialDecryption |  9,411,644 |
| 1×  combineDecryption       |    609,652 |
| **Round total at n=32**     | **34,438,537** |

The largest absolute saving comes from `finalizeRound` (−28.6%), driven by
the smaller calldata transcript and the smaller per-contributor digest input.
The contribution savings are smaller than before because the BabyJubJub point
accumulation overhead now dominates the calldata savings at small n.

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

---

## Gas Cost Model (MaxN-aware)

Gas per call breaks down into N-independent and N-dependent parts:

1. **Groth16 BN254 pairing verification** (~207 k base + ~6.65 k per public
   input). Independent of N; this is the floor for every proof-gated call.
2. **Calldata-direct transcript verification** (~30 gas per word). Linear in
   transcript word count: `O(N²)` for `finalizeRound`, `O(N)` for the others.
3. **Cold SSTOREs** (22.1 k each). Linear in actual `n`, independent of MaxN.
   Mostly: 1 per `submitContribution`, `n` at finalize for share commitments,
   2-3 per partial-decryption / reveal call.
4. **Lottery setup**: `createRound` ≈ 194 k + `n × claimSlot` ≈ 118 k each,
   distributed across nodes. Independent of MaxN.
5. **Per-contributor digest in finalize**: `keccak256` over each contributor's
   `2N`-word slice in calldata, ~`30 + 6×(2N)` gas per iteration. Linear in MaxN.

Halving MaxN cuts items 2 and 5 by roughly the same factor, which is why
`finalizeRound` shrinks the most when you switch to MaxN=16.

---

## How to Reproduce

```bash
# 1. Set MaxN in circuits/common/sizes.go AND solidity/src/DKGManager.sol::MAX_N
# 2. Recompile circuits + Solidity + Go bindings
make circuits

# 3. Build Docker images
docker compose -f testnet/docker-compose.yml build deployer dkg-node

# 4. Start chain + nodes, run scenario
docker compose -f testnet/docker-compose.yml up -d anvil deployer
DKG_NODE_COUNT=16 DKG_THRESHOLD=11 \
  docker compose -f testnet/docker-compose.yml up -d --scale dkg-node=16 dkg-node
DKG_NODE_COUNT=16 DKG_THRESHOLD=11 \
  docker compose -f testnet/docker-compose.yml --profile runner run --rm dkg-runner
```
