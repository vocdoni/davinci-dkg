// Central registry of every React Query key used in the app. Keeping them
// here is the only way to make `queryClient.invalidateQueries({queryKey:
// QueryKeys.foo})` greppable and to avoid the classic stringly-typed bug
// where two callers spell the same key slightly differently and miss each
// other's invalidations.

export const QueryKeys = {
  chain: ['chain'] as const,
  blockNumber: ['chain', 'blockNumber'] as const,

  epochsRecent: (limit: number) => ['epochs', 'recent', limit] as const,
  epoch: (id: `0x${string}`) => ['epochs', id] as const,
  epochEvents: (id: `0x${string}`, fromBlock?: bigint) =>
    ['epochs', id, 'events', fromBlock?.toString() ?? 'all'] as const,

  registryNodes: ['registry', 'nodes'] as const,
  registryStats: ['registry', 'stats'] as const,

  decryption: (id: `0x${string}`, ix: number) => ['epochs', id, 'decryption', ix] as const,
} as const
