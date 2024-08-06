// src/App.js

import React from 'react';
import { Route, Routes, Navigate } from 'react-router-dom';
import Login from './components/Login';
import AdminDashboard from './components/AdminDashboard';
import UserDashboard from './components/UserDashboard';
import Users from './components/Users';

const App = () => {
  return (
    <Routes>
      {/* Redirect root path to /login */}
      <Route path="/" element={<Navigate to="/login" />} />
      
      {/* Route for Login */}
      <Route path="/login" element={<Login />} />
      
      {/* Route for Admin Dashboard */}
      <Route path="/admin-dashboard" element={<AdminDashboard />} />
      {/* Route for user dashboard */}
      <Route path="/user-dashboard" element={<UserDashboard />} />
      {/* Route for Users */}
      <Route path="/users" element={<Users />} />
      
      {/* Catch-all redirect to /login */}
      <Route path="*" element={<Navigate to="/login" />} />
    </Routes>
  );
};

export default App;
