import React, { useState } from 'react';
import { FaBars } from 'react-icons/fa';
import { useNavigate } from 'react-router-dom';

const Sidebar = () => {
  const [isOpen, setIsOpen] = useState(false);
  const navigate = useNavigate();

  const toggleSidebar = () => setIsOpen(!isOpen);
  const goToDashboard = () => navigate('/admin-dashboard');
  const goToUsers = () => navigate('/users');

  return (
    <div className="relative">
      <button className="p-2 text-white" onClick={toggleSidebar}>
        <FaBars />
      </button>
      {isOpen && (
        <div className="absolute top-0 left-0 w-48 bg-gray-800 text-white p-4 rounded shadow-lg">
          <button onClick={goToDashboard} className="block mb-2">Dashboard</button>
          <button onClick={goToUsers} className="block">Users</button>
        </div>
      )}
    </div>
  );
};

export default Sidebar;
