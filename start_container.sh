#!/bin/bash

# Function to stop the container if it is running
docker pull postgres:latest

docker pull mennaa/go-backend

docker pull maryam197/frontend:latest

docker network create my-network

stop_container_if_running() {
  container_name=$1
  if [ "$(docker ps -q -f name=${container_name})" ]; then
    echo "Stopping running container: ${container_name}..."
    docker stop ${container_name}  # Stop the running container
  fi
}

# Function to start the container if it exists, or create it if it doesn't exist
start_or_create_container() {
  container_name=$1
  container_exists=$(docker ps -aq -f name=${container_name})

  if [ -n "$container_exists" ]; then
    echo "Starting existing container: ${container_name}..."
    docker start ${container_name}  # Start the existing container
  else
    echo "Creating new container: ${container_name}..."
    # The command to create the container should go here
    # Example:
    if [ "$container_name" == "my-postgres-db" ]; then
      docker run --name my-postgres-db --network my-network -v C:/Users/DELL/Downloads/Volume:/var/lib/postgresql/data -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=asdqwe123 -e POSTGRES_DB=bosta -p 5432:5432 -d postgres:latest
    elif [ "$container_name" == "my-go-backend" ]; then
      docker run --name my-go-backend --network my-network -e DB_HOST=my-postgres-db -e DB_PORT=5432 -e DB_USER=postgres -e DB_PASSWORD=asdqwe123 -e DB_NAME=bosta -p 8080:8080 -d mennaa/go-backend
    elif [ "$container_name" == "frontend-container" ]; then
      docker run --name frontend-container --network my-network -p 8081:80 -d maryam197/frontend:latest
    fi
  fi
}

# Start PostgreSQL container
stop_container_if_running "my-postgres-db"
start_or_create_container "my-postgres-db"
echo "PostgreSQL container started or resumed."

# Wait for PostgreSQL to initialize
sleep 15

# Start Backend container
stop_container_if_running "my-go-backend"
start_or_create_container "my-go-backend"
echo "Backend container started or resumed."

# Wait for Backend to initialize
sleep 15

# Start Frontend container
stop_container_if_running "frontend-container"
start_or_create_container "frontend-container"
echo "Frontend container started or resumed."
