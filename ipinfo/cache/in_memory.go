package cache // import "github.com/ipinfo/go-ipinfo/ipinfo/cache"

import (
	"time"

	"github.com/patrickmn/go-cache"
)

const defaultExpiration = 24 * time.Hour

type InMemory struct {
	cache      *cache.Cache
	expiration time.Duration
}

func NewInMemory() *InMemory {
	return &InMemory{
		cache:      cache.New(-1, defaultExpiration),
		expiration: defaultExpiration,
	}
}

func (c *InMemory) WithExpiration(d time.Duration) *InMemory {
	c.expiration = d
	return c
}

func (c *InMemory) Get(key string) (interface{}, error) {
	v, found := c.cache.Get(key)
	if !found {
		return nil, ErrNotFound
	}
	return v, nil
}

func (c *InMemory) Set(key string, value interface{}) error {
	c.cache.Set(key, value, c.expiration)
	return nil
}

// Check if InMemory implements cache.Interface
var _ Interface = (*InMemory)(nil)
