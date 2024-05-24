import React, { useState, useEffect } from 'react'
import {useHistory} from "react-router-dom";
import {useParams} from "react-router-dom";
import UserService from '../services/UserService';

const CreateUserComponent = () => {
    const {id} = useParams()
    const user = {firstName: '', lastName: '', middleName: '', gender: '', email: '', contact: '', civilStatus: '', address: '', birthday: ''}
    const [currentUser, setCurrentUser] = useState(user)
    const history = useHistory()
    useEffect(() => {
        if(id !== '_add') {
            UserService.getUserById(id).then((res) => {
                setCurrentUser(res)
            })

        }
    }, [id])
    return (

        <div>
            {currentUser ? <h1>Update User Component</h1> : <h1>Create User Component</h1>}
            <div>
                <div>First name</div>
                <input type={"text"} id={'first-name'} defaultValue={currentUser.firstName}
                       onChange={(e) => {
                           currentUser.firstName = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Middle name</div>
                <input type={"text"} id={'middle-name'} defaultValue={currentUser.middleName}
                       onChange={(e) => {
                           currentUser.middleName = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Last name</div>
                <input type={"text"} id={'last-name'} defaultValue={currentUser.lastName}
                       onChange={(e) => {
                           currentUser.lastName = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Gender</div>
                <input type={"text"} id={'gender'} defaultValue={currentUser.gender}
                       onChange={(e) => {
                           currentUser.gender = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Email</div>
                <input type={"text"} id={'email'} defaultValue={currentUser.email}
                       onChange={(e) => {
                           currentUser.email = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Contact</div>
                <input type={"text"} id={'contact'} defaultValue={currentUser.contact}
                       onChange={(e) => {
                           currentUser.contact = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Civil Status</div>
                <input type={"text"} id={'civil-status'} defaultValue={currentUser.civilStatus}
                       onChange={(e) => {
                           currentUser.civilStatus = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Address</div>
                <input type={"text"} id={'address'} defaultValue={currentUser.address}
                       onChange={(e) => {
                           currentUser.address = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <div>
                <div>Birthday</div>
                <input type={"text"} id={'birthday'} defaultValue={currentUser.birthday}
                       onChange={(e) => {
                           currentUser.birthday = e.target.value
                           setCurrentUser(currentUser)
                       }}/>
            </div>
            <button onClick={() => {
                if (id === '_add') {
                    UserService.createUser(currentUser).then((res) => {
                        console.log(res)
                    })
                } else {
                    UserService.updateUser(currentUser, id).then((res) => {
                        console.log(res)
                    })
                }
                history.push('/users')
            }}>Submit
            </button>
            <button onClick={() => {
                history.push('/users')
            }}>Cancel
            </button>
        </div>
    )
}

export default CreateUserComponent