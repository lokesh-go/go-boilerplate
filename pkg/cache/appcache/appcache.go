package appcache

import (
	"time"

	appCache "github.com/patrickmn/go-cache"

	"github.com/lokesh-go/go-boilerplate/internal/config"
)

// AppCache ...
type AppCache struct {
	Cache *appCache.Cache
}

// New ...
// TODO :: Need decoupling here
func New(config config.Methods) (c *AppCache) {
	// Initialise
	defaultExpiration := time.Duration(config.Get().DAL.Cache.AppCache.DefaultExpirationInSeconds) * time.Second
	cleanupInterval := time.Duration(config.Get().DAL.Cache.AppCache.CleanupIntervalInMinutes) * time.Minute
	cache := appCache.New(defaultExpiration, cleanupInterval)

	// Returns
	return &AppCache{
		Cache: cache,
	}
}
