import { useQuery } from '@tanstack/react-query';
import { getContract, type Address } from 'viem';
import { dkgManagerAbi, dkgRegistryAbi } from './abi';
import { getClient, loadConfig, type RuntimeConfig } from './client';

export function useConfig() {
  return useQuery({
    queryKey: ['config'],
    queryFn: loadConfig,
    staleTime: Infinity,
    refetchInterval: false,
  });
}

async function manager(cfg: RuntimeConfig) {
  const client = await getClient();
  return getContract({
    address: cfg.managerAddress,
    abi: dkgManagerAbi,
    client,
  });
}

async function registry(cfg: RuntimeConfig) {
  const client = await getClient();
  return getContract({
    address: cfg.registryAddress,
    abi: dkgRegistryAbi,
    client,
  });
}

export function useChainTip() {
  return useQuery({
    queryKey: ['chain-tip'],
    queryFn: async () => {
      const client = await getClient();
      const block = await client.getBlock();
      return { number: block.number, timestamp: block.timestamp };
    },
  });
}

export function useRoundNonce() {
  return useQuery({
    queryKey: ['round-nonce'],
    queryFn: async () => {
      const cfg = await loadConfig();
      const m = await manager(cfg);
      const nonce = (await m.read.roundNonce()) as bigint;
      return nonce;
    },
  });
}

function buildRoundId(prefix: number, nonce: bigint): `0x${string}` {
  // bytes12 = uint32 prefix || uint64 nonce
  const prefHex = prefix.toString(16).padStart(8, '0');
  const nonceHex = nonce.toString(16).padStart(16, '0');
  return `0x${prefHex}${nonceHex}` as `0x${string}`;
}

export function useRoundPrefix() {
  return useQuery({
    queryKey: ['round-prefix'],
    queryFn: async () => {
      const cfg = await loadConfig();
      const m = await manager(cfg);
      const pref = (await m.read.ROUND_PREFIX()) as number;
      return pref;
    },
    staleTime: Infinity,
  });
}

export function useRound(roundId: `0x${string}` | undefined) {
  return useQuery({
    queryKey: ['round', roundId],
    enabled: !!roundId,
    queryFn: async () => {
      const cfg = await loadConfig();
      const m = await manager(cfg);
      const data = (await m.read.getRound([roundId!])) as any;
      const participants = (await m.read.selectedParticipants([roundId!])) as Address[];
      return { ...data, participants };
    },
  });
}

const RING_BUFFER_SIZE = 64n;

export function useRecentRounds(limit = 20) {
  const nonceQ = useRoundNonce();
  const prefQ = useRoundPrefix();
  return useQuery({
    queryKey: ['recent-rounds', nonceQ.data?.toString(), prefQ.data, limit],
    enabled: nonceQ.data !== undefined && prefQ.data !== undefined,
    queryFn: async () => {
      const cfg = await loadConfig();
      const m = await manager(cfg);
      const nonce = nonceQ.data!;
      const pref = prefQ.data!;
      const ids: `0x${string}`[] = [];
      // roundNonce is post-increment: the most recent round has nonce == roundNonce,
      // and the first round ever created has nonce 1 (nonce 0 is unused). The ring
      // buffer keeps the latest ROUND_HISTORY_SIZE entries.
      if (nonce === 0n) return [] as Array<{ id: `0x${string}`; round: any }>;
      const start = nonce;
      const minNonce = start > RING_BUFFER_SIZE ? start - RING_BUFFER_SIZE + 1n : 1n;
      for (let i = start; i >= minNonce && ids.length < limit; i--) {
        ids.push(buildRoundId(pref, i));
        if (i === 1n) break;
      }
      const rounds = await Promise.all(
        ids.map(async (id) => {
          try {
            const round = (await m.read.getRound([id])) as any;
            return { id, round };
          } catch {
            return { id, round: null };
          }
        }),
      );
      return rounds.filter((r) => r.round && Number(r.round.status) !== 0);
    },
  });
}

export function useRegistry() {
  return useQuery({
    queryKey: ['registry-count'],
    queryFn: async () => {
      const cfg = await loadConfig();
      const r = await registry(cfg);
      const count = (await r.read.nodeCount()) as bigint;
      return count;
    },
  });
}

export function useActiveNodeCount() {
  return useQuery({
    queryKey: ['registry-active-count'],
    queryFn: async () => {
      const cfg = await loadConfig();
      const r = await registry(cfg);
      const count = (await r.read.activeCount()) as bigint;
      return count;
    },
  });
}

export function useInactivityWindow() {
  return useQuery({
    queryKey: ['registry-inactivity-window'],
    queryFn: async () => {
      const cfg = await loadConfig();
      const r = await registry(cfg);
      const window = (await r.read.INACTIVITY_WINDOW()) as bigint;
      return window;
    },
    staleTime: Infinity,
    refetchInterval: false,
  });
}

export function useRoundEvents(roundId: `0x${string}` | undefined) {
  return useQuery({
    queryKey: ['round-events', roundId],
    enabled: !!roundId,
    queryFn: async () => {
      const cfg = await loadConfig();
      const client = await getClient();
      const fromBlock = cfg.startBlock !== undefined ? BigInt(cfg.startBlock) : 0n;
      const events = await client.getContractEvents({
        address: cfg.managerAddress,
        abi: dkgManagerAbi,
        fromBlock,
        toBlock: 'latest',
      });
      return events.filter(
        (ev) => 'args' in ev && (ev.args as { roundId?: string }).roundId === roundId,
      );
    },
  });
}

export function useRegistryNodes() {
  const countQ = useRegistry();
  return useQuery({
    queryKey: ['registry-nodes', countQ.data?.toString()],
    enabled: countQ.data !== undefined,
    queryFn: async () => {
      const cfg = await loadConfig();
      const client = await getClient();
      const fromBlock = cfg.startBlock !== undefined ? BigInt(cfg.startBlock) : 0n;
      const events = await client.getContractEvents({
        address: cfg.registryAddress,
        abi: dkgRegistryAbi,
        eventName: 'NodeRegistered',
        fromBlock,
        toBlock: 'latest',
      });
      const operators = Array.from(
        new Set(
          events
            .map((e) => (e as { args?: { operator?: string } }).args?.operator)
            .filter((op): op is string => typeof op === 'string')
            .map((op) => op.toLowerCase()),
        ),
      );
      const r = await registry(cfg);
      const nodes = await Promise.all(
        operators.map(async (op) => {
          try {
            const node = (await r.read.getNode([op as Address])) as any;
            return node;
          } catch {
            return null;
          }
        }),
      );
      return nodes.filter((n) => n !== null);
    },
  });
}
