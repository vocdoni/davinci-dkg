import { useEffect, useState } from 'react'
import { Box, Drawer, HStack, IconButton, Portal, Stack, Text } from '@chakra-ui/react'
import { LuMenu, LuX } from 'react-icons/lu'
import { NavLink, useLocation } from 'react-router-dom'
import { useDebugMode } from '~hooks/use-debug-mode'
import { ChainPill } from './ChainPill'

interface NavItem {
  to: string
  label: string
  /** Optional one-line caption shown under the label inside the drawer. */
  hint?: string
}

interface Props {
  items: NavItem[]
}

// Slide-in drawer for the < lg viewports. Mirrors what the desktop header
// shows on the right: full nav, chain status pill, debug toggle. The
// connect button stays visible in the header itself so users can sign in
// without opening the menu.
//
// The drawer auto-closes when the route changes — without it, tapping a
// link would route correctly but leave the panel covering the new page.
export function MobileNav({ items }: Props) {
  const [open, setOpen] = useState(false)
  const location = useLocation()
  const { enabled: debug, toggle: toggleDebug } = useDebugMode()

  useEffect(() => {
    setOpen(false)
  }, [location.pathname])

  return (
    <Drawer.Root
      open={open}
      onOpenChange={(d) => setOpen(d.open)}
      placement='end'
      size='xs'
    >
      <Drawer.Trigger asChild>
        <IconButton
          aria-label={open ? 'Close menu' : 'Open menu'}
          variant='ghost'
          size='sm'
          color='ink.2'
          borderRadius='full'
          borderWidth='1px'
          borderColor='border.subtle'
          _hover={{ bg: 'surface.raised', color: 'ink.0' }}
        >
          {open ? <LuX /> : <LuMenu />}
        </IconButton>
      </Drawer.Trigger>
      <Portal>
        <Drawer.Backdrop bg='rgba(8, 8, 10, 0.65)' backdropFilter='blur(4px)' />
        <Drawer.Positioner>
          <Drawer.Content
            bg='canvas.deep'
            borderLeftWidth='1px'
            borderColor='rule'
            maxW={{ base: '85vw', sm: '320px' }}
          >
            <Drawer.Header borderBottomWidth='1px' borderColor='rule' py={4}>
              <HStack justify='space-between' align='center'>
                <HStack gap={2}>
                  <Box w='6px' h='6px' borderRadius='full' bg='accent.fg' />
                  <Text fontFamily='mono' fontSize='sm' color='ink.0' fontWeight={500}>
                    davinci-dkg
                  </Text>
                </HStack>
                <Drawer.CloseTrigger asChild>
                  <IconButton
                    aria-label='Close menu'
                    variant='ghost'
                    size='xs'
                    color='ink.3'
                    _hover={{ color: 'ink.0' }}
                  >
                    <LuX />
                  </IconButton>
                </Drawer.CloseTrigger>
              </HStack>
            </Drawer.Header>

            <Drawer.Body py={4} px={3}>
              <Stack gap={1} as='nav'>
                {items.map((item) => (
                  <NavLink key={item.to} to={item.to} end={item.to === '/'}>
                    {({ isActive }) => (
                      <Box
                        position='relative'
                        px={3}
                        py={2.5}
                        borderRadius='md'
                        bg={isActive ? 'accent.bg' : 'transparent'}
                        transition='background 0.15s ease'
                        _hover={{ bg: isActive ? 'accent.bg' : 'surface.raised' }}
                      >
                        {/* Active rule on the left edge — quiet but unambiguous. */}
                        {isActive && (
                          <Box
                            position='absolute'
                            left={0}
                            top='25%'
                            bottom='25%'
                            w='2px'
                            bg='accent.fg'
                            borderRightRadius='full'
                          />
                        )}
                        <Text
                          fontSize='sm'
                          fontWeight={isActive ? 500 : 400}
                          color={isActive ? 'ink.0' : 'ink.2'}
                        >
                          {item.label}
                        </Text>
                        {item.hint && (
                          <Text fontSize='2xs' color='ink.4' mt={0.5}>
                            {item.hint}
                          </Text>
                        )}
                      </Box>
                    )}
                  </NavLink>
                ))}
              </Stack>
            </Drawer.Body>

            <Drawer.Footer
              borderTopWidth='1px'
              borderColor='rule'
              py={4}
              px={4}
              flexDirection='column'
              alignItems='stretch'
              gap={3}
            >
              <Box>
                <Text
                  fontFamily='mono'
                  fontSize='2xs'
                  color='ink.4'
                  letterSpacing='0.08em'
                  textTransform='uppercase'
                  mb={2}
                >
                  Network
                </Text>
                {/* Self-aligns inside the drawer's narrow column. */}
                <HStack>
                  <ChainPill />
                </HStack>
              </Box>

              <HStack
                as='button'
                onClick={toggleDebug}
                px={3}
                py={2}
                borderRadius='md'
                borderWidth='1px'
                borderColor={debug ? 'rgba(134, 239, 172, 0.30)' : 'border.subtle'}
                bg={debug ? 'live.bg' : 'transparent'}
                color={debug ? 'live.fg' : 'ink.2'}
                fontSize='xs'
                fontFamily='sans'
                fontWeight={500}
                cursor='pointer'
                justify='space-between'
                _hover={{ bg: debug ? 'live.bg' : 'surface.raised' }}
                transition='background 0.15s, color 0.15s'
              >
                <Text>Debug mode</Text>
                <Text fontFamily='mono' fontSize='2xs' textTransform='uppercase' letterSpacing='0.06em'>
                  {debug ? 'on' : 'off'}
                </Text>
              </HStack>
            </Drawer.Footer>
          </Drawer.Content>
        </Drawer.Positioner>
      </Portal>
    </Drawer.Root>
  )
}
