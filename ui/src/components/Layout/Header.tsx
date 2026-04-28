import { Box, Container, Flex, HStack, Spacer, Text } from '@chakra-ui/react'
import { NavLink } from 'react-router-dom'
import { Routes } from '~router/routes'
import { ConnectButton } from './ConnectButton'
import { DebugModeToggle } from './DebugModeToggle'
import { NetworkBadge } from './NetworkBadge'

const navItems: { to: string; label: string }[] = [
  { to: Routes.home, label: 'Overview' },
  { to: Routes.rounds, label: 'Rounds' },
  { to: Routes.registry, label: 'Registry' },
  { to: Routes.playground, label: 'Playground' },
  { to: Routes.runNode, label: 'Run a node' },
  { to: Routes.sdk, label: 'SDK' },
  { to: Routes.settings, label: 'Settings' },
]

// Single-line top bar:
//
//   [wordmark]              [nav · centered]              [actions]
//
// The wordmark is restrained — small mono with a single accent dot —
// so it doesn't fight the page heading on landing/section pages. The
// nav sits in the middle with subtle hover + active states (no pills,
// no large coloured backgrounds). Actions on the right are minimised
// chrome: a small chain status pill + debug toggle + connect button.
export function Header() {
  return (
    <Box
      as='header'
      position='sticky'
      top={0}
      zIndex={20}
      bg='rgba(10, 10, 12, 0.78)'
      backdropFilter='saturate(180%) blur(20px)'
      borderBottomWidth='1px'
      borderColor='rule'
    >
      <Container maxW='7xl' px={{ base: 4, md: 6 }} py={3.5}>
        <Flex align='center' gap={{ base: 4, md: 6 }}>
          <NavLink to={Routes.home} aria-label='davinci-dkg'>
            <HStack gap={2} align='center'>
              <Box w='6px' h='6px' borderRadius='full' bg='accent.fg' />
              <Text
                fontFamily='mono'
                fontSize='sm'
                fontWeight={500}
                color='ink.0'
                letterSpacing='-0.01em'
              >
                davinci-dkg
              </Text>
            </HStack>
          </NavLink>

          <Spacer />

          {/* Centered nav. Underline-on-active. Hover lifts the colour
              from ink.3 to ink.0; no chip background. */}
          <HStack
            as='nav'
            gap={{ base: 0, md: 1 }}
            display={{ base: 'none', md: 'flex' }}
          >
            {navItems.map((item) => (
              <NavLink key={item.to} to={item.to} end={item.to === Routes.home}>
                {({ isActive }) => (
                  <Box
                    position='relative'
                    px={3}
                    py={1.5}
                    fontSize='sm'
                    fontWeight={isActive ? 500 : 400}
                    color={isActive ? 'ink.0' : 'ink.3'}
                    transition='color 0.15s ease'
                    _hover={{ color: 'ink.1' }}
                    _after={{
                      content: '""',
                      position: 'absolute',
                      left: 3,
                      right: 3,
                      bottom: '-13px',
                      height: '1px',
                      bg: isActive ? 'accent.fg' : 'transparent',
                      transition: 'background 0.15s ease',
                    }}
                  >
                    {item.label}
                  </Box>
                )}
              </NavLink>
            ))}
          </HStack>

          <Spacer />

          <HStack gap={2}>
            <Box display={{ base: 'none', lg: 'block' }}>
              <NetworkBadge />
            </Box>
            <DebugModeToggle />
            <ConnectButton />
          </HStack>
        </Flex>
      </Container>
    </Box>
  )
}
