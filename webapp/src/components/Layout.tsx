import { Box, Container, Flex, HStack, Heading, Link, Spacer, Text } from '@chakra-ui/react';
import { Link as RouterLink, useLocation } from 'react-router-dom';
import { useChainTip, useConfig } from '../lib/hooks';

const NAV = [
  { to: '/', label: 'Overview' },
  { to: '/rounds', label: 'Rounds' },
  { to: '/registry', label: 'Registry' },
  { to: '/settings', label: 'Settings' },
];

export function Layout({ children }: { children: React.ReactNode }) {
  const loc = useLocation();
  const cfg = useConfig();
  const tip = useChainTip();
  return (
    <Flex direction="column" minH="100vh">
      <Box bg="gray.900" borderBottom="1px solid" borderColor="gray.700" py={3}>
        <Container maxW="container.xl">
          <Flex align="center">
            <Heading size="md" color="cyan.300" fontFamily="mono">
              DAVINCI DKG Explorer
            </Heading>
            <HStack spacing={5} ml={10}>
              {NAV.map((n) => {
                const active =
                  n.to === '/' ? loc.pathname === '/' : loc.pathname.startsWith(n.to);
                return (
                  <Link
                    as={RouterLink}
                    key={n.to}
                    to={n.to}
                    color={active ? 'cyan.300' : 'gray.300'}
                    fontWeight={active ? 'semibold' : 'normal'}
                    _hover={{ color: 'cyan.200', textDecoration: 'none' }}
                  >
                    {n.label}
                  </Link>
                );
              })}
            </HStack>
            <Spacer />
            <HStack spacing={4} fontFamily="mono" fontSize="sm" color="gray.400">
              {cfg.data && (
                <Text>
                  chain{' '}
                  <Text as="span" color="gray.200">
                    {cfg.data.chainName} ({cfg.data.chainId})
                  </Text>
                </Text>
              )}
              {tip.data && (
                <Text>
                  block{' '}
                  <Text as="span" color="green.300">
                    #{tip.data.number.toString()}
                  </Text>
                </Text>
              )}
            </HStack>
          </Flex>
        </Container>
      </Box>
      <Box flex="1" py={6}>
        <Container maxW="container.xl">{children}</Container>
      </Box>
      <Box bg="gray.900" borderTop="1px solid" borderColor="gray.700" py={3}>
        <Container maxW="container.xl">
          <Text fontSize="xs" color="gray.500" textAlign="center" fontFamily="mono">
            DAVINCI DKG protocol explorer · reads on-chain state via JSON-RPC · 4s polling
          </Text>
        </Container>
      </Box>
    </Flex>
  );
}
