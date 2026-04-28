import { useState } from 'react'
import { Box, HStack, IconButton, Text, Tooltip } from '@chakra-ui/react'
import { LuCopy, LuCheck } from 'react-icons/lu'

interface Props {
  value: string | undefined | null
  /** Override default head/tail truncation lengths. */
  head?: number
  tail?: number
  /** Render the full hash in plain text instead of truncating. */
  full?: boolean
}

// Compact, copyable display of a hash/address. Shows the truncated form by
// default; full value on hover (tooltip), copy on click.
//
// The middle ellipsis is rendered as a typeset hairline rather than three
// dots — easier to read at small sizes and matches the editorial type
// system. The copy icon is hidden until the row is hovered, keeping
// dense tables uncluttered.
export function HashCell({ value, head, tail, full }: Props) {
  const [copied, setCopied] = useState(false)
  if (!value) {
    return (
      <Text fontFamily='mono' fontSize='xs' color='ink.4'>
        —
      </Text>
    )
  }

  const onCopy = async () => {
    try {
      await navigator.clipboard.writeText(value)
      setCopied(true)
      setTimeout(() => setCopied(false), 1200)
    } catch {
      /* clipboard denied — silently no-op */
    }
  }

  if (full) {
    return (
      <HStack
        gap={1.5}
        align='center'
        role='group'
      >
        <Text fontFamily='mono' fontSize='xs' color='ink.1' wordBreak='break-all'>
          {value}
        </Text>
        <CopyBtn copied={copied} onClick={onCopy} />
      </HStack>
    )
  }

  // Truncated: render head, hairline, tail in three pieces so the typeset
  // bridge is visually distinct from a literal "..." in a hash.
  const headLen = head ?? 6
  const tailLen = tail ?? 4
  const truncated = value.length > headLen + tailLen + 4
  if (!truncated) {
    return (
      <HStack gap={1.5} align='center' role='group'>
        <Text fontFamily='mono' fontSize='xs' color='ink.1'>
          {value}
        </Text>
        <CopyBtn copied={copied} onClick={onCopy} />
      </HStack>
    )
  }

  const head_ = value.slice(0, 2 + headLen)
  const tail_ = value.slice(-tailLen)

  return (
    <Tooltip.Root openDelay={400}>
      <Tooltip.Trigger asChild>
        <HStack gap={1.5} align='center' role='group' as='button' onClick={onCopy}>
          <HStack gap={0} align='center'>
            <Text fontFamily='mono' fontSize='xs' color='ink.1'>
              {head_}
            </Text>
            <Box
              as='span'
              w='10px'
              h='1px'
              bg='ink.4'
              mx='3px'
              opacity={0.7}
              alignSelf='center'
            />
            <Text fontFamily='mono' fontSize='xs' color='ink.1'>
              {tail_}
            </Text>
          </HStack>
          <CopyBtn copied={copied} onClick={onCopy} />
        </HStack>
      </Tooltip.Trigger>
      <Tooltip.Positioner>
        <Tooltip.Content fontFamily='mono' fontSize='2xs'>
          {value}
        </Tooltip.Content>
      </Tooltip.Positioner>
    </Tooltip.Root>
  )
}

function CopyBtn({ copied, onClick }: { copied: boolean; onClick: () => void }) {
  return (
    <IconButton
      aria-label={copied ? 'Copied' : 'Copy'}
      size='2xs'
      variant='ghost'
      onClick={(e) => {
        e.stopPropagation()
        onClick()
      }}
      color={copied ? 'live.fg' : 'ink.4'}
      bg='transparent'
      opacity={copied ? 1 : 0}
      _groupHover={{ opacity: 1 }}
      _hover={{ color: copied ? 'live.fg' : 'accent.fg', bg: 'transparent' }}
      transition='opacity 0.15s, color 0.15s'
    >
      {copied ? <LuCheck /> : <LuCopy />}
    </IconButton>
  )
}
