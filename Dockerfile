# Use the official Node.js 14 image as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Install the Go dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Expose the port on which the application will run
EXPOSE 8080

# Start the application
CMD [ "go", "run", "./..." ]