import { Address, KeyValue } from '~kit'
import { useRuntimeConfig } from '~config/config-context'
import { paths } from '~routes/paths'
import { DocsLayout, type DocsSection } from './DocsLayout'
import { Bullets, C, Code, Em, Ext, Internal, Note, P, Section, Sub } from './prose'

const REPO = 'https://github.com/vocdoni/davinci-dkg'

const SECTIONS: DocsSection[] = [
  { id: 'install', title: 'Install' },
  { id: 'roles', title: 'Who does what' },
  { id: 'client', title: 'Read-only client' },
  { id: 'epochs', title: 'Reading an epoch' },
  { id: 'watch', title: 'Watching for events' },
  { id: 'register', title: 'Registering an application' },
  { id: 'encrypt', title: 'Encrypting a value' },
  { id: 'submit', title: 'Submitting a ciphertext' },
  { id: 'reveal', title: 'Revealing the organizer secret' },
  { id: 'read', title: 'Reading the plaintext' },
  { id: 'reference', title: 'Reference' },
]

export function DocsSdkPage() {
  const config = useRuntimeConfig()
  const manager = config.managerAddress
  return (
    <DocsLayout
      label='Docs'
      title='SDK'
      description='@vocdoni/davinci-dkg-sdk is a TypeScript wrapper around the contracts and the BabyJubJub cryptography. It runs against any EVM JSON-RPC endpoint with no service in between, and an organizer never needs a prover.'
      sections={SECTIONS}
    >
      <Section id='install' title='Install'>
        <P>
          ECMAScript-modules-only (ESM) TypeScript package. <C>viem</C> is a peer dependency; <C>@zk-kit/baby-jubjub</C>{' '}
          provides the curve arithmetic and is pure TypeScript, so nothing needs a native build or a Node polyfill in
          the browser.
        </P>
        <Code>{`pnpm add @vocdoni/davinci-dkg-sdk viem`}</Code>
        <P>
          Inside this repository the package is consumed through a file dependency, which is what this explorer does:
        </P>
        <Code caption='package.json'>{`"@vocdoni/davinci-dkg-sdk": "link:../sdk"`}</Code>
      </Section>

      <Section id='roles' title='Who does what'>
        <P>
          Two roles, and the SDK is aimed squarely at the second. <Em>Operators</Em> run the node binary: they open each
          epoch on the contract&rsquo;s block cadence, win committee slots by lottery, publish contributions, finalize,
          activate the pool keys, and later answer ciphertexts with partial decryptions. <Em>Organizers</Em> never
          create or schedule an epoch. They register an application against whichever epoch is already Live — claiming
          one of its pool keys <C>P_j</C> — encrypt under <C>PK_aid</C>, submit ciphertexts, and, for an{' '}
          <Em>organizer-locked</Em> application, reveal <C>sk_org</C> once, when they decide the plaintexts may exist. An{' '}
          <Em>automatic</Em> application has no organizer key at all: <C>PK_aid = P_j</C> and the committee decrypts on
          its own.
        </P>
        <P>
          Only an organizer-locked registration carries a proof of possession of <C>sk_org</C> — a Schnorr proof,
          verified on chain. Submitting a ciphertext proves nothing, and the reveal is two calldata words the contract
          checks itself. An organizer therefore needs no circuit artifacts and no prover at all: everything expensive
          lives on the committee&rsquo;s side.
        </P>
      </Section>

      <Section id='client' title='Read-only client'>
        <P>
          All reads go through <C>DKGClient</C>. It needs a viem <C>PublicClient</C> and the address of the deployed{' '}
          <C>DKGManager</C>; the registry and the app manager are discovered from it on first use.
        </P>
        <Code caption='client.ts'>{`import { createPublicClient, http } from 'viem'
import { DKGClient } from '@vocdoni/davinci-dkg-sdk'

const publicClient = createPublicClient({
  transport: http('${config.rpcUrl}'),
})

export const dkg = new DKGClient({
  publicClient,
  managerAddress: '${manager}',
})`}</Code>
      </Section>

      <Section id='epochs' title='Reading an epoch'>
        <P>
          Epoch identifiers are 12-byte values: a 4-byte chain prefix and an 8-byte nonce. Build one with{' '}
          <C>buildEpochId</C>, or pass a known one as a hex string. Most applications do not choose at all. They take
          the newest epoch whose status is <C>Live</C> and whose pool still has an activated, unclaimed key.
        </P>
        <Code caption='epoch.ts'>{`import { buildEpochId, EpochPhase, roundStatusLabel } from '@vocdoni/davinci-dkg-sdk'

const nonce   = await dkg.epochNonce()
const prefix  = await dkg.roundPrefix()
const epochId = buildEpochId(prefix, nonce - 1n)

const epoch = await dkg.getEpoch(epochId)
console.log(roundStatusLabel(epoch.status), epoch.policy.threshold, epoch.policy.committeeSize)

// Once Live, the pool is dealt one key at a time.
if (epoch.status === EpochPhase.Live) {
  const { nextIndex, activated } = await dkg.getPoolStatus(epochId)
  const free = (activated >> nextIndex) & 1        // is the next key activated?
  // Twisted-Edwards (TE) coordinates, the form the SDK works in.
  const pj = await dkg.getPoolKey(epochId, nextIndex)      // [x, y]
  const root = await dkg.getPoolShareRoot(epochId, nextIndex)
}`}</Code>
      </Section>

      <Section id='watch' title='Watching for events'>
        <P>
          The monitor helpers wrap viem event logs and polling for long-running processes and backend jobs. Each returns
          an unsubscribe function.
        </P>
        <Code caption='watch.ts'>{`import {
  watchNewEpochs,
  watchEpochLive,
  waitForPoolKeyActivated,
  watchCiphertextSubmitted,
} from '@vocdoni/davinci-dkg-sdk'

const stop = watchNewEpochs(dkg, (epochId, organizer) => { /* … */ })

// Fires once per epoch, before any key is activated.
watchEpochLive(dkg, epochId, (contributionCount) => { /* … */ })

// Polls getPoolStatus until key j is activated (registration needs it).
await waitForPoolKeyActivated(dkg, epochId, 0)

// Ciphertexts as they land. c1/c2 come back in TE form.
watchCiphertextSubmitted(dkg, epochId, (event) => {
  console.log(event.aid, event.ciphertextIndex, event.c1, event.c2)
})

stop()`}</Code>
      </Section>

      <Section id='register' title='Registering an application'>
        <P>
          Every ciphertext belongs to a registered application, and an unregistered <C>aid</C> reverts. Registration
          claims the epoch&rsquo;s next activated pool key (<C>PoolExhausted()</C> / <C>PoolKeyNotActive()</C>{' '}
          otherwise) and, in the default <Em>organizer-locked</Em> mode, binds an organizer key on top of it so the
          encryption key becomes <C>PK_aid = P_j + PK_org</C>. Only <C>PK_org</C> and the Schnorr proof of possession
          go on chain — <C>sk_org</C> never leaves the process that drew it until you choose to reveal it. Every policy
          field is optional; the defaults are organizer-locked, registrant-only submission, no cap, no block window and
          no decryption window.
        </P>
        <Code caption='register.ts'>{`import { createWalletClient, http } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { DKGWriter, randomOrganizerSecret } from '@vocdoni/davinci-dkg-sdk'

const walletClient = createWalletClient({
  account: privateKeyToAccount('0x<your-private-key>'),
  transport: http('${config.rpcUrl}'),
})

const writer = new DKGWriter({
  publicClient,
  walletClient,
  managerAddress: '${manager}',
})

const skOrg = randomOrganizerSecret()   // store this BEFORE you register
const aid   = '0x…'                     // bytes32, non-zero, top three bits clear

await writer.registerApplication(epochId, aid, {
  mode:            AppMode.OrganizerLocked, // the default
  openSubmission:  false,                   // true = anyone may submitCiphertext
  submitters:      [],                      // allow-list (max 32); empty = the caller only
  maxCiphertexts:  0,                       // 0 = unlimited
  notBeforeBlock:  0n,                      // submitCiphertext block window, 0 = none
  notAfterBlock:   0n,
  decryptNotBefore: 0n,                     // decryption window, unix seconds; 0 = unbounded
  decryptNotAfter:  0n,                     // must lie in the future when set
}, skOrg)

const app = await dkg.getApplication(epochId, aid)
app.poolIndex        // j — the pool key this application claimed`}</Code>
        <Note tone='warn'>
          <Em>Keep sk_org.</Em> For an organizer-locked application it is the only decryption capability until you
          reveal it, it is never transmitted, and nothing on chain can reconstruct it. Lose it and every ciphertext ever
          submitted under that <C>aid</C> is permanently undecryptable — the committee threshold alone cannot open them,
          by design.
        </Note>
        <Sub>Automatic mode</Sub>
        <P>
          With <C>mode: AppMode.Automatic</C> there is no organizer key: pass no <C>skOrg</C>, the writer sends the
          identity and no Schnorr words, and <C>PK_aid</C> is the pool key alone. Committee nodes decrypt without
          anyone&rsquo;s intervention, so confidentiality of such an application rests on the committee threshold only.
          The <Em>decryption window</Em> is the way to bound that exposure: outside{' '}
          <C>decryptNotBefore → decryptNotAfter</C> the contract refuses partials and combines with{' '}
          <C>DecryptionNotOpen()</C> / <C>DecryptionClosed()</C>.
        </P>
        <Code caption='register-automatic.ts'>{`import { AppMode } from '@vocdoni/davinci-dkg-sdk'

const now = BigInt(Math.floor(Date.now() / 1000))
await writer.registerApplication(epochId, aid, {
  mode:             AppMode.Automatic,
  openSubmission:   true,
  decryptNotBefore: now + 3600n,              // opens in an hour
  decryptNotAfter:  now + 7n * 24n * 3600n,   // closes in a week
})

const app = await dkg.getApplication(epochId, aid)
app.policy.mode      // AppMode.Automatic
app.organizerPK      // [0n, 1n] — the identity, no organizer key
await dkg.isDecryptionOpen(epochId, aid) // false before it opens and after it closes`}</Code>
      </Section>

      <Section id='encrypt' title='Encrypting a value'>
        <P>
          ElGamal on BabyJubJub, entirely client-side. <C>getApplicationKey</C> resolves <C>PK_aid</C> for a registered
          application — its pool key, plus <C>PK_org</C> when organizer-locked — in TE form, and <C>encrypt</C>{' '}
          encrypts under it. <C>encryptForApplication</C> does the same from the two halves if you already hold them.
        </P>
        <Code caption='encrypt.ts'>{`import { encrypt, encryptForApplication } from '@vocdoni/davinci-dkg-sdk'

// P_j (automatic) or P_j + PK_org (organizer-locked), read off the contract.
const pkAid = await dkg.getApplicationKey(epochId, aid)

// 'message' is a non-negative integer below 2^50 (~1.13e15): the cap the
// committee's BSGS dlog can invert. Larger values are unrecoverable.
const ciphertext = await encrypt(42n, pkAid)
// → { c1: [x, y], c2: [x, y] }

// Equivalent, from the halves:
const app = await dkg.getApplication(epochId, aid)
const pj  = await dkg.getPoolKey(epochId, app.poolIndex)
const same = await encryptForApplication(42n, pj, app.policy.mode === AppMode.Automatic ? undefined : app.organizerPK)`}</Code>
        <P>
          There is no proof of knowledge of the randomness, because whoever submits an aggregated tally cannot know it.
          Cross-application replay is stopped by the per-application pool keys instead. The matching client-side{' '}
          <C>decrypt(ct, privKey)</C> helper — for tests and direct single-key recovery, not the threshold path — caps
          at <C>2³²</C> so its table stays browser-sized.
        </P>
      </Section>

      <Section id='submit' title='Submitting a ciphertext'>
        <P>
          <C>DKGWriter</C> extends <C>DKGClient</C> with the three organizer writes, plus calls an organizer normally
          never makes: <C>createEpoch</C> (permissionless but cadence-gated; the nodes fire it themselves),{' '}
          <C>activatePoolKey</C> (the nodes do that too) and <C>abortEpoch</C> (only accepted for a provably dead
          epoch). The contract assigns the ciphertext index, so <C>submitCiphertext</C> waits for the receipt and reads
          it back out of the event.
        </P>
        <Code caption='submit.ts'>{`const { hash, ciphertextIndex } = await writer.submitCiphertext(epochId, aid, ciphertext)
console.log('submitted in', hash, 'as ciphertext #', ciphertextIndex)`}</Code>
      </Section>

      <Section id='reveal' title='Revealing the organizer secret'>
        <P>
          The committee&rsquo;s partials are not enough on their own for an organizer-locked application: the combine
          proof needs <C>sk_org</C>, and it takes it from the chain. The organizer reveals it <Em>once</Em>, whenever it
          decides the plaintexts may exist — keeping it until a poll closes costs nothing but a decision. From that block
          on every ciphertext under the <C>aid</C>, past and future, is combined by the committee on its own; there is
          no per-ciphertext share. An automatic application skips this step entirely.
        </P>
        <Code caption='reveal.ts'>{`await writer.revealOrganizerSecret(epochId, aid, skOrg)

// No proof: the contract checks sk_org·G == PK_org (InvalidOrganizerSecret()
// otherwise), stores the scalar, emits OrganizerSecretRevealed(eid, aid, sk)
// and reverts AlreadyRevealed() on a second call.
const app = await dkg.getApplication(epochId, aid)
app.organizerSecret  // skOrg — public on chain from here on`}</Code>
        <P>
          Anyone may relay a reveal — the secret is what authenticates it — and it cannot be undone. Before sending,
          check the scalar against <C>organizerPublicKey(skOrg)</C> if it came from a paste: a wrong one costs a revert,
          the right one costs the application its lock.
        </P>
      </Section>

      <Section id='read' title='Reading the plaintext'>
        <P>
          With <C>t</C> partials on chain and the application unlocked — automatic, or revealed — any node combines
          them inside the decryption window and publishes the plaintext. <C>decryptionProgress</C> says whether the
          ciphertext and its combine are on chain; <C>waitForCombinedDecryption</C> polls until the record completes.
        </P>
        <Code caption='read.ts'>{`import { decryptionProgress, waitForCombinedDecryption } from '@vocdoni/davinci-dkg-sdk'

const progress = await decryptionProgress(dkg, epochId, aid, ciphertextIndex)
console.log(progress.combined ? 'combined' : 'waiting on t partials (and the reveal, if organizer-locked)')

const record = await waitForCombinedDecryption(dkg, epochId, aid, ciphertextIndex, {
  intervalMs: 3000,
  timeoutMs: 5 * 60_000,
})
if (record.completed) console.log('plaintext:', record.plaintext.toString())

// Or read it directly once it has landed.
const m = await dkg.getPlaintext(epochId, aid, ciphertextIndex)`}</Code>
        <Note>
          The <Internal to={paths.playground()}>playground</Internal> runs exactly this sequence in the browser, with
          the transcripts of each proof printed next to the transaction that carried them.
        </Note>
      </Section>

      <Section id='reference' title='Reference'>
        <Sub>This deployment</Sub>
        <KeyValue
          columns={2}
          items={[
            { label: 'network', value: config.chainName },
            { label: 'chain id', value: config.chainId, mono: true },
            { label: 'managerAddress', value: <Address value={manager} full /> },
            { label: 'rpc endpoint', value: config.rpcUrl, mono: true },
          ]}
        />
        <Sub>Further reading</Sub>
        <Bullets
          items={[
            <>
              <Ext href={`${REPO}/tree/main/sdk`}>SDK source</Ext> — every export documented inline.
            </>,
            <>
              <Ext href={`${REPO}/tree/main/sdk/tests`}>SDK tests</Ext> — copy-pasteable end-to-end usage.
            </>,
            <>
              <Ext href={`${REPO}/tree/main/solidity/src/interfaces`}>Contract interfaces</Ext> — the integration
              contract: method signatures and event schemas.
            </>,
            <>
              <Ext href='https://viem.sh/'>viem documentation</Ext>.
            </>,
          ]}
        />
      </Section>
    </DocsLayout>
  )
}
