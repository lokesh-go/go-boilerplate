<p align="center"><img src="https://user-images.githubusercontent.com/1092882/60883564-20142380-a268-11e9-988a-d98fb639adc6.png" alt="webgo gopher" width="256px"/></p>

# go-boilerplate

Golang microservice boilerplate supporting both REST and gRPC APIs.

## Directory Structure

```bash
.
├── api/                # API definitions and documentation
│   ├── proto/         # gRPC Protocol buffer definitions
│   │   └── v1/       # API versioning
│   │       ├── user/
│   │       └── auth/
│   │
│   ├── rest/         # REST API definitions
│   │   ├── routes/   # Route definitions
│   │   │   ├── v1/  # Version-specific routes
│   │   │   └── routes.go
│   │   │
│   │   └── swagger/ # Swagger/OpenAPI documentation
│   │       └── v1/
│   │
│   └── docs/         # API documentation
│
├── build/            # Build artifacts and packaging
│
├── cmd/              # Main applications
│   ├── api-server/  # Combined REST + gRPC server
│   │   └── main.go
│   │
│   ├── worker/      # Background job processor
│   │   └── main.go
│   │
│   └── cli/         # Command-line tools
│       └── main.go
│
├── docker/          # Containerization
│
├── docs/            # Design and user documents
│
├── internal/        # Private application code
│   ├── config/     # Configuration management
│   │
│   ├── dal/        # Data Access Layer
│   │   ├── postgres/
│   │   └── redis/
│   │
│   ├── models/     # Shared domain models
│   │   ├── user.go
│   │   └── auth.go
│   │
│   ├── server/     # Server implementations
│   │   ├── grpc/   # gRPC server
│   │   │   ├── handlers/
│   │   │   └── middleware/
│   │   │
│   │   ├── http/   # REST server
│   │   │   ├── handlers/
│   │   │   ├── middleware/
│   │   │   └── server.go
│   │   │
│   │   └── server.go
│   │
│   └── service/    # Business logic layer
│       ├── user/
│       └── auth/
│
├── scripts/        # Build and utility scripts
│
├── test/           # Additional test data and test suites
│
├── go.mod          # Module dependencies
│
└── README.md       # Project documentation
```

## Directory Significance

### `api/`
- Contains API definitions and documentation
- Separates API contract from implementation
- Includes both REST and gRPC API definitions
- Versioned API specifications

### `cmd/`
- Contains main applications
- Each subdirectory is a main package
- Directory name matches executable name
- Minimal code, mainly calling packages from `internal/`
- Examples:
  - `api-server`: Combined REST + gRPC server
  - `worker`: Background job processor
  - `migrator`: Database migration tool
  - `cli`: Command-line tools

### `internal/`
- Private application code
- Cannot be imported by other projects
- Contains core business logic
- Organized by domain and responsibility

### Supporting Directories
- `build/`: Build artifacts and packaging
- `docker/`: Containerization
- `docs/`: Design and user documents
- `scripts/`: Build and utility scripts
- `test/`: Additional test data and test suites

## Installation

Install individual applications:
```bash
go install github.com/lokesh-go/go-boilerplate/cmd/api-server@latest
go install github.com/lokesh-go/go-boilerplate/cmd/worker@latest
```

## Development

1. Clone the repository
2. Install dependencies
3. Run the development server:
```bash
go run cmd/api-server/main.go
```

## Testing

Run tests:
```bash
go test ./...
```