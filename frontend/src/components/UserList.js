import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import {LOCAL_API} from '../common'
const UserList = () => {
    const [users, setUsers] = useState([]);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const token = localStorage.getItem('token');
                if (!token) {
                    navigate('/login');
                    return;
                }

                const response = await axios.get(`${LOCAL_API}/users`, {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                });
                console.log('Users data:', response.data); //   debug
            
                setUsers(response.data);
            } catch (err) {
                if (err.response?.status === 401) {
                    localStorage.removeItem('token');
                    localStorage.removeItem('user');
                    navigate('/login');
                } else {
                    setError(err.response?.data?.error || 'An error occurred');
                }
            }
        };

        fetchUsers();
    }, [navigate]);
    const formatDate = (dateString) => {
        if (!dateString) return 'N/A';
        const date = new Date(dateString);
        if (isNaN(date.getTime())) return 'Invalid Date';
        return date.toLocaleDateString('en-US', {
            year: 'numeric',
            month: 'long',
            day: 'numeric',
            hour: '2-digit',
            minute: '2-digit'
        });
    };
    return (
        <div className="user-list-container">
            <h2 className='text-2xl font-semibold'>Users</h2>
            {/* {error && <div className="error-message">{error}</div>} */}
            
            <div className='divide-y divide-gray-200'>
            <table className="user-table">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>Name</th>
                        <th>Email</th>
                        <th>Created At</th>
                    </tr>
                </thead>
                <tbody>
                    {users && users.map(user => (
                        <tr key={user.id}>
                            <td>{user.id || 'N/A'}</td>
                            <td>{user.username || 'N/A'}</td>
                            <td>{user.email || 'N/A'}</td>
                            <td>{formatDate(user.created_at)}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
            </div>
           
        </div>
    );
};

export default UserList;