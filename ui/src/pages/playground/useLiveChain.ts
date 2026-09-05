// The live chain: a wagmi wallet plus the SDK's `DKGWriter`.
//
// This is the only file in the playground that imports wagmi or RainbowKit —
// demo mode renders a different component tree and never mounts it, so
// `?demo=1` cannot end up asking a browser extension for anything.
//
// Reads of the pipeline do not go through the SDK's client: every piece of
// state the stepper shows (the pool, the ciphertext the contract stored, the
// partials, the reveal, the combine) is already in the explorer's indexer,
// and reading it from there keeps the playground on the same snapshot as the
// rest of the app. The application key is the one exception — it is read
// straight off the contract with `getApplicationKey`, so the value encrypted
// under is exactly what the chain says, not what the indexer has caught up to.
// Only the three organizer writes go through the writer.

import { useCallback, useMemo } from 'react'
import { useAccount, useChainId, usePublicClient, useWalletClient } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { AppMode, DKGClient, DKGWriter, type ElGamalCiphertext, fromRTEtoTE } from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import { useRuntimeConfig } from '~config/config-context'
import { useApplication, useEpoch, useIndexer, useStore } from '~data/hooks'
import { POOL_SIZE } from '~indexer/types'
import type {
  ApplicationKeys,
  DecryptionView,
  PlaygroundChain,
  PlaygroundTarget,
  RegisterArgs,
  RegisterResult,
  RevealArgs,
  SubmitArgs,
  TxRecord,
} from './types'
type WriterConfig = ConstructorParameters<typeof DKGWriter>[0]
type ClientConfig = ConstructorParameters<typeof DKGClient>[0]

export function useLiveChain(target: PlaygroundTarget): PlaygroundChain {
  const config = useRuntimeConfig()
  const { address, isConnected } = useAccount()
  const walletChainId = useChainId()
  const publicClient = usePublicClient()
  const { data: walletClient } = useWalletClient()
  const { openConnectModal } = useConnectModal()
  const { refresh, headBlock } = useIndexer()
  const staggerBlocks = useStore().chain.staggerBlocks || 3

  const epoch = useEpoch(target.epochId ?? undefined)
  const application = useApplication(target.epochId ?? undefined, target.aid ?? undefined)

  // Reads need no wallet; writes do.
  const client = useMemo(() => {
    if (!publicClient) return null
    return new DKGClient({
      publicClient: publicClient as unknown as ClientConfig['publicClient'],
      managerAddress: config.managerAddress,
    })
  }, [publicClient, config.managerAddress])

  const writer = useMemo(() => {
    if (!publicClient || !walletClient?.account) return null
    return new DKGWriter({
      publicClient: publicClient as unknown as WriterConfig['publicClient'],
      walletClient: walletClient as unknown as WriterConfig['walletClient'],
      managerAddress: config.managerAddress,
    })
  }, [publicClient, walletClient, config.managerAddress])

  /** Wait for the receipt so the step can show a real block and gas figure. */
  const settle = useCallback(
    async (hash: Hex): Promise<TxRecord> => {
      const receipt = await publicClient?.waitForTransactionReceipt({ hash })
      void refresh()
      return {
        hash,
        block: receipt ? Number(receipt.blockNumber) : null,
        gasUsed: receipt ? Number(receipt.gasUsed) : null,
        simulated: false,
      }
    },
    [publicClient, refresh]
  )

  const register = useCallback(
    async ({
      aid,
      skOrg,
      mode,
      submitters,
      maxCiphertexts,
      decryptNotBefore,
      decryptNotAfter,
      nonce,
    }: RegisterArgs): Promise<RegisterResult> => {
      if (!writer || !client || !target.epochId) throw new Error('Connect a wallet first')
      // Everything not named here takes the writer's default: no open
      // submission, no block window.
      const hash = await writer.registerApplication(
        target.epochId,
        aid,
        {
          mode: mode === 'automatic' ? AppMode.Automatic : AppMode.OrganizerLocked,
          submitters,
          maxCiphertexts,
          decryptNotBefore: BigInt(decryptNotBefore),
          decryptNotAfter: BigInt(decryptNotAfter),
        },
        skOrg ?? undefined,
        nonce
      )
      const tx = await settle(hash)
      const poolIndex = await client.getAppPoolIndex(target.epochId, aid)
      return { tx, poolIndex }
    },
    [writer, client, target.epochId, settle]
  )

  const applicationKeys = useCallback(
    async (aid: Hex): Promise<ApplicationKeys> => {
      if (!client || !target.epochId) throw new Error('No RPC client yet')
      const app = await client.getApplication(target.epochId, aid)
      if (!app.exists) throw new Error('The application is not registered yet')
      const [poolKey, key] = await Promise.all([
        client.getPoolKey(target.epochId, app.poolIndex),
        client.getApplicationKey(target.epochId, aid),
      ])
      return {
        poolIndex: app.poolIndex,
        poolKey,
        organizerPK: app.policy.mode === AppMode.Automatic ? null : app.organizerPK,
        key,
      }
    },
    [client, target.epochId]
  )

  const submitCiphertext = useCallback(
    async ({ aid, ciphertext }: SubmitArgs) => {
      if (!writer || !target.epochId) throw new Error('Connect a wallet first')
      const result = await writer.submitCiphertext(target.epochId, aid, ciphertext)
      void refresh()
      return {
        tx: {
          hash: result.hash,
          block: Number(result.receipt.blockNumber),
          gasUsed: Number(result.receipt.gasUsed),
          simulated: false,
        },
        ciphertextIndex: result.ciphertextIndex,
      }
    },
    [writer, target.epochId, refresh]
  )

  const revealSecret = useCallback(
    async ({ aid, skOrg }: RevealArgs): Promise<TxRecord> => {
      if (!writer || !target.epochId) throw new Error('Connect a wallet first')
      const hash = await writer.revealOrganizerSecret(target.epochId, aid, skOrg)
      return settle(hash)
    },
    [writer, target.epochId, settle]
  )

  const decryption = useCallback(
    (ciphertextIndex: number | null): DecryptionView | null => {
      if (ciphertextIndex == null || !application) return null
      const row = application.ciphertexts.find((ct) => ct.index === ciphertextIndex)
      if (!row) return null
      const app = application.row
      return {
        threshold: row.threshold,
        committeeSize: row.committeeSize,
        staggerBlocks,
        ciphertextBlock: row.block,
        partials: row.partials.map((partial) => ({
          participantIndex: partial.participantIndex,
          participant: partial.participant,
          block: partial.block,
          wave: partial.wave,
          tx: partial.tx,
        })),
        reveal: {
          required: app.mode === 'organizer-locked',
          done: app.unlocked,
          block: app.revealBlock,
          tx: app.revealTx,
        },
        combined: {
          done: row.combined.done,
          block: row.combined.block,
          tx: row.combined.tx,
          plaintext: row.combined.plaintext,
        },
        // The store keeps the calldata (RTE) words; the local ciphertext is TE.
        onChain: { c1: fromRTEtoTE(row.c1.x, row.c1.y), c2: fromRTEtoTE(row.c2.x, row.c2.y) } as ElGamalCiphertext,
      }
    },
    [application, staggerBlocks]
  )

  const wrongChain = isConnected && walletChainId !== config.chainId
  const problem = !isConnected
    ? 'No wallet connected'
    : wrongChain
      ? `Wallet is on chain ${walletChainId}; this deployment is chain ${config.chainId}`
      : !writer
        ? 'Wallet client is not ready yet'
        : null

  const pool = epoch?.finalization
    ? { activated: epoch.poolActivated, claimed: epoch.poolClaimed, size: POOL_SIZE }
    : null
  const poolActivated = pool?.activated ?? -1
  const poolClaimed = pool?.claimed ?? -1

  return useMemo<PlaygroundChain>(
    () => ({
      kind: 'live',
      headBlock,
      pool: poolActivated < 0 ? null : { activated: poolActivated, claimed: poolClaimed, size: POOL_SIZE },
      wallet: {
        connected: Boolean(isConnected && address),
        address: address ?? null,
        label: address ?? 'wallet',
        connect: () => openConnectModal?.(),
        chainOk: !wrongChain,
        problem,
      },
      register,
      applicationKeys,
      submitCiphertext,
      revealSecret,
      decryption,
    }),
    [
      headBlock,
      poolActivated,
      poolClaimed,
      isConnected,
      address,
      openConnectModal,
      wrongChain,
      problem,
      register,
      applicationKeys,
      submitCiphertext,
      revealSecret,
      decryption,
    ]
  )
}
