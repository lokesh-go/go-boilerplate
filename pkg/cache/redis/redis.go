package redis

// RedisCache ...
type RedisCache struct{}

// New ...
func New(config *Config) (*RedisCache, error) {
	// Returns
	return &RedisCache{}, nil
}
