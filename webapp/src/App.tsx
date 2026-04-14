import { Route, Routes } from 'react-router-dom';
import { Layout } from './components/Layout';
import { Home } from './pages/Home';
import { Playground } from './pages/Playground';
import { Registry } from './pages/Registry';
import { RoundDetail } from './pages/RoundDetail';
import { Rounds } from './pages/Rounds';
import { Settings } from './pages/Settings';

export default function App() {
  return (
    <Layout>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/rounds" element={<Rounds />} />
        <Route path="/rounds/:id" element={<RoundDetail />} />
        <Route path="/registry" element={<Registry />} />
        <Route path="/playground" element={<Playground />} />
        <Route path="/settings" element={<Settings />} />
      </Routes>
    </Layout>
  );
}
