package cache

import "github.com/coocood/freecache"

var cacheSize = 100 * 1024 * 1024

var cache = freecache.NewCache(cacheSize)

// SetCache
func SetCache(key, value []byte, expireSeconds int) error {
	return cache.Set(key, value, expireSeconds)
}

// GetCache
func GetCache(key []byte) ([]byte, error) {
	return cache.Get(key)
}

// DelCache
func DelCache(key []byte) bool {
	return cache.Del(key)
}
