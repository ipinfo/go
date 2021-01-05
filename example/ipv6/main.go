package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	core, err := ipinfo.GetIPInfo(net.ParseIP("2a03:2880:f10a:83:face:b00c:0:25de"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", core)
}
