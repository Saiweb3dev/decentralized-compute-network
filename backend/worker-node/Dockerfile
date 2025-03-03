# Build stage
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies
RUN apk add --no-cache git

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /worker-node ./cmd/worker

# Final stage
FROM alpine:3.17

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates

# Copy binary from builder
COPY --from=builder /worker-node /usr/local/bin/

# Create a non-root user
RUN adduser -D -g '' appuser
USER appuser

# Run the application
ENTRYPOINT ["worker-node"]