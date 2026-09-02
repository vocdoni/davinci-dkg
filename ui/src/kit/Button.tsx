import { forwardRef, type ButtonHTMLAttributes, type ReactNode } from 'react'
import { buttonClasses, type ButtonSize, type ButtonVariant } from './button-styles'
import { cn } from '~lib/cn'

export interface ButtonProps extends Omit<ButtonHTMLAttributes<HTMLButtonElement>, 'className'> {
  variant?: ButtonVariant
  size?: ButtonSize
  /** Swaps the leading icon for a spinner and disables the button. */
  loading?: boolean
  iconLeft?: ReactNode
  iconRight?: ReactNode
  className?: string
}

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(function Button(
  { variant = 'secondary', size = 'md', loading = false, iconLeft, iconRight, className, children, ...rest },
  ref
) {
  return (
    <button
      ref={ref}
      type={rest.type ?? 'button'}
      {...rest}
      disabled={rest.disabled || loading}
      className={buttonClasses(variant, size, className)}
    >
      {loading ? (
        <span className='h-3.5 w-3.5 animate-spin rounded-full border-[1.5px] border-current border-t-transparent' />
      ) : (
        iconLeft
      )}
      {children}
      {iconRight}
    </button>
  )
})

export interface ButtonLinkProps {
  href: string
  variant?: ButtonVariant
  size?: ButtonSize
  external?: boolean
  className?: string
  children?: ReactNode
}

/** An `<a>` wearing the button skin. For explorer links and downloads. */
export function ButtonLink({
  href,
  variant = 'secondary',
  size = 'md',
  external,
  className,
  children,
}: ButtonLinkProps) {
  return (
    <a
      href={href}
      target={external ? '_blank' : undefined}
      rel={external ? 'noreferrer noopener' : undefined}
      className={cn(buttonClasses(variant, size, className))}
    >
      {children}
    </a>
  )
}
