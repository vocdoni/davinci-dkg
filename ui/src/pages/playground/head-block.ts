// One defensive read of the chain head.
//
// The synthetic fixture currently publishes `chain.headBlock` (and
// `lastIndexedBlock`) as `NaN`, which silently poisons every block arithmetic
// downstream — a countdown renders "NaN", and a simulator seeded with it puts
// its first transaction in block `NaN`, after which every comparison against a
// real block number is false. Rather than trust the field, take it when it is
// a number and otherwise recover the head from the newest event in the store.

import type { IndexerStore } from '~indexer/types'

export function resolveHeadBlock(store: IndexerStore): number {
  const head = store.chain.headBlock
  if (Number.isFinite(head)) return head
  const indexed = store.lastIndexedBlock
  if (Number.isFinite(indexed)) return indexed
  let newest = Number.isFinite(store.chain.deployBlock) ? store.chain.deployBlock : 0
  for (const event of store.events) {
    if (event.block > newest) newest = event.block
  }
  return newest
}
