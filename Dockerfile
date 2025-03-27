# Этап сборки
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd

# Финальный контейнер
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

# Порт, используемый Gin
EXPOSE 5000

CMD ["./main"]
