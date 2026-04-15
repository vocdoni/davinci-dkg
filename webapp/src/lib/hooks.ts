import { useQuery } from '@tanstack/react-query';
import { buildRoundId } from '@vocdoni/davinci-dkg-sdk';
import { getDKGClient, loadConfig, type RuntimeConfig } from './client';

export function useConfig() {
  return useQuery<RuntimeConfig>({
    queryKey: ['config'],
    queryFn: loadConfig,
    staleTime: Infinity,
    refetchInterval: false,
  });
}

export function useChainTip() {
  return useQuery({
    queryKey: ['chain-tip'],
    queryFn: async () => {
      const client = await getDKGClient();
      const block = await client.publicClient.getBlock();
      return { number: block.number, timestamp: block.timestamp };
    },
  });
}

/**
 * Current block number, refreshed every 12 seconds (≈ one Sepolia block).
 * Returns `undefined` while loading.
 */
export function useBlockNumber() {
  return useQuery<bigint>({
    queryKey: ['block-number'],
    queryFn: async () => {
      const client = await getDKGClient();
      return client.blockNumber();
    },
    refetchInterval: 12_000,
    staleTime: 10_000,
  });
}

export function useRoundNonce() {
  return useQuery({
    queryKey: ['round-nonce'],
    queryFn: async () => {
      const client = await getDKGClient();
      return client.roundNonce();
    },
  });
}

export function useRoundPrefix() {
  return useQuery({
    queryKey: ['round-prefix'],
    queryFn: async () => {
      const client = await getDKGClient();
      return client.roundPrefix();
    },
    staleTime: Infinity,
  });
}

export function useRound(roundId: `0x${string}` | undefined) {
  return useQuery({
    queryKey: ['round', roundId],
    enabled: !!roundId,
    queryFn: async () => {
      const client = await getDKGClient();
      const [round, participants] = await Promise.all([
        client.getRound(roundId!),
        client.selectedParticipants(roundId!),
      ]);
      return { ...round, participants };
    },
  });
}

export function useRecentRounds(limit = 20) {
  return useQuery({
    queryKey: ['recent-rounds', limit],
    queryFn: async () => {
      const client = await getDKGClient();
      return client.getRecentRounds(limit);
    },
  });
}

export function useRegistry() {
  return useQuery({
    queryKey: ['registry-count'],
    queryFn: async () => {
      const client = await getDKGClient();
      return client.nodeCount();
    },
  });
}

export function useActiveNodeCount() {
  return useQuery({
    queryKey: ['registry-active-count'],
    queryFn: async () => {
      const client = await getDKGClient();
      return client.activeCount();
    },
  });
}

export function useInactivityWindow() {
  return useQuery({
    queryKey: ['registry-inactivity-window'],
    queryFn: async () => {
      const client = await getDKGClient();
      return client.inactivityWindow();
    },
    staleTime: Infinity,
    refetchInterval: false,
  });
}

export function useRoundEvents(roundId: `0x${string}` | undefined) {
  const configQ = useConfig();
  return useQuery({
    queryKey: ['round-events', roundId],
    enabled: !!roundId && configQ.data !== undefined,
    queryFn: async () => {
      const client = await getDKGClient();
      const fromBlock = configQ.data?.startBlock !== undefined
        ? BigInt(configQ.data.startBlock)
        : 0n;
      return client.getAllRoundEvents(roundId!, fromBlock);
    },
  });
}

export function useRegistryNodes() {
  const configQ = useConfig();
  return useQuery({
    queryKey: ['registry-nodes'],
    enabled: configQ.data !== undefined,
    queryFn: async () => {
      const client = await getDKGClient();
      const fromBlock = configQ.data?.startBlock !== undefined
        ? BigInt(configQ.data.startBlock)
        : 0n;
      return client.getRegistryNodes(fromBlock);
    },
  });
}

// ── Round ID helper (re-exported for pages that build IDs from nonce) ─────────
export { buildRoundId };
