import { Box, HStack, Stack, Text } from '@chakra-ui/react'
import type { Epoch } from '@vocdoni/davinci-dkg-sdk'
import { blocksToDuration } from '~lib/format'

interface Props {
  epoch: Epoch
  /** Chain head; drives the "you are here" marker. */
  currentBlock: bigint | null
  /** `EPOCH_DURATION_BLOCKS`; the Live window runs to `startBlock + duration`. */
  durationBlocks: bigint | null
}

interface Window {
  key: string
  label: string
  from: bigint
  to: bigint
  color: string
  hint: string
}

/**
 * The epoch drawn on a block axis: the three fixed Preparation windows and the
 * Live window that fills the rest, sized in proportion to their block counts,
 * with a marker at the chain head.
 *
 * The phase dots elsewhere say *which* phase an epoch is in; this says how much
 * of it is left, which is the question someone waiting for a key actually has.
 * Built out of boxes rather than SVG so it inherits the theme tokens and reflows
 * with the container for free.
 */
export function BlockWindowTimeline({ epoch, currentBlock, durationBlocks }: Props) {
  const start = epoch.startBlock
  const selectionEnd = epoch.policy.committeeSelectionDeadlineBlock
  const assemblyEnd = epoch.policy.keyAssemblyDeadlineBlock
  const liveAt = epoch.policy.liveNotBeforeBlock
  // Without the duration immutable the Live window has no right edge; fall back
  // to the end of the finalize gap so the Preparation windows still render.
  const end = durationBlocks != null ? start + durationBlocks : liveAt

  const windows: Window[] = [
    {
      key: 'selection',
      label: 'Committee selection',
      from: start,
      to: selectionEnd,
      color: 'accent.dim',
      hint: 'lottery claims',
    },
    {
      key: 'assembly',
      label: 'Key assembly',
      from: selectionEnd,
      to: assemblyEnd,
      color: 'accent.fg',
      hint: 'contributions',
    },
    {
      key: 'gap',
      label: 'Finalize gap',
      from: assemblyEnd,
      to: liveAt,
      color: 'border.strong',
      hint: 'finalize proof',
    },
    {
      key: 'live',
      label: 'Live',
      from: liveAt,
      to: end,
      color: 'live.fg',
      hint: 'apps encrypt & decrypt',
    },
  ]

  const total = Number(end - start)
  if (total <= 0) return null

  const span = (w: Window) => Math.max(0, Number(w.to - w.from))
  const pctOf = (block: bigint) => (Number(block - start) / total) * 100
  const headPct =
    currentBlock == null ? null : Math.min(100, Math.max(0, pctOf(currentBlock)))
  const past = (w: Window) => currentBlock != null && currentBlock >= w.to
  const isCurrent = (w: Window) =>
    currentBlock != null && currentBlock >= w.from && currentBlock < w.to

  return (
    <Stack gap={3}>
      <Box position='relative' pt={5}>
        {/* Head marker. Sits above the bar so it never hides a window edge. */}
        {headPct != null && currentBlock! >= start && currentBlock! <= end && (
          <Box
            position='absolute'
            top={0}
            bottom={0}
            left={`${headPct}%`}
            w='1px'
            bg='ink.1'
            opacity={0.75}
            zIndex={1}
          >
            <Text
              position='absolute'
              top='-2px'
              left='50%'
              transform='translateX(-50%)'
              whiteSpace='nowrap'
              className='dkg-tabular'
              fontFamily='mono'
              fontSize='2xs'
              color='ink.1'
              bg='surface'
              px={1}
            >
              #{currentBlock!.toString()}
            </Text>
          </Box>
        )}

        <HStack gap='2px' align='stretch' h='26px'>
          {windows.map((w) => {
            const width = (span(w) / total) * 100
            if (width <= 0) return null
            return (
              <Box
                key={w.key}
                flex={`0 0 ${width}%`}
                bg={w.color}
                opacity={past(w) ? 0.35 : isCurrent(w) ? 1 : 0.6}
                borderRadius='2px'
                position='relative'
                title={`${w.label}: blocks ${w.from.toString()}–${w.to.toString()} (${span(w)} blocks, ${w.hint})`}
              />
            )
          })}
        </HStack>
      </Box>

      {/* Legend rows: each window with its block range and wall-clock estimate. */}
      <HStack gap={{ base: 3, md: 5 }} wrap='wrap'>
        {windows.map((w) => (
          <Stack key={w.key} gap={0.5} minW='130px'>
            <HStack gap={2}>
              <Box w='8px' h='8px' borderRadius='2px' bg={w.color} opacity={isCurrent(w) ? 1 : 0.6} />
              <Text
                fontFamily='mono'
                fontSize='2xs'
                color={isCurrent(w) ? 'ink.0' : 'ink.3'}
                letterSpacing='0.04em'
                textTransform='uppercase'
              >
                {w.label}
              </Text>
            </HStack>
            <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.4'>
              #{w.from.toString()} → #{w.to.toString()}
            </Text>
            <Text fontSize='2xs' color='ink.4'>
              {span(w)} blocks · {blocksToDuration(span(w))}
            </Text>
          </Stack>
        ))}
      </HStack>
    </Stack>
  )
}
