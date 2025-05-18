# Project variables
APP_NAME=go-api-microservice
ENTRYPOINT=cmd/go-api-microservice/main.go
DOCKER_IMAGE=$(APP_NAME):latest
DOCKERFILE=docker/Dockerfile
PORT=80
INTERNAL_PORT=8080
ENV ?= dev

# Default target
.PHONY: run
run:
	@echo "Running $(APP_NAME) with ENV=$(ENV)"
	ENV=$(ENV) go run $(ENTRYPOINT)

.PHONY: build-docker
build-docker:
	@echo "Building Docker image $(DOCKER_IMAGE)"
	docker build --build-arg ENV=$(ENV) -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .

.PHONY: run-docker
run-docker:
	@echo "Running Docker container $(DOCKER_IMAGE)"
	docker run --rm -p $(PORT):$(PORT) -p $(INTERNAL_PORT):$(INTERNAL_PORT) --env ENV=$(ENV) $(DOCKER_IMAGE)

.PHONY: clean
clean:
	@echo "Cleaning build artifacts"
	rm -f $(APP_NAME)

.PHONY: help
help:
	@echo "Makefile commands:"
	@echo "  run           - Run using go run (default ENV=dev)"
	@echo "  build-docker  - Build Docker image"
	@echo "  run-docker    - Run Docker container with exposed port"
	@echo "  clean         - Remove local build artifacts"
	@echo "  help          - Show this help message"
