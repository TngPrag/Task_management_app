const { createApp } = require('./app/setup');

async function startServer() {
  const app = await createApp();
  const port = process.env.PORT || 8903;
  app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
  });
}

startServer();
