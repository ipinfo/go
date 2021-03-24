package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	bogon, err := ipinfo.GetIPBogon(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", bogon)
}
