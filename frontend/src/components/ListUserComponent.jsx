import React, { useState, useEffect } from 'react'
import UserService from '../services/UserService'

const ListUserComponent = () => {
    const [users, setUsers] = useState([])

    useEffect(() => {
        // console.log("Inside useEffect")
        UserService.getUsers().then((res) => {
            // console.log(res)
            setUsers(res)
        })
    }, [])

    const addUser = () => {
        UserService.createUser().then((res) => {
            setUsers([...users, res.data])
        })
    }

    // const updateUser = () => {
    //     UserService.updateUser(id).then((res) => {
    //         setUsers(users.map((user) => (user.id === res.data.id)? res.data : user))
    //     })
    // }

    return (
        <div>
            <p>Users List</p>
            <button className="btn btn-primary" onClick={() => {}}>
                Add User
            </button>
            <table className="table">
                <thead>
                    <tr>
                        <th>Id</th>
                        <th>Name</th>
                        <th>Email</th>
                        <th>Update user </th>
                        <th>Delete user</th>
                        <th>View user</th>
                    </tr>
                </thead>
                <tbody>
                    {
                        users.map((user) =>
                            <tr key={user.id}>
                                <td>{user.id}</td>
                                <td>{user.name}</td>
                                <td>{user.email}</td>
                                <td><button className="btn btn-info" onClick={() => addUser()}>Update</button></td>
                                <td><button className="btn btn-danger">Delete</button></td>
                                <td><button className="btn btn-info">View</button></td>
                            </tr>
                        )
                    }
                </tbody>
            </table>
        </div>
    )
}

export default ListUserComponent