# Variables
DOCKER_COMPOSE = docker-compose
SQLC_LOCATION = internals/db/sqlc.yaml

# Load Variables from .env
ifneq (,$(wildcard .env))
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


# Help
help:
	@echo "Available commands:"
	@echo "  make up       - Start the services in detached mode"
	@echo "  make down     - Stop the services and remove containers"
	@echo "  make restart  - Restart the services"
	@echo "  make logs     - View logs for the services"
	@echo "  make ps       - List running services"
