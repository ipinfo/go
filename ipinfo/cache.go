package ipinfo

import (
	"sync"

	"github.com/ipinfo/go/ipinfo/cache"
)

// Cache represents the internal cache used by the IPinfo client.
type Cache struct {
	cache.Interface
	requestLocks sync.Map
}

// NewCache creates a new cache given a specific engine.
func NewCache(engine cache.Interface) *Cache {
	return &Cache{Interface: engine}
}
