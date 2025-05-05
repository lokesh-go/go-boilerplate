package cache

import (
	"github.com/lokesh-go/go-boilerplate/internal/config"

	appCachePkg "github.com/lokesh-go/go-boilerplate/pkg/cache/appcache"
	redisCachePkg "github.com/lokesh-go/go-boilerplate/pkg/cache/redis"
)

type cacheMethod string

const (
	appCache     cacheMethod = "appCache"
	redisCache   cacheMethod = "redisCache"
	defaultCache cacheMethod = appCache
)

// Methods ...
type Methods interface {
	Get(key string, val interface{}) (err error)
}

type cache struct {
	appCache   *appCachePkg.AppCache
	redisCache *redisCachePkg.RedisCache
}

// New ...
func New(config config.Methods) (Methods, error) {
	var ac *appCachePkg.AppCache
	var rc *redisCachePkg.RedisCache

	switch defaultCache {
	case redisCache:
		{
			// Cache disabled
			if !config.Get().DAL.Cache.Redis.Enabled {
				return nil, nil
			}

			// Initialise redis cache
			conn, err := redisCachePkg.New(config)
			if err != nil {
				return nil, err
			}
			rc = conn
		}
	case appCache:
		{
			// Cache disabled
			if !config.Get().DAL.Cache.AppCache.Enabled {
				return nil, nil
			}

			// Initialise app cache
			ac = appCachePkg.New(config)
		}
	}

	// Returns
	return &cache{
		appCache:   ac,
		redisCache: rc,
	}, nil
}
