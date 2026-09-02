import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import '@fontsource-variable/inter'
import '@fontsource/jetbrains-mono/400.css'
import '@fontsource/jetbrains-mono/500.css'
import './styles/index.css'
import { App } from './App'

const rootEl = document.getElementById('root')
if (!rootEl) throw new Error('No #root element in index.html')

createRoot(rootEl).render(
  <StrictMode>
    <App />
  </StrictMode>
)
