package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	city, err := ipinfo.GetIPCity(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", city)
}
