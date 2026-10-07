# DAVINCI DKG

Non-interactive distributed key generation for EVM chains. A rotating committee of operators
generates threshold ElGamal keys, applications encrypt under them, and the committee decrypts on
chain when the application allows it. It is the key layer of [DAVINCI](https://davinci.vote) and
works for any contract that needs a public key whose secret nobody holds.

[![Build and Test](https://github.com/vocdoni/davinci-dkg/actions/workflows/main.yml/badge.svg)](https://github.com/vocdoni/davinci-dkg/actions/workflows/main.yml)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](LICENSE)

## Overview

Operators run `davinci-dkg-node` and register in an on-chain registry. For every epoch a public
lottery draws a committee of `n` of them. Each member submits one contribution that deals 16
independent pool keys, and one finalization proof stores all 16 public keys on chain; from then
on any `t` members can decrypt under any of them. Every protocol step (contribution,
finalization, partial decryption, combine) is a single transaction checked on chain by a Groth16
proof, so there are no complaint rounds and no dispute phase.

Applications register against a live epoch and each claims one pool key. An application is
either **organizer-locked**, where its key also includes an organizer key and nothing decrypts
until the organizer publishes that secret, or **automatic**, where the committee decrypts as soon
as the decryption window opens. Whoever the application's policy admits submits ElGamal
ciphertexts; committee nodes post partial decryptions, one of them combines them, and the
plaintext is stored on chain for everyone.

```
createEpoch ─► claimSlot (lottery) ─► submitContribution ─► finalizeEpoch ─► Live
registerApplication ─► submitCiphertext ─► [revealOrganizerSecret] ─► t × submitPartialDecryption ─► combineDecryption
```

| Component | Path | Purpose |
|---|---|---|
| `davinci-dkg-node` | `cmd/davinci-dkg-node`, `node/` | Operator daemon: registers, joins committees, deals, finalizes, decrypts |
| `dkgapp` | `cmd/dkgapp` | Command line client for applications: register, encrypt, reveal, read plaintexts |
| Contracts | `solidity/` | `DKGRegistry`, `DKGManager`, `DKGAppManager` and four Groth16 verifiers |
| Circuits | `circuits/` | Groth16 circuits over BN254 (gnark) |
| TypeScript SDK | `sdk/` | `@vocdoni/davinci-dkg-sdk`: read and write the contracts, encrypt in the browser |
| Explorer | `ui/` | Web explorer and playground for any deployment |

In DAVINCI the [ProcessRegistry](https://github.com/vocdoni/davinci-contracts) registers one DKG
application per voting process, voters encrypt their ballots under its key, the
[sequencer](https://github.com/vocdoni/davinci-sequencer) aggregates them, and the committee
decrypts the final tally. The protocol is described in [docs/protocol.md](docs/protocol.md).

## Quick start

Run a node on the Gnosis Chain deployment. You need Docker and an EVM account with a little xDAI
for gas.

```bash
git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg
cp .env.example .env        # set DAVINCI_DKG_PRIVKEY=0x...
docker compose --profile node up -d
docker compose --profile node logs -f node
```

On first start the node derives its BabyJubJub key from the operator key and registers it (one
transaction), downloads the pinned circuit artifacts (about 1 GB, verified against hashes built
into the binary) and then takes part in every epoch it is drawn into. The compose profile also
starts Watchtower, which keeps the node on the latest stable release.

## Usage

### Running a node

A node needs 2 cores and 4 GB of RAM (8 GB recommended), about 1.5 GB of disk for the artifacts,
and an operator key funded for gas. Every setting is a flag or a `DAVINCI_DKG_*` environment
variable (dots and dashes become underscores); `davinci-dkg-node --help` lists them all.

| Variable | Flag | Default | Description |
|---|---|---|---|
| `DAVINCI_DKG_PRIVKEY` | `--privkey` | | Operator key (hex) that signs every transaction. Required |
| `DAVINCI_DKG_WEB3_RPC` | `--web3.rpc` | the network's public endpoints | Comma-separated JSON-RPC endpoints. Set at least two: the node rotates off rate-limited ones |
| `DAVINCI_DKG_NETWORK` | `--network` | `gnosis` | Built-in deployment: `gnosis` or `sepolia` (no public endpoints built in) |
| `DAVINCI_DKG_MANAGER` | `--manager` | | `DKGManager` address of any other deployment; overrides the network |
| `DAVINCI_DKG_DATADIR` | `--datadir` | `~/.davinci-dkg` | Node state, one subdirectory per deployment |
| `DAVINCI_DKG_ARTIFACTS_DIR` | | `~/.davinci/artifacts` | Circuit artifact cache |
| `DAVINCI_DKG_AUTO_CREATE_EPOCHS` | `--auto-create-epochs` | `true` | Create the next epoch when it is due and abort dead ones |
| `DAVINCI_DKG_LOG_LEVEL` | `--log.level` | `info` | `debug`, `info`, `warn` or `error` |

The node checks that the endpoints serve the network's chain and refuses to start otherwise
(`--manager` skips the check). It serves no HTTP; watch it through its logs or the explorer.
Prebuilt binaries for Linux are attached to each [release](https://github.com/vocdoni/davinci-dkg/releases),
and `go build ./cmd/davinci-dkg-node` builds one from source. Resource figures, upgrades, the data
directory and hosting on Railway are covered in [docs/node.md](docs/node.md).

### Using the DKG from an application

Each application is identified by a 32-byte `aid = salt << 160 | registrant`: the low 160 bits
are the address that registers it (the contract refuses an id in another account's namespace),
and a salt below 2⁹² keeps it inside the BN254 scalar field. `dkgapp register` and the SDK's
`randomAid(account)` build one for you. An organizer-locked application prints an organizer secret at
registration: **store it**. It is not derivable from anything on chain, and without it the
application can never be decrypted.

`dkgapp` covers the whole flow from the command line:

```bash
go build -o dkgapp ./cmd/dkgapp
export DAVINCI_DKG_NETWORK=gnosis DAVINCI_DKG_PRIVKEY=0x...

./dkgapp epoch                                                  # newest epoch and its key pool
./dkgapp register                                               # prints the aid, the epoch id and the organizer secret
AID=0x...                                                       # the aid it printed
./dkgapp encrypt   -epoch <epoch> -aid $AID -m 42               # prints the ciphertext index
./dkgapp reveal    -epoch <epoch> -aid $AID -org-secret <secret>   # opens the application, once
./dkgapp plaintext -epoch <epoch> -aid $AID -index 1 -wait 10m
```

`register -mode automatic` creates an application without an organizer key, which the committee
decrypts without a `reveal`. `register` also takes `-submitters 0xA,0xB` (an exclusive allow-list
of up to 32 addresses) or `-open` (anyone may submit); with neither only the registrant submits.
`-max` caps the number of ciphertexts and `-decrypt-from` / `-decrypt-until` bound the decryption
window (RFC 3339 or a duration from now, such as `24h`). Plaintexts are small integers: the
committee recovers values below 2⁵⁰.

The same flow in TypeScript, with the [SDK](sdk/README.md):

```ts
import { createPublicClient, createWalletClient, http } from 'viem';
import { gnosis } from 'viem/chains';
import { privateKeyToAccount } from 'viem/accounts';
import { AppMode, DKGWriter, getNetwork, randomAid, waitForCombinedDecryption } from '@vocdoni/davinci-dkg-sdk';

const net = getNetwork('gnosis');
const publicClient = createPublicClient({ chain: gnosis, transport: http(net.rpcUrls[0]) });
const walletClient = createWalletClient({
  chain: gnosis, transport: http(net.rpcUrls[0]), account: privateKeyToAccount('0x...'),
});
const dkg = new DKGWriter({ publicClient, walletClient, managerAddress: net.managerAddress });

const epochId = '0x...';   // a Live epoch with a free pool key, see `dkgapp epoch`
const aid = randomAid(dkg.walletClient.account!.address);   // salt << 160 | your address
await dkg.waitForTransaction(await dkg.registerApplication(epochId, aid, { mode: AppMode.Automatic }));

const { ciphertextIndex } = await dkg.encryptAndSubmit(epochId, aid, 42n);
await waitForCombinedDecryption(dkg, epochId, aid, ciphertextIndex, { timeoutMs: 600_000 });
console.log(await dkg.getPlaintext(epochId, aid, ciphertextIndex)); // 42n
```

A contract reads the key with `DKGAppManager.getApplicationKey(epochId, aid)` and submits with
`DKGManager.submitCiphertext(epochId, aid, c1x, c1y, c2x, c2y)`; see
[solidity/README.md](solidity/README.md). `submitCiphertext` takes no proof, so an
application whose participants could gain by copying or shifting someone else's ciphertext must
bind ciphertexts to their authors itself. [docs/use-cases.md](docs/use-cases.md) walks through
auctions, scheduled disclosure, lotteries and private aggregation.

### Running the explorer

The explorer in `ui/` is a static web app that reads a deployment straight from a JSON-RPC
endpoint, with no backend. `make ui-dev` serves it for the Gnosis deployment on
`http://localhost:5174`; [ui/README.md](ui/README.md) covers configuration and hosting.

### Deployments

| Network | Chain id | `DKGManager` | Circuits |
|---|---|---|---|
| Gnosis Chain (default) | 100 | `0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF` | [`circuits-v6`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v6) |
| Sepolia (testnet) | 11155111 | `0xc73b7a868eca6ac7e3e647e2665aa16a793cf551` | [`circuits-v6`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v6) |

Both are built into the node, `dkgapp` and the SDK. The other contracts are resolved from the
manager on chain; their addresses, the epoch parameters and how to deploy your own are in
[docs/deployments.md](docs/deployments.md).

## Documentation

- [docs/protocol.md](docs/protocol.md): epochs, committee selection, application keys, threshold
  decryption, proofs and contracts.
- [docs/node.md](docs/node.md): operating a node: configuration, resources, upgrades, Railway.
- [docs/deployments.md](docs/deployments.md): contract addresses and deploying the contracts.
- [docs/use-cases.md](docs/use-cases.md): application patterns beyond voting.
- [docs/pool-keys.md](docs/pool-keys.md): normative transcript encodings and proof statements.
- [BENCHMARKS.md](BENCHMARKS.md): constraints, proving times, memory, gas and RPC load.
- [sdk/README.md](sdk/README.md), [ui/README.md](ui/README.md),
  [solidity/README.md](solidity/README.md): the SDK, the explorer and the contracts.

## Development

Requires Go 1.25+, Foundry, Node.js 20 with pnpm 10, and Docker for the integration tests.

```bash
make build                                            # Go binaries
go test ./node/... ./crypto/... ./web3/...            # fast unit tests
(cd solidity && forge build && forge test)            # contracts
(cd sdk && pnpm install && pnpm build && pnpm test)   # SDK
make testnet-up                                       # local Anvil chain with three nodes
```

[CONTRIBUTING.md](CONTRIBUTING.md) covers the full test suites, the circuit pipeline and the
parts that must stay in sync across Go, Solidity and TypeScript.

## License

[GNU Affero General Public License v3.0](LICENSE).
