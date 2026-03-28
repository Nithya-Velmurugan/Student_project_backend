# -------- BUILD STAGE --------
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Only copy go mod first (cache optimization)
COPY go.mod go.sum ./
RUN go mod download

# Copy remaining code
COPY . .

# Build binary
RUN go build -o app ./services/cmd

# -------- FINAL STAGE --------
FROM alpine:latest

WORKDIR /app

# Copy only binary (VERY SMALL IMAGE)
COPY --from=builder /app/app .

EXPOSE 8082

CMD ["./app"]