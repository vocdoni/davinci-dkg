import type { RuntimeConfig } from '~types/index'

/**
 * Lower bound for every historical `getLogs` scan the explorer runs.
 *
 * Prefers `deployBlock` (the block `DKGManager` was deployed at, set from
 * `DEPLOY_BLOCK` by `scripts/render-ui-config.sh`) and falls back to the older
 * `startBlock` key so an existing `/config.json` keeps working. `0n` means
 * "unknown": the SDK then clamps the scan to a recent-block window instead of
 * walking back to genesis, which would be slow and would still be correct.
 */
export function scanFromBlock(config: RuntimeConfig): bigint {
  const block = config.deployBlock ?? config.startBlock ?? 0
  return block > 0 ? BigInt(block) : 0n
}
