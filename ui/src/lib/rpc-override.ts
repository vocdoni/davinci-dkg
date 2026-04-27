// localStorage-backed RPC override. Same key as the legacy webapp so users
// who set an override there carry it over after the cutover.

const KEY = 'dkg-explorer:rpc-url'

export function readRpcOverride(): string | null {
  try {
    return localStorage.getItem(KEY)
  } catch {
    return null
  }
}

export function writeRpcOverride(url: string | null) {
  try {
    if (url && url.trim() !== '') {
      localStorage.setItem(KEY, url.trim())
    } else {
      localStorage.removeItem(KEY)
    }
  } catch {
    // private browsing; the override just won't persist.
  }
}
