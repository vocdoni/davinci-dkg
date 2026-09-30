# Deployments

The public deployments below are built into the node (`--network`), `dkgapp` and the SDK
(`KNOWN_NETWORKS`). Only the `DKGManager` address is configured anywhere; the other contracts are
read from it on chain.

## Gnosis Chain

Chain id 100. The default network of the node. Deployed at block 48,483,860 with the
[`circuits-v6`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v6) artifacts; every
contract is source-verified on [Gnosisscan](https://gnosisscan.io).

| Contract | Address |
|---|---|
| DKGManager | `0x9999F38Ff8Bf959E98Ddd5D4551f82775219c01B` |
| DKGAppManager | `0xd4d8f9708c380d81aec294b199081c5d2c782087` |
| DKGRegistry | `0x45ab8b64633076ddc020b12d1f1325fa55f629c5` |
| ContributionVerifier | `0x6d198bc613205957444b53a09bb22ed7bc650912` |
| FinalizeVerifier | `0xc354ea7f3ef6db4ca0b89a1a5a6395c2d6126b38` |
| PartialDecryptVerifier | `0x0f19886ee73fd74e3f88ce3a061490facd7561db` |
| DecryptCombineVerifier | `0x2980e664edef91f554cc75b15cb8eeea61586644` |

| Parameter | Value |
|---|---|
| `EPOCH_DURATION_BLOCKS` | 17,280 (about 24 h at 5 s blocks) |
| `COMMITTEE_SELECTION_BLOCKS` / `KEY_ASSEMBLY_BLOCKS` / `FINALIZE_GAP_BLOCKS` | 8 / 12 / 1 |
| `MIN_THRESHOLD` / `MIN_COMMITTEE_SIZE` / `MAX_LOTTERY_ALPHA_BPS` | 2 / 3 / 20000 |

With these windows an epoch is `Live` about two minutes after `createEpoch`.

## Sepolia

Chain id 11155111, a testnet. It has no public RPC endpoints built in: pass `--web3.rpc`.
Deployed at block 11,668,198 with the `circuits-v6` artifacts.

| Contract | Address |
|---|---|
| DKGManager | `0xc73b7a868eca6ac7e3e647e2665aa16a793cf551` |
| DKGAppManager | `0x7d5c48696b8c29638cc7819403d5989d9b5b8a94` |
| DKGRegistry | `0x6c4b8da67746677c3b0cb3122663186eea504737` |
| ContributionVerifier | `0x40954188b8b63c0829c34e8d9b9835416bf51f6e` |
| FinalizeVerifier | `0xa96519f22018ad8a9d91f8ec75431c2108d54ff3` |
| PartialDecryptVerifier | `0xf61356cb3cfdfca7cfe0a1dfc13eed9cad7d5a33` |
| DecryptCombineVerifier | `0xeef22b5b2174ce3d7c9d09b45892429216ecaefb` |

| Parameter | Value |
|---|---|
| `EPOCH_DURATION_BLOCKS` | 7,200 (about 24 h) |
| `COMMITTEE_SELECTION_BLOCKS` / `KEY_ASSEMBLY_BLOCKS` / `FINALIZE_GAP_BLOCKS` | 100 / 150 / 10 |
| `MIN_THRESHOLD` / `MIN_COMMITTEE_SIZE` / `MAX_LOTTERY_ALPHA_BPS` | 2 / 3 / 20000 |
| `INACTIVITY_WINDOW` | 50,400 blocks |

## Deploying your own

`make solidity-deploy` runs `solidity/deploy_all.sh`, which runs `forge test`, deploys the four
verifiers, `DKGRegistry`, `DKGManager` and `DKGAppManager`, wires them together and writes the
addresses to `solidity/.last_deployed_addresses.env`. It reads its settings from the environment
or from `.env` at the repository root:

| Variable | Required | Description |
|---|---|---|
| `RPC_URL`, `CHAIN_ID`, `PRIVATE_KEY` | yes | Target chain and deployer key |
| `MIN_THRESHOLD`, `MIN_COMMITTEE_SIZE`, `MAX_LOTTERY_ALPHA_BPS` | yes | Floors `createEpoch` enforces |
| `EPOCH_DURATION_BLOCKS` | no | Epoch cadence (default 100) |
| `COMMITTEE_SELECTION_BLOCKS`, `KEY_ASSEMBLY_BLOCKS`, `FINALIZE_GAP_BLOCKS` | no | Preparation windows (default 25, 25, 5) |
| `INACTIVITY_WINDOW` | no | Registry inactivity window in blocks (default 50,400) |
| `ETHERSCAN_API_KEY`, `VERIFIER_URL` | no | Verify the sources on an Etherscan-compatible explorer |
| `SKIP_TESTS=1` | no | Skip the `forge test` gate |

The window defaults suit a local chain; size them for the target chain's block time. The
verifiers embed the verifying keys of the circuits in `config/circuit_artifacts.go`, so nodes
running a release built from the same commit can serve the deployment. Point them at it with
`--manager <DKGManager> --web3.rpc <endpoints>`.
