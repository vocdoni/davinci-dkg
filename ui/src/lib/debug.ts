// Tiny localStorage-backed flag for debug mode. Living outside React so
// modules that aren't components (loaders, error helpers) can read it too
// without going through context.

const KEY = 'dkg-ui:debug'

export function readDebugMode(): boolean {
  try {
    return localStorage.getItem(KEY) === '1'
  } catch {
    return false
  }
}

export function writeDebugMode(enabled: boolean) {
  try {
    if (enabled) localStorage.setItem(KEY, '1')
    else localStorage.removeItem(KEY)
  } catch {
    // Private browsing mode etc. — debug mode just won't persist.
  }
}
