const axios = require('axios');

/**
 * GetUserRole sends an HTTP request to get the user's role.
 * @param {string} token - The authorization token.
 * @returns {Promise<string>} - A promise that resolves to the user's role.
 */
async function GetUserRole(token) {
  const url = 'http://localhost:8980/task_manager_app/auth_service/v0.1/role/read';

  try {
    const response = await axios.get(url, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    });

    if (response.status !== 200) {
      throw new Error(`Failed to get user role: status code ${response.status}`);
    }

    return response.data.role;
  } catch (error) {
    console.error(`Error getting user role: ${error.message}`);
    throw error;
  }
}

module.exports = {
  GetUserRole,
};
