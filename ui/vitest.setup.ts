import '@testing-library/jest-dom/vitest'

// jsdom implements neither of these, and both are touched as soon as a
// component tree with the theme provider mounts: next-themes reads
// `matchMedia` to follow the OS colour scheme, and the charts measure their
// container with a ResizeObserver. Stubbing them here keeps component tests to
// `render(<Theme>…)` with no per-file boilerplate.
if (typeof window !== 'undefined') {
  if (!window.matchMedia) {
    window.matchMedia = ((query: string) => ({
      matches: false,
      media: query,
      onchange: null,
      addListener: () => {},
      removeListener: () => {},
      addEventListener: () => {},
      removeEventListener: () => {},
      dispatchEvent: () => false,
    })) as typeof window.matchMedia
  }
  if (!window.ResizeObserver) {
    window.ResizeObserver = class {
      observe() {}
      unobserve() {}
      disconnect() {}
    } as unknown as typeof window.ResizeObserver
  }
}

// jsdom's CSS parser chokes on the `@layer` rules Chakra emits and dumps the
// entire stylesheet to stderr for every component test. Nothing is broken —
// jsdom simply skips the rule — so drop that one message and let everything
// else through.
const consoleError = console.error
console.error = (...args: unknown[]) => {
  const first = args[0]
  const message = first instanceof Error ? first.message : String(first ?? '')
  if (message.includes('Could not parse CSS stylesheet')) return
  consoleError(...args)
}
