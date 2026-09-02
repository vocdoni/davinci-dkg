import { Link } from 'react-router-dom'
import { useRuntimeConfig } from '~config/config-context'
import { checksum } from '~lib/address'
import { explorerAddressUrl, explorerBlockUrl, explorerTxUrl } from '~lib/explorer'
import { shortHash } from '~lib/format'
import { cn } from '~lib/cn'
import { CopyButton } from './CopyButton'
import { ExternalIcon } from './icons'
import { Tooltip } from './Tooltip'

export interface AddressProps {
  value: string
  /** Characters kept either side of the ellipsis. */
  chars?: number
  /** Render the full checksummed address instead of truncating. */
  full?: boolean
  copy?: boolean
  /** Show the block-explorer link icon. */
  explorer?: boolean
  /** Internal route for the address text (usually `/operators/0x…`). */
  to?: string
  className?: string
}

/**
 * Checksummed address, truncated in the middle, with copy and explorer
 * affordances. The full value is always in the tooltip and the `title`, so a
 * truncated address is never a dead end.
 */
export function Address({ value, chars = 4, full = false, copy = true, explorer = true, to, className }: AddressProps) {
  const config = useRuntimeConfig()
  const checksummed = checksum(value)
  const text = full ? checksummed : shortHash(checksummed, chars, chars)
  const href = explorer ? explorerAddressUrl(config.explorerUrl, checksummed) : null
  return (
    <span className={cn('inline-flex min-w-0 items-center gap-1 font-mono text-[12px]', className)}>
      <Tooltip content={checksummed}>
        {to ? (
          <Link to={to} title={checksummed} onClick={(e) => e.stopPropagation()} className='truncate text-silver transition-colors hover:text-emerald'>
            {text}
          </Link>
        ) : (
          <span title={checksummed} className='truncate text-silver'>
            {text}
          </span>
        )}
      </Tooltip>
      {copy ? <CopyButton value={checksummed} label='Copy address' /> : null}
      {href ? <ExplorerLink href={href} label='View address on the block explorer' /> : null}
    </span>
  )
}

export interface HashProps {
  value: string
  chars?: number
  full?: boolean
  copy?: boolean
  /** Optional href for the hash text itself. */
  href?: string
  className?: string
}

/** Any 0x value that isn't an address: epoch ids, aids, digests, points. */
export function Hash({ value, chars = 6, full = false, copy = true, href, className }: HashProps) {
  const text = full ? value : shortHash(value, chars, 4)
  return (
    <span className={cn('inline-flex min-w-0 items-center gap-1 font-mono text-[12px]', className)}>
      <Tooltip content={value}>
        {href ? (
          <Link to={href} title={value} onClick={(e) => e.stopPropagation()} className='truncate text-silver transition-colors hover:text-emerald'>
            {text}
          </Link>
        ) : (
          <span title={value} className={cn('text-silver', full ? 'break-all' : 'truncate')}>
            {text}
          </span>
        )}
      </Tooltip>
      {copy ? <CopyButton value={value} label='Copy value' /> : null}
    </span>
  )
}

export interface TxCellProps {
  hash: string
  chars?: number
  copy?: boolean
  className?: string
}

/** Transaction hash → block explorer. */
export function TxCell({ hash, chars = 6, copy = false, className }: TxCellProps) {
  const config = useRuntimeConfig()
  const href = explorerTxUrl(config.explorerUrl, hash)
  return (
    <span className={cn('inline-flex min-w-0 items-center gap-1 font-mono text-[12px]', className)}>
      <Tooltip content={hash}>
        {href ? (
          <a
            href={href}
            target='_blank'
            rel='noreferrer noopener'
            className='truncate text-silver transition-colors hover:text-emerald'
          >
            {shortHash(hash, chars, 4)}
          </a>
        ) : (
          <span className='truncate text-silver'>{shortHash(hash, chars, 4)}</span>
        )}
      </Tooltip>
      {copy ? <CopyButton value={hash} label='Copy transaction hash' /> : null}
    </span>
  )
}

export interface BlockCellProps {
  block: bigint | number | null | undefined
  /** Appended after the number, e.g. "· 3 min ago". */
  suffix?: string
  className?: string
}

/** Block number → block explorer, tabular so columns of them line up. */
export function BlockCell({ block, suffix, className }: BlockCellProps) {
  const config = useRuntimeConfig()
  if (block == null) return <span className={cn('font-mono text-[12px] text-ash', className)}>—</span>
  const text = block.toString()
  const href = explorerBlockUrl(config.explorerUrl, block)
  return (
    <span className={cn('inline-flex items-center gap-1 font-mono text-[12px] tnum', className)}>
      {href ? (
        <a
          href={href}
          target='_blank'
          rel='noreferrer noopener'
          className='text-silver transition-colors hover:text-emerald'
        >
          {text}
        </a>
      ) : (
        <span className='text-silver'>{text}</span>
      )}
      {suffix ? <span className='text-ash'>{suffix}</span> : null}
    </span>
  )
}

function ExplorerLink({ href, label }: { href: string; label: string }) {
  return (
    <Tooltip content={label}>
      <a
        href={href}
        target='_blank'
        rel='noreferrer noopener'
        aria-label={label}
        onClick={(e) => e.stopPropagation()}
        className='inline-flex shrink-0 items-center rounded-sm p-1 text-ash transition-colors hover:bg-onyx hover:text-ghost'
      >
        <ExternalIcon size={13} />
      </a>
    </Tooltip>
  )
}
