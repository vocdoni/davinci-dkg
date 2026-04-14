# DAVINCI DKG — Benchmarks

Two compile-time builds are compared:

- **MaxN = 32** — committees up to 32 participants per round.
- **MaxN = 16** — committees up to 16 participants per round (current default).

Both share identical contract logic, lottery selection, and circuit
optimizations. The only difference is the single source-of-truth constant
`circuits/common.MaxN` (mirrored in `solidity/src/DKGManager.sol::MAX_N`).
Threshold per run is `t = ⌈2n/3⌉`.

---

## Circuit Constraint Counts

| Circuit | MaxN = 32 | MaxN = 16 | Δ | Scaling |
|---|---|---|---|---|
| Contribution    | 2,836,382 | 802,476 | **−71.7%** | O(N²) |
| Finalize        | 2,490,861 | 625,522 | **−74.9%** | O(N²) |
| DecryptCombine  |    84,314 |  43,635 | −48.3%     | O(N)  |
| RevealShare     |     3,342 |   1,904 | −43.0%     | O(N)  |
| PartialDecrypt  |    20,361 |  20,361 | —          | O(1)  |
| RevealSubmit    |     2,346 |   2,346 | —          | O(1)  |
| **Total**       | **5,437,606** | **1,496,244** | **−72.5%** | |

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
| 4  | 3  | 407,460 |   590,581 | 427,131 | 356,288 |
| 8  | 6  | 417,324 |   754,440 | 427,155 | 392,804 |
| 12 | 8  | 426,432 |   926,624 | 427,155 | 417,152 |
| 16 | 11 | 436,332 | 1,123,645 | 427,155 | 453,584 |

### MaxN = 32

| n | t | submitContribution | finalizeRound | submitPartialDecryption | combineDecryption |
|---|---|---|---|---|---|
| 4  | 3  | 500,533 | 1,056,772 | 425,105 | 373,929 |
| 8  | 6  | 559,110 | 1,221,403 | 425,069 | 410,469 |
| 12 | 8  | 616,911 | 1,394,615 | 425,069 | 434,793 |
| 16 | 11 | 675,508 | 1,593,124 | 425,069 | 471,261 |
| 20 | 14 | 734,085 | 1,809,766 | 425,033 | 507,825 |
| 24 | 16 | 791,922 | 2,026,049 | 425,033 | 532,137 |
| 28 | 19 | 850,494 | 2,276,365 | 425,057 | 568,641 |
| 32 | 22 | 909,011 | 2,494,009 | 425,069 | 605,097 |

### Side-by-side at the sizes both builds support (n = 4, 8, 12, 16)

| n  | Call | MaxN=32 | MaxN=16 | Δ |
|---|---|---|---|---|
| 4  | submitContribution | 500,533   |   407,460 | −18.6% |
| 4  | finalizeRound      | 1,056,772 |   590,581 | **−44.1%** |
| 4  | combineDecryption  | 373,929   |   356,288 | −4.7%  |
| 8  | submitContribution | 559,110   |   417,324 | −25.4% |
| 8  | finalizeRound      | 1,221,403 |   754,440 | **−38.2%** |
| 8  | combineDecryption  | 410,469   |   392,804 | −4.3%  |
| 12 | submitContribution | 616,911   |   426,432 | −30.9% |
| 12 | finalizeRound      | 1,394,615 |   926,624 | **−33.6%** |
| 12 | combineDecryption  | 434,793   |   417,152 | −4.1%  |
| 16 | submitContribution | 675,508   |   436,332 | **−35.4%** |
| 16 | finalizeRound      | 1,593,124 | 1,123,645 | **−29.5%** |
| 16 | combineDecryption  | 471,261   |   453,584 | −3.7%  |

`submitPartialDecryption` is essentially identical between builds (~425–427 k);
its dominant cost is the Groth16 verifier base which doesn't scale with N.

### Setup overhead (createRound + n × claimSlot)

`createRound` ≈ 194 k in both builds. `claimSlot` ≈ 120 k average per node
(first claimer pays the seed-resolve, last claimer pays the committee snapshot).
Setup overhead scales linearly with `n` and is essentially identical between
the two builds — none of it depends on MaxN.

### Whole-round totals at n = 16

| Phase | MaxN = 32 | MaxN = 16 | Saved |
|---|---|---|---|
| Setup (create + 16× claim)  |  2,108,488 |  2,118,801 |     −10,313 |
| 16× submitContribution      | 10,808,128 |  6,981,312 |   3,826,816 |
| 1×  finalizeRound           |  1,593,124 |  1,123,645 |     469,479 |
| 11× submitPartialDecryption |  4,675,759 |  4,698,705 |     −22,946 |
| 1×  combineDecryption       |    471,261 |    453,584 |      17,677 |
| **Round total at n=16**     | **19,656,760** | **15,376,047** | **4,280,713 (−21.8%)** |

The largest absolute saving comes from the 16 contribution calls (each shrinks
~239 k). The largest relative saving is in `finalizeRound` (−29.5%), driven by
the smaller calldata transcript and the smaller per-contributor digest input.

---

## Why MaxN = 16 is cheaper

1. **Smaller transcripts** → cheaper BRLC stream. The on-chain BRLC commitment
   is a streaming linear combination over the transcript at ~30 gas per word.

   | Transcript        | Word count | MaxN = 32 | MaxN = 16 |
   |---|---|---|---|
   | `submitContribution` | `8N`     | 256       | **128** |
   | `finalizeRound`      | `2N²+5N` | 2,208     | **592**  |
   | `combineDecryption`  | `4+3N`   | 100       | **52**   |
   | `reconstructSecret`  | `2N`     | 64        | **32**   |

   The `finalizeRound` transcript shrinks **3.7×**, which is why finalize gas
   drops the most in absolute terms.

2. **Smaller per-contributor digest in `finalizeRound`.** Each contributor's
   commitment slice is `2N` words = `64N` bytes:

   | Build  | Per-contributor digest input | `keccak256` cost / iter |
   |---|---|---|
   | MaxN=32 | 2,048 bytes | ~6.5 k gas |
   | MaxN=16 |   512 bytes | ~2.0 k gas |

3. **Smaller circuits** → ~4× constraint reduction in `contribution` and
   `finalize` (both dominated by `O(N²)` scalar multiplications inside
   `CommitmentPolynomialValue`), proportionally smaller proving keys, ~3-4×
   faster proving.

4. **What does NOT change**: Groth16 verifier base cost (~270 k per call,
   independent of N), cold SSTORE prices (22.1 k each), the lottery setup, and
   `submitPartialDecryption` (its cost is 100% verifier-base + 1 keccak + a
   couple of SSTOREs, none of which scale with N).

---

## Trade-offs

| Concern | MaxN = 32 | MaxN = 16 |
|---|---|---|
| Max committee per round           | 32        | **16** (hard cap) |
| `finalizeRound` worst case        | ~2.55 M (n=32) | ~1.12 M (n=16) |
| Whole round at n=16               | ~19.7 M   | **~15.4 M** (−22%) |
| Whole round at n=32               | ~33.0 M   | not supported |
| Per-node proof time (contribution)| ~3.4 s    | **~0.9 s** |
| Total proving wall-clock at full house | ~109 s | **~14 s** |
| Proving key size                  | ~2-4 GB   | **~0.5-1 GB** |
| Trusted-setup time                | minutes   | seconds |

**Choose MaxN = 16** if your protocol caps at 16 participants and you want the
cheapest per-round gas, the fastest provers, and the smallest proving-key
download.

**Choose MaxN = 32** if you need rounds of up to 32 participants under any
circumstances. The extra capacity costs ~3-4× in per-node proving time and
~22% more whole-round gas at n=16, but you maintain only one build.

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

Pipeline takes ~1 minute at MaxN=16, ~3 minutes at MaxN=32.

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

## Practical Limits

- **Block gas limit** (Anvil default): 30,000,000. Worst-case per call:
  - MaxN=32, n=32: `finalizeRound` ≈ 2.55 M (~8.5% of a block).
  - MaxN=16, n=16: `finalizeRound` ≈ 1.12 M (~3.7% of a block).
- **Throughput bottleneck** is proof generation, not on-chain gas. All calls
  fit in a single block in both builds.
- **Long-term storage** is bounded by a fixed-size ring buffer
  (`ROUND_HISTORY_SIZE = 64`). On every `createRound`, the oldest live round is
  evicted (`delete rounds[…]`), partially refunded by EIP-3529. Off-chain
  consumers reconstruct historical round data from the event log.
- **Memory** for proving:
  - MaxN=32: ~2-4 GB peak.
  - MaxN=16: ~0.5-1 GB peak.

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
