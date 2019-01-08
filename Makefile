# Docker parameters
DOCKERCMD=docker
DOCKERCOMPOSECMD=docker-compose

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

# App parameters
CONTAINER_TAG=go-docker-api-boilerplate
DOCKERFILE=Dockerfile
MAIN_PATH=cmd/api/main.go
DEPLOY_FOLDER=deploy
DEPLOY_BINARY_NAME=$(DEPLOY_FOLDER)/application
ZIP_PATH=$(DEPLOY_FOLDER)/deploy-$(shell date +'%Y%m%d-%H%M%S').zip

default:
	@echo "=============building API============="
	$(DOCKERCMD) build -f $(DOCKERFILE) -t $(CONTAINER_TAG) .

up: default
	@echo "=============starting API locally============="
	$(DOCKERCOMPOSECMD) up -d

logs:
	$(DOCKERCOMPOSECMD) logs -f

down:
	$(DOCKERCOMPOSECMD) down

test:
	$(GOTEST) -v -cover ./...

clean: down
	@echo "=============cleaning up============="
	$(DOCKERCMD) system prune -f
	$(DOCKERCMD) volume prune -f

deploy-prod:
	GOARCH=amd64 GOOS=linux $(GOBUILD) -o $(DEPLOY_BINARY_NAME) $(MAIN_PATH)
	cp .env.prod $(DEPLOY_FOLDER)/.env
	zip $(ZIP_PATH) $(DEPLOY_BINARY_NAME) $(DEPLOY_FOLDER)/.env $(DEPLOY_FOLDER)/Procfile