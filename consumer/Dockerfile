FROM golang:1.24.1-alpine3.21 as build

RUN apk add --no-cache git

WORKDIR /src

RUN go mod init go-consumer
RUN go get github.com/rabbitmq/amqp091-go

COPY main.go .
RUN go build -o consumer main.go


FROM alpine as runtime

COPY --from=build /src/consumer /app/consumer


CMD [ "/app/consumer"]
