import React from 'react';
import './App.css';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom'
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListUserComponent from "./components/ListUserComponent";
import CreateUserComponent from "./components/CreateUserComponent";

// Task 12: Write code for App function here(Update routes in the switch statement)
function App() {
      return (
        <div>
            <Router>
                  <HeaderComponent />
                    <Switch>
                        <Route path = "/" exact component =
                            {ListUserComponent}></Route>
                        <Route path = "/users" component =
                            {ListUserComponent}></Route>
                        <Route path = "/add-user/:id" component =
                            {CreateUserComponent}></Route>
                    </Switch>
                  <FooterComponent />
            </Router>
        </div>
        
      );
    }
    
    export default App;