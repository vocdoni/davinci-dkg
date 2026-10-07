# Pool keys: encodings and proof statements

The normative layouts of the proof-carrying calls. Go, Solidity, the TypeScript SDK and the
explorer must agree on every one of them bit for bit; `tests/vectors/*.json` pins them (see
[Vectors](#protocol-constants-and-vectors)). The protocol itself is described in
[protocol.md](protocol.md).

Constants: `MaxK = 16` pool keys per epoch, `MaxN = 32` committee members, `MaxT = 32`
threshold, `MerkleDepth = 5`. BRLC domains `davinci-dkg:contribution:v2`,
`davinci-dkg:finalize:v2` and `davinci-dkg:decrypt-combine:v1`; share-encryption KDF domain
`davinci-dkg/share-encryption/v2`. `p` is the BN254 scalar field (the circuits' native field),
`r` the BabyJubJub subgroup order, `H` Poseidon `MultiHash`.

## Why one key per application

If every application shared one epoch key, the committee's partials `d_i·C1` would be the same
for every application: anyone able to register an application could copy any `C1` into it and
learn `sk_ep·C1`, a decryption oracle across applications. Dealing `MaxK` independent keys per
epoch and binding each application to its own key closes it: a ciphertext copied from one
application decrypts under an unrelated secret.

## Sizes

| Name | Go | Solidity | Value |
|---|---|---|---|
| Committee bound | `ccommon.MaxN` | `MAX_N` | 32 |
| Pool size | `ccommon.MaxK` | `MAX_K` | 16 |
| Merkle depth | `ccommon.MerkleDepth` | `MERKLE_DEPTH` | 5 (`log2 MaxN`) |

Every epoch deals exactly `MaxK` keys; unused keys cost only calldata. `MaxN` must be a power of
two because the share-commitment tree has `2^MerkleDepth = MaxN` leaves. Pool-key indexes are
`uint8`: keys `0…15`, a spent cursor `poolNext = 16`, and index-plus-one application markers.

## Keys and modes

- `P_j = Σ_d A_{d,j,0}` over the accepted contributors `d`, for `j ∈ [0, MaxK)`.
- Application key: `PK_aid = P_j` (automatic) or `PK_aid = P_j + PK_org` (organizer-locked),
  with `j` claimed at registration.
- Automatic: `organizerPK` is stored as the identity `(0, 1)` and `organizerSecret = 0`.
- Organizer-locked: `PK_org` with a Schnorr proof of possession. `requireDecryptionOpen` reverts
  `OrganizerSecretNotRevealed` until `revealOrganizerSecret`, so no partial and no combine of a
  locked application exists on chain before the reveal.
- Decryption window: `decryptNotBefore` / `decryptNotAfter` (unix seconds, 0 = unbounded) gate
  partials and combines. Submission is gated by `notBeforeBlock` / `notAfterBlock`, the
  submitter policy, `maxCiphertexts` and `decryptNotAfter`.

## Epoch flow

1. `createEpoch`: cadence-gated, also allowed when the newest epoch is `Live` with
   `poolNext >= MAX_K − 1`, or `Aborted`. An epoch still selecting its committee or assembling
   its keys cannot be pre-empted.
2. `claimSlot`: the lottery.
3. `submitContribution`: one proof dealing all `MaxK` polynomials over a compact transcript of
   `K·(2t+n) + 5n` words.
4. `finalizeEpoch(eid, transcriptDigest, transcript, proof, input)`: one `circuits/finalize`
   proof; stores every pool key and the Merkle root of the whole committee's share commitments
   for each key, freezes the accepted contributor set, sets the epoch `Live` and emits
   `EpochLive(eid, acceptedCount)`.
5. `registerApplication` claims the next unclaimed key and reverts `PoolExhausted` once all
   `MaxK` are claimed.

Nodes race on a seed-derived finalize stagger anchored at `liveNotBeforeBlock`; the first proof
to land makes the epoch `Live` and the others see `AlreadyLive`.

## Contribution proof (`circuits/contribution`)

Public inputs, in order (8):
`[eid, threshold, committeeSize, contributorIndex, commitmentsHash, encryptedSharesHash, challenge, transcriptCommitment]`

Require `1 ≤ t ≤ n ≤ N` and `1 ≤ contributorIndex ≤ n`. The contract binds these values to the
epoch policy and the sender's committee position. `K` is fixed by the circuit, verifier and
contract release.

Private witness: `Commitments[MaxK][MaxT]`, `ConstantTerms[MaxK]`,
`CommitmentPreimages[MaxK][MaxT−1]`, `RecipientIndexes[MaxN]`, `RecipientPubKeys[MaxN]`,
`EncryptionNonces[MaxN]`, `Ephemerals[MaxN]`, `Shares[MaxK][MaxN]`, `MaskedShares[MaxK][MaxN]`.

Only the constant term `a_{j,0}` of each polynomial is a witness, range-checked to `[0, r)` and
proven as the discrete logarithm of `C_{j,0}`, which puts the pool-key contribution in the prime
subgroup. The higher coefficients are not: the Feldman checks below hold for every active
recipient, `n ≥ t` of them at positions `1..n`, and any `t` consistent shares interpolate the
polynomial, so knowledge of the shares is knowledge of the coefficients. For that argument to
hold every commitment must lie in the prime subgroup, which `CommitmentPreimages` certify: for
`m ≥ 1` the circuit checks that `Q_{j,m}` is on the curve and that `8·Q_{j,m}` equals the
commitment (three constrained doublings; the image of multiplication by the cofactor is exactly
`⟨G⟩`). Without it the pair `C_1 + T, C_2 + T` with `T = (0, p−1)` of order two passes every
Feldman check, since `x·T + x²·T = x(x+1)·T = O`. An honest dealer sets `Q = [8⁻¹ mod r]·C`
(`CofactorPreimageNative`) and the identity for inactive slots.

Recipient slot `i` is committee member `i+1`: the circuit asserts `RecipientIndexes[i] = i+1` for
active slots and evaluates the Feldman polynomial at that constant. For each active share it
checks `s_{j,i} < r`, `s_{j,i}·G = Σ_{m<t} (i+1)^m C_{j,m}`, and the encryption below.

One ephemeral key and ECDH secret per recipient, shared by all `MaxK` keys. For recipient `i` and
key `j`:

```
seed_i        = H(domain, eid, contributorIndex<<16 | recipientIndex_i, shared_i.x, shared_i.y)
mask[j][i]    = H(seed_i, j)
masked[j][i]  = s[j][i] + mask[j][i]  (mod p)
```

`domain = "davinci-dkg/share-encryption/v2"`. The mask is a Poseidon output, uniform in `F_p`, so
the masked word is a one-time pad over `F_p`. The recipient computes `s = masked − mask (mod p)`
and rejects `s ≥ r` (`crypto/shareenc`: `ShareMaskSeed`, `ShareMask`). No reduction modulo `r`
happens anywhere.

Digests:

```
keyDigest[j]        = H(A[j][0].x, A[j][0].y, …, A[j][MaxN-1].x, A[j][MaxN-1].y)   // identity (0,1) for m >= t
commitmentsHash     = H(eid, contributorIndex, threshold, keyDigest[0], …, keyDigest[MaxK-1])
rowDigest[i]        = H(idx_i, pk_i.x, pk_i.y, eph_i.x, eph_i.y, ms[0][i], …, ms[MaxK-1][i])   // zeros/(0,1) when i >= n
encryptedSharesHash = H(eid, contributorIndex, committeeSize, rowDigest[0], …, rowDigest[MaxN-1])
```

Transcript: `L_C = K·(2t+n) + 5n` words and `transcript.length == 32·L_C`. No padding travels in
calldata; the length is a function of the epoch's public `(t, n)`:

```
[0, 2Kt)          commitments, key-major: for j in [0, K-1], then m in [0, t-1]: A[j][m].x, A[j][m].y
[2Kt, 2Kt+n)      recipientIndexes           (word i MUST equal i+1)
[2Kt+n, 2Kt+3n)   recipientPubKeys (x, y)
[2Kt+3n, 2Kt+5n)  ephemerals (x, y)
[2Kt+5n, L_C)     maskedShares, key-major: for j in [0, K-1], then i in [0, n-1]: ms[j][i]
```

The circuit's BRLC fold uses a gate `b = [entry is active]` derived from the public counts.
Starting from `(acc, power, count) = (0, ρ, 0)`, for each candidate word `v`:

```
acc   ← acc + b·power·v
power ← power·(1 + b·(ρ−1))
count ← count + b
```

and requires `count == L_C` and `acc == transcriptCommitment`; an inactive entry neither
contributes nor advances the exponent. The contract streams exactly `L_C` calldata words through
`BRLC.commitCalldata`, yielding `Σ ρ^(q+1)·w[q]`. The challenge anchor is
`keccak256(commitmentsHash ‖ encryptedSharesHash ‖ keccak256(transcript))`, so the challenge
depends on both the prover's digests and the calldata.

## Finalization proof (`circuits/finalize`)

`FinalizeCircuit` proves, over the accepted contributors listed in the transcript, that each
contributor's stored `commitmentsHash` is reproduced from its commitments for every key, and that
the aggregates `Ā[j][m] = Σ_{d<a} A[d][j][m]` (identity for `m >= t`), the pool keys
`P[j] = Ā[j][0]` and the share commitments `D[j][i] = Σ_m (i+1)^m·Ā[j][m]` (for `i < n`, identity
otherwise) are the values the contract stores.

Public inputs, in order (7):
`[eid, threshold, committeeSize, acceptedCount, transcriptDigest, challenge, transcriptCommitment]`

Private witness: `ParticipantIndexes[MaxN]`, `ContributionHashes[MaxN]`,
`Commitments[MaxN][MaxK][MaxN]` (dealer, key, coefficient), `AggregateCommitments[MaxK][MaxN]`,
`ShareCommitments[MaxK][MaxN]`.

Require `1 ≤ t ≤ a ≤ n ≤ N`. Each active dealer row `d < a` has a participant index in `[1, n]`,
unique, naming an accepted contributor, and recomputes that dealer's `commitmentsHash` once from
all `K` key digests (inactive scalars zero, inactive points `(0, 1)`). Inactive rows contribute the
identity or zero everywhere; exactly `a` unique accepted rows rule out omitted dealers.

The transcript has a fixed `L_F = 2N + K·(2+2N)` words (1,120 at `N = 32`, `K = 16`):

```
[0, N)        participantIndexes         (0 for rows >= a)
[N, 2N)       contributionHashes         (0 for rows >= a)
then for key j in [0, K-1], a (2+2N)-word row:
    P[j].x, P[j].y, D[j][0].x, D[j][0].y, …, D[j][N-1].x, D[j][N-1].y
                                   (D[j][i] = (0,1) for i >= n)
```

Accepted dealer rows may appear in any order; builders should emit ascending indexes. Over these
exact words:

```
R   = H(0, I[0], …, I[N−1], h[0], …, h[N−1])
B_j = H(1, j, P[j].x, P[j].y, D[j][0].x, D[j][0].y, …)
T   = H(2, eid, t, n, a, K, L_F, R, B_0, …, B_(K−1))
```

Tags `0, 1, 2` are field integers. Require `transcriptDigest == T`, the plain BRLC commitment over
all `L_F` words, and the challenge anchor `keccak256(transcriptDigest ‖ keccak256(transcript))`.
The digest must stay in the anchor: with `keccak(keccak(transcript))` alone the challenge would
not depend on the witness, the calldata region the contract reads `P_j` from would be bound to
the proof only by the linear BRLC relation, and a finalizer could search for a calldata transcript
carrying a forged `P_j` that still verifies.

`finalizeEpoch` checks, in order:

1. Direct call: `msg.sender == tx.origin && msg.sender.code.length == 0` (`DirectCallRequired`).
2. The epoch exists (`InvalidEpoch`), is not `Live` (`AlreadyLive`) and is in `KeyAssembly`
   (`InvalidPhase`).
3. `block.number >= liveNotBeforeBlock`, `acceptedCount >= minValidContributions`, and
   `1 ≤ t ≤ a ≤ n ≤ N` (`InvalidProofInput`).
4. Transcript length `32·L_F`, input length 224 bytes, proof length 256 bytes; the seven public
   inputs are canonical and positions `0…4` match the state and the digest argument.
5. The challenge derived from `keccak256(transcriptDigest ‖ keccak256(transcript))` under
   `davinci-dkg:finalize:v2` equals input 5, and the BRLC over all `L_F` calldata words equals
   input 6.
6. Every active row names a distinct accepted contributor and carries its stored
   `commitmentsHash`.
7. Inactive share slots hold `(0, 1)`, and the Groth16 proof verifies.
8. Every Merkle root is computed and every key and root stored, then the epoch becomes `Live`. A
   revert leaves no partial state.

## Share-commitment tree

For each key `j` the tree has `MaxN` leaves, leaf `i` being member `i+1`'s share commitment
`D[j][i]`:

```
leaf  = keccak256(0x00 ‖ D.x ‖ D.y)
empty = keccak256("davinci-dkg:merkle-empty:v1")
node  = keccak256(0x01 ‖ left ‖ right)
```

`finalizeEpoch` stores the root as `poolShareRoots[eid][j]`. Every member is a leaf, not only the
contributors, because a member that did not contribute still received a share from every accepted
dealer.

## Partial decryption

`submitPartialDecryption` carries a trailing `bytes32[] shareProof` of `MERKLE_DEPTH` siblings,
bottom-up. The contract checks the path of `keccak256(0x00 ‖ pi[6] ‖ pi[7])` at leaf index
`participantIndex − 1` against `poolShareRoots[eid][appPoolIndex[eid][aid]]`, and calls
`requireDecryptionOpen`. Any committee member may post, contributor or not.

## Combine (`circuits/decryptcombine`)

Public inputs, in order (9):
`[eid, aid, ctIdx, threshold, shareCount, combineHash, plaintext, challenge, transcriptCommitment]`

Private witness: `CiphertextC1`, `CiphertextC2`, `Plaintext`, `OrganizerPK`, `OrganizerSecret`,
indexes, partials and Lagrange coefficients. Constraints: `OrganizerSecret < r`,
`OrganizerPK == OrganizerSecret·G`, `Δ = OrganizerSecret·C1` and `C2 == m·G + Σ λ_k δ_k + Δ`. The
Lagrange coefficients are pinned to the canonical vector at 0 of the qualifying set, the unique
solution of the Vandermonde system over the first `shareCount` indexes, so a prover cannot
substitute another interpolation. Automatic applications use `OrganizerPK = (0, 1)` and
`OrganizerSecret = 0`.

```
combineHash = H(eid, aid, ctIdx, threshold, shareCount, C1.x, C1.y, C2.x, C2.y, PK_org.x, PK_org.y, [idx_k, δ_k.x, δ_k.y]…)
transcript  = [C1.x, C1.y, C2.x, C2.y, PK_org.x, PK_org.y, idx[0..N), (δ.x, δ.y)[0..N)]   // COMBINE_TRANSCRIPT_WORDS = 6 + 3·MaxN
```

The contract requires transcript words 4 and 5 to equal the application's `organizerPK`, applies
`requireDecryptionOpen`, and derives the challenge from
`keccak256(combineHash ‖ plaintext ‖ keccak256(transcript))`.

On every proof-carrying call the calldata BRLC rejects a word `>= p`, so a transcript has exactly
one encoding and the values the contract reads straight from calldata (pool keys, share
commitments, partials) are canonical field elements.

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
revealOrganizerSecret(eid, aid, sk)   // permissionless; locked applications only; once; checks sk·G == organizerPK
getApplication / getApplicationKey / getOrganizerPK / requireDecryptionOpen / requireCanSubmitCiphertext / getRegisteredAids
event ApplicationRegistered(eid, aid, creator, pkX, pkY, mode, poolIndex)
event OrganizerSecretRevealed(eid, aid, sk)
errors: PoolExhausted, InvalidOrganizerSecret, InvalidPolicy, DecryptionClosed, DecryptionNotOpen, OrganizerSecretNotRevealed, AlreadyRevealed
```

`aid` must be non-zero, below the BN254 scalar field and carry the registrant in its low 160
bits (`aid = salt << 160 | msg.sender`, `salt < 2^92`), else registration reverts
`InvalidApplication`. Automatic registration ignores the key and Schnorr arguments and stores
`(0, 1)`; locked registration verifies the Schnorr proof of possession (domain
`davinci-dkg:organizer-register:v1`). Registration calls `DKGManager.claimPoolKey(eid, aid)`,
callable only by the app manager, which assigns the next unclaimed key, advances the cursor,
records the index-plus-one marker and emits `PoolKeyClaimed`.

`requireDecryptionOpen(eid, aid)` reverts `DecryptionNotOpen` before `decryptNotBefore`,
`DecryptionClosed` after `decryptNotAfter`, and `OrganizerSecretNotRevealed` for a locked
application whose `organizerSecret` is still 0. `requireCanSubmitCiphertext` gates submission by
the block window, the submitter policy, `maxCiphertexts` and `decryptNotAfter` only, so a
ciphertext may be submitted before decryption opens.

Manager views: `getPoolKey(eid, j)` (requires `Live` and `j < MAX_K`), `getPoolStatus(eid)` (the
`poolNext` cursor), `getPoolShareRoot(eid, j)`, `getAppPoolIndex(eid, aid)`.

## Protocol constants and vectors

`internal/protocol/protocol.go` is the source of the Fiat–Shamir domain strings, mirrored by
`solidity/src/libraries/DKGProtocol.sol` and `sdk/src/protocol.ts`. `cmd/protocol-vectors`
writes `tests/vectors/*.json` (protocol constants, compact contribution transcripts, finalization
transcripts) from the Go code; the SDK, Foundry and explorer tests assert against them, and
`make vectors-check` fails when they drift.

## Circuit toolchain

The circuits are compiled with gnark v0.16.3 and gnark-crypto v0.21.0. gnark releases through
v0.15.0 have an unsound variable-base twisted-Edwards `ScalarMul`, a hinted decomposition that a
prover can satisfy for any output point (fixed in v0.16.0); the project requires v0.16.2 or later
and never downgrades. No circuit uses a hinted scalar
multiplication: variable-base products go through `ccommon.ScalarMulVarBits`, a double-and-add
over constrained bits, and fixed-base products through `ccommon.FixedBaseMulBits`.
