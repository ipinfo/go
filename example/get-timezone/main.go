package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/v2/ipinfo"
)

func main() {
	timezone, err := ipinfo.GetIPTimezone(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", timezone)
}
