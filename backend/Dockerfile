# Use an official Golang runtime as a parent image
FROM golang:1.22 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Build the Go application
RUN go build -o /app/bin/auth cmd/main.go

# Start a new stage for the runtime image
FROM golang:1.22

# Copy the built binary from the builder stage
COPY --from=builder /app/bin/auth /app/bin/auth

# Set the working directory and execute the binary
WORKDIR /app
CMD ["/app/bin/auth"]
