# Use the official Golang image as a base
FROM golang:1.22

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files for dependency management
COPY go.mod ./

# Download all dependencies; layer caching helps speed up future builds if dependencies don't change
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o go-lru-cache ./cmd/main.go

# Expose a port if your application serves over HTTP or similar (optional)
# EXPOSE 8080 

# Start the compiled binary
CMD ["./go-lru-cache"]
