package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	var c *ipinfo.LiteClient

	if token := os.Getenv("TOKEN"); token != "" {
		c = ipinfo.NewLiteClient(nil, nil, token)
	} else {
		c = ipinfo.DefaultLiteClient
	}

	/* default to user IP */
	if len(os.Args) == 1 {
		info, err := c.GetIPInfo(nil)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v\n", info)
		return
	}

	for _, s := range os.Args[1:] {
		ip := net.ParseIP(s)
		if ip == nil {
			continue
		}
		info, err := c.GetIPInfo(ip)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%v\n", info)
	}
}
