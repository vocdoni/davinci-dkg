import { EpochPhase, type Epoch } from '@vocdoni/davinci-dkg-sdk'

// Plain-English summaries and phase math for an Epoch. Pure functions so they
// can drive both the epochs-list cards and the epoch-detail header without
// duplicating logic.

export type EpochPhase =
  | 'committee-selection'
  | 'key-assembly'
  | 'live'
  | 'completed'
  | 'aborted'
  | 'unknown'

export function roundPhase(epoch: Epoch): EpochPhase {
  switch (epoch.status) {
    case EpochPhase.CommitteeSelection:
      return 'committee-selection'
    case EpochPhase.KeyAssembly:
      return 'key-assembly'
    case EpochPhase.Live:
      return 'live'
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
    case 'committee-selection':
      return 'Committee Selection'
    case 'key-assembly':
      return 'Key Assembly'
    case 'live':
      return 'Live'
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
    case EpochPhase.CommitteeSelection: {
      const remaining = currentBlock ? Number(epoch.policy.committeeSelectionDeadlineBlock - currentBlock) : null
      const claimed = epoch.claimedCount
      const size = epoch.policy.committeeSize
      const blocks = remaining != null && remaining > 0 ? ` (closes in ~${remaining} blocks)` : ''
      return `Waiting for nodes to claim committee slots — ${claimed}/${size} claimed${blocks}.`
    }
    case EpochPhase.KeyAssembly: {
      const need = epoch.policy.minValidContributions
      const have = epoch.contributionCount
      if (have >= need) {
        const block = epoch.policy.liveNotBeforeBlock
        return `Threshold met (${have}/${need}). Epoch goes Live at block #${block.toString()}.`
      }
      return `Awaiting contributions — ${have}/${need} accepted so far.`
    }
    case EpochPhase.Live:
      return 'Collective public key is live; apps can register and ciphertexts can be decrypted.'
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
    case 'committee-selection':
      return 'yellow'
    case 'key-assembly':
      return 'blue'
    case 'live':
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
export const phaseSequence: EpochPhase[] = ['committee-selection', 'key-assembly', 'live', 'completed']

/**
 * An epoch is dead when its current phase deadline has passed without the
 * participation the next phase needs. The contract doesn't auto-flip the
 * epoch to Aborted in this case — it just sits stuck until someone calls
 * `abortEpoch` — so the UI has to recognise it and surface the failure.
 *
 * Mirrors `DKGManager.abortEpoch`'s dead-epoch predicate exactly: anyone may
 * abort an epoch for which this returns non-null, and abort reverts for any
 * other epoch.
 *   - CommitteeSelection: the selection deadline passed. (Filling the
 *     committee moves the epoch to KeyAssembly, so still being in this phase
 *     past the deadline means the committee never filled.)
 *   - KeyAssembly: the assembly deadline passed with fewer than
 *     minValidContributions accepted.
 *
 * Returns null when the epoch is healthy or already aborted/completed.
 */
export function roundFailure(
  epoch: Epoch,
  currentBlock: bigint | null
): { kind: 'committee-selection' | 'key-assembly'; have: number; need: number; total: number } | null {
  if (currentBlock == null) return null
  if (epoch.status === EpochPhase.CommitteeSelection) {
    if (currentBlock > epoch.policy.committeeSelectionDeadlineBlock) {
      const have = epoch.claimedCount
      const need = epoch.policy.committeeSize
      return { kind: 'committee-selection', have, need, total: epoch.policy.committeeSize }
    }
  } else if (epoch.status === EpochPhase.KeyAssembly) {
    if (currentBlock > epoch.policy.keyAssemblyDeadlineBlock) {
      const have = epoch.contributionCount
      const need = epoch.policy.minValidContributions
      if (have < need) {
        return { kind: 'key-assembly', have, need, total: epoch.policy.committeeSize }
      }
    }
  }
  return null
}
