import { darkTheme, type Theme } from '@rainbow-me/rainbowkit'

const base = darkTheme({
  accentColor: '#00d992',
  accentColorForeground: '#050507',
  borderRadius: 'small',
  overlayBlur: 'small',
})

/**
 * RainbowKit's modal, dressed in the explorer's tokens: obsidian scrim, carbon
 * panel, charcoal hairlines, emerald accent, Inter. Only the colours the modal
 * actually shows are overridden — the rest of RainbowKit's dark theme is
 * already close enough.
 */
export const walletTheme: Theme = {
  ...base,
  fonts: { body: "'Inter Variable', Inter, system-ui, sans-serif" },
  radii: {
    ...base.radii,
    modal: '8px',
    menuButton: '6px',
    actionButton: '6px',
    connectButton: '6px',
    modalMobile: '8px',
  },
  colors: {
    ...base.colors,
    accentColor: '#00d992',
    accentColorForeground: '#050507',
    modalBackground: '#101010',
    modalBorder: '#3d3a39',
    modalText: '#ffffff',
    modalTextSecondary: '#8a8380',
    modalTextDim: '#5c5855',
    menuItemBackground: '#1a1a1a',
    actionButtonBorder: '#3d3a39',
    actionButtonBorderMobile: '#3d3a39',
    actionButtonSecondaryBackground: '#1a1a1a',
    closeButton: '#b8b3b0',
    closeButtonBackground: '#1a1a1a',
    generalBorder: '#3d3a39',
    profileForeground: '#101010',
    profileAction: '#1a1a1a',
    profileActionHover: '#3d3a39',
    connectButtonBackground: '#101010',
    connectButtonInnerBackground: '#1a1a1a',
    connectButtonText: '#ffffff',
    connectButtonTextError: '#f85149',
    error: '#f85149',
    standby: '#e3b341',
  },
}
