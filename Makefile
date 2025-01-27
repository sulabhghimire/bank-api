# Variables
DOCKER_COMPOSE = docker compose
SQLC_LOCATION = internals/db/sqlc.yaml

# Load Variables from .env
ifneq (,$(wildcard app.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

#Goose
migrate-up:
	goose up

migrate-down:
	goose down

#Database
database-create:
	@docker exec -it $(POSTGRES_CONTAINER_NAME) createdb --username=$(DB_USER) --owner=admin $(DB_NAME)

database-delete:
	@docker exec -it $(POSTGRES_CONTAINER_NAME) dropdb --username=$(DB_USER) --owner=admin $(DB_NAME)
 

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
	go test -v -cover ./...

server:
	go run ./cmd/bank-api/main.go

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