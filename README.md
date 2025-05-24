# go-api-microservice

A Go microservice boilerplate that supports both REST and gRPC APIs. No need to worry about project structure—just clone the repo and start writing your business logic.

## Core Packages

### HTTP Framework

- **[github.com/gin-gonic/gin](https://github.com/gin-gonic/gin) v1.10.0**
  - High-performance HTTP web framework
  - Provides middleware support and routing
  - Built-in input validation and error handling

### Application Cache

- **[github.com/patrickmn/go-cache](https://github.com/patrickmn/go-cache) v2.1.0**
  - In-memory key-value store
  - Supports TTL and automatic cleanup
  - Thread-safe operations

### Configuration Management

- **[gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3) v3.0.1**
  - YAML configuration parser
  - Supports environment variable interpolation
  - Type-safe configuration loading

## Project Structure

```bash
.
├── api/               # API layer: REST and gRPC endpoints
│   ├── grpc/          # gRPC service implementations
│   │   └── grpc.go    # gRPC service entry point
│   │
│   └── rest/           # REST API components
│       ├── handler/    # HTTP request handlers
│       ├── middleware/ # Request processing middleware
│       └── router/     # URL routing definitions
│
├── cmd/                 # Application entry points
│   ├── go-api-microservice/  # Main API server executable
│   │   └── main.go      # Server initialization
│   │
│   └── cli/          # Command-line tools
│       └── main.go   # CLI entry point
│
├── docker/           # Container configuration
│   └── Dockerfile    # Multi-stage build definition
│
├── internal/        # Private application code
│   ├── bootstrap/   # App initialization & setup
│   │   ├── env.go   # Environment configuration
│   │   └── init.go  # Bootstrap sequence
│   │
│   ├── config/       # Configuration management
│   │   ├── config.go # Config loading logic
│   │   └── models.go # Config structure definitions
│   │
│   ├── dal/                # Data Access Layer
│   │   ├── cache/          # Caching implementations
│   │   │   ├── appcache/   # Local memory cache
│   │   │   └── rediscache/ # Redis cache client
│   │   │
│   │   ├── db/             # Database implementations
│   │   │   ├── mongo/      # MongoDB client & operations
│   │   │   └── mysql/      # MySQL client & operations
│   │   │
│   │   └── entities/       # Domain model definitions
│   │       ├── address/    # Address related models
│   │       └── user/       # User related models
│   │
│   ├── dependencies/       # Dependency injection container
│   │
│   ├── models/      # Shared domain models
│   │
│   │── modules/     # Business logic for all entities
│   │
│   └── server/       # Server implementations
│       ├── grpc/     # gRPC server setup
│       ├── http/     # HTTP server setup
│       └── server.go # Server interface definitions
│
├── pkg/              # Shared libraries and utilities
│   ├── cache/        # Generic caching interfaces
│   ├── db/           # Database utilities and interfaces
│   ├── logger/       # Logging utilities and abstractions
│   └── utils/        # Common helper and utility functions

├── scripts/          # Build and deployment scripts
│   ├── build.sh      # Build automation
│   └── run.sh        # Runtime scripts
│
├── test/             # Test suites and test data
|
|-- go.mod
|
|-- go.sum
|
|__ Makefile

```

## Getting Started

### Prerequisites

- Go 1.23 or higher

### Project Setup

1. Clone the repository:

```bash
git clone https://github.com/lokesh-go/go-api-microservice.git
cd go-api-microservice
```

2. Install dependencies:

```bash
go mod download
```

3. Set up environment variables:

```bash
Project supports three enviroments.

# dev ( Default )
export ENV=dev

# test
export ENV=test

# prod
export ENV=prod
```

4. Run the application:

```bash
make run
```

5. Run the application using Docker:

```bash
# Build docker image
make build-docker

# Run docker image
make run-docker
```

6. Ping your application:

```bash
# Health check for your application
http://localhost/ping
```

## License

This project is licensed under the [MIT License](./LICENSE) © 2025 Lokesh Chandra.