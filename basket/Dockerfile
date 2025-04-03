FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download


COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o basket-service ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/basket-service .

EXPOSE 8083

CMD ["./basket-service"]