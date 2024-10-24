import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import {LOCAL_API} from '../common'

const TodoList = () => {
    const [todos, setTodos] = useState([]);
    const [newTodo, setNewTodo] = useState({ title: '', description: '' });
    const [error, setError] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        fetchTodos();
        console.log(todos)
    }, [todos]);

    const fetchTodos = async () => {
        try {
            const token = localStorage.getItem('token');
            if (!token) {
                navigate('/login');
                return;
            }

            const response = await axios.get(`${LOCAL_API}/todos`, { 
                withCredentials: true,
                headers: { 'Authorization': `Bearer ${token}`,
                 'Access-Control-Allow-Origin': '*',
                  'Content-Type': 'application/json'
                }
            });
            setTodos(response.data);
        } catch (err) {
            if (err.response?.status === 401) {
                localStorage.removeItem('token');
                navigate('/login');
            } else {
                setError(err.response?.data?.error || 'An error occurred');
            }
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const token = localStorage.getItem('token');
            const response = await axios.post(`${LOCAL_API}/todos`, newTodo, {
                headers: { 
                    'Authorization': `Bearer ${token}`, 
                    'Access-Control-Allow-Origin': '*',
                    'Content-Type': 'application/json'
                 },
                 withCredentials: true
            });
            setTodos([...todos, response.data]);
            setNewTodo({ title: '', description: '' });
        } catch (err) {
            setError(err.response?.data?.error || 'An error occurred');
        }
    };

    const handleToggleStatus = async (todoId, currentStatus) => {
        try {
            const token = localStorage.getItem('token');
            await axios.put(`${LOCAL_API}/todos/${todoId}`, 
                { status: !currentStatus },
                { headers: { 
                    'Authorization': `Bearer ${token}`, 
                    'Access-Control-Allow-Origin': '*',
                    'Content-Type': 'application/json'
                 },
                 withCredentials: true,
                },
            );
            fetchTodos();
        } catch (err) {
            setError(err.response?.data?.error || 'An error occurred');
        }
    };

    const handleDelete = async (todoId) => {
        try {
            const token = localStorage.getItem('token');
            await axios.delete(`${LOCAL_API}/todos/${todoId}`, {
                headers: { 
                    'Authorization': `Bearer ${token}`,
                    'Access-Control-Allow-Origin': '*',
                    'Content-Type': 'application/json' 
                },
                withCredentials: true,
            });
            setTodos(todos.filter(todo => todo.id !== todoId));
        } catch (err) {
            setError(err.response?.data?.error || 'An error occurred');
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('token');
        localStorage.removeItem('user');
        navigate('/login');
    };

    return (
        <div className="todo-container">
            <div className="header">
                <h2 className='text-2xl font-semibold'>Todo List</h2>
                <button onClick={handleLogout} className="logout-btn">Logout</button>
            </div>

            {error && <div className="error-message">{error}</div>}

            <form onSubmit={handleSubmit} className="todo-form">
                <input
                    type="text"
                    value={newTodo.title}
                    onChange={(e) => setNewTodo({...newTodo, title: e.target.value})}
                    placeholder="Todo title"
                    required
                />
                <input
                    type="text"
                    value={newTodo.description}
                    onChange={(e) => setNewTodo({...newTodo, description: e.target.value})}
                    placeholder="Description (optional)"
                />
                <button type="submit">Add Todo</button>
            </form>

            <div className="todo-list">
                {todos.map(todo => (
                    <div key={todo.id} className={`todo-item ${todo.status ? 'completed' : ''}`}>
                        <div className="todo-content">
                            <h3>{todo.title}</h3>
                            <p>{todo.description}</p>
                        </div>
                        <div className="todo-actions">
                            <input
                                type="checkbox"
                                checked={todo.status || false}
                                onChange={() => handleToggleStatus(todo.id, todo.status)}
                            />
 

                            <div className="relative">
							<button onClick={() => handleDelete(todo.id)} className="bg-blue-500 text-white rounded-md px-2 py-1">Delete</button>
						</div>

                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default TodoList;