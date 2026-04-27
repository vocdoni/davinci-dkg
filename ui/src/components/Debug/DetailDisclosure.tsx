import { useState, type ReactNode } from 'react'
import { Box, Button, Collapsible } from '@chakra-ui/react'
import { LuChevronDown, LuChevronUp } from 'react-icons/lu'
import { useDebugMode } from '~hooks/use-debug-mode'

interface Props {
  /** Plain-English label for the disclosure trigger. */
  title?: string
  /** Children: the technical detail block. Always renders inside a styled box. */
  children: ReactNode
}

// Wraps technical content (raw hashes, JSON dumps, BigInt coords) so the
// default UX hides it. Auto-expands when global debug mode is on.
//
// Use anywhere the underlying value would otherwise leak protocol jargon
// into the primary view — e.g. event arg dumps, transcript hashes, raw
// curve coordinates. Cf. UI_PLAN.md §4.
export function DetailDisclosure({ title = 'Show technical details', children }: Props) {
  const { enabled: debug } = useDebugMode()
  const [open, setOpen] = useState(debug)

  // When debug mode flips on globally, force this disclosure open. When it
  // flips off we leave the local state alone — a researcher may want to
  // keep one panel open even after toggling debug off.
  const isOpen = debug || open

  return (
    <Box mt={2}>
      <Button
        variant='ghost'
        size='xs'
        onClick={() => setOpen((v) => !v)}
        color='gray.400'
        _hover={{ color: 'cyan.300', bg: 'transparent' }}
        px={0}
      >
        {isOpen ? <LuChevronUp /> : <LuChevronDown />}
        <Box as='span' ml={1}>
          {title}
        </Box>
      </Button>
      <Collapsible.Root open={isOpen}>
        <Collapsible.Content>
          <Box
            mt={2}
            p={3}
            borderWidth='1px'
            borderColor='gray.800'
            borderRadius='md'
            bg='gray.900'
            fontSize='xs'
            fontFamily='mono'
            color='gray.300'
            overflowX='auto'
          >
            {children}
          </Box>
        </Collapsible.Content>
      </Collapsible.Root>
    </Box>
  )
}
