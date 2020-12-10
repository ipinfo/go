package main

import (
	"fmt"
	"log"

	"github.com/ipinfo/go/ipinfo"
)

func main() {
	org, err := ipinfo.GetIPOrg(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", org)
}
