import axios from 'axios';

const API_BASE_URL = 'http://localhost:3000'; // Replace with your API base URL

// Fetch tasks
export const listTasks = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/tasks`);
    return response.data;
  } catch (error) {
    console.error('Error fetching tasks:', error);
    throw error;
  }
};

// Create a new task
export const createTask = async (taskData) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/tasks`, taskData);
    return response.data;
  } catch (error) {
    console.error('Error creating task:', error);
    throw error;
  }
};

// Delete a specific task
export const deleteTask = async (taskId) => {
  try {
    const response = await axios.delete(`${API_BASE_URL}/tasks/${taskId}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting task:', error);
    throw error;
  }
};

// Delete all tasks
export const deleteAllTasks = async () => {
  try {
    const response = await axios.delete(`${API_BASE_URL}/tasks`);
    return response.data;
  } catch (error) {
    console.error('Error deleting all tasks:', error);
    throw error;
  }
};

// Fetch users
export const listUsers = async () => {
  try {
    const response = await axios.get(`${API_BASE_URL}/users`);
    return response.data;
  } catch (error) {
    console.error('Error fetching users:', error);
    throw error;
  }
};
