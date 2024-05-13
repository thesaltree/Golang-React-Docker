import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';

// Task 12: Write code for App function here(Update routes in the switch statement)
function App() {
      return (
        <div>
            <Router>
                  <HeaderComponent />
                    <Switch> 

                    </Switch>
                  <FooterComponent />
            </Router>
        </div>
        
      );
    }
    
    export default App;