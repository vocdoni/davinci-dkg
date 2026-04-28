import { useState, type ReactNode } from 'react'
import { Box, chakra, Collapsible, HStack, Text } from '@chakra-ui/react'
import { LuPlus, LuMinus } from 'react-icons/lu'
import { useDebugMode } from '~hooks/use-debug-mode'

// chakra.button gives us style props on a real <button>, sidestepping the
// polymorphic-as-prop type juggling that <Box as='button'> demands.
const TriggerBtn = chakra('button', {
  base: {
    bg: 'transparent',
    border: 'none',
    p: 0,
    cursor: 'pointer',
    textAlign: 'left',
  },
})

interface Props {
  /** Plain-English label for the disclosure trigger. */
  title?: string
  /** Children: the technical detail block. Always renders inside a styled box. */
  children: ReactNode
}

// Editorial appendix block. Reads like a footnote in a paper:
//   ── trigger    small mono "+ Show technical details" / "− Hide …",
//                 colour-shifts to gold on hover.
//   ── body       hairline-bordered card with an inset shadow, content
//                 in mono.
//
// Auto-opens when the global "debug mode" toggle is on; manual open/close
// is preserved when debug mode flips off so a researcher can keep one
// panel open while everything else collapses.
export function DetailDisclosure({ title = 'Show technical details', children }: Props) {
  const { enabled: debug } = useDebugMode()
  const [open, setOpen] = useState(debug)
  const isOpen = debug || open
  const label = isOpen ? title.replace(/^Show\b/i, 'Hide') : title

  return (
    <Box mt={3}>
      <TriggerBtn
        type='button'
        onClick={() => setOpen((v) => !v)}
        aria-expanded={isOpen}
        _hover={{ '& .dkg-disclosure-label': { color: 'accent.fg' } }}
      >
        <HStack gap={2}>
          <Box
            w='14px'
            h='14px'
            borderRadius='full'
            borderWidth='1px'
            borderColor='border.strong'
            display='flex'
            alignItems='center'
            justifyContent='center'
            color='ink.3'
            fontSize='8px'
          >
            {isOpen ? <LuMinus /> : <LuPlus />}
          </Box>
          <Text
            className='dkg-disclosure-label'
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.06em'
            textTransform='uppercase'
            transition='color 0.15s'
          >
            {label}
          </Text>
        </HStack>
      </TriggerBtn>
      <Collapsible.Root open={isOpen}>
        <Collapsible.Content>
          <Box
            mt={3}
            p={4}
            borderWidth='1px'
            borderColor='border.subtle'
            borderRadius='md'
            bg='surface.sunken'
            fontSize='xs'
            fontFamily='mono'
            color='ink.2'
            overflowX='auto'
            boxShadow='inset'
          >
            {children}
          </Box>
        </Collapsible.Content>
      </Collapsible.Root>
    </Box>
  )
}
