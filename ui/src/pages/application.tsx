import { useMemo } from 'react'
import { Link, useParams } from 'react-router-dom'
import { useApplication, useIndexer, usePartialMatrix, useStore } from '~data/hooks'
import { useRuntimeConfig } from '~config/config-context'
import { useCopy } from '~hooks/use-copy'
import type { EpochPhaseName } from '~indexer/types'
import {
  Address,
  Badge,
  BlockCell,
  Button,
  Callout,
  Hash,
  KeyValue,
  Panel,
  SectionHeader,
  SkeletonText,
  Stack,
  StatCell,
  StatRow,
  TxCell,
  buttonClasses,
  type BadgeTone,
  type KeyValueItem,
} from '~kit'
import { bigIntToHex } from '~lib/format'
import { paths } from '~routes/paths'
import { CiphertextTable } from './application/ciphertexts'
import { ApplicationPartialMatrix } from './application/matrix'
import { applicationPublicKey, formatPointPair } from './application/keys'
import { ReleaseSharePanel } from './application/release-share'

const PHASE_TONE: Record<EpochPhaseName, BadgeTone> = {
  none: 'neutral',
  'committee-selection': 'neutral',
  'key-assembly': 'warn',
  live: 'ok',
  aborted: 'danger',
  completed: 'neutral',
}

/**
 * One application: its record, every ciphertext under it, who answered which,
 * and the organizer's own tools.
 */
export function ApplicationPage() {
  const { epoch = '', aid = '' } = useParams()
  const detail = useApplication(epoch, aid)
  const matrix = usePartialMatrix(epoch, aid)
  const store = useStore()
  const { scanning } = useIndexer()
  const { demo } = useRuntimeConfig()
  const { copied, copy } = useCopy()

  const pkAid = useMemo(
    () => applicationPublicKey(detail?.epoch?.collectivePublicKey, detail?.application.organizerPK),
    [detail]
  )

  if (!detail) {
    if (scanning) return <SkeletonText lines={8} className='max-w-2xl' />
    return (
      <Stack>
        <SectionHeader
          size='page'
          label='Application'
          title='Unknown application'
          description='No application with this id is registered against this epoch in the indexed range.'
        />
        <Callout tone='warn' title='Not found'>
          <p className='font-mono text-[12px] break-all'>
            epoch {epoch || '—'} · aid {aid || '—'}
          </p>
          <p className='mt-2'>
            Applications are keyed by (epoch, aid): the same id under another epoch is a different application.{' '}
            <Link to={paths.applications()} className='text-emerald hover:underline'>
              Browse all applications
            </Link>
            .
          </p>
        </Callout>
      </Stack>
    )
  }

  const { row, application, summary } = detail
  const epochEntity = detail.epoch
  const record: KeyValueItem[] = [
    {
      label: 'Epoch',
      value: (
        <span className='flex items-center justify-end gap-2'>
          <Link
            to={paths.epoch(row.epoch)}
            className='font-mono text-[12px] text-silver transition-colors hover:text-emerald'
          >
            {epochEntity ? `#${epochEntity.nonce}` : row.epoch}
          </Link>
          {epochEntity ? (
            <Badge size='sm' tone={PHASE_TONE[epochEntity.status]} dot={epochEntity.status === 'live'}>
              {epochEntity.status}
            </Badge>
          ) : null}
        </span>
      ),
    },
    { label: 'Application id', value: <Hash value={row.aid} chars={12} /> },
    { label: 'Organizer', value: <Address value={row.creator} to={paths.operator(row.creator)} /> },
    {
      label: 'Registered',
      value: <BlockCell block={row.createdBlock} />,
      hint: row.createdTx ? <TxCell hash={row.createdTx} /> : undefined,
    },
    {
      label: 'Authorized submitter',
      value: row.authorizedSubmitter ? <Address value={row.authorizedSubmitter} /> : '—',
      hint: 'the only address the contract accepts ciphertexts from',
    },
    {
      label: 'Ciphertext cap',
      value: row.maxCiphertexts ? row.maxCiphertexts : '∞',
      mono: true,
      hint: row.maxCiphertexts ? undefined : 'uncapped (0 on chain)',
    },
    {
      label: 'Decryption window',
      value:
        (row.notBeforeBlock ?? 0) === 0 && (row.notAfterBlock ?? 0) === 0
          ? 'open'
          : `${row.notBeforeBlock || 'any'} → ${row.notAfterBlock || 'any'}`,
      mono: true,
      hint: 'blocks between which ciphertexts are accepted',
    },
    { label: 'PK_org x', value: <Hash value={bigIntToHex(application.organizerPK.x)} chars={10} /> },
    { label: 'PK_org y', value: <Hash value={bigIntToHex(application.organizerPK.y)} chars={10} /> },
    {
      label: 'PK_aid x',
      value: pkAid ? <Hash value={bigIntToHex(pkAid.x)} chars={10} /> : '—',
      hint: pkAid ? undefined : 'PK_ep is not on chain until the epoch is live',
    },
    {
      label: 'PK_aid y',
      value: pkAid ? <Hash value={bigIntToHex(pkAid.y)} chars={10} /> : '—',
    },
  ]

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Application'
        title={<Hash value={row.aid} chars={10} className='text-[20px]' />}
        description='Everything registered under this (epoch, application id): the record, every ciphertext and its decryption pipeline, and the organizer tools. The application key is PK_aid = PK_ep + PK_org — the committee alone cannot open it.'
        actions={
          <>
            <Button
              size='sm'
              variant='secondary'
              disabled={!pkAid}
              onClick={() => {
                if (pkAid) copy(formatPointPair(pkAid))
              }}
            >
              {copied ? 'Copied PK_aid' : 'Copy PK_aid'}
            </Button>
            <Link to={paths.playground({ epoch: row.epoch, aid: row.aid })} className={buttonClasses('ghost', 'sm')}>
              Resume in playground
            </Link>
          </>
        }
      />

      <StatRow>
        <StatCell label='Ciphertexts' value={summary.total} mono hint={`cap ${row.maxCiphertexts || '∞'}`} />
        <StatCell
          label='Threshold met'
          value={summary.thresholdMet}
          mono
          hint={
            epochEntity?.policy
              ? `t = ${epochEntity.policy.threshold} of ${epochEntity.policy.committeeSize}`
              : undefined
          }
        />
        <StatCell
          label='Shares released'
          value={summary.withShare}
          mono
          tone={summary.withShare < summary.total ? 'warn' : 'default'}
          hint={`${Math.max(0, summary.total - summary.withShare)} still withheld`}
        />
        <StatCell
          label='Decrypted'
          value={summary.combined}
          mono
          tone='accent'
          hint={
            summary.total > 0
              ? `${Math.round((summary.combined / summary.total) * 100)}% of this application`
              : undefined
          }
        />
      </StatRow>

      <Panel label='Record' title='On-chain application record'>
        <KeyValue items={record} columns={2} />
      </Panel>

      <Panel
        label='Pipeline'
        title='Ciphertexts'
        description='Submitted → partials → organizer share → combined, with the block and transaction behind every step.'
        bodyClassName='p-0'
      >
        <CiphertextTable rows={detail.ciphertexts} />
      </Panel>

      <Panel
        label='Decryption'
        title='Partials by committee member'
        description={`Rows are the ${matrix?.rows.length ?? 0} committee members in slot order, columns are this application's ciphertexts. Colour is the decryption wave (stagger ${matrix?.staggerBlocks ?? store.chain.staggerBlocks} blocks).`}
      >
        <ApplicationPartialMatrix matrix={matrix} />
      </Panel>

      <Panel
        label='Organizer'
        title='Release organizer share'
        description='Δ = sk_org·C1 with a Chaum–Pedersen DLEQ against PK_org. Computed in this tab; the secret is never transmitted.'
      >
        <ReleaseSharePanel
          key={`${row.epoch}:${row.aid}`}
          epoch={row.epoch}
          aid={row.aid}
          ciphertexts={detail.ciphertexts}
          managerAddress={store.chain.managerAddress}
          appManagerAddress={store.chain.appManagerAddress}
          demo={demo}
        />
      </Panel>
    </Stack>
  )
}
