import { Box } from '@chakra-ui/react'

// Stringify any value (BigInt-safe) for display inside a DetailDisclosure.
// JSON.stringify can't serialise BigInt by default, so we patch it here
// rather than asking every caller to remember.
function stringify(value: unknown): string {
  return JSON.stringify(value, (_k, v) => (typeof v === 'bigint' ? v.toString() : v), 2)
}

export function RawJson({ value }: { value: unknown }) {
  return (
    <Box as='pre' whiteSpace='pre-wrap' wordBreak='break-all' m={0}>
      {stringify(value)}
    </Box>
  )
}
