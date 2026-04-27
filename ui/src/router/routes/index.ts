// Single source of truth for every URL the app can render. Components must
// reference these constants via generatePath / Link rather than hard-coding
// strings, so a route rename is a one-place change.

export const Routes = {
  home: '/',
  rounds: '/rounds',
  round: '/rounds/:id',
  registry: '/registry',
  playground: '/playground',
  runNode: '/run-a-node',
  sdk: '/sdk',
  settings: '/settings',
} as const

export type RouteKey = keyof typeof Routes
