async function checkHealth(req, res) {
    res.json({ status: 'Task manager service is running Ok!' });
  }
  
  module.exports = {
    checkHealth,
  };
  