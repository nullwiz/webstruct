import Navbar from './components/navbar'
import { useNavigate } from 'react-router-dom'

function App() {
  const navigate = useNavigate()

  const arrayPage = () => {
    navigate('/arrays')
  }
  return (
    <div className="overflow-hidden">
      <Navbar />
      <div className="hero min-h-screen">
        <div className="hero-content text-center">
          <div className="max-w-md">
            <h1 className="text-5xl font-bold">Hello there!</h1>
            <p className="py-6">Pick from any of the structures</p>
            <button className="btn btn-primary" onClick={arrayPage}>
              {' '}
              Get Started
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default App
