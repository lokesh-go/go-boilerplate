package server

import (
	grpcServer "github.com/lokesh-go/go-boilerplate/internal/server/grpc"
	httpServer "github.com/lokesh-go/go-boilerplate/internal/server/http"
)

// Start
func Start() (err error) {
	// Starts http server
	httpServer.Start()

	// Starts grpc server
	grpcServer.Start()

	// Returns
	return nil
}
