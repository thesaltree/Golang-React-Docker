import React, { useState, useEffect } from 'react'
import { useHistory } from 'react-router-dom';
import UserService from '../services/UserService'

const ListUserComponent = () => {
    const [users, setUsers] = useState([])
    const history = useHistory()

    useEffect(() => {
        UserService.getUsers().then((res) => {
            setUsers(res)
        })
    }, [])

    const addUser = () => {
        console.log('add user')
        history.push('/add-user')
    }

    const updateUser = (id) => {
        history.push(`/add-user/${id}`)
    }

    const deleteUser = (id) => {
        UserService.deleteUser(id).then((res) => {
            setUsers(users.filter((user) => user.id!== id))
        })
    }

    const viewUser = (id) => {
        history.push(`/view-user/${id}`)
    }

    return (
        <div>
            <p>Users List</p>
            <button className="btn btn-primary" onClick={() => {addUser()}}>
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
                                <td><button className="btn btn-info" onClick={() => updateUser(user.id)}>Update</button></td>
                                <td><button className="btn btn-danger" onClick={() => deleteUser(user.id)}>Delete</button></td>
                                <td><button className="btn btn-info" onClick={() => viewUser(user.id)}>View</button></td>
                            </tr>
                        )
                    }
                </tbody>
            </table>
        </div>
    )
}

export default ListUserComponent