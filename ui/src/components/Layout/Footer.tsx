import { Box, Container, HStack, Link, Text } from '@chakra-ui/react'

// Editorial colophon: typeset rule above, mono on left (publication info),
// brand attribution centered, source link on right. The "♥" is a glyph
// drawn in the accent — looks like a printer's flourish, not an emoji.
export function Footer() {
  return (
    <Box as='footer' mt={{ base: 16, md: 24 }} pb={{ base: 6, md: 8 }}>
      <Container maxW='5xl' px={{ base: 5, md: 8 }}>
        <Box className='dkg-rule' mb={5} />
        <HStack
          justify='space-between'
          fontFamily='mono'
          fontSize='2xs'
          color='ink.3'
          letterSpacing='0.04em'
          wrap='wrap'
          gap={3}
        >
          <Text textTransform='uppercase'>davinci-dkg · explorer</Text>

          <HStack gap={1.5}>
            <Text>Made with</Text>
            <Box
              as='span'
              color='accent.fg'
             
              fontSize='sm'
              lineHeight='0'
              transform='translateY(2px)'
            >
              ♥
            </Box>
            <Text>by</Text>
            <Link
              href='https://vocdoni.io'
              target='_blank'
              rel='noopener noreferrer'
              color='ink.1'
              textTransform='uppercase'
              _hover={{ color: 'accent.fg' }}
              transition='color 0.15s'
            >
              vocdoni
            </Link>
          </HStack>

          <Link
            href='https://github.com/vocdoni/davinci-dkg'
            target='_blank'
            rel='noopener noreferrer'
            color='ink.3'
            textTransform='uppercase'
            _hover={{ color: 'ink.1' }}
            transition='color 0.15s'
          >
            github ↗
          </Link>
        </HStack>
      </Container>
    </Box>
  )
}
