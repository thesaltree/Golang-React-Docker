import React, { useState, useEffect } from 'react'
import {useParams} from "react-router-dom";
import UserService from '../services/UserService';

const CreateUserComponent = () => {
    const {id} = useParams()
    const [user, setUser] = useState(undefined)
    useEffect(() => {
        if(id) {
            UserService.getUserById(id).then((res) => {
                setUser(res)
                console.log(res)
            })

        }
    }, [id])
    return (
        <div>
            <h1>Create User Component</h1>
        </div>
    )
}

export default CreateUserComponent