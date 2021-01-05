package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	core, err := ipinfo.GetIPInfo(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", core)
}
