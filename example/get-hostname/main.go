package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	hostname, err := ipinfo.GetIPHostname(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", hostname)
}
