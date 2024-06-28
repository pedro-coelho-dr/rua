# Use the official Go image
FROM golang:1.17-alpine

# Install necessary build tools
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/api

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]