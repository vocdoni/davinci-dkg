import { Address, BlockCell, Callout, Hash, KeyValue, Panel, TxCell, type KeyValueItem } from '~kit'
import type { EpochDetail } from '~indexer/selectors'
import { bigIntToHex } from '~lib/format'
import { paths } from '~routes/paths'
import { blockOrNull, transcriptWords } from '~pages/epochs/cadence'

/**
 * PK_ep and the transaction that produced it. The collective key is the whole
 * point of the epoch: `PK_aid = PK_ep + PK_org`, so every application key on
 * this epoch is derived from the coordinates below.
 */
export function KeyPanel({ detail }: { detail: EpochDetail }) {
  const { finalization, collectivePublicKey, row } = detail
  const words = transcriptWords(row.committeeSize)

  if (!finalization) {
    return (
      <Panel label='Key' title='Collective public key' description='Produced by finalizeEpoch, with a Groth16 proof.'>
        <Callout tone='info' title='No collective key yet'>
          The epoch has not been finalized. finalizeEpoch may run once the key-assembly window has closed, the finalize
          gap has passed, and at least m_min = {row.minValidContributions} contributions are on chain.
        </Callout>
      </Panel>
    )
  }

  const items: KeyValueItem[] = [
    {
      label: 'PK_ep.x',
      value: collectivePublicKey ? (
        <Hash value={bigIntToHex(collectivePublicKey.x)} chars={10} />
      ) : (
        <span className='text-ash'>not read yet</span>
      ),
      mono: true,
    },
    {
      label: 'PK_ep.y',
      value: collectivePublicKey ? (
        <Hash value={bigIntToHex(collectivePublicKey.y)} chars={10} />
      ) : (
        <span className='text-ash'>not read yet</span>
      ),
      mono: true,
      hint: 'twisted-Edwards form, as returned by getCollectivePublicKey',
    },
    { label: 'PK_ep hash', value: <Hash value={finalization.collectivePublicKeyHash} chars={10} />, mono: true },
    {
      label: 'aggregate commitments',
      value: <Hash value={finalization.aggregateCommitmentsHash} chars={10} />,
      mono: true,
    },
    { label: 'share commitments', value: <Hash value={finalization.shareCommitmentHash} chars={10} />, mono: true },
    {
      label: 'finalizer',
      value: finalization.by ? (
        <Address value={finalization.by} chars={6} to={paths.operator(finalization.by)} />
      ) : (
        <span className='text-ash'>resolving…</span>
      ),
      hint: 'EpochLive names nobody — this is the transaction sender',
    },
    {
      label: 'finalized',
      value: (
        <span className='inline-flex items-center gap-2'>
          <BlockCell block={blockOrNull(finalization.block)} />
          {finalization.tx ? <TxCell hash={finalization.tx} /> : null}
        </span>
      ),
    },
    {
      label: 'gas',
      value: finalization.gasUsed != null ? finalization.gasUsed.toLocaleString() : '—',
      mono: true,
      hint: 'one Groth16 verification over the whole transcript',
    },
    {
      label: 'transcript',
      value: `${words.toLocaleString()} words`,
      mono: true,
      hint: `2·n² + 5·n at n = ${row.committeeSize} — ${(words * 32).toLocaleString()} bytes of calldata`,
    },
  ]

  return (
    <Panel
      label='Key'
      title='Collective public key'
      description='finalizeEpoch proved the aggregate against every accepted contribution and flipped the epoch to Live.'
    >
      <KeyValue items={items} columns={2} />
    </Panel>
  )
}
