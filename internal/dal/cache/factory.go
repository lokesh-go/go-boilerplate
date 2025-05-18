package cache

import (
	"github.com/lokesh-go/go-api-service/internal/config"
	appCacheDAL "github.com/lokesh-go/go-api-service/internal/dal/cache/appcache"
	redisCacheDAL "github.com/lokesh-go/go-api-service/internal/dal/cache/rediscache"
)

type cacheType string

const (
	appCache     cacheType = "appCache"
	redisCache   cacheType = "redisCache"
	cacheEnabled cacheType = appCache
)

// New Cache
func New(config config.Methods) (Cache, error) {
	// Switch to the cache
	switch cacheEnabled {
	case redisCache:
		{
			// Create redis connection
			return redisCacheDAL.New(config)
		}
	default:
		{
			// Create appCache connection
			return appCacheDAL.New(config), nil
		}
	}
}
