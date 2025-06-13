FROM golang:1.23-alpine

WORKDIR /app

COPY ./build/task-api-linux-amd64 .
COPY ./utils/config/config.ini .

EXPOSE 8080

CMD ["./task-api-linux-amd64", "--config", "config.ini"]