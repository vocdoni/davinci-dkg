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

Two new constant-cost calls and one change since the last benchmark:

- **`submitCiphertext` ≈ 65,800 gas** (independent of `n`). Includes the
  on-curve + canonical-coord check on `(c1, c2)` (~2k gas), one cold SSTORE
  of `keccak256(c1, c2)` (22.1k), increment of `ciphertextCount`, and a
  7-topic event. Measured 65,820 at n=4 and 65,784 at n=32 — the small
  delta is storage-slot warmth noise.
- **`createRound` ≈ 205 k** (was ~194 k). The `DecryptionPolicy` struct
  occupies 2 extra SSTOREs on the round creation (≈ 9k gas); this is a
  one-time per-round cost, independent of committee size.
- **`combineDecryption` +25 k gas flat** across all `n`. Cost split:
  ~22 k for the cold SSTORE of the recovered plaintext into
  `CombinedDecryptionRecord.plaintext` (new — previously only `completed`
  was stored), plus ~3 k for the `SLOAD` of the stored ciphertext hash and
  `keccak256` over the transcript's first 128 bytes that binds the combine
  proof to the submitted ciphertext.

### MaxN = 16

| n | t | submitContribution | finalizeRound | submitPartialDecryption | combineDecryption |
|---|---|---|---|---|---|
| 4  | 3  | 454,701 |   610,016 | 427,826 | 358,595 |
| 8  | 6  | 464,469 |   774,651 | 427,802 | 395,672 |
| 12 | 8  | 473,625 |   947,599 | 427,826 | 420,426 |
| 16 | 11 | 483,525 | 1,145,456 | 427,814 | 457,539 |

*(MaxN=16 numbers are pre-submitCiphertext; add ~66k for the new
`submitCiphertext` call and ~+25k to `combineDecryption` for the plaintext
persistence to compute current totals. The ring is expected to shift by the
same constants as MaxN=32 below.)*

### MaxN = 32

| n | t | submitContribution | finalizeRound | submitPartialDecryption | **submitCiphertext** | combineDecryption |
|---|---|---|---|---|---|---|
| 4  | 3  | 491,260 | 1,062,412 | 427,832 | 65,820 | 400,112 |
| 16 | 11 | 520,100 | 1,604,296 | 427,880 | 65,820 | 499,032 |
| 32 | 22 | 558,896 | 2,563,850 | 427,820 | 65,784 | 635,077 |

`createRound` was 204,835 gas in all three runs (independent of `n`).

Rows at n=8, 12, 20, 24, 28 were not re-measured in this pass; they shift
uniformly from the pre-change table above by **+25,430 gas in
`combineDecryption`** and **+65,800 gas for the new `submitCiphertext`
call**. `submitContribution` / `finalizeRound` / `submitPartialDecryption`
shift by <0.1% (measurement noise).

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

### Whole-round totals at n = 16 (MaxN = 32)

Post-change measured figures:

| Phase | Gas |
|---|---|
| createRound                  |    204,835 |
| 16× claimSlot                |  1,905,000 (≈118k avg × 16) |
| 16× submitContribution       |  8,321,600 (520,100 × 16) |
| 1×  finalizeRound            |  1,604,296 |
| 1×  submitCiphertext         |     65,820 |
| 11× submitPartialDecryption  |  4,706,680 (427,880 × 11) |
| 1×  combineDecryption        |    499,032 |
| **Round total at n=16**      | **17,307,263** |

### Whole-round totals at n = 32 (MaxN = 32)

| Phase | Gas |
|---|---|
| createRound                  |    204,835 |
| 32× claimSlot                |  3,776,000 (≈118k avg × 32) |
| 32× submitContribution       | 17,884,672 (558,896 × 32) |
| 1×  finalizeRound            |  2,563,850 |
| 1×  submitCiphertext         |     65,784 |
| 22× submitPartialDecryption  |  9,412,040 (427,820 × 22) |
| 1×  combineDecryption        |    635,077 |
| **Round total at n=32**      | **34,542,258** |

Net effect vs. previous (pre-DecryptionPolicy / pre-submitCiphertext) build:
n=16 +98 k (+0.6%), n=32 +104 k (+0.3%). The increase is dominated by the
new `submitCiphertext` call and the on-chain plaintext persistence, both of
which are constants independent of `n`. Per-node per-phase costs are
unchanged.

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

0. **Ciphertext submission** (new). `submitCiphertext` is a flat ~66 k: on-curve
   + canonical-field check on the two BabyJubJub points (~2 k), one cold SSTORE
   for `keccak256(c1,c2)` (22.1 k), one cold SSTORE bump of `ciphertextCount`
   (warmed to 5 k after first write), and a 7-topic event (~3 k). Independent
   of N.
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
