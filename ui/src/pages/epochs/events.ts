// Event presentation shared by the overview feed and the epoch event log.
//
// The explorer never hides an event type, so every name in `EventDataMap` has
// a one-line projection of its payload here. Keeping it pure (string in,
// string out) is what makes it testable without rendering a table.

import type { IndexedEvent, IndexedEventName } from '~indexer/types'
import { shortHash } from '~lib/format'

export type EventTone = 'ok' | 'danger' | 'warn' | 'neutral'

/**
 * Emerald marks the events that mean "a protocol milestone landed"; red the
 * two failures; amber the reveal, which changes who can decrypt. Everything
 * else is routine and stays neutral.
 */
export const EVENT_TONE: Record<IndexedEventName, EventTone> = {
  NodeRegistered: 'neutral',
  NodeUpdated: 'neutral',
  NodeMarkedActive: 'neutral',
  NodeReaped: 'danger',
  NodeReactivated: 'neutral',
  ManagerSet: 'neutral',
  EpochCreated: 'neutral',
  SeedResolved: 'neutral',
  SlotClaimed: 'neutral',
  CommitteeFilled: 'ok',
  ContributionSubmitted: 'neutral',
  EpochLive: 'ok',
  PoolKeyActivated: 'ok',
  PoolKeyClaimed: 'neutral',
  CiphertextSubmitted: 'neutral',
  PartialDecryptionSubmitted: 'neutral',
  DecryptionCombined: 'ok',
  EpochAborted: 'danger',
  ApplicationRegistered: 'neutral',
  OrganizerSecretRevealed: 'warn',
}

export const EVENT_TONE_CLASS: Record<EventTone, string> = {
  ok: 'text-emerald',
  danger: 'text-red',
  warn: 'text-amber',
  neutral: 'text-silver',
}

/**
 * The payload fields that do not already have a column of their own, in one
 * line. Slots are 0-based and participant indexes 1-based, so a claim prints
 * both — that pairing is the thing readers get wrong.
 */
export function eventSummary(event: IndexedEvent): string {
  switch (event.name) {
    case 'NodeRegistered':
      return 'BabyJubJub key registered'
    case 'NodeUpdated':
      return 'encryption key replaced'
    case 'NodeMarkedActive':
      return `active at block ${event.data.atBlock.toLocaleString()}`
    case 'NodeReaped':
      return `reaped · last active ${event.data.lastActiveBlock.toLocaleString()}`
    case 'NodeReactivated':
      return 'back in the active set'
    case 'ManagerSet':
      return `manager ${shortHash(event.data.manager, 6, 4)}`
    case 'EpochCreated':
      return `start ${event.data.startBlock.toLocaleString()} · seed block ${event.data.seedBlock.toLocaleString()}`
    case 'SeedResolved':
      return `seed ${shortHash(event.data.seed, 8, 6)}`
    case 'SlotClaimed':
      return `slot ${event.data.slot} · participant index ${event.data.slot + 1}`
    case 'CommitteeFilled':
      return 'every slot claimed — key assembly opens'
    case 'ContributionSubmitted':
      return `participant index ${event.data.contributorIndex}`
    case 'EpochLive':
      return `${event.data.contributionCount} contributions frozen · pool opens`
    case 'PoolKeyActivated':
      return `key ${event.data.keyIndex} · P (${shortHash(bigintHex(event.data.key.x), 6, 4)}, …)`
    case 'PoolKeyClaimed':
      return `key ${event.data.keyIndex} claimed`
    case 'CiphertextSubmitted':
      return `ciphertext ${event.data.ciphertextIndex}`
    case 'PartialDecryptionSubmitted':
      return `ciphertext ${event.data.ciphertextIndex} · participant index ${event.data.participantIndex}`
    case 'DecryptionCombined':
      return `ciphertext ${event.data.ciphertextIndex} · m = ${event.data.plaintext.toString()}`
    case 'EpochAborted':
      return 'committee never filled'
    case 'ApplicationRegistered':
      return event.data.mode === 'automatic'
        ? `automatic · pool key ${event.data.poolIndex}`
        : `organizer-locked · pool key ${event.data.poolIndex} · PK_org (${shortHash(bigintHex(event.data.organizerPK.x), 6, 4)}, …)`
    case 'OrganizerSecretRevealed':
      return 'sk_org revealed · the committee combines on its own'
  }
}

function bigintHex(value: bigint): string {
  return `0x${value.toString(16).padStart(64, '0')}`
}
