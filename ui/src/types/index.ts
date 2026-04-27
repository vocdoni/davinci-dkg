export interface RuntimeConfig {
  rpcUrl: string
  managerAddress: `0x${string}`
  registryAddress?: `0x${string}`
  chainId: number
  chainName: string
  startBlock?: number
}
