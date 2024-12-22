# Step 1: Build the React app
FROM node:16 AS build

# Set the working directory for the build process
WORKDIR /app

# Copy package.json and package-lock.json for dependency installation
COPY frontend/package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the frontend code
COPY frontend/ .

# Build the React app for production
RUN npm run build

# Step 2: Use Nginx to serve the built files
FROM nginx:alpine

# Create required directories for Nginx to avoid permission issues
RUN mkdir -p /var/cache/nginx/client_temp && \
    chown -R nginx:nginx /var/cache/nginx

# Copy the build files from the previous stage into the Nginx directory
COPY --from=build /app/build /usr/share/nginx/html

# Expose port 80 for the frontend
EXPOSE 80

# Set appropriate permissions for Nginx to serve the files
RUN chown -R nginx:nginx /usr/share/nginx/html

# Start the Nginx server to serve the frontend app
CMD ["nginx", "-g", "daemon off;"]
