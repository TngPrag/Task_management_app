const axios = require('axios');

async function notifyUser(email, title, description, deadline, authToken) {
  const url = 'http://localhost:3000/task_app/user_manager_service/api/v0.1/user/notify'; // Replace with your actual URL
  const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${authToken}`
  };
  const body = {
    email,
    title,
    description,
    deadline
  };

  try {
    const response = await axios.post(url, body, { headers });
    if (response.status === 200) {
      return 'User has been notified successfully about the task assignment';
    } else {
      throw new Error(`Failed to notify user: ${response.data}`);
    }
  } catch (error) {
    throw new Error(`Error notifying user: ${error.message}`);
  }
}

module.exports = {
    notifyUser,
  };