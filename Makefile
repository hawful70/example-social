include .env
MIGRATIONS_PATH = ./cmd/migrate/migrations

clean:
	@echo "Cleaning..."
	@docker compose -f .container/docker-compose.yml down --remove-orphans
	@docker volume rm -f container_social-postgres-data
	@docker volume rm -f container_social-redis-data

migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

seed: 
	@go run cmd/migrate/seed/main.go

infras:
	@docker compose -f .container/docker-compose.yml down --remove-orphans
	@docker compose -f .container/docker-compose.yml up -d