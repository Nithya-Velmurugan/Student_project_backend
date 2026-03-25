# Build stage
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go Module files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./services/cmd/main.go

# Run stage
FROM alpine:latest

# Add ca-certificates for secure connections (often needed for Go apps)
RUN apk --no-cache add ca-certificates

# Working directory for the application
WORKDIR /root/

# Copy the pre-built binary
COPY --from=builder /app/main .

# Expose port
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
