import clsx, { type ClassValue } from 'clsx'

/**
 * Class-name join. Deliberately just `clsx` — no tailwind-merge: kit
 * components put the caller's `className` last so a caller's utility wins by
 * source order, which is enough for the handful of overrides we do.
 */
export function cn(...inputs: ClassValue[]): string {
  return clsx(inputs)
}
