import { Box, Container, Flex, HStack, Heading, Spacer } from '@chakra-ui/react'
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

export function Header() {
  return (
    <Box
      as='header'
      borderBottomWidth='1px'
      borderColor='gray.800'
      bg='gray.950'
      position='sticky'
      top={0}
      zIndex={10}
      backdropFilter='blur(12px)'
    >
      <Container maxW='7xl' py={3}>
        <Flex align='center' gap={6}>
          <NavLink to={Routes.home}>
            <Heading size='md' color='cyan.300' letterSpacing='tight'>
              davinci-dkg
            </Heading>
          </NavLink>
          <HStack gap={1} fontSize='sm'>
            {navItems.map((item) => (
              <NavLink key={item.to} to={item.to} end={item.to === Routes.home}>
                {({ isActive }) => (
                  <Box
                    px={3}
                    py={1.5}
                    borderRadius='md'
                    color={isActive ? 'cyan.300' : 'gray.400'}
                    bg={isActive ? 'gray.900' : 'transparent'}
                    _hover={{ color: 'cyan.200', bg: 'gray.900' }}
                    transition='all 0.15s'
                  >
                    {item.label}
                  </Box>
                )}
              </NavLink>
            ))}
          </HStack>
          <Spacer />
          <HStack gap={3}>
            <NetworkBadge />
            <DebugModeToggle />
            <ConnectButton />
          </HStack>
        </Flex>
      </Container>
    </Box>
  )
}
