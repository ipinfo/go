package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/ipinfo/go/v2/ipinfo/cache"
)

type dummyCacheEngine struct {
	cache map[string]interface{}
}

func newDummyCacheEngine() *dummyCacheEngine {
	return &dummyCacheEngine{
		cache: make(map[string]interface{}),
	}
}

func (c *dummyCacheEngine) Get(key string) (interface{}, error) {
	log.Printf("[CACHE]: Trying to get value for key %q", key)
	if v, ok := c.cache[key]; ok {
		log.Printf("[CACHE]: Got value for key %q=%q", key, v)
		return v, nil
	}
	log.Printf("[CACHE]: Key %q not found", key)
	return nil, cache.ErrNotFound
}

func (c *dummyCacheEngine) Set(key string, value interface{}) error {
	log.Printf("[CACHE]: Setting value for key %q=%q", key, value)
	c.cache[key] = value
	return nil
}

var dummyCache = ipinfo.NewCache(newDummyCacheEngine())

func main() {
	ipinfo.SetCache(dummyCache)
	ip := net.ParseIP("8.8.8.8")

	for i := 0; i < 2; i++ {
		fmt.Println([]string{"Actual requests", "From cache"}[i])
		if v, err := ipinfo.GetIPInfo(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("IP: %v\n", v)
		}
	}
}
