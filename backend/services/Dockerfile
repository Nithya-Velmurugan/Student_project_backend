FROM golang:1.22

WORKDIR /app
COPY . .

RUN go mod init student-service
RUN go mod tidy
RUN go build -o main ./cmd

CMD ["./main"]
