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
  managerAddress: '0x92d324254ef12e4392d54c771b121cc976682340',
})`}
        </CodeBlock>
      </Section>

      <Section heading='Reading a epoch'>
        <Text fontSize='sm' color='ink.2'>
          Epoch identifiers are 12-byte values formed from a 4-byte chain prefix and an 8-byte
          nonce. Build one with <Code>buildEpochId</Code> or pass a known one as a hex string.
        </Text>
        <CodeBlock language='tsx'>
          {`import { buildEpochId, EpochPhase, roundStatusLabel } from '@vocdoni/davinci-dkg-sdk'

// Latest epoch on the chain.
const nonce  = await dkg.epochNonce()
const prefix = await dkg.roundPrefix()
const epochId = buildEpochId(prefix, nonce - 1n)

const epoch = await dkg.getEpoch(epochId)
console.log({
  status: roundStatusLabel(epoch.status),
  threshold: \`\${epoch.policy.threshold} of \${epoch.policy.committeeSize}\`,
  contributions: \`\${epoch.contributionCount} / \${epoch.policy.minValidContributions}\`,
})

// Once finalized, the public key is exposed on-chain:
if (epoch.status === EpochPhase.Live) {
  const pk = await dkg.getCollectivePublicKey(epochId)
  console.log('shared key', pk.x.toString(16), pk.y.toString(16))
}`}
        </CodeBlock>
      </Section>

      <Section heading='Watching new epochs'>
        <Text fontSize='sm' color='ink.2'>
          The monitor helpers wrap react-query–friendly polling and viem event logs. Use them in a
          long-running process or a backend job.
        </Text>
        <CodeBlock language='tsx'>
          {`import { watchNewEpochs, watchEpochLive, watchCiphertextSubmitted } from '@vocdoni/davinci-dkg-sdk'

const stop = watchNewEpochs(dkg, (epochId, organizer) => {
  console.log('new epoch', epochId, 'created by', organizer)
})

watchEpochLive(dkg, '0x82...0001', (collectivePublicKeyHash) => {
  console.log('epoch finalized; key hash:', collectivePublicKeyHash)
})

// Ciphertexts as they land. pokValid is the same check every committee
// node runs before releasing a partial decryption.
watchCiphertextSubmitted(dkg, '0x82...0001', ({ aid, ciphertextIndex, pokValid }) => {
  console.log('ciphertext', ciphertextIndex, 'for app', aid, pokValid ? 'ok' : 'INVALID PROOF')
})

// Later, to clean up:
stop()`}
        </CodeBlock>
      </Section>

      <Section heading='Encrypting for the committee'>
        <Text fontSize='sm' color='ink.2'>
          Once an epoch is Live, anyone can encrypt for it. ElGamal on BabyJubJub runs entirely
          client-side. Use <Code>encryptWithProof</Code>: besides the ciphertext it returns a
          Schnorr proof that you know the encryption randomness, bound to the epoch and
          application id. Every committee node verifies that proof before it releases a partial
          decryption, so a ciphertext submitted without one is never decrypted — that is what
          stops a copied <Code>C1</Code> from being used as a decryption oracle.
        </Text>
        <CodeBlock language='tsx'>
          {`import { encryptWithProof } from '@vocdoni/davinci-dkg-sdk'

const ZERO_AID = '0x' + '00'.repeat(32)   // bare epoch key; or a registered app's aid
const pk = await dkg.getCollectivePublicKey(epochId)

// 'message' must be a non-negative integer strictly below 2^50 (≈ 1.13e15).
// That's the upper bound the committee's BSGS dlog can recover; submitting
// anything larger leaves the ciphertext unrecoverable.
const { ciphertext, pok } = await encryptWithProof(epochId, ZERO_AID, 42n, [pk.x, pk.y])
// ciphertext = { c1: [x, y], c2: [x, y] } — both points on BabyJubJub
// pok        = { ax, ay, z }              — proof of knowledge of the randomness`}
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
          <Code>createEpoch</Code> (four policy fields, bounded by the deployment's{' '}
          <Code>getEpochBounds()</Code>), <Code>submitCiphertext</Code>, and{' '}
          <Code>abortEpoch</Code> (permissionless, but only accepted for a dead epoch whose
          selection or key-assembly deadline has passed). The contract assigns the ciphertext
          index; <Code>submitCiphertext</Code> waits for the receipt and returns it.
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
  managerAddress: '0x92d324254ef12e4392d54c771b121cc976682340',
})

// Verifies the proof locally, sends, waits for the receipt and reads the
// on-chain-assigned index from the CiphertextSubmitted event.
const { hash, ciphertextIndex } = await writer.submitCiphertext(
  epochId, ZERO_AID, ciphertext, pok,
)
console.log('submitted in', hash, 'as ciphertext #', ciphertextIndex)`}
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

const record = await waitForCombinedDecryption(dkg, epochId, ZERO_AID, ciphertextIndex, {
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
