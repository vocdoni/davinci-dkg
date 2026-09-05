# Pool keys v4: per-application committee-held keys

Status: implementation spec, v4 (2026-09-05). Every layout below is normative;
Go, Solidity, the TypeScript SDK and the explorer must agree bit for bit.
v4 changes, relative to v3.1: an atomic, proof-carrying `finalizeEpoch` that
activates all `MaxK = 16` keys at once (replacing the proof-less finalize plus
per-key `activatePoolKey` proofs), a compact contribution transcript of
`K·(2t+n) + 5n` words (no padding in calldata), BRLC domains
`davinci-dkg:contribution:v2` and `davinci-dkg:finalize:v2`, and the deletion
of every activation-bitmap / activation-state field.

## Why

With one epoch key, the committee's partials `d_i·C1` are the same for every
application, so anyone who can register an application can copy any `C1`
into it and learn `sk_ep·C1`. The organizer key masked that for
organizer-locked applications; an automatic application that published its
organizer secret had no mask at all (see README, "treat an automatic
ciphertext as public").

Fix: every application gets its own committee-held key `P_j`, dealt in the
epoch's DKG as one of `MaxK` polynomials. Partials are now `e_{j,i}·C1` with
`e_{j,i}` the member's share of `P_j`, so a ciphertext copied into another
application yields a value under a different key and is useless. The
organizer key becomes optional and, when present, is revealed once.

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

Private witness: `Coefficients[MaxK][MaxN]`, `Commitments[MaxK][MaxN]`,
`RecipientIndexes[MaxN]`, `RecipientPubKeys[MaxN]`, `EncryptionNonces[MaxN]`,
`Ephemerals[MaxN]`, `Shares[MaxK][MaxN]`, `MaskedShares[MaxK][MaxN]`,
`MaskQuotients[MaxK][MaxN]`, `ShareMasks[MaxK][MaxN]`,
`MaskedShareCarries[MaxK][MaxN]`.

One ephemeral / ECDH secret per recipient, shared by all `MaxK` keys. Per key
`j` and recipient `i`:

```
rawMask[j][i] = ShareMaskHash(eid, contributorIndex, recipientIndex_i, shared_i.x, shared_i.y, j)
```

`ShareMaskHash` gains the trailing key index as a Poseidon input of the meta
hash: `meta = H(domain, eid, contributorIndex<<16 | recipientIndex, keyIndex)`,
`raw = H(meta, shared.x, shared.y)`. `crypto/shareenc` mirrors it
(`EncryptShare*` / `DecryptShare*` take `keyIndex uint8`).

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

## Batched finalization proof (`circuits/finalize`, replaces `circuits/poolkey`)

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
scalars zero, inactive points `(0, 1)`); there is no `OtherKeyDigests`
shortcut any more. Inactive rows contribute identity / zero everywhere;
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
discipline as contribution and combine (see "Why the digest is in the anchor",
v3.1): with `keccak(keccak(transcript))` alone the challenge would depend on
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
else `InvalidProofInput`), `getPoolStatus(eid) → (nextIndex)` (no activation
bitmap — `poolNext` is all the status there is), `getPoolShareRoot(eid, j)`,
`getAppPoolIndex(eid, aid)`. Removed: `getCollectivePublicKey`,
`submitOrganizerShare`, `getOrganizerShareHash`, `OrganizerShareSubmitted`,
`PoolKeyActivated`, `PoolKeyAlreadyActive`.

## Protocol constants and vectors

`internal/protocol/protocol.go`, `DKGProtocol.sol`, `sdk/src/protocol.ts`:
remove `DOMAIN_ORGANIZER_SHARE_V1`; the BRLC transcript domain strings are
`davinci-dkg:contribution:v2`, `davinci-dkg:finalize:v2` (replaces
`davinci-dkg:poolkey:v1`) and `davinci-dkg:decrypt-combine:v1` (unchanged).
`tests/vectors/*.json` are regenerated; the organizer-share DLEQ vectors
disappear.

## Circuit toolchain

Every circuit is compiled with gnark v0.16.3 / gnark-crypto v0.21.0
(`go.mod`). Every gnark release up to v0.15.0, and the snapshot pinned
before (`v0.14.1-0.20260126…`), has an unsound variable-base twisted-Edwards
`ScalarMul`: the fake-GLV decomposition check `s1 + s2·s = k·order` is
evaluated in the native field with the quotient `k` a free hint output, so a
prover can make the gadget return any point (reproduced with verifying
Groth16 proofs; fixed upstream in gnark v0.16.0, PR #1765, without an
advisory — the weaker cofactor-torsion offset is IACR ePrint 2026/1776).
Never downgrade gnark below v0.16.2. No circuit uses the hinted gadget any
more: every variable-base multiplication is `ccommon.ScalarMulVar`, a
hint-free double-and-add. The v4 changes (batched finalization circuit,
compact contribution transcript, `MaxK = 16`) change every compiled R1CS, so
all four circuits — contribution, finalize, partialdecrypt, decryptcombine —
are recompiled and their hashes re-pinned in `config/circuit_artifacts.go`
with this release.

## Costs to measure

Per contribution, per finalization, per partial (with path), per combine, per
registration, per reveal; circuit constraints and proving times for all four
circuits at MaxN = 32, MaxK = 16.

### v4 estimates (superseded by measurements)

The numbers below were **estimates** made before v4 was compiled; the measured
values are in BENCHMARKS.md ("v4: batched finalization and compact contributions":
contribution 5,904,167 constraints / 3.54 s / 9.3 GB peak, finalize 2,328,130 / 2.26 s /
5.5 GB, finalize 2.20–2.81 M gas, contribution 0.50–1.31 M gas at n = 4…32). Kept for the record:

- Finalization circuit: expect approximately **2.6–2.9 M R1CS constraints**,
  budgeting 3.1 M pending compilation (estimate).
- At `t = n = 32` the contribution calldata grows to ~55,108 ABI bytes;
  check signed-transaction / RPC limits (Geth's usual 128 KiB transaction-pool
  limit is not a consensus guarantee).
- Estimated gas (Cancun, ±25% uncertainty):

  | Call | (n, t) | transcript / full ABI bytes | estimated gas |
  |---|---|---:|---:|
  | Finalize | (4,3) | 35,840 / 36,580 | ~2.1 M |
  | Finalize | (32,22) | 35,840 / 36,580 | ~2.8 M |
  | Contribution | (4,3) | 5,760 / 6,596 | ~0.56 M |
  | Contribution | (32,22) | 44,032 / 44,868 | ~1.28 M |

  Finalization includes 48 cold key/root writes (~1.06 M gas), one verifier,
  16 Merkle trees and one dealer-validation pass. On EIP-7623 chains the
  large contribution (44,032-byte transcript) rises to about **1.77 M**
  (`max(normalGas, 21000 + 10·(zeroBytes + 4·nonzeroBytes))`).

### Measured (this worktree)

`go run ./cmd/constraints` (MaxN = 32, MaxK = 16, gnark v0.16.3):

| Circuit        | constraints |
|-----------------|------------:|
| Contribution    |   5,904,167 |
| Finalize        |   2,328,130 |
| PartialDecrypt  |      29,026 |
| DecryptCombine  |     287,338 |

The v4 contribution circuit is 5,904,167 constraints at `MaxK = 16` (vs. 3,060,692
in v3.1 at `MaxK = 8`): 3.54 s and 9.3 GB peak resident memory per proof on the
benchmark machine, against 1.57 s and 5.0 GB. The `Finalize` circuit is
2,328,130 constraints, 2.26 s and 5.5 GB, below the 2.6–2.9 M estimate above.

