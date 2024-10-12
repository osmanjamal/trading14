# Start from the official Go image
FROM golang:1.17-alpine

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/server

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]