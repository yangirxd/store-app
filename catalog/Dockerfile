FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o catalog-service ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/catalog-service .

EXPOSE 8081

CMD ["./catalog-service"]