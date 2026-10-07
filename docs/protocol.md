# Protocol

How DAVINCI DKG generates keys and decrypts. The byte-level encodings and proof statements are in
[pool-keys.md](pool-keys.md); gas and proving costs in [BENCHMARKS.md](../BENCHMARKS.md).

Keys, shares and ciphertexts live on BabyJubJub over the BN254 scalar field. Circuits hash with
Poseidon, contracts derive Fiat–Shamir challenges with keccak256, shares and ciphertexts are
ElGamal-encrypted, and every proof is Groth16 over BN254.

## Epochs

An epoch is one DKG run. Its `n` committee members jointly deal `MAX_K = 16` independent pool
keys `P_0 … P_15`, and any `t` of them can decrypt under any key. Epochs start at a fixed cadence
of `EPOCH_DURATION_BLOCKS`.

```
startBlock                                                                 endBlock
│ ─── Preparation (short, fixed) ─────►◄──────── Service (the rest) ───────│
│ CommitteeSelection │ KeyAssembly  │gap│              Live                 │
├────────────────────┼──────────────┼───┼───────────────────────────────────┤
│ claimSlot          │ submit-      │   │ registerApplication               │
│ (lottery)          │ Contribution │   │ submitCiphertext                  │
│                    │ (deals all   │   │ revealOrganizerSecret             │
│                    │  16 keys)    │   │ submitPartialDecryption           │
│                    │              │   │ combineDecryption                 │
└────────────────────┴──────────────┴───┴───────────────────────────────────┘
                                        ▲ finalizeEpoch: stores the 16 keys, sets Live
```

- **CommitteeSelection** (`COMMITTEE_SELECTION_BLOCKS`): the lottery fills `n` committee slots.
- **KeyAssembly** (`KEY_ASSEMBLY_BLOCKS`): each member submits one Feldman VSS contribution
  that deals all 16 polynomials at once, with a single Groth16 proof.
- **Finalize gap** (`FINALIZE_GAP_BLOCKS`), after which anyone may call `finalizeEpoch`. One
  proof over the accepted contributions reproduces every pool key and the Merkle root of the
  committee's share commitments for each key, and the contract stores them atomically and marks
  the epoch `Live`.

The window lengths are absolute block counts set at deployment, so a day-long epoch has the same
short preparation as a short one. A `Live` epoch keeps serving decryptions after its cadence
ends, while the next epoch prepares.

`createEpoch` is permissionless but cadence-gated: it reverts before `nextEpochStartBlock()`,
except when the newest epoch is `Live` with at most one unclaimed key or is `Aborted`. Nodes race
to call it after a random delay; the first call lands and the rest revert cheaply. The caller
picks `(t, n, minValidContributions, α)` within the deployment's floors (`MIN_THRESHOLD`,
`MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS`) and `t ≤ MAX_T`, and the contract requires
`minValidContributions ≥ t`. Nodes size the committee as three quarters of the active operators,
capped at 32, with a majority threshold.

An epoch serves at most 16 applications. Once its keys are claimed, `registerApplication` reverts
`PoolExhausted` until the next epoch is `Live`, one preparation window later. Because
registration is permissionless, anyone can claim fifteen keys and force the next epoch to start
early; the cost is bounded but it is an amplification, and a fee or allow-list is a deployment's
own choice.

Epoch phases: `None`, `CommitteeSelection`, `KeyAssembly`, `Live`, `Aborted` (`Completed` is
reserved).

## Committee selection

The lottery is replayable by anyone and has no coordinator:

- `R = registry.activeCount()` is snapshotted at `createEpoch`, `α = lotteryAlphaBps / 10000`;
- `seed = blockhash(startBlock + SEED_DELAY_BLOCKS)`, resolved by the first `claimSlot`;
- an operator is eligible iff `keccak256(seed ‖ address) < α · n · 2²⁵⁶ / R`.

Eligible operators claim first-come-first-served until `n` slots are filled, which snapshots the
committee and moves the epoch to `KeyAssembly`. Only operators registered before `createEpoch`
may claim, so fresh identities cannot be ground against a known seed.

An epoch whose committee does not fill in time, or whose key assembly closes with fewer than
`minValidContributions`, is dead: anyone may call `abortEpoch` once the deadline has passed. An
epoch that can still be finalized cannot be aborted. Nodes running with `--auto-create-epochs`
abort a dead newest epoch and create the next one, backing off when several epochs in a row
were aborted.

## Applications

A `Live` epoch hosts one encryption context per application, keyed by a 32-byte `aid` chosen by
the registrant inside its own namespace: `aid = salt << 160 | registrant`. The contract reverts
`InvalidApplication` unless the low 160 bits of `aid` are the caller, so nobody can register an
id in another account's namespace (Gnosis deployment; the Sepolia deployment `0xc73b…` predates
[#14](https://github.com/vocdoni/davinci-dkg/issues/14) and does not enforce this check), and an
integrator whose ids are predictable (a hash of a process id, say) cannot be front-run into
`ApplicationAlreadyExists`. `aid` is also a public
input of the decryption proofs, so it must be non-zero and below the BN254 scalar field; a salt
below `2^92` guarantees that. Registration itself is open to anyone, and ids are first come,
first served per epoch within a namespace.

`registerApplication` claims the next unclaimed pool key `P_j` and fixes the application's mode:

- **Organizer-locked** (default). The organizer publishes `PK_org = sk_org · G` with a Schnorr
  proof of possession and keeps `sk_org`. The application key is `PK_aid = P_j + PK_org`. The
  contract refuses every partial decryption and combine until the organizer calls
  `revealOrganizerSecret` once, for the whole application.
- **Automatic**. No organizer key: `PK_aid = P_j`. The committee decrypts as soon as `t` partials
  land inside the decryption window. Use it for values that are meant to be opened, such as a
  tally.

`DKGAppManager.getApplicationKey(epochId, aid)` returns `PK_aid`. Because each application has its own pool
key, a ciphertext copied from one application into another decrypts under an unrelated secret
and yields nothing useful.

For an organizer-locked application:

- **Losing `sk_org` makes the application permanently undecryptable.** Back it up at
  registration.
- Until the reveal, not even `t` colluding members can open a ciphertext on chain, and the
  organizer sees results no earlier than anyone else.
- The reveal is one-time and application-wide: every past and future ciphertext of the
  application becomes decryptable.
- Never reuse an organizer secret across applications: revealing it for one opens the other.

The policy is fixed at registration. `openSubmission` lets anyone submit; otherwise `submitters`
is an exclusive allow-list of up to 32 addresses; with neither, only the registrant submits.
`maxCiphertexts` caps the count, `notBeforeBlock` / `notAfterBlock` bound submission, and
`decryptNotBefore` / `decryptNotAfter` (unix seconds, `0` = unbounded) bound decryption.
Contradictory policies revert `InvalidPolicy()`.

## Threshold decryption

A ciphertext `(C₁, C₂)` under `PK_aid` is submitted with `submitCiphertext(epochId, aid, c1x,
c1y, c2x, c2y)`. The contract checks the submitter, the block window, the cap and
`decryptNotAfter`, checks the points are canonical, on the curve and not the identity, assigns
the next index and emits `CiphertextSubmitted`. It takes no proof of knowledge of the encryption
randomness: that is what allows homomorphic aggregation, since whoever submits an aggregated
tally cannot know its randomness.

1. Every committee member `i` watches the event and posts `δ_i = e_{j,i} · C₁`, where `e_{j,i}`
   is its share of `P_j`, with a Groth16 proof of the Chaum–Pedersen relation and a Merkle path
   of its share commitment against the root stored at finalization. Members that were selected
   but did not contribute also hold shares, so decryption tolerates `n − t` absent members.
   Nodes post in waves of `t`, so an honest ciphertext costs `t` partials rather than `n`.
2. Partials and combines are accepted only while `decryptNotBefore ≤ now ≤ decryptNotAfter`
   (`DecryptionNotOpen()`, `DecryptionClosed()`) and, for an organizer-locked application, after
   the reveal (`OrganizerSecretNotRevealed()`).
3. Once `t` partials are on chain, a node calls `combineDecryption` with a proof that the
   Lagrange interpolation of the partials is correct, that the caller knows the secret behind
   `PK_org` (zero for automatic applications) and that `m · G + Σ λ_k · δ_k + sk_org · C₁ = C₂`.
4. The plaintext `m` is stored on chain and read with `getPlaintext(epochId, aid, index)`.

The combine recovers `m` by baby-step giant-step: nodes recover plaintexts below 2⁵⁰, the SDK
below 2³². A plaintext above the cap is never combined.

The contract skips the prime-subgroup check on `C₁` to keep submission cheap; nodes perform it
before computing a partial, since a partial over a small-order `C₁` would leak a member's share
modulo the cofactor. A ciphertext whose plaintext is out of range marks its (application,
submitter) pair as tainted for the epoch, so one bad submitter cannot make the committee burn a
search on every ciphertext of an application.

**What the window guarantees.** The decryption window and the reveal bound what the contract
accepts and what honest nodes post. They do not bind `t` colluding members, who hold shares and
can compute partials off chain at any time. For an automatic application that is the whole
confidentiality assumption; for a locked one they still lack `sk_org`.

## Proofs

Four gnark circuits. Each is bound to its calldata by a Fiat–Shamir challenge that the contract
derives from keccak256 over the calldata and the circuit's Poseidon digests, and that the circuit
checks through a random linear combination of every transcript word (BRLC). A transcript word
cannot differ between what the contract read and what the prover proved, and a word that is not
a canonical field element is rejected.

| Circuit | Proves | Public inputs |
|---|---|---:|
| Contribution | 16 polynomials of degree `< t`: commitments in the prime subgroup, each recipient's share consistent with them, each share encrypted to its recipient | 8 |
| Finalize | for up to 32 accepted dealers: their commitments match the stored digests, the aggregate pool keys and every member's share commitment for every key | 7 |
| PartialDecrypt | the Chaum–Pedersen relation of one partial against a committed share | 15 |
| DecryptCombine | the Lagrange interpolation over the qualifying set, knowledge of the organizer secret and the decryption equation | 9 |

Shares travel masked in the BN254 scalar field with a one-time pad derived from an ECDH secret
between dealer and recipient. The compiled circuits and their keys are published on the CDN
(`https://davinci-assets.fra1.cdn.digitaloceanspaces.com/dkg/circuits-v6/`) with the GitHub
release as fallback, and pinned by SHA-256 in `config/circuit_artifacts.go`. A node downloads
them on first start and verifies every file against those hashes, again whenever it loads a
proving key, and never compiles a circuit at runtime.

## Contracts

Three contracts, deployed as `DKGRegistry → DKGManager → DKGAppManager` and wired with
`DKGRegistry.setManager` and `DKGManager.setAppManager`. The split only keeps each contract under
the EIP-170 size limit.

| Contract | Owns |
|---|---|
| `DKGRegistry` | Operator keys (BabyJubJub, checked on curve and in the prime subgroup) and liveness: `heartbeat`, `reactivate`, `reap` |
| `DKGManager` | Epoch lifecycle, pool keys, ciphertexts, partial and combined decryptions |
| `DKGAppManager` | `registerApplication`, `revealOrganizerSecret` and the submission and decryption gates the manager consults |

`solidity/src/interfaces/*.sol` hold the full signatures and events. `submitContribution` and
`finalizeEpoch` only accept direct calls from an externally owned account, because nodes recover
contributions from transaction calldata.

| Constant | Where | Value | Notes |
|---|---|---|---|
| `EPOCH_DURATION_BLOCKS` | `DKGManager` constructor | deployment | Epoch cadence |
| `COMMITTEE_SELECTION_BLOCKS`, `KEY_ASSEMBLY_BLOCKS`, `FINALIZE_GAP_BLOCKS` | `DKGManager` constructor arguments | deployment | Preparation windows, stored as deadline offsets |
| `MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS` | `DKGManager` constructor | deployment | Floors for `createEpoch` |
| `INACTIVITY_WINDOW` | `DKGRegistry` constructor | deployment | Blocks without activity before `reap` |
| `MAX_N`, `MAX_T` | `Sizes.sol` | 32 | Committee and threshold caps, fixed by the circuits |
| `MAX_K` | `Sizes.sol` | 16 | Pool keys per epoch |
| `MERKLE_DEPTH` | `Sizes.sol` | 5 | `log2(MAX_N)` |
| `SEED_DELAY_BLOCKS` | `Sizes.sol` | 1 | Lottery seed offset |
| `MAX_SUBMITTERS` | `DKGAppManager` | 32 | Allow-list cap |

The values of the public deployments are in [deployments.md](deployments.md).

## References

- The construction and its security arguments: [NI-DKG paper](https://eprint.iacr.org/2026/552).
- J. Groth, *On the Size of Pairing-based Non-interactive Arguments*, EUROCRYPT 2016.
- P. Feldman, *A Practical Scheme for Non-interactive Verifiable Secret Sharing*, FOCS 1987.
