# FROM alpine:latest

# WORKDIR /app

# COPY app .

# EXPOSE 8082

# CMD ["./app"]

FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./services/cmd

EXPOSE 8082

CMD ["./main"]