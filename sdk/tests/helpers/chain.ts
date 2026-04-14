// Chain manipulation helpers for tests running against Anvil.

import type { PublicClient } from 'viem';

/**
 * Mine `count` blocks instantly via the `anvil_mine` RPC method.
 * Only works on an Anvil devnet.
 */
export async function mineBlocks(client: PublicClient, count: number): Promise<void> {
  await client.request({
    method: 'anvil_mine' as any,
    params: [`0x${count.toString(16)}`] as any,
  });
}

/**
 * Mine enough blocks that the seed block (seedBlock = createdAtBlock + seedDelay)
 * is in the past, so claimSlot can resolve the blockhash.
 *
 * @param client      viem PublicClient
 * @param seedBlock   the round's seedBlock value (bigint)
 */
export async function mineUntilSeedAvailable(
  client: PublicClient,
  seedBlock: bigint,
): Promise<void> {
  const current = await client.getBlockNumber();
  if (current <= seedBlock) {
    await mineBlocks(client, Number(seedBlock - current) + 1);
  }
}
