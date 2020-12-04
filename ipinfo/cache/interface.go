package cache

import "errors"

var (
	// ErrNotFound means that the key was not found.
	ErrNotFound = errors.New("key not found")
)

// Interface is the cache interface that all cache implementations must adhere
// to at the minimum to be usable in the IPinfo client.
type Interface interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
