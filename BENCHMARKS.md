# DAVINCI DKG — Benchmarks

Reference figures for the **MaxN = 32** production build, with MaxN = 16 and
MaxN = 48 columns so operators can size proving infrastructure for other
committee bounds. The four Groth16 circuits are Contribution, Finalize,
PartialDecrypt and DecryptCombine.

> **Caveat — single-party trusted setup.** The reference artifacts come from a
> single-party Groth16 setup. A multi-party ceremony produces fresh pk/vk
> values; constraint counts and gas are unchanged by the ceremony (the verifier
> bytecode shifts by a few hundred bytes of vkey constants).

---

## Circuit constraint counts

Measured with `go run ./cmd/constraints` (`<circuit>.Compile()` +
`ccs.GetNbConstraints()`), September 2026.

| Circuit         | MaxN = 16 | MaxN = 32 | MaxN = 48 |
|-----------------|----------:|----------:|----------:|
| Contribution    |   254,270 |   541,218 |   866,120 |
| Finalize        |    65,343 |   211,600 |   444,799 |
| PartialDecrypt  |    22,061 |    22,061 |    22,061 |
| DecryptCombine  |   104,328 |   238,207 |   409,969 |

Contribution and Finalize scale with `MaxN × t` because every polynomial
evaluation is a Horner chain of `t` short scalar multiplications by the 6-bit
recipient index. PartialDecrypt is one DLEQ proof and does not depend on N.
DecryptCombine grows with `MaxN` fixed-base multiplications plus `MaxN²`
six-bit multiplications of the Vandermonde identity that pins the Lagrange
vector (the check is what makes the proof sound against a prover who knows the
discrete log of one partial; it grew the circuit from 88k to 238k at MaxN = 32).
The contribution share digest is one Poseidon per recipient row plus one over
the row digests, which keeps every sponge below its 256-input cap at any MaxN.

## Proof generation time

Wall-clock per single proof, median of five runs, gnark parallelising over
all 32 logical threads of an **AMD Ryzen 9 9950X3D (64 GiB RAM)**, idle
host. The time includes witness solving but not loading the proving key
(hundreds of ms once per process). Command:
`go test ./circuits/... -run XXX -bench '^BenchmarkProve$' -benchtime=5x`.

| Circuit         | MaxN = 16 | MaxN = 32 | MaxN = 48 |
|-----------------|----------:|----------:|----------:|
| Contribution    |    137 ms |    374 ms |    421 ms |
| Finalize        |     78 ms |    191 ms |    365 ms |
| PartialDecrypt  |     25 ms |     29 ms |     26 ms |
| DecryptCombine  |     70 ms |    130 ms |    220 ms |

Contribution grows little from MaxN = 32 to 48 because both sizes round up to
the same 2²⁰ FFT domain. A node therefore spends well under a second of CPU
per epoch (one contribution, possibly one finalization) and under 200 ms per
ciphertext it helps decrypt.

## On-chain gas (Sepolia, real verifiers)

Measured on the Sepolia deployment described in the README (manager
`0xb64f2d0870d2285f662e295f8a48adce79ccc56c`, deployed at block 11,616,449,
`MIN_THRESHOLD=2`, `MIN_COMMITTEE_SIZE=3`, `MAX_LOTTERY_ALPHA_BPS=20000`,
windows 100/25/25/5 blocks) with **MaxN = 32**, a committee of `n = 3`,
`t = 2` drawn from four operators at α = 1.5, epoch
`0x58c59e3f0000000000000001`. Gas is `gasUsed` from the receipt, so it
includes the pairing check, calldata and the intrinsic cost. The
N-dependent calldata (contribution `8·MaxN` words, finalization
`2·MaxN² + 5·MaxN` words) does not depend on `n`, so these are the costs any
`n ≤ 32` pays per call.

| Call                        |                 Gas | Paid by                                   | Evidence (tx) |
|-----------------------------|--------------------:|-------------------------------------------|---------------|
| `registerKey`               |             322,034 | operator, once (+21k for the first ever)  | `0x45b9b51d…` |
| `createEpoch`               |             150,232 | race winner, per epoch                    | `0xc8257e89…` |
| `claimSlot`                 | 103,677 – 175,345   | each member (seed resolution / key snapshot on first / last claim) | `0xaea30641…`, `0xa0dcf58a…` |
| `submitContribution`        |             462,584 | each member                               | `0xc1096c41…` |
| `finalizeEpoch`             |           1,112,348 | one member, per epoch                     | `0xe3d89391…` |
| `registerApplication`       |             167,722 | organizer, mode 0                         | `0x8f2a3f02…` |
| `registerApplicationCoDec`  |             375,499 | organizer, mode 1 (Schnorr PoK on chain)  | `0xda16ef20…` |
| `submitCiphertext`          |             468,630 | submitter (subgroup check + Schnorr PoK on chain) | `0x8ce94ee4…` |
| `submitPartialDecryption`   | 388,406 – 405,458   | each of `t` members                       | `0xaded2308…`, `0x703bc15f…` |
| `submitOrganizerShare`      |             435,851 | organizer, per ciphertext (mode 1)        | `0x35656ff8…` |
| `combineDecryption`         | 425,680 / 430,511   | one member, per ciphertext (mode 0 / 1)   | `0xa48c1b18…`, `0x014493d0…` |
| deployment (4 verifiers + 3 contracts) | 12,764,762 | deployer, once                    | `broadcast/DeployAll.s.sol/11155111/run-latest.json` |

Reading the table:

* Every proof-carrying call pays roughly 250k gas for the Groth16 verification
  itself (four pairings plus the public-input multiplications, EIP-1108
  prices); calldata explains the rest, which is why the contribution
  (256 words) and the finalization (2,208 words) dominate.
* An epoch with `n` members costs about `1.26M + n × 0.61M` gas at MaxN = 32:
  3.1M for the `n = 3` committee above, about 20.7M for `n = 32`.
* Decrypting one ciphertext costs 0.47M (submission) + `t` × ~0.4M (partials)
  + 0.43M (combine): 1.7M at `t = 2` in mode 0, 2.1M in mode 1 with the
  organizer share. The on-chain prime-subgroup check (~168k) and Schnorr proof
  of knowledge (~202k, measured in isolation with full-size scalars) account
  for ~370k of the submission; they reject malformed ciphertexts before any
  member spends a partial on them.
* `registerKey` is dominated by the Schnorr verification in extended
  twisted-Edwards coordinates; the affine implementation cost 1.27M.

## How to reproduce

* Constraints: `go run ./cmd/constraints` after setting `MaxN` in
  `circuits/common/sizes.go` (and `MAX_N` in `solidity/src/libraries/Sizes.sol`
  for the contracts).
* Proving times: `DAVINCI_ARTIFACTS_DIR=/tmp/bench-$N go test ./circuits/... -run XXX -bench '^BenchmarkProve$' -benchtime=5x`
  on an idle host; the first run performs the setup.
* Gas: run the flow on any deployment (`make testnet-up`, or the Sepolia
  deployment with `cmd/dkgapp`) and read `gasUsed` from the receipts, e.g.
  `cast receipt --json <tx> | jq .gasUsed`. `forge test --gas-report` gives the
  same per-function figures against mock verifiers, i.e. without the ~250k
  pairing check.
