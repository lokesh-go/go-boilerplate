package mongo

import (
	"github.com/lokesh-go/go-boilerplate/internal/config"
	mongoPkg "github.com/lokesh-go/go-boilerplate/pkg/db/mongo"
)

// MongoDB connection
type mongoDB struct {
	conn *mongoPkg.MongoDB
}

// New mongoDB connection
func New(config config.Methods) (*mongoDB, error) {
	// Forms mongoDB config
	mongoDBConfig := &mongoPkg.Config{}

	// Creats mongoDB connection
	conn, err := mongoPkg.New(mongoDBConfig)
	if err != nil {
		return nil, err
	}

	// Returns
	return &mongoDB{
		conn: conn,
	}, nil
}
