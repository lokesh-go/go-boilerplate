# Project variables
VERSION := $(shell cat VERSION)
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)
APP_NAME := go-api-microservice
ENTRYPOINT := cmd/go-api-microservice
BUILD_DIR := .build
BINARY_NAME := $(APP_NAME)-$(VERSION)
DOCKER_IMAGE := $(APP_NAME):$(VERSION)
DOCKERFILE := docker/Dockerfile
PORT := 80
INTERNAL_PORT := 8080
ENV ?= dev
DEFAULT_BRANCH ?= master

.PHONY: build run build-docker run-docker clean test format release help

build:
	@echo "🛠️  Building $(APP_NAME) for $(GOOS)/$(GOARCH) [ENV=$(ENV)]"
	@mkdir -p $(BUILD_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) \
	go build -ldflags "-s -w" \
	-o $(BUILD_DIR)/$(BINARY_NAME) ./$(ENTRYPOINT)

run:
	@echo "🚀 Running $(APP_NAME) with ENV=$(ENV)"
	ENV="$(ENV)" go run $(ENTRYPOINT)/main.go

build-docker:
	@echo "🐳 Building Docker image: $(DOCKER_IMAGE)"
	docker build --build-arg ENV=$(ENV) -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .

run-docker:
	@echo "🐳 Running Docker container: $(DOCKER_IMAGE)"
	docker run --rm \
		-p $(PORT):$(PORT) \
		-p $(INTERNAL_PORT):$(INTERNAL_PORT) \
		--env ENV=$(ENV) \
		$(DOCKER_IMAGE)

clean:
	@echo "🧹 Cleaning build artifacts"
	@go clean -cache -modcache
	@rm -rf $(BUILD_DIR)
	@echo "✅ Clean complete."

test:
	@echo "🧪 Running test cases"
	go test ./test/...

format:
	@echo "🧼 Formatting source code"
	go fmt ./...

release:
	@echo "🔁 Switching to $(DEFAULT_BRANCH) branch..."
	@git checkout $(DEFAULT_BRANCH)
	@git pull origin $(DEFAULT_BRANCH)

	@echo "🌿 Creating bump/release branch..."
	@git checkout -b bump/release

	@echo "🚀 Bumping version using: $(bump)"
	@./scripts/bump_version.sh $(bump)

	@git add VERSION
	@git commit -m "chore: bump version to $(shell cat VERSION)"
	@git tag $(shell cat VERSION)

	@echo "📤 Pushing release branch and tag..."
	@git push origin bump/release
	@git push origin $(shell cat VERSION)

	@echo "✅ Release pushed successfully!"

help:
	@echo "📘 Available Makefile commands:"
	@echo "  build         - Build Go binary into .build/ [default ENV=dev]"
	@echo "  run           - Run app using go run"
	@echo "  build-docker  - Build Docker image from Dockerfile"
	@echo "  run-docker    - Run Docker container locally"
	@echo "  clean         - Clean Go caches and .build directory"
	@echo "  test          - Run test suite"
	@echo "  format        - Format code using go fmt"
	@echo "  release       - Usage: make release bump=patch|minor|major"
	@echo "  help          - Show this help message"
