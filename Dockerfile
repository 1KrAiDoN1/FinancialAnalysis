# Используем многоэтапную сборку
FROM golang:1.24.3-alpine AS builder

# Устанавливаем инструменты для сборки
RUN apk add --no-cache git build-base

# Устанавливаем migrate
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
WORKDIR /finance_app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o /finance_app/finance-app ./cmd/app/main.go

# Финальный образ
FROM alpine:latest

WORKDIR /finance_app
# Устанавливаем инструменты для работы с БД и миграциями
RUN apk add --no-cache postgresql-client bash

# Копируем бинарник и миграции
COPY --from=builder /finance_app/finance-app .
COPY --from=builder /finance_app/migrations ./migrations
COPY --from=builder /finance_app/internal/config /finance_app/internal/config
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate

# Проверяем установку migrate
RUN migrate -version

# # Устанавливаем инструменты для работы с БД (опционально)
# RUN apk add --no-cache postgresql-client

# Порт приложения
EXPOSE 8081

# Команда запуска
CMD ["sh", "-c", "while ! pg_isready -h db -U ${POSTGRES_USER}; do sleep 1; done && migrate -path /finance_app/migrations -database ${DB_URL} up && ./finance-app"]
