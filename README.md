# Store App - Микросервисная E-commerce Платформа

Современная платформа электронной коммерции, построенная на микросервисной архитектуре с использованием Go и Docker.

## 🏗 Архитектура

Приложение состоит из 4 микросервисов:

- **Auth Service** (Порт: 8085)
  - Управление аутентификацией и авторизацией пользователей
  - Выдача JWT токенов
  - Регистрация и вход пользователей

- **Catalog Service** (Порт: 8081)
  - Управление каталогом товаров
  - CRUD операции с товарами
  - Публичное API для списка товаров

- **Basket Service** (Порт: 8083)
  - Управление корзиной покупок
  - Добавление/удаление товаров из корзины
  - Персональные корзины для пользователей

- **Orders Service** (Порт: 8084)
  - Обработка заказов
  - Интеграция с Kafka для событийной архитектуры
  - Взаимодействие с Catalog service для проверки цен

## 🛠 Технологии

- **Бэкенд**: Go 1.24
- **API**: REST с фреймворком Gin
- **Документация**: Swagger/OpenAPI
- **База данных**: PostgreSQL
- **Брокер сообщений**: Kafka
- **Контейнеризация**: Docker
- **API шлюз**: Traefik
- **UI для Kafka**: Kafka UI

## ⚙️ Требования

- Docker
- Docker Compose
- Go 1.24 (для разработки)

## 🚀 Быстрый старт

1. Клонируйте репозиторий:
```bash
git clone https://github.com/yangirxd/store-app.git
cd store-app
```

2. Настройте переменные окружения:
```bash
cp .env.example .env
```

3. Запустите приложение:
```bash
docker-compose up --build
```

## 🌐 Endpoints сервисов

- **API Gateway**: `http://localhost:80`
- **Auth Service**: `http://localhost:8085`
- **Catalog Service**: `http://localhost:8081`
- **Basket Service**: `http://localhost:8083`
- **Orders Service**: `http://localhost:8084`
- **Kafka UI**: `http://localhost:8086`
- **Traefik Dashboard**: `http://localhost:8080`
- **Единый Swagger UI**: `http://localhost/swagger/`

## 📚 Документация API

Swagger документация доступна в двух вариантах:
1. Единый Swagger UI для всех сервисов: `http://localhost/swagger/`
2. Отдельная документация для каждого сервиса:
   - Auth: `http://localhost:8085/swagger/index.html`
   - Catalog: `http://localhost:8081/swagger/index.html`
   - Basket: `http://localhost:8083/swagger/index.html`
   - Orders: `http://localhost:8084/swagger/index.html`

## 💾 Структура баз данных

Приложение использует отдельные базы данных PostgreSQL для каждого сервиса:
- `auth_db`: Данные аутентификации
- `basket_db`: Данные корзины
- `catalog_db`: Каталог товаров
- `orders_db`: Информация о заказах

## 👨‍💻 Разработка

1. Запустите зависимости:
```bash
docker-compose up -d db kafka
```

2. Запустите каждый сервис:
```bash
cd auth && go run cmd/main.go
cd catalog && go run cmd/main.go
cd basket && go run cmd/main.go
cd orders && go run cmd/main.go
```

## 📁 Структура проекта

```
store-app/
├── auth/           # Сервис аутентификации
├── basket/         # Сервис корзины
├── catalog/        # Сервис каталога
├── orders/         # Сервис заказов
├── docker-compose.yml
└── .env
```

Каждый сервис имеет структуру:
```
service/
├── api/           # HTTP обработчики и маршрутизация
├── cmd/           # Точка входа в приложение
├── db/            # Конфигурация базы данных
├── docs/          # Swagger документация
├── domain/        # Доменные модели
├── repository/    # Слой доступа к данным
├── service/       # Бизнес-логика
└── Dockerfile
```

## 🔐 Безопасность

- Все сервисы используют JWT для аутентификации
- Пароли хешируются перед сохранением
- Traefik обеспечивает безопасную маршрутизацию

## 🔄 Event-Driven взаимодействие

- Orders Service публикует события в Kafka при создании заказов
- Catalog Service обновляет остатки товаров при получении событий
- Kafka UI для мониторинга очередей сообщений

## 🤝 Contributing

1. Fork репозиторий
2. Создайте ветку для фичи (`git checkout -b feature/amazing-feature`)
3. Зафиксируйте изменения (`git commit -m 'Add amazing feature'`)
4. Push в ветку (`git push origin feature/amazing-feature`)
5. Откройте Pull Request

## 📜 Лицензия

MIT License

## 📞 Контакты

yangirxd - [@yangirxd](https://github.com/yangirxd)