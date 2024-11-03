# Build Stage
FROM golang:1.22 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go app with CGO disabled for static binary
ENV CGO_ENABLED=0
RUN go build -o gowebserver gws.go

# Final Stage
FROM debian:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/gowebserver .

# Command to run the executable
CMD ["./gowebserver"]

