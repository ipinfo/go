package cache // import "github.com/ipinfo/go/ipinfo/cache"

import "errors"

// Errors
var (
	ErrNotFound = errors.New("key not found")
)

type Interface interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
