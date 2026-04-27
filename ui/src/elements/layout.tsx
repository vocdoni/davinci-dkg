import { Box, Container } from '@chakra-ui/react'
import { Outlet, ScrollRestoration } from 'react-router-dom'
import { Header } from '~components/Layout/Header'
import { Footer } from '~components/Layout/Footer'

// Public app shell. Sticky Header, scrollable content, Footer at the bottom
// of the page. Each route element renders inside the <Outlet/>.
export function Layout() {
  return (
    <Box minH='100vh' display='flex' flexDirection='column'>
      <Header />
      <Box as='main' flex='1'>
        <Container maxW='7xl' py={8}>
          <Outlet />
        </Container>
      </Box>
      <Footer />
      <ScrollRestoration />
    </Box>
  )
}
