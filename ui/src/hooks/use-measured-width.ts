import { useCallback, useEffect, useRef, useState } from 'react'

/**
 * Measures the rendered width of a element and keeps it in state.
 *
 * The charts draw in SVG user units that map 1:1 to CSS pixels, so they need
 * the real container width rather than a viewBox scale — scaling a viewBox
 * would blow up the type along with the geometry, and a 10px axis label is
 * 10px whatever the container does.
 *
 * Returns `[ref, width]`; width is `null` until the first measurement, which
 * is the caller's cue to render a skeleton of the right height instead of a
 * chart at a guessed size.
 */
export function useMeasuredWidth<T extends HTMLElement = HTMLDivElement>(): [(node: T | null) => void, number | null] {
  const [width, setWidth] = useState<number | null>(null)
  const observer = useRef<ResizeObserver | null>(null)

  const ref = useCallback((node: T | null) => {
    observer.current?.disconnect()
    if (!node) return
    setWidth(node.getBoundingClientRect().width)
    // jsdom (and very old browsers) have no ResizeObserver; the single
    // measurement above is still enough to render.
    if (typeof ResizeObserver === 'undefined') return
    observer.current = new ResizeObserver((entries) => {
      const w = entries[0]?.contentRect.width
      if (w != null) setWidth(w)
    })
    observer.current.observe(node)
  }, [])

  useEffect(() => () => observer.current?.disconnect(), [])

  return [ref, width]
}
