import { useEffect, useRef, useState } from 'react'
import { Box, HStack, Text } from '@chakra-ui/react'
import { useEstimatedBlock } from '~hooks/use-estimated-block'

interface Props {
  /** Target block number to count down to. */
  target: bigint | number
  /** What's happening at the target block — "until contributions close" etc. */
  label: string
  /** Average seconds per block on this chain. Defaults to 12s (mainnet/sepolia). */
  blockTimeSeconds?: number
}

// Live "X blocks (~MM:SS) until <label>" indicator. Editorial register:
// large mono digits for the time remaining, then a hairline divider, then
// the label in italic body serif. Reads like an inline footnote rather than
// a generic "X minutes left" countdown.
export function Countdown({ target, label, blockTimeSeconds = 12 }: Props) {
  const estimated = useEstimatedBlock(blockTimeSeconds)
  // Re-render every second so the "MM:SS" string ticks even when the
  // estimated block hasn't bumped yet (estimated only changes once per
  // blockTimeSeconds on average).
  const [, force] = useState(0)
  const t = useRef<ReturnType<typeof setInterval> | null>(null)
  useEffect(() => {
    t.current = setInterval(() => force((n) => n + 1), 1000)
    return () => {
      if (t.current) clearInterval(t.current)
    }
  }, [])

  if (estimated == null) {
    return (
      <Row tone='dim' time='—' label={label} blocks={null} />
    )
  }

  const targetBig = typeof target === 'bigint' ? target : BigInt(target)
  const blocksLeft = Number(targetBig - estimated)

  if (blocksLeft <= 0) {
    return <Row tone='live' time='ready' label={label} blocks={null} />
  }

  const secondsLeft = blocksLeft * blockTimeSeconds
  const mm = Math.floor(secondsLeft / 60)
  const ss = secondsLeft % 60
  const time = mm > 0 ? `${mm.toString().padStart(2, '0')}:${ss.toString().padStart(2, '0')}` : `0:${ss.toString().padStart(2, '0')}`

  return <Row tone='accent' time={time} label={label} blocks={blocksLeft} />
}

function Row({
  tone,
  time,
  label,
  blocks,
}: {
  tone: 'accent' | 'live' | 'dim'
  time: string
  label: string
  blocks: number | null
}) {
  const timeColor = tone === 'live' ? 'live.fg' : tone === 'dim' ? 'ink.4' : 'accent.fg'
  return (
    <HStack gap={3} align='baseline' flexWrap='wrap'>
      <Text
        className='dkg-tabular'
        fontFamily='mono'
        fontSize='md'
        fontWeight={600}
        color={timeColor}
        letterSpacing='0.02em'
      >
        {time}
      </Text>
      {blocks != null && (
        <>
          <Box w='1px' h='12px' bg='border.strong' alignSelf='center' />
          <Text className='dkg-tabular' fontFamily='mono' fontSize='2xs' color='ink.3'>
            {blocks} {blocks === 1 ? 'blk' : 'blks'}
          </Text>
        </>
      )}
      <Text fontSize='sm' color='ink.3'>
        {tone === 'live' ? `${label} now` : label}
      </Text>
    </HStack>
  )
}
