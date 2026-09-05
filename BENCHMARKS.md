# DAVINCI DKG — Benchmarks

Reference figures for the **v4 production build** (MaxN = 32, MaxK = 16: a pool of sixteen
committee-held application keys per epoch, finalized by one proof, see `docs/pool-keys.md`).
The four Groth16 circuits are Contribution, Finalize, PartialDecrypt and DecryptCombine. The
v3.1 tables further down (eight keys, per-key activation proofs) and the pre-upgrade tables are
kept for comparison; `MaxN` must be a power of two (the share-commitment Merkle tree has
`2^MERKLE_DEPTH = MaxN` leaves).

> **Caveat — single-party trusted setup.** The reference artifacts come from a
> single-party Groth16 setup. A multi-party ceremony produces fresh pk/vk
> values; constraint counts and gas are unchanged by the ceremony (the verifier
> bytecode shifts by a few hundred bytes of vkey constants).

> **Two toolchains.** The v4 section at the top is the current build (gnark
> v0.16.3 / gnark-crypto v0.21.0, measured 2026-09-05); the v3.1 sections
> (eight keys, per-key activation proofs) and the sections marked
> **pre-upgrade** were measured with the previously pinned gnark
> snapshot (`v0.14.1-0.20260126…`) and the first pool-key cut, and are kept
> for comparison. That snapshot — like every gnark release up to v0.15.0 —
> has an unsound variable-base twisted-Edwards `ScalarMul`: the fake-GLV
> decomposition check leaves its quotient unconstrained, so a prover can make
> the gadget return any point (fixed upstream in gnark v0.16.0, PR #1765,
> without an advisory; the weaker cofactor-torsion offset is IACR ePrint
> 2026/1776). Every circuit is recompiled and re-pinned against v0.16.3 and no
> circuit uses the hinted gadget any more; the v3.1 activation proof also adds
> a public input (`transcriptDigest`) and one calldata word.

---

## Current release (v4): batched finalization and compact contributions (MaxN = 32, MaxK = 16, gnark v0.16.3)

Measured 2026-09-05 on the v4 build now on main and deployed on Sepolia: one proof-carrying `finalizeEpoch` activates
all 16 keys, contributions carry the unpadded transcript of `MaxK·(2t+n) + 5n` words, MaxN = 32.
Same machine and method as the v3.1 tables (Anvil with the real verifiers, receipts' `gasUsed`;
proving mean of five, 32 threads).

| Circuit | Constraints | Public inputs | Proving time |
|---|---:|---:|---:|
| Contribution (16 keys, compact fold) | 5,904,167 | 8 | 3,540 ms |
| Finalize (all 16 keys) | 2,328,130 | 7 | 2,259 ms |
| PartialDecrypt | 29,026 | 15 | unchanged |
| DecryptCombine | 287,338 | 9 | unchanged |

| Call | n = 4, t = 3 | n = 32, t = 22 | Notes |
|---|---:|---:|---|
| `createEpoch` | 133,269 | 133,269 | 150,280 for the first seed |
| `claimSlot` (avg) | 137,463 | 121,021 | |
| `submitContribution` | 500,070 | 1,314,499 | 16 keys; 1.5 KB of calldata per member per key |
| `finalizeEpoch` | 2,201,937 | 2,810,545 | verifier + 16 keys + 16 roots (48 cold stores) |
| `registerApplication` locked / automatic | 619,074 / 232,394 | 620,994 / 232,394 | locked includes the 168k organizer-key subgroup check |
| `revealOrganizerSecret` | 225,902 | 223,750 | |
| `submitCiphertext` | 102,770 | 102,770 | |
| `submitPartialDecryption` | 402,251 | 402,223 | |
| `combineDecryption` | 410,614 | 483,955 | |

Sweep (n = 8 … 28): contribution 622,141 / 722,287 / 844,682 / 967,364 / 1,068,474 / 1,191,409;
finalize 2,288,989 / 2,375,993 / 2,462,841 / 2,549,557 / 2,636,513 / 2,723,469.

An epoch yielding 16 keys costs about `0.15 + 4 × 0.14 + 4 × 0.50 + 2.20 ≈ 4.9 M` gas at n = 4
(0.31 M per key, against 0.95 M per key in v3.1 and 3.1 M for a single-key epoch) and about
`32 × (0.12 + 1.31) + 2.81 ≈ 48.9 M` at n = 32 (3.1 M per key, against 5.1 M). The price is on the
node: the contribution proof takes 3.5 s instead of 1.6 s and the proving key doubles; the contribution
prover peaks at 9.3 GB resident (5.0 GB at MaxK = 8) and the finalize prover at 5.5 GB (`/usr/bin/time -v`
around the benchmark, proving key and circuit loaded). With MaxK = 8 the same design would keep the v3.1 proving cost and land
near 0.5 M per key at n = 4.

**Hardware guidance:** a v0.5.0 node with all four circuits and their proving keys preloaded
sits at 8.5–9.6 GB of resident memory at rest (docker stats on the three Sepolia seed nodes
right after start, against 3.1 GB for a v0.4 node), and peaks at about 9.8 GB during its first
contribution proof (sampled every five seconds across the three nodes). State the requirement as
**16 GB of RAM minimum, more is safer** — the 9.3 GB proving peak sits on top of the ~9 GB
resting footprint, so a node with less RAM will swap or be killed mid-proof.

## Circuit constraint counts after the gnark upgrade (v3.1, superseded by v4, 2026-09-04)

Recompiled with gnark v0.16.3 / gnark-crypto v0.21.0 after the pool-key v3.1
changes (activation transcript digest, whole-committee share commitments,
canonical Lagrange mask) and with every variable-base scalar multiplication
done by the hint-free `ccommon.ScalarMulVar` gadget instead of gnark's hinted
fake-GLV (the pinned snapshot's version was unsound; the fixed one commits and
would have added a pairing to every proof). `go run ./cmd/constraints`, MaxN = 32,
MaxK = 8:

| Circuit         | constraints |
|-----------------|------------:|
| Contribution    |   3,060,692 |
| PoolKey         |     187,495 |
| PartialDecrypt  |      29,026 |
| DecryptCombine  |     287,338 |

## Proof generation time (v3.1, superseded by v4, gnark v0.16.3)

Wall-clock per single proof, mean of five runs (`go test ./circuits/... -run XXX
-bench '^BenchmarkProve$' -benchtime=5x`, which reports the mean), gnark
parallelising over all 32 logical threads of an **AMD Ryzen 9 9950X3D (64 GiB
RAM)**, idle host, 2026-09-05. The time includes witness solving but not
loading the proving key (a few seconds once per process for the contribution
key, which is on the order of a gigabyte; the node preloads all four).

| Circuit         | MaxN = 32 | pre-upgrade |
|-----------------|----------:|------------:|
| Contribution    |  1,574 ms |    1,687 ms |
| PoolKey         |    181 ms |      246 ms |
| PartialDecrypt  |     32 ms |       28 ms |
| DecryptCombine  |    213 ms |      134 ms |

The two large circuits prove faster than before despite more constraints; the
two small ones pay for the hint-free multiplications (a 254-step constrained
double-and-add per variable-base product instead of a hinted fake-GLV). A node
still spends under two seconds of CPU per epoch on its contribution, under a
fifth of a second per key it activates, and about a quarter of a second per
ciphertext it combines.

## Single-key baseline (MaxK = 1, same toolchain, for comparison)

Measured 2026-09-05 with the v3.1 code at commit d8c186a, `MaxK = 1` / `MAX_K = 1`, the same
gnark v0.16.3 toolchain, MaxN = 32, on the same machine, so the pool can be compared with a
single-key deployment without mixing toolchains.

| Circuit | Constraints, MaxK = 1 | Constraints, MaxK = 8 | Proving, MaxK = 1 | Proving, MaxK = 8 |
|---|---:|---:|---:|---:|
| Contribution | 587,022 | 3,060,692 | 447 ms | 1,574 ms |
| Pool-key activation | 181,881 | 187,495 | 277 ms | 181 ms |
| PartialDecrypt | 29,026 | 29,026 | — | 32 ms |
| DecryptCombine | 287,338 | 287,338 | — | 213 ms |

Proving times are the mean of five (`-benchtime=5x`); the activation figure at MaxK = 1 was taken
while other jobs ran and is noisy.

| Call | MaxK = 1, n = 4 | MaxK = 1, n = 32 | MaxK = 8, n = 4 | MaxK = 8, n = 32 |
|---|---:|---:|---:|---:|
| `submitContribution` | 467,536 | 535,120 | 702,013 | 944,425 |
| `activatePoolKey` | 508,266 | 809,066 | 508,137 | 808,913 |
| `submitPartialDecryption` | 402,320 | 402,340 | 402,332 | 402,340 |
| `combineDecryption` | 410,639 | 484,028 | 410,651 | 483,980 |

Registration and reveal are not comparable across the two runs because the MaxK = 1 build includes
the audit fixes (the organizer-key subgroup check adds about 168k gas to a locked registration).

Reading it: a contribution dealing eight keys costs 1.5–1.8× a single-key one, so per key the pool
is 4.5–5.3× cheaper on contributions; activation is per key in both. An epoch that yields one usable
key costs about 3.1 M gas at n = 4 (creation, four claims, four contributions, finalization, one
activation) against 0.95 M per key for a fully used eight-key epoch.

## On-chain gas (Anvil, real verifiers, v3.1 — superseded by v4)

Measured on 2026-09-05 with the verifiers generated from the re-pinned gnark
v0.16.3 artifacts, same procedure as the pre-upgrade table below
(`RUN_BENCHMARKS=true` single-node profile and `RUN_BENCHMARKS_MULTI=true`
sweep, each in its own process, harness windows widened through
`COMMITTEE_SELECTION_BLOCKS=50 KEY_ASSEMBLY_BLOCKS=120 EPOCH_DURATION_BLOCKS=250`),
**MaxN = 32, MaxK = 8**, `gasUsed` from the receipts.

| Call                        | n = 4, t = 3 | n = 32, t = 22 | Paid by |
|-----------------------------|-------------:|---------------:|---------|
| `createEpoch`               |      133,313 |        133,313 | race winner, per epoch (150,324 for the epoch that resolves the first seed) |
| `claimSlot`                 |      138,831 |        119,505 | each member (average; the first claim pays the seed resolution, 205,116) |
| `submitContribution`        |      702,013 |        944,425 | each member (8 polynomials; +8.7k per extra member of calldata) |
| `finalizeEpoch`             |       34,806 |         34,806 | one member, per epoch (no proof) |
| `activatePoolKey`           |      508,137 |        808,913 | one member, per key (+10.7k per extra member) |
| `registerApplication`       | 453,167 / 234,269 | 455,715 / 234,257 | organizer, per application (organizer-locked / automatic) |
| `revealOrganizerSecret`     |      223,307 |        228,439 | organizer, once per locked application |
| `submitCiphertext`          |      102,814 |        102,802 | authorised submitter |
| `submitPartialDecryption`   |      402,332 |        402,340 | each of `t` members (5-word Merkle path included) |
| `combineDecryption`         |      410,651 |        483,980 | one member, per ciphertext (+3.9k per extra partial) |

Intermediate committee sizes (8, 12, 16, 20, 24, 28): contribution 738,217 /
768,421 / 804,793 / 841,141 / 871,729 / 908,053; activation 551,297 / 593,797 /
636,945 / 680,093 / 722,533 / 765,717; combine 422,234 / 429,923 / 441,550 /
453,118 / 460,818 / 472,483. The other calls do not depend on `n`.

Against the pre-upgrade table: activation +24.5k (+5 %), contribution +15.6k
(+2.3 %), combine +3.0k, reveal +2.5k, partial +1.6k; every other call moves by
less than 0.1 %. The v3.1 changes that touch these calls are one more public
input and the digest word for activation, a tagged Merkle tree over all 32
member slots (activation, partial), the reveal-gate storage read (partial,
combine), and verifiers that now receive the proof as `bytes`.

Reading the table:

* Every proof-carrying call pays roughly 250k gas for the Groth16 verification
  itself; calldata explains the rest. A contribution carries `3·MaxK·MaxN +
  5·MaxN = 928` words, most of them commitments and masked shares of the
  eight keys; the activation transcript is `6·MaxN = 192` words plus the row
  checks against every accepted contribution.
* An epoch with `n` members and all eight keys activated costs about
  `n × 0.70M + 8 × 0.51M` at `n = 4` (6.9M, i.e. 0.86M per key) and
  `32 × 0.94M + 8 × 0.81M ≈ 37M` at `n = 32` (4.6M per key). Keys are
  activated lazily, two ahead of demand, so an epoch that serves few
  applications pays for few activations.
* Decrypting one ciphertext costs 0.10M (submission) + `t` × 0.40M (partials)
  + 0.41–0.48M (combine): 1.7M at `t = 3`. There is no organizer transaction
  per ciphertext; an organizer-locked application pays one 0.22M reveal when
  it opens.
* `registerApplication` in organizer-locked mode is dominated by the Schnorr
  proof of possession (~200k); the automatic mode skips it. `revealOrganizerSecret`
  pays one fixed-base multiplication to check the secret against `PK_org`.

## Circuit constraint counts (pre-upgrade)

Rank-1 constraint system (R1CS) sizes, measured with
`go run ./cmd/constraints` (`<circuit>.Compile()` + `ccs.GetNbConstraints()`),
2026-09-04, with the gnark v0.14 snapshot and the 7-input activation circuit
(the v3.1 table above has the current counts).

| Circuit         | MaxN = 16 | MaxN = 32 |
|-----------------|----------:|----------:|
| Contribution    | 1,359,732 | 3,015,892 |
| PoolKey         |    66,859 |   215,291 |
| PartialDecrypt  |    22,026 |    22,026 |
| DecryptCombine  |   107,259 |   241,138 |

The contribution circuit deals all `MaxK` polynomials in one proof: per key
and recipient it evaluates the Feldman polynomial (a Horner chain of `t`
six-bit multiplications), one fixed-base multiplication for the share and one
Poseidon share mask, so it scales with `MaxK × MaxN × t`; the ECDH secret per
recipient is shared by every key. PoolKey recomputes each contributor's
two-level Poseidon commitment digest for the activated key (about 2,400
absorptions at MaxN = 32) and the Vandermonde share commitments, and replaces
the old Finalize circuit (211,600 constraints) at a similar size while
dropping the `2·N²`-word contributor matrix from calldata. PartialDecrypt is
unchanged. DecryptCombine proves knowledge of the organizer secret instead of
verifying a Chaum–Pedersen transcript and is 2.4% smaller than before.

## Proof generation time (pre-upgrade)

Wall-clock per single proof, median of five runs, gnark parallelising over
all 32 logical threads of an **AMD Ryzen 9 9950X3D (64 GiB RAM)**, idle
host, 2026-09-04, with the gnark v0.14 snapshot (the v3.1 table above has the
current times). The time includes witness solving but not loading the
proving key (a few seconds once per process for the contribution key).
Command: `go test ./circuits/... -run XXX -bench '^BenchmarkProve$' -benchtime=5x`.

| Circuit         | MaxN = 32 |
|-----------------|----------:|
| Contribution    |  1,687 ms |
| PoolKey         |    246 ms |
| PartialDecrypt  |     28 ms |
| DecryptCombine  |    134 ms |

A node therefore spends under two seconds of CPU per epoch on its
contribution plus a quarter of a second per key it activates, and under 200
ms per ciphertext it helps decrypt. The contribution circuit's Groth16 setup
takes about ten minutes on this host and its proving key is on the order of
a gigabyte; the circuit package's tests need `-timeout 120m`. Peak memory
during setup and proving has not been measured yet.

## On-chain gas (Anvil, real verifiers, pre-upgrade)

Measured on 2026-09-04 with the real verifier contracts of the first pool-key
cut on the Docker test
stack (`RUN_INTEGRATION_TESTS=true RUN_BENCHMARKS=true go test ./tests -run TestGasProfiles`
for the single-node profile and `RUN_BENCHMARKS_MULTI=true` for the sweep
over committee sizes, each in its own process, with the harness epoch
windows widened through `COMMITTEE_SELECTION_BLOCKS=50 KEY_ASSEMBLY_BLOCKS=120
EPOCH_DURATION_BLOCKS=250` so that 32 sequential proofs fit), **MaxN = 32,
MaxK = 8**, reading `gasUsed` from the receipts.

| Call                        | n = 4, t = 3 | n = 32, t = 22 | Paid by |
|-----------------------------|-------------:|---------------:|---------|
| `createEpoch`               |      133,291 |        133,291 | race winner, per epoch (150,302 for the epoch that resolves the first seed) |
| `claimSlot`                 |      138,831 |        119,505 | each member (average; the first claim pays the seed resolution) |
| `submitContribution`        |      686,441 |        928,865 | each member (8 polynomials; +8.7k per extra member of calldata) |
| `finalizeEpoch`             |       34,781 |         34,781 | one member, per epoch (no proof) |
| `activatePoolKey`           |      483,639 |        797,131 | one member, per key (+11.2k per extra member) |
| `registerApplication`       |      452,505 / 234,235 | 452,481 / 234,235 | organizer, per application (organizer-locked / automatic) |
| `revealOrganizerSecret`     |      220,759 |        221,905 | organizer, once per locked application |
| `submitCiphertext`          |      102,792 |        102,792 | authorised submitter |
| `submitPartialDecryption`   |      400,767 |        400,773 | each of `t` members (5-word Merkle path included) |
| `combineDecryption`         |      407,651 |        481,004 | one member, per ciphertext |

Intermediate committee sizes (8, 12, 16, 20, 24, 28) interpolate linearly:
contribution 722,657 / 752,885 / 789,233 / 825,617 / 856,181 / 892,517 and
activation 528,655 / 572,927 / 617,919 / 662,851 / 707,111 / 752,091.

The previous release's Sepolia figures (single epoch key, per-ciphertext
organizer share) were: contribution 462,523, finalize 1,112,337,
registerApplication 407,793, partial 381,604–398,704, organizer share
87,991, combine 430,432, deployment 11,865,129 (not re-measured).

## A 32-node fleet under load (Anvil, two hosts)

Measured with `tests/battery` on 2026-09-03, before pool keys (single epoch
key, per-ciphertext organizer shares): 32 operators split over two
32-core hosts, committee n = 24, t = 16, m_min = 20, 2-second blocks,
300-block epochs. Eight organizers register concurrently and submit six
ciphertexts each (48 in one block), releasing shares immediately, after six
blocks, or never for a quarter of them.

| | quiet fleet | 4 local nodes restarted, Anvil paused 20 s, 6 remote nodes down 90 s |
|---|---:|---:|
| ciphertexts decrypted / released | 40 / 40 | 40 / 40 |
| withheld shares never combined | 8 / 8 | 8 / 8 |
| partials on chain per ciphertext | 16 for 41 of 48, never above 23 | 16 to 23 |
| submit to combine, average | 17 blocks (35 s) | 23 blocks (65 s) |
| throughput | 0.77 ciphertexts/s | 0.42 ciphertexts/s |

Gas on this fleet: contribution 513 k and finalize 2.14 M at n = 24; partial
399 k, combine 430 k, registerApplication 390 k, organizer share 88 k.

A node sends its partial and combine transactions without waiting for the
receipt, so one poll serves every pending ciphertext; later partial waves
only step in when the earlier ones have stopped landing partials; and the
combine (discrete-log search plus proof) runs off the service loop, one at
a time per node, yielding to a contribution or finalization in progress.
Every adversarial scenario of the battery is rejected as designed: tampered,
replayed and relayed shares; ciphertexts outside the window, over the cap,
off-curve, in a small-order subgroup, undecryptable or copied from another
application; late registration, duplicate claims, contributions and
partials; finalizing early; aborting a healthy epoch.

Two costs are deliberate. An undecryptable ciphertext makes each member of
the combine rotation run one bounded 2^50 baby-step giant-step search (about
20 s of CPU and a 256 MB table built once per process), after which the node
ignores the rest of that application's ciphertexts for the epoch, so an
attacker pays one search per registration rather than per ciphertext. A
withheld organizer share parks its slot in every node at no recurring cost
until a share event wakes it.

## Groth16 versus PLONK (universal setup)

Measured on 2026-09-03 with the single-key circuits that preceded pool keys
(the Finalize row is the old finalization circuit); the backend comparison
is unchanged in kind.

Measured on 2026-09-03 with gnark v0.16.3 / gnark-crypto v0.21.0 at
`MaxN = 32` (`go test ./circuits/... -bench BenchmarkBackends -benchtime=1x`,
Ryzen 9 9950X3D, 32 threads, average of 3 proofs; PLONK with an unsafe test
structured reference string (SRS), which does not change prover cost).
On-chain gas is the deployed Groth16 combine verifier on a real proof versus
gnark's exported PLONK verifier for the same statement (11 public inputs),
both under `via_ir`, `optimizer_runs = 1`.

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
default; a Groth16 multi-party computation (MPC) ceremony for the four
circuits is the cheaper way to remove the single-party setup.

The same run showed that upgrading gnark from the pinned v0.14 snapshot to
v0.16.3 changes the compiled R1CS (partial +19 %, combine +11 %,
contribution +5 %, finalize unchanged) and therefore the pinned artifact
hashes. That upgrade has since been made — the snapshot's `ScalarMul` was
unsound, see the caveat at the top — and ships with the pool-key circuit
release; the constraint, proving-time and gas tables above predate it.

## How to reproduce

* Constraints: `go run ./cmd/constraints` after setting `MaxN` in
  `circuits/common/sizes.go` (and `MAX_N` in `solidity/src/libraries/Sizes.sol`
  for the contracts). The current pipeline is `make circuits` (compile all four
  circuits — Contribution, Finalize, PartialDecrypt, DecryptCombine — regenerate
  the Solidity verifiers, patch `config/circuit_artifacts.go` and rebuild).
* Proving times (contribution + finalize are the two expensive ones):
  `DAVINCI_ARTIFACTS_DIR=/tmp/bench-$N go test ./circuits/contribution -run XXX -bench '^BenchmarkProve$' -benchtime=5x`
  and `DAVINCI_ARTIFACTS_DIR=/tmp/bench-$N go test ./circuits/finalize -run XXX -bench '^BenchmarkProve$' -benchtime=5x`
  on an idle host; the first run performs the setup. (The full suite is
  `go test ./circuits/... -run XXX -bench '^BenchmarkProve$' -benchtime=5x`.)
* Gas: run the flow on any deployment (`make testnet-up`, or the Sepolia
  public testnet with `--network sepolia`) with `cmd/dkgapp`
  (`register`, `encrypt`, `reveal`, `plaintext`)
  and read `gasUsed` from the receipts, e.g. `cast receipt --json <tx> | jq
  .gasUsed`. `forge test --gas-report` gives the same per-function figures
  against mock verifiers, i.e. without the ~250k pairing check.
