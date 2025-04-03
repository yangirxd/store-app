FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o orders-service ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/orders-service .

EXPOSE 8084

CMD ["./orders-service"]