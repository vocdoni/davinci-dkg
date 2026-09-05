// Event catalogue and log normalisation.
//
// The indexer asks for every event of the three contracts in one `getLogs`
// filter per chunk, so the ABI fragments live here in one list and the raw
// decoded logs are turned into the flat `IndexedEvent` union the reducer
// understands. Event names are unique across DKGRegistry / DKGManager /
// DKGAppManager, which is what lets a single filter cover all three.

import { dkgAppManagerAbi, dkgManagerAbi, dkgRegistryAbi, fromRTEtoTE } from '@vocdoni/davinci-dkg-sdk'
import type { AbiEvent } from 'viem'
import { appModeName, type Address, type Hex, type IndexedEvent, type IndexedEventName, type Point } from './types'

/** Every event fragment of the three contracts, in ABI order. */
export const ALL_EVENT_ABIS: AbiEvent[] = [
  ...(dkgRegistryAbi as readonly unknown[]),
  ...(dkgManagerAbi as readonly unknown[]),
  ...(dkgAppManagerAbi as readonly unknown[]),
].filter((item): item is AbiEvent => (item as AbiEvent).type === 'event')

export const REGISTRY_EVENT_ABIS: AbiEvent[] = (dkgRegistryAbi as readonly unknown[]).filter(
  (item): item is AbiEvent => (item as AbiEvent).type === 'event',
)
export const MANAGER_EVENT_ABIS: AbiEvent[] = (dkgManagerAbi as readonly unknown[]).filter(
  (item): item is AbiEvent => (item as AbiEvent).type === 'event',
)
export const APP_MANAGER_EVENT_ABIS: AbiEvent[] = (dkgAppManagerAbi as readonly unknown[]).filter(
  (item): item is AbiEvent => (item as AbiEvent).type === 'event',
)

const KNOWN_EVENTS = new Set<string>([
  'NodeRegistered',
  'NodeUpdated',
  'NodeMarkedActive',
  'NodeReaped',
  'NodeReactivated',
  'ManagerSet',
  'EpochCreated',
  'SeedResolved',
  'SlotClaimed',
  'CommitteeFilled',
  'ContributionSubmitted',
  'EpochLive',
  'PoolKeyClaimed',
  'CiphertextSubmitted',
  'PartialDecryptionSubmitted',
  'DecryptionCombined',
  'EpochAborted',
  'ApplicationRegistered',
  'OrganizerSecretRevealed',
])

export function isKnownEvent(name: string | undefined): name is IndexedEventName {
  return name != null && KNOWN_EVENTS.has(name)
}

/** The shape of a viem log after decoding, loosened to what we actually read. */
export interface RawLog {
  eventName?: string
  args?: Record<string, unknown>
  blockNumber?: bigint | null
  transactionHash?: Hex | null
  logIndex?: number | null
  address?: string
}

function num(v: unknown): number {
  if (typeof v === 'bigint') return Number(v)
  if (typeof v === 'number') return v
  return Number(v ?? 0)
}

function big(v: unknown): bigint {
  if (typeof v === 'bigint') return v
  if (typeof v === 'number') return BigInt(v)
  if (typeof v === 'string' && v.length > 0) return BigInt(v)
  return 0n
}

function addr(v: unknown): Address {
  return String(v ?? '0x0000000000000000000000000000000000000000').toLowerCase() as Address
}

function hex(v: unknown): Hex {
  return String(v ?? '0x') as Hex
}

function point(x: unknown, y: unknown): Point {
  return { x: big(x), y: big(y) }
}

/**
 * Turn one decoded log into an `IndexedEvent`, or null when the log is not
 * one of ours (a stray topic collision, or an event added after this build).
 */
export function normalizeLog(log: RawLog): IndexedEvent | null {
  const name = log.eventName
  if (!isKnownEvent(name)) return null
  const a = log.args ?? {}
  const envelope = {
    block: num(log.blockNumber),
    tx: (log.transactionHash ?? null) as Hex | null,
    logIndex: log.logIndex ?? 0,
  }
  const epochId = a.epochId != null ? (hex(a.epochId).toLowerCase() as Hex) : null
  const aid = a.aid != null ? (hex(a.aid).toLowerCase() as Hex) : null

  switch (name) {
    case 'NodeRegistered':
    case 'NodeUpdated':
      return {
        ...envelope,
        name,
        epoch: null,
        aid: null,
        actor: addr(a.operator),
        data: { operator: addr(a.operator), pubX: big(a.pubX), pubY: big(a.pubY) },
      }
    case 'NodeMarkedActive':
      return {
        ...envelope,
        name,
        epoch: null,
        aid: null,
        actor: addr(a.operator),
        data: { operator: addr(a.operator), atBlock: num(a.atBlock) },
      }
    case 'NodeReaped':
      return {
        ...envelope,
        name,
        epoch: null,
        aid: null,
        actor: addr(a.operator),
        data: { operator: addr(a.operator), lastActiveBlock: num(a.lastActiveBlock) },
      }
    case 'NodeReactivated':
      return {
        ...envelope,
        name,
        epoch: null,
        aid: null,
        actor: addr(a.operator),
        data: { operator: addr(a.operator) },
      }
    case 'ManagerSet':
      return {
        ...envelope,
        name,
        epoch: null,
        aid: null,
        actor: null,
        data: { manager: addr(a.manager) },
      }
    case 'EpochCreated':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid: null,
        actor: addr(a.organizer),
        data: {
          epochId: epochId as Hex,
          organizer: addr(a.organizer),
          startBlock: num(a.startBlock),
          seedBlock: num(a.seedBlock),
          lotteryThreshold: big(a.lotteryThreshold),
        },
      }
    case 'SeedResolved':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid: null,
        actor: null,
        data: { epochId: epochId as Hex, seed: hex(a.seed) },
      }
    case 'SlotClaimed':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid: null,
        actor: addr(a.claimer),
        data: { epochId: epochId as Hex, claimer: addr(a.claimer), slot: num(a.slot) },
      }
    case 'CommitteeFilled':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid: null,
        actor: null,
        data: { epochId: epochId as Hex },
      }
    case 'ContributionSubmitted':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid: null,
        actor: addr(a.contributor),
        data: {
          epochId: epochId as Hex,
          contributor: addr(a.contributor),
          contributorIndex: num(a.contributorIndex),
          commitmentsHash: hex(a.commitmentsHash),
          encryptedSharesHash: hex(a.encryptedSharesHash),
        },
      }
    case 'EpochLive':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid: null,
        actor: null,
        data: { epochId: epochId as Hex, contributionCount: num(a.contributionCount) },
      }
    case 'PoolKeyClaimed':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid,
        actor: null,
        data: { epochId: epochId as Hex, aid: aid as Hex, keyIndex: num(a.keyIndex) },
      }
    case 'CiphertextSubmitted':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid,
        actor: addr(a.submitter),
        data: {
          epochId: epochId as Hex,
          aid: aid as Hex,
          ciphertextIndex: num(a.ciphertextIndex),
          submitter: addr(a.submitter),
          c1: point(a.c1x, a.c1y),
          c2: point(a.c2x, a.c2y),
        },
      }
    case 'PartialDecryptionSubmitted':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid,
        actor: addr(a.participant),
        data: {
          epochId: epochId as Hex,
          aid: aid as Hex,
          participant: addr(a.participant),
          participantIndex: num(a.participantIndex),
          ciphertextIndex: num(a.ciphertextIndex),
          delta: point(a.deltaX, a.deltaY),
        },
      }
    case 'DecryptionCombined':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid,
        actor: null,
        data: {
          epochId: epochId as Hex,
          aid: aid as Hex,
          ciphertextIndex: num(a.ciphertextIndex),
          combineHash: hex(a.combineHash),
          plaintext: big(a.plaintext),
        },
      }
    case 'EpochAborted':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid: null,
        actor: null,
        data: { epochId: epochId as Hex },
      }
    case 'ApplicationRegistered': {
      // The log carries the on-chain (RTE) words; the store keeps TE form so
      // it composes with the SDK's ElGamal helpers, exactly like
      // `DKGClient.getApplication`.
      const [x, y] = fromRTEtoTE(big(a.organizerPKx), big(a.organizerPKy))
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid,
        actor: addr(a.creator),
        data: {
          epochId: epochId as Hex,
          aid: aid as Hex,
          creator: addr(a.creator),
          organizerPK: { x, y },
          mode: appModeName(num(a.mode)),
          poolIndex: num(a.poolIndex),
        },
      }
    }
    case 'OrganizerSecretRevealed':
      return {
        ...envelope,
        name,
        epoch: epochId,
        aid,
        actor: null,
        data: { epochId: epochId as Hex, aid: aid as Hex, organizerSecret: big(a.organizerSecret) },
      }
    default:
      return null
  }
}

/** Chronological order: block, then log index within the block. */
export function compareEvents(a: IndexedEvent, b: IndexedEvent): number {
  if (a.block !== b.block) return a.block - b.block
  return a.logIndex - b.logIndex
}
