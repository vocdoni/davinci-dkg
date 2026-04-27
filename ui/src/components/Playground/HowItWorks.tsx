import { Box, HStack, Stack, Text } from '@chakra-ui/react'
import type { ReactNode } from 'react'
import { LuChevronRight } from 'react-icons/lu'
import { DetailDisclosure } from '~components/Debug/DetailDisclosure'

interface FlowStep {
  icon: ReactNode
  label: string
}

interface Props {
  /** One-paragraph plain-English explanation. */
  body: ReactNode
  /** Optional flow strip — a row of (icon, label) pairs separated by chevrons. */
  flow?: FlowStep[]
}

// Per-step "How this works" disclosure. The body is plain English (no
// protocol jargon), and the flow strip — when supplied — gives a quick
// visual of the actors and order of operations using lucide icons. No SVG
// authoring required, but still gives the page a graphical anchor.
//
// Closed by default so the step UI stays compact; debug mode (header
// terminal icon) auto-expands it via DetailDisclosure.
export function HowItWorks({ body, flow }: Props) {
  return (
    <DetailDisclosure title='How this step works'>
      <Stack gap={4}>
        <Text fontSize='xs' color='gray.300' lineHeight='1.6' fontFamily='body'>
          {body}
        </Text>
        {flow && flow.length > 0 && (
          <HStack gap={0} wrap='wrap' justify='center'>
            {flow.map((s, i) => (
              <HStack key={`${s.label}-${i}`} gap={0}>
                <Stack align='center' gap={1} px={3} py={2}>
                  <Box
                    w={9}
                    h={9}
                    borderRadius='full'
                    bg='gray.800'
                    color='cyan.300'
                    display='flex'
                    alignItems='center'
                    justifyContent='center'
                    fontSize='lg'
                  >
                    {s.icon}
                  </Box>
                  <Text fontSize='2xs' color='gray.400' textAlign='center' maxW='90px'>
                    {s.label}
                  </Text>
                </Stack>
                {i < flow.length - 1 && (
                  <Box color='gray.600'>
                    <LuChevronRight />
                  </Box>
                )}
              </HStack>
            ))}
          </HStack>
        )}
      </Stack>
    </DetailDisclosure>
  )
}
