import { Box, Container, HStack, Link, Text } from '@chakra-ui/react'
import { LuHeart } from 'react-icons/lu'

export function Footer() {
  return (
    <Box as='footer' borderTopWidth='1px' borderColor='gray.800' mt={12} py={5} bg='gray.950'>
      <Container maxW='7xl'>
        <HStack justify='space-between' fontSize='xs' color='gray.500' wrap='wrap' gap={3}>
          <HStack gap={1.5} wrap='wrap'>
            <Text>Made with</Text>
            <Box as='span' color='red.400' display='inline-flex' alignItems='center'>
              <LuHeart fill='currentColor' />
            </Box>
            <Text>by</Text>
            <Link
              href='https://vocdoni.io'
              target='_blank'
              rel='noopener noreferrer'
              color='cyan.400'
              _hover={{ color: 'cyan.300', textDecoration: 'underline' }}
            >
              vocdoni
            </Link>
          </HStack>
          <HStack gap={4}>
            <Link
              href='https://github.com/vocdoni/davinci-dkg'
              target='_blank'
              rel='noopener noreferrer'
              color='gray.400'
              _hover={{ color: 'cyan.300' }}
            >
              github
            </Link>
          </HStack>
        </HStack>
      </Container>
    </Box>
  )
}
