const swaggerJsdoc = require('swagger-jsdoc');
const path = require('path');

const swaggerOptions = {
  definition: {
    openapi: '3.0.0',
    info: {
      title: 'Task Management API',
      version: '0.1.0',
      description: 'API documentation for the Task Management service',
      contact: {
        name: 'Tsegay Negassi',
        email: 'tng.nat2023@gmail.com',
      },
    },
    servers: [
      {
        url: 'http://localhost:3000',
        description: 'Local server',
      },
    ],
  },
  apis: [path.join(__dirname, 'logic/handlers/*.js')], // Path to your JSDoc comments
};

const swaggerSpec = swaggerJsdoc(swaggerOptions);

module.exports = swaggerSpec;
