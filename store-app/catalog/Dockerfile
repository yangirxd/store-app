# Этап 1: Сборка бинарного файла
FROM golang:1.24 AS builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Компилируем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o catalog-service ./cmd

# Этап 2: Создаём минимальный образ
FROM alpine:latest

WORKDIR /app

# Копируем скомпилированный бинарный файл из предыдущего этапа
COPY --from=builder /app/catalog-service .

# Указываем порт
EXPOSE 8081

# Запускаем приложение
CMD ["./catalog-service"]