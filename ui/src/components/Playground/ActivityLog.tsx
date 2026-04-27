import { useEffect, useRef, useState } from 'react'
import { Box, Button, HStack, Text } from '@chakra-ui/react'
import { LuCheck, LuClipboardCopy } from 'react-icons/lu'

export type LogLevel = 'info' | 'success' | 'error' | 'tx' | 'chain' | 'crypto'

export interface LogEntry {
  ts: number
  level: LogLevel
  msg: string
}

const colorByLevel: Record<LogLevel, string> = {
  info: 'gray.300',
  success: 'green.300',
  error: 'red.400',
  tx: 'yellow.300',
  chain: 'cyan.300',
  crypto: 'purple.300',
}

function entryToText(e: LogEntry): string {
  const ts = new Date(e.ts).toISOString()
  return `[${ts}] [${e.level.padEnd(7)}] ${e.msg}`
}

// Terminal-style chronological log of everything the playground does
// on-chain or in-crypto. Shown by default for debug-mode users; collapsible
// for everyone else. Auto-scrolls to the newest entry. The Copy button
// dumps the whole log as plain text — handy for pasting into a bug report.
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
        <Text fontSize='2xs' color='gray.500'>
          {entries.length} {entries.length === 1 ? 'entry' : 'entries'}
        </Text>
        <Button
          size='2xs'
          variant='ghost'
          onClick={onCopy}
          disabled={entries.length === 0}
          color='gray.400'
          _hover={{ color: 'cyan.300', bg: 'transparent' }}
        >
          <HStack gap={1}>
            {copied ? <LuCheck /> : <LuClipboardCopy />}
            <Text fontSize='2xs'>{copied ? 'Copied' : 'Copy'}</Text>
          </HStack>
        </Button>
      </HStack>
      <Box
        ref={ref}
        bg='black'
        p={3}
        borderRadius='md'
        fontFamily='mono'
        fontSize='xs'
        maxH='320px'
        overflowY='auto'
        borderWidth='1px'
        borderColor='gray.800'
      >
        {entries.length === 0 && (
          <Text color='gray.600'>Activity will appear here as you walk through the steps…</Text>
        )}
        {entries.map((e, i) => (
          <Box key={i} color={colorByLevel[e.level]} mb='1px'>
            <Text as='span' color='gray.600' userSelect='none'>
              [{new Date(e.ts).toLocaleTimeString('en', { hour12: false })}]{' '}
            </Text>
            {e.msg}
          </Box>
        ))}
      </Box>
    </Box>
  )
}
