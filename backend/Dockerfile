# backend/Dockerfile

# Stage 1: Build the Go binary
FROM golang:1.18 as builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Run the Go binary on a minimal image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Ensure the binary is executable
RUN chmod +x ./main

# Run the binary
CMD ["./main"]
