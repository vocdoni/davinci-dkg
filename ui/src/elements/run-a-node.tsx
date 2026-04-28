import { Badge, Box, Heading, HStack, Link, List, Stack, Text } from '@chakra-ui/react'
import { LuExternalLink } from 'react-icons/lu'
import type { ReactNode } from 'react'
import { CodeBlock } from '~components/ui/CodeBlock'

// Operator-facing "how to run a node" page. Currently Sepolia-only because
// that's the only deployment in config/networks.go.
//
// Style: documentation, not marketing. Sections are numbered, not iconed
// in colored bubbles; the page reads top-to-bottom like a README.
export function RunNode() {
  return (
    <Stack gap={10} maxW='3xl'>
      <Box pb={4} borderBottomWidth='1px' borderColor='gray.800'>
        <HStack gap={3} mb={2} align='baseline' wrap='wrap'>
          <Heading size='lg' color='gray.100'>
            Running a node
          </Heading>
          <Badge colorPalette='gray' variant='outline' fontSize='2xs'>
            sepolia only
          </Badge>
        </HStack>
        <Text color='gray.400' fontSize='sm' lineHeight='1.6'>
          Procedure for joining the davinci-dkg committee on Sepolia. The node binary ships as a
          single Docker image; no Go, Node, or build toolchain is required on the host. The
          explorer UI is a separate, optional container — see step 4 below.
        </Text>
      </Box>

      <Section heading='Prerequisites'>
        <List.Root gap={2} fontSize='sm' color='gray.300' pl={5}>
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
        <Text fontSize='sm' color='gray.300'>
          The compose file and the example env live at the repository root. Nothing is built locally
          — everything is pulled from the published images.
        </Text>
        <CodeBlock language='bash'>
          {`git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg`}
        </CodeBlock>
      </Section>

      <Section heading='2. Configure the node'>
        <Text fontSize='sm' color='gray.300'>
          Copy the example env file and fill in your private key. The Sepolia network preset
          resolves all contract addresses for you.
        </Text>
        <CodeBlock language='bash'>{`cp .env.example .env
$EDITOR .env`}</CodeBlock>
        <Text fontSize='sm' color='gray.300'>
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
        <Text fontSize='sm' color='gray.300'>
          One command brings up the node and Watchtower (which auto-updates the image when new
          versions are published). The container restarts on failure by default.
        </Text>
        <CodeBlock language='bash'>{`docker compose --profile node up -d`}</CodeBlock>
        <Text fontSize='sm' color='gray.300'>
          Tail the logs to confirm it is running:
        </Text>
        <CodeBlock language='bash'>{`docker compose --profile node logs -f node`}</CodeBlock>
        <Text fontSize='sm' color='gray.300'>
          On first boot the node:
        </Text>
        <List.Root gap={1.5} fontSize='sm' color='gray.300' pl={5}>
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
        <Text fontSize='sm' color='gray.300'>
          The node binary does not serve any HTTP — it only talks to the chain. To browse rounds,
          the registry, and the playground from a wallet, host the explorer separately. The
          explorer is a Vite static bundle whose chain config is baked in at build time.
        </Text>
        <Text fontSize='sm' color='gray.300'>
          The compose <Code>ui</Code> profile bind-mounts the locally-built bundle into stock{' '}
          <Code>nginx:alpine</Code>:
        </Text>
        <CodeBlock language='bash'>
          {`# 1. Build the bundle with the chain config you want.
make ui-build \\
  RPC_URL=https://eth-sepolia.public.blastapi.io \\
  MANAGER_ADDRESS=0xd3ef727b695b21e108497c36f9dcec52d741298a \\
  CHAIN_ID=11155111 CHAIN_NAME=sepolia

# 2. Run nginx alongside the node.
docker compose --profile node --profile ui up -d
# explorer reachable at http://<your-host>:8082/`}
        </CodeBlock>
        <Text fontSize='sm' color='gray.300'>
          For a public deployment, point DigitalOcean App Platform at the spec checked into the
          repo (<Code>ui/.do/davinci-dkg-ui.yaml</Code>) and edit the <Code>BUILD_TIME</Code> env
          values to retarget the chain. App Platform builds the Dockerfile per push and serves
          the static files from its edge — no nginx in the loop.
        </Text>
        <CodeBlock language='bash'>{`doctl apps create --spec ui/.do/davinci-dkg-ui.yaml`}</CodeBlock>
      </Section>

      <Section heading='5. Maintenance'>
        <List.Root gap={2.5} fontSize='sm' color='gray.300' pl={5}>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='gray.100'>
              Upgrades.
            </Text>{' '}
            Watchtower follows the <Code>latest</Code> tag and recreates the container when a new
            version ships. Pin <Code>DAVINCI_DKG_TAG=v0.1.0</Code> in <Code>.env</Code> (or remove
            the watchtower service from <Code>docker-compose.yml</Code>) for manual control.
          </List.Item>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='gray.100'>
              State.
            </Text>{' '}
            The node mounts a Docker volume named <Code>run</Code> for its data directory. All
            per-round state is rebuilt from on-chain records on restart, so a stop/start mid-round
            is safe.
          </List.Item>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='gray.100'>
              Stop.
            </Text>{' '}
            <Code>docker compose --profile node down</Code>. Add <Code>-v</Code> to also wipe the
            volume.
          </List.Item>
        </List.Root>
      </Section>

      <Section heading='Limits'>
        <List.Root gap={2.5} fontSize='sm' color='gray.300' pl={5}>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='gray.100'>
              Committee size.
            </Text>{' '}
            Capped at <Code>32</Code> per round (the circuits' compile-time <Code>MaxN</Code>). The
            contract enforces this in <Code>createRound</Code>; raising the cap requires a fresh
            trusted setup, redeployed verifier contracts, and a redeployed manager.
          </List.Item>
          <List.Item>
            <Text as='span' fontWeight='semibold' color='gray.100'>
              Recent rounds buffer.
            </Text>{' '}
            The contract keeps a ring buffer of the most recent 64 rounds for the explorer's
            "Rounds" page. Older rounds remain valid on-chain but are not enumerated.
          </List.Item>
        </List.Root>
      </Section>

      <Section heading='Other networks'>
        <Text fontSize='sm' color='gray.300'>
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
