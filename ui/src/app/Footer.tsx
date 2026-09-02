import { Link } from 'react-router-dom'
import { paths } from '~routes/paths'
import { PageContainer } from '~kit'

/** Design-system credit, source links and the route to the `/kit` showcase. */
export function Footer() {
  return (
    <footer className='mt-16 border-t border-charcoal py-6'>
      <PageContainer className='flex flex-wrap items-center justify-between gap-4'>
        <p className='text-[12px] text-ash'>
          davinci-dkg explorer · design system by{' '}
          <a
            href='https://github.com/VoltAgent/voltagent'
            target='_blank'
            rel='noreferrer noopener'
            className='text-pewter transition-colors hover:text-emerald'
          >
            VoltAgent
          </a>{' '}
          · Inter &amp; JetBrains Mono
        </p>
        <nav className='flex flex-wrap items-center gap-4 text-[12px]'>
          <Link to={paths.docsProtocol()} className='text-pewter transition-colors hover:text-emerald'>
            Protocol
          </Link>
          <Link to={paths.docsRunANode()} className='text-pewter transition-colors hover:text-emerald'>
            Run a node
          </Link>
          <Link to={paths.docsSdk()} className='text-pewter transition-colors hover:text-emerald'>
            SDK
          </Link>
          <a
            href='https://github.com/vocdoni/davinci-dkg'
            target='_blank'
            rel='noreferrer noopener'
            className='text-pewter transition-colors hover:text-emerald'
          >
            Source
          </a>
          <Link to={paths.kit()} className='text-pewter transition-colors hover:text-emerald'>
            Design kit
          </Link>
        </nav>
      </PageContainer>
    </footer>
  )
}
