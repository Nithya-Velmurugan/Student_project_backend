FROM golang:1.24-alpine

WORKDIR /app

COPY . .

# ❌ REMOVE THIS LINE
# RUN go mod init student-service

RUN go mod tidy
RUN go build -o app ./services/cmd

EXPOSE 8080

CMD ["go", "run", "./services/cmd/main.go"]