# Davinci DKG SDK

TypeScript SDK for the Davinci DKG smart contracts. Covers read/write access to `DKGManager`, `DKGRegistry` and `DKGAppManager`, ElGamal encryption on BabyJubJub under per-application pool keys, the pool's claim status, the organizer's one-time secret reveal, helpers for polling epoch status and decryption results, and the transcript codecs (compact contribution, batched finalization, share-commitment Merkle tree) that read the key pool back from calldata.

## Installation

```sh
npm install @vocdoni/davinci-dkg-sdk
# or
pnpm add @vocdoni/davinci-dkg-sdk
```

Requires `viem` as a peer dependency.

## Quick start

```ts
import { createPublicClient, createWalletClient, http, defineChain } from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import {
  DKGClient,
  DKGWriter,
  AppMode,
  buildElGamal,
  encryptForApplication,
  randomOrganizerSecret,
  decryptionProgress,
  waitForEpochPhase,
  waitForDecryption,
  watchNewEpochs,
  watchCiphertextSubmitted,
  EpochPhase,
  buildEpochId,
} from '@vocdoni/davinci-dkg-sdk';

const chain = defineChain({ id: 1337, name: 'Anvil', nativeCurrency: { name: 'Ether', symbol: 'ETH', decimals: 18 }, rpcUrls: { default: { http: ['http://localhost:8545'] } } });

const publicClient = createPublicClient({ chain, transport: http() });
const account = privateKeyToAccount('0x...');
const walletClient = createWalletClient({ chain, transport: http(), account });

const client = new DKGClient({
  publicClient,
  managerAddress: '0x...',
  registryAddress: '0x...',
});

const writer = new DKGWriter({
  publicClient,
  walletClient,
  managerAddress: '0x...',
  registryAddress: '0x...',
});
```

## Reading on-chain state

```ts
// Network overview
const blockNumber = await client.blockNumber();
const totalNodes  = await client.nodeCount();
const activeNodes = await client.activeCount();
const nonce       = await client.epochNonce();

// Epoch details
const epoch = await client.getEpoch(epochId);
console.log(epoch.status);           // 1=CommitteeSelection, 2=KeyAssembly, 3=Live …
console.log(epoch.policy.threshold);

// Deploy-time bounds createEpoch enforces (MIN_THRESHOLD, MIN_COMMITTEE_SIZE, MAX_LOTTERY_ALPHA_BPS)
const bounds = await client.getEpochBounds();

// Participants and contributions
const participants = await client.selectedParticipants(epochId);
const contrib      = await client.getContribution(epochId, participantAddress);

// Pool keys (MaxK = 16 dealt per epoch, all stored by finalizeEpoch; each application claims one)
const { nextIndex } = await client.getPoolStatus(epochId);     // the claim cursor, MAX_K once spent
const poolKey = await client.getPoolKey(epochId, 0);           // reverts InvalidPhase before Live
const poolKeys = await client.getPoolKeys(epochId);            // all MAX_K keys, TE form
const pkAid   = await client.getApplicationKey(epochId, aid);  // P_j (automatic) or P_j + PK_org (locked)
const finalize = await client.getFinalizeTranscript(epochId);  // decoded finalizeEpoch calldata: keys + every D_{j,i}
const proof = await client.getShareProof(epochId, 0, 1);       // member 1's Merkle path under key 0

// Decryption state (ciphertexts are keyed by application id; there is no bare epoch-key path)
const count    = await client.ciphertextCount(epochId, aid);
const partial  = await client.getPartialDecryption(epochId, aid, participantIndex, ciphertextIndex);
const combined = await client.getCombinedDecryption(epochId, aid, ciphertextIndex);

// Registry (registeredAtBlock: the node only enters lotteries of epochs created after it)
const node = await client.getNode(operatorAddress);
```

## Writing transactions

```ts
// Create an epoch. Permissionless once `nextEpochStartBlock()` is reached
// (or earlier, when the newest epoch is Live with at most one unclaimed
// pool key, or Aborted).
// Only the four policy fields are sent; phase deadlines are derived on-chain
// from EPOCH_DURATION_BLOCKS. The values must satisfy `getEpochBounds()` plus
// 1 ≤ threshold ≤ minValidContributions ≤ committeeSize ≤ MaxN and
// lotteryAlphaBps ≥ 10000, or the call reverts InvalidPolicy().
const hash = await writer.createEpoch({
  threshold:             2,
  committeeSize:         3,
  minValidContributions: 2,
  lotteryAlphaBps:       15000,   // 1.5× over-subscription
});
await writer.waitForTransaction(hash);

// Derive the epoch ID from the nonce after creation
const nonce   = await client.epochNonce();
const epochId = await client.buildEpochId(nonce);

// Register a DKG node (the SDK derives the public key and the Schnorr PoK)
await writer.registerKey(babyJubPrivateKey);

// Claim a slot (DKG node role — after the seed block has been mined)
await writer.claimSlot(epochId);

// Register an application once the epoch is Live; it claims the next
// unclaimed pool key (reverts PoolExhausted once all 16 are taken; an epoch
// serves at most 16 applications, the next epoch refills the pool).
// Organizer-locked (the default) needs a one-time reveal before the
// committee can combine; automatic has no organizer key at all — read
// "Applications and the organizer secret" before choosing.
const skOrg = randomOrganizerSecret();              // store this somewhere safe (locked mode)
const aid   = '0x…';                                // your bytes32 application id
await writer.registerApplication(epochId, aid, {
  mode:             AppMode.OrganizerLocked,  // or AppMode.Automatic
  openSubmission:   false,                    // true: anyone may submitCiphertext
  submitters:       [],                       // exclusive allow-list (≤ 32); empty and not open = you only
  maxCiphertexts:   0,                        // 0 = unlimited (256 hard cap)
  notBeforeBlock:   0n,
  notAfterBlock:    0n,
  decryptNotBefore: 0n,                       // unix seconds; 0 = unbounded
  decryptNotAfter:  0n,
}, skOrg);

// Encrypt under PK_aid = P_j (automatic) or P_j + PK_org (locked) and
// publish. The contract assigns the index (1, 2, … per aid).
const pkAid = await client.getApplicationKey(epochId, aid);
const ciphertext = await encryptForApplication(42n, pkAid);
const { ciphertextIndex } = await writer.submitCiphertext(epochId, aid, ciphertext);

// Locked mode only, once per application, whenever the organizer is ready:
// reveal `sk_org` so the committee can combine by itself — for every past
// and future ciphertext of the application, not just this one. Automatic
// applications skip this step entirely.
await writer.revealOrganizerSecret(epochId, aid, skOrg);

// Read the recovered plaintext (nodes combine automatically once threshold
// partial decryptions are on-chain, the decryption window is open, and —
// for a locked application — the secret has been revealed).
const plaintext = await client.getPlaintext(epochId, aid, ciphertextIndex);

// Abort a dead epoch. Permissionless, but only succeeds once the selection
// deadline passed without a full committee, or key assembly closed with
// fewer than minValidContributions accepted.
await writer.abortEpoch(epochId);
```

## Monitoring rounds

```ts
import { waitForEpochPhase, waitForDecryption, watchNewEpochs } from '@vocdoni/davinci-dkg-sdk';

// Poll until an epoch reaches the Live status
await waitForEpochPhase(client, epochId, EpochPhase.Live, {
  intervalMs: 2000,
  timeoutMs:  120_000,
});

// Poll until the ciphertext has been decrypted on-chain
const record = await waitForDecryption(client, epochId, aid, ciphertextIndex);
console.log(record.completed); // true

// One-shot pipeline snapshot: partials so far, whether the decryption
// window is open, and whether it has been combined.
const progress = await decryptionProgress(client, epochId, aid, ciphertextIndex);
console.log(progress.combined);

// Subscribe to new epochs in real time (returns unsubscribe fn)
const unsub = watchNewEpochs(client, (epochId, organizer) => {
  console.log('new epoch', epochId, 'by', organizer);
});
unsub(); // stop watching

// Subscribe to ciphertexts as they land. `c1`/`c2` come back in TE form.
watchCiphertextSubmitted(client, epochId, ({ aid, ciphertextIndex, c1, c2 }) => {
  console.log('ciphertext', aid, ciphertextIndex, c1, c2);
});
```

## ElGamal encryption

The SDK exposes ElGamal encryption/decryption on the BabyJubJub curve via [`@zk-kit/baby-jubjub`](https://github.com/zk-kit/zk-kit/tree/main/packages/baby-jubjub), a pure TypeScript implementation that runs synchronously in the browser.

```ts
import { buildElGamal } from '@vocdoni/davinci-dkg-sdk';

const elgamal = await buildElGamal();

// Generate a key pair
const { privKey, pubKey } = elgamal.generateKeyPair();

// Encrypt a small integer (plaintext must fit in the BabyJubJub subgroup order)
const ciphertext = elgamal.encrypt(42n, pubKey);
// { c1: [bigint, bigint], c2: [bigint, bigint] }

// Decrypt (baby-step/giant-step DLOG — works for values < 2^32)
const plaintext = elgamal.decrypt(ciphertext, privKey);
console.log(plaintext); // 42n

// Point arithmetic
const doubled = elgamal.addPoint(pubKey, pubKey);
const via2    = elgamal.mulPoint(pubKey, 2n);

// Compact serialization (compressed point ↔ bigint)
const packed   = elgamal.packPoint(pubKey);   // bigint
const unpacked = elgamal.unpackPoint(packed); // [bigint, bigint]
```

The `flow` module provides higher-level wrappers for typical usage:

```ts
import {
  encrypt,
  encryptForApplication,
  applicationKey,
  decrypt,
} from '@vocdoni/davinci-dkg-sdk';

// Encrypt under a raw key (local use — nothing on chain can open this)
const ciphertext = await encrypt(42n, someKey);

// Encrypt for on-chain submission: PK_aid = P_j (automatic) or P_j + PK_org (locked)
const pkAid = await client.getApplicationKey(epochId, aid);
const ct = await encryptForApplication(42n, pkAid);
// …or, if you already hold the pool key point and the organizer key locally
const pkAidLocal = applicationKey(poolKeyPoint, app.organizerPK);
```

## Applications and the organizer secret

Ciphertexts are always submitted under a registered application id (`aid`); there is no
epoch-key path, and an unregistered `aid` reverts. Registration claims the next unclaimed
**pool key** `P_j` — one of the `MaxK = 16` keys the epoch's DKG deals, one polynomial per
key inside the same contribution proofs, all proven and stored by the one `finalizeEpoch` —
and binds it into the application's key:

```
PK_aid = P_j                    (automatic — no organizer key at all)
PK_aid = P_j + PK_org           (organizer-locked)   PK_org = sk_org · G
```

Because every application claims a distinct `P_j`, a ciphertext copied out of one application
and re-submitted under another decrypts under an unrelated key: the result is garbage, not the
original plaintext. **This closes the cross-application decryption oracle** that existed when
every application shared one epoch key `PK_ep` — previously, an automatic application's
published `sk_org` let anyone learn `sk_ep · C1` for a `C1` they never had a right to; that path
no longer exists.

An application is registered in one of two **modes**, fixed by `policy.mode`:

- `AppMode.OrganizerLocked` (the default). `writer.registerApplication(epochId, aid, policy,
  skOrg)` derives `PK_org` from `skOrg` and builds the Schnorr proof of possession the contract
  verifies (domain `davinci-dkg:organizer-register:v1`). Only `PK_org` and the proof go on
  chain; `skOrg` stays with you until you choose to call
  `writer.revealOrganizerSecret(epochId, aid, skOrg)` — a **one-time, whole-application** act,
  not per-ciphertext. Before the reveal, no one — not even `t` colluding committee members —
  can decrypt any ciphertext of the application, and the contract refuses every partial and
  combine (`OrganizerSecretNotRevealed()`), so the organizer learns results together with
  everyone else: it decides *when* the application opens, never *which* ciphertexts. After the
  reveal the committee combines by itself for every ciphertext, past and future, once the
  decryption window is open.
- `AppMode.Automatic`. There is **no organizer key at all**: `registerApplication` ignores the
  key and Schnorr arguments, and the contract stores the identity point `(0, 1)` as `PK_org`
  and `0` as `organizerSecret`, so `PK_aid = P_j` directly. The committee alone can decrypt, as
  soon as `t` partials land and the decryption window is open — there is no reveal step, ever.
  Confidentiality of such an application is exactly the `t`-of-`n` committee assumption.

> **Keep `sk_org` of an organizer-locked application until you mean to reveal it.** It is the
> application's only decryption capability before the reveal, it is never transmitted, and the
> SDK cannot recover it. **If you lose it before revealing, every ciphertext ever submitted
> under that `aid` is permanently undecryptable** — the committee threshold alone cannot open
> them, by design. Draw it with `randomOrganizerSecret()` and persist it at registration.
> Revealing is irreversible and application-wide: once posted, anyone (together with the
> committee) can open every past and future ciphertext of that application — within an
> application the organizer is trusted and accountable; the guarantee the DKG adds is *across*
> applications. Never reuse an organizer secret across two locked applications: revealing it
> for one opens the other too.

The SDK takes the secret as a plain `bigint` parameter everywhere. It deliberately does not
derive it from a wallet signature: that would tie the application's decryptability to one
wallet's exact signing behaviour, which no wallet guarantees across versions.

The rest of the policy is fixed at registration too. **Submission**: `openSubmission: true`
lets anyone call `submitCiphertext`; otherwise `submitters` is an exclusive allow-list of up to
32 addresses (the registrant is not implicitly on it), and when both are empty only the
registrant may submit. **Submission window**: `notBeforeBlock` / `notAfterBlock` (blocks, `0n` =
unbounded) plus `maxCiphertexts`; `submitCiphertext` is gated by that window, the submitter
policy and `decryptNotAfter` (a ciphertext nobody may ever decrypt is refused), and may land
before decryption opens. **Decryption window**: `decryptNotBefore` / `decryptNotAfter` (unix
seconds, `0n` = unbounded) make `submitPartialDecryption` and `combineDecryption` revert
`DecryptionNotOpen()` before the window opens and `DecryptionClosed()` once it has passed; a
locked application additionally reverts `OrganizerSecretNotRevealed()` until its organizer has
revealed. The window bounds what the contract accepts and what honest nodes post — it does not
bind `t` colluding members, who hold shares regardless (for a locked application they still lack
`sk_org`), but since nothing is accepted before the window or the reveal there are no on-chain
partials from before it either. An empty window (`notBefore` after `notAfter` when both are
set), open submission with a non-empty list, a zero address on the list or more than 32 entries
revert `InvalidPolicy()`.

There is **no proof of knowledge of the ElGamal randomness**. The submitter of an aggregated
tally cannot know its randomness, so such a proof is incompatible with homomorphic
aggregation; cross-application replay is stopped by the pool key itself now, not by the
organizer key — a `C1` copied into another application decrypts under that application's own
`P_j`, an unrelated secret, so the result is meaningless there.

To decrypt a ciphertext:

1. The submitter publishes it with `submitCiphertext(...)` — the contract assigns the next
   index for `(epochId, aid)`, stores `keccak256(c1, c2)` and emits `CiphertextSubmitted`
   carrying the raw coordinates.
2. For a locked application, the organizer calls `revealOrganizerSecret` once, whenever they
   choose — automatic applications skip this step entirely. Nothing below happens before it.
3. DKG nodes watch the event and, once the decryption window is open (and the secret revealed,
   for a locked application), each call `submitPartialDecryption` on `DKGManager` with a Merkle
   proof of their share commitment against the claimed pool key's on-chain root. The tree covers
   every committee member — leaves `keccak256(0x00 ‖ x ‖ y)`, empty leaves
   `keccak256("davinci-dkg:merkle-empty:v1")`, nodes `keccak256(0x01 ‖ l ‖ r)`, depth 5 — so a
   member that did not contribute can still post.
4. Once `t` partials are on chain, the decryption window is open, and — for a locked
   application — the secret has been revealed, any DKG node calls `combineDecryption`; the
   proof is bound to the stored ciphertext hash so combine cannot be mounted against a
   different ciphertext.
5. `DecryptionCombined` is emitted, `getCombinedDecryption` returns `completed: true`, and
   `getPlaintext(epochId, aid, ciphertextIndex)` returns the recovered plaintext `uint256`.

## Full DKG flow overview

```
[Anyone]    createEpoch(threshold, n, minValid, alphaBps)   ← permissionless, cadence-gated;
               │                                               also allowed early when the
               │                                               newest epoch is Live with ≤ 1
               │                                               unclaimed key, or Aborted
               ▼  (seed block mined)
[DKG Node]  claimSlot(epochId)           ← lottery via on-chain blockhash seed; only
               │                            nodes registered before the epoch qualify
               ▼  (committee full)
[DKG Node]  submitContribution(...)      ← ZK proof dealing all MaxK = 16 pool-key
               │                            polynomials at once; compact calldata,
               │                            MaxK·(2t+n) + 5n words
               ▼  (key-assembly deadline + live-not-before block reached)
[Anyone]    finalizeEpoch(eid, transcriptDigest, transcript, proof, input)
               │                          ← one ZK proof over the accepted contributor
               │                            set; stores all MaxK keys P_j and the
               │                            committee's share-commitment Merkle roots,
               │                            then flips the epoch to Live (direct EOA
               │                            call only: the 1,120-word transcript lives
               │                            in the calldata)
               ▼  Epoch.status = Live — every key exists
[Organizer] registerApplication(eid, aid, policy, skOrg)  ← claims the next unclaimed
               │                                             key j; locked: Schnorr PoP of
               │                                             sk_org; automatic: no organizer
               │                                             key at all; PK_aid = P_j (+PK_org)
               ▼
[Anyone]    getApplicationKey(eid, aid) → PK_aid     ← simple contract read
               │
               ▼
[Anyone]    encryptForApplication(m, PK_aid)         ← ElGamal in the browser
               │
               ▼
[Submitter] submitCiphertext(epochId, aid, ct)       ← gated by the app's AppPolicy;
               │                                        index assigned on-chain;
               │                                        emits CiphertextSubmitted
               ▼
[Organizer] revealOrganizerSecret(eid, aid, sk)      ← locked apps only, once, whenever
               │                                        the organizer chooses; automatic
               │                                        apps skip this entirely; no partial
               │                                        is accepted before it
               ▼  (decryption window open AND, if locked, revealed)
[DKG Node]  submitPartialDecryption(..., shareProof) ← picked up from the event;
               │                                        Merkle path against the pool
               │                                        key's share-commitment root
               ▼  (t partials on chain)
[DKG Node]  combineDecryption(...)                   ← proof bound to the stored ct hash;
               │                                        emits DecryptionCombined(plaintext)
               ▼
[Anyone]    getPlaintext(epochId, aid, idx)          ← plaintext is on-chain

(dead epoch: selection deadline passed without a full committee, or key assembly
 closed below minValidContributions)  →  [Anyone] abortEpoch(epochId)
```

> **Pool keys:** Each epoch's DKG deals `MaxK = 16` independent keys `P_0…P_15` inside the same
> contribution proofs used for the committee itself. The single proof-carrying `finalizeEpoch`
> proves and stores all of them (plus one share-commitment Merkle root per key) atomically, so
> `Live` means every key exists; each application claims exactly one at registration.
> `client.getPoolStatus(epochId)` reports the claim cursor; `client.getPoolKey(epochId, j)` /
> `client.getPoolKeys(epochId)` return the keys' `(x, y)`; `client.getApplicationKey(epochId, aid)`
> returns the application's `PK_aid` directly; `client.getFinalizeTranscript(epochId)` decodes the
> finalization calldata (every key and every member's share commitment `D_{j,i}` at word
> `64 + 66·j + 2 + 2·i`) and `client.getShareProof(epochId, j, participantIndex)` builds the
> Merkle path a partial decryption carries, checked against the stored root.

## API reference

### `DKGClient`

| Method | Description |
|--------|-------------|
| `getEpoch(epochId)` | Full epoch struct (no per-epoch decryption policy any more) |
| `getEpochBounds()` | Deploy-time `createEpoch` bounds: `minThreshold`, `minCommitteeSize`, `maxLotteryAlphaBps` |
| `selectedParticipants(epochId)` | Addresses that claimed a slot |
| `getContribution(epochId, address)` | Contribution record |
| `getPartialDecryption(epochId, aid, participantIndex, idx)` | Partial decryption record |
| `getCombinedDecryption(epochId, aid, idx)` | Combined decryption record (includes `plaintext`) |
| `getPlaintext(epochId, aid, idx)` | Recovered plaintext scalar `uint256` |
| `getCiphertextHash(epochId, aid, idx)` | `keccak256(c1,c2)` of the submitted ciphertext |
| `ciphertextCount(epochId, aid)` | Ciphertexts accepted so far under `(epochId, aid)`; indices run 1…count |
| `getApplication(epochId, aid)` | Cached `Application` record (creator, `PK_org` in TE form — identity for automatic —, `organizerSecret` — `0` until a locked app's organizer reveals it, always `0` for automatic —, `poolIndex`, `AppPolicy`) |
| `getApplicationKey(epochId, aid)` | The application's `PK_aid` directly: `P_j` (automatic) or `P_j + PK_org` (locked) |
| `getPoolKey(epochId, j)` | One pool key's `(x, y)`, TE form; reverts `InvalidPhase` before the epoch is Live |
| `getPoolKeys(epochId)` | All `MAX_K` pool keys of a Live epoch, by index |
| `getPoolStatus(epochId)` | `{ nextIndex }` — the claim cursor (`MAX_K` once the pool is spent) |
| `getFinalizeTranscript(epochId)` | The decoded `finalizeEpoch` calldata: pool keys, every member's share commitment, digest, public inputs, finalizer; null before Live |
| `getShareProof(epochId, j, participantIndex)` | Member `participantIndex`'s leaf, `MERKLE_DEPTH` siblings and root under key `j`, from the finalization calldata and checked against `getPoolShareRoot` |
| `getPoolShareRoot(epochId, j)` | Merkle root of key `j`'s per-member share commitments |
| `getAppPoolIndex(epochId, aid)` | Which pool key index an application claimed |
| `getAppManagerAddress()` | Resolve the `DKGAppManager` address |
| `getCiphertextSubmittedEvents(epochId, opts?)` | Historical `CiphertextSubmitted` logs (raw C1/C2 coords) |
| `getDecryptionCombinedEvents(epochId, opts?)` | Historical `DecryptionCombined` logs (carries `plaintext`) |
| `getNode(address)` | Registry node record (includes `registeredAtBlock`) |
| `nodeCount()` / `activeCount()` | Registry stats |
| `isActive(address)` | Node liveness check |
| `inactivityWindow()` | Blocks after which a silent node can be reaped |
| `blockNumber()` | Current block |
| `epochNonce()` | Nonce of the most recent epoch |
| `buildEpochId(nonce)` | Derive an epoch ID from a nonce |
| `getEpochDurationBlocks()` / `getNextEpochStartBlock()` / `getLastEpochStartBlock()` | Cadence reads |
| `getEpochCreatedEvents(opts?)` | Historical EpochCreated logs |
| `getEpochLiveEvents(epochId)` | Historical EpochLive logs (includes `transactionHash`) |
| `getAllEpochEvents(epochId, fromBlock?)` | All DKGManager events for a specific epoch |
| `getRecentEpochs(limit?)` | Most recent N epochs (default 20) as `EpochEntry[]` |
| `getRegistryNodes(fromBlock?)` | All registered nodes via NodeRegistered events |
| `roundPrefix()` | Fetch the immutable EPOCH_PREFIX constant |
| `watchManagerEvents(handler)` | Live event stream (returns unsubscribe fn) |
| `watchRegistryEvents(handler)` | Live registry events |

### `DKGWriter` (extends `DKGClient`)

All `DKGClient` methods plus:

| Method | Description |
|--------|-------------|
| `createEpoch({ threshold, committeeSize, minValidContributions, lotteryAlphaBps })` | Create a new DKG epoch (permissionless, cadence-gated, bounded by `getEpochBounds()`) |
| `claimSlot(epochId)` | Claim a committee slot |
| `submitContribution(...)` | Submit VSS contribution + ZK proof, dealing all `MaxK` pool-key polynomials at once |
| `finalizeEpoch(epochId, transcriptDigest, transcript, proof, input)` | Finalize the epoch with the `circuits/finalize` ZK proof over the whole key pool (7 public inputs; `transcriptDigest` is the proof's Poseidon transcript digest, folded into the Fiat–Shamir anchor); stores every key and share root and flips the epoch to Live; permissionless, direct EOA call only |
| `submitCiphertext(epochId, aid, ciphertext)` | Publish a ciphertext; waits for the receipt and returns `{ hash, receipt, ciphertextIndex }` with the on-chain-assigned index |
| `submitPartialDecryption(...)` | Submit partial decryption + ZK proof, with a Merkle proof of the share commitment against the claimed pool key's root; reverts until the decryption window is open and, for a locked application, the organizer has revealed |
| `combineDecryption(epochId, aid, idx, combineHash, plaintext, ...)` | Combine partial decryptions + ZK proof; stores plaintext |
| `registerApplication(epochId, aid, policy, skOrg)` | Register an application, claiming the next unclaimed pool key (reverts `PoolExhausted` once all `MAX_K` are taken); derives `PK_org` and, for `AppMode.OrganizerLocked`, the Schnorr proof of possession; `AppMode.Automatic` sends no organizer material at all |
| `revealOrganizerSecret(epochId, aid, skOrg)` | Reveal the organizer secret once (locked applications only); from then on the committee combines by itself for every past and future ciphertext of the application |
| `abortEpoch(epochId)` | Abort a dead epoch (permissionless; reverts unless the epoch can no longer progress) |
| `registerKey(privateKey)` | Register a DKG node in the registry (derives the key + Schnorr PoK) |
| `updateKey(privateKey)` | Update node public key |
| `heartbeat()` | Refresh node liveness |
| `reactivate()` | Rejoin after being reaped |
| `reap(operator)` | Permissionlessly reap a stale node |
| `waitForTransaction(hash)` | Wait for tx receipt |
| `createRoundAndWait(policy)` | createEpoch + wait |

### Monitor utilities

| Export | Description |
|--------|-------------|
| `waitForEpochPhase(client, epochId, status, opts?)` | Poll until epoch status reached |
| `waitForDecryption(client, epochId, aid, idx, opts?)` | Poll until decryption complete |
| `watchNewEpochs(client, onEpoch, fromBlock?)` | Subscribe to new epochs |
| `watchEpochLive(client, epochId, onFinalized)` | Subscribe to finalization |
| `watchCiphertextSubmitted(client, epochId, onCiphertext, opts?)` | Subscribe to ciphertexts (C1/C2 coords in TE form) |
| `decryptionProgress(client, epochId, aid, idx)` | One-shot pipeline snapshot (partials so far, decryption window, combined) |
| `watchDecryptionCombined(client, epochId, idx, onCombined, opts?)` | Subscribe to decryption |
| `networkSummary(client)` | Block, node counts, epoch nonce |

### Flow helpers

Higher-level helpers built on top of the primitives above:

| Export | Description |
|--------|-------------|
| `encrypt(message, pubKey, k?)` | ElGamal encrypt under any BabyJubJub key |
| `encryptForApplication(message, pkAid, k?)` | ElGamal encrypt under an application's `PK_aid` |
| `decrypt(ciphertext, privKey)` | ElGamal decrypt via BSGS DLOG (values < 2^32) |
| `waitForCombinedDecryption(client, epochId, aid, idx, opts?)` | Wait for on-chain decryption to complete |
| `demonstrateEncryptDecryptFlow(client, epochId, aid, pubKey, plaintext, idx)` | End-to-end demo flow |

### Proof helpers

| Export | Description |
|--------|-------------|
| `proveOperator` / `verifyOperatorSchnorr` | Operator registration Schnorr proof of possession |
| `proveOrganizer` / `verifyOrganizerSchnorr` | Organizer registration Schnorr proof of possession |
| `verifyDleq` / `dleqChallenge` | Committee partial-decryption Chaum-Pedersen check |
| `applicationKey(poolKey, pkOrg)` | `PK_aid = P_j + PK_org` (pass the identity point for automatic apps) |
| `randomOrganizerSecret()` | Fresh `sk_org` from the platform CSPRNG |

### ElGamal interface

| Method | Description |
|--------|-------------|
| `generateKeyPair()` | Return `{ privKey: bigint, pubKey: BabyJubPoint }` |
| `randomScalar()` | Uniformly random scalar in the BabyJubJub subgroup |
| `encrypt(msg, pubKey, k?)` | Encrypt a small integer; `k` is an optional ephemeral scalar |
| `decrypt(ciphertext, privKey)` | BSGS DLOG; works for msg < 2^32 |
| `packPoint(p)` | Compress a curve point to a single `bigint` |
| `unpackPoint(packed)` | Decompress back to `[bigint, bigint]` |
| `mulPoint(point, scalar)` | Scalar multiplication |
| `addPoint(a, b)` | Curve point addition |

## Development

```sh
pnpm install
pnpm run check             # type-check only (no emit)
pnpm run build             # emit to dist/
pnpm run test              # unit + fixture tests
pnpm run test:integration  # end-to-end tests against a live chain (requires RUN_INTEGRATION_TESTS=true)
pnpm run test:watch        # watch mode
```
