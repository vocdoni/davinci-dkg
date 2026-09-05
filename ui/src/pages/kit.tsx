import { useMemo, useState } from 'react'
import {
  Address,
  Badge,
  BlockCell,
  Button,
  ButtonLink,
  Callout,
  Card,
  CardBody,
  CardHeader,
  CopyButton,
  DataTable,
  Dialog,
  EmptyState,
  ExternalIcon,
  Hash,
  Input,
  KeyValue,
  Pagination,
  Panel,
  Popover,
  ProgressBar,
  SectionHeader,
  Select,
  Skeleton,
  SkeletonText,
  Stack,
  Stat,
  StatCell,
  StatRow,
  Tabs,
  Timeline,
  TimelineRow,
  Toggle,
  Tooltip,
  TxCell,
  type AnyColumnDef,
} from '~kit'
import { BlockTimeline, CadenceStrip, Donut, Gauge, Matrix, Sparkline, StackedBars, waveColor } from '~kit/charts'
import { CHART_COLORS } from '~kit/charts'
import {
  demoActivity,
  demoCadence,
  demoMatrix,
  demoOperators,
  demoSparkline,
  demoWindows,
  type DemoOperator,
} from './kit-data'

const SAMPLE_ADDRESS = '0x3f9b338706a31f26d49159478015c8aaeab908ad'
const SAMPLE_EPOCH = '0x0102030405060708090a0b0c'
const SAMPLE_TX = `0x${'9f'.repeat(32)}`
const SAMPLE_AID = `0x${'2c'.repeat(32)}`

/**
 * Every primitive and every chart on one page, with inline sample data. This
 * is the design review surface: if something looks wrong here, it looks wrong
 * everywhere, because the pages are built only from these parts.
 */
export function KitPage() {
  const operators = useMemo(() => demoOperators(300), [])
  const activity = useMemo(() => demoActivity(30), [])
  const matrix = useMemo(() => demoMatrix(64, 16, 4), [])
  const cadence = useMemo(() => demoCadence(10), [])

  const [dialogOpen, setDialogOpen] = useState(false)
  const [toggle, setToggle] = useState(true)
  const [page, setPage] = useState(2)

  const columns = useMemo<AnyColumnDef<DemoOperator>[]>(
    () => [
      {
        id: 'address',
        header: 'Operator',
        accessorKey: 'address',
        meta: { width: '190px' },
        cell: ({ row }) => <Address value={row.original.address} copy={false} />,
      },
      {
        id: 'status',
        header: 'Status',
        accessorKey: 'status',
        meta: { width: '90px' },
        cell: ({ row }) => (
          <Badge
            size='sm'
            dot
            tone={row.original.status === 'active' ? 'ok' : row.original.status === 'idle' ? 'warn' : 'danger'}
          >
            {row.original.status}
          </Badge>
        ),
      },
      {
        id: 'registeredAt',
        header: 'Registered',
        accessorKey: 'registeredAt',
        meta: { numeric: true, width: '120px' },
        cell: ({ row }) => <BlockCell block={row.original.registeredAt} />,
      },
      { id: 'epochs', header: 'Epochs', accessorKey: 'epochs', meta: { numeric: true, width: '80px' } },
      { id: 'claims', header: 'Claims', accessorKey: 'claims', meta: { numeric: true, width: '80px' } },
      {
        id: 'contributions',
        header: 'Contributions',
        accessorKey: 'contributions',
        meta: { numeric: true, width: '110px', headerTooltip: 'Accepted contributions to a committee transcript' },
      },
      {
        id: 'participation',
        header: 'Participation',
        accessorFn: (row) => (row.claims === 0 ? -1 : row.contributions / row.claims),
        meta: { numeric: true, width: '130px' },
        cell: ({ row }) =>
          row.original.claims === 0 ? (
            <span className='text-ash'>—</span>
          ) : (
            <ProgressBar
              value={row.original.contributions}
              total={row.original.claims}
              label={false}
              size='sm'
              className='w-24'
            />
          ),
      },
      {
        id: 'trend',
        header: 'Trend',
        meta: { width: '110px' },
        cell: ({ row }) => <Sparkline values={row.original.trend} width={90} height={20} />,
      },
    ],
    []
  )

  return (
    <Stack className='gap-10'>
      <SectionHeader
        size='page'
        label='Design system'
        title='Kit'
        description='Every primitive and chart the explorer is built from, rendered with sample data. Tokens come from ui/design — obsidian canvas, one emerald accent, charcoal hairlines, Inter and JetBrains Mono.'
        actions={
          <ButtonLink href='https://github.com/vocdoni/davinci-dkg' external variant='ghost' size='sm'>
            Source <ExternalIcon size={12} />
          </ButtonLink>
        }
      />

      <Section
        title='Buttons'
        note='Primary is the single most important action on a screen. Everything else is secondary or subtle.'
      >
        <div className='flex flex-wrap items-center gap-3'>
          <Button variant='primary'>Primary</Button>
          <Button variant='ghost'>Ghost</Button>
          <Button variant='secondary'>Secondary</Button>
          <Button variant='subtle'>Subtle</Button>
          <Button variant='danger'>Danger</Button>
          <Button variant='primary' loading>
            Submitting
          </Button>
          <Button variant='secondary' disabled>
            Disabled
          </Button>
        </div>
        <div className='mt-4 flex flex-wrap items-center gap-3'>
          <Button size='sm' variant='primary'>
            Small
          </Button>
          <Button size='md' variant='primary'>
            Medium
          </Button>
          <Button size='lg' variant='primary'>
            Large
          </Button>
          <Button size='icon' variant='secondary' aria-label='External'>
            <ExternalIcon />
          </Button>
        </div>
      </Section>

      <Section title='Badges' note='Emerald for live/ok, amber for warnings, red for failure, charcoal for neutral.'>
        <div className='flex flex-wrap items-center gap-3'>
          <Badge tone='ok' dot>
            live
          </Badge>
          <Badge tone='accent'>selected</Badge>
          <Badge tone='warn' dot>
            assembling
          </Badge>
          <Badge tone='danger' dot>
            aborted
          </Badge>
          <Badge tone='neutral'>closed</Badge>
          <Badge tone='ok' size='sm'>
            sm
          </Badge>
        </div>
      </Section>

      <Section
        title='Surfaces'
        note='Level 0 obsidian page, level 1 carbon cards, level 2 onyx hover, charcoal hairlines. Borders over shadows.'
      >
        <div className='grid gap-4 md:grid-cols-3'>
          <Card hover>
            <div className='text-[15px] font-semibold text-ghost'>Carbon card</div>
            <p className='mt-1.5 text-[13px] leading-relaxed text-ash'>
              Level 1 surface with a charcoal hairline. Hover lifts the border to warm gray and the surface to onyx.
            </p>
          </Card>
          <Card level='onyx'>
            <div className='text-[15px] font-semibold text-ghost'>Onyx card</div>
            <p className='mt-1.5 text-[13px] leading-relaxed text-ash'>
              Level 2 surface, for nested panels and hover states.
            </p>
          </Card>
          <Card flush>
            <CardHeader label='Panel' title='With a header' actions={<Badge tone='ok'>ok</Badge>} />
            <CardBody className='text-[13px] text-ash'>
              Header, hairline, body. The shape of every explorer panel.
            </CardBody>
          </Card>
        </div>
        <div className='mt-4 grid gap-4 md:grid-cols-2'>
          <Callout tone='info' title='Info'>
            The indexer is still scanning: 62% of the history is loaded.
          </Callout>
          <Callout tone='warn' title='Secret not revealed'>
            The organizer has not revealed its secret; these ciphertexts cannot be combined.
          </Callout>
          <Callout tone='ok' title='Finalized'>
            The collective key is on chain and the epoch is live.
          </Callout>
          <Callout tone='danger' title='Aborted'>
            No committee reached the minimum size before the window closed.
          </Callout>
        </div>
      </Section>

      <Section title='Stats'>
        <StatRow>
          <StatCell label='Live epochs' value='3' hint='of 148 created' tone='accent' />
          <StatCell label='Operators active' value='271' hint='of 300 registered' />
          <StatCell label='Committee' value='64' hint='threshold t = 33' mono />
          <StatCell label='Decrypted' value='1 284' hint='all time' mono />
        </StatRow>
        <div className='mt-4 grid gap-4 md:grid-cols-3'>
          <Card>
            <Stat label='Next epoch' value='412' hint='blocks remaining' mono />
          </Card>
          <Card>
            <Stat label='Contributions' value='31 / 64' hint='t = 33 not reached' tone='warn' mono />
          </Card>
          <Card>
            <Stat label='Loading' value='—' loading />
          </Card>
        </div>
      </Section>

      <Section title='Inputs'>
        <div className='grid max-w-3xl gap-4 md:grid-cols-2'>
          <Input label='Application id' placeholder='0x…' mono hint='32 bytes, below the BN254 scalar field.' />
          <Input label='Ciphertext cap' type='number' defaultValue={8} hint='Maximum ciphertexts under this aid.' />
          <Input label='Organizer secret' placeholder='0x…' mono error='Not a valid scalar' />
          <Select
            label='Phase'
            options={[
              { value: 'all', label: 'All phases' },
              { value: 'live', label: 'Live' },
              { value: 'assembly', label: 'Key assembly' },
              { value: 'closed', label: 'Closed' },
            ]}
          />
          <Toggle
            checked={toggle}
            onChange={setToggle}
            label='Advanced transcripts'
            hint='Print the PoP and reveal transcripts.'
          />
        </div>
      </Section>

      <Section
        title='Chain values'
        note='Addresses are checksummed and truncated in the middle; the full value is in the tooltip, the title and the clipboard.'
      >
        <div className='flex flex-wrap items-center gap-x-8 gap-y-3'>
          <Labelled label='Address'>
            <Address value={SAMPLE_ADDRESS} />
          </Labelled>
          <Labelled label='Address (full)'>
            <Address value={SAMPLE_ADDRESS} full explorer={false} />
          </Labelled>
          <Labelled label='Epoch id'>
            <Hash value={SAMPLE_EPOCH} />
          </Labelled>
          <Labelled label='Application id'>
            <Hash value={SAMPLE_AID} chars={8} />
          </Labelled>
          <Labelled label='Transaction'>
            <TxCell hash={SAMPLE_TX} copy />
          </Labelled>
          <Labelled label='Block'>
            <BlockCell block={11_903_412n} suffix='· 2 min ago' />
          </Labelled>
          <Labelled label='Copy'>
            <CopyButton value={SAMPLE_ADDRESS} />
          </Labelled>
        </div>
      </Section>

      <Section title='Progress, key-values and timeline'>
        <div className='grid gap-6 md:grid-cols-2'>
          <div className='space-y-4'>
            <ProgressBar value={48} total={64} threshold={33} label='claims' />
            <ProgressBar value={22} total={64} threshold={33} label='contributions' tone='warn' />
            <ProgressBar value={64} total={64} threshold={33} label='partials' />
            <ProgressBar value={4} total={8} label='ciphertexts combined' tone='neutral' />
            <Pagination page={page} pageCount={12} pageSize={25} total={287} onPageChange={setPage} />
          </div>
          <KeyValue
            items={[
              { label: 'epoch id', value: <Hash value={SAMPLE_EPOCH} copy={false} />, mono: true },
              { label: 'nonce', value: '148', mono: true },
              { label: 'threshold', value: '33 of 64', mono: true },
              { label: 'seed block', value: <BlockCell block={11_898_000} /> },
              { label: 'creator', value: <Address value={SAMPLE_ADDRESS} copy={false} /> },
              { label: 'α', value: '0.65535', mono: true, hint: 'lottery slack' },
            ]}
          />
        </div>
        <div className='mt-6 max-w-2xl'>
          <Timeline>
            <TimelineRow tone='ok' title='SlotClaimed — slot 0' meta='11 898 041' right={<TxCell hash={SAMPLE_TX} />} />
            <TimelineRow
              tone='ok'
              title='ContributionSubmitted'
              meta='11 898 210'
              right={<TxCell hash={SAMPLE_TX} />}
            />
            <TimelineRow tone='warn' title='OrganizerSecretRevealed' meta='11 901 004' description='aid 0x2c2c…2c2c' />
            <TimelineRow tone='muted' last title='CiphertextCombined (pending)' meta='—' />
          </Timeline>
        </div>
      </Section>

      <Section title='Overlays'>
        <div className='flex flex-wrap items-center gap-3'>
          <Tooltip content='Radix tooltip, 11 px, onyx surface'>
            <Button variant='secondary'>Hover me</Button>
          </Tooltip>
          <Button variant='ghost' onClick={() => setDialogOpen(true)}>
            Open dialog
          </Button>
          <Popover trigger={<Button variant='secondary'>Popover</Button>}>
            <div className='space-y-1 text-[13px] text-silver'>
              <div className='label-caps px-2 py-1 text-[10px] text-pewter'>Filter</div>
              {['All phases', 'Live', 'Key assembly', 'Closed'].map((label) => (
                <button key={label} className='block w-full rounded-sm px-2 py-1.5 text-left hover:bg-onyx'>
                  {label}
                </button>
              ))}
            </div>
          </Popover>
          <Dialog
            open={dialogOpen}
            onOpenChange={setDialogOpen}
            title='Reveal organizer secret'
            description='The secret is checked against PK_org in this tab before it is sent; afterwards the committee decrypts alone.'
            footer={
              <>
                <Button variant='secondary' onClick={() => setDialogOpen(false)}>
                  Cancel
                </Button>
                <Button variant='primary' onClick={() => setDialogOpen(false)}>
                  Reveal
                </Button>
              </>
            }
          >
            <div className='space-y-4'>
              <Input label='Organizer secret' placeholder='0x…' mono />
              <Callout tone='warn'>Revealing the secret makes every ciphertext of this application decryptable by the committee.</Callout>
            </div>
          </Dialog>
        </div>
        <div className='mt-6'>
          <Tabs
            items={[
              {
                value: 'committee',
                label: 'Committee',
                meta: '64',
                content: <p className='text-[13px] text-ash'>Committee panel.</p>,
              },
              {
                value: 'apps',
                label: 'Applications',
                meta: '3',
                content: <p className='text-[13px] text-ash'>Applications table.</p>,
              },
              {
                value: 'events',
                label: 'Event log',
                meta: '412',
                content: <p className='text-[13px] text-ash'>Every event of the epoch.</p>,
              },
            ]}
          />
        </div>
      </Section>

      <Section title='Loading and empty states'>
        <div className='grid gap-4 md:grid-cols-3'>
          <Card>
            <SkeletonText lines={4} />
          </Card>
          <Card className='space-y-3'>
            <Skeleton className='h-6 w-32' />
            <Skeleton className='h-24 w-full' rounded='md' />
          </Card>
          <Card flush>
            <EmptyState
              compact
              title='No applications yet'
              description='Register one from the playground to see it here.'
              action={
                <Button size='sm' variant='ghost'>
                  Open playground
                </Button>
              }
            />
          </Card>
        </div>
      </Section>

      <Panel
        label='Table'
        title='Operators — 300 rows, virtualised'
        description='Sticky header, click-to-sort, mono numeric columns, windowed body. Sorting and scrolling stay instant at registry scale.'
        bodyClassName='p-0'
      >
        <DataTable
          data={operators}
          columns={columns}
          virtualized
          rowHeight={44}
          maxHeight={420}
          getRowId={(row) => row.address}
          initialSorting={[{ id: 'contributions', desc: true }]}
        />
      </Panel>

      <div className='grid gap-6 lg:grid-cols-2'>
        <Panel
          label='Table'
          title='Compact, unvirtualised'
          description='The same component under ~50 rows.'
          bodyClassName='p-0'
        >
          <DataTable data={operators.slice(0, 8)} columns={columns.slice(0, 5)} maxHeight={360} />
        </Panel>
        <Panel label='Table' title='Loading and empty' bodyClassName='p-0'>
          <DataTable data={[]} columns={columns.slice(0, 4)} loading loadingRows={4} />
          <DataTable data={[]} columns={columns.slice(0, 4)} />
        </Panel>
      </div>

      <SectionHeader
        label='Charts'
        title='Hand-rolled SVG'
        description='No chart library. Every chart has a legend, a skeleton and a tooltip, and its maths is unit-tested.'
      />

      <div className='grid gap-6 lg:grid-cols-2'>
        <Panel label='Overview' title='Activity per epoch' description='Stacked bars over the last 30 epochs.'>
          <StackedBars
            data={activity}
            series={[
              { key: 'claims', label: 'claims' },
              { key: 'contributions', label: 'contributions' },
              { key: 'ciphertexts', label: 'ciphertexts' },
              { key: 'partials', label: 'partials' },
            ]}
            height={220}
          />
        </Panel>

        <Panel label='Epoch' title='Lifecycle on the block axis' description='The four windows and the current block.'>
          <BlockTimeline windows={demoWindows} current={11_902_400} height={110} />
        </Panel>

        <Panel
          label='Overview'
          title='Epoch cadence'
          description='Epochs staggered on a shared block axis, coloured by phase.'
        >
          <CadenceStrip epochs={cadence} current={11_902_400} height={110} />
        </Panel>

        <Panel label='Operators' title='Status split'>
          <div className='flex flex-wrap items-center gap-8'>
            <Donut
              slices={[
                { label: 'active', value: 271 },
                { label: 'idle', value: 21 },
                { label: 'reaped', value: 8, color: CHART_COLORS.red },
              ]}
              centerValue={300}
              centerLabel='registered'
              size={168}
              className='max-w-[260px]'
            />
            <Gauge
              value={0.000065535}
              label='τ / 2²⁵⁶'
              caption='α = 6.5535 · N = 300'
              reference={0.5}
              className='max-w-[240px]'
            />
          </div>
        </Panel>
      </div>

      <Panel
        label='Epoch'
        title='Decryption matrix — 64 members × 16 ciphertexts'
        description='Cells are coloured by the wave a partial arrived in: emerald answered first, slate trailed. Row labels stay put while the grid scrolls inside the panel.'
      >
        <Matrix
          rows={matrix.rows}
          columns={matrix.columns}
          cells={matrix.cells.map((cell) => ({
            row: cell.row,
            col: cell.col,
            color: waveColor(cell.wave, matrix.waves),
            detail: (
              <div className='font-mono text-[10px]'>
                <div className='text-ghost'>{matrix.rows[cell.row]}</div>
                <div className='text-ash'>
                  ciphertext {cell.col} · wave {cell.wave} · block {cell.block.toLocaleString()}
                </div>
              </div>
            ),
          }))}
          cellSize={12}
          legend={Array.from({ length: matrix.waves }, (_, w) => ({
            label: `wave ${w}`,
            color: waveColor(w, matrix.waves),
          })).concat([{ label: 'no partial', color: CHART_COLORS.onyx }])}
        />
      </Panel>

      <div className='grid gap-6 lg:grid-cols-2'>
        <Panel label='Operator' title='Sparklines' description='Fixed-size, unmeasured — they live in table cells.'>
          <div className='flex flex-wrap items-center gap-6'>
            <Sparkline values={demoSparkline} />
            <Sparkline values={demoSparkline} width={140} height={36} color={CHART_COLORS.teal} />
            <Sparkline values={[5, 5, 5, 5]} area={false} />
            <Sparkline values={[]} />
          </div>
        </Panel>
        <Panel label='Charts' title='Skeletons' description='What every chart shows before its first measurement.'>
          <div className='space-y-4'>
            <StackedBars data={activity} series={[{ key: 'claims', label: 'claims' }]} loading height={120} />
            <BlockTimeline windows={demoWindows} loading height={90} />
          </div>
        </Panel>
      </div>
    </Stack>
  )
}

function Section({ title, note, children }: { title: string; note?: string; children: React.ReactNode }) {
  return (
    <section>
      <SectionHeader label='Kit' title={title} description={note} className='mb-5' />
      {children}
    </section>
  )
}

function Labelled({ label, children }: { label: string; children: React.ReactNode }) {
  return (
    <div>
      <div className='label-caps mb-1.5 text-[10px] text-pewter'>{label}</div>
      {children}
    </div>
  )
}
