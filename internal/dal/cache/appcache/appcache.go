package appcache

import (
	"github.com/lokesh-go/go-boilerplate/internal/config"
	appCachePkg "github.com/lokesh-go/go-boilerplate/pkg/cache/appcache"
)

// AppCache conn
type AppCache struct {
	conn *appCachePkg.AppCache
}

// New app cache
func New(config config.Methods) *AppCache {
	// Forms appCache config
	ac := &appCachePkg.Config{
		DefaultExpirationInSeconds: config.Get().DAL.Cache.AppCache.DefaultExpirationInSeconds,
		CleanupIntervalInMinutes:   config.Get().DAL.Cache.AppCache.CleanupIntervalInMinutes,
	}

	// Get the appCache connection
	return &AppCache{
		conn: appCachePkg.New(ac),
	}
}
