import { Box, Stack, Text } from '@chakra-ui/react'
import type { ReactNode } from 'react'

interface Props {
  label: string
  value: ReactNode
  hint?: ReactNode
}

// Tiny presentational primitive used by Home and Registry to render the
// row of KPI cards uniformly.
export function StatCard({ label, value, hint }: Props) {
  return (
    <Box borderWidth='1px' borderColor='gray.800' p={4} borderRadius='md' bg='gray.900'>
      <Stack gap={1}>
        <Text fontSize='xs' color='gray.500'>
          {label}
        </Text>
        <Text fontSize='lg' fontWeight='semibold' fontFamily='mono'>
          {value}
        </Text>
        {hint && (
          <Text fontSize='xs' color='gray.500'>
            {hint}
          </Text>
        )}
      </Stack>
    </Box>
  )
}
