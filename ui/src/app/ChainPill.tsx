import { useRuntimeConfig } from '~config/config-context'
import { useLatestBlock } from '~hooks/use-latest-block'
import { Address, Badge, Tooltip } from '~kit'
import { cn } from '~lib/cn'

/**
 * Chain identity in the top bar: network name, live head, manager address.
 * The pulsing dot is the only "is it alive?" signal the explorer needs.
 */
export function ChainPill({ className }: { className?: string }) {
  const config = useRuntimeConfig()
  const { blockNumber, isError } = useLatestBlock()

  return (
    <div className={cn('flex items-center gap-2.5 rounded-pill border border-charcoal bg-carbon px-3 py-1', className)}>
      <Tooltip content={`chain id ${config.chainId}`}>
        <span className='flex items-center gap-1.5 text-[11px] font-medium whitespace-nowrap text-silver'>
          <span
            className={cn(
              'h-1.5 w-1.5 rounded-full',
              isError ? 'bg-red' : blockNumber == null ? 'bg-amber' : 'animate-skeleton bg-emerald'
            )}
          />
          {config.chainName}
        </span>
      </Tooltip>
      <span aria-hidden='true' className='h-3 w-px bg-charcoal' />
      <Tooltip content={isError ? 'The RPC endpoint is not answering' : 'Latest block'}>
        <span className='font-mono text-[11px] tnum text-pewter'>
          {blockNumber == null ? '—' : `#${blockNumber.toString()}`}
        </span>
      </Tooltip>
      <span aria-hidden='true' className='hidden h-3 w-px bg-charcoal lg:block' />
      <span className='hidden lg:block'>
        <Address value={config.managerAddress} chars={4} copy={false} />
      </span>
      {config.demo ? (
        <Badge tone='warn' size='sm' className='ml-0.5'>
          demo
        </Badge>
      ) : null}
    </div>
  )
}
