package ipinfo

import (
	"sync"

	"github.com/ipinfo/go/ipinfo/cache"
)

type Cache struct {
	cache.Interface
	requestLocks sync.Map
}

func NewCache(engine cache.Interface) *Cache {
	return &Cache{Interface: engine}
}

type EvaluatorFunc func() (interface{}, error)

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
