const axios = require('axios');

/**
 * Middleware to retrieve user profile from an external service.
 * @param {string} token - The authorization token.
 * @returns {Promise<Object>} The user profile.
 */
async function authenticate(token) {
  try {
    const response = await axios.get('http://localhost:8981/task_app/user_manager_service/api/v0.1/user/verify', {
      headers: {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    return response.data;
  } catch (error) {
    throw new Error('Failed to retrieve user profile');
  }
}

module.exports = { authenticate };
