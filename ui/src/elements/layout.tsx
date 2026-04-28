import { Box, Container } from '@chakra-ui/react'
import { Outlet, ScrollRestoration } from 'react-router-dom'
import { Header } from '~components/Layout/Header'
import { Footer } from '~components/Layout/Footer'

// Centered, narrow page shell. Container caps at 5xl (~1024px) — wide
// enough for tables and the playground two-column layout, narrow enough
// that prose pages stay readable without an additional inner maxW.
export function Layout() {
  return (
    <Box minH='100vh' display='flex' flexDirection='column' bg='canvas'>
      <Header />
      <Box as='main' flex='1'>
        <Container maxW='5xl' px={{ base: 5, md: 6 }} pt={{ base: 8, md: 12 }} pb={{ base: 12, md: 16 }}>
          <Outlet />
        </Container>
      </Box>
      <Footer />
      <ScrollRestoration />
    </Box>
  )
}
