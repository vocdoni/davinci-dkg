import * as RadixTooltip from '@radix-ui/react-tooltip'
import type { ReactNode } from 'react'
import { cn } from '~lib/cn'

/** Mount once, near the root. Radix requires a provider above every tooltip. */
export function TooltipProvider({ children }: { children: ReactNode }) {
  return (
    <RadixTooltip.Provider delayDuration={120} skipDelayDuration={300}>
      {children}
    </RadixTooltip.Provider>
  )
}

export interface TooltipProps {
  content: ReactNode
  side?: 'top' | 'right' | 'bottom' | 'left'
  align?: 'start' | 'center' | 'end'
  /** Rendered as-is; must accept a ref (Radix asChild). */
  children: ReactNode
  className?: string
}

export function Tooltip({ content, side = 'top', align = 'center', children, className }: TooltipProps) {
  if (content == null || content === '') return <>{children}</>
  return (
    <RadixTooltip.Root>
      <RadixTooltip.Trigger asChild>{children}</RadixTooltip.Trigger>
      <RadixTooltip.Portal>
        <RadixTooltip.Content
          side={side}
          align={align}
          sideOffset={6}
          collisionPadding={8}
          className={cn(
            'z-50 max-w-xs rounded-sm border border-charcoal bg-onyx px-2.5 py-1.5',
            'text-[11px] leading-snug text-silver shadow-[0_8px_24px_rgba(0,0,0,0.55)]',
            className
          )}
        >
          {content}
          <RadixTooltip.Arrow className='fill-onyx' width={8} height={4} />
        </RadixTooltip.Content>
      </RadixTooltip.Portal>
    </RadixTooltip.Root>
  )
}
