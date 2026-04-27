import { useState } from 'react'
import { HStack, IconButton, Text, Tooltip } from '@chakra-ui/react'
import { LuCopy, LuCheck } from 'react-icons/lu'
import { shortHash } from '~lib/format'

interface Props {
  value: string | undefined | null
  /** Override default head/tail truncation lengths. */
  head?: number
  tail?: number
  /** Render the full hash in plain text instead of truncating. */
  full?: boolean
}

// Compact, copyable display of a hash/address. Shows the truncated form by
// default; copy-on-click + tooltip with the full value.
export function HashCell({ value, head, tail, full }: Props) {
  const [copied, setCopied] = useState(false)
  if (!value) return <Text fontSize='xs' color='gray.500'>—</Text>

  const display = full ? value : shortHash(value, head, tail)
  const onCopy = async () => {
    try {
      await navigator.clipboard.writeText(value)
      setCopied(true)
      setTimeout(() => setCopied(false), 1200)
    } catch {
      /* ignore */
    }
  }

  return (
    <HStack gap={1} align='center'>
      <Tooltip.Root>
        <Tooltip.Trigger asChild>
          <Text fontFamily='mono' fontSize='xs' color='gray.300'>
            {display}
          </Text>
        </Tooltip.Trigger>
        <Tooltip.Positioner>
          <Tooltip.Content>{value}</Tooltip.Content>
        </Tooltip.Positioner>
      </Tooltip.Root>
      <IconButton aria-label='Copy' size='2xs' variant='ghost' onClick={onCopy}>
        {copied ? <LuCheck /> : <LuCopy />}
      </IconButton>
    </HStack>
  )
}
