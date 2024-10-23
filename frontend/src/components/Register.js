import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import {LOCAL_API} from '../common'

const Register = () => {
    const navigate = useNavigate();
    const [formData, setFormData] = useState({
        name: '',
        email: '',
        password: ''
    });
    const [errors, setErrors] = useState({});

    const handleChange = (e) => {
        const { name, value } = e.target;
        setFormData(prevState => ({
            ...prevState,
            [name]: value
        }));
        // Clear error when user starts typing
        if (errors[name]) {
            setErrors(prevState => ({
                ...prevState,
                [name]: ''
            }));
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post(`${LOCAL_API}/register`, formData);
            if (response.data.status === 'success') {
                navigate('/login');
            }
        } catch (err) {
            if (err.response?.data?.errors) {
                const serverErrors = {};
                err.response.data.errors.forEach(error => {
                    serverErrors[error.field.toLowerCase()] = error.message;
                });
                setErrors(serverErrors);
            } else {
                setErrors({
                    general: err.response?.data?.message || 'An error occurred'
                });
            }
        }
    };

    return (
        <div className="register-container">
           <h1 className="text-2xl font-semibold">Register</h1>  
            {errors.general && (
                <div className="error-message">{errors.general}</div>
            )}
            <div className="divide-y divide-gray-200">
            <form onSubmit={handleSubmit} className='py-8 text-base leading-6 space-y-4 text-gray-700 sm:text-lg sm:leading-7'>
                <div className="form-group relative">
                    <label className="absolute left-0 -top-3.5 text-gray-600 text-sm peer-placeholder-shown:text-base peer-placeholder-shown:text-gray-440 peer-placeholder-shown:top-2 transition-all peer-focus:-top-3.5 peer-focus:text-gray-600 peer-focus:text-sm">Name:</label>
                    <input  autoComplete="off" className="peer placeholder-transparent h-10 w-full border-b-2 border-gray-300 text-gray-900 focus:outline-none focus:borer-rose-600"
                        type="text"
                        name="name"
                        value={formData.name}
                        onChange={handleChange}
                    />
                    {errors.name && (
                        <div className="error-message">{errors.name}</div>
                    )}
                </div>
                <div className="form-group relative">
                    <label className="absolute left-0 -top-3.5 text-gray-600 text-sm peer-placeholder-shown:text-base peer-placeholder-shown:text-gray-440 peer-placeholder-shown:top-2 transition-all peer-focus:-top-3.5 peer-focus:text-gray-600 peer-focus:text-sm">Email:</label>
                    <input  autoComplete="off" className="peer placeholder-transparent h-10 w-full border-b-2 border-gray-300 text-gray-900 focus:outline-none focus:borer-rose-600"
                        type="email"
                        name="email"
                        value={formData.email}
                        onChange={handleChange}
                    />
                    {errors.email && (
                        <div className="error-message">{errors.email}</div>
                    )}
                </div>
                <div className="form-group relative">
                    <label className="absolute left-0 -top-3.5 text-gray-600 text-sm peer-placeholder-shown:text-base peer-placeholder-shown:text-gray-440 peer-placeholder-shown:top-2 transition-all peer-focus:-top-3.5 peer-focus:text-gray-600 peer-focus:text-sm">Password:</label>
                    <input  autoComplete="off" className="peer placeholder-transparent h-10 w-full border-b-2 border-gray-300 text-gray-900 focus:outline-none focus:borer-rose-600"
                        type="password"
                        name="password"
                        value={formData.password}
                        onChange={handleChange}
                    />
                    {errors.password && (
                        <div className="error-message">{errors.password}</div>
                    )}
                </div>
                <div className="relative">
                <button type="submit" className='bg-blue-500 text-white rounded-md px-2 py-1'>Register</button>
                </div>
                
            </form></div>
        </div>
        
    );
};

export default Register;
