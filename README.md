# Subscription Billing Service

REST API for managing users, plans, subscriptions, payments, and subscription history.

## Stack

- Go
- PostgreSQL
- pgx
- Docker Compose
- golang-migrate
- JWT
- Swagger
- slog
- Chi

## First Steps

```sh
docker compose up -d
go run ./cmd/api
```

Health check:

```sh
curl http://localhost:8080/health
```
