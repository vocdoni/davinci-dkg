import { Address, KeyValue } from '~kit'
import { useRuntimeConfig } from '~config/config-context'
import { DocsLayout, type DocsSection } from './DocsLayout'
import { Bullets, C, Code, Em, Ext, Note, P, Section, Steps, Sub } from './prose'

const PAPER = 'https://eprint.iacr.org/2026/552'
const DAVINCI = 'https://davinci.vote'
const REPO = 'https://github.com/vocdoni/davinci-dkg'

const SECTIONS: DocsSection[] = [
  { id: 'overview', title: 'Overview' },
  { id: 'lifecycle', title: 'Epoch lifecycle' },
  { id: 'lottery', title: 'Committee lottery' },
  { id: 'keygen', title: 'Contribution and finalization' },
  { id: 'applications', title: 'Applications and organizer keys' },
  { id: 'decryption', title: 'Threshold decryption' },
  { id: 'security', title: 'What holds, and what does not' },
  { id: 'deployment', title: 'This deployment' },
]

export function DocsProtocolPage() {
  const config = useRuntimeConfig()
  return (
    <DocsLayout
      label='Docs'
      title='Protocol'
      description='Non-interactive distributed key generation on an EVM chain: how an epoch is born, how its committee is drawn, how applications get their own key, and what it takes to open a single ciphertext.'
      sections={SECTIONS}
    >
      <Section id='overview' title='Overview'>
        <P>
          A committee of independent operators jointly generates an <Em>ElGamal</Em> public key on{' '}
          <Em>BabyJubJub</Em>, a SNARK-friendly elliptic curve over the BN254 scalar field. The matching private key is
          never assembled anywhere: ciphertexts are opened by combining partial decryptions, and ElGamal&rsquo;s additive
          homomorphism lets ciphertexts be aggregated <Em>before</Em> anything is decrypted — so the result of a vote
          can be revealed while the individual ballots stay sealed.
        </P>
        <P>
          Every step is a single self-contained transaction carrying a Groth16 proof. There is no complaint phase and no
          dispute window: an invalid contribution is rejected at submission time by the on-chain verifier, not audited
          afterwards. That is what the &ldquo;non-interactive&rdquo; in NI-DKG buys.
        </P>
        <P>
          NI-DKG is a building block of the <Ext href={DAVINCI}>DAVINCI</Ext> voting protocol. The construction and its
          security proofs are in the <Ext href={PAPER}>NI-DKG paper</Ext>; the reference implementation, contracts and
          circuits are in the <Ext href={REPO}>davinci-dkg repository</Ext>.
        </P>
        <Sub>Three contracts</Sub>
        <P>
          <C>DKGRegistry</C> holds operator identities and liveness. <C>DKGManager</C> owns the epoch state machine,
          ciphertexts and decryption. <C>DKGAppManager</C> owns per-application registration and organizer shares. The
          split exists only to keep each contract under EIP-170 — logically the last two share one storage — and the
          registry and app manager are resolved from the manager on chain, so only one address ever needs configuring.
        </P>
      </Section>

      <Section id='lifecycle' title='Epoch lifecycle'>
        <P>
          An <Em>epoch</Em> is one DKG run. It produces a single collective public key <C>PK_ep</C> shared by <C>n</C>{' '}
          committee members, any <C>t</C> of which can decrypt. Epochs are scheduled on a fixed block cadence and split
          into a short <Em>preparation</Em> and a long <Em>service</Em> window.
        </P>
        <Code caption='epoch windows'>{`  startBlock                                                          endBlock
  │ ── Preparation (small, fixed) ──►  ◄──────── Service ───────────► │
  │ CommitteeSelection │ KeyAssembly │ gap │        Live              │
  ├────────────────────┼─────────────┼─────┼──────────────────────────┤
  │ claimSlot          │submitContrib│     │ registerApplication /    │
  │ (lottery)          │  (Groth16)  │     │ submitCiphertext /       │
  │                    │             │     │ submitPartialDecryption /│
  │                    │             │     │ submitOrganizerShare /   │
  │                    │             │     │ combineDecryption        │
  └────────────────────┴─────────────┴─────┴──────────────────────────┘
                                      ▲
                            finalizeEpoch (Groth16)
                            flips KeyAssembly → Live`}</Code>
        <P>
          Each preparation window is an <Em>absolute</Em> block count, not a fraction of the epoch: the lottery is one
          keccak per claimer and a contribution is one transaction per member, so a fixed budget is the right shape. A
          multi-day epoch keeps the same short preparation and spends the rest in service.
        </P>
        <Sub>Who starts an epoch</Sub>
        <P>
          <C>createEpoch</C> is permissionless but cadence-gated: it reverts unless{' '}
          <C>block.number &ge; nextEpochStartBlock()</C>. In production every node races to fire it the moment the
          window opens, with jitter, so most losing calls are never even sent. <Em>No application schedules an epoch</Em>
          , and no operator has to either — the set produces them by itself.
        </P>
        <Sub>Deploy-time parameters</Sub>
        <P>
          The four block constants and the policy floors are <C>DKGManager</C> constructor immutables. Defaults below;
          this deployment&rsquo;s real values are readable on the manager.
        </P>
        <KeyValue
          columns={2}
          items={[
            { label: 'EPOCH_DURATION_BLOCKS', value: '100', mono: true, hint: '≈20 min at 12 s — the cadence anchor' },
            { label: 'COMMITTEE_SELECTION_BLOCKS', value: '25', mono: true, hint: 'lottery window' },
            { label: 'KEY_ASSEMBLY_BLOCKS', value: '25', mono: true, hint: 'contribution window' },
            { label: 'FINALIZE_GAP_BLOCKS', value: '5', mono: true, hint: 'cooldown before finalizeEpoch' },
            { label: 'SEED_DELAY_BLOCKS', value: '1', mono: true, hint: 'seed = blockhash(startBlock + this)' },
            { label: 'MaxN', value: '32', mono: true, hint: 'compile-time committee cap, mirrors the circuits' },
            { label: 'INACTIVITY_WINDOW', value: '50,400', mono: true, hint: '≈7 days — heartbeat window before reap' },
          ]}
        />
        <Note>
          Phases exposed by <C>EpochPhase</C>: <C>None</C>, <C>CommitteeSelection</C>, <C>KeyAssembly</C>, <C>Live</C>,{' '}
          <C>Aborted</C>. A <C>Completed</C> value is reserved but unused in the live state machine.
        </Note>
      </Section>

      <Section id='lottery' title='Committee lottery'>
        <P>
          Every epoch draws a fresh committee from the registry, so no organizer can prefer specific operators. With{' '}
          <C>n = committeeSize</C>, <C>α = lotteryAlphaBps / 10000</C> (the oversubscription factor), <C>R</C> the
          registry&rsquo;s <C>activeCount()</C> snapshotted at <C>createEpoch</C>, and <C>seed</C> the block hash
          resolved on the first <C>claimSlot</C>, an operator is admissible exactly when:
        </P>
        <Code caption='eligibility'>{`keccak256(seed ‖ msg.sender) < (α · n · 2²⁵⁶) / R`}</Code>
        <P>
          Admissible operators then race first-come-first-served until <C>n</C> slots are filled, at which point the
          epoch advances to key assembly on its own. Anyone can replay the keccak and check the result — no ZK proof, no
          coordinator, nothing to trust. Because <C>R</C> and the registry are snapshotted at creation, only operators
          registered <Em>before</Em> the epoch existed may claim, so fresh identities cannot be ground against a
          revealed seed.
        </P>
        <P>
          If the committee never fills, the epoch is dead: anyone may record the abort, and the next scheduled epoch
          opens regardless. An epoch that could still be finalized cannot be aborted by anybody. Whoever wins the{' '}
          <C>createEpoch</C> race picks <C>(t, n, minValidContributions, α)</C>, which is why the deployment pins floors
          — <C>MIN_THRESHOLD</C>, <C>MIN_COMMITTEE_SIZE</C> and a ceiling <C>MAX_LOTTERY_ALPHA_BPS</C>.
        </P>
      </Section>

      <Section id='keygen' title='Contribution and finalization'>
        <Steps
          items={[
            <>
              Each committee member publishes a Feldman VSS contribution: polynomial commitments plus shares encrypted
              to the other members&rsquo; registry keys, with a Groth16 proof that the two are consistent. One
              transaction, verified on chain at submission.
            </>,
            <>
              The contract accumulates accepted contributions into the collective key{' '}
              <C>PK_ep = Σ a&#95;&#123;i,0&#125;·G</C> — the sum of every contributor&rsquo;s zeroth commitment — and
              records each member&rsquo;s share commitment <C>D_i</C>.
            </>,
            <>
              Once at least <C>minValidContributions</C> are in and the finalize gap has passed, any member calls{' '}
              <C>finalizeEpoch</C> with a Groth16 proof over the aggregate. The epoch flips to <C>Live</C> and{' '}
              <C>PK_ep</C> becomes readable on chain.
            </>,
          ]}
        />
        <P>
          The private key <C>sk_ep</C> is never assembled — not by the contract, not by the finalizer, not by any single
          member. It only ever exists as <C>n</C> Shamir shares <C>d_i</C>, of which <C>t</C> suffice to act.
        </P>
      </Section>

      <Section id='applications' title='Applications and organizer keys'>
        <P>
          A Live epoch hosts many independent encryption contexts — one per <Em>application</Em>, named by a 32-byte{' '}
          <C>aid</C> the integrator chooses. Because <C>aid</C> is a public input of every decryption proof it must be
          non-zero and below the BN254 scalar modulus: clear the top three bits of a random or hashed id.
        </P>
        <P>
          There is exactly one registration path, and it is not optional. <C>registerApplication</C> publishes{' '}
          <C>PK_org = sk_org·G</C> together with a Schnorr proof of possession of <C>sk_org</C> (domain{' '}
          <C>davinci-dkg:organizer-register:v1</C>), verified on chain. The application key is then:
        </P>
        <Code caption='application key'>{`PK_aid = PK_ep + PK_org`}</Code>
        <P>
          so opening a ciphertext needs <Em>both</Em> the committee threshold and the organizer. The committee alone
          only ever recovers <C>sk_ep·C1</C>. <C>policy.authorizedSubmitter == address(0)</C> resolves to the
          registering address — there is no open submission and no bare epoch-key path.
        </P>
        <Note tone='warn'>
          Losing <C>sk_org</C> makes the application <Em>permanently undecryptable</Em>. It is never transmitted and
          nothing on chain can reconstruct it. Conversely, the organizer can decrypt any ciphertext of its own
          application by combining its <C>Δ</C> with the committee&rsquo;s published partials: within an application the
          organizer is trusted and accountable; across applications, secrecy rests on DDH over the organizer keys.
        </Note>
        <P>
          This is what replaces a proof of knowledge of the encryption randomness. A ciphertext copied out of one
          application and decrypted under another yields <C>sk_ep·C1</C>, useless without the target
          application&rsquo;s <C>sk_org·C1</C> — so <C>submitCiphertext</C> needs no proof, which is in turn what makes
          homomorphic aggregation possible: whoever submits an aggregated tally cannot know its randomness.
        </P>
      </Section>

      <Section id='decryption' title='Threshold decryption'>
        <Steps
          items={[
            <>
              The authorised submitter publishes <C>(C1, C2)</C> with <C>submitCiphertext</C>. The contract checks the
              points are canonical, on-curve and non-identity, assigns the next index for that <C>aid</C>, and emits{' '}
              <C>CiphertextSubmitted</C>. It deliberately skips the prime-order subgroup check (~2M gas); every
              committee node performs it off chain before computing anything, and that off-chain check is load-bearing —
              a cofactor <C>C1</C> would leak <C>d_i mod h</C> from a node that skipped it.
            </>,
            <>
              Each selected node publishes its partial <C>δ_i = d_i·C1</C> with a Groth16 proof of a Chaum&ndash;Pedersen
              DLEQ binding <C>δ_i</C> to its share commitment <C>D_i</C> and to <C>C1</C>; the DLEQ transcript is hashed
              in-circuit with Poseidon. Nodes stagger by slot, so partials arrive in <Em>waves</Em> and only the first{' '}
              <C>t</C> members of a seed-derived rotation actually spend gas.
            </>,
            <>
              The organizer publishes <C>Δ = sk_org·C1</C> with a Chaum&ndash;Pedersen DLEQ whose challenge is a{' '}
              <Em>keccak</Em> rather than a Poseidon — cheap enough for a browser-only organizer, and recomputed by the
              contract from calldata. The contract stores only <C>keccak256(Δ ‖ A1 ‖ A2 ‖ z)</C> and{' '}
              <Em>never verifies the DLEQ itself</Em>.
            </>,
            <>
              Once <C>t</C> partials <Em>and</Em> an organizer share are on chain, anyone calls{' '}
              <C>combineDecryption</C>. Its Groth16 proof attests three things at once: that{' '}
              <C>Σ λ_k · δ_k</C> Lagrange-interpolates correctly, that the organizer&rsquo;s DLEQ verifies against the
              registered <C>PK_org</C> and the challenge <C>e</C> the contract pinned, and that{' '}
              <C>m·G + Σ λ_k · δ_k + Δ = C2</C>. The recovered scalar <C>m</C> lands on chain, readable via{' '}
              <C>getPlaintext</C>.
            </>,
          ]}
        />
        <P>
          A malformed organizer share cannot brick a ciphertext: re-submission overwrites the stored hash until the
          plaintext lands, and nodes simply skip a share whose DLEQ does not verify and re-check on the next tick.
          Anyone may relay a share — it is self-authenticating.
        </P>
        <Sub>Plaintext range</Sub>
        <P>
          <C>m</C> is recovered by baby-step / giant-step discrete-log inversion. The committee node caps at{' '}
          <C>2⁵⁰</C> (≈1.13&times;10¹⁵, ~1 GB table); the browser SDK caps at <C>2³²</C> so its table stays around
          16 MB. Submitting a plaintext above the relevant cap leaves the ciphertext unrecoverable.
        </P>
      </Section>

      <Section id='security' title='What holds, and what does not'>
        <Sub>Properties</Sub>
        <Bullets
          items={[
            <>
              <Em>No trusted dealer.</Em> The collective key is generated jointly; corrupting fewer than <C>t</C>{' '}
              operators reveals nothing about <C>sk_ep</C>.
            </>,
            <>
              <Em>Non-interactive.</Em> Every step is one self-contained transaction. Invalid contributions are rejected
              at submission, not disputed afterwards.
            </>,
            <>
              <Em>Verified on chain.</Em> Contribution, finalization, partial decryption and combination are each gated
              by a Groth16 verifier. Correctness is enforced by the EVM.
            </>,
            <>
              <Em>Permissionless committee.</Em> Anyone can register and be drawn by the lottery; the selection is a
              keccak anyone can replay.
            </>,
            <>
              <Em>Per-application isolation.</Em> One committee serves many applications, and a ciphertext moved between
              them decrypts to nothing useful.
            </>,
          ]}
        />
        <Sub>Risks, stated plainly</Sub>
        <Bullets
          items={[
            <>
              <Em>The organizer is a single point of failure for its own application.</Em> Lose <C>sk_org</C> and every
              ciphertext under that <C>aid</C> is undecryptable forever; keep it and you can open any of them the moment{' '}
              <C>t</C> partials exist, without asking anyone.
            </>,
            <>
              <Em>A colluding threshold breaks confidentiality.</Em> <C>t</C> operators plus the organizer recover the
              plaintext of anything. The lottery makes the set unpredictable, not incorruptible.
            </>,
            <>
              <Em>No proof of knowledge on ciphertexts.</Em> Anyone authorised may submit arbitrary <C>(C1, C2)</C>.
              Cross-application replay is neutralised by the organizer key, but a malformed ciphertext inside its own
              application is the submitter&rsquo;s problem.
            </>,
            <>
              <Em>The subgroup check is off chain.</Em> The contract does not perform it; a node that skips it leaks
              part of its share. This is a node-implementation obligation, not a contract guarantee.
            </>,
            <>
              <Em>Liveness is not guaranteed.</Em> A committee that fails to fill or to reach{' '}
              <C>minValidContributions</C> produces a dead epoch that anyone may abort; a committee that stops answering
              leaves ciphertexts undecrypted until a later epoch is used.
            </>,
            <>
              <Em>Plaintexts above the BSGS cap are lost.</Em> Not detectable at submission time — the ciphertext is
              accepted and simply never inverts.
            </>,
            <>
              <Em>Lottery entropy is a block hash.</Em> A proposer with a stake in the outcome has one bit of influence
              per block over the seed, bounded by the oversubscription factor <C>α</C> and the registry size.
            </>,
          ]}
        />
      </Section>

      <Section id='deployment' title='This deployment'>
        <P>
          Everything below is read from the explorer&rsquo;s runtime configuration, not baked into this page — the same
          bundle serves any deployment.
        </P>
        <KeyValue
          columns={2}
          items={[
            { label: 'network', value: config.chainName },
            { label: 'chain id', value: config.chainId, mono: true },
            { label: 'DKGManager', value: <Address value={config.managerAddress} full /> },
            { label: 'deploy block', value: config.deployBlock.toLocaleString(), mono: true },
            {
              label: 'registry / app manager',
              value: 'resolved from the manager on chain',
            },
            { label: 'block explorer', value: config.explorerUrl ?? 'not configured' },
          ]}
        />
      </Section>
    </DocsLayout>
  )
}
