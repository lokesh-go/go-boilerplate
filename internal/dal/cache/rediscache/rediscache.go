package rediscache

import (
	"github.com/lokesh-go/go-boilerplate/internal/config"
	redisPkg "github.com/lokesh-go/go-boilerplate/pkg/cache/redis"
)

// RedisCache conn
type redisCache struct {
	conn *redisPkg.RedisCache
}

// New redis cache
func New(config config.Methods) (*redisCache, error) {
	// Forms redis config
	rc := &redisPkg.Config{}

	// Get the redis client
	conn, err := redisPkg.New(rc)
	if err != nil {
		return nil, err
	}

	// Returns
	return &redisCache{
		conn: conn,
	}, nil
}
