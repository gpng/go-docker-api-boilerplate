# Docker parameters
DOCKERCMD=docker
DOCKERCOMPOSECMD=docker-compose

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

# App parameters
CONTAINER_TAG=go-docker-api-boilerplate
DOCKERFILE=Dockerfile.dev
MAIN_FOLDER=cmd/api
MAIN_PATH=$(MAIN_FOLDER)/main.go

# Goose parameters for migrations
GOOSECMD=goose
DBSTRING="host=${DB_HOST} user=${DB_USER} dbname=${DB_NAME} sslmode=disable password=${DB_PASSWORD}"
MIGRATIONSPATH=db/migrations/

default: build up logs

build:
	@echo "=============building API============="
	$(DOCKERCMD) build -f $(DOCKERFILE) -t $(CONTAINER_TAG) .

up:
	@echo "=============starting API locally============="
	$(DOCKERCOMPOSECMD) up -d

logs:
	$(DOCKERCOMPOSECMD) logs -f

run:
	go build -o bin/application $(MAIN_PATH) && ./bin/application -docs

down:
	$(DOCKERCOMPOSECMD) down --remove-orphans

test:
	godotenv -f .test.env $(GOTEST) -cover ./... -count=1

clean: down
	@echo "=============cleaning up============="
	$(DOCKERCMD) system prune -f
	$(DOCKERCMD) volume prune -f

run-prod:
	$(DOCKERCMD) build -t chanced-api-eb .
	docker run -p 4000:5000 chanced-api-eb

gen-docs:
	swag init -g $(MAIN_PATH)

migrate:
	$(GOOSECMD) -dir $(MIGRATIONSPATH) postgres $(DBSTRING) up

rollback:
	$(GOOSECMD) -dir $(MIGRATIONSPATH) postgres $(DBSTRING) down

rollback-all:
	$(GOOSECMD) -dir $(MIGRATIONSPATH) postgres $(DBSTRING) down-to 0

reset: rollback-all migrate