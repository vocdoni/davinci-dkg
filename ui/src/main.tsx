import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { App } from './App'
import './theme/global.css'

const rootEl = document.getElementById('root')
if (!rootEl) throw new Error('No #root element in index.html')

createRoot(rootEl).render(
  <StrictMode>
    <App />
  </StrictMode>
)
