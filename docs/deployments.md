# Deployments

The public deployments below are built into the node (`--network`), `dkgapp` and the SDK
(`KNOWN_NETWORKS`). Only the `DKGManager` address is configured anywhere; the other contracts are
read from it on chain.

## OP Mainnet

Chain id 10. The default network of the node, `dkgapp` and the SDK since the Gnosis deployment
retired. Deployed on 2026-10-15 at block 158,245,366 with the
[`circuits-v6`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v6) artifacts, from
a dedicated deployer key, [`0x5eb38f2D16CEaabFF6a63C17505a14b45624f537`](https://optimistic.etherscan.io/address/0x5eb38f2D16CEaabFF6a63C17505a14b45624f537),
used for nothing else (not the operators' keys). All nine transactions landed in that block. The
runtime code of every contract equals the Gnosis deployment's with the immutables masked, and the
four verifiers equal Gnosis's byte for byte. Sources are verified on Sourcify (exact match; the
manager a full match); not yet on Optimistic Etherscan.

| Contract | Address |
|---|---|
| DKGManager | `0x9E799A1D22710056E62c849e3FFE19Dd01db2D41` |
| DKGAppManager | `0x1e7f0bec007828e2e0a1a469d9d3e148888fc39a` |
| DKGRegistry | `0x5ac7ec4c9a2100f705d108841a170074fe95a998` |
| ContributionVerifier | `0xc05077531c6d7dd0cd0265b7dc677676ad12e23f` |
| FinalizeVerifier | `0xa10fe6cc25b71493a5073354328296b4218bce39` |
| PartialDecryptVerifier | `0x6a5fc58dded56a7093c776032fc87c98d3e5c933` |
| DecryptCombineVerifier | `0x0610a2f9cc701e3faad5817b11f89040b934205c` |

Deployment block: 158,245,366 (2026-10-15).

| Transaction | Hash | Gas |
|---|---|---:|
| ContributionVerifier | `0xeeb5e45a61ea98509620f0a32d77b13858e0e4c32094f73d84c19036d22bfa1d` | 1,286,050 |
| FinalizeVerifier | `0x99679901915009a2159c5a00893d1b925aebed57d20581defa6a487f02da3076` | 1,242,131 |
| PartialDecryptVerifier | `0x700a72b446b5ad8b0dcffc91c71bc324e6c4ee71c10bd041f677588714644b34` | 1,580,373 |
| DecryptCombineVerifier | `0x3faa94dd858a8701b6393d319e38efd19437f6c87690b70f6da025f8c5ce5ca2` | 1,322,348 |
| DKGRegistry | `0xc3851eca0ea2843ea894ac6f64258842d81c007d0258ad5ccf4316293c97a45d` | 1,230,014 |
| DKGManager | `0x03eb26001de8009cf3cfd9903e4a12fb5e2e35ffe8fd82de61847cb857b47aa3` | 4,070,003 |
| DKGRegistry.setManager | `0x9950b864049ff33c7cb87e0331e5676c3de6b078498b0a0368fd01d12d13b799` | 45,186 |
| DKGAppManager | `0x705ce3a8bb7cd739c300e416e89eab424f6c146299dd2b139c60b57897f79578` | 2,035,220 |
| DKGManager.setAppManager | `0xc5e44b571d8e3494698ca27feeab59fb720f43dbe4369df6a1cc030a11f56625` | 44,179 |

12,855,504 gas at 2,422 wei per gas (0.000000031 ETH) plus 0.00000058 ETH of OP L1 data fee.
`setManager` and `setAppManager` are one-shot, so the suite has no owner.

OP has 2-second blocks, so the windows are the Gnosis ones (5-second blocks) in the same
wall-clock time:

| Parameter | Value |
|---|---|
| `EPOCH_DURATION_BLOCKS` | 43,200 (24 h at 2 s blocks) |
| `COMMITTEE_SELECTION_BLOCKS` / `KEY_ASSEMBLY_BLOCKS` / `FINALIZE_GAP_BLOCKS` | 90 / 30 / 1 (3 min / 1 min / 2 s) |
| `MIN_THRESHOLD` / `MIN_COMMITTEE_SIZE` / `MAX_LOTTERY_ALPHA_BPS` | 2 / 3 / 20000 |
| `INACTIVITY_WINDOW` (registry) | 126,000 blocks (about 2.9 days) |

The preset also carries the node's `--decrypt-lookback-blocks` default for OP, 126,000 blocks: the
same span as the registry's inactivity window. Without a preset the default is 50,400 blocks.

The preset's public endpoints (`https://mainnet.optimism.io`, `https://optimism-rpc.publicnode.com`)
are rate-limited, and publicnode refuses state older than about 600 blocks and old log ranges
(`mainnet.optimism.io` served 10,000-block log ranges 136,000 blocks back on 2026-10-15). A
long-lived node lists a private provider first in
`--web3.rpc` and the public ones after it as failover; the node rotates off refusing endpoints.

## Gnosis Chain (retired)

Chain id 100. The former default network of the node, no longer built in: run it with
`--manager` and `--web3.rpc`, and the explorer still serves it until it moves to OP. Deployed at block 48,632,905 with the
[`circuits-v6`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v6) artifacts; every
contract is source-verified on [Gnosisscan](https://gnosisscan.io) and Sourcify. It replaced the
first Gnosis deployment on 2026-10-07 to ship the application-id namespace of
[#14](https://github.com/vocdoni/davinci-dkg/issues/14).

| Contract | Address |
|---|---|
| DKGManager | `0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF` |
| DKGAppManager | `0x81bac6b9aae85311741204c22cbaab96f03b567a` |
| DKGRegistry | `0x393049828bc565152c223730ce67574f15a7a16a` |
| ContributionVerifier | `0x67b7c4daa3db8b84e817d1a3d57cb8c40f59935f` |
| FinalizeVerifier | `0x682ffa2a6e049e0dc889ceb7038fffc3bbf36d7d` |
| PartialDecryptVerifier | `0x7ee28e810086590192e93e6046d90fe4e97f0f60` |
| DecryptCombineVerifier | `0x1fd445c2e5700a0791ca70ac1bc3447b32acc56e` |

| Parameter | Value |
|---|---|
| `EPOCH_DURATION_BLOCKS` | 17,280 (about 24 h at 5 s blocks) |
| `COMMITTEE_SELECTION_BLOCKS` / `KEY_ASSEMBLY_BLOCKS` / `FINALIZE_GAP_BLOCKS` | 36 / 12 / 1 |
| `MIN_THRESHOLD` / `MIN_COMMITTEE_SIZE` / `MAX_LOTTERY_ALPHA_BPS` | 2 / 3 / 20000 |

With these windows an epoch is `Live` about four minutes after `createEpoch`. The selection
window was 8 blocks on the first deployment, and 9 of its first 20 epochs aborted because the
committee did not fill in time.

**Retired:** DKGManager `0x9999F38Ff8Bf959E98Ddd5D4551f82775219c01B` (block 48,483,860, registry
`0x45ab8b64633076ddc020b12d1f1325fa55f629c5`, app manager `0xd4d8f9708c380d81aec294b199081c5d2c782087`).
Its applications and plaintexts stay on chain, but no node serves it since v0.10.0.

## Sepolia

Chain id 11155111, a testnet. It has no public RPC endpoints built in: pass `--web3.rpc`.
Deployed at block 11,668,198 with the `circuits-v6` artifacts. **This deployment predates the
application-id namespace fix of [#14](https://github.com/vocdoni/davinci-dkg/issues/14)**: the
contract does not enforce that `aid`'s low 160 bits match the registrant's address.

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

## Moving the OP preset

A change to the contracts' rules (such as the application-id namespace of
[#14](https://github.com/vocdoni/davinci-dkg/issues/14)) needs a fresh deployment. The old one keeps
its applications and plaintexts on chain, but its nodes leave it once the preset moves. The same
steps took the preset from its placeholder to the first OP deployment.

1. **Deploy** from a clean checkout of the release commit, with Foundry on `PATH`. Do not run
   `make circuits` first: the committed verifiers must stay those of the pinned `circuits-v6`
   artifacts. Write the settings to a `0600` `solidity/.env` and delete it afterwards. These are the
   OP Mainnet parameters:

   ```bash
   RPC_URL=https://mainnet.optimism.io
   CHAIN_ID=10
   PRIVATE_KEY=<deployer key>
   MIN_THRESHOLD=2
   MIN_COMMITTEE_SIZE=3
   MAX_LOTTERY_ALPHA_BPS=20000
   EPOCH_DURATION_BLOCKS=43200
   COMMITTEE_SELECTION_BLOCKS=90
   KEY_ASSEMBLY_BLOCKS=30
   FINALIZE_GAP_BLOCKS=1
   INACTIVITY_WINDOW=126000
   ETHERSCAN_API_KEY=<Etherscan v2 key, for Optimistic Etherscan verification>
   ```

   ```bash
   make solidity-deploy && rm solidity/.env
   ```

   The deploy script prints `RPC_URL` and writes it to `.last_deployed_addresses.env`, so use a
   public endpoint here, never a provider URL that embeds a key.
2. **Check it.** `solidity/.last_deployed_addresses.env` lists the seven addresses. Read the
   manager's deployment block and the parameters back:

   ```bash
   . solidity/.last_deployed_addresses.env
   jq -r --arg m "${MANAGER,,}" '.receipts[] | select((.contractAddress // "" | ascii_downcase) == $m)
     | .blockNumber' solidity/broadcast/DeployAll.s.sol/10/run-latest.json | xargs cast to-dec
   for f in "REGISTRY()(address)" "appManager()(address)" "EPOCH_DURATION_BLOCKS()(uint256)" \
     "MIN_THRESHOLD()(uint16)" "MIN_COMMITTEE_SIZE()(uint16)" "MAX_LOTTERY_ALPHA_BPS()(uint16)"; do
     cast call "$MANAGER" "$f" --rpc-url https://mainnet.optimism.io; done
   cast call "$REGISTRY" "INACTIVITY_WINDOW()(uint64)" --rpc-url https://mainnet.optimism.io
   ```

3. **Move the preset.** `scripts/move-optimism-preset.sh solidity/.last_deployed_addresses.env
   <deploy block> [YYYY-MM-DD]` fills the `optimism` preset in `config/networks.go` and
   `sdk/src/networks.ts` (manager and start block), the OP Mainnet tables in this file and the
   README networks row. It edits nothing else; check with `bash scripts/move-optimism-preset.test.sh`.
   Review the diff, run the checks it prints, list a replaced manager under a "Retired" note here,
   and commit. The explorer (`ui/public/config.json`, `scripts/render-ui-config.sh`,
   `ui/.do/davinci-dkg-ui.yaml`) moves to OP separately, once the deployment is live: on DigitalOcean
   App Platform set `CHAIN_ID=10`, `CHAIN_NAME=optimism`, `EXPLORER_URL=https://optimistic.etherscan.io`,
   `RPC_URL=https://mainnet.optimism.io`, `MANAGER_ADDRESS` and `DEPLOY_BLOCK`.
4. **Release** a stable `vX.Y.Z`. `latest` moves, so compose nodes under Watchtower switch to the
   new manager on their own, register in its registry at start (`EnsureRegistered`) and start
   from an empty `<datadir>/10-<manager>`. Railway has no Watchtower: roll every service with
   `scripts/railway-roll-node.sh <service id> ghcr.io/vocdoni/davinci-dkg:vX.Y.Z`. For the
   explorer on DigitalOcean App Platform: the app's build-time env (`MANAGER_ADDRESS`,
   `DEPLOY_BLOCK`) overrides the repo file, so apply the values from `ui/.do/davinci-dkg-ui.yaml`
   to the app's environment spec, trigger a redeploy, and purge the Cloudflare cache.
5. **Verify.** Every node logs `self: registry row` and `node running` against the new manager
   (`scripts/railway-node-status.sh <service id>`). `cast call $REGISTRY "activeCount()(uint64)"`
   reaches the fleet size. The first epoch goes `Live` about four minutes after a node creates it (the windows are 3 min + 1 min).
   It needs `MIN_COMMITTEE_SIZE` (3) active operators, all claiming within the 90-block selection
   window, and at least `minValidContributions` contributions. Then register an automatic
   application and decrypt a value end to end with `dkgapp`.
