# Makefile для Finance App
include .env
export $(shell sed 's/=.*//' .env)

.PHONY: build up down logs migrate-up migrate-down docker-full

# Переменные
DOCKER_COMPOSE = docker-compose
DB_URL = postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5433/${POSTGRES_DB}?sslmode=disable
DOCKER_DB_URL = postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable

# Локальная разработка
build:
	go build -o finance-app ./cmd/app/main.go

up: build migrate-up
	./finance-app

migrate-up:
	migrate -path ./migrations -database "${DB_URL}" up

migrate-down:
	migrate -path ./migrations -database "${DB_URL}" down

# Docker-команды
docker-build:
	${DOCKER_COMPOSE} build

docker-up:
	${DOCKER_COMPOSE} up -d

docker-down:
	${DOCKER_COMPOSE} down

docker-logs:
	${DOCKER_COMPOSE} logs -f finance_app

docker-full: docker-down docker-build docker-up

# Миграции внутри Docker
docker-migrate-up:
	docker-compose exec finance_app migrate -path /finance_app/migrations -database "${DOCKER_DB_URL}" up

docker-migrate-down:
	docker-compose exec finance_app migrate -path /finance_app/migrations -database "${DOCKER_DB_URL}" down