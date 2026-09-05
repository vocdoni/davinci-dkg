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
import { RevealSecretPanel } from './application/reveal-secret'
import { DecryptionWindowCell, ModeBadge } from './applications/cells'
import { MODE_HELP, WINDOW_LABEL, describeDecryptionWindow, submissionPolicy } from './applications/policy'

const PHASE_TONE: Record<EpochPhaseName, BadgeTone> = {
  none: 'neutral',
  'committee-selection': 'neutral',
  'key-assembly': 'warn',
  live: 'ok',
  aborted: 'danger',
  completed: 'neutral',
}

/**
 * One application: its record, the pool key it encrypts under, every
 * ciphertext under it, who answered which, and the organizer's one tool —
 * the reveal.
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
    () =>
      detail
        ? applicationPublicKey(
            detail.row.poolKey,
            detail.application.mode === 'automatic' ? null : detail.application.organizerPK
          )
        : null,
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
  const automatic = application.mode === 'automatic'
  const revealed = !automatic && application.organizerSecret != null
  const submission = submissionPolicy(row)
  const window = describeDecryptionWindow(row.decryptNotBefore, row.decryptNotAfter)
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
    {
      label: 'Mode',
      value: <ModeBadge mode={application.mode} size='md' />,
      hint: MODE_HELP[application.mode],
    },
    {
      label: 'Pool key',
      value: row.poolIndex != null ? `key ${row.poolIndex}` : '—',
      mono: true,
      hint:
        row.poolIndex != null
          ? `P_${row.poolIndex}, claimed at registration out of the epoch’s pool`
          : 'claim not indexed yet',
    },
    { label: 'Organizer', value: <Address value={row.creator} to={paths.operator(row.creator)} /> },
    {
      label: 'Registered',
      value: <BlockCell block={row.createdBlock} />,
      hint: row.createdTx ? <TxCell hash={row.createdTx} /> : undefined,
    },
    {
      label: 'Submission policy',
      value:
        submission.kind === 'unknown' ? (
          '—'
        ) : submission.kind === 'open' ? (
          <Badge size='sm' tone='warn'>
            open
          </Badge>
        ) : submission.kind === 'registrant' ? (
          'registrant only'
        ) : (
          <span className='flex flex-col items-end gap-1'>
            {submission.submitters.map((submitter) => (
              <Address key={submitter} value={submitter} />
            ))}
          </span>
        ),
      hint:
        submission.kind === 'open'
          ? 'anyone may call submitCiphertext'
          : submission.kind === 'allow-list'
            ? `allow-list of ${submission.submitters.length} — the only addresses the contract accepts ciphertexts from`
            : 'only the registering address may submit ciphertexts',
    },
    {
      label: 'Ciphertext cap',
      value: row.maxCiphertexts ? row.maxCiphertexts : '∞',
      mono: true,
      hint: row.maxCiphertexts ? undefined : 'uncapped (0 on chain)',
    },
    {
      label: 'Submission window',
      value:
        (row.notBeforeBlock ?? 0) === 0 && (row.notAfterBlock ?? 0) === 0
          ? 'any block'
          : `${row.notBeforeBlock || 'any'} → ${row.notAfterBlock || 'any'}`,
      mono: true,
      hint: 'blocks between which ciphertexts are accepted',
    },
    {
      label: 'Decryption window',
      value: <DecryptionWindowCell notBefore={row.decryptNotBefore} notAfter={row.decryptNotAfter} />,
      hint:
        window == null
          ? undefined
          : window.state === 'unbounded'
            ? automatic
              ? 'no window (0, 0 on chain): partials and combines are accepted at any time'
              : 'no window (0, 0 on chain): partials and combines are accepted at any time once sk_org is revealed'
            : `decryptNotBefore → decryptNotAfter, unix seconds on chain · ${WINDOW_LABEL[window.state]} — outside it submitPartialDecryption and combineDecryption revert${automatic ? '' : ', and before the reveal too'}`,
    },
    {
      label: 'P_j x',
      value: row.poolKey ? <Hash value={bigIntToHex(row.poolKey.x)} chars={10} /> : '—',
      hint: row.poolKey ? undefined : 'the epoch’s pool keys have not been read yet',
    },
    {
      label: 'P_j y',
      value: row.poolKey ? <Hash value={bigIntToHex(row.poolKey.y)} chars={10} /> : '—',
    },
    // The record is a two-column grid, so every x/y pair must start on an
    // even item: the coordinates come first and the reveal state closes the
    // list rather than splitting PK_aid across two rows.
    ...(automatic
      ? []
      : [
          { label: 'PK_org x', value: <Hash value={bigIntToHex(application.organizerPK.x)} chars={10} /> },
          { label: 'PK_org y', value: <Hash value={bigIntToHex(application.organizerPK.y)} chars={10} /> },
        ]),
    {
      label: 'PK_aid x',
      value: pkAid ? <Hash value={bigIntToHex(pkAid.x)} chars={10} /> : '—',
      hint: pkAid ? (automatic ? 'PK_aid = P_j' : 'PK_aid = P_j + PK_org') : 'needs the pool key',
    },
    {
      label: 'PK_aid y',
      value: pkAid ? <Hash value={bigIntToHex(pkAid.y)} chars={10} /> : '—',
    },
    ...(automatic
      ? []
      : [
          {
            label: 'Organizer secret',
            value: revealed ? (
              <span className='flex items-center justify-end gap-2'>
                <Badge size='sm' tone='ok'>
                  revealed
                </Badge>
                {row.revealBlock != null ? <BlockCell block={row.revealBlock} /> : null}
                {row.revealTx ? <TxCell hash={row.revealTx} chars={4} /> : null}
              </span>
            ) : (
              <Badge size='sm' tone='warn'>
                kept
              </Badge>
            ),
            hint: revealed
              ? 'sk_org is on chain: the committee combines on its own'
              : 'sk_org is not on chain: the contract refuses every partial and combine until the organizer reveals it',
          },
        ]),
  ]

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Application'
        title={<Hash value={row.aid} chars={10} className='text-[20px]' />}
        description={
          automatic
            ? 'Everything registered under this (epoch, application id): the record, every ciphertext and its decryption pipeline. The application key is its pool key alone, PK_aid = P_j — there is no organizer key, so t partials inside the decryption window are all a plaintext takes.'
            : 'Everything registered under this (epoch, application id): the record, every ciphertext and its decryption pipeline, and the organizer’s reveal. The application key is PK_aid = P_j + PK_org, and the contract refuses every partial and combine (OrganizerSecretNotRevealed) until sk_org is on chain — its ciphertexts sit at awaiting reveal.'
        }
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
          label='Awaiting reveal'
          value={summary.awaitingReveal}
          mono
          tone={summary.awaitingReveal > 0 ? 'warn' : 'default'}
          hint={
            automatic
              ? 'n/a: no organizer key'
              : revealed
                ? 'the secret is on chain'
                : 'partials and combines refused until the reveal'
          }
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
        description={
          automatic
            ? 'Submitted → partials → combined, with the block and transaction behind every step. Each partial carries a Merkle proof of the member’s share of this application’s pool key.'
            : 'Submitted → awaiting reveal → partials → combined, with the block and transaction behind every step. Until sk_org is on chain the contract refuses every partial and combine of this application; each partial then carries a Merkle proof of the member’s share of its pool key.'
        }
        bodyClassName='p-0'
      >
        <CiphertextTable rows={detail.ciphertexts} />
      </Panel>

      <Panel
        label='Decryption'
        title='Partials by committee member'
        description={`Rows are the ${matrix?.rows.length ?? 0} committee members in slot order, columns are this application's ciphertexts. Colour is the decryption wave, counted from 0 in windows of ${matrix?.staggerBlocks ?? store.chain.staggerBlocks} blocks from the moment the ciphertext became decryptable${automatic ? '' : ' (the reveal, for the ciphertexts that were waiting on it)'}.`}
      >
        <ApplicationPartialMatrix matrix={matrix} />
      </Panel>

      {automatic ? (
        <Panel
          label='Organizer'
          title='No organizer step'
          description='This application is automatic: it has no organizer key, so there is nothing to reveal.'
        >
          <Callout tone='info' title='The committee decrypts on its own'>
            <p>
              PK_org is the identity (0, 1) and the combine proof runs with a zero organizer secret. Once t partials are
              in — and the decryption window is open — a committee node combines and the plaintext lands on chain.
              Confidentiality of this application rests on the committee threshold alone.
            </p>
          </Callout>
        </Panel>
      ) : revealed ? (
        <Panel
          label='Organizer'
          title='Organizer secret revealed'
          description='revealOrganizerSecret ran once; from then on the committee combines every ciphertext of this application by itself.'
        >
          <Callout tone='ok' title='Nothing left for the organizer to do'>
            <p className='flex flex-wrap items-center gap-2'>
              Revealed at
              {row.revealBlock != null ? <BlockCell block={row.revealBlock} /> : <span>an unindexed block</span>}
              {row.revealTx ? <TxCell hash={row.revealTx} /> : null}
            </p>
            <p className='mt-2 font-mono text-[12px] break-all text-silver'>
              sk_org = {application.organizerSecret?.toString()}
            </p>
          </Callout>
        </Panel>
      ) : (
        <Panel
          label='Organizer'
          title='Reveal organizer secret'
          description='One transaction, once: the contract checks sk_org·G = PK_org and stores the scalar. Until then it refuses every partial decryption and combine of this application (OrganizerSecretNotRevealed); from that block on every ciphertext — past and future — is decryptable by t committee members inside the decryption window.'
        >
          <RevealSecretPanel
            key={`${row.epoch}:${row.aid}`}
            epoch={row.epoch}
            aid={row.aid}
            organizerPK={application.organizerPK}
            ciphertexts={summary.total}
            managerAddress={store.chain.managerAddress}
            appManagerAddress={store.chain.appManagerAddress}
            demo={demo}
          />
        </Panel>
      )}
    </Stack>
  )
}
