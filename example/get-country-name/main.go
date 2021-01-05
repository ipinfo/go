package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	countryName, err := ipinfo.GetIPCountryName(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", countryName)
}
