import { Link, useParams } from 'react-router-dom'
import { ButtonLink, Callout, Card, EmptyState, Hash, SectionHeader, Skeleton, SkeletonText, Stack } from '~kit'
import { useEpoch, useIndexer, useNetworkStats, usePartialMatrix } from '~data/hooks'
import { paths } from '~routes/paths'
import { ApplicationsPanel } from './epoch/ApplicationsPanel'
import { CommitteePanel } from './epoch/CommitteePanel'
import { DecryptionMatrix } from './epoch/DecryptionMatrix'
import { EpochHeader } from './epoch/EpochHeader'
import { EventLogPanel } from './epoch/EventLogPanel'
import { PoolPanel } from './epoch/PoolPanel'
import { LotteryPanel } from './epoch/LotteryPanel'
import { RawPanel } from './epoch/RawPanel'

/**
 * Everything the chain knows about one epoch, in lifecycle order: the windows,
 * the lottery that picked the committee, the committee itself, the pool of
 * keys it dealt, the applications it serves and every decryption in flight.
 */
export function EpochPage() {
  const { id = '' } = useParams()
  const indexer = useIndexer()
  const stats = useNetworkStats()
  const detail = useEpoch(id)
  const matrix = usePartialMatrix(id)

  if (!detail) {
    return indexer.scanning || indexer.status.phase === 'loading' ? <EpochSkeleton id={id} /> : <EpochNotFound id={id} />
  }

  return (
    <Stack>
      <EpochHeader
        detail={detail}
        head={stats.headBlock}
        blockTimeSeconds={stats.blockTimeSeconds}
        epochDurationBlocks={stats.epochDurationBlocks}
      />
      <LotteryPanel detail={detail} />
      <CommitteePanel detail={detail} />
      <PoolPanel detail={detail} />
      <ApplicationsPanel applications={detail.applications} />
      <DecryptionMatrix matrix={matrix} applications={detail.applications} />
      <EventLogPanel events={detail.events} />
      <RawPanel detail={detail} />
    </Stack>
  )
}

function EpochSkeleton({ id }: { id: string }) {
  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Epoch'
        title={id ? <Hash value={id} full copy={false} className='text-lg' /> : 'Epoch'}
        description='Reading this epoch out of the log — the indexer is still scanning.'
      />
      <Card className='space-y-4'>
        <Skeleton className='h-6 w-48' />
        <Skeleton className='h-28 w-full' rounded='md' />
        <SkeletonText lines={4} />
      </Card>
      <Card className='space-y-4'>
        <Skeleton className='h-6 w-40' />
        <Skeleton className='h-40 w-full' rounded='md' />
      </Card>
    </Stack>
  )
}

function EpochNotFound({ id }: { id: string }) {
  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Epoch'
        title='Unknown epoch'
        description='No epoch with this id has been created by the configured manager.'
      />
      <Callout tone='warn' title='Nothing indexed under this id'>
        <p className='mt-1'>
          The id below produced no <code className='font-mono'>EpochCreated</code> log in the scanned range. Check the
          bytes12 id, or that the explorer is pointed at the deployment that created it.
        </p>
        <p className='mt-2 font-mono text-[12px] break-all text-ash'>{id || '(no id in the URL)'}</p>
      </Callout>
      <EmptyState
        title='Browse the epochs instead'
        description='Every epoch the manager has created, newest first.'
        action={
          <Link to={paths.epochs()}>
            <ButtonLink href={paths.epochs()} variant='ghost' size='sm'>
              Go to epochs
            </ButtonLink>
          </Link>
        }
      />
    </Stack>
  )
}
