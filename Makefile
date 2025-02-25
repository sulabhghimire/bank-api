# Variables
DOCKER_COMPOSE = docker compose
SQLC_LOCATION = internals/db/sqlc.yaml

# Load Variables from .env
ifneq (,$(wildcard app.env))
    include app.env
    export $(shell sed 's/=.*//' app.env)
endif

#Goose
migrate-up-all:
	goose up

migrate-up-one:
	goose up-by-one

migrate-down-one:
	goose down

migrate-down-all:
	goose down-to 0

#Database
database-create:
	@docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=admin $(DB_NAME)

database-delete:
	@docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb --username=$(DB_USER) $(DB_NAME)
 

#SQLC
sqlc:
	sqlc generate

# Docker
docker-up:
	@$(DOCKER_COMPOSE) up -d

docker-down:
	@$(DOCKER_COMPOSE) down

docker-restart:
	@$(DOCKER_COMPOSE) down
	@$(DOCKER_COMPOSE) up -d

docker-logs:
	@$(DOCKER_COMPOSE) logs -f

docker-ps:
	@$(DOCKER_COMPOSE) ps

# GO
test:
	@echo "Running migrations..."
	@env GOOSE_DRIVER=postgres GOOSE_DBSTRING=postgresql://admin:admin@localhost:5433/test_db GOOSE_MIGRATION_DIR=./internals/db/migrations goose up
	@echo "Migrations completed."

	@echo "Running tests..."
	@env DB_DRIVER=postgres DB_SOURCE=postgresql://admin:admin@localhost:5433/test_db?sslmode=disable go test -v -cover ./...
	@echo "Tests completed."

test-ci-cd:
	@env DB_DRIVER=postgres DB_SOURCE=postgresql://admin:admin@localhost:5433/test_db?sslmode=disable \
	go test -v -cover ./...

server:
	go run ./cmd/bank-api/main.go

# MOCK GEN
mock:
	@mockgen --package mockdb --build_flags=--mod=mod --destination internals/db/mock/store.go github.com/sulabhghimire/bank-api/internals/db/sqlc Store 

# Help
help:
	@echo "Available commands"
	@echo "************************************************************************"
	@echo "Docker commands"
	@echo "  make docker-up       - Start the services in detached mode"
	@echo "  make docker-down     - Stop the services and remove containers"
	@echo "  make docker-restart  - Restart the services"
	@echo "  make docker-logs     - View logs for the services"
	@echo "  make docker-ps       - List running services"
	@echo "************************************************************************"
	@echo "SQLC commands"
	@echo "  make sqlc            - Generates the queries via sqlc"
	@echo "************************************************************************"
	@echo "Database commands"
	@echo "  make database-create - Creates the required database for the project"
	@echo "  make database-delete - Deletes the created database by the project"
	@echo "************************************************************************"
	@echo "Goose commands"
	@echo "  make migrate-up	  - Makes one up migration to the database"
	@echo "  make migrate-down	  - Makes one down migration to the database"
	@echo "************************************************************************"
	@echo "GO commands"
	@echo "  make test			  - Run all the tests for all packages"
	@echo "  make server		  - Run GO HTTP Server"
	@echo "************************************************************************"