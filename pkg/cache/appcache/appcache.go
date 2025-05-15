package appcache

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

// AppCache ...
type AppCache struct {
	Cache *goCache.Cache
}

// New ...
func New(config *Config) *AppCache {
	// Override config
	if config != nil && config.DefaultExpirationInSeconds != 0 {
		defaultExpirationInSeconds = config.DefaultExpirationInSeconds
	}
	if config != nil && config.CleanupIntervalInMinutes != 0 {
		cleanupIntervalInMinutes = config.CleanupIntervalInMinutes
	}

	// Initialise
	defaultExpiration := time.Duration(defaultExpirationInSeconds) * time.Second
	cleanupInterval := time.Duration(cleanupIntervalInMinutes) * time.Minute
	cache := goCache.New(defaultExpiration, cleanupInterval)

	// Returns
	return &AppCache{
		Cache: cache,
	}
}
