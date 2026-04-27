// Central registry of every React Query key used in the app. Keeping them
// here is the only way to make `queryClient.invalidateQueries({queryKey:
// QueryKeys.foo})` greppable and to avoid the classic stringly-typed bug
// where two callers spell the same key slightly differently and miss each
// other's invalidations.

export const QueryKeys = {
  chain: ['chain'] as const,
  blockNumber: ['chain', 'blockNumber'] as const,

  roundsRecent: (limit: number) => ['rounds', 'recent', limit] as const,
  round: (id: `0x${string}`) => ['rounds', id] as const,
  roundEvents: (id: `0x${string}`, fromBlock?: bigint) =>
    ['rounds', id, 'events', fromBlock?.toString() ?? 'all'] as const,

  registryNodes: ['registry', 'nodes'] as const,
  registryStats: ['registry', 'stats'] as const,

  decryption: (id: `0x${string}`, ix: number) => ['rounds', id, 'decryption', ix] as const,
} as const
