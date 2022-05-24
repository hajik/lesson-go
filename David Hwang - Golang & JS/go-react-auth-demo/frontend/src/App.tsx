import { useEffect, useState } from 'react';
import './App.css';
import Header from './components/Header';
import Footer from './components/Footer';
import {Container} from 'react-bootstrap';
import {BrowserRouter as Router, Route} from 'react-router-dom';
import HomeScreen from './screens/HomeScreen';
import SignupScreen from './screens/SignupScreen';
import LoginScreen from './screens/LoginScreen';

function App() {

  const [firstName, setFirstName] = useState('')


  useEffect(() => {
    (
      async () => {
        const response = await fetch('https:localhost:8081/api/user', {
          headers: {'Content-Type': 'application/json'},
          credentials: 'include',
        })
       
        const data = await response.json()
        setFirstName(data.firstname)
        
      }
    )()
  })

  return (
    <Router>
      <Header firstName={firstName} setFirstName={setFirstName} />
        <main>
          <Container>
            <Route 
            path='/'>
              <HomeScreen firstName={firstName} />
            </Route>
            {/* <Route path='/signup' component={SignupScreen} />
            <Route path='/siginin' component={LoginScreen} /> */}
          </Container>
        </main>
      <Footer />
    </Router>
  );
}

export default App;
