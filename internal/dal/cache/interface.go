package cache

// Cache ...
type Cache interface {
	Get(key string, val interface{}) (err error)
}
