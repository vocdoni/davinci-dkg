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
| PartialDecrypt  |    22,026 |    22,026 |    22,026 |
| DecryptCombine  |   113,200 |   247,076 |   418,823 |

Contribution and Finalize scale with `MaxN × t` because every polynomial
evaluation is a Horner chain of `t` short scalar multiplications by the 6-bit
recipient index. PartialDecrypt is one committee DLEQ proof and does not
depend on N. DecryptCombine grows with `MaxN` fixed-base multiplications plus
`MaxN²` six-bit multiplications of the Vandermonde identity that pins the
Lagrange vector; it also verifies the organizer's Chaum–Pedersen proof (two
double-base equations, about 9k constraints) so that the organizer never
needs a prover. The contribution share digest is one Poseidon per recipient
row plus one over the row digests, which keeps every sponge below its
256-input cap at any MaxN.

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
| PartialDecrypt  |     26 ms |     25 ms |     26 ms |
| DecryptCombine  |     75 ms |    134 ms |    222 ms |

Contribution grows little from MaxN = 32 to 48 because both sizes round up to
the same 2²⁰ FFT domain. A node therefore spends well under a second of CPU
per epoch (one contribution, possibly one finalization) and under 200 ms per
ciphertext it helps decrypt.

## On-chain gas (Sepolia, real verifiers)

Measured on the Sepolia deployment described in the README (manager
`0x3f9b338706a31f26d49159478015c8aaeab908ad`, deployed at block 11,619,019,
`MIN_THRESHOLD=2`, `MIN_COMMITTEE_SIZE=3`, `MAX_LOTTERY_ALPHA_BPS=20000`,
windows 100/25/25/5 blocks) with **MaxN = 32**, a committee of `n = 3`,
`t = 2` drawn from four operators at α = 1.5, epoch
`0x2f1105e90000000000000001`. Gas is `gasUsed` from the receipt, so it
includes the pairing check, calldata and the intrinsic cost. The
N-dependent calldata (contribution `8·MaxN` words, finalization
`2·MaxN² + 5·MaxN` words, combine `12 + 3·MaxN` words) does not depend on
`n`, so these are the costs any `n ≤ 32` pays per call. Every transaction is
an event on the deployment above; the hashes given are examples.

| Call                        |                 Gas | Paid by                                        | Evidence (tx) |
|-----------------------------|--------------------:|------------------------------------------------|---------------|
| `registerKey`               |             322,112 | operator, once (+17k for the first ever)       | `NodeRegistered` events from block 11,619,019 |
| `createEpoch`               |             150,279 | race winner, per epoch                         | `EpochCreated` |
| `claimSlot`                 | 103,725 – 175,520   | each member (seed resolution / key snapshot on first / last claim) | `SlotClaimed` |
| `submitContribution`        |             462,523 | each member                                    | `ContributionSubmitted` |
| `finalizeEpoch`             |           1,112,337 | one member, per epoch                          | `0x68137083…` |
| `registerApplication`       |             407,793 | organizer, per application (Schnorr proof of possession on chain) | `0xe1a0230e…` |
| `submitCiphertext`          |    96,001 / 78,901  | authorised submitter (first / later ciphertext) | `0x2d086c97…`, `0xc620ee49…` |
| `submitPartialDecryption`   | 381,604 – 398,704   | each of `t` members                            | `0x2064118e…`, `0xde371a95…` |
| `submitOrganizerShare`      |    87,991 / 70,879  | organizer, per ciphertext (first / overwrite)  | `0x900886c6…`, `0x85642284…` |
| `combineDecryption`         |             430,432 | one member, per ciphertext                     | `0x9cc53df3…` |
| deployment (4 verifiers + 3 contracts) | 11,865,129 | deployer, once                       | `broadcast/DeployAll.s.sol/11155111/run-latest.json` |

Reading the table:

* Every proof-carrying call pays roughly 250k gas for the Groth16 verification
  itself (four pairings plus the public-input multiplications, EIP-1108
  prices); calldata explains the rest, which is why the contribution
  (256 words) and the finalization (2,208 words) dominate.
* An epoch with `n` members costs about `1.26M + n × 0.61M` gas at MaxN = 32:
  3.1M for the `n = 3` committee above, about 20.7M for `n = 32`.
* Decrypting one ciphertext costs 0.10M (submission) + `t` × ~0.39M
  (partials) + 0.09M (organizer share) + 0.43M (combine): 1.4M at `t = 2`.
  The organizer's two transactions together cost less than a fifth of one
  partial; its Chaum–Pedersen proof is stored as a hash and verified inside
  the combine proof, which is why the organizer needs neither a prover nor
  half a million gas of on-chain curve arithmetic.
* `registerApplication` and `registerKey` are dominated by a Schnorr
  verification in extended twisted-Edwards coordinates (~200k); the affine
  implementation cost 1.27M for `registerKey`.

## A 32-node fleet under load (Anvil, two hosts)

Measured on 2026-09-03 with `tests/battery`: 32 operators split over two
32-core hosts, committee n = 24, t = 16, m_min = 20, 2-second blocks, 300-block
epochs. Eight organizers register concurrently and submit six ciphertexts
each (48 in one block), releasing shares immediately, after six blocks, or
never for a quarter of them.

| | before the fix | after |
|---|---|---|
| ciphertexts decrypted / released | 40 / 40 | 40 / 40 |
| withheld shares never combined | 8 / 8 | 8 / 8 |
| partials on chain per ciphertext | 18–24 (t = 16) | 16 for 41 of 48, ≤ 23 |
| submit → combine latency (avg) | 50 blocks / 100 s | 17 blocks / 35 s |
| throughput | 0.30 ct/s | 0.77 ct/s |
| gas: contribution / finalize (n = 24) | 513 k / 2.14 M | same |
| gas: partial / combine / register / share | 399 k / 430 k / 390 k / 88 k | same |
| same swarm with 4 local nodes restarted, Anvil paused 20 s and 6 remote nodes down 90 s mid-run | — | 40 / 40 decrypted, 23 blocks avg, 0.42 ct/s |

The fix (`node/decrypt.go`): partial and combine transactions are sent
without waiting for their receipt, later waves only fire when the earlier
ones have stopped landing partials, and the combine (dlog + proof) runs off
the service loop, one at a time per node. Every adversarial scenario
(tampered, replayed and relayed shares; ciphertexts outside the window, over
the cap, off-curve, in a small-order subgroup, undecryptable or copied from
another application; late registration, duplicate claims, contributions and
partials; early finalize; abort of a healthy epoch) is rejected as designed.
Two costs are accepted and documented: an undecryptable ciphertext makes each
member of the combine rotation run one bounded 2^50 BSGS search (~20 s of
CPU, 256 MB table built once per process), after which the node ignores the
rest of that application's ciphertexts for the epoch (one search per
registration, not per ciphertext); and a withheld share parks its slot in
every node (no per-tick cost) until an organizer share event wakes it.

## Groth16 versus PLONK (universal setup)

Measured on 2026-09-03 with gnark v0.16.3 / gnark-crypto v0.21.0 at
`MaxN = 32` (`go test ./circuits/... -bench BenchmarkBackends -benchtime=1x`,
Ryzen 9 9950X3D, 32 threads, average of 3 proofs; PLONK with an unsafe test
SRS, which does not change prover cost). On-chain gas is the deployed Groth16
combine verifier on a real proof versus gnark's exported PLONK verifier for
the same statement (11 public inputs), both under `via_ir`, `optimizer_runs = 1`.

| circuit | Groth16 R1CS | PLONK gates | Groth16 prove | PLONK prove | proof bytes G16 / PLONK |
|---|---:|---:|---:|---:|---:|
| Contribution | 566,419 | 1,242,994 | 0.37 s | 12.53 s | 256 / 864 |
| Finalize | 211,600 | 929,175 | 0.18 s | 4.12 s | 256 / 864 |
| PartialDecrypt | 26,186 | 58,040 | 0.03 s | 0.46 s | 256 / 864 |
| DecryptCombine | 274,615 | 606,441 | 0.21 s | 4.96 s | 256 / 864 |

| | Groth16 | PLONK |
|---|---:|---:|
| combine verifier gas (11 inputs) | 259,694 | 283,509 |
| proof calldata | 256 B | 864 B (+9.7 k gas) |
| proving key, contribution | 86 MB | 134 MB + 134 MB SRS |
| verifying key | ~0.8 kB | 34 kB |
| setup | per circuit (four ceremonies, or one party) | one universal KZG SRS of 2^21 points (Aztec / Hermez transcripts qualify) |

PLONK removes the per-circuit trusted setup at the price of 2.2–4.4× more
constraints and 15–34× slower proving in gnark (contribution 0.4 s → 12.5 s
on 32 threads; a node's partial 30 ms → 0.5 s), about 13 % more gas per
verification once calldata is counted, and 3.4× larger proofs. Functionally
everything fits the protocol windows even on modest hardware, and the
migration is mechanical (`scs` builder, `plonk` backend, gnark's PLONK
Solidity export, 864-byte proofs and a `Verify(bytes, uint256[])` ABI), so
the choice is purely setup trust versus prover cost. Groth16 stays the
default; a Groth16 MPC ceremony for the four circuits is the cheaper way to
remove the single-party setup.

The same run showed that upgrading gnark from the pinned v0.14 snapshot to
v0.16.3 changes the compiled R1CS (partial +19 %, combine +11 %,
contribution +5 %, finalize unchanged) and therefore the pinned artifact
hashes: the upgrade has to ship with a circuit re-release.

## How to reproduce

* Constraints: `go run ./cmd/constraints` after setting `MaxN` in
  `circuits/common/sizes.go` (and `MAX_N` in `solidity/src/libraries/Sizes.sol`
  for the contracts).
* Proving times: `DAVINCI_ARTIFACTS_DIR=/tmp/bench-$N go test ./circuits/... -run XXX -bench '^BenchmarkProve$' -benchtime=5x`
  on an idle host; the first run performs the setup.
* Gas: run the flow on any deployment (`make testnet-up`, or the Sepolia
  deployment with `cmd/dkgapp`: `register`, `encrypt`, `share`, `plaintext`)
  and read `gasUsed` from the receipts, e.g. `cast receipt --json <tx> | jq
  .gasUsed`. `forge test --gas-report` gives the same per-function figures
  against mock verifiers, i.e. without the ~250k pairing check.
