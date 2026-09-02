import { useState } from 'react'
import { Box, HStack, Stack, Text } from '@chakra-ui/react'
import { useMeasuredWidth } from '~hooks/use-measured-width'
import { ChartAxisInk, ChartGrid, ChartLabelInk, type ChartSeriesSpec } from './palette'

export interface BarGroup {
  /** Stable identity for React keys. */
  key: string
  /** Short label under the group (epoch nonce, short address, …). */
  label: string
  /** Optional second line in the tooltip (block height, full address, …). */
  sublabel?: string
  /** One entry per series key; missing keys count as zero. */
  values: Record<string, number>
}

interface Props {
  series: ChartSeriesSpec[]
  groups: BarGroup[]
  /** Plot height in px, excluding the legend. */
  height?: number
  /** Shown instead of the plot when there is nothing to draw. */
  emptyMessage?: string
  /** Accessible description of what the chart shows. */
  ariaLabel: string
}

const PAD = { top: 10, right: 8, bottom: 32, left: 32 }
/** Gap between the bars of one group, in px — the spacer that keeps fills apart. */
const BAR_GAP = 2
/** Below this per-group width the plot scrolls sideways instead of squeezing. */
const MIN_GROUP_W = 46
/** Rough px per character of the 9px mono axis font, used to thin x labels. */
const LABEL_CHAR_W = 5.6

/**
 * Grouped bar chart, hand-drawn in SVG.
 *
 * Deliberately not a charting library: the whole thing is ~150 lines, it
 * inherits the editorial palette instead of fighting a theme adapter, and it
 * adds nothing to the bundle. Marks are thin, rounded at the data end and
 * anchored to the baseline; the grid is recessive; identity is carried by the
 * legend *and* the tooltip, never by colour alone.
 */
export function BarChart({ series, groups, height = 180, emptyMessage, ariaLabel }: Props) {
  const [ref, width] = useMeasuredWidth<HTMLDivElement>()
  const [hovered, setHovered] = useState<number | null>(null)

  const legend = (
    <HStack gap={4} wrap='wrap' mt={3}>
      {series.map((s) => (
        <HStack key={s.key} gap={2}>
          <Box w='9px' h='9px' borderRadius='2px' bg={s.color} flexShrink={0} />
          <Text fontSize='2xs' color='ink.3' fontFamily='mono' letterSpacing='0.04em'>
            {s.label}
          </Text>
        </HStack>
      ))}
    </HStack>
  )

  if (groups.length === 0) {
    return (
      <Box>
        <Box
          ref={ref}
          h={`${height}px`}
          borderWidth='1px'
          borderColor='border.subtle'
          borderRadius='md'
          bg='surface.sunken'
          display='flex'
          alignItems='center'
          justifyContent='center'
        >
          <Text fontSize='xs' color='ink.4'>
            {emptyMessage ?? 'Nothing to chart yet.'}
          </Text>
        </Box>
        {legend}
      </Box>
    )
  }

  // First paint has no measurement yet — reserve the exact height so the page
  // doesn't jump when the plot appears.
  if (width == null) {
    return (
      <Box>
        <Box ref={ref} h={`${height}px`} borderRadius='md' bg='surface.sunken' opacity={0.5} />
        {legend}
      </Box>
    )
  }

  // Wide data scrolls in its own container rather than compressing to a smear;
  // the page itself never scrolls horizontally.
  const w = Math.max(280, width, groups.length * MIN_GROUP_W + PAD.left + PAD.right)
  const plotW = w - PAD.left - PAD.right
  const plotH = height - PAD.top - PAD.bottom

  const rawMax = Math.max(
    0,
    ...groups.flatMap((g) => series.map((s) => g.values[s.key] ?? 0)),
  )
  // Even integer ceiling, so 0 / half / max are all whole numbers.
  const yMax = Math.max(2, Math.ceil(rawMax / 2) * 2)
  const ticks = [0, yMax / 2, yMax]

  const groupW = plotW / groups.length
  const innerPad = Math.min(12, groupW * 0.22)
  const barsW = Math.max(4, groupW - innerPad)
  const barW = Math.max(2, (barsW - BAR_GAP * (series.length - 1)) / series.length)

  const longestLabel = groups.reduce((n, g) => Math.max(n, g.label.length), 0)
  const labelEvery = Math.max(1, Math.ceil((longestLabel * LABEL_CHAR_W + 8) / groupW))

  const y = (v: number) => PAD.top + plotH - (v / yMax) * plotH
  const hoveredGroup = hovered != null ? groups[hovered] : null

  return (
    <Box>
      <Box ref={ref} position='relative' overflowX='auto' overflowY='hidden'>
        <svg
          width={w}
          height={height}
          role='img'
          aria-label={ariaLabel}
          style={{ display: 'block', overflow: 'visible' }}
        >
          {/* Recessive grid + y ticks. */}
          {ticks.map((t) => (
            <g key={t}>
              <line
                x1={PAD.left}
                x2={PAD.left + plotW}
                y1={y(t)}
                y2={y(t)}
                stroke={ChartGrid}
                strokeWidth={1}
              />
              <text
                x={PAD.left - 6}
                y={y(t) + 3}
                textAnchor='end'
                fontSize={9}
                fontFamily='JetBrains Mono, monospace'
                fill={ChartAxisInk}
              >
                {t}
              </text>
            </g>
          ))}

          {groups.map((g, i) => {
            const gx = PAD.left + i * groupW
            return (
              <g key={g.key}>
                {hovered === i && (
                  <rect
                    x={gx}
                    y={PAD.top}
                    width={groupW}
                    height={plotH}
                    fill='rgba(236, 232, 223, 0.045)'
                  />
                )}
                {series.map((s, j) => {
                  const v = g.values[s.key] ?? 0
                  // Zero draws nothing: a 1px stub would read as "a little",
                  // and the tooltip already reports the exact value.
                  if (v <= 0) return null
                  const h = Math.max(2, (v / yMax) * plotH)
                  const x = gx + innerPad / 2 + j * (barW + BAR_GAP)
                  const top = PAD.top + plotH - h
                  return (
                    <path
                      key={s.key}
                      d={roundedTopBar(x, top, barW, h, Math.min(3, barW / 2))}
                      fill={s.color}
                      opacity={hovered == null || hovered === i ? 1 : 0.45}
                    />
                  )
                })}
                {i % labelEvery === 0 && (
                  <text
                    x={gx + groupW / 2}
                    y={height - PAD.bottom + 14}
                    textAnchor='middle'
                    fontSize={9}
                    fontFamily='JetBrains Mono, monospace'
                    fill={ChartLabelInk}
                  >
                    {g.label}
                  </text>
                )}
                {/* Hit target: the whole column, so the hover works even where
                    a series is zero. */}
                <rect
                  x={gx}
                  y={PAD.top}
                  width={groupW}
                  height={plotH}
                  fill='transparent'
                  onMouseEnter={() => setHovered(i)}
                  onMouseLeave={() => setHovered((cur) => (cur === i ? null : cur))}
                >
                  <title>
                    {`${g.label}${g.sublabel ? ` (${g.sublabel})` : ''}: ` +
                      series.map((s) => `${s.label} ${g.values[s.key] ?? 0}`).join(', ')}
                  </title>
                </rect>
              </g>
            )
          })}

          {/* Baseline, drawn last so bars sit on it. */}
          <line
            x1={PAD.left}
            x2={PAD.left + plotW}
            y1={PAD.top + plotH}
            y2={PAD.top + plotH}
            stroke={ChartGrid}
            strokeWidth={1}
          />
        </svg>

        {hoveredGroup && (
          <Box
            position='absolute'
            top='0'
            left={`${clamp(
              PAD.left + (hovered! + 0.5) * groupW,
              90,
              Math.max(90, w - 90),
            )}px`}
            transform='translateX(-50%)'
            pointerEvents='none'
            bg='canvas.deep'
            borderWidth='1px'
            borderColor='border'
            borderRadius='md'
            px={3}
            py={2}
            minW='150px'
            zIndex={2}
            boxShadow='0 6px 18px rgba(0,0,0,0.45)'
          >
            <Stack gap={1.5}>
              <Text fontFamily='mono' fontSize='2xs' color='ink.1' letterSpacing='0.04em'>
                {hoveredGroup.label}
              </Text>
              {hoveredGroup.sublabel && (
                <Text fontFamily='mono' fontSize='2xs' color='ink.4'>
                  {hoveredGroup.sublabel}
                </Text>
              )}
              {series.map((s) => (
                <HStack key={s.key} gap={2} justify='space-between'>
                  <HStack gap={2}>
                    <Box w='7px' h='7px' borderRadius='2px' bg={s.color} flexShrink={0} />
                    <Text fontSize='2xs' color='ink.3'>
                      {s.label}
                    </Text>
                  </HStack>
                  <Text
                    className='dkg-tabular'
                    fontFamily='mono'
                    fontSize='2xs'
                    color='ink.0'
                  >
                    {hoveredGroup.values[s.key] ?? 0}
                  </Text>
                </HStack>
              ))}
            </Stack>
          </Box>
        )}
      </Box>
      {legend}
    </Box>
  )
}

/** Bar path with the two data-end corners rounded and the baseline square. */
function roundedTopBar(x: number, y: number, w: number, h: number, r: number): string {
  if (h <= 0) return ''
  const rad = Math.min(r, h)
  return [
    `M${x},${y + h}`,
    `L${x},${y + rad}`,
    `Q${x},${y} ${x + rad},${y}`,
    `L${x + w - rad},${y}`,
    `Q${x + w},${y} ${x + w},${y + rad}`,
    `L${x + w},${y + h}`,
    'Z',
  ].join(' ')
}

function clamp(v: number, lo: number, hi: number): number {
  return Math.min(hi, Math.max(lo, v))
}
