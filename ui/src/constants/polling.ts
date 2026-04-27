// Centralized polling cadences, in milliseconds.
//
// Keeping these in one place makes it easy to tune for different deployments
// (a public Sepolia RPC tolerates much less aggressive polling than a local
// Anvil) and makes the cost model visible.

export const Polling = {
  /** Default React Query refetchInterval for read-only views. */
  default: 4000,
  /** Faster cadence for the playground while a round is in flight. */
  playgroundRound: 2000,
  /** Cadence for watching a ciphertext's combined-decryption record. */
  decryption: 3000,
  /** Block number poll cadence for the header. */
  blockNumber: 4000,
} as const
