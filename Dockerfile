# -------- BUILD STAGE --------
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./services/cmd

# -------- FINAL STAGE --------
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8082

CMD ["./app"]