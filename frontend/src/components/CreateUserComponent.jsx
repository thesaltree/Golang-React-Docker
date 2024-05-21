import React, { useState, useEffect } from 'react'
import {useParams} from "react-router-dom";
import UserService from '../services/UserService';

const CreateUserComponent = () => {
    const {id} = useParams()
    const user = {id: '', firstName: '', lastName: '', middleName: '', gender: '', email: '', contact: '', civilStatus: '', address: '', birthday: ''}
    const [currentUser, setCurrentUser] = useState(user)
    useEffect(() => {
        if(id !== '_add') {
            UserService.getUserById(id).then((res) => {
                setCurrentUser(res)
                console.log('in func',res)
            })

        }
    }, [])

    console.log('userrr', currentUser, id)
    return (

        <div>
            {currentUser ? <h1>Update User Component</h1> : <h1>Create User Component</h1>}
            <div>
                <div>First name</div>
                <input type={"text"} id={'first-name'} value={currentUser.firstName} onChange={(e) => {currentUser.firstName = e.target.value
                setCurrentUser(currentUser)}}/>
            </div>
            <div>
                <div>Middle name</div>
                <input type={"text"} id={'middle-name'} defaultValue={currentUser.middleName}/>
            </div>
            <div>
                <div>Last name</div>
                <input type={"text"} id={'last-name'} defaultValue={currentUser.lastName}/>
            </div>
            <div>
                <div>Last name</div>
                <input type={"text"} id={'last-name'} defaultValue={currentUser.lastName}/>
            </div>
        </div>
    )
}

export default CreateUserComponent