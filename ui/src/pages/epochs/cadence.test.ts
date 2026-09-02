import { describe, expect, it } from 'vitest'
import type { NetworkStats } from '~indexer/selectors'
import type { EpochEntity, EpochPhaseName, EpochPolicy } from '~indexer/types'
import {
  blockOrNull,
  countdownTo,
  elapsedSince,
  formatCountdown,
  nextEpochCountdown,
  phaseCountdown,
  transcriptWords,
} from './cadence'

const policy: EpochPolicy = {
  threshold: 33,
  committeeSize: 64,
  minValidContributions: 40,
  lotteryAlphaBps: 15_000,
  committeeSelectionDeadlineBlock: 1_025,
  keyAssemblyDeadlineBlock: 1_050,
  liveNotBeforeBlock: 1_055,
}

function epoch(status: EpochPhaseName, overrides: Partial<EpochEntity> = {}): EpochEntity {
  return {
    id: '0x2f1105e9000000000000000a',
    nonce: 10,
    creator: '0x0000000000000000000000000000000000000001',
    createdBlock: 1_000,
    createdTx: null,
    startBlock: 1_000,
    seedBlock: 1_001,
    seed: null,
    seedResolvedBlock: null,
    lotteryThreshold: 1n,
    status,
    policy,
    registrySnapshot: 300,
    committee: [],
    committeeFilledBlock: null,
    abortedBlock: null,
    slots: [],
    contributions: [],
    finalization: null,
    collectivePublicKey: null,
    shareCommitmentHashes: [],
    applications: [],
    counts: { claims: 0, contributions: 0, ciphertexts: 0, partials: 0, combines: 0, applications: 0 },
    events: [],
    stateBlock: 0,
    ...overrides,
  }
}

function stats(overrides: Partial<NetworkStats> = {}): NetworkStats {
  return {
    chainId: 11155111,
    chainName: 'sepolia',
    managerAddress: '0x0000000000000000000000000000000000000002',
    registryAddress: null,
    appManagerAddress: null,
    explorerUrl: '',
    headBlock: 1_010,
    lastIndexedBlock: 1_010,
    deployBlock: 0,
    blockTimeSeconds: 12,
    epochDurationBlocks: 300,
    nextEpochStartBlock: null,
    blocksToNextEpoch: null,
    operatorsRegistered: 300,
    operatorsActive: 271,
    operatorsInactive: 29,
    inactivityWindow: 50_400,
    epochs: 8,
    epochsLive: 3,
    epochsAborted: 1,
    newestEpoch: epoch('live'),
    thresholdInForce: 33,
    committeeSizeInForce: 64,
    applications: 4,
    ciphertexts: 32,
    ciphertextsDecrypted: 20,
    partials: 400,
    contributions: 300,
    claims: 320,
    events: 900,
    ...overrides,
  }
}

describe('countdownTo', () => {
  it('counts the blocks left and never goes negative', () => {
    expect(countdownTo('x', 1_100, 1_010)).toEqual({ label: 'x', targetBlock: 1_100, blocks: 90, passed: false })
    expect(countdownTo('x', 1_000, 1_010)).toEqual({ label: 'x', targetBlock: 1_000, blocks: 0, passed: true })
  })

  it('is null when either end of the interval is unknown', () => {
    expect(countdownTo('x', null, 1_010)).toBeNull()
    expect(countdownTo('x', 1_100, Number.NaN)).toBeNull()
  })
})

describe('phaseCountdown', () => {
  it('races the deadline of the phase the epoch is actually in', () => {
    expect(phaseCountdown(epoch('committee-selection'), 1_010, 300)).toMatchObject({
      label: 'selection closes',
      targetBlock: 1_025,
      blocks: 15,
    })
    expect(phaseCountdown(epoch('key-assembly'), 1_010, 300)).toMatchObject({
      targetBlock: 1_050,
      blocks: 40,
    })
    // Live counts down to startBlock + EPOCH_DURATION_BLOCKS, not to a policy field.
    expect(phaseCountdown(epoch('live'), 1_010, 300)).toMatchObject({ targetBlock: 1_300, blocks: 290 })
  })

  it('has nothing to wait for in a terminal phase', () => {
    expect(phaseCountdown(epoch('aborted'), 1_010, 300)).toBeNull()
    expect(phaseCountdown(epoch('completed'), 1_010, 300)).toBeNull()
    expect(phaseCountdown(null, 1_010, 300)).toBeNull()
  })

  it('falls back to null when the duration immutable is unknown', () => {
    expect(phaseCountdown(epoch('live'), 1_010, null)).toBeNull()
  })
})

describe('nextEpochCountdown', () => {
  it('prefers the manager’s own nextEpochStartBlock', () => {
    expect(nextEpochCountdown(stats({ nextEpochStartBlock: 1_200 }))).toEqual({
      label: 'next epoch',
      targetBlock: 1_200,
      blocks: 190,
      passed: false,
      source: 'chain',
    })
  })

  it('derives the cadence from the newest epoch when the manager has not been read', () => {
    expect(nextEpochCountdown(stats())).toEqual({
      label: 'next epoch',
      targetBlock: 1_300,
      blocks: 290,
      passed: false,
      source: 'cadence',
    })
  })

  it('is null with no epoch and no schedule', () => {
    expect(nextEpochCountdown(stats({ newestEpoch: null }))).toBeNull()
  })
})

describe('formatting helpers', () => {
  it('prints blocks and a wall-clock estimate', () => {
    expect(formatCountdown(countdownTo('x', 1_060, 1_010), 12)).toBe('50 blocks · ~10 min')
    expect(formatCountdown(countdownTo('x', 1_000, 1_010), 12)).toBe('due')
    expect(formatCountdown(null)).toBe('—')
  })

  it('turns a "live since" block into an age', () => {
    expect(elapsedSince(1_000, 1_010, 12)).toEqual({ blocks: 10, text: '~2 min ago' })
    expect(elapsedSince(1_010, 1_010, 12)).toEqual({ blocks: 0, text: 'just now' })
    expect(elapsedSince(null, 1_010, 12)).toBeNull()
  })

  it('rejects non-blocks so no page ever prints NaN', () => {
    expect(blockOrNull(Number.NaN)).toBeNull()
    expect(blockOrNull(undefined)).toBeNull()
    expect(blockOrNull(0)).toBe(0)
  })
})

describe('transcriptWords', () => {
  it('is 2·n² + 5·n', () => {
    expect(transcriptWords(64)).toBe(2 * 64 * 64 + 5 * 64)
    expect(transcriptWords(1)).toBe(7)
    expect(transcriptWords(0)).toBe(0)
  })
})
