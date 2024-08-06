import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';

const Login = () => {
  const [username, setUsername] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await fetch('http://localhost:8981/task_app/user_manager_service/api/v0.1/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ user_name: username, email, password }),
      });

      if (response.ok) {
        const data = await response.json();
        localStorage.setItem('token', data.token);

        // Assume data.role contains the user's role (e.g., 'admin' or 'user')
        if (data.role === 'admin') {
          navigate('/admin-dashboard'); // Redirect to admin dashboard
        } else if (data.role === 'user') {
          navigate('/user-dashboard'); // Redirect to user dashboard
        } else {
          throw new Error('Unknown role');
        }
      } else {
        throw new Error('Login failed');
      }
    } catch (err) {
      alert('Invalid username, email, or password');
    }
  };

  return (
    <div className="shadow-lg rounded-lg px-8 pt-6 pb-8 mb-4 w-full max-w-sm login-container">
      <h1 className="text-3xl font-bold text-center mb-6 text-gray-100">Login</h1>
      <form onSubmit={handleSubmit}>
        <div className="mb-4">
          <label className="block text-gray-300 text-sm font-semibold mb-2" htmlFor="username">
            Username
          </label>
          <div className="flex items-center bg-gray-800 rounded shadow-md border border-gray-700">
            <span className="px-3 text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path d="M10 11a5 5 0 100-10 5 5 0 000 10zm-8 7a8 8 0 1116 0H2z" />
              </svg>
            </span>
            <input
              className="appearance-none rounded w-full py-2 px-3 text-gray-200 leading-tight focus:outline-none focus:shadow-outline bg-gray-800"
              id="username"
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
        </div>
        <div className="mb-4">
          <label className="block text-gray-300 text-sm font-semibold mb-2" htmlFor="email">
            Email
          </label>
          <div className="flex items-center bg-gray-800 rounded shadow-md border border-gray-700">
            <span className="px-3 text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path d="M2 5a2 2 0 012-2h12a2 2 0 012 2v10a2 2 0 01-2 2H4a2 2 0 01-2-2V5zm2 0v10h12V5H4zm4 4h4a1 1 0 010 2H8a1 1 0 010-2z" />
              </svg>
            </span>
            <input
              className="appearance-none rounded w-full py-2 px-3 text-gray-200 leading-tight focus:outline-none focus:shadow-outline bg-gray-800"
              id="email"
              type="email"
              placeholder="Email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
        </div>
        <div className="mb-6">
          <label className="block text-gray-300 text-sm font-semibold mb-2" htmlFor="password">
            Password
          </label>
          <div className="flex items-center bg-gray-800 rounded shadow-md border border-gray-700">
            <span className="px-3 text-gray-400">
              <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                <path
                  fillRule="evenodd"
                  d="M10 2a6 6 0 00-6 6v3H3a1 1 0 000 2h1v4a2 2 0 002 2h8a2 2 0 002-2v-4h1a1 1 0 100-2h-1V8a6 6 0 00-6-6zM7 8a3 3 0 016 0v3H7V8zm7 7H6v-4h8v4z"
                  clipRule="evenodd"
                />
              </svg>
            </span>
            <input
              className="appearance-none rounded w-full py-2 px-3 text-gray-200 mb-3 leading-tight focus:outline-none focus:shadow-outline bg-gray-800"
              id="password"
              type="password"
              placeholder="**********"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
        </div>
        <div className="flex items-center justify-between">
          <button
            className="bg-teal-600 hover:bg-teal-700 text-gray-100 font-bold py-2 px-4 rounded-full focus:outline-none focus:shadow-outline w-full"
            type="submit"
          >
            Sign In
          </button>
        </div>
        <div className="mt-4 text-center">
          <button className="inline-block align-baseline font-semibold text-sm text-blue-400 hover:text-blue-500">
            Forgot Password?
          </button>
        </div>
      </form>
    </div>
  );
};

export default Login;
