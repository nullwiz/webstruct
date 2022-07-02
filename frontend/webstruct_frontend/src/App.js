import Navbar from './components/navbar';
import {useNavigate} from 'react-router-dom';

function App() {
  const navigate = useNavigate();
  
  const arrayPage = () =>{
    navigate('/arrays')
  }
  return (
    <div class="overflow-hidden">
      <Navbar/>
    <div class="hero min-h-screen bg-base-200">
      <div class="hero-content text-center">
        <div class="max-w-md">
          <h1 class="text-5xl font-bold">Hello there!</h1>
          <p class="py-6">Pick from any of the structures</p>
          <button class="btn btn-primary" onClick={arrayPage}> Get Started</button>
        </div>
      </div>
    </div>
    </div>
    );
}

export default App;
