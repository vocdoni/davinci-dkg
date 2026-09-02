import '@testing-library/jest-dom/vitest'

// jsdom implements neither of these, and both are touched as soon as a kit
// component mounts: Radix reads `matchMedia`, and every chart measures its
// container with a ResizeObserver. Stubbing them here keeps component tests to
// a bare `render(<Thing />)` with no per-file boilerplate.
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
