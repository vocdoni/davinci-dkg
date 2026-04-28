import { createSystem, defaultConfig, defineConfig } from '@chakra-ui/react'

// "Quiet protocol tool" theme.
//
// Two families only — Manrope for everything UI/prose, JetBrains Mono
// for code/numbers/hashes. Single warm-gold accent used sparingly. The
// rest of the palette is bone-on-ink.
const config = defineConfig({
  globalCss: {
    'html, body, #root': {
      bg: 'canvas',
      color: 'ink.1',
      minHeight: '100vh',
      colorScheme: 'dark',
      fontFamily: 'body',
      fontFeatureSettings: '"liga" 1, "calt" 1, "kern" 1',
    },
    '::selection': {
      bg: 'rgba(212, 165, 116, 0.28)',
      color: 'ink.0',
    },
    '*::-webkit-scrollbar': { width: '10px', height: '10px' },
    '*::-webkit-scrollbar-track': { bg: 'transparent' },
    '*::-webkit-scrollbar-thumb': {
      bg: 'rgba(255, 255, 255, 0.06)',
      borderRadius: '999px',
      border: '2px solid transparent',
      backgroundClip: 'padding-box',
    },
    '*::-webkit-scrollbar-thumb:hover': {
      bg: 'rgba(255, 255, 255, 0.12)',
    },
    // Tabular figures for fixed-width digit contexts (block heights,
    // counters, timestamps). Stops the page from twitching as numbers
    // tick.
    '.dkg-tabular': {
      fontVariantNumeric: 'tabular-nums',
      fontFeatureSettings: '"tnum" 1',
    },
    // Hairline divider used by the header + footer + section breaks.
    '.dkg-rule': {
      height: '1px',
      width: '100%',
      bg: 'rule',
      borderRadius: '999px',
    },
    // Note: keyframes (`dkgPhasePulse`) live in src/theme/global.css —
    // Chakra v3's globalCss type-checker is too strict to accept them
    // as plain object keys.
  },
  theme: {
    tokens: {
      fonts: {
        body: { value: "Manrope, system-ui, -apple-system, sans-serif" },
        heading: { value: "Manrope, system-ui, -apple-system, sans-serif" },
        sans: { value: "Manrope, system-ui, -apple-system, sans-serif" },
        mono: { value: "'JetBrains Mono', ui-monospace, SFMono-Regular, Menlo, monospace" },
      },
      colors: {
        ink: {
          950: { value: '#06070a' },
          900: { value: '#0a0a0c' },
          850: { value: '#0f0f13' },
          800: { value: '#13131a' },
          700: { value: '#191921' },
          600: { value: '#22222b' },
          500: { value: '#2a2a35' },
          400: { value: '#3a3a47' },
        },
        bone: {
          50: { value: '#f6f1e7' },
          100: { value: '#ece8df' },
          200: { value: '#d9d3c4' },
          300: { value: '#b8b3a6' },
          400: { value: '#8d887a' },
          500: { value: '#6b6759' },
          600: { value: '#4a4639' },
        },
        gold: {
          50: { value: '#faecd0' },
          100: { value: '#f0d9aa' },
          200: { value: '#e5c184' },
          300: { value: '#d4a574' },
          400: { value: '#b8884e' },
          500: { value: '#8a6438' },
          600: { value: '#5e4426' },
          700: { value: '#3a2a18' },
        },
        phosphor: {
          200: { value: '#bef0d4' },
          300: { value: '#86efac' },
          400: { value: '#4ade80' },
          500: { value: '#22c55e' },
          700: { value: '#15803d' },
        },
        coral: {
          200: { value: '#f5b8a8' },
          300: { value: '#f0846a' },
          400: { value: '#e0664c' },
          500: { value: '#b54632' },
        },
        amber: {
          300: { value: '#f0c674' },
          400: { value: '#d4a045' },
        },
      },
    },
    semanticTokens: {
      colors: {
        canvas: { value: '{colors.ink.900}' },
        'canvas.deep': { value: '{colors.ink.950}' },
        surface: { value: '{colors.ink.800}' },
        'surface.raised': { value: '{colors.ink.700}' },
        'surface.sunken': { value: '{colors.ink.850}' },
        'surface.mono': { value: '#0c0c10' },
        border: { value: '{colors.ink.600}' },
        'border.subtle': { value: '{colors.ink.700}' },
        'border.strong': { value: '{colors.ink.500}' },
        rule: { value: 'rgba(236, 232, 223, 0.08)' },
        'ink.0': { value: '{colors.bone.50}' },
        'ink.1': { value: '{colors.bone.100}' },
        'ink.2': { value: '{colors.bone.300}' },
        'ink.3': { value: '{colors.bone.400}' },
        'ink.4': { value: '{colors.bone.500}' },
        'accent.fg': { value: '{colors.gold.300}' },
        'accent.bright': { value: '{colors.gold.200}' },
        'accent.dim': { value: '{colors.gold.500}' },
        'accent.deep': { value: '{colors.gold.700}' },
        'accent.bg': { value: 'rgba(212, 165, 116, 0.08)' },
        'accent.bg.strong': { value: 'rgba(212, 165, 116, 0.14)' },
        'accent.border': { value: 'rgba(212, 165, 116, 0.30)' },
        'live.fg': { value: '{colors.phosphor.300}' },
        'live.bright': { value: '{colors.phosphor.200}' },
        'live.dim': { value: '{colors.phosphor.700}' },
        'live.bg': { value: 'rgba(134, 239, 172, 0.08)' },
        'danger.fg': { value: '{colors.coral.300}' },
        'danger.bg': { value: 'rgba(240, 132, 106, 0.08)' },
        'danger.border': { value: 'rgba(240, 132, 106, 0.30)' },
        'warn.fg': { value: '{colors.amber.300}' },
        'warn.bg': { value: 'rgba(240, 198, 116, 0.08)' },
      },
      shadows: {
        inset: { value: 'inset 0 1px 0 0 rgba(255, 255, 255, 0.03)' },
      },
    },
  },
})

export const system = createSystem(defaultConfig, config)
