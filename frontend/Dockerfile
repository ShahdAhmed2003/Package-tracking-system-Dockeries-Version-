# Use the official Node.js image
FROM node:18-alpine

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY frontend/package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Serve the application using a lightweight web server
RUN npm install -g serve
CMD ["serve", "-s", "build", "-l", "3000"]

# Expose the port
EXPOSE 3000
