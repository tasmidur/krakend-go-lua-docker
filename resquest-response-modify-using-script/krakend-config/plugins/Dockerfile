# Use the official Golang image as a build stage
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .


# Run go mod tidy to clean up go.mod and go.sum
RUN go mod tidy

# Build the Go app
# RUN go build -buildmode=plugin -o modifier.so modifier.go
RUN go build -buildmode=plugin -o myplugin.so myplugin.go
# Use a minimal image to copy the plugin
# FROM alpine:latest

# # Copy the built plugin from the builder stage
# COPY --from=builder /app/modifier.so /etc/krakend/plugins/modifier.so
