# Makefile for managing the Docker Compose workflow of the fullstack-auth application

# Specify default target.
all: build

# Phony targets for cleanliness.
.PHONY: build up down restart logs ps frontend backend database

# Build the application.
build:
	docker-compose build

# Start up all services in the background.
up:
	docker-compose up -d

# Stop all services.
down:
	docker-compose down

# Restart all services.
restart: down up

# Follow log output from containers.
logs:
	docker-compose logs -f

# List all running containers related to the compose configuration.
ps:
	docker-compose ps

# Shortcut for working with the frontend service.
frontend:
	docker-compose exec frontend sh

# Shortcut for working with the backend service.
backend:
	docker-compose exec backend sh

# Shortcut for accessing the database command line.
database:
	docker-compose exec database psql -U postgres
