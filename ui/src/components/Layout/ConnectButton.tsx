import { ConnectButton as RkConnectButton } from '@rainbow-me/rainbowkit'

// Re-export point for the wallet connect button. Encapsulating the import
// here means a future swap of the underlying wallet UI (RainbowKit → AppKit
// → custom) only touches this file.
export function ConnectButton() {
  return <RkConnectButton chainStatus='icon' showBalance={false} accountStatus={{ smallScreen: 'avatar', largeScreen: 'address' }} />
}
