const axios = require('axios');

async function notifyUser(user_id, title, description, deadline, authToken) {
  const url = 'http://localhost:8981/task_app/user_manager_service/api/v0.1/user/notify'; // Updated port
  const headers = {
    'Content-Type': 'application/json',
    'Authorization': `Bearer ${authToken}`
  };
  const body = {
    "user_id":user_id,
    "title": title,
    "description": description,
    "deadline": deadline
  };

  try {
    const response = await axios.post(url, body, { headers });
    if (response.status === 200) {
      return 'User has been notified successfully about the task assignment';
    } else {
      throw new Error(`Failed to notify user: ${response.data}`);
    }
  } catch (error) {
    console.error('Error response data:', error.response.data);
    throw new Error(`Error notifying user: ${error.message}`);
  }
}

module.exports = {
  notifyUser,
};
