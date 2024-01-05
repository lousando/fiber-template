FROM golang:1.21.5-alpine

COPY . /app
WORKDIR /app

ENV IS_PROD="true"
EXPOSE 8080

RUN go build -o bin/main

CMD ["/app/bin/main"]