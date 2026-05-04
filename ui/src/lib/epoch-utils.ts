import { EpochPhase, type Epoch } from '@vocdoni/davinci-dkg-sdk'

// Plain-English summaries and phase math for a Epoch. Pure functions so they
// can drive both the epochs-list cards and the epoch-detail header without
// duplicating logic.

export type EpochPhase = 'registration' | 'contribution' | 'finalized' | 'completed' | 'aborted' | 'unknown'

export function roundPhase(epoch: Epoch): EpochPhase {
  switch (epoch.status) {
    case EpochPhase.Registration:
      return 'registration'
    case EpochPhase.Contribution:
      return 'contribution'
    case EpochPhase.Finalized:
      return 'finalized'
    case EpochPhase.Completed:
      return 'completed'
    case EpochPhase.Aborted:
      return 'aborted'
    default:
      return 'unknown'
  }
}

/** Human-readable label per epoch phase. */
export function roundPhaseLabel(phase: EpochPhase): string {
  switch (phase) {
    case 'registration':
      return 'Registration'
    case 'contribution':
      return 'Contribution'
    case 'finalized':
      return 'Finalized'
    case 'completed':
      return 'Completed'
    case 'aborted':
      return 'Aborted'
    case 'unknown':
      return 'Unknown'
  }
}

/** One-sentence "what is this epoch waiting for" for the epoch detail header. */
export function roundSummary(epoch: Epoch, currentBlock: bigint | null): string {
  switch (epoch.status) {
    case EpochPhase.Registration: {
      const remaining = currentBlock ? Number(epoch.policy.registrationDeadlineBlock - currentBlock) : null
      const claimed = epoch.claimedCount
      const size = epoch.policy.committeeSize
      const blocks = remaining != null && remaining > 0 ? ` (closes in ~${remaining} blocks)` : ''
      return `Waiting for nodes to claim committee slots — ${claimed}/${size} claimed${blocks}.`
    }
    case EpochPhase.Contribution: {
      const need = epoch.policy.minValidContributions
      const have = epoch.contributionCount
      if (have >= need) {
        const block = epoch.policy.finalizeNotBeforeBlock
        return `Threshold met (${have}/${need}). Finalize unlocks at block #${block.toString()}.`
      }
      return `Awaiting contributions — ${have}/${need} accepted so far.`
    }
    case EpochPhase.Finalized:
      return 'Epoch finalized. Collective public key is live; awaiting ciphertext submissions.'
    case EpochPhase.Completed:
      return 'Epoch completed. All ciphertexts have been threshold-decrypted.'
    case EpochPhase.Aborted:
      return 'Epoch was aborted before completion.'
    default:
      return 'Unknown epoch status.'
  }
}

/** Color palette key for the StatusBadge per phase. Keeps Chakra colour choices centralised. */
export function roundPhaseColor(phase: EpochPhase): string {
  switch (phase) {
    case 'registration':
      return 'yellow'
    case 'contribution':
      return 'blue'
    case 'finalized':
      return 'cyan'
    case 'completed':
      return 'green'
    case 'aborted':
      return 'red'
    case 'unknown':
      return 'gray'
  }
}

/** Steps in the canonical phase timeline, in order. Drives PhaseTimeline rendering. */
export const phaseSequence: EpochPhase[] = ['registration', 'contribution', 'finalized', 'completed']

/**
 * A epoch has effectively failed when its current phase deadline has passed
 * without enough nodes participating to make the next phase viable. The
 * contract doesn't auto-flip the epoch to Aborted in this case — it just
 * sits stuck — so the UI has to recognise it and surface the failure.
 *
 * Returns null when the epoch is healthy or already aborted/completed.
 */
export function roundFailure(
  epoch: Epoch,
  currentBlock: bigint | null
): { kind: 'registration' | 'contribution'; have: number; need: number; total: number } | null {
  if (currentBlock == null) return null
  if (epoch.status === EpochPhase.Registration) {
    if (currentBlock > epoch.policy.registrationDeadlineBlock) {
      const have = epoch.claimedCount
      const need = epoch.policy.minValidContributions
      if (have < need) {
        return { kind: 'registration', have, need, total: epoch.policy.committeeSize }
      }
    }
  } else if (epoch.status === EpochPhase.Contribution) {
    if (currentBlock > epoch.policy.contributionDeadlineBlock) {
      const have = epoch.contributionCount
      const need = epoch.policy.minValidContributions
      if (have < need) {
        return { kind: 'contribution', have, need, total: epoch.policy.committeeSize }
      }
    }
  }
  return null
}
