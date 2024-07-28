import { Layout } from './components/Layout'
import { Home } from './components/Home'
import { Login, LoginCallBack } from './components/Login'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'

function App() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/auth/google/callback" element={<LoginCallBack />} />
        </Routes>
      </Layout>
    </Router>
  )
}

export default App
