# DAVINCI DKG SDK

TypeScript SDK for [DAVINCI DKG](../README.md): read and write the `DKGManager`, `DKGRegistry` and
`DKGAppManager` contracts, register applications, encrypt under their keys and follow
decryptions. Built on [viem](https://viem.sh); the ElGamal code runs in Node.js and in the browser.

[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](../LICENSE)

## Overview

- `DKGClient` reads epochs, pool keys, applications, ciphertexts, decryptions and the registry.
- `DKGWriter` extends it with transactions: create and abort epochs, register applications,
  submit ciphertexts, reveal organizer secrets, and the node-side calls.
- Helpers encrypt under an application key, wait for epochs and decryptions, and watch events.
- Codecs decode contribution and finalization transcripts from calldata and rebuild the
  share-commitment Merkle paths, checked against the vectors the Go code generates.

Points are handled in twisted Edwards form (the form `@zk-kit/baby-jubjub` uses) and converted at
the contract boundary.

## Quick start

The package is not published to npm yet. Build it from a checkout and add it as a local
dependency (Node.js 20, pnpm 10):

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg/sdk && pnpm install && pnpm build
cd /path/to/your-app && pnpm add viem file:/path/to/davinci-dkg/sdk
```

Register an automatic application on the Gnosis deployment, encrypt a value and wait for the
committee to decrypt it:

```ts
import { createPublicClient, createWalletClient, http } from 'viem';
import { gnosis } from 'viem/chains';
import { privateKeyToAccount } from 'viem/accounts';
import {
  AppMode, DKGWriter, EpochPhase, MAX_K, getNetwork, randomAid, waitForCombinedDecryption,
} from '@vocdoni/davinci-dkg-sdk';

const net = getNetwork('gnosis');
const transport = http(net.rpcUrls[0]);
const dkg = new DKGWriter({
  publicClient: createPublicClient({ chain: gnosis, transport }),
  walletClient: createWalletClient({ chain: gnosis, transport, account: privateKeyToAccount('0x...') }),
  managerAddress: net.managerAddress,
});

// The newest Live epoch that still has a free pool key.
let epochId: `0x${string}` | undefined;
for (const { id, epoch } of await dkg.getRecentEpochs(8)) {
  if (epoch.status === EpochPhase.Live && (await dkg.getPoolStatus(id)).nextIndex < MAX_K) {
    epochId = id;
    break;
  }
}
if (!epochId) throw new Error('no Live epoch with a free pool key');

const aid = randomAid();
await dkg.waitForTransaction(await dkg.registerApplication(epochId, aid, { mode: AppMode.Automatic }));

const { ciphertextIndex } = await dkg.encryptAndSubmit(epochId, aid, 42n);
await waitForCombinedDecryption(dkg, epochId, aid, ciphertextIndex, { timeoutMs: 600_000 });
console.log(await dkg.getPlaintext(epochId, aid, ciphertextIndex)); // 42n
```

`registryAddress` and `appManagerAddress` are optional in the client config; they are read from
the manager on first use.

## Usage

### Networks

`KNOWN_NETWORKS` mirrors the node's presets: chain id, `DKGManager`, deployment block and public
RPC endpoints. `getNetwork(name)` returns one (`gnosis`, `sepolia`), `NODE_DEFAULT_NETWORK` is the
node's default (`gnosis`), and `findNetwork(chainId, manager)` goes the other way. Sepolia has no
public endpoints built in.

### Applications

`registerApplication(epochId, aid, policy, skOrg?)` claims the epoch's next free pool key. Every
policy field is optional:

| Field | Default | Meaning |
|---|---|---|
| `mode` | `AppMode.OrganizerLocked` | `OrganizerLocked`: nothing decrypts until the organizer reveals `sk_org`. `Automatic`: no organizer key |
| `openSubmission` | `false` | Anyone may submit ciphertexts |
| `submitters` | `[]` | Exclusive allow-list, up to 32 addresses; empty and not open means the registrant only |
| `maxCiphertexts` | `0` | Cap on ciphertexts, `0` = unlimited |
| `notBeforeBlock`, `notAfterBlock` | `0n` | Submission window in blocks, `0n` = unbounded |
| `decryptNotBefore`, `decryptNotAfter` | `0n` | Decryption window in unix seconds, `0n` = unbounded |

An organizer-locked application needs an organizer secret:

```ts
import { randomOrganizerSecret } from '@vocdoni/davinci-dkg-sdk';

const skOrg = randomOrganizerSecret();          // persist it before registering
await dkg.registerApplication(epochId, aid, {}, skOrg);
// ... ciphertexts are submitted ...
await dkg.revealOrganizerSecret(epochId, aid, skOrg);   // once: opens every ciphertext of the application
```

**Losing `skOrg` before the reveal makes every ciphertext of the application permanently
undecryptable.** The SDK only sends `PK_org` and a proof of possession at registration. The
reveal is irreversible and application-wide, and an organizer secret must never be reused across
applications. The `aid` must be non-zero and below the BN254 scalar field (`randomAid()` produces
one); registration is open to anyone, so check that it succeeded.

### Encrypting

`encryptAndSubmit(epochId, aid, m)` reads the application key, encrypts and submits in one call
and returns `{ hash, receipt, ciphertextIndex, ciphertext }`. To encrypt without submitting:

```ts
import { encryptForApplication } from '@vocdoni/davinci-dkg-sdk';

const pkAid = await dkg.getApplicationKey(epochId, aid);   // P_j, or P_j + PK_org when locked
const ciphertext = await encryptForApplication(42n, pkAid);
const { ciphertextIndex } = await dkg.submitCiphertext(epochId, aid, ciphertext);
```

Plaintexts are small integers: the committee recovers values below 2⁵⁰, the SDK's own `decrypt`
below 2³². Ciphertexts can be added homomorphically before submission.

### Waiting and watching

```ts
import { decryptionProgress, waitForEpochPhase, watchCiphertextSubmitted, watchNewEpochs } from '@vocdoni/davinci-dkg-sdk';

await waitForEpochPhase(dkg, epochId, EpochPhase.Live, { timeoutMs: 300_000 });
const { combined, plaintext } = await decryptionProgress(dkg, epochId, aid, ciphertextIndex);
const open = await dkg.isDecryptionOpen(epochId, aid);   // window open and, if locked, revealed

const stop = watchNewEpochs(dkg, (id, creator) => console.log('epoch', id, 'by', creator));
watchCiphertextSubmitted(dkg, epochId, ({ aid, ciphertextIndex }) => console.log(aid, ciphertextIndex));
stop();
```

### Epochs and operators

`createEpoch({ threshold, committeeSize, minValidContributions, lotteryAlphaBps })` is
permissionless once `getNextEpochStartBlock()` is reached, and earlier when the newest epoch is
`Live` with at most one free key or `Aborted`; the values must respect `getEpochBounds()`.
`abortEpoch(epochId)` succeeds only on an epoch that can no longer progress. Nodes normally do
both. `registerKey`, `claimSlot`, `heartbeat`, `reactivate` and `reap` cover the operator side;
contributions, finalizations and partial decryptions come from `davinci-dkg-node`.

## API reference

### `DKGClient`

| Method | Returns |
|---|---|
| `getEpoch(epochId)`, `getRecentEpochs(limit?)` | Epoch record; the newest epochs |
| `epochNonce()`, `buildEpochId(nonce)`, `roundPrefix()` | Epoch numbering |
| `getEpochDurationBlocks()`, `getNextEpochStartBlock()`, `getLastEpochStartBlock()`, `getEpochBounds()` | Cadence and `createEpoch` floors |
| `selectedParticipants(epochId)`, `getContribution(epochId, address)` | Committee and contributions |
| `getPoolStatus(epochId)` | `{ nextIndex }`, the next free pool key (`MAX_K` once spent) |
| `getPoolKey(epochId, j)`, `getPoolKeys(epochId)`, `getPoolShareRoot(epochId, j)` | Pool keys and share-commitment roots of a Live epoch |
| `getApplication(epochId, aid)`, `getApplicationKey(epochId, aid)`, `getAppPoolIndex(epochId, aid)`, `getOrganizerPK(epochId, aid)` | Application record and key |
| `isDecryptionOpen(epochId, aid)` | Whether partials and combines are accepted now |
| `ciphertextCount(epochId, aid)`, `getCiphertextHash(epochId, aid, index)` | Submitted ciphertexts (indices start at 1) |
| `getPartialDecryption(epochId, aid, member, index)`, `getCombinedDecryption(epochId, aid, index)`, `getPlaintext(epochId, aid, index)` | Decryption state; `getPlaintext` is `0n` until combined |
| `getFinalizeTranscript(epochId)`, `getShareProof(epochId, j, member)` | Decoded finalization calldata; a member's Merkle path |
| `getNode(address)`, `nodeCount()`, `activeCount()`, `isActive(address)`, `inactivityWindow()`, `getRegistryNodes(fromBlock?)` | Registry |
| `get*Events(...)`, `getAllEpochEvents(epochId)`, `watchManagerEvents(handler)`, `watchRegistryEvents(handler)` | Event history and subscriptions |

### `DKGWriter`

| Method | Description |
|---|---|
| `registerApplication(epochId, aid, policy, skOrg?)` | Register an application and claim a pool key |
| `revealOrganizerSecret(epochId, aid, skOrg)` | Publish the organizer secret, once |
| `submitCiphertext(epochId, aid, ciphertext)`, `encryptAndSubmit(epochId, aid, m)` | Submit; waits for the receipt and returns the assigned index |
| `createEpoch(policy)`, `createRoundAndWait(policy)`, `abortEpoch(epochId)` | Epoch lifecycle |
| `registerKey(sk)`, `updateKey(sk)`, `heartbeat()`, `reactivate()`, `reap(operator)`, `claimSlot(epochId)` | Operator calls |
| `submitContribution(...)`, `finalizeEpoch(...)`, `submitPartialDecryption(...)`, `combineDecryption(...)` | Proof-carrying calls, as made by nodes |
| `waitForTransaction(hash)` | Wait for a receipt |

### Functions

| Export | Description |
|---|---|
| `encrypt(m, pubKey)`, `encryptForApplication(m, poolKey, pkOrg?)`, `decrypt(ct, sk)` | ElGamal on BabyJubJub |
| `buildElGamal()` | Key generation, encryption, point arithmetic and packing |
| `applicationKey(poolKey, pkOrg?)` | `P_j + PK_org`, or `P_j` without an organizer key |
| `randomAid()`, `randomOrganizerSecret()` | Fresh application id and organizer secret |
| `waitForEpochPhase`, `waitForDecryption`, `waitForCombinedDecryption`, `waitForPoolKey` | Polling helpers |
| `watchNewEpochs`, `watchEpochLive`, `watchCiphertextSubmitted`, `watchDecryptionCombined` | Event subscriptions; each returns an unsubscribe function |
| `decryptionProgress`, `networkSummary` | One-shot snapshots |
| `proveOperator`, `proveOrganizer`, `verifyOperatorSchnorr`, `verifyOrganizerSchnorr`, `verifyDleq` | Schnorr proofs of possession and the partial-decryption check |
| `decodeContributionCalldata`, `decodeFinalizeCalldata`, `merklePath`, `shareProof`, ... | Transcript and Merkle codecs |
| `KNOWN_NETWORKS`, `getNetwork`, `findNetwork`, `NODE_DEFAULT_NETWORK` | Network presets |
| `dkgManagerAbi`, `dkgRegistryAbi`, `dkgAppManagerAbi` | Contract ABIs |

## Development

```bash
pnpm install
pnpm build              # emit to dist/
pnpm check              # type-check sources and tests
pnpm test               # unit and vector tests
pnpm test:integration   # end-to-end against a local chain; needs Docker and the circuit artifacts
```

The integration tests start their own Anvil stack and use `cmd/sdk-test-fixture` to produce
proofs. See [CONTRIBUTING.md](../CONTRIBUTING.md).

## License

[GNU Affero General Public License v3.0](../LICENSE).
