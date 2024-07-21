const express = require('express');
const statusMonitor = require('express-status-monitor');
const bodyParser = require('body-parser');
const taskRouter = require('../routers/router');
const { initializeDB } = require('../fs/fs_conn');

async function createApp() {
  await initializeDB();
  const app = express();
  app.use(statusMonitor());
  app.use(bodyParser.json());
  app.use('/task_app/task_manager_service/api/v0.1/', taskRouter);
  return app;
}

module.exports = { createApp };
