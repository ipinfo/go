package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	region, err := ipinfo.GetIPRegion(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", region)
}
