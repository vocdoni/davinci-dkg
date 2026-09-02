import { useState } from 'react'
import { Link, NavLink, useLocation } from 'react-router-dom'
import { NAV_ITEMS, paths } from '~routes/paths'
import { Button, MenuIcon, CloseIcon, PageContainer } from '~kit'
import { cn } from '~lib/cn'
import { ChainPill } from './ChainPill'
import { GlobalSearch } from './GlobalSearch'
import { WalletButton } from './WalletButton'

function isActive(pathname: string, match: string): boolean {
  return match === '/' ? pathname === '/' : pathname === match || pathname.startsWith(`${match}/`)
}

/**
 * Sticky top bar: brand, primary nav, global search, chain identity, wallet.
 * Below 1024 px the nav folds into a disclosure and the search takes its own
 * row — everything above that is the desktop layout the design targets.
 */
export function TopBar() {
  const { pathname } = useLocation()
  const [open, setOpen] = useState(false)

  return (
    <header className='sticky top-0 z-40 border-b border-charcoal bg-obsidian/92 backdrop-blur-md'>
      <PageContainer className='flex h-14 items-center gap-4'>
        <Link to={paths.home()} className='flex shrink-0 items-center gap-2'>
          <span className='text-[15px] font-bold tracking-tight text-emerald'>davinci-dkg</span>
          <span className='label-caps hidden rounded-pill border border-emerald/20 bg-emerald/8 px-2 py-[2px] text-[9px] text-emerald sm:inline'>
            explorer
          </span>
        </Link>

        <nav className='hidden shrink-0 items-center gap-0.5 lg:flex'>
          {NAV_ITEMS.map((item) => (
            <NavLink
              key={item.to}
              to={item.to}
              className={cn(
                'rounded-sm px-3 py-1.5 text-[13px] font-medium transition-colors',
                isActive(pathname, item.match) ? 'text-emerald' : 'text-pewter hover:text-ghost'
              )}
            >
              {item.label}
            </NavLink>
          ))}
        </nav>

        <div className='ml-auto flex min-w-0 items-center gap-3'>
          <GlobalSearch className='hidden w-72 min-w-32 shrink xl:block' />
          <ChainPill className='hidden md:flex' />
          <WalletButton />
          <Button
            size='icon'
            variant='subtle'
            className='lg:hidden'
            aria-label={open ? 'Close menu' : 'Open menu'}
            aria-expanded={open}
            onClick={() => setOpen((v) => !v)}
          >
            {open ? <CloseIcon /> : <MenuIcon />}
          </Button>
        </div>
      </PageContainer>

      <PageContainer className='pb-3 xl:hidden'>
        <GlobalSearch />
      </PageContainer>

      {open ? (
        <PageContainer className='border-t border-charcoal py-2 lg:hidden'>
          <nav className='flex flex-col'>
            {NAV_ITEMS.map((item) => (
              <NavLink
                key={item.to}
                to={item.to}
                onClick={() => setOpen(false)}
                className={cn(
                  'rounded-sm px-2 py-2 text-[13px] font-medium transition-colors',
                  isActive(pathname, item.match) ? 'text-emerald' : 'text-pewter hover:text-ghost'
                )}
              >
                {item.label}
              </NavLink>
            ))}
          </nav>
          <div className='mt-2 border-t border-charcoal pt-3 md:hidden'>
            <ChainPill className='w-fit' />
          </div>
        </PageContainer>
      ) : null}
    </header>
  )
}
