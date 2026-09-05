// Cross-impl vector test for the UI.
//
// `ui/tests/vectors/*.json` is a copy of the canonical `tests/vectors/*.json`
// that `cmd/protocol-vectors` generates from the Go side. This test gives the
// copy a consumer: it pins the protocol constants the UI renders and fails if
// the mirror has drifted from the canonical files — so a stale copy is a red
// test rather than silently wrong documentation.

import { describe, it, expect } from 'vitest'
import { existsSync, readFileSync } from 'node:fs'
import { resolve } from 'node:path'
import {
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
  DomainContributionTranscriptV1,
  DomainPoolKeyTranscriptV1,
  DomainDecryptCombineTranscriptV1,
  SUBGROUP_ORDER,
  BN254_Q,
} from '@vocdoni/davinci-dkg-sdk'

// Vitest runs with the package root as cwd; `import.meta.url` is an http URL
// under the jsdom environment, so it cannot be used to locate files here.
const MIRROR = resolve(process.cwd(), 'tests/vectors/protocol.json')
const CANONICAL = resolve(process.cwd(), '../tests/vectors/protocol.json')

interface ProtocolFile {
  domains: Record<string, { preimage: string; keccak256: string; bn254Reduced: string }>
  bn254Q: string
  subgroupOrderL: string
}

const raw = readFileSync(MIRROR, 'utf8')
const f = JSON.parse(raw) as ProtocolFile

describe('protocol vectors mirror', () => {
  it('is byte-identical to the canonical tests/vectors copy', () => {
    if (!existsSync(CANONICAL)) {
      // Published UI builds ship without the repo root; nothing to compare to.
      return
    }
    expect(raw).toBe(readFileSync(CANONICAL, 'utf8'))
  })

  it('pins the domain digests the SDK exports', () => {
    expect(DomainOperatorRegisterV1).toBe(f.domains.OperatorRegisterV1.keccak256)
    expect(DomainOrganizerRegisterV1).toBe(f.domains.OrganizerRegisterV1.keccak256)
    expect(DomainContributionTranscriptV1).toBe(f.domains.ContributionTranscriptV1.keccak256)
    expect(DomainPoolKeyTranscriptV1).toBe(f.domains.PoolKeyTranscriptV1.keccak256)
    expect(DomainDecryptCombineTranscriptV1).toBe(f.domains.DecryptCombineTranscriptV1.keccak256)
    expect(f.domains.OrganizerRegisterV1.preimage).toBe('davinci-dkg:organizer-register:v1')
    // The per-ciphertext organizer share, and its domain, are gone: the
    // organizer reveals `sk_org` once instead.
    expect(f.domains.OrganizerShareV1).toBeUndefined()
  })

  it('pins the field and the BabyJubJub subgroup order', () => {
    expect(BN254_Q.toString()).toBe(f.bn254Q)
    expect(SUBGROUP_ORDER.toString()).toBe(f.subgroupOrderL)
  })
})
