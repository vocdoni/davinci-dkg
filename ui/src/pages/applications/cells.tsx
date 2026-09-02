// Cells the applications table renders on its own.

import type { ApplicationRow } from '~indexer/selectors'

/** `notBefore → notAfter`, with 0 meaning "unbounded" on chain. */
export function WindowCell({ row }: { row: ApplicationRow }) {
  const from = row.notBeforeBlock ?? 0
  const to = row.notAfterBlock ?? 0
  if (from === 0 && to === 0) return <span className='text-ash'>open</span>
  return (
    <span className='font-mono text-[12px] tnum text-silver'>
      {from === 0 ? 'any' : from} <span className='text-ash'>→</span> {to === 0 ? 'any' : to}
    </span>
  )
}
