import { Box, Heading, HStack, Link, List, Stack, Text } from '@chakra-ui/react'
import { LuExternalLink } from 'react-icons/lu'
import type { ReactNode } from 'react'
import { CodeBlock } from '~components/ui/CodeBlock'
import { PageHeader } from '~components/Layout/PageHeader'

export function Sdk() {
  return (
    <Stack gap={{ base: 10, md: 14 }}>
      <PageHeader
        title='SDK reference'
        subtitle={
          <>
            <Code>@vocdoni/davinci-dkg-sdk</Code> is a TypeScript wrapper around the on-chain
            contracts and the BabyJubJub cryptography. It runs entirely against any EVM JSON-RPC
            endpoint; no extra services in between.
          </>
        }
        action={
          <HStack gap={2} fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.08em' textTransform='uppercase'>
            <Box w='6px' h='6px' borderRadius='full' bg='accent.fg' />
            <Text color='ink.2'>TypeScript</Text>
          </HStack>
        }
      />

      <Section heading='Install'>
        <Text fontSize='sm' color='ink.2'>
          ESM-only TypeScript package. Depends on <Code>viem</Code> and{' '}
          <Code>@zk-kit/baby-jubjub</Code>.
        </Text>
        <CodeBlock language='bash'>{`pnpm add @vocdoni/davinci-dkg-sdk viem
# or: npm install @vocdoni/davinci-dkg-sdk viem
# or: yarn add @vocdoni/davinci-dkg-sdk viem`}</CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          Until the package is published to npm, the monorepo can be consumed directly via a
          relative file dependency (this is what the explorer in this page does):
        </Text>
        <CodeBlock language='json'>{`"@vocdoni/davinci-dkg-sdk": "file:../davinci-dkg/sdk"`}</CodeBlock>
      </Section>

      <Section heading='Read-only client'>
        <Text fontSize='sm' color='ink.2'>
          All read calls go through <Code>DKGClient</Code>. It needs a viem{' '}
          <Code>PublicClient</Code> and the address of the deployed <Code>DKGManager</Code>{' '}
          contract; the registry address is auto-discovered if omitted.
        </Text>
        <CodeBlock caption='client.ts' language='tsx'>
          {`import { createPublicClient, http } from 'viem'
import { sepolia } from 'viem/chains'
import { DKGClient } from '@vocdoni/davinci-dkg-sdk'

const publicClient = createPublicClient({
  chain: sepolia,
  transport: http('https://eth-sepolia.public.blastapi.io'),
})

export const dkg = new DKGClient({
  publicClient,
  managerAddress: '0x01ee71fdce1705c8823f9f8b2f312100165fdd70',
})`}
        </CodeBlock>
      </Section>

      <Section heading='Reading a round'>
        <Text fontSize='sm' color='ink.2'>
          Round identifiers are 12-byte values formed from a 4-byte chain prefix and an 8-byte
          nonce. Build one with <Code>buildRoundId</Code> or pass a known one as a hex string.
        </Text>
        <CodeBlock language='tsx'>
          {`import { buildRoundId, RoundStatus, roundStatusLabel } from '@vocdoni/davinci-dkg-sdk'

// Latest round on the chain.
const nonce  = await dkg.roundNonce()
const prefix = await dkg.roundPrefix()
const roundId = buildRoundId(prefix, nonce - 1n)

const round = await dkg.getRound(roundId)
console.log({
  status: roundStatusLabel(round.status),
  threshold: \`\${round.policy.threshold} of \${round.policy.committeeSize}\`,
  contributions: \`\${round.contributionCount} / \${round.policy.minValidContributions}\`,
})

// Once finalized, the public key is exposed on-chain:
if (round.status === RoundStatus.Finalized) {
  const pk = await dkg.getCollectivePublicKey(roundId)
  console.log('shared key', pk.x.toString(16), pk.y.toString(16))
}`}
        </CodeBlock>
      </Section>

      <Section heading='Watching new rounds'>
        <Text fontSize='sm' color='ink.2'>
          The monitor helpers wrap react-query–friendly polling and viem event logs. Use them in a
          long-running process or a backend job.
        </Text>
        <CodeBlock language='tsx'>
          {`import { watchNewRounds, watchRoundFinalized } from '@vocdoni/davinci-dkg-sdk'

const stop = watchNewRounds(dkg, (entry) => {
  console.log('new round', entry.id, entry.round.policy.threshold,
              'of', entry.round.policy.committeeSize)
})

watchRoundFinalized(dkg, '0x82...0001', (event) => {
  console.log('round finalized; key hash:', event.collectivePublicKeyHash)
})

// Later, to clean up:
stop()`}
        </CodeBlock>
      </Section>

      <Section heading='Encrypting for the committee'>
        <Text fontSize='sm' color='ink.2'>
          Once a round is finalized, anyone can encrypt for it. ElGamal on BabyJubJub is provided
          by <Code>buildElGamal</Code> and operates entirely client-side.
        </Text>
        <CodeBlock language='tsx'>
          {`import { buildElGamal } from '@vocdoni/davinci-dkg-sdk'

const eg = await buildElGamal()
const pk = await dkg.getCollectivePublicKey(roundId)

// 'message' must be a non-negative integer strictly below 2^50 (≈ 1.13e15).
// That's the upper bound the committee's BSGS dlog can recover; submitting
// anything larger leaves the round unrecoverable.
const ciphertext = eg.encrypt(42n, [pk.x, pk.y])
// ciphertext = { c1: [x, y], c2: [x, y] } — both points on BabyJubJub`}
        </CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          The matching client-side <Code>decrypt(ct, privKey)</Code> helper (used in tests
          and direct-key recovery, not in the threshold flow) caps at 2<sup>32</sup> ≈ 4.3
          billion — the SDK's table fits in ~16 MB so it stays browser-friendly. The
          on-chain threshold path always uses the committee's higher 2<sup>50</sup> cap.
        </Text>
      </Section>

      <Section heading='Submitting a ciphertext'>
        <Text fontSize='sm' color='ink.2'>
          Chain-writing operations require a viem <Code>WalletClient</Code>. Wrap it with{' '}
          <Code>DKGWriter</Code>, which extends <Code>DKGClient</Code> with{' '}
          <Code>createRound</Code>, <Code>submitCiphertext</Code>, and <Code>abortRound</Code>.
        </Text>
        <CodeBlock language='tsx'>
          {`import { createWalletClient, http } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { sepolia } from 'viem/chains'
import { DKGWriter } from '@vocdoni/davinci-dkg-sdk'

const account = privateKeyToAccount('0x<your-private-key>')
const walletClient = createWalletClient({
  account,
  chain: sepolia,
  transport: http('https://eth-sepolia.public.blastapi.io'),
})

const writer = new DKGWriter({
  publicClient,
  walletClient,
  managerAddress: '0x01ee71fdce1705c8823f9f8b2f312100165fdd70',
})

const txHash = await writer.submitCiphertext(
  roundId, /* ciphertextIndex */ 1,
  ciphertext.c1[0], ciphertext.c1[1],
  ciphertext.c2[0], ciphertext.c2[1],
)
await writer.waitForTransaction(txHash)`}
        </CodeBlock>
      </Section>

      <Section heading='Awaiting committee decryption'>
        <Text fontSize='sm' color='ink.2'>
          After submission, the committee picks up the new ciphertext, posts partial decryptions,
          and combines them into the final plaintext on-chain. The flow helper polls the contract
          until <Code>completed === true</Code> and returns the recovered value.
        </Text>
        <CodeBlock language='tsx'>
          {`import { waitForCombinedDecryption } from '@vocdoni/davinci-dkg-sdk'

const record = await waitForCombinedDecryption(dkg, roundId, 1, {
  intervalMs: 3000,
  timeoutMs: 5 * 60_000,
})
if (record.completed) {
  console.log('plaintext:', record.plaintext.toString())
}`}
        </CodeBlock>
      </Section>

      <Section heading='References'>
        <List.Root gap={1.5} fontSize='sm' pl={5}>
          <List.Item>
            <ExternalLink href='https://github.com/vocdoni/davinci-dkg/tree/main/sdk'>
              SDK source — every export documented inline.
            </ExternalLink>
          </List.Item>
          <List.Item>
            <ExternalLink href='https://github.com/vocdoni/davinci-dkg/tree/main/sdk/tests'>
              SDK tests — copy-pasteable end-to-end usage examples.
            </ExternalLink>
          </List.Item>
          <List.Item>
            <ExternalLink href='https://github.com/vocdoni/davinci-dkg#readme'>
              davinci-dkg README — protocol overview and contract reference.
            </ExternalLink>
          </List.Item>
          <List.Item>
            <ExternalLink href='https://viem.sh/'>viem documentation.</ExternalLink>
          </List.Item>
        </List.Root>
      </Section>
    </Stack>
  )
}

function Section({ heading, children }: { heading: string; children: ReactNode }) {
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
        {heading}
      </Heading>
      <Stack gap={4}>{children}</Stack>
    </Box>
  )
}

function Code({ children }: { children: ReactNode }) {
  return (
    <Box
      as='code'
      bg='surface.sunken'
      borderWidth='1px'
      borderColor='border.subtle'
      px='0.4em'
      py='0.1em'
      borderRadius='sm'
      fontFamily='mono'
      fontSize='0.86em'
      color='accent.bright'
    >
      {children}
    </Box>
  )
}

function ExternalLink({ href, children }: { href: string; children: ReactNode }) {
  return (
    <Link
      href={href}
      target='_blank'
      rel='noopener noreferrer'
      color='accent.fg'
      borderBottomWidth='1px'
      borderColor='accent.border'
      _hover={{ color: 'accent.bright', borderColor: 'accent.fg' }}
      transition='color 0.15s, border-color 0.15s'
    >
      <HStack gap={1.5} display='inline-flex' align='center'>
        <Box as='span'>{children}</Box>
        <LuExternalLink />
      </HStack>
    </Link>
  )
}
