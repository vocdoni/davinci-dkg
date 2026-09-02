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

      <Section heading='Reading an epoch'>
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

// Ciphertexts as they land. c1/c2 come back in TE form, ready to hand
// straight to writer.submitOrganizerShare.
watchCiphertextSubmitted(dkg, '0x82...0001', ({ aid, ciphertextIndex, c1, c2 }) => {
  console.log('ciphertext', ciphertextIndex, 'for app', aid, c1, c2)
})

// Later, to clean up:
stop()`}
        </CodeBlock>
      </Section>

      <Section heading='Registering an application'>
        <Text fontSize='sm' color='ink.2'>
          Ciphertexts always belong to a registered application; there is no bare epoch-key
          path. Registration binds an organizer key so the encryption key becomes{' '}
          <Code>PK_aid = PK_ep + PK_org</Code>: opening a ciphertext then needs both the
          committee threshold and the organizer's share. Only <Code>PK_org</Code> and a Schnorr
          proof of possession go on chain — <Code>sk_org</Code> never leaves the browser.
        </Text>
        <CodeBlock language='tsx'>
          {`import { randomOrganizerSecret } from '@vocdoni/davinci-dkg-sdk'

const skOrg = randomOrganizerSecret()   // store this before you register
const aid   = '0x…'                     // your bytes32 application id

await writer.registerApplication(epochId, aid, {
  authorizedSubmitter: '0x0000000000000000000000000000000000000000', // = the caller
  maxCiphertexts:      0,     // 0 = unlimited
  notBeforeBlock:      0n,
  notAfterBlock:       0n,
}, skOrg)`}
        </CodeBlock>
        <Text fontSize='sm' color='danger.fg'>
          <strong>Keep sk_org.</strong> It is the application's only decryption capability, it is
          never transmitted, and nothing on chain can reconstruct it. Lose it and every
          ciphertext ever submitted under that <Code>aid</Code> is permanently undecryptable —
          the committee threshold alone cannot open them, by design.
        </Text>
      </Section>

      <Section heading='Encrypting for an application'>
        <Text fontSize='sm' color='ink.2'>
          ElGamal on BabyJubJub runs entirely client-side.{' '}
          <Code>encryptForApplication</Code> takes the epoch key and the application's{' '}
          <Code>PK_org</Code> (both in TE form, both readable from the client) and encrypts
          under their sum.
        </Text>
        <CodeBlock language='tsx'>
          {`import { encryptForApplication } from '@vocdoni/davinci-dkg-sdk'

const pk  = await dkg.getCollectivePublicKey(epochId)
const app = await dkg.getApplication(epochId, aid)

// 'message' must be a non-negative integer strictly below 2^50 (≈ 1.13e15).
// That's the upper bound the committee's BSGS dlog can recover; submitting
// anything larger leaves the ciphertext unrecoverable.
const ciphertext = await encryptForApplication(42n, [pk.x, pk.y], app.organizerPK)
// ciphertext = { c1: [x, y], c2: [x, y] } — both points on BabyJubJub`}
        </CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          There is no proof of knowledge of the encryption randomness. The submitter of an
          aggregated tally cannot know its randomness, so such a proof is incompatible with
          homomorphic aggregation; a <Code>C1</Code> copied into another application and
          decrypted there only yields <Code>sk_ep·C1</Code>, useless without that application's{' '}
          <Code>sk_org·C1</Code>.
        </Text>
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

// Sends, waits for the receipt and reads the on-chain-assigned index from
// the CiphertextSubmitted event.
const { hash, ciphertextIndex } = await writer.submitCiphertext(epochId, aid, ciphertext)
console.log('submitted in', hash, 'as ciphertext #', ciphertextIndex)`}
        </CodeBlock>
      </Section>

      <Section heading='Releasing the organizer share'>
        <Text fontSize='sm' color='ink.2'>
          The committee's partial decryptions are not enough on their own. The organizer
          publishes <Code>Δ = sk_org·C1</Code> with a Chaum–Pedersen DLEQ proving the same
          secret relates <Code>(G, PK_org)</Code> and <Code>(C1, Δ)</Code>. Until that lands,{' '}
          <Code>combineDecryption</Code> reverts <Code>OrganizerShareMissing()</Code>.
        </Text>
        <CodeBlock language='tsx'>
          {`await writer.submitOrganizerShare(epochId, aid, ciphertextIndex, ciphertext, skOrg)

// The challenge is a keccak256 the contract recomputes and the combine
// circuit consumes:
//   e = keccak256(DOMAIN_ORGANIZER_SHARE_V1 ‖ eid ‖ aid ‖ uint256(ctIdx)
//                 ‖ PK_org ‖ C1 ‖ Δ ‖ A1 ‖ A2) mod q
//   z = w + e·sk_org mod q
// The contract stores only keccak256(Δ ‖ A1 ‖ A2 ‖ z); the DLEQ itself is
// verified inside the committee's combine proof.`}
        </CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          Anyone may relay a share — it is self-authenticating — and re-submission overwrites
          until the ciphertext is combined, so a malformed share cannot brick a ciphertext.
        </Text>
      </Section>

      <Section heading='Awaiting committee decryption'>
        <Text fontSize='sm' color='ink.2'>
          With the threshold of partial decryptions and the organizer share both on chain, any
          node combines them and publishes the plaintext. The flow helper polls the contract
          until <Code>completed === true</Code> and returns the recovered value;{' '}
          <Code>decryptionProgress</Code> tells you which of the two halves is still missing.
        </Text>
        <CodeBlock language='tsx'>
          {`import { waitForCombinedDecryption, decryptionProgress } from '@vocdoni/davinci-dkg-sdk'

const progress = await decryptionProgress(dkg, epochId, aid, ciphertextIndex)
console.log(progress.organizerShare ? 'waiting on the committee' : 'waiting on the organizer')

const record = await waitForCombinedDecryption(dkg, epochId, aid, ciphertextIndex, {
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
