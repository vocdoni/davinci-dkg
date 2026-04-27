import { useEffect, useRef, useState } from 'react'
import { HStack, Text } from '@chakra-ui/react'
import { LuClock } from 'react-icons/lu'
import { useEstimatedBlock } from '~hooks/use-estimated-block'

interface Props {
  /** Target block number to count down to. */
  target: bigint | number
  /** What's happening at the target block — "until contributions close" etc. */
  label: string
  /** Average seconds per block on this chain. Defaults to 12s (mainnet/sepolia). */
  blockTimeSeconds?: number
}

// Live "X blocks (~M:SS) left until <label>" indicator. Combines the slow
// 4s block-poll with a 1Hz tick to render a smoothly-counting timer.
//
// Once the target is reached, the component renders a green "now" pill so
// the user knows the gate just opened — useful right at phase boundaries.
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
      <HStack gap={2} fontSize='xs' color='gray.500'>
        <LuClock />
        <Text>… {label}</Text>
      </HStack>
    )
  }

  const targetBig = typeof target === 'bigint' ? target : BigInt(target)
  const blocksLeft = Number(targetBig - estimated)

  if (blocksLeft <= 0) {
    return (
      <HStack gap={2} fontSize='xs' color='green.300'>
        <LuClock />
        <Text fontWeight='semibold'>{label} now</Text>
      </HStack>
    )
  }

  const secondsLeft = blocksLeft * blockTimeSeconds
  const mm = Math.floor(secondsLeft / 60)
  const ss = secondsLeft % 60
  const human = mm > 0 ? `${mm}m ${ss.toString().padStart(2, '0')}s` : `${ss}s`

  return (
    <HStack gap={2} fontSize='xs' color='cyan.300'>
      <LuClock />
      <Text fontFamily='mono'>
        {human} ({blocksLeft} {blocksLeft === 1 ? 'block' : 'blocks'}) {label}
      </Text>
    </HStack>
  )
}
