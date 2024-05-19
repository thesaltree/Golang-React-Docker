import React from 'react';
import './App.css';
import {BrowserRouter as Router, Switch} from 'react-router-dom'
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListUserComponent from "./components/ListUserComponent";

// Task 12: Write code for App function here(Update routes in the switch statement)
function App() {
      return (
        <div>
            <Router>
                  <HeaderComponent />
                    <Switch>
                        <ListUserComponent />
                    </Switch>
                  <FooterComponent />
            </Router>
        </div>
        
      );
    }
    
    export default App;