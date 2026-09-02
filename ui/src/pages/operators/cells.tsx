// Cells the operators table and the operator page share.

import type { ReactNode } from 'react'
import { Badge, Button, Hash, Popover, Tooltip } from '~kit'
import type { OperatorRow } from '~indexer/selectors'
import type { Point } from '~indexer/types'
import { bigIntToHex, blocksToDuration, shortHash } from '~lib/format'
import { formatCompact } from '~kit/charts'

/**
 * Swallows a click before it reaches a clickable table row.
 *
 * Block and transaction cells are links to the block explorer; without this a
 * click on one both opens the explorer and navigates the row underneath.
 */
export function StopClick({ children }: { children: ReactNode }) {
  return (
    <span className='inline-flex min-w-0 items-center' onClick={(e) => e.stopPropagation()}>
      {children}
    </span>
  )
}

export function OperatorStatusBadge({ row, size = 'sm' }: { row: OperatorRow; size?: 'sm' | 'md' }) {
  if (row.status === 'active') {
    return (
      <Badge
        size={size}
        dot
        tone={row.reapable ? 'warn' : 'ok'}
        title={row.reapable ? 'Active on chain, but past the inactivity window — anyone may reap it' : undefined}
      >
        {row.reapable ? 'idle' : 'active'}
      </Badge>
    )
  }
  if (row.status === 'inactive') {
    return (
      <Badge size={size} tone='danger' title='Reaped for inactivity; must reactivate before it can claim again'>
        inactive
      </Badge>
    )
  }
  return (
    <Badge size={size} tone='neutral'>
      unregistered
    </Badge>
  )
}

/** BabyJubJub encryption key, collapsed to its x prefix until asked for. */
export function OperatorKeyCell({ pubKey }: { pubKey: Point | null }) {
  if (!pubKey) return <span className='text-ash'>—</span>
  const x = bigIntToHex(pubKey.x)
  const y = bigIntToHex(pubKey.y)
  return (
    <StopClick>
      <Popover
        align='end'
        trigger={
          <Button size='sm' variant='subtle' className='px-2 font-mono text-[11px]'>
            {shortHash(x, 3, 3)}
          </Button>
        }
      >
        <div className='flex flex-col gap-2 p-1'>
          <div className='label-caps text-[10px] text-pewter'>BabyJubJub key (TE)</div>
          <div>
            <div className='text-[10px] text-ash'>x</div>
            <Hash value={x} chars={16} />
          </div>
          <div>
            <div className='text-[10px] text-ash'>y</div>
            <Hash value={y} chars={16} />
          </div>
        </div>
      </Popover>
    </StopClick>
  )
}

export function LastActiveCell({ row, blockTimeSeconds }: { row: OperatorRow; blockTimeSeconds: number }) {
  if (row.lastActiveBlock <= 0 || row.idleBlocks == null) return <span className='text-ash'>never</span>
  return (
    <Tooltip content={`block ${row.lastActiveBlock} · ${blocksToDuration(row.idleBlocks, blockTimeSeconds)} ago`}>
      <span className={row.reapable ? 'text-amber' : undefined}>{formatCompact(row.idleBlocks)}</span>
    </Tooltip>
  )
}
