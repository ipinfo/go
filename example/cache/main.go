package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/ipinfo"
	"github.com/ipinfo/go/ipinfo/cache"
)

type DummyCacheEngine struct {
	cache map[string]interface{}
}

func NewDummyCacheEngine() *DummyCacheEngine {
	return &DummyCacheEngine{
		cache: make(map[string]interface{}),
	}
}

func (c *DummyCacheEngine) Get(key string) (interface{}, error) {
	log.Printf("[CACHE]: Trying to get value for key %q", key)
	if v, ok := c.cache[key]; ok {
		log.Printf("[CACHE]: Got value for key %q=%q", key, v)
		return v, nil
	}
	log.Printf("[CACHE]: Key %q not found", key)
	return nil, cache.ErrNotFound
}

func (c *DummyCacheEngine) Set(key string, value interface{}) error {
	log.Printf("[CACHE]: Setting value for key %q=%q", key, value)
	c.cache[key] = value
	return nil
}

var DummyCache = ipinfo.NewCache(NewDummyCacheEngine())

func main() {
	ipinfo.SetCache(DummyCache)
	ip := net.ParseIP("8.8.8.8")

	for i := 0; i < 2; i++ {
		if v, err := ipinfo.GetIP(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("IP:", v)
		}

		if v, err := ipinfo.GetHostname(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Hostname:", v)
		}

		if v, err := ipinfo.GetOrganization(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Organization:", v)
		}

		if v, err := ipinfo.GetCity(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("City:", v)
		}

		if v, err := ipinfo.GetRegion(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Region:", v)
		}

		if v, err := ipinfo.GetCountry(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Country:", v)
		}

		if v, err := ipinfo.GetLocation(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Location:", v)
		}

		if v, err := ipinfo.GetPhone(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Phone:", v)
		}

		if v, err := ipinfo.GetPostal(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Println("Postal:", v)
		}
	}
}
