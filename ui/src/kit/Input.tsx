import { forwardRef, useId, type InputHTMLAttributes, type ReactNode, type SelectHTMLAttributes } from 'react'
import { cn } from '~lib/cn'
import { ChevronDownIcon } from './icons'

const FIELD = [
  'w-full rounded-sm border border-charcoal bg-transparent text-[13px] text-ghost',
  'placeholder:text-ash transition-[border-color,box-shadow] duration-150',
  'focus:border-emerald focus:shadow-[0_0_8px_rgba(0,217,146,0.15)] focus:outline-none',
  'disabled:cursor-not-allowed disabled:opacity-50',
].join(' ')

export interface FieldShellProps {
  label?: ReactNode
  hint?: ReactNode
  error?: ReactNode
  className?: string
  children: ReactNode
  htmlFor?: string
}

/** Label + control + hint/error. Shared by Input, Select and Toggle. */
export function Field({ label, hint, error, className, children, htmlFor }: FieldShellProps) {
  return (
    <div className={cn('min-w-0', className)}>
      {label ? (
        <label htmlFor={htmlFor} className='label-caps mb-1.5 block text-[11px] text-pewter'>
          {label}
        </label>
      ) : null}
      {children}
      {error ? <p className='mt-1.5 text-[11px] text-red'>{error}</p> : null}
      {!error && hint ? <p className='mt-1.5 text-[11px] text-ash'>{hint}</p> : null}
    </div>
  )
}

export interface InputProps extends Omit<InputHTMLAttributes<HTMLInputElement>, 'size'> {
  label?: ReactNode
  hint?: ReactNode
  error?: ReactNode
  /** JetBrains Mono + tabular numbers, for addresses, hashes and amounts. */
  mono?: boolean
  size?: 'sm' | 'md'
  iconLeft?: ReactNode
  /** Trailing slot: a unit, a button, a spinner. */
  addonRight?: ReactNode
  wrapperClassName?: string
}

export const Input = forwardRef<HTMLInputElement, InputProps>(function Input(
  { label, hint, error, mono, size = 'md', iconLeft, addonRight, className, wrapperClassName, ...rest },
  ref
) {
  // Generated when the caller gives no id, so a label is always associated
  // with its control even in a loop of unnamed fields.
  const fallbackId = useId()
  const id = rest.id ?? fallbackId
  return (
    <Field label={label} hint={hint} error={error} className={wrapperClassName} htmlFor={id}>
      <div className='relative flex items-center'>
        {iconLeft ? <span className='pointer-events-none absolute left-3 text-ash'>{iconLeft}</span> : null}
        <input
          ref={ref}
          {...rest}
          id={id}
          className={cn(
            FIELD,
            size === 'sm' ? 'h-8 px-3' : 'h-9 px-3.5',
            iconLeft && 'pl-9',
            addonRight && 'pr-9',
            mono && 'font-mono tnum',
            error && 'border-red/60',
            className
          )}
        />
        {addonRight ? <span className='absolute right-2 flex items-center text-ash'>{addonRight}</span> : null}
      </div>
    </Field>
  )
})

export interface SelectOption {
  value: string
  label: string
  disabled?: boolean
}

export interface SelectProps extends Omit<SelectHTMLAttributes<HTMLSelectElement>, 'size'> {
  label?: ReactNode
  hint?: ReactNode
  error?: ReactNode
  options: SelectOption[]
  size?: 'sm' | 'md'
  wrapperClassName?: string
}

/** Native select — keyboard and mobile behaviour for free, styled to match. */
export const Select = forwardRef<HTMLSelectElement, SelectProps>(function Select(
  { label, hint, error, options, size = 'md', className, wrapperClassName, ...rest },
  ref
) {
  const fallbackId = useId()
  const id = rest.id ?? fallbackId
  return (
    <Field label={label} hint={hint} error={error} className={wrapperClassName} htmlFor={id}>
      <div className='relative flex items-center'>
        <select
          ref={ref}
          {...rest}
          id={id}
          className={cn(
            FIELD,
            'cursor-pointer appearance-none pr-8',
            size === 'sm' ? 'h-8 px-3' : 'h-9 px-3.5',
            error && 'border-red/60',
            className
          )}
        >
          {options.map((opt) => (
            <option key={opt.value} value={opt.value} disabled={opt.disabled} className='bg-carbon text-ghost'>
              {opt.label}
            </option>
          ))}
        </select>
        <ChevronDownIcon className='pointer-events-none absolute right-2.5 text-ash' />
      </div>
    </Field>
  )
})

export interface ToggleProps {
  checked: boolean
  onChange: (checked: boolean) => void
  label?: ReactNode
  hint?: ReactNode
  disabled?: boolean
  className?: string
}

/** Switch. Emerald when on, onyx track when off. */
export function Toggle({ checked, onChange, label, hint, disabled, className }: ToggleProps) {
  return (
    <label
      className={cn('flex cursor-pointer items-center gap-3', disabled && 'cursor-not-allowed opacity-50', className)}
    >
      <button
        type='button'
        role='switch'
        aria-checked={checked}
        disabled={disabled}
        onClick={() => onChange(!checked)}
        className={cn(
          'relative h-4.5 w-8 shrink-0 rounded-pill border transition-colors duration-150',
          checked ? 'border-emerald bg-emerald/25' : 'border-charcoal bg-onyx'
        )}
      >
        <span
          className={cn(
            'absolute top-0.5 h-3 w-3 rounded-full transition-[left] duration-150',
            checked ? 'left-4 bg-emerald' : 'left-0.5 bg-warm-gray'
          )}
        />
      </button>
      {label || hint ? (
        <span className='min-w-0'>
          {label ? <span className='block text-[13px] text-silver'>{label}</span> : null}
          {hint ? <span className='block text-[11px] text-ash'>{hint}</span> : null}
        </span>
      ) : null}
    </label>
  )
}
