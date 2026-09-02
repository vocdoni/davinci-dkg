export interface RuntimeConfig {
  rpcUrl: string
  managerAddress: `0x${string}`
  registryAddress?: `0x${string}`
  chainId: number
  chainName: string
  startBlock?: number
  /**
   * Block the `DKGManager` was deployed at. Every historical log scan (operator
   * statistics, epoch activity) starts here, so a correct value is the
   * difference between a handful of `eth_getLogs` calls and a walk over the
   * whole chain. `0` (or absent) means "unknown": the SDK then clamps the scan
   * to a recent-block window instead of going back to genesis.
   */
  deployBlock?: number
  /** Block-explorer base URL (no trailing slash), e.g. `https://sepolia.etherscan.io`. */
  explorerUrl?: string
}
