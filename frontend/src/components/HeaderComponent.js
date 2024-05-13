import React, { Component } from 'react'
import 'bootstrap/dist/css/bootstrap.min.css';
class HeaderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {

        }
    }

    render() {
        return (
            <div>
                <header>
                    <nav 
                    className="header  navbar-dark bg-dark">
                        <div>
                            <a href="/users"
                                className="navbar-brand">
                                    User CRUD using Golang
                            </a></div>
                    </nav>
                </header>
            </div>
        )
    }
}

export default HeaderComponent