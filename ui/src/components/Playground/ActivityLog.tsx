import { useEffect, useRef, useState } from 'react'
import { Box, Button, HStack, Text } from '@chakra-ui/react'
import { LuCheck, LuClipboardCopy } from 'react-icons/lu'

export type LogLevel = 'info' | 'success' | 'error' | 'tx' | 'chain' | 'crypto'

export interface LogEntry {
  ts: number
  level: LogLevel
  msg: string
}

// Each level maps to one of the editorial-palette tokens. Keep the spread
// narrow so the log reads as one piece of typography, not a kaleidoscope.
const colorByLevel: Record<LogLevel, string> = {
  info: 'ink.2',
  success: 'live.fg',
  error: 'danger.fg',
  tx: 'amber.300',
  chain: 'accent.bright',
  crypto: 'live.bright',
}

function entryToText(e: LogEntry): string {
  const ts = new Date(e.ts).toISOString()
  return `[${ts}] [${e.level.padEnd(7)}] ${e.msg}`
}

// Terminal-style chronological log of everything the playground does
// on-chain or in-crypto. Editorial dark plate (deeper than surrounding
// surfaces), JetBrains Mono, fixed-width timestamps. Auto-scrolls to the
// newest entry. The Copy button dumps the whole log as plain text — handy
// for pasting into a bug report.
export function ActivityLog({ entries }: { entries: LogEntry[] }) {
  const ref = useRef<HTMLDivElement>(null)
  const [copied, setCopied] = useState(false)

  useEffect(() => {
    if (ref.current) ref.current.scrollTop = ref.current.scrollHeight
  }, [entries])

  const onCopy = async () => {
    if (entries.length === 0) return
    const text = entries.map(entryToText).join('\n')
    try {
      await navigator.clipboard.writeText(text)
      setCopied(true)
      setTimeout(() => setCopied(false), 1500)
    } catch {
      console.warn('Clipboard write failed; printing log to console:\n', text)
    }
  }

  return (
    <Box>
      <HStack justify='space-between' mb={2}>
        <Text fontFamily='mono' fontSize='2xs' color='ink.4' letterSpacing='0.04em'>
          {entries.length} {entries.length === 1 ? 'entry' : 'entries'}
        </Text>
        <Button
          size='2xs'
          variant='ghost'
          onClick={onCopy}
          disabled={entries.length === 0}
          color={copied ? 'live.fg' : 'ink.3'}
          fontFamily='sans'
          _hover={{ color: copied ? 'live.fg' : 'accent.fg', bg: 'transparent' }}
        >
          <HStack gap={1}>
            {copied ? <LuCheck /> : <LuClipboardCopy />}
            <Text fontSize='2xs'>{copied ? 'Copied' : 'Copy'}</Text>
          </HStack>
        </Button>
      </HStack>
      <Box
        ref={ref}
        bg='canvas.deep'
        p={3}
        borderRadius='md'
        borderWidth='1px'
        borderColor='border.subtle'
        fontFamily='mono'
        fontSize='2xs'
        lineHeight='1.6'
        maxH='320px'
        overflowY='auto'
        boxShadow='inset 0 0 0 1px rgba(0,0,0,0.3)'
      >
        {entries.length === 0 && (
          <Text color='ink.4' fontSize='xs'>
            Activity will appear here as you walk through the steps…
          </Text>
        )}
        {entries.map((e, i) => (
          <Box key={i} color={colorByLevel[e.level]} className='dkg-tabular'>
            <Text as='span' color='ink.4' userSelect='none' mr={2}>
              {new Date(e.ts).toLocaleTimeString('en', { hour12: false })}
            </Text>
            {e.msg}
          </Box>
        ))}
      </Box>
    </Box>
  )
}
