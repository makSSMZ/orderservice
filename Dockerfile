# Используем Go 1.24 для сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o orderservice ./cmd/main.go

FROM alpine:latest

RUN apk add --no-cache bash
WORKDIR /app
COPY --from=builder /app/orderservice .
COPY ./web ./web

COPY wait-for-it.sh /app/wait-for-it.sh
RUN chmod +x /app/wait-for-it.sh

ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_USER=order_user
ENV DB_PASSWORD=password
ENV DB_NAME=orders_db
ENV KAFKA_BROKER=kafka:9092
ENV KAFKA_TOPIC=order-topic

CMD ["./orderservice"]
