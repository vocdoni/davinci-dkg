import { RouterProvider } from 'react-router-dom'
import { router } from './routes/root'

export function Router() {
  return <RouterProvider router={router} />
}
