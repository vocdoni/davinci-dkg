import { Box, IconButton, Tooltip } from '@chakra-ui/react'
import { LuTerminal } from 'react-icons/lu'
import { useDebugMode } from '~hooks/use-debug-mode'

// Header affordance for toggling debug mode. Quiet by default; when on, a
// small phosphor dot lights up the corner of the icon so the active state
// is unmistakable from across the room without drawing the eye when off.
export function DebugModeToggle() {
  const { enabled, toggle } = useDebugMode()
  return (
    <Tooltip.Root>
      <Tooltip.Trigger asChild>
        <IconButton
          aria-label={enabled ? 'Disable debug mode' : 'Enable debug mode'}
          aria-pressed={enabled}
          variant='ghost'
          size='sm'
          onClick={toggle}
          color={enabled ? 'live.fg' : 'ink.3'}
          bg={enabled ? 'live.bg' : 'transparent'}
          borderRadius='full'
          borderWidth='1px'
          borderColor={enabled ? 'rgba(134, 239, 172, 0.30)' : 'border.subtle'}
          _hover={{ color: enabled ? 'live.bright' : 'ink.1', bg: enabled ? 'live.bg' : 'surface.raised' }}
          position='relative'
        >
          <LuTerminal />
          {enabled && (
            <Box
              position='absolute'
              top='-1px'
              right='-1px'
              w='6px'
              h='6px'
              borderRadius='full'
              bg='live.fg'
              boxShadow='0 0 8px rgba(134, 239, 172, 0.7)'
            />
          )}
        </IconButton>
      </Tooltip.Trigger>
      <Tooltip.Positioner>
        <Tooltip.Content fontFamily='sans' fontSize='xs'>
          {enabled ? 'Debug mode on — raw protocol data visible' : 'Debug mode off'}
        </Tooltip.Content>
      </Tooltip.Positioner>
    </Tooltip.Root>
  )
}
