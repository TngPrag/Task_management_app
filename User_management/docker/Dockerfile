# Use the official Golang image as a base image
FROM golang:1.22.1

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o user_management_service main.go

# Expose the port on which the service will run
EXPOSE 8981

# Command to run the application
CMD ["./user_management_service"]
