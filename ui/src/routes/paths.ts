// Every URL the explorer can render, in one place. Components build links with
// these helpers rather than string literals, so a rename is a one-file change.

export const patterns = {
  home: '/',
  epochs: '/epochs',
  epoch: '/epochs/:id',
  operators: '/operators',
  operator: '/operators/:address',
  applications: '/applications',
  application: '/applications/:epoch/:aid',
  playground: '/playground',
  docsProtocol: '/docs/protocol',
  docsRunANode: '/docs/run-a-node',
  docsSdk: '/docs/sdk',
  kit: '/kit',
} as const

export const paths = {
  home: () => patterns.home,
  epochs: () => patterns.epochs,
  epoch: (id: string) => `/epochs/${id}`,
  operators: () => patterns.operators,
  operator: (address: string) => `/operators/${address}`,
  applications: () => patterns.applications,
  application: (epochId: string, aid: string) => `/applications/${epochId}/${aid}`,
  playground: (params?: { epoch?: string; aid?: string }) => {
    const search = new URLSearchParams()
    if (params?.epoch) search.set('epoch', params.epoch)
    if (params?.aid) search.set('aid', params.aid)
    const qs = search.toString()
    return qs ? `${patterns.playground}?${qs}` : patterns.playground
  },
  docsProtocol: () => patterns.docsProtocol,
  docsRunANode: () => patterns.docsRunANode,
  docsSdk: () => patterns.docsSdk,
  kit: () => patterns.kit,
} as const

export interface NavItem {
  label: string
  to: string
  /** Marks the item active for any path under this prefix. */
  match: string
}

/** Primary navigation, in bar order. */
export const NAV_ITEMS: NavItem[] = [
  { label: 'Overview', to: patterns.home, match: '/' },
  { label: 'Epochs', to: patterns.epochs, match: '/epochs' },
  { label: 'Operators', to: patterns.operators, match: '/operators' },
  { label: 'Applications', to: patterns.applications, match: '/applications' },
  { label: 'Playground', to: patterns.playground, match: '/playground' },
  { label: 'Docs', to: patterns.docsProtocol, match: '/docs' },
]
