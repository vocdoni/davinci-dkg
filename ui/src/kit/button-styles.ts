import { cn } from '~lib/cn'

export type ButtonVariant = 'primary' | 'ghost' | 'secondary' | 'subtle' | 'danger'
export type ButtonSize = 'sm' | 'md' | 'lg' | 'icon'

const VARIANTS: Record<ButtonVariant, string> = {
  // Emerald fill, obsidian text. The single most important action on a screen.
  primary:
    'bg-emerald text-obsidian border border-transparent font-semibold hover:opacity-85 hover:shadow-[0_0_20px_rgba(0,217,146,0.2)]',
  // Emerald text on an emerald hairline: interactive without competing.
  ghost: 'bg-transparent text-emerald border border-emerald/70 hover:bg-emerald/8 hover:border-emerald',
  // Ghost text, charcoal hairline: everything else.
  secondary: 'bg-transparent text-ghost border border-charcoal hover:border-warm-gray hover:bg-onyx',
  // Borderless; toolbars, table row actions, icon buttons.
  subtle: 'bg-transparent text-pewter border border-transparent hover:text-ghost hover:bg-onyx',
  // The one destructive treatment. Used sparingly (spec §5).
  danger: 'bg-transparent text-red border border-red/40 hover:bg-red/10 hover:border-red',
}

const SIZES: Record<ButtonSize, string> = {
  sm: 'h-7 gap-1.5 px-3 text-xs',
  md: 'h-8 gap-1.5 px-4 text-[13px]',
  lg: 'h-10 gap-2 px-6 text-sm',
  icon: 'h-8 w-8 justify-center p-0',
}

/**
 * Shared button surface. Exported separately from the component so `<Link>`s
 * and `<a>`s can wear the same skin without a polymorphic `as` prop.
 */
export function buttonClasses(variant: ButtonVariant = 'secondary', size: ButtonSize = 'md', className?: string) {
  return cn(
    'inline-flex select-none items-center whitespace-nowrap rounded-sm font-medium leading-none no-underline',
    'transition-[background-color,border-color,color,opacity,box-shadow] duration-150 active:scale-[0.97]',
    'disabled:pointer-events-none disabled:opacity-40',
    VARIANTS[variant],
    SIZES[size],
    className
  )
}
