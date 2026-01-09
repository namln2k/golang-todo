# =========================
# Stage 1: Build
# =========================
FROM golang:1.25-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install build tools
RUN apk add --no-cache git tzdata

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Compile the binary as a static executable
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server


# =========================
# Stage 2: Run
# =========================
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/server .

# Expose the port the app runs on (adjust if different)
EXPOSE 8888


# Ensure non-root security best practice
RUN adduser -D appuser
USER appuser

# Command to run the executable
CMD ["./server"]
