// Cross-impl vector test for the UI.
//
// `ui/tests/vectors/*.json` is a copy of the canonical `tests/vectors/*.json`
// that `cmd/protocol-vectors` generates from the Go side. This test gives the
// copy a consumer: it pins the protocol constants the UI renders and the
// organizer-share encoding the browser performs, and it fails if the mirror
// has drifted from the canonical files — so a stale copy is a red test rather
// than silently wrong documentation.

import { describe, it, expect } from 'vitest'
import { existsSync, readFileSync } from 'node:fs'
import { resolve } from 'node:path'
import {
  DomainOperatorRegisterV1,
  DomainOrganizerRegisterV1,
  DomainDLEQV1,
  DomainOrganizerShareV1,
  SUBGROUP_ORDER,
  organizerShareChallenge,
  proveOrganizerShare,
  verifyOrganizerShare,
  fromRTEtoTE,
  fromTEtoRTE,
  type BabyJubPoint,
} from '@vocdoni/davinci-dkg-sdk'
import { Base8, mulPointEscalar } from '@zk-kit/baby-jubjub'
import type { Hex } from 'viem'

// Vitest runs with the package root as cwd; `import.meta.url` is an http URL
// under the jsdom environment, so it cannot be used to locate files here.
const MIRROR = resolve(process.cwd(), 'tests/vectors/protocol.json')
const CANONICAL = resolve(process.cwd(), '../tests/vectors/protocol.json')

interface ProtocolFile {
  domains: Record<string, { preimage: string; keccak256: string; bn254Reduced: string }>
  bn254Q: string
  subgroupOrderL: string
  organizerShare: {
    domain: string
    epochId: string
    aid: string
    ctIdx: number
    skOrg: string
    w: string
    pkOrgX: string
    pkOrgY: string
    c1x: string
    c1y: string
    deltaX: string
    deltaY: string
    a1x: string
    a1y: string
    a2x: string
    a2y: string
    e: string
    z: string
  }
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
    expect(DomainDLEQV1).toBe(f.domains.DLEQV1.keccak256)
    expect(DomainOrganizerShareV1).toBe(f.domains.OrganizerShareV1.keccak256)
    expect(f.domains.OrganizerShareV1.preimage).toBe('davinci-dkg:organizer-share:v1')
  })

  it('pins the BabyJubJub subgroup order', () => {
    expect(SUBGROUP_ORDER.toString()).toBe(f.subgroupOrderL)
  })
})

describe('organizer share encoding (as the browser computes it)', () => {
  const v = f.organizerShare
  const epochId = v.epochId as Hex
  const aid = v.aid as Hex
  const pkOrg: BabyJubPoint = [BigInt(v.pkOrgX), BigInt(v.pkOrgY)]
  const c1: BabyJubPoint = [BigInt(v.c1x), BigInt(v.c1y)]
  const delta: BabyJubPoint = [BigInt(v.deltaX), BigInt(v.deltaY)]
  const a1: BabyJubPoint = [BigInt(v.a1x), BigInt(v.a1y)]
  const a2: BabyJubPoint = [BigInt(v.a2x), BigInt(v.a2y)]

  it('binds the OrganizerShareV1 domain', () => {
    expect(v.domain).toBe(DomainOrganizerShareV1)
  })

  it('PK_org is sk_org·G in on-chain coordinates', () => {
    const te = mulPointEscalar(Base8, BigInt(v.skOrg))
    expect(fromTEtoRTE(te[0], te[1])).toEqual(pkOrg)
  })

  it('reproduces the Go challenge and verifies the proof', () => {
    expect(organizerShareChallenge(epochId, aid, v.ctIdx, pkOrg, c1, delta, a1, a2).toString()).toBe(
      v.e,
    )
    expect(
      verifyOrganizerShare(epochId, aid, v.ctIdx, pkOrg, c1, delta, { a1, a2, z: BigInt(v.z) }),
    ).toBe(true)
  })

  it('the in-browser prover reproduces every word', () => {
    const c1TE = fromRTEtoTE(c1[0], c1[1]) as BabyJubPoint
    const share = proveOrganizerShare(epochId, aid, v.ctIdx, BigInt(v.skOrg), c1TE, BigInt(v.w))
    expect(share.delta).toEqual(delta)
    expect(share.a1).toEqual(a1)
    expect(share.a2).toEqual(a2)
    expect(share.z.toString()).toBe(v.z)
  })
})
