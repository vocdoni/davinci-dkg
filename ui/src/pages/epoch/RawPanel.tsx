import { KeyValue, Panel, type KeyValueItem } from '~kit'
import type { EpochDetail } from '~indexer/selectors'
import { bigIntToHex } from '~lib/format'

const fmt = (value: number | null | undefined): string =>
  value == null || !Number.isFinite(value) ? '—' : value.toLocaleString()

/**
 * The epoch struct as the contract holds it, plus the counters the store keeps.
 * Nothing here is derived — this is the row you compare against `getEpoch`.
 */
export function RawPanel({ detail }: { detail: EpochDetail }) {
  const { epoch } = detail
  const policy = epoch.policy

  const items: KeyValueItem[] = [
    { label: 'threshold', value: fmt(policy?.threshold), mono: true },
    { label: 'committeeSize', value: fmt(policy?.committeeSize), mono: true },
    { label: 'minValidContributions', value: fmt(policy?.minValidContributions), mono: true },
    { label: 'lotteryAlphaBps', value: fmt(policy?.lotteryAlphaBps), mono: true },
    { label: 'committeeSelectionDeadlineBlock', value: fmt(policy?.committeeSelectionDeadlineBlock), mono: true },
    { label: 'keyAssemblyDeadlineBlock', value: fmt(policy?.keyAssemblyDeadlineBlock), mono: true },
    { label: 'liveNotBeforeBlock', value: fmt(policy?.liveNotBeforeBlock), mono: true },
    { label: 'startBlock', value: fmt(epoch.startBlock), mono: true },
    { label: 'seedBlock', value: fmt(epoch.seedBlock), mono: true },
    { label: 'lotteryThreshold', value: bigIntToHex(epoch.lotteryThreshold), mono: true },
    { label: 'registrySnapshot', value: fmt(epoch.registrySnapshot), mono: true },
    { label: 'committeeFilledBlock', value: fmt(epoch.committeeFilledBlock), mono: true },
    { label: 'abortedBlock', value: fmt(epoch.abortedBlock), mono: true },
    { label: 'status', value: epoch.status, mono: true },
    {
      label: 'counts',
      value: `${epoch.counts.claims} claims · ${epoch.counts.contributions} contributions · ${epoch.counts.ciphertexts} ciphertexts · ${epoch.counts.partials} partials · ${epoch.counts.combines} combines`,
      mono: true,
    },
    {
      label: 'stateBlock',
      value: fmt(epoch.stateBlock),
      mono: true,
      hint: 'block at which the struct was last read on chain',
    },
  ]

  return (
    <Panel
      label='Raw'
      title='Epoch record'
      description='The policy immutables and the on-chain struct behind every number on this page.'
    >
      <KeyValue items={items} columns={2} />
    </Panel>
  )
}
