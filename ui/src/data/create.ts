// One factory for both modes, so the app shell never branches on `demo`
// beyond passing the flag.

import type { PublicClient } from 'viem'
import { createDemoDataSource, type DemoSourceOptions } from '../fixtures/demo'
import { createLiveDataSource, type DataSource } from './source'
import type { Address } from '../indexer/types'
import type { KVStore } from '../indexer/persist'

/** The fields the explorer reads out of `public/config.json`. */
export interface ExplorerConfig {
  chainId: number
  chainName?: string
  managerAddress: Address
  deployBlock: number
  rpcUrl?: string
  explorerUrl?: string
  /** Optional overrides; both are read from the manager when omitted. */
  registryAddress?: Address
  appManagerAddress?: Address
  /** `RuntimeConfig.demo` — accepted so the whole config can be passed in. */
  demo?: boolean
}

export interface CreateDataSourceOptions {
  /** Overrides `config.demo` when given. */
  demo?: boolean
  /** Required for the live source. */
  client?: PublicClient | null
  config: ExplorerConfig
  pollIntervalMs?: number
  chunkSize?: number
  blockTimeSeconds?: number
  staggerBlocks?: number
  /** Persistence backend; `null` disables the IndexedDB cache. */
  kv?: KVStore | null
  demoOptions?: DemoSourceOptions
}

export function createDataSource(options: CreateDataSourceOptions): DataSource {
  const { client, config } = options
  const demo = options.demo ?? config.demo ?? false
  if (demo || !client) {
    return createDemoDataSource({
      chainId: config.chainId,
      chainName: config.chainName,
      managerAddress: config.managerAddress,
      deployBlock: config.deployBlock,
      explorerUrl: config.explorerUrl,
      staggerBlocks: options.staggerBlocks,
      ...options.demoOptions,
    })
  }
  return createLiveDataSource({
    client,
    chainId: config.chainId,
    chainName: config.chainName,
    managerAddress: config.managerAddress,
    deployBlock: config.deployBlock,
    explorerUrl: config.explorerUrl,
    registryAddress: config.registryAddress,
    appManagerAddress: config.appManagerAddress,
    pollIntervalMs: options.pollIntervalMs,
    chunkSize: options.chunkSize,
    blockTimeSeconds: options.blockTimeSeconds,
    staggerBlocks: options.staggerBlocks,
    kv: options.kv,
  })
}
