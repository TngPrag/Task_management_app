import React, { useState } from 'react';
import { FaUserCircle, FaArrowLeft } from 'react-icons/fa';
import { Link } from 'react-router-dom';

const UsersPage = () => {
  const [users, setUsers] = useState([
    { id: 1, fullName: 'User A', userName: 'usera', password: 'password123', email: 'usera@example.com' },
    { id: 2, fullName: 'User B', userName: 'userb', password: 'password456', email: 'userb@example.com' },
    { id: 3, fullName: 'User C', userName: 'userc', password: 'password789', email: 'userc@example.com' },
  ]);
  const [newUser, setNewUser] = useState({ fullName: '', userName: '', password: '', email: '' });
  const [modalOpen, setModalOpen] = useState(false);

  const handleDeleteUser = (userId) => {
    if (window.confirm('Are you sure you want to delete this user?')) {
      setUsers(users.filter(user => user.id !== userId));
    }
  };

  const handleDeleteAllUsers = () => {
    if (window.confirm('Are you sure you want to delete all users?')) {
      setUsers([]);
    }
  };

  const handleCreateUser = () => {
    if (newUser.fullName && newUser.userName && newUser.password && newUser.email) {
      setUsers([...users, { id: Date.now(), ...newUser }]);
      setNewUser({ fullName: '', userName: '', password: '', email: '' });
      setModalOpen(false);
    } else {
      alert('Please fill in all fields');
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewUser({ ...newUser, [name]: value });
  };

  return (
    <div className="flex flex-col h-screen">
      {/* Header */}
      <header className="bg-gray-800 text-gray-100 p-4 flex justify-between items-center">
        <Link to="/admin-dashboard">
          <button className="text-gray-100 flex items-center">
            <FaArrowLeft size={24} />
            <span className="ml-2">Back</span>
          </button>
        </Link>
        {/* Admin Profile */}
        <div className="flex items-center">
          <FaUserCircle size={40} className="text-gray-400 mr-2" />
          <span className="text-gray-300">Admin</span>
        </div>
      </header>

      {/* Main Content */}
      <main className="flex-1 p-6">
        <div className="flex justify-between mb-4">
          {/* Delete All Button */}
          <button
            className="bg-red-600 text-white px-4 py-2 rounded"
            onClick={handleDeleteAllUsers}
          >
            Delete All
          </button>
        </div>

        {/* Users Table */}
        <div className="overflow-x-auto mb-20">
          <table className="min-w-full bg-gray-800 text-gray-100 border border-gray-600 text-sm">
            <thead>
              <tr className="w-full border-b border-gray-600">
                <th className="px-4 py-2 text-left">Full Name</th>
                <th className="px-4 py-2 text-left">User Name</th>
                <th className="px-4 py-2 text-left">Email</th>
                <th className="px-4 py-2 text-left">Password</th>
                <th className="px-4 py-2 text-left">Actions</th>
              </tr>
            </thead>
            <tbody>
              {users.map(user => (
                <tr key={user.id} className="border-b border-gray-600">
                  <td className="px-4 py-2">{user.fullName}</td>
                  <td className="px-4 py-2">{user.userName}</td>
                  <td className="px-4 py-2">{user.email}</td>
                  <td className="px-4 py-2">{user.password}</td>
                  <td className="px-4 py-2">
                    <button
                      className="bg-red-600 text-white px-4 py-2 rounded"
                      onClick={() => handleDeleteUser(user.id)}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {/* Add User Button */}
        <div className="fixed bottom-4 right-4">
          <button
            className="p-4 bg-blue-600 text-white rounded-full shadow-lg"
            onClick={() => setModalOpen(true)}
          >
            +
          </button>
        </div>

        {/* Create User Modal */}
        {modalOpen && (
          <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
            <div className="bg-gray-800 p-6 rounded w-full max-w-sm space-y-4">
              <h2 className="text-white text-xl font-bold mb-4">Create New User</h2>
              <input
                type="text"
                name="fullName"
                value={newUser.fullName}
                onChange={handleInputChange}
                placeholder="Full Name"
                className="w-full p-2 bg-gray-700 text-white rounded"
              />
              <input
                type="text"
                name="userName"
                value={newUser.userName}
                onChange={handleInputChange}
                placeholder="User Name"
                className="w-full p-2 bg-gray-700 text-white rounded"
              />
              <input
                type="password"
                name="password"
                value={newUser.password}
                onChange={handleInputChange}
                placeholder="Password"
                className="w-full p-2 bg-gray-700 text-white rounded"
              />
              <input
                type="email"
                name="email"
                value={newUser.email}
                onChange={handleInputChange}
                placeholder="Email"
                className="w-full p-2 bg-gray-700 text-white rounded"
              />
              <div className="flex justify-between space-x-2">
                <button
                  className="w-1/2 bg-blue-600 text-white p-2 rounded"
                  onClick={handleCreateUser}
                >
                  Create
                </button>
                <button
                  className="w-1/2 bg-red-600 text-white p-2 rounded"
                  onClick={() => setModalOpen(false)}
                >
                  Cancel
                </button>
              </div>
            </div>
          </div>
        )}
      </main>
    </div>
  );
};

export default UsersPage;
