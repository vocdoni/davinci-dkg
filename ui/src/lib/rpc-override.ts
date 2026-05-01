// localStorage-backed RPC override. Lets users override the bundled RPC
// endpoint without rebuilding the image.

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
