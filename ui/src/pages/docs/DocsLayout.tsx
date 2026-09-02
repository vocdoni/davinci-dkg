import type { ReactNode } from 'react'
import { NavLink } from 'react-router-dom'
import { SectionHeader } from '~kit'
import { paths } from '~routes/paths'
import { cn } from '~lib/cn'

export interface DocsSection {
  id: string
  title: string
}

const PAGES = [
  { to: paths.docsProtocol(), label: 'Protocol' },
  { to: paths.docsRunANode(), label: 'Run a node' },
  { to: paths.docsSdk(), label: 'SDK' },
]

/**
 * The three docs pages share one frame: the page switcher on top, the prose in
 * a measured column, and an "on this page" rail on the right that is built
 * from the same `sections` array the page renders its anchors from — so an
 * outline entry without a section (or the reverse) is a compile-time mistake,
 * not a dead link.
 */
export function DocsLayout({
  label,
  title,
  description,
  sections,
  children,
}: {
  label: string
  title: string
  description: string
  sections: DocsSection[]
  children: ReactNode
}) {
  return (
    <div className='flex flex-col gap-8'>
      <SectionHeader size='page' label={label} title={title} description={description} />

      <nav className='flex gap-1 border-b border-charcoal'>
        {PAGES.map((page) => (
          <NavLink
            key={page.to}
            to={page.to}
            className={({ isActive }) =>
              cn(
                '-mb-px border-b-2 px-3 py-2 text-[13px] font-medium no-underline transition-colors',
                isActive
                  ? 'border-emerald text-emerald'
                  : 'border-transparent text-pewter hover:text-ghost'
              )
            }
          >
            {page.label}
          </NavLink>
        ))}
      </nav>

      <div className='grid gap-10 lg:grid-cols-[minmax(0,1fr)_200px]'>
        <div className='flex min-w-0 flex-col gap-12'>{children}</div>
        <aside className='hidden lg:block'>
          <div className='sticky top-8'>
            <div className='label-caps mb-3 text-[10px] text-pewter'>On this page</div>
            <ul className='m-0 flex list-none flex-col gap-1.5 border-l border-charcoal p-0'>
              {sections.map((section) => (
                <li key={section.id}>
                  <a
                    href={`#${section.id}`}
                    className='-ml-px block border-l border-transparent pl-3 text-[12px] leading-snug text-ash no-underline transition-colors hover:border-emerald hover:text-emerald'
                  >
                    {section.title}
                  </a>
                </li>
              ))}
            </ul>
          </div>
        </aside>
      </div>
    </div>
  )
}
