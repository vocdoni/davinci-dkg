import { Box, Heading, HStack, Link, List, Stack, Text } from '@chakra-ui/react'
import { LuExternalLink } from 'react-icons/lu'
import type { ReactNode } from 'react'
import { CodeBlock } from '~components/ui/CodeBlock'
import { PageHeader } from '~components/Layout/PageHeader'

export function RunNode() {
  return (
    <Stack gap={{ base: 10, md: 14 }}>
      <PageHeader
        title='Running a node'
        subtitle='Procedure for joining the davinci-dkg committee on Sepolia. The node binary ships as a single Docker image; no Go, Node, or build toolchain is required on the host. The explorer UI is a separate, optional container — see step 4.'
        action={
          <HStack gap={2} fontFamily='mono' fontSize='2xs' color='ink.3' letterSpacing='0.08em' textTransform='uppercase'>
            <Box w='6px' h='6px' borderRadius='full' bg='accent.fg' />
            <Text color='ink.2'>Sepolia</Text>
          </HStack>
        }
      />

      <Section heading='Prerequisites'>
        <List.Root gap={2} fontSize='sm' color='ink.2' pl={5}>
          <List.Item>
            <Code>docker</Code> ≥ 24 and <Code>docker compose</Code> v2.
          </List.Item>
          <List.Item>
            An Ethereum private key with a small amount of <strong>Sepolia ETH</strong> for
            transaction fees (registry registration, slot claims, contributions, partial
            decryptions). Two faucets:
            <List.Root gap={1} pl={5} mt={1.5} listStyle='disc'>
              <List.Item>
                <ExternalLink href='https://sepolia-faucet.pk910.de/'>
                  sepolia-faucet.pk910.de
                </ExternalLink>{' '}
                — proof-of-work mining faucet.
              </List.Item>
              <List.Item>
                <ExternalLink href='https://cloud.google.com/application/web3/faucet/ethereum/sepolia'>
                  Google Cloud Sepolia faucet
                </ExternalLink>
                .
              </List.Item>
            </List.Root>
          </List.Item>
          <List.Item>
            A Sepolia JSON-RPC endpoint. Public providers work for testing; for a long-lived node a
            dedicated provider (Alchemy, Infura, QuickNode, your own node) is more reliable.
          </List.Item>
        </List.Root>
      </Section>

      <Section heading='1. Clone the repository'>
        <Text fontSize='sm' color='ink.2'>
          The compose file and the example env live at the repository root. Nothing is built locally
          — everything is pulled from the published images.
        </Text>
        <CodeBlock language='bash'>
          {`git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg`}
        </CodeBlock>
      </Section>

      <Section heading='2. Configure the node'>
        <Text fontSize='sm' color='ink.2'>
          Copy the example env file and fill in your private key. The Sepolia network preset
          resolves all contract addresses for you.
        </Text>
        <CodeBlock language='bash'>{`cp .env.example .env
$EDITOR .env`}</CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          Minimum required entries:
        </Text>
        <CodeBlock caption='.env' language='bash'>
          {`# Sepolia RPC (replace with your own provider for production).
DAVINCI_DKG_WEB3_RPC=https://eth-sepolia.public.blastapi.io

# Hex private key (must start with 0x). This wallet pays node tx fees and
# is your operator identity in the registry.
DAVINCI_DKG_PRIVKEY=0x<your-private-key>

# Use the well-known Sepolia deployment — no addresses needed.
DAVINCI_DKG_NETWORK=sepolia

# Optional: bind the embedded explorer to a non-default port.
# DAVINCI_DKG_WEBAPP_LISTEN=0.0.0.0:8081

# Optional: tell the embedded explorer to advertise a host-reachable RPC
# URL when accessed from a different machine than the node.
# DAVINCI_DKG_WEBAPP_PUBLIC_RPC=https://your-rpc.example`}
        </CodeBlock>
      </Section>

      <Section heading='3. Start the node'>
        <Text fontSize='sm' color='ink.2'>
          One command brings up the node and Watchtower (which auto-updates the image when new
          versions are published). The container restarts on failure by default.
        </Text>
        <CodeBlock language='bash'>{`docker compose --profile node up -d`}</CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          Tail the logs to confirm it is running:
        </Text>
        <CodeBlock language='bash'>{`docker compose --profile node logs -f node`}</CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          On first boot the node:
        </Text>
        <List.Root gap={1.5} fontSize='sm' color='ink.2' pl={5}>
          <List.Item>Derives a BabyJubJub encryption key from your Ethereum private key.</List.Item>
          <List.Item>
            Registers in <Code>DKGRegistry</Code> (one transaction) — or skips if it is already
            registered and active.
          </List.Item>
          <List.Item>
            Polls the chain for round events at the configured interval (default: 20 seconds).
          </List.Item>
        </List.Root>
      </Section>

      <Section heading='4. (Optional) Host an explorer alongside it'>
        <Text fontSize='sm' color='ink.2'>
          The node binary does not serve any HTTP — it only talks to the chain. To browse rounds,
          the registry, and the playground from a wallet, host the explorer separately. The
          explorer is a Vite static bundle whose chain config is baked in at build time.
        </Text>
        <Text fontSize='sm' color='ink.2'>
          The compose <Code>ui</Code> profile bind-mounts the locally-built bundle into stock{' '}
          <Code>nginx:alpine</Code>:
        </Text>
        <CodeBlock language='bash'>
          {`# 1. Build the bundle with the chain config you want.
make ui-build \\
  RPC_URL=https://eth-sepolia.public.blastapi.io \\
  MANAGER_ADDRESS=0x01ee71fdce1705c8823f9f8b2f312100165fdd70 \\
  CHAIN_ID=11155111 CHAIN_NAME=sepolia

# 2. Run nginx alongside the node.
docker compose --profile node --profile ui up -d
# explorer reachable at http://<your-host>:8082/`}
        </CodeBlock>
        <Text fontSize='sm' color='ink.2'>
          For a public deployment, point DigitalOcean App Platform at the spec checked into the
          repo (<Code>ui/.do/davinci-dkg-ui.yaml</Code>) and edit the <Code>BUILD_TIME</Code> env
          values to retarget the chain. App Platform builds the Dockerfile per push and serves
          the static files from its edge — no nginx in the loop.
        </Text>
        <CodeBlock language='bash'>{`doctl apps create --spec ui/.do/davinci-dkg-ui.yaml`}</CodeBlock>
      </Section>

      <Section heading='5. Maintenance'>
        <List.Root gap={2.5} fontSize='sm' color='ink.2' pl={5}>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='ink.0'>
              Upgrades.
            </Text>{' '}
            Watchtower follows the <Code>latest</Code> tag and recreates the container when a new
            version ships. Pin <Code>DAVINCI_DKG_TAG=v0.1.0</Code> in <Code>.env</Code> (or remove
            the watchtower service from <Code>docker-compose.yml</Code>) for manual control.
          </List.Item>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='ink.0'>
              State.
            </Text>{' '}
            The node mounts a Docker volume named <Code>run</Code> for its data directory. All
            per-round state is rebuilt from on-chain records on restart, so a stop/start mid-round
            is safe.
          </List.Item>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='ink.0'>
              Stop.
            </Text>{' '}
            <Code>docker compose --profile node down</Code>. Add <Code>-v</Code> to also wipe the
            volume.
          </List.Item>
        </List.Root>
      </Section>

      <Section heading='Limits'>
        <List.Root gap={2.5} fontSize='sm' color='ink.2' pl={5}>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='ink.0'>
              Committee size.
            </Text>{' '}
            Capped at <Code>32</Code> per round (the circuits' compile-time <Code>MaxN</Code>). The
            contract enforces this in <Code>createRound</Code>; raising the cap requires a fresh
            trusted setup, redeployed verifier contracts, and a redeployed manager.
          </List.Item>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='ink.0'>
              Recent rounds buffer.
            </Text>{' '}
            The contract keeps a ring buffer of the most recent 64 rounds for the explorer's
            "Rounds" page. Older rounds remain valid on-chain but are not enumerated.
          </List.Item>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='ink.0'>
              Plaintext range.
            </Text>{' '}
            Combined ciphertexts must encode a non-negative integer strictly below{' '}
            <Code>2^50</Code> (≈ 1.13 × 10<sup>15</sup>). Recovery uses baby-step / giant-step
            DLOG with a precomputed table — about ~30–60 s and ~1–2 GB heap on first decrypt,
            then the table is reused for every subsequent round. A node never decrypting pays
            nothing. Submitting larger plaintexts leaves the round permanently unrecoverable.
          </List.Item>
        </List.Root>
      </Section>

      <Section heading='Other networks'>
        <Text fontSize='sm' color='ink.2'>
          Sepolia is the only network with a published deployment so far. As more land,{' '}
          <Code>config/networks.go</Code> will gain entries and{' '}
          <Code>DAVINCI_DKG_NETWORK=sepolia</Code> can be swapped for the new name without other
          changes.
        </Text>
      </Section>

      <Section heading='References'>
        <List.Root gap={1.5} fontSize='sm' pl={5}>
          <List.Item>
            <ExternalLink href='https://github.com/vocdoni/davinci-dkg#readme'>
              README — full installation and configuration reference.
            </ExternalLink>
          </List.Item>
          <List.Item>
            <ExternalLink href='https://github.com/vocdoni/davinci-dkg/blob/main/.env.example'>
              .env.example — every available setting with inline documentation.
            </ExternalLink>
          </List.Item>
          <List.Item>
            <ExternalLink href='https://github.com/vocdoni/davinci-dkg/releases'>
              Releases — pinned tags and changelogs.
            </ExternalLink>
          </List.Item>
          <List.Item>
            <ExternalLink href='https://github.com/vocdoni/davinci-dkg/issues'>
              Issue tracker.
            </ExternalLink>
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
