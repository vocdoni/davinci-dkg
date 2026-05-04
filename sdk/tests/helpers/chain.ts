// Chain manipulation helpers for tests running against Anvil.

import type { Address, PublicClient } from 'viem';

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
 * @param seedBlock   the epoch's seedBlock value (bigint)
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

/**
 * Honor the on-chain cadence guard before calling createEpoch: the manager
 * reverts unless `block.number >= nextEpochStartBlock()`. When SDK tests
 * run on the same Anvil after the Go fixture has already minted epochs,
 * `lastEpochStartBlock` is non-zero and the guard fires until enough blocks
 * are mined. This is a test-only concession; production callers wait for
 * cadence naturally.
 */
export async function mineUntilEpochAllowed(
  client: PublicClient,
  managerAddress: Address,
): Promise<void> {
  const nextStart = (await client.readContract({
    address: managerAddress,
    abi: [{
      type: 'function',
      name: 'nextEpochStartBlock',
      stateMutability: 'view',
      inputs: [],
      outputs: [{ type: 'uint64' }],
    }],
    functionName: 'nextEpochStartBlock',
  })) as bigint;
  const current = await client.getBlockNumber();
  if (current < nextStart) {
    await mineBlocks(client, Number(nextStart - current));
  }
}
