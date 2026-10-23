# Deployments

The public deployments below are built into the node (`--network`), `dkgapp` and the SDK
(`KNOWN_NETWORKS`). Only the `DKGManager` address is configured anywhere; the other contracts are
read from it on chain.

## OP Mainnet

Chain id 10. The default network of the node, `dkgapp` and the SDK since the Gnosis deployment
retired. Deployed on 2026-10-23 at block 158,592,323 with the
[`circuits-v6`](https://github.com/vocdoni/davinci-dkg/releases/tag/circuits-v6) artifacts, from
a dedicated deployer key, [`0x5eb38f2D16CEaabFF6a63C17505a14b45624f537`](https://optimistic.etherscan.io/address/0x5eb38f2D16CEaabFF6a63C17505a14b45624f537),
used for nothing else (not the operators' keys), from the v0.11.1 sources: the Solidity of the
2026-10-15 deployment, whose four verifiers it equals byte for byte. Sources are verified on
Optimistic Etherscan.

| Contract | Address |
|---|---|
| DKGManager | `0x401381a9711B473baf8757404B5514537A5216d7` |
| DKGAppManager | `0x94bead76b89f33e4168940ac7cefbb5cc75f69f0` |
| DKGRegistry | `0x08e7916c3b91ffeb13214dde2542643f8df96af0` |
| ContributionVerifier | `0xe2eb033c756b334c60bf1680a11cbd521f0aa51a` |
| FinalizeVerifier | `0x6a0975823def5b9a70c490967aee9bb5708969d2` |
| PartialDecryptVerifier | `0x99c9816f6c11dcd631a8d2229e7886d3f4829068` |
| DecryptCombineVerifier | `0x460dbd0c3863617b74b0367e2258ec2d5202b9a7` |

Deployment block: 158,592,323 (2026-10-23).

| Transaction | Hash | Gas |
|---|---|---:|
| ContributionVerifier | `0xfc687aba047c274204cdb771670ecac6ed491605ec7ba6802553b828604577eb` | 1,286,050 |
| FinalizeVerifier | `0xd5c08c55425ef32c5f9216e5e89d79bb0fb0f35c09061d844d636d2f4feab0d1` | 1,242,131 |
| PartialDecryptVerifier | `0xb5e5d283b7babed079a2d8935c30f8744e08960a196475ec1d4b527dde39fb78` | 1,580,373 |
| DecryptCombineVerifier | `0xf9125833555cbb650b6fc3ed448bf3ac1de8d76f6efbd12aa6ab0be9b04781ed` | 1,322,348 |
| DKGRegistry | `0x28639ac9f91c34d2f99ff896a6b0ec6ca48e0c9da9170cdd804086a580aadafe` | 1,230,014 |
| DKGManager | `0xc61e4f448aee47c746099a491288956191715beb99320437d9d2b96ad7f0f337` | 4,070,027 |
| DKGRegistry.setManager | `0x7bc3c0d983a625185d1ace83b9cb9b96f9c2a5208f48d7fc5a46fa2ed4bd92e8` | 45,198 |
| DKGAppManager | `0x9cf47270530564e483fd0535f96995775340aee3b2e90d4511d6c65bd604a2d9` | 2,035,232 |
| DKGManager.setAppManager | `0x1046139374d1c75386aca5f55d0975ce6fbb38128a38ed2e57e9d419e9dfb5ac` | 44,191 |

12,855,564 gas (0.000000004 ETH) plus 0.00000108 ETH of OP L1 data fee.
`setManager` and `setAppManager` are one-shot, so the suite has no owner.

Nodes on chain 10 tip 0.001 gwei when the RPC cannot suggest a tip, cap `maxFeePerGas` (and the
tip) at 0.02 gwei, bumps included (5x the worst OP base fee in 4,145 samples, 0.0039 gwei), and
space their finalize and decryption rotations 8 blocks apart (16 s; 3 blocks, 15 s, on Gnosis).
Each operator holds at least 0.1 mETH and is topped up monthly: the pool reserves gas limit ×
cap over the transactions in flight (a finalize, 2.6M gas, plus a contribution, 0.55M: 0.063
mETH at 0.02 gwei), and running costs about 0.0015 mETH a day.

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

The preset's public endpoints (`https://mainnet.optimism.io`, `https://optimism.gateway.tenderly.co`)
are rate-limited. Both serve old log ranges and receipts (`mainnet.optimism.io` served 10,000-block
log ranges 136,000 blocks back on 2026-10-15), which a member rebuilding its key share and every
transaction wait need. `https://optimism-rpc.publicnode.com` is not in the preset: it refuses state
older than about 600 blocks, log ranges older than about 10,000 blocks and every receipt (HTTP 403).
A long-lived node lists a private provider first in `--web3.rpc` and the public ones after it as
failover; the node always prefers the earliest endpoint that works and moves a refused,
rate-limited or failed read to the next one within the call.

**Retired:** DKGManager `0x9E799A1D22710056E62c849e3FFE19Dd01db2D41` (block 158,245,366, 2026-10-15,
registry `0x5ac7ec4c9a2100f705d108841a170074fe95a998`, app manager `0x1e7f0bec007828e2e0a1a469d9d3e148888fc39a`).
Its applications and plaintexts stay on chain, but no node serves it since v0.11.2.

## Gnosis Chain (retired)

Chain id 100. The former default network of the node, no longer built in: run it with
`--manager` and `--web3.rpc` (the explorer serves OP since 2026-10-15). Deployed at block 48,632,905 with the
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
   and commit. The explorer's config (`ui/public/config.json`, the fallbacks in
   `scripts/render-ui-config.sh` and its test, `ui/.do/davinci-dkg-ui.yaml`) moves in the same
   release with `MANAGER_ADDRESS` and `DEPLOY_BLOCK`; its `RPC_URL` reaches every visitor's browser,
   so it stays a public endpoint (`https://mainnet.optimism.io` serves its log scans).
4. **Release** a stable `vX.Y.Z`. `latest` moves, so compose nodes under Watchtower switch to the
   new manager on their own, register in its registry at start (`EnsureRegistered`) and start
   from an empty `<datadir>/10-<manager>`. Railway has no Watchtower: roll every service with
   `scripts/railway-roll-node.sh <service id> ghcr.io/vocdoni/davinci-dkg:vX.Y.Z`; a service that
   changes network as well (as the Gnosis fleet did for OP) moves with
   `scripts/railway-repoint-node.sh <service name> --plan|--apply`, which also sets
   `DAVINCI_DKG_NETWORK` and the RPC list (a private endpoint first), drops variables that would
   override the preset and the fixed epoch policy (the adaptive one sizes committees from the
   registry), keeps the volume and the key, and saves what it replaced for `--rollback`
   (self-test: `scripts/railway-repoint-node.test.sh`). For the
   explorer on DigitalOcean App Platform: the app's build-time env (`MANAGER_ADDRESS`,
   `DEPLOY_BLOCK`) overrides the repo file, so apply the values from `ui/.do/davinci-dkg-ui.yaml`
   to the app's environment spec, trigger a redeploy, and purge the Cloudflare cache.
5. **Verify.** Every node logs `self: registry row` and `node running` against the new manager
   (`scripts/railway-node-status.sh <service id>`). `cast call $REGISTRY "activeCount()(uint64)"`
   reaches the fleet size. The first epoch goes `Live` about four minutes after a node creates it (the windows are 3 min + 1 min).
   It needs `MIN_COMMITTEE_SIZE` (3) active operators, all claiming within the 90-block selection
   window, and at least `minValidContributions` contributions. Then register an automatic
   application and decrypt a value end to end with `dkgapp`.
