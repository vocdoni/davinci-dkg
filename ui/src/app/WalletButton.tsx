import { ConnectButton } from '@rainbow-me/rainbowkit'
import { useRuntimeConfig } from '~config/config-context'
import { Badge, Button } from '~kit'

/**
 * Wallet connect, wearing the kit's Button rather than RainbowKit's own skin.
 * In demo mode there is no chain to connect to, so it says so instead.
 */
export function WalletButton() {
  const { demo } = useRuntimeConfig()
  if (demo) {
    return (
      <Badge tone='warn' title='Demo mode runs off the synthetic fixture; no wallet is connected'>
        demo wallet
      </Badge>
    )
  }

  return (
    <ConnectButton.Custom>
      {({ account, chain, openAccountModal, openChainModal, openConnectModal, mounted }) => {
        const ready = mounted
        const connected = ready && account && chain
        return (
          <div aria-hidden={!ready} className={ready ? undefined : 'pointer-events-none opacity-0'}>
            {!connected ? (
              <Button variant='primary' size='sm' onClick={openConnectModal}>
                Connect
              </Button>
            ) : chain.unsupported ? (
              <Button variant='danger' size='sm' onClick={openChainModal}>
                Wrong network
              </Button>
            ) : (
              <Button variant='secondary' size='sm' onClick={openAccountModal} className='font-mono'>
                {account.displayName}
              </Button>
            )}
          </div>
        )
      }}
    </ConnectButton.Custom>
  )
}
