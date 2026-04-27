import { useState } from 'react'
import { Button, HStack, Text } from '@chakra-ui/react'
import { LuClipboardCopy, LuCheck } from 'react-icons/lu'
import { useLocation } from 'react-router-dom'
import { useAccount } from 'wagmi'
import { useConfig } from '~providers/ConfigProvider'
import { buildErrorReport } from '~lib/error-report'

interface Props {
  error: unknown
  /** Extra context (e.g. roundId on a round detail page). */
  extra?: { roundId?: string; blockNumber?: bigint | number }
}

// One-click "Copy error report". Bundles route, chain, wallet, and any
// caller-supplied extras into a markdown blob the user can paste into a
// GitHub issue. Dramatically lowers the friction of bug reporting and
// almost always gives the maintainer enough to reproduce.
export function ErrorReportButton({ error, extra }: Props) {
  const config = useConfig()
  const { address } = useAccount()
  const location = useLocation()
  const [copied, setCopied] = useState(false)

  const onCopy = async () => {
    const report = buildErrorReport(error, {
      route: location.pathname + location.search,
      chainId: config.chainId,
      chainName: config.chainName,
      walletAddress: address,
      roundId: extra?.roundId,
      blockNumber: extra?.blockNumber,
      buildVersion: import.meta.env.VITE_BUILD_VERSION as string | undefined,
    })
    try {
      await navigator.clipboard.writeText(report)
      setCopied(true)
      setTimeout(() => setCopied(false), 1500)
    } catch {
      // fall back to logging — clipboard can be denied in some contexts
      console.warn('Clipboard write failed; printing report to console:\n', report)
    }
  }

  return (
    <Button size='xs' variant='outline' colorPalette='red' onClick={onCopy}>
      <HStack gap={1.5}>
        {copied ? <LuCheck /> : <LuClipboardCopy />}
        <Text fontSize='xs'>{copied ? 'Copied' : 'Copy error report'}</Text>
      </HStack>
    </Button>
  )
}
