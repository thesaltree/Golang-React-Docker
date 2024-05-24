import React from 'react';
import './App.css';
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom'
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import ListUserComponent from "./components/ListUserComponent";
import CreateUserComponent from "./components/CreateUserComponent";
import ViewUserComponent from "./components/ViewUserComponent";

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
                        <Route path = "/view-user/:id" component =
                            {ViewUserComponent}></Route>
                    </Switch>
                  <FooterComponent />
            </Router>
        </div>
        
      );
    }
    
    export default App;