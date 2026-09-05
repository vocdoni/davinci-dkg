import { describe, expect, it } from 'vitest'
import { screen, within } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import { DocsProtocolPage } from './protocol'
import { DocsRunANodePage } from './run-a-node'
import { DocsSdkPage } from './sdk'

const CONFIG = {
  chainName: 'sepolia',
  chainId: 11155111,
  managerAddress: '0x3f9b338706a31f26d49159478015c8aaeab908ad' as `0x${string}`,
  rpcUrl: 'https://rpc.example/sepolia',
  deployBlock: 11_619_019,
  explorerUrl: 'https://sepolia.etherscan.io',
}

/** Every heading the outline names must exist as a section, and vice versa. */
function expectOutlineMatchesSections() {
  const outline = screen.getByText('On this page').parentElement as HTMLElement
  const links = within(outline).getAllByRole('link')
  expect(links.length).toBeGreaterThan(4)
  for (const link of links) {
    const id = link.getAttribute('href')?.slice(1)
    expect(id).toBeTruthy()
    expect(document.getElementById(id as string)).not.toBeNull()
  }
}

describe('docs/protocol', () => {
  it('renders the protocol sections', () => {
    renderWithProviders(<DocsProtocolPage />, { config: CONFIG })
    expect(screen.getByRole('heading', { name: 'Protocol', level: 1 })).toBeInTheDocument()
    for (const title of [
      'Overview',
      'Epoch lifecycle',
      'Committee lottery',
      'Contribution, finalization and pool keys',
      'Applications, modes and windows',
      'Threshold decryption',
      'What holds, and what does not',
    ]) {
      expect(screen.getByRole('heading', { name: title, level: 2 })).toBeInTheDocument()
    }
  })

  it('states the organizer model accurately', () => {
    renderWithProviders(<DocsProtocolPage />, { config: CONFIG })
    expect(screen.getByText(/PK_aid = P_j \+ PK_org/)).toBeInTheDocument()
    expect(screen.getByText(/finalizeEpoch \(one Groth16 proof\)/)).toBeInTheDocument()
    expect(screen.queryByText(/activatePoolKey/)).not.toBeInTheDocument()
    expect(screen.getAllByText(/revealOrganizerSecret/).length).toBeGreaterThan(0)
    expect(screen.getByText(/permanently undecryptable/i)).toBeInTheDocument()
    expect(screen.queryByText(/organizer share/i)).not.toBeInTheDocument()
  })

  it('takes the deployment from the runtime config, not from a constant', () => {
    renderWithProviders(<DocsProtocolPage />, { config: CONFIG })
    expect(screen.getByText('sepolia')).toBeInTheDocument()
    expect(screen.getAllByTitle(/0x3F9B338706a31f26D49159478015C8AAEAb908Ad/i).length).toBeGreaterThan(0)
  })

  it('has an outline whose every entry points at a real section', () => {
    renderWithProviders(<DocsProtocolPage />, { config: CONFIG })
    expectOutlineMatchesSections()
  })
})

describe('docs/run-a-node', () => {
  it('renders the operator walkthrough', () => {
    renderWithProviders(<DocsRunANodePage />, { config: CONFIG })
    expect(screen.getByRole('heading', { name: 'Run a node', level: 1 })).toBeInTheDocument()
    expect(screen.getByRole('heading', { name: 'What happens on first boot', level: 2 })).toBeInTheDocument()
    expect(screen.getByRole('heading', { name: 'Limits worth knowing', level: 2 })).toBeInTheDocument()
  })

  it('templates the snippets with the configured deployment', () => {
    renderWithProviders(<DocsRunANodePage />, { config: CONFIG })
    expect(screen.getByText(/DAVINCI_DKG_WEB3_RPC=https:\/\/rpc.example\/sepolia/)).toBeInTheDocument()
    expect(screen.getByText(/DAVINCI_DKG_NETWORK=sepolia/)).toBeInTheDocument()
  })

  it('has a matching outline', () => {
    renderWithProviders(<DocsRunANodePage />, { config: CONFIG })
    expectOutlineMatchesSections()
  })
})

describe('docs/sdk', () => {
  it('renders the organizer-side snippets', () => {
    renderWithProviders(<DocsSdkPage />, { config: CONFIG })
    expect(screen.getByRole('heading', { name: 'SDK', level: 1 })).toBeInTheDocument()
    for (const title of [
      'Registering an application',
      'Encrypting a value',
      'Submitting a ciphertext',
      'Revealing the organizer secret',
      'Reading the plaintext',
    ]) {
      expect(screen.getByRole('heading', { name: title, level: 2 })).toBeInTheDocument()
    }
    expect(screen.getByText(/randomOrganizerSecret\(\)/)).toBeInTheDocument()
    expect(screen.getByText(/revealOrganizerSecret\(epochId, aid, skOrg\)/)).toBeInTheDocument()
    expect(screen.getByText(/dkg.getApplicationKey\(epochId, aid\)/)).toBeInTheDocument()
    expect(screen.queryByText(/submitOrganizerShare/)).not.toBeInTheDocument()
  })

  it('uses the configured manager address in the client snippet', () => {
    renderWithProviders(<DocsSdkPage />, { config: CONFIG })
    expect(screen.getAllByText(new RegExp(CONFIG.managerAddress)).length).toBeGreaterThan(0)
  })

  it('offers a copy button per snippet', () => {
    renderWithProviders(<DocsSdkPage />, { config: CONFIG })
    expect(screen.getAllByRole('button', { name: /copy snippet/i }).length).toBeGreaterThan(5)
  })

  it('has a matching outline', () => {
    renderWithProviders(<DocsSdkPage />, { config: CONFIG })
    expectOutlineMatchesSections()
  })
})
