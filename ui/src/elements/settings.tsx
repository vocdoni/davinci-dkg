import { useState } from 'react'
import { Box, Button, Field, Heading, HStack, Input, Stack, Text } from '@chakra-ui/react'
import { useConfig } from '~providers/ConfigProvider'
import { useDebugMode } from '~hooks/use-debug-mode'
import { readRpcOverride, writeRpcOverride } from '~lib/rpc-override'
import { HashCell } from '~components/ui/HashCell'

export function Settings() {
  const config = useConfig()
  const { enabled, setEnabled } = useDebugMode()
  const [rpcInput, setRpcInput] = useState<string>(() => readRpcOverride() ?? '')
  const baseRpc = config.rpcUrl
  const overrideActive = (readRpcOverride() ?? '') !== ''

  const onSave = () => {
    writeRpcOverride(rpcInput.trim() || null)
    // Cheapest correct way to apply: full reload. The DKG client and viem
    // PublicClient are constructed once per page load against the active
    // config, so swapping the URL in place would otherwise leave stale
    // connections in flight.
    window.location.reload()
  }

  const onReset = () => {
    writeRpcOverride(null)
    setRpcInput('')
    window.location.reload()
  }

  return (
    <Stack gap={8}>
      <Box>
        <Heading size='lg'>Settings</Heading>
        <Text color='gray.400' fontSize='sm' mt={1}>
          UI-only preferences. Changes are stored in this browser via <code>localStorage</code>; nothing
          is sent on-chain.
        </Text>
      </Box>

      {/* RPC override */}
      <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={5}>
        <Stack gap={3}>
          <Box>
            <Heading size='sm'>RPC endpoint</Heading>
            <Text fontSize='xs' color='gray.500' mt={1}>
              Override the JSON-RPC URL used for all read calls. Useful when the default endpoint is
              rate-limited or you want to point at a private node.
            </Text>
          </Box>
          <Field.Root>
            <Field.Label>JSON-RPC URL</Field.Label>
            <Input
              value={rpcInput}
              onChange={(e) => setRpcInput(e.target.value)}
              placeholder={baseRpc}
              fontFamily='mono'
              fontSize='sm'
            />
            <Field.HelperText>
              Default from <code>/config.json</code>: {baseRpc}
            </Field.HelperText>
          </Field.Root>
          <HStack>
            <Button colorPalette='cyan' size='sm' onClick={onSave}>
              Save and reload
            </Button>
            {overrideActive && (
              <Button variant='outline' size='sm' onClick={onReset}>
                Reset to default
              </Button>
            )}
          </HStack>
        </Stack>
      </Box>

      {/* Debug mode */}
      <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={5}>
        <Stack gap={3}>
          <Box>
            <Heading size='sm'>Debug mode</Heading>
            <Text fontSize='xs' color='gray.500' mt={1}>
              Auto-expand all "Show technical details" disclosures and surface raw protocol data
              (event args, transcript hashes, BabyJubJub coordinates) inline. Off by default.
            </Text>
          </Box>
          <HStack>
            <Button
              size='sm'
              colorPalette={enabled ? 'red' : 'cyan'}
              variant={enabled ? 'subtle' : 'solid'}
              onClick={() => setEnabled(!enabled)}
            >
              {enabled ? 'Disable debug mode' : 'Enable debug mode'}
            </Button>
            <Text fontSize='xs' color='gray.500'>
              Status: {enabled ? 'on' : 'off'}
            </Text>
          </HStack>
        </Stack>
      </Box>

      {/* Chain info */}
      <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={5}>
        <Stack gap={3}>
          <Heading size='sm'>Network</Heading>
          <KV label='Chain' value={`${config.chainName} (id ${config.chainId})`} />
          <KV label='DKG Manager' value={<HashCell value={config.managerAddress} head={6} tail={6} />} />
          {config.registryAddress && (
            <KV label='DKG Registry' value={<HashCell value={config.registryAddress} head={6} tail={6} />} />
          )}
          {config.startBlock != null && (
            <KV label='Start block' value={`#${config.startBlock}`} />
          )}
        </Stack>
      </Box>

      {/* About */}
      <Box borderWidth='1px' borderColor='gray.800' borderRadius='md' bg='gray.900' p={5}>
        <Stack gap={1}>
          <Heading size='sm'>About</Heading>
          <Text fontSize='xs' color='gray.500'>
            build {(import.meta.env.VITE_BUILD_VERSION as string | undefined) ?? 'dev'}
          </Text>
        </Stack>
      </Box>
    </Stack>
  )
}

function KV({ label, value }: { label: string; value: React.ReactNode }) {
  return (
    <HStack gap={3} fontSize='sm'>
      <Text color='gray.500' minW='32'>
        {label}
      </Text>
      <Box>{value}</Box>
    </HStack>
  )
}
