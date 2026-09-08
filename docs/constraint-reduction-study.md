# Constraint reduction study (v4 circuits)

Status: study, 2026-09-08, nothing implemented. Numbers are measured on `main`
(gnark v0.16.3, `MaxK = 16`, `MaxN = 32`) with `go run ./cmd/circuit-profile`
and with micro-circuits compiled for each gadget; a second, independent
source-level budget by the protocol-architect review agrees with them within 1%.

## Where the constraints go

| Circuit | Constraints | Dominant terms |
|---|---:|---|
| Contribution | 5,904,167 | 512 shares × 7.9 k; 512 coefficient commitments × 3.0 k; 32 ECDH × 3.8 k |
| Finalize | 2,328,130 | re-hashing 32 dealers × 16 keys × 32 commitment points (1.40 M); 512 Horner evaluations (0.60 M) |
| DecryptCombine | 287,338 | 34 fixed-base and 33 variable-base multiplications, Lagrange checks |
| PartialDecrypt | 29,026 | five variable-base and three fixed-base multiplications |

Unit costs (BN254 R1CS, measured):

| Gadget | Constraints | Notes |
|---|---:|---|
| `FixedBaseMul` (4-bit windows, nested `Lookup2`) | 2,342 | includes a 254-bit `ToBinary`; ≈ 570 of it is zero-window select handling |
| fixed-base, 2-bit windows, bilinear select | 1,393 | one `Mul` per window shared by both coordinates; verified against gnark-crypto |
| `CommitmentPolynomialValue`, 32 coefficients, variable 6-bit index | 2,295 | 31 × `ScalarMulSmallScalar` + Add |
| same with a constant index | 1,024 at `x = 7`; 1,164 averaged over `x = 1..32` | what `finalize` already pays |
| `ScalarMulVar` (254-bit double-and-add) | 3,800 | |
| range check `x ≤ r − 1` | 672 | a 251-bit decomposition with the comparison on its bits would cost ≈ 502 |
| `ReduceToSubgroupOrder` + `AddModSubgroupOrder` | 2,034 | on top of the share range check |
| `ShareMaskHash` (two Poseidon calls) | 553 | |
| Poseidon of two inputs | 241 | Poseidon costs ≈ 43 constraints per absorbed element in `MultiHash` |
| `AssertPointOnCurve` | 4 | |

Contribution budget from those units (matches the profiler within 1%):

| Region | Count | Per item | Total |
|---|---:|---:|---:|
| Coefficient commitments `C = a·G` (fixed-base + range check + on-curve) | 512 | 3,022 | 1.55 M |
| Ephemerals `R_i = r_i·G` | 32 | 2,344 | 0.08 M |
| ECDH `r_i·PK_i` (`ScalarMulVar`, already one per recipient for all 16 keys) | 32 | 3,800 | 0.12 M |
| Per share: range 672 + Horner 2,295 + `s·G` 2,342 + mask hash 553 + mod-r mask 2,034 + selects | 512 | 7,904 | 4.05 M |
| Digests, BRLC fold, masks | | | 0.10 M |

Finalize: 61% re-derives every accepted dealer's Poseidon commitment digest from
witness points (32 × 16 × 64 field elements), 26% is the constant-index Horner
over the aggregated commitments, 4% aggregation additions, 3% on-curve checks.
It contains no fixed-base or secret-scalar multiplication and no keccak: the
Merkle roots are computed by the contract after the proof.

Two facts frame every idea below. Runtime masking does not remove compiled
constraints: an epoch with `n = 4, t = 3` proves the circuit compiled for
`N = 32, K = 16`. And Groth16 verification gas is independent of circuit size
as long as the public-input count and proof shape stay: a smaller circuit buys
proving time, memory and artifact size, not gas.

## Ideas, measured or modelled

"Circuit-local" means only the Go circuits, witness builders and regenerated
verifier change (a circuit release either way); "protocol" means the node, SDK
and vectors change too; "contract" means Solidity beyond the verifier.

### I1. Stop proving `C_{j,m} = a_{j,m}·G` for m ≥ 1, certify the subgroup instead (−1.49 M, 25%)

Feldman consistency is already proven per share: `s_{j,i}·G = Σ_{m<t} (i+1)^m
C_{j,m}` for every active recipient, `n ≥ t` of them at the committee
positions `1..n` the contract pins. Any `t` consistent shares interpolate the
polynomial, so knowledge of the witness implies knowledge of the coefficients
and the 496 explicit fixed-base proofs add nothing to extraction. Bare removal
is not enough, though: with the order-two point `T = (0, p−1)`, the pair
`C_1 = a_1·G − T`, `C_2 = a_2·G + T` passes every Feldman check (`x·(−T) +
x²·T = x(x−1)·T = O` for every integer `x`) while `C_1, C_2` have no discrete
logarithm to `G`. No contract, finalize, node or SDK path was found that turns
this into a wrong key or share, but it contradicts the relation the paper
states (`05-succinct.tex`, `docs/pool-keys.md` §witness). Close it with a
cofactor-preimage certificate per higher commitment: a private on-curve point
`Q_m` with `8·Q_m = C̃_m` (three constrained doublings; the image of
multiplication by 8 is exactly the prime-order subgroup), 21 constraints each,
10,416 in total. With every commitment in `⟨G⟩` the interpolation argument
gives the coefficient discrete logarithms and the dealing statement survives
unchanged. Keep the 16 constant-term proofs as a conservative choice (with
consecutive positions the binomial identity `F(0) = Σ_{i≤t} (−1)^{i+1} C(t,i)
F(i)` already forces `C_0` torsion-free). Circuit-local.

### I2. Mask shares in the native field (−1.04 M, 18%)

Today `masked = s + (H mod r) mod r`, which costs a quotient witness, two extra
range checks and a carry per share. Use `masked = s + H mod p`, `p` the BN254
scalar field the circuit computes in: one constraint. `H` is a Poseidon output,
uniform in `F_p`, so the masked word is a one-time pad over `F_p`; the recipient
computes `s = masked − H mod p` and gets the range-checked `s < r`. The word is
still a canonical field element, so BRLC and the contract are untouched.
Protocol change in the meaning of the word: an old decoder subtracting mod `r`
gets a wrong share silently, so the release must refuse mixed versions
(domain bump). Reached independently by both reviewers.

### I5. 2-bit bilinear fixed-base multiplication (−0.53 M, 9%)

`X = c00 + b0·(c10−c00) + b1·(c01−c00) + b0·b1·(c11−c10−c01+c00)` selects among
four constants with one multiplication shared by both coordinates; 127 windows
plus 126 complete additions, no zero-window selects. 1,393 against 2,342,
verified to agree with gnark-crypto (including 0 and r−1) and to reject a wrong
point. The architect's 3-bit shared-monomial variant models at ≈ 1,349; not
worth the extra table. Also helps `partialdecrypt` and `decryptcombine`.
Circuit-local.

### I6. Constant recipient index (−0.58 M, 10%)

The committee snapshot already fixes member `i`'s index as `i + 1`, and the
contract checks the transcript's index words against it. Evaluate Horner at
the constant `i + 1` and assert the witness index equals it when the slot is
active: 1,164 on average instead of 2,295 per share, transcript unchanged.
Circuit-local.
(Removing the index words from the transcript would save `n` calldata words
but changes `Sizes.sol` and the snapshot hash; not worth it.)

### I4. One bit decomposition per scalar (−0.36 M with I1, −0.98 M without)

A share is decomposed to bits twice, by its range check and again inside its
fixed-base multiplication, and the range check runs against a 254-bit
constant although `r` has 251 bits. One 251-bit decomposition with the
comparison on those bits (≈ 502) feeding the fixed-base gadget saves ≈ 678
per scalar; the same holds for the coefficients while they are still proven,
and for the nonce shared by `R_i` and the ECDH product. Circuit-local; keep
canonicality (`s < r`), not just congruence.

### I3. Per-recipient mask seed (−0.14 M, 2%)

`seed_i = H(domain, roundHash, indexes, S_i.x, S_i.y)` once per recipient,
`mask_{j,i} = H(seed_i, j)`: 32 × 553 + 512 × 241 instead of 512 × 553.
Protocol change; bundle with I2's domain bump.

### I8. Bounded scalar Feldman check instead of curve arithmetic (architect's route, −2.0 to −2.1 M)

Keep `C = a·G` for every coefficient and prove `s + q·r = Σ_{m<t} a_m (i+1)^m`
over the integers with bounded limbs (radix 2^80, four limbs per scalar, every
product is witness-limb × constant so it is linear; a 156-bit quotient
suffices; signed carries range-checked), instead of the point Horner and
`s·G`. Removes 2.37 M and adds an estimated 0.25 to 0.40 M. Not a native-field
equation with a free quotient, which would repeat the fake-GLV class of error.
Circuit-local, transcript-identical. Mutually exclusive with I1 (I8 needs the
coefficient proofs I1 removes). Introduces the one new non-native integer
gadget of this study and therefore the most review.

### I7. Separate threshold cap `MaxT = 17` (−0.25 M contribution, −0.97 M finalize)

`MaxCoefficients = MaxN = 32` today, so every Horner runs 31 steps and every
commitment digest absorbs 32 points although `t = n/2 + 1 = 17` at `n = 32`.
A separate `MaxT` halves both, and takes finalize below the `2^21` FFT
boundary (2,097,152), which roughly halves its proving memory and time beyond
the constraint count. Product decision: a policy with `t > MaxT` (a 2/3
threshold at full committee would need 22) becomes impossible. Changes
`sizes.go`, `Sizes.sol`, SDK, UI, vectors.

### F1. Constant addition chains in finalize (−0.10 M, 4%)

`ScalarMulSmallScalar` builds `2P, 3P` and walks six-bit windows even when the
scalar is the compile-time constant `1..32`. Fixed double/add chains per
constant take the Horner region from 596 k to about 496 k. Circuit-local.

### F2. Flatten the per-dealer digest hierarchy (−0.08 to −0.17 M finalize)

One recursive `MultiHash` over a dealer's 1,027 words instead of 16 key
digests plus an outer hash (≈ 42.2 k vs 44.8 k per dealer), or a fixed-width
sponge for long vectors. Protocol change to the digest definitions in both
circuits, node and SDK; only worth it if it is what crosses `2^21` when I7 is
rejected.

### Not worth doing

- Batching Feldman checks with the BRLC challenge: the right-hand side needs
  variable-base multiplications by the challenge powers (3.8 k each), more
  than the 512 fixed-base products it would replace.
- Compressed points in the finalize digests: hashing `y` costs 43
  constraints, extracting its sign costs 255, and dropping the sign lets a
  griefing finalizer flip points and store an unusable key.
- Sharing one ephemeral across all recipients (≈ 73 k), shorter ECDH nonces,
  XOR or point-valued share encodings, Poseidon2 for its own sake.
- Removing the finalize transcript digest (53 k): docs/pool-keys.md explains
  the forgery it prevents.
- Reducing `MaxK`: same design cost per usable application, twice the epochs.
  Reducing `MaxN` to 16 (contribution ≈ 2.65 M, finalize ≈ 0.60 M) is a
  product decision about committee size, not an optimization.
- Splitting the dealing into two proofs of 8 keys: same constraints, half the
  peak memory, twice the gas.

## Combined effect

| Set | Contribution | Finalize | Scope |
|---|---:|---:|---|
| today | 5.90 M | 2.33 M | |
| Route A: I1 with certificates + I2 + I3 + I4 + I5 + I6 | ≈ 1.8 M (−69%) | 2.33 M | circuit + node + SDK; contract untouched beyond the verifier |
| Route B: I8 + I2 + I3 + I4 + I5 + I6 | ≈ 1.8 M (−69%) | 2.33 M | same scope, one new integer gadget |
| circuit-local only (I1, I4, I5, I6) | ≈ 2.9 M (−51%) | 2.33 M | no node/SDK change |
| Route A + F1 constant chains | ≈ 1.7 M | 2.23 M | circuit-local additions |
| Route A + F1 + I7 (`MaxT = 17`, 34-word digests) | ≈ 1.45 M (−75%) | ≈ 1.36 M (−42%, below 2^21) | `Sizes.sol`, policy cap, digest definition |
| Route A + F1 + I7 keeping the 64-word digest with a constant identity tail | ≈ 1.45 M | ≈ 1.55 M | `Sizes.sol`, policy cap only |

Proving cost follows the constraint count: the contribution proving key would
fall from 800 MB to roughly 250 MB, the proof from 3.5 s to about 1 s on the
benchmark host, and the node's peak from 7 GB to roughly 2.5 to 3 GB (FFT
domains are rounded, so measure rather than divide). Gas does not move.

## Verdict (both reviewers, 2026-09-08)

**Route A with subgroup certificates and constant addition chains**, no new
integer-limb machinery. It keeps the Feldman equations that are already
exercised end to end and adds one small standard gadget; Route B's soundness
would rest on every limb, carry and quotient bound of a new relation.

`MaxT = 17` is an acceptable deployment profile rather than a free
optimization: the adaptive policy already picks `t = 17, m_min = 22` at
`n = 32`, and the cap is what takes finalize below the `2^21` boundary. It
trades reconstruction threshold for availability (at `n = 32`: 16 vs 21
colluders tolerated, 15 vs 10 absent decryptors tolerated) and forbids a 2/3
threshold at full committee. If adopted, `createEpoch` must reject
`threshold > MAX_T` so nobody can create an epoch the circuits cannot serve;
dealer, recipient, combine and Merkle capacities stay at 32.

Sequencing that keeps each release reviewable:

1. Circuit-local: I5 bilinear fixed-base, I4 shared bits, I6 constant index,
   I1 with certificates, F1 chains. ≈ 2.8 M contribution, 2.23 M finalize;
   transcript, node and SDK untouched; a circuit release.
2. Protocol: I2 native-field masks with I3 seeds under a bumped domain, and
   the `MaxT` cap with 34-word digests if the threshold restriction is
   accepted. ≈ 1.45 M and 1.36 M.

Adversarial regressions the implementation must carry: the torsion pair above
fails with certificates; off-curve or wrong preimages fail even when the
witness hints are overridden; wrong shares, wrong constant commitments and
substituted recipient indexes fail; `t = 1`, `t = n`, `n = 32` with masked
inactive slots; the fixed-base gadget rejects non-boolean bits and out-of-range
scalars, not only wrong points. Native evaluators reduce powers mod `r` while
the circuit multiplies by the integer index; they agree today because
`32^31 < r`, which must be re-checked if any bound grows.

Side finding (fixed separately): `crypto/group.IsOnCurve` delegated to a
decoder that never checked the curve equation, so `(0, 0)` passed as
on-curve; the circuits were unaffected because they check the equation
themselves.
