import { useMemo } from 'react'
import { ButtonLink, Panel, SectionHeader, Stack } from '~kit'
import { useActivity, useEpochs, useEventFeed, useIndexer, useNetworkStats } from '~data/hooks'
import { paths } from '~routes/paths'
import { EventTable } from './epochs/EventTable'
import { ActivityPanel } from './overview/ActivityPanel'
import { CadencePanel } from './overview/CadencePanel'
import { HeaderStrip, StatusCards } from './overview/Summary'
import { IndexerStatus } from './overview/IndexerStatus'

/**
 * The network at a glance: which deployment, what block, what the newest epoch
 * is waiting for, what the committee has been doing, and the last twenty
 * things that happened on chain. Everything is a link to the page that has the
 * full record.
 */
export function OverviewPage() {
  const indexer = useIndexer()
  const stats = useNetworkStats()
  const activity = useActivity(30)
  const epochs = useEpochs({ limit: 30 })
  const feed = useEventFeed(20)

  // Nothing has landed yet *and* the indexer is still working: that is the only
  // state where a skeleton is honest. An empty deployment must show "empty".
  const loading = indexer.scanning || indexer.status.phase === 'loading'
  const events = useMemo(() => feed.map((entry) => entry.event), [feed])

  return (
    <Stack>
      <SectionHeader
        size='page'
        label='Network'
        title='Overview'
        description='Non-interactive DKG on chain: every epoch, every committee and every decryption, reconstructed from the logs of the three contracts.'
        actions={
          <ButtonLink href={paths.epochs()} variant='ghost' size='sm'>
            Browse epochs
          </ButtonLink>
        }
      />

      <IndexerStatus indexer={indexer} />
      <HeaderStrip stats={stats} loading={loading} />
      <StatusCards stats={stats} loading={loading && stats.epochs === 0} />

      <ActivityPanel buckets={activity} loading={loading && activity.length === 0} />
      <CadencePanel rows={epochs} head={stats.headBlock} loading={loading && epochs.length === 0} />

      <Panel
        label='Log'
        title='Latest events'
        description='The twenty newest logs across the registry, the manager and the application manager.'
        bodyClassName='p-0'
        actions={
          <span className='font-mono text-[11px] text-ash'>{stats.events.toLocaleString()} indexed</span>
        }
      >
        <EventTable
          events={events}
          showEpoch
          loading={loading && events.length === 0}
          maxHeight={560}
          emptyTitle='No events indexed yet'
          emptyDescription='Nothing has been logged by this deployment in the scanned range.'
        />
      </Panel>
    </Stack>
  )
}
