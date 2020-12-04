package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	var c *ipinfo.Client

	if token := os.Getenv("TOKEN"); token != "" {
		c = ipinfo.NewClient(nil, nil, token)
	} else {
		c = ipinfo.DefaultClient
	}

	/* default to user IP */
	if len(os.Args) == 1 {
		core, err := c.GetIpInfo(nil)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v\n", core)
		return
	}

	for _, s := range os.Args[1:] {
		ip := net.ParseIP(s)
		if ip == nil {
			continue
		}
		core, err := c.GetIpInfo(ip)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v\n", core)
	}
}
