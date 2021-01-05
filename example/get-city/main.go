package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	city, err := ipinfo.GetIPCity(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", city)
}
