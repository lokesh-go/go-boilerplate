package server

import (
	"github.com/lokesh-go/go-api-microservice/internal/dependencies"
	grpcServer "github.com/lokesh-go/go-api-microservice/internal/server/grpc"
	httpServer "github.com/lokesh-go/go-api-microservice/internal/server/http"
)

// Start
func Start(d *dependencies.Deps) (err error) {
	// Starts http server
	err = httpServer.New(d).Start()
	if err != nil {
		return err
	}

	// Starts grpc server
	grpcServer.Start()

	// Returns
	return nil
}
