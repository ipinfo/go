package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	country, err := ipinfo.GetIPCountry(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", country)
}
