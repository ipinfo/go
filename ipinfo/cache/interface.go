package cache

import "errors"

var (
	ErrNotFound = errors.New("key not found")
)

type Interface interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}
