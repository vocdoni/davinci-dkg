# Davinci DKG SDK

TypeScript SDK for the Davinci DKG smart contracts. Covers read/write access to `DKGManager`, `DKGRegistry` and `DKGAppManager`, ElGamal encryption on BabyJubJub under per-application keys, the organizer's decryption share, and helpers for polling epoch status and decryption results.

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

// Decryption state (ciphertexts are keyed by application id; use 0x00…00 for the bare epoch key)
const count    = await client.ciphertextCount(epochId, aid);
const partial  = await client.getPartialDecryption(epochId, aid, participantIndex, ciphertextIndex);
const combined = await client.getCombinedDecryption(epochId, aid, ciphertextIndex);

// Registry (registeredAtBlock: the node only enters lotteries of epochs created after it)
const node = await client.getNode(operatorAddress);
```

## Writing transactions

```ts
// Create an epoch. Permissionless once `nextEpochStartBlock()` is reached.
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

// Register an application once the epoch is Live. Every application is
// organizer co-decryption: keep `skOrg`, it is the only decryption capability.
const skOrg = randomOrganizerSecret();              // store this somewhere safe
const aid   = '0x…';                                // your bytes32 application id
await writer.registerApplication(epochId, aid, {
  authorizedSubmitter: '0x0000000000000000000000000000000000000000', // = you
  maxCiphertexts:      0,
  notBeforeBlock:      0n,
  notAfterBlock:       0n,
}, skOrg);

// Encrypt under PK_aid = PK_ep + PK_org and publish. The contract assigns the
// index (1, 2, … per aid).
const pk  = await client.getCollectivePublicKey(epochId);
const app = await client.getApplication(epochId, aid);
const ciphertext = await encryptForApplication(42n, [pk.x, pk.y], app.organizerPK);
const { ciphertextIndex } = await writer.submitCiphertext(epochId, aid, ciphertext);

// Release the organizer share so the committee can combine.
await writer.submitOrganizerShare(epochId, aid, ciphertextIndex, ciphertext, skOrg);

// Read the recovered plaintext (nodes combine automatically once threshold
// partial decryptions and the organizer share are on-chain).
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

// Where is a ciphertext stuck? `organizerShare` tells "waiting for the
// committee" apart from "waiting for the organizer".
const progress = await decryptionProgress(client, epochId, aid, ciphertextIndex);
console.log(progress.organizerShare, progress.combined);

// Subscribe to new epochs in real time (returns unsubscribe fn)
const unsub = watchNewEpochs(client, (epochId, organizer) => {
  console.log('new epoch', epochId, 'by', organizer);
});
unsub(); // stop watching

// Subscribe to ciphertexts as they land. `c1`/`c2` come back in TE form, so
// they can be handed straight to `writer.submitOrganizerShare`.
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
  waitForCollectivePublicKeyHash,
} from '@vocdoni/davinci-dkg-sdk';

// Encrypt under a raw key (local use — nothing on chain can open this)
const ciphertext = await encrypt(42n, someKey);

// Encrypt for on-chain submission: PK_aid = PK_ep + PK_org
const ct = await encryptForApplication(42n, [pk.x, pk.y], app.organizerPK);
// …or derive the key yourself
const pkAid = applicationKey([pk.x, pk.y], app.organizerPK);

// Wait for an epoch to go Live and return the on-chain public key hash
const hash = await waitForCollectivePublicKeyHash(client, epochId);
```

## Applications and the organizer secret

Ciphertexts are always submitted under a registered application id (`aid`); there is no
epoch-key path, and an unregistered `aid` reverts. Registration binds an organizer key:

```
PK_aid = PK_ep + PK_org        PK_org = sk_org · G
```

`writer.registerApplication(epochId, aid, policy, skOrg)` derives `PK_org` from `skOrg` and
builds the Schnorr proof of possession the contract verifies (domain
`davinci-dkg:organizer-register:v1`). Only `PK_org` and the proof go on chain.

> **Keep `sk_org`.** It is the application's only decryption capability, it is never
> transmitted, and the SDK cannot recover it. **If you lose it, every ciphertext ever
> submitted under that `aid` is permanently undecryptable** — the committee threshold alone
> cannot open them, by design. Draw it with `randomOrganizerSecret()` and persist it before
> you register. Conversely, anyone holding it can (together with the committee) open every
> ciphertext of the application: within an application the organizer is trusted and
> accountable; the guarantee the DKG adds is *across* applications.

The SDK takes the secret as a plain `bigint` parameter everywhere. It deliberately does not
derive it from a wallet signature: that would tie the application's decryptability to one
wallet's exact signing behaviour, which no wallet guarantees across versions.

To open a ciphertext the organizer publishes `Δ = sk_org · C1` with a Chaum-Pedersen DLEQ
proving the same `sk_org` relates `(G, PK_org)` and `(C1, Δ)`:

```ts
await writer.submitOrganizerShare(epochId, aid, ciphertextIndex, ciphertext, skOrg);
```

The challenge is a keccak256 over
`DOMAIN_ORGANIZER_SHARE_V1 ‖ eid ‖ aid ‖ uint256(ctIdx) ‖ PK_org ‖ C1 ‖ Δ ‖ A1 ‖ A2`, reduced
mod the BabyJubJub subgroup order, over the on-chain (RTE) words. The contract recomputes it
and stores only `keccak256(Δ ‖ A1 ‖ A2 ‖ z)`; the DLEQ itself is verified inside the
committee's combine SNARK. `proveOrganizerShare` / `verifyOrganizerShare` /
`organizerShareChallenge` are exported so integrators and auditors can build or re-check a
share offline. Anyone may relay a share, and re-submission overwrites until the ciphertext is
combined, so a malformed share cannot brick a ciphertext.

There is **no proof of knowledge of the ElGamal randomness**. The submitter of an aggregated
tally cannot know its randomness, so such a proof is incompatible with homomorphic
aggregation; cross-application replay is stopped by the organizer key instead — a `C1` copied
into another application and decrypted there only yields `sk_ep·C1`, useless without that
application's `sk_org·C1`.

To decrypt a ciphertext:

1. The submitter publishes it with `submitCiphertext(...)` — the contract assigns the next
   index for `(epochId, aid)`, stores `keccak256(c1, c2)` and emits `CiphertextSubmitted`
   carrying the raw coordinates.
2. DKG nodes watch that event and each call `submitPartialDecryption` on `DKGManager`.
3. The organizer calls `submitOrganizerShare` on `DKGAppManager`.
4. Once the threshold is met **and** the share is stored, any DKG node calls
   `combineDecryption`; it reverts `OrganizerShareMissing()` otherwise, and the proof is bound
   to the stored ciphertext hash so combine cannot be mounted against a different ciphertext.
5. `DecryptionCombined` is emitted, `getCombinedDecryption` returns `completed: true`, and
   `getPlaintext(epochId, aid, ciphertextIndex)` returns the recovered plaintext `uint256`.

## Full DKG flow overview

```
[Anyone]    createEpoch(threshold, n, minValid, alphaBps)   ← permissionless, cadence-gated
               │
               ▼  (seed block mined)
[DKG Node]  claimSlot(epochId)           ← lottery via on-chain blockhash seed; only
               │                            nodes registered before the epoch qualify
               ▼  (committee full)
[DKG Node]  submitContribution(...)      ← ZK proof of VSS shares; contract
               │                            adds commitment[0] to collective key
               ▼  (key-assembly deadline + live-not-before block reached)
[DKG Node]  finalizeEpoch(...)           ← ZK proof aggregating all commitments;
               │                            nodes auto-call this on a deterministic
               │                            stagger derived from the lottery seed,
               │                            so only one node submits per epoch.
               │                            collectivePublicKeyHash emitted.
               ▼  Epoch.status = Live
[Anyone]    getCollectivePublicKey(epochId) → {x, y}   ← simple contract read
               │
               ▼
[Organizer] registerApplication(eid, aid, policy, skOrg)  ← Schnorr PoP of sk_org
               │                                             PK_aid = PK_ep + PK_org
               ▼
[Anyone]    encryptForApplication(m, PK_ep, PK_org)  ← ElGamal in the browser
               │
               ▼
[Submitter] submitCiphertext(epochId, aid, ct)       ← gated by the app's AppPolicy;
               │                                        index assigned on-chain;
               │                                        emits CiphertextSubmitted
               ▼
[DKG Node]  submitPartialDecryption(...)             ← picked up from the event
               │
[Organizer] submitOrganizerShare(...)                ← Δ = sk_org·C1 + DLEQ
               ▼  (threshold met AND share stored)
[DKG Node]  combineDecryption(...)                   ← proof bound to stored ct hash and
               │                                        to the stored organizer share;
               │                                        emits DecryptionCombined(plaintext)
               ▼
[Anyone]    getPlaintext(epochId, aid, idx)          ← plaintext is on-chain

(dead epoch: selection deadline passed without a full committee, or key assembly
 closed below minValidContributions)  →  [Anyone] abortEpoch(epochId)
```

> **Collective public key:** The contract accumulates the key incrementally as contributions are accepted — each contributor's `commitment[0]` point is added on-chain during `submitContribution`. The `(x, y)` coordinates are available at any time via `client.getCollectivePublicKey(epochId)`, a simple view-call that requires no calldata parsing. `EpochLive` emits `collectivePublicKeyHash` (keccak256 of the final key) for integrity verification.

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
| `getApplication(epochId, aid)` | Cached `Application` record (creator, `PK_org` in TE form, `AppPolicy`) |
| `getOrganizerShareHash(epochId, aid, idx)` | `keccak256` of the stored organizer share, `0x00…00` if none |
| `hasOrganizerShare(epochId, aid, idx)` | Whether the organizer share is on chain |
| `getOrganizerShareEvents(epochId, aid, opts?)` | Historical `OrganizerShareSubmitted` logs (Δ, A1, A2, z) |
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
| `getCollectivePublicKey(epochId)` | Fetch the collective public key `(x, y)` from the on-chain accumulator (available after the first contribution is accepted) |
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
| `submitContribution(...)` | Submit VSS contribution + ZK proof |
| `finalizeEpoch(...)` | Finalize epoch + ZK proof |
| `submitCiphertext(epochId, aid, ciphertext)` | Publish a ciphertext; waits for the receipt and returns `{ hash, receipt, ciphertextIndex }` with the on-chain-assigned index |
| `submitPartialDecryption(...)` | Submit partial decryption + ZK proof |
| `combineDecryption(epochId, aid, idx, combineHash, plaintext, ...)` | Combine partial decryptions + ZK proof; stores plaintext |
| `registerApplication(epochId, aid, policy, skOrg)` | Register an application; derives `PK_org` and the Schnorr proof of possession |
| `submitOrganizerShare(epochId, aid, idx, ciphertext, skOrg)` | Release `Δ = sk_org·C1` with its DLEQ |
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
| `waitForOrganizerShare(client, epochId, aid, idx, opts?)` | Poll until the organizer share is on chain |
| `decryptionProgress(client, epochId, aid, idx)` | One-shot pipeline snapshot (ciphertext / organizer share / combined) |
| `watchDecryptionCombined(client, epochId, idx, onCombined, opts?)` | Subscribe to decryption |
| `networkSummary(client)` | Block, node counts, epoch nonce |

### Flow helpers

Higher-level helpers built on top of the primitives above:

| Export | Description |
|--------|-------------|
| `encrypt(message, pubKey, k?)` | ElGamal encrypt under any BabyJubJub key |
| `encryptForApplication(message, pkEp, pkOrg, k?)` | ElGamal encrypt under `PK_aid = PK_ep + PK_org` |
| `decrypt(ciphertext, privKey)` | ElGamal decrypt via BSGS DLOG (values < 2^32) |
| `waitForCollectivePublicKeyHash(client, epochId, opts?)` | Wait for Live; return on-chain key hash |
| `waitForCombinedDecryption(client, epochId, aid, idx, opts?)` | Wait for on-chain decryption to complete |
| `demonstrateEncryptDecryptFlow(client, epochId, aid, pubKey, plaintext, idx)` | End-to-end demo flow |

### Proof helpers

| Export | Description |
|--------|-------------|
| `proveOperator` / `verifyOperatorSchnorr` | Operator registration Schnorr proof of possession |
| `proveOrganizer` / `verifyOrganizerSchnorr` | Organizer registration Schnorr proof of possession |
| `proveOrganizerShare` / `verifyOrganizerShare` / `organizerShareChallenge` | Organizer decryption share (Δ = sk_org·C1) and its DLEQ |
| `verifyDleq` / `dleqChallenge` | Committee partial-decryption Chaum-Pedersen check |
| `applicationKey(pkEp, pkOrg)` | `PK_aid = PK_ep + PK_org` |
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
