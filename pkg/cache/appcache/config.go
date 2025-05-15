package appcache

var (
	defaultExpirationInSeconds int64 = 300
	cleanupIntervalInMinutes   int64 = 10
)

// Config ...
type Config struct {
	DefaultExpirationInSeconds int64
	CleanupIntervalInMinutes   int64
}
