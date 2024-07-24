const express = require('express');
const statusMonitor = require('express-status-monitor');
const bodyParser = require('body-parser');
const taskRouter = require('../routers/router');
const { initializeDB } = require('../fs/fs_conn');
const swaggerUi = require('swagger-ui-express');
const swaggerSpec = require('../swagger'); // Path to your Swagger config

async function createApp() {
  await initializeDB();
  const app = express();
  
  app.use(statusMonitor());
  app.use(bodyParser.json());
  app.use('/task_app/task_manager_service/api/v0.1/', taskRouter);

  // Serve Swagger docs
  app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerSpec));

  return app;
}

module.exports = { createApp };
