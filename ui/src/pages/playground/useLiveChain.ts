// The live chain: a wagmi wallet plus the SDK's `DKGWriter`.
//
// This is the only file in the playground that imports wagmi or RainbowKit —
// demo mode renders a different component tree and never mounts it, so
// `?demo=1` cannot end up asking a browser extension for anything.
//
// Reads do not go through the SDK's client: every piece of state the stepper
// shows (the epoch key, the ciphertext the contract stored, the partials,
// the share, the combine) is already in the explorer's indexer, and reading it
// from there keeps the playground on the same snapshot as the rest of the app.
// Only the three organizer writes go through the writer.

import { useCallback, useMemo } from 'react'
import { useAccount, useChainId, usePublicClient, useWalletClient } from 'wagmi'
import { useConnectModal } from '@rainbow-me/rainbowkit'
import { DKGWriter, type ElGamalCiphertext, fromRTEtoTE } from '@vocdoni/davinci-dkg-sdk'
import type { Hex } from 'viem'
import { useRuntimeConfig } from '~config/config-context'
import { useApplication, useEpoch, useIndexer, useStore } from '~data/hooks'
import type {
  DecryptionView,
  PlaygroundChain,
  PlaygroundTarget,
  RegisterArgs,
  ShareArgs,
  SubmitArgs,
  TxRecord,
} from './types'
type WriterConfig = ConstructorParameters<typeof DKGWriter>[0]

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
    async ({ aid, skOrg, authorizedSubmitter, maxCiphertexts, nonce }: RegisterArgs): Promise<TxRecord> => {
      if (!writer || !target.epochId) throw new Error('Connect a wallet first')
      const hash = await writer.registerApplication(
        target.epochId,
        aid,
        {
          authorizedSubmitter,
          maxCiphertexts,
          notBeforeBlock: 0n,
          notAfterBlock: 0n,
        },
        skOrg,
        nonce
      )
      return settle(hash)
    },
    [writer, target.epochId, settle]
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

  const releaseShare = useCallback(
    async ({ aid, ciphertext, ciphertextIndex, skOrg, nonce }: ShareArgs): Promise<TxRecord> => {
      if (!writer || !target.epochId) throw new Error('Connect a wallet first')
      const hash = await writer.submitOrganizerShare(
        target.epochId,
        aid,
        ciphertextIndex,
        ciphertext,
        skOrg,
        nonce
      )
      return settle(hash)
    },
    [writer, target.epochId, settle]
  )

  const decryption = useCallback(
    (ciphertextIndex: number | null): DecryptionView | null => {
      if (ciphertextIndex == null || !application) return null
      const row = application.ciphertexts.find((ct) => ct.index === ciphertextIndex)
      if (!row) return null
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
        share: { present: row.share.present, block: row.share.block, tx: row.share.tx },
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

  const epochKey = epoch?.collectivePublicKey ?? null

  return useMemo<PlaygroundChain>(
    () => ({
      kind: 'live',
      headBlock,
      epochKey: epochKey ? [epochKey.x, epochKey.y] : null,
      wallet: {
        connected: Boolean(isConnected && address),
        address: address ?? null,
        label: address ?? 'wallet',
        connect: () => openConnectModal?.(),
        chainOk: !wrongChain,
        problem,
      },
      register,
      submitCiphertext,
      releaseShare,
      decryption,
    }),
    [
      headBlock,
      epochKey,
      isConnected,
      address,
      openConnectModal,
      wrongChain,
      problem,
      register,
      submitCiphertext,
      releaseShare,
      decryption,
    ]
  )
}
