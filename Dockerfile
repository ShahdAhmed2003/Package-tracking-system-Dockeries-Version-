# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

#set the working directory inside the container
WORKDIR /app

#copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

#download and cache Go modules
RUN go mod download

#copy the application source code to the container
COPY . .

#to set the working directory to the directory that contains the main.go file
WORKDIR /app/cmd/api

#to build the Go application
RUN go build -o /app/main .

#stage 2:lightweight final image
FROM alpine:3.18

#the working directory for the application
WORKDIR /app

#copy the built application from the builder stage
COPY --from=builder /app/main .

#the port Go application runs on
EXPOSE 8080

#to run the application
CMD ["./main"]