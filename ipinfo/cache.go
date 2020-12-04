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

// EvaluatorFunc is a function called when a cache miss occurs to fill the
// cache with the missing value for a key.
type EvaluatorFunc func() (interface{}, error)

// GetOrRequest attempts to first retrieve the value from the cache, and if it
// isn't found, calls `evaluator` to get it and backfill the cache.
func (c *Cache) GetOrRequest(
	key string,
	evaluator EvaluatorFunc,
) (interface{}, error) {
	value, _ := c.requestLocks.LoadOrStore(key, &sync.Mutex{})
	mutex := value.(*sync.Mutex)
	mutex.Lock()
	defer func() {
		c.requestLocks.Delete(key)
		mutex.Unlock()
	}()
	value, err := c.Get(key)
	if err == nil {
		return value, nil
	}
	if err == cache.ErrNotFound {
		value, err := evaluator()
		if err != nil {
			return nil, err
		}
		err = c.Set(key, value)
		return value, err
	}
	return nil, err
}
