const axios = require('axios');

/**
 * VerifyPolicy sends an HTTP request to verify if a policy allows a certain action.
 * @param {string} token - The authorization token.
 * @param {string} sub - The subject.
 * @param {string} obj - The object.
 * @param {string} act - The action.
 * @returns {Promise<boolean>} - A promise that resolves to whether the action is allowed.
 */
async function verifyPolicy(token, sub, obj, act) {
  const url = 'http://localhost:8980/task_app/authz_service/api/v0.1/policy/check_permission';

  const payload = {
    sub: sub,
    object: obj,
    action: act,
  };

  try {
    const response = await axios.post(url, payload, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    });

    if (response.status !== 200) {
      throw new Error(`Failed to verify policy: status code ${response.status}`);
    }

    return response.data;
  } catch (error) {
    console.error(`Error verifying policy: ${error}`);
    throw error;
  }
}

module.exports = {
  verifyPolicy,
};
