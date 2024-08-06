# Use a specific Node.js version as a base image
FROM node:22.4.1-slim

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json from the parent directory
COPY ../package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code from the parent directory
COPY ../ ./

# Expose the port the app runs on (adjust if needed)
EXPOSE 8903

# Command to run the application (adjust if needed)
CMD ["node", "main.js"]
