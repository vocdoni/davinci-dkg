import { createSystem, defaultConfig, defineConfig } from '@chakra-ui/react'

// Minimal Chakra v3 system: dark-first, Inter for prose, JetBrains Mono for
// code. Recipes / semantic tokens can be added under ./recipes once we have
// concrete components asking for them — premature theming is a maintenance
// trap.
const config = defineConfig({
  globalCss: {
    'html, body, #root': {
      bg: 'gray.950',
      color: 'gray.100',
      minHeight: '100vh',
    },
  },
  theme: {
    tokens: {
      fonts: {
        body: { value: "Inter, system-ui, sans-serif" },
        heading: { value: "Inter, system-ui, sans-serif" },
        mono: { value: "'JetBrains Mono', ui-monospace, SFMono-Regular, Menlo, monospace" },
      },
    },
  },
})

export const system = createSystem(defaultConfig, config)
