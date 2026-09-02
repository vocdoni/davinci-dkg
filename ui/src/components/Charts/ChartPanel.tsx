import { Box, HStack, Stack, Text } from '@chakra-ui/react'
import type { ReactNode } from 'react'

interface Props {
  title: string
  /** One line saying what the reader is looking at. */
  caption?: ReactNode
  /** Reserved plot height; the skeleton uses it so nothing jumps on load. */
  height?: number
  loading?: boolean
  children: ReactNode
}

/**
 * Panel shell shared by the explorer's charts: mono small-caps title, a caption
 * that explains the measurement, and a fixed-height skeleton so a slow log scan
 * never reflows the page around it.
 */
export function ChartPanel({ title, caption, height = 180, loading, children }: Props) {
  return (
    <Box
      borderWidth='1px'
      borderColor='border.subtle'
      borderRadius='lg'
      bg='surface'
      p={{ base: 4, md: 5 }}
      boxShadow='inset'
    >
      <Stack gap={caption ? 1 : 3}>
        <HStack justify='space-between' align='baseline' gap={3} wrap='wrap'>
          <Text
            fontFamily='mono'
            fontSize='2xs'
            color='ink.3'
            letterSpacing='0.08em'
            textTransform='uppercase'
          >
            {title}
          </Text>
        </HStack>
        {caption && (
          <Text fontSize='xs' color='ink.4' lineHeight='1.5' mb={2} maxW='68ch'>
            {caption}
          </Text>
        )}
        {loading ? <ChartSkeleton height={height} /> : children}
      </Stack>
    </Box>
  )
}

function ChartSkeleton({ height }: { height: number }) {
  return (
    <Box
      h={`${height}px`}
      borderRadius='md'
      bg='surface.sunken'
      css={{ animation: 'dkgSkeletonPulse 1.6s ease-in-out infinite' }}
      aria-hidden
    />
  )
}
