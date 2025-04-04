# Use official golang image as a parent image
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy the modules and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o app cmd/server/main.go

# Use a smaller image for the final build
FROM alpine:latest

# Set working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Copy env file
COPY .env .

# Expose the port 8080 (app)
EXPOSE 8080

# Run the app, redis, and mysql server
CMD ["/root/app"]