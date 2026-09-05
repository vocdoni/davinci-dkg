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
  { id: 'keygen', title: 'Contribution, finalization and pool keys' },
  { id: 'applications', title: 'Applications, modes and windows' },
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
      description='Non-interactive distributed key generation on an EVM chain: how an epoch is born, how its committee is drawn, how each application gets a key of its own out of the epoch’s pool, and what it takes to open a single ciphertext.'
      sections={SECTIONS}
    >
      <Section id='overview' title='Overview'>
        <P>
          A committee of independent operators jointly generates a <Em>pool</Em> of <Em>ElGamal</Em> public keys on{' '}
          <Em>BabyJubJub</Em>, a SNARK-friendly elliptic curve over the BN254 scalar field — one key per application
          that registers against the epoch. No matching private key is ever assembled anywhere: ciphertexts are opened
          by combining partial decryptions, and ElGamal&rsquo;s additive homomorphism
          lets ciphertexts be aggregated <Em>before</Em> anything is decrypted — so the result of a vote can be revealed
          while the individual ballots stay sealed.
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
          ciphertexts and decryption. <C>DKGAppManager</C> owns per-application registration and the organizer reveal. The
          split exists only to keep each contract under EIP-170 — logically the last two share one storage — and the
          registry and app manager are resolved from the manager on chain, so only one address ever needs configuring.
        </P>
      </Section>

      <Section id='lifecycle' title='Epoch lifecycle'>
        <P>
          An <Em>epoch</Em> is one DKG run. It deals a pool of <C>MaxK = 8</C> keys <C>P_0 … P_7</C>, each shared by
          the same <C>n</C> committee members, any <C>t</C> of which can decrypt under it. Epochs are scheduled on a
          fixed block cadence and split into a short <Em>preparation</Em> and a long <Em>service</Em> window; the next
          one may also open early once the pool is down to its last key.
        </P>
        <Code caption='epoch windows'>{`  startBlock                                                          endBlock
  │ ── Preparation (small, fixed) ──►  ◄──────── Service ───────────► │
  │ CommitteeSelection │ KeyAssembly │ gap │        Live              │
  ├────────────────────┼─────────────┼─────┼──────────────────────────┤
  │ claimSlot          │submitContrib│     │ activatePoolKey (×8) /   │
  │ (lottery)          │  (Groth16)  │     │ registerApplication /    │
  │                    │             │     │ submitCiphertext /       │
  │                    │             │     │ submitPartialDecryption /│
  │                    │             │     │ revealOrganizerSecret /  │
  │                    │             │     │ combineDecryption        │
  └────────────────────┴─────────────┴─────┴──────────────────────────┘
                                      ▲
                            finalizeEpoch (no proof)
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
          window opens, with jitter, so most losing calls are never even sent.{' '}
          <Em>No application schedules an epoch</Em>, and no operator has to either — the set produces them by itself.
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
            { label: 'MaxK', value: '8', mono: true, hint: 'pool keys dealt per epoch, one per application' },
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

      <Section id='keygen' title='Contribution, finalization and pool keys'>
        <Steps
          items={[
            <>
              Each committee member publishes one Feldman verifiable secret sharing (VSS) contribution dealing{' '}
              <C>MaxK</C> polynomials at once: <C>MaxK</C> sets of commitments plus, per recipient, <C>MaxK</C> shares
              encrypted to that member&rsquo;s registry key under one shared ECDH secret, with a Groth16 proof that
              commitments and shares are consistent. One transaction, verified on chain at submission.
            </>,
            <>
              Once at least <C>minValidContributions</C> are in and the finalize gap has passed, anyone calls{' '}
              <C>finalizeEpoch</C> — <Em>no proof</Em>. It freezes the accepted contributor set, flips the epoch to{' '}
              <C>Live</C> and emits <C>EpochLive(eid, contributionCount)</C>. No key exists yet.
            </>,
            <>
              Each pool key is then <Em>activated</Em> separately: <C>activatePoolKey(eid, j, …)</C> carries a Groth16
              proof that <C>P_j = Σ A&#95;&#123;i,j,0&#125;</C> is the sum of the accepted contributors&rsquo; zeroth
              commitments for key <C>j</C>, and stores <C>P_j</C> together with a Merkle root of every member&rsquo;s
              share commitment <C>D&#95;&#123;j,i&#125;</C>. Permissionless, any order, one proof per key; nodes keep
              a couple of keys activated ahead of demand.
            </>,
          ]}
        />
        <P>
          No private key is ever assembled — not by the contract, not by the finalizer, not by any single member. Each{' '}
          <C>P_j</C> only ever exists as <C>n</C> Shamir shares <C>e&#95;&#123;j,i&#125;</C>, of which <C>t</C> suffice
          to act, and a member proves its partial against the Merkle root rather than against per-member storage.
        </P>
      </Section>

      <Section id='applications' title='Applications, modes and windows'>
        <P>
          A Live epoch hosts up to <C>MaxK</C> independent encryption contexts — one per <Em>application</Em>, named
          by a 32-byte <C>aid</C> chosen by whoever registers it. Because <C>aid</C> is a public input of every
          decryption proof it must be non-zero and below the BN254 scalar modulus: clear the top three bits of a random
          or hashed id.
        </P>
        <P>
          There is exactly one registration path, and it is not optional. <C>registerApplication</C> claims the
          epoch&rsquo;s next unused <Em>activated</Em> pool key <C>P_j</C> (<C>PoolExhausted()</C> when none is left,{' '}
          <C>PoolKeyNotActive()</C> when the next one is not proven yet) and records <C>poolIndex = j</C>. The
          application key is then:
        </P>
        <Code caption='application key'>{`PK_aid = P_j            // automatic
PK_aid = P_j + PK_org   // organizer-locked`}</Code>
        <P>
          Whether an organizer key sits on top of the pool key is the application&rsquo;s <Em>mode</Em>, fixed at
          registration:
        </P>
        <Bullets
          items={[
            <>
              <Em>Organizer-locked</Em> (the default). The registration carries <C>PK_org = sk_org·G</C> and a Schnorr
              proof of possession of <C>sk_org</C> (domain <C>davinci-dkg:organizer-register:v1</C>), verified on
              chain, and nothing else about the secret. The committee cannot combine anything until the organizer
              calls <C>revealOrganizerSecret</C> — once, permissionlessly relayable, checked as{' '}
              <C>sk_org·G = PK_org</C>; from that block on it combines by itself. The organizer does nothing per
              ciphertext.
            </>,
            <>
              <Em>Automatic.</Em> No organizer key at all: <C>organizerPK</C> is stored as the identity <C>(0, 1)</C>{' '}
              and the combine proof runs with a zero secret. <C>t</C> partials inside the decryption window are all a
              plaintext takes, and confidentiality rests on the committee threshold only.
            </>,
          ]}
        />
        <P>
          The rest of the policy is fixed at registration too. <Em>Submission</Em> is open to anyone (
          <C>openSubmission</C>), to an allow-list of at most 32 addresses (<C>submitters</C>), or — when both are empty
          — to the registrant only; a block window and a ciphertext cap apply on top. The <Em>decryption window</Em>{' '}
          <C>decryptNotBefore</C> → <C>decryptNotAfter</C> (unix seconds, 0 = unbounded on that side) gates partials
          and combines in both modes: before it opens they revert <C>DecryptionNotOpen()</C>, after it closes{' '}
          <C>DecryptionClosed()</C>, and the close, once set, must lie in the future at registration. A contradictory
          policy reverts <C>InvalidPolicy()</C>. Every ciphertext belongs to a registered application.
        </P>
        <Note tone='warn'>
          Losing the <C>sk_org</C> of an organizer-locked application before the reveal makes it{' '}
          <Em>permanently undecryptable</Em>. It is never transmitted and nothing on chain can reconstruct it.
          Conversely, the organizer can decrypt any ciphertext of its own application from <C>t</C> published partials
          and its own secret, and the reveal is irreversible: once <C>sk_org</C> is on chain every ciphertext under the{' '}
          <C>aid</C> — past and future — is open to the committee threshold.
        </Note>
        <P>
          The pool is what replaces a proof of knowledge of the encryption randomness. A ciphertext copied out of one
          application into another is decrypted under a different <C>P_j</C> and yields nothing useful — so{' '}
          <C>submitCiphertext</C> needs no proof, which is in turn what makes homomorphic aggregation possible: whoever
          submits an aggregated tally cannot know its randomness. The organizer key no longer carries that burden, which
          is why it may be absent or revealed.
        </P>
      </Section>

      <Section id='decryption' title='Threshold decryption'>
        <Steps
          items={[
            <>
              A permitted submitter publishes <C>(C1, C2)</C> with <C>submitCiphertext</C>. The contract checks the
              points are canonical, on-curve and non-identity, assigns the next index for that <C>aid</C>, and emits{' '}
              <C>CiphertextSubmitted</C>. It deliberately skips the prime-order subgroup check (~2M gas); every
              committee node performs it off chain before computing anything, and that off-chain check is load-bearing —
              a cofactor <C>C1</C> would leak <C>d_i mod h</C> from a node that skipped it.
            </>,
            <>
              Each selected node publishes its partial <C>δ_i = e&#95;&#123;j,i&#125;·C1</C> — its share of{' '}
              <Em>this application&rsquo;s</Em> pool key — with a Groth16 proof of a Chaum&ndash;Pedersen
              discrete-logarithm equality (DLEQ) binding <C>δ_i</C> to its share commitment <C>D&#95;&#123;j,i&#125;</C>{' '}
              and to <C>C1</C>, plus a Merkle path (<C>MERKLE_DEPTH = 5</C> siblings) proving that commitment sits at
              leaf <C>i − 1</C> of the key&rsquo;s root. The DLEQ transcript is hashed in-circuit with Poseidon. Nodes
              stagger by slot, so partials arrive in <Em>waves</Em> and only the first <C>t</C> members of a
              seed-derived rotation actually spend gas.
            </>,
            <>
              For an organizer-locked application the organizer reveals <C>sk_org</C> once with{' '}
              <C>revealOrganizerSecret</C> — two calldata words and no proof; the contract checks{' '}
              <C>sk_org·G = PK_org</C> and stores the scalar. An automatic application has nothing to reveal.
            </>,
            <>
              Once <C>t</C> partials are on chain and the application is unlocked, anyone calls{' '}
              <C>combineDecryption</C>. Its Groth16 proof attests three things at once: that <C>Σ λ_k · δ_k</C>{' '}
              Lagrange-interpolates correctly, that it knows the organizer secret behind the registered <C>PK_org</C>{' '}
              (zero for an automatic application), and that <C>m·G + Σ λ_k · δ_k + sk_org·C1 = C2</C>. The recovered
              scalar <C>m</C> lands on chain, readable via <C>getPlaintext</C>.
            </>,
          ]}
        />
        <P>
          Every one of these steps is gated by the application&rsquo;s decryption window: a partial or a combine outside
          it reverts, so a closed window is a hard stop even for a revealed or automatic application.
        </P>
        <Sub>Plaintext range</Sub>
        <P>
          <C>m</C> is recovered by baby-step giant-step (BSGS) discrete-log inversion. The committee node caps at{' '}
          <C>2⁵⁰</C> (≈1.13&times;10¹⁵) and builds a 256 MB table once per process. The browser SDK caps at <C>2³²</C>,
          so its table stays around 16 MB. Submitting a plaintext above the relevant cap leaves the ciphertext
          unrecoverable.
        </P>
      </Section>

      <Section id='security' title='What holds, and what does not'>
        <Sub>Properties</Sub>
        <Bullets
          items={[
            <>
              <Em>No trusted dealer.</Em> Every pool key is generated jointly; corrupting fewer than <C>t</C> operators
              reveals nothing about any of them.
            </>,
            <>
              <Em>Non-interactive.</Em> Every step is one self-contained transaction. Invalid contributions are rejected
              at submission, not disputed afterwards.
            </>,
            <>
              <Em>Verified on chain.</Em> Contribution, pool-key activation, partial decryption and combination are
              each gated by a Groth16 verifier; finalization needs none because it proves nothing. Correctness is
              enforced by the EVM.
            </>,
            <>
              <Em>Permissionless committee.</Em> Anyone can register and be drawn by the lottery; the selection is a
              keccak anyone can replay.
            </>,
            <>
              <Em>Per-application isolation.</Em> One committee serves up to <C>MaxK</C> applications per epoch, each
              under its own key, and a ciphertext moved between them decrypts to nothing useful — with or without an
              organizer key.
            </>,
          ]}
        />
        <Sub>Risks, stated plainly</Sub>
        <Bullets
          items={[
            <>
              <Em>The organizer is a single point of failure for its own application.</Em> Lose <C>sk_org</C> before
              the reveal and every ciphertext under that <C>aid</C> is undecryptable forever; keep it and you can open
              any of them the moment <C>t</C> partials exist, without asking anyone. And the reveal cannot be undone.
            </>,
            <>
              <Em>A colluding threshold breaks confidentiality.</Em> <C>t</C> operators plus the organizer recover the
              plaintext of anything — and for an automatic or a revealed application <C>t</C> operators alone do. The
              lottery makes the set unpredictable, not incorruptible; the decryption window bounds the exposure in time.
            </>,
            <>
              <Em>No proof of knowledge on ciphertexts.</Em> Anyone authorised may submit arbitrary <C>(C1, C2)</C>.
              Cross-application replay is neutralised by the per-application pool keys, but a malformed ciphertext
              inside its own application is the submitter&rsquo;s problem.
            </>,
            <>
              <Em>The subgroup check is off chain.</Em> The contract does not perform it; a node that skips it leaks
              part of its share. This is a node-implementation obligation, not a contract guarantee.
            </>,
            <>
              <Em>Liveness is not guaranteed.</Em> A committee that fails to fill or to reach{' '}
              <C>minValidContributions</C> produces a dead epoch that anyone may abort; a committee that stops answering
              leaves ciphertexts undecrypted until a later epoch is used. And an epoch serves at most <C>MaxK</C>{' '}
              applications — past that, registrations wait for the next one.
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
