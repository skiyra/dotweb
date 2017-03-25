package cache

import (
	"github.com/devfeel/dotweb/cache/redis"
	"github.com/devfeel/dotweb/cache/runtime"
)

type Cache interface {
	// Exist return true if value cached by given key
	Exists(key string) bool
	// Get returns value by given key
	Get(key string) interface{}
	// GetString returns value string format by given key
	GetString(key string) string
	// GetInt returns value int format by given key
	GetInt(key string) int
	// GetInt64 returns value int64 format by given key
	GetInt64(key string) int64
	// Set cache value by given key
	Set(key string, v interface{}, ttl int64) error
	// Incr increases int64-type value by given key as a counter
	// if key not exist, before increase set value with zero
	Incr(key string) (int64, error)
	// Decr decreases int64-type value by given key as a counter
	// if key not exist, before increase set value with zero
	Decr(key string) (int64, error)
	// Delete delete cache item by given key
	Delete(key string) error
	// ClearAll clear all cache items
	ClearAll() error
}

//new runtime cache
func NewRuntimeCache() Cache {
	return runtime.NewRuntimeCache()
}

//new redis cache
//must set serverIp like "redis://:password@10.0.1.11:6379/0"
func NewRedisCache(serverIp string) Cache {
	return redis.NewRedisCache(serverIp)
}
