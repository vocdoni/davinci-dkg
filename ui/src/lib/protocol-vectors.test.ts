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
  DomainContributionTranscriptV2,
  DomainFinalizeTranscriptV2,
  DomainDecryptCombineTranscriptV1,
  SUBGROUP_ORDER,
  BN254_Q,
  MAX_K,
  MAX_N,
  FINALIZE_TRANSCRIPT_WORDS,
  MERKLE_EMPTY_LEAF,
} from '@vocdoni/davinci-dkg-sdk'
import { MAX_COMMITTEE, POOL_SIZE } from '~indexer/types'
import { FINALIZE_TRANSCRIPT_WORDS as UI_FINALIZE_TRANSCRIPT_WORDS, contributionTranscriptWords } from '~pages/epochs/cadence'

// Vitest runs with the package root as cwd; `import.meta.url` is an http URL
// under the jsdom environment, so it cannot be used to locate files here.
const MIRROR_DIR = resolve(process.cwd(), 'tests/vectors')
const CANONICAL_DIR = resolve(process.cwd(), '../tests/vectors')
const FILES = ['protocol.json', 'schnorr.json', 'dleq.json', 'contribution_compact.json', 'finalize_transcript.json']

interface ProtocolFile {
  domains: Record<string, { preimage: string; keccak256: string; bn254Reduced: string }>
  bn254Q: string
  subgroupOrderL: string
}

interface CompactFile {
  maxN: number
  maxK: number
  vectors: Array<{ threshold: number; committeeSize: number; offsets: { words: number } }>
}

interface FinalizeFile {
  maxN: number
  maxK: number
  transcriptWords: number
  merkleEmptyLeaf: string
}

const read = (name: string): string => readFileSync(resolve(MIRROR_DIR, name), 'utf8')
const f = JSON.parse(read('protocol.json')) as ProtocolFile

describe('protocol vectors mirror', () => {
  it('is byte-identical to the canonical tests/vectors copy', () => {
    if (!existsSync(CANONICAL_DIR)) {
      // Published UI builds ship without the repo root; nothing to compare to.
      return
    }
    for (const name of FILES) {
      expect(read(name), name).toBe(readFileSync(resolve(CANONICAL_DIR, name), 'utf8'))
    }
  })

  it('pins the domain digests the SDK exports', () => {
    expect(DomainOperatorRegisterV1).toBe(f.domains.OperatorRegisterV1.keccak256)
    expect(DomainOrganizerRegisterV1).toBe(f.domains.OrganizerRegisterV1.keccak256)
    expect(DomainContributionTranscriptV2).toBe(f.domains.ContributionTranscriptV2.keccak256)
    expect(DomainFinalizeTranscriptV2).toBe(f.domains.FinalizeTranscriptV2.keccak256)
    expect(DomainDecryptCombineTranscriptV1).toBe(f.domains.DecryptCombineTranscriptV1.keccak256)
    expect(f.domains.OrganizerRegisterV1.preimage).toBe('davinci-dkg:organizer-register:v1')
    expect(f.domains.ContributionTranscriptV2.preimage).toBe('davinci-dkg:contribution:v2')
    expect(f.domains.FinalizeTranscriptV2.preimage).toBe('davinci-dkg:finalize:v2')
    // The per-ciphertext organizer share, the v1 contribution transcript and
    // the per-key activation are gone with their domains.
    expect(f.domains.OrganizerShareV1).toBeUndefined()
    expect(f.domains.ContributionTranscriptV1).toBeUndefined()
    expect(f.domains.PoolKeyTranscriptV1).toBeUndefined()
  })

  it('pins the field and the BabyJubJub subgroup order', () => {
    expect(BN254_Q.toString()).toBe(f.bn254Q)
    expect(SUBGROUP_ORDER.toString()).toBe(f.subgroupOrderL)
  })

  it('pins the circuit bounds and transcript sizes the explorer renders', () => {
    const compact = JSON.parse(read('contribution_compact.json')) as CompactFile
    const finalize = JSON.parse(read('finalize_transcript.json')) as FinalizeFile
    expect(POOL_SIZE).toBe(MAX_K)
    expect(MAX_COMMITTEE).toBe(MAX_N)
    expect(compact.maxK).toBe(POOL_SIZE)
    expect(compact.maxN).toBe(MAX_COMMITTEE)
    expect(finalize.maxK).toBe(POOL_SIZE)
    expect(finalize.maxN).toBe(MAX_COMMITTEE)
    expect(UI_FINALIZE_TRANSCRIPT_WORDS).toBe(FINALIZE_TRANSCRIPT_WORDS)
    expect(UI_FINALIZE_TRANSCRIPT_WORDS).toBe(finalize.transcriptWords)
    expect(MERKLE_EMPTY_LEAF).toBe(finalize.merkleEmptyLeaf)
    for (const v of compact.vectors) {
      expect(contributionTranscriptWords(v.threshold, v.committeeSize)).toBe(v.offsets.words)
    }
  })
})
