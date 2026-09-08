# Pool keys: per-application committee-held keys

Status: implementation spec. Every layout below is normative; Go, Solidity,
the TypeScript SDK and the explorer must agree bit for bit. Constants:
`MaxK = 16` pool keys per epoch, `MaxN = 32` committee members, `MaxT = 32`
threshold, `MerkleDepth = 5`; BRLC domains `davinci-dkg:contribution:v2`,
`davinci-dkg:finalize:v2`, `davinci-dkg:decrypt-combine:v1`; share-encryption
KDF domain `davinci-dkg/share-encryption/v2`.

## Why

If every application shared one epoch key, the committee's partials `d_i·C1`
would be the same for every application: anyone able to register an
application could copy any `C1` into it and learn `sk_ep·C1`, a decryption
oracle across applications. Dealing `MaxK` independent keys per epoch and
binding every application to its own key closes it: a ciphertext copied from
one application decrypts under an unrelated secret and yields garbage.

## Constants

| name | Go | Solidity | value |
|---|---|---|---|
| committee bound | `ccommon.MaxN` | `MAX_N` | 32 |
| pool size (keys per epoch) | `ccommon.MaxK` | `MAX_K` | 16 |
| Merkle depth | `ccommon.MerkleDepth` | `MERKLE_DEPTH` | 5 (= log2 MaxN) |

Every epoch deals exactly `MaxK` keys; unused keys cost only calldata.
`MaxN` must be a power of two (the share-commitment Merkle tree has
`2^MerkleDepth = MaxN` leaves), so the supported committee bounds are 16, 32
and 64, not 48. The existing `uint8` pool-key indexes accommodate keys
`0…15`, `poolNext = 16`, and the application index-plus-one markers.
Activation bitmaps are gone: `finalizeEpoch` stores every key and its
share-commitment root at once, so a `Live` epoch needs no activation step.

## Keys and modes

- `P_j = Σ_d A_{d,j,0}` over the accepted contributors `d`, `j ∈ [0, MaxK)`.
- Application key: `PK_aid = P_j` (automatic) or `PK_aid = P_j + PK_org`
  (organizer-locked). `j` is claimed at registration, one key per application.
- Automatic: no organizer key. `organizerPK` is stored as the identity
  `(0, 1)`, `organizerSecret = 0`.
- Organizer-locked: `PK_org` with the Schnorr proof of possession as today.
  The organizer later calls `revealOrganizerSecret` once; from then on the
  committee combines by itself. There is no per-ciphertext organizer share.
  Until the reveal, `requireDecryptionOpen` reverts
  `OrganizerSecretNotRevealed`, so **no partial and no combine of a locked
  application exists on chain before it**: the organizer learns every result
  together with everyone else and only decides *when* the application opens,
  never *which* ciphertexts — the contract enforces that, not node policy.
- Decryption window: `decryptNotBefore` / `decryptNotAfter` (unix seconds,
  0 = unbounded). Submission (`submitCiphertext`) is gated by the block window
  (`notBeforeBlock` / `notAfterBlock`), the submitter policy and
  `decryptNotAfter`; `decryptNotBefore` and, for a locked application, the
  reveal gate partials and combines (`requireDecryptionOpen`). The window
  bounds what the contract accepts and what honest nodes post; it does not
  bind `t` colluding members, who hold shares and can compute partials off
  chain whenever they like (for a locked application they still lack
  `sk_org` until the reveal). Because the contract refuses partials before
  `decryptNotBefore` and before the reveal, there are none on chain to
  collect from before the window either.

## Epoch flow

1. `createEpoch` — unchanged, plus: allowed before the cadence only when the
   newest epoch is `Live` with at most one unclaimed key (`poolNext >= MAX_K - 1`),
   or `Aborted`. An epoch that is still selecting its committee or assembling
   its keys cannot be pre-empted.
2. `claimSlot` — unchanged.
3. `submitContribution` — one proof dealing all `MaxK` polynomials (below),
   now with a compact transcript of `K·(2t+n) + 5n` words (no padding in calldata).
4. `finalizeEpoch(eid, transcriptDigest, transcript, proof, input)` — **one
   proof activates the whole pool**: it verifies a single `circuits/finalize`
   Groth16 proof, stores every pool key and the Merkle root of the whole
   committee's share commitments for each key, and flips the epoch `Live`.
   Emits `EpochLive(eid, contributionCount)`. Freezes the accepted
   contributor set.
5. `registerApplication` claims the next unclaimed key; a `Live` epoch serves
   at most `MaxK` (16) applications before `registerApplication` reverts
   `PoolExhausted` until the next epoch has gone through its preparation
   window (pool exhaustion).

Nodes race on a deterministic finalize stagger (seed-derived, anchored at
`liveNotBeforeBlock`): the node whose slot is due reconstructs the accepted
contributions, proves the batched finalization and submits `finalizeEpoch`;
the first proof to land makes the epoch `Live` with every key and share root
stored, and the rest see `AlreadyLive` and stop. Nodes create the next
epoch early when the newest epoch has at most one unclaimed key
(`poolNext >= MAX_K - 1`) or is `Aborted`, bypassing the normal cadence.

Two consequences of a fixed pool are accepted for now:

- **Pool exhaustion.** An epoch serves at most `MaxK` (16) applications.
  Once its keys are claimed, registrations revert `PoolExhausted` until the
  next epoch is `Live`, which takes a full preparation window (committee
  selection, key assembly, finalize gap, one finalization proof). A burst of
  more than sixteen registrations in one epoch waits for the next epoch.
- **Registration-driven epoch amplification.** Registration is
  permissionless, so anyone registering fifteen automatic applications (no
  organizer key needed) drives `poolNext` to `MAX_K - 1` (15) and lets the
  next epoch be created before the cadence — every committee member then
  pays a contribution again. The cost is bounded (one extra epoch per fifteen
  registrations, each of which pays gas) but it is an amplification. A
  registration fee or an allow-list on `registerApplication` is future work.

## Contribution proof (`circuits/contribution`)

Public inputs remain exactly eight, in the existing order:
`[eid, threshold, committeeSize, contributorIndex, commitmentsHash, encryptedSharesHash, challenge, transcriptCommitment]`

Require `1 ≤ t ≤ n ≤ N` and `1 ≤ contributorIndex ≤ n`. The contract binds
these values to epoch policy and sender membership. `K` is fixed by the
circuit/verifier and contract release, not caller-selected.

Private witness: `Commitments[MaxK][MaxT]`, `ConstantTerms[MaxK]`,
`CommitmentPreimages[MaxK][MaxT−1]`, `RecipientIndexes[MaxN]`,
`RecipientPubKeys[MaxN]`, `EncryptionNonces[MaxN]`, `Ephemerals[MaxN]`,
`Shares[MaxK][MaxN]`, `MaskedShares[MaxK][MaxN]` (`MaxT = MaxN = 32` today,
`sizes.go` / `Sizes.sol`; `createEpoch` rejects `t > MAX_T`).

Only the constant term `a_{j,0}` of each polynomial is a witness, range-checked
to `[0, r)` and proven as the discrete logarithm of `C_{j,0}`, which pins the
pool-key contribution to the prime subgroup. The higher coefficients are not:
the Feldman checks below hold for every active recipient, `n ≥ t` of them at
the committee positions `1..n`, and any `t` consistent shares interpolate the
polynomial, so knowledge of the shares is knowledge of the coefficients. For
that argument to speak about discrete logarithms every commitment must lie in
the prime subgroup, which is what `CommitmentPreimages` certify: for `m ≥ 1`
the circuit checks `Q_{j,m}` is on the curve and `8·Q_{j,m}` equals the
(threshold-masked) commitment, three constrained doublings, the image of
multiplication by the cofactor being exactly `⟨G⟩`. Without it the pair
`C_1 + T, C_2 + T` with `T = (0, p−1)` of order two passes every Feldman check
(`x·T + x²·T = x(x+1)·T = O`). An honest dealer sets `Q = [8⁻¹ mod r]·C`
(`CofactorPreimageNative`) and the identity for inactive slots.

Recipient slot `i` of the committee snapshot is member `i+1`; the circuit
asserts `RecipientIndexes[i] = i+1` for active slots and evaluates the Feldman
polynomial at that constant. Per active share it checks `s_{j,i} < r`,
`s_{j,i}·G = Σ_{m<t} (i+1)^m C_{j,m}` and the encryption below.

One ephemeral / ECDH secret per recipient, shared by all `MaxK` keys. Per
recipient `i` and key `j` (`p` the BN254 scalar field, the circuit's native
field):

```
seed_i        = H(domain, eid, contributorIndex<<16 | recipientIndex_i, shared_i.x, shared_i.y)
mask[j][i]    = H(seed_i, j)
masked[j][i]  = s[j][i] + mask[j][i]  (mod p)
```

`domain = "davinci-dkg/share-encryption/v2"`. The mask is a Poseidon output,
uniform in `F_p`, so the masked word is a one-time pad over `F_p`; the
recipient computes `s = masked − mask (mod p)` and rejects `s ≥ r`
(`crypto/shareenc` mirrors both hashes: `ShareMaskSeed`, `ShareMask`). The
word is a canonical field element like every other transcript word. No
reduction to `r` happens anywhere.

Digests (Poseidon `MultiHash`, every input masked exactly as today):

```
keyDigest[j]    = MultiHash(A[j][0].x, A[j][0].y, …, A[j][MaxN-1].x, A[j][MaxN-1].y)   // identity (0,1) for m >= t
commitmentsHash = MultiHash(eid, contributorIndex, threshold, keyDigest[0], …, keyDigest[MaxK-1])
rowDigest[i]    = MultiHash(idx_i, pk_i.x, pk_i.y, eph_i.x, eph_i.y, ms[0][i], …, ms[MaxK-1][i])   // zeros/(0,1) when i >= n
encryptedSharesHash = MultiHash(eid, contributorIndex, committeeSize, rowDigest[0], …, rowDigest[MaxN-1])
```

Compact transcript: `L_C = K·(2t+n) + 5n` words. Require
`transcript.length == 32·L_C`. No padding travels in calldata; the length is
a function of the epoch's public policy `(t, n)`, not of the circuit bounds:

```
[0, 2Kt)          commitments, key-major: for j in [0, K-1], then m in [0, t-1]: A[j][m].x, A[j][m].y
[2Kt, 2Kt+n)      recipientIndexes           (word i MUST equal i+1)
[2Kt+n, 2Kt+3n)   recipientPubKeys (x, y)
[2Kt+3n, 2Kt+5n)  ephemerals (x, y)
[2Kt+5n, L_C)     maskedShares, key-major: for j in [0, K-1], then i in [0, n-1]: ms[j][i]
```

The BRLC fold uses a gate `b = [entry is active]` derived from the bounded
public counts: initialize `(acc, power, count) = (0, ρ, 0)` and for each
candidate word `v`:

```
acc   ← acc + b·power·v
power ← power·(1 + b·(ρ−1))
count ← count + b
```

Require `count == L_C` and `acc == transcriptCommitment`; an inactive entry
neither contributes nor advances the exponent. The contract streams exactly
`L_C` calldata words using the canonical `BRLC.commitCalldata`, yielding
`Σ ρ^(q+1)·w[q]`. The challenge anchor is unchanged:
`keccak256(commitmentsHash ‖ encryptedSharesHash ‖ keccak256(transcript))` —
the same Fiat–Shamir discipline as finalization and combine, so every
proof-carrying call anchors its challenge on the prover's digests *and* the
calldata.

## Batched finalization proof (`circuits/finalize`)

`FinalizeCircuit` proves, over the accepted contributors listed in the
transcript, that each contributor's on-chain `commitmentsHash` is reproduced
from its commitments for **every** key, that the aggregates
`Ā[j][m] = Σ_{d<a} A[d][j][m]` (identity for `m >= t`), the pool keys
`P[j] = Ā[j][0]`, and the share commitments `D[j][i] = Σ_m (i+1)^m·Ā[j][m]`
(for `i < n`; identity for the rest), are the values the contract stores.

Public inputs, in order (7):
`[eid, threshold, committeeSize, acceptedCount, transcriptDigest, challenge, transcriptCommitment]`

Private witness: `ParticipantIndexes[MaxN]`, `ContributionHashes[MaxN]`,
`Commitments[MaxN][MaxK][MaxN]` (indexed dealer/key/coefficient),
`AggregateCommitments[MaxK][MaxN]`, `ShareCommitments[MaxK][MaxN]`.

Require `1 ≤ t ≤ a ≤ n ≤ N`. For each active dealer row `d < a`: its
participant index is in `[1, n]` and unique, names an accepted contributor,
and recomputes that dealer's outer `commitmentsHash` **once** from all `K`
key digests (the per-key digest absorbs the padded vectors — inactive
scalars zero, inactive points `(0, 1)`). Inactive rows contribute identity / zero everywhere;
exactly `a` unique accepted rows prevent omitted dealers.

The fixed finalization transcript has `L_F = 2N + K·(2+2N)` words (= 1,120
at N = 32, K = 16):

```
[0, N)        participantIndexes         (0 for rows >= a)
[N, 2N)        contributionHashes        (0 for rows >= a)
then for key j in [0, K-1], a (2+2N)-word row:
    P[j].x, P[j].y, D[j][0].x, D[j][0].y, …, D[j][N-1].x, D[j][N-1].y
                                   (D[j][i] = (0,1) for i >= n)
```

Accepted dealer rows may appear in any order, matching the contract's frozen
contributor set; builders SHOULD emit ascending indexes. Let `H` be the
existing Poseidon `MultiHash`. Over these exact masked words:

```
R   = H(0, I[0], …, I[N−1], h[0], …, h[N−1])
B_j = H(1, j, P[j].x, P[j].y, D[j][0].x, D[j][0].y, …)
T   = H(2, eid, t, n, a, K, L_F, R, B_0, …, B_(K−1))
```

Tags `0, 1, 2` are field integers. Require `transcriptDigest == T`, the
ordinary BRLC commitment over all `L_F` words, and the challenge anchor
`keccak256(transcriptDigest ‖ keccak256(transcript))` — the same anchor
discipline as contribution and combine: with `keccak(keccak(transcript))`
alone the challenge would depend on
the calldata only, and a permissionless finalizer could grind a calldata
transcript carrying a forged `P_j` that still verifies.

Contract `finalizeEpoch(eid, transcriptDigest, transcript, proof, input)`
checks, in order:

1. Direct-call gate: `msg.sender == tx.origin && msg.sender.code.length == 0`
   (`DirectCallRequired`).
2. Epoch exists (`InvalidEpoch`), is not already `Live` (`AlreadyLive`), and
   is in `KeyAssembly` (`InvalidPhase`).
3. `block.number >= liveNotBeforeBlock` (the positive finalize gap:
   contributions remain accepted through their deadline; finalization happens
   after it), `acceptedCount >= minValidContributions`, and the count bounds
   `1 ≤ t ≤ a ≤ n ≤ N` (`InvalidProofInput`).
4. Transcript length `32·L_F`, input length `224` (7 × 32 bytes), proof
   length `256`; decode the 7 canonical public inputs and bind positions
   `0…4` to state and the digest argument.
5. Derive the challenge from
   `keccak256(transcriptDigest ‖ keccak256(transcript))` with the
   `davinci-dkg:finalize:v2` domain and check it against input position 5;
   canonical-stream the BRLC over all `L_F` calldata words and compare with
   input position 6.
6. Row validation: every active row names a distinct accepted contributor
   under its index and carries that contributor's stored
   `commitmentsHash`; exactly `a` unique rows prevent omitted dealers.
7. Inactive share slots hold the identity `(0,1)`; verify the pinned
   `circuits/finalize` Groth16 proof.
8. Compute every Merkle root, store every key and root, then set the epoch
   `Live`. Reverts must leave no partial state.

Emits `EpochLive(eid, acceptedCount)`.

## Partial decryption

Circuit unchanged. `submitPartialDecryption` gains a trailing
`bytes32[] calldata shareProof` (length `MERKLE_DEPTH`, siblings bottom-up).
The contract checks the Merkle path of `keccak(0x00 ‖ pi[6] ‖ pi[7])` at leaf
index `participantIndex - 1` against
`poolShareRoots[eid][appPoolIndex[eid][aid]]` — the root that `finalizeEpoch`
stored for the application's claimed key — and `requireDecryptionOpen`
(which also reverts `OrganizerSecretNotRevealed` for a locked application
whose organizer has not revealed). Any committee member — contributor or
not — may post. `epochShareCommitmentHashes` is removed.

Nodes park a locked application's slots *before* posting any partial and
wake them on `OrganizerSecretRevealed`, rescanning the application's
ciphertexts from its registration block so nothing submitted while parked is
missed.

## Combine

Circuit `decryptcombine`: the organizer DLEQ is replaced by knowledge of the
organizer secret. Public inputs, in order (9):
`[eid, aid, ctIdx, threshold, shareCount, combineHash, plaintext, challenge, transcriptCommitment]`

Private: `CiphertextC1`, `CiphertextC2`, `Plaintext`, `OrganizerPK`,
`OrganizerSecret`, indexes, partials, Lagrange coefficients. Constraints:
`OrganizerSecret < r_bjj`, `OrganizerPK == OrganizerSecret·G`,
`Δ = OrganizerSecret·C1`, `C2 == m·G + Σ λ_k δ_k + Δ`; the Lagrange
coefficients are pinned to the **canonical** vector at 0 of the qualifying
set — the unique solution of the Vandermonde system over the first
`shareCount` indexes, masked by `shareCount` — so a prover cannot substitute
a different valid-looking interpolation. Automatic applications use
`OrganizerPK = (0,1)`, `OrganizerSecret = 0`.

```
combineHash = MultiHash(eid, aid, ctIdx, threshold, shareCount, C1.x, C1.y, C2.x, C2.y, PK_org.x, PK_org.y, [idx_k, δ_k.x, δ_k.y]…)
transcript  = [C1.x, C1.y, C2.x, C2.y, PK_org.x, PK_org.y, idx[0..N), (δ.x, δ.y)[0..N)]   // COMBINE_TRANSCRIPT_WORDS = 6 + 3·MaxN
```

Contract: `_verifyOrganizerWords`, the stored share hash, the `e`
recomputation and `DOMAIN_ORGANIZER_SHARE_V1` are removed. `w[4..5]` must
equal the application's `organizerPK`; `requireDecryptionOpen` applies
(window and, for a locked application, the reveal). Challenge anchor
`keccak(combineHash ‖ plaintext ‖ keccak(transcript))` unchanged. On every
proof-carrying call the BRLC commitment over calldata refuses a
non-canonical word (a word `>= p`, the BN254 scalar field), so a transcript
has exactly one encoding and the values the contract reads straight from
calldata (the pool key, the share commitments, the partials) are canonical
field elements.

## App manager

```
enum AppMode { OrganizerLocked, Automatic }
struct AppPolicy {
  AppMode   mode;
  bool      openSubmission;
  address[] submitters;        // <= 32, exclusive; empty = registrant only
  uint16    maxCiphertexts;
  uint64    notBeforeBlock;    // submitCiphertext window, blocks
  uint64    notAfterBlock;
  uint64    decryptNotBefore;  // unix seconds, 0 = none
  uint64    decryptNotAfter;   // unix seconds, 0 = none
}
struct Application {
  address creator; Point organizerPK; uint256 organizerSecret; uint8 poolIndex;
  AppPolicy policy; uint64 createdAtBlock; bool exists;
}
registerApplication(eid, aid, policy, pkOrgX, pkOrgY, schnorrAx, schnorrAy, schnorrZ)
revealOrganizerSecret(eid, aid, sk)   // permissionless; locked apps only; once; checks sk·G == organizerPK
getApplication / getOrganizerPK / requireDecryptionOpen / requireCanSubmitCiphertext / getRegisteredAids
event ApplicationRegistered(eid, aid, creator, pkX, pkY, mode, poolIndex)
event OrganizerSecretRevealed(eid, aid, sk)
errors: PoolExhausted, InvalidOrganizerSecret, InvalidPolicy, DecryptionClosed, DecryptionNotOpen, OrganizerSecretNotRevealed, AlreadyRevealed
```

Automatic registration ignores the key and Schnorr arguments and stores
`(0, 1)`. Locked registration verifies the Schnorr PoP as today. Registration
calls `IDKGManager(MANAGER).claimPoolKey(eid, aid)` (only the app manager may
call it) which assigns the next unclaimed key index, increments the pool
cursor, preserves the index-plus-one application-key marker, and emits
`PoolKeyClaimed`. It reverts `PoolExhausted` when all `MAX_K` keys are taken;
because `finalizeEpoch` proves and stores the whole pool atomically, every
unclaimed key of a `Live` epoch is usable — there is no activation state.

`requireDecryptionOpen(eid, aid)` (consulted by `submitPartialDecryption`
and `combineDecryption`) reverts `DecryptionNotOpen` before
`decryptNotBefore`, `DecryptionClosed` after `decryptNotAfter`, and
`OrganizerSecretNotRevealed` for an organizer-locked application whose
`organizerSecret` is still `0`. `requireCanSubmitCiphertext` gates
submission by the block window, the submitter policy, `maxCiphertexts` and
`decryptNotAfter` only — a ciphertext may be submitted before decryption
opens.

Manager views: `getPoolKey(eid, j) → (x, y)` (requires `Live` and `j < MAX_K`,
else `InvalidProofInput`), `getPoolStatus(eid) → (nextIndex)` (the `poolNext`
cursor), `getPoolShareRoot(eid, j)`, `getAppPoolIndex(eid, aid)`.

## Protocol constants and vectors

`internal/protocol/protocol.go` is the source of truth for the Fiat–Shamir
domain strings, mirrored by `solidity/src/libraries/DKGProtocol.sol` and
`sdk/src/protocol.ts`: `davinci-dkg:contribution:v2`, `davinci-dkg:finalize:v2`
and `davinci-dkg:decrypt-combine:v1`. `cmd/protocol-vectors` writes
`tests/vectors/*.json` (protocol constants, a compact contribution transcript,
a finalization transcript) from the Go side; the SDK tests and the Foundry
tests assert against them, and CI fails if `make vectors` changes them.

## Circuit toolchain

The circuits are compiled with gnark v0.16.3 / gnark-crypto v0.21.0
(`go.mod`) and require gnark ≥ v0.16.2: older releases have an unsound
variable-base twisted-Edwards `ScalarMul` (a hinted fake-GLV decomposition
that a prover can satisfy for any output point). No circuit uses a hinted
scalar multiplication: every variable-base product is
`ccommon.ScalarMulVarBits`, a plain double-and-add over constrained bits, and
every fixed-base product is `ccommon.FixedBaseMulBits`. `make circuits`
recompiles the four circuits, runs the Groth16 setup, rewrites the Solidity
verifying keys and pins the artifact hashes in `config/circuit_artifacts.go`;
a node never compiles a circuit, it stream-verifies the release artifacts
against those hashes. Constraint counts, proving times, memory and gas are in
`BENCHMARKS.md`.
