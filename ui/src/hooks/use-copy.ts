import { useCallback, useEffect, useRef, useState } from 'react'

/**
 * Copy-to-clipboard with a 1.5 s "copied" flash. Falls back to a hidden
 * textarea + `execCommand` because the async clipboard API is unavailable on
 * insecure origins — and a node operator inspecting an explorer over plain
 * http:// on a LAN is a normal thing to do.
 */
export function useCopy(resetMs = 1500): { copied: boolean; copy: (value: string) => void } {
  const [copied, setCopied] = useState(false)
  const timer = useRef<ReturnType<typeof setTimeout> | null>(null)

  useEffect(
    () => () => {
      if (timer.current) clearTimeout(timer.current)
    },
    []
  )

  const flash = useCallback(() => {
    setCopied(true)
    if (timer.current) clearTimeout(timer.current)
    timer.current = setTimeout(() => setCopied(false), resetMs)
  }, [resetMs])

  const copy = useCallback(
    (value: string) => {
      const legacy = () => {
        const el = document.createElement('textarea')
        el.value = value
        el.setAttribute('readonly', '')
        el.style.position = 'fixed'
        el.style.opacity = '0'
        document.body.appendChild(el)
        el.select()
        try {
          document.execCommand('copy')
          flash()
        } finally {
          document.body.removeChild(el)
        }
      }
      if (navigator.clipboard?.writeText) {
        navigator.clipboard.writeText(value).then(flash, legacy)
        return
      }
      legacy()
    },
    [flash]
  )

  return { copied, copy }
}
