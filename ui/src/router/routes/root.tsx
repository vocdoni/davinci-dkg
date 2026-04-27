import { lazy } from 'react'
import { createBrowserRouter, Navigate } from 'react-router-dom'
import { Routes } from './index'
import { SuspenseLoader } from '../SuspenseLoader'
import { Layout } from '~elements/layout'
import { ErrorElement } from '~elements/error'

// Each route element is lazy-loaded so a Phase-1 deployment with only the
// home page implemented doesn't pay the bundle cost of the (yet-unwritten)
// playground or registry chunks.
const Home = lazy(() => import('~elements/home').then((m) => ({ default: m.Home })))
const RoundsList = lazy(() => import('~elements/rounds/index').then((m) => ({ default: m.RoundsList })))
const RoundView = lazy(() => import('~elements/rounds/view').then((m) => ({ default: m.RoundView })))
const Registry = lazy(() => import('~elements/registry').then((m) => ({ default: m.Registry })))
const Playground = lazy(() => import('~elements/playground').then((m) => ({ default: m.Playground })))
const RunNode = lazy(() => import('~elements/run-a-node').then((m) => ({ default: m.RunNode })))
const Sdk = lazy(() => import('~elements/sdk').then((m) => ({ default: m.Sdk })))
const Settings = lazy(() => import('~elements/settings').then((m) => ({ default: m.Settings })))

export const router = createBrowserRouter([
  {
    path: Routes.home,
    element: <Layout />,
    errorElement: <ErrorElement />,
    children: [
      {
        index: true,
        element: (
          <SuspenseLoader>
            <Home />
          </SuspenseLoader>
        ),
      },
      {
        path: Routes.rounds,
        element: (
          <SuspenseLoader>
            <RoundsList />
          </SuspenseLoader>
        ),
      },
      {
        path: Routes.round,
        element: (
          <SuspenseLoader>
            <RoundView />
          </SuspenseLoader>
        ),
      },
      {
        path: Routes.registry,
        element: (
          <SuspenseLoader>
            <Registry />
          </SuspenseLoader>
        ),
      },
      {
        path: Routes.playground,
        element: (
          <SuspenseLoader>
            <Playground />
          </SuspenseLoader>
        ),
      },
      {
        path: Routes.runNode,
        element: (
          <SuspenseLoader>
            <RunNode />
          </SuspenseLoader>
        ),
      },
      {
        path: Routes.sdk,
        element: (
          <SuspenseLoader>
            <Sdk />
          </SuspenseLoader>
        ),
      },
      {
        path: Routes.settings,
        element: (
          <SuspenseLoader>
            <Settings />
          </SuspenseLoader>
        ),
      },
      // Catch-all: anything we don't know about goes home rather than
      // dumping a router error in front of the user.
      { path: '*', element: <Navigate to={Routes.home} replace /> },
    ],
  },
])
