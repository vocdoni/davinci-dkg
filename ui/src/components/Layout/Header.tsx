import { Box, Container, Flex, HStack, Text } from '@chakra-ui/react'
import { NavLink } from 'react-router-dom'
import { Routes } from '~router/routes'
import { ChainPill } from './ChainPill'
import { ConnectButton } from './ConnectButton'
import { DebugModeToggle } from './DebugModeToggle'
import { MobileNav } from './MobileNav'

interface NavItem {
  to: string
  label: string
  hint?: string
  /** Smallest breakpoint at which this item appears inline. Items
   *  without an inlineFrom only appear in the hamburger drawer. */
  inlineFrom?: 'md' | 'lg' | 'xl'
}

// Primary nav surface area.
//
//   Overview / Rounds / Registry / Playground are the everyday-use
//   destinations and stay inline from `md`. Run a node, SDK and
//   Settings are secondary; they only fit inline once we reach `xl`.
//   Until then, the hamburger drawer carries them.
const navItems: NavItem[] = [
  { to: Routes.home, label: 'Overview', hint: 'Protocol summary', inlineFrom: 'md' },
  { to: Routes.rounds, label: 'Rounds', hint: 'On-chain ring buffer', inlineFrom: 'md' },
  { to: Routes.registry, label: 'Registry', hint: 'Operator nodes', inlineFrom: 'md' },
  { to: Routes.playground, label: 'Playground', hint: 'Interactive walkthrough', inlineFrom: 'md' },
  { to: Routes.runNode, label: 'Run a node', hint: 'Operator handbook', inlineFrom: 'xl' },
  { to: Routes.sdk, label: 'SDK', hint: 'TypeScript reference', inlineFrom: 'xl' },
  { to: Routes.settings, label: 'Settings', hint: 'UI preferences', inlineFrom: 'xl' },
]

// Single-row, breakpoint-aware top bar.
//
//   base..sm  (<768) :  [wordmark]                              [connect][≡]
//   md        (≥768) :  [wordmark]   nav (4 primary)            [connect][≡]
//   lg        (≥1024):  [wordmark]   nav (4 primary)   [pill compact][⌨][connect][≡]
//   xl        (≥1280):  [wordmark]   nav (all 7)         [pill full ][⌨][connect]
//
// Everything below is padding/gap tuning. The bar is always exactly one
// Flex row — never wraps. The hamburger drawer carries every secondary
// affordance until xl, where the layout has room for them inline.
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
      <Container maxW='none' px={{ base: 3, md: 5, lg: 6 }} py={3}>
        <Flex align='center' gap={{ base: 2, md: 3, lg: 4 }} minH='40px'>
          <Brand />

          {/* Inline nav — natural width (no flex stretch). The right
              cluster is pushed to the end by `ml='auto'`. */}
          <HStack
            as='nav'
            ml={{ md: 4, lg: 4, xl: 6 }}
            gap={{ md: 0, lg: 1 }}
            display={{ base: 'none', md: 'flex' }}
          >
            {navItems.map((item) => (
              <DesktopNavLink
                key={item.to}
                to={item.to}
                label={item.label}
                inlineFrom={item.inlineFrom}
              />
            ))}
          </HStack>

          {/* Right cluster. Auto-margin pushes it against the right edge.
              Chain pill is compact (no contract) at lg, full at xl.
              Debug toggle appears at lg+. The hamburger drops away only
              at xl, where every nav item can fit inline. */}
          <HStack gap={2} flexShrink={0} ml='auto'>
            <Box display={{ base: 'none', xl: 'block' }}>
              <ChainPill />
            </Box>
            <Box display={{ base: 'none', lg: 'block', xl: 'none' }}>
              <ChainPill compact />
            </Box>
            <Box display={{ base: 'none', lg: 'block' }}>
              <DebugModeToggle />
            </Box>
            <ConnectButton />
            <Box display={{ base: 'block', xl: 'none' }}>
              <MobileNav items={navItems} />
            </Box>
          </HStack>
        </Flex>
      </Container>
    </Box>
  )
}

// Compact wordmark — accent dot + mono label. Same across breakpoints.
function Brand() {
  return (
    <NavLink to={Routes.home} aria-label='davinci-dkg'>
      <HStack gap={2} align='center'>
        <Box w='6px' h='6px' borderRadius='full' bg='accent.fg' flexShrink={0} />
        <Text
          fontFamily='mono'
          fontSize='sm'
          fontWeight={500}
          color='ink.0'
          letterSpacing='-0.01em'
          whiteSpace='nowrap'
        >
          davinci-dkg
        </Text>
      </HStack>
    </NavLink>
  )
}

// Inline nav link. The `inlineFrom` knob controls when the item appears
// — `undefined` means "drawer-only, never inline".
function DesktopNavLink({
  to,
  label,
  inlineFrom,
}: {
  to: string
  label: string
  inlineFrom?: 'md' | 'lg' | 'xl'
}) {
  const display = !inlineFrom
    ? 'none'
    : inlineFrom === 'md'
      ? { base: 'none', md: 'flex' }
      : inlineFrom === 'lg'
        ? { base: 'none', lg: 'flex' }
        : { base: 'none', xl: 'flex' }
  return (
    <NavLink to={to} end={to === Routes.home}>
      {({ isActive }) => (
        <Box
          display={display}
          position='relative'
          px={{ md: 2, lg: 3 }}
          py={1.5}
          fontSize='sm'
          fontWeight={isActive ? 500 : 400}
          color={isActive ? 'ink.0' : 'ink.3'}
          transition='color 0.15s ease'
          whiteSpace='nowrap'
          _hover={{ color: 'ink.1' }}
          _after={{
            content: '""',
            position: 'absolute',
            left: { md: 2, lg: 3 },
            right: { md: 2, lg: 3 },
            bottom: '-13px',
            height: '1px',
            bg: isActive ? 'accent.fg' : 'transparent',
            transition: 'background 0.15s ease',
          }}
        >
          {label}
        </Box>
      )}
    </NavLink>
  )
}
