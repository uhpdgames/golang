
import React , {useEffect, useState} from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Register from './components/Register';
import UserList from './components/UserList';
import TodoList from './components/TodoList';

import axios from 'axios';
 import {LOCAL_API} from './common'



const PrivateRoute = ({ children }) => {
    const token = localStorage.getItem('token');
    return token ? children : <Navigate to="/login" />;
};

function App() {
    //test

    const [user, setUser] = useState([]);
    const fetchUsers = async (user) => {
        if(user) return;
        
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                
                return;
            }

            const response = await axios.get(`${LOCAL_API}/users`);
            console.log('Users data:', response.data); //   debug
            setUser(response.data);

        } catch (err) {
            if (err.response?.status === 401) {
                localStorage.removeItem('token');
                localStorage.removeItem('user');
               
            } else {
            
            }
        }
    };

    useEffect(() => {
        console.log('Users data:' ); //   debug

        fetchUsers(user);
    },[user]);

    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route path="/login" element={<Login />} />
                    <Route path="/register" element={<Register />} />
                    <Route 
                        path="/users" 
                        element={
                            <PrivateRoute>
                                <UserList />
                            </PrivateRoute>
                        } 
                    />
                    <Route 
                        path="/todos" 
                        element={
                            <PrivateRoute>
                                <TodoList />
                            </PrivateRoute>
                        } 
                    />
                    <Route path="/" element={<Navigate to="/todos" />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;