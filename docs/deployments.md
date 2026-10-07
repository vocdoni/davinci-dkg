# Deployments

The public deployments below are built into the node (`--network`), `dkgapp` and the SDK
(`KNOWN_NETWORKS`). Only the `DKGManager` address is configured anywhere; the other contracts are
read from it on chain.

## Gnosis Chain

Chain id 100. The default network of the node. Deployed at block 48,632,905 with the
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

## Moving the Gnosis deployment

A change to the contracts' rules (such as the application-id namespace of
[#14](https://github.com/vocdoni/davinci-dkg/issues/14)) needs a fresh deployment. The old one keeps
its applications and plaintexts on chain, but its nodes leave it once the preset moves.

1. **Deploy** from a clean checkout of the release commit, with Foundry on `PATH`. Do not run
   `make circuits` first: the committed verifiers must stay those of the pinned `circuits-v6`
   artifacts. Write the settings to a `0600` `solidity/.env` and delete it afterwards. These are the
   current Gnosis parameters:

   ```bash
   RPC_URL=https://rpc.gnosischain.com
   CHAIN_ID=100
   PRIVATE_KEY=<deployer key>
   MIN_THRESHOLD=2
   MIN_COMMITTEE_SIZE=3
   MAX_LOTTERY_ALPHA_BPS=20000
   EPOCH_DURATION_BLOCKS=17280
   COMMITTEE_SELECTION_BLOCKS=36
   KEY_ASSEMBLY_BLOCKS=12
   FINALIZE_GAP_BLOCKS=1
   INACTIVITY_WINDOW=50400
   ETHERSCAN_API_KEY=<Etherscan v2 key, for Gnosisscan verification>
   ```

   ```bash
   make solidity-deploy && rm solidity/.env
   ```

   `rpc.gnosischain.com` is the endpoint that accepts the deployment; publicnode and drpc cap
   `eth_getLogs` ranges, which the explorer needs later, not the deployment.
2. **Check it.** `solidity/.last_deployed_addresses.env` lists the seven addresses. Read the
   manager's deployment block and the parameters back:

   ```bash
   . solidity/.last_deployed_addresses.env
   jq -r --arg m "${MANAGER,,}" '.receipts[] | select((.contractAddress // "" | ascii_downcase) == $m)
     | .blockNumber' solidity/broadcast/DeployAll.s.sol/100/run-latest.json | xargs cast to-dec
   for f in "REGISTRY()(address)" "appManager()(address)" "EPOCH_DURATION_BLOCKS()(uint256)" \
     "MIN_THRESHOLD()(uint16)" "MIN_COMMITTEE_SIZE()(uint16)" "MAX_LOTTERY_ALPHA_BPS()(uint16)"; do
     cast call "$MANAGER" "$f" --rpc-url https://rpc.gnosischain.com; done
   cast call "$REGISTRY" "INACTIVITY_WINDOW()(uint64)" --rpc-url https://rpc.gnosischain.com
   ```

3. **Move the preset.** `scripts/move-gnosis-preset.sh solidity/.last_deployed_addresses.env
   <deploy block>` rewrites every place that names the Gnosis deployment: `config/networks.go`,
   `sdk/src/networks.ts`, `ui/public/config.json`, the fallbacks in `scripts/render-ui-config.sh`,
   `ui/.do/davinci-dkg-ui.yaml`, this file, the README and their tests. Review the diff, run the
   checks it prints, list the old manager under a "Retired" note here, and commit.
4. **Release** a stable `vX.Y.Z`. `latest` moves, so compose nodes under Watchtower switch to the
   new manager on their own, register in its registry at start (`EnsureRegistered`) and start
   from an empty `<datadir>/100-<manager>`. Railway has no Watchtower: roll every service with
   `scripts/railway-roll-node.sh <service id> ghcr.io/vocdoni/davinci-dkg:vX.Y.Z`. The explorer
   redeploys from `main` on DigitalOcean App Platform.
5. **Verify.** Every node logs `self: registry row` and `node running` against the new manager
   (`scripts/railway-node-status.sh <service id>`). `cast call $REGISTRY "activeCount()(uint64)"`
   reaches the fleet size. The first epoch goes `Live` about four minutes after a node creates it.
   It needs `MIN_COMMITTEE_SIZE` (3) active operators, all claiming within the 36-block selection
   window, and at least `minValidContributions` contributions. Then register an automatic
   application and decrypt a value end to end with `dkgapp`.
