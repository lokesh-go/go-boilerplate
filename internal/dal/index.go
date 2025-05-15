package dal

import "github.com/lokesh-go/go-boilerplate/internal/dal/cache"

// Data Access Layer
type Dal struct {
	Cache cache.Cache
}

// Initialize
func Initialize() (err error) {
	// Connects cache

	// Connects DB

	// Returns
	return nil
}
