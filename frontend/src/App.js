
import React from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Login from './components/Login';
import Register from './components/Register';
import UserList from './components/UserList';
import TodoList from './components/TodoList';

const PrivateRoute = ({ children }) => {
    const token = localStorage.getItem('token');
    return token ? children : <Navigate to="/login" />;
};

function App() {
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