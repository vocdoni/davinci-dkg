export interface RuntimeConfig {
  rpcUrl: string
  managerAddress: `0x${string}`
  registryAddress?: `0x${string}`
  chainId: number
  chainName: string
  startBlock?: number
  /** Block-explorer base URL (no trailing slash), e.g. `https://sepolia.etherscan.io`. */
  explorerUrl?: string
}
