APP_NAME=subs-api
DATABASE_URL?=postgres://postgres:postgres@localhost:5432/subscription_service?sslmode=disable

.PHONY: run test migrate-up migrate-down docker-up docker-down

run:
	go run ./cmd/api

test:
	go test ./...

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

docker-up:
	docker compose up -d

docker-down:
	docker compose down
