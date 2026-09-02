import { useCopy } from '~hooks/use-copy'
import { cn } from '~lib/cn'
import { CheckIcon, CopyIcon } from './icons'
import { Tooltip } from './Tooltip'

export interface CopyButtonProps {
  value: string
  /** Accessible name; defaults to "Copy". */
  label?: string
  size?: number
  className?: string
}

/** Icon-only copy affordance. Flashes a check for 1.5 s. */
export function CopyButton({ value, label = 'Copy', size = 13, className }: CopyButtonProps) {
  const { copied, copy } = useCopy()
  return (
    <Tooltip content={copied ? 'Copied' : label}>
      <button
        type='button'
        aria-label={label}
        onClick={(e) => {
          e.stopPropagation()
          e.preventDefault()
          copy(value)
        }}
        className={cn(
          'inline-flex shrink-0 items-center justify-center rounded-sm p-1 transition-colors',
          copied ? 'text-emerald' : 'text-ash hover:bg-onyx hover:text-ghost',
          className
        )}
      >
        {copied ? <CheckIcon size={size} /> : <CopyIcon size={size} />}
      </button>
    </Tooltip>
  )
}
