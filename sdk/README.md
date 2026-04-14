# Davinci DKG SDK

TypeScript SDK for interacting with the Davinci DKG smart contracts. Provides read/write access to the `DKGManager` and `DKGRegistry` contracts, ElGamal encryption on BabyJubJub, and utilities for monitoring round lifecycle and decryption.

## Installation

```sh
npm install @vocdoni/davinci-dkg-sdk
# or
pnpm add @vocdoni/davinci-dkg-sdk
```

**Peer dependencies:** `viem`

## Quick start

```ts
import { createPublicClient, createWalletClient, http, defineChain } from 'viem';
import { privateKeyToAccount } from 'viem/accounts';
import {
  DKGClient,
  DKGWriter,
  buildElGamal,
  waitForRoundStatus,
  RoundStatus,
  buildRoundId,
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
const nonce       = await client.roundNonce();

// Round details
const round = await client.getRound(roundId);
console.log(round.status);           // 1=Registration, 2=Contribution, 3=Finalized …
console.log(round.policy.threshold);

// Participants and contributions
const participants = await client.selectedParticipants(roundId);
const contrib      = await client.getContribution(roundId, participantAddress);

// Decryption state
const partial  = await client.getPartialDecryption(roundId, participant, ciphertextIndex);
const combined = await client.getCombinedDecryption(roundId, ciphertextIndex);

// Registry
const node = await client.getNode(operatorAddress);
```

## Writing transactions

```ts
// Create a round (organizer role)
const currentBlock = await client.blockNumber();
const hash = await writer.createRound({
  threshold:                 2,
  committeeSize:             3,
  minValidContributions:     2,
  lotteryAlphaBps:           15000,   // 1.5× over-subscription
  seedDelay:                 1,        // blocks before seed is available
  registrationDeadlineBlock: currentBlock + 25n,
  contributionDeadlineBlock: currentBlock + 50n,
  disclosureAllowed:         false,
});
await writer.waitForTransaction(hash);

// Derive the round ID from the nonce that was current at round creation
const nonce  = await client.roundNonce();
const roundId = await client.buildRoundId(nonce - 1n);

// Register a DKG node
await writer.registerKey(pubX, pubY);

// Claim a slot (DKG node role — after seedDelay blocks have passed)
await writer.claimSlot(roundId);

// Abort (organizer only, when below minimum contributions)
await writer.abortRound(roundId);
```

## Monitoring rounds

```ts
import { waitForRoundStatus, waitForDecryption, watchNewRounds } from '@vocdoni/davinci-dkg-sdk';

// Poll until a round reaches Finalized status
await waitForRoundStatus(client, roundId, RoundStatus.Finalized, {
  intervalMs: 2000,
  timeoutMs:  120_000,
});

// Poll until ciphertext 1 has been decrypted on-chain
const record = await waitForDecryption(client, roundId, 1);
console.log(record.completed); // true

// Subscribe to new rounds in real time (returns unsubscribe fn)
const unsub = watchNewRounds(client, (roundId, organizer) => {
  console.log('new round', roundId, 'by', organizer);
});
unsub(); // stop watching
```

## ElGamal encryption

The SDK exposes ElGamal encryption/decryption on the BabyJubJub curve, backed by [`@zk-kit/baby-jubjub`](https://github.com/zk-kit/zk-kit/tree/main/packages/baby-jubjub) — pure TypeScript, synchronous, browser-native.

```ts
import { buildElGamal } from '@vocdoni/davinci-dkg-sdk';

const elgamal = await buildElGamal();

// Generate a key pair
const { privKey, pubKey } = elgamal.generateKeyPair();

// Encrypt a small integer (plaintext must fit in the BabyJubJub subgroup order)
const ciphertext = elgamal.encrypt(42n, pubKey);
// { c1: [bigint, bigint], c2: [bigint, bigint] }

// Decrypt (brute-force DLOG — works for values < 2^20)
const plaintext = elgamal.decrypt(ciphertext, privKey);
console.log(plaintext); // 42n

// Point arithmetic
const doubled = elgamal.addPoint(pubKey, pubKey);
const via2    = elgamal.mulPoint(pubKey, 2n);

// Compact serialization (compressed point ↔ bigint)
const packed   = elgamal.packPoint(pubKey);   // bigint
const unpacked = elgamal.unpackPoint(packed); // [bigint, bigint]
```

The `flow` module provides convenience wrappers for the most common patterns:

```ts
import { encrypt, decrypt, waitForCollectivePublicKeyHash } from '@vocdoni/davinci-dkg-sdk';

// Encrypt with the collective public key
const ciphertext = await encrypt(42n, collectivePubKey);

// Wait for a round to finalize and return the on-chain public key hash
const hash = await waitForCollectivePublicKeyHash(client, roundId);
```

In the real DKG protocol the private key is never held by a single party. To decrypt a ciphertext:

1. DKG nodes each call `submitPartialDecryption` on the `DKGManager` contract.
2. Once the threshold is met, any party calls `combineDecryption`.
3. The `DecryptionCombined` event is emitted and `getCombinedDecryption` returns `completed: true`.

## Full DKG flow overview

```
[Organizer] createRound(policy)
               │
               ▼  (seedDelay blocks mined)
[DKG Node]  claimSlot(roundId)           ← lottery via on-chain blockhash seed
               │
               ▼  (registration deadline passes)
[DKG Node]  submitContribution(...)      ← ZK proof of VSS shares
               │
               ▼  (contribution deadline passes)
[DKG Node]  finalizeRound(...)           ← ZK proof aggregating all commitments
               │                            collectivePublicKeyHash emitted
               ▼  Round.status = Finalized
[Anyone]    encrypt(plaintext, collectivePubKey)  ← ElGamal; ciphertext published off-chain
               │
               ▼
[DKG Node]  submitPartialDecryption(...)
               │
               ▼  (threshold met)
[DKG Node]  combineDecryption(...)       ← DecryptionCombined event
               │
               ▼  Round.status = Completed
[Anyone]    getCombinedDecryption(roundId, idx)  ← completed: true
```

> **Collective public key:** The `RoundFinalized` event contains `collectivePublicKeyHash` — the keccak256 hash of the collective key point. The actual `(x, y)` coordinates are encoded in the `transcript` argument of the `finalizeRound` calldata and can be recovered by decoding that transaction.

## API reference

### `DKGClient`

| Method | Description |
|--------|-------------|
| `getRound(roundId)` | Full round struct |
| `selectedParticipants(roundId)` | Addresses that claimed a slot |
| `getContribution(roundId, address)` | Contribution record |
| `getPartialDecryption(roundId, address, idx)` | Partial decryption record |
| `getCombinedDecryption(roundId, idx)` | Combined decryption record |
| `getRevealedShare(roundId, address)` | Revealed share (disclosure mode) |
| `getNode(address)` | Registry node record |
| `nodeCount()` / `activeCount()` | Registry stats |
| `isActive(address)` | Node liveness check |
| `blockNumber()` | Current block |
| `roundNonce()` | Next round nonce |
| `buildRoundId(nonce)` | Derive round ID from nonce |
| `getRoundCreatedEvents(opts?)` | Historical RoundCreated logs |
| `getRoundFinalizedEvents(roundId)` | Historical RoundFinalized logs |
| `getAllRoundEvents(roundId, fromBlock?)` | All DKGManager events for a specific round |
| `getRecentRounds(limit?)` | Most recent N rounds (default 20) as `RoundEntry[]` |
| `getRegistryNodes(fromBlock?)` | All registered nodes via NodeRegistered events |
| `roundPrefix()` | Fetch the immutable ROUND_PREFIX constant |
| `watchManagerEvents(handler)` | Live event stream (returns unsubscribe fn) |
| `watchRegistryEvents(handler)` | Live registry events |

### `DKGWriter` (extends `DKGClient`)

All `DKGClient` methods plus:

| Method | Description |
|--------|-------------|
| `createRound(policy)` | Create a new DKG round |
| `claimSlot(roundId)` | Claim a committee slot |
| `extendRegistration(roundId)` | Extend registration deadline (organizer) |
| `submitContribution(...)` | Submit VSS contribution + ZK proof |
| `finalizeRound(...)` | Finalize round + ZK proof |
| `submitPartialDecryption(...)` | Submit partial decryption + ZK proof |
| `combineDecryption(...)` | Combine partial decryptions + ZK proof |
| `submitRevealedShare(...)` | Disclose share (disclosure mode) |
| `reconstructSecret(...)` | Reconstruct secret from shares |
| `abortRound(roundId)` | Abort a round (organizer) |
| `registerKey(pubX, pubY)` | Register a DKG node in the registry |
| `updateKey(pubX, pubY)` | Update node public key |
| `heartbeat()` | Refresh node liveness |
| `reactivate()` | Rejoin after being reaped |
| `reap(operator)` | Permissionlessly reap a stale node |
| `waitForTransaction(hash)` | Wait for tx receipt |
| `createRoundAndWait(policy)` | createRound + wait |

### Monitor utilities

| Export | Description |
|--------|-------------|
| `waitForRoundStatus(client, roundId, status, opts?)` | Poll until round status reached |
| `waitForDecryption(client, roundId, idx, opts?)` | Poll until decryption complete |
| `watchNewRounds(client, onRound, fromBlock?)` | Subscribe to new rounds |
| `watchRoundFinalized(client, roundId, onFinalized)` | Subscribe to finalization |
| `watchDecryptionCombined(client, roundId, idx, onCombined)` | Subscribe to decryption |
| `networkSummary(client)` | Block, node counts, round nonce |

### Flow helpers

High-level helpers that compose the primitives above:

| Export | Description |
|--------|-------------|
| `encrypt(message, pubKey, k?)` | ElGamal encrypt via collective public key |
| `decrypt(ciphertext, privKey)` | ElGamal decrypt (brute-force DLOG, values < 2^20) |
| `waitForCollectivePublicKeyHash(client, roundId, opts?)` | Wait for Finalized; return on-chain key hash |
| `waitForCombinedDecryption(client, roundId, idx, opts?)` | Wait for on-chain decryption to complete |
| `demonstrateEncryptDecryptFlow(client, roundId, pubKey, plaintext, idx)` | End-to-end demo flow |

### ElGamal interface

| Method | Description |
|--------|-------------|
| `generateKeyPair()` | Return `{ privKey: bigint, pubKey: BabyJubPoint }` |
| `randomScalar()` | Uniformly random scalar in the BabyJubJub subgroup |
| `encrypt(msg, pubKey, k?)` | Encrypt a small integer; `k` is the ephemeral scalar |
| `decrypt(ciphertext, privKey)` | Brute-force DLOG (works for msg < 2^20) |
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
