// Cells the applications tables and the application page render on their own.

import { Address, Badge, Hash } from '~kit'
import type { ApplicationRow } from '~indexer/selectors'
import type { AppModeName } from '~indexer/types'
import { bigIntToHex } from '~lib/format'
import { MODE_HELP, MODE_TONE, WINDOW_LABEL, WINDOW_TONE, describeDecryptionWindow, submissionPolicy } from './policy'

export function ModeBadge({ mode, size = 'sm' }: { mode: AppModeName; size?: 'sm' | 'md' }) {
  return (
    <Badge size={size} tone={MODE_TONE[mode]} title={MODE_HELP[mode]}>
      {mode}
    </Badge>
  )
}

export function SubmissionCell({ row }: { row: ApplicationRow }) {
  const policy = submissionPolicy(row)
  switch (policy.kind) {
    case 'unknown':
      return <span className='text-ash'>—</span>
    case 'open':
      return (
        <Badge size='sm' tone='warn' title='policy.openSubmission — anyone may call submitCiphertext'>
          open
        </Badge>
      )
    case 'registrant':
      return (
        <span className='text-pewter' title='policy.submitters is empty: only the registering address may submit'>
          registrant only
        </span>
      )
    case 'allow-list':
      return policy.submitters.length === 1 ? (
        <Address value={policy.submitters[0]} copy={false} explorer={false} />
      ) : (
        <span className='text-silver' title={policy.submitters.join('\n')}>
          {policy.submitters.length} addresses
        </span>
      )
  }
}

/**
 * The block window `notBeforeBlock → notAfterBlock` in which `submitCiphertext`
 * is accepted; 0 on both sides is how the contract spells "any block".
 */
export function WindowCell({ row }: { row: Pick<ApplicationRow, 'notBeforeBlock' | 'notAfterBlock'> }) {
  const from = row.notBeforeBlock
  const to = row.notAfterBlock
  if (from == null || to == null) return <span className='text-ash'>—</span>
  if (from === 0 && to === 0) return <span className='whitespace-nowrap text-ash'>any block</span>
  return (
    <span className='flex flex-col font-mono text-[12px] leading-tight tnum text-silver'>
      <span className='whitespace-nowrap'>
        {from === 0 ? 'any' : from.toLocaleString()} <span className='text-ash'>→</span>
      </span>
      <span className='whitespace-nowrap'>{to === 0 ? 'any' : to.toLocaleString()}</span>
    </span>
  )
}

/**
 * The decryption window as dates, with whether the chain accepts partials and
 * combines right now. `compact` stacks the two bounds on two lines (dates
 * only, the full date-time in the title) so a table row holds it; the default
 * is the one-line form the application page uses.
 */
export function DecryptionWindowCell({
  notBefore,
  notAfter,
  compact = false,
}: {
  notBefore: number | null
  notAfter: number | null
  compact?: boolean
}) {
  const window = describeDecryptionWindow(notBefore, notAfter, Date.now(), compact ? 'date' : 'datetime')
  if (!window) return <span className='text-ash'>—</span>
  const from = window.from ?? 'any time'
  const until = window.until ?? 'no deadline'
  const badge =
    window.state !== 'unbounded' ? (
      <Badge
        size='sm'
        tone={WINDOW_TONE[window.state]}
        title={
          window.state === 'closed'
            ? 'Past decryptNotAfter: partials and combines revert DecryptionClosed()'
            : window.state === 'not-yet-open'
              ? 'Before decryptNotBefore: partials and combines revert DecryptionNotOpen()'
              : 'Partials and combines are accepted (an organizer-locked application also needs its reveal)'
        }
      >
        {WINDOW_LABEL[window.state]}
      </Badge>
    ) : null
  if (window.state === 'unbounded') {
    return <span className='whitespace-nowrap font-mono text-[12px] tnum text-ash'>any time</span>
  }
  if (compact) {
    return (
      <span className='inline-flex items-center gap-2' title={`${window.fromFull ?? 'any time'} → ${window.untilFull ?? 'no deadline'}`}>
        <span className='flex flex-col font-mono text-[12px] leading-tight tnum text-silver'>
          <span className='whitespace-nowrap'>
            {from} <span className='text-ash'>→</span>
          </span>
          <span className='whitespace-nowrap'>{until}</span>
        </span>
        {badge}
      </span>
    )
  }
  return (
    <span className='inline-flex items-center gap-2 whitespace-nowrap font-mono text-[12px] tnum text-silver'>
      <span>
        {from} <span className='text-ash'>→</span> {until}
      </span>
      {badge}
    </span>
  )
}

/**
 * Which of the epoch's pool keys the application encrypts under: the index on
 * one line, `P_j.x` under it, so the cell holds in a narrow column.
 */
export function PoolKeyCell({ row }: { row: ApplicationRow }) {
  if (row.poolIndex == null) return <span className='text-ash'>—</span>
  return (
    <span className='flex flex-col font-mono text-[12px] leading-tight tnum'>
      <span className='whitespace-nowrap text-silver'>key {row.poolIndex}</span>
      {row.poolKey ? (
        <Hash value={bigIntToHex(row.poolKey.x)} chars={3} copy={false} className='text-[11px]' />
      ) : (
        <span className='text-[11px] text-ash' title='pool key not read yet'>
          …
        </span>
      )}
    </span>
  )
}

/** Whether the committee may combine on its own: n/a, revealed or still held. */
export function SecretCell({ row }: { row: ApplicationRow }) {
  if (row.mode === 'automatic') {
    return (
      <span className='text-ash' title='Automatic: there is no organizer key'>
        none
      </span>
    )
  }
  return row.unlocked ? (
    <Badge size='sm' tone='ok' title='sk_org is on chain — the committee combines on its own'>
      revealed
    </Badge>
  ) : (
    <Badge size='sm' tone='warn' title='sk_org not revealed — the contract refuses partials and combines until it is'>
      kept
    </Badge>
  )
}

/** Ciphertexts submitted over the policy cap (`∞` when uncapped). */
export function CiphertextCountCell({ row }: { row: Pick<ApplicationRow, 'ciphertexts' | 'maxCiphertexts'> }) {
  const cap = row.maxCiphertexts
  const capText = cap == null ? '…' : cap === 0 ? '∞' : String(cap)
  return (
    <span
      className='whitespace-nowrap font-mono text-[12px] tnum'
      title={cap ? `${row.ciphertexts} submitted · cap ${cap}` : `${row.ciphertexts} submitted · uncapped`}
    >
      {row.ciphertexts}
      <span className='text-ash'> / {capText}</span>
    </span>
  )
}
