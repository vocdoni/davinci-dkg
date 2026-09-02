// One line of English per indexed event.
//
// The operator timeline is the page that has to render *every* event type, so
// the mapping lives here as a pure function: a new event in the ABI shows up
// as a compile error in this switch rather than as a blank row.

import type { EpochId, IndexedEvent } from '~indexer/types'
import type { TimelineTone } from '~kit'
import { shortHash } from '~lib/format'
import { paths } from '~routes/paths'

export interface EventDescription {
  title: string
  detail?: string
  tone: TimelineTone
  /** In-app route for the entity this event belongs to. */
  href?: string
}

export type NonceLookup = (id: EpochId) => number | null | undefined

export function epochLabel(id: EpochId | null, nonceOf?: NonceLookup): string {
  if (!id) return 'the registry'
  const nonce = nonceOf?.(id)
  return nonce != null ? `epoch #${nonce}` : `epoch ${shortHash(id, 6, 4)}`
}

function hrefFor(event: IndexedEvent): string | undefined {
  if (event.epoch && event.aid) return paths.application(event.epoch, event.aid)
  if (event.epoch) return paths.epoch(event.epoch)
  return undefined
}

/** Title, tone and link for one event, as the timeline renders it. */
export function describeEvent(event: IndexedEvent, nonceOf?: NonceLookup): EventDescription {
  const where = epochLabel(event.epoch, nonceOf)
  const href = hrefFor(event)
  switch (event.name) {
    case 'NodeRegistered':
      return { title: 'Registered in the registry', detail: 'published a BabyJubJub encryption key', tone: 'ok' }
    case 'NodeUpdated':
      return { title: 'Updated its encryption key', tone: 'neutral' }
    case 'NodeMarkedActive':
      return { title: 'Marked active', detail: `liveness recorded at block ${event.data.atBlock}`, tone: 'ok' }
    case 'NodeReaped':
      return {
        title: 'Reaped for inactivity',
        detail: `last active at block ${event.data.lastActiveBlock}`,
        tone: 'danger',
      }
    case 'NodeReactivated':
      return { title: 'Reactivated', tone: 'ok' }
    case 'ManagerSet':
      return { title: 'Registry manager set', detail: event.data.manager, tone: 'muted' }
    case 'EpochCreated':
      return { title: `Created ${where}`, detail: `seed block ${event.data.seedBlock}`, tone: 'neutral', href }
    case 'SeedResolved':
      return {
        title: `Lottery seed resolved for ${where}`,
        detail: shortHash(event.data.seed, 10, 6),
        tone: 'neutral',
        href,
      }
    case 'SlotClaimed':
      return {
        title: `Claimed slot ${event.data.slot} in ${where}`,
        detail: `participant index ${event.data.slot + 1}`,
        tone: 'neutral',
        href,
      }
    case 'CommitteeFilled':
      return { title: `Committee filled in ${where}`, tone: 'neutral', href }
    case 'ContributionSubmitted':
      return {
        title: `Submitted contribution ${event.data.contributorIndex} in ${where}`,
        detail: `commitments ${shortHash(event.data.commitmentsHash, 6, 4)}`,
        tone: 'ok',
        href,
      }
    case 'EpochLive':
      return { title: `Finalized ${where}`, detail: 'the collective key is assembled', tone: 'ok', href }
    case 'EpochAborted':
      return { title: `Aborted ${where}`, detail: 'the epoch could no longer progress', tone: 'danger', href }
    case 'CiphertextSubmitted':
      return {
        title: `Submitted ciphertext ${event.data.ciphertextIndex}`,
        detail: `${where} · application ${shortHash(event.data.aid, 6, 4)}`,
        tone: 'neutral',
        href,
      }
    case 'PartialDecryptionSubmitted':
      return {
        title: `Published a partial for ciphertext ${event.data.ciphertextIndex}`,
        detail: `participant ${event.data.participantIndex} · ${where}`,
        tone: 'ok',
        href,
      }
    case 'DecryptionCombined':
      return {
        title: `Combined ciphertext ${event.data.ciphertextIndex}`,
        detail: `plaintext ${event.data.plaintext.toString()}`,
        tone: 'ok',
        href,
      }
    case 'ApplicationRegistered':
      return {
        title: `Registered application ${shortHash(event.data.aid, 6, 4)}`,
        detail: where,
        tone: 'neutral',
        href,
      }
    case 'OrganizerShareSubmitted':
      return {
        title: `Released the organizer share for ciphertext ${event.data.ciphertextIndex}`,
        detail: `${where} · application ${shortHash(event.data.aid, 6, 4)}`,
        tone: 'ok',
        href,
      }
    default:
      return { title: (event as IndexedEvent).name, tone: 'muted' }
  }
}
