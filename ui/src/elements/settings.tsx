import { useState, type ReactNode } from 'react'
import { Box, Button, Field, Grid, GridItem, Heading, HStack, Input, Stack, Text } from '@chakra-ui/react'
import { useConfig } from '~providers/ConfigProvider'
import { useDebugMode } from '~hooks/use-debug-mode'
import { readRpcOverride, writeRpcOverride } from '~lib/rpc-override'
import { HashCell } from '~components/ui/HashCell'
import { PageHeader } from '~components/Layout/PageHeader'

export function Settings() {
  const config = useConfig()
  const { enabled, setEnabled } = useDebugMode()
  const [rpcInput, setRpcInput] = useState<string>(() => readRpcOverride() ?? '')
  const baseRpc = config.rpcUrl
  const overrideActive = (readRpcOverride() ?? '') !== ''

  const onSave = () => {
    writeRpcOverride(rpcInput.trim() || null)
    window.location.reload()
  }

  const onReset = () => {
    writeRpcOverride(null)
    setRpcInput('')
    window.location.reload()
  }

  return (
    <Stack gap={{ base: 10, md: 14 }}>
      <PageHeader
        title='Settings'
        subtitle={
          <>
            UI-only preferences. Changes are stored in this browser via{' '}
            <Box as='code' fontFamily='mono' fontSize='0.92em' color='ink.2'>
              localStorage
            </Box>
            ; nothing is sent on-chain.
          </>
        }
      />

      <Section title='RPC endpoint'>
        <Text fontSize='sm' color='ink.2' lineHeight='1.6' mb={5} maxW='62ch'>
          Override the JSON-RPC URL used for all read calls. Useful when the default endpoint is
          rate-limited or you want to point at a private node.
        </Text>
        <Field.Root>
          <Field.Label fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.08em' textTransform='uppercase'>
            JSON-RPC URL
          </Field.Label>
          <Input
            value={rpcInput}
            onChange={(e) => setRpcInput(e.target.value)}
            placeholder={baseRpc}
            fontFamily='mono'
            fontSize='sm'
            bg='surface.sunken'
            borderColor='border.subtle'
            _hover={{ borderColor: 'border' }}
            _focus={{ borderColor: 'accent.fg', boxShadow: '0 0 0 1px var(--chakra-colors-accent-fg)' }}
            mt={1.5}
          />
          <Field.HelperText fontSize='xs' color='ink.3' mt={2}>
            Default from <Box as='code' fontFamily='mono' color='ink.2'>/config.json</Box>: {baseRpc}
          </Field.HelperText>
        </Field.Root>
        <HStack mt={5} gap={2}>
          <Button
            size='sm'
            onClick={onSave}
            bg='accent.fg'
            color='canvas.deep'
            fontFamily='sans'
            fontWeight={500}
            _hover={{ bg: 'accent.bright' }}
          >
            Save and reload
          </Button>
          {overrideActive && (
            <Button
              variant='outline'
              size='sm'
              onClick={onReset}
              borderColor='border'
              color='ink.2'
              fontFamily='sans'
              _hover={{ bg: 'surface.raised', color: 'ink.0' }}
            >
              Reset to default
            </Button>
          )}
        </HStack>
      </Section>

      <Section title='Debug mode'>
        <Text fontSize='sm' color='ink.2' lineHeight='1.6' mb={5} maxW='62ch'>
          Auto-expand all "Show technical details" disclosures and surface raw protocol data —
          event args, transcript hashes, BabyJubJub coordinates — inline. Off by default.
        </Text>
        <HStack gap={3}>
          <Button
            size='sm'
            onClick={() => setEnabled(!enabled)}
            bg={enabled ? 'transparent' : 'live.fg'}
            color={enabled ? 'live.fg' : 'canvas.deep'}
            borderWidth='1px'
            borderColor={enabled ? 'rgba(134, 239, 172, 0.40)' : 'transparent'}
            fontFamily='sans'
            fontWeight={500}
            _hover={{ bg: enabled ? 'live.bg' : 'live.bright' }}
          >
            {enabled ? 'Disable debug mode' : 'Enable debug mode'}
          </Button>
          <Text fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.06em' textTransform='uppercase'>
            Status: {enabled ? 'On' : 'Off'}
          </Text>
        </HStack>
      </Section>

      <Section title='Connected chain'>
        <Stack gap={4} mt={1}>
          <KV label='Chain' value={`${config.chainName} · id ${config.chainId}`} />
          <KV label='DKG manager' value={<HashCell value={config.managerAddress} head={6} tail={6} />} />
          {config.registryAddress && (
            <KV label='DKG registry' value={<HashCell value={config.registryAddress} head={6} tail={6} />} />
          )}
          {config.startBlock != null && <KV label='Start block' value={`#${config.startBlock}`} />}
        </Stack>
      </Section>

      <Section title='About'>
        <Text fontFamily='mono' fontSize='xs' color='ink.3'>
          build{' '}
          <Box as='span' color='ink.1'>
            {(import.meta.env.VITE_BUILD_VERSION as string | undefined) ?? 'dev'}
          </Box>
        </Text>
      </Section>
    </Stack>
  )
}

function Section({ title, children }: { title: string; children: ReactNode }) {
  return (
    <Box as='section'>
      <Heading
        as='h2'
        fontSize={{ base: 'lg', md: 'xl' }}
        fontWeight={500}
        color='ink.0'
        letterSpacing='-0.01em'
        mb={4}
      >
        {title}
      </Heading>
      {children}
    </Box>
  )
}

function KV({ label, value }: { label: string; value: ReactNode }) {
  return (
    <Grid templateColumns={{ base: '120px 1fr', md: '160px 1fr' }} gap={4} alignItems='baseline'>
      <GridItem>
        <Text fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.08em' textTransform='uppercase'>
          {label}
        </Text>
      </GridItem>
      <GridItem fontSize='sm' color='ink.0'>
        {value}
      </GridItem>
    </Grid>
  )
}
