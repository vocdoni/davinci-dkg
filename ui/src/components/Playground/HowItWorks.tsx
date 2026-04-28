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

// Per-step "How this works" appendix. Plain-English body in italic body
// serif; optional flow strip below as a row of icon nodes connected by
// chevrons. The flow circles are subtle gold-bordered, not filled, so the
// strip reads as a diagram rather than as a row of buttons.
export function HowItWorks({ body, flow }: Props) {
  return (
    <DetailDisclosure title='How this step works'>
      <Stack gap={5}>
        <Text fontSize='sm' color='ink.2' lineHeight='1.65'>
          {body}
        </Text>
        {flow && flow.length > 0 && (
          <Box overflowX='auto' mx={{ base: -2, md: 0 }} px={{ base: 2, md: 0 }}>
            <HStack gap={0} wrap='nowrap' justify={{ base: 'flex-start', md: 'center' }} py={2} minW='min-content'>
              {flow.map((s, i) => (
                <HStack key={`${s.label}-${i}`} gap={0} flexShrink={0}>
                  <Stack align='center' gap={2} px={{ base: 2, md: 4 }} flexShrink={0}>
                    <Box
                      w={{ base: 9, md: 10 }}
                      h={{ base: 9, md: 10 }}
                      borderRadius='full'
                      bg='surface.raised'
                      borderWidth='1px'
                      borderColor='accent.border'
                      color='accent.fg'
                      display='flex'
                      alignItems='center'
                      justifyContent='center'
                      fontSize='md'
                      flexShrink={0}
                    >
                      {s.icon}
                    </Box>
                    <Text
                      fontFamily='mono'
                      fontSize='2xs'
                      color='ink.3'
                      textAlign='center'
                      maxW={{ base: '88px', md: '110px' }}
                      letterSpacing='0.04em'
                      lineHeight='1.3'
                    >
                      {s.label}
                    </Text>
                  </Stack>
                  {i < flow.length - 1 && (
                    <Box color='ink.4' alignSelf='center' mb={5} flexShrink={0}>
                      <LuChevronRight />
                    </Box>
                  )}
                </HStack>
              ))}
            </HStack>
          </Box>
        )}
      </Stack>
    </DetailDisclosure>
  )
}
