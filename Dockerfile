FROM alpine:latest

WORKDIR /app

COPY app .

EXPOSE 8082

CMD ["./app"]