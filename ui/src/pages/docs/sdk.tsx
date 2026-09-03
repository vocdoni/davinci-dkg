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
  { id: 'share', title: 'Releasing the organizer share' },
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
          and later answer ciphertexts with partial decryptions. <Em>Organizers</Em> never create or schedule an epoch.
          They register an application against whichever epoch is already Live, encrypt under{' '}
          <C>PK_aid = PK_ep + PK_org</C>, submit ciphertexts, and release their own decryption share when they decide
          the plaintext may exist.
        </P>
        <P>
          Only registration carries a proof of possession of <C>sk_org</C> — a Schnorr proof, verified on chain.
          Submitting a ciphertext proves nothing, and the organizer share is a keccak-challenge Chaum&ndash;Pedersen
          discrete-logarithm equality (DLEQ) proof computed client-side. An organizer therefore needs no circuit
          artifacts and no prover at all: everything expensive lives on the committee&rsquo;s side.
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
          the newest epoch whose status is <C>Live</C>, which the operator set has already produced and finalized.
        </P>
        <Code caption='epoch.ts'>{`import { buildEpochId, EpochPhase, roundStatusLabel } from '@vocdoni/davinci-dkg-sdk'

const nonce   = await dkg.epochNonce()
const prefix  = await dkg.roundPrefix()
const epochId = buildEpochId(prefix, nonce - 1n)

const epoch = await dkg.getEpoch(epochId)
console.log(roundStatusLabel(epoch.status), epoch.policy.threshold, epoch.policy.committeeSize)

// Once Live, the collective key is a plain view call.
if (epoch.status === EpochPhase.Live) {
  // Twisted-Edwards (TE) coordinates, the form every contract call uses.
  const pkEp = await dkg.getCollectivePublicKey(epochId)   // { x, y }
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
  watchCiphertextSubmitted,
} from '@vocdoni/davinci-dkg-sdk'

const stop = watchNewEpochs(dkg, (epochId, organizer) => { /* … */ })

watchEpochLive(dkg, epochId, (collectivePublicKeyHash) => { /* … */ })

// Ciphertexts as they land. c1/c2 come back in TE form, ready to hand
// straight to writer.submitOrganizerShare.
watchCiphertextSubmitted(dkg, epochId, (event) => {
  console.log(event.aid, event.ciphertextIndex, event.c1, event.c2)
})

stop()`}</Code>
      </Section>

      <Section id='register' title='Registering an application'>
        <P>
          Every ciphertext belongs to a registered application, and an unregistered <C>aid</C> reverts. Registration
          binds an organizer key so the encryption key becomes <C>PK_aid = PK_ep + PK_org</C>. Only <C>PK_org</C> and
          the Schnorr proof of possession go on chain — <C>sk_org</C> never leaves the process that drew it.
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
  authorizedSubmitter: '0x0000000000000000000000000000000000000000', // = the caller
  maxCiphertexts:      0,     // 0 = unlimited
  notBeforeBlock:      0n,
  notAfterBlock:       0n,
}, skOrg)`}</Code>
        <Note tone='warn'>
          <Em>Keep sk_org.</Em> It is the application&rsquo;s only decryption capability, it is never transmitted, and
          nothing on chain can reconstruct it. Lose it and every ciphertext ever submitted under that <C>aid</C> is
          permanently undecryptable — the committee threshold alone cannot open them, by design.
        </Note>
      </Section>

      <Section id='encrypt' title='Encrypting a value'>
        <P>
          ElGamal on BabyJubJub, entirely client-side. <C>encryptForApplication</C> takes the epoch key and the
          application&rsquo;s <C>PK_org</C> — both readable from the client, both in TE form — and encrypts under their
          sum.
        </P>
        <Code caption='encrypt.ts'>{`import { encryptForApplication } from '@vocdoni/davinci-dkg-sdk'

const pkEp = await dkg.getCollectivePublicKey(epochId)
const app  = await dkg.getApplication(epochId, aid)

// 'message' is a non-negative integer below 2^50 (~1.13e15): the cap the
// committee's BSGS dlog can invert. Larger values are unrecoverable.
const ciphertext = await encryptForApplication(42n, [pkEp.x, pkEp.y], app.organizerPK)
// → { c1: [x, y], c2: [x, y] }`}</Code>
        <P>
          There is no proof of knowledge of the randomness, because whoever submits an aggregated tally cannot know it.
          Cross-application replay is stopped by the organizer key instead. The matching client-side{' '}
          <C>decrypt(ct, privKey)</C> helper — for tests and direct single-key recovery, not the threshold path — caps
          at <C>2³²</C> so its table stays browser-sized.
        </P>
      </Section>

      <Section id='submit' title='Submitting a ciphertext'>
        <P>
          <C>DKGWriter</C> extends <C>DKGClient</C> with the three organizer writes, plus two calls an organizer
          normally never makes: <C>createEpoch</C> (permissionless but cadence-gated; the nodes fire it themselves) and{' '}
          <C>abortEpoch</C> (only accepted for a provably dead epoch). The contract assigns the ciphertext index, so{' '}
          <C>submitCiphertext</C> waits for the receipt and reads it back out of the event.
        </P>
        <Code caption='submit.ts'>{`const { hash, ciphertextIndex } = await writer.submitCiphertext(epochId, aid, ciphertext)
console.log('submitted in', hash, 'as ciphertext #', ciphertextIndex)`}</Code>
      </Section>

      <Section id='share' title='Releasing the organizer share'>
        <P>
          The committee&rsquo;s partials are not enough on their own. The organizer publishes <C>Δ = sk_org·C1</C> with
          a Chaum&ndash;Pedersen DLEQ proving the same secret relates <C>(G, PK_org)</C> and <C>(C1, Δ)</C>. Until that
          lands, <C>combineDecryption</C> reverts <C>OrganizerShareMissing()</C> — so withholding the share until a poll
          closes costs nothing but a decision.
        </P>
        <Code caption='share.ts'>{`await writer.submitOrganizerShare(epochId, aid, ciphertextIndex, ciphertext, skOrg)

// The challenge is a keccak256 the contract recomputes and the combine
// circuit consumes:
//   e = keccak256(DOMAIN_ORGANIZER_SHARE_V1 ‖ eid ‖ aid ‖ uint256(ctIdx)
//                 ‖ PK_org ‖ C1 ‖ Δ ‖ A1 ‖ A2) mod q
//   z = w + e·sk_org mod q
// The contract stores only keccak256(Δ ‖ A1 ‖ A2 ‖ z); the DLEQ itself is
// verified inside the committee's combine proof.`}</Code>
        <P>
          Anyone may relay a share — it is self-authenticating — and re-submission overwrites until the ciphertext is
          combined, so a malformed share cannot brick a ciphertext. <C>verifyOrganizerShare</C> is exported if you want
          to check one before or after it goes on chain.
        </P>
      </Section>

      <Section id='read' title='Reading the plaintext'>
        <P>
          With the threshold of partials and the organizer share both on chain, any node combines them and publishes the
          plaintext. <C>decryptionProgress</C> says which of the two halves is still missing;{' '}
          <C>waitForCombinedDecryption</C> polls until the record completes.
        </P>
        <Code caption='read.ts'>{`import { decryptionProgress, waitForCombinedDecryption } from '@vocdoni/davinci-dkg-sdk'

const progress = await decryptionProgress(dkg, epochId, aid, ciphertextIndex)
console.log(progress.organizerShare ? 'waiting on the committee' : 'waiting on the organizer')

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
