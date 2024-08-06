import React, { useState } from 'react';
import { FaUserCircle } from 'react-icons/fa';
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

const UserDashboard = () => {
  const [tasks] = useState([
    {
      task_name: 'Task 1',
      task_description: 'Description for Task 1',
      task_status: 'Pending',
      task_schedule: '2024-07-01',
      assigned_user_name: 'User A'
    },
    {
      task_name: 'Task 2',
      task_description: 'Description for Task 2',
      task_status: 'Scheduled',
      task_schedule: '2024-07-05',
      assigned_user_name: 'User A'
    },
    {
      task_name: 'Task 3',
      task_description: 'Description for Task 3',
      task_status: 'Completed',
      task_schedule: '2024-06-30',
      assigned_user_name: 'User A'
    }
  ]);

  // Analytics Data
  const analyticsData = {
    completed: tasks.filter(task => task.task_status === 'Completed').length,
    pending: tasks.filter(task => task.task_status === 'Pending').length,
    scheduled: tasks.filter(task => task.task_status === 'Scheduled').length
  };

  // Sample data for the bar chart
  const data = {
    labels: ['Pending', 'Scheduled', 'Completed'],
    datasets: [
      {
        label: 'Number of Tasks',
        data: [
          analyticsData.pending,
          analyticsData.scheduled,
          analyticsData.completed
        ],
        backgroundColor: [
          'rgba(75, 192, 192, 0.2)',
          'rgba(153, 102, 255, 0.2)',
          'rgba(255, 159, 64, 0.2)'
        ],
        borderColor: [
          'rgba(75, 192, 192, 1)',
          'rgba(153, 102, 255, 1)',
          'rgba(255, 159, 64, 1)'
        ],
        borderWidth: 1
      }
    ]
  };

  // Sample data for the circles
  const circleData = [
    { label: 'Pending', value: analyticsData.pending, color: '#ff6384' },
    { label: 'In-progress', value: analyticsData.scheduled, color: '#36a2eb' },
    { label: 'Completed', value: analyticsData.completed, color: '#ffce56' }
  ];

  return (
    <div className="flex flex-col p-6 space-y-8">
      {/* Admin Profile */}
      <div className="flex items-center justify-end mb-4">
        <FaUserCircle size={40} className="text-gray-400 mr-2" />
        <span className="text-gray-300">User</span>
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
              <th className="px-4 py-2 text-left">Task Name</th>
              <th className="px-4 py-2 text-left">Description</th>
              <th className="px-4 py-2 text-left">Status</th>
              <th className="px-4 py-2 text-left">Schedule</th>
            </tr>
          </thead>
          <tbody>
            {tasks.map((task, index) => (
              <tr key={index} className="border-b border-gray-600">
                <td className="px-4 py-2">{task.task_name}</td>
                <td className="px-4 py-2">{task.task_description}</td>
                <td className="px-4 py-2">{task.task_status}</td>
                <td className="px-4 py-2">{task.task_schedule}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default UserDashboard;
