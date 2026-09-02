import { useCallback, useState, type ReactNode } from 'react'

export interface ChartTooltipState {
  x: number
  y: number
  content: ReactNode
}

/**
 * Cursor-follow tooltip state for a chart. Coordinates are relative to the
 * chart's own container, so the caller subtracts the container rect once and
 * every mark reuses it.
 */
export function useChartTooltip() {
  const [tooltip, setTooltip] = useState<ChartTooltipState | null>(null)

  const show = useCallback(
    (event: { clientX: number; clientY: number; currentTarget: Element }, content: ReactNode) => {
      const host = event.currentTarget.closest('[data-chart-root]') ?? event.currentTarget
      const rect = host.getBoundingClientRect()
      setTooltip({ x: event.clientX - rect.left, y: event.clientY - rect.top, content })
    },
    []
  )

  const hide = useCallback(() => setTooltip(null), [])

  return { tooltip, show, hide }
}
