# Этап 1: Сборка бинарного файла
FROM golang:1.24 AS builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код
COPY . .

# Компилируем приложение
# Предполагается, что main.go находится в cmd/auth/
RUN CGO_ENABLED=0 GOOS=linux go build -o basket-service ./cmd

# Этап 2: Создаём минимальный образ
FROM alpine:latest

WORKDIR /app

# Копируем скомпилированный бинарный файл из предыдущего этапа
COPY --from=builder /app/basket-service .

# Указываем порт
EXPOSE 8083

# Запускаем приложение
CMD ["./basket-service"]