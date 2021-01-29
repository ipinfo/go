package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	anycast, err := ipinfo.GetIPAnycast(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", anycast)
}
