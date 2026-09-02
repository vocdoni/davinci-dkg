import type { ReactNode } from 'react'
import { Link } from 'react-router-dom'
import { CopyButton } from '~kit'
import { cn } from '~lib/cn'

// The kit has no long-form typography — it is built for tables, panels and
// charts — so the docs bring their own: one heading, one paragraph, one list,
// one code block. Everything below draws only on the theme tokens, so a change
// to the palette still reaches these pages.

/** A section with an anchor the "on this page" rail links to. */
export function Section({ id, title, children }: { id: string; title: string; children: ReactNode }) {
  return (
    <section id={id} className='scroll-mt-24'>
      <h2 className='mb-4 text-lg font-semibold tracking-tight text-ghost'>
        <a href={`#${id}`} className='no-underline hover:text-emerald'>
          {title}
        </a>
      </h2>
      <div className='flex flex-col gap-4'>{children}</div>
    </section>
  )
}

export function Sub({ children }: { children: ReactNode }) {
  return <h3 className='mt-2 text-[15px] font-semibold text-silver'>{children}</h3>
}

export function P({ children, className }: { children: ReactNode; className?: string }) {
  return <p className={cn('max-w-[70ch] text-[14px] leading-relaxed text-ash', className)}>{children}</p>
}

/** Inline literal: a flag, a constant, a function name. */
export function C({ children }: { children: ReactNode }) {
  return (
    <code className='rounded-sm border border-charcoal bg-onyx px-1 py-px font-mono text-[0.86em] text-emerald'>
      {children}
    </code>
  )
}

export function Em({ children }: { children: ReactNode }) {
  return <strong className='font-medium text-silver'>{children}</strong>
}

export function Bullets({ items }: { items: ReactNode[] }) {
  return (
    <ul className='m-0 flex max-w-[70ch] list-disc flex-col gap-2 pl-5 text-[14px] leading-relaxed text-ash marker:text-charcoal'>
      {items.map((item, i) => (
        <li key={i}>{item}</li>
      ))}
    </ul>
  )
}

export function Steps({ items }: { items: ReactNode[] }) {
  return (
    <ol className='m-0 flex max-w-[70ch] list-decimal flex-col gap-2 pl-5 text-[14px] leading-relaxed text-ash marker:font-mono marker:text-pewter'>
      {items.map((item, i) => (
        <li key={i}>{item}</li>
      ))}
    </ol>
  )
}

/** Fenced code on a carbon surface, with a caption bar and a copy button. */
export function Code({ children, caption }: { children: string; caption?: string }) {
  return (
    <div className='overflow-hidden rounded-md border border-charcoal bg-carbon'>
      <div className='flex items-center justify-between gap-3 border-b border-charcoal px-3 py-1.5'>
        <span className='label-caps text-[10px] text-pewter'>{caption ?? 'shell'}</span>
        <CopyButton value={children} label='Copy snippet' />
      </div>
      <pre className='m-0 overflow-x-auto p-4 text-[12.5px] leading-relaxed text-silver scroll-slim'>
        <code>{children}</code>
      </pre>
    </div>
  )
}

/** A link to another page of the explorer. */
export function Internal({ to, children }: { to: string; children: ReactNode }) {
  return (
    <Link
      to={to}
      className='border-b border-emerald/40 text-emerald no-underline transition-colors hover:border-emerald'
    >
      {children}
    </Link>
  )
}

export function Ext({ href, children }: { href: string; children: ReactNode }) {
  return (
    <a
      href={href}
      target='_blank'
      rel='noreferrer noopener'
      className='border-b border-emerald/40 text-emerald no-underline transition-colors hover:border-emerald'
    >
      {children}
    </a>
  )
}

/** A short aside — a caveat that must not be missed but is not a whole section. */
export function Note({ tone = 'info', children }: { tone?: 'info' | 'warn'; children: ReactNode }) {
  return (
    <div
      className={cn(
        'max-w-[70ch] rounded-md border-l-2 py-2 pl-4 text-[13px] leading-relaxed',
        tone === 'warn' ? 'border-l-amber bg-amber/[0.03] text-silver' : 'border-l-charcoal text-ash'
      )}
    >
      {children}
    </div>
  )
}
