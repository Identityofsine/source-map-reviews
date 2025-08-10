# Use multi-stage build for security and smaller image size
FROM golang:1.23.1-bookworm AS builder

# Creates an app directory to hold your app's source code
WORKDIR /usr/src/app

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/server/main.go

# Use a minimal base image for the final stage
FROM alpine:3.19

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /usr/src/app/app .
# Copy configuration files
COPY --from=builder /usr/src/app/config ./config
# Copy migrations folder (required for database setup)
COPY --from=builder /usr/src/app/migrations ./migrations

# Tells Docker which network port your container listens on
EXPOSE 3001 

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3001/health || exit 1

# Specifies the executable command that runs when the container starts
CMD ["./app"]
