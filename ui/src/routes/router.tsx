import { lazy, Suspense, type ReactNode } from 'react'
import { createBrowserRouter } from 'react-router-dom'
import { Shell } from '~app/Shell'
import { RouteError } from '~pages/error'
import { NotFoundPage } from '~pages/not-found'
import { SkeletonText } from '~kit'
import { patterns } from './paths'

// Route elements are lazy so the wallet-heavy playground and the /kit showcase
// never land in the first chunk a visitor downloads.
const OverviewPage = lazy(() => import('~pages/overview').then((m) => ({ default: m.OverviewPage })))
const EpochsPage = lazy(() => import('~pages/epochs').then((m) => ({ default: m.EpochsPage })))
const EpochPage = lazy(() => import('~pages/epoch').then((m) => ({ default: m.EpochPage })))
const OperatorsPage = lazy(() => import('~pages/operators').then((m) => ({ default: m.OperatorsPage })))
const OperatorPage = lazy(() => import('~pages/operator').then((m) => ({ default: m.OperatorPage })))
const ApplicationsPage = lazy(() => import('~pages/applications').then((m) => ({ default: m.ApplicationsPage })))
const ApplicationPage = lazy(() => import('~pages/application').then((m) => ({ default: m.ApplicationPage })))
const PlaygroundPage = lazy(() => import('~pages/playground').then((m) => ({ default: m.PlaygroundPage })))
const DocsProtocolPage = lazy(() => import('~pages/docs/protocol').then((m) => ({ default: m.DocsProtocolPage })))
const DocsRunANodePage = lazy(() => import('~pages/docs/run-a-node').then((m) => ({ default: m.DocsRunANodePage })))
const DocsSdkPage = lazy(() => import('~pages/docs/sdk').then((m) => ({ default: m.DocsSdkPage })))
const KitPage = lazy(() => import('~pages/kit').then((m) => ({ default: m.KitPage })))

function Loading() {
  return <SkeletonText lines={6} className='max-w-2xl' />
}

const page = (element: ReactNode) => <Suspense fallback={<Loading />}>{element}</Suspense>

export const router = createBrowserRouter([
  {
    path: '/',
    element: <Shell />,
    errorElement: <RouteError />,
    children: [
      { index: true, element: page(<OverviewPage />) },
      { path: patterns.epochs, element: page(<EpochsPage />) },
      { path: patterns.epoch, element: page(<EpochPage />) },
      { path: patterns.operators, element: page(<OperatorsPage />) },
      { path: patterns.operator, element: page(<OperatorPage />) },
      { path: patterns.applications, element: page(<ApplicationsPage />) },
      { path: patterns.application, element: page(<ApplicationPage />) },
      { path: patterns.playground, element: page(<PlaygroundPage />) },
      { path: patterns.docsProtocol, element: page(<DocsProtocolPage />) },
      { path: patterns.docsRunANode, element: page(<DocsRunANodePage />) },
      { path: patterns.docsSdk, element: page(<DocsSdkPage />) },
      { path: patterns.kit, element: page(<KitPage />) },
      { path: '*', element: <NotFoundPage /> },
    ],
  },
])
