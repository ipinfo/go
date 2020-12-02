package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/ipinfo/go/ipinfo"
	"github.com/ipinfo/go/ipinfo/cache"
)

func main() {
	ipinfo.SetCache(
		ipinfo.NewCache(
			cache.NewInMemory().WithExpiration(5 * time.Minute),
		),
	)
	ip := net.ParseIP("8.8.8.8")

	for i := 0; i < 2; i++ {
		fmt.Println([]string{"Actual requests", "From cache"}[i])
		if v, err := ipinfo.GetIpInfo(ip); err != nil {
			log.Println(err)
		} else {
			fmt.Printf("IP: %v\n", v)
		}
	}
}
