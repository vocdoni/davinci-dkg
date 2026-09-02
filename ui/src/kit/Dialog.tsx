import * as RadixDialog from '@radix-ui/react-dialog'
import type { ReactNode } from 'react'
import { cn } from '~lib/cn'
import { CloseIcon } from './icons'

export interface DialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  title: ReactNode
  description?: ReactNode
  footer?: ReactNode
  /** Optional element that opens the dialog (rendered with `asChild`). */
  trigger?: ReactNode
  size?: 'sm' | 'md' | 'lg'
  children?: ReactNode
}

const SIZES = { sm: 'max-w-sm', md: 'max-w-lg', lg: 'max-w-3xl' } as const

/** Modal. Obsidian scrim, carbon panel, charcoal hairline — no drop shadow drama. */
export function Dialog({
  open,
  onOpenChange,
  title,
  description,
  footer,
  trigger,
  size = 'md',
  children,
}: DialogProps) {
  return (
    <RadixDialog.Root open={open} onOpenChange={onOpenChange}>
      {trigger ? <RadixDialog.Trigger asChild>{trigger}</RadixDialog.Trigger> : null}
      <RadixDialog.Portal>
        <RadixDialog.Overlay className='fixed inset-0 z-40 bg-obsidian/80 backdrop-blur-sm' />
        <RadixDialog.Content
          className={cn(
            'fixed top-1/2 left-1/2 z-50 w-[calc(100vw-2rem)] -translate-x-1/2 -translate-y-1/2',
            'rounded-md border border-charcoal bg-carbon',
            SIZES[size]
          )}
        >
          <div className='flex items-start justify-between gap-4 border-b border-charcoal px-5 py-4'>
            <div className='min-w-0'>
              <RadixDialog.Title className='text-[15px] font-semibold text-ghost'>{title}</RadixDialog.Title>
              {description ? (
                <RadixDialog.Description className='mt-1 text-[13px] leading-relaxed text-ash'>
                  {description}
                </RadixDialog.Description>
              ) : null}
            </div>
            <RadixDialog.Close
              aria-label='Close'
              className='shrink-0 rounded-sm p-1 text-pewter transition-colors hover:bg-onyx hover:text-ghost'
            >
              <CloseIcon />
            </RadixDialog.Close>
          </div>
          <div className='max-h-[70vh] overflow-y-auto p-5 scroll-slim'>{children}</div>
          {footer ? <div className='flex justify-end gap-2 border-t border-charcoal px-5 py-3'>{footer}</div> : null}
        </RadixDialog.Content>
      </RadixDialog.Portal>
    </RadixDialog.Root>
  )
}
