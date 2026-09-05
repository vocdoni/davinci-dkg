import { Address, KeyValue } from '~kit'
import { useRuntimeConfig } from '~config/config-context'
import { DocsLayout, type DocsSection } from './DocsLayout'
import { Bullets, C, Code, Em, Ext, Note, P, Section, Steps, Sub } from './prose'

const REPO = 'https://github.com/vocdoni/davinci-dkg'

const SECTIONS: DocsSection[] = [
  { id: 'prerequisites', title: 'Prerequisites' },
  { id: 'install', title: 'Clone and configure' },
  { id: 'start', title: 'Start the node' },
  { id: 'first-boot', title: 'What happens on first boot' },
  { id: 'explorer', title: 'Host an explorer alongside it' },
  { id: 'maintenance', title: 'Maintenance' },
  { id: 'limits', title: 'Limits worth knowing' },
  { id: 'networks', title: 'This network' },
  { id: 'references', title: 'References' },
]

export function DocsRunANodePage() {
  const config = useRuntimeConfig()
  return (
    <DocsLayout
      label='Docs'
      title='Run a node'
      description='Join the committee. The node ships as a single Docker image — no Go, no Node, no build toolchain on the host — and you become eligible for every epoch drawn after you register.'
      sections={SECTIONS}
    >
      <Section id='prerequisites' title='Prerequisites'>
        <Bullets
          items={[
            <>
              <C>docker</C> ≥ 24 and <C>docker compose</C> v2.
            </>,
            <>
              An Ethereum Virtual Machine (EVM) private key, holding a little {config.chainName} native currency for
              fees. The node pays for registry registration, slot claims, contributions and partial decryptions, only
              for the phases it actually takes part in. For Sepolia, the{' '}
              <Ext href='https://sepolia-faucet.pk910.de/'>pk910 proof-of-work faucet</Ext> and the{' '}
              <Ext href='https://cloud.google.com/application/web3/faucet/ethereum/sepolia'>Google Cloud faucet</Ext>{' '}
              both work.
            </>,
            <>
              A JSON-RPC endpoint. Public providers are fine to try it out; a long-lived node wants a dedicated provider
              or your own node.
            </>,
          ]}
        />
      </Section>

      <Section id='install' title='Clone and configure'>
        <P>
          The compose file and the example environment live at the repository root. Nothing is built locally —
          everything is pulled from published images.
        </P>
        <Code>{`git clone https://github.com/vocdoni/davinci-dkg.git
cd davinci-dkg
cp .env.example .env && $EDITOR .env`}</Code>
        <P>The minimum set of entries:</P>
        <Code caption='.env'>{`# JSON-RPC endpoint (use your own provider in production).
DAVINCI_DKG_WEB3_RPC=${config.rpcUrl}

# Hex private key (0x-prefixed). This wallet pays the node's fees and is
# your operator identity in the registry.
DAVINCI_DKG_PRIVKEY=0x<your-private-key>

# Named deployment: contract addresses are built into the binary.
DAVINCI_DKG_NETWORK=${config.chainName}

# Any other network: point at the manager instead and the node resolves
# the registry and the app manager from it on chain.
# DAVINCI_DKG_MANAGER=${config.managerAddress}`}</Code>
        <Note>
          Every flag has a <C>DAVINCI_DKG_…</C> environment equivalent; run <C>davinci-dkg-node --help</C> for the full
          list. Release binaries and source builds are configured exactly the same way.
        </Note>
      </Section>

      <Section id='start' title='Start the node'>
        <P>
          One command brings up the node and Watchtower, which recreates the container when a new image is published.
          The container restarts on failure by default.
        </P>
        <Code>{`docker compose --profile node up -d
docker compose --profile node logs -f node`}</Code>
      </Section>

      <Section id='first-boot' title='What happens on first boot'>
        <Steps
          items={[
            <>
              The node derives a BabyJubJub encryption key from your EVM key and registers (or updates) it in{' '}
              <C>DKGRegistry</C> — one transaction, skipped if you are already registered and active.
            </>,
            <>
              It prints a startup banner with the chain head, registry statistics and its own <C>self:</C> row, then
              polls <C>DKGManager</C> on the configured interval: <C>DAVINCI_DKG_POLL_INTERVAL</C>, which the example{' '}
              <C>.env</C> sets to 20 s and the binary defaults to 5 s.
            </>,
            <>
              It races the other nodes to call <C>createEpoch</C> once the cadence window opens, with random jitter so
              most of the losing calls are never sent. Epochs are produced by the operator set itself; no application
              and no operator has to schedule them.
            </>,
            <>
              From then on it reacts to every phase it is eligible for: claim a slot when the lottery admits it, submit
              a contribution, finalize the whole key pool with one proof, answer ciphertexts with partial decryptions,
              and combine when its turn comes in the seed-derived rotation. It heartbeats and reactivates by itself
              before the inactivity window expires.
            </>,
          ]}
        />
        <Note tone='warn'>
          Registration takes effect for <Em>future</Em> epochs only. An epoch&rsquo;s lottery runs over the registry as
          it was snapshotted at <C>createEpoch</C>, so a <C>claimSlot</C> against an older epoch reverts with{' '}
          <C>NotInSnapshot</C>. Expect to join your first committee in the epoch <Em>after</Em> the one running when you
          came up.
        </Note>
      </Section>

      <Section id='explorer' title='Host an explorer alongside it'>
        <P>
          The node binary serves no HTTP — it only talks to the chain. This explorer is a separate static bundle whose
          chain configuration is a single <C>config.json</C> rendered at build or deploy time, so one image can be
          pointed anywhere.
        </P>
        <Code>{`# Build the bundle against the deployment you want.
make ui-build \\
  RPC_URL=${config.rpcUrl} \\
  MANAGER_ADDRESS=${config.managerAddress} \\
  CHAIN_ID=${config.chainId} CHAIN_NAME=${config.chainName}

# Serve it next to the node (stock nginx, bind-mounted bundle).
docker compose --profile node --profile ui up -d
# → http://<your-host>:8082/`}</Code>
        <P>
          Or run the published image <C>ghcr.io/vocdoni/davinci-dkg-ui</C> directly. For a hosted deployment, the
          repository carries a DigitalOcean App Platform spec at <C>ui/.do/davinci-dkg-ui.yaml</C>; edit its build-time
          environment values to retarget the chain.
        </P>
      </Section>

      <Section id='maintenance' title='Maintenance'>
        <Bullets
          items={[
            <>
              <Em>Upgrades.</Em> Watchtower follows the <C>latest</C> tag. Pin <C>DAVINCI_DKG_TAG=v0.1.0</C> in{' '}
              <C>.env</C>, or drop the watchtower service, for manual control.
            </>,
            <>
              <Em>State.</Em> The node mounts a volume for its data directory, but all per-epoch state is rebuilt from
              on-chain records on restart — stopping and starting mid-epoch is safe.
            </>,
            <>
              <Em>Circuit artifacts.</Em> Proving keys are downloaded once and cached under <C>~/.davinci/artifacts</C>{' '}
              (inside the container, on the mounted volume). The node verifies their SHA-256 against the hashes pinned
              in the binary and falls back to a local trusted setup only if the download fails.
            </>,
            <>
              <Em>Stop.</Em> <C>docker compose --profile node down</C>; add <C>-v</C> to wipe the volume too.
            </>,
            <>
              <Em>Cost.</Em> You pay gas only for the phases you take part in. The per-call breakdown is in{' '}
              <Ext href={`${REPO}/blob/main/BENCHMARKS.md`}>BENCHMARKS.md</Ext>.
            </>,
          ]}
        />
      </Section>

      <Section id='limits' title='Limits worth knowing'>
        <Bullets
          items={[
            <>
              <Em>Committee size.</Em> Capped at <C>32</C> per epoch — the circuits&rsquo; compile-time <C>MaxN</C>,
              mirrored in <C>Sizes.sol</C> and enforced by <C>createEpoch</C>. Raising it needs a fresh trusted setup,
              redeployed verifiers and a redeployed manager.
            </>,
            <>
              <Em>Ciphertext checks.</Em> Ciphertexts are plain <C>(C1, C2)</C> calldata with no proof of knowledge of
              the randomness. Before releasing a partial the node checks that <C>C1</C> is in the prime-order subgroup —
              the contract does not — which is what stops a small-order point from leaking a share.
            </>,
            <>
              <Em>Organizer secrets.</Em> A ciphertext of an organizer-locked application is only combinable once its
              organizer has called <C>revealOrganizerSecret</C>; until then the node parks the application after posting
              its partials and wakes it on the reveal event, so a kept secret costs nothing per tick. An automatic
              application needs no reveal. Both are gated by the application&rsquo;s decryption window.
            </>,
            <>
              <Em>Dead epochs.</Em> An epoch whose committee never filled, or whose key assembly closed below{' '}
              <C>minValidContributions</C>, sits until anyone calls <C>abortEpoch</C> — permissionless, and it reverts
              for a healthy epoch. The next epoch opens on the normal cadence regardless.
            </>,
            <>
              <Em>Plaintext range.</Em> Combined ciphertexts must encode a non-negative integer below <C>2⁵⁰</C>.
              Recovery uses a precomputed baby-step giant-step (BSGS) table of 256 MB, built once on the first decrypt
              and reused afterwards. A node that never decrypts never builds it.
            </>,
          ]}
        />
      </Section>

      <Section id='networks' title='This network'>
        <P>
          The values this explorer is pointed at. A named network resolves its addresses from the binary; anything else
          only needs the manager, since the registry and app manager are read from it on chain.
        </P>
        <KeyValue
          columns={2}
          items={[
            { label: 'network', value: config.chainName },
            { label: 'chain id', value: config.chainId, mono: true },
            { label: 'DKGManager', value: <Address value={config.managerAddress} full /> },
            { label: 'rpc endpoint', value: config.rpcUrl, mono: true },
            { label: 'deploy block', value: config.deployBlock.toLocaleString(), mono: true },
          ]}
        />
        <Sub>Other networks</Sub>
        <P>
          New deployments are added to <C>config/networks.go</C>, after which <C>DAVINCI_DKG_NETWORK=&lt;name&gt;</C>{' '}
          works without any other change.
        </P>
      </Section>

      <Section id='references' title='References'>
        <Bullets
          items={[
            <>
              <Ext href={`${REPO}#readme`}>README</Ext> — full installation and configuration reference.
            </>,
            <>
              <Ext href={`${REPO}/blob/main/.env.example`}>.env.example</Ext> — every setting with inline documentation.
            </>,
            <>
              <Ext href={`${REPO}/releases`}>Releases</Ext> — pinned tags and changelogs.
            </>,
            <>
              <Ext href={`${REPO}/issues`}>Issue tracker</Ext>.
            </>,
          ]}
        />
      </Section>
    </DocsLayout>
  )
}
