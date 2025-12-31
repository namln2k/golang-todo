# Build stage
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Compile the binary as a static executable
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

# Runtime stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /main .

# Expose the port the app runs on (adjust if different)
EXPOSE 8888

# Command to run the executable
CMD ["./main"]
