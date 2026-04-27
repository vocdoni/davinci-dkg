import { IconButton, Tooltip } from '@chakra-ui/react'
import { LuTerminal } from 'react-icons/lu'
import { useDebugMode } from '~hooks/use-debug-mode'

// Header affordance for toggling debug mode. The icon is intentionally
// understated — power users will discover it; everyone else won't be
// distracted.
export function DebugModeToggle() {
  const { enabled, toggle } = useDebugMode()
  return (
    <Tooltip.Root>
      <Tooltip.Trigger asChild>
        <IconButton
          aria-label='Toggle debug mode'
          variant={enabled ? 'subtle' : 'ghost'}
          colorPalette={enabled ? 'cyan' : 'gray'}
          size='sm'
          onClick={toggle}
        >
          <LuTerminal />
        </IconButton>
      </Tooltip.Trigger>
      <Tooltip.Positioner>
        <Tooltip.Content>{enabled ? 'Debug mode on' : 'Debug mode off — show technical details'}</Tooltip.Content>
      </Tooltip.Positioner>
    </Tooltip.Root>
  )
}
