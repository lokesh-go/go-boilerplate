package main

import (
	"log"

	"github.com/lokesh-go/go-boilerplate/internal/bootstrap"
)

// API Server start point
func main() {
	// Initialize the server
	err := bootstrap.Initialize()
	if err != nil {
		log.Fatalf("api server initialization failed: %v", err)
	}
}
