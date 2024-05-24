import React, { useState, useEffect } from 'react'
import {useHistory} from "react-router-dom";
import {useParams} from "react-router-dom";
import UserService from '../services/UserService';


const ViewUserComponent = () => {
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
            User Details
            <div>
                <div>First name</div>
                {currentUser.firstName}
            </div>
            <div>
                <div>Middle name</div>
                {currentUser.middleName}
            </div>
            <div>
                <div>Last name</div>
                {currentUser.lastName}
            </div>
            <div>
                <div>Gender</div>
                {currentUser.gender}
            </div>
            <div>
                <div>Email</div>
                {currentUser.email}
            </div>
            <div>
                <div>Contact</div>
                {currentUser.contact}
            </div>
            <div>
                <div>Civil Status</div>
                {currentUser.civilStatus}
            </div>
            <div>
                <div>Address</div>
                {currentUser.address}
            </div>
            <div>
                <div>Birthday</div>
                {currentUser.birthday}
            </div>
            <button onClick={() => {
                history.push('/users')
            }}>Back
            </button>
        </div>
    )
}

export default ViewUserComponent