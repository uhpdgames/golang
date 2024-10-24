import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import {LOCAL_API} from '../common'
import {Link} from 'react-router-dom';

const Login = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        email: '',
        password: ''
    });
    const [error, setError] = useState('');

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.name]: e.target.value
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post(`${LOCAL_API}/login`, formData, {
                headers: { 
                    'Access-Control-Allow-Origin': '*',
                    'Content-Type': 'application/json'
                 },
                 withCredentials: true
            });
            
            // token
            localStorage.setItem('token', response.data.token);
            localStorage.setItem('user', JSON.stringify(response.data.user));
            
            // to list users 
            navigate('/users');
        } catch (err) {
            setError(err.response?.data?.error || 'An error occurred');
        }
    };

    const allList = ['/login','/register','/users','/todos']
    return (
        <div className="login-container">

            <div>
            <ul className='flex'>
            {allList.map((l) => (
                 
                <li className='py-2 my-4' key={Math.random()}>
                    <Link   to={l}>{l}</Link>
                </li>
               
                ))}
        </ul>
            <Link to='login' />
        
        
        
            </div>
            <h2 className='text-2xl font-semibold'>Login</h2>
            {error && <div className="error-message">{error}</div>}
            <div className="divide-y divide-gray-200">
            <form onSubmit={handleSubmit} className='py-8 text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7'>
                <div className="relative form-group">
                    <label className="absolute left-0 -top-3.5 text-gray-600 text-sm peer-placeholder-shown:text-base peer-placeholder-shown:text-gray-440 peer-placeholder-shown:top-2 transition-all peer-focus:-top-3.5 peer-focus:text-gray-600 peer-focus:text-sm">Email:</label>
                    <input  autoComplete="off" className="peer placeholder-transparent h-10 w-full border-b-2 border-gray-300 text-gray-900 focus:outline-none focus:borer-rose-600"
                        type="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="relative form-group">
                    <label className="absolute left-0 -top-3.5 text-gray-600 text-sm peer-placeholder-shown:text-base peer-placeholder-shown:text-gray-440 peer-placeholder-shown:top-2 transition-all peer-focus:-top-3.5 peer-focus:text-gray-600 peer-focus:text-sm">Password:</label>
                    <input  autoComplete="off" className="peer placeholder-transparent h-10 w-full border-b-2 border-gray-300 text-gray-900 focus:outline-none focus:borer-rose-600"
                        type="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                        required
                    />
                </div>
                <div className="relative">
							<button className="bg-blue-500 text-white rounded-md px-2 py-1">Login</button>
						</div>
                 
            </form>
            </div>
        </div>
    );
};

export default Login;