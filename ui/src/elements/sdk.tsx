import { Badge, Box, Heading, HStack, Link, List, Stack, Text } from '@chakra-ui/react'
import { LuExternalLink } from 'react-icons/lu'
import type { ReactNode } from 'react'
import { CodeBlock } from '~components/ui/CodeBlock'

// Developer-facing "use the SDK" page. Mirrors the structure of run-a-node:
// a numbered/lettered sequence of sections with minimal code snippets.
//
// Style: documentation. Each snippet is the minimum to be useful — no
// boilerplate, no decorative chrome.
export function Sdk() {
  return (
    <Stack gap={10} maxW='3xl'>
      <Box pb={4} borderBottomWidth='1px' borderColor='gray.800'>
        <HStack gap={3} mb={2} align='baseline' wrap='wrap'>
          <Heading size='lg' color='gray.100'>
            SDK reference
          </Heading>
          <Badge colorPalette='gray' variant='outline' fontSize='2xs'>
            typescript
          </Badge>
        </HStack>
        <Text color='gray.400' fontSize='sm' lineHeight='1.6'>
          <Code>@vocdoni/davinci-dkg-sdk</Code> is a TypeScript wrapper around the on-chain
          contracts and the BabyJubJub cryptography. It runs entirely against any EVM JSON-RPC
          endpoint; no extra services in between.
        </Text>
      </Box>

      <Section heading='Install'>
        <Text fontSize='sm' color='gray.300'>
          ESM-only TypeScript package. Depends on <Code>viem</Code> and{' '}
          <Code>@zk-kit/baby-jubjub</Code>.
        </Text>
        <CodeBlock language='bash'>{`pnpm add @vocdoni/davinci-dkg-sdk viem
# or: npm install @vocdoni/davinci-dkg-sdk viem
# or: yarn add @vocdoni/davinci-dkg-sdk viem`}</CodeBlock>
        <Text fontSize='sm' color='gray.300'>
          Until the package is published to npm, the monorepo can be consumed directly via a
          relative file dependency (this is what the explorer in this page does):
        </Text>
        <CodeBlock language='json'>{`"@vocdoni/davinci-dkg-sdk": "file:../davinci-dkg/sdk"`}</CodeBlock>
      </Section>

      <Section heading='Read-only client'>
        <Text fontSize='sm' color='gray.300'>
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
  managerAddress: '0xd3ef727b695b21e108497c36f9dcec52d741298a',
})`}
        </CodeBlock>
      </Section>

      <Section heading='Reading a round'>
        <Text fontSize='sm' color='gray.300'>
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
        <Text fontSize='sm' color='gray.300'>
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
        <Text fontSize='sm' color='gray.300'>
          Once a round is finalized, anyone can encrypt for it. ElGamal on BabyJubJub is provided
          by <Code>buildElGamal</Code> and operates entirely client-side.
        </Text>
        <CodeBlock language='tsx'>
          {`import { buildElGamal } from '@vocdoni/davinci-dkg-sdk'

const eg = await buildElGamal()
const pk = await dkg.getCollectivePublicKey(roundId)

// 'message' is a small integer in the BabyJubJub scalar field (here, < 2^20
// so the brute-force decoder can recover it cheaply on-chain).
const ciphertext = eg.encrypt(42n, [pk.x, pk.y])
// ciphertext = { c1: [x, y], c2: [x, y] } — both points on BabyJubJub`}
        </CodeBlock>
      </Section>

      <Section heading='Submitting a ciphertext'>
        <Text fontSize='sm' color='gray.300'>
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
  managerAddress: '0xd3ef727b695b21e108497c36f9dcec52d741298a',
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
        <Text fontSize='sm' color='gray.300'>
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
      <Heading size='md' color='gray.100' mb={3}>
        {heading}
      </Heading>
      <Stack gap={3}>{children}</Stack>
    </Box>
  )
}

function Code({ children }: { children: ReactNode }) {
  return (
    <Box
      as='code'
      bg='gray.800'
      px={1.5}
      py={0.5}
      borderRadius='sm'
      fontFamily='mono'
      fontSize='xs'
      color='gray.100'
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
      color='cyan.300'
      _hover={{ color: 'cyan.200', textDecoration: 'underline' }}
    >
      <HStack gap={1.5} display='inline-flex' align='center'>
        <Box as='span'>{children}</Box>
        <LuExternalLink />
      </HStack>
    </Link>
  )
}
