package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	loc, err := ipinfo.GetIPLocation(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", loc)
}
