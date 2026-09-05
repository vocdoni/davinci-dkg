# Pool keys: per-application committee-held keys

Status: implementation spec, v3.1 (2026-09-04). Every layout below is normative;
Go, Solidity, the TypeScript SDK and the explorer must agree bit for bit. v3.1
changes, relative to the first pool-key cut: a prover digest in the activation
proof's Fiat–Shamir anchor, share commitments for the whole committee with
domain-separated Merkle leaves, a contract-enforced organizer-locked gate, a
narrower early-epoch rule, a canonical Lagrange vector in the combine, and a
gnark upgrade that recompiles every circuit (see "Circuit toolchain").

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
| pool size (keys per epoch) | `ccommon.MaxK` | `MAX_K` | 8 |
| Merkle depth | `ccommon.MerkleDepth` | `MERKLE_DEPTH` | 5 (= log2 MaxN) |

Every epoch deals exactly `MaxK` keys; unused keys cost only calldata.
`MaxN` must be a power of two (the share-commitment Merkle tree has
`2^MerkleDepth = MaxN` leaves), so the supported committee bounds are 16, 32
and 64, not 48.

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
   newest epoch is `Live` with at most one unclaimed key
   (`poolNext >= MAX_K - 1`), or `Aborted`. An epoch that is still selecting
   its committee or assembling its keys cannot be pre-empted.
2. `claimSlot` — unchanged.
3. `submitContribution` — one proof dealing `MaxK` polynomials (below).
4. `finalizeEpoch(eid)` — **no proof**. Requires KeyAssembly,
   `block.number >= liveNotBeforeBlock`, `contributionCount >= minValidContributions`.
   Sets `Live`, emits `EpochLive(eid, contributionCount)`. Freezes the
   accepted contributor set.
5. `activatePoolKey(eid, j, transcriptDigest, transcript, proof, input)` —
   permissionless, one per key, any order, only while `Live`. Stores `P_j`
   and the Merkle root of the whole committee's share commitments for key
   `j`. Emits `PoolKeyActivated(eid, j, x, y)`.
6. `registerApplication` claims the next unused **activated** key.

Nodes activate key 0 as soon as the epoch is Live and keep
`ActivateAhead` (default 2) activated-but-unclaimed keys available, in the
same seed-derived rotation as the old auto-finalize. Nodes create the next
epoch early when the newest epoch has fewer than `ActivateAhead` unclaimed
keys and the contract allows it.

Two consequences of a fixed pool are accepted for now:

- **Pool exhaustion.** An epoch serves at most `MaxK` (8) applications.
  Once its keys are claimed, registrations revert `PoolExhausted` until the
  next epoch is `Live` and has activated a key, which takes a full
  preparation window (committee selection, key assembly, finalize gap) plus
  one activation. Nodes start that refill early (above), but a burst of more
  than eight registrations in one epoch waits.
- **Registration-driven epoch amplification.** Registration is
  permissionless, so anyone registering seven automatic applications (no
  organizer key needed) drives `poolNext` to `MAX_K - 1` and lets the next
  epoch be created before the cadence — every committee member then pays a
  contribution again. The cost is bounded (one extra epoch per seven
  registrations, each of which pays gas) but it is an amplification. A
  registration fee or an allow-list on `registerApplication` is future work.

## Contribution proof (`circuits/contribution`)

Public inputs, in order (8, unchanged count):
`[eid, threshold, committeeSize, contributorIndex, commitmentsHash, encryptedSharesHash, challenge, transcriptCommitment]`

Private witness: `Coefficients[MaxK][MaxN]`, `Commitments[MaxK][MaxN]`,
`RecipientIndexes[MaxN]`, `RecipientPubKeys[MaxN]`, `EncryptionNonces[MaxN]`,
`Ephemerals[MaxN]`, `Shares[MaxK][MaxN]`, `MaskedShares[MaxK][MaxN]`,
`MaskQuotients[MaxK][MaxN]`, `ShareMasks[MaxK][MaxN]`, `MaskedShareCarries[MaxK][MaxN]`.

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

Transcript words (`CONTRIB_TRANSCRIPT_WORDS = 3·MaxK·MaxN + 5·MaxN`):

```
[0, 2KN)          commitments, key-major: for j, for m: A[j][m].x, A[j][m].y
[2KN, 2KN+N)      recipientIndexes
[2KN+N, 2KN+3N)   recipientPubKeys (x, y)
[2KN+3N, 2KN+5N)  ephemerals (x, y)
[2KN+5N, 3KN+5N)  maskedShares, key-major: for j, for i: ms[j][i]
```

Contract: signature unchanged. The committee prefix check covers
`[2KN·32, (2KN+3N)·32)` bytes. `ContributionRecord.commitmentsHash` stores the
Poseidon `commitmentsHash` (public input 4); the old keccak
`commitmentVectorDigest` is gone. BRLC challenge anchor unchanged:
`keccak(commitmentsHash ‖ encryptedSharesHash ‖ keccak(transcript))`.

## Pool-key activation proof (`circuits/poolkey`, replaces `circuits/finalize`)

Statement, for key `j`: over the accepted contributors listed in the
transcript, the contributor's on-chain `commitmentsHash` is reproduced from
its commitments for key `j` plus the digests of its other keys; the
aggregate `Ā[m] = Σ_i A_i[j][m]`; `P_j = Ā[0]`; `D_i = Σ_m idx_i^m · Ā[m]`.

Public inputs, in order (8):
`[eid, threshold, committeeSize, acceptedCount, keyIndex, transcriptDigest, challenge, transcriptCommitment]`

```
transcriptDigest = MultiHash(eid, keyIndex, w[0], …, w[6·MaxN − 1])   // the masked transcript words below
```

Private witness: `ParticipantIndexes[MaxN]`, `ContributionHashes[MaxN]`,
`KeyCommitments[MaxN][MaxN]` (contributor × coefficient, key `j` only),
`OtherKeyDigests[MaxN][MaxK]` (all `MaxK` digests of each contributor; slot
`j` is ignored and replaced by the recomputed one), `AggregateCommitments[MaxN]`,
`ShareCommitments[MaxN]`.

Constraints: `threshold <= acceptedCount <= committeeSize <= MaxN`,
`keyIndex < MaxK`; for each active contributor slot `i < acceptedCount`:
`recomputed = MultiHash(eid, idx_i, threshold, digests with slot j := MultiHash(KeyCommitments[i]…))`
equals `ContributionHashes[i]`; aggregate `Ā[m] = Σ_i A_i[j][m]` (identity
for `m >= threshold`); for every **committee position** `p = i + 1`,
`i < committeeSize`, the Vandermonde check `D_p = Σ_m p^m · Ā[m]` (identity
for `i >= committeeSize`); `transcriptDigest` equals the Poseidon `MultiHash`
over `eid`, `keyIndex` and the masked transcript words; inactive slots
contribute identity / zero everywhere.

Share commitments cover the **whole committee**, not only the contributors:
a member that claimed a slot but never contributed still received a share of
every accepted dealer's polynomial (each contribution encrypts to all `n`
recipients), so it holds `e_{j,p}` for every key and may post partials.
Decryption liveness is therefore `n − t` absent members, not `m − t`
(`m` = accepted contributions).

Transcript words (`POOLKEY_TRANSCRIPT_WORDS = 6·MaxN`):

```
[0, N)     participantIndexes            (0 when inactive)
[N, 2N)    contributionHashes            (0 when inactive)
[2N, 4N)   aggregateCommitments (x, y)   ((0,1) for m >= t)
[4N, 6N)   shareCommitments D_p (x, y)   (slot i holds committee position p = i + 1 for i < committeeSize; (0,1) for the rest)
```

BRLC domain `davinci-dkg:poolkey:v1` (replaces `davinci-dkg:finalize:v1`);
challenge anchor `keccak(transcriptDigest ‖ keccak(transcript))`, the same
shape as contribution (`keccak(digests ‖ keccak(transcript))`) and combine.

Why the digest is in the anchor (v3.1). The first cut anchored the challenge
on `keccak(keccak(transcript))` alone: the challenge depended on the calldata
only, never on the words inside the proof. The contract re-checks rows
`[0, 2N)` against on-chain state, but the aggregate and share commitments in
`[2N, 6N)` — including `P_j` itself, which the contract reads from calldata
and stores — are prover-chosen words bound to the proven witness only through
the BRLC equality `Σ c^k·w_k = Σ c^k·w'_k`. With `c` independent of the
witness words that is one linear relation over `4·MaxN` free words, and a
permissionless activator can grind calldata and witness transcripts that
satisfy it with a different `P_j` in the aggregate slot: a
generalized-birthday search whose cost is nowhere near the `2^128` the
transcript commitment is meant to provide, and whose result is a forged pool
key with a valid proof. Folding `transcriptDigest` into the anchor makes `c`
depend on the witness words through a collision-resistant hash: a calldata
transcript that differs from the proven one changes `c` and breaks the BRLC
equality unless the prover also finds a Poseidon collision. Activation now
follows the same discipline as contribution and combine — every
proof-carrying call anchors its challenge on prover digests *and* calldata.

Contract `activatePoolKey(eid, j, transcriptDigest, transcript, proof, input)`:
for `i < contributionCount`: index in `[1, committeeSize]`, no duplicates,
participant accepted with that index,
`contributionHashes[i] == epochContributions[participant].commitmentsHash`;
`input[5] == transcriptDigest`; `challenge` derived from
`keccak(transcriptDigest ‖ keccak(transcript))`. Verify proof and BRLC. Store
`poolKeys[eid][j] = aggregate[0]` and `poolShareRoots[eid][j] = root` where

```
leaf[p-1] = keccak256(0x00 ‖ D_p.x ‖ D_p.y)              for committee position p <= committeeSize
leaf[p-1] = keccak256("davinci-dkg:merkle-empty:v1")     for p > committeeSize
node      = keccak256(0x01 ‖ left ‖ right)               MERKLE_DEPTH (5) levels, MAX_N (32) leaves
```

Leaves and internal nodes are domain-separated by the prefix byte, and an
empty leaf is a fixed non-zero constant, so a leaf cannot be presented as a
node (or the other way round) and an absent position cannot be confused with
a commitment to the identity point.

## Partial decryption

Circuit unchanged. `submitPartialDecryption` gains a trailing
`bytes32[] calldata shareProof` (length `MERKLE_DEPTH`, siblings bottom-up).
The contract checks the Merkle path of `keccak(0x00 ‖ pi[6] ‖ pi[7])` at leaf
index `participantIndex - 1` against
`poolShareRoots[eid][appPoolIndex[eid][aid]]`, and `requireDecryptionOpen`
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
errors: PoolExhausted, PoolKeyNotActive, InvalidOrganizerSecret, InvalidPolicy, DecryptionClosed, DecryptionNotOpen, OrganizerSecretNotRevealed, AlreadyRevealed
```

Automatic registration ignores the key and Schnorr arguments and stores
`(0, 1)`. Locked registration verifies the Schnorr PoP as today. Registration
calls `IDKGManager(MANAGER).claimPoolKey(eid, aid)` (only the app manager may
call it) which returns the index, reverts `PoolExhausted` when none is left
and `PoolKeyNotActive` when the next key is not activated yet, and records
`appPoolIndex[eid][aid]`.

`requireDecryptionOpen(eid, aid)` (consulted by `submitPartialDecryption`
and `combineDecryption`) reverts `DecryptionNotOpen` before
`decryptNotBefore`, `DecryptionClosed` after `decryptNotAfter`, and
`OrganizerSecretNotRevealed` for an organizer-locked application whose
`organizerSecret` is still `0`. `requireCanSubmitCiphertext` gates
submission by the block window, the submitter policy, `maxCiphertexts` and
`decryptNotAfter` only — a ciphertext may be submitted before decryption
opens.

Manager views: `getPoolKey(eid, j) → (x, y)` (reverts `PoolKeyNotActive`),
`getPoolStatus(eid) → (nextIndex, activated bitmap)`, `getPoolShareRoot(eid, j)`,
`getAppPoolIndex(eid, aid)`. Removed: `getCollectivePublicKey`,
`submitOrganizerShare`, `getOrganizerShareHash`, `OrganizerShareSubmitted`.

## Protocol constants and vectors

`internal/protocol/protocol.go`, `DKGProtocol.sol`, `sdk/src/protocol.ts`:
remove `DOMAIN_ORGANIZER_SHARE_V1`; the BRLC transcript domain strings are
`davinci-dkg:contribution:v1` (unchanged), `davinci-dkg:poolkey:v1` (new,
replaces `davinci-dkg:finalize:v1`) and `davinci-dkg:decrypt-combine:v1`
(unchanged). `tests/vectors/*.json` are regenerated; the organizer-share
DLEQ vectors disappear.

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
hint-free double-and-add. The upgrade changes every compiled R1CS, so all
four circuits are recompiled and their hashes re-pinned in
`config/circuit_artifacts.go` with this release.

## Costs to measure

Per contribution, per activation, per partial (with path), per combine, per
registration, per reveal; circuit constraints and proving times for all four
circuits at MaxN = 32, MaxK = 8.
