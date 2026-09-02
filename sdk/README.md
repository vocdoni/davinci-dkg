# Davinci DKG SDK

TypeScript SDK for the Davinci DKG smart contracts. Covers read/write access to `DKGManager`, `DKGRegistry` and `DKGAppManager`, ElGamal encryption on BabyJubJub (with the submitter's proof of knowledge the committee requires), and helpers for polling epoch status and decryption results.

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
  waitForEpochPhase,
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

// Publish a ciphertext for threshold decryption once the epoch is Live.
// `encryptWithProof` binds a Schnorr proof of knowledge of the ElGamal
// randomness to (epochId, aid); committee nodes only decrypt ciphertexts
// whose proof verifies. The contract assigns the index (1, 2, … per aid).
const ZERO_AID = '0x' + '00'.repeat(32);            // bare epoch key
const pk = await client.getCollectivePublicKey(epochId);
const { ciphertext, pok } = await encryptWithProof(epochId, ZERO_AID, 42n, [pk.x, pk.y]);
const { ciphertextIndex } = await writer.submitCiphertext(epochId, ZERO_AID, ciphertext, pok);

// Read the recovered plaintext (nodes combine automatically once threshold
// partial decryptions are on-chain).
const plaintext = await client.getPlaintext(epochId, ZERO_AID, ciphertextIndex);

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
const record = await waitForDecryption(client, epochId, ZERO_AID, ciphertextIndex);
console.log(record.completed); // true

// Subscribe to new epochs in real time (returns unsubscribe fn)
const unsub = watchNewEpochs(client, (epochId, organizer) => {
  console.log('new epoch', epochId, 'by', organizer);
});
unsub(); // stop watching

// Subscribe to ciphertexts as they land; `pokValid` tells you whether the
// committee will decrypt it (nodes skip ciphertexts with a bad proof).
watchCiphertextSubmitted(client, epochId, ({ aid, ciphertextIndex, pokValid }) => {
  console.log('ciphertext', aid, ciphertextIndex, pokValid ? 'ok' : 'INVALID PROOF');
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
import { encrypt, encryptWithProof, decrypt, waitForCollectivePublicKeyHash } from '@vocdoni/davinci-dkg-sdk';

// Encrypt with the collective public key (no proof — for local use only)
const ciphertext = await encrypt(42n, collectivePubKey);

// Encrypt for on-chain submission: also proves knowledge of the randomness
// for exactly this (epochId, aid). Required by `writer.submitCiphertext`.
const { ciphertext: ct, pok } = await encryptWithProof(epochId, aid, 42n, collectivePubKey);

// Wait for an epoch to go Live and return the on-chain public key hash
const hash = await waitForCollectivePublicKeyHash(client, epochId);
```

The proof of knowledge is a Schnorr proof over BabyJubJub that the submitter knows `r` with
`C1 = r·G`, bound to `keccak256(DOMAIN_CIPHERTEXT_POK_V1 ‖ epochId ‖ aid ‖ C1 ‖ C2 ‖ A)` where the
coordinates are the on-chain (RTE) words. It mirrors `crypto/elgamal` in the Go node byte for
byte (`proveCiphertext` / `verifyCiphertextPoK` are exported for auditors). The contract verifies
it (and that `C1` lies in the prime-order subgroup) before accepting the ciphertext, and every
committee node verifies it again before releasing a partial decryption, so a ciphertext without
a valid proof — for instance a `C1` copied from another application's ciphertext as a decryption
oracle — is rejected at submission and never decrypted.

In the DKG protocol the private key is never held by a single party. To decrypt a ciphertext:

1. The consumer publishes the ciphertext on-chain with `submitCiphertext(...)` — the contract
   assigns the next index for `(epochId, aid)`, stores `keccak256(c1, c2)` and emits a
   `CiphertextSubmitted` event carrying the raw coordinates and the proof of knowledge.
2. DKG nodes watch that event, verify the proof, and each call `submitPartialDecryption` on the
   `DKGManager` contract.
3. Once the threshold is met, any DKG node calls `combineDecryption`; the proof is bound to the
   stored ciphertext hash, so combine cannot be mounted against a different ciphertext.
4. The `DecryptionCombined` event is emitted, `getCombinedDecryption` returns `completed: true`,
   and `getPlaintext(epochId, aid, ciphertextIndex)` returns the recovered plaintext `uint256`.

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
[Anyone]    encryptWithProof(epochId, aid, m, pk)   ← ElGamal + PoK in the browser
               │
               ▼
[Submitter] submitCiphertext(epochId, aid, ct, pok) ← gated by the app's AppPolicy;
               │                                        index assigned on-chain;
               │                                        emits CiphertextSubmitted
               ▼
[DKG Node]  submitPartialDecryption(...)             ← picked up from the event
               │                                        after verifying the PoK
               ▼  (threshold met)
[DKG Node]  combineDecryption(...)                   ← proof bound to stored ct hash;
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
| `getApplication(epochId, aid)` | Cached `Application` record (mode, S, PK_org, AppPolicy) |
| `getCiphertextSubmittedEvents(epochId, opts?)` | Historical `CiphertextSubmitted` logs (raw C1/C2 coords, PoK and `pokValid`) |
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
| `submitCiphertext(epochId, aid, ciphertext, pok)` | Publish a ciphertext with its proof of knowledge; waits for the receipt and returns `{ hash, receipt, ciphertextIndex }` with the on-chain-assigned index |
| `submitPartialDecryption(...)` | Submit partial decryption + ZK proof |
| `combineDecryption(epochId, aid, idx, combineHash, plaintext, ...)` | Combine partial decryptions + ZK proof; stores plaintext |
| `registerApplication(epochId, aid, policy)` | Register a mode-0 (public derivation) application |
| `registerApplicationCoDec(epochId, aid, policy, pkOrgX, pkOrgY, ax, ay, z)` | Register a mode-1 (organizer co-decryption) application |
| `submitOrganizerShare(...)` | Submit the organizer's Δ_org share (mode 1) |
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
| `watchCiphertextSubmitted(client, epochId, onCiphertext, opts?)` | Subscribe to ciphertexts (coords, PoK, `pokValid`) |
| `watchDecryptionCombined(client, epochId, idx, onCombined, opts?)` | Subscribe to decryption |
| `networkSummary(client)` | Block, node counts, epoch nonce |

### Flow helpers

Higher-level helpers built on top of the primitives above:

| Export | Description |
|--------|-------------|
| `encrypt(message, pubKey, k?)` | ElGamal encrypt via collective public key (no proof) |
| `encryptWithProof(epochId, aid, message, pubKey, k?)` | ElGamal encrypt + proof of knowledge of the randomness for `(epochId, aid)` |
| `decrypt(ciphertext, privKey)` | ElGamal decrypt via BSGS DLOG (values < 2^32) |
| `waitForCollectivePublicKeyHash(client, epochId, opts?)` | Wait for Live; return on-chain key hash |
| `waitForCombinedDecryption(client, epochId, aid, idx, opts?)` | Wait for on-chain decryption to complete |
| `demonstrateEncryptDecryptFlow(client, epochId, aid, pubKey, plaintext, idx)` | End-to-end demo flow |

### Proof helpers

| Export | Description |
|--------|-------------|
| `proveCiphertext(epochId, aid, ciphertext, r)` | Schnorr PoK of the ElGamal randomness (TE ciphertext in, RTE proof out) |
| `verifyCiphertextPoK(epochId, aid, c1x, c1y, c2x, c2y, pok)` | Verify a ciphertext PoK over on-chain (RTE) words — the committee's check |
| `proveOperator` / `verifyOperatorSchnorr` | Operator registration Schnorr PoK |
| `proveOrganizer` / `verifyOrganizerSchnorr` | Organizer (mode 1) Schnorr PoK |
| `verifyDleq` / `dleqChallenge` | Chaum-Pedersen partial-decryption proof check |

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
