import React, { useState } from 'react';
import { FaBars, FaUserCircle } from 'react-icons/fa';
import { Link } from 'react-router-dom';
import { Bar } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale
} from 'chart.js';

// Register Chart.js components
ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale);

const AdminDashboard = () => {
  const [sidebarOpen, setSidebarOpen] = useState(true);
  const [modalOpen, setModalOpen] = useState(false);
  const [newTask, setNewTask] = useState({
    title: '',
    description: '',
    deadline: '',
    userId: ''
  });
  const [tasks, setTasks] = useState([
    {
      task_name: 'Task 1',
      task_description: 'Description for Task 1',
      task_status: 'Pending',
      task_schedule: '2024-07-01',
      assigned_user_name: 'User A',
      assigned_task_email: 'usera@example.com'
    },
    {
      task_name: 'Task 2',
      task_description: 'Description for Task 2',
      task_status: 'In-progress',
      task_schedule: '2024-07-05',
      assigned_user_name: 'User B',
      assigned_task_email: 'userb@example.com'
    },
    {
      task_name: 'Task 3',
      task_description: 'Description for Task 3',
      task_status: 'Completed',
      task_schedule: '2024-06-30',
      assigned_user_name: 'User C',
      assigned_task_email: 'userc@example.com'
    }
  ]);
  const [users] = useState([
    { id: '1', name: 'User A' },
    { id: '2', name: 'User B' },
    { id: '3', name: 'User C' }
  ]);

  const handleCreateTask = () => {
    setTasks([...tasks, { ...newTask, task_status: 'Pending', task_name: newTask.title, assigned_user_name: users.find(user => user.id === newTask.userId)?.name || 'Unassigned', assigned_task_email: 'email@example.com' }]);
    setModalOpen(false);
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setNewTask({ ...newTask, [name]: value });
  };

  const handleDeleteTask = (taskName) => {
    if (window.confirm('Are you sure you want to delete this task?')) {
      setTasks(tasks.filter(task => task.task_name !== taskName));
    }
  };

  const handleDeleteAllTasks = () => {
    if (window.confirm('Are you sure you want to delete all tasks?')) {
      setTasks([]);
    }
  };

  const handleUpdateSchedule = (taskName, newSchedule) => {
    const updatedTasks = tasks.map(task =>
      task.task_name === taskName ? { ...task, task_schedule: newSchedule } : task
    );
    setTasks(updatedTasks);
  };

  // Sample data for the bar chart
  const data = {
    labels: ['Pending', 'In-progress', 'Completed', 'Total Tasks'],
    datasets: [
      {
        label: 'Number of Tasks',
        data: [tasks.filter(task => task.task_status === 'Pending').length,
               tasks.filter(task => task.task_status === 'In-progress').length,
               tasks.filter(task => task.task_status === 'Completed').length,
               tasks.length],
        backgroundColor: [
          'rgba(75, 192, 192, 0.2)',
          'rgba(153, 102, 255, 0.2)',
          'rgba(255, 159, 64, 0.2)',
          'rgba(255, 99, 132, 0.2)'
        ],
        borderColor: [
          'rgba(75, 192, 192, 1)',
          'rgba(153, 102, 255, 1)',
          'rgba(255, 159, 64, 1)',
          'rgba(255, 99, 132, 1)'
        ],
        borderWidth: 1
      }
    ]
  };

  // Sample data for the circles
  const circleData = [
    { label: 'Pending', value: tasks.filter(task => task.task_status === 'Pending').length, color: '#ff6384' },
    { label: 'In-progress', value: tasks.filter(task => task.task_status === 'In-progress').length, color: '#36a2eb' },
    { label: 'Completed', value: tasks.filter(task => task.task_status === 'Completed').length, color: '#ffce56' },
    { label: 'Total Tasks', value: tasks.length, color: '#4bc0c0' }
  ];

  return (
    <div className="flex">
      {/* Sidebar */}
      <div
        className={`fixed top-0 left-0 h-full bg-gray-800 text-gray-100 w-64 transform ${sidebarOpen ? 'translate-x-0' : '-translate-x-full'} transition-transform duration-300 ease-in-out z-40`}
        style={{ maxWidth: '16rem' }}
      >
        <div className="p-4">
          <h2 className="text-xl font-bold mb-4">Menu</h2>
          <Link to="/users" className="block py-2 px-4 flex items-center space-x-2 hover:bg-gray-700">
            <span>Users</span>
          </Link>
        </div>
        {/* Close Button */}
        <div className="absolute top-4 right-4">
          <button
            className="p-2 text-gray-100"
            onClick={() => setSidebarOpen(false)}
          >
            <FaBars size={24} />
          </button>
        </div>
      </div>

      {/* Main Content */}
      <div className={`flex-1 p-6 space-y-8 relative ${sidebarOpen ? 'ml-64' : ''}`}>
        {/* Sidebar Button */}
        <div className={`absolute top-4 left-4 z-50 ${sidebarOpen ? 'hidden' : 'block'}`}>
          <button
            className="p-2 text-gray-100"
            onClick={() => setSidebarOpen(true)}
          >
            <FaBars size={24} />
          </button>
        </div>

        {/* Admin Profile */}
        <div className="flex items-center justify-end mb-4">
          <FaUserCircle size={40} className="text-gray-400 mr-2" />
          <span className="text-gray-300">Admin</span>
        </div>

        {/* Analytics Circles */}
        <div className="flex gap-6 mb-8 justify-center flex-wrap">
          {circleData.map((data, index) => (
            <div key={index} className="flex flex-col items-center space-y-2">
              <div className="w-20 h-20 flex items-center justify-center rounded-full" style={{ backgroundColor: data.color }}>
                <span className="text-white text-xl font-bold">{data.value}</span>
              </div>
              <p className="text-white font-semibold">{data.label}</p>
            </div>
          ))}
        </div>

        {/* Bar Chart */}
        <div className="w-full max-w-4xl mb-8">
          <Bar data={data} options={{ responsive: true, maintainAspectRatio: false }} />
        </div>

        {/* Tasks Table */}
        <div className="overflow-x-auto mb-20">
          <table className="min-w-full bg-gray-800 text-gray-100 border border-gray-600 text-sm">
            <thead>
              <tr className="w-full border-b border-gray-600">
                <th className="px-4 py-2 text-left">Task Title</th>
                <th className="px-4 py-2 text-left">Description</th>
                <th className="px-4 py-2 text-left">Status</th>
                <th className="px-4 py-2 text-left">Deadline</th>
                <th className="px-4 py-2 text-left">Assigned User</th>
                <th className="px-4 py-2 text-left">Assigned Email</th>
                <th className="px-4 py-2 text-left">Actions</th>
                <th className="px-4 py-2 text-left">
                  <button
                    className="bg-red-600 text-white px-4 py-2 rounded"
                    onClick={handleDeleteAllTasks}
                  >
                    Delete All
                  </button>
                </th>
              </tr>
            </thead>
            <tbody>
              {tasks.map((task, index) => (
                <tr key={index} className="border-b border-gray-600">
                  <td className="px-4 py-2">{task.task_name}</td>
                  <td className="px-4 py-2">{task.task_description}</td>
                  <td className="px-4 py-2">{task.task_status}</td>
                  <td className="px-4 py-2">
                    <input
                      type="date"
                      value={task.task_schedule}
                      onChange={(e) => handleUpdateSchedule(task.task_name, e.target.value)}
                      className="bg-gray-700 text-white rounded px-2 py-1"
                    />
                  </td>
                  <td className="px-4 py-2">{task.assigned_user_name}</td>
                  <td className="px-4 py-2">{task.assigned_task_email}</td>
                  <td className="px-4 py-2">
                    <button
                      className="bg-red-600 text-white px-4 py-2 rounded"
                      onClick={() => handleDeleteTask(task.task_name)}
                    >
                      Delete
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {/* Add Task Button */}
        <div className="fixed bottom-4 right-4">
          <button
            className="p-4 bg-blue-600 text-white rounded-full shadow-lg"
            onClick={() => setModalOpen(true)}
          >
            +
          </button>
        </div>

        {/* Create Task Modal */}
        {modalOpen && (
          <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
            <div className="bg-gray-800 p-6 rounded w-full max-w-sm space-y-4">
              <h2 className="text-white text-xl font-bold mb-4">Create New Task</h2>
              <input
                type="text"
                name="title"
                value={newTask.title}
                onChange={handleInputChange}
                placeholder="Title"
                className="w-full p-2 bg-gray-700 text-white rounded"
              />
              <input
                type="text"
                name="description"
                value={newTask.description}
                onChange={handleInputChange}
                placeholder="Description"
                className="w-full p-2 bg-gray-700 text-white rounded"
              />
              <input
                type="date"
                name="deadline"
                value={newTask.deadline}
                onChange={handleInputChange}
                className="w-full p-2 bg-gray-700 text-white rounded"
              />
              <select
                name="userId"
                value={newTask.userId}
                onChange={handleInputChange}
                className="w-full p-2 bg-gray-700 text-white rounded"
              >
                <option value="">Select User</option>
                {users.map(user => (
                  <option key={user.id} value={user.id}>{user.name}</option>
                ))}
              </select>
              <div className="flex justify-between space-x-2">
                <button
                  className="w-1/2 bg-blue-600 text-white p-2 rounded"
                  onClick={handleCreateTask}
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
      </div>
    </div>
  );
};

export default AdminDashboard;
