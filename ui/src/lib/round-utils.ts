import { RoundStatus, type Round } from '@vocdoni/davinci-dkg-sdk'

// Plain-English summaries and phase math for a Round. Pure functions so they
// can drive both the rounds-list cards and the round-detail header without
// duplicating logic.

export type RoundPhase = 'registration' | 'contribution' | 'finalized' | 'completed' | 'aborted' | 'unknown'

export function roundPhase(round: Round): RoundPhase {
  switch (round.status) {
    case RoundStatus.Registration:
      return 'registration'
    case RoundStatus.Contribution:
      return 'contribution'
    case RoundStatus.Finalized:
      return 'finalized'
    case RoundStatus.Completed:
      return 'completed'
    case RoundStatus.Aborted:
      return 'aborted'
    default:
      return 'unknown'
  }
}

/** Human-readable label per round phase. */
export function roundPhaseLabel(phase: RoundPhase): string {
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

/** One-sentence "what is this round waiting for" for the round detail header. */
export function roundSummary(round: Round, currentBlock: bigint | null): string {
  switch (round.status) {
    case RoundStatus.Registration: {
      const remaining = currentBlock ? Number(round.policy.registrationDeadlineBlock - currentBlock) : null
      const claimed = round.claimedCount
      const size = round.policy.committeeSize
      const blocks = remaining != null && remaining > 0 ? ` (closes in ~${remaining} blocks)` : ''
      return `Waiting for nodes to claim committee slots — ${claimed}/${size} claimed${blocks}.`
    }
    case RoundStatus.Contribution: {
      const need = round.policy.minValidContributions
      const have = round.contributionCount
      if (have >= need) {
        const block = round.policy.finalizeNotBeforeBlock
        return `Threshold met (${have}/${need}). Finalize unlocks at block #${block.toString()}.`
      }
      return `Awaiting contributions — ${have}/${need} accepted so far.`
    }
    case RoundStatus.Finalized:
      return 'Round finalized. Collective public key is live; awaiting ciphertext submissions.'
    case RoundStatus.Completed:
      return 'Round completed. All ciphertexts have been threshold-decrypted.'
    case RoundStatus.Aborted:
      return 'Round was aborted before completion.'
    default:
      return 'Unknown round status.'
  }
}

/** Color palette key for the StatusBadge per phase. Keeps Chakra colour choices centralised. */
export function roundPhaseColor(phase: RoundPhase): string {
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
export const phaseSequence: RoundPhase[] = ['registration', 'contribution', 'finalized', 'completed']
