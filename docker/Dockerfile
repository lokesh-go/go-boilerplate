# -------- Stage 1: Build --------
FROM golang:1.23 as builder

# Set environment variables to build a static binary
ENV CGO_ENABLED=0 GOOS=linux GOARCH=arm64

# Set the working directory inside the container
WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary
RUN go build -o /app/go-api-microservice ./cmd/go-api-microservice/main.go


# -------- Stage 2: Run --------
FROM gcr.io/distroless/static:nonroot

# Accept environment as a build argument
ARG ENV=dev
ENV ENV=${ENV}

# Set working directory to /app (same as project root)
WORKDIR /app

# Copy only the binary and the required config files
COPY --from=builder /app/go-api-microservice .
COPY --from=builder /app/internal/config/env/${ENV} /app/internal/config/env/${ENV}

# Expose the app port (change if needed)
EXPOSE 80 8080

# Use a non-root user (automatically set in distroless)
ENTRYPOINT ["/app/go-api-microservice"]