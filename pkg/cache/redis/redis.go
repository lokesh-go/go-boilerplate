package redis

import "github.com/lokesh-go/go-boilerplate/internal/config"

// RedisCache ...
type RedisCache struct{}

// New ...
// TODO :: Need decoupling here
func New(config config.Methods) (r *RedisCache, err error) {
	// Returns
	return &RedisCache{}, nil
}
