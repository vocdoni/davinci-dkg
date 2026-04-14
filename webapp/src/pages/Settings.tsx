import {
  Box,
  Button,
  Heading,
  HStack,
  Input,
  SimpleGrid,
  Tag,
  Text,
  VStack,
  useToast,
} from '@chakra-ui/react';
import { useQueryClient } from '@tanstack/react-query';
import { useEffect, useState } from 'react';
import { HashCell } from '../components/HashCell';
import { getRpcOverride, loadBaseConfig, setRpcOverride } from '../lib/client';
import { useConfig } from '../lib/hooks';

function Row({ label, children }: { label: string; children: React.ReactNode }) {
  return (
    <Box>
      <Text fontSize="xs" textTransform="uppercase" color="gray.500" mb={1}>
        {label}
      </Text>
      <Box fontFamily="mono" fontSize="sm">
        {children}
      </Box>
    </Box>
  );
}

export function Settings() {
  const cfg = useConfig();
  const qc = useQueryClient();
  const toast = useToast();
  const [draft, setDraft] = useState('');
  const [override, setOverride] = useState<string | null>(null);
  const [defaultRpc, setDefaultRpc] = useState<string>('');

  useEffect(() => {
    setOverride(getRpcOverride());
    loadBaseConfig()
      .then((b) => setDefaultRpc(b.rpcUrl))
      .catch(() => undefined);
  }, []);

  useEffect(() => {
    if (cfg.data) setDraft(cfg.data.rpcUrl);
  }, [cfg.data?.rpcUrl]);

  const apply = async (url: string | null) => {
    setRpcOverride(url);
    setOverride(getRpcOverride());
    await qc.invalidateQueries();
    toast({
      title: url ? 'RPC endpoint updated' : 'RPC override cleared',
      description: url ?? `Using default: ${defaultRpc}`,
      status: 'success',
      duration: 2000,
      isClosable: true,
    });
  };

  if (!cfg.data) return <Text color="gray.400">Loading…</Text>;

  return (
    <VStack align="stretch" spacing={5}>
      <Heading size="lg">Settings</Heading>

      <Box bg="gray.800" p={5} borderRadius="md" borderWidth="1px" borderColor="gray.700">
        <Heading size="md" mb={4}>
          RPC endpoint
        </Heading>
        <VStack align="stretch" spacing={3}>
          <Box>
            <Text fontSize="xs" textTransform="uppercase" color="gray.500" mb={1}>
              Active
            </Text>
            <HStack>
              <Text fontFamily="mono" fontSize="sm" color="cyan.200">
                {cfg.data.rpcUrl}
              </Text>
              {override ? (
                <Tag size="sm" colorScheme="purple">
                  override
                </Tag>
              ) : (
                <Tag size="sm" colorScheme="gray">
                  default
                </Tag>
              )}
            </HStack>
            {defaultRpc && defaultRpc !== cfg.data.rpcUrl && (
              <Text fontSize="xs" color="gray.500" mt={1} fontFamily="mono">
                default from server: {defaultRpc}
              </Text>
            )}
          </Box>

          <HStack>
            <Input
              value={draft}
              onChange={(e) => setDraft(e.target.value)}
              placeholder="http://host:8545"
              aria-label="RPC endpoint URL"
              fontFamily="mono"
              fontSize="sm"
              bg="gray.900"
              borderColor="gray.700"
            />
            <Button
              colorScheme="cyan"
              onClick={() => apply(draft)}
              isDisabled={!draft || draft === cfg.data.rpcUrl}
            >
              Save
            </Button>
            <Button
              variant="outline"
              onClick={() => {
                setDraft(defaultRpc);
                apply(null);
              }}
              isDisabled={!override}
            >
              Reset
            </Button>
          </HStack>
          <Text fontSize="xs" color="gray.500">
            Stored in your browser (localStorage). Changes take effect immediately and only affect
            this browser.
          </Text>
        </VStack>
      </Box>

      <Box bg="gray.800" p={5} borderRadius="md" borderWidth="1px" borderColor="gray.700">
        <Heading size="md" mb={4}>
          Chain
        </Heading>
        <SimpleGrid columns={{ base: 1, md: 2 }} spacing={4}>
          <Row label="Chain">
            <Text>
              {cfg.data.chainName} ({cfg.data.chainId})
            </Text>
          </Row>
          <Row label="DKGManager">
            <HashCell value={cfg.data.managerAddress} full />
          </Row>
          <Row label="DKGRegistry">
            <HashCell value={cfg.data.registryAddress} full />
          </Row>
          {cfg.data.startBlock !== undefined && (
            <Row label="Start block">
              <Text>#{cfg.data.startBlock}</Text>
            </Row>
          )}
        </SimpleGrid>
      </Box>
    </VStack>
  );
}
